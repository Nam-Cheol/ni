# ni-grill Planning Challenge

This docs-only example shows how `ni-grill` challenges a draft NI plan before
`ni-end`.

Expected `ni status`: not claimed. This example is a transcript fixture, not a
trusted CLI workspace.

## What this proves

- `ni-grill` reads `docs/plan/**`, `.ni/contract.json`, and `ni status --proof --next-questions`.
- If deterministic blockers exist, it uses those before adding critique.
- It asks focused `GRILL-*` questions about vague decisions, acceptance
  evidence, risk, non-goals, handoff, and docs/contract sync.
- It does not execute downstream work, implement the product, or approve lock
  by model judgment.
- User answers are recorded through the same planning update discipline as
  `ni-start`: docs, contract, and session together, followed by status proof.

## Files

- `01-draft-plan.md`: draft plan excerpt with weak accepted content.
- `02-grill-questions.md`: model output after reading status and applying grill pressure.
- `03-user-answers.md`: user answers to focused grill questions.
- `04-docs-contract-delta.md`: planning updates and proof shape after answers.
- `05-status-after-grill.md`: status proof summary after the grill update.

## Boundary

`ni-grill` challenges planning quality before lock. It does not execute work.
Skills are UX; CLI is authority.

Covered by documentation review and `bash scripts/quality.sh`.
