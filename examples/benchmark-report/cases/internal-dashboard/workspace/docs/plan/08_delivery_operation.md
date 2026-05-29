# Delivery and operation

## Delivery surfaces

- Web dashboard, requested but not yet implementation-ready.

## Initial delivery

No dashboard delivery is authorized in this benchmark case. The only checked-in
artifact is an isolated ni planning workspace and report evidence showing that
readiness remains blocked before execution.

## Operating model

- Planning docs are committed to git.
- Contract JSON is committed to git.
- `ni status` may be run for readiness proof.
- `ni end` and `ni run` are not run unless the workspace reaches `READY` or
  `READY_WITH_DEFERRALS`.
- Downstream agents, dashboard builds, model APIs, queues, shell adapters, and
  PR or release automation are out of scope.
