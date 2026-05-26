# Decision log

## DEC-001: Internal refund policy page is authoritative

Status: accepted

The internal refund policy page is the source of truth for refund eligibility.
If ticket text conflicts with policy, the assistant escalates.

## DEC-002: Transcript fixtures are the first evaluation method

Status: accepted

The first evaluations are eligible refund, ineligible refund, and ambiguous
policy escalation transcript fixtures.

## DEC-003: human-team is the first prompt target

Status: accepted

The first compiled handoff target after lock is `human-team`.

## DEC-004: Production policy owner

Status: deferred

The production owner for maintaining the refund policy page is still deferred.
The deferral is visible in the contract and lock context instead of being
silently resolved.
