# Contributing to ni

Thank you for helping improve `ni`.

`ni` is a Project Intent Compiler for AI Agents. The current product is
`ni-kernel`: a deterministic pre-runtime layer that creates, validates, locks,
and compiles planning contracts before any downstream execution starts.

`ni` is not an execution runtime. It should protect the Intent Lock Protocol,
not grow into the thing that runs the work.

## Product Boundary

The kernel owns:

- planning docs and `.ni/contract.json` sync,
- deterministic readiness validation,
- `.ni/plan.lock.json` lock/hash correctness,
- bounded prompt compilation,
- inert downstream seed material derived from locked intent.

The kernel must not own:

- task runners,
- SPEC runners,
- Codex execution,
- shell adapters,
- queues,
- agent teams,
- PR automation,
- release automation,
- downstream runtime state.

Contributions that add task runners, SPEC runners, Codex exec, queues, agent
teams, or PR/release automation to core are out of scope for `ni-kernel`.

## Preferred Contributions

Good kernel contributions usually improve one of these areas:

- readiness rules,
- docs/contract sync validation,
- examples,
- target seed formats,
- benchmark fixtures,
- documentation clarity,
- lock/hash correctness,
- conversation authoring UX.

Target seed material is welcome when it remains derived, inert, and locked-plan
dependent. If a change starts, schedules, tracks, or completes downstream work,
it belongs outside `ni-kernel`.

## Before Opening an Issue or PR

Please check whether the proposal belongs in `ni-kernel` or in downstream seed
material. A useful test is:

```text
Does this validate or compile locked intent, or does it execute work?
```

If it executes work, it is outside the current core boundary.

For bugs, include:

- `ni` version or commit,
- command and flags,
- expected result,
- actual result,
- workspace shape, including relevant `docs/plan/**`, `.ni/contract.json`, and
  `.ni/plan.lock.json` state when safe to share.

Do not share secrets, credentials, proprietary planning contracts, or sensitive
prompts in public issues.

## Pull Requests

Keep changes small and tied to one coherent intent. When code changes exist,
include validation evidence from the relevant checks.

For repository validation, run:

```bash
bash scripts/quality.sh
```

When Go files are touched, also run:

```bash
go test ./...
```

Do not weaken acceptance criteria, risks, mitigations, evaluations, or
non-goals to make validation pass.

## No Contributor License Agreement

This project does not add a contributor license agreement in the current
contribution flow.
