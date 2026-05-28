package lock

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"ni/internal/core/docstore"
)

func TestCreateAtWritesLockfile(t *testing.T) {
	dir := readyProject(t)
	lockedAt := time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC)

	lockfile, err := CreateAt(dir, lockedAt)
	if err != nil {
		t.Fatalf("CreateAt returned error: %v", err)
	}

	if lockfile.Schema != Schema {
		t.Fatalf("expected schema %q, got %q", Schema, lockfile.Schema)
	}
	if lockfile.LockedAt != "2026-05-26T01:02:03Z" {
		t.Fatalf("unexpected lock timestamp %q", lockfile.LockedAt)
	}
	if lockfile.Readiness.Status != "READY" {
		t.Fatalf("expected READY, got %q", lockfile.Readiness.Status)
	}
	if len(lockfile.SourceOfTruth) == 0 || lockfile.SourceOfTruth[0] != ".ni/plan.lock.json" {
		t.Fatalf("unexpected source of truth: %#v", lockfile.SourceOfTruth)
	}
	if len(lockfile.SourceOfTruth) != 5 || lockfile.SourceOfTruth[3] != ".ni/session.json" {
		t.Fatalf("expected session state below docs/plan in source of truth: %#v", lockfile.SourceOfTruth)
	}
	if len(lockfile.Files) != 13 {
		t.Fatalf("expected 13 hashed files, got %d", len(lockfile.Files))
	}
	for _, file := range lockfile.Files {
		if file.Path == ".ni/contract.json" && len(file.SHA256) == 64 {
			return
		}
	}
	t.Fatalf("expected contract hash in %#v", lockfile.Files)
}

func TestCreateAtPersistsJSON(t *testing.T) {
	dir := readyProject(t)

	lockfile, err := CreateAt(dir, time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC))
	if err != nil {
		t.Fatalf("CreateAt returned error: %v", err)
	}

	data, err := os.ReadFile(lockfile.Path)
	if err != nil {
		t.Fatalf("reading lockfile: %v", err)
	}
	var persisted Lockfile
	if err := json.Unmarshal(data, &persisted); err != nil {
		t.Fatalf("parsing lockfile JSON: %v", err)
	}
	if persisted.Schema != Schema {
		t.Fatalf("expected persisted schema %q, got %q", Schema, persisted.Schema)
	}
}

func TestCreateAtRefusesBlockedReadiness(t *testing.T) {
	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}

	_, err := CreateAt(dir, time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC))
	if err == nil {
		t.Fatal("expected blocked readiness error")
	}
	if _, statErr := os.Stat(filepath.Join(dir, ".ni", "plan.lock.json")); !os.IsNotExist(statErr) {
		t.Fatalf("lockfile should not exist, stat err: %v", statErr)
	}
}

func readyProject(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}
	data, err := os.ReadFile(filepath.Join("..", "readiness", "testdata", "ready.json"))
	if err != nil {
		t.Fatalf("reading ready fixture: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), data, 0o644); err != nil {
		t.Fatalf("writing ready contract: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "10_open_questions.md"), []byte("# Open questions\n\nNo open questions are listed in this fixture.\n"), 0o644); err != nil {
		t.Fatalf("writing ready open question doc: %v", err)
	}
	return dir
}
