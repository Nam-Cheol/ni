# Delivery and operation

## Delivery surfaces

- document
- workflow
- human_service

## Initial delivery

Initial delivery is a checked-in benchmark case with planning docs, a contract,
status proof, next-question evidence, a measurement table, and not_measured
boundaries. It is not a research protocol approved for fieldwork.

## Operating model

- Planning docs are committed to git.
- Contract JSON is committed to git.
- `ni status --proof --next-questions` is the authority for measured
  readiness.
- If readiness is `BLOCKED`, do not lock and do not compile a prompt.
- A later task may collect user answers and rerun status, but this task keeps
  the initial blockers unresolved.
