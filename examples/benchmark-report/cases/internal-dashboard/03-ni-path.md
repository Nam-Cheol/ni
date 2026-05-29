# ni Path

This case records the expected ni intent-lock path for the vague dashboard
request. It does not claim that a checked-in dashboard workspace reached
`READY`, `ni end`, or `ni run`.

## ni-start Questions

- Who are the primary users, and what decision should the dashboard help them
  make?
- Which account signals and source systems are allowed for the first version?
- What does "needs attention" mean in observable terms?
- What acceptance checks must pass before the planning meeting?
- Which dashboard behaviors are explicitly out of scope for this iteration?
- What privacy, access-control, or data-freshness constraints apply?

## Readiness Blockers

- Primary users and actor outcomes are not accepted yet.
- Delivery surface is assumed but not confirmed.
- Source systems, required account fields, and freshness constraints are not
  accepted yet.
- Customer-data exposure, incorrect prioritization, and stale-signal risks are
  not mitigated yet.
- Acceptance criteria for priority ranking, usability, performance, and
  planning-meeting review are missing.
- Non-goals such as CRM replacement, forecasting, workflow automation, and
  write-back behavior are not explicit.

## Docs and Contract Expectations

- `docs/plan/01_actors_outcomes.md` names the dashboard users and decisions.
- `docs/plan/02_capabilities.md` records accepted capabilities for account
  visibility, attention prioritization, and planning-meeting review.
- `.ni/contract.json` links each accepted capability to requirements,
  evaluations, risks, and artifacts.
- `docs/plan/05_constraints.md` records privacy, access-control, data source,
  and freshness constraints.
- `docs/plan/06_risks_security.md` records high-severity risks with
  mitigations.
- `docs/plan/07_evaluation_contract.md` records review evidence for
  correctness, freshness, usability, and meeting acceptance.
- `docs/plan/08_delivery_operation.md` records the delivery surface and
  non-execution handoff boundary.
- `docs/plan/10_open_questions.md` has no blocker open questions before lock.

## Lock Readiness Improvement

The ni path improves readiness only if it keeps the plan blocked until these
records exist or are explicitly and safely deferred. A future completed run
would need authoritative `ni status` output, then `ni end`, then a bounded
`ni run` prompt. Those artifacts are not present in this docs-only case.
