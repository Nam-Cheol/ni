# 06. status proof

## Model using ni-start

I will run the readiness gate. The CLI result controls what happens next.

```bash
ni status --dir ./refund-triage-plan --proof --next-questions
```

## CLI result

```text
NI Intent Readiness: BLOCKED

Blockers:
- OQ-001 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.

Passed checks:
- Required docs exist.
- Contract JSON is valid.
- Project purpose is recorded.
- Actors and outcomes are recorded.
- Delivery surface is recorded.

Execution must not start.

Next questions:
1. OQ-001 is blocking readiness. What answer would resolve it: an accepted decision, a deferral with reason, not_applicable, or keeping it blocking with the missing information named?
```

## Model response

`ni status` reports `BLOCKED`, so I will not suggest `ni-end` yet. The
first-run blockers are resolved, and the next planning turn should answer
`OQ-001` or explicitly keep it blocking.

## Boundary shown

The model records the user's first-run answers, runs status again, and does not
override readiness. A blocker question prevents lock even though the first
purpose, actor/outcome, and delivery-surface gaps are now filled.
