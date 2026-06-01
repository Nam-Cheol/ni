# Decision log

## DEC-001: Use a fresh isolated benchmark workspace for the research-protocol initial readiness measurement.

Status: accepted

Rationale: The second v0.5 benchmark case should not reuse the existing locked
research-protocol example as a new measurement. This workspace is isolated
under examples/benchmark-report/cases/research-protocol/workspace.

## DEC-002: Measure the initial vague research-protocol request only; do not resolve blockers in this task.

Status: accepted

Rationale: This task measures whether ni blocks unclear research intent before
fieldwork or handoff. OQ-001 through OQ-005 remain open blockers.

## DEC-003: Treat any later lock or prompt compilation as out of scope until blocker answers are provided and ni status is rerun.

Status: accepted

Rationale: A bounded prompt is allowed only after a valid lock in a later task.
This initial measurement records lock as no, bounded prompt as no, and prompt
count as not_measured when status is BLOCKED.
