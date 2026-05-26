# 001 - Reframe docs to ni-kernel

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Make documentation consistently describe `ni` as a project intent compiler.

Scope:
- README.md
- AGENTS.md
- docs/NI_BLUEPRINT.md
- docs/08_ROADMAP.md
- docs/00_START_HERE.md
- .ni/contract.json if needed

Required result:
- The first implementation target is contract/readiness/lock/prompt.
- Existing execution-harness concepts are preserved only as later generated harness work.

Non-goals:
- Do not implement Go code.
- Do not add Codex adapter.
- Do not add shell adapter.
- Do not add SPEC runner.

Validation:
- bash scripts/quality.sh
