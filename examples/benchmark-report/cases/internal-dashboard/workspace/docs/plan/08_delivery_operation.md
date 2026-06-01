# Delivery and operation

## Delivery surfaces

- Document: an isolated benchmark planning workspace and benchmark report
  evidence files.

## Initial delivery

No dashboard delivery is authorized in this benchmark case. The checked-in
artifacts are an isolated ni planning workspace and report evidence showing the
measured readiness result after applying the user-provided answers.

The minimum accepted artifact is a completed answer packet containing all
required `OQ` fields, clear pass/fail criteria, explicit non-goals, and enough
evidence references for a reviewer to validate readiness.

## Operating model

- Planning docs are committed to git.
- Contract JSON is committed to git.
- `ni status` may be run for readiness proof.
- `ni end` and `ni run` are not run unless the workspace reaches `READY` or
  `READY_WITH_DEFERRALS`.
- Downstream agents, dashboard builds, model APIs, queues, shell adapters, and
  PR or release automation are out of scope.
- Acceptance may be approved by the planning owner or designated benchmark case
  reviewer. If no person is assigned yet, the approval owner remains unassigned
  until assigned by the project lead.
