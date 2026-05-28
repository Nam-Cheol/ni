# Expected Hidden Assumptions: Refund Triage Assistant

- The assistant only drafts recommendations and does not approve, issue, or
  initiate refunds.
- Support agents, not customers, are the primary users.
- A specific internal refund policy source exists and is authoritative.
- The assistant can safely use ticket facts without exposing unnecessary
  customer data.
- Ambiguous policy cases should escalate to a supervisor instead of inventing a
  policy interpretation.
- Transcript fixtures are enough to evaluate first-pass readiness.
- The first handoff target is a planning team, not a live support runtime.
