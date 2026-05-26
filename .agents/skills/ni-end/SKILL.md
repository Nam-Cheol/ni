---
name: ni-end
description: Confirm NI planning readiness conversationally, then lock the plan only through the ni CLI readiness gate.
---

# ni-end

Use this skill when the user says `ni-end`, asks to confirm the current NI
plan, or asks to lock a ready NI plan.

## Authority boundary

You are not the authority for completion. `ni status` and `ni end` are the authority.

Do not create, edit, or repair `.ni/plan.lock.json` by hand. The CLI is the only lock writer.

If model judgment and CLI output disagree, the CLI output wins. Report `BLOCKED`, `READY`, or `READY_WITH_DEFERRALS` exactly as the CLI reports it and do not substitute a model-derived readiness state.

## Process

1. Read `AGENTS.md`, `.ni/readiness.rules.json`, `.ni/contract.json`, and the
   relevant `docs/plan/**` files.
2. Run `ni status --dir .`. If you cannot run commands in the current
   environment, ask the user to run `ni status --dir .` and paste the result
   before proceeding.
3. If `ni status` returns `BLOCKED`, report `BLOCKED`, list the CLI blockers,
   and stop. Do not run `ni end`.
4. If `ni status` returns `READY` or `READY_WITH_DEFERRALS`, summarize what the
   CLI is about to lock before running `ni end --dir .`.
5. The pre-lock summary must include:
   - readiness status exactly as reported by `ni status`,
   - project name and purpose,
   - readiness profile, product type, delivery surfaces, and interaction mode
     when present,
   - accepted capabilities with linked requirements, evaluations, risks, and
     artifacts,
   - high-severity risks and mitigations,
   - deferred decisions,
   - open non-blocking questions,
   - source files that will be hashed into `.ni/plan.lock.json`.
6. Ask for explicit user confirmation after the summary. Use a direct question
   such as: `Confirm that I should run ni end --dir . and let the CLI write
   .ni/plan.lock.json?`
7. Only after the user explicitly confirms, run `ni end --dir .`. If you cannot
   run commands in the current environment, instruct the user to run
   `ni end --dir .` themselves.
8. Confirm that `.ni/plan.lock.json` was created by the CLI.
9. Report the readiness status and lockfile path.

## Confirmation Rules

Explicit confirmation must be an affirmative user response to the specific lock
question in the current conversation. A prior request such as "run ni-end" or
"lock the plan" starts the confirmation flow; it is not by itself permission to
create the lock.

If the user changes planning intent during confirmation, stop the lock flow and
return to planning. The changed intent must be captured through the planning
authoring process and checked again with `ni status --dir .`.

If `READY_WITH_DEFERRALS` is reported, do not hide the deferrals. List each
deferred decision and each open non-blocking question before asking for
confirmation. The user may still confirm the lock only because the CLI has
classified those items as non-blocking.

## Do not

- Do not write `.ni/plan.lock.json` manually.
- Do not declare readiness from model judgment alone.
- Do not bypass `ni status`.
- Do not modify accepted docs during lock unless the user explicitly resumes planning.
- Do not auto-lock from chat without the explicit confirmation step.
- Do not execute implementation or downstream harness work.
- Do not add execution adapters or remote automation.
