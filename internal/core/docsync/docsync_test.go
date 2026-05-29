package docsync

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"ni/internal/core/contract"
)

func TestCheckFixtures(t *testing.T) {
	for _, tt := range []struct {
		name      string
		wantCount int
		wantText  string
	}{
		{"consistent", 0, ""},
		{"docs_cap_missing_contract", 1, "CAP-999"},
		{"contract_cap_missing_docs", 1, "accepted capability CAP-002"},
		{"accepted_capability_lacks_docs_explanation", 1, "no explanatory body"},
		{"decision_conflicts_contract", 1, "DEC-001 status"},
		{"risk_lacks_docs_explanation", 1, "lacks `Mitigation: ...`"},
		{"decision_log_contradicts_contract", 1, "contradicts the contract decision title"},
		{"resolved_open_question_still_blocker", 1, "still shown as a blocker"},
		{"evaluation_references_missing_capability", 1, "references missing capability CAP-999"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			root := filepath.Join("testdata", tt.name)
			c := loadFixtureContract(t, root)

			findings := Check(root, c)
			if len(findings) != tt.wantCount {
				t.Fatalf("expected %d finding(s), got %#v", tt.wantCount, findings)
			}
			if tt.wantText == "" {
				return
			}
			if !strings.Contains(findings[0].Message, tt.wantText) {
				t.Fatalf("expected finding containing %q, got %#v", tt.wantText, findings)
			}
		})
	}
}

func TestDiagnosticFields(t *testing.T) {
	root := filepath.Join("testdata", "docs_cap_missing_contract")
	c := loadFixtureContract(t, root)

	findings := Check(root, c)
	if len(findings) != 1 {
		t.Fatalf("expected one finding, got %#v", findings)
	}
	diagnostic := findings[0].Diagnostic
	if diagnostic.ID != "CAP-999" {
		t.Fatalf("expected diagnostic id CAP-999, got %#v", diagnostic)
	}
	if diagnostic.Location == "" || diagnostic.Problem == "" || diagnostic.WhyItMatters == "" || diagnostic.SuggestedRepair == "" {
		t.Fatalf("expected stable diagnostic fields, got %#v", diagnostic)
	}
	if !diagnostic.BlocksEnd {
		t.Fatalf("expected sync diagnostic to block ni-end, got %#v", diagnostic)
	}
}

func TestFirstRunSyncDiagnostics(t *testing.T) {
	for _, tt := range []struct {
		name         string
		mutate       func(*firstRunFixture)
		wantID       string
		wantProblem  string
		wantLocation string
	}{
		{
			name: "purpose_docs_missing_contract",
			mutate: func(f *firstRunFixture) {
				f.Contract.Project.Purpose = "TODO"
				f.ProjectBrief = projectBrief("conversation_product", []string{"conversation", "document"}, "Plan a support-agent refund triage assistant.")
			},
			wantID:       "SYNC-014",
			wantProblem:  "Project purpose is documented but missing from .ni/contract.json.",
			wantLocation: "docs/plan/00_project_brief.md",
		},
		{
			name: "purpose_contract_missing_docs",
			mutate: func(f *firstRunFixture) {
				f.ProjectBrief = projectBrief("conversation_product", []string{"conversation", "document"}, "TODO")
			},
			wantID:       "SYNC-014",
			wantProblem:  "Project purpose is recorded in .ni/contract.json but not explained in docs.",
			wantLocation: ".ni/contract.json:project.purpose",
		},
		{
			name: "actors_docs_missing_contract",
			mutate: func(f *firstRunFixture) {
				f.Contract.Project.Purpose = "TODO"
				f.Contract.Capabilities[0].Title = "TODO"
				f.Contract.Requirements[0].Title = "TODO"
				f.Contract.Decisions[0].Title = "TODO"
				f.ProjectBrief = projectBrief("conversation_product", []string{"conversation", "document"}, "TODO")
				f.Actors = actorsDoc("Support agent", "gets a draft refund recommendation")
			},
			wantID:       "SYNC-015",
			wantProblem:  "Actors or outcomes are documented but missing from .ni/contract.json.",
			wantLocation: "docs/plan/01_actors_outcomes.md",
		},
		{
			name: "actors_contract_missing_docs",
			mutate: func(f *firstRunFixture) {
				f.Actors = actorsDoc("TODO", "TODO")
			},
			wantID:       "SYNC-015",
			wantProblem:  "Actors or outcomes are recorded in .ni/contract.json but not explained in docs.",
			wantLocation: "docs/plan/01_actors_outcomes.md",
		},
		{
			name: "delivery_docs_missing_contract",
			mutate: func(f *firstRunFixture) {
				f.Contract.ProductType = "software"
				f.Contract.DeliverySurfaces = []string{"cli"}
				f.ProjectBrief = projectBrief("conversation_product", []string{"conversation", "document"}, f.Contract.Project.Purpose)
				f.Delivery = deliveryDoc([]string{"conversation", "document"})
			},
			wantID:       "SYNC-016",
			wantProblem:  "Delivery surface differs between docs and contract.",
			wantLocation: ".ni/contract.json:delivery_surfaces",
		},
		{
			name: "delivery_contract_missing_docs",
			mutate: func(f *firstRunFixture) {
				f.ProjectBrief = projectBrief("conversation_product", nil, f.Contract.Project.Purpose)
				f.Delivery = deliveryDoc(nil)
			},
			wantID:       "SYNC-016",
			wantProblem:  "Delivery surface is recorded in .ni/contract.json but not explained in docs.",
			wantLocation: "docs/plan/08_delivery_operation.md",
		},
		{
			name: "open_question_resolved_in_docs_still_blocker_contract",
			mutate: func(f *firstRunFixture) {
				f.Contract.OpenQuestions = []contract.OpenQuestion{
					{ID: "OQ-001", Title: "Which policy source is authoritative?", Blocker: true, Status: "open"},
				}
				f.OpenQuestions = "# Open questions\n\n## OQ-001: Which policy source is authoritative?\n\nBlocker: false\n\nStatus: resolved\n"
			},
			wantID:       "OQ-001",
			wantProblem:  "open question OQ-001 is resolved in docs but still marked as a blocker in .ni/contract.json.",
			wantLocation: "docs/plan/10_open_questions.md",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			root := writeFirstRunFixture(t, tt.mutate)
			c := loadFixtureContract(t, root)

			findings := Check(root, c)
			finding := requireFinding(t, findings, tt.wantID)
			if finding.Diagnostic.Problem != tt.wantProblem {
				t.Fatalf("expected problem %q, got %#v", tt.wantProblem, finding.Diagnostic)
			}
			if !strings.Contains(finding.Diagnostic.Location, tt.wantLocation) {
				t.Fatalf("expected location containing %q, got %#v", tt.wantLocation, finding.Diagnostic)
			}
			if finding.Diagnostic.WhyItMatters == "" || finding.Diagnostic.SuggestedRepair == "" || !finding.Diagnostic.BlocksEnd {
				t.Fatalf("expected actionable blocking diagnostic, got %#v", finding.Diagnostic)
			}
		})
	}
}

func TestFirstRunConsistentDocsContractPasses(t *testing.T) {
	root := writeFirstRunFixture(t, nil)
	c := loadFixtureContract(t, root)

	findings := Check(root, c)
	if len(findings) != 0 {
		t.Fatalf("expected no findings, got %#v", findings)
	}
}

func TestCheckDoesNotModifyFiles(t *testing.T) {
	root := filepath.Join("testdata", "decision_conflicts_contract")
	c := loadFixtureContract(t, root)
	path := filepath.Join(root, DecisionDoc)
	before, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading fixture doc: %v", err)
	}

	_ = Check(root, c)

	after, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading fixture doc after check: %v", err)
	}
	if string(after) != string(before) {
		t.Fatalf("sync check modified %s", path)
	}
}

type firstRunFixture struct {
	Contract      contract.Contract
	ProjectBrief  string
	Actors        string
	Delivery      string
	OpenQuestions string
}

func writeFirstRunFixture(t *testing.T, mutate func(*firstRunFixture)) string {
	t.Helper()

	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, ".ni"), 0o755); err != nil {
		t.Fatalf("creating .ni: %v", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "docs", "plan"), 0o755); err != nil {
		t.Fatalf("creating docs/plan: %v", err)
	}

	fixture := firstRunFixture{
		Contract: contract.Contract{
			Schema:           contract.Schema,
			ReadinessProfile: "prototype",
			ProductType:      "conversation_product",
			DeliverySurfaces: []string{"conversation", "document"},
			InteractionMode:  "human_to_system",
			Project: contract.Project{
				ID:      "first-run-sync",
				Name:    "First Run Sync",
				Purpose: "Plan a support-agent refund triage assistant that drafts recommendations.",
				Status:  "draft",
			},
			NonGoals: []contract.NonGoal{
				{ID: "NG-001", Title: "Do not issue refunds."},
			},
			Capabilities: []contract.Capability{
				{
					ID:           "CAP-001",
					Title:        "Draft refund recommendations for support agents.",
					Status:       "accepted",
					Requirements: []string{"REQ-001"},
					Evaluations:  []string{"EVAL-001"},
					Risks:        []string{},
					Artifacts:    []string{"ART-001"},
				},
			},
			Requirements: []contract.Requirement{
				{ID: "REQ-001", Title: "Support agents receive a draft recommendation.", Status: "accepted"},
			},
			Decisions: []contract.Decision{
				{ID: "DEC-001", Title: "Deliver as a conversation product with a document handoff.", Status: "accepted"},
			},
			Risks: []contract.Risk{},
			Evaluations: []contract.Evaluation{
				{ID: "EVAL-001", Title: "Transcript fixture review", Method: "review refund triage transcripts"},
			},
			Artifacts: []contract.Artifact{
				{ID: "ART-001", Path: "docs/plan/02_capabilities.md", Kind: "capability_plan"},
			},
			OpenQuestions: []contract.OpenQuestion{},
		},
	}
	fixture.ProjectBrief = projectBrief(fixture.Contract.ProductType, fixture.Contract.DeliverySurfaces, fixture.Contract.Project.Purpose)
	fixture.Actors = actorsDoc("Support agent", "gets a draft refund recommendation")
	fixture.Delivery = deliveryDoc(fixture.Contract.DeliverySurfaces)
	fixture.OpenQuestions = "# Open questions\n\nNo open questions are listed.\n"

	if mutate != nil {
		mutate(&fixture)
	}
	data, err := json.MarshalIndent(fixture.Contract, "", "  ")
	if err != nil {
		t.Fatalf("marshaling fixture contract: %v", err)
	}
	writeFixtureFile(t, root, ".ni/contract.json", append(data, '\n'))
	writeFixtureFile(t, root, ProjectBriefDoc, []byte(fixture.ProjectBrief))
	writeFixtureFile(t, root, ActorsDoc, []byte(fixture.Actors))
	writeFixtureFile(t, root, DeliveryDoc, []byte(fixture.Delivery))
	writeFixtureFile(t, root, CapabilityDoc, []byte("# Capabilities\n\n## CAP-001: Draft refund recommendations for support agents.\n\nSupport agents receive draft refund recommendations.\n"))
	writeFixtureFile(t, root, RiskDoc, []byte("# Risks and security\n\nNo accepted risks are listed.\n"))
	writeFixtureFile(t, root, EvaluationDoc, []byte("# Evaluation contract\n\n## EVAL-001: Transcript fixture review\n\nMethod: review refund triage transcripts\n"))
	writeFixtureFile(t, root, DecisionDoc, []byte("# Decision log\n\n## DEC-001: Deliver as a conversation product with a document handoff.\n\nStatus: accepted\n"))
	writeFixtureFile(t, root, OpenQuestionDoc, []byte(fixture.OpenQuestions))
	return root
}

func projectBrief(productType string, surfaces []string, purpose string) string {
	return "# Project brief\n\n## Product type\n\n" + valueOrTODO(productType) + "\n\n## Delivery surfaces\n\n" + bulletListOrTODO(surfaces) + "\n\n## Purpose\n\n" + valueOrTODO(purpose) + "\n"
}

func actorsDoc(actor string, outcome string) string {
	return "# Actors and outcomes\n\n## Actors\n\n- " + valueOrTODO(actor) + "\n\n## Outcomes\n\n- " + valueOrTODO(outcome) + "\n"
}

func deliveryDoc(surfaces []string) string {
	return "# Delivery and operation\n\n## Delivery surfaces\n\n" + bulletListOrTODO(surfaces) + "\n\n## Initial delivery\n\nPlanning docs and contract records are reviewed before lock.\n"
}

func valueOrTODO(value string) string {
	if strings.TrimSpace(value) == "" {
		return "TODO"
	}
	return value
}

func bulletListOrTODO(values []string) string {
	if len(values) == 0 {
		return "- TODO"
	}
	return "- " + strings.Join(values, "\n- ")
}

func writeFixtureFile(t *testing.T, root string, relPath string, data []byte) {
	t.Helper()

	path := filepath.Join(root, relPath)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("creating fixture directory: %v", err)
	}
	if err := os.WriteFile(path, data, 0o644); err != nil {
		t.Fatalf("writing fixture file %s: %v", relPath, err)
	}
}

func requireFinding(t *testing.T, findings []Finding, id string) Finding {
	t.Helper()

	for _, finding := range findings {
		if finding.Diagnostic.ID == id {
			return finding
		}
	}
	t.Fatalf("expected finding %s in %#v", id, findings)
	return Finding{}
}

func loadFixtureContract(t *testing.T, root string) contract.Contract {
	t.Helper()

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("loading fixture contract: %v", err)
	}
	return c
}
