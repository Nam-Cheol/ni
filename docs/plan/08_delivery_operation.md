# Delivery and operation

## Delivery

The v1 kernel ships as a local Go CLI with repository-local Codex skills and Markdown planning docs.

Tracked planning state:

- `docs/plan/**`
- `.ni/contract.json`
- `.ni/readiness.rules.json`
- `.ni/readiness.profiles.json`
- `.ni/plan.lock.json`

Generated or inert state:

- `.ni/generated/**`
- `.ni/feedback.jsonl`
- `.ni/pressure.json`
- `.ni/harness.candidates.json`
- `.ni/amendments/**`
- `.ni/locks/**`

## Operating model

1. Edit planning docs and `.ni/contract.json`.
2. Run `ni status`.
3. Resolve blockers without weakening accepted criteria.
4. Run `ni end` for a first lock or `ni relock` after an applied amendment.
5. Run `ni run --target <target> --max-chars 4000` or `ni export --target <target> --out <dir>`.
6. Treat downstream feedback as inert until it becomes an explicit amendment.

## Validation

When Go code exists, run:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

For this v1 planning contract, also verify:

```bash
ni status --dir .
ni run --dir . --target codex --max-chars 4000
```
