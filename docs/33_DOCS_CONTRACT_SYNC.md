# Docs-contract sync

Conversation-authored planning docs are reviewable narrative. `.ni/contract.json`
is the structured contract. `ni status` checks that the two forms do not drift
for lock-critical planning IDs.

## Rule

`R012` validates these deterministic mappings:

| Contract data | Planning doc | Required docs shape |
| --- | --- | --- |
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
- If a resolved open question is still marked as a blocker, change `Blocker:
  false` or reopen it consistently in docs and contract.
- If an evaluation references a missing capability, add the capability to the
  contract or point the evaluation docs at an existing capability.

`ni end` calls the same readiness gate as `ni status`, so sync failures prevent
locking. `ni run` remains a prompt compiler only and does not perform sync
repairs.
