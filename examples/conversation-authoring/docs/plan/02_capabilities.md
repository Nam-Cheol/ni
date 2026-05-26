# Capabilities

## CAP-001: Draft refund recommendations

Status: accepted

The assistant drafts refund recommendations for support agents. It cites the
ticket facts used and the relevant internal refund policy page section. It does
not approve, issue, or initiate refunds.

Linked records: `REQ-001`, `REQ-002`, `EVAL-001`, `RISK-001`, `RISK-002`,
`ART-001`.

## CAP-002: Escalate ambiguous or conflicting cases

Status: accepted

The assistant escalates when policy is ambiguous, ticket facts conflict with
policy, or supervisor judgment is required. Escalation preserves the agent's
next action without inventing policy.

Linked records: `REQ-003`, `REQ-004`, `EVAL-002`, `RISK-001`, `RISK-003`,
`ART-002`.

## CAP-003: Maintain docs and contract from conversation

Status: accepted

The model updates human planning docs and `.ni/contract.json` together after
conversation changes planning state. The CLI remains responsible for readiness,
lock, and prompt compilation.

Linked records: `REQ-005`, `REQ-006`, `EVAL-003`, `RISK-004`, `ART-003`.
