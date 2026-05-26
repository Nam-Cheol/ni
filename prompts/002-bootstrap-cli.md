# 002 - Bootstrap CLI

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement the smallest Go CLI scaffold.

Goal:
Create `ni --help` and `ni version` only.

Scope:
- go.mod
- cmd/ni/main.go
- internal/version or equivalent minimal package if useful
- scripts/quality.sh update only if needed

Non-goals:
- No ni init.
- No ni status.
- No lockfile.
- No prompt compiler.
- No adapters.

Validation:
- gofmt -w .
- go test ./...
- go run ./cmd/ni --help
- go run ./cmd/ni version
- bash scripts/quality.sh
