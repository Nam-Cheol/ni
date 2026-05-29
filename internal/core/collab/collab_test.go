package collab

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"ni/internal/core/lock"
)

func TestDiffReportsNonConflictingParallelChanges(t *testing.T) {
	result, err := Diff(fixture("base.json"), fixture("non_conflicting_parallel_head.json"))
	if err != nil {
		t.Fatalf("diff failed: %v", err)
	}
	if len(result.Changes) == 0 {
		t.Fatal("expected added parallel planning changes")
	}
	if !hasChange(result.Changes, "added", "capability", "CAP-002") {
		t.Fatalf("expected CAP-002 added change, got %#v", result.Changes)
	}

	conflicts, err := Conflicts(fixture("base.json"), fixture("non_conflicting_parallel_head.json"))
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	if len(conflicts.Conflicts) != 0 {
		t.Fatalf("expected non-conflicting parallel additions, got %#v", conflicts.Conflicts)
	}
}

func TestConflictsDetectDecisionChanges(t *testing.T) {
	result, err := Conflicts(fixture("base.json"), fixture("conflicting_decision_head.json"))
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	requireConflict(t, result.Conflicts, "same_id_changed", "DEC-001")
	requireConflict(t, result.Conflicts, "contradictory_decision", "DEC-002")
	if !strings.Contains(FormatConflicts(result), "collaboration conflicts") {
		t.Fatalf("expected human conflict output, got %q", FormatConflicts(result))
	}
}

func TestConflictsDetectWeakenedAcceptanceCriteria(t *testing.T) {
	result, err := Conflicts(fixture("base.json"), fixture("weakened_acceptance_head.json"))
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	requireConflict(t, result.Conflicts, "acceptance_weakened", "REQ-001")
}

func TestConflictsDetectRemovedCapabilityWithLiveReferences(t *testing.T) {
	dir := t.TempDir()
	base := readContractFixture(t, "base.json")
	head := readContractFixture(t, "base.json")
	delete(head, "capabilities")
	head["capabilities"] = []any{}
	basePath := writeContract(t, dir, "base.json", base)
	headPath := writeContract(t, dir, "head.json", head)

	result, err := Conflicts(basePath, headPath)
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	requireConflict(t, result.Conflicts, "removed_capability_referenced", "CAP-001")
}

func TestConflictsDetectLoweredRiskSeverityWithoutNewMitigation(t *testing.T) {
	dir := t.TempDir()
	base := readContractFixture(t, "base.json")
	head := readContractFixture(t, "base.json")
	risks := head["risks"].([]any)
	risk := risks[0].(map[string]any)
	risk["severity"] = "low"
	headPath := writeContract(t, dir, "head.json", head)
	basePath := writeContract(t, dir, "base.json", base)

	result, err := Conflicts(basePath, headPath)
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	requireConflict(t, result.Conflicts, "risk_severity_lowered", "RISK-001")
}

func TestConflictsDetectLockHashMismatch(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, ".ni"), 0o755); err != nil {
		t.Fatalf("creating .ni: %v", err)
	}
	copyFile(t, fixture("base.json"), filepath.Join(dir, ".ni", "contract.json"))
	for _, doc := range []string{
		"docs/plan/00_project_brief.md",
		"docs/plan/01_actors_outcomes.md",
		"docs/plan/03_interaction_contract.md",
		"docs/plan/04_domain_state.md",
		"docs/plan/05_constraints.md",
		"docs/plan/08_delivery_operation.md",
		"docs/plan/09_execution_strategy.md",
		"docs/plan/10_open_questions.md",
	} {
		path := filepath.Join(dir, doc)
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatalf("creating doc dir: %v", err)
		}
		if err := os.WriteFile(path, []byte("# fixture\n"), 0o644); err != nil {
			t.Fatalf("writing doc: %v", err)
		}
	}
	writePlanDoc(t, dir, "docs/plan/02_capabilities.md", "# Capabilities\n\n## CAP-001: Prompt compiler\n\nCompile prompts.\n")
	writePlanDoc(t, dir, "docs/plan/00_project_brief.md", "# Project brief\n\n## Product type\n\nsoftware\n\n## Delivery surfaces\n\n- cli\n\n## Purpose\n\nExercise collaboration conflict detection.\n")
	writePlanDoc(t, dir, "docs/plan/01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- CLI user: reviews collaboration conflict detection.\n- CLI: reports lock hash mismatches.\n\n## Outcomes\n\n- Exercise collaboration conflict detection.\n")
	writePlanDoc(t, dir, "docs/plan/06_risks_security.md", "# Risks and security\n\n## RISK-001: Prompt may exceed budget\n\nSeverity: high\n\nMitigation: Enforce maximum prompt length.\n")
	writePlanDoc(t, dir, "docs/plan/07_evaluation_contract.md", "# Evaluation contract\n\n## EVAL-001: Prompt budget check\n\nMethod: automated test\n")
	writePlanDoc(t, dir, "docs/plan/08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- cli\n\n## Initial delivery\n\nThe CLI prompt compiler is reviewed before lock.\n")
	writePlanDoc(t, dir, "docs/plan/11_decision_log.md", "# Decision log\n\n## DEC-001: Use prompt compiler only\n\nStatus: accepted\n")
	if _, err := lock.CreateAt(dir, time.Date(2026, 5, 26, 0, 0, 0, 0, time.UTC)); err != nil {
		t.Fatalf("creating lock: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "02_capabilities.md"), []byte("# changed\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	result, err := Conflicts(dir, fixture("base.json"))
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	requireConflict(t, result.Conflicts, "lock_hash_mismatch", "docs/plan/02_capabilities.md")
}

func TestFormatsEmptyResults(t *testing.T) {
	diff, err := Diff(fixture("base.json"), fixture("base.json"))
	if err != nil {
		t.Fatalf("diff failed: %v", err)
	}
	if got := FormatDiff(diff); got != "no contract changes\n" {
		t.Fatalf("expected empty diff text, got %q", got)
	}
	conflicts, err := Conflicts(fixture("base.json"), fixture("base.json"))
	if err != nil {
		t.Fatalf("conflicts failed: %v", err)
	}
	if got := FormatConflicts(conflicts); got != "no collaboration conflicts\n" {
		t.Fatalf("expected empty conflicts text, got %q", got)
	}
}

func fixture(name string) string {
	return filepath.Join("testdata", name)
}

func hasChange(changes []Change, kind string, entityType string, id string) bool {
	for _, change := range changes {
		if change.Kind == kind && change.EntityType == entityType && change.ID == id {
			return true
		}
	}
	return false
}

func requireConflict(t *testing.T, conflicts []Conflict, code string, id string) {
	t.Helper()
	for _, conflict := range conflicts {
		if conflict.Code == code && conflict.ID == id {
			return
		}
	}
	t.Fatalf("expected conflict %s %s, got %#v", code, id, conflicts)
}

func readContractFixture(t *testing.T, name string) map[string]any {
	t.Helper()
	data, err := os.ReadFile(fixture(name))
	if err != nil {
		t.Fatalf("reading fixture: %v", err)
	}
	var payload map[string]any
	if err := json.Unmarshal(data, &payload); err != nil {
		t.Fatalf("parsing fixture: %v", err)
	}
	return payload
}

func writeContract(t *testing.T, dir string, name string, payload map[string]any) string {
	t.Helper()
	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		t.Fatalf("marshaling contract: %v", err)
	}
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, append(data, '\n'), 0o644); err != nil {
		t.Fatalf("writing contract: %v", err)
	}
	return path
}

func copyFile(t *testing.T, from string, to string) {
	t.Helper()
	data, err := os.ReadFile(from)
	if err != nil {
		t.Fatalf("reading %s: %v", from, err)
	}
	if err := os.WriteFile(to, data, 0o644); err != nil {
		t.Fatalf("writing %s: %v", to, err)
	}
}

func writePlanDoc(t *testing.T, dir string, relPath string, content string) {
	t.Helper()
	path := filepath.Join(dir, relPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("creating doc dir: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writing %s: %v", relPath, err)
	}
}
