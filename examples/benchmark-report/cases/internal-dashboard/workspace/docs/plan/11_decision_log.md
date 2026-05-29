# Decision log

## DEC-001: Treat the dashboard request as planning evidence, not build scope

Status: accepted

Rationale: The original request is intentionally vague. The benchmark case must
show pre-runtime readiness blocking before downstream execution starts.

## DEC-002: Use the benchmark workspace only for ni CLI readiness proof

Status: accepted

Rationale: The case may run `ni init` and `ni status`; it must not run
`ni end` or `ni run` while blocker questions remain open, and it must not
implement the dashboard.
