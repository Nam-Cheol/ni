package harness

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"ni/internal/core/docstore"
	"ni/internal/core/lock"
	"ni/internal/core/pressure"
)

func TestPlanBuildsGeneratedHarnessProposal(t *testing.T) {
	dir := lockedHarnessProject(t)

	proposal, err := Plan(dir)
	if err != nil {
		t.Fatalf("Plan returned error: %v", err)
	}
	if proposal.Schema != Schema {
		t.Fatalf("expected schema %q, got %q", Schema, proposal.Schema)
	}
	if !strings.HasPrefix(proposal.SourceLockHash, "sha256:") {
		t.Fatalf("expected source lock hash, got %q", proposal.SourceLockHash)
	}
	if len(proposal.SelectedCapabilities) != 2 {
		t.Fatalf("expected selected capabilities, got %#v", proposal.SelectedCapabilities)
	}
	if len(proposal.WorkPackets) != 2 {
		t.Fatalf("expected work packets, got %#v", proposal.WorkPackets)
	}
	if len(proposal.Validations) == 0 {
		t.Fatal("expected validations")
	}
	if len(proposal.EvidenceLocations) == 0 {
		t.Fatal("expected evidence locations")
	}
	if len(proposal.KnownRisks) == 0 {
		t.Fatal("expected known risks")
	}
	if len(proposal.NonGoals) == 0 {
		t.Fatal("expected non-goals")
	}
}

func TestPlanRequiresLockfile(t *testing.T) {
	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), []byte(harnessContract), 0o644); err != nil {
		t.Fatalf("writing harness contract: %v", err)
	}
	writeHarnessDocs(t, dir)

	_, err := Plan(dir)
	if err == nil {
		t.Fatal("expected missing lockfile error")
	}
	if !strings.Contains(err.Error(), "missing lockfile") {
		t.Fatalf("expected missing lockfile error, got %v", err)
	}
}

func TestPlanRefusesStaleLock(t *testing.T) {
	dir := lockedHarnessProject(t)
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("changed\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	_, err := Plan(dir)
	if err == nil {
		t.Fatal("expected stale lock error")
	}
	if !strings.Contains(err.Error(), "BLOCKED: lock hash mismatch") {
		t.Fatalf("expected lock hash mismatch, got %v", err)
	}
}

func TestFormatText(t *testing.T) {
	dir := lockedHarnessProject(t)
	proposal, err := Plan(dir)
	if err != nil {
		t.Fatalf("Plan returned error: %v", err)
	}

	text := FormatText(proposal)
	if !strings.Contains(text, "generated harness proposal") {
		t.Fatalf("missing heading: %q", text)
	}
	if !strings.Contains(text, "WP-001") {
		t.Fatalf("missing work packet: %q", text)
	}
}

func TestCandidateLifecycleRequiresValidationBeforeUserAcceptance(t *testing.T) {
	dir := lockedHarnessProject(t)
	writeAcceptedHarnessPressure(t, dir)

	candidate, err := ProposeFromPressure(dir, "P-001")
	if err != nil {
		t.Fatalf("ProposeFromPressure returned error: %v", err)
	}
	if candidate.Status != StatusProposed {
		t.Fatalf("expected proposed status, got %q", candidate.Status)
	}
	if candidate.ExecutesInsideKernel {
		t.Fatalf("candidate must not execute inside ni core")
	}

	if _, err := AcceptCandidate(dir, candidate.ID); err == nil {
		t.Fatal("expected accepting an unvalidated candidate to fail")
	}

	evidencePath := filepath.Join(dir, "evidence.txt")
	if err := os.WriteFile(evidencePath, []byte("validated by test\n"), 0o644); err != nil {
		t.Fatalf("writing evidence: %v", err)
	}
	candidate, err = ValidateCandidate(dir, candidate.ID, "evidence.txt")
	if err != nil {
		t.Fatalf("ValidateCandidate returned error: %v", err)
	}
	if candidate.Status != StatusValidatedCandidate {
		t.Fatalf("expected validated candidate status, got %q", candidate.Status)
	}
	if candidate.ValidationEvidencePath != "evidence.txt" {
		t.Fatalf("expected evidence path, got %q", candidate.ValidationEvidencePath)
	}

	candidate, err = AcceptCandidate(dir, candidate.ID)
	if err != nil {
		t.Fatalf("AcceptCandidate returned error: %v", err)
	}
	if candidate.Status != StatusUserAccepted {
		t.Fatalf("expected user accepted status, got %q", candidate.Status)
	}

	ledger, err := Candidates(dir)
	if err != nil {
		t.Fatalf("Candidates returned error: %v", err)
	}
	if ledger.ActiveRuleID != candidate.ID {
		t.Fatalf("expected active rule id %q, got %q", candidate.ID, ledger.ActiveRuleID)
	}
}

func TestCandidateCannotSkipUserAcceptance(t *testing.T) {
	dir := lockedHarnessProject(t)
	writeAcceptedHarnessPressure(t, dir)

	candidate, err := ProposeFromPressure(dir, "P-001")
	if err != nil {
		t.Fatalf("ProposeFromPressure returned error: %v", err)
	}
	evidencePath := filepath.Join(dir, "evidence.txt")
	if err := os.WriteFile(evidencePath, []byte("validated by test\n"), 0o644); err != nil {
		t.Fatalf("writing evidence: %v", err)
	}
	candidate, err = ValidateCandidate(dir, candidate.ID, "evidence.txt")
	if err != nil {
		t.Fatalf("ValidateCandidate returned error: %v", err)
	}

	ledger := CandidateLedger{
		Schema:       CandidateLedgerSchema,
		ActiveRuleID: candidate.ID,
		Candidates: []Candidate{
			{
				ID:                        candidate.ID,
				Status:                    StatusValidatedCandidate,
				Target:                    candidate.Target,
				IntendedDownstreamRuntime: candidate.IntendedDownstreamRuntime,
				RequiredEvidence:          candidate.RequiredEvidence,
				Constraints:               candidate.Constraints,
				ForbiddenBehavior:         candidate.ForbiddenBehavior,
				RelatedLockHash:           candidate.RelatedLockHash,
				RelatedPressureIDs:        candidate.RelatedPressureIDs,
				ValidationEvidencePath:    candidate.ValidationEvidencePath,
				RequiresUserAcceptance:    true,
				ExecutesInsideKernel:      false,
			},
		},
	}
	if err := ledger.Validate(); err == nil {
		t.Fatal("expected active rule without user acceptance to fail validation")
	}
}

func lockedHarnessProject(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), []byte(harnessContract), 0o644); err != nil {
		t.Fatalf("writing harness contract: %v", err)
	}
	writeHarnessDocs(t, dir)
	if _, err := lock.CreateAt(dir, time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC)); err != nil {
		t.Fatalf("creating lock: %v", err)
	}
	return dir
}

func writeHarnessDocs(t *testing.T, dir string) {
	t.Helper()
	writeHarnessDoc(t, dir, "00_project_brief.md", "# Project brief\n\n## Product type\n\nsoftware\n\n## Delivery surfaces\n\n- cli\n\n## Purpose\n\nExercise generated harness planning.\n")
	writeHarnessDoc(t, dir, "01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: reviews generated harness proposals.\n- CLI: validates readiness.\n\n## Outcomes\n\n- Exercise generated harness planning.\n")
	writeHarnessDoc(t, dir, "02_capabilities.md", "# Capabilities\n\n## CAP-001: First capability\n\nFirst harness capability.\n\n## CAP-002: Second capability\n\nSecond harness capability.\n")
	writeHarnessDoc(t, dir, "06_risks_security.md", "# Risks and security\n\n## RISK-001: Risk\n\nSeverity: high\n\nMitigation: Keep generated harness read-only.\n")
	writeHarnessDoc(t, dir, "07_evaluation_contract.md", "# Evaluation contract\n\n## EVAL-001: Evaluation one\n\nMethod: go test ./...\n\n## EVAL-002: Evaluation two\n\nMethod: bash scripts/quality.sh\n")
	writeHarnessDoc(t, dir, "08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- cli\n\n## Initial delivery\n\nThe harness fixture is reviewed before lock.\n")
	writeHarnessDoc(t, dir, "10_open_questions.md", "# Open questions\n\nNo open questions are listed in this fixture.\n")
	writeHarnessDoc(t, dir, "11_decision_log.md", "# Decision log\n\n## DEC-001: Decision\n\nStatus: accepted\n")
}

func writeHarnessDoc(t *testing.T, dir string, name string, content string) {
	t.Helper()
	path := filepath.Join(dir, "docs", "plan", name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writing harness doc %s: %v", name, err)
	}
}

func writeAcceptedHarnessPressure(t *testing.T, dir string) {
	t.Helper()

	ledger := pressure.Ledger{
		Schema: pressure.Schema,
		Items: []pressure.Item{
			{
				ID:                     "P-001",
				Kind:                   pressure.KindHarnessCandidate,
				Status:                 pressure.StatusAccepted,
				EvidenceRefs:           []string{"test:evidence"},
				RelatedCapabilities:    []string{"CAP-001"},
				RelatedRisks:           []string{"RISK-001"},
				ProposedAction:         "Create an inert downstream harness proposal for CAP-001.",
				RequiresUserAcceptance: true,
			},
		},
	}
	if err := pressure.Save(dir, ledger); err != nil {
		t.Fatalf("writing accepted harness pressure: %v", err)
	}
}

const harnessContract = `{
  "schema": "ni.contract.v0",
  "readiness_profile": "prototype",
  "project": {
    "id": "harness-fixture",
    "name": "Harness Fixture",
    "purpose": "Exercise generated harness planning.",
    "status": "draft"
  },
  "non_goals": [
    {
      "id": "NG-001",
      "title": "Do not execute the harness."
    }
  ],
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "First capability",
      "status": "accepted",
      "requirements": ["REQ-001"],
      "evaluations": ["EVAL-001"],
      "risks": ["RISK-001"],
      "artifacts": ["ART-001"]
    },
    {
      "id": "CAP-002",
      "title": "Second capability",
      "status": "accepted",
      "dependencies": ["CAP-001"],
      "requirements": ["REQ-002"],
      "evaluations": ["EVAL-002"],
      "risks": [],
      "artifacts": ["ART-002"]
    }
  ],
  "requirements": [
    {"id": "REQ-001", "title": "Requirement one", "status": "accepted"},
    {"id": "REQ-002", "title": "Requirement two", "status": "accepted"}
  ],
  "decisions": [
    {"id": "DEC-001", "title": "Decision", "status": "accepted"}
  ],
  "risks": [
    {
      "id": "RISK-001",
      "title": "Risk",
      "severity": "high",
      "status": "accepted",
      "mitigation": "Keep generated harness read-only."
    }
  ],
  "evaluations": [
    {"id": "EVAL-001", "title": "Evaluation one", "method": "go test ./..."},
    {"id": "EVAL-002", "title": "Evaluation two", "method": "bash scripts/quality.sh"}
  ],
  "artifacts": [
    {"id": "ART-001", "path": "cmd/ni/main.go", "kind": "code"},
    {"id": "ART-002", "path": "internal/core/harness", "kind": "code"}
  ],
  "open_questions": []
}
`
