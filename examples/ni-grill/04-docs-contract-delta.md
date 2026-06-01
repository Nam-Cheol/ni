# 04. Docs And Contract Delta

The model updates only the affected planning records:

- `docs/plan/02_capabilities.md`: keep `CAP-001` accepted and clarify that
  completion depends on supervisor checklist review.
- `docs/plan/05_constraints.md`: add non-goals for live ticket systems,
  queues, shell/Codex adapters, customer contact, and runtime integrations.
- `docs/plan/06_risks_security.md`: add or update a high-severity data-handling
  risk with mitigation.
- `docs/plan/07_evaluation_contract.md`: expand `EVAL-001` manual inspection
  into the supervisor checklist.
- `.ni/contract.json`: update the matching `CAP`, `REQ`, `EVAL`, `RISK`, and
  `NG` records without renumbering unrelated IDs.
- `.ni/session.json`: record that `GRILL-001` through `GRILL-004` were answered
  and that the next step is status proof.

Planning proof:

```text
Planning proof:
- User input captured:
  "Supervisor checklist review is the acceptance evidence; data handling is a high risk; live systems, queues, adapters, runtime integration, refund approval, and customer contact are excluded."
- Interpreted planning records:
  - Capabilities: CAP-001 remains accepted with supervisor checklist evidence.
  - Requirements: REQ-001 clarified around supervisor-acceptable recommendations.
  - Evaluations: EVAL-001 manual inspection now lists concrete checklist items.
  - Risks: RISK-002 data handling is high severity with synthetic/redacted document-only mitigation.
  - Non-goals: NG-002 excludes live ticket systems, queues, shell/Codex adapters, customer contact, and runtime integrations.
  - Decisions: none newly accepted beyond the clear exclusions and evidence answer.
  - Assumptions: none promoted from uncertainty.
- Updated planning artifacts:
  - docs/plan/02_capabilities.md: CAP-001 evidence clarified.
  - docs/plan/05_constraints.md: NG-002 added.
  - docs/plan/06_risks_security.md: RISK-002 added.
  - docs/plan/07_evaluation_contract.md: EVAL-001 checklist expanded.
  - .ni/contract.json: matching CAP, REQ, EVAL, RISK, and NG records updated.
  - .ni/session.json: grill findings answered.
- Status result:
  - before: pending exact CLI output from ni status --proof --next-questions.
  - after: pending exact CLI output from ni status --proof --next-questions.
- Remaining blockers:
  - determined by the next CLI status run.
- Next question group:
  - determined by the next CLI status run.
```
