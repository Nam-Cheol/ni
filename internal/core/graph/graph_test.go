package graph

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"ni/internal/core/docstore"
	"ni/internal/core/lock"
)

func TestProposeBuildsCapabilityArtifactGraph(t *testing.T) {
	dir := graphProject(t)

	proposal, err := Propose(dir)
	if err != nil {
		t.Fatalf("Propose returned error: %v", err)
	}
	if proposal.Project != "graph-fixture" {
		t.Fatalf("expected graph-fixture project, got %q", proposal.Project)
	}
	requireNode(t, proposal, "CAP-001", "capability")
	requireNode(t, proposal, "ART-001", "artifact")
	requireEdge(t, proposal, "CAP-001", "CAP-002", "depends_on")
	requireEdge(t, proposal, "CAP-001", "ART-001", "produces")
}

func TestFormatText(t *testing.T) {
	dir := graphProject(t)
	proposal, err := Propose(dir)
	if err != nil {
		t.Fatalf("Propose returned error: %v", err)
	}

	text := FormatText(proposal)
	if !strings.Contains(text, "work graph proposal for graph-fixture") {
		t.Fatalf("missing heading: %q", text)
	}
	if !strings.Contains(text, "CAP-001 -> CAP-002 [depends_on]") {
		t.Fatalf("missing dependency edge: %q", text)
	}
}

func TestProposeRefusesStaleLock(t *testing.T) {
	dir := graphProject(t)
	if _, err := lock.CreateAt(dir, time.Date(2026, 5, 26, 1, 2, 3, 0, time.UTC)); err != nil {
		t.Fatalf("creating lock: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, "docs", "plan", "00_project_brief.md"), []byte("changed\n"), 0o644); err != nil {
		t.Fatalf("changing locked doc: %v", err)
	}

	_, err := Propose(dir)
	if err == nil {
		t.Fatal("expected stale lock error")
	}
	if !strings.Contains(err.Error(), "BLOCKED: lock hash mismatch") {
		t.Fatalf("expected lock hash mismatch error, got %v", err)
	}
}

func graphProject(t *testing.T) string {
	t.Helper()

	dir := t.TempDir()
	if _, err := docstore.Init(dir); err != nil {
		t.Fatalf("initializing project: %v", err)
	}
	if err := os.WriteFile(filepath.Join(dir, ".ni", "contract.json"), []byte(graphContract), 0o644); err != nil {
		t.Fatalf("writing graph contract: %v", err)
	}
	writeGraphDoc(t, dir, "00_project_brief.md", "# Project brief\n\n## Product type\n\nsoftware\n\n## Delivery surfaces\n\n- cli\n\n## Purpose\n\nExercise work graph proposal planning.\n")
	writeGraphDoc(t, dir, "01_actors_outcomes.md", "# Actors and outcomes\n\n## Actors\n\n- User: reviews graph proposals.\n- CLI: validates readiness.\n\n## Outcomes\n\n- Exercise work graph proposal planning.\n")
	writeGraphDoc(t, dir, "02_capabilities.md", "# Capabilities\n\n## CAP-001: First capability\n\nFirst graph node.\n\n## CAP-002: Second capability\n\nSecond graph node.\n")
	writeGraphDoc(t, dir, "06_risks_security.md", "# Risks and security\n\nNo accepted risks are listed in this fixture.\n")
	writeGraphDoc(t, dir, "07_evaluation_contract.md", "# Evaluation contract\n\n## EVAL-001: Evaluation\n\nMethod: fixture\n")
	writeGraphDoc(t, dir, "08_delivery_operation.md", "# Delivery and operation\n\n## Delivery surfaces\n\n- cli\n\n## Initial delivery\n\nThe graph fixture is reviewed before lock.\n")
	writeGraphDoc(t, dir, "10_open_questions.md", "# Open questions\n\nNo open questions are listed in this fixture.\n")
	writeGraphDoc(t, dir, "11_decision_log.md", "# Decision log\n\n## DEC-001: Decision\n\nStatus: accepted\n")
	return dir
}

func writeGraphDoc(t *testing.T, dir string, name string, content string) {
	t.Helper()
	path := filepath.Join(dir, "docs", "plan", name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("writing graph doc %s: %v", name, err)
	}
}

func requireNode(t *testing.T, proposal Proposal, id string, kind string) {
	t.Helper()
	for _, node := range proposal.Nodes {
		if node.ID == id && node.Kind == kind {
			return
		}
	}
	t.Fatalf("missing node %s/%s in %#v", id, kind, proposal.Nodes)
}

func requireEdge(t *testing.T, proposal Proposal, from string, to string, kind string) {
	t.Helper()
	for _, edge := range proposal.Edges {
		if edge.From == from && edge.To == to && edge.Kind == kind {
			return
		}
	}
	t.Fatalf("missing edge %s -> %s/%s in %#v", from, to, kind, proposal.Edges)
}

const graphContract = `{
  "schema": "ni.contract.v0",
  "readiness_profile": "prototype",
  "project": {
    "id": "graph-fixture",
    "name": "Graph Fixture",
    "purpose": "Exercise graph proposal.",
    "status": "draft"
  },
  "non_goals": [
    {
      "id": "NG-001",
      "title": "Do not run work packets."
    }
  ],
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "First capability",
      "status": "accepted",
      "requirements": [
        "REQ-001"
      ],
      "evaluations": [
        "EVAL-001"
      ],
      "risks": [],
      "artifacts": [
        "ART-001"
      ]
    },
    {
      "id": "CAP-002",
      "title": "Second capability",
      "status": "accepted",
      "dependencies": [
        "CAP-001"
      ],
      "requirements": [
        "REQ-002"
      ],
      "evaluations": [
        "EVAL-001"
      ],
      "risks": [],
      "artifacts": [
        "ART-002"
      ]
    }
  ],
  "requirements": [
    {
      "id": "REQ-001",
      "title": "Requirement one",
      "status": "accepted"
    },
    {
      "id": "REQ-002",
      "title": "Requirement two",
      "status": "accepted"
    }
  ],
  "decisions": [
    {
      "id": "DEC-001",
      "title": "Decision",
      "status": "accepted"
    }
  ],
  "risks": [],
  "evaluations": [
    {
      "id": "EVAL-001",
      "title": "Evaluation",
      "method": "fixture"
    }
  ],
  "artifacts": [
    {
      "id": "ART-001",
      "path": "docs/plan/",
      "kind": "planning_docs"
    },
    {
      "id": "ART-002",
      "path": ".ni/contract.json",
      "kind": "contract"
    }
  ],
  "open_questions": []
}
`
