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

## DEC-003: Scope readiness to the benchmark planning artifact

Status: accepted

Rationale: The user-provided answers for `OQ-001` through `OQ-004` focus on
planning-meeting artifact readiness. They do not answer real customer account
dashboard product requirements, so the benchmark may only claim readiness for
the benchmark packet and its review evidence.

## DEC-004: Keep approval role explicit even when the named person is unassigned

Status: accepted

Rationale: Acceptance may be approved by the planning owner or designated
benchmark case reviewer. A specific person can remain unassigned until the
project lead assigns one without turning the benchmark artifact readiness
decision back into a blocker.
