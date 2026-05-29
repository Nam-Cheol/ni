# Actors and outcomes

## Actors

- Customer team: requested as the audience, but the exact role group is still a
  blocker because "customer team" could mean customer success, support, sales,
  account management, or leadership.
- Planning model: records the vague request, visible assumptions, risks, and
  blocker questions in `docs/plan/**`, `.ni/contract.json`, and
  `.ni/session.json`.
- NI CLI: validates readiness, prevents lock while blocker questions remain
  open, and compiles prompts only after a valid lock exists.

## Outcomes

- The customer-team actor should eventually know which accounts need attention,
  but the decision definition and observable account signals are not accepted
  yet.
- The planning state should make hidden assumptions visible before a downstream
  actor starts building.
- Execution must not start while the target users, source systems, data
  boundary, and review evidence remain unresolved.
