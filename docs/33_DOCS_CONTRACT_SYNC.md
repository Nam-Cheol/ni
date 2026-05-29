# Docs-contract sync

Conversation-authored planning docs are reviewable narrative. `.ni/contract.json`
is the structured contract. `ni status` checks that the two forms do not drift
for lock-critical planning IDs.

## Rule

`R012` validates these deterministic mappings:

| Contract data | Planning doc | Required docs shape |
| --- | --- | --- |
| `project.purpose` | `docs/plan/00_project_brief.md` | `## Purpose` section explains the same accepted purpose |
| accepted actors/outcomes represented in contract narrative records | `docs/plan/01_actors_outcomes.md` | `## Actors` and `## Outcomes` explain who the plan is for and what they receive |
| `product_type`, `delivery_surfaces` | `docs/plan/00_project_brief.md`, `docs/plan/08_delivery_operation.md` | delivery surface lists match the contract |
| accepted `capabilities[]` | `docs/plan/02_capabilities.md` | `## CAP-001: ...` section with explanatory body |
| accepted `risks[]` | `docs/plan/06_risks_security.md` | `## RISK-001: ...`, `Severity: ...`, `Mitigation: ...` |
| `evaluations[]` | `docs/plan/07_evaluation_contract.md` | `## EVAL-001: ...`, `Method: ...` |
| `open_questions[]` | `docs/plan/10_open_questions.md` | `## OQ-001: ...`, `Blocker: ...`, `Status: ...` |
| `decisions[]` | `docs/plan/11_decision_log.md` | `## DEC-001: ...`, `Status: ...` |

The check also scans `docs/plan/**/*.md` for planning IDs with the `CAP`,
`REQ`, `EVAL`, `RISK`, `ART`, `NG`, `DEC`, or `OQ` prefixes and rejects
docs-only IDs that are missing from the contract. Decision status and risk
severity must match the contract value. Decision headings are checked for simple
deterministic polarity contradictions such as `Use X` versus `Do not use X`.
Resolved open questions must not still be shown as blockers. Evaluation docs
may include explicit capability fields, but any referenced capability ID must
exist in the contract.

## First-run sync diagnostics

After the first `ni-start` card, the user's answers must be recorded in both
the human docs and `.ni/contract.json`. `ni status --proof` reports these
first-run drift diagnostics as `R012` blocker issues with stable diagnostic
IDs:

| Diagnostic ID | Drift detected | Blocks ni-end |
| --- | --- | --- |
| `SYNC-014` | project purpose is documented but missing from the contract, recorded in the contract but missing from docs, or differs between the two | yes |
| `SYNC-015` | actors/outcomes are documented but missing from contract narrative records, recorded in contract but missing from docs, or differ between the two | yes |
| `SYNC-016` | delivery surface is documented but missing from the contract, recorded in the contract but missing from docs, or differs between the two | yes |

Example proof item:

```text
- Project purpose differs between docs and contract.
  ID: SYNC-014
  Location: docs/plan/00_project_brief.md and .ni/contract.json:project.purpose
  Problem: Project purpose differs between docs and contract.
  Why it matters: ni cannot safely lock a plan when the human-readable purpose and machine-readable contract disagree.
  Suggested repair: Update the stale side so both describe the same project purpose.
  Blocks ni-end: true
```

Delivery-surface drift uses the same shape:

```text
- Delivery surface differs between docs and contract.
  ID: SYNC-016
  Location: docs/plan/08_delivery_operation.md and .ni/contract.json:delivery_surfaces
  Problem: Delivery surface differs between docs and contract.
  Why it matters: ni-run cannot produce a safe handoff if human docs and the contract disagree about how the project will be delivered.
  Suggested repair: Update the stale side so both list the same delivery surface, or mark the surface deferred with an explicit reason.
  Blocks ni-end: true
```

These checks do not infer missing intent. If the user's answer is uncertain,
the repair is to keep the uncertainty visible as an assumption, blocker open
question, or deferral rather than silently accepting a stale value.

## Diagnostic shape

`ni status --proof` prints each sync failure with stable repair fields:

- `ID`
- `Location`
- `Problem`
- `Why it matters`
- `Suggested repair`
- `Blocks ni-end`

`ni status --json --proof` includes the same values under
`sync_diagnostic` on the `issues[]` and `proof[]` entries for `R012`.

## Repair workflow

Sync failures are blocker readiness issues. They do not rewrite files.

Use the `R012` message as the repair instruction:

- If docs mention an unknown ID, add the matching contract item or remove/rename
  the docs section.
- If the contract contains an accepted capability or risk missing from docs, add
  the matching section in the mapped file.
- If an accepted capability has only a heading, add a short explanatory body.
- If decision status or risk severity conflicts, update the stale source.
- If a decision heading contradicts the contract title, revise the stale source
  so the accepted decision has one polarity.
- If risk mitigation or evaluation method is missing from docs, add the
  deterministic field line.
- If a resolved open question is still marked as a blocker in docs or contract,
  change the stale side to `Blocker: false` and `Status: resolved`, or reopen
  it consistently when it still blocks lock readiness.
- If an evaluation references a missing capability, add the capability to the
  contract or point the evaluation docs at an existing capability.
- If `SYNC-014`, `SYNC-015`, or `SYNC-016` appears after first-run authoring,
  update the stale side so purpose, actors/outcomes, and delivery surface match
  before attempting `ni end`.

`ni end` calls the same readiness gate as `ni status`, so sync failures prevent
locking. `ni run` remains a prompt compiler only and does not perform sync
repairs.
