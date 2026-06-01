# Lessons

Research-protocol case가 중요한 이유는 original request가 단순히 product detail이
부족한 또 다른 software surface가 아니기 때문이다. 이 request는 resident,
safety, privacy, consent, field operation에 영향을 줄 수 있는 작업을 계획하게
한다. 그래서 premature handoff의 위험은 잘못된 화면을 만드는 것보다 크다.

## Why This Case Matters

- ni가 software-only가 아님을 보여준다.
- Human team handoff 전, 그리고 research fieldwork planning 전 intent
  readiness를 검증한다.
- ni를 execution 또는 fieldwork system으로 만들지 않고 research planning을
  contract boundary에 둔다.
- Vague request가 practical하게 들려도 safe research planning에 필요한 evidence가
  빠질 수 있음을 보여준다.

## What ni Surfaced

- Research question ambiguity: request는 무엇을 배우고, 비교하고, 어떤 evidence를
  acceptance로 볼지 명시하지 않았다.
- Participant and observation scope: eligible area, group, observation window가
  accepted 상태가 아니었다.
- Privacy, consent, and data boundaries: participant contact, personal data,
  translation, accessibility, storage rule을 명시해야 했다.
- Field safety and weather stop rules: extreme heat에서는 fieldwork를 중단하거나
  바꿔야 하는 조건이 필요하다.
- Reviewer and acceptance evidence: handoff 전에 named review owner와 clear
  pass/fail evidence가 필요했다.

## Why Synthetic Answers Are Acceptable Here

- Benchmark fixture data라고 명확히 표시되어 있다.
- `BLOCKED`에서 `READY`로 가는 ni readiness transition을 test하는 데만 사용된다.
- Real approval, fieldwork authorization, ethics review, research quality
  evidence가 아니다.
- 실제 study가 ready라고 가장하지 않고도 status, lock, bounded prompt behavior를
  검증하게 해준다.

## Why This Strengthens ni's Positioning

- ni는 AI agent가 coding을 시작하기 전뿐 아니라 human team이 ownership을 갖기
  전에도 유용하다.
- ni는 unclear intent가 safety와 governance risk를 만들 수 있는 non-software
  planning에도 유용하다.
- ni는 execution-oriented가 아니다. Intent를 gate하고, accepted planning
  artifact를 lock하고, bounded handoff seed material을 compile한다.
- ni는 unresolved scope, safety, privacy, reviewer question을 handoff 전에 드러내
  research fieldwork planning을 도울 수 있다.

## Practical Takeaways

- Original `BLOCKED` proof와 resolved fixture proof를 나란히 보존한다.
- Workspace가 `READY`에 도달한 뒤에도 `not_measured` boundary를 보존한다.
- Research approval과 fieldwork authorization은 ni readiness의 side effect가
  아니라 external governance step으로 다룬다.
- 이 benchmark의 일부로 compiled prompt를 실행하지 않는다. 실행은 Intent Lock
  Protocol이 아니라 downstream system을 측정하게 된다.
