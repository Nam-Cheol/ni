---
name: ni-run
description: Compile a 4000-character-or-less conversational handoff prompt from a locked NI project plan.
---

# ni-run

Use this skill when the user says `ni-run` or asks to produce the NI handoff
prompt for a downstream target.

## Authority boundary

`ni run` is a prompt compiler in v0. It is not an execution command.

Do not reimplement prompt compilation in the skill. The CLI verifies lock hashes and enforces the prompt budget.

If `ni run` reports a missing or stale lock, the skill must not produce a replacement prompt from memory. Report the CLI result as `BLOCKED` and stop.

## Process

1. Read `AGENTS.md` and confirm `.ni/plan.lock.json` exists.
2. Infer the target only from explicit user intent or locked project context.
3. Ask the user for a target if it is not inferable.
4. Run or request `ni run --dir . --target <target> --max-chars 4000`.
5. Use `--out .ni/generated/<target>.prompt.txt` only when the user asks for a
   file or the prompt is too long to show comfortably.
6. If `ni run` reports a missing or stale lock, report `BLOCKED` and stop.
7. Confirm the prompt is 4000 characters or less.
8. Confirm the prompt references authoritative files instead of embedding all docs.
9. Show the generated prompt, or show the output path if `--out` was used.
10. State clearly that `ni` compiled a prompt only and did not execute
    implementation.

## Target Selection

Use a target immediately when the user names one, such as `codex`,
`human-team`, `hyper-run`, `namba-ai`, `ouroboros`, `spec-kit`, or `generic`.

If the user does not name a target, inspect locked project context such as
`docs/plan/09_execution_strategy.md`, `docs/plan/11_decision_log.md`,
`.ni/contract.json`, and recent conversation. Use the target only when it is
clearly identified as the current handoff target.

Default to `codex` only when project context says Codex is the current
experiment or selected downstream target. Do not choose `codex` merely because
the model is Codex.

If multiple targets are plausible, ask one short question and do not run
`ni run` until the user chooses.

## Output Shape

When the prompt is printed to stdout, include the command and then the generated
prompt. When `--out` is used, include the command and path. In both cases, end
with a boundary statement like:

```text
`ni` compiled this prompt only. It did not execute implementation, Codex, shell
commands, adapters, queues, PR automation, or downstream runtime work.
```

## Do not

- Do not implement the product directly from this skill.
- Do not execute Codex or shell commands as part of v0 `ni run`.
- Do not ignore a stale lock hash.
- Do not change locked planning docs.
- Do not create execution adapters or remote automation.
