# Target Export Conformance

Target exports are downstream seed material. They must help another tool start
from a locked NI plan without creating runtime-owned packets inside the NI
export directory.

The NI kernel remains authoritative for readiness, lock hashes, and
source-of-truth order. Exported files are derived and mutable handoff material.

## General rules

Every target export must:

- require `.ni/plan.lock.json`,
- verify locked hashes before writing output,
- stop with `BLOCKED` on any lock mismatch,
- write only target seed files,
- avoid calling downstream tools or external binaries,
- avoid creating runtime-owned state directories,
- preserve locked acceptance criteria and risk mitigations.

Every target export must not:

- execute downstream work,
- create queues, task runners, or lifecycle state,
- turn target-specific language into NI-owned workflow state,
- weaken locked planning constraints to make an export appear usable.

## Target rules

### hyper-run

The Hyper Run export may create:

```text
plan.md
ni-context.md
readiness-expectations.md
evidence-requirements.md
first-run-focus.md
```

It must not create Hyper Run runtime packet paths such as:

```text
.hyper/goals/
.hyper/goals/GOAL-0001/
tasks.md
evidence.md
review.md
next.md
```

### namba-ai

The namba-ai export may create:

```text
planning.md
ni-lock-summary.md
capability-map.md
evaluation-map.md
risk-map.md
suggested-spec-boundaries.md
```

`suggested-spec-boundaries.md` must remain graph-oriented seed material. It may
name boundary candidates and `depends_on` edges, but it must not create a
mandatory sequential SPEC execution chain.

Forbidden namba-ai runtime or chain state includes:

```text
.namba/
.namba/specs/
SPEC-001.md
SPEC-002.md
SPEC_SEQUENCE.md
specs/
tasks.md
run.md
sync.md
pr.md
land.md
```

### ouroboros

The Ouroboros export may create:

```text
ouroboros-seed-notes.md
```

It must not create execute, evaluate, or evolve runtime state such as:

```text
.ouroboros/
.ouroboros/runtime/
execute
execute.md
evaluate
evaluate.md
evolve
evolve.md
runtime/
```

### spec-kit

The Spec Kit export may create:

```text
spec-kit-seed-notes.md
```

It must not create slash-command workflow state such as:

```text
.specify/
.specify/specs/
.specify/memory/
.github/prompts/
.claude/commands/
.codex/commands/
slash-commands.md
commands/
specify.md
plan.md
tasks.md
```

## Enforcement

Conformance is checked in two layers:

- `go test ./...` includes CLI integration coverage for each target boundary.
- `scripts/smoke.sh` runs `scripts/check-target-conformance.py` against each
  smoke export directory.

The checker is path-oriented by design. Export content may become more useful,
but runtime-owned packet names and directories must stay outside NI exports.
