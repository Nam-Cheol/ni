package readiness

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"ni/internal/core/contract"
	"ni/internal/core/docstore"
)

func TestEvaluateReadyFixture(t *testing.T) {
	dir := initFixtureProject(t, "ready.json")

	result := Evaluate(dir)
	if result.Status != StatusReady {
		t.Fatalf("expected READY, got %#v", result)
	}
	if result.Profile != "prototype" {
		t.Fatalf("expected prototype profile, got %#v", result)
	}
}

func TestEvaluateUniversalProductFixtures(t *testing.T) {
	for _, tt := range []struct {
		fixture     string
		productType string
		surface     string
		guidance    string
	}{
		{"conversation_product.json", "conversation_product", "conversation", "conversation turns"},
		{"research_protocol.json", "research_protocol", "document", "hypothesis"},
		{"software_cli.json", "software", "cli", "readiness authority remains"},
	} {
		t.Run(tt.fixture, func(t *testing.T) {
			dir := initFixtureProject(t, tt.fixture)

			result := Evaluate(dir)
			if result.Status != StatusReady {
				t.Fatalf("expected READY, got %#v", result)
			}
			if result.ProductType != tt.productType {
				t.Fatalf("expected product type %s, got %#v", tt.productType, result)
			}
			if !containsString(result.DeliverySurfaces, tt.surface) {
				t.Fatalf("expected surface %s, got %#v", tt.surface, result.DeliverySurfaces)
			}
			if !guidanceContains(result.Guidance, tt.guidance) {
				t.Fatalf("expected guidance containing %q, got %#v", tt.guidance, result.Guidance)
			}
		})
	}
}

func TestEvaluateReadyWithDeferralsFixture(t *testing.T) {
	dir := initFixtureProject(t, "ready_with_deferrals.json")

	result := Evaluate(dir)
	if result.Status != StatusReadyWithDeferrals {
		t.Fatalf("expected READY_WITH_DEFERRALS, got %#v", result)
	}
}

func TestEvaluateConceptProfileTreatsTraceabilityGapAsDeferral(t *testing.T) {
	dir := initFixtureProject(t, "capability_without_evaluation.json")
	setContractProfile(t, dir, "concept")

	result := Evaluate(dir)
	if result.Status != StatusReadyWithDeferrals {
		t.Fatalf("expected READY_WITH_DEFERRALS, got %#v", result)
	}
	requireIssueSeverity(t, result, "R004", "deferral")
}

func TestEvaluateProductionProfileTreatsOpenDeferralsAsBlockers(t *testing.T) {
	dir := initFixtureProject(t, "ready_with_deferrals.json")
	setContractProfile(t, dir, "production")

	result := Evaluate(dir)
	if result.Status != StatusBlocked {
		t.Fatalf("expected BLOCKED, got %#v", result)
	}
	requireIssueSeverity(t, result, "D001", "blocker")
	requireIssueSeverity(t, result, "D002", "blocker")
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

func TestNextQuestionsFromRuleFailures(t *testing.T) {
	for _, tt := range []struct {
		name      string
		fixture   string
		ruleID    string
		reference string
		want      string
	}{
		{
			name:      "missing capability evaluation",
			fixture:   "capability_without_evaluation.json",
			ruleID:    "R004",
			reference: "CAP-001",
			want:      "what evidence proves this capability works",
		},
		{
			name:      "high risk mitigation",
			fixture:   "high_risk_without_mitigation.json",
			ruleID:    "R006",
			reference: "RISK-001",
			want:      "what mitigation, owner, or explicit accepted-risk decision is required",
		},
		{
			name:      "blocker open question",
			fixture:   "blocker_open_question.json",
			ruleID:    "R009",
			reference: "OQ-001",
			want:      "what decision resolves this blocker",
		},
		{
			name:      "missing non-goal",
			fixture:   "missing_non_goal.json",
			ruleID:    "R010",
			reference: "",
			want:      "what must this project explicitly avoid",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dir := initFixtureProject(t, tt.fixture)

			questions := NextQuestions(Evaluate(dir))
			question := requireQuestion(t, questions, tt.ruleID)
			if tt.reference != "" && !containsString(question.References, tt.reference) {
				t.Fatalf("expected reference %s, got %#v", tt.reference, question)
			}
			if !strings.Contains(question.Question, tt.want) {
				t.Fatalf("expected question containing %q, got %#v", tt.want, question)
			}
			if strings.Contains(question.Question, "implement") {
				t.Fatalf("question should not imply implementation, got %#v", question)
			}
		})
	}
}

func TestEvaluateConsistentDocsContractSyncFixture(t *testing.T) {
	dir := initSyncFixtureProject(t, "consistent")

	result := Evaluate(dir)
	if result.Status != StatusReady {
		t.Fatalf("expected READY, got %#v", result)
	}
}

func TestEvaluateBlocksDocsContractSyncFixture(t *testing.T) {
	dir := initSyncFixtureProject(t, "decision_conflicts_contract")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R012")
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
	if c, err := contract.Load(data); err == nil {
		writePlanDocsForContract(t, dir, c)
	}
	return dir
}

func initSyncFixtureProject(t *testing.T, fixture string) string {
	t.Helper()

	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing sync fixture project: %v", err)
	}
	root := filepath.Join("..", "docsync", "testdata", fixture)
	copyFixtureFile(t, root, dir, ".ni/contract.json")
	copyFixtureFile(t, root, dir, "docs/plan/02_capabilities.md")
	copyFixtureFile(t, root, dir, "docs/plan/06_risks_security.md")
	copyFixtureFile(t, root, dir, "docs/plan/07_evaluation_contract.md")
	copyFixtureFile(t, root, dir, "docs/plan/11_decision_log.md")
	return dir
}

func copyFixtureFile(t *testing.T, fixtureRoot string, dir string, relPath string) {
	t.Helper()

	data, err := os.ReadFile(filepath.Join(fixtureRoot, relPath))
	if err != nil {
		t.Fatalf("reading fixture %s: %v", relPath, err)
	}
	if err := os.WriteFile(filepath.Join(dir, relPath), data, 0o644); err != nil {
		t.Fatalf("writing fixture %s: %v", relPath, err)
	}
}

func writeContract(t *testing.T, dir string, data []byte) {
	t.Helper()

	path := filepath.Join(dir, ".ni", "contract.json")
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("writing contract fixture: %v", err)
	}
}

func writePlanDocsForContract(t *testing.T, dir string, c contract.Contract) {
	t.Helper()

	var capabilities strings.Builder
	capabilities.WriteString("# Capabilities\n\n")
	for _, capability := range c.Capabilities {
		capabilities.WriteString("## ")
		capabilities.WriteString(capability.ID)
		capabilities.WriteString(": ")
		capabilities.WriteString(capability.Title)
		capabilities.WriteString("\n\nDescribe the accepted capability.\n\n")
	}
	writePlanDoc(t, dir, "02_capabilities.md", capabilities.String())

	var risks strings.Builder
	risks.WriteString("# Risks and security\n\n")
	for _, risk := range c.Risks {
		risks.WriteString("## ")
		risks.WriteString(risk.ID)
		risks.WriteString(": ")
		risks.WriteString(risk.Title)
		risks.WriteString("\n\nSeverity: ")
		risks.WriteString(risk.Severity)
		risks.WriteString("\n\nMitigation: ")
		risks.WriteString(risk.Mitigation)
		risks.WriteString("\n\n")
	}
	writePlanDoc(t, dir, "06_risks_security.md", risks.String())

	var evaluations strings.Builder
	evaluations.WriteString("# Evaluation contract\n\n")
	for _, evaluation := range c.Evaluations {
		evaluations.WriteString("## ")
		evaluations.WriteString(evaluation.ID)
		evaluations.WriteString(": ")
		evaluations.WriteString(evaluation.Title)
		evaluations.WriteString("\n\nMethod: ")
		evaluations.WriteString(evaluation.Method)
		evaluations.WriteString("\n\n")
	}
	writePlanDoc(t, dir, "07_evaluation_contract.md", evaluations.String())

	var decisions strings.Builder
	decisions.WriteString("# Decision log\n\n")
	for _, decision := range c.Decisions {
		decisions.WriteString("## ")
		decisions.WriteString(decision.ID)
		decisions.WriteString(": ")
		decisions.WriteString(decision.Title)
		decisions.WriteString("\n\nStatus: ")
		decisions.WriteString(decision.Status)
		decisions.WriteString("\n\n")
	}
	writePlanDoc(t, dir, "11_decision_log.md", decisions.String())
}

func writePlanDoc(t *testing.T, dir string, name string, content string) {
	t.Helper()

	path := filepath.Join(dir, "docs", "plan", name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writing plan doc %s: %v", name, err)
	}
}

func setContractProfile(t *testing.T, dir string, readinessProfile string) {
	t.Helper()

	path := filepath.Join(dir, ".ni", "contract.json")
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading contract fixture: %v", err)
	}
	updated := strings.Replace(string(data), `"readiness_profile": "prototype"`, `"readiness_profile": "`+readinessProfile+`"`, 1)
	if updated == string(data) {
		t.Fatalf("contract fixture did not contain prototype profile")
	}
	if err := os.WriteFile(path, []byte(updated), 0o644); err != nil {
		t.Fatalf("writing updated contract fixture: %v", err)
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

func requireIssueSeverity(t *testing.T, result Result, ruleID string, severity string) {
	t.Helper()

	for _, issue := range result.Issues {
		if issue.RuleID == ruleID {
			if issue.Severity != severity {
				t.Fatalf("expected %s severity %s, got %#v", ruleID, severity, issue)
			}
			return
		}
	}
	t.Fatalf("expected issue %s, got %#v", ruleID, result.Issues)
}

func requireQuestion(t *testing.T, questions []NextQuestion, ruleID string) NextQuestion {
	t.Helper()

	for _, question := range questions {
		if question.RuleID == ruleID {
			return question
		}
	}
	t.Fatalf("expected question %s, got %#v", ruleID, questions)
	return NextQuestion{}
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func guidanceContains(values []string, want string) bool {
	for _, value := range values {
		if strings.Contains(value, want) {
			return true
		}
	}
	return false
}
