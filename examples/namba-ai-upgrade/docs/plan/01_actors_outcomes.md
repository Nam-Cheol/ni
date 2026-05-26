# Actors and outcomes

## Actors

- Planning owner: maintains the locked namba-ai upgrade intent and decides when planning changes require amendment or relock.
- NI CLI: validates readiness, writes the lockfile, verifies lock hashes, and compiles prompt seed material.
- Codex downstream user: may paste the compiled prompt into Codex manually after lock verification.
- namba-ai downstream maintainer: may inspect the locked plan and decide implementation packets outside NI.
- Human reviewer: checks non-goals, collaboration boundaries, risks, and validation coverage before execution starts.

## Outcomes

- The upgrade plan names current limitations without modifying namba-ai.
- SDD collaboration and sequence problems are described as dependency and conflict concerns.
- Codex remains a downstream consumer, not a kernel dependency or invoked runtime.
- Generated material can seed Codex, namba-ai-oriented planning, human-team review, or generated harness proposals while staying derived from the lock.
