# Delivery and operation

## Delivery

The v0.2 kernel ships as a local Go CLI with repository-local Codex skills and Markdown planning docs. The primary authoring surface after `ni init` is model-user conversation, not manual contract editing commands.

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

1. Run `ni init` to create the initial planning structure.
2. Use model-user conversation with `ni-start` to update `docs/plan/**`, `.ni/contract.json`, and bounded `.ni/session.json` continuity state together.
3. Run `ni status` to get deterministic readiness gaps.
4. Resolve blockers without weakening accepted criteria.
5. Run `ni end` for a first lock or `ni relock` after an applied amendment.
6. Run `ni run --target <target> --max-chars 4000` to compile a bounded handoff prompt.
7. Treat downstream feedback as inert until it becomes an explicit amendment.

## Validation

When Go code exists, run:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

For this v0.2 planning contract, also verify:

```bash
ni status --dir .
ni run --dir . --target codex --max-chars 4000
```
