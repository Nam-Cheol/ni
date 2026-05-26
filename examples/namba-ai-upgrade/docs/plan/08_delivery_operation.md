# Delivery and operation

## Delivery surfaces

- CLI: NI commands validate, lock, and compile planning output.
- Document: planning docs, contract JSON, lockfile, and generated prompt are file artifacts.
- Workflow: downstream work can derive proposals after lock verification.

## Initial delivery

The initial delivery for this example is a locked planning contract plus `generated/codex.goal.prompt.txt`.

## Downstream targets

- Codex target: compiled prompt for manual downstream Codex use.
- namba-ai target: future seed guidance for namba-ai-oriented planning, not namba-ai execution.
- Human-team target: reviewer handoff for scope, risk, validation, and ownership.
- Generated harness proposal: optional derived work graph and evidence plan outside NI-owned runtime state.

## Operating model

- Planning docs are committed to git.
- Contract JSON is committed to git.
- `.ni/plan.lock.json` is written only by `ni end`.
- Generated files are derived from a valid lock and may be regenerated.
