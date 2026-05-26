package main

import (
	"bytes"
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

func TestUnknownCommand(t *testing.T) {
	var stderr bytes.Buffer
	code := run([]string{"status"}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "unknown command: status") {
		t.Fatalf("expected unknown command error, got %q", stderr.String())
	}
}
