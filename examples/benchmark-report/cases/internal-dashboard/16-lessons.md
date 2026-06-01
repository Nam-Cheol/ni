# Lessons

This case is useful because the original request feels buildable. "Build a
dashboard for the customer team" sounds like a normal implementation prompt,
but the benchmark showed that the handoff was not ready until four blocker
questions were answered and accepted.

## Lessons From the BLOCKED State

- A plausible web surface can hide the real product decision.
- "Customer team" is not specific enough to identify the user, reviewer, or
  accountable planning owner.
- "Needs attention" is not an observable signal until the source fields,
  threshold or ranking rule, and freshness expectations are accepted.
- Privacy and access-control boundaries must be accepted before a customer data
  artifact can be trusted.
- A planning meeting can be a valid delivery boundary, but only when the
  artifact and pass/fail evidence are explicit.

## Lessons From the READY State

- `READY` became valid only after the scope shifted to benchmark
  planning-meeting artifact readiness.
- The lock is useful because it freezes the accepted artifact contract and the
  source-of-truth order before prompt compilation.
- The 4000-character prompt proves bounded handoff seed generation, not
  downstream success.
- The strongest evidence is the transition itself: hidden assumptions became
  visible blockers, then accepted artifact answers, then isolated lock data.

## Practical Takeaways

- Keep direct-to-agent risk notes next to CLI proof so the reader can see what
  was prevented.
- Preserve `not_measured` cells instead of filling them with optimism.
- State the delivery surface. In this case it is `document`, not a web
  dashboard.
- Treat product-readiness claims as a separate future measurement, not a side
  effect of artifact readiness.
- Do not run the compiled prompt as part of this benchmark; execution would
  measure a different system boundary.

## Reusable Pattern

For future benchmark cases, record the same sequence:

| Step | Required evidence |
| --- | --- |
| Vague request | Source request and direct-to-agent risk notes |
| Blocked ni path | `ni status --proof --next-questions` output and blocker list |
| Resolution path | Required answers, expected planning updates, and unsafe assumptions avoided |
| Answered variant | Status proof showing `READY` or `READY_WITH_DEFERRALS` |
| Lock | Isolated lock summary and source list |
| Prompt | Bounded prompt summary and character count |
| Limits | Remaining `not_measured` claims and non-execution confirmation |

The pattern should remain qualitative until repeated trials, independent
reviewers, and outcome measurements actually exist.
