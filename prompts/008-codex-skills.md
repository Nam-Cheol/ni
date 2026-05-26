# 008 - Codex skills

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Refine repository-local Codex skills for NI.

Goal:
The skills support the planning UX while preserving CLI authority.

Scope:
- .agents/skills/ni-start/SKILL.md
- .agents/skills/ni-end/SKILL.md
- .agents/skills/ni-run/SKILL.md
- optional static validation script for skill metadata

Required behavior:
- Each skill has name and description metadata.
- ni-start updates docs/plan and .ni/contract.json only.
- ni-end relies on ni status / ni end.
- ni-run relies on ni run and does not implement directly.

Non-goals:
- Do not add execution adapter.
- Do not add remote automation.

Validation:
- gofmt -w . if Go files changed
- go test ./... if Go exists
- bash scripts/quality.sh
