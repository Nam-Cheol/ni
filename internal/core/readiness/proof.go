package readiness

import (
	"strings"
)

func Proof(result Result) []ProofItem {
	if len(result.Issues) == 0 {
		return []ProofItem{{
			RuleID:   "READY",
			Severity: "ready",
			Message:  "All readiness, sync, and conflict rules passed for profile " + result.Profile + ".",
		}}
	}

	items := make([]ProofItem, 0, len(result.Issues))
	for _, issue := range result.Issues {
		items = append(items, ProofItem{
			RuleID:         issue.RuleID,
			Severity:       issue.Severity,
			References:     issueReferences(issue),
			Message:        proofMessage(issue),
			SyncDiagnostic: issue.SyncDiagnostic,
		})
	}
	return items
}

func proofMessage(issue Issue) string {
	if issue.SyncDiagnostic != nil {
		return ensureSentence(issue.SyncDiagnostic.Problem)
	}
	references := issueReferences(issue)
	ref := issue.RuleID
	if len(references) > 0 {
		ref = references[0]
	}

	switch issue.RuleID {
	case "R004":
		return ref + " is accepted but has no linked EVAL."
	case "R006":
		return ref + " is high severity but has no mitigation."
	case "R009":
		return ref + " is marked as blocker."
	case "R013":
		return ensureSentence(issue.Message)
	case "R014":
		return "R014 Project purpose is missing."
	case "R015":
		return "R015 Actors or outcomes are missing."
	case "R016":
		return "R016 Delivery surface is missing."
	default:
		return ensureSentence(issue.Message)
	}
}

func ensureSentence(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if strings.HasSuffix(value, ".") || strings.HasSuffix(value, "!") || strings.HasSuffix(value, "?") {
		return value
	}
	return value + "."
}
