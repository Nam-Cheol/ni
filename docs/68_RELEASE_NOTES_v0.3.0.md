# ni v0.3.0 - Project Intent Compiler for AI Agents

Tag suggestion: `v0.3.0`

Summary: Don't run the agent yet. Compile the intent first.

These notes are draft release notes for a future public `ni` GitHub Release.
They do not publish a release, create a tag, upload binaries, or claim package
manager availability. Release binaries and curl installer availability:
available only after a tagged GitHub Release contains assets and checksums.

## Why v0.3.0

`v0.2.0` was used as a planning and differentiation label. The repository now
also contains the visual and distribution packaging work for the next public
surface: README product-pamphlet direction, local SVG assets, release-binary
distribution language, and model workspace pack docs. That makes `v0.3.0` the
first public GitHub Release.

## Category

`ni` is a Project Intent Compiler for AI Agents. It turns planning
conversation into a checked project contract, locks accepted intent, and
compiles bounded handoff material before downstream execution begins.

The core mechanism is the Intent Lock Protocol: a deterministic pre-runtime
control layer for deciding when intent is ready, what downstream actors may
trust, and when execution must stop because intent changed.

## Included

- Project Intent Compiler positioning.
- README as product pamphlet, with deeper technical detail in docs.
- Local deterministic SVG visual system for README and repository assets.
- Intent Lock Protocol docs and source-of-truth rules.
- `ni init`, `ni status`, `ni end`, and `ni run`.
- Status proof report through `ni status --proof`.
- Locked plan hash validation.
- Ambiguous prompt blocked demo.
- Non-software demos for planning outside software delivery.
- Model workspace packs for Codex- and Claude-style planning UX.
- Source-first usage through `go run ./cmd/ni ...`, local build, and local
  install.
- Release binary pipeline configuration for future GitHub Release assets.

## Not Included

- Task runner.
- SPEC runner.
- Multi-agent execution layer.
- Codex exec adapter.
- Shell adapter.
- Queue.
- PR automation.
- Release automation inside `ni-kernel`.
- Downstream execution runtime.
- Package manager distribution.

## Distribution

Available usage for this release preparation is source-first: `go run
./cmd/ni ...`, `make build`, and `make install-local`.

The repository includes GoReleaser configuration for future release assets:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

The release workflow must also generate `ni_<version>_checksums.txt`. These
notes do not claim hosted release assets, Homebrew support, Scoop support,
published binary packages, or curl installer availability before the GitHub
Release contains real assets and checksums.

These notes do not claim Homebrew support, Scoop support, package-manager
distribution, global model-pack installation, hosted release assets, or public
curl installer availability.

## Validation Commands

Run these checks before any manual tag or GitHub Release step:

```bash
bash scripts/release-dry-run.sh
bash scripts/quality.sh
go test ./...
```

## Release Boundary

`ni run` compiles a prompt only. It does not execute Codex, shells, model APIs,
queues, adapters, SPEC workflows, or multi-agent systems.

The tag-triggered repository release workflow is distribution infrastructure.
It is not part of `ni` runtime behavior and does not add release commands
inside the product.
