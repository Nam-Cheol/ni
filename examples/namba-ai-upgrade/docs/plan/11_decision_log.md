# Decision log

## DEC-001: Use NI as the source-of-truth planning compiler

Status: accepted

Rationale: The upgrade needs locked intent and validation before namba-ai or any downstream harness receives implementation seed material.

## DEC-002: Represent SDD work as graph-compatible boundaries

Status: accepted

Rationale: A total-order SPEC chain would obscure independent work and make collaboration conflicts harder to detect.

## DEC-003: Compile the first downstream artifact with `ni run --target codex`

Status: accepted

Rationale: The task asks for a Codex target prompt while explicitly excluding Codex exec and namba-ai execution.
