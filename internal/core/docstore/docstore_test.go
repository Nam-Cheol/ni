package docstore

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestInitCreatesRequiredFiles(t *testing.T) {
	dir := t.TempDir()

	result, err := Init(dir)
	if err != nil {
		t.Fatalf("Init returned error: %v", err)
	}

	paths := RequiredPaths()
	if len(result.Created) != len(paths) {
		t.Fatalf("expected %d created files, got %d", len(paths), len(result.Created))
	}
	for _, path := range paths {
		if _, err := os.Stat(filepath.Join(dir, path)); err != nil {
			t.Fatalf("expected %s to exist: %v", path, err)
		}
	}
}

func TestInitWithProfileWritesSelectedReadinessProfile(t *testing.T) {
	dir := t.TempDir()

	if _, err := InitWithProfile(dir, "production"); err != nil {
		t.Fatalf("InitWithProfile returned error: %v", err)
	}

	data, err := os.ReadFile(filepath.Join(dir, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("reading contract: %v", err)
	}
	var contract struct {
		ReadinessProfile string `json:"readiness_profile"`
	}
	if err := json.Unmarshal(data, &contract); err != nil {
		t.Fatalf("unmarshaling contract: %v", err)
	}
	if contract.ReadinessProfile != "production" {
		t.Fatalf("expected production profile, got %q", contract.ReadinessProfile)
	}

	if _, err := os.Stat(filepath.Join(dir, ".ni", "readiness.profiles.json")); err != nil {
		t.Fatalf("expected readiness profiles file: %v", err)
	}
}

func TestInitPreservesExistingFiles(t *testing.T) {
	dir := t.TempDir()
	existingPath := filepath.Join(dir, "docs", "plan", "00_project_brief.md")
	if err := os.MkdirAll(filepath.Dir(existingPath), 0o755); err != nil {
		t.Fatalf("creating existing file dir: %v", err)
	}
	if err := os.WriteFile(existingPath, []byte("user content\n"), 0o644); err != nil {
		t.Fatalf("creating existing file: %v", err)
	}

	result, err := Init(dir)
	if err != nil {
		t.Fatalf("Init returned error: %v", err)
	}

	content, err := os.ReadFile(existingPath)
	if err != nil {
		t.Fatalf("reading existing file: %v", err)
	}
	if string(content) != "user content\n" {
		t.Fatalf("existing file was overwritten: %q", string(content))
	}
	if len(result.Existing) != 1 || result.Existing[0] != "docs/plan/00_project_brief.md" {
		t.Fatalf("expected existing file to be reported, got %#v", result.Existing)
	}
}
