# Docs-contract sync

Conversation-authored planning docs are reviewable narrative. `.ni/contract.json`
is the structured contract. `ni status` checks that the two forms do not drift
for lock-critical planning IDs.

## Rule

`R012` validates these deterministic mappings:

| Contract data | Planning doc | Required docs shape |
| --- | --- | --- |
| accepted `capabilities[]` | `docs/plan/02_capabilities.md` | `## CAP-001: ...` section |
| accepted `risks[]` | `docs/plan/06_risks_security.md` | `## RISK-001: ...`, `Severity: ...`, `Mitigation: ...` |
| `evaluations[]` | `docs/plan/07_evaluation_contract.md` | `## EVAL-001: ...`, `Method: ...` |
| `decisions[]` | `docs/plan/11_decision_log.md` | `## DEC-001: ...`, `Status: ...` |

The check also rejects matching doc sections whose IDs do not exist in the
contract. Decision status and risk severity must match the contract value.

## Repair workflow

Sync failures are blocker readiness issues. They do not rewrite files.

Use the `R012` message as the repair instruction:

- If docs mention an unknown ID, add the matching contract item or remove/rename
  the docs section.
- If the contract contains an accepted capability or risk missing from docs, add
  the matching section in the mapped file.
- If decision status or risk severity conflicts, update the stale source.
- If risk mitigation or evaluation method is missing from docs, add the
  deterministic field line.

`ni end` calls the same readiness gate as `ni status`, so sync failures prevent
locking. `ni run` remains a prompt compiler only and does not perform sync
repairs.
