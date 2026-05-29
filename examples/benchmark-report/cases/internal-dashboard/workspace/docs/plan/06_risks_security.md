# Risks and security

## RISK-001: Hidden implementation assumptions

Severity: high

Mitigation: Keep actor, attention-definition, source-system, and acceptance
questions as blocker open questions until the user answers or explicitly
defers them with rationale.

Risk: a direct downstream actor could build a dashboard around guessed users,
metrics, charts, deadlines, and success criteria.

## RISK-002: Unsafe customer-data exposure

Severity: high

Mitigation: Block lock until allowed source systems, required fields, data
freshness, privacy review, and access-control constraints are explicit.

Risk: a dashboard could expose sensitive customer health signals, account
status, or support context to the wrong audience if the data boundary is
guessed.

## RISK-003: Incorrect or stale account prioritization

Severity: high

Mitigation: Block lock until "needs attention" is defined with observable
signals, freshness expectations, and review evidence for planning-meeting use.

Risk: customer-team decisions may be distorted if attention ranking is based on
stale, incomplete, or undefined signals.
