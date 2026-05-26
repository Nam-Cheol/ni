# Project brief

## Product type

Software upgrade planning.

## Delivery surfaces

- CLI: namba-ai and NI command handoffs remain command-oriented, but this example does not run namba-ai.
- Document: locked planning docs and generated prompt artifacts carry the handoff.
- Workflow: downstream teams or agents may derive work packets after lock verification.

## Purpose

Create a locked NI planning contract for upgrading namba-ai before any runtime, SPEC runner, or Codex execution harness is used.

## Current namba-ai limitations

- Upgrade planning can become coupled to runtime commands before intent is locked.
- Acceptance criteria and evidence expectations can be implied by a SPEC sequence instead of traced to stable CAP, REQ, EVAL, and RISK IDs.
- Collaboration is difficult when branches or agents edit sequential SPEC steps with overlapping decisions.
- Codex-oriented work can appear to be the source of truth when it should only consume locked planning seed material.
- Downstream target material can blur into mutable execution state if lock hashes and ownership boundaries are not explicit.

## Problem

namba-ai needs an upgrade plan that benefits from SDD-style structure without inheriting a mandatory linear execution model. The planning layer must establish intent, non-goals, risks, validations, and downstream handoff targets before implementation work starts.

## Success definition

The example is successful when `ni status` is ready or ready with deferrals, `ni end` creates `.ni/plan.lock.json`, and `ni run --target codex` writes a prompt under `generated/` without executing namba-ai or Codex.
