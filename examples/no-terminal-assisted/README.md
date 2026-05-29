# No-Terminal Assisted Draft

## 1. Purpose

This example shows a no-terminal assisted drafting flow for `ni`. A user starts
in a model workspace before installing the CLI, asks the model to draft
`docs/plan/**` and `.ni/contract.json`, and keeps the result marked not locked
until later CLI validation.

It is intentionally docs-only. There is no lockfile, no hash proof, and no
compiled handoff prompt.

## 2. Short flow

1. User starts in a model workspace and loads or copies `ni-start` guidance.
2. Model asks the first-run card:
   - What should this project change, for whom, and why does it matter?
   - Who are the primary actors, and what outcome should each one get?
   - What is the likely delivery surface?
3. User answers with an initial project idea and any clear exclusions.
4. Model drafts `docs/plan/00_project_brief.md` and `.ni/contract.json`.
5. The draft is marked "not locked" and keeps assumptions, non-goals, and
   blocker questions visible.
6. Later, CLI validation is required before readiness, locking, hash trust,
   prompt compilation, or downstream seed generation.

## 3. What this proves

- Assisted planning can capture docs and a draft contract before local CLI
  access exists.
- No-terminal users can start planning, but cannot finish a trusted lock
  without CLI validation.
- Model judgment is not readiness, lock, hash, or prompt compilation
  authority.
- Model workspace packs can guide drafting, but they do not execute downstream
  work and do not override CLI readiness.
- This example does not make deterministic validation claims.

## 4. Product type / surface

- `product_type`: draft `workflow`
- `delivery_surface`: draft `document`
- Expected `ni status`: not claimed until a teammate, CI job, or local CLI
  setup runs the command.
- Expected `ni run` target: not applicable.

## 5. Files

- `docs/plan/00_project_brief.md`: human-readable planning notes.
- `.ni/contract.json`: a draft contract aligned with the docs.
- Assumptions and blocker questions that stay visible instead of being treated
  as accepted decisions.

## 6. No-terminal checklist

- Copy or load `ni-start` guidance.
- Describe the project idea.
- Ask the model to create a `docs/plan/**` draft.
- Ask the model to draft `.ni/contract.json`.
- Mark uncertain statements as assumptions or open questions.
- Mark explicit exclusions as non-goals.
- Do not treat the draft as locked.
- Later validate with the CLI or a teammate.

## 7. Team handoff path

1. A user without terminal access drafts with the model.
2. A teammate with the CLI runs `ni status --proof --next-questions`.
3. The teammate returns blockers, proof, and questions.
4. The user continues planning with the model.
5. Lock happens only after deterministic CLI validation clears blockers and the
   user confirms the accepted plan.

## 8. Commands

This example is intentionally docs-only. From the repository root:

```bash
test -f examples/no-terminal-assisted/README.md
test -f examples/no-terminal-assisted/README.ko.md
test -f examples/no-terminal-assisted/docs/plan/00_project_brief.md
test -f examples/no-terminal-assisted/.ni/contract.json
```

## 9. Expected output

The `test` commands should exit successfully. Do not describe the draft as
`READY`, `READY_WITH_DEFERRALS`, or `BLOCKED` unless a trusted CLI run is
available and quoted separately.

## 10. demo-check coverage

Covered by `bash scripts/demo-check.sh` as a docs-only example.

The demo check verifies required files and boundary wording. It does not run
`ni status`, `ni end`, or `ni run` for this intentionally assisted draft.

## 11. Korean companion

Korean companion docs exist: `README.ko.md`.

## 12. Graduate before handoff

Use full `ni` before this draft guides implementation or downstream seed
generation:

1. Install `ni` through the release binary path or curl installer.
2. Run `ni status --proof --next-questions`.
3. Fix blockers in docs and the contract draft.
4. Run `ni end` only after readiness and user confirmation.
5. Run `ni run` only after a valid lock exists.

## 13. Non-execution boundary

This example does not add a web service, model API call, runtime execution,
shell adapter, queue, or model-authoritative skill.
