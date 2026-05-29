package readiness

import (
	"fmt"
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

func TestEvaluateBlocksConflictingAcceptedDecision(t *testing.T) {
	dir := initFixtureProject(t, "conflicting_decision.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R013")
}

func TestEvaluateBlocksMissingNonGoal(t *testing.T) {
	dir := initFixtureProject(t, "missing_non_goal.json")

	result := Evaluate(dir)
	requireIssue(t, result, StatusBlocked, "R010")
}

func TestProofFromRuleFailures(t *testing.T) {
	for _, tt := range []struct {
		name    string
		fixture string
		ruleID  string
		want    string
	}{
		{"missing evaluation", "capability_without_evaluation.json", "R004", "CAP-001 is accepted but has no linked EVAL."},
		{"high risk", "high_risk_without_mitigation.json", "R006", "RISK-001 is high severity but has no mitigation."},
		{"blocker open question", "blocker_open_question.json", "R009", "OQ-001 is marked as blocker."},
		{"conflicting decision", "conflicting_decision.json", "R013", "DEC-002 conflicts with DEC-001."},
	} {
		t.Run(tt.name, func(t *testing.T) {
			dir := initFixtureProject(t, tt.fixture)

			item := requireProof(t, Proof(Evaluate(dir)), tt.ruleID)
			if item.Message != tt.want {
				t.Fatalf("expected proof %q, got %#v", tt.want, item)
			}
		})
	}
}

func TestProofFromDocsContractSyncMismatch(t *testing.T) {
	dir := initSyncFixtureProject(t, "decision_conflicts_contract")

	item := requireProof(t, Proof(Evaluate(dir)), "R012")
	if item.Severity != "blocker" {
		t.Fatalf("expected blocker sync proof, got %#v", item)
	}
	if !strings.Contains(item.Message, "DEC-001 status") {
		t.Fatalf("expected docs/contract mismatch proof, got %#v", item)
	}
	if item.SyncDiagnostic == nil || item.SyncDiagnostic.ID != "DEC-001" || item.SyncDiagnostic.Location == "" || item.SyncDiagnostic.SuggestedRepair == "" || !item.SyncDiagnostic.BlocksEnd {
		t.Fatalf("expected stable sync diagnostic fields, got %#v", item.SyncDiagnostic)
	}
}

func TestProofReadyWithDeferralsPlan(t *testing.T) {
	dir := initFixtureProject(t, "ready_with_deferrals.json")

	proof := Proof(Evaluate(dir))
	requireProof(t, proof, "D001")
	requireProof(t, proof, "D002")
}

func TestProofReadyPlan(t *testing.T) {
	dir := initFixtureProject(t, "ready.json")

	proof := Proof(Evaluate(dir))
	if len(proof) != 1 {
		t.Fatalf("expected one ready proof item, got %#v", proof)
	}
	if proof[0].RuleID != "READY" || !strings.Contains(proof[0].Message, "All readiness, sync, and conflict rules passed") {
		t.Fatalf("expected ready proof, got %#v", proof)
	}
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
			want:      "What evidence would prove this capability is complete",
		},
		{
			name:      "high risk mitigation",
			fixture:   "high_risk_without_mitigation.json",
			ruleID:    "R006",
			reference: "RISK-001",
			want:      "What mitigation would reduce or monitor it",
		},
		{
			name:      "blocker open question",
			fixture:   "blocker_open_question.json",
			ruleID:    "R009",
			reference: "OQ-001",
			want:      "What answer would resolve it",
		},
		{
			name:      "conflicting decision",
			fixture:   "conflicting_decision.json",
			ruleID:    "R013",
			reference: "DEC-001",
			want:      "Which accepted decision should be revised",
		},
		{
			name:      "missing non-goal",
			fixture:   "missing_non_goal.json",
			ruleID:    "R010",
			reference: "",
			want:      "What explicit non-goal should bound the plan",
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

func TestNextQuestionsFocusedOutputFixture(t *testing.T) {
	dir := initFixtureProject(t, "question_output.json")
	writePlanDoc(t, dir, "01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: TODO\n\n## Outcomes\n\n- TODO\n")
	writePlanDoc(t, dir, "08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- TODO\n")

	questions := NextQuestions(Evaluate(dir))
	want := map[string]string{
		"R014": "project.purpose is missing a concrete purpose. What should change, for whom, and why does it matter?",
		"R004": "CAP-001 has no evaluation. What evidence would prove this capability is complete: a test, review checklist, demo condition, user approval, or an explicit deferral?",
		"R006": "RISK-001 is high severity and has no mitigation. What mitigation would reduce or monitor it, who owns it, or should this become an explicit accepted-risk decision?",
		"D001": "DEC-001 is deferred. Should it remain deferred with a reason, become an accepted or rejected decision, or be marked not_applicable?",
		"R009": "OQ-001 is blocking readiness. What answer would resolve it: an accepted decision, a deferral with reason, not_applicable, or keeping it blocking with the missing information named?",
		"R010": "No non-goal is recorded. What explicit non-goal should bound the plan, or why is this boundary not_applicable?",
		"R015": "docs/plan/01_actors_outcomes.md is missing an actor or outcome. Which actor needs what outcome, and should that record be accepted, kept as evidence, deferred, or marked not_applicable?",
		"R016": "docs/plan/08_delivery_operation.md is missing a delivery surface. Should the plan target a CLI, web app, conversation, document, workflow, research protocol, human service, another surface, or a deferral with reason?",
	}
	for ruleID, question := range want {
		got := requireQuestion(t, questions, ruleID)
		if got.Question != question {
			t.Fatalf("expected %s question %q, got %#v", ruleID, question, got)
		}
		if strings.Contains(strings.ToLower(got.Question), "implement") {
			t.Fatalf("question should not imply implementation, got %#v", got)
		}
	}
}

func TestNextQuestionsDocsContractMismatchIncludesRepairChoices(t *testing.T) {
	dir := initSyncFixtureProject(t, "decision_conflicts_contract")

	question := requireQuestion(t, NextQuestions(Evaluate(dir)), "R012")
	if !containsString(question.References, "DEC-001") {
		t.Fatalf("expected DEC-001 reference, got %#v", question)
	}
	if !strings.Contains(question.Question, "Which source is correct") || !strings.Contains(question.Question, "update docs") || !strings.Contains(question.Question, "update the contract") {
		t.Fatalf("expected docs/contract repair choices, got %#v", question)
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
	c, err := contract.LoadFile(filepath.Join(dir, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("loading sync fixture contract: %v", err)
	}
	writePlanDoc(t, dir, "00_project_brief.md", "# Project brief\n\n## Product type\n\n"+c.ProductType+"\n\n## Delivery surfaces\n\n- "+strings.Join(c.DeliverySurfaces, "\n- ")+"\n\n## Purpose\n\n"+c.Project.Purpose+"\n")
	writePlanDoc(t, dir, "01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: reviews sync fixtures.\n- CLI: validates readiness.\n\n## Outcomes\n\n- "+c.Project.Purpose+"\n")
	writePlanDoc(t, dir, "08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- "+strings.Join(c.DeliverySurfaces, "\n- ")+"\n\n## Initial delivery\n\nPlanning docs and contract records are reviewed before lock.\n")
	writeOpenQuestionDocForContract(t, dir, c)
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

	writePlanDoc(t, dir, "00_project_brief.md", "# Project brief\n\n## Product type\n\n"+c.ProductType+"\n\n## Delivery surfaces\n\n- "+strings.Join(c.DeliverySurfaces, "\n- ")+"\n\n## Purpose\n\n"+c.Project.Purpose+"\n")

	writePlanDoc(t, dir, "01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: reviews and accepts the planning contract.\n- CLI: validates deterministic readiness.\n\n## Outcomes\n\n- "+c.Project.Purpose+"\n")

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

	writeOpenQuestionDocForContract(t, dir, c)

	writePlanDoc(t, dir, "08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- "+strings.Join(c.DeliverySurfaces, "\n- ")+"\n\n## Initial delivery\n\nPlanning docs and contract records are reviewed before lock.\n")
}

func writeOpenQuestionDocForContract(t *testing.T, dir string, c contract.Contract) {
	t.Helper()

	var openQuestions strings.Builder
	openQuestions.WriteString("# Open questions\n\n")
	for _, openQuestion := range c.OpenQuestions {
		openQuestions.WriteString("## ")
		openQuestions.WriteString(openQuestion.ID)
		openQuestions.WriteString(": ")
		openQuestions.WriteString(openQuestion.Title)
		openQuestions.WriteString("\n\nBlocker: ")
		openQuestions.WriteString(strings.ToLower(fmt.Sprint(openQuestion.Blocker)))
		openQuestions.WriteString("\n\nStatus: ")
		openQuestions.WriteString(openQuestion.Status)
		openQuestions.WriteString("\n\n")
	}
	writePlanDoc(t, dir, "10_open_questions.md", openQuestions.String())
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

func requireProof(t *testing.T, proof []ProofItem, ruleID string) ProofItem {
	t.Helper()

	for _, item := range proof {
		if item.RuleID == ruleID {
			return item
		}
	}
	t.Fatalf("expected proof %s, got %#v", ruleID, proof)
	return ProofItem{}
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
