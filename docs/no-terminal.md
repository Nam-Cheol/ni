# No-Terminal Planning

You can benefit from `ni`'s planning method before you install the CLI. The
useful part is the discipline: make intent explicit, keep docs and a contract
draft aligned, surface blockers, and avoid handing work to an agent until the
plan has been checked.

No-terminal planning is not the same as validated `ni`. Deterministic readiness,
locking, lock hash verification, and prompt compilation require the CLI.

## Three Levels

| Level | How it works | What you can trust |
| --- | --- | --- |
| Full `ni` | The CLI is installed and `ni status`, `ni end`, and `ni run` are available. | Authoritative readiness, lock creation, lock hash checks, and bounded prompt compilation. |
| Model pack assisted | Claude/Codex-style skills guide planning-doc authoring and contract drafting. A user, teammate, or CI runner should run CLI validation before lock. | Useful model-assisted drafting, but readiness and lock claims are only authoritative after CLI output. |
| Read-only method | Copy the Intent Lock checklist or these instructions into a model session and ask it to reason through the plan. | Useful for learning and self-review; not authoritative and not equivalent to validated `ni`. |

## Manual Flow

1. Download or copy the relevant skill instructions from this repository:
   `packages/claude-skills`, `packages/codex-skills`, or `.agents/skills`.
2. Ask the model to create a `docs/plan` draft for your project. It should cover
   purpose, actors, capabilities, requirements, decisions, risks, evaluations,
   non-goals, constraints, artifacts, and open questions.
3. Ask the model to maintain a `.ni/contract.json` draft alongside the docs.
   Tentative or conflicting statements should stay as assumptions, draft
   records, or open questions.
4. Keep a visible checklist of blockers. Open blocker questions must prevent
   lock, even if the plan feels close.
5. Later, ask a teammate, CI job, or local setup with the CLI to run
   `ni status`. If the result is blocked, continue the planning conversation.
6. When `ni status` reports that the plan is ready, use `ni end` to create the
   lock. Only the CLI should create `.ni/plan.lock.json`.
7. Use `ni run` to compile the final handoff prompt. `ni run` compiles text; it
   does not execute shell commands, agents, queues, or downstream work.

## Intent Lock Checklist

Use this checklist when you are working without a terminal:

- Is the project purpose explicit?
- Are actors and outcomes named?
- Does every capability trace to at least one requirement and evaluation?
- Are non-goals visible?
- Are high-severity risks paired with mitigations?
- Are open questions marked clearly, especially blockers?
- Are accepted decisions separated from assumptions and rejected options?
- Are expected artifacts named?
- Is the downstream handoff bounded to planning output, not runtime execution?

This checklist is a learning and drafting aid. It can help a model ask better
questions, but it does not replace `ni status`, `ni end`, or `ni run`.

## Boundary

No-terminal planning must not add a hosted web app, model API calls, runtime
execution, shell adapters, queues, or automation behavior. It is a docs-first
way to start the Intent Lock method while preserving the kernel boundary:

```text
model pack or copied checklist -> draft docs and contract
ni CLI -> deterministic readiness, lock, hash proof, prompt compile
```
