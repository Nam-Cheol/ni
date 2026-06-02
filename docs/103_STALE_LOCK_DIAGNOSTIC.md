# Stale Lock Diagnostic

## Current status

ni now surfaces stale existing locks on CLI diagnostic surfaces. The diagnostic
label is `LOCK-STALE`.

`ni status --proof --next-questions` may still report `READY` for the current
planning contract, but it now warns when an existing `.ni/plan.lock.json` no
longer matches the current lockable planning inputs.

## What stale means

A stale lock means:

- `.ni/plan.lock.json` exists; and
- the current lockable planning inputs produce hashes that differ from the
  hashes recorded in `.ni/plan.lock.json`.

The comparison uses the same lock verification path used by `ni run`.
Lockable inputs are `.ni/contract.json` and the required `docs/plan/**` files
recorded by the lock. `.ni/session.json` remains mutable continuity state below
locked docs and is not hashed by the current lock semantics.

## What stale does not mean

Stale does not prove:

- implementation failure;
- product unreadiness;
- benchmark failure;
- downstream execution failure;
- adoption, cost, latency, quality, or downstream agent performance.

It means only that the current planning inputs and the existing locked plan no
longer match, so downstream handoff must not rely on the stale lock.

## User-facing behavior

`ni status --proof --next-questions` keeps readiness semantics separate from
lock freshness. If the current planning contract is ready, readiness can remain
`READY`, but stale existing locks appear under `Warnings` with `LOCK-STALE`.

The warning text is:

```text
WARNING: LOCK-STALE existing lock is stale. Current planning inputs differ from .ni/plan.lock.json.
```

When available, the diagnostic also names the first mismatched lockable input
path and gives the recovery action.

`ni end` preserves the normal first-lock flow when no lock exists. If an
existing lock is current, it says the lock is current and that `ni end` will
refresh through the CLI readiness flow. If an existing lock is stale, it says
`ni end` is the CLI-authoritative relock step after changed intent has been
reviewed.

`ni run` continues to refuse stale handoff. It does not relock automatically,
does not mutate `.ni/plan.lock.json`, and does not execute downstream work. Its
stale-lock refusal includes recovery guidance.

## Recovery flow

```text
review changed intent
-> run ni status --proof --next-questions
-> run ni end
-> run ni run --max-chars 4000
```

The review step is human and planning-state work. Update `docs/plan/**` and
`.ni/contract.json` together when intent changes, then rely on the CLI gates.
For practical user workflows, see
[`104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md`](104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md).

## Authority boundary

- CLI is authority.
- Skills are UX.
- Skills do not determine readiness.
- Skills do not lock or relock.
- No-terminal assisted workflow does not provide deterministic validation.
- `ni run` compiles a bounded handoff prompt and does not execute downstream
  work.

## Test coverage

This task adds focused stale-lock coverage for:

- no lock exists: `ni status --proof --next-questions` does not emit
  `LOCK-STALE`;
- lock exists and is current: `ni status --proof --next-questions` does not
  emit `LOCK-STALE`;
- lock exists and planning inputs change: `ni status --proof --next-questions`
  emits `LOCK-STALE`;
- lock exists and planning inputs change: `ni run` refuses stale handoff and
  includes recovery guidance;
- after relock in a temporary fixture: stale warning disappears and `ni run`
  compiles the bounded handoff again.

## Remaining follow-up candidates

- Add broader changed-intent fixtures.
- Improve no-terminal stale-lock explanation.
- Add model workspace stale-lock wording verification.
