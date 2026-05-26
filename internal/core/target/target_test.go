package target

import (
	"strings"
	"testing"
)

func TestListIncludesBuiltins(t *testing.T) {
	want := []string{"generic", "codex", "human-team", "hyper-run", "namba-ai", "ouroboros", "spec-kit"}
	got := List()
	if len(got) != len(want) {
		t.Fatalf("expected %d targets, got %d: %#v", len(want), len(got), got)
	}
	for i, name := range want {
		if got[i].Name != name {
			t.Fatalf("target %d: expected %q, got %q", i, name, got[i].Name)
		}
	}
}

func TestLookupDefaultsToGeneric(t *testing.T) {
	item, err := Lookup("")
	if err != nil {
		t.Fatalf("Lookup returned error: %v", err)
	}
	if item.Name != Generic {
		t.Fatalf("expected generic target, got %q", item.Name)
	}
}

func TestLookupRejectsUnsupportedTarget(t *testing.T) {
	_, err := Lookup("shell")
	if err == nil {
		t.Fatal("expected unsupported target error")
	}
	if !strings.Contains(err.Error(), `unsupported target "shell"`) {
		t.Fatalf("expected clear target error, got %v", err)
	}
	if !strings.Contains(err.Error(), "codex") || !strings.Contains(err.Error(), "spec-kit") {
		t.Fatalf("expected valid target names in error, got %v", err)
	}
}
