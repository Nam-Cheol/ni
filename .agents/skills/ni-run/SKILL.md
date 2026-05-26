---
name: ni-run
description: Compile a 4000-character-or-less execution goal prompt from a locked NI project plan.
---

# ni-run

Use this skill when the user says `ni-run` or asks to produce the NI execution prompt.

## Authority boundary

`ni run` is a prompt compiler in v0. It is not an execution command.

## Process

1. Verify `.ni/plan.lock.json` exists.
2. Run `ni run` if the CLI exists.
3. Confirm the prompt is 4000 characters or less.
4. The prompt must reference authoritative files instead of embedding all docs.
5. The prompt must tell the execution agent to derive a project-specific harness before implementation.

## Do not

- Do not implement the product directly from this skill.
- Do not execute Codex or shell commands as part of v0 `ni run`.
- Do not ignore a stale lock hash.
- Do not change locked planning docs.
