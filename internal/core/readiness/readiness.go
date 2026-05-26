package readiness

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ni/internal/core/contract"
)

type Status string

const (
	StatusBlocked            Status = "BLOCKED"
	StatusReadyWithDeferrals Status = "READY_WITH_DEFERRALS"
	StatusReady              Status = "READY"
)

type Result struct {
	Status Status  `json:"status"`
	Issues []Issue `json:"issues"`
}

type Issue struct {
	RuleID   string `json:"rule_id"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type rulesFile struct {
	RequiredDocs []string `json:"required_docs"`
}

func Evaluate(dir string) Result {
	root := filepath.Clean(dir)
	issues := make([]Issue, 0)

	for _, path := range requiredDocs(root) {
		if err := requireFile(root, path); err != nil {
			issues = append(issues, block("R001", err.Error()))
		}
	}

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	if err != nil {
		issues = append(issues, block("R002", err.Error()))
		return resultFromIssues(issues)
	}

	issues = append(issues, evaluateContract(c)...)
	return resultFromIssues(issues)
}

func evaluateContract(c contract.Contract) []Issue {
	var issues []Issue

	evaluations := make(map[string]contract.Evaluation, len(c.Evaluations))
	for _, evaluation := range c.Evaluations {
		evaluations[evaluation.ID] = evaluation
		if strings.TrimSpace(evaluation.Method) == "" {
			issues = append(issues, block("R005", fmt.Sprintf("%s has no method", evaluation.ID)))
		}
	}

	if len(c.Capabilities) == 0 {
		issues = append(issues, block("R003", "at least one capability is required"))
	}
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		if !hasLinkedEvaluation(capability, evaluations) {
			issues = append(issues, block("R004", fmt.Sprintf("%s has no linked evaluation", capability.ID)))
		}
		if len(capability.Artifacts) == 0 && len(capability.Requirements) == 0 {
			issues = append(issues, block("R007", fmt.Sprintf("%s has no artifact or requirement", capability.ID)))
		}
	}

	for _, risk := range c.Risks {
		if strings.EqualFold(risk.Severity, "high") && strings.TrimSpace(risk.Mitigation) == "" {
			issues = append(issues, block("R006", fmt.Sprintf("%s is high severity and has no mitigation", risk.ID)))
		}
	}

	for _, decision := range c.Decisions {
		if !validDecisionStatus(decision.Status) {
			issues = append(issues, block("R008", fmt.Sprintf("%s has invalid status %q", decision.ID, decision.Status)))
		}
		if decision.Status == "deferred" {
			issues = append(issues, deferIssue("D001", fmt.Sprintf("%s is deferred", decision.ID)))
		}
	}

	for _, openQuestion := range c.OpenQuestions {
		if isClosed(openQuestion.Status) {
			continue
		}
		if openQuestion.Blocker {
			issues = append(issues, block("R009", fmt.Sprintf("%s is a blocker open question", openQuestion.ID)))
		} else {
			issues = append(issues, deferIssue("D002", fmt.Sprintf("%s remains open", openQuestion.ID)))
		}
	}

	if len(c.NonGoals) == 0 {
		issues = append(issues, block("R010", "at least one non-goal is required"))
	}

	return issues
}

func hasLinkedEvaluation(capability contract.Capability, evaluations map[string]contract.Evaluation) bool {
	for _, id := range capability.Evaluations {
		if _, ok := evaluations[id]; ok {
			return true
		}
	}
	return false
}

func validDecisionStatus(status string) bool {
	switch status {
	case "accepted", "deferred", "rejected", "not_applicable":
		return true
	default:
		return false
	}
}

func isClosed(status string) bool {
	switch status {
	case "closed", "resolved":
		return true
	default:
		return false
	}
}

func resultFromIssues(issues []Issue) Result {
	status := StatusReady
	for _, issue := range issues {
		if issue.Severity == "blocker" {
			status = StatusBlocked
			break
		}
		if issue.Severity == "deferral" {
			status = StatusReadyWithDeferrals
		}
	}
	return Result{Status: status, Issues: issues}
}

func block(ruleID string, message string) Issue {
	return Issue{RuleID: ruleID, Severity: "blocker", Message: message}
}

func deferIssue(ruleID string, message string) Issue {
	return Issue{RuleID: ruleID, Severity: "deferral", Message: message}
}

func requireFile(root string, path string) error {
	info, err := os.Stat(filepath.Join(root, path))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("missing required planning doc %s", path)
		}
		return fmt.Errorf("cannot read required planning doc %s: %w", path, err)
	}
	if info.IsDir() {
		return fmt.Errorf("required planning doc %s is a directory", path)
	}
	return nil
}

func RequiredDocs(root string) []string {
	return requiredDocs(filepath.Clean(root))
}

func requiredDocs(root string) []string {
	data, err := os.ReadFile(filepath.Join(root, ".ni", "readiness.rules.json"))
	if err == nil {
		var rules rulesFile
		if json.Unmarshal(data, &rules) == nil && len(rules.RequiredDocs) > 0 {
			return rules.RequiredDocs
		}
	}
	return defaultRequiredDocs()
}

func defaultRequiredDocs() []string {
	return []string{
		"docs/plan/00_project_brief.md",
		"docs/plan/01_actors_outcomes.md",
		"docs/plan/02_capabilities.md",
		"docs/plan/03_interaction_contract.md",
		"docs/plan/04_domain_state.md",
		"docs/plan/05_constraints.md",
		"docs/plan/06_risks_security.md",
		"docs/plan/07_evaluation_contract.md",
		"docs/plan/08_delivery_operation.md",
		"docs/plan/09_execution_strategy.md",
		"docs/plan/10_open_questions.md",
		"docs/plan/11_decision_log.md",
	}
}
