# 04. user answers

## User

It only drafts recommendations for support agents. It must never issue refunds
or contact customers.

Readiness should be proven with three transcript fixtures:

- eligible refund,
- ineligible refund,
- policy ambiguity that escalates to a supervisor.

Also keep downstream execution out of scope for this plan. I only want the
kernel planning flow, lock, and prompt handoff.

## Accepted from this answer

- `CAP-001`: draft refund recommendations.
- `CAP-002`: escalate ambiguous or conflicting cases.
- `NG-001`: do not issue, approve, or initiate refunds.
- `NG-003`: do not add an execution runtime, shell adapter, Codex adapter,
  task queue, or live support integration.
- `EVAL-001` and `EVAL-002`: transcript fixture evaluations.

## Still open

- `OQ-001`: which refund policy source is authoritative?

The model can now update docs and contract records, but `OQ-001` remains a
blocker until answered.
