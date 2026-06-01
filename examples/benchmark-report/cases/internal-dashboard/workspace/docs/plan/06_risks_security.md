# Risks and security

## RISK-001: Hidden implementation assumptions

Severity: high

Mitigation: Record the scope decision that `OQ-001` through `OQ-004` resolve
benchmark artifact readiness only, preserve dashboard implementation and product
quality as non-goals, and require testable planning-meeting pass/fail criteria
before lock.

Risk: a direct downstream actor could build a dashboard around guessed users,
metrics, charts, deadlines, and success criteria.

## RISK-002: Unsafe customer-data exposure

Severity: high

Mitigation: Restrict the packet to project documentation, benchmark case files,
review notes, planning artifacts, issue or PR references, approved source
material required for benchmark validation, non-sensitive operational context,
and minimal evidence references. Prohibit personal data, credentials, tokens,
private customer data, confidential business metrics, production secrets,
sensitive raw logs, and unrelated production telemetry.

Risk: a dashboard could expose sensitive customer health signals, account
status, or support context to the wrong audience if the data boundary is
guessed.

## RISK-003: Incorrect or stale account prioritization

Severity: high

Mitigation: Define attention and ranking only for planning-meeting readiness:
decision supported, pass/fail criteria, minimum artifact, approval owner, and
source/privacy constraints rank above optional notes. Mark stale or uncertain
source status as stale or TBD instead of treating it as current.

Risk: customer-team decisions may be distorted if attention ranking is based on
stale, incomplete, or undefined signals.

## RISK-004: False product-readiness claim

Severity: high

Mitigation: State in the plan, contract, evidence files, and benchmark docs
that readiness applies only to the benchmark planning artifact. Do not claim
dashboard implementation quality, downstream agent performance, production
release readiness, rework reduction, adoption, cost, latency, or statistical
effect size.

Risk: reviewers could mistake a ready benchmark planning packet for proof that
the dashboard product itself is ready to build or release.
