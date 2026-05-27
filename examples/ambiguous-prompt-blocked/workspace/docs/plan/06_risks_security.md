# Risks and security

## RISK-001: Hidden implementation assumptions

Severity: high

Mitigation: Keep unresolved workflow and data-boundary questions as blocker open
questions until the user answers.

Risk: a direct agent may implement a dashboard based on guessed team, workflow,
data source, and success criteria.

## RISK-002: Unsafe customer-data handling

Severity: high

Mitigation: Block lock until the allowed data source and non-goals for live
systems are explicit.

Risk: customer data handling may be unsafe if the source and sensitivity
boundary are guessed.
