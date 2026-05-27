# Actors and outcomes

## Actors

- User: asks for a customer-issue dashboard.
- Planning model: captures intent into docs and contract without declaring
  readiness.
- CLI: validates readiness and reports blocker questions.
- Downstream implementation agent: must wait until a plan is locked.

## Outcomes

- Ambiguous execution is blocked before implementation starts.
- Focused questions identify the missing workflow and data boundary.
- Any later Codex prompt is derived from a locked contract, not hidden chat
  assumptions.
