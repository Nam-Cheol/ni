# 05. docs and contract delta

## Model using ni-start

I will update the human docs and machine contract together from the confirmed
answers.

Changed planning files:

- `docs/plan/00_project_brief.md`: add refund triage purpose, who it serves,
  and why it matters.
- `docs/plan/01_actors_outcomes.md`: add support agent, supervisor, planning
  model, and NI CLI roles.
- `docs/plan/02_capabilities.md`: add `CAP-001`, `CAP-002`, and `CAP-003`.
- `docs/plan/03_interaction_contract.md`: capture conversation and refund
  triage boundaries.
- `docs/plan/08_delivery_operation.md`: record `conversation` and `document`
  as the delivery surfaces.
- `docs/plan/05_constraints.md`: add non-goals for refunds, customer contact,
  and execution runtime.
- `docs/plan/06_risks_security.md`: add refund authority and data risks with
  mitigations.
- `docs/plan/07_evaluation_contract.md`: add transcript fixture evaluations.
- `docs/plan/10_open_questions.md`: add blocker `OQ-001`.
- `.ni/contract.json`: add `project.purpose`, `product_type`,
  `delivery_surfaces`, and matching `CAP`, `REQ`, `EVAL`, `RISK`, `ART`,
  `NG`, and `OQ` records.
- `.ni/session.json`: refresh bounded continuity state below contract
  authority, including active readiness profile, product type, delivery
  surfaces, recent decisions, pending questions, and last readiness status.

## Planning proof shown to the user

```text
Planning proof:
- User input captured:
  "Refund triage should become a documented recommendation workflow for support agents and supervisors, delivered as a conversation product with a document handoff, with refunds/customer contact/runtime execution excluded."
- Interpreted planning records:
  - Purpose: refund triage planning should move ad hoc judgment into a lockable recommendation workflow.
  - Actors/outcomes: support agents get draft recommendations; supervisors get ambiguity escalations; the planning model keeps docs and contract synchronized; the NI CLI validates readiness.
  - Delivery surface: conversation and document.
  - Capabilities: CAP-001 draft refund recommendations; CAP-002 escalate ambiguous or conflicting cases.
  - Requirements: REQ-001 recommendations only; REQ-002 escalation on policy ambiguity.
  - Risks: RISK-001 refund authority confusion; RISK-002 customer or policy data mishandling.
  - Evaluations: EVAL-001 and EVAL-002 transcript fixtures.
  - Decisions: none newly accepted beyond the clear scope exclusions.
  - Assumptions: none promoted to accepted decisions.
  - Non-goals: NG-001 no refund issuing or approval; NG-003 no execution runtime, shell adapter, Codex adapter, queue, or live support integration.
  - Open questions: OQ-001 authoritative refund policy source remains blocking.
- Updated planning artifacts:
  - docs/plan/00_project_brief.md: refund triage purpose recorded.
  - docs/plan/01_actors_outcomes.md: support agent, supervisor, planning model, and NI CLI roles recorded.
  - docs/plan/02_capabilities.md: CAP-001 and CAP-002 recorded.
  - docs/plan/03_interaction_contract.md: conversation and recommendation boundaries recorded.
  - docs/plan/05_constraints.md: refund, customer contact, and runtime non-goals recorded.
  - docs/plan/06_risks_security.md: refund authority and data risks recorded.
  - docs/plan/07_evaluation_contract.md: transcript fixture evaluations recorded.
  - docs/plan/08_delivery_operation.md: conversation and document delivery surfaces recorded.
  - docs/plan/10_open_questions.md: OQ-001 recorded as blocking.
  - .ni/contract.json: project.purpose, product_type, delivery_surfaces, CAP, REQ, EVAL, RISK, ART, NG, and OQ records updated.
  - .ni/session.json: active focus, pending questions, and last readiness status refreshed.
- Status result:
  - before: BLOCKED because R014, R015, R016, and OQ-001 were unresolved.
  - after: BLOCKED because OQ-001 remains a blocker.
- Remaining blockers:
  - OQ-001: authoritative refund policy source is still unresolved.
- Next question group:
  - Open blockers.
```

## Example synchronized records

Human docs describe:

```text
CAP-001: Draft refund recommendations
Status: accepted
Linked records: REQ-001, REQ-002, EVAL-001, RISK-001, RISK-002, ART-001
```

The contract contains corresponding machine records:

```json
{
  "product_type": "conversation_product",
  "delivery_surfaces": ["conversation", "document"],
  "project": {
    "purpose": "Plan a support-agent assistant that drafts refund recommendations from tickets and policy, escalates ambiguity, and excludes refund approval, customer contact, and runtimes."
  },
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "Draft refund recommendations for support agents.",
      "status": "accepted",
      "requirements": ["REQ-001", "REQ-002"],
      "evaluations": ["EVAL-001"],
      "risks": ["RISK-001", "RISK-002"],
      "artifacts": ["ART-001"]
    }
  ]
}
```

## Boundary shown

The user still does not edit JSON by hand or type contract authoring commands.
The model-maintained contract is the machine-readable companion to the docs,
and the CLI will validate it next.
