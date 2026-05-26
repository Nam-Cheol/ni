# 012 - Codex exec experiment later

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Only after v0 works, add a documented experiment for running the compiled prompt with Codex exec.

Goal:
Document and optionally script a safe local experiment.

Scope:
- docs/experiments/codex-exec.md
- optional scripts/run-codex-prompt.sh

Required safety:
- Use workspace-write sandbox.
- Do not use dangerous bypass options.
- Require a clean git tree before running.
- Require valid lock hash before running.

Non-goals:
- Do not make Codex exec part of the kernel.
- Do not auto-run from `ni run`.

Validation:
- bash scripts/quality.sh
