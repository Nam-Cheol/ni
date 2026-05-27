# Why ni

`ni` exists because AI agents often start from prompts that feel actionable
before the project intent is explicit, accepted, validated, locked, and
unchanged.

The product claim is small but strict: compile intent before execution starts.

## The Problem

A vague request can hide the details that determine whether downstream work is
trustworthy:

- who the project is for;
- what must be accepted;
- which risks need mitigation;
- what is explicitly out of scope;
- which questions should block execution;
- whether the accepted plan has changed.

Most tools add control after a prompt, spec, worklist, or execution loop
already exists. At that point a downstream agent or team must either stop and
ask questions, silently invent intent, or continue from stale planning context.

`ni` moves the control point earlier. It asks whether project intent has earned
a trustworthy handoff before any downstream actor starts work.

## The Payoff

The output of `ni-kernel` is a locked, versioned, verifiable project contract.
Downstream actors may consume bounded prompts or inert seed material only after
readiness and lock checks pass.

That gives a project four useful properties:

- ambiguous intent blocks before execution;
- accepted intent has a stable lock hash;
- changed intent stops downstream handoff;
- prompts and seed exports stay derived from the locked contract.

## What ni Blocks

`ni` blocks downstream handoff when intent is not trustworthy yet:

- accepted capabilities without linked evaluations;
- high-severity risks without mitigation;
- open blocker questions;
- conflicting accepted decisions;
- missing or invalid required planning records;
- stale locks where current files no longer match `.ni/plan.lock.json`;
- target prompt compilation before a valid lock exists.

`BLOCKED` is not a failure mode to hide. It is the product refusing to start
execution from unclear or stale intent.

## Boundary

`ni` is not:

- a task runner;
- a spec runner;
- a multi-agent execution layer;
- a host-specific adapter;
- a queue;
- a shell adapter;
- release automation;
- PR automation;
- a competitor to downstream tools.

The kernel owns planning contracts, readiness gates, lockfiles, prompt
compilation, and inert downstream seed exports. It does not own downstream
execution state.

## Non-Software Projects

`ni` is not limited to software specs. The
[Neighborhood Cooling Study Protocol](../examples/research-protocol/) example
plans a research protocol with `product_type: research_protocol`, a `document`
delivery surface, protocol review evaluations, and a human-team handoff prompt.

It does not collect data, run analysis, deploy sensors, or execute fieldwork.
That boundary is the point: `ni` compiles intent for the project surface, then
hands off only when the locked plan is valid.

## Benchmark Methodology

The benchmark method compares a direct-to-agent prompt against the intent
quality available after the Intent Lock Protocol captures, validates, locks,
and compiles a plan.

It is not an execution benchmark. It does not run Codex, shell commands, model
APIs, downstream agents, queues, adapters, or harnesses. It measures readiness
gaps, hidden assumptions, non-goal coverage, stale plan detection, and prompt
boundedness before execution starts.

The full method lives in [Benchmark Protocol](43_BENCHMARK_PROTOCOL.md). The
fixture corpus lives in
[`testdata/benchmark/vague-requests/`](../testdata/benchmark/vague-requests/),
and the sample report template lives in
[`examples/benchmark-report/sample-report.md`](../examples/benchmark-report/sample-report.md).
