# Target: namba-ai

`ni export --target namba-ai --out <dir>` writes a seed package for downstream namba-ai-oriented planning after the NI plan is locked.

The export is derived material. The NI kernel remains authoritative for readiness, lock hashes, and source-of-truth order.

## Preconditions

The command must:

- require `.ni/plan.lock.json`,
- verify locked hashes before reading the plan,
- stop with `BLOCKED` on any lock mismatch,
- read `.ni/contract.json` and locked `docs/plan/**` only as locked planning context,
- avoid calling external binaries.

## Seed files

The namba-ai target may write these Markdown files:

```text
planning.md
ni-lock-summary.md
capability-map.md
evaluation-map.md
risk-map.md
suggested-spec-boundaries.md
```

These files are handoff seed material. They are not NI-owned runtime state.

## Suggested spec boundaries

`suggested-spec-boundaries.md` is a proposal, not a required sequential SPEC chain.

It should use graph and dependency language:

```text
boundary candidate
depends_on
acceptance_refs
validation_refs
risk_refs
artifact_refs
```

A downstream tool may split, merge, or reorder boundaries when dependency edges and locked acceptance criteria are preserved.

## Forbidden behavior

The namba-ai export must not:

- call `namba`,
- implement `namba run`,
- implement `namba sync`,
- implement `namba pr`,
- implement `namba land`,
- create mandatory sequential SPEC execution order,
- add Codex-only assumptions.

## Boundary

This command does not run namba-ai, implement namba-ai lifecycle behavior, or add an execution loop to NI core.

If a future target needs richer downstream material, it should remain generated seed or handoff content derived from `.ni/plan.lock.json`. Any mutable runtime state belongs outside the NI kernel.
