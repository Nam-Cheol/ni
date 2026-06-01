# 01. Draft Plan

The draft project is a conversation product that helps support supervisors
review refund triage recommendations before a handoff prompt is locked.

Excerpted planning state:

```text
CAP-001: Draft refund triage recommendations
Status: accepted
Linked records: REQ-001, EVAL-001, RISK-001

REQ-001: Recommendations should be useful for support work.

EVAL-001: Review the output.
Method: manual inspection

RISK-001: Refund authority confusion
Severity: high
Mitigation: Supervisors review recommendations before action.

NG-001: The product does not issue refunds or contact customers.

DEC-001: Use conversation and document surfaces.
Status: accepted
```

The draft is close enough to challenge because it has accepted capability,
requirement, evaluation, risk, non-goal, and decision records.

It is still weak:

- `REQ-001` says "useful" without defining the approving actor.
- `EVAL-001` names manual inspection but not the review checklist.
- `RISK-001` mitigates refund authority but not policy-data handling.
- `NG-001` excludes execution actions but does not mention queue or adapter
  scope drift.
- `DEC-001` names surfaces but does not say which artifact is handed off.
