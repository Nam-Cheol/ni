# Session resume

NI planning must be resumable from persisted project files, not hidden chat
memory. A later model session should be able to enter a planning workspace,
read the authoritative records, summarize where planning stopped, and ask the
remaining focused questions.

Resume is part of `ni-start` authoring. It does not add an execution harness,
task runner, database, vector store, or transcript archive.

## Resume inputs

At the start of a resumed planning session, `ni-start` reads:

- `AGENTS.md` for repository and model authority rules,
- `.ni/contract.json` for machine-readable planning state,
- `docs/plan/**` for human-readable planning state,
- `.ni/session.json` when it exists,
- `.ni/plan.lock.json` when it exists,
- `ni status --dir . --next-questions` when available.

The model must not depend on private memory from a prior chat. Any state needed
for resume must be present in those files or reconstructed from them.

## Source-of-truth order

Session state is below the contract and planning docs. After a lock exists, the
source-of-truth order remains:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If `.ni/session.json` conflicts with `.ni/contract.json`, the contract wins.
The model reports the conflict and updates or ignores the stale session entry in
the next planning edit. Session state never resolves a lock mismatch and never
marks a plan ready.

If `.ni/session.json` conflicts with locked docs, the lock and matching docs
win. If the lock hash check fails, resume stops with `BLOCKED`.

## Resume with session state

When `.ni/session.json` exists, `ni-start` may use it as a planning aid for:

- active planning focus,
- last planning summary,
- pending questions,
- recent decisions and risks,
- last readiness status and blockers,
- recently updated docs.

Before repeating any session claim, the model verifies it against the contract,
docs, lock state, and CLI readiness result. A useful resumed summary names
which session entries were confirmed and which were stale or superseded.

## Resume without session state

If `.ni/session.json` is missing, empty, invalid, or too stale to trust,
`ni-start` reconstructs continuity from `docs/plan/**` and
`.ni/contract.json`.

The reconstructed summary should include:

- purpose and delivery surface,
- accepted and draft capabilities,
- decisions, non-goals, assumptions, and risks,
- open blocker questions,
- missing trace links or evaluation gaps,
- the latest readiness state from the CLI when available.

After the next meaningful planning edit, `ni-start` should recreate or refresh
`.ni/session.json` with bounded continuity state. It must not backfill a raw
transcript by default.

## Conflict handling

Resume conflict handling is deterministic:

1. Read the contract and plan docs before trusting session state.
2. Compare session focus, pending questions, decisions, risks, and blockers
   against matching IDs in the contract and docs.
3. If the same ID disagrees, report the conflict and use the contract value.
4. If a session entry references an ID missing from the contract, treat it as
   stale or draft context until the docs and contract prove otherwise.
5. If a lock hash mismatch exists, stop and report `BLOCKED`.

Examples:

- Session says `OQ-003` is pending, but the contract marks `OQ-003` resolved:
  report that the session file is stale and do not ask the resolved question.
- Session says `DEC-002` is accepted, but the contract marks it draft: report
  the conflict and treat `DEC-002` as draft.
- Session says readiness was `READY`, but `ni status` reports `BLOCKED`:
  report the CLI status and keep planning open.

## Example transcript

See
[`examples/conversation-authoring/session-resume.md`](../examples/conversation-authoring/session-resume.md)
for an illustrative resumed planning flow:

- first planning session,
- docs updated,
- session ends,
- later session resumes,
- model summarizes previous state from persisted files,
- model asks pending questions.

