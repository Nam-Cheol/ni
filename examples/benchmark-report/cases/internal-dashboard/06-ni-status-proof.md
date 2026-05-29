# ni Status Proof

Command run from the repository root on 2026-05-29:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
```

Exact output:

```text
NI Intent Readiness: BLOCKED

Blockers:
- OQ-001 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.
- OQ-002 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.
- OQ-003 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.
- OQ-004 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.

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
- At least one non-goal is recorded.
- Docs and contract are synchronized.

Execution must not start.

Next questions:
Open blockers:
1. OQ-001: OQ-001 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
   Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason
2. OQ-002: OQ-002 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
   Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason
3. OQ-003: OQ-003 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
   Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason

1 additional lower-priority question(s) remain after these top 3.
```

Result: the benchmark workspace remains `BLOCKED`. No lock was created and no
handoff prompt was compiled.
