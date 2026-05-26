package target

import (
	"fmt"
	"strings"
)

const Generic = "generic"

type Target struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Artifact    string `json:"artifact"`
}

var builtins = []Target{
	{
		Name:        Generic,
		Description: "General implementation prompt for any downstream worker.",
		Artifact:    "prompt",
	},
	{
		Name:        "codex",
		Description: "Codex-oriented implementation prompt that remains a seed artifact.",
		Artifact:    "prompt",
	},
	{
		Name:        "human-team",
		Description: "Human-team handoff prompt for planning, implementation, and validation.",
		Artifact:    "handoff",
	},
	{
		Name:        "hyper-run",
		Description: "Hyper Run seed prompt without runtime packets or execution.",
		Artifact:    "seed",
	},
	{
		Name:        "namba-ai",
		Description: "Namba AI seed prompt for downstream planning handoff.",
		Artifact:    "seed",
	},
	{
		Name:        "ouroboros",
		Description: "Ouroboros seed notes for downstream recursive planning handoff.",
		Artifact:    "seed",
	},
	{
		Name:        "spec-kit",
		Description: "Spec Kit seed notes for downstream specification work.",
		Artifact:    "seed",
	},
}

var byName = func() map[string]Target {
	items := make(map[string]Target, len(builtins))
	for _, item := range builtins {
		items[item.Name] = item
	}
	return items
}()

func List() []Target {
	return append([]Target(nil), builtins...)
}

func Lookup(name string) (Target, error) {
	if strings.TrimSpace(name) == "" {
		name = Generic
	}
	if item, ok := byName[name]; ok {
		return item, nil
	}
	return Target{}, fmt.Errorf("unsupported target %q (valid: %s)", name, NamesText())
}

func NamesText() string {
	names := make([]string, 0, len(builtins))
	for _, item := range List() {
		names = append(names, item.Name)
	}
	return strings.Join(names, ", ")
}
