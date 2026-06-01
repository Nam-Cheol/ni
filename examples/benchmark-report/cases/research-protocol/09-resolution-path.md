# Resolution Path

This path describes how a later resolved variant could move from `BLOCKED` to a
possible `READY` or `READY_WITH_DEFERRALS` state without inventing answers or
weakening ni's gates. It does not resolve the blockers in this task.

At the time this path was created, the research-protocol benchmark remained
`BLOCKED`, unlocked, and without a compiled bounded prompt.

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
| `OQ-001` | Confirm the research question, supported decision, decision owner, final study artifact, and what the study must not decide. | Update project brief, actors/outcomes, capabilities, evaluation contract, open questions, `.ni/contract.json`, and `.ni/session.json` if present. | Avoids guessing whether the study chooses intervention locations, produces a full statistical report, or answers all heat-risk questions. |
| `OQ-002` | Confirm observation units, participant groups if any, inclusion/exclusion criteria, locations or neighborhood types, observation window, sampling or selection rule, and out-of-scope populations or locations. | Update actors/outcomes, capabilities, domain state, constraints, evaluation contract, open questions, `.ni/contract.json`, and `.ni/session.json` if present. | Avoids recruiting vulnerable participants without criteria, treating neighborhoods as interchangeable, or treating any hot day/location as valid. |
| `OQ-003` | Confirm allowed and prohibited data, consent approach, privacy constraints, storage/retention rule, translation needs, accessibility needs, sensitive-data boundaries, and out-of-scope data. | Update interaction contract, domain state, constraints, risks/security, evaluation contract, open questions, `.ni/contract.json`, and `.ni/session.json` if present. | Avoids unnecessary identifiable data, consent-free assumptions, ignored translation/accessibility needs, and unbounded sensitive-note retention. |
| `OQ-004` | Confirm field team safety rules, heat/weather stop conditions, exposure or shift limits, hydration/rest requirements, vulnerable-group safeguards, emergency escalation path, and cancellation conditions. | Update constraints, risks/security, delivery/operation, execution strategy, open questions, `.ni/contract.json`, and `.ni/session.json` if present. | Avoids unsafe heat exposure, continuing observations in unsafe conditions, and treating vulnerable groups without extra safeguards. |
| `OQ-005` | Confirm reviewer or approval owner, review audience, minimum protocol artifact, acceptance evidence, pass/fail criteria, pre-fieldwork readiness checklist, what is not required, and who may approve moving forward. | Update actors/outcomes, evaluation contract, delivery/operation, open questions, decision log, `.ni/contract.json`, and `.ni/session.json` if present. | Avoids treating a draft as fieldwork-ready, treating informal agreement as approval, or starting data collection before acceptance evidence exists. |

## Refusal Conditions

`ni end` should still be refused if any blocker remains open without an
accepted deferral rationale, if docs and contract are out of sync, if accepted
requirements conflict, if high-severity risks lack mitigation, if required
evaluations are missing, or if user confirmation for lock does not exist.

`ni run` should become allowed only after `.ni/plan.lock.json` exists, locked
hashes are valid, the requested target is supported, and the compiled prompt
can stay within the configured character bound. If intent changes after lock,
execution must stop until the amendment or relock flow restores a valid lock.

## Boundary

This path does not prove research quality, fieldwork safety, intervention
effectiveness, or downstream performance. It only describes how explicit user
answers could later make the isolated benchmark workspace eligible for another
authoritative status check.
