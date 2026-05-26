package graph

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"ni/internal/core/contract"
	"ni/internal/core/lock"
)

type Proposal struct {
	Project string `json:"project"`
	Nodes   []Node `json:"nodes"`
	Edges   []Edge `json:"edges"`
}

type Node struct {
	ID    string `json:"id"`
	Kind  string `json:"kind"`
	Label string `json:"label"`
}

type Edge struct {
	From string `json:"from"`
	To   string `json:"to"`
	Kind string `json:"kind"`
}

func Propose(dir string) (Proposal, error) {
	root := filepath.Clean(dir)
	if err := verifyLockIfPresent(root); err != nil {
		return Proposal{}, err
	}

	c, err := contract.LoadFile(filepath.Join(root, ".ni", "contract.json"))
	if err != nil {
		return Proposal{}, err
	}

	nodes := make([]Node, 0, len(c.Capabilities)+len(c.Artifacts))
	edges := make([]Edge, 0)
	artifactIDs := make(map[string]bool, len(c.Artifacts))

	for _, artifact := range c.Artifacts {
		artifactIDs[artifact.ID] = true
		nodes = append(nodes, Node{ID: artifact.ID, Kind: "artifact", Label: artifact.Path})
	}
	for _, capability := range c.Capabilities {
		nodes = append(nodes, Node{ID: capability.ID, Kind: "capability", Label: capability.Title})
		for _, dependency := range capability.Dependencies {
			edges = append(edges, Edge{From: dependency, To: capability.ID, Kind: "depends_on"})
		}
		for _, artifact := range capability.Artifacts {
			if artifactIDs[artifact] {
				edges = append(edges, Edge{From: capability.ID, To: artifact, Kind: "produces"})
			}
		}
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].Kind == nodes[j].Kind {
			return nodes[i].ID < nodes[j].ID
		}
		return nodes[i].Kind < nodes[j].Kind
	})
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].From == edges[j].From {
			if edges[i].To == edges[j].To {
				return edges[i].Kind < edges[j].Kind
			}
			return edges[i].To < edges[j].To
		}
		return edges[i].From < edges[j].From
	})

	return Proposal{Project: c.Project.ID, Nodes: nodes, Edges: edges}, nil
}

func FormatText(p Proposal) string {
	var b strings.Builder
	fmt.Fprintf(&b, "work graph proposal for %s\n\n", p.Project)
	b.WriteString("nodes:\n")
	for _, node := range p.Nodes {
		fmt.Fprintf(&b, "- %s [%s] %s\n", node.ID, node.Kind, node.Label)
	}
	b.WriteString("\nedges:\n")
	if len(p.Edges) == 0 {
		b.WriteString("- none\n")
	} else {
		for _, edge := range p.Edges {
			fmt.Fprintf(&b, "- %s -> %s [%s]\n", edge.From, edge.To, edge.Kind)
		}
	}
	return b.String()
}

func verifyLockIfPresent(root string) error {
	lockPath := filepath.Join(root, ".ni", "plan.lock.json")
	if _, err := os.Stat(lockPath); err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	verification, err := lock.Verify(root)
	if err != nil {
		return err
	}
	if len(verification.Mismatches) > 0 {
		return fmt.Errorf("BLOCKED: lock hash mismatch for %s", verification.Mismatches[0].Path)
	}
	return nil
}
