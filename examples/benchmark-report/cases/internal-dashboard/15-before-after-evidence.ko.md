# Before/After Evidence

이 문서는 internal-dashboard benchmark의 한 전환을 기록한다.

```text
BLOCKED vague request -> answered benchmark artifact readiness -> READY isolated lock -> bounded prompt
```

이 증거는 pre-runtime evidence다. Dashboard implementation, downstream agent,
model API, shell adapter, queue, runtime harness가 시작되기 전의 readiness
artifact만 비교한다.

## Readiness Transition

| Stage | Readiness | Lock | ni-run prompt | Prompt count | Meaning |
| --- | --- | --- | --- | --- | --- |
| Vague request | BLOCKED | No | No | not_measured | Not safe to hand off |
| Answered artifact-readiness case | READY | Yes, isolated workspace only | Yes | 4000 | Safe to hand off as benchmark planning artifact |

## Original Vague Request State

- Readiness: `BLOCKED`
- Blockers: `OQ-001`부터 `OQ-004`
- Lock: no
- `ni-run` prompt: no
- Prompt count: `not_measured`

Evidence:

- `01-original-request.md`
- `02-direct-to-agent-risk.md`
- `06-ni-status-proof.md`
- `07-ni-next-questions.md`
- `08-blocker-analysis.md`

의미: original request는 그럴듯하지만 handoff하기 안전하지 않았다. Primary
dashboard user, supported decision, observable attention signal, source,
freshness, privacy, access boundary, planning-meeting acceptance evidence가
accepted 상태가 아니었다.

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

의미: user-provided answer는 benchmark planning-meeting artifact blocker를
해결했다. Isolated workspace는 `READY`에 도달했고
`examples/benchmark-report/cases/internal-dashboard/workspace/.ni/plan.lock.json`에
lock되었으며, 정확히 4000자의 bounded generic prompt를 만들었다.

## Claim Boundary

| Measured | Not measured |
| --- | --- |
| answer 전 `BLOCKED`; artifact-readiness answer 뒤 `READY`; isolated workspace lock; 4000-character prompt | dashboard product readiness; dashboard implementation quality; downstream agent performance; rework reduction; adoption; cost; latency; statistical effect |

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

`READY`는 benchmark planning-meeting artifact readiness에만 적용된다. 즉,
isolated case workspace가 benchmark report를 위한 inert prompt seed를 lock하고
compile하기에 충분한 accepted planning evidence를 갖췄다는 뜻이다.

`READY`는 dashboard product readiness를 주장하지 않는다. Dashboard
implementation quality는 `not_measured`로 남는다. Downstream agent performance도
`not_measured`다. Real user impact도 `not_measured`다.

## What ni Improved

- Hidden assumption을 명시적으로 만들었다.
- Premature handoff를 막았다.
- `OQ-001`부터 `OQ-004`의 acceptance evidence를 요구했다.
- Privacy, freshness, source, access boundary를 contract에 넣도록 했다.
- Dashboard implementation을 scope 밖에 유지했다.
- Isolated lock 뒤에만 bounded prompt를 만들었다.

## What ni Did Not Prove

- Implementation quality.
- Adoption.
- Rework reduction.
- Cost savings.
- Latency improvement.
- Statistical effect size.
- Downstream agent quality.

## Non-Execution Confirmation

Generated prompt는 실행하지 않았다. Dashboard는 구현하지 않았다. Downstream
agent는 호출하지 않았다. Model API도 호출하지 않았다. Shell adapter, queue,
runtime execution path, PR automation, release automation, empirical-product
claim도 추가하지 않았다.
