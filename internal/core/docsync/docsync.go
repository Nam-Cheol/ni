package docsync

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"ni/internal/core/contract"
)

const (
	CapabilityDoc = "docs/plan/02_capabilities.md"
	RiskDoc       = "docs/plan/06_risks_security.md"
	EvaluationDoc = "docs/plan/07_evaluation_contract.md"
	DecisionDoc   = "docs/plan/11_decision_log.md"
)

type Finding struct {
	Message string
}

type section struct {
	id     string
	path   string
	fields map[string]string
}

var sectionHeading = regexp.MustCompile(`^##\s+((CAP|EVAL|RISK|DEC)-[0-9]{3})(?::\s*(.*))?\s*$`)
var fieldLine = regexp.MustCompile(`^([A-Za-z][A-Za-z ]*):\s*(.*)$`)

func Check(root string, c contract.Contract) []Finding {
	root = filepath.Clean(root)
	var findings []Finding

	capabilitySections := parseDoc(root, CapabilityDoc, "CAP")
	riskSections := parseDoc(root, RiskDoc, "RISK")
	evaluationSections := parseDoc(root, EvaluationDoc, "EVAL")
	decisionSections := parseDoc(root, DecisionDoc, "DEC")

	findings = append(findings, checkCapabilities(c, capabilitySections)...)
	findings = append(findings, checkEvaluations(c, evaluationSections)...)
	findings = append(findings, checkRisks(c, riskSections)...)
	findings = append(findings, checkDecisions(c, decisionSections)...)

	sort.Slice(findings, func(i, j int) bool {
		return findings[i].Message < findings[j].Message
	})
	return findings
}

func checkCapabilities(c contract.Contract, sections map[string]section) []Finding {
	contractIDs := make(map[string]contract.Capability, len(c.Capabilities))
	for _, capability := range c.Capabilities {
		contractIDs[capability.ID] = capability
	}

	var findings []Finding
	for _, id := range sortedSectionIDs(sections) {
		if _, ok := contractIDs[id]; !ok {
			findings = append(findings, finding("%s mentions %s but .ni/contract.json has no capability with that id; add %s to contract or remove/rename the docs section", CapabilityDoc, id, id))
		}
	}
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		if _, ok := sections[capability.ID]; !ok {
			findings = append(findings, finding("accepted capability %s is in .ni/contract.json but missing from %s; add a `## %s: ...` section or remove/defer it in contract", capability.ID, CapabilityDoc, capability.ID))
		}
	}
	return findings
}

func checkEvaluations(c contract.Contract, sections map[string]section) []Finding {
	contractIDs := make(map[string]contract.Evaluation, len(c.Evaluations))
	for _, evaluation := range c.Evaluations {
		contractIDs[evaluation.ID] = evaluation
	}

	var findings []Finding
	for _, id := range sortedSectionIDs(sections) {
		if _, ok := contractIDs[id]; !ok {
			findings = append(findings, finding("%s mentions %s but .ni/contract.json has no evaluation with that id; add %s to contract or remove/rename the docs section", EvaluationDoc, id, id))
		}
	}
	for _, evaluation := range c.Evaluations {
		doc, ok := sections[evaluation.ID]
		if !ok {
			findings = append(findings, finding("evaluation %s is in .ni/contract.json but missing from %s; add a `## %s: ...` section or remove it in contract", evaluation.ID, EvaluationDoc, evaluation.ID))
			continue
		}
		if strings.TrimSpace(doc.fields["method"]) == "" {
			findings = append(findings, finding("evaluation %s lacks `Method: ...` in %s; add the deterministic evaluation method to the docs section", evaluation.ID, EvaluationDoc))
		}
	}
	return findings
}

func checkRisks(c contract.Contract, sections map[string]section) []Finding {
	contractIDs := make(map[string]contract.Risk, len(c.Risks))
	for _, risk := range c.Risks {
		contractIDs[risk.ID] = risk
	}

	var findings []Finding
	for _, id := range sortedSectionIDs(sections) {
		if _, ok := contractIDs[id]; !ok {
			findings = append(findings, finding("%s mentions %s but .ni/contract.json has no risk with that id; add %s to contract or remove/rename the docs section", RiskDoc, id, id))
		}
	}
	for _, risk := range c.Risks {
		if risk.Status != "accepted" {
			continue
		}
		doc, ok := sections[risk.ID]
		if !ok {
			findings = append(findings, finding("accepted risk %s is in .ni/contract.json but missing from %s; add a `## %s: ...` section with severity and mitigation", risk.ID, RiskDoc, risk.ID))
			continue
		}
		docSeverity := normalizeValue(doc.fields["severity"])
		if docSeverity == "" {
			findings = append(findings, finding("accepted risk %s lacks `Severity: ...` in %s; add the contract severity to the docs section", risk.ID, RiskDoc))
		} else if docSeverity != normalizeValue(risk.Severity) {
			findings = append(findings, finding("%s severity for %s is %q but .ni/contract.json says %q; update one source so they match", RiskDoc, risk.ID, doc.fields["severity"], risk.Severity))
		}
		if strings.TrimSpace(doc.fields["mitigation"]) == "" {
			findings = append(findings, finding("accepted risk %s lacks `Mitigation: ...` in %s; add the docs explanation for how the risk is handled", risk.ID, RiskDoc))
		}
	}
	return findings
}

func checkDecisions(c contract.Contract, sections map[string]section) []Finding {
	contractIDs := make(map[string]contract.Decision, len(c.Decisions))
	for _, decision := range c.Decisions {
		contractIDs[decision.ID] = decision
	}

	var findings []Finding
	for _, id := range sortedSectionIDs(sections) {
		if _, ok := contractIDs[id]; !ok {
			findings = append(findings, finding("%s mentions %s but .ni/contract.json has no decision with that id; add %s to contract or remove/rename the docs section", DecisionDoc, id, id))
		}
	}
	for _, decision := range c.Decisions {
		doc, ok := sections[decision.ID]
		if !ok {
			findings = append(findings, finding("decision %s is in .ni/contract.json but missing from %s; add a `## %s: ...` section or remove it in contract", decision.ID, DecisionDoc, decision.ID))
			continue
		}
		docStatus := normalizeValue(doc.fields["status"])
		if docStatus == "" {
			findings = append(findings, finding("decision %s lacks `Status: ...` in %s; add the contract decision status to the docs section", decision.ID, DecisionDoc))
		} else if docStatus != normalizeValue(decision.Status) {
			findings = append(findings, finding("%s status for %s is %q but .ni/contract.json says %q; update one source so they match", DecisionDoc, decision.ID, doc.fields["status"], decision.Status))
		}
	}
	return findings
}

func parseDoc(root string, relPath string, prefix string) map[string]section {
	data, err := os.ReadFile(filepath.Join(root, relPath))
	if err != nil {
		return map[string]section{}
	}

	sections := make(map[string]section)
	var current *section
	for _, rawLine := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(rawLine)
		if strings.HasPrefix(line, "## ") {
			match := sectionHeading.FindStringSubmatch(line)
			if match == nil || match[2] != prefix {
				current = nil
				continue
			}
			sec := section{id: match[1], path: relPath, fields: map[string]string{}}
			sections[sec.id] = sec
			current = &sec
			continue
		}
		if current == nil {
			continue
		}
		match := fieldLine.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		key := normalizeKey(match[1])
		if key == "" {
			continue
		}
		stored := sections[current.id]
		stored.fields[key] = strings.TrimSpace(match[2])
		sections[current.id] = stored
	}
	return sections
}

func sortedSectionIDs(sections map[string]section) []string {
	ids := make([]string, 0, len(sections))
	for id := range sections {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func normalizeKey(value string) string {
	return strings.ToLower(strings.Join(strings.Fields(value), " "))
}

func normalizeValue(value string) string {
	return strings.ToLower(strings.TrimSpace(value))
}

func finding(format string, args ...any) Finding {
	return Finding{Message: fmt.Sprintf(format, args...)}
}
