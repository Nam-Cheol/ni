---
name: ni-end
description: Review NI readiness in Claude-compatible workspaces and lock only through the NI CLI.
---

# ni-end

Use this skill when the user asks Claude to review readiness, finish planning,
or lock an NI plan.

`ni-end` is a confirmation workflow. It does not decide readiness and does not
write the lockfile by hand.

## Authority Boundary

`ni status` and `ni end` are the authority.

Do not write `.ni/plan.lock.json` manually. Do not repair a stale lock by
editing JSON. If a lock hash mismatch exists, stop and report `BLOCKED`.

If command execution is unavailable, ask the user to run the command and paste
the exact output. Do not substitute model judgment for CLI output.

## Process

1. Locate the project root from the user's workspace context.
2. Read `AGENTS.md`, `.ni/readiness.rules.json`, `.ni/contract.json`, and the
   relevant `docs/plan/**` files.
3. Run or request:

```bash
ni status --dir .
```

4. If the CLI reports `BLOCKED`, report `BLOCKED`, list the CLI blockers, and
   stop. Do not run `ni end`.
5. If the CLI reports `READY` or `READY_WITH_DEFERRALS`, summarize the contract
   that will be locked.
6. Include the readiness status exactly as reported by the CLI, project
   purpose, accepted capabilities, linked requirements, evaluations, risks,
   mitigations, artifacts, deferred decisions, non-blocking open questions, and
   source files that will be hashed.
7. Ask for explicit confirmation:

```text
Confirm that I should run ni end --dir . and let the CLI write .ni/plan.lock.json?
```

8. Only after explicit confirmation, run or request:

```bash
ni end --dir .
```

9. Confirm that the CLI created `.ni/plan.lock.json` and report the lock path.

## Boundaries

Do not execute implementation.

Do not call Claude APIs.

Do not auto-lock from a vague approval such as "looks good." The user must
confirm the exact lock command after seeing the summary.

Do not modify accepted planning docs during the lock flow unless the user
explicitly returns to planning.

Do not add downstream execution behavior, shell/Codex adapters, model
orchestration, queues, evidence runners, or automation.

