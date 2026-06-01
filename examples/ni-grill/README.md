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
- It labels findings with `Critical`, `High`, `Medium`, `Low`, or `Note` and
  keeps output to at most 5 findings by default.
- It can review benchmark cases after `ni status` says `READY` to pressure-test
  claim boundaries, `not_measured` sections, and handoff wording.
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
- `06-internal-dashboard-grill.md`: benchmark grill against the
  `internal-dashboard` case after isolated `READY` proof.
- `07-research-protocol-grill.md`: benchmark grill against the
  `research-protocol` case after isolated `READY` proof.
- `08-grill-lessons.md`: lessons from dogfooding `ni-grill` on benchmark
  evidence.
- `09-ni-project-grill.md`: read-only dogfood report from applying `ni-grill`
  to ni's own current planning state.

## Boundary

`ni-grill` challenges planning quality before lock. It does not execute work.
Skills are UX; CLI is authority. Benchmark grills are claim-boundary reviews;
they do not create new empirical claims, run generated prompts, call model
APIs, implement products, perform fieldwork, or make `ni-grill` authoritative
over the CLI.

Default grill output is budgeted: at most 5 findings, with at most 3
`Critical`/`High` findings shown first when they exist. Lower-priority findings
should be summarized instead of listed exhaustively.

Covered by documentation review and `bash scripts/quality.sh`.
