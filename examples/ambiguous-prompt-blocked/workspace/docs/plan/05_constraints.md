# Constraints

## Hard constraints

- Do not execute Codex.
- Do not implement the dashboard product.
- Do not add shell scripts that run downstream tools.
- Do not add queue or task-runner behavior.
- Do not fake benchmark results.
- Prompt output from `ni run` must be 4000 characters or less.

## Source-of-truth constraint

After lock, `.ni/plan.lock.json` outranks `.ni/contract.json`, `docs/plan/**`,
session state, and chat history.
