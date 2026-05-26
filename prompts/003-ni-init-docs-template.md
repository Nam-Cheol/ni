# 003 - ni init docs template

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement `ni init`.

Goal:
`ni init --dir <path>` creates initial planning docs and `.ni` skeleton in an empty or existing directory without destroying user files.

Scope:
- cmd/ni command routing
- internal/core/docstore or equivalent
- embedded or file-based templates
- tests using temporary directories

Required generated files:
- docs/plan/*.md
- .ni/project.json
- .ni/contract.json
- .ni/readiness.rules.json

Non-goals:
- No readiness validation yet beyond basic write checks.
- No lockfile.
- No prompt compiler.
- No adapters.

Validation:
- gofmt -w .
- go test ./...
- go run ./cmd/ni init --dir /tmp/ni-init-test
- bash scripts/quality.sh
