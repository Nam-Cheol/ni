# Delivery and operation

## Delivery

The v0.2 kernel ships as a local Go CLI with repository-local Codex skills and Markdown planning docs. The primary authoring surface after `ni init` is model-user conversation, not manual contract editing commands.

The v0.2 differentiation surface is documentation and proof-asset driven: README.md, README.ko.md, positioning docs, demos, benchmark protocol, status proof, target story, and release readiness checklist explain and verify the Intent Lock Protocol without starting downstream execution.

The v0.3 packaging surface adds README pamphlet strategy, SVG-first visual
identity, source-first plus release-binary distribution planning, and Codex- and
Claude-style model workspace packs. These surfaces help users understand and
adopt the kernel; they do not add `ni-kernel` runtime execution.

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

Packaging and distribution references:

- `README.md`
- `README.ko.md`
- `docs/52_README_PAMPHLET_STRATEGY.md`
- `docs/53_DISTRIBUTION_STRATEGY.md`
- `docs/54_HOMEBREW_DISTRIBUTION.md`
- `docs/22_INSTALL.md`
- `docs/69_MANUAL_RELEASE_STEPS.md`
- `docs/55_MODEL_WORKSPACE_PACKS.md`
- `docs/56_CLAUDE_SKILL_PACK.md`
- `docs/57_CODEX_SKILL_PACK.md`
- `docs/58_VISUAL_ASSETS.md`

## Operating model

1. Run `ni init` to create the initial planning structure.
2. Use model-user conversation with `ni-start` to update `docs/plan/**`, `.ni/contract.json`, and bounded `.ni/session.json` continuity state together.
3. Run `ni status` to get deterministic readiness gaps.
4. Resolve blockers without weakening accepted criteria.
5. Run `ni end` for a first lock or `ni relock` after an applied amendment.
6. Run `ni run --target <target> --max-chars 4000` to compile a bounded handoff prompt.
7. Treat downstream feedback as inert until it becomes an explicit amendment.
8. Use the differentiation proof assets to show why locked intent should precede downstream agent, human-team, or harness work.
9. Keep public packaging claims factual: release binaries come before non-Go
   curl and package-manager availability, model workspace packs remain UX, and
   no-terminal mode remains assisted unless CLI proof is supplied.

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
