# ni v0.3.0 - Project Intent Compiler for AI Agents

Tag suggestion: `v0.3.0`

Summary: Don't run the agent yet. Compile the intent first.

이 문서는 첫 public `ni` GitHub Release를 위한 draft release notes다. Release를
publish하거나, tag를 만들거나, binary를 upload하거나, package manager
availability를 claim하지 않는다. Release binaries와 curl installer
availability는 tagged GitHub Release가 assets와 checksums를 포함한 뒤에만
available하다.

## Why v0.3.0

`v0.2.0`은 planning 및 differentiation label로 사용되었다. Repository는 이제
README product pamphlet 방향, local SVG assets, source-first distribution
language, model workspace pack docs 같은 visual/distribution packaging work도
포함한다. 그래서 첫 public GitHub Release candidate version은 `v0.3.0`이다.

## Category

`ni`는 AI Agents를 위한 Project Intent Compiler다. Planning conversation을
checked project contract로 만들고, accepted intent를 lock한 뒤, downstream
execution이 시작되기 전에 bounded handoff material을 compile한다.

핵심 mechanism은 Intent Lock Protocol이다. 이는 intent가 ready인지, downstream
actor가 무엇을 trust할 수 있는지, intent가 바뀌었을 때 execution을 언제 멈춰야
하는지를 결정하는 deterministic pre-runtime control layer다.

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

이 release preparation에서 available usage는 source-first path다:
`go run ./cmd/ni ...`, `make build`, `make install-local`.

Repository에는 future release assets를 위한 GoReleaser configuration이 있다:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

Release workflow는 `ni_<version>_checksums.txt`도 generate해야 한다. 이 문서는
GitHub Release가 실제 assets와 checksums를 포함하기 전에 hosted release assets,
Homebrew support, Scoop support, published binary packages, curl installer
availability를 claim하지 않는다.

## Validation Commands

manual tag 또는 GitHub Release step 전에 아래 checks를 실행한다:

```bash
bash scripts/release-dry-run.sh
bash scripts/quality.sh
go test ./...
```

## Release Boundary

`ni run`은 prompt만 compile한다. Codex, shells, model APIs, queues, adapters,
SPEC workflows, multi-agent systems를 실행하지 않는다.

Tag-triggered repository release workflow는 distribution infrastructure다. 이는
`ni` runtime behavior의 일부가 아니며 product 내부에 release commands를 추가하지
않는다.
