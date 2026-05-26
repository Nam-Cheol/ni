# Actors and outcomes

## Actors

- Support agent: uses the assistant output to decide the next human action.
- Supervisor: reviews ambiguous, conflicting, or exception-heavy refund cases.
- Planning model: maintains `docs/plan/**`, `.ni/contract.json`, and bounded
  session state from conversation.
- NI CLI: validates readiness, writes the lock, and compiles prompts.

## Outcomes

- Support agents receive recommendation drafts with cited ticket facts and
  cited policy source.
- Ambiguous or conflicting cases are escalated instead of guessed.
- Planning state remains visible in docs and machine contract records.
- The user never needs to type `contract add`, `contract set`, or other
  user-facing contract authoring commands.
