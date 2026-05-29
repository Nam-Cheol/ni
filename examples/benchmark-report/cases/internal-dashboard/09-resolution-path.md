# Resolution Path

This path describes how a future resolved variant could move from `BLOCKED` to
`READY` or `READY_WITH_DEFERRALS` without weakening ni's gates. It is not a
resolution of the current benchmark case. The current workspace remains
blocked, unlocked, and without a compiled prompt.

## Required Answers

| Blocker | Required answer | Expected planning update | Unsafe assumption avoided |
| --- | --- | --- | --- |
| `OQ-001` | Confirm the primary dashboard user and the specific decision the dashboard supports. | Update `docs/plan/01_actors_outcomes.md`, `docs/plan/02_capabilities.md`, `docs/plan/10_open_questions.md`, and `.ni/contract.json` with accepted actor/outcome records and any linked requirements. | Avoids guessing which customer-team role uses the dashboard or which decision it should guide. |
| `OQ-002` | Define observable "needs attention" signals, thresholds, ordering criteria, or review rules for v0. | Update `docs/plan/02_capabilities.md`, `docs/plan/04_domain_state.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, and `.ni/contract.json` with accepted signal and evaluation records. | Avoids inventing account-health metrics, ranking formulas, or alert criteria. |
| `OQ-003` | Confirm allowed source systems, account fields or field categories, freshness rules, privacy constraints, and access controls. | Update `docs/plan/04_domain_state.md`, `docs/plan/05_constraints.md`, `docs/plan/06_risks_security.md`, `docs/plan/10_open_questions.md`, and `.ni/contract.json` with accepted data-boundary, mitigation, and access records. | Avoids assuming sensitive customer data may be exposed or that stale data is acceptable. |
| `OQ-004` | Confirm the next-planning-meeting evidence: timing, audience, minimum artifact, and pass/fail acceptance checks. | Update `docs/plan/07_evaluation_contract.md`, `docs/plan/08_delivery_operation.md`, `docs/plan/10_open_questions.md`, and `.ni/contract.json` with accepted evidence and delivery-readiness records. | Avoids treating any prototype, mockup, memo, or live dashboard as sufficient without user-confirmed evidence. |

## Expected Sequence

| Step | Action | Expected result |
| --- | --- | --- |
| 1 | User answers `OQ-001`. | Primary user is explicit. |
| 2 | User answers `OQ-002`. | Attention signals and ranking criteria are explicit. |
| 3 | User answers `OQ-003`. | Source systems, privacy, and access constraints are explicit. |
| 4 | User answers `OQ-004`. | Meeting acceptance evidence is explicit. |
| 5 | Update `docs/plan/**` and `.ni/contract.json` together. | Planning docs and machine contract reflect the same accepted answers, with stable IDs preserved where possible. |
| 6 | Update risks and evaluations. | High-severity risks remain mitigated, and every capability still maps to at least one evaluation. |
| 7 | Run `ni status`. | May become `READY` or `READY_WITH_DEFERRALS` if no new blockers, conflicts, sync errors, or unmitigated high risks appear. |
| 8 | Run `ni end` only after explicit confirmation. | A lock may be created only if the CLI readiness gate allows it. |
| 9 | Run `ni run` only after a valid lock exists. | A bounded handoff prompt may be compiled, and prompt character count can then be measured. |

## Expected Contract Updates

A future resolved variant should update `.ni/contract.json` only from confirmed
user answers. Expected changes may include:

- marking answered blocker questions resolved or replacing them with accepted
  decisions, requirements, constraints, evaluations, and risks;
- adding or refining accepted requirements for actor, attention signals, data
  boundary, access control, freshness, and meeting evidence;
- adding evaluation methods for signal correctness, freshness/access review,
  usability or meeting-readiness evidence, and traceability;
- preserving high-severity risk mitigations rather than deleting risks merely
  to pass validation;
- keeping non-goals explicit so the benchmark does not become dashboard
  implementation, live integration, model execution, or downstream automation.

## Refusal Conditions

`ni-end` should still be refused if any blocker remains open without an
accepted deferral rationale, if docs and contract are out of sync, if accepted
requirements conflict, if high-severity risks lack mitigation, if required
evaluations are missing, or if the user has not confirmed the plan for lock.

`ni-run` should become allowed only after `.ni/plan.lock.json` exists, locked
hashes are valid, the requested target is supported, and the compiled prompt
can stay within the configured character bound. If intent changes after lock,
execution must stop until the amendment or relock flow restores a valid lock.

## Current Boundary

The current benchmark remains `BLOCKED`. No lock or prompt was created. The
resolution path is planning guidance for a future variant, not evidence that
the current case is ready.
