---
name: ni-end
description: Validate NI planning readiness and lock the plan only through the ni CLI readiness gate.
---

# ni-end

Use this skill when the user says `ni-end` or asks to confirm the current NI plan.

## Authority boundary

You are not the authority for completion. `ni status` and `ni end` are the authority.

Do not create, edit, or repair `.ni/plan.lock.json` by hand. The CLI is the only lock writer.

If model judgment and CLI output disagree, the CLI output wins. Report `BLOCKED`, `READY`, or `READY_WITH_DEFERRALS` exactly as the CLI reports it and do not substitute a model-derived readiness state.

## Process

1. Read `AGENTS.md`, `.ni/readiness.rules.json`, and `.ni/contract.json`.
2. Run `ni status --dir .`.
3. If `ni status` returns `BLOCKED`, report blockers and stop.
4. If `ni status` returns `READY` or `READY_WITH_DEFERRALS`, run `ni end --dir .`.
5. Confirm that `.ni/plan.lock.json` was created by the CLI.
6. Report the readiness status and lockfile path.

## Do not

- Do not write `.ni/plan.lock.json` manually unless implementing the CLI feature.
- Do not declare readiness from model judgment alone.
- Do not modify accepted docs during lock unless the user explicitly resumes planning.
- Do not add execution adapters or remote automation.
