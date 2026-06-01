# Decision log

## DEC-001: Use a fresh isolated benchmark workspace for the research-protocol initial readiness measurement.

Status: accepted

Rationale: The second v0.5 benchmark case should not reuse the existing locked
research-protocol example as a new measurement. This workspace is isolated
under examples/benchmark-report/cases/research-protocol/workspace.

## DEC-002: Measure the initial vague research-protocol request only; do not resolve blockers in this task.

Status: rejected

Rationale: This was correct for the initial measurement task, but it is
superseded by the current resolved benchmark task that explicitly applies
synthetic fixture answers to OQ-001 through OQ-005.

## DEC-003: Treat any later lock or prompt compilation as out of scope until blocker answers are provided and ni status is rerun.

Status: rejected

Rationale: This is superseded by DEC-005. The resolved benchmark task provides
fixture answers and reruns status, so a bounded prompt is allowed only if the
isolated workspace first locks through `ni end`.

## DEC-004: Apply synthetic benchmark fixture answers to OQ-001 through OQ-005.

Status: accepted

Rationale: The current task explicitly asks the assistant to fill the research
protocol blockers with labeled synthetic fixture answers instead of waiting for
real user-provided research approval.

## DEC-005: Lock and compile only inside the isolated research-protocol benchmark workspace after readiness clears.

Status: accepted

Rationale: If `ni status` reports `READY` or `READY_WITH_DEFERRALS`, `ni end`
and `ni run --max-chars 4000` may run only against
examples/benchmark-report/cases/research-protocol/workspace. This does not
authorize fieldwork, downstream agents, model APIs, shell adapters, or research
implementation.
