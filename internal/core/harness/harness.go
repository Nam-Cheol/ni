package harness

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
	"ni/internal/core/pressure"
)

const Schema = "ni.generated_harness.v0"

const CandidateLedgerSchema = "ni.harness_candidates.v0"

const (
	StatusProposed           = "proposed"
	StatusCandidate          = "candidate"
	StatusValidatedCandidate = "validated_candidate"
	StatusUserAccepted       = "user_accepted"
	StatusActiveRule         = "active_rule"
	StatusRetired            = "retired"
)

type Proposal struct {
	Schema               string             `json:"schema"`
	SourceLockHash       string             `json:"source_lock_hash"`
	SelectedCapabilities []string           `json:"selected_capabilities"`
	WorkPackets          []WorkPacket       `json:"work_packets"`
	Validations          []Validation       `json:"validations"`
	EvidenceLocations    []EvidenceLocation `json:"evidence_locations"`
	KnownRisks           []KnownRisk        `json:"known_risks"`
	NonGoals             []NonGoal          `json:"non_goals"`
}

type WorkPacket struct {
	ID                string   `json:"id"`
	Title             string   `json:"title"`
	Capabilities      []string `json:"capabilities"`
	DependsOn         []string `json:"depends_on"`
	Artifacts         []string `json:"artifacts"`
	Validations       []string `json:"validations"`
	EvidenceLocations []string `json:"evidence_locations"`
	KnownRisks        []string `json:"known_risks"`
}

type Validation struct {
	ID       string `json:"id"`
	Method   string `json:"method"`
	Required bool   `json:"required"`
}

type EvidenceLocation struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
	Path string `json:"path"`
}

type KnownRisk struct {
	ID         string `json:"id"`
	Severity   string `json:"severity"`
	Mitigation string `json:"mitigation"`
}

type NonGoal struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type CandidateLedger struct {
	Schema       string      `json:"schema"`
	ActiveRuleID string      `json:"active_rule_id,omitempty"`
	Candidates   []Candidate `json:"candidates"`
}

type Candidate struct {
	ID                        string   `json:"id"`
	Status                    string   `json:"status"`
	Target                    string   `json:"target"`
	IntendedDownstreamRuntime string   `json:"intended_downstream_runtime"`
	RequiredEvidence          []string `json:"required_evidence"`
	Constraints               []string `json:"constraints"`
	ForbiddenBehavior         []string `json:"forbidden_behavior"`
	RelatedLockHash           string   `json:"related_lock_hash"`
	RelatedPressureIDs        []string `json:"related_pressure_ids"`
	ValidationEvidencePath    string   `json:"validation_evidence_path,omitempty"`
	RequiresUserAcceptance    bool     `json:"requires_user_acceptance"`
	ExecutesInsideKernel      bool     `json:"executes_inside_kernel"`
}

func Plan(dir string) (Proposal, error) {
	root := filepath.Clean(dir)
	verification, err := lock.Verify(root)
	if err != nil {
		return Proposal{}, err
	}
	if len(verification.Mismatches) > 0 {
		return Proposal{}, fmt.Errorf("BLOCKED: lock hash mismatch for %s", verification.Mismatches[0].Path)
	}

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	if err != nil {
		return Proposal{}, err
	}
	sourceHash, err := fileSHA256(filepath.Join(root, ".ni", "plan.lock.json"))
	if err != nil {
		return Proposal{}, err
	}

	return buildProposal(c, sourceHash), nil
}

func Candidates(dir string) (CandidateLedger, error) {
	if _, _, err := verifyLockedPlan(dir); err != nil {
		return CandidateLedger{}, err
	}
	return loadCandidates(StorePath(dir))
}

func ProposeFromPressure(dir string, pressureID string) (Candidate, error) {
	root, lockHash, err := verifyLockedPlan(dir)
	if err != nil {
		return Candidate{}, err
	}
	pressureID = strings.TrimSpace(pressureID)
	if pressureID == "" {
		return Candidate{}, fmt.Errorf("missing pressure id")
	}

	ledger, err := pressure.Load(root)
	if err != nil {
		return Candidate{}, err
	}
	item, ok := findPressureItem(ledger, pressureID)
	if !ok {
		return Candidate{}, fmt.Errorf("unknown pressure id %q", pressureID)
	}
	if item.Kind != pressure.KindHarnessCandidate {
		return Candidate{}, fmt.Errorf("pressure %s is %s, not %s", pressureID, item.Kind, pressure.KindHarnessCandidate)
	}
	if item.Status != pressure.StatusAccepted {
		return Candidate{}, fmt.Errorf("pressure %s must be accepted before proposing a harness candidate", pressureID)
	}

	candidates, err := loadCandidates(StorePath(root))
	if err != nil {
		return Candidate{}, err
	}
	candidate := Candidate{
		ID:                        nextCandidateID(candidates.Candidates),
		Status:                    StatusProposed,
		Target:                    item.ProposedAction,
		IntendedDownstreamRuntime: "downstream runtime selected outside ni core",
		RequiredEvidence: []string{
			"Evidence path supplied with ni harness validate",
		},
		Constraints: []string{
			"Derived from the locked planning contract.",
			"Related lock hash must remain valid before candidate lifecycle changes.",
			"Downstream runtime state is mutable and not owned by ni core.",
		},
		ForbiddenBehavior: []string{
			"execute shell commands",
			"call Codex",
			"call Hyper Run",
			"call namba-ai",
			"call GitHub or CI",
			"weaken readiness gates",
		},
		RelatedLockHash:        "sha256:" + lockHash,
		RelatedPressureIDs:     []string{pressureID},
		RequiresUserAcceptance: true,
		ExecutesInsideKernel:   false,
	}
	candidate.normalize()
	if err := candidate.Validate(); err != nil {
		return Candidate{}, err
	}
	candidates.Candidates = append(candidates.Candidates, candidate)
	if err := SaveCandidates(root, candidates); err != nil {
		return Candidate{}, err
	}
	return candidate, nil
}

func ValidateCandidate(dir string, candidateID string, evidencePath string) (Candidate, error) {
	root, _, err := verifyLockedPlan(dir)
	if err != nil {
		return Candidate{}, err
	}
	candidateID = strings.TrimSpace(candidateID)
	if candidateID == "" {
		return Candidate{}, fmt.Errorf("missing candidate id")
	}
	cleanEvidencePath, err := validateEvidencePath(root, evidencePath)
	if err != nil {
		return Candidate{}, err
	}

	ledger, err := loadCandidates(StorePath(root))
	if err != nil {
		return Candidate{}, err
	}
	for i := range ledger.Candidates {
		if ledger.Candidates[i].ID != candidateID {
			continue
		}
		switch ledger.Candidates[i].Status {
		case StatusProposed, StatusCandidate:
		default:
			return Candidate{}, fmt.Errorf("candidate %s cannot be validated from status %s", candidateID, ledger.Candidates[i].Status)
		}
		ledger.Candidates[i].Status = StatusValidatedCandidate
		ledger.Candidates[i].ValidationEvidencePath = cleanEvidencePath
		ledger.Candidates[i].normalize()
		if err := SaveCandidates(root, ledger); err != nil {
			return Candidate{}, err
		}
		return ledger.Candidates[i], nil
	}
	return Candidate{}, fmt.Errorf("unknown candidate id %q", candidateID)
}

func AcceptCandidate(dir string, candidateID string) (Candidate, error) {
	root, _, err := verifyLockedPlan(dir)
	if err != nil {
		return Candidate{}, err
	}
	candidateID = strings.TrimSpace(candidateID)
	if candidateID == "" {
		return Candidate{}, fmt.Errorf("missing candidate id")
	}

	ledger, err := loadCandidates(StorePath(root))
	if err != nil {
		return Candidate{}, err
	}
	for i := range ledger.Candidates {
		if ledger.Candidates[i].ID != candidateID {
			continue
		}
		if ledger.Candidates[i].Status != StatusValidatedCandidate {
			return Candidate{}, fmt.Errorf("candidate %s requires validation before user acceptance", candidateID)
		}
		ledger.Candidates[i].Status = StatusUserAccepted
		ledger.ActiveRuleID = candidateID
		ledger.Candidates[i].normalize()
		if err := SaveCandidates(root, ledger); err != nil {
			return Candidate{}, err
		}
		return ledger.Candidates[i], nil
	}
	return Candidate{}, fmt.Errorf("unknown candidate id %q", candidateID)
}

func RetireCandidate(dir string, candidateID string) (Candidate, error) {
	root, _, err := verifyLockedPlan(dir)
	if err != nil {
		return Candidate{}, err
	}
	candidateID = strings.TrimSpace(candidateID)
	if candidateID == "" {
		return Candidate{}, fmt.Errorf("missing candidate id")
	}

	ledger, err := loadCandidates(StorePath(root))
	if err != nil {
		return Candidate{}, err
	}
	for i := range ledger.Candidates {
		if ledger.Candidates[i].ID != candidateID {
			continue
		}
		ledger.Candidates[i].Status = StatusRetired
		if ledger.ActiveRuleID == candidateID {
			ledger.ActiveRuleID = ""
		}
		ledger.Candidates[i].normalize()
		if err := SaveCandidates(root, ledger); err != nil {
			return Candidate{}, err
		}
		return ledger.Candidates[i], nil
	}
	return Candidate{}, fmt.Errorf("unknown candidate id %q", candidateID)
}

func SaveCandidates(dir string, ledger CandidateLedger) error {
	if _, _, err := verifyLockedPlan(dir); err != nil {
		return err
	}
	ledger.normalize()
	if err := ledger.Validate(); err != nil {
		return err
	}
	path := StorePath(dir)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(data, '\n'), 0o644)
}

func StorePath(dir string) string {
	root := strings.TrimSpace(dir)
	if root == "" {
		root = "."
	}
	return filepath.Join(filepath.Clean(root), ".ni", "harness.candidates.json")
}

func FormatText(p Proposal) string {
	var b strings.Builder
	fmt.Fprintf(&b, "generated harness proposal\n")
	fmt.Fprintf(&b, "source_lock_hash: %s\n\n", p.SourceLockHash)
	b.WriteString("selected capabilities:\n")
	for _, id := range p.SelectedCapabilities {
		fmt.Fprintf(&b, "- %s\n", id)
	}
	b.WriteString("\nwork packets:\n")
	for _, packet := range p.WorkPackets {
		fmt.Fprintf(&b, "- %s: %s\n", packet.ID, packet.Title)
		fmt.Fprintf(&b, "  capabilities: %s\n", strings.Join(packet.Capabilities, ", "))
		fmt.Fprintf(&b, "  validations: %s\n", strings.Join(packet.Validations, ", "))
	}
	b.WriteString("\nnon-goals:\n")
	for _, nonGoal := range p.NonGoals {
		fmt.Fprintf(&b, "- %s: %s\n", nonGoal.ID, nonGoal.Title)
	}
	return b.String()
}

func FormatCandidates(ledger CandidateLedger) string {
	if len(ledger.Candidates) == 0 {
		return "no harness candidates\n"
	}
	var b strings.Builder
	fmt.Fprintf(&b, "harness candidates: %d item(s)\n", len(ledger.Candidates))
	if ledger.ActiveRuleID != "" {
		fmt.Fprintf(&b, "active_rule_id: %s\n", ledger.ActiveRuleID)
	}
	for _, candidate := range ledger.Candidates {
		fmt.Fprintf(&b, "%s\t%s\t%s\truntime:%s", candidate.ID, candidate.Status, candidate.Target, candidate.IntendedDownstreamRuntime)
		if candidate.ValidationEvidencePath != "" {
			fmt.Fprintf(&b, "\tevidence:%s", candidate.ValidationEvidencePath)
		}
		if len(candidate.RelatedPressureIDs) > 0 {
			fmt.Fprintf(&b, "\tpressure:%s", strings.Join(candidate.RelatedPressureIDs, ","))
		}
		b.WriteByte('\n')
	}
	return b.String()
}

func buildProposal(c contract.Contract, sourceHash string) Proposal {
	evidence := buildEvidence(c.Artifacts)
	validations := buildValidations(c.Evaluations)
	risks := buildRisks(c.Risks)
	nonGoals := buildNonGoals(c.NonGoals)

	evidenceIDs := make([]string, 0, len(evidence))
	for _, item := range evidence {
		evidenceIDs = append(evidenceIDs, item.ID)
	}

	capabilityToPacket := make(map[string]string, len(c.Capabilities))
	packetNumber := 1
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		capabilityToPacket[capability.ID] = fmt.Sprintf("WP-%03d", packetNumber)
		packetNumber++
	}

	selected := make([]string, 0, len(capabilityToPacket))
	packets := make([]WorkPacket, 0, len(capabilityToPacket))
	for _, capability := range c.Capabilities {
		if capability.Status != "accepted" {
			continue
		}
		selected = append(selected, capability.ID)
		packets = append(packets, WorkPacket{
			ID:                capabilityToPacket[capability.ID],
			Title:             capability.Title,
			Capabilities:      []string{capability.ID},
			DependsOn:         dependencyPackets(capability.Dependencies, capabilityToPacket),
			Artifacts:         capability.Artifacts,
			Validations:       capability.Evaluations,
			EvidenceLocations: evidenceIDs,
			KnownRisks:        capability.Risks,
		})
	}

	return Proposal{
		Schema:               Schema,
		SourceLockHash:       "sha256:" + sourceHash,
		SelectedCapabilities: selected,
		WorkPackets:          packets,
		Validations:          validations,
		EvidenceLocations:    evidence,
		KnownRisks:           risks,
		NonGoals:             nonGoals,
	}
}

func loadCandidates(path string) (CandidateLedger, error) {
	file, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return CandidateLedger{Schema: CandidateLedgerSchema, Candidates: []Candidate{}}, nil
		}
		return CandidateLedger{}, err
	}
	defer file.Close()

	var ledger CandidateLedger
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&ledger); err != nil {
		return CandidateLedger{}, fmt.Errorf("malformed harness candidate ledger JSON: %w", err)
	}
	ledger.normalize()
	if err := ledger.Validate(); err != nil {
		return CandidateLedger{}, err
	}
	return ledger, nil
}

func (l *CandidateLedger) normalize() {
	if l.Schema == "" {
		l.Schema = CandidateLedgerSchema
	}
	if l.Candidates == nil {
		l.Candidates = []Candidate{}
	}
	for i := range l.Candidates {
		l.Candidates[i].normalize()
	}
}

func (l CandidateLedger) Validate() error {
	if strings.TrimSpace(l.Schema) == "" {
		return fmt.Errorf("harness candidate ledger missing schema")
	}
	if l.Schema != CandidateLedgerSchema {
		return fmt.Errorf("unsupported harness candidate ledger schema %q", l.Schema)
	}
	seen := map[string]struct{}{}
	activeSeen := l.ActiveRuleID == ""
	for _, candidate := range l.Candidates {
		if err := candidate.Validate(); err != nil {
			return err
		}
		if _, ok := seen[candidate.ID]; ok {
			return fmt.Errorf("duplicate harness candidate id %q", candidate.ID)
		}
		seen[candidate.ID] = struct{}{}
		if candidate.ID == l.ActiveRuleID {
			if candidate.Status != StatusUserAccepted && candidate.Status != StatusActiveRule {
				return fmt.Errorf("active rule candidate %s has status %s", candidate.ID, candidate.Status)
			}
			activeSeen = true
		}
	}
	if !activeSeen {
		return fmt.Errorf("active rule id %q does not match a candidate", l.ActiveRuleID)
	}
	return nil
}

func (c *Candidate) normalize() {
	if c.RequiredEvidence == nil {
		c.RequiredEvidence = []string{}
	}
	if c.Constraints == nil {
		c.Constraints = []string{}
	}
	if c.ForbiddenBehavior == nil {
		c.ForbiddenBehavior = []string{}
	}
	if c.RelatedPressureIDs == nil {
		c.RelatedPressureIDs = []string{}
	}
}

func (c Candidate) Validate() error {
	var missing []string
	if strings.TrimSpace(c.ID) == "" {
		missing = append(missing, "id")
	}
	if strings.TrimSpace(c.Status) == "" {
		missing = append(missing, "status")
	}
	if strings.TrimSpace(c.Target) == "" {
		missing = append(missing, "target")
	}
	if strings.TrimSpace(c.IntendedDownstreamRuntime) == "" {
		missing = append(missing, "intended_downstream_runtime")
	}
	if strings.TrimSpace(c.RelatedLockHash) == "" {
		missing = append(missing, "related_lock_hash")
	}
	if len(c.RequiredEvidence) == 0 {
		missing = append(missing, "required_evidence")
	}
	if len(c.Constraints) == 0 {
		missing = append(missing, "constraints")
	}
	if len(c.ForbiddenBehavior) == 0 {
		missing = append(missing, "forbidden_behavior")
	}
	if len(c.RelatedPressureIDs) == 0 {
		missing = append(missing, "related_pressure_ids")
	}
	if len(missing) > 0 {
		return fmt.Errorf("harness candidate missing required field(s): %s", strings.Join(missing, ", "))
	}
	if !validCandidateStatus(c.Status) {
		return fmt.Errorf("invalid harness candidate status %q", c.Status)
	}
	if !strings.HasPrefix(c.RelatedLockHash, "sha256:") {
		return fmt.Errorf("related_lock_hash must use sha256 prefix")
	}
	if !c.RequiresUserAcceptance {
		return fmt.Errorf("harness candidate must require user acceptance")
	}
	if c.ExecutesInsideKernel {
		return fmt.Errorf("harness candidate must not execute inside ni core")
	}
	if (c.Status == StatusValidatedCandidate || c.Status == StatusUserAccepted || c.Status == StatusActiveRule) && strings.TrimSpace(c.ValidationEvidencePath) == "" {
		return fmt.Errorf("harness candidate %s requires validation evidence path", c.ID)
	}
	return nil
}

func validCandidateStatus(status string) bool {
	switch status {
	case StatusProposed, StatusCandidate, StatusValidatedCandidate, StatusUserAccepted, StatusActiveRule, StatusRetired:
		return true
	default:
		return false
	}
}

func findPressureItem(ledger pressure.Ledger, id string) (pressure.Item, bool) {
	for _, item := range ledger.Items {
		if item.ID == id {
			return item, true
		}
	}
	return pressure.Item{}, false
}

func nextCandidateID(items []Candidate) string {
	max := 0
	for _, item := range items {
		if !strings.HasPrefix(item.ID, "HC-") {
			continue
		}
		var n int
		if _, err := fmt.Sscanf(item.ID, "HC-%03d", &n); err == nil && n > max {
			max = n
		}
	}
	return fmt.Sprintf("HC-%03d", max+1)
}

func validateEvidencePath(root string, path string) (string, error) {
	path = strings.TrimSpace(path)
	if path == "" {
		return "", fmt.Errorf("missing evidence path")
	}
	checkPath := path
	if !filepath.IsAbs(checkPath) {
		checkPath = filepath.Join(root, path)
	}
	info, err := os.Stat(checkPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("evidence path does not exist: %s", path)
		}
		return "", err
	}
	if info.IsDir() {
		return "", fmt.Errorf("evidence path is a directory: %s", path)
	}
	return filepath.Clean(path), nil
}

func verifyLockedPlan(dir string) (string, string, error) {
	root := filepath.Clean(dir)
	verification, err := lock.Verify(root)
	if err != nil {
		return "", "", err
	}
	if len(verification.Mismatches) > 0 {
		return "", "", fmt.Errorf("BLOCKED: lock hash mismatch for %s", verification.Mismatches[0].Path)
	}
	sourceHash, err := fileSHA256(filepath.Join(root, ".ni", "plan.lock.json"))
	if err != nil {
		return "", "", err
	}
	return root, sourceHash, nil
}

func dependencyPackets(dependencies []string, capabilityToPacket map[string]string) []string {
	packets := make([]string, 0, len(dependencies))
	for _, dependency := range dependencies {
		if packet, ok := capabilityToPacket[dependency]; ok {
			packets = append(packets, packet)
		}
	}
	return packets
}

func buildValidations(items []contract.Evaluation) []Validation {
	validations := make([]Validation, 0, len(items))
	for _, item := range items {
		validations = append(validations, Validation{ID: item.ID, Method: item.Method, Required: true})
	}
	return validations
}

func buildEvidence(items []contract.Artifact) []EvidenceLocation {
	evidence := make([]EvidenceLocation, 0, len(items)+1)
	for i, item := range items {
		evidence = append(evidence, EvidenceLocation{
			ID:   fmt.Sprintf("EVID-%03d", i+1),
			Kind: item.Kind,
			Path: item.Path,
		})
	}
	evidence = append(evidence, EvidenceLocation{
		ID:   fmt.Sprintf("EVID-%03d", len(evidence)+1),
		Kind: "command_output",
		Path: "validation command output",
	})
	return evidence
}

func buildRisks(items []contract.Risk) []KnownRisk {
	risks := make([]KnownRisk, 0, len(items))
	for _, item := range items {
		risks = append(risks, KnownRisk{ID: item.ID, Severity: item.Severity, Mitigation: item.Mitigation})
	}
	return risks
}

func buildNonGoals(items []contract.NonGoal) []NonGoal {
	nonGoals := make([]NonGoal, 0, len(items))
	for _, item := range items {
		nonGoals = append(nonGoals, NonGoal{ID: item.ID, Title: item.Title})
	}
	return nonGoals
}

func fileSHA256(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}
