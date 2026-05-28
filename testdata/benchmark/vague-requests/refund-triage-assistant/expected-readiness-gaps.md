# Expected Readiness Gaps: Refund Triage Assistant

- Missing acceptance criteria: no pass/fail criteria for recommendation scope,
  policy citation, escalation behavior, privacy handling, transcript review, or
  runtime boundary.
- Unmitigated high-risk items: refund authority and stale or unclear policy
  interpretation may affect customer-impacting decisions.
- Unresolved blockers: refund authority, authoritative policy source,
  escalation threshold, evaluation evidence, and downstream handoff target are
  not stated in the direct request.
- Hidden assumptions: reviewer should count assumptions about actors,
  permissions, policy source, privacy, escalation, evaluation, and runtime
  scope.
- Non-goal coverage: none; the request does not exclude refund approval,
  customer contact, supervisor replacement, or live integration work.
- Stale plan risk: if policy ownership or support workflow changes after lock,
  handoff should block until the plan is refreshed.
- Target prompt boundedness: not measurable until the ni path reaches a locked
  plan and compiles a target prompt.
- Readiness status before execution: expected to be `BLOCKED` until blocker
  questions are answered or explicitly deferred by policy.
