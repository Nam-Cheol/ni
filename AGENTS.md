# AGENTS.md

## Critical instruction

Do not implement `ni` as a task runner first.

The initial product is `ni-kernel`: a Project Intent Compiler for AI Agents
that creates, validates, locks, and compiles planning contracts before any
execution harness runs.

Its core mechanism is the Intent Lock Protocol: a deterministic pre-runtime
control layer that defines how planning conversations become a project
contract, when the contract is ready to lock, how the accepted plan is hashed,
what downstream actors may trust, and when execution must stop because intent
changed.

## Product architecture

Use this boundary throughout the repository:

```text
ni-kernel
  docs contract
  readiness gate
  lockfile
  prompt compiler
  source-of-truth rule

ni-downstream-seeds
  project-specific work graph
  project-specific evaluation plan
  project-specific evidence rules
  project-specific adapter notes
```

The kernel is authoritative. Downstream seed material is derived and mutable.

## Explicit boundary rules

- `ni` is not a project growth runtime.
- `ni` is not a SPEC runner.
- `ni` is not a multi-agent execution layer.
- `ni` must not copy Hyper Run `run` or `complete` behavior into core.
- `ni` may generate downstream-compatible seed material only after `.ni/plan.lock.json` exists and locked hashes are valid.
- Downstream seed material may include prompts, work-graph proposals,
  evaluation-plan proposals, evidence-rule notes, harness seed proposals, or
  handoff packets. It must not become kernel-owned execution state.

## Authority rules

1. Skills are UX; the CLI is authority.
2. A model may draft docs, detect gaps, propose work graphs, and propose downstream seed material.
3. A model may not declare readiness without `ni status`.
4. A model may not lock a plan without `ni end`.
5. A model may not weaken acceptance criteria to pass validation.
6. A model may not silently edit locked planning docs.
7. If a lock hash mismatch exists, stop and report `BLOCKED`.

## Conversation authoring rules

`ni init` creates the workspace. After initialization, the primary authoring UX
is sustained model-user conversation through docs and skills, not user-entered
contract editing commands.

When authoring, models must:

- extract purpose, actors, capabilities, requirements, decisions, risks,
  evaluations, non-goals, constraints, artifacts, and open questions from the
  conversation;
- update `docs/plan/**` and `.ni/contract.json` together when a turn changes
  planning state;
- treat tentative, inferred, conflicting, or incomplete statements as draft
  records, assumptions, or open questions instead of accepted decisions;
- preserve stable contract IDs and trace capability links to requirements,
  evaluations, risks, and artifacts;
- use `ni status`, `ni end`, and `ni run` as authoritative gates and compilers;
- never declare docs complete from model judgment alone.

Do not add user-facing contract `add`, `list`, or `set` commands for early
authoring. Users should not need to manually edit `.ni/contract.json`; models
and skills maintain it from conversation while the CLI validates the result.

## Model edit discipline

When a model updates planning state, it must keep the diff minimal, visible,
and tied to the current conversation turn.

- A model may propose a doc update from conversation, but ambiguous, tentative,
  inferred, conflicting, or incomplete statements must remain assumptions,
  draft records, or open questions until the user confirms them.
- A model must not convert ambiguous user statements into accepted decisions.
- A model must not weaken risks, mitigations, requirements, evaluations, or
  non-goals to reach readiness.
- A model must not silently delete planning records. It should mark records
  rejected, deferred, resolved, or not applicable when preserving history
  matters.
- A model must not edit locked planning docs except through the amendment or
  relock flow.
- A model should show a short change summary after updating docs, naming the
  changed files, affected records, and remaining blockers or assumptions.

See `docs/37_MODEL_EDIT_SAFETY.md` for examples of good and bad planning
updates.

## Source-of-truth precedence

After `.ni/plan.lock.json` exists, use this order:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

## Initial command roadmap

Implement in this order:

```text
ni --help
ni version
ni init
ni status
ni end
ni run
```

`ni run` must initially compile a prompt only. Do not make it execute Codex or shell commands.

## Development rules

- Keep changes small.
- One prompt should map to one coherent commit.
- Prefer deterministic validation over model judgment.
- Every capability should map to at least one evaluation.
- High-severity risks require mitigation.
- Open blocker questions must prevent locking.
- Prompt output from `ni run` must be 4000 characters or less.

## Validation expectations

When Go code exists, run:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

Before Go code exists, run:

```bash
bash scripts/quality.sh
```

## Forbidden early work

Do not add these before `ni status`, `ni end`, and `ni run` work:

- shell adapter,
- Codex adapter,
- evidence runner,
- queue,
- PR automation,
- release automation,
- plugin system,
- TUI or web UI.
