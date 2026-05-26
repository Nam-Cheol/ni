# Domain and state model

## Core entities

```text
namba-ai upgrade intent
current limitation
planning contract
capability
requirement
decision
risk
mitigation
evaluation
artifact
downstream target
boundary candidate
dependency edge
lockfile
compiled prompt
```

## Authority order

After `.ni/plan.lock.json` exists:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > generated prompt > chat transcript > model guess
```

Generated prompt files are derived artifacts. They are never the source of truth for planning state.

## Collaboration state

Collaboration changes are contract changes when they alter accepted requirements, decisions, risks, evaluations, non-goals, or artifacts. Parallel planning is valid only when the dependency graph and conflict checks preserve locked acceptance criteria.
