# ni v0.4.0 Release Preflight

Date: 2026-05-29

Release: `ni v0.4.0 - Conversation Authoring Hardening`

Status: Preflight record only. This document does not create a tag, push tags,
publish a GitHub Release, upload assets, mark Homebrew available, or add
runtime execution behavior.

## Scope Confirmation

`v0.4.0` covers adoption hardening after `v0.3.0`:

- first-run proof text for `R014`, `R015`, and `R016`;
- first-run `ni-start` conversation card;
- `SYNC-014`, `SYNC-015`, and `SYNC-016` docs/contract sync diagnostics;
- grouped and prioritized next questions;
- grouped-question usage in `ni-start` and model workspace packs;
- example coverage and maintained Korean companion docs;
- no-terminal assisted workflow hardening;
- benchmark case study expansion;
- Homebrew deferred to `v0.5`;
- conversation proof capture;
- v0.4.0 release planning.

Not included:

- task runner;
- SPEC runner;
- execution harness;
- Codex exec adapter;
- shell adapter;
- downstream agents;
- queue;
- PR automation;
- release automation inside `ni-kernel`;
- Homebrew availability.

No runtime execution behavior is part of this release. `ni run` remains a
bounded prompt compiler and must not execute shells, agents, queues, adapters,
or downstream work.

## Git State

Observed during preflight:

| Check | Result |
| --- | --- |
| Branch | `main` |
| Working tree | clean before preflight edits |
| `HEAD == origin/main` | yes before preflight edits: `6f80074eb09fbcc5e81b64abd98f172ea69c7e66` |
| Recent tags | `v0.3.0`, `v0.1.0` |
| `v0.4.0` tag absent before release | yes |

The preflight task may add this document, its Korean companion, and release
check updates. Re-check the final working tree before creating any tag.

## Version State

Expected release version: `0.4.0`.

The checked-in development version source is
`internal/version/version.go`, where `Version` defaults to `0.0.0-dev`.
`cmd/ni/main.go` prints that value for `ni version`.

GoReleaser injects the release value through `.goreleaser.yaml`:

```text
-s -w -X ni/internal/version.Version={{ .Version }}
```

Therefore the repository should not hard-code `0.4.0` into
`internal/version/version.go`. A binary release built by GoReleaser from tag
`v0.4.0` should report `0.4.0` with `ni version`.

Required preflight conclusion: binary release should report `0.4.0`.

## Release Workflow

`.github/workflows/release.yml` runs on tag pushes matching:

```yaml
push:
  tags:
    - "v*"
```

The workflow checks out the repository with full history, sets up Go from
`go.mod`, runs:

- `go test ./...`;
- `bash scripts/quality.sh`;
- `bash scripts/release-check.sh`;
- `goreleaser/goreleaser-action@v6` with `release --clean`.

This confirms the release workflow will run on a `v0.4.0` tag push. The manual
release operator must still verify the resulting assets and checksums.

## Validation Commands

Run before tagging:

- `go test ./...`
- `bash scripts/quality.sh`
- `bash scripts/smoke.sh`
- `bash scripts/demo-check.sh`
- `bash scripts/install-check.sh`
- `bash scripts/release-check.sh`
- `bash scripts/fresh-install-check.sh` if present
- `bash scripts/check-skill-packs.sh`
- `bash scripts/package-claude-skills.sh`
- `bash scripts/package-codex-skills.sh`
- `bash scripts/release-dry-run.sh`

GoReleaser local dry run:

- If GoReleaser is installed, run `goreleaser check` and
  `goreleaser release --snapshot --clean`.
- If GoReleaser is not installed, record that the local GoReleaser portion was
  skipped and must be covered by GitHub Actions or another machine with
  GoReleaser installed.

## Validation Results

Current-task results:

| Command | Result |
| --- | --- |
| `go test ./...` | passed |
| `bash scripts/quality.sh` | passed |
| `bash scripts/smoke.sh` | passed |
| `bash scripts/demo-check.sh` | passed |
| `bash scripts/install-check.sh` | passed |
| `bash scripts/release-check.sh` | passed |
| `bash scripts/fresh-install-check.sh` | passed with network access; validates current published `v0.3.0` release assets because `v0.4.0` assets do not exist before release |
| `bash scripts/check-skill-packs.sh` | passed |
| `bash scripts/package-claude-skills.sh` | passed; created `dist/ni-claude-skills.zip` |
| `bash scripts/package-codex-skills.sh` | passed; created `dist/ni-codex-skills.zip` |
| `bash scripts/release-dry-run.sh` | passed without creating tags, pushing, or publishing |
| `goreleaser check` | skipped; `goreleaser` is not installed locally |
| `goreleaser release --snapshot --clean` | skipped; `goreleaser` is not installed locally |

Local GoReleaser config validation is limited to repository checks in
`bash scripts/release-check.sh` because the `goreleaser` binary is not
installed on this machine. The GoReleaser check and snapshot archive build must
be covered by GitHub Actions or another machine with GoReleaser installed
before publishing release assets.

## Manual Release Steps

Run these only after the preflight passes:

```bash
git status --short --branch
git fetch origin --tags
git tag --list v0.4.0
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/demo-check.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
bash scripts/fresh-install-check.sh
bash scripts/check-skill-packs.sh
bash scripts/package-claude-skills.sh
bash scripts/package-codex-skills.sh
bash scripts/release-dry-run.sh
git tag -a v0.4.0 -m "ni v0.4.0"
git push origin v0.4.0
```

After the tag push:

1. Wait for the GitHub Actions release workflow to finish.
2. Verify release assets and `ni_0.4.0_checksums.txt`.
3. Verify the current-platform binary with `ni --help` and `ni version`.
4. Verify the curl installer against `--version 0.4.0`.
5. Update install docs only if availability status changes after verification.

## Guardrails

- No tag creation during preflight.
- No tag push during preflight.
- No GitHub Release publication during preflight.
- No Homebrew availability claim.
- No runtime execution behavior.
- No shell adapter.
- No Codex exec adapter.
- No downstream agents.
- No queues.
- No user-facing contract `add`, `list`, or `set` commands.
- No `ni end`.
- No `ni relock`.
- No manual `.ni/plan.lock.json` edits.
