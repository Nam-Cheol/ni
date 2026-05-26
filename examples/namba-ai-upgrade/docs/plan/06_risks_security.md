# Risks and security

## RISK-001: Planning may invent namba-ai behavior

Severity: high

Mitigation: Treat this example as upgrade planning only. Require downstream implementers to inspect namba-ai repository evidence before making namba-ai changes.

## RISK-002: Sequential SPEC pressure may hide collaboration conflicts

Severity: high

Mitigation: Use graph and dependency language. Require conflict review for changed decisions, requirements, risks, evaluations, and artifacts.

## RISK-003: Codex handoff may be mistaken for NI executing Codex

Severity: high

Mitigation: State that `ni run` compiles prompt text only. Codex use, if any, is manual and downstream.

## RISK-004: Dogfooding may pull execution behavior into NI

Severity: high

Mitigation: Keep this example limited to init, status, end, and run. Non-goals reject adapters, queues, SPEC runners, and execution tasks.

## RISK-005: Generated seed material may drift from the lock

Severity: medium

Mitigation: Generated files point back to `.ni/plan.lock.json`; downstream users stop with `BLOCKED` on lock mismatch.
