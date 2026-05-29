# ni Path

This case records a real ni intent-lock path for the vague dashboard request.
The checked-in workspace is intentionally pre-runtime and remains blocked until
the missing dashboard intent is answered.

Workspace:
`examples/benchmark-report/cases/internal-dashboard/workspace/`

## ni-start Questions

- Who are the primary users, and what decision should the dashboard help them
  make?
- Which account signals and source systems are allowed for the first version?
- What does "needs attention" mean in observable terms?
- What acceptance checks must pass before the planning meeting?
- Which dashboard behaviors are explicitly out of scope for this iteration?
- What privacy, access-control, or data-freshness constraints apply?

## Readiness Blockers

- `OQ-001`: primary dashboard user and supported decision are not accepted yet.
- `OQ-002`: "needs attention" is not defined in observable account signals.
- `OQ-003`: source systems, account fields, freshness rules, privacy
  constraints, and access controls are not accepted yet.
- `OQ-004`: planning-meeting acceptance evidence is not accepted yet.

## Docs and Contract Expectations

- `docs/plan/01_actors_outcomes.md` records the requested customer-team actor
  and keeps exact role/outcome as blocker intent.
- `docs/plan/02_capabilities.md` records accepted planning capabilities for
  capturing the request and blocking readiness.
- `.ni/contract.json` links accepted planning capabilities to requirements,
  evaluations, risks, and artifacts.
- `docs/plan/05_constraints.md` records the non-execution boundary and the
  unresolved data constraints.
- `docs/plan/06_risks_security.md` records three high-severity risks with
  mitigations that preserve the blockers instead of weakening them.
- `docs/plan/07_evaluation_contract.md` records planning-capture review and
  blocked-readiness proof.
- `docs/plan/08_delivery_operation.md` records that no dashboard delivery is
  authorized in this benchmark case.
- `docs/plan/10_open_questions.md` keeps four blocker open questions before
  lock.

## Measured CLI Evidence

Command run from the repository root on 2026-05-29:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
```

Readiness result:

```text
NI Intent Readiness: BLOCKED
```

The status proof records four blocker open questions, no deferrals, no
warnings, valid contract JSON, passing traceability, mitigated high risks,
recorded non-goals, and synchronized docs/contract records. See
`06-ni-status-proof.md` and `07-ni-next-questions.md`.

## Lock Readiness Improvement

The ni path improved readiness visibility by converting hidden dashboard
assumptions into explicit blocker questions and preventing execution. Because
the authoritative status is `BLOCKED`, this benchmark did not run `ni end`,
`ni relock`, or `ni run`; no lockfile or bounded handoff prompt exists for this
case.
