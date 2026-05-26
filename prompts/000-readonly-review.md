# 000 - Read-only review

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Review the repository as read-only.

Goal:
Confirm that the repository is positioned as `ni-kernel`: a project intent compiler, not a task runner.

Scope:
- Read README.md, AGENTS.md, docs/NI_BLUEPRINT.md, docs/08_ROADMAP.md, .ni/contract.json, and prompts/.
- Do not edit files.

Output:
- Summarize the architecture.
- Identify any file that still suggests runner-first design.
- Recommend the next prompt to run.
