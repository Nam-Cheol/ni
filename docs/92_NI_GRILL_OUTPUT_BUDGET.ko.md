# NI Grill Output Budget

`ni-grill` critique는 사용자가 한 turn 안에 답할 수 있을 만큼 짧을 때 유용하다.
이 문서는 grill output의 severity model, output budget, finding format을 정의한다.

Boundary는 그대로다:

```text
Skills are UX; CLI is authority.
```

`ni-grill`은 advisory planning pressure를 줄 수 있다. 하지만 second readiness
gate가 되거나, lock을 approve하거나, model judgment만으로 lock을 거부하거나,
generated prompts를 실행하거나, plan을 ready처럼 보이게 하려고 acceptance
criteria를 약화해서는 안 된다.

## Why The Budget Exists

Planning critique는 intent를 해결하는 속도보다 빨리 커질 수 있다. 같은 무게의 긴
concern 목록은 사용자가 `ni-end` 전에 무엇을 고쳐야 하는지 보기 어렵게 만든다.

Output budget은 `ni-grill`이 다음을 지키게 한다:

- deterministic `ni status` blockers를 먼저 보존한다;
- editorial issue보다 highest-risk planning concerns를 먼저 보여준다;
- 질문을 한 turn에 답할 수 있게 유지한다;
- CLI blockers와 advisory findings를 구분한다;
- lower-priority issues는 쏟아내지 않고 요약한다.

## Severity Model

Severity는 `ni-grill`의 advisory planning pressure다. CLI readiness와 같지 않다.
Lock은 `ni status`, user confirmation, `ni end`만 govern한다.

다음 labels를 그대로 사용한다:

| Severity | Meaning | Examples |
| --- | --- | --- |
| `Critical` | CLI가 완전히 잡지 못할 수 있는 serious planning-quality 또는 claim-boundary problem이라서 `ni-end` 전에 해결해야 할 가능성이 높다. | Evidence가 뒷받침하지 않는 public claim; privacy/safety boundary ambiguity; evaluation은 있지만 accepted capability의 evidence wording이 약함; benchmark가 `READY`지만 `not_measured` boundary가 숨겨짐; downstream handoff가 implementation authorization처럼 읽힘. |
| `High` | Lock 전에 중요하지만 user confirmation, deferral, explicit non-goal로 해결할 수 있다. | Actor/outcome이 너무 넓음; prose의 delivery surface ambiguity; risk mitigation이 vague함; target handoff가 불명확함. |
| `Medium` | Clarity를 높이고 rework를 줄이지만, 사용자가 tradeoff를 받아들이면 next lock을 막지 않을 수 있다. | Wording을 더 precise하게 만들 수 있음; evidence가 reviewer를 더 명확히 이름 붙일 수 있음; assumptions를 open questions로 올리는 것이 좋음. |
| `Low` | Editorial 또는 maintainability improvement. | Duplicate wording; docs organization; minor example clarity. |
| `Note` | Observation only. 사용자가 refine하고 싶지 않으면 question이 필요 없다. | 잘 처리된 non-goal; clear claim boundary; 반복할 만한 useful pattern. |

## Output Budget

기본적으로 `ni-grill`은 한 turn에 최대 5 findings만 보여준다.

Rules:

- Advisory findings보다 `ni status`의 deterministic CLI blockers를 먼저 보여준다.
- `Critical` 또는 `High` findings가 있으면 최대 3개까지만 먼저 보여준다.
- 한 turn에 user-facing questions를 5개 넘게 묻지 않는다.
- 꼭 필요하지 않으면 너무 많은 categories를 섞지 않는다.
- 설명에 도움이 될 때가 아니면 deterministic blocker 전체를 반복하지 않는다.
- 더 많은 issues가 있으면 `N additional lower-priority findings were not shown.`라고 쓴다.

`ni status`가 `BLOCKED`이면 deterministic blockers를 우선하고 secondary grill
critique는 짧게 유지한다. `ni status`가 `READY` 또는 `READY_WITH_DEFERRALS`이면
claim quality, public handoff, risk clarity, overclaim prevention에 집중한다.

## Prioritization

Findings는 다음 순서로 정렬한다:

1. `ni status`의 deterministic CLI blockers;
2. `Critical` ni-grill findings;
3. `High` ni-grill findings;
4. acceptance evidence gaps;
5. privacy/security/safety risks;
6. claim-boundary risks;
7. non-goal 또는 scope-drift risks;
8. handoff ambiguity;
9. `Medium` 또는 `Low` editorial issues;
10. `Note` observations.

## Finding Format

다음 shape를 사용한다:

```text
Grill findings:
1. GRILL-001 — <severity> — <category>
   Affected: <file path or planning ID>
   Concern: <specific concern>
   Why it matters: <why downstream handoff or lock quality could suffer>
   Question: <user-facing question>
   Answer shape: <expected answer form>
   Suggested action: <resolve / defer / mark non-goal / clarify / keep as note>
   Blocks ni-end: <CLI decides / likely yes / maybe / no>
```

Finding이 deterministic readiness와 대응하면 `Blocks ni-end: CLI decides`를
사용한다. Lock 전에 해결해야 할 severe planning-quality issue에만
`Blocks ni-end: likely yes`를 사용한다. User-confirmable tradeoff에는
`Blocks ni-end: maybe`를 사용한다. Clarity, editorial, observation findings에는
`Blocks ni-end: no`를 사용한다.

## Language Behavior

User-facing questions는 사용자의 latest substantive language로 묻는다. IDs,
commands, paths, schema keys, target names, status constants, severity labels는
그대로 보존한다. 주변 prose에 번역을 덧붙일 수는 있다. 예:
`GRILL-001`, `R014`, `OQ-001`, `SYNC-014`, `ni status`,
`.ni/contract.json`, `READY`, `READY_WITH_DEFERRALS`, `BLOCKED`, `Critical`,
`High`, `Medium`, `Low`, `Note`.

## Good And Bad Output

Bad:

- 같은 weight로 12 findings를 나열한다.
- "Can you clarify this?"처럼 generic question을 묻는다.
- 모든 `ni status` blockers를 반복한 뒤 관련 없는 critiques를 많이 추가한다.
- Model judgment만으로 "do not lock"이라고 말한다.
- Planning challenge에 execution advice를 섞는다.

Good:

- 충분하면 top 3 findings만 보여준다.
- Severity와 category를 label한다.
- Specific user-facing question과 answer shape를 제공한다.
- CLI authority를 보존한다.
- 생략한 lower-priority findings를 요약한다.

## Before And After

Bad:

```text
GRILL-001: Evidence is vague. Can you clarify?
GRILL-002: Risk is unclear. Can you clarify?
GRILL-003: Handoff is unclear. Can you clarify?
GRILL-004: Non-goal is unclear. Can you clarify?
GRILL-005: Claim is unclear. Can you clarify?
GRILL-006: The docs could be better.
```

Good:

```text
Grill findings:
1. GRILL-001 — Critical — claim boundary
   Affected: examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md
   Concern: `READY` transition can be quoted without the nearby
   `not_measured` research-approval boundary.
   Why it matters: Reader가 benchmark artifact readiness를 real fieldwork
   authorization으로 오해할 수 있다.
   Question: Transition row에 real research approval, fieldwork authorization,
   research quality, intervention effectiveness가 여전히 `not_measured`라고
   써야 하는가?
   Answer shape: yes/no plus exact row wording, 또는 existing scope note가
   충분하다는 rationale.
   Suggested action: clarify
   Blocks ni-end: maybe

2 additional lower-priority findings were not shown.
```

## CLI Authority Boundary

`ni-grill`은 어떤 finding이 lock 전에 해결되는 편이 좋다고 말할 수 있지만,
readiness와 locking은 여전히 CLI로 route해야 한다:

```text
ni status --dir . --proof --next-questions
ni end --dir .
```

`ni status`가 `BLOCKED`를 보고하면 deterministic blockers를 보고하고 `ni-end`
전에 멈춘다. Lock hash mismatch가 있으면 멈추고 `BLOCKED`를 보고한다.
