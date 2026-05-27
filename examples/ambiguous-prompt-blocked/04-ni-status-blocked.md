# 04. ni status blocked

Run:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
```

Expected output excerpt:

```text
BLOCKED
profile: prototype
product type: software
delivery surfaces: web
interaction mode: human_to_system
blocker R009: OQ-001 is a blocker open question
blocker R009: OQ-002 is a blocker open question
```

This is the payoff: ni refuses to treat a plausible implementation guess as a
ready plan.
