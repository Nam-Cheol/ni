package main

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestHelp(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"--help"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "ni is a project intent compiler") {
		t.Fatalf("help output did not describe ni: %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "ni version") {
		t.Fatalf("help output did not mention version command: %q", stdout.String())
	}
}

func TestVersion(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"version"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if got := strings.TrimSpace(stdout.String()); got != "0.0.0-dev" {
		t.Fatalf("expected version 0.0.0-dev, got %q", got)
	}
}

func TestInit(t *testing.T) {
	dir := t.TempDir()
	var stdout bytes.Buffer

	code := run([]string{"init", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	required := []string{
		"docs/plan/00_project_brief.md",
		".ni/project.json",
		".ni/contract.json",
		".ni/readiness.rules.json",
	}
	for _, path := range required {
		if _, err := os.Stat(filepath.Join(dir, path)); err != nil {
			t.Fatalf("expected %s to exist: %v", path, err)
		}
	}
	if !strings.Contains(stdout.String(), "initialized ni planning workspace") {
		t.Fatalf("expected init summary, got %q", stdout.String())
	}
}

func TestStatus(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "BLOCKED") {
		t.Fatalf("expected blocked status for template project, got %q", stdout.String())
	}
}

func TestStatusJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"status": "BLOCKED"`) {
		t.Fatalf("expected JSON blocked status, got %q", stdout.String())
	}
}

func TestEndBlocked(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	var stderr bytes.Buffer
	code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "readiness is BLOCKED") {
		t.Fatalf("expected blocked error, got %q", stderr.String())
	}
}

func TestEndCreatesLockfile(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"end", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "locked plan") {
		t.Fatalf("expected lock summary, got %q", stdout.String())
	}
	if _, err := os.Stat(filepath.Join(dir, ".ni", "plan.lock.json")); err != nil {
		t.Fatalf("expected lockfile: %v", err)
	}
}

func TestUnknownCommand(t *testing.T) {
	var stderr bytes.Buffer
	code := run([]string{"run"}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "unknown command: run") {
		t.Fatalf("expected unknown command error, got %q", stderr.String())
	}
}

func writeReadyContractForCLI(t *testing.T, dir string) {
	t.Helper()

	contract := map[string]any{
		"schema": "ni.contract.v0",
		"project": map[string]any{
			"id":      "cli-fixture",
			"name":    "CLI Fixture",
			"purpose": "Exercise ni end.",
			"status":  "draft",
		},
		"non_goals": []any{
			map[string]any{"id": "NG-001", "title": "Do not execute work."},
		},
		"capabilities": []any{
			map[string]any{
				"id":           "CAP-001",
				"title":        "Capability",
				"status":       "accepted",
				"requirements": []string{"REQ-001"},
				"evaluations":  []string{"EVAL-001"},
				"risks":        []string{},
				"artifacts":    []string{"ART-001"},
			},
		},
		"requirements":   []any{map[string]any{"id": "REQ-001", "title": "Requirement", "status": "accepted"}},
		"decisions":      []any{map[string]any{"id": "DEC-001", "title": "Decision", "status": "accepted"}},
		"risks":          []any{},
		"evaluations":    []any{map[string]any{"id": "EVAL-001", "title": "Evaluation", "method": "fixture"}},
		"artifacts":      []any{map[string]any{"id": "ART-001", "path": "docs/plan/", "kind": "planning_docs"}},
		"open_questions": []any{},
	}
	data, err := json.MarshalIndent(contract, "", "  ")
	if err != nil {
		t.Fatalf("marshaling ready contract: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), append(data, '\n'), 0o644); err != nil {
		t.Fatalf("writing ready contract: %v", err)
	}
}
