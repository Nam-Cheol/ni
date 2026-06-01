# Before/After Evidence

This package records one internal-dashboard benchmark transition:

```text
BLOCKED vague request -> answered benchmark artifact readiness -> READY isolated lock -> bounded prompt
```

The evidence is pre-runtime. It compares readiness artifacts before any
dashboard implementation, downstream agent, model API, shell adapter, queue, or
runtime harness starts work.

## Readiness Transition

| Stage | Readiness | Lock | ni-run prompt | Prompt count | Meaning |
| --- | --- | --- | --- | --- | --- |
| Vague request | BLOCKED | No | No | not_measured | Not safe to hand off |
| Answered artifact-readiness case | READY | Yes, isolated workspace only | Yes | 4000 | Safe to hand off as benchmark planning artifact |

## Original Vague Request State

- Readiness: `BLOCKED`
- Blockers: `OQ-001` through `OQ-004`
- Lock: no
- `ni-run` prompt: no
- Prompt count: `not_measured`

Evidence:

- `01-original-request.md`
- `02-direct-to-agent-risk.md`
- `06-ni-status-proof.md`
- `07-ni-next-questions.md`
- `08-blocker-analysis.md`

Meaning: the original request was plausible but not safe to hand off. The
primary dashboard user, supported decision, observable attention signals,
source/freshness/privacy/access boundaries, and planning-meeting acceptance
evidence were not accepted.

## Answered Artifact-Readiness State

- Readiness: `READY`
- Blockers: none
- Lock: yes, isolated benchmark workspace only
- `ni-run` prompt: yes
- Prompt count: `4000`

Evidence:

- `10-answer-packet.md`
- `11-resolved-status-proof.md`
- `12-resolved-next-questions.md`
- `13-lock-summary.md`
- `14-bounded-prompt-summary.md`

Meaning: the user-provided answers resolved the benchmark planning-meeting
artifact blockers. The isolated workspace reached `READY`, was locked at
`examples/benchmark-report/cases/internal-dashboard/workspace/.ni/plan.lock.json`,
and produced a bounded generic prompt of exactly 4000 characters.

## Claim Boundary

| Measured | Not measured |
| --- | --- |
| `BLOCKED` before answers; `READY` after artifact-readiness answers; isolated workspace lock; 4000-character prompt | dashboard product readiness; dashboard implementation quality; downstream agent performance; rework reduction; adoption; cost; latency; statistical effect |

| Claim | Supported? | Evidence | Notes |
| --- | --- | --- | --- |
| ni exposes unclear intent | Yes | BLOCKED proof with OQ-001~OQ-004 | Pre-runtime evidence |
| ni prevents premature handoff | Yes | No lock/no prompt before answers | Planning evidence |
| ni can produce bounded handoff after lock | Yes | 4000-char prompt after isolated lock | Artifact-readiness only |
| dashboard product is ready | No | not_measured | Out of scope |
| dashboard product readiness is proven | No | not_measured | Artifact-readiness only |
| downstream agent will perform well | No | not_measured | No downstream execution |

## Remaining not_measured

| Area | Status | Reason |
| --- | --- | --- |
| dashboard implementation quality | not_measured | no dashboard was built |
| dashboard product readiness | not_measured | no dashboard product was validated |
| downstream agent performance | not_measured | no agent was run |
| rework reduction | not_measured | no repeated trial |
| adoption | not_measured | no external usage data |
| cost/latency | not_measured | no runtime measurement |
| statistical effect | not_measured | no repeated quantitative study |

## Critical Scope Note

`READY` applies only to benchmark planning-meeting artifact readiness. It means
the isolated case workspace has enough accepted planning evidence to lock and
compile an inert prompt seed for the benchmark report.

`READY` does not claim dashboard product readiness. Dashboard implementation
quality remains `not_measured`. Downstream agent performance remains
`not_measured`. Real user impact remains `not_measured`.

## What ni Improved

- Made hidden assumptions explicit.
- Blocked premature handoff.
- Forced acceptance evidence for `OQ-001` through `OQ-004`.
- Forced privacy, freshness, source, and access boundaries into the contract.
- Kept dashboard implementation out of scope.
- Created a bounded prompt only after an isolated lock existed.

## What ni Did Not Prove

- Implementation quality.
- Adoption.
- Rework reduction.
- Cost savings.
- Latency improvement.
- Statistical effect size.
- Downstream agent quality.

## Non-Execution Confirmation

No generated prompt was executed. No dashboard was implemented. No downstream
agent was called. No model API was called. No shell adapter, queue, runtime
execution path, PR automation, release automation, or empirical-product claim
was added.
