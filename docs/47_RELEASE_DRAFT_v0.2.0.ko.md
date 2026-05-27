# ni v0.2.0 — Project Intent Compiler for AI Agents

Tag suggestion: `v0.2.0`

Summary: Don't run the agent yet. Compile the intent first.

이 문서는 future source-first `ni` release를 위한 GitHub release note draft다.
release를 publish하거나, tag를 만들거나, binary를 upload하거나, package manager
availability를 claim하지 않는다.

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

## Not Included

- Execution runtime.
- Codex adapter.
- Shell adapter.
- SPEC runner.
- Task queue.
- Multi-agent orchestration.
- PR/release automation.
- Binary package distribution.

## Source-First Distribution

이 draft는 repository source 사용만 전제한다: `go run ./cmd/ni ...`, local build,
또는 checked-out source tree에서의 local install. Hosted release assets,
Homebrew support, GoReleaser support, published binary packages를 claim하지
않는다.

## Validation Commands

manual tag 또는 GitHub release step 전에 아래 checks를 실행한다:

```bash
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
```

demo-check script가 나중에 추가되면, release notes를 publish하기 전에 optional
extra gate로 실행한다.

## Release Boundary

`ni run`은 prompt만 compile한다. Codex, shells, model APIs, queues, adapters,
SPEC workflows, multi-agent systems를 실행하지 않는다.

이 draft는 release automation을 추가하지 않는다. Tag 또는 GitHub release는 exact
release commit에서 local validation gate와 CI가 pass한 뒤 수동으로 만들어야 한다.
