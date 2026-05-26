---
name: ni-run
description: Compile a 4000-character-or-less execution goal prompt from a locked NI project plan.
---

# ni-run

Use this skill when the user says `ni-run` or asks to produce the NI execution prompt.

## Authority boundary

`ni run` is a prompt compiler in v0. It is not an execution command.

Do not reimplement prompt compilation in the skill. The CLI verifies lock hashes and enforces the prompt budget.

If `ni run` reports a missing or stale lock, the skill must not produce a replacement prompt from memory. Report the CLI result as `BLOCKED` and stop.

## Process

1. Read `AGENTS.md` and confirm `.ni/plan.lock.json` exists.
2. Run `ni run --dir . --max-chars 4000`, or use `--out .ni/generated/goal.prompt.txt` when the user asks for a file.
3. If `ni run` reports a missing or stale lock, report `BLOCKED` and stop.
4. Confirm the prompt is 4000 characters or less.
5. Confirm the prompt references authoritative files instead of embedding all docs.
6. Report the command and output path, if any.

## Do not

- Do not implement the product directly from this skill.
- Do not execute Codex or shell commands as part of v0 `ni run`.
- Do not ignore a stale lock hash.
- Do not change locked planning docs.
- Do not create execution adapters or remote automation.
