package pressure

import "testing"

func TestLoadMissingLedgerReturnsEmptySchema(t *testing.T) {
	ledger, err := Load(t.TempDir())
	if err != nil {
		t.Fatalf("load pressure ledger: %v", err)
	}
	if ledger.Schema != Schema {
		t.Fatalf("expected schema %q, got %q", Schema, ledger.Schema)
	}
	if len(ledger.Items) != 0 {
		t.Fatalf("expected no pressure items, got %#v", ledger.Items)
	}
	if got := FormatText(ledger); got != "no pressure recorded\n" {
		t.Fatalf("expected empty text output, got %q", got)
	}
}

func TestAddObservedPressureItems(t *testing.T) {
	dir := t.TempDir()
	ledger, err := AddObserved(dir, []Item{
		{
			Kind:                KindValidationGap,
			EvidenceRefs:        []string{"feedback:codex:test"},
			RelatedCapabilities: []string{"CAP-001"},
			ProposedAction:      "Add a CLI smoke test.",
		},
	})
	if err != nil {
		t.Fatalf("add observed pressure: %v", err)
	}
	if len(ledger.Items) != 1 {
		t.Fatalf("expected one pressure item, got %#v", ledger.Items)
	}
	item := ledger.Items[0]
	if item.ID != "P-001" {
		t.Fatalf("expected first pressure id P-001, got %q", item.ID)
	}
	if item.Status != StatusObserved {
		t.Fatalf("expected observed status, got %q", item.Status)
	}
	if !item.RequiresUserAcceptance {
		t.Fatalf("expected pressure item to require user acceptance")
	}
}

func TestPromoteAdvancesOnlyOneExplicitStep(t *testing.T) {
	dir := t.TempDir()
	if _, err := AddObserved(dir, []Item{{Kind: KindPlanningGap, ProposedAction: "Clarify plan."}}); err != nil {
		t.Fatalf("add observed pressure: %v", err)
	}

	item, err := Promote(dir, "P-001")
	if err != nil {
		t.Fatalf("promote pressure: %v", err)
	}
	if item.Status != StatusRepeated {
		t.Fatalf("expected first promote to repeated, got %q", item.Status)
	}
	item, err = Promote(dir, "P-001")
	if err != nil {
		t.Fatalf("promote pressure second time: %v", err)
	}
	if item.Status != StatusPromotable {
		t.Fatalf("expected second promote to promotable, got %q", item.Status)
	}
	item, err = Promote(dir, "P-001")
	if err != nil {
		t.Fatalf("promote pressure third time: %v", err)
	}
	if item.Status != StatusAccepted {
		t.Fatalf("expected third promote to accepted, got %q", item.Status)
	}
}

func TestRetirePressureItem(t *testing.T) {
	dir := t.TempDir()
	if _, err := AddObserved(dir, []Item{{Kind: KindHarnessCandidate, ProposedAction: "Consider reusable harness seed."}}); err != nil {
		t.Fatalf("add observed pressure: %v", err)
	}

	item, err := Retire(dir, "P-001")
	if err != nil {
		t.Fatalf("retire pressure: %v", err)
	}
	if item.Status != StatusRetired {
		t.Fatalf("expected retired status, got %q", item.Status)
	}
	if _, err := Promote(dir, "P-001"); err == nil {
		t.Fatalf("expected retired pressure promotion to fail")
	}
}
