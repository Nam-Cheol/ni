# Target: Spec Kit

`ni export --target spec-kit --out <dir>` writes seed notes for downstream Spec Kit-oriented specification work after the NI plan is locked.

The export is derived material. The NI kernel remains authoritative for readiness, lock hashes, and source-of-truth order.

## Preconditions

The command must:

- require `.ni/plan.lock.json`,
- verify locked hashes before reading the plan,
- stop with `BLOCKED` on any lock mismatch,
- read `.ni/contract.json` and locked `docs/plan/**` only as locked planning context,
- avoid calling external binaries.

## Seed files

The Spec Kit target may write this Markdown file:

```text
spec-kit-seed-notes.md
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

Spec Kit-specific references must be written as seed context, not as commands or NI-owned lifecycle state.

## Forbidden behavior

The Spec Kit export must not:

- implement slash commands,
- create coding task lists as NI core state,
- run downstream tools,
- become a spec-first coding operating system,
- weaken locked acceptance criteria or risk mitigations.

## Boundary

This command does not run Spec Kit, implement Spec Kit command behavior, or add an execution loop to NI core.

If a future target needs richer downstream material, it should remain generated seed or handoff content derived from `.ni/plan.lock.json`. Any mutable runtime state belongs outside the NI kernel.
