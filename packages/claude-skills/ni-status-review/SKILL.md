---
name: ni-status-review
description: Explain NI status proof output in Claude-compatible workspaces without becoming a second readiness engine.
---

# ni-status-review

Use this skill when the user asks Claude to review, explain, or act on
`ni status` output.

`ni-status-review` is a review and planning-navigation skill. It is not a
readiness gate.

## Authority Boundary

Skills are UX; CLI is authority.

`ni status` is the authority for `BLOCKED`, `READY_WITH_DEFERRALS`, and
`READY`.

- Skills may help draft or explain proof-related planning text.
- Skills may help draft amended planning text.
- Skills may help explain `LOCK-STALE`.
- Skills do not determine readiness.
- Skills do not lock plans.
- Skills do not lock or relock.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- Skills do not update `.ni/plan.lock.json`.

`LOCK-STALE` means the existing lock no longer matches current planning inputs.
Recovery order: `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000`.

Run or request one of these commands:

```bash
ni status --dir . --proof --next-questions
ni status --dir . --proof --json --next-questions
```

If command execution is unavailable, ask the user to paste the exact CLI
output. Do not infer a status from docs alone.

If a lock hash mismatch exists, stop and report `BLOCKED`.

## Review Process

1. Preserve the status exactly as reported by the CLI.
2. Group blockers by affected planning record IDs.
3. Explain what each blocker means in plain language.
4. Identify whether the next action belongs in conversation authoring
   (`ni-start`), lock confirmation (`ni-end`), or prompt compilation
   (`ni-run`).
5. Prefer the CLI-provided `next_questions` when asking the user what to
   resolve next.
6. Do not rewrite blockers into weaker requirements.

## Output Shape

Report:

- CLI status exactly as printed.
- Blocking issues or deferrals.
- The highest-impact one to three next questions.
- The file areas likely to change, such as `docs/plan/**` and
  `.ni/contract.json`.

If the status is `READY` or `READY_WITH_DEFERRALS`, say that the next step is
`ni-end` confirmation. Do not claim the plan is locked until `ni end` succeeds.

## Boundaries

Do not execute implementation.

Do not call Claude APIs.

Do not edit `.ni/plan.lock.json` manually.

Do not create downstream execution state, adapters, queues, model orchestration,
or automation.
