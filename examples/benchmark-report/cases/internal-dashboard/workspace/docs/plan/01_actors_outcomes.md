# Actors and outcomes

## Actors

- Planning meeting owner, product lead, or internal operations lead: primary
  user who needs to decide whether the internal-dashboard benchmark case is
  ready to use in a planning meeting.
- Engineering lead, QA or reviewer, documentation maintainer, and relevant
  stakeholders: secondary users who need to understand the benchmark result
  without inspecting implementation details directly.
- Planning model: records the vague request, visible assumptions, risks, and
  blocker questions in `docs/plan/**`, `.ni/contract.json`, and
  `.ni/session.json`.
- NI CLI: validates readiness, prevents lock while blocker questions remain
  open, and compiles prompts only after a valid lock exists.

## Outcomes

- The supported decision is whether the internal-dashboard benchmark case has
  enough structure, evidence, and acceptance criteria to be used in a planning
  meeting.
- The decision timing is before or during the next planning meeting where the
  benchmark case is reviewed for readiness.
- The packet must not decide final product direction, implementation scope,
  production release readiness, unresolved blocker resolution, or dashboard
  implementation quality.
- Execution must not start from this benchmark workspace; it may only produce
  planning evidence, an isolated lock if allowed, and a bounded prompt if a
  valid lock exists.
