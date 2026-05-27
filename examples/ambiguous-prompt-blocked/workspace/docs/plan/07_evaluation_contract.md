# Evaluation contract

## EVAL-001: Planning record review

Method: Review docs/plan and .ni/contract.json to confirm the vague request is represented as planning state, not executable implementation scope.

Evidence: the workspace contains blocker open questions and no downstream
implementation artifacts.

## EVAL-002: Blocked readiness check

Method: Run ni status for the workspace and verify it reports BLOCKED with R009 blocker open-question issues.

Evidence: `ni status --dir examples/ambiguous-prompt-blocked/workspace` should
report `BLOCKED`.
