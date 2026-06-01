# Before/After Evidence

This package records the research-protocol benchmark transition:

```text
vague research request -> BLOCKED -> synthetic fixture answers -> READY isolated lock -> bounded prompt
```

The evidence is pre-runtime benchmark evidence only. It does not execute the
generated prompt, run fieldwork, call downstream agents, call model APIs, or
claim real research approval.

## Table 1: Readiness Transition

| Stage | Readiness | Lock | ni-run prompt | Prompt count | Meaning |
| --- | --- | --- | --- | --- | --- |
| Vague research request | BLOCKED | No | No | not_measured | Not safe to hand off |
| Synthetic answered protocol artifact | READY | Yes, isolated workspace only | Yes | 4000 | Safe to hand off as benchmark planning artifact |

## Original Vague Request State

- Readiness: `BLOCKED`
- Blockers: `OQ-001` through `OQ-005`
- Lock: no
- `ni-run` prompt: no
- Prompt count: `not_measured`
- Fieldwork: not performed
- Research approval: not claimed

Evidence:

- `01-original-request.md`
- `02-direct-to-agent-risk.md`
- `06-ni-status-proof.md`
- `07-ni-next-questions.md`
- `08-blocker-analysis.md`

Meaning: the original research request was plausible but not safe to hand off.
The research question, participant or observation scope, consent/privacy/data
boundaries, field safety stop rules, vulnerable-group safeguards, reviewer
owner, and acceptance evidence were not accepted.

## Synthetic Fixture Answered State

- Readiness: `READY`
- Blockers: none
- Lock: yes, isolated benchmark workspace only
- `ni-run` prompt: yes
- Prompt count: `4000`
- Synthetic fixture data: yes
- Real research approval: no
- Fieldwork authorization: no

Evidence:

- `10-answer-packet.md`
- `11-resolved-status-proof.md`
- `12-resolved-next-questions.md`
- `13-lock-summary.md`
- `14-bounded-prompt-summary.md`

Meaning: synthetic fixture answers resolved the benchmark planning blockers
inside the isolated workspace. The workspace reached `READY`, was locked at
`examples/benchmark-report/cases/research-protocol/workspace/.ni/plan.lock.json`,
and produced a bounded prompt of exactly 4000 characters.

## Critical Scope Note

`READY` applies only to synthetic benchmark protocol planning artifact
readiness.

`READY` does not claim real research approval. It does not authorize fieldwork.
It does not prove research quality. It does not prove intervention
effectiveness. The benchmark proves that ni exposes and gates
research-planning readiness gaps before handoff.

## Table 2: Claim Boundary

| Measured | Not measured |
| --- | --- |
| `BLOCKED` before synthetic answers; `READY` after synthetic benchmark protocol artifact answers; isolated workspace lock; 4000-character prompt | real research approval; fieldwork authorization; research quality; intervention effectiveness; downstream agent performance; rework reduction; adoption; cost; latency; statistical effect |

| Claim | Supported? | Evidence | Notes |
| --- | --- | --- | --- |
| ni exposes unclear research intent | Yes | BLOCKED proof with OQ-001~OQ-005 | Pre-runtime evidence |
| ni prevents premature research handoff | Yes | No lock/no prompt before answers | Planning evidence |
| ni can produce bounded handoff after lock | Yes | 4000-char prompt after isolated lock | Synthetic benchmark artifact only |
| research protocol is approved | No | not_measured | Out of scope |
| real research approval exists | No | not_measured | Synthetic benchmark only |
| fieldwork is authorized | No | not_measured | Out of scope |
| downstream agent will perform well | No | not_measured | No downstream execution |

## Table 3: Remaining not_measured

| Area | Status | Reason |
| --- | --- | --- |
| real fieldwork | not_measured | no fieldwork was performed |
| participant recruitment | not_measured | no recruitment occurred |
| data collection | not_measured | no data was collected |
| research quality | not_measured | no real review board approval or field validation |
| intervention effectiveness | not_measured | no intervention was tested |
| downstream agent performance | not_measured | no downstream agent was run |
| rework reduction | not_measured | no repeated trial |
| adoption | not_measured | no external usage data |
| cost/latency | not_measured | no runtime measurement |
| statistical effect size | not_measured | no repeated quantitative study |

## Non-Execution Confirmation

No generated prompt was executed. No fieldwork was performed. No participant
recruitment, data collection, intervention placement, downstream agent call,
model API call, shell adapter, queue, PR automation, release automation, or
runtime execution path was added.
