# Blocker Analysis

This note explains why the internal-dashboard benchmark remains `BLOCKED`.
It does not answer the blocker questions, resolve them, defer them, lock the
workspace, compile a prompt, or authorize downstream dashboard work.

`BLOCKED` is a valid benchmark result. In this case, ni prevented premature
handoff by making readiness gaps explicit before implementation. The case did
not prove dashboard implementation quality; it proved that missing intent was
captured as auditable blocker evidence.

| Blocker | What is unknown? | Why it blocks lock | What kind of answer would resolve it? | What would be unsafe to assume? |
| --- | --- | --- | --- | --- |
| `OQ-001` | The exact primary dashboard user and the decision the dashboard should support are not accepted. | Without the actor and decision, requirements, layout, permissions, metrics, and meeting evidence would be built around guessed intent. | A user-confirmed role or audience and a concrete decision the dashboard must help that audience make. | That "customer team" means customer success managers, executives, support, sales, or any other role; that the decision is renewal risk, escalation, outreach, forecasting, or meeting reporting. |
| `OQ-002` | "Needs attention" is not defined as observable account signals or ranking criteria. | Without attention signals, the dashboard cannot have trustworthy acceptance criteria for account priority, sorting, alerts, or correctness. | User-confirmed signals, thresholds, ordering rules, or review criteria that define when an account needs attention for v0. | That attention means product usage drop, open support escalation, unpaid invoice, renewal date, health score, executive request, or any specific ranking formula. |
| `OQ-003` | Allowed source systems, account fields, freshness rules, privacy constraints, and access controls are not accepted. | Customer-account data and health signals can create privacy, security, and stale-data risk; lock would be unsafe without an explicit data boundary. | User-confirmed source systems, field list or field categories, freshness expectations, access audience, and privacy/security constraints for v0. | That trusted data already exists, that any CRM/support/billing/product data may be used, that sensitive fields may be shown, or that all customer-team members have access. |
| `OQ-004` | The required acceptance evidence for the next planning meeting is not accepted. | "Ready for the next planning meeting" does not define pass/fail evidence, date, audience, scope, or review artifact; lock would allow an unbounded handoff. | User-confirmed meeting date or timing, review audience, minimum useful artifact, and acceptance checks for showing that v0 planning is ready. | That a prototype, static mockup, planning memo, metric table, live dashboard, or implementation plan is enough; that "easy to use" has no explicit evidence. |

## Interpretation

The blocker set is useful because it stops execution at the correct boundary.
The direct prompt hides important assumptions; the ni path records them as open
questions tied to risks, evaluations, non-goals, and synchronized planning
state.

This evidence does not measure rework reduction, adoption, cost, latency,
statistical effect, downstream agent performance, or dashboard quality. It also
does not show that a future implementation would be correct. It shows that
`ni status` refused to treat ambiguous intent as ready.

## Current Result

- Readiness: `BLOCKED`
- Workspace locked: no
- Bounded prompt compiled: no
- Prompt character count: `not_measured`
- Downstream execution: none
