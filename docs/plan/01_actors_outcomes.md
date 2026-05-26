# Actors and outcomes

## Actors

- User: owns scope, accepts amendments, and decides when to lock or relock.
- Planning model: drafts docs, detects gaps, and proposes contract changes without declaring readiness.
- CLI: validates readiness, verifies hashes, writes locks, and compiles bounded prompts or seed material.
- Downstream target: consumes locked prompts or exports without becoming kernel-owned state.
- Reviewer: compares planning states and resolves deterministic conflicts before relock.

## Outcomes

- The user has a v1 planning contract with accepted capabilities, requirements, evaluations, risks, and non-goals.
- Readiness is determined by `ni status`, not by model judgment.
- Execution starts from a locked contract and hash-verified planning docs.
- Feedback and pressure are visible without silently mutating accepted criteria.
- Parallel planning changes can be reviewed before they overwrite locked intent.
