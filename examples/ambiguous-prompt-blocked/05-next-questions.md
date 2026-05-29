# 05. Next questions

Run:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --proof --next-questions
```

Expected focused questions:

```text
NI Intent Readiness: BLOCKED

Next questions:
Open blockers:
1. OQ-001: OQ-001 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
   Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason
2. OQ-002: OQ-002 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
   Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason
```

Human-facing version:

1. Which team is this for, and what is the first workflow the dashboard must
   support?
2. What issue data may v0 use, and what customer-data boundary must it respect?

These are narrow enough to answer, but strong enough to prevent hidden scope
creep.
