# Resolution Path

이 path는 later resolved variant가 answer를 발명하거나 ni gate를 약화하지 않고
`BLOCKED`에서 가능한 `READY` 또는 `READY_WITH_DEFERRALS` 상태로 이동할 수 있는
방법을 설명한다. 이 task에서 blocker를 resolve하지 않는다.

이 path를 작성한 시점에 research-protocol benchmark는 `BLOCKED`, unlocked,
bounded prompt 없음 상태로 남아 있다.

| Step | Action | Expected result |
| --- | --- | --- |
| 1 | User answers `OQ-001`. | Research question, supported decision, and final artifact become explicit. |
| 2 | User answers `OQ-002`. | Participant/observation scope and locations become explicit. |
| 3 | User answers `OQ-003`. | Consent, privacy, data handling, translation, accessibility boundaries become explicit. |
| 4 | User answers `OQ-004`. | Field safety, stop rules, and vulnerable-group safeguards become explicit. |
| 5 | User answers `OQ-005`. | Reviewer, acceptance evidence, and pre-fieldwork readiness criteria become explicit. |
| 6 | Update isolated workspace `docs/plan` and `.ni/contract.json`. | Human-readable docs and machine-readable contract align. |
| 7 | Run `ni status` in isolated workspace. | Status may become `READY` or `READY_WITH_DEFERRALS` if no new blockers appear. |
| 8 | Run `ni end` only after user confirmation. | Isolated workspace lock may be created. |
| 9 | Run `ni run` only after lock. | Bounded handoff prompt may be compiled. |

## Required Answer Classes

| Blocker | Required answer | Expected planning update | Unsafe assumption avoided |
| --- | --- | --- | --- |
| `OQ-001` | Research question, supported decision, decision owner, final study artifact, study가 결정하지 말아야 할 것을 확인한다. | Project brief, actors/outcomes, capabilities, evaluation contract, open questions, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. | study가 intervention location을 선택하거나 full statistical report를 만들거나 모든 heat-risk question에 답한다고 추측하지 않는다. |
| `OQ-002` | Observation unit, 필요한 경우 participant group, inclusion/exclusion criteria, location 또는 neighborhood type, observation window, sampling 또는 selection rule, out-of-scope population/location을 확인한다. | Actors/outcomes, capabilities, domain state, constraints, evaluation contract, open questions, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. | Criteria 없이 vulnerable participant를 모집하거나 neighborhood를 interchangeable하게 보거나 아무 hot day/location이 valid하다고 보지 않는다. |
| `OQ-003` | Allowed/prohibited data, consent approach, privacy constraints, storage/retention rule, translation needs, accessibility needs, sensitive-data boundaries, out-of-scope data를 확인한다. | Interaction contract, domain state, constraints, risks/security, evaluation contract, open questions, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. | 불필요한 identifiable data, consent-free assumption, translation/accessibility need 무시, unbounded sensitive-note retention을 피한다. |
| `OQ-004` | Field team safety rules, heat/weather stop conditions, exposure 또는 shift limits, hydration/rest requirements, vulnerable-group safeguards, emergency escalation path, cancellation conditions를 확인한다. | Constraints, risks/security, delivery/operation, execution strategy, open questions, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. | Unsafe heat exposure, unsafe condition에서 observation 계속하기, vulnerable group을 extra safeguard 없이 다루기를 피한다. |
| `OQ-005` | Reviewer 또는 approval owner, review audience, minimum protocol artifact, acceptance evidence, pass/fail criteria, pre-fieldwork readiness checklist, required가 아닌 것, moving forward approve 권한자를 확인한다. | Actors/outcomes, evaluation contract, delivery/operation, open questions, decision log, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. | Draft를 fieldwork-ready로 보거나 informal agreement를 approval로 보거나 acceptance evidence 전에 data collection을 시작하지 않는다. |

## Refusal Conditions

Open blocker가 accepted deferral rationale 없이 남아 있거나, docs와 contract가
out of sync이거나, accepted requirement가 conflict하거나, high-severity risk에
mitigation이 없거나, required evaluation이 없거나, lock을 위한 user
confirmation이 없다면 `ni end`는 계속 refuse되어야 한다.

`.ni/plan.lock.json`이 존재하고, locked hash가 valid하며, requested target이
지원되고, compiled prompt가 configured character bound 안에 머무를 수 있을 때만
`ni run`이 허용되어야 한다. Lock 이후 intent가 바뀌면 valid lock이 amendment
또는 relock flow로 복구될 때까지 execution은 멈춰야 한다.

## Boundary

이 path는 research quality, fieldwork safety, intervention effectiveness,
downstream performance를 증명하지 않는다. 명시적인 user answer가 나중에 isolated
benchmark workspace를 다시 authoritative status check할 수 있게 만드는 방법만
설명한다.
