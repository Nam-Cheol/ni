# Target Conformance

Target outputs are downstream seed material. They may help Codex, Spec Kit,
Hyper Run, Ouroboros, namba-ai, or a human team start from a locked NI plan, but
they must not create runtime execution state owned by those downstream tools.

NI remains a pre-runtime intent lock layer. The kernel validates readiness,
writes `.ni/plan.lock.json`, verifies hashes, and compiles bounded handoff
artifacts. Execution, queues, slash-command workflows, lifecycle state, and
tool-owned packets belong outside NI.

## Universal rules

Every target output must:

- be derived from `.ni/plan.lock.json`, `.ni/contract.json`, and locked
  `docs/plan/**` content,
- require `.ni/plan.lock.json` before output is produced,
- verify lock hashes before writing output,
- refuse stale locks with `BLOCKED`,
- preserve accepted requirements, risk mitigations, non-goals, blocker handling,
  and source-of-truth order,
- stay inert: prompt, handoff, notes, maps, or seed documents only.

Every target output must not:

- call downstream CLIs or external binaries,
- create downstream runtime packet directories,
- create queues, task-runner state, slash-command state, lifecycle state, PR
  automation, or adapter execution state,
- weaken the locked contract to make a target output appear usable.

## Prompt targets

`codex` output is a bounded prompt only. `ni run --target codex` may print or
write a prompt whose artifact is `prompt`; it must not invoke Codex, create
`.codex/commands`, write queue state, or create execution scripts.

`human-team` output is a bounded handoff prompt only. `ni run --target
human-team` may print or write a handoff artifact for PM/dev/design/research
coordination; it must not create owner databases, team workflow state, queues,
or NI-owned orchestration.

Both prompt targets must remain within the prompt budget, defaulting to 4000
characters.

## Seed export targets

### hyper-run

The Hyper Run export may create these seed files:

```text
plan.md
ni-context.md
readiness-expectations.md
evidence-requirements.md
first-run-focus.md
```

It must not create Hyper Run runtime packet paths, including:

```text
.hyper/
.hyper/goals/
.hyper/goals/GOAL-0001/
tasks.md
evidence.md
review.md
next.md
```

### namba-ai

The namba-ai export may create these seed files:

```text
planning.md
ni-lock-summary.md
capability-map.md
evaluation-map.md
risk-map.md
suggested-spec-boundaries.md
```

`suggested-spec-boundaries.md` must remain graph-oriented seed material. It may
name candidate boundaries and `depends_on` edges, but it must not create a
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

The Ouroboros export may create this seed file:

```text
ouroboros-seed-notes.md
```

It must not create execute, evaluate, or evolve runtime state, including:

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

The Spec Kit export may create this seed file:

```text
spec-kit-seed-notes.md
```

It must not create slash-command workflow state, including:

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

Conformance is enforced in three layers:

- `go test ./...` covers stale lock refusal, seed-only export paths, and
  prompt-only or handoff-only target output.
- `scripts/smoke.sh` runs `scripts/check-target-conformance.py` against smoke
  exports for each seed export target.
- `scripts/demo-check.sh` verifies public demos without invoking downstream
  runtimes.

The checker is intentionally path-oriented. Seed content may evolve, but
runtime-owned packet names and directories must stay outside NI outputs.
