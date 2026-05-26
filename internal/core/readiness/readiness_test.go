package readiness

import (
	"os"
	"path/filepath"
	"testing"

	"ni/internal/core/docstore"
)

func TestEvaluateReadyFixture(t *testing.T) {
	dir := initFixtureProject(t, "ready.json")

	result := Evaluate(dir)
	if result.Status != StatusReady {
		t.Fatalf("expected READY, got %#v", result)
	}
}

func TestEvaluateReadyWithDeferralsFixture(t *testing.T) {
	dir := initFixtureProject(t, "ready_with_deferrals.json")

	result := Evaluate(dir)
	if result.Status != StatusReadyWithDeferrals {
		t.Fatalf("expected READY_WITH_DEFERRALS, got %#v", result)
	}
}

func TestEvaluateBlocksMissingPlanningDoc(t *testing.T) {
	dir := initFixtureProject(t, "ready.json")
	if err := os.Remove(filepath.Join(dir, "docs", "plan", "02_capabilities.md")); err != nil {
		t.Fatalf("removing planning doc: %v", err)
	}

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R001")
}

func TestEvaluateBlocksInvalidContract(t *testing.T) {
	dir := initFixtureProject(t, "ready.json")
	writeContract(t, dir, []byte(`{"schema":`))

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R002")
}

func TestEvaluateBlocksMissingCapability(t *testing.T) {
	dir := initFixtureProject(t, "missing_capability.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R003")
}

func TestEvaluateBlocksCapabilityWithoutEvaluation(t *testing.T) {
	dir := initFixtureProject(t, "capability_without_evaluation.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R004")
}

func TestEvaluateBlocksEvaluationWithoutMethod(t *testing.T) {
	dir := initFixtureProject(t, "evaluation_without_method.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R005")
}

func TestEvaluateBlocksHighRiskWithoutMitigation(t *testing.T) {
	dir := initFixtureProject(t, "high_risk_without_mitigation.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R006")
}

func TestEvaluateBlocksAcceptedCapabilityWithoutArtifactOrRequirement(t *testing.T) {
	dir := initFixtureProject(t, "capability_without_artifact_or_requirement.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R007")
}

func TestEvaluateBlocksInvalidDecisionStatus(t *testing.T) {
	dir := initFixtureProject(t, "invalid_decision_status.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R008")
}

func TestEvaluateBlocksBlockerOpenQuestion(t *testing.T) {
	dir := initFixtureProject(t, "blocker_open_question.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R009")
}

func TestEvaluateBlocksMissingNonGoal(t *testing.T) {
	dir := initFixtureProject(t, "missing_non_goal.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R010")
}

func initFixtureProject(t *testing.T, fixture string) string {
	t.Helper()

	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing fixture project: %v", err)
	}
	data, err := os.ReadFile(filepath.Join("testdata", fixture))
	if err != nil {
		t.Fatalf("reading fixture %s: %v", fixture, err)
	}
	writeContract(t, dir, data)
	return dir
}

func writeContract(t *testing.T, dir string, data []byte) {
	t.Helper()

	path := filepath.Join(dir, ".ni", "contract.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("writing contract fixture: %v", err)
	}
}

func requireIssue(t *testing.T, result Result, status Status, ruleID string) {
	t.Helper()

	if result.Status != status {
		t.Fatalf("expected status %s, got %#v", status, result)
	}
	for _, issue := range result.Issues {
		if issue.RuleID == ruleID {
			return
		}
	}
	t.Fatalf("expected issue %s, got %#v", ruleID, result.Issues)
}
