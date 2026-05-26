# Execution strategy

## Target

The first downstream target is `human-team`, selected in `DEC-003`.

## Handoff material

`ni run --target human-team --max-chars 4000` compiles a prompt that points the
team to the locked NI sources:

- `.ni/plan.lock.json`
- `.ni/contract.json`
- `docs/plan/**`

## Boundary

The compiled prompt is seed material only. It does not start an agent, run shell
commands, execute Codex, create adapters, or manage a task queue.
