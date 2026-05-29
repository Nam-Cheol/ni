# Resolution Path

이 path는 future resolved variant가 ni gate를 약화하지 않고 `BLOCKED`에서
`READY` 또는 `READY_WITH_DEFERRALS`로 이동할 수 있는 방법을 설명한다. 현재
benchmark case를 resolve하는 문서가 아니다. 현재 workspace는 blocked,
unlocked 상태이며 compiled prompt가 없다.

## Required Answers

| Blocker | Required answer | Expected planning update | Unsafe assumption avoided |
| --- | --- | --- | --- |
| `OQ-001` | Primary dashboard user와 dashboard가 지원하는 specific decision을 확인한다. | `docs/plan/01_actors_outcomes.md`, `docs/plan/02_capabilities.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`에 accepted actor/outcome record와 linked requirement를 반영한다. | 어떤 customer-team role이 dashboard를 쓰는지 또는 어떤 decision을 guide해야 하는지 추측하지 않는다. |
| `OQ-002` | v0의 observable "needs attention" signal, threshold, ordering criteria, review rule을 정의한다. | `docs/plan/02_capabilities.md`, `docs/plan/04_domain_state.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`에 accepted signal과 evaluation record를 반영한다. | Account-health metric, ranking formula, alert criteria를 발명하지 않는다. |
| `OQ-003` | Allowed source system, account field 또는 field category, freshness rule, privacy constraint, access control을 확인한다. | `docs/plan/04_domain_state.md`, `docs/plan/05_constraints.md`, `docs/plan/06_risks_security.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`에 accepted data-boundary, mitigation, access record를 반영한다. | Sensitive customer data를 노출해도 된다고 가정하거나 stale data를 허용한다고 가정하지 않는다. |
| `OQ-004` | 다음 planning meeting evidence, 즉 timing, audience, minimum artifact, pass/fail acceptance check를 확인한다. | `docs/plan/07_evaluation_contract.md`, `docs/plan/08_delivery_operation.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`에 accepted evidence와 delivery-readiness record를 반영한다. | Prototype, mockup, memo, live dashboard 중 아무거나 충분하다고 user confirmation 없이 판단하지 않는다. |

## Expected Sequence

| Step | Action | Expected result |
| --- | --- | --- |
| 1 | User answers `OQ-001`. | Primary user is explicit. |
| 2 | User answers `OQ-002`. | Attention signals and ranking criteria are explicit. |
| 3 | User answers `OQ-003`. | Source systems, privacy, and access constraints are explicit. |
| 4 | User answers `OQ-004`. | Meeting acceptance evidence is explicit. |
| 5 | `docs/plan/**`와 `.ni/contract.json`을 함께 update한다. | Planning docs와 machine contract가 같은 accepted answer를 반영하고, 가능한 한 stable ID를 보존한다. |
| 6 | Risk와 evaluation을 update한다. | High-severity risk는 계속 mitigated 상태이고, 모든 capability는 적어도 하나의 evaluation에 mapping된다. |
| 7 | `ni status`를 실행한다. | 새로운 blocker, conflict, sync error, unmitigated high risk가 없으면 `READY` 또는 `READY_WITH_DEFERRALS`가 될 수 있다. |
| 8 | Explicit confirmation 뒤에만 `ni end`를 실행한다. | CLI readiness gate가 허용할 때만 lock이 생성될 수 있다. |
| 9 | Valid lock이 존재한 뒤에만 `ni run`을 실행한다. | Bounded handoff prompt가 compile될 수 있고 prompt character count를 측정할 수 있다. |

## Expected Contract Updates

Future resolved variant는 confirmed user answer에서만 `.ni/contract.json`을
update해야 한다. Expected change는 다음을 포함할 수 있다.

- answered blocker question을 resolved로 표시하거나 accepted decision,
  requirement, constraint, evaluation, risk로 대체한다;
- actor, attention signal, data boundary, access control, freshness, meeting
  evidence에 대한 accepted requirement를 추가하거나 refine한다;
- signal correctness, freshness/access review, usability 또는 meeting-readiness
  evidence, traceability를 위한 evaluation method를 추가한다;
- validation을 통과하기 위해 risk를 삭제하지 않고 high-severity risk mitigation을
  보존한다;
- benchmark가 dashboard implementation, live integration, model execution,
  downstream automation으로 변하지 않도록 non-goal을 명시적으로 유지한다.

## Refusal Conditions

Open blocker가 accepted deferral rationale 없이 남아 있거나, docs와 contract가
out of sync이거나, accepted requirement가 conflict하거나, high-severity risk에
mitigation이 없거나, required evaluation이 없거나, user가 lock을 확인하지
않았다면 `ni-end`는 계속 refuse되어야 한다.

`.ni/plan.lock.json`이 존재하고, locked hash가 valid하며, requested target이
지원되고, compiled prompt가 configured character bound 안에 머무를 수 있을
때만 `ni-run`이 허용되어야 한다. Lock 이후 intent가 바뀌면 valid lock이
amendment 또는 relock flow로 복구될 때까지 execution은 멈춰야 한다.

## Current Boundary

현재 benchmark는 `BLOCKED`로 남아 있다. Lock이나 prompt는 생성되지 않았다.
이 resolution path는 future variant를 위한 planning guidance이지 현재 case가
ready라는 evidence가 아니다.
