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

1. Start with a model pack or copied instructions from this repository:
   `packages/claude-skills`, `packages/codex-skills`, or `.agents/skills`.
2. Ask the model to create a `docs/plan` draft for your project. It should cover
   purpose, actors, capabilities, requirements, decisions, risks, evaluations,
   non-goals, constraints, artifacts, and open questions.
3. Ask the model to draft `.ni/contract.json` alongside the docs. Treat it as a
   model-maintained draft, not authoritative state.
4. Mark assumptions and open questions explicitly. Tentative, conflicting, or
   incomplete statements should not become accepted decisions.
5. Later, ask a teammate, CI job, or local setup with the CLI to run
   `ni status`. If the result is blocked, continue the planning conversation.
6. Do not treat model judgment as a lock. When `ni status` reports that the
   plan is ready, use `ni end` to create the lock. Only the CLI should create
   `.ni/plan.lock.json`.
7. Use `ni run` to compile the final handoff prompt. `ni run` compiles text; it
   does not execute shell commands, agents, queues, or downstream work.

## No-Terminal Assisted Checklist

Use this checklist when you are starting without a local CLI:

- Start with a model pack or copied instructions.
- Create a `docs/plan` draft.
- Draft `.ni/contract.json` alongside the docs.
- Mark assumptions and open questions, especially blockers.
- Later validate with the CLI, a teammate, or a trusted runner.
- Do not treat model judgment as a lock.

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

## When to Graduate to Full ni

Move from no-terminal assisted drafting to full `ni` as soon as the plan might
guide implementation, budget, review, or downstream seed generation. In
particular, use the CLI before you claim readiness, create or trust a lockfile,
verify plan hashes, compile a bounded handoff prompt, or ask another actor to
start work from the plan.

If you cannot run the CLI locally, hand the draft to a teammate, CI job, or
trusted runner that can execute `ni status`, `ni end`, and `ni run`. Until that
happens, the workspace is useful for learning and drafting only.

See `examples/no-terminal-assisted/` for a docs-only example that keeps the
draft useful without claiming deterministic validation.

## Boundary

No-terminal planning must not add a hosted web app, model API calls, runtime
execution, shell adapters, queues, or automation behavior. It is a docs-first
way to start the Intent Lock method while preserving the kernel boundary:

```text
model pack or copied checklist -> draft docs and contract
ni CLI -> deterministic readiness, lock, hash proof, prompt compile
```
