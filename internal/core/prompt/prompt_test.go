package prompt

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unicode/utf8"

	"ni/internal/core/docstore"
	"ni/internal/core/lock"
)

func TestCompileRequiresLockfile(t *testing.T) {
	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}

	_, err := Compile(Options{Dir: dir, MaxChars: 4000})
	if err == nil {
		t.Fatal("expected missing lockfile error")
	}
	if !strings.Contains(err.Error(), "missing lockfile") {
		t.Fatalf("expected missing lockfile error, got %v", err)
	}
}

func TestCompileRefusesStaleLock(t *testing.T) {
	dir := readyLockedProject(t)
	path := filepath.Join(dir, "docs", "plan", "00_project_brief.md")
	if err := os.WriteFile(path, []byte("changed\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	_, err := Compile(Options{Dir: dir, MaxChars: 4000})
	if err == nil {
		t.Fatal("expected stale lock error")
	}
	if !strings.Contains(err.Error(), "BLOCKED: lock hash mismatch") {
		t.Fatalf("expected lock hash mismatch error, got %v", err)
	}
}

func TestCompileHonorsPromptBudget(t *testing.T) {
	dir := readyLockedProject(t)

	result, err := Compile(Options{Dir: dir, MaxChars: 240})
	if err != nil {
		t.Fatalf("Compile returned error: %v", err)
	}
	if utf8.RuneCountInString(result.Prompt) > 240 {
		t.Fatalf("prompt exceeded budget: %d", utf8.RuneCountInString(result.Prompt))
	}
	if strings.Contains(result.Prompt, "# Project brief") {
		t.Fatalf("prompt should not paste full planning docs: %q", result.Prompt)
	}
}

func TestCompileWritesOutFile(t *testing.T) {
	dir := readyLockedProject(t)
	out := filepath.Join(dir, ".ni", "generated", "goal.prompt.txt")

	result, err := Compile(Options{Dir: dir, Out: out, MaxChars: 4000})
	if err != nil {
		t.Fatalf("Compile returned error: %v", err)
	}
	if result.Out != out {
		t.Fatalf("expected out %q, got %q", out, result.Out)
	}
	data, err := os.ReadFile(out)
	if err != nil {
		t.Fatalf("reading prompt output: %v", err)
	}
	if !strings.Contains(string(data), "Generic target prompt") {
		t.Fatalf("unexpected prompt output: %q", string(data))
	}
}

func TestCompileTargetsStayWithinPromptBudget(t *testing.T) {
	dir := readyLockedProject(t)

	for _, targetName := range []string{"generic", "codex", "human-team"} {
		t.Run(targetName, func(t *testing.T) {
			result, err := Compile(Options{Dir: dir, Target: targetName, MaxChars: DefaultMaxChars})
			if err != nil {
				t.Fatalf("Compile returned error: %v", err)
			}
			if utf8.RuneCountInString(result.Prompt) > DefaultMaxChars {
				t.Fatalf("prompt exceeded budget: %d", utf8.RuneCountInString(result.Prompt))
			}
			if !strings.Contains(result.Prompt, "Target: "+targetName) {
				t.Fatalf("expected target marker in prompt, got %q", result.Prompt)
			}
		})
	}
}

func TestCompileTargetTemplatesContainRequiredAuthorityRules(t *testing.T) {
	dir := readyLockedProject(t)

	for _, targetName := range []string{"generic", "codex", "human-team"} {
		t.Run(targetName, func(t *testing.T) {
			result, err := Compile(Options{Dir: dir, Target: targetName, MaxChars: DefaultMaxChars})
			if err != nil {
				t.Fatalf("Compile returned error: %v", err)
			}
			assertPromptContains(t, result.Prompt,
				".ni/plan.lock.json is authoritative",
				".ni/contract.json",
				"docs/plan",
				"Do not weaken acceptance criteria",
				"on lock mismatch report BLOCKED",
				"If there are conflicting requirements, report BLOCKED",
				"pre-runtime boundary",
			)
		})
	}
}

func TestCompileCodexTemplateIsPasteReadyWithoutAutoInvocation(t *testing.T) {
	dir := readyLockedProject(t)

	result, err := Compile(Options{Dir: dir, Target: "codex", MaxChars: DefaultMaxChars})
	if err != nil {
		t.Fatalf("Compile returned error: %v", err)
	}
	assertPromptContains(t, result.Prompt,
		"Paste this prompt into Codex",
		"Do not ask Namba Intent to invoke Codex automatically",
		"namba-intent run only compiled this prompt",
	)
}

func TestCompileHumanTeamTemplateNamesHandoffRoles(t *testing.T) {
	dir := readyLockedProject(t)

	result, err := Compile(Options{Dir: dir, Target: "human-team", MaxChars: DefaultMaxChars})
	if err != nil {
		t.Fatalf("Compile returned error: %v", err)
	}
	assertPromptContains(t, result.Prompt,
		"PM/dev/design/research team",
		"PM:",
		"Dev:",
		"Design/research:",
		"team handoff",
	)
}

func TestCompileRejectsUnsupportedTarget(t *testing.T) {
	dir := readyLockedProject(t)

	_, err := Compile(Options{Dir: dir, Target: "shell", MaxChars: DefaultMaxChars})
	if err == nil {
		t.Fatal("expected unsupported target error")
	}
	if !strings.Contains(err.Error(), `unsupported target "shell"`) {
		t.Fatalf("expected unsupported target error, got %v", err)
	}
}

func TestCompileValidatesTargetBeforeLockfile(t *testing.T) {
	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}

	_, err := Compile(Options{Dir: dir, Target: "shell", MaxChars: DefaultMaxChars})
	if err == nil {
		t.Fatal("expected unsupported target error")
	}
	if !strings.Contains(err.Error(), `unsupported target "shell"`) {
		t.Fatalf("expected unsupported target error, got %v", err)
	}
}

func readyLockedProject(t *testing.T) string {
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
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("# Project brief\n\n## Product type\n\nsoftware\n\n## Delivery surfaces\n\n- cli\n\n## Purpose\n\nExercise readiness rules.\n"), 0o644); err != nil {
		t.Fatalf("writing ready project brief: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "01_actors_outcomes.md"), []byte("# Actors and outcomes\n\n## Actors\n\n- User: reviews the prompt fixture.\n- CLI: validates readiness.\n\n## Outcomes\n\n- Exercise readiness rules.\n"), 0o644); err != nil {
		t.Fatalf("writing ready actors doc: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "08_delivery_operation.md"), []byte("# Delivery and operation\n\n## Delivery surfaces\n\n- cli\n\n## Initial delivery\n\nThe prompt fixture is reviewed before lock.\n"), 0o644); err != nil {
		t.Fatalf("writing ready delivery doc: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "10_open_questions.md"), []byte("# Open questions\n\nNo open questions are listed in this fixture.\n"), 0o644); err != nil {
		t.Fatalf("writing ready open question doc: %v", err)
	}
	if _, err := lock.CreateAt(dir, time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC)); err != nil {
		t.Fatalf("creating lock: %v", err)
	}
	return dir
}

func assertPromptContains(t *testing.T, prompt string, want ...string) {
	t.Helper()
	for _, item := range want {
		if !strings.Contains(prompt, item) {
			t.Fatalf("expected prompt to contain %q, got %q", item, prompt)
		}
	}
}
