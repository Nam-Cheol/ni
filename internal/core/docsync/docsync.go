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
	CapabilityDoc   = "docs/plan/02_capabilities.md"
	RiskDoc         = "docs/plan/06_risks_security.md"
	EvaluationDoc   = "docs/plan/07_evaluation_contract.md"
	OpenQuestionDoc = "docs/plan/10_open_questions.md"
	DecisionDoc     = "docs/plan/11_decision_log.md"
)

type Finding struct {
	Message    string
	Diagnostic Diagnostic
}

type Diagnostic struct {
	ID              string `json:"id"`
	Location        string `json:"location"`
	Problem         string `json:"problem"`
	WhyItMatters    string `json:"why_it_matters"`
	SuggestedRepair string `json:"suggested_repair"`
	BlocksEnd       bool   `json:"blocks_ni_end"`
}

type section struct {
	id          string
	path        string
	line        int
	title       string
	fields      map[string]string
	explanation bool
}

type mention struct {
	id   string
	path string
	line int
}

var sectionHeading = regexp.MustCompile(`^##\s+((CAP|EVAL|RISK|DEC|OQ)-[0-9]{3})(?::\s*(.*))?\s*$`)
var fieldLine = regexp.MustCompile(`^([A-Za-z][A-Za-z ]*):\s*(.*)$`)
var planID = regexp.MustCompile(`\b(CAP|REQ|EVAL|RISK|ART|NG|DEC|OQ)-[0-9]{3}\b`)

func Check(root string, c contract.Contract) []Finding {
	root = filepath.Clean(root)
	var findings []Finding

	capabilitySections := parseDoc(root, CapabilityDoc, "CAP")
	riskSections := parseDoc(root, RiskDoc, "RISK")
	evaluationSections := parseDoc(root, EvaluationDoc, "EVAL")
	openQuestionSections := parseDoc(root, OpenQuestionDoc, "OQ")
	decisionSections := parseDoc(root, DecisionDoc, "DEC")

	findings = append(findings, checkUnknownDocIDs(root, c)...)
	findings = append(findings, checkCapabilities(c, capabilitySections)...)
	findings = append(findings, checkEvaluations(c, evaluationSections)...)
	findings = append(findings, checkRisks(c, riskSections)...)
	findings = append(findings, checkOpenQuestions(c, openQuestionSections)...)
	findings = append(findings, checkDecisions(c, decisionSections)...)

	sort.Slice(findings, func(i, j int) bool {
		if findings[i].Diagnostic.ID != findings[j].Diagnostic.ID {
			return findings[i].Diagnostic.ID < findings[j].Diagnostic.ID
		}
		if findings[i].Diagnostic.Location != findings[j].Diagnostic.Location {
			return findings[i].Diagnostic.Location < findings[j].Diagnostic.Location
		}
		return findings[i].Diagnostic.Problem < findings[j].Diagnostic.Problem
	})
	return findings
}

func checkUnknownDocIDs(root string, c contract.Contract) []Finding {
	known := contractIDs(c)
	var findings []Finding
	for _, mention := range docMentions(root) {
		if known[mention.id] {
			continue
		}
		findings = append(findings, diagnosticFinding(Diagnostic{
			ID:              mention.id,
			Location:        location(mention.path, mention.line),
			Problem:         fmt.Sprintf("%s is mentioned in docs/plan but is missing from .ni/contract.json.", mention.id),
			WhyItMatters:    "Docs and contract are both lock sources; an unexplained docs-only ID makes the accepted plan ambiguous.",
			SuggestedRepair: fmt.Sprintf("Add %s to the matching contract collection, or remove or rename the docs reference if it is not accepted planning state.", mention.id),
			BlocksEnd:       true,
		}))
	}
	return findings
}

func checkCapabilities(c contract.Contract, sections map[string]section) []Finding {
	contractIDs := make(map[string]contract.Capability, len(c.Capabilities))
	for _, capability := range c.Capabilities {
		contractIDs[capability.ID] = capability
	}

	var findings []Finding
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		doc, ok := sections[capability.ID]
		if !ok {
			findings = append(findings, missingDocsFinding(capability.ID, contractLocation("capabilities", capability.ID), fmt.Sprintf("accepted capability %s is not explained in %s.", capability.ID, CapabilityDoc), "Accepted capabilities define downstream behavior; without docs, reviewers cannot see the human-readable intent behind the contract record.", fmt.Sprintf("Add a `## %s: ...` section to %s, or defer/remove the capability in .ni/contract.json.", capability.ID, CapabilityDoc)))
			continue
		}
		if !doc.explanation {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              capability.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("accepted capability %s has a docs heading but no explanatory body.", capability.ID),
				WhyItMatters:    "Accepted capabilities need reviewable narrative, not just matching IDs.",
				SuggestedRepair: "Add a short deterministic explanation of the behavior, boundary, or outcome under the capability heading.",
				BlocksEnd:       true,
			}))
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
	capabilityIDs := contractCapabilityIDs(c)
	for _, evaluation := range c.Evaluations {
		doc, ok := sections[evaluation.ID]
		if !ok {
			findings = append(findings, missingDocsFinding(evaluation.ID, contractLocation("evaluations", evaluation.ID), fmt.Sprintf("evaluation %s is not explained in %s.", evaluation.ID, EvaluationDoc), "Evaluations are the evidence plan; missing docs hide how accepted capabilities will be checked.", fmt.Sprintf("Add a `## %s: ...` section with `Method: ...`, or remove the evaluation from .ni/contract.json.", evaluation.ID)))
			continue
		}
		if strings.TrimSpace(doc.fields["method"]) == "" {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              evaluation.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("evaluation %s lacks `Method: ...` in docs.", evaluation.ID),
				WhyItMatters:    "Evidence is not trustworthy unless the evaluation method is explicit.",
				SuggestedRepair: "Add a deterministic `Method: ...` line to the evaluation docs section.",
				BlocksEnd:       true,
			}))
		}
		for _, capID := range referencedCapabilityIDs(doc) {
			if !capabilityIDs[capID] {
				findings = append(findings, diagnosticFinding(Diagnostic{
					ID:              evaluation.ID,
					Location:        location(doc.path, doc.line),
					Problem:         fmt.Sprintf("evaluation %s references missing capability %s.", evaluation.ID, capID),
					WhyItMatters:    "Evaluation coverage must trace to real capabilities before downstream actors can rely on it.",
					SuggestedRepair: fmt.Sprintf("Add %s to .ni/contract.json capabilities, or replace the docs reference with an existing capability ID.", capID),
					BlocksEnd:       true,
				}))
			}
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
	for _, risk := range c.Risks {
		if risk.Status != "accepted" {
			continue
		}
		doc, ok := sections[risk.ID]
		if !ok {
			findings = append(findings, missingDocsFinding(risk.ID, contractLocation("risks", risk.ID), fmt.Sprintf("accepted risk %s lacks docs discussion in %s.", risk.ID, RiskDoc), "Risks shape lock safety; missing docs hide the accepted risk context and mitigation from reviewers.", fmt.Sprintf("Add a `## %s: ...` section with `Severity: ...` and `Mitigation: ...`.", risk.ID)))
			continue
		}
		docSeverity := normalizeValue(doc.fields["severity"])
		if docSeverity == "" {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              risk.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("accepted risk %s lacks `Severity: ...` in docs.", risk.ID),
				WhyItMatters:    "Risk severity controls how reviewers judge the lock boundary.",
				SuggestedRepair: "Add the contract severity to the risk docs section.",
				BlocksEnd:       true,
			}))
		} else if docSeverity != normalizeValue(risk.Severity) {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              risk.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("%s severity is %q but .ni/contract.json says %q.", risk.ID, doc.fields["severity"], risk.Severity),
				WhyItMatters:    "Conflicting risk severity makes the accepted risk posture ambiguous.",
				SuggestedRepair: "Update the stale source so the docs severity and contract severity match.",
				BlocksEnd:       true,
			}))
		}
		if strings.TrimSpace(doc.fields["mitigation"]) == "" {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              risk.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("accepted risk %s lacks `Mitigation: ...` in docs.", risk.ID),
				WhyItMatters:    "Risk discussion is actionable only when the mitigation is visible in docs.",
				SuggestedRepair: "Add a `Mitigation: ...` line explaining how the risk is handled.",
				BlocksEnd:       true,
			}))
		}
	}
	return findings
}

func checkOpenQuestions(c contract.Contract, sections map[string]section) []Finding {
	contractIDs := make(map[string]contract.OpenQuestion, len(c.OpenQuestions))
	for _, openQuestion := range c.OpenQuestions {
		contractIDs[openQuestion.ID] = openQuestion
	}

	var findings []Finding
	for _, openQuestion := range c.OpenQuestions {
		doc, ok := sections[openQuestion.ID]
		if !ok {
			findings = append(findings, missingDocsFinding(openQuestion.ID, contractLocation("open_questions", openQuestion.ID), fmt.Sprintf("open question %s is not explained in %s.", openQuestion.ID, OpenQuestionDoc), "Open questions control blockers and deferrals; missing docs hide unresolved intent from review.", fmt.Sprintf("Add a `## %s: ...` section with `Blocker: ...` and `Status: ...`, or remove it from .ni/contract.json.", openQuestion.ID)))
			continue
		}
		if isClosed(openQuestion.Status) && openQuestion.Blocker {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              openQuestion.ID,
				Location:        contractLocation("open_questions", openQuestion.ID),
				Problem:         fmt.Sprintf("resolved open question %s is still marked as a blocker in .ni/contract.json.", openQuestion.ID),
				WhyItMatters:    "A resolved blocker should not keep ni-end blocked or confuse the next planning turn.",
				SuggestedRepair: "Set `blocker` to false, or reopen the question if it still blocks lock readiness.",
				BlocksEnd:       true,
			}))
		}
		docStatus := normalizeValue(doc.fields["status"])
		docBlocker := normalizeValue(doc.fields["blocker"])
		if (isClosed(openQuestion.Status) || isClosed(docStatus)) && docBlocker == "true" {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              openQuestion.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("resolved open question %s is still shown as a blocker in docs.", openQuestion.ID),
				WhyItMatters:    "Resolved blocker questions must stop appearing as lock blockers once intent is settled.",
				SuggestedRepair: "Change `Blocker: false` in docs, or reopen the question consistently in docs and contract.",
				BlocksEnd:       true,
			}))
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
	for _, decision := range c.Decisions {
		doc, ok := sections[decision.ID]
		if !ok {
			findings = append(findings, missingDocsFinding(decision.ID, contractLocation("decisions", decision.ID), fmt.Sprintf("decision %s is not explained in %s.", decision.ID, DecisionDoc), "Decisions define what downstream actors may trust; missing docs remove reviewable rationale.", fmt.Sprintf("Add a `## %s: ...` section with `Status: ...`, or remove the decision from .ni/contract.json.", decision.ID)))
			continue
		}
		docStatus := normalizeValue(doc.fields["status"])
		if docStatus == "" {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              decision.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("decision %s lacks `Status: ...` in docs.", decision.ID),
				WhyItMatters:    "Decision status controls whether downstream actors may depend on the decision.",
				SuggestedRepair: "Add the contract decision status to the docs section.",
				BlocksEnd:       true,
			}))
		} else if docStatus != normalizeValue(decision.Status) {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              decision.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("%s status is %q but .ni/contract.json says %q.", decision.ID, doc.fields["status"], decision.Status),
				WhyItMatters:    "Contradictory decision status changes whether downstream actors may rely on the decision.",
				SuggestedRepair: "Update the stale source so the docs status and contract status match.",
				BlocksEnd:       true,
			}))
		}
		if decisionTitlesContradict(doc.title, decision.Title) {
			findings = append(findings, diagnosticFinding(Diagnostic{
				ID:              decision.ID,
				Location:        location(doc.path, doc.line),
				Problem:         fmt.Sprintf("decision log title for %s contradicts the contract decision title.", decision.ID),
				WhyItMatters:    "A contradictory accepted decision gives downstream actors incompatible instructions.",
				SuggestedRepair: "Revise the docs heading or contract title so the accepted decision has one polarity.",
				BlocksEnd:       true,
			}))
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
	for idx, rawLine := range strings.Split(string(data), "\n") {
		line := strings.TrimSpace(rawLine)
		if strings.HasPrefix(line, "## ") {
			match := sectionHeading.FindStringSubmatch(line)
			if match == nil || match[2] != prefix {
				current = nil
				continue
			}
			sec := section{id: match[1], path: relPath, line: idx + 1, title: strings.TrimSpace(match[3]), fields: map[string]string{}}
			sections[sec.id] = sec
			current = &sec
			continue
		}
		if current == nil {
			continue
		}
		if line == "" {
			continue
		}
		match := fieldLine.FindStringSubmatch(line)
		if match == nil {
			stored := sections[current.id]
			stored.explanation = true
			sections[current.id] = stored
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

func docMentions(root string) []mention {
	var mentions []mention
	base := filepath.Join(root, "docs", "plan")
	_ = filepath.WalkDir(base, func(path string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		seenByLine := map[string]struct{}{}
		for idx, rawLine := range strings.Split(string(data), "\n") {
			seenByLine = map[string]struct{}{}
			if isEvaluationCapabilityField(rel, rawLine) {
				continue
			}
			for _, id := range planID.FindAllString(rawLine, -1) {
				key := fmt.Sprintf("%s:%d:%s", rel, idx+1, id)
				if _, ok := seenByLine[key]; ok {
					continue
				}
				seenByLine[key] = struct{}{}
				mentions = append(mentions, mention{id: id, path: rel, line: idx + 1})
			}
		}
		return nil
	})
	sort.Slice(mentions, func(i, j int) bool {
		if mentions[i].id != mentions[j].id {
			return mentions[i].id < mentions[j].id
		}
		if mentions[i].path != mentions[j].path {
			return mentions[i].path < mentions[j].path
		}
		return mentions[i].line < mentions[j].line
	})
	return mentions
}

func isEvaluationCapabilityField(rel string, rawLine string) bool {
	if rel != EvaluationDoc {
		return false
	}
	match := fieldLine.FindStringSubmatch(strings.TrimSpace(rawLine))
	if match == nil {
		return false
	}
	switch normalizeKey(match[1]) {
	case "capability", "capabilities", "capability id", "capability ids", "linked capabilities":
		return true
	default:
		return false
	}
}

func contractIDs(c contract.Contract) map[string]bool {
	ids := map[string]bool{}
	for _, item := range c.NonGoals {
		ids[item.ID] = true
	}
	for _, item := range c.Capabilities {
		ids[item.ID] = true
	}
	for _, item := range c.Requirements {
		ids[item.ID] = true
	}
	for _, item := range c.Decisions {
		ids[item.ID] = true
	}
	for _, item := range c.Risks {
		ids[item.ID] = true
	}
	for _, item := range c.Evaluations {
		ids[item.ID] = true
	}
	for _, item := range c.Artifacts {
		ids[item.ID] = true
	}
	for _, item := range c.OpenQuestions {
		ids[item.ID] = true
	}
	return ids
}

func contractCapabilityIDs(c contract.Contract) map[string]bool {
	ids := make(map[string]bool, len(c.Capabilities))
	for _, capability := range c.Capabilities {
		ids[capability.ID] = true
	}
	return ids
}

func referencedCapabilityIDs(doc section) []string {
	var ids []string
	for _, key := range []string{"capability", "capabilities", "capability id", "capability ids", "linked capabilities"} {
		for _, id := range planID.FindAllString(doc.fields[key], -1) {
			if strings.HasPrefix(id, "CAP-") {
				ids = append(ids, id)
			}
		}
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

func missingDocsFinding(id string, loc string, problem string, why string, repair string) Finding {
	return diagnosticFinding(Diagnostic{
		ID:              id,
		Location:        loc,
		Problem:         problem,
		WhyItMatters:    why,
		SuggestedRepair: repair,
		BlocksEnd:       true,
	})
}

func diagnosticFinding(d Diagnostic) Finding {
	return Finding{
		Message:    fmt.Sprintf("%s at %s: %s Suggested repair: %s Blocks ni-end: %t", d.ID, d.Location, d.Problem, d.SuggestedRepair, d.BlocksEnd),
		Diagnostic: d,
	}
}

func location(path string, line int) string {
	if line <= 0 {
		return path
	}
	return fmt.Sprintf("%s:%d", path, line)
}

func contractLocation(collection string, id string) string {
	return fmt.Sprintf(".ni/contract.json:%s[%s]", collection, id)
}

func isClosed(status string) bool {
	switch normalizeValue(status) {
	case "closed", "resolved":
		return true
	default:
		return false
	}
}

func decisionTitlesContradict(docsTitle string, contractTitle string) bool {
	leftPolarity, leftSubject, leftOK := decisionPolarity(docsTitle)
	rightPolarity, rightSubject, rightOK := decisionPolarity(contractTitle)
	return leftOK && rightOK && leftSubject == rightSubject && leftPolarity != rightPolarity
}

func decisionPolarity(title string) (int, string, bool) {
	text := wordText(title)
	negativePrefixes := []string{"do not use ", "do not require ", "must not ", "disable ", "disallow ", "avoid "}
	positivePrefixes := []string{"use ", "require ", "must ", "enable ", "allow "}
	for _, prefix := range negativePrefixes {
		if strings.HasPrefix(text, prefix) {
			return -1, strings.TrimSpace(strings.TrimPrefix(text, prefix)), true
		}
	}
	for _, prefix := range positivePrefixes {
		if strings.HasPrefix(text, prefix) {
			return 1, strings.TrimSpace(strings.TrimPrefix(text, prefix)), true
		}
	}
	return 0, "", false
}

func wordText(value string) string {
	value = strings.ToLower(value)
	var b strings.Builder
	lastSpace := true
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			b.WriteRune(r)
			lastSpace = false
			continue
		}
		if !lastSpace {
			b.WriteByte(' ')
			lastSpace = true
		}
	}
	return strings.TrimSpace(b.String())
}
