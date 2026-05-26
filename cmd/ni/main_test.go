package main

import (
	"bytes"
	"encoding/json"
	"io/fs"
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
		".ni/pressure.json",
		".ni/harness.candidates.json",
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

func TestFeedbackAddAndListJSON(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	for _, fixture := range []string{"codex.json", "human-team.json"} {
		var stdout bytes.Buffer
		code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", fixture)}, &stdout, &bytes.Buffer{})
		if code != 0 {
			t.Fatalf("expected exit code 0 adding %s, got %d", fixture, code)
		}
		if !strings.Contains(stdout.String(), "recorded feedback from") {
			t.Fatalf("expected feedback add summary, got %q", stdout.String())
		}
	}

	var textOut bytes.Buffer
	code := run([]string{"feedback", "list", "--dir", dir}, &textOut, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(textOut.String(), "codex") || !strings.Contains(textOut.String(), "human-team") {
		t.Fatalf("expected text feedback list to include both fixtures, got %q", textOut.String())
	}

	var stdout bytes.Buffer
	code = run([]string{"feedback", "list", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload []map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected valid feedback JSON, got %v: %q", err, stdout.String())
	}
	if len(payload) != 2 {
		t.Fatalf("expected two feedback entries, got %#v", payload)
	}
	if payload[0]["source_target"] != "codex" || payload[1]["source_target"] != "human-team" {
		t.Fatalf("expected codex and human-team feedback, got %#v", payload)
	}
}

func TestPressureStatusEmptyState(t *testing.T) {
	dir := t.TempDir()

	var stdout bytes.Buffer
	code := run([]string{"pressure", "status", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "no pressure recorded") {
		t.Fatalf("expected empty pressure status, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"pressure", "status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var payload map[string]any
	if err := json.Unmarshal(stdout.Bytes(), &payload); err != nil {
		t.Fatalf("expected pressure JSON, got %v: %q", err, stdout.String())
	}
	if payload["schema"] != "ni.pressure.v0" {
		t.Fatalf("expected pressure schema, got %#v", payload)
	}
	if items, ok := payload["items"].([]any); !ok || len(items) != 0 {
		t.Fatalf("expected no pressure items, got %#v", payload["items"])
	}
}

func TestFeedbackAddCreatesObservedPressure(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}

	code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code = run([]string{"pressure", "status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	var ledger struct {
		Schema string `json:"schema"`
		Items  []struct {
			ID                     string   `json:"id"`
			Kind                   string   `json:"kind"`
			Status                 string   `json:"status"`
			EvidenceRefs           []string `json:"evidence_refs"`
			RelatedCapabilities    []string `json:"related_capabilities"`
			RelatedRisks           []string `json:"related_risks"`
			ProposedAction         string   `json:"proposed_action"`
			RequiresUserAcceptance bool     `json:"requires_user_acceptance"`
		} `json:"items"`
	}
	if err := json.Unmarshal(stdout.Bytes(), &ledger); err != nil {
		t.Fatalf("expected pressure ledger JSON, got %v: %q", err, stdout.String())
	}
	if ledger.Schema != "ni.pressure.v0" {
		t.Fatalf("expected pressure schema, got %q", ledger.Schema)
	}
	if len(ledger.Items) != 4 {
		t.Fatalf("expected four pressure items from fixture, got %#v", ledger.Items)
	}
	seenKinds := map[string]bool{}
	for _, item := range ledger.Items {
		if item.ID == "" || item.ProposedAction == "" {
			t.Fatalf("expected pressure item fields to be set, got %#v", item)
		}
		if item.Status != "observed" {
			t.Fatalf("expected feedback pressure to start observed, got %#v", item)
		}
		if !item.RequiresUserAcceptance {
			t.Fatalf("expected pressure item to require user acceptance, got %#v", item)
		}
		if len(item.EvidenceRefs) != 1 || !strings.HasPrefix(item.EvidenceRefs[0], "feedback:codex:") {
			t.Fatalf("expected feedback evidence ref, got %#v", item)
		}
		if len(item.RelatedCapabilities) != 1 || item.RelatedCapabilities[0] != "CAP-021" {
			t.Fatalf("expected related capability from feedback, got %#v", item)
		}
		if item.RelatedRisks == nil {
			t.Fatalf("expected related_risks field to be present, got %#v", item)
		}
		seenKinds[item.Kind] = true
	}
	for _, kind := range []string{"recurring_blocker", "validation_gap", "planning_gap"} {
		if !seenKinds[kind] {
			t.Fatalf("expected pressure kind %s in %#v", kind, ledger.Items)
		}
	}
}

func TestPressurePromoteRequiresExplicitCommand(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	if code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("feedback add expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"pressure", "promote", "P-001", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "promoted P-001 to repeated") {
		t.Fatalf("expected one-step promotion summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"pressure", "status", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"status": "repeated"`) {
		t.Fatalf("expected promoted status in ledger, got %q", stdout.String())
	}
	if strings.Contains(stdout.String(), `"status": "accepted"`) {
		t.Fatalf("pressure became accepted without enough explicit promotes: %q", stdout.String())
	}
}

func TestPressureRetire(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	if code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("feedback add expected exit code 0, got %d", code)
	}

	var stdout bytes.Buffer
	code := run([]string{"pressure", "retire", "P-001", "--dir", dir}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "retired P-001") {
		t.Fatalf("expected retire summary, got %q", stdout.String())
	}
}

func TestFeedbackAddIsInert(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	contractPath := filepath.Join(dir, ".ni", "contract.json")
	lockPath := filepath.Join(dir, ".ni", "plan.lock.json")
	beforeContract := readFileForCLI(t, contractPath)
	beforeLock := readFileForCLI(t, lockPath)

	code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if got := readFileForCLI(t, contractPath); !bytes.Equal(got, beforeContract) {
		t.Fatalf("feedback add changed contract.json")
	}
	if got := readFileForCLI(t, lockPath); !bytes.Equal(got, beforeLock) {
		t.Fatalf("feedback add changed plan.lock.json")
	}
}

func TestFeedbackAddBlocksOnLockMismatch(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("changed after lock\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	var stderr bytes.Buffer
	code := run([]string{"feedback", "add", "--dir", dir, "--file", filepath.Join("..", "..", "testdata", "feedback", "codex.json")}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "BLOCKED: lock hash mismatch") {
		t.Fatalf("expected lock mismatch block, got %q", stderr.String())
	}
}

func TestExportHyperRunCreatesSeedDocsOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "hyper-seed")
	var stdout bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "hyper-run", "--out", out}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "exported hyper-run seed package") {
		t.Fatalf("expected export summary, got %q", stdout.String())
	}

	wantFiles := []string{
		"plan.md",
		"ni-context.md",
		"readiness-expectations.md",
		"evidence-requirements.md",
		"first-run-focus.md",
	}
	for _, name := range wantFiles {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("expected seed doc %s: %v", name, err)
		}
		if len(data) == 0 {
			t.Fatalf("expected seed doc %s to have content", name)
		}
	}

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	if len(entries) != len(wantFiles) {
		t.Fatalf("expected seed docs only, got %d entries", len(entries))
	}

	forbidden := map[string]struct{}{
		filepath.Join(".hyper", "goals", "GOAL-0001"): {},
		"tasks.md":    {},
		"evidence.md": {},
		"review.md":   {},
		"next.md":     {},
	}
	err = filepath.WalkDir(out, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(out, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		if _, ok := forbidden[rel]; ok {
			t.Fatalf("export created forbidden runtime packet path %s", rel)
		}
		return nil
	})
	if err != nil {
		t.Fatalf("walking export directory: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, ".hyper", "goals")); !os.IsNotExist(err) {
		t.Fatalf("expected no .hyper/goals directory, stat err: %v", err)
	}
}

func TestExportHyperRunRejectsExistingRuntimePacketDirectory(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "hyper-seed")
	if err := os.MkdirAll(filepath.Join(out, ".hyper", "goals", "GOAL-0001"), 0o755); err != nil {
		t.Fatalf("creating stale runtime directory: %v", err)
	}

	var stderr bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "hyper-run", "--out", out}, &bytes.Buffer{}, &stderr)
	if code != 1 {
		t.Fatalf("expected exit code 1, got %d", code)
	}
	if !strings.Contains(stderr.String(), "non-seed entry") {
		t.Fatalf("expected non-seed rejection, got %q", stderr.String())
	}
}

func TestExportNambaAICreatesSeedDocsOnly(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}

	out := filepath.Join(dir, "namba-seed")
	var stdout bytes.Buffer
	code := run([]string{"export", "--dir", dir, "--target", "namba-ai", "--out", out}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "exported namba-ai seed package") {
		t.Fatalf("expected export summary, got %q", stdout.String())
	}

	wantFiles := []string{
		"planning.md",
		"ni-lock-summary.md",
		"capability-map.md",
		"evaluation-map.md",
		"risk-map.md",
		"suggested-spec-boundaries.md",
	}
	for _, name := range wantFiles {
		data, err := os.ReadFile(filepath.Join(out, name))
		if err != nil {
			t.Fatalf("expected seed doc %s: %v", name, err)
		}
		if len(data) == 0 {
			t.Fatalf("expected seed doc %s to have content", name)
		}
	}

	planning, err := os.ReadFile(filepath.Join(out, "planning.md"))
	if err != nil {
		t.Fatalf("reading planning seed: %v", err)
	}
	for _, want := range []string{
		"namba-ai-oriented planning seed material",
		"NI does not call namba",
		"run, sync, pr, or land behavior",
		"does not add Codex-only assumptions",
	} {
		if !strings.Contains(string(planning), want) {
			t.Fatalf("expected planning.md to contain %q, got %q", want, string(planning))
		}
	}

	boundaries, err := os.ReadFile(filepath.Join(out, "suggested-spec-boundaries.md"))
	if err != nil {
		t.Fatalf("reading suggested boundaries: %v", err)
	}
	for _, want := range []string{
		"proposal, not a required sequential SPEC chain",
		"candidate graph boundaries",
		"depends_on",
	} {
		if !strings.Contains(string(boundaries), want) {
			t.Fatalf("expected suggested boundaries to contain %q, got %q", want, string(boundaries))
		}
	}

	entries, err := os.ReadDir(out)
	if err != nil {
		t.Fatalf("reading export directory: %v", err)
	}
	if len(entries) != len(wantFiles) {
		t.Fatalf("expected seed docs only, got %d entries", len(entries))
	}
	for _, forbidden := range []string{"run.md", "sync.md", "pr.md", "land.md"} {
		if _, err := os.Stat(filepath.Join(out, forbidden)); !os.IsNotExist(err) {
			t.Fatalf("expected no namba execution file %s, stat err: %v", forbidden, err)
		}
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

func TestHarnessCandidateCommands(t *testing.T) {
	dir := t.TempDir()
	if code := run([]string{"init", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("init expected exit code 0, got %d", code)
	}
	writeReadyContractForCLI(t, dir)
	if code := run([]string{"end", "--dir", dir}, &bytes.Buffer{}, &bytes.Buffer{}); code != 0 {
		t.Fatalf("end expected exit code 0, got %d", code)
	}
	writeAcceptedHarnessPressureForCLI(t, dir)

	var stdout bytes.Buffer
	code := run([]string{"harness", "propose", "--dir", dir, "--from-pressure", "P-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected propose exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "proposed harness candidate HC-001") {
		t.Fatalf("expected propose summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "accept", "--dir", dir, "HC-001"}, &stdout, &bytes.Buffer{})
	if code == 0 {
		t.Fatal("expected accept before validation to fail")
	}

	evidencePath := filepath.Join(dir, "harness-evidence.txt")
	if err := os.WriteFile(evidencePath, []byte("validated by CLI test\n"), 0o644); err != nil {
		t.Fatalf("writing evidence: %v", err)
	}
	stdout.Reset()
	code = run([]string{"harness", "validate", "--dir", dir, "HC-001", "--evidence", "harness-evidence.txt"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected validate exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "validated harness candidate HC-001 to validated_candidate") {
		t.Fatalf("expected validate summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "accept", "--dir", dir, "HC-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected accept exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "accepted harness candidate HC-001 as user_accepted") {
		t.Fatalf("expected accept summary, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "candidates", "--dir", dir, "--json"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected candidates exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), `"active_rule_id": "HC-001"`) {
		t.Fatalf("expected active rule id in candidates JSON, got %q", stdout.String())
	}

	stdout.Reset()
	code = run([]string{"harness", "retire", "--dir", dir, "HC-001"}, &stdout, &bytes.Buffer{})
	if code != 0 {
		t.Fatalf("expected retire exit code 0, got %d", code)
	}
	if !strings.Contains(stdout.String(), "retired harness candidate HC-001") {
		t.Fatalf("expected retire summary, got %q", stdout.String())
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

func writeAcceptedHarnessPressureForCLI(t *testing.T, dir string) {
	t.Helper()

	data := []byte(`{
  "schema": "ni.pressure.v0",
  "items": [
    {
      "id": "P-001",
      "kind": "harness_candidate",
      "status": "accepted",
      "evidence_refs": ["cli:test"],
      "related_capabilities": ["CAP-001"],
      "related_risks": ["RISK-001"],
      "proposed_action": "Create an inert downstream harness proposal for CLI tests.",
      "requires_user_acceptance": true
    }
  ]
}
`)
	if err := os.WriteFile(filepath.Join(dir, ".ni", "pressure.json"), data, 0o644); err != nil {
		t.Fatalf("writing pressure ledger: %v", err)
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

func readFileForCLI(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading %s: %v", path, err)
	}
	return data
}
