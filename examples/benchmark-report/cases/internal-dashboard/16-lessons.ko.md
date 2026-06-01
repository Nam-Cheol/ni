# Lessons

이 case가 유용한 이유는 original request가 얼핏 buildable하게 보이기 때문이다.
"Build a dashboard for the customer team"은 평범한 implementation prompt처럼
들리지만, benchmark는 네 blocker question이 답변되고 accepted 되기 전에는
handoff가 준비되지 않았음을 보여줬다.

## Lessons From the BLOCKED State

- 그럴듯한 web surface도 실제 product decision을 숨길 수 있다.
- "Customer team"만으로는 user, reviewer, accountable planning owner를 특정할
  수 없다.
- "Needs attention"은 source field, threshold 또는 ranking rule, freshness
  expectation이 accepted 되기 전까지 observable signal이 아니다.
- Customer data artifact를 신뢰하려면 privacy와 access-control boundary가 먼저
  accepted 되어야 한다.
- Planning meeting은 유효한 delivery boundary가 될 수 있지만 artifact와
  pass/fail evidence가 explicit해야 한다.

## Lessons From the READY State

- `READY`는 scope가 benchmark planning-meeting artifact readiness로 이동한 뒤에만
  유효했다.
- Lock은 prompt compilation 전에 accepted artifact contract와 source-of-truth
  order를 고정하기 때문에 유용하다.
- 4000자 prompt는 bounded handoff seed generation을 증명할 뿐 downstream
  success를 증명하지 않는다.
- 가장 강한 증거는 transition 자체다. Hidden assumption이 visible blocker가
  되었고, accepted artifact answer가 되었으며, 이후 isolated lock data가
  되었다.

## Practical Takeaways

- Reader가 무엇이 방지되었는지 볼 수 있도록 direct-to-agent risk note를 CLI
  proof 옆에 둔다.
- 낙관으로 채우는 대신 `not_measured` cell을 보존한다.
- Delivery surface를 명시한다. 이 case에서는 web dashboard가 아니라
  `document`다.
- Product-readiness claim은 artifact readiness의 부산물이 아니라 별도의 future
  measurement로 다룬다.
- 이 benchmark의 일부로 compiled prompt를 실행하지 않는다. 실행은 다른 system
  boundary를 측정한다.

## Reusable Pattern

Future benchmark case에는 같은 sequence를 기록한다.

| Step | Required evidence |
| --- | --- |
| Vague request | Source request and direct-to-agent risk notes |
| Blocked ni path | `ni status --proof --next-questions` output and blocker list |
| Resolution path | Required answers, expected planning updates, and unsafe assumptions avoided |
| Answered variant | Status proof showing `READY` or `READY_WITH_DEFERRALS` |
| Lock | Isolated lock summary and source list |
| Prompt | Bounded prompt summary and character count |
| Limits | Remaining `not_measured` claims and non-execution confirmation |

이 pattern은 repeated trial, independent reviewer, outcome measurement가 실제로
존재하기 전까지 qualitative 상태로 남아야 한다.
