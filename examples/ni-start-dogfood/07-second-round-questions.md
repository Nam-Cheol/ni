# 07. second round questions

## Model using ni-start

Focused question from the readiness gate:

1. For `OQ-001`, which refund policy source should `REQ-001`, `REQ-002`,
   `EVAL-001`, and `EVAL-002` treat as authoritative?

## User

Use the internal refund policy page as the source of truth. If it conflicts
with the ticket text, escalate to a supervisor.

## Model using ni-start

I will resolve `OQ-001`, add `DEC-001`, update the escalation requirement, add
the stale-policy risk, and run `ni status` again.

Changed records:

- `DEC-001`: internal refund policy page is authoritative.
- `REQ-003`: escalate policy ambiguity or ticket-policy conflict.
- `RISK-003`: stale or unclear policy page can create incorrect
  recommendations.
- `OQ-001`: resolved by `DEC-001`.
- `OQ-002`: support dashboard target remains open but non-blocking.

```bash
ni status --dir ./refund-triage-plan --next-questions
```

## CLI result

```text
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document

deferral D001: DEC-004 is deferred
deferral D002: OQ-002 remains open
```

## Boundary shown

The model can now suggest moving to `ni-end` because the CLI reports
`READY_WITH_DEFERRALS`. The deferrals remain visible instead of being hidden or
weakened.
