# 05. Status After Grill

After updating docs, contract, and session, the model runs or requests:

```bash
ni status --dir . --proof --next-questions
```

Example summary:

```text
NI Intent Readiness: READY_WITH_DEFERRALS

Blockers:
- None.

Deferrals:
- DEC-004 remains deferred.

Warnings:
- Deferred decisions must not be treated as accepted downstream instructions.

Next questions:
Handoff deferrals:
1. DEC-004: Does this deferred decision affect the next handoff, or should it remain visible without blocking?
   Answer shape: keep deferred with reason, resolve as accepted/rejected/not_applicable, or keep blocking with reason
```

The model may explain this output, but it must not claim readiness unless this
status came from the CLI. `ni-grill` still does not lock the plan. If the user
wants to lock, the next step is `ni-end` confirmation.
