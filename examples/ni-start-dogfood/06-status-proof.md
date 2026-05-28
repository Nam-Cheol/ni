# 06. status proof

## Model using ni-start

I will run the readiness gate. The CLI result controls what happens next.

```bash
ni status --dir ./refund-triage-plan --proof --next-questions
```

## CLI result

```text
NI Intent Readiness: BLOCKED

Proof:
- OQ-001 is a blocker open question.

Execution must not start.

Next questions:
1. Which refund policy source is authoritative for refund eligibility?
```

## Model response

`ni status` reports `BLOCKED`, so I will not suggest `ni-end` yet. The next
planning turn should answer `OQ-001` or explicitly keep it blocking.

## Boundary shown

The model does not override readiness. A blocker question prevents lock even
though the docs already look useful to a human reader.
