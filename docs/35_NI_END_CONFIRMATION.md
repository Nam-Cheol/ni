# ni-end confirmation

`ni-end` is the conversational confirmation step before `ni end` writes
`.ni/plan.lock.json`.

The skill improves the user experience around locking, but it does not become a
readiness authority. The CLI remains authoritative for both status and lock
creation:

```text
ni status --dir . -> conversational summary -> explicit confirmation -> ni end --dir .
```

## Required Flow

1. Read the current planning sources: `AGENTS.md`, `.ni/readiness.rules.json`,
   `.ni/contract.json`, and relevant `docs/plan/**` files.
2. Run or request `ni status --dir .`.
3. If the status is `BLOCKED`, report the blockers and stop.
4. If the status is `READY` or `READY_WITH_DEFERRALS`, summarize the contract
   that would be locked.
5. List deferred decisions and open non-blocking questions.
6. Ask the user for explicit confirmation to run `ni end --dir .`.
7. Only after confirmation, run or instruct `ni end --dir .`.
8. Verify that `.ni/plan.lock.json` exists and report the CLI-created lock path.

The model must never create, edit, repair, or synthesize `.ni/plan.lock.json`.

## Optional Pre-Lock Grill

Before starting this confirmation flow, a user may run `ni-grill` to challenge
the draft plan. `ni-grill` is optional UX: it can ask focused questions about
weak assumptions, vague decisions, missing acceptance evidence, risks,
non-goals, handoff ambiguity, and docs/contract sync.

`ni-grill` does not replace this flow. If it changes planning state, the plan
must pass `ni status --dir .` again before `ni-end` summarizes it for lock
confirmation.

## Summary Contents

The pre-lock summary should be compact but complete enough for a user to know
what they are approving:

- readiness status exactly as reported by `ni status`,
- project name and purpose,
- readiness profile, product type, delivery surfaces, and interaction mode,
- accepted capabilities and their linked requirements, evaluations, risks, and
  artifacts,
- high-severity risks and mitigations,
- deferred decisions,
- open non-blocking questions,
- source files that will be hashed into the lock.

`READY_WITH_DEFERRALS` is lockable only because the CLI says the remaining
items are deferrals. The summary must name those deferrals before asking for
confirmation.

## Confirmation Contract

A user command like "run ni-end" or "lock the plan" starts the confirmation
flow. It is not enough to create a lock. The user must answer the specific
confirmation prompt after seeing the summary.

Accepted confirmation examples:

```text
Yes, run ni end.
Confirmed.
Proceed with ni end --dir .
```

Non-confirmation examples:

```text
Looks good.
What about the risk section?
Change CAP-002 first.
```

If the user changes planning intent, stop the lock flow and return to
conversation authoring. The updated plan must pass `ni status --dir .` before
`ni-end` can ask for lock confirmation again.

## Blocked Behavior

When `ni status` reports `BLOCKED`, `ni-end` must not summarize the plan as
ready and must not run `ni end`. It should report the CLI blockers directly and
hand the user back to planning.

If a lock hash mismatch is present, report `BLOCKED` and stop. Do not read the
stale lock as approval to proceed, and do not manually repair the lockfile.

## Boundary

`ni-end` does not execute implementation, generate harness state, weaken
acceptance criteria, resolve planning questions by model judgment, or bypass
the readiness gate. It only helps the user review the CLI-ready contract before
the CLI writes the lock.
