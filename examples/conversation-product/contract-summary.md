# Contract summary: Travel Concierge Triage

## Identity

- Product type: `conversation_product`
- Delivery surface: `conversation`
- Interaction mode: `human_to_system`
- Readiness profile: `prototype`
- Lock status: `READY`

## Purpose

Plan a bounded travel intake conversation that clarifies trip intent, catches
unsupported requests, and hands a concise brief to a human concierge.

## Accepted capabilities

- `CAP-001`: Collect trip goals, constraints, and traveler preferences through
  a bounded intake conversation.
- `CAP-002`: Detect unsupported, sensitive, or unsafe requests and move to
  human handoff language.
- `CAP-003`: Produce a concise human concierge handoff brief from the completed
  conversation.

## Evaluation model

- `EVAL-001`: Transcript fixture review for complete intake coverage.
- `EVAL-002`: Unsupported request escalation review.
- `EVAL-003`: Human concierge handoff rubric.

The evaluations are transcript review and checklist based. They are not
automated software tests.

## Non-goal boundary

ni does not deploy an assistant, contact vendors, make bookings, charge payment
methods, or provide regulated advice. `ni run` only compiles the locked plan
into a downstream handoff prompt.
