# Target: Hyper Run

`ni export --target hyper-run --out <dir>` writes a seed package for a downstream Hyper Run workflow after the NI plan is locked.

The export is derived material. The NI kernel remains authoritative for readiness, lock hashes, and source-of-truth order.

## Preconditions

The command must:

- require `.ni/plan.lock.json`,
- verify locked hashes before reading the plan,
- stop with `BLOCKED` on any lock mismatch,
- read `.ni/contract.json` and locked `docs/plan/**` only as locked planning context,
- avoid calling external binaries.

## Seed files

The Hyper Run target may write these Markdown files:

```text
plan.md
ni-context.md
readiness-expectations.md
evidence-requirements.md
first-run-focus.md
```

These files are handoff seed material. They are not NI-owned runtime state.

## Forbidden output

The Hyper Run export must not create:

```text
.hyper/goals/GOAL-0001/
tasks.md
evidence.md
review.md
next.md
```

Those names are reserved for downstream runtime packets. NI may describe expectations that a downstream runtime should honor, but it must not create or execute those packets.

## Boundary

This command does not run Hyper Run, implement Hyper Run completion behavior, or add an execution loop to NI core.

If a future target needs richer downstream material, it should remain generated seed or handoff content derived from `.ni/plan.lock.json`. Any mutable runtime state belongs outside the NI kernel.
