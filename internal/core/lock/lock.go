package lock

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"

	"ni/internal/core/readiness"
)

const Schema = "ni.lock.v0"

type Lockfile struct {
	Schema        string           `json:"schema"`
	LockedAt      string           `json:"locked_at"`
	SourceOfTruth []string         `json:"source_of_truth"`
	Readiness     ReadinessSummary `json:"readiness"`
	Files         []FileHash       `json:"files"`
	Path          string           `json:"-"`
}

type ReadinessSummary struct {
	Status string `json:"status"`
}

type FileHash struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

type Verification struct {
	Lockfile   Lockfile
	Mismatches []Mismatch
}

type Mismatch struct {
	Path     string
	WantHash string
	GotHash  string
}

func Create(dir string) (Lockfile, error) {
	return CreateAt(dir, time.Now().UTC())
}

func CreateAt(dir string, lockedAt time.Time) (Lockfile, error) {
	root := filepath.Clean(dir)
	status := readiness.Evaluate(root)
	if status.Status == readiness.StatusBlocked {
		return Lockfile{}, fmt.Errorf("readiness is BLOCKED; refusing to lock")
	}

	files, err := hashFiles(root, lockPaths(root))
	if err != nil {
		return Lockfile{}, err
	}

	lockfile := Lockfile{
		Schema:   Schema,
		LockedAt: lockedAt.UTC().Format(time.RFC3339),
		SourceOfTruth: []string{
			".ni/plan.lock.json",
			".ni/contract.json",
			"docs/plan/**",
			"chat transcript",
			"model guess",
		},
		Readiness: ReadinessSummary{Status: string(status.Status)},
		Files:     files,
	}

	data, err := json.MarshalIndent(lockfile, "", "  ")
	if err != nil {
		return Lockfile{}, err
	}
	lockPath := filepath.Join(root, ".ni", "plan.lock.json")
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o755); err != nil {
		return Lockfile{}, err
	}
	if err := os.WriteFile(lockPath, append(data, '\n'), 0o644); err != nil {
		return Lockfile{}, err
	}
	lockfile.Path = lockPath
	return lockfile, nil
}

func LoadFile(path string) (Lockfile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Lockfile{}, err
	}
	var lockfile Lockfile
	if err := json.Unmarshal(data, &lockfile); err != nil {
		return Lockfile{}, fmt.Errorf("malformed lockfile JSON: %w", err)
	}
	if lockfile.Schema != Schema {
		return Lockfile{}, fmt.Errorf("unsupported lockfile schema %q", lockfile.Schema)
	}
	lockfile.Path = path
	return lockfile, nil
}

func Verify(dir string) (Verification, error) {
	root := filepath.Clean(dir)
	lockPath := filepath.Join(root, ".ni", "plan.lock.json")
	lockfile, err := LoadFile(lockPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Verification{}, fmt.Errorf("missing lockfile %s", lockPath)
		}
		return Verification{}, err
	}

	var mismatches []Mismatch
	for _, file := range lockfile.Files {
		got, err := fileSHA256(filepath.Join(root, file.Path))
		if err != nil {
			mismatches = append(mismatches, Mismatch{Path: file.Path, WantHash: file.SHA256, GotHash: "missing"})
			continue
		}
		if got != file.SHA256 {
			mismatches = append(mismatches, Mismatch{Path: file.Path, WantHash: file.SHA256, GotHash: got})
		}
	}

	return Verification{Lockfile: lockfile, Mismatches: mismatches}, nil
}

func lockPaths(root string) []string {
	paths := []string{".ni/contract.json"}
	paths = append(paths, readiness.RequiredDocs(root)...)
	sort.Strings(paths)
	return paths
}

func hashFiles(root string, paths []string) ([]FileHash, error) {
	files := make([]FileHash, 0, len(paths))
	for _, path := range paths {
		sum, err := fileSHA256(filepath.Join(root, path))
		if err != nil {
			return nil, fmt.Errorf("hash %s: %w", path, err)
		}
		files = append(files, FileHash{
			Path:   path,
			SHA256: sum,
		})
	}
	return files, nil
}

func fileSHA256(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}
