# Release Pipeline

This document describes the non-publishing release validation path for `ni`.
The release pipeline is repository distribution infrastructure. It is not
`ni-kernel` runtime behavior and must not become a task runner, SPEC runner,
multi-agent execution layer, adapter, queue, or project execution harness.

## Current Status

- Release binary availability: available for verified v0.3.0 GitHub Release
  assets and checksums.
- Curl installer availability: available for the verified v0.3.0 installer path
  against published release assets.
- Homebrew availability: not available.
- Local release validation: available through
  `bash scripts/release-dry-run.sh`.

## GitHub Release Workflow

The workflow lives at `.github/workflows/release.yml`.

- Trigger: `push` events for tags matching `v*` only.
- Permissions: `contents: write`, so GoReleaser can create release assets after
  a tag is pushed.
- Required steps:
  - check out the repository with `actions/checkout@v4`;
  - set up Go with `actions/setup-go@v5` and `go.mod`;
  - run `go test ./...`;
  - run `bash scripts/quality.sh`;
  - run `bash scripts/release-check.sh`;
  - run GoReleaser with `release --clean`.

The workflow must not be run from branches, pull requests, schedules, or manual
dispatch before this release path is intentionally expanded.

## GoReleaser Archive Matrix

`.goreleaser.yaml` builds `./cmd/ni` as the `ni` binary with `CGO_ENABLED=0`.

| OS | Arch | Archive |
| --- | --- | --- |
| linux | amd64 | `namba-intent_<version>_linux_amd64.tar.gz` |
| linux | arm64 | `namba-intent_<version>_linux_arm64.tar.gz` |
| darwin | amd64 | `namba-intent_<version>_darwin_amd64.tar.gz` |
| darwin | arm64 | `namba-intent_<version>_darwin_arm64.tar.gz` |
| windows | amd64 | `namba-intent_<version>_windows_amd64.zip` |

Windows arm64 is intentionally ignored. GoReleaser also writes
`namba-intent_<version>_checksums.txt`.

## Local Dry Run

Run the dry run before any tag is pushed:

```bash
bash scripts/release-dry-run.sh
```

The dry run executes:

```text
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/demo-check.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
goreleaser check, if GoReleaser is installed
goreleaser release --snapshot --clean, if GoReleaser is installed
```

If GoReleaser is not installed locally, the script prints explicit installation
and rerun instructions. The GoReleaser portion must not be treated as silently
passed when it was not executed.

## Stop Conditions

Stop before tagging if any validation fails, if the workflow no longer runs only
on `v*` tags, if the GoReleaser matrix changes without an explicit release
decision, or if documentation claims release binary, curl installer, or Homebrew
availability before those paths are verified. Homebrew remains planned until an
external package exists and is verified.
