# Constraints

## Hard constraints

- `ni run` prompt output must be 4000 characters or less when `--max-chars 4000` is used.
- Readiness must be rule-based, not model-feeling-based.
- Lockfile hash mismatch must block prompt compilation and target export.
- Codex is an adapter or UX target, not the kernel.
- Target exports must be seed material only.
- Feedback, pressure, and harness candidates must not silently change locked planning docs.
- Relock must require an explicit applied amendment when a prior lock exists.
- Collaboration checks must be deterministic and contract-local.
- After `ni init`, v0.2 authoring must flow through model-user conversation that updates docs and contract together.
- User-facing contract `add`, `list`, or `set` commands must not become the v0.2 primary authoring UX.
- Differentiation proof assets must remain pre-runtime evidence: demos, benchmark protocols, proof reports, target stories, README sync, and release checklists must not execute downstream agents or become kernel-owned runtime state.

## Kernel boundary

`ni-kernel` owns docs contract, readiness gate, lockfile, prompt compiler, target registry, inert feedback and pressure ledgers, amendment/relock, and collaboration conflict checks.

`ni-generated-harness` owns project-specific work graphs, evaluation plans, evidence rules, adapters, and runtime decisions.

## Forbidden v0.2 behavior

- Do not add a shell adapter.
- Do not add a Codex execution adapter.
- Do not add an evidence runner.
- Do not add a queue.
- Do not add PR automation.
- Do not add release automation.
- Do not add a plugin system.
- Do not add a TUI or web UI.
- Do not add primary contract editing commands that make users hand-maintain `.ni/contract.json`.
