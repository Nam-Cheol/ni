# Resolved status proof

Command:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions
```

Output:

```text
NI Intent Readiness: READY

Blockers:
- None.

Deferrals:
- None.

Warnings:
- None.

Passed checks:
- Required docs exist.
- Contract JSON is valid.
- Readiness profile definitions are valid.
- Capability and evaluation traceability rules passed.
- High-severity risks have mitigation.
- Decision statuses are valid and accepted decisions do not conflict.
- No blocker open questions are present.
- At least one non-goal is recorded.
- Docs and contract are synchronized.

Execution may proceed only after lock.
```

Interpretation: `OQ-001` through `OQ-005` were resolved inside the isolated
benchmark workspace using synthetic benchmark fixture answers. This result is
not real fieldwork approval, not actual research authorization, not proof of
research quality, and not empirical evidence.
