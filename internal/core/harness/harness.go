package harness

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
)

const Schema = "ni.generated_harness.v0"

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
