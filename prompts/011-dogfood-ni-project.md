# 011 - Dogfood NI on NI

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Use NI's own docs and contract to lock the next increment and compile a prompt.

Goal:
Prove that NI can plan NI.

Scope:
- Run `ni status`.
- Resolve blockers if any.
- Run `ni end`.
- Run `ni run --max-chars 4000 --out .ni/generated/goal.prompt.txt`.
- Review generated prompt quality.

Non-goals:
- Do not execute the generated prompt yet.
- Do not add adapters.

Validation:
- gofmt -w . if Go changed
- go test ./... if Go exists
- bash scripts/quality.sh
- wc -m .ni/generated/goal.prompt.txt must be <= 4000
