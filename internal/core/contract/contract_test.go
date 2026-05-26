package contract

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadFileValidContract(t *testing.T) {
	c, err := LoadFile(filepath.Join("..", "..", "..", ".ni", "contract.json"))
	if err != nil {
		t.Fatalf("LoadFile returned error: %v", err)
	}
	if c.Schema != Schema {
		t.Fatalf("expected schema %q, got %q", Schema, c.Schema)
	}
	if c.Project.ID != "ni" {
		t.Fatalf("expected project id ni, got %q", c.Project.ID)
	}
	if c.ReadinessProfile != "prototype" {
		t.Fatalf("expected readiness profile prototype, got %q", c.ReadinessProfile)
	}
	if c.ProductType != DefaultProductType {
		t.Fatalf("expected default product type %q, got %q", DefaultProductType, c.ProductType)
	}
	if len(c.DeliverySurfaces) != 1 || c.DeliverySurfaces[0] != DefaultDeliverySurface {
		t.Fatalf("expected default delivery surface %q, got %#v", DefaultDeliverySurface, c.DeliverySurfaces)
	}
}

func TestLoadProductFields(t *testing.T) {
	c, err := Load([]byte(`{
	  "schema": "ni.contract.v0",
	  "readiness_profile": "prototype",
	  "product_type": "conversation_product",
	  "delivery_surfaces": ["conversation"],
	  "interaction_mode": "human_to_human",
	  "project": {
	    "id": "p",
	    "name": "P",
	    "purpose": "test",
	    "status": "draft"
	  },
	  "non_goals": [],
	  "capabilities": [],
	  "requirements": [],
	  "decisions": [],
	  "risks": [],
	  "evaluations": [],
	  "artifacts": [],
	  "open_questions": []
	}`))
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if c.ProductType != "conversation_product" {
		t.Fatalf("expected conversation product, got %q", c.ProductType)
	}
	if len(c.DeliverySurfaces) != 1 || c.DeliverySurfaces[0] != "conversation" {
		t.Fatalf("expected conversation surface, got %#v", c.DeliverySurfaces)
	}
	if c.InteractionMode != "human_to_human" {
		t.Fatalf("expected interaction mode, got %q", c.InteractionMode)
	}
}

func TestLoadDefaultsProductFieldsWhenNoValueSupplied(t *testing.T) {
	c, err := Load([]byte(`{
	  "schema": "ni.contract.v0",
	  "readiness_profile": "prototype",
	  "project": {
	    "id": "p",
	    "name": "P",
	    "purpose": "test",
	    "status": "draft"
	  },
	  "non_goals": [],
	  "capabilities": [],
	  "requirements": [],
	  "decisions": [],
	  "risks": [],
	  "evaluations": [],
	  "artifacts": [],
	  "open_questions": []
	}`))
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if c.ProductType != "software" {
		t.Fatalf("expected default software product type, got %q", c.ProductType)
	}
	if len(c.DeliverySurfaces) != 1 || c.DeliverySurfaces[0] != "cli" {
		t.Fatalf("expected default cli surface, got %#v", c.DeliverySurfaces)
	}
	if c.InteractionMode != "human_to_system" {
		t.Fatalf("expected default interaction mode, got %q", c.InteractionMode)
	}
}

func TestLoadRejectsUnsupportedProductType(t *testing.T) {
	_, err := Load([]byte(`{
	  "schema": "ni.contract.v0",
	  "readiness_profile": "prototype",
	  "product_type": "website_only",
	  "delivery_surfaces": ["web"],
	  "interaction_mode": "human_to_system",
	  "project": {
	    "id": "p",
	    "name": "P",
	    "purpose": "test",
	    "status": "draft"
	  },
	  "non_goals": [],
	  "capabilities": [],
	  "requirements": [],
	  "decisions": [],
	  "risks": [],
	  "evaluations": [],
	  "artifacts": [],
	  "open_questions": []
	}`))
	if err == nil {
		t.Fatal("expected unsupported product type error")
	}
	if !strings.Contains(err.Error(), "unsupported product_type") {
		t.Fatalf("expected product type error, got %v", err)
	}
}

func TestLoadMalformedJSON(t *testing.T) {
	_, err := Load([]byte(`{"schema":`))
	if err == nil {
		t.Fatal("expected malformed JSON error")
	}
	if !strings.Contains(err.Error(), "malformed contract JSON") {
		t.Fatalf("expected malformed JSON context, got %v", err)
	}
}

func TestLoadMissingTopLevelFields(t *testing.T) {
	_, err := Load([]byte(`{"schema":"ni.contract.v0","project":{"id":"p"}}`))
	if err == nil {
		t.Fatal("expected missing field error")
	}
	for _, field := range []string{"readiness_profile", "project.name", "project.purpose", "project.status", "capabilities", "open_questions"} {
		if !strings.Contains(err.Error(), field) {
			t.Fatalf("expected missing field %s in error %q", field, err.Error())
		}
	}
}

func TestUnsupportedSchema(t *testing.T) {
	_, err := Load([]byte(`{
	  "schema": "ni.contract.v9",
	  "readiness_profile": "prototype",
	  "project": {
	    "id": "p",
	    "name": "P",
	    "purpose": "test",
	    "status": "draft"
	  },
	  "non_goals": [],
	  "capabilities": [],
	  "requirements": [],
	  "decisions": [],
	  "risks": [],
	  "evaluations": [],
	  "artifacts": [],
	  "open_questions": []
	}`))
	if err == nil {
		t.Fatal("expected unsupported schema error")
	}
	if !strings.Contains(err.Error(), "unsupported contract schema") {
		t.Fatalf("expected unsupported schema error, got %v", err)
	}
}

func TestValidateIDPrefix(t *testing.T) {
	for _, tt := range []struct {
		id     string
		prefix string
	}{
		{"CAP-001", "CAP"},
		{"REQ-001", "REQ"},
		{"DEC-001", "DEC"},
		{"RISK-001", "RISK"},
		{"EVAL-001", "EVAL"},
		{"ART-001", "ART"},
		{"OQ-001", "OQ"},
	} {
		if err := ValidateIDPrefix(tt.id, tt.prefix); err != nil {
			t.Fatalf("ValidateIDPrefix(%q, %q) returned error: %v", tt.id, tt.prefix, err)
		}
	}
}

func TestValidateIDPrefixRejectsWrongPrefix(t *testing.T) {
	err := ValidateIDPrefix("REQ-001", "CAP")
	if err == nil {
		t.Fatal("expected prefix error")
	}
	if !strings.Contains(err.Error(), "CAP- prefix") {
		t.Fatalf("expected CAP prefix error, got %v", err)
	}
}

func TestLoadRejectsBadCapabilityReferencePrefix(t *testing.T) {
	_, err := Load([]byte(`{
	  "schema": "ni.contract.v0",
	  "readiness_profile": "prototype",
	  "project": {
	    "id": "p",
	    "name": "P",
	    "purpose": "test",
	    "status": "draft"
	  },
	  "non_goals": [],
	  "capabilities": [
	    {
	      "id": "CAP-001",
	      "title": "capability",
	      "status": "accepted",
	      "requirements": ["CAP-002"],
	      "evaluations": [],
	      "risks": [],
	      "artifacts": []
	    }
	  ],
	  "requirements": [],
	  "decisions": [],
	  "risks": [],
	  "evaluations": [],
	  "artifacts": [],
	  "open_questions": []
	}`))
	if err == nil {
		t.Fatal("expected reference prefix error")
	}
	if !strings.Contains(err.Error(), "REQ- prefix") {
		t.Fatalf("expected REQ prefix error, got %v", err)
	}
}

func TestLoadRejectsUnsupportedReadinessProfile(t *testing.T) {
	_, err := Load([]byte(`{
	  "schema": "ni.contract.v0",
	  "readiness_profile": "launch",
	  "project": {
	    "id": "p",
	    "name": "P",
	    "purpose": "test",
	    "status": "draft"
	  },
	  "non_goals": [],
	  "capabilities": [],
	  "requirements": [],
	  "decisions": [],
	  "risks": [],
	  "evaluations": [],
	  "artifacts": [],
	  "open_questions": []
	}`))
	if err == nil {
		t.Fatal("expected readiness profile error")
	}
	if !strings.Contains(err.Error(), "unsupported readiness profile") {
		t.Fatalf("expected readiness profile error, got %v", err)
	}
}

func TestLoadFileReportsMissingProjectField(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "contract.json")
	if err := os.WriteFile(path, []byte(`{"schema":"ni.contract.v0"}`), 0o644); err != nil {
		t.Fatalf("writing test contract: %v", err)
	}

	_, err := LoadFile(path)
	if err == nil {
		t.Fatal("expected missing project field error")
	}
	if !strings.Contains(err.Error(), "project.id") {
		t.Fatalf("expected project.id in error, got %v", err)
	}
}
