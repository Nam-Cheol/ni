# Constraints

## Non-goals

- NG-001: Do not issue refunds, approve refunds, or contact customers.
- NG-002: Do not replace supervisor judgment when policy is ambiguous or ticket
  facts conflict with policy.
- NG-003: Do not add an execution runtime, shell adapter, Codex adapter, task
  queue, or live support integration.

## Planning constraints

- The user does not type contract authoring commands.
- The model keeps `docs/plan/**` and `.ni/contract.json` synchronized.
- The CLI validates readiness and writes locks.
- `ni run` compiles prompts only; it does not execute Codex, shell commands,
  queues, adapters, or implementation work.
