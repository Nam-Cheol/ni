# Constraints

## Hard constraints

- Readiness must be rule-based, not model-feeling-based.
- Prompt output from `ni run` must be 4000 characters or less.
- The benchmark workspace must not implement a dashboard, scaffold frontend
  code, call model APIs, execute downstream agents, or run dashboard build
  commands.
- The repository root `.ni/plan.lock.json` must not be modified by this case.
- `ni end` and `ni run` must not run for this workspace while `ni status`
  reports `BLOCKED`.
- The accepted scope is benchmark planning-meeting artifact readiness, not
  dashboard product readiness.
- The packet must not include personal data, credentials, tokens, private
  customer data, confidential business metrics, production secrets, raw logs
  with sensitive identifiers, or source data not required for benchmark
  acceptance.

## Planning constraints

- "Customer team", account health, dashboard UI, product implementation, and
  production release readiness remain out of scope for this resolved benchmark
  artifact case.
- Author may fill the packet. Reviewer may verify completeness and consistency.
  Planning owner may approve acceptance. Read access should be limited to
  project members who need the benchmark case for planning or review.
- Evidence references must be minimal, scoped, and reviewable, without copying
  sensitive source content into the packet.
