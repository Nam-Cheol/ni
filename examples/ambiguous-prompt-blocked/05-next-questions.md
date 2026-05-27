# 05. Next questions

Run:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected focused questions:

```text
question R009 OQ-001: For OQ-001, what decision resolves this blocker, should it be deferred, or why must it remain blocking?
question R009 OQ-002: For OQ-002, what decision resolves this blocker, should it be deferred, or why must it remain blocking?
```

Human-facing version:

1. Which team is this for, and what is the first workflow the dashboard must
   support?
2. What issue data may v0 use, and what customer-data boundary must it respect?

These are narrow enough to answer, but strong enough to prevent hidden scope
creep.
