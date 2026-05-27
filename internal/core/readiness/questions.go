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
		question.Question = "For " + ref + ", what planning content is needed for the missing required doc, or should this remain a blocker?"
	case "R003":
		question.Question = "For R003, what capability should this plan accept first, or should capability definition be deferred?"
	case "R004":
		question.Question = "For " + ref + ", what evidence proves this capability works, or should that evidence be deferred?"
	case "R005":
		question.Question = "For " + ref + ", what deterministic method should evaluate this evidence, or should the evaluation be deferred?"
	case "R006":
		question.Question = "For " + ref + ", what mitigation, owner, or explicit accepted-risk decision is required?"
	case "R007":
		question.Question = "For " + ref + ", what requirement or artifact should this capability trace to, or should that trace be deferred?"
	case "R008":
		question.Question = "For " + ref + ", should the decision be accepted, deferred, rejected, or not_applicable?"
	case "R009":
		question.Question = "For " + ref + ", what decision resolves this blocker, should it be deferred, or why must it remain blocking?"
	case "R010":
		question.Question = "For R010, what must this project explicitly avoid?"
	case "R011":
		question.Question = "For R011, what readiness profile correction is needed before the gate can be trusted?"
	case "R012":
		question.Question = "For " + ref + ", which source should be corrected so docs and contract agree?"
	case "R013":
		question.Question = "For " + ref + ", which accepted decision should be revised, rejected, or split to remove the conflict?"
	case "D001":
		question.Question = "For " + ref + ", should this deferred decision remain deferred, become accepted or rejected, or be not_applicable?"
	case "D002":
		question.Question = "For " + ref + ", should this open question be resolved, deferred, or left open with a reason?"
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

func trimReference(token string) string {
	return strings.TrimFunc(token, func(r rune) bool {
		return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_' || r == '/' || r == '.')
	})
}

func looksLikeReference(value string) bool {
	if strings.Contains(value, "/") || strings.Contains(value, ".md") || strings.Contains(value, ".json") {
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
