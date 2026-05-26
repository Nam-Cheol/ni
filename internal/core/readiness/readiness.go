package readiness

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/docsync"
	readinessprofile "ni/internal/core/profile"
)

type Status string

const (
	StatusBlocked            Status = "BLOCKED"
	StatusReadyWithDeferrals Status = "READY_WITH_DEFERRALS"
	StatusReady              Status = "READY"
)

type Result struct {
	Status           Status   `json:"status"`
	Profile          string   `json:"profile"`
	ProductType      string   `json:"product_type,omitempty"`
	DeliverySurfaces []string `json:"delivery_surfaces,omitempty"`
	InteractionMode  string   `json:"interaction_mode,omitempty"`
	Guidance         []string `json:"guidance,omitempty"`
	Issues           []Issue  `json:"issues"`
}

type Issue struct {
	RuleID   string `json:"rule_id"`
	Severity string `json:"severity"`
	Message  string `json:"message"`
}

type rulesFile struct {
	RequiredDocs []string `json:"required_docs"`
}

type profilesFile struct {
	Schema         string              `json:"schema"`
	DefaultProfile string              `json:"default_profile"`
	Profiles       []profileRulesEntry `json:"profiles"`
}

type profileRulesEntry struct {
	Name          string            `json:"name"`
	IssueSeverity map[string]string `json:"issue_severity"`
}

type profileRules struct {
	DefaultProfile string
	IssueSeverity  map[string]map[string]string
}

func Evaluate(dir string) Result {
	root := filepath.Clean(dir)
	issues := make([]Issue, 0)

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	activeProfile := readinessprofile.Default
	if err == nil {
		activeProfile = c.ReadinessProfile
	}

	rules, rulesErr := loadProfileRules(root)
	if rulesErr != nil {
		issues = append(issues, block("R011", rulesErr.Error()))
	}

	for _, path := range requiredDocs(root) {
		if err := requireFile(root, path); err != nil {
			issues = append(issues, issue(rules, activeProfile, "R001", err.Error()))
		}
	}

	if err != nil {
		issues = append(issues, block("R002", err.Error()))
		return resultFromIssues(activeProfile, issues)
	}

	issues = append(issues, evaluateContract(c, rules)...)
	for _, finding := range docsync.Check(root, c) {
		issues = append(issues, issue(rules, activeProfile, "R012", finding.Message))
	}
	result := resultFromIssues(activeProfile, issues)
	result.ProductType = c.ProductType
	result.DeliverySurfaces = c.DeliverySurfaces
	result.InteractionMode = c.InteractionMode
	result.Guidance = guidanceFor(c)
	return result
}

func evaluateContract(c contract.Contract, rules profileRules) []Issue {
	var issues []Issue

	evaluations := make(map[string]contract.Evaluation, len(c.Evaluations))
	for _, evaluation := range c.Evaluations {
		evaluations[evaluation.ID] = evaluation
		if strings.TrimSpace(evaluation.Method) == "" {
			issues = append(issues, issue(rules, c.ReadinessProfile, "R005", fmt.Sprintf("%s has no method", evaluation.ID)))
		}
	}

	if len(c.Capabilities) == 0 {
		issues = append(issues, issue(rules, c.ReadinessProfile, "R003", "at least one capability is required"))
	}
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		if !hasLinkedEvaluation(capability, evaluations) {
			issues = append(issues, issue(rules, c.ReadinessProfile, "R004", fmt.Sprintf("%s has no linked evaluation", capability.ID)))
		}
		if len(capability.Artifacts) == 0 && len(capability.Requirements) == 0 {
			issues = append(issues, issue(rules, c.ReadinessProfile, "R007", fmt.Sprintf("%s has no artifact or requirement", capability.ID)))
		}
	}

	for _, risk := range c.Risks {
		if strings.EqualFold(risk.Severity, "high") && strings.TrimSpace(risk.Mitigation) == "" {
			issues = append(issues, issue(rules, c.ReadinessProfile, "R006", fmt.Sprintf("%s is high severity and has no mitigation", risk.ID)))
		}
	}

	for _, decision := range c.Decisions {
		if !validDecisionStatus(decision.Status) {
			issues = append(issues, issue(rules, c.ReadinessProfile, "R008", fmt.Sprintf("%s has invalid status %q", decision.ID, decision.Status)))
		}
		if decision.Status == "deferred" {
			issues = append(issues, issue(rules, c.ReadinessProfile, "D001", fmt.Sprintf("%s is deferred", decision.ID)))
		}
	}

	for _, openQuestion := range c.OpenQuestions {
		if isClosed(openQuestion.Status) {
			continue
		}
		if openQuestion.Blocker {
			issues = append(issues, issue(rules, c.ReadinessProfile, "R009", fmt.Sprintf("%s is a blocker open question", openQuestion.ID)))
		} else {
			issues = append(issues, issue(rules, c.ReadinessProfile, "D002", fmt.Sprintf("%s remains open", openQuestion.ID)))
		}
	}

	if len(c.NonGoals) == 0 {
		issues = append(issues, issue(rules, c.ReadinessProfile, "R010", "at least one non-goal is required"))
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

func resultFromIssues(activeProfile string, issues []Issue) Result {
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
	return Result{Status: status, Profile: activeProfile, Issues: issues}
}

func guidanceFor(c contract.Contract) []string {
	guidance := []string{
		fmt.Sprintf("product_type=%s changes planning guidance only; readiness authority remains the shared gate.", c.ProductType),
	}
	switch c.ProductType {
	case "conversation_product":
		guidance = append(guidance, "Cover conversation turns, failure handling, transcript evaluation, and human handoff.")
	case "research_protocol":
		guidance = append(guidance, "Cover hypothesis, data handling, method, analysis, ethics, and reproducibility evidence.")
	case "operations_process":
		guidance = append(guidance, "Cover roles, handoffs, triggers, exceptions, service levels, and operating evidence.")
	case "education_program":
		guidance = append(guidance, "Cover learners, outcomes, curriculum flow, assessment, and facilitator materials.")
	case "document_product":
		guidance = append(guidance, "Cover readers, document structure, review criteria, publishing format, and maintenance.")
	case "physical_product":
		guidance = append(guidance, "Cover materials, safety, production or service process, logistics, and physical validation.")
	case "mixed":
		guidance = append(guidance, "Cover each delivery surface explicitly and keep cross-surface acceptance criteria traceable.")
	default:
		guidance = append(guidance, "Cover interfaces, runtime boundaries, validation evidence, and operational constraints.")
	}
	for _, surface := range c.DeliverySurfaces {
		switch surface {
		case "conversation":
			guidance = append(guidance, "conversation surface: specify turn boundaries, memory expectations, refusals, and escalation.")
		case "document":
			guidance = append(guidance, "document surface: specify audience, structure, review workflow, and publication format.")
		case "workflow":
			guidance = append(guidance, "workflow surface: specify roles, state transitions, handoffs, and exception handling.")
		case "human_service":
			guidance = append(guidance, "human_service surface: specify service roles, scripts or playbooks, and quality checks.")
		case "physical":
			guidance = append(guidance, "physical surface: specify physical constraints, safety checks, and validation evidence.")
		}
	}
	return guidance
}

func block(ruleID string, message string) Issue {
	return Issue{RuleID: ruleID, Severity: "blocker", Message: message}
}

func deferIssue(ruleID string, message string) Issue {
	return Issue{RuleID: ruleID, Severity: "deferral", Message: message}
}

func issue(rules profileRules, activeProfile string, ruleID string, message string) Issue {
	severity := rules.severity(activeProfile, ruleID)
	if severity == "blocker" {
		return block(ruleID, message)
	}
	return deferIssue(ruleID, message)
}

func (rules profileRules) severity(activeProfile string, ruleID string) string {
	if profileRules, ok := rules.IssueSeverity[activeProfile]; ok {
		if severity, ok := profileRules[ruleID]; ok && severity != "" {
			return severity
		}
	}
	if profileRules, ok := rules.IssueSeverity[readinessprofile.Default]; ok {
		if severity, ok := profileRules[ruleID]; ok && severity != "" {
			return severity
		}
	}
	return "blocker"
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

func loadProfileRules(root string) (profileRules, error) {
	rules := defaultProfileRules()
	data, err := os.ReadFile(filepath.Join(root, ".ni", "readiness.profiles.json"))
	if err != nil {
		if os.IsNotExist(err) {
			return rules, nil
		}
		return rules, fmt.Errorf("cannot read readiness profiles: %w", err)
	}

	var file profilesFile
	if err := json.Unmarshal(data, &file); err != nil {
		return rules, fmt.Errorf("malformed readiness profiles JSON: %w", err)
	}
	if file.Schema != "ni.readiness.profiles.v0" {
		return rules, fmt.Errorf("unsupported readiness profiles schema %q", file.Schema)
	}
	if err := readinessprofile.Validate(file.DefaultProfile); err != nil {
		return rules, fmt.Errorf("invalid default readiness profile: %w", err)
	}
	if len(file.Profiles) == 0 {
		return rules, fmt.Errorf("readiness profiles must define at least one profile")
	}

	loaded := profileRules{
		DefaultProfile: file.DefaultProfile,
		IssueSeverity:  make(map[string]map[string]string, len(file.Profiles)),
	}
	for _, entry := range file.Profiles {
		if err := readinessprofile.Validate(entry.Name); err != nil {
			return rules, err
		}
		if len(entry.IssueSeverity) == 0 {
			return rules, fmt.Errorf("readiness profile %q has no issue severity map", entry.Name)
		}
		loaded.IssueSeverity[entry.Name] = make(map[string]string, len(entry.IssueSeverity))
		for ruleID, severity := range entry.IssueSeverity {
			if severity != "blocker" && severity != "deferral" {
				return rules, fmt.Errorf("readiness profile %q rule %s has invalid severity %q", entry.Name, ruleID, severity)
			}
			loaded.IssueSeverity[entry.Name][ruleID] = severity
		}
	}
	for _, name := range readinessprofile.Names() {
		if _, ok := loaded.IssueSeverity[name]; !ok {
			return rules, fmt.Errorf("readiness profiles missing %q definition", name)
		}
	}
	return loaded, nil
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

func defaultProfileRules() profileRules {
	return profileRules{
		DefaultProfile: readinessprofile.Default,
		IssueSeverity: map[string]map[string]string{
			"concept": {
				"R001": "blocker",
				"R002": "blocker",
				"R003": "deferral",
				"R004": "deferral",
				"R005": "deferral",
				"R006": "blocker",
				"R007": "deferral",
				"R008": "blocker",
				"R009": "blocker",
				"R010": "deferral",
				"R012": "blocker",
				"D001": "deferral",
				"D002": "deferral",
			},
			"prototype": {
				"R001": "blocker",
				"R002": "blocker",
				"R003": "blocker",
				"R004": "blocker",
				"R005": "blocker",
				"R006": "blocker",
				"R007": "blocker",
				"R008": "blocker",
				"R009": "blocker",
				"R010": "blocker",
				"R012": "blocker",
				"D001": "deferral",
				"D002": "deferral",
			},
			"mvp": {
				"R001": "blocker",
				"R002": "blocker",
				"R003": "blocker",
				"R004": "blocker",
				"R005": "blocker",
				"R006": "blocker",
				"R007": "blocker",
				"R008": "blocker",
				"R009": "blocker",
				"R010": "blocker",
				"R012": "blocker",
				"D001": "deferral",
				"D002": "deferral",
			},
			"beta": {
				"R001": "blocker",
				"R002": "blocker",
				"R003": "blocker",
				"R004": "blocker",
				"R005": "blocker",
				"R006": "blocker",
				"R007": "blocker",
				"R008": "blocker",
				"R009": "blocker",
				"R010": "blocker",
				"R012": "blocker",
				"D001": "deferral",
				"D002": "deferral",
			},
			"production": {
				"R001": "blocker",
				"R002": "blocker",
				"R003": "blocker",
				"R004": "blocker",
				"R005": "blocker",
				"R006": "blocker",
				"R007": "blocker",
				"R008": "blocker",
				"R009": "blocker",
				"R010": "blocker",
				"R012": "blocker",
				"D001": "blocker",
				"D002": "blocker",
			},
		},
	}
}
