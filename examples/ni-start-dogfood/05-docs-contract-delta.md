# 05. docs and contract delta

## Model using ni-start

I will update the human docs and machine contract together from the confirmed
answers.

Changed planning files:

- `docs/plan/00_project_brief.md`: add refund triage purpose and success
  definition.
- `docs/plan/01_actors_outcomes.md`: add support agent, supervisor, planning
  model, and NI CLI roles.
- `docs/plan/02_capabilities.md`: add `CAP-001`, `CAP-002`, and `CAP-003`.
- `docs/plan/03_interaction_contract.md`: capture conversation and refund
  triage boundaries.
- `docs/plan/05_constraints.md`: add non-goals for refunds, customer contact,
  and execution runtime.
- `docs/plan/06_risks_security.md`: add refund authority and data risks with
  mitigations.
- `docs/plan/07_evaluation_contract.md`: add transcript fixture evaluations.
- `docs/plan/10_open_questions.md`: add blocker `OQ-001`.
- `.ni/contract.json`: add matching `CAP`, `REQ`, `EVAL`, `RISK`, `ART`,
  `NG`, and `OQ` records.
- `.ni/session.json`: refresh bounded continuity state below contract
  authority, including active readiness profile, product type, delivery
  surfaces, recent decisions, pending questions, and last readiness status.

## Example synchronized records

Human docs describe:

```text
CAP-001: Draft refund recommendations
Status: accepted
Linked records: REQ-001, REQ-002, EVAL-001, RISK-001, RISK-002, ART-001
```

The contract contains the corresponding machine record:

```json
{
  "id": "CAP-001",
  "title": "Draft refund recommendations for support agents.",
  "status": "accepted",
  "requirements": ["REQ-001", "REQ-002"],
  "evaluations": ["EVAL-001"],
  "risks": ["RISK-001", "RISK-002"],
  "artifacts": ["ART-001"]
}
```

## Boundary shown

The user still does not edit JSON by hand or type contract authoring commands.
The model-maintained contract is the machine-readable companion to the docs,
and the CLI will validate it next.
