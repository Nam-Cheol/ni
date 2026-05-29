package readiness

import (
	"sort"
	"strings"
	"unicode"

	"ni/internal/core/docsync"
)

const maxPrimaryQuestions = 3

type questionCandidate struct {
	question NextQuestion
	priority int
	index    int
	key      string
}

func NextQuestions(result Result) []NextQuestion {
	candidates := prioritizedQuestionCandidates(result)
	questions := make([]NextQuestion, 0, maxPrimaryQuestions)
	for _, candidate := range candidates {
		if len(questions) >= maxPrimaryQuestions {
			break
		}
		questions = append(questions, candidate.question)
	}
	return questions
}

func OmittedNextQuestionCount(result Result) int {
	candidates := prioritizedQuestionCandidates(result)
	if len(candidates) <= maxPrimaryQuestions {
		return 0
	}
	return len(candidates) - maxPrimaryQuestions
}

func prioritizedQuestionCandidates(result Result) []questionCandidate {
	raw := make([]questionCandidate, 0, len(result.Issues))
	hasFirstRunSync := false
	firstRunRules := map[string]bool{}
	for idx, issue := range result.Issues {
		if isFirstRunRule(issue.RuleID) {
			firstRunRules[issue.RuleID] = true
		}
		if isFirstRunSyncIssue(issue) {
			hasFirstRunSync = true
		}
		if result.Status == StatusBlocked && issue.Severity != "blocker" {
			continue
		}
		if question := questionForIssue(issue); question.Question != "" {
			raw = append(raw, questionCandidate{
				question: question,
				priority: questionPriority(issue),
				index:    idx,
				key:      questionKey(question),
			})
		}
	}

	if !hasFirstRunSync && firstRunRules["R014"] && firstRunRules["R015"] && firstRunRules["R016"] {
		return firstRunCardCandidates(result.Issues)
	}

	sort.SliceStable(raw, func(i, j int) bool {
		if raw[i].priority != raw[j].priority {
			return raw[i].priority < raw[j].priority
		}
		return raw[i].index < raw[j].index
	})

	var deduped []questionCandidate
	seen := map[string]bool{}
	for _, candidate := range raw {
		if candidate.key == "" {
			candidate.key = questionKey(candidate.question)
		}
		if seen[candidate.key] {
			continue
		}
		if hasFirstRunSync && isFirstRunRule(candidate.question.RuleID) {
			continue
		}
		seen[candidate.key] = true
		deduped = append(deduped, candidate)
	}
	return deduped
}

func firstRunCardCandidates(issues []Issue) []questionCandidate {
	byRule := map[string]Issue{}
	for _, issue := range issues {
		if isFirstRunRule(issue.RuleID) {
			byRule[issue.RuleID] = issue
		}
	}
	order := []string{"R014", "R015", "R016"}
	candidates := make([]questionCandidate, 0, len(order))
	for idx, ruleID := range order {
		issue, ok := byRule[ruleID]
		if !ok {
			continue
		}
		question := questionForIssue(issue)
		question.Group = "First-run card"
		candidates = append(candidates, questionCandidate{
			question: question,
			priority: 1,
			index:    idx,
			key:      questionKey(question),
		})
	}
	return candidates
}

func questionForIssue(issue Issue) NextQuestion {
	references := issueReferences(issue)
	if issue.SyncDiagnostic != nil && issue.SyncDiagnostic.ID != "" {
		references = prioritizeReference(references, issue.SyncDiagnostic.ID)
	}
	question := NextQuestion{
		RuleID:     issue.RuleID,
		Severity:   issue.Severity,
		References: references,
	}
	ref := issue.RuleID
	if len(references) > 0 {
		ref = references[0]
	}

	switch issue.RuleID {
	case "R001":
		question.Group = "Plan source repairs"
		question.Question = ref + " is missing. What planning content belongs in this required doc, or should the missing doc remain a blocker?"
		question.AnswerShape = "restore doc / recreate from template / keep blocker with reason"
	case "R003":
		question.Group = "Capability definition"
		question.Question = "No capability is accepted yet. What user-visible capability should be recorded first, or should capability definition be explicitly deferred?"
		question.AnswerShape = "capability title plus accepted, draft, or deferred status"
	case "R004":
		question.Group = "Evaluation evidence"
		question.Question = ref + " has no evaluation. What evidence would prove this capability is complete?"
		question.AnswerShape = "test, review checklist, demo condition, user approval, protocol check, or manual inspection"
	case "R005":
		question.Group = "Evaluation evidence"
		question.Question = ref + " has no evaluation method. What deterministic method should check the evidence, or should the method be deferred with a reason?"
		question.AnswerShape = "test, review checklist, demo condition, protocol check, manual inspection, or deferral with reason"
	case "R006":
		question.Group = "Risk decisions"
		question.Question = ref + " is high severity and has no mitigation. What mitigation would reduce or monitor it, who owns it, or should this become an explicit accepted-risk decision?"
		question.AnswerShape = "mitigation plus owner, monitoring plan, accepted-risk decision, or explicit deferral"
	case "R007":
		question.Group = "Traceability"
		question.Question = ref + " has no requirement or artifact trace. Which requirement or artifact anchors it, or should the trace be deferred with a reason?"
		question.AnswerShape = "requirement ID, artifact ID, or deferral with reason"
	case "R008":
		question.Group = "Decision repairs"
		question.Question = ref + " has an invalid decision status. Should it be accepted, deferred, rejected, or marked not_applicable?"
		question.AnswerShape = "accepted / deferred with reason / rejected / not_applicable"
	case "R009":
		question.Group = "Open blockers"
		question.Question = ref + " is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?"
		question.AnswerShape = "accepted decision, deferral with reason, not_applicable, or keep blocking with reason"
	case "R010":
		question.Group = "Scope boundaries"
		question.Question = "What must this project explicitly avoid so downstream work does not drift in scope?"
		question.AnswerShape = "one or more non-goals, or not_applicable with reason"
	case "R011":
		question.Group = "Readiness profile repairs"
		question.Question = "The readiness profile cannot be trusted. What profile file correction is needed before status output should guide planning?"
		question.AnswerShape = "profile file repair or keep blocker with reason"
	case "R012":
		return syncRepairQuestion(issue, references, ref)
	case "R013":
		question.Group = "Decision repairs"
		question.Question = ref + " is part of a decision conflict. Which accepted decision should be revised, rejected, split, or marked not_applicable?"
		question.AnswerShape = "revise one decision / reject one decision / split decisions / not_applicable"
	case "R014":
		question.Group = "First-run card"
		question.Question = "What should this project change, for whom, and why does it matter?"
		question.AnswerShape = "one or two sentences describing the desired reality change"
	case "R015":
		question.Group = "First-run card"
		question.Question = "Who are the primary actors, and what outcome should each one get?"
		question.AnswerShape = "actor -> expected outcome"
	case "R016":
		question.Group = "First-run card"
		question.Question = "What is the likely delivery surface?"
		question.AnswerShape = "CLI, web app, conversation, document, workflow, research protocol, human service, or deferred with reason"
	case "D001":
		if issue.Severity != "deferral" {
			return question
		}
		question.Group = "Handoff deferrals"
		question.Question = ref + " is deferred. Does this deferred decision affect the next handoff, or should it remain visible without blocking?"
		question.AnswerShape = "affects handoff and must resolve / remains deferred with reason"
	case "D002":
		if issue.Severity != "deferral" {
			return question
		}
		question.Group = "Handoff deferrals"
		question.Question = ref + " remains open. Does this open question affect the next handoff, or should it remain visible without blocking?"
		question.AnswerShape = "resolve now / defer with reason / leave open and non-blocking"
	default:
		question.Group = "Planning repairs"
		question.Question = "For " + ref + ", what planning decision is needed to address " + issue.RuleID + "?"
		question.AnswerShape = "decision, evidence, deferral, not_applicable, or keep blocker with reason"
	}
	return question
}

func syncRepairQuestion(issue Issue, references []string, ref string) NextQuestion {
	question := NextQuestion{
		RuleID:      issue.RuleID,
		Severity:    issue.Severity,
		Group:       "Sync repairs",
		References:  references,
		Question:    ref + " differs between docs and contract. Which source is correct, and should the repair update docs, update the contract, defer the record, or mark it not_applicable?",
		AnswerShape: "update docs / update contract / revise both / defer with reason / keep blocker with reason",
	}
	if issue.SyncDiagnostic == nil {
		return question
	}
	diag := issue.SyncDiagnostic
	question.Location = diag.Location
	question.AnswerShape = "update contract / revise docs / revise both / keep blocker with reason"
	switch diag.ID {
	case "SYNC-014":
		question.Question = firstRunSyncQuestion(diag, "project purpose")
	case "SYNC-015":
		question.Question = firstRunSyncQuestion(diag, "actors and outcomes")
	case "SYNC-016":
		question.Question = firstRunSyncQuestion(diag, "delivery surface")
	default:
		question.Question = diag.Problem + " Which source is correct, and should the repair update docs, update the contract, defer the record, or keep this blocking?"
	}
	return question
}

func firstRunSyncQuestion(diag *docsync.Diagnostic, field string) string {
	problem := strings.TrimSuffix(diag.Problem, ".")
	if strings.Contains(diag.Problem, "documented but missing") {
		return problem + ". Should .ni/contract.json be updated to match the docs, or is the docs text only a draft?"
	}
	if strings.Contains(diag.Problem, "recorded in .ni/contract.json but not explained in docs") {
		return problem + ". Should docs/plan be updated to explain the contract value, or is the contract value no longer accepted?"
	}
	if strings.Contains(diag.Problem, "differs between docs and contract") {
		return problem + ". Which " + field + " is correct, and should the stale side be updated or kept blocking?"
	}
	return problem + ". Should the stale side be updated, revised, or kept blocking with a reason?"
}

func questionPriority(issue Issue) int {
	if isFirstRunSyncIssue(issue) {
		return 10
	}
	switch issue.RuleID {
	case "R012":
		return 20
	case "R006":
		return 30
	case "R004", "R005":
		return 40
	case "R010":
		return 50
	case "R009":
		return 60
	case "R014", "R015", "R016":
		return 70
	case "R003", "R007":
		return 80
	case "R008", "R013":
		return 90
	case "D001", "D002":
		return 120
	default:
		return 100
	}
}

func isFirstRunRule(ruleID string) bool {
	return ruleID == "R014" || ruleID == "R015" || ruleID == "R016"
}

func isFirstRunSyncIssue(issue Issue) bool {
	if issue.RuleID != "R012" || issue.SyncDiagnostic == nil {
		return false
	}
	switch issue.SyncDiagnostic.ID {
	case "SYNC-014", "SYNC-015", "SYNC-016":
		return true
	default:
		return false
	}
}

func questionKey(question NextQuestion) string {
	if len(question.References) > 0 {
		return question.RuleID + ":" + question.References[0]
	}
	return question.RuleID + ":" + question.Question
}

func issueReferences(issue Issue) []string {
	seen := map[string]struct{}{}
	var refs []string
	for _, token := range strings.Fields(issue.Message) {
		ref := trimReference(token)
		if ref == "" || !looksLikeReference(ref) {
			continue
		}
		if _, ok := seen[ref]; ok {
			continue
		}
		seen[ref] = struct{}{}
		refs = append(refs, ref)
	}
	sort.Strings(refs)
	return refs
}

func prioritizeReference(references []string, preferred string) []string {
	ordered := []string{preferred}
	for _, ref := range references {
		if ref != preferred {
			ordered = append(ordered, ref)
		}
	}
	return ordered
}

func trimReference(token string) string {
	return strings.TrimFunc(token, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/' || r == '.')
	})
}

func looksLikeReference(value string) bool {
	if strings.Contains(value, "/") || strings.Contains(value, ".md") || strings.Contains(value, ".json") || strings.HasPrefix(value, "project.") {
		return true
	}
	parts := strings.Split(value, "-")
	if len(parts) < 2 {
		return false
	}
	hasLetter := false
	hasDigit := false
	for _, r := range value {
		if unicode.IsLetter(r) {
			hasLetter = true
		}
		if unicode.IsDigit(r) {
			hasDigit = true
		}
	}
	return hasLetter && hasDigit
}
