# ni v0.2.0 — Project Intent Compiler for AI Agents

Tag suggestion: `v0.2.0`

Summary: Don't run the agent yet. Compile the intent first.

This is a draft GitHub release note for a future `ni` release. It does not
publish a release, create a tag, upload binaries, or claim package manager
availability. Release binary assets are only available after a tagged GitHub
Release publishes them.

Current install status is tracked in [Install ni](22_INSTALL.md). This v0.2.0
draft is historical and does not override the verified v0.3.0 release binary
availability.

## Category

`ni` is a Project Intent Compiler for AI Agents. It turns planning
conversation into a checked project contract, locks the accepted intent, and
compiles bounded handoff material before downstream execution begins.

The core mechanism is the Intent Lock Protocol: a deterministic pre-runtime
control layer for deciding when intent is ready, what downstream actors may
trust, and when execution must stop because intent changed.

## Included

- Intent Lock Protocol.
- Conversation-driven authoring model.
- `ni init`, `ni status`, `ni end`, and `ni run`.
- Status proof report through `ni status --proof`.
- Locked plan hash validation.
- Target prompt and export seed outputs.
- Ambiguous prompt blocked demo.
- Non-software planning demos.
- Benchmark protocol.
- Korean companion README.
- Release binary pipeline for Linux, macOS, and Windows GitHub Release assets.

## Not Included

- Execution runtime.
- Codex adapter.
- Shell adapter.
- SPEC runner.
- Task queue.
- Multi-agent orchestration.
- Product-level PR/release automation or release commands.
- Package manager distribution.

## Distribution

This draft keeps source, local build, and local install paths supported:
`go run ./cmd/ni ...`, `make build`, and `make install-local`.

The repository includes GoReleaser configuration for future release assets:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

The release pipeline also generates `ni_<version>_checksums.txt`. This draft
does not claim Homebrew support, Scoop support, or published binary packages;
see [Install ni](22_INSTALL.md) for current release binary availability.

## Validation Commands

Run these checks before any manual tag or GitHub release step:

```bash
goreleaser check
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/release-check.sh
```

If GoReleaser is not installed locally, validate `.goreleaser.yaml` through the
tag-triggered release workflow before publishing assets.

## Release Boundary

`ni run` compiles a prompt only. It does not execute Codex, shells, model APIs,
queues, adapters, SPEC workflows, or multi-agent systems.

The tag-triggered repository release workflow is distribution infrastructure.
It is not part of `ni` runtime behavior and does not add release commands inside
the product.
