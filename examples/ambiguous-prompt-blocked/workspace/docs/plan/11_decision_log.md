# Decision log

## DEC-001: Planning state before implementation

Status: accepted

Rationale: the user asked for a dashboard, but the team, workflow, and data
source are still blocker questions. The accepted action is planning capture,
not implementation.

## DEC-002: Downstream prompts require lock

Status: accepted

Rationale: a Codex target prompt may be compiled only after the plan is ready
and locked. This preserves ni-kernel as a pre-runtime control layer.
