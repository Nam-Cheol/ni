# Session state

`.ni/session.json` is persistent planning memory for model-assisted NI
authoring. It lets a later model session resume from explicit state instead of
hidden chat history.

Session state is a planning aid, not authority. It does not replace
`.ni/contract.json`, `docs/plan/**`, `ni status`, or `.ni/plan.lock.json`.

## Authority order

After `.ni/plan.lock.json` exists, use this order:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

`.ni/session.json` must not override locked docs, repair a stale lock, or mark
docs complete. If it conflicts with the contract or planning docs, the contract
and docs win, and the session state should be corrected.

Raw transcript is not the source of truth. The session file stores only a
bounded summary and selected planning records by default.

## Schema

The published schema is `schema/ni.session.v0.json`. `ni init` creates an empty
session file with these fields:

- `active_planning_focus`
- `last_planning_summary`
- `pending_questions`
- `recent_decisions`
- `recent_risks`
- `last_readiness_status`
- `last_readiness_blockers`
- `last_updated_docs`
- `authority_order`
- `notes`

`pending_questions`, `recent_decisions`, and `recent_risks` are short carryover
records for planning continuity. They should point back to stable contract IDs
when those IDs exist.

## ni-start maintenance

At the start of a planning turn, `ni-start` should read `.ni/session.json`
after the authoritative contract and docs. It may use the file to summarize
where the previous turn left off, but it must verify important claims against
`.ni/contract.json`, `docs/plan/**`, and `ni status`.

After a meaningful planning update, `ni-start` should update `.ni/session.json`
with the latest focus, summary, pending questions, recent decisions, recent
risks, readiness result, blockers, and docs changed in that turn.
