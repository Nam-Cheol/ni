# Execution strategy

## Kernel strategy

`ni` does not execute implementation work. It compiles locked planning context into bounded prompts and seed artifacts.

For v0.2, `ni run` compiles a handoff prompt only. It must not call Codex, shell commands, adapters, queues, PR automation, or downstream runtimes.

## Target strategy

`ni run --target <target>` produces a prompt for a downstream actor. The built-in prompt targets are:

```text
generic
codex
human-team
hyper-run
namba-ai
ouroboros
spec-kit
```

`ni export --target <target> --out <dir>` may write target-specific seed packages for downstream tools after lock verification in later seed workflows; it is not the primary v0.2 authoring path.

## Feedback strategy

Downstream runs may produce observations, requested changes, or harness ideas. `ni feedback add` records those observations without changing accepted planning state.

## Pressure strategy

Pressure items summarize repeated or important planning pressure. A pressure item is not accepted scope until the user promotes it and an amendment changes the contract.

## Harness strategy

Generated harness candidates are proposals for downstream work. They must remain derived, require validation evidence, and never execute from inside `ni-kernel`.

## Amendment strategy

Locked docs are changed through explicit amendment and relock. If a lock hash mismatch exists, prompt compilation and exports stop with `BLOCKED`.
