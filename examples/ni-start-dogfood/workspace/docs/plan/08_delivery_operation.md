# Delivery operation

## Authoring workflow

1. `ni init` creates the workspace and template planning state.
2. The user provides a vague refund triage idea in conversation.
3. `ni-start` asks focused questions from the highest-impact gaps.
4. The user confirms scope, non-goals, and transcript evidence.
5. The model updates `docs/plan/**`, `.ni/contract.json`, and
   `.ni/session.json`.
6. `ni status` reports blocker gaps until the policy source question is
   resolved.
7. `ni-end` summarizes the CLI-ready plan and asks for explicit confirmation.
8. `ni run` compiles a target prompt from the locked plan.

## Operating boundary

This example stops at a locked plan and compiled prompt. No execution runtime,
adapters, queue, PR automation, or live support workflow is created.
