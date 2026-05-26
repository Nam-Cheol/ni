# Codex Exec Experiment

This is a local experiment for running the prompt compiled by `ni run` with
`codex exec`. It is not part of `ni-kernel`, and `ni run` must not invoke Codex
or any shell execution adapter.

## Purpose

Use this experiment only after the v0 kernel flow works:

```bash
go run ./cmd/ni status --dir .
go run ./cmd/ni end --dir .
go run ./cmd/ni run --dir . --out .ni/generated/goal.prompt.txt
```

The kernel remains responsible for the planning contract, readiness gate,
lockfile, source-of-truth rule, and prompt compiler. Codex remains an external
consumer of the compiled prompt.

## Safety Rules

- Run only from a clean git tree.
- Compile the prompt through `ni run` immediately before invoking Codex.
- Let `ni run` verify the lock hash; if it reports `BLOCKED`, stop.
- Use `codex exec --sandbox workspace-write`.
- Do not use `--dangerously-bypass-approvals-and-sandbox`.
- Do not use `--dangerously-bypass-hook-trust`.
- Do not add this path to `ni run`.

## Manual Experiment

From the repository root:

```bash
test -z "$(git status --porcelain=v1)"
go run ./cmd/ni run --dir . --out .ni/generated/goal.prompt.txt
codex exec --sandbox workspace-write --cd . - < .ni/generated/goal.prompt.txt
```

The first command requires a clean git tree. The second command validates the
lock hash and writes the prompt. The third command starts Codex with the
compiled prompt under the workspace-write sandbox.

## Scripted Experiment

The optional helper script performs the same checks:

```bash
bash scripts/run-codex-prompt.sh
```

The script refuses to run when the git tree is dirty, when `ni run` cannot
validate the lock, or when `codex` is not installed.

## Non-Goals

- This experiment does not make Codex exec part of `ni-kernel`.
- This experiment does not add a Codex adapter.
- This experiment does not make `ni run` execute anything.
- This experiment does not create a queue, evidence runner, or release
  automation.
