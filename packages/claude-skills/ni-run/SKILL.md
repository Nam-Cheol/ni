---
name: ni-run
description: Compile a bounded NI handoff prompt in Claude-compatible workspaces without executing downstream work.
---

# ni-run

Use this skill when the user asks Claude to produce a handoff prompt from a
locked NI plan.

`namba-intent run` is a prompt compiler in v0. It is not an execution command.

## Authority Boundary

Skills are UX; CLI is authority.

Do not reimplement prompt compilation in the skill.

The `namba-intent` CLI verifies lock hashes and enforces the prompt budget. If `namba-intent run`
reports a missing lock, stale lock, or hash mismatch, report `BLOCKED` and
stop.

Do not edit `.ni/plan.lock.json` manually. Do not change locked planning docs
to make prompt compilation pass.

`LOCK-STALE` means the existing lock no longer matches current planning inputs.
- Skills may help draft amended planning text.
- Skills may help explain `LOCK-STALE`.
- Skills do not determine readiness.
- Skills do not lock or relock.
- Skills do not replace `namba-intent status`, `namba-intent end`, or `namba-intent run`.
- Skills do not update `.ni/plan.lock.json`.

Recovery order: `review changed intent -> namba-intent status --proof --next-questions -> namba-intent end -> namba-intent run --max-chars 4000`.

## Process

1. Locate the project root from the user's workspace context.
2. Confirm `.ni/plan.lock.json` exists.
3. Infer the target only from explicit user intent or locked project context.
   Ask one short target question if multiple targets are plausible.
4. Run or request:

```bash
namba-intent run --dir . --target <target> --max-chars 4000
```

5. Use `--out .ni/generated/<target>.prompt.txt` only when the user asks for a
   file or the prompt is too long to show comfortably.
6. Confirm the compiled prompt is 4000 characters or less.
7. Show the prompt or output path.
8. State clearly that `namba-intent` compiled a prompt only and did not execute
   implementation.

## Target Selection

Use the target the user names, such as `claude-code`, `codex`, `human-team`, or
`generic`.

If no target is named, inspect locked planning context such as
`docs/plan/09_execution_strategy.md`, `docs/plan/11_decision_log.md`, and
`.ni/contract.json`. Do not choose a target merely because the current model is
Claude.

## Boundaries

Do not execute implementation.

Do not execute Codex or shell commands as part of v0 `namba-intent run`.

Do not call Claude APIs.

Do not run downstream commands, shell/Codex adapters, model orchestration,
queues, evidence runners, PR automation, or target runtimes.

Do not synthesize a replacement prompt from memory when the CLI refuses to
compile one.
