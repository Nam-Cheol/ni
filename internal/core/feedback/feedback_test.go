package feedback

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestAddAndList(t *testing.T) {
	dir := t.TempDir()
	input := filepath.Join(dir, "codex.json")
	if err := os.WriteFile(input, []byte(`{
  "source_target": "codex",
  "related_lock_hash": "sha256:test-lock",
  "related_capabilities": ["CAP-021"],
  "observed_blockers": ["blocker"],
  "validation_gaps": ["gap"],
  "recurring_failures": ["failure"],
  "suggested_contract_changes": ["change"],
  "raw_summary": "summary"
}`), 0o644); err != nil {
		t.Fatalf("writing input: %v", err)
	}

	entry, err := Add(AddOptions{
		Dir:  dir,
		File: input,
		Now:  time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC),
	})
	if err != nil {
		t.Fatalf("add feedback: %v", err)
	}
	if entry.Schema != Schema {
		t.Fatalf("expected schema %q, got %q", Schema, entry.Schema)
	}

	entries, err := List(dir)
	if err != nil {
		t.Fatalf("list feedback: %v", err)
	}
	if len(entries) != 1 {
		t.Fatalf("expected one entry, got %#v", entries)
	}
	if entries[0].SourceTarget != "codex" {
		t.Fatalf("expected codex feedback, got %#v", entries[0])
	}
}

func TestAddRejectsUnknownFields(t *testing.T) {
	dir := t.TempDir()
	input := filepath.Join(dir, "log.json")
	if err := os.WriteFile(input, []byte(`{
  "source_target": "codex",
  "raw_summary": "summary",
  "raw_log": "arbitrary logs stay outside feedback ingest"
}`), 0o644); err != nil {
		t.Fatalf("writing input: %v", err)
	}

	if _, err := Add(AddOptions{Dir: dir, File: input}); err == nil {
		t.Fatalf("expected unknown field rejection")
	}
}
