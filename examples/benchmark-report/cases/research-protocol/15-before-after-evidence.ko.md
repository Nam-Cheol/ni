# Before/After Evidence

이 문서는 research-protocol benchmark transition을 기록한다.

```text
vague research request -> BLOCKED -> synthetic fixture answers -> READY isolated lock -> bounded prompt
```

이 evidence는 pre-runtime benchmark evidence일 뿐이다. Generated prompt를
실행하지 않고, fieldwork를 수행하지 않으며, downstream agent나 model API를
호출하지 않고, real research approval을 주장하지 않는다.

## Table 1: Readiness Transition

| Stage | Readiness | Lock | ni-run prompt | Prompt count | Meaning |
| --- | --- | --- | --- | --- | --- |
| Vague research request | BLOCKED | No | No | not_measured | Not safe to hand off |
| Synthetic answered protocol artifact | READY | Yes, isolated workspace only | Yes | 4000 | Safe to hand off as benchmark planning artifact |

## Original Vague Request State

- Readiness: `BLOCKED`
- Blockers: `OQ-001`부터 `OQ-005`
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

의미: original research request는 그럴듯했지만 handoff하기에 안전하지 않았다.
Research question, participant 또는 observation scope, consent/privacy/data
boundary, field safety stop rule, vulnerable-group safeguard, reviewer owner,
acceptance evidence가 accepted 상태가 아니었다.

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

의미: synthetic fixture answer가 isolated workspace 안에서 benchmark planning
blocker를 해결했다. Workspace는 `READY`에 도달했고,
`examples/benchmark-report/cases/research-protocol/workspace/.ni/plan.lock.json`
에 lock되었으며, 정확히 4000자 bounded prompt를 생성했다.

## Critical Scope Note

`READY`는 synthetic benchmark protocol planning artifact readiness에만
적용된다.

`READY`는 real research approval을 주장하지 않는다. Fieldwork를 승인하지
않는다. Research quality를 증명하지 않는다. Intervention effectiveness를
증명하지 않는다. 이 benchmark가 증명하는 것은 ni가 handoff 전에
research-planning readiness gap을 드러내고 gate한다는 점이다.

## Table 2: Claim Boundary

| Claim | Supported? | Evidence | Notes |
| --- | --- | --- | --- |
| ni exposes unclear research intent | Yes | BLOCKED proof with OQ-001~OQ-005 | Pre-runtime evidence |
| ni prevents premature research handoff | Yes | No lock/no prompt before answers | Planning evidence |
| ni can produce bounded handoff after lock | Yes | 4000-char prompt after isolated lock | Synthetic benchmark artifact only |
| research protocol is approved | No | not_measured | Out of scope |
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

Generated prompt는 실행하지 않았다. Fieldwork, participant recruitment, data
collection, intervention placement, downstream agent call, model API call,
shell adapter, queue, PR automation, release automation, runtime execution path는
추가하지 않았다.
