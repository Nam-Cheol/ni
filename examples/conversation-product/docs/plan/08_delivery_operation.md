# Delivery and operation

## Delivery surfaces

- conversation

## Initial delivery

The initial delivery is locked planning material plus generated prompts for a human team and a generic/Codex downstream worker. It does not include a deployed assistant.

## Operating model

- Planning docs are committed to git.
- Contract JSON is committed to git.
- The lockfile is authoritative after `ni end`.
- Generated prompts are derived seed material and may be regenerated from the lock.
