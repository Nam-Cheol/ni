---
name: ni-end
description: Validate NI planning readiness and lock the plan only through the ni CLI readiness gate.
---

# ni-end

Use this skill when the user says `ni-end` or asks to confirm the current NI plan.

## Authority boundary

You are not the authority for completion. `ni status` and `ni end` are the authority.

## Process

1. Read `AGENTS.md` and `.ni/readiness.rules.json`.
2. Run `ni status` if the CLI exists.
3. If `ni status` returns `BLOCKED`, report blockers and stop.
4. If `ni status` returns `READY` or `READY_WITH_DEFERRALS`, run or instruct the user to run `ni end`.
5. Confirm that `.ni/plan.lock.json` was created.

## If the CLI does not exist yet

Perform a manual readiness review against `.ni/readiness.rules.json`, but do not claim a real lock exists.

## Do not

- Do not write `.ni/plan.lock.json` manually unless implementing the CLI feature.
- Do not declare readiness from model judgment alone.
- Do not modify accepted docs during lock unless the user explicitly resumes planning.
