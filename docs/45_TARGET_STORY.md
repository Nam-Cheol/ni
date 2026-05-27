# Target Story

`ni` works upstream of downstream humans, agents, and harnesses.

It does not replace Codex, Spec Kit, Hyper Run, Ouroboros, or namba-ai. It gives
them a locked intent contract to obey.

Targets are consumption shapes for a locked plan. They are not integrations
that `ni` executes, runtime adapters that `ni` owns, or lifecycle state that
becomes part of `ni-kernel`.

```text
conversation -> docs/plan + .ni/contract.json -> ni status -> ni end -> locked plan -> target prompt or seed material -> downstream use
```

## Kernel rule

The kernel is authoritative for:

- planning docs,
- `.ni/contract.json`,
- readiness validation,
- `.ni/plan.lock.json`,
- lock hash verification,
- bounded prompt compilation,
- seed export generation.

Target outputs are derived from a locked plan. They are mutable downstream
artifacts. They do not become NI kernel execution state.

`ni run` and `ni export` must verify `.ni/plan.lock.json` hashes before writing
or printing target output. If a locked file hash does not match, target
generation stops with `BLOCKED`.

## Target meanings

### generic

The `generic` target is a general downstream implementation prompt. It is for
any person, agent, or tool that needs a concise handoff from the locked intent
contract without host-specific assumptions.

### codex

The `codex` target is a bounded implementation prompt seed for Codex. It gives
Codex locked purpose, requirements, constraints, risks, evaluations, and
non-goals, but it does not call `codex exec` or turn `ni` into a Codex adapter.

### human-team

The `human-team` target is a planning handoff for people. It should help a PM,
developer, researcher, reviewer, or operations team understand what has been
accepted, what remains constrained, and what evidence will matter downstream.

### hyper-run

The `hyper-run` target is seed material for a downstream Hyper Run workflow. It
may describe the locked plan, readiness expectations, evidence requirements,
and first-run focus. It must not create `.hyper/goals` runtime packets or copy
Hyper Run `run` or `complete` behavior into `ni-kernel`.

### namba-ai

The `namba-ai` target is planning seed material and suggested graph boundaries.
It may propose capability, evaluation, risk, and boundary maps derived from the
lock. It must not create mandatory sequential SPEC execution, call namba-ai, or
make Codex-only workflow assumptions kernel state.

### spec-kit

The `spec-kit` target is an upstream intent summary for Spec Kit-oriented work.
It gives Spec Kit locked planning context to consume, but it does not create
Spec Kit-owned workflow state, implement slash commands, or make a SPEC file
the source of truth over the NI lock.

### ouroboros

The `ouroboros` target is upstream intent notes for an Ouroboros-oriented
workflow. It may summarize locked capabilities, constraints, risks, and
evaluations, but it does not create Agent OS execution state or implement an
Ouroboros lifecycle inside `ni`.

## Non-goals

Target support must not:

- call downstream tools,
- add target adapters,
- add execution scripts,
- create `.hyper/goals`,
- create Spec Kit runtime state,
- create Ouroboros Agent OS execution state,
- create namba-ai sequential SPEC execution,
- turn downstream seed material into kernel-owned execution state.

## Practical rule

If the output tells another actor what locked intent to obey, it can be a target
shape. If it starts, schedules, tracks, or completes downstream work, it belongs
outside `ni-kernel`.
