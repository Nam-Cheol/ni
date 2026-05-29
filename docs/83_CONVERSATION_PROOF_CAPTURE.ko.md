# Conversation proof capture

Conversation proof capture는 `ni-start`가 의미 있는 authoring update 뒤에 보여야
하는 짧은 audit trail이다. 사용자의 최신 답변이 planning state를 어떻게 바꾸었는지,
어떤 file과 contract record가 바뀌었는지, 그리고 edit 전후에 CLI readiness gate가
무엇을 말했는지를 설명한다.

목적은 사용자가 model vibes를 믿지 않고 planning progress를 검토할 수 있게 하는
것이다. 이 proof는 intent authoring에 대한 것이다. Execution evidence가 아니며,
downstream work를 실행하지 않고, model을 readiness authority로 만들지 않는다.

## Planning proof가 하는 일

Planning proof는 한 authoring turn에 대한 user-visible summary다.

- 사용자의 답변을 짧게 paraphrase한다.
- model이 record type별로 무엇을 planning state로 해석했는지 보여준다.
- 실제로 변경된 planning artifact를 보여준다.
- 실제로 변경된 contract field 또는 ID를 보여준다.
- turn에서 영향을 받은 decision, assumption, non-goal, open question을 보여준다.
- edit 전후의 `ni status --proof --next-questions` 결과를 보여준다.
- CLI가 반환한 다음 highest-priority question group을 보여준다.

간결해야 한다. Hidden chain-of-thought를 노출하면 안 된다. 실제로 일어나지 않은
file change, contract ID, readiness, lock state를 주장하면 안 된다.

## Execution evidence와의 차이

Planning proof는 conversation이 docs와 contract로 이동한 흔적을 기록한다. 이것은
"어떤 intent가 바뀌었는가?"와 "readiness gate가 무엇을 말했는가?"에 답한다.

Execution evidence는 implementation이 올바르게 실행됐는지를 다룬다. 그것은
`ni-kernel` 밖의 영역이다. Planning proof를 만들기 위해 runtime execution,
downstream agents, shell adapters, queues, or an execution evidence loop must
not be added.

## ni-start 필수 block

의미 있는 authoring update 뒤 `ni-start`는
`ni status --dir . --proof --next-questions`를 다시 실행하거나 요청한 다음, 다음
형태의 block을 보고해야 한다.

```text
Planning proof:
- User input captured:
  "<short paraphrase of user answer>"
- Interpreted planning records:
  - Purpose: ...
  - Actors/outcomes: ...
  - Delivery surface: ...
  - Capabilities: CAP-001 ...
  - Requirements: REQ-001 ...
  - Risks: RISK-001 ...
  - Evaluations: EVAL-001 ...
  - Decisions: DEC-001 accepted/deferred/rejected if applicable
  - Assumptions: ASM-001 or open question if applicable
  - Non-goals: NG-001 if applicable
  - Open questions: OQ-001 ...
- Updated planning artifacts:
  - docs/plan/00_project_brief.md: purpose clarified
  - docs/plan/01_actors_outcomes.md: primary actors added
  - docs/plan/03_interaction_contract.md: delivery surface recorded
  - .ni/contract.json: project.purpose, actors/outcomes, delivery_surfaces updated
  - .ni/session.json: active focus and pending questions updated
- Status result:
  - before: BLOCKED because R014/R015/R016
  - after: BLOCKED/READY_WITH_DEFERRALS/READY because ...
- Remaining blockers:
  - OQ-001 ...
- Next question group:
  - Sync repairs / Risk decisions / Evaluation evidence / Open blockers / none
```

변경되지 않은 record type은 생략하거나 `none`이라고 쓴다. 변경된 file이 없다면
`No planning artifacts were updated.`라고 쓴다. Command execution이 불가능하면
before 또는 after status가 사용자나 trusted runner의 exact CLI output을 기다리고
있다고 표시해야 한다.

## 사용자가 읽는 방법

사용자는 proof를 다음 순서로 확인한다.

1. Paraphrase가 사용자의 의도와 맞는가?
2. Tentative statement가 assumption 또는 open question으로 남아 있는가?
3. Clear exclusion이 non-goal로 기록되었는가?
4. Changed files와 contract fields가 stated interpretation과 맞는가?
5. After-status가 `ni status --proof --next-questions`에서 온 것인가?
6. Next question group이 CLI의 highest-priority group인가?

Proof가 docs와 contract disagreement를 말하면, 다음 turn은 stale side를 고치거나
그 disagreement를 blocker로 유지해야 한다. Sync diagnostic이 readiness를 막고
있는 동안 `ni-end`로 진행하면 안 된다.

## CLI validation 없이 신뢰하면 안 되는 것

Model-only proof block을 readiness, lock, handoff authority로 신뢰하면 안 된다.
Model은 edit을 summarize할 수 있지만 권한은 다음에 있다.

- `ni status`가 `BLOCKED`, `READY_WITH_DEFERRALS`, `READY`를 결정한다.
- `ni end`가 `.ni/plan.lock.json`을 생성한다.
- `ni run`이 lock hash를 검증하고 bounded prompt를 compile한다.

No-terminal mode에서 planning proof는 draft audit trail이다. Model이 무엇을 바꾸려
했는지 검토하는 데는 유용하지만, drafted docs와 contract를 CLI run이 validate한
뒤에만 trusted 상태가 된다.
