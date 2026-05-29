package readiness

import (
	"sort"
	"strings"
	"unicode"
)

func NextQuestions(result Result) []NextQuestion {
	questions := make([]NextQuestion, 0, len(result.Issues))
	for _, issue := range result.Issues {
		if question := questionForIssue(issue); question.Question != "" {
			questions = append(questions, question)
		}
	}
	return questions
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
		question.Question = ref + " is missing. What planning content belongs in this required doc, or should the missing doc remain a blocker?"
	case "R003":
		question.Question = "No capability is accepted yet. What user-visible capability should be recorded first, or should capability definition be explicitly deferred?"
	case "R004":
		question.Question = ref + " has no evaluation. What evidence would prove this capability is complete: a test, review checklist, demo condition, user approval, or an explicit deferral?"
	case "R005":
		question.Question = ref + " has no evaluation method. What deterministic method should check the evidence, or should the method be deferred with a reason?"
	case "R006":
		question.Question = ref + " is high severity and has no mitigation. What mitigation would reduce or monitor it, who owns it, or should this become an explicit accepted-risk decision?"
	case "R007":
		question.Question = ref + " has no requirement or artifact trace. Which requirement or artifact anchors it, or should the trace be deferred with a reason?"
	case "R008":
		question.Question = ref + " has an invalid decision status. Should it be accepted, deferred, rejected, or marked not_applicable?"
	case "R009":
		question.Question = ref + " is blocking readiness. What answer would resolve it: an accepted decision, a deferral with reason, not_applicable, or keeping it blocking with the missing information named?"
	case "R010":
		question.Question = "No non-goal is recorded. What explicit non-goal should bound the plan, or why is this boundary not_applicable?"
	case "R011":
		question.Question = "The readiness profile cannot be trusted. What profile file correction is needed before status output should guide planning?"
	case "R012":
		question.Question = ref + " differs between docs and contract. Which source is correct, and should the repair update docs, update the contract, defer the record, or mark it not_applicable?"
	case "R013":
		question.Question = ref + " is part of a decision conflict. Which accepted decision should be revised, rejected, split, or marked not_applicable?"
	case "R014":
		question.Question = ref + " is missing a concrete purpose. What should change, for whom, and why does it matter?"
	case "R015":
		question.Question = ref + " is missing an actor or outcome. Which actor needs what outcome, and should that record be accepted, kept as evidence, deferred, or marked not_applicable?"
	case "R016":
		question.Question = ref + " is missing a delivery surface. Should the plan target a CLI, web app, conversation, document, workflow, research protocol, human service, another surface, or a deferral with reason?"
	case "D001":
		question.Question = ref + " is deferred. Should it remain deferred with a reason, become an accepted or rejected decision, or be marked not_applicable?"
	case "D002":
		question.Question = ref + " remains open. Should it be answered, deferred with a reason, marked not_applicable, or left open with the missing information named?"
	default:
		question.Question = "For " + ref + ", what planning decision is needed to address " + issue.RuleID + "?"
	}
	return question
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
