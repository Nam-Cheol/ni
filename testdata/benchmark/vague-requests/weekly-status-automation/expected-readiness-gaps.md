# Expected Readiness Gaps: Weekly Status Automation

- Missing acceptance criteria: no criteria for source systems, summary quality,
  blocker detection, permissions, delivery timing, or human review.
- Unmitigated high-risk items: sensitive information leakage, false blocker
  claims, notification spam, and unauthorized tool access are high-risk without
  mitigation.
- Unresolved blockers: audience, data sources, output channel, schedule,
  timezone, and approval flow are undefined.
- Hidden assumptions: reviewer should count assumptions about tool access,
  source of truth, schedule, distribution, sensitivity, and automation level.
- Non-goal coverage: none; the request does not exclude autonomous posting,
  task reassignment, calendar changes, or cross-team surveillance.
- Stale plan risk: tool permissions, team membership, or reporting cadence
  changes after lock should block execution.
- Target prompt boundedness: not measurable until the ni path reaches a locked
  plan and compiles a target prompt.
- Readiness status before execution: expected to be `BLOCKED` until blockers
  are answered.
