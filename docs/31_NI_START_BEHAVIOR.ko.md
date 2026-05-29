# ni-start 동작

`ni-start`는 NI 계획을 지속 대화로 작성하는 모드다. `ni init` 이후 사용자가
`.ni/contract.json`을 직접 편집하지 않아도, 모델이 대화를 바탕으로 계획
상태를 유지하도록 돕는다.

이 skill은 CLI gate를 대체하지 않는다. skill은 계획 상태를 쓰고, CLI는
검증, 잠금, 컴파일의 권한이다.

장기 계획을 이어가기 위한 resume mode도 지원한다. resume mode는 숨겨진
chat memory가 아니라 저장된 project file에서 계획 연속성을 재구성하거나
검증한다.

```text
ni init -> ni-start conversation -> docs/plan + .ni/contract.json -> ni status
```

## 책임

계획 turn을 시작할 때 `ni-start`는 현재 계획 source를 읽는다.

- `AGENTS.md`: repository authority rule,
- `docs/plan/**`: 사람이 읽는 계획 상태,
- `.ni/contract.json`: machine-readable 계획 상태,
- `.ni/session.json`: 권한이 없는 carryover context,
- `.ni/plan.lock.json`: 있을 경우 locked-plan authority 확인.

그 다음 추가 입력을 요청하기 전에 현재 상태를 요약한다. 시작 요약에는 반드시
다음 항목이 포함되어야 한다.

1. 현재 목적,
2. 활성 readiness profile,
3. product type과 delivery surface,
4. accepted capability,
5. unresolved blocker question,
6. 최근 decision,
7. 다음으로 권장하는 planning focus.

`.ni/session.json`의 주장은 contract, docs, 그리고
`ni status --dir . --proof --next-questions` 결과와 대조해야 한다.

## First-run conversation card

fresh workspace에서 `ni init` 직후 `ni status --proof --next-questions`는
보통 첫 intent blocker를 보고한다.

- `R014 Project purpose is missing`
- `R015 Actors or outcomes are missing`
- `R016 Delivery surface is missing`

`ni-start`는 이 blocker를 사용자가 막힌 느낌을 받는 상태 메시지가 아니라
간결한 opening planning card로 바꿔야 한다. card는 `ni`가 막힌 이유가 초기
intent가 아직 lock할 만큼 명시적이지 않기 때문이며, implementation은 아직
시작되지 않았다고 설명해야 한다.

권장 문구:

```text
ni is blocked because the initial project intent is not explicit enough to lock
yet. I need three things before execution can safely start: what reality this
project should change, who it is for, and how it will be delivered.

Implementation has not started. This is still planning.
```

그 다음 first-run blocker를 중심으로 최대 3개의 focused question만 묻는다.

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?

CLI가 template open question도 함께 보여줄 수 있지만, generic brainstorming이
이 세 가지 first-run question을 밀어내면 안 된다. 사용자가 답하면
`ni-start`는 purpose를 `docs/plan/00_project_brief.md`와 `project.purpose`에,
actor와 outcome을 `docs/plan/01_actors_outcomes.md`와 matching contract
record에, delivery surface를 `docs/plan/08_delivery_operation.md`,
`product_type`, `delivery_surfaces`에 적절히 기록한다.

명확한 제외 사항은 non-goal로 기록한다. 불확실하거나 tentative하거나 vague한
답은 사용자가 확인할 때까지 assumption이나 blocker open question으로 남겨야
한다. model은 readiness를 통과시키려고 vague answer를 accepted decision으로
바꾸면 안 된다. 답을 기록한 뒤에는 다시
`ni status --dir . --proof --next-questions`를 실행하거나 요청하고, CLI 결과를
다음 권한으로 사용한다.

그 status 결과가 `SYNC-014`, `SYNC-015`, `SYNC-016`을 보고하면 first-run
답변이 docs/contract pair에 일관되게 기록되지 않은 것이다. `ni-start`는 stale한
쪽을 고치거나, 불확실한 intent를 assumption, blocker open question, deferral로
명시해야 한다. 이 sync diagnostic이 readiness를 막고 있는 동안 `ni-end`로
진행하면 안 된다.

`ni status --next-questions`는 prompt를 `First-run card`, `Sync repairs`,
`Risk decisions`, `Evaluation evidence`, `Scope boundaries`, `Open blockers`
같은 heading으로 묶는다. `ni-start`는 사용자에게 질문할 때 CLI가 반환한 group,
ID, location, answer-shape field를 반드시 보존하고, CLI가 반환한 top question만
물어야 한다.

## Grouped next-question 처리

`ni-start`는 계획 turn 시작 시와 의미 있는 authoring update 뒤에 다음 명령을
실행하거나 요청해야 한다.

```bash
ni status --dir . --proof --next-questions
```

grouped `Next questions` section이 있으면 이것이 conversation의 primary driver다.
`ni-start`는 CLI 순서대로 group을 읽고, highest-priority group을 선택하며,
compact `First-run card`가 아닌 한 turn마다 최대 1개 group만 물어야 한다. 한 번에
묻는 primary question은 최대 3개다. deterministic next question이 있을 때 넓은
brainstorming question을 새로 만들면 안 된다.

model response에서는 group label을 보여주거나 보존해야 한다. CLI가 `Location`이나
`Answer shape`을 출력하면, 사용자가 기대된 형태로 답할 수 있도록 그 field도
충분히 보존한다. readiness는 model judgment가 아니라 deterministic CLI gate에
의해 blocked 상태로 남는다.

group별 규칙:

- `First-run card`: purpose, actor/outcome, delivery surface를 묻는 compact
  3-question card를 사용한다. unrelated lower-priority question을 추가하지 않는다.
- `Sync repairs`: repair question을 묻고, contract update, docs revise,
  both revise, blocker 유지와 reason 중 무엇인지 묻는다. 한쪽에 이미 useful content가
  있으면 broad planning을 다시 시작하지 않는다. `SYNC-014`, `SYNC-015`,
  `SYNC-016` repair question이 있으면 matching `R014`, `R015`, `R016`을 다시 묻지
  않는다.
- `Risk decisions`: mitigation, owner, monitoring plan, accepted-risk decision,
  explicit deferral을 묻는다. readiness 통과를 위해 risk severity를 낮추자고 제안하지
  않는다.
- `Evaluation evidence`: capability 완료를 증명하는 evidence를 묻는다. answer
  shape로 test, review checklist, demo condition, user approval, protocol check,
  manual inspection을 제시한다.
- `Scope boundaries`: explicit non-goal을 묻고, non-goal이 scope drift를 막는다고
  설명한다.
- `Open blockers`: resolve, defer with reason, keep blocking 중 무엇인지 묻는다.
  open question을 accepted decision으로 조용히 바꾸면 안 된다.

사용자가 선택된 group에 답하면 `ni-start`는 `docs/plan/**`,
`.ni/contract.json`, `.ni/session.json`을 함께 업데이트한 뒤
`ni status --dir . --proof --next-questions`를 다시 실행하거나 요청한다.

## Resume mode

나중의 model session이 계획을 이어갈 때, `ni-start`는 일반 turn과 같은
authoritative input에서 시작한다. `.ni/session.json`이 있으면 active focus,
last summary, pending question, recent decision과 risk, last readiness status,
최근 변경 docs를 planning aid로 사용할 수 있다.

모든 session claim은 `.ni/contract.json`, `docs/plan/**`, lock state, 그리고
가능하면 `ni status --dir . --proof --next-questions`로 검증해야 한다. session
file이 contract record와 충돌하면 contract가 우선이고 충돌을 보고한다.
locked plan과 충돌하면 lock과 locked docs가 우선이다. lock hash mismatch는
turn을 `BLOCKED`로 중단한다.

`.ni/session.json`이 없거나, invalid, empty, stale이면 `ni-start`는
`docs/plan/**`, `.ni/contract.json`, CLI readiness output에서 resume summary를
재구성한다. 다음 의미 있는 planning edit은 `.ni/session.json`을 bounded
continuity state로 다시 만들거나 갱신해야 한다. 기본적으로 raw transcript
전체를 저장하지 않는다.

## Gap detection

`ni-start`는 현재 contract와 docs에서 필요한 계획 영역이 빠졌는지 확인해야
한다. readiness-blocking interview prompt의 첫 source는
`ni status --proof --next-questions`다. 흔한 gap은 다음과 같다.

- purpose, actor, outcome, delivery surface가 아직 TODO인 경우,
- capability에 requirement나 evaluation이 없는 경우,
- accepted capability의 artifact가 없는 경우,
- high-severity risk에 mitigation이 없는 경우,
- scope나 acceptance criteria에 영향을 주는 blocker question,
- constraint나 non-goal이 요청된 behavior와 충돌하는 경우,
- docs와 `.ni/contract.json`이 같은 record에 대해 불일치하는 경우.

질문은 readiness를 막는 gap에 집중해야 한다. CLI가 grouped next question을
반환하면 highest-priority group을 먼저 묻고, turn마다 최대 1개 group만 묻고,
한 번에 primary question은 최대 3개만 묻는다. deterministic next question이
있는 동안 넓고 일반적인 brainstorming 질문은 피한다. 예를 들어 "계획에 또
무엇을 넣을까요?" 대신 다음처럼 묻는다.

```text
For CAP-002, what evidence proves this capability works, or should that evidence be deferred?
```

질문은 관련 ID를 보존하고, implementation work를 암시하지 않으며, acceptance를
압박하지 않아야 한다. 적절한 planning outcome이면 `deferred`나
`not_applicable`도 허용해야 한다.

## Persistence Rules

사용자가 답한 뒤 `ni-start`는 같은 authoring pass에서 두 planning form을
갱신하고 session state를 새로 고친다.

- `docs/plan/**`: 사람이 읽는 계획 설명,
- `.ni/contract.json`: CLI가 검증하는 stable ID, status, trace link,
- `.ni/session.json`: latest focus, short summary, active readiness profile,
  product type과 delivery surface, pending question, recent decision, recent
  risk, readiness status, readiness blocker, 이 turn에서 바뀐 docs.

대화가 변경한 purpose, actor, capability, requirement, decision, risk,
evaluation, non-goal, constraint, artifact, assumption, open question을 기록한다.
tentative 또는 inferred statement는 사용자가 확인할 때까지 assumption이나
open question으로 남긴다. session state는 planning aid일 뿐이다. locked docs를
override하거나 docs complete를 표시하거나 raw chat log 전체를 기본 저장하면 안
된다.

기존 ID는 안정적으로 유지한다. 구분되는 record가 필요할 때만 새 ID를 붙인다.
obsolete record는 history 보존이 중요하면 rejected, deferred, not applicable로
표시한다.

## Readiness Handoff

의미 있는 update 뒤 `ni-start`는 다음 명령을 실행하거나 요청한다.

```bash
ni status --dir . --proof --next-questions
```

status 결과가 권한이다. `BLOCKED`면 `ni-start`는 planning을 열어 두고
`next_questions`의 highest-priority group을 묻는다. CLI가 다음 질문을 반환하지
않으면 readiness issue를 직접 보여준다. `READY` 또는 `READY_WITH_DEFERRALS`면
`ni-end`로 이동하자고 제안할 수 있다.

`ni-start`는 model judgment만으로 completion을 선언하면 안 된다.

## Boundaries

`ni-start`가 하면 안 되는 일:

- user-facing contract `add`, `list`, `set` command 추가,
- implementation work 실행,
<!-- ni-boundary-allow: explicit negative boundary list item. -->
- SPEC runner behavior 생성,
- shell, Codex, queue, evidence-runner, downstream runtime adapter 생성,
- downstream runtime 직접 호출,
- `.ni/plan.lock.json` 생성 또는 편집,
- validation을 통과시키려고 accepted requirement나 evaluation 약화.

skill은 downstream seed idea를 planning content로 제안할 수만 있다. downstream
harness material은 derived and mutable 상태이며 kernel-owned execution state가
아니다.
