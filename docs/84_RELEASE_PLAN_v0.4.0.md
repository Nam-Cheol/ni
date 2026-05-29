# ni v0.4.0 — Conversation Authoring Hardening

Date: 2026-05-29

Status: Draft release plan only. This document does not publish a release,
create a tag, push tags, create a GitHub Release, upload assets, mark
Homebrew available, or add runtime execution.

## Version Decision

The latest repository tags are `v0.3.0` and `v0.1.0`. No `v0.4.0` tag is
present, so the next planned release is `v0.4.0`. If a `v0.4.0` tag appears
before this plan is used, prepare `v0.4.1` instead.

## Release Summary

`ni v0.4.0` hardens the adoption path after installation. The main improvement
is conversation-to-contract authoring quality, not runtime execution.

The release story is:

```text
ni init -> ni-start conversation -> docs/contract/session updates -> ni status --proof --next-questions
```

`ni` remains `ni-kernel`: a pre-runtime Project Intent Compiler for AI Agents.
The kernel owns the docs contract, readiness gate, lockfile, prompt compiler,
and source-of-truth rule. It does not execute downstream work.

## Included Changes

### Conversation authoring

- First-run `ni-start` conversation card for fresh `R014`, `R015`, and `R016`
  blockers.
- Planning proof capture so model-authored turns can show the exact CLI
  readiness basis after docs, contract, and session updates.
- Model edit safety improvements for minimal, visible, conversation-tied
  planning diffs; ambiguous statements stay assumptions, drafts, or open
  questions until confirmed.

### Readiness proof

- Concrete first-run proof text for `R014`, `R015`, and `R016`.
- Better blocker `Why it matters` and `Next` wording in
  `ni status --proof`.
- `ni status --proof --next-questions` remains the visible explanation loop
  after a blocked planning turn.

### Docs/contract sync

- `SYNC-014` for purpose drift between docs and `.ni/contract.json`.
- `SYNC-015` for actor/outcome drift between docs and contract narrative
  records.
- `SYNC-016` for delivery-surface drift between docs and contract.
- Sync diagnostics in proof output, with stable `ID`, `Location`, `Problem`,
  `Why it matters`, `Suggested repair`, and `Blocks ni-end` fields.
- JSON proof/status output may include `sync_diagnostic` on affected `R012`
  `issues[]` and `proof[]` entries.

### Next questions

- Grouped question output for the next planning turn.
- Prioritization and deduplication so repair questions do not shadow or repeat
  lower-value generic prompts.
- First-run card questions for purpose, actors/outcomes, and delivery surface.
- Sync repair questions for `SYNC-014`, `SYNC-015`, and `SYNC-016`.
- Risk-decision, evaluation-evidence, scope-boundary, and open-blocker groups.

### Model workspace packs

- Grouped questions are wired into `ni-start` skills and packaged workspace
  guidance.
- The rule remains: Skills are UX; CLI is authority.
- Packs must not claim downstream execution, runtime readiness, or a second
  readiness engine.

### Examples

- Improved example coverage matrix and demo expectations.
- Korean companion docs where companion docs are maintained.
- `ni-start` dogfood updates for first-run cards, grouped questions,
  docs/contract/session updates, and re-status loops.
- No-terminal assisted flow hardening for drafting before trusted CLI proof.
- Intent-readiness benchmark case study expansion without fake empirical
  claims or model API calls.

### Distribution

- Release binary remains Available after release asset and checksum
  verification.
- Curl installer remains Available after verification against the new release.
- Homebrew remains Planned and deferred to v0.5 unless separately implemented
  and tested before a future release plan changes that status.

## Not Included

- No task runner.
- No SPEC runner.
- No execution harness.
- No Codex exec adapter.
- No shell adapter.
- No downstream agents.
- No queue.
- No PR automation.
- No release automation inside `ni-kernel`.
- No Homebrew availability in this release unless separately implemented and
  tested.

These exclusions are release boundaries. They keep `ni-kernel` focused on the
Intent Lock Protocol rather than turning it into a runtime.

## Compatibility Notes

- CLI command names remain stable for `ni init`, `ni status`, `ni end`, and
  `ni run`.
- Human-readable `ni status --proof` and `--next-questions` output changed:
  proof text is more specific, next questions are grouped, and first-run or
  sync-repair cards may alter snapshot/docs expectations.
- JSON output has additive shape changes when matching flags are requested:
  `next_questions` is present only with `--next-questions`, `proof` is present
  only with `--proof`, and `sync_diagnostic` may appear on `R012` `issues[]`
  and `proof[]` entries.
- Human-only proof grouping, `Why it matters`, and `Next` wording are not added
  to the proof JSON schema.
- `ni run` remains a prompt compiler only and must still produce bounded prompt
  output, not execute shells, agents, queues, adapters, or downstream work.

## Validation Checklist

Before tagging or publishing, collect evidence from:

- [ ] `go test ./...`
- [ ] `bash scripts/quality.sh`
- [ ] `bash scripts/smoke.sh`
- [ ] `bash scripts/demo-check.sh`
- [ ] `bash scripts/install-check.sh`
- [ ] `bash scripts/release-check.sh`
- [ ] `bash scripts/fresh-install-check.sh` if present
- [ ] `bash scripts/check-skill-packs.sh`
- [ ] package Claude skills: `bash scripts/package-claude-skills.sh`
- [ ] package Codex skills: `bash scripts/package-codex-skills.sh`
- [ ] release dry run before tag: `bash scripts/release-dry-run.sh`

The checklist is evidence collection. It does not publish release assets, push
tags, update a Homebrew tap, or mark unverified distribution paths Available.

## Manual Release Steps

1. Ensure the git tree is clean and the release plan matches the intended
   release version.
2. Run the validation checklist and preserve the important proof output.
3. Run the release dry run before creating any tag.
4. Create an annotated tag for `v0.4.0`.
5. Push the tag.
6. Wait for the GitHub Actions release workflow to finish.
7. Verify release assets and checksums.
8. Verify the current-platform binary from the release assets.
9. Verify the curl installer against the new release.
10. Only then update docs if an availability status changes.

Do not run these manual release steps while preparing this plan.

## Release Risk Notes

- Proof wording changes could affect snapshots, walkthroughs, and docs that
  quote status output.
- Examples and docs need to stay aligned with grouped next-question output.
- Model workspace packs remain UX, not authority; only CLI gates decide
  readiness, lock, and prompt compilation.
- Homebrew remains Planned, so docs must not imply package-manager
  availability.
- No-terminal remains assisted, not deterministic; readiness still requires
  trusted CLI proof.

## Next Release Direction

`v0.5` candidates:

- Homebrew decision and implementation if it is still useful.
- Stronger benchmark evidence and more transparent case-study reporting.
- More product surfaces and example coverage beyond software-only planning.
- Continued conversation authoring reliability work.
- Optional landing page after the installation and adoption surfaces are
  stable.
- Downstream integrations only as separate packages, not `ni-kernel`
  execution behavior.
