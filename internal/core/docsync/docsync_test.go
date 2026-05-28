package docsync

import (
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

func loadFixtureContract(t *testing.T, root string) contract.Contract {
	t.Helper()

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("loading fixture contract: %v", err)
	}
	return c
}
