# Blocker 분석

이 문서는 internal-dashboard benchmark가 왜 `BLOCKED`로 남아 있는지 설명한다.
Blocker question에 답하지 않고, resolved 또는 deferred로 표시하지 않으며,
workspace를 lock하거나 prompt를 compile하지 않고, downstream dashboard 작업도
허가하지 않는다.

`BLOCKED`는 유효한 benchmark result다. 이 case에서 ni는 구현 전에 readiness
gap을 명시적으로 드러내어 premature handoff를 막았다. 이 case가 dashboard
implementation quality를 증명한 것은 아니다. Missing intent가 auditable
blocker evidence로 기록되었다는 점을 증명한다.

| Blocker | 무엇이 unknown인가? | 왜 lock을 막는가? | 어떤 answer가 resolve할 수 있는가? | 피해야 할 unsafe assumption |
| --- | --- | --- | --- | --- |
| `OQ-001` | 정확한 primary dashboard user와 dashboard가 지원해야 할 decision이 accepted 상태가 아니다. | Actor와 decision이 없으면 requirement, layout, permission, metric, meeting evidence가 guessed intent 위에 만들어진다. | User가 확인한 role 또는 audience와, 그 audience가 내려야 하는 concrete decision. | "customer team"이 customer success manager, executive, support, sales 등 어떤 role을 뜻한다고 가정하는 것. Decision이 renewal risk, escalation, outreach, forecasting, meeting reporting이라고 가정하는 것. |
| `OQ-002` | "needs attention"이 observable account signal이나 ranking criteria로 정의되지 않았다. | Attention signal이 없으면 account priority, sorting, alert, correctness에 대한 trustworthy acceptance criteria가 없다. | v0에서 account가 attention이 필요한 시점을 정의하는 user-confirmed signal, threshold, ordering rule, review criteria. | Attention이 product usage drop, open support escalation, unpaid invoice, renewal date, health score, executive request, 특정 ranking formula라고 가정하는 것. |
| `OQ-003` | Allowed source system, account field, freshness rule, privacy constraint, access control이 accepted 상태가 아니다. | Customer-account data와 health signal은 privacy, security, stale-data risk를 만들 수 있으므로 explicit data boundary 없이는 lock이 unsafe하다. | v0의 source system, field list 또는 field category, freshness expectation, access audience, privacy/security constraint에 대한 user-confirmed answer. | Trusted data가 이미 있다고 가정하거나, CRM/support/billing/product data를 마음대로 사용할 수 있다고 가정하거나, sensitive field를 보여도 된다고 가정하거나, 모든 customer-team member가 access할 수 있다고 가정하는 것. |
| `OQ-004` | 다음 planning meeting에 필요한 acceptance evidence가 accepted 상태가 아니다. | "ready for the next planning meeting"은 pass/fail evidence, date, audience, scope, review artifact를 정의하지 않으므로 lock이 unbounded handoff를 허용하게 된다. | User가 확인한 meeting date 또는 timing, review audience, minimum useful artifact, v0 planning readiness acceptance check. | Prototype, static mockup, planning memo, metric table, live dashboard, implementation plan 중 아무거나 충분하다고 가정하거나, "easy to use"에 explicit evidence가 필요 없다고 가정하는 것. |

## 해석

이 blocker set은 올바른 boundary에서 execution을 멈추기 때문에 유용하다.
Direct prompt는 중요한 assumption을 숨긴다. ni path는 그 assumption을 risk,
evaluation, non-goal, synchronized planning state와 연결된 open question으로
기록한다.

이 evidence는 rework reduction, adoption, cost, latency, statistical effect,
downstream agent performance, dashboard quality를 측정하지 않는다. 또한 future
implementation이 correct하다는 점도 보여주지 않는다. 보여주는 것은
`ni status`가 ambiguous intent를 ready로 취급하지 않았다는 사실이다.

## 현재 결과

- Readiness: `BLOCKED`
- Workspace locked: no
- Bounded prompt compiled: no
- Prompt character count: `not_measured`
- Downstream execution: none
