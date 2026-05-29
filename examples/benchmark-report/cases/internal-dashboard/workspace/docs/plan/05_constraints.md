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

## Planning constraints

- "Customer team", "needs attention", "easy to use", and "next planning
  meeting" are not accepted implementation requirements until clarified.
- Customer account data, source systems, freshness, privacy, and access control
  must remain blocker questions until confirmed.
