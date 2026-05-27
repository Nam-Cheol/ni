# 08. ni-end confirmation

## User

Invoke `ni-end`.

## Model using ni-end

I will check the readiness gate before asking for lock confirmation.

```bash
ni status --dir ./refund-triage-plan
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

## Model confirmation request

The CLI says this prototype plan can lock with visible deferrals. The lock will
hash `.ni/contract.json` and `docs/plan/**`.

Confirm that I should run `ni end --dir ./refund-triage-plan` and let the CLI
write `.ni/plan.lock.json`?

## User

Confirmed. Run `ni end`.

## CLI result

```text
locked plan at .ni/plan.lock.json
status READY_WITH_DEFERRALS
```

## Boundary shown

The model does not write the lock by hand. The CLI writes the lock after
readiness and explicit user confirmation.
