# 04. user answers

## User

This project should change refund triage from an ad hoc support-agent judgment
call into a documented recommendation workflow. It is for support agents who
need a draft recommendation, supervisors who review ambiguous cases, and the
planning model that must keep the intent lockable before any downstream work.

Deliver it as a conversation product with a document handoff. It should only
draft recommendations for support agents. It must never issue refunds, approve
refunds, or contact customers.

Readiness for the recommendation behavior should eventually be proven with
three transcript fixtures:

- eligible refund,
- ineligible refund,
- policy ambiguity that escalates to a supervisor.

Also keep downstream execution out of scope for this plan. I only want the
kernel planning flow, lock, and prompt handoff.

## Accepted from this answer

- `project.purpose`: plan a support-agent refund triage assistant that turns
  ticket facts and policy into draft recommendations while keeping lock and
  handoff authority in NI.
- Actors and outcomes: support agent receives a draft recommendation;
  supervisor receives escalations for ambiguous or conflicting cases; planning
  model maintains synchronized docs and contract records; NI CLI validates
  readiness before lock.
- Delivery surfaces: `conversation` and `document`.
- `CAP-001`: draft refund recommendations.
- `CAP-002`: escalate ambiguous or conflicting cases.
- `NG-001`: do not issue, approve, or initiate refunds.
- `NG-003`: do not add an execution runtime, shell adapter, Codex adapter,
  task queue, or live support integration.
- `EVAL-001` and `EVAL-002`: transcript fixture evaluations.

## Still open

- `OQ-001`: which refund policy source is authoritative?

The model can now update docs and contract records for purpose, actors,
outcomes, delivery surface, initial capabilities, evaluations, and non-goals.
`OQ-001` remains a blocker until answered.
