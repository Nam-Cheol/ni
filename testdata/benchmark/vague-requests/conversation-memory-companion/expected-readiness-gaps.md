# Expected Readiness Gaps: Conversation Memory Companion

- Missing acceptance criteria: no criteria for memory accuracy, consent,
  retrieval relevance, deletion, project boundaries, or user trust.
- Unmitigated high-risk items: privacy leakage, unwanted inference, stale memory,
  and hallucinated preferences are high-risk without mitigation.
- Unresolved blockers: target environment, memory scope, sensitive-data policy,
  and user controls are undefined.
- Hidden assumptions: reviewer should count assumptions about storage,
  cross-project access, consent, tone, retention, and failure behavior.
- Non-goal coverage: none; the request does not exclude autonomous actions,
  surveillance, shared memory, or third-party integrations.
- Stale plan risk: memory policies and user preferences can change after lock,
  so stale intent must block handoff.
- Target prompt boundedness: not measurable until the ni path reaches a locked
  plan and compiles a target prompt.
- Readiness status before execution: expected to be `BLOCKED` until blockers
  are answered.
