# ni v0.3.0 - Project Intent Compiler for AI Agents

Tag: `v0.3.0`

Summary: Don't run the agent yet. Compile the intent first.

These notes describe the first public `ni` GitHub Release. The release includes
OS/architecture binary archives and `ni_0.3.0_checksums.txt`. They do not claim
package manager availability or public curl installer availability.

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
- Source usage through `go run ./cmd/ni ...`, local build, and local install.
- Release binary archives for Linux, macOS, and Windows with checksums.

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

Available usage for this release includes source, local build, local install,
and manual GitHub Release binary download with checksum verification.

The release includes these asset patterns:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

The release also includes `ni_0.3.0_checksums.txt`. Manual binary installation
requires choosing the OS/arch archive, downloading the checksum file from the
same release, verifying the checksum, unpacking the binary, and running
`ni --help` and `ni version`.

These notes do not claim Homebrew support, Scoop support, package-manager
distribution, global model-pack installation, or public curl installer
availability.

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
