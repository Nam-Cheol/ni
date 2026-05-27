# Expected Readiness Gaps: Partner Sync CLI

- Missing acceptance criteria: no criteria for sync semantics, duplicate
  detection, authentication, dry run, error handling, logging, or rollback.
- Unmitigated high-risk items: data corruption, duplicate creation, unauthorized
  writes, and partner data exposure are high-risk without mitigation.
- Unresolved blockers: source of truth, API contract, user environment, record
  schema, and conflict policy are undefined.
- Hidden assumptions: reviewer should count assumptions about identifiers,
  operation type, access, data contracts, safety modes, and audit needs.
- Non-goal coverage: none; the request does not exclude background jobs,
  bidirectional sync, deletes, merge tooling, or admin UI.
- Stale plan risk: API schema or partner data policy changes after lock should
  block execution.
- Target prompt boundedness: not measurable until the ni path reaches a locked
  plan and compiles a target prompt.
- Readiness status before execution: expected to be `BLOCKED` until blockers
  are answered.
