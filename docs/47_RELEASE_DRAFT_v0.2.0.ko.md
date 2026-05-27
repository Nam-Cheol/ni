# ni v0.2.0 — Project Intent Compiler for AI Agents

Tag suggestion: `v0.2.0`

Summary: Don't run the agent yet. Compile the intent first.

이 문서는 future `ni` release를 위한 GitHub release note draft다. release를
publish하거나, tag를 만들거나, binary를 upload하거나, package manager
availability를 claim하지 않는다. Release binary assets는 tagged GitHub Release가
publish한 뒤에만 available하다.

## Category

`ni`는 AI agent를 위한 Project Intent Compiler다. planning conversation을
checked project contract로 만들고, accepted intent를 lock한 뒤, downstream
execution이 시작되기 전에 bounded handoff material을 compile한다.

핵심 mechanism은 Intent Lock Protocol이다. 이는 intent가 ready인지, downstream
actor가 무엇을 trust할 수 있는지, intent가 바뀌었을 때 execution을 언제 멈춰야
하는지를 결정하는 deterministic pre-runtime control layer다.

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

이 draft는 source, local build, local install path를 계속 support한다:
`go run ./cmd/ni ...`, `make build`, `make install-local`.

Repository에는 future release assets를 위한 GoReleaser configuration이 있다:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

Release pipeline은 `ni_<version>_checksums.txt`도 generate한다. 이 draft는
hosted release assets가 already available하다고 claim하지 않으며, Homebrew
support, Scoop support, published binary packages도 claim하지 않는다.

## Validation Commands

manual tag 또는 GitHub release step 전에 아래 checks를 실행한다:

```bash
goreleaser check
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/release-check.sh
```

GoReleaser가 local에 설치되어 있지 않다면, assets publish 전에 tag-triggered
release workflow에서 `.goreleaser.yaml`을 validate한다.

## Release Boundary

`ni run`은 prompt만 compile한다. Codex, shells, model APIs, queues, adapters,
SPEC workflows, multi-agent systems를 실행하지 않는다.

Tag-triggered repository release workflow는 distribution infrastructure다. 이는
`ni` runtime behavior의 일부가 아니며 product 내부에 release commands를 추가하지
않는다.
