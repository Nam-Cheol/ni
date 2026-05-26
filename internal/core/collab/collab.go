package collab

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
)

const (
	DiffSchema     = "ni.collaboration.diff.v0"
	ConflictSchema = "ni.collaboration.conflicts.v0"
)

type Endpoint struct {
	Input        string `json:"input"`
	Root         string `json:"root,omitempty"`
	ContractPath string `json:"contract_path"`
	LockPath     string `json:"lock_path,omitempty"`
}

type Change struct {
	Kind       string          `json:"kind"`
	EntityType string          `json:"entity_type"`
	ID         string          `json:"id"`
	Message    string          `json:"message"`
	Before     json.RawMessage `json:"before,omitempty"`
	After      json.RawMessage `json:"after,omitempty"`
}

type DiffResult struct {
	Schema  string   `json:"schema"`
	Base    Endpoint `json:"base"`
	Head    Endpoint `json:"head"`
	Changes []Change `json:"changes"`
}

type Conflict struct {
	Code       string `json:"code"`
	Severity   string `json:"severity"`
	EntityType string `json:"entity_type,omitempty"`
	ID         string `json:"id,omitempty"`
	Message    string `json:"message"`
	BaseValue  string `json:"base_value,omitempty"`
	HeadValue  string `json:"head_value,omitempty"`
}

type ConflictResult struct {
	Schema    string     `json:"schema"`
	Base      Endpoint   `json:"base"`
	Head      Endpoint   `json:"head"`
	Conflicts []Conflict `json:"conflicts"`
}

type state struct {
	endpoint   Endpoint
	contract   contract.Contract
	mismatches []lock.Mismatch
}

func Diff(basePath string, headPath string) (DiffResult, error) {
	base, head, err := resolvePair(basePath, headPath)
	if err != nil {
		return DiffResult{}, err
	}
	changes := diffContracts(base.contract, head.contract)
	if changes == nil {
		changes = []Change{}
	}
	return DiffResult{
		Schema:  DiffSchema,
		Base:    base.endpoint,
		Head:    head.endpoint,
		Changes: changes,
	}, nil
}

func Conflicts(basePath string, headPath string) (ConflictResult, error) {
	base, head, err := resolvePair(basePath, headPath)
	if err != nil {
		return ConflictResult{}, err
	}
	conflicts := semanticConflicts(base, head)
	if conflicts == nil {
		conflicts = []Conflict{}
	}
	return ConflictResult{
		Schema:    ConflictSchema,
		Base:      base.endpoint,
		Head:      head.endpoint,
		Conflicts: conflicts,
	}, nil
}

func FormatDiff(result DiffResult) string {
	var b strings.Builder
	if len(result.Changes) == 0 {
		b.WriteString("no contract changes\n")
		return b.String()
	}
	b.WriteString("contract diff\n")
	for _, change := range result.Changes {
		fmt.Fprintf(&b, "- %s %s %s: %s\n", change.Kind, change.EntityType, change.ID, change.Message)
	}
	return b.String()
}

func FormatConflicts(result ConflictResult) string {
	var b strings.Builder
	if len(result.Conflicts) == 0 {
		b.WriteString("no collaboration conflicts\n")
		return b.String()
	}
	b.WriteString("collaboration conflicts\n")
	for _, conflict := range result.Conflicts {
		id := conflict.ID
		if id == "" {
			id = conflict.EntityType
		}
		if id == "" {
			id = conflict.Code
		}
		fmt.Fprintf(&b, "- %s %s: %s\n", conflict.Severity, id, conflict.Message)
	}
	return b.String()
}

func resolvePair(basePath string, headPath string) (state, state, error) {
	base, err := resolve(basePath)
	if err != nil {
		return state{}, state{}, fmt.Errorf("load base: %w", err)
	}
	head, err := resolve(headPath)
	if err != nil {
		return state{}, state{}, fmt.Errorf("load head: %w", err)
	}
	return base, head, nil
}

func resolve(input string) (state, error) {
	if strings.TrimSpace(input) == "" {
		return state{}, fmt.Errorf("missing path")
	}
	clean := filepath.Clean(input)
	info, err := os.Stat(clean)
	if err != nil {
		return state{}, err
	}
	if info.IsDir() {
		return resolveProjectRoot(input, clean, filepath.Join(clean, ".ni", "contract.json"), filepath.Join(clean, ".ni", "plan.lock.json"))
	}

	var probe struct {
		Schema string `json:"schema"`
	}
	data, err := os.ReadFile(clean)
	if err != nil {
		return state{}, err
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return state{}, fmt.Errorf("malformed JSON: %w", err)
	}
	switch probe.Schema {
	case contract.Schema:
		root, lockPath := rootFromContractPath(clean)
		return resolveProjectRoot(input, root, clean, lockPath)
	case lock.Schema:
		root := rootFromLockPath(clean)
		return resolveProjectRoot(input, root, filepath.Join(root, ".ni", "contract.json"), clean)
	default:
		return state{}, fmt.Errorf("unsupported planning file schema %q", probe.Schema)
	}
}

func resolveProjectRoot(input string, root string, contractPath string, lockPath string) (state, error) {
	loaded, err := contract.LoadFile(contractPath)
	if err != nil {
		return state{}, err
	}
	endpoint := Endpoint{
		Input:        input,
		Root:         root,
		ContractPath: contractPath,
	}
	var mismatches []lock.Mismatch
	if lockPath != "" {
		if _, err := os.Stat(lockPath); err == nil {
			endpoint.LockPath = lockPath
			mismatches, err = verifyLock(root, lockPath)
			if err != nil {
				return state{}, err
			}
		} else if !os.IsNotExist(err) {
			return state{}, err
		}
	}
	return state{endpoint: endpoint, contract: loaded, mismatches: mismatches}, nil
}

func rootFromContractPath(path string) (string, string) {
	dir := filepath.Dir(path)
	if filepath.Base(dir) == ".ni" {
		root := filepath.Dir(dir)
		return root, filepath.Join(root, ".ni", "plan.lock.json")
	}
	return dir, ""
}

func rootFromLockPath(path string) string {
	dir := filepath.Dir(path)
	if filepath.Base(dir) == "locks" && filepath.Base(filepath.Dir(dir)) == ".ni" {
		return filepath.Dir(filepath.Dir(dir))
	}
	if filepath.Base(dir) == ".ni" {
		return filepath.Dir(dir)
	}
	return filepath.Dir(filepath.Dir(path))
}

func verifyLock(root string, lockPath string) ([]lock.Mismatch, error) {
	lockfile, err := lock.LoadFile(lockPath)
	if err != nil {
		return nil, err
	}
	var mismatches []lock.Mismatch
	for _, file := range lockfile.Files {
		got, err := fileSHA256(filepath.Join(root, file.Path))
		if err != nil {
			mismatches = append(mismatches, lock.Mismatch{Path: file.Path, WantHash: file.SHA256, GotHash: "missing"})
			continue
		}
		if got != file.SHA256 {
			mismatches = append(mismatches, lock.Mismatch{Path: file.Path, WantHash: file.SHA256, GotHash: got})
		}
	}
	return mismatches, nil
}

func fileSHA256(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func diffContracts(base contract.Contract, head contract.Contract) []Change {
	var changes []Change
	changes = append(changes, diffItems("non_goal", nonGoalItems(base), nonGoalItems(head))...)
	changes = append(changes, diffItems("capability", capabilityItems(base), capabilityItems(head))...)
	changes = append(changes, diffItems("requirement", requirementItems(base), requirementItems(head))...)
	changes = append(changes, diffItems("decision", decisionItems(base), decisionItems(head))...)
	changes = append(changes, diffItems("risk", riskItems(base), riskItems(head))...)
	changes = append(changes, diffItems("evaluation", evaluationItems(base), evaluationItems(head))...)
	changes = append(changes, diffItems("artifact", artifactItems(base), artifactItems(head))...)
	changes = append(changes, diffItems("open_question", openQuestionItems(base), openQuestionItems(head))...)
	sortChanges(changes)
	return changes
}

type item struct {
	id    string
	title string
	value any
}

func diffItems(entityType string, base map[string]item, head map[string]item) []Change {
	var changes []Change
	ids := unionIDs(base, head)
	for _, id := range ids {
		baseItem, inBase := base[id]
		headItem, inHead := head[id]
		switch {
		case !inBase:
			changes = append(changes, Change{
				Kind:       "added",
				EntityType: entityType,
				ID:         id,
				Message:    fmt.Sprintf("added %s", headItem.title),
				After:      mustJSON(headItem.value),
			})
		case !inHead:
			changes = append(changes, Change{
				Kind:       "removed",
				EntityType: entityType,
				ID:         id,
				Message:    fmt.Sprintf("removed %s", baseItem.title),
				Before:     mustJSON(baseItem.value),
			})
		default:
			before := mustJSON(baseItem.value)
			after := mustJSON(headItem.value)
			if string(before) != string(after) {
				changes = append(changes, Change{
					Kind:       "modified",
					EntityType: entityType,
					ID:         id,
					Message:    fmt.Sprintf("changed %s", headItem.title),
					Before:     before,
					After:      after,
				})
			}
		}
	}
	return changes
}

func semanticConflicts(base state, head state) []Conflict {
	var conflicts []Conflict
	for _, mismatch := range base.mismatches {
		conflicts = append(conflicts, lockMismatchConflict("base", mismatch))
	}
	for _, mismatch := range head.mismatches {
		conflicts = append(conflicts, lockMismatchConflict("head", mismatch))
	}
	conflicts = append(conflicts, changedSameIDConflicts(base.contract, head.contract)...)
	conflicts = append(conflicts, contradictoryDecisionConflicts(base.contract, head.contract)...)
	conflicts = append(conflicts, removedCapabilityReferenceConflicts(base.contract, head.contract)...)
	conflicts = append(conflicts, loweredRiskConflicts(base.contract, head.contract)...)
	conflicts = append(conflicts, weakenedAcceptanceConflicts(base.contract, head.contract)...)
	sortConflicts(conflicts)
	return conflicts
}

func lockMismatchConflict(side string, mismatch lock.Mismatch) Conflict {
	return Conflict{
		Code:      "lock_hash_mismatch",
		Severity:  "BLOCKED",
		ID:        mismatch.Path,
		Message:   fmt.Sprintf("%s lock hash mismatch for %s without relock", side, mismatch.Path),
		BaseValue: mismatch.WantHash,
		HeadValue: mismatch.GotHash,
	}
}

func changedSameIDConflicts(base contract.Contract, head contract.Contract) []Conflict {
	var conflicts []Conflict
	for _, change := range diffContracts(base, head) {
		if change.Kind != "modified" {
			continue
		}
		switch change.EntityType {
		case "capability", "requirement", "decision", "risk", "evaluation":
			conflicts = append(conflicts, Conflict{
				Code:       "same_id_changed",
				Severity:   "BLOCKED",
				EntityType: change.EntityType,
				ID:         change.ID,
				Message:    fmt.Sprintf("same %s ID changed between base and head", change.EntityType),
				BaseValue:  compactRaw(change.Before),
				HeadValue:  compactRaw(change.After),
			})
		}
	}
	return conflicts
}

func contradictoryDecisionConflicts(base contract.Contract, head contract.Contract) []Conflict {
	baseDecisions := decisionItems(base)
	var conflicts []Conflict
	for _, decision := range head.Decisions {
		if _, existed := baseDecisions[decision.ID]; existed || decision.Status != "accepted" {
			continue
		}
		headPolarity, headSubject, ok := decisionPolarity(decision.Title)
		if !ok {
			continue
		}
		for _, baseDecision := range base.Decisions {
			if baseDecision.Status != "accepted" {
				continue
			}
			basePolarity, baseSubject, ok := decisionPolarity(baseDecision.Title)
			if !ok {
				continue
			}
			if baseSubject == headSubject && basePolarity != headPolarity {
				conflicts = append(conflicts, Conflict{
					Code:       "contradictory_decision",
					Severity:   "BLOCKED",
					EntityType: "decision",
					ID:         decision.ID,
					Message:    fmt.Sprintf("%s contradicts accepted decision %s", decision.ID, baseDecision.ID),
					BaseValue:  baseDecision.Title,
					HeadValue:  decision.Title,
				})
			}
		}
	}
	return conflicts
}

func removedCapabilityReferenceConflicts(base contract.Contract, head contract.Contract) []Conflict {
	headCapabilities := capabilityItems(head)
	headEvaluations := evaluationItems(head)
	headArtifacts := artifactItems(head)
	var conflicts []Conflict
	for _, cap := range base.Capabilities {
		if _, ok := headCapabilities[cap.ID]; ok {
			continue
		}
		for _, id := range cap.Evaluations {
			if _, ok := headEvaluations[id]; ok {
				conflicts = append(conflicts, Conflict{
					Code:       "removed_capability_referenced",
					Severity:   "BLOCKED",
					EntityType: "capability",
					ID:         cap.ID,
					Message:    fmt.Sprintf("removed capability %s while evaluation %s still exists", cap.ID, id),
					BaseValue:  id,
					HeadValue:  id,
				})
			}
		}
		for _, id := range cap.Artifacts {
			if _, ok := headArtifacts[id]; ok {
				conflicts = append(conflicts, Conflict{
					Code:       "removed_capability_referenced",
					Severity:   "BLOCKED",
					EntityType: "capability",
					ID:         cap.ID,
					Message:    fmt.Sprintf("removed capability %s while artifact %s still exists", cap.ID, id),
					BaseValue:  id,
					HeadValue:  id,
				})
			}
		}
	}
	return conflicts
}

func loweredRiskConflicts(base contract.Contract, head contract.Contract) []Conflict {
	headRisks := riskItems(head)
	var conflicts []Conflict
	for _, baseRisk := range base.Risks {
		headItem, ok := headRisks[baseRisk.ID]
		if !ok {
			continue
		}
		headRisk := headItem.value.(contract.Risk)
		if severityRank(headRisk.Severity) >= severityRank(baseRisk.Severity) {
			continue
		}
		if mitigationChanged(baseRisk.Mitigation, headRisk.Mitigation) {
			continue
		}
		conflicts = append(conflicts, Conflict{
			Code:       "risk_severity_lowered",
			Severity:   "BLOCKED",
			EntityType: "risk",
			ID:         baseRisk.ID,
			Message:    fmt.Sprintf("risk severity lowered from %s to %s without new mitigation or amendment reason", baseRisk.Severity, headRisk.Severity),
			BaseValue:  baseRisk.Severity,
			HeadValue:  headRisk.Severity,
		})
	}
	return conflicts
}

func weakenedAcceptanceConflicts(base contract.Contract, head contract.Contract) []Conflict {
	headRequirements := requirementItems(head)
	var conflicts []Conflict
	for _, baseReq := range base.Requirements {
		headItem, ok := headRequirements[baseReq.ID]
		if !ok {
			continue
		}
		headReq := headItem.value.(contract.Requirement)
		if baseReq.Status == "accepted" && headReq.Status != "accepted" {
			conflicts = append(conflicts, Conflict{
				Code:       "acceptance_weakened",
				Severity:   "BLOCKED",
				EntityType: "requirement",
				ID:         baseReq.ID,
				Message:    fmt.Sprintf("accepted requirement status weakened from %s to %s", baseReq.Status, headReq.Status),
				BaseValue:  baseReq.Status,
				HeadValue:  headReq.Status,
			})
			continue
		}
		if titleWeakened(baseReq.Title, headReq.Title) {
			conflicts = append(conflicts, Conflict{
				Code:       "acceptance_weakened",
				Severity:   "BLOCKED",
				EntityType: "requirement",
				ID:         baseReq.ID,
				Message:    "acceptance criterion wording was weakened",
				BaseValue:  baseReq.Title,
				HeadValue:  headReq.Title,
			})
		}
	}
	return conflicts
}

func mitigationChanged(base string, head string) bool {
	base = strings.TrimSpace(base)
	head = strings.TrimSpace(head)
	if head == "" {
		return false
	}
	return head != base || strings.Contains(strings.ToLower(head), "amend")
}

func severityRank(value string) int {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "critical":
		return 4
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	default:
		return 0
	}
}

func titleWeakened(base string, head string) bool {
	baseNorm := wordText(base)
	headNorm := wordText(head)
	pairs := [][2]string{
		{"must", "should"},
		{"must", "may"},
		{"must", "can"},
		{"shall", "should"},
		{"shall", "may"},
		{"required", "optional"},
		{"always", "usually"},
		{"cannot", "can"},
		{"must not", "may"},
		{"without", "with"},
	}
	for _, pair := range pairs {
		if containsPhrase(baseNorm, pair[0]) && containsPhrase(headNorm, pair[1]) && !containsPhrase(headNorm, pair[0]) {
			return true
		}
	}
	return false
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

func containsPhrase(text string, phrase string) bool {
	return strings.Contains(" "+text+" ", " "+phrase+" ")
}

func nonGoalItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.NonGoals))
	for _, value := range c.NonGoals {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func capabilityItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.Capabilities))
	for _, value := range c.Capabilities {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func requirementItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.Requirements))
	for _, value := range c.Requirements {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func decisionItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.Decisions))
	for _, value := range c.Decisions {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func riskItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.Risks))
	for _, value := range c.Risks {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func evaluationItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.Evaluations))
	for _, value := range c.Evaluations {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func artifactItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.Artifacts))
	for _, value := range c.Artifacts {
		items[value.ID] = item{id: value.ID, title: value.Path, value: value}
	}
	return items
}

func openQuestionItems(c contract.Contract) map[string]item {
	items := make(map[string]item, len(c.OpenQuestions))
	for _, value := range c.OpenQuestions {
		items[value.ID] = item{id: value.ID, title: value.Title, value: value}
	}
	return items
}

func unionIDs(base map[string]item, head map[string]item) []string {
	seen := map[string]struct{}{}
	for id := range base {
		seen[id] = struct{}{}
	}
	for id := range head {
		seen[id] = struct{}{}
	}
	ids := make([]string, 0, len(seen))
	for id := range seen {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func mustJSON(value any) json.RawMessage {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return json.RawMessage(data)
}

func compactRaw(value json.RawMessage) string {
	var out bytes.Buffer
	if err := json.Compact(&out, value); err != nil {
		return string(value)
	}
	return out.String()
}

func sortChanges(changes []Change) {
	sort.Slice(changes, func(i, j int) bool {
		if changes[i].EntityType != changes[j].EntityType {
			return changes[i].EntityType < changes[j].EntityType
		}
		if changes[i].ID != changes[j].ID {
			return changes[i].ID < changes[j].ID
		}
		return changes[i].Kind < changes[j].Kind
	})
}

func sortConflicts(conflicts []Conflict) {
	sort.Slice(conflicts, func(i, j int) bool {
		if conflicts[i].Code != conflicts[j].Code {
			return conflicts[i].Code < conflicts[j].Code
		}
		if conflicts[i].EntityType != conflicts[j].EntityType {
			return conflicts[i].EntityType < conflicts[j].EntityType
		}
		return conflicts[i].ID < conflicts[j].ID
	})
}
