# ni

[English](README.md) | [한국어](README.ko.md)

ni는 AI agent를 위한 Project Intent Compiler다.

agent를 아직 실행하지 마라. 먼저 의도를 컴파일하라.

`ni`는 Codex, Claude, Spec Kit, Hyper Run, namba-ai, generated harness, 또는
human team이 실행을 시작하기 전에 planning conversation을 잠긴, versioned,
검증 가능한 project contract로 바꾼다.

현재 제품은 `ni-kernel`이다. 이것은 intent를 위한 deterministic pre-runtime
control layer이며, 구현을 실행하는 하네스가 아니다.

```text
conversation -> docs/plan + .ni/contract.json -> ni status -> ni end -> locked intent -> ni run
```

## What Problem ni Solves

Agents는 실행 가능해 보이지만 trustworthy execution에 필요한 intent를 숨기는
prompt를 자주 받는다:

- 이 project가 누구를 위한 것인지,
- 무엇이 accepted되어야 하는지,
- 어떤 risks에 mitigation이 필요한지,
- 무엇이 명시적으로 scope 밖인지,
- 어떤 questions가 execution을 막아야 하는지,
- accepted plan이 바뀌었는지.

대부분의 tools는 prompt, spec, worklist, runtime loop가 이미 존재한 뒤에
agent를 제어하려고 한다. `ni`는 control을 더 앞쪽으로 옮긴다. `ni`는
downstream actor가 work를 시작하기 전에 project intent가 explicit, accepted,
validated, locked, unchanged 상태인지 묻는다.

## Core Idea: Intent Lock Protocol

[Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md)은 planning conversation이
project contract가 되는 방식, contract가 lock될 준비가 되는 시점, accepted
plan이 hash되는 방식, downstream actors가 trust할 수 있는 것, intent가
바뀌었기 때문에 handoff가 멈춰야 하는 시점을 정의한다.

kernel이 소유하는 것은 다음이다:

- `docs/plan/**`
- `.ni/contract.json`
- deterministic readiness validation
- `.ni/plan.lock.json`
- lock hash verification
- bounded prompt compilation
- inert downstream seed exports

plan이 locked된 뒤 source-of-truth precedence는 다음과 같다:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

locked hash가 더 이상 일치하지 않으면 target handoff는 `BLOCKED`로 멈춘다.

## 5-Minute Demo

`ni`를 가장 빠르게 이해하는 방법은
[Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) demo를 보는
것이다.

이 demo는 다음 request에서 시작한다:

```text
Build me a dashboard for my team.
```

direct-to-agent 경로라면 users, data, workflow, non-goals, success criteria에
관한 hidden assumptions를 만들어내야 한다. `ni` 경로는 request를 planning
intent로 기록한 뒤, blocker questions가 열려 있는 동안 executable한 것으로
취급하기를 거부한다.

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected result:

```text
BLOCKED
```

핵심은 이것이다. 모호한 execution은 agent가 시작하기 전에 blocked된다.

## Non-Software Proof

`ni`는 software spec generator가 아니다. `ni`는 어떤 product surface에
대해서도 project intent를 컴파일한다.

[Neighborhood Cooling Study Protocol](examples/research-protocol/)을 실행해보라:

```bash
go run ./cmd/ni status --dir examples/research-protocol
go run ./cmd/ni run --dir examples/research-protocol --target human-team --out examples/research-protocol/generated/human-team.prompt.md
```

이 locked example은 app이 아니라 research protocol을 계획한다. 여기에는
`product_type: research_protocol`, `document` delivery surface, protocol review
evaluations, human-team handoff prompt가 있다. 데이터를 수집하거나, analysis를
실행하거나, sensors를 배포하거나, fieldwork를 실행하지 않는다.

## Core Flow

planning workspace를 만든다:

```bash
go run ./cmd/ni init --dir <path> --profile prototype
```

sustained model-user conversation을 사용해 `docs/plan/**`과
`.ni/contract.json`을 함께 유지한다. Skills와 models는 UX다. CLI가 authority다.

```bash
go run ./cmd/ni status --dir <path>
go run ./cmd/ni end --dir <path>
go run ./cmd/ni run --dir <path> --target codex --max-chars 4000
```

`ni run`은 prompt를 출력하거나 쓴다. 그 prompt를 실행하지 않는다.

## What ni Blocks

`ni`는 intent가 아직 trust할 수 있는 상태가 아닐 때 downstream handoff를
block한다:

- linked evaluations가 없는 accepted capabilities,
- mitigation이 없는 high-severity risks,
- open blocker questions,
- conflicting accepted decisions,
- 누락되었거나 invalid한 required planning records,
- current files가 `.ni/plan.lock.json`과 더 이상 일치하지 않는 stale locks,
- valid lock이 존재하기 전의 target prompt compilation.

[Benchmark Protocol](docs/43_BENCHMARK_PROTOCOL.md)은 downstream agents를
실행하지 않고 direct-to-agent prompts와 locked `ni` intent를 비교하는 방법을
설명한다.

## Commands Summary

core path는 다음이다:

```text
ni init -> ni status -> ni end -> ni run
```

다른 implemented kernel commands는 targets를 inspect하고, locked seed
material을 export하고, inert feedback과 pressure를 기록하고, explicit
amendments를 관리하고, planning states를 비교하고, inert graph 또는 harness
material을 제안한다.

전체 내용은 [command reference](docs/commands.ko.md)를 참고하라.

## Targets Summary

Targets는 locked plan을 소비하는 shapes다. `ni`가 실행하는 integrations도,
`ni`가 소유하는 runtime adapters도, `ni-kernel`의 일부가 되는 lifecycle state도
아니다.

Built-in targets는 다음을 포함한다:

- `generic`, `codex`, `human-team` prompt 또는 handoff targets,
- `hyper-run`, `namba-ai`, `ouroboros`, `spec-kit` seed targets.

target별 boundary는 [Target Story](docs/45_TARGET_STORY.md)를 참고하라.

## What ni Is Not

`ni`는 다음이 아니다:

- task runner,
- SPEC runner,
- multi-agent execution layer,
- Codex adapter,
- queue,
- shell adapter,
- release automation,
- PR automation,
- Hyper Run, Spec Kit, Ouroboros, 또는 namba-ai.

Downstream prompts, seed packages, harness proposals는 derived and mutable이다.
이것들은 kernel-owned execution state가 되지 않는다.

## Examples and Docs

먼저 읽을 곳:

- [Positioning](docs/40_POSITIONING.md)
- [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md)
- [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/)
- [Neighborhood Cooling Study Protocol](examples/research-protocol/)
- [Command Reference](docs/commands.md)
- [Korean Command Reference](docs/commands.ko.md)
- [Benchmark Protocol](docs/43_BENCHMARK_PROTOCOL.md)
- [Target Story](docs/45_TARGET_STORY.md)
- [v0.2.0 Release Draft](docs/47_RELEASE_DRAFT_v0.2.0.ko.md)

## Development and Release Status

`ni`는 현재 source-first다.
Release status: package distribution이나 published binary release를 claim하지 않는다.
Package publishing, Homebrew taps, GoReleaser, automated release tooling은 현재
kernel scope 밖이다.

source에서는 `go run`을 사용하고, `make build`로 local binary를 build하거나
`make install-local`로 local install mode를 사용할 수 있다. source, local build, local install mode
상세 정보는 [docs/22_INSTALL.md](docs/22_INSTALL.md)를 참고하라.

Public demo verification:

```bash
bash scripts/demo-check.sh
```

Repository validation:

```bash
bash scripts/quality.sh
```

CI validation은 `.github/workflows/ci.yml`에 정의되어 있으며 Go tests, quality
checks, smoke tests를 실행한다.

Source/build/install verification:

```bash
bash scripts/install-check.sh
```

`ni`는 [MIT License](LICENSE)에 따라 licensed된다. Contribution guidelines는
[CONTRIBUTING.ko.md](CONTRIBUTING.ko.md)에 있다. Release readiness notes는
[docs/46_RELEASE_READINESS.ko.md](docs/46_RELEASE_READINESS.ko.md)에 있고,
project security policy는 [SECURITY.md](SECURITY.md)에 있다.
