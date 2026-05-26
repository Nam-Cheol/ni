# 005 - Readiness status

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement `ni status`.

Goal:
Return `BLOCKED`, `READY_WITH_DEFERRALS`, or `READY` using deterministic rules.

Scope:
- internal/core/readiness
- `ni status --dir <path>`
- `ni status --json`
- testdata fixtures

Rules to implement:
- Required planning docs exist.
- `.ni/contract.json` is valid.
- At least one capability exists.
- Every accepted capability has at least one linked evaluation.
- Every evaluation has method.
- Every high-severity risk has mitigation.
- Every accepted capability has at least one artifact or requirement.
- Decision statuses are valid.
- Blocker open questions prevent lock.
- At least one non-goal exists.

Validation:
- gofmt -w .
- go test ./...
- go run ./cmd/ni status --dir .
- go run ./cmd/ni status --dir . --json
- bash scripts/quality.sh
