# ni

[English](README.md) | [한국어](README.ko.md)

Project Intent Compiler for AI Agents.

Don't run the agent yet. Compile the intent first.

`ni` turns planning conversations into a locked, versioned, verifiable project
contract before Codex, Claude, Spec Kit, Hyper Run, namba-ai, a generated
harness, or a human team starts execution.

The current product is `ni-kernel`: a deterministic pre-runtime control layer
for intent, not an execution harness.

```text
conversation -> docs/plan + .ni/contract.json -> ni status -> ni end -> locked intent -> ni run
```

## What Problem ni Solves

Agents are often handed prompts that sound actionable while still hiding the
intent needed for trustworthy execution:

- who the project is for,
- what must be accepted,
- which risks need mitigation,
- what is explicitly out of scope,
- which questions should block execution,
- whether the accepted plan has changed.

Most tools try to control the agent after a prompt, spec, worklist, or runtime
loop exists. `ni` moves control earlier. It asks whether project intent is
explicit, accepted, validated, locked, and unchanged before any downstream
actor starts work.

## Core Idea: Intent Lock Protocol

The [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) defines how a
planning conversation becomes a project contract, when that contract is ready
to lock, how the accepted plan is hashed, what downstream actors may trust, and
when handoff must stop because intent changed.

The kernel owns:

- `docs/plan/**`
- `.ni/contract.json`
- deterministic readiness validation
- `.ni/plan.lock.json`
- lock hash verification
- bounded prompt compilation
- inert downstream seed exports

After a plan is locked, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If locked hashes no longer match, target handoff stops with `BLOCKED`.

## 5-Minute Demo

The fastest way to understand `ni` is the
[Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) demo.

It starts from:

```text
Build me a dashboard for my team.
```

A direct-to-agent path would force hidden assumptions about users, data,
workflow, non-goals, and success criteria. The `ni` path records the request as
planning intent, then refuses to treat it as executable while blocker questions
remain open.

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected result:

```text
BLOCKED
```

That is the point: ambiguous execution is blocked before an agent starts.

## Non-Software Proof

`ni` is not a software spec generator. It compiles project intent for any
product surface.

Try the [Neighborhood Cooling Study Protocol](examples/research-protocol/):

```bash
go run ./cmd/ni status --dir examples/research-protocol
go run ./cmd/ni run --dir examples/research-protocol --target human-team --out examples/research-protocol/generated/human-team.prompt.md
```

This locked example plans a research protocol, not an app. It has
`product_type: research_protocol`, a `document` delivery surface, protocol
review evaluations, and a human-team handoff prompt. It does not collect data,
run analysis, deploy sensors, or execute fieldwork.

## Core Flow

Create a planning workspace:

```bash
go run ./cmd/ni init --dir <path> --profile prototype
```

Use sustained model-user conversation to maintain `docs/plan/**` and
`.ni/contract.json` together. Skills and models are UX. The CLI is authority.

```bash
go run ./cmd/ni status --dir <path>
go run ./cmd/ni end --dir <path>
go run ./cmd/ni run --dir <path> --target codex --max-chars 4000
```

`ni run` prints or writes a prompt. It does not execute that prompt.

## What ni Blocks

`ni` blocks downstream handoff when intent is not trustworthy yet:

- accepted capabilities without linked evaluations,
- high-severity risks without mitigation,
- open blocker questions,
- conflicting accepted decisions,
- missing or invalid required planning records,
- stale locks where current files no longer match `.ni/plan.lock.json`,
- target prompt compilation before a valid lock exists.

The [Benchmark Protocol](docs/43_BENCHMARK_PROTOCOL.md) describes how to
compare direct-to-agent prompts against locked `ni` intent without running
downstream agents.
The fixture corpus lives in
[`testdata/benchmark/vague-requests/`](testdata/benchmark/vague-requests/),
and the sample report template lives in
[`examples/benchmark-report/sample-report.md`](examples/benchmark-report/sample-report.md).

## Commands Summary

The core path is:

```text
ni init -> ni status -> ni end -> ni run
```

Other implemented kernel commands inspect targets, export locked seed material,
record inert feedback and pressure, manage explicit amendments, compare
planning states, and propose inert graph or harness material.

See the full [command reference](docs/commands.md).

## Targets Summary

Targets are consumption shapes for a locked plan. They are not integrations
that `ni` executes, runtime adapters that `ni` owns, or lifecycle state that
becomes part of `ni-kernel`.

Built-in targets include:

- `generic`, `codex`, and `human-team` prompt or handoff targets,
- `hyper-run`, `namba-ai`, `ouroboros`, and `spec-kit` seed targets.

See the [Target Story](docs/45_TARGET_STORY.md) for target-by-target
boundaries.

## What ni Is Not

`ni` is not:

- a task runner,
- a SPEC runner,
- a multi-agent execution layer,
- a Codex adapter,
- a queue,
- a shell adapter,
- release automation,
- PR automation,
- Hyper Run, Spec Kit, Ouroboros, or namba-ai.

Downstream prompts, seed packages, and harness proposals are derived and
mutable. They do not become kernel-owned execution state.

## Examples and Docs

Start here:

- [Positioning](docs/40_POSITIONING.md)
- [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md)
- [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/)
- [Neighborhood Cooling Study Protocol](examples/research-protocol/)
- [Command Reference](docs/commands.md)
- [Benchmark Protocol](docs/43_BENCHMARK_PROTOCOL.md)
- [Benchmark Report Template](examples/benchmark-report/sample-report.md)
- [Target Story](docs/45_TARGET_STORY.md)
- [v0.2.0 Release Draft](docs/47_RELEASE_DRAFT_v0.2.0.md)

## Development and Release Status

`ni` is currently source-first.
Release status: does not claim package distribution or a published binary release.
Package publishing, Homebrew taps, GoReleaser, and automated release tooling are
outside the current kernel scope.

Use `go run` from source, build a local binary with `make build`, or install
locally with `make install-local`. See [docs/22_INSTALL.md](docs/22_INSTALL.md)
for source, local build, or local install details.

Public demo verification:

```bash
bash scripts/demo-check.sh
```

Repository validation:

```bash
bash scripts/quality.sh
```

CI validation is defined in `.github/workflows/ci.yml` and runs Go tests,
quality checks, and smoke tests.

Source/build/install verification:

```bash
bash scripts/install-check.sh
```

`ni` is licensed under the [MIT License](LICENSE). Release readiness notes live
in [docs/46_RELEASE_READINESS.md](docs/46_RELEASE_READINESS.md), and the
project security policy is [SECURITY.md](SECURITY.md).
