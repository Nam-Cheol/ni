# Target: Ouroboros

`ni export --target ouroboros --out <dir>` writes seed notes for downstream Ouroboros-oriented recursive planning after the NI plan is locked.

The export is derived material. The NI kernel remains authoritative for readiness, lock hashes, and source-of-truth order.

## Preconditions

The command must:

- require `.ni/plan.lock.json`,
- verify locked hashes before reading the plan,
- stop with `BLOCKED` on any lock mismatch,
- read `.ni/contract.json` and locked `docs/plan/**` only as locked planning context,
- avoid calling external binaries.

## Seed files

The Ouroboros target may write this Markdown file:

```text
ouroboros-seed-notes.md
```

This file is handoff seed material. It is not NI-owned runtime state.

## Seed-note contents

The seed notes should contain only:

- locked contract summary,
- capabilities,
- constraints,
- risks,
- evaluation contract,
- non-goals,
- source-of-truth references.

Ouroboros-specific references must be written as seed context, not as commands or NI-owned lifecycle state.

## Forbidden behavior

The Ouroboros export must not:

- implement interview,
- implement crystallize,
- implement execute,
- implement evaluate,
- implement evolve,
- create an Ouroboros lifecycle in NI core,
- run downstream agents.

## Boundary

This command does not run Ouroboros, implement Ouroboros lifecycle behavior, or add an execution loop to NI core.

If a future target needs richer downstream material, it should remain generated seed or handoff content derived from `.ni/plan.lock.json`. Any mutable runtime state belongs outside the NI kernel.
