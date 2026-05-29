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
ni status --dir ./refund-triage-plan --proof --next-questions
```

## CLI result

```text
NI Intent Readiness: READY_WITH_DEFERRALS

Proof:
- DEC-004 is deferred.
- OQ-002 remains open.

Execution may proceed only after lock; deferrals remain explicit.

Next questions:
Handoff deferrals:
1. DEC-004: DEC-004 is deferred. Does this deferred decision affect the next handoff, or should it remain visible without blocking?
   Answer shape: affects handoff and must resolve / remains deferred with reason
2. OQ-002: OQ-002 remains open. Does this open question affect the next handoff, or should it remain visible without blocking?
   Answer shape: resolve now / defer with reason / leave open and non-blocking
```

## Boundary shown

The model can now suggest moving to `ni-end` because the CLI reports
`READY_WITH_DEFERRALS`. The deferrals remain visible instead of being hidden or
weakened.
