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
		".ni/readiness.profiles.json",
	}
	for _, path := range required {
		if _, err := os.Stat(filepath.Join(dir, path)); err != nil {
			t.Fatalf("expected %s to exist: %v", path, err)
		}
	}
	if !strings.Contains(stdout.String(), "initialized ni planning workspace") {
		t.Fatalf("expected init summary, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "readiness profile: prototype") {
		t.Fatalf("expected prototype profile summary, got %q", stdout.String())
	}
}

func TestInitWithProfile(t *testing.T) {
	dir := t.TempDir()
	var stdout bytes.Buffer

	code := run([]string{"init", "--dir", dir, "--profile", "concept"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	data, err := os.ReadFile(filepath.Join(dir, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("reading contract: %v", err)
	}
	if !strings.Contains(string(data), `"readiness_profile": "concept"`) {
		t.Fatalf("expected concept readiness profile, got %q", string(data))
	}
	if !strings.Contains(stdout.String(), "readiness profile: concept") {
		t.Fatalf("expected concept profile summary, got %q", stdout.String())
	}
}

func TestInitWithProductTypeAndSurface(t *testing.T) {
	dir := t.TempDir()
	var stdout bytes.Buffer

	code := run([]string{"init", "--dir", dir, "--product-type", "conversation_product", "--surface", "conversation", "--interaction-mode", "human_to_human"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	data, err := os.ReadFile(filepath.Join(dir, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("reading contract: %v", err)
	}
	if !strings.Contains(string(data), `"product_type": "conversation_product"`) {
		t.Fatalf("expected conversation product type, got %q", string(data))
	}
	if !strings.Contains(string(data), `"delivery_surfaces": ["conversation"]`) {
		t.Fatalf("expected conversation surface, got %q", string(data))
	}
	if !strings.Contains(stdout.String(), "product type: conversation_product") {
		t.Fatalf("expected product type summary, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "delivery surfaces: conversation") {
		t.Fatalf("expected surface summary, got %q", stdout.String())
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
	if !strings.Contains(stdout.String(), "profile: prototype") {
		t.Fatalf("expected active profile in status output, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "product type: software") {
		t.Fatalf("expected product type in status output, got %q", stdout.String())
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
	if !strings.Contains(stdout.String(), `"profile": "prototype"`) {
		t.Fatalf("expected JSON profile, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), `"product_type": "software"`) {
		t.Fatalf("expected JSON product type, got %q", stdout.String())
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

func TestRunWritesPrompt(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, ".ni", "generated", "goal.prompt.txt")
	var stdout bytes.Buffer
	code := run([]string{"run", "--dir", dir, "--out", out, "--max-chars", "4000"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "compiled prompt") {
		t.Fatalf("expected compile summary, got %q", stdout.String())
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("expected prompt file: %v", err)
	}
	if !strings.Contains(string(data), "Source of truth") {
		t.Fatalf("expected compiled prompt content, got %q", string(data))
	}
}

func TestRunWithTarget(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	for _, targetName := range []string{"generic", "codex", "human-team"} {
		t.Run(targetName, func(t *testing.T) {
			var stdout bytes.Buffer
			code := run([]string{"run", "--dir", dir, "--target", targetName}, &stdout, &bytes.Buffer{})
			if code != 0 {
				t.Fatalf("expected exit code 0, got %d", code)
			}
			if !strings.Contains(stdout.String(), "Target: "+targetName) {
				t.Fatalf("expected target output, got %q", stdout.String())
			}
			if len([]rune(stdout.String())) > 4000 {
				t.Fatalf("expected prompt <= 4000 chars, got %d", len([]rune(stdout.String())))
			}
		})
	}
}

func TestRunRejectsUnsupportedTarget(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stderr bytes.Buffer
	code := run([]string{"run", "--dir", dir, "--target", "shell"}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), `unsupported target "shell"`) {
		t.Fatalf("expected unsupported target error, got %q", stderr.String())
	}
}

func TestTargets(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"targets"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	for _, name := range []string{"generic", "codex", "human-team", "hyper-run", "namba-ai", "ouroboros", "spec-kit"} {
		if !strings.Contains(stdout.String(), name) {
			t.Fatalf("expected target %q in output, got %q", name, stdout.String())
		}
	}
}

func TestTargetsJSON(t *testing.T) {
	var stdout bytes.Buffer
	code := run([]string{"targets", "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload []map[string]string
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid JSON targets, got %v: %q", err, stdout.String())
	}
	if len(payload) != 7 {
		t.Fatalf("expected 7 targets, got %d: %#v", len(payload), payload)
	}
	if payload[0]["name"] != "generic" {
		t.Fatalf("expected built-in target order, got %#v", payload)
	}
}

func TestGraph(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"graph", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "work graph proposal") {
		t.Fatalf("expected graph output, got %q", stdout.String())
	}
	if !strings.Contains(stdout.String(), "CAP-001") {
		t.Fatalf("expected capability node, got %q", stdout.String())
	}
}

func TestGraphJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"graph", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"nodes"`) {
		t.Fatalf("expected JSON graph output, got %q", stdout.String())
	}
}

func TestHarnessPlan(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"harness", "plan", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "generated harness proposal") {
		t.Fatalf("expected harness output, got %q", stdout.String())
	}
}

func TestHarnessPlanJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"harness", "plan", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"source_lock_hash"`) {
		t.Fatalf("expected JSON harness output, got %q", stdout.String())
	}
}

func TestUnknownCommand(t *testing.T) {
	var stderr bytes.Buffer
	code := run([]string{"serve"}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "unknown command: serve") {
		t.Fatalf("expected unknown command error, got %q", stderr.String())
	}
}

func writeReadyContractForCLI(t *testing.T, dir string) {
	t.Helper()

	contract := map[string]any{
		"schema":            "ni.contract.v0",
		"readiness_profile": "prototype",
		"product_type":      "software",
		"delivery_surfaces": []string{"cli"},
		"interaction_mode":  "human_to_system",
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
