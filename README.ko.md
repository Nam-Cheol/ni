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

먼저 category docs를 읽어보라:

- [Positioning](docs/40_POSITIONING.md)
- [Differentiation map](docs/41_DIFFERENTIATION.md)
- [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md)

## What Problem ni Solves

Agents는 실행 가능해 보이지만 중요한 intent를 숨긴 prompt에서 시작하라는
요청을 자주 받는다:

- 이 project는 누구를 위한 것인가?
- work가 accepted되려면 무엇이 참이어야 하는가?
- 어떤 risks에 mitigation이 필요한가?
- 무엇이 명시적으로 scope 밖인가?
- 어떤 questions가 execution을 막아야 하는가?
- accepted 이후 plan이 바뀌었는가?

대부분의 tools는 prompt, spec, worklist, runtime loop가 이미 존재한 뒤에
agent를 제어하려고 한다. `ni`는 control을 더 앞쪽으로 옮긴다. `ni`는
downstream actor가 work를 시작하기 전에 project intent가 explicit,
accepted, validated, locked, unchanged 상태인지 묻는다.

## The ni Answer: Intent Lock Protocol

[Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md)은 `ni-kernel`의 핵심
mechanism이다. 이 protocol은 다음을 정의한다:

1. planning conversation이 project contract가 되는 방식,
2. contract가 lock될 준비가 되는 시점,
3. accepted plan이 hash되는 방식,
4. downstream actors가 trust할 수 있는 것,
5. intent가 바뀌었기 때문에 execution이 멈춰야 하는 시점.

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

locked hash가 더 이상 일치하지 않으면 `ni run`, `ni export`, feedback,
pressure, downstream handoff commands는 `BLOCKED`로 멈춘다.

## 5-Minute Demo: Ambiguous Prompt Blocked

`ni`를 가장 빠르게 이해하는 방법은
[ambiguous prompt blocked demo](examples/ambiguous-prompt-blocked/)를 보는
것이다.

이 demo는 모호한 request에서 시작한다:

```text
Build me a dashboard for my team.
```

direct-to-agent 경로라면 agent가 users, data, workflow, non-goals, success
criteria에 관한 hidden assumptions를 만들어내야 한다. `ni` 경로는 request를
planning intent로 기록한 뒤, blocker questions가 열려 있는 동안 이를
executable한 것으로 취급하기를 거부한다.

blocked workspace를 실행해보라:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected result:

```text
BLOCKED
```

핵심은 이것이다. 모호한 execution은 agent가 시작하기 전에 blocked된다.

## Non-Software Demo

`ni`는 software spec generator가 아니다. `ni`는 어떤 product surface에
대해서도 project intent를 컴파일한다.

[Neighborhood Cooling Study Protocol](examples/research-protocol/)을 실행해보라:

```bash
go run ./cmd/ni status --dir examples/research-protocol
go run ./cmd/ni run --dir examples/research-protocol --target human-team --out examples/research-protocol/generated/human-team.prompt.md
```

이 locked example은 app이 아니라 research protocol을 계획한다. 여기에는
`product_type: research_protocol`, `document` delivery surface, protocol
review evaluations, human-team handoff prompt가 있다. 데이터를 수집하거나,
analysis를 실행하거나, sensors를 배포하거나, fieldwork를 실행하지 않는다.

또 다른 non-software example은
[Travel Concierge Triage](examples/conversation-product/)이다. 이것은 chatbot
배포나 travel booking 없이 human concierge handoff를 컴파일하는 conversation
product다.

## Core Flow

planning workspace를 만든다:

```bash
go run ./cmd/ni init --dir <path> --profile prototype
```

sustained model-user conversation을 사용해 `docs/plan/**`과
`.ni/contract.json`을 함께 유지한다. Skills와 models는 UX다. CLI가 authority다.

readiness를 확인한다:

```bash
go run ./cmd/ni status --dir <path>
go run ./cmd/ni status --dir <path> --proof --next-questions
```

model은 status를 설명할 수 있지만 override할 수는 없다. status가 `BLOCKED`면
execution은 시작하면 안 된다.

readiness가 통과한 뒤에만 lock한다:

```bash
go run ./cmd/ni end --dir <path>
```

valid lock에서 bounded downstream prompt를 컴파일한다:

```bash
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

[docs/43_BENCHMARK_PROTOCOL.md](docs/43_BENCHMARK_PROTOCOL.md)의 benchmark
protocol은 downstream agents를 실행하지 않고 direct-to-agent prompts와 locked
`ni` intent를 비교하는 방법을 설명한다.

## Core Commands

### Help and Version

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
```

### init

`ni init`은 planning docs와 `.ni` skeleton을 만든다. contract editing
session을 시작하지 않는다. 의도된 authoring flow는 workspace creation 이후의
model-assisted conversation이다.

```bash
go run ./cmd/ni init --dir <path>
go run ./cmd/ni init --dir <path> --profile concept
go run ./cmd/ni init --dir <path> --product-type conversation_product --surface conversation --interaction-mode human_to_system
```

Supported readiness profiles:

```text
concept
prototype
mvp
beta
production
```

Supported product types:

```text
software
conversation_product
research_protocol
operations_process
education_program
document_product
physical_product
mixed
```

Supported delivery surfaces:

```text
web
cli
api
conversation
document
workflow
human_service
physical
```

이 fields는 planning과 status output을 안내한다. runtime stages나 execution
behavior를 만들지 않는다.

### status

`ni status`는 deterministic rules로 readiness를 평가한다.

```bash
go run ./cmd/ni status --dir <path>
go run ./cmd/ni status --dir <path> --json
go run ./cmd/ni status --dir <path> --proof
go run ./cmd/ni status --dir <path> --proof --json
go run ./cmd/ni status --dir <path> --next-questions
go run ./cmd/ni status --dir <path> --json --next-questions
```

Status values:

```text
BLOCKED
READY_WITH_DEFERRALS
READY
```

`--proof`가 있으면 `ni status`는 readiness, docs/contract sync,
accepted-decision conflict checks에서 나온 rule-level evidence를 출력한다.

`--next-questions`가 있으면 `ni status`는 readiness rule failures에서 concise
planning questions를 도출해 planning conversation이 다음 specific gap을 다룰 수
있게 한다.

### end

`ni end`는 ready plan을 lock한다.

```bash
go run ./cmd/ni end --dir <path>
```

readiness gate를 실행하고, `BLOCKED`를 거부하며, `.ni/contract.json`과 required
`docs/plan/**` files에 대한 hashes를 포함해 `.ni/plan.lock.json`을 쓴다.
`.ni/session.json`은 locked docs보다 아래에 있는 mutable planning aid이므로
hash되지 않는다.

### run

`ni run`은 locked plan에서 prompt를 컴파일한다.

```bash
go run ./cmd/ni run --dir <path>
go run ./cmd/ni run --dir <path> --target codex
go run ./cmd/ni run --dir <path> --target human-team --out <file>
go run ./cmd/ni run --dir <path> --max-chars 2400
```

Prompt output은 configured maximum 안에 머물러야 하며, 초기값은 4000
characters다. `ni run`은 Codex, shell commands, agents, queues, adapters를
실행하지 않는다.

## Targets and Exports

Targets는 locked plan을 소비하는 shapes다. `ni`가 실행하는 integrations도,
`ni`가 소유하는 runtime adapters도, `ni-kernel`의 일부가 되는 lifecycle state도
아니다.

target별 boundary는 [target story](docs/45_TARGET_STORY.md)를 참고하라.

지원되는 prompt/export targets를 나열한다:

```bash
go run ./cmd/ni targets
go run ./cmd/ni targets --json
```

Built-in targets:

```text
generic     prompt   general downstream implementation prompt
codex       prompt   bounded implementation prompt seed
human-team  handoff  planning handoff for people
hyper-run   seed     seed material, not .hyper/goals runtime packets
namba-ai    seed     planning seed and suggested graph boundaries
ouroboros   seed     upstream intent notes, not Agent OS execution state
spec-kit    seed     upstream intent summary, not Spec Kit workflow state
```

`ni export`는 supported downstream targets를 위한 locked-plan seed packages를
쓴다. 이 outputs는 locked plan에서 파생되며 mutable downstream artifacts로
남는다.

```bash
go run ./cmd/ni export --dir <path> --target hyper-run --out <dir>
go run ./cmd/ni export --dir <path> --target namba-ai --out <dir>
go run ./cmd/ni export --dir <path> --target ouroboros --out <dir>
go run ./cmd/ni export --dir <path> --target spec-kit --out <dir>
```

Export는 `.ni/plan.lock.json`을 요구하고, locked hashes를 검증하며, stale
plans를 `BLOCKED`로 거부한다. Seed Markdown만 쓴다. external runtimes를
호출하거나, downstream runtime packets를 만들거나, target adapters를 추가하지
않는다.

## Additional Kernel Commands

이 commands는 같은 boundary를 유지한다. kernel records와 inert proposals를
읽거나 쓰지만 downstream work를 실행하지 않는다.

### feedback

`ni feedback`은 contract나 lock을 mutate하지 않고 downstream observations를
기록한다.

```bash
go run ./cmd/ni feedback add --dir <path> --file testdata/feedback/codex.json
go run ./cmd/ni feedback list --dir <path>
go run ./cmd/ni feedback list --dir <path> --json
```

Feedback은 `.ni/feedback.jsonl`에 append되고 observed pressure items로
translated된다. 이것은 future planning cycle을 위한 evidence이지, automatic
contract change가 아니다.

### pressure

`ni pressure`는 readiness rules를 그 자체로 바꾸지 않으면서 recurring planning
pressure를 추적한다.

```bash
go run ./cmd/ni pressure status --dir <path>
go run ./cmd/ni pressure promote P-001 --dir <path>
go run ./cmd/ni pressure retire P-001 --dir <path>
```

Promotion은 explicit하고 staged된다:

```text
observed -> repeated -> promotable -> accepted
```

Accepted pressure도 locked contract를 바꾸기 전에 human planning decision을
필요로 한다.

### amend and relock

Locked planning docs는 silently edited되면 안 된다. Amendments를 사용해 locked
plan이 왜 바뀌었는지 설명한 다음 relock한다.

```bash
go run ./cmd/ni amend create --dir <path> --title "Clarify acceptance criteria"
go run ./cmd/ni amend list --dir <path>
go run ./cmd/ni amend show AMEND-001 --dir <path>
go run ./cmd/ni amend apply AMEND-001 --dir <path>
go run ./cmd/ni relock --dir <path>
```

Applied amendment에는 reason, affected docs 또는 contract IDs, proposed
changes, risk impact, readiness impact가 포함되어야 한다. `ni relock`은
applied amendment가 없는 stale locks를 거부하고 blocked readiness도 거부한다.

### diff and conflicts

`ni diff`와 `ni conflicts`는 planning states를 resolve하거나 mutate하지 않고
비교한다.

```bash
go run ./cmd/ni diff --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock> --json
```

Inputs는 project directory, `.ni/contract.json`, `.ni/plan.lock.json` 중 하나일
수 있다. `ni conflicts`는 stale locks, conflicting decisions, weakened accepted
requirements, mitigation context 없는 risk severity reductions를 포함한 blocking
semantic conflicts를 발견하면 nonzero로 종료한다.

### graph and harness

`ni graph`와 `ni harness`는 locked contract에서 optional downstream work를
설명한다. command names와 달리, 이 outputs는 inert seed/proposal material이다.
task runner, evidence runner, queue, adapter, kernel-owned execution state가
아니다.

```bash
go run ./cmd/ni graph --dir <path>
go run ./cmd/ni harness plan --dir <path>
go run ./cmd/ni harness candidates --dir <path>
go run ./cmd/ni harness propose --dir <path> --from-pressure P-001
go run ./cmd/ni harness validate CAND-001 --dir <path> --evidence <path>
go run ./cmd/ni harness accept CAND-001 --dir <path>
go run ./cmd/ni harness retire CAND-001 --dir <path>
```

kernel은 valid lock에서 work graphs, evaluation-plan proposals, evidence-rule
notes, downstream handoff material을 컴파일할 수 있다. 그것들을 실행해서는 안
된다.

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

`ni`가 host enhancers, SDD toolkits, coding-agent operating systems, execution
growth runtimes와 어떻게 다른지는
[differentiation map](docs/41_DIFFERENTIATION.md)을 참고하라.

## Examples

Complete example workspaces는 `examples/`에 있다:

- [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/): vague requests를
  위한 core blocking demo.
- [Neighborhood Cooling Study Protocol](examples/research-protocol/):
  human-team과 generic prompt artifacts가 있는 non-software research protocol.
- [Travel Concierge Triage](examples/conversation-product/): human-team과 Codex
  prompt artifacts가 있는 conversation product.
- [Conversation Authoring Fixture](examples/conversation-authoring/): `ni init`
  이후 model-maintained docs와 contract records를 보여주는 end-to-end
  transcript이며, CLI validation, lock, prompt compilation을 포함한다.
- [Namba AI Upgrade](examples/namba-ai-upgrade/): software product planning
  example.
- [Benchmark Report Template](examples/benchmark-report/): pre-runtime intent
  readiness benchmark를 위한 manual report template.

## Development Status

`ni`는 현재 source-first다. Package publishing, Homebrew taps, GoReleaser,
automated release tooling은 현재 kernel scope 밖이다.

source에서 CLI를 직접 실행해볼 때는 `go run`을 사용한다:

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
go run ./cmd/ni status --dir .
```

local binary를 `bin/ni`로 build한다:

```bash
make build
./bin/ni --help
./bin/ni version
```

기본적으로 local binary를 `~/.local/bin/ni`에 install한다:

```bash
make install-local
~/.local/bin/ni version
```

다른 install location을 선택하려면 `PREFIX` 또는 `BINDIR`를 override한다. 설치
상세 정보는 [docs/22_INSTALL.md](docs/22_INSTALL.md)를 참고하라. Korean
companion README는 [README.ko.md](README.ko.md)에서 함께 유지된다.

## JSON Schemas

NI state files를 위한 versioned JSON Schemas는 `schema/`에 있다:

```text
schema/ni.project.v0.json
schema/ni.contract.v0.json
schema/ni.lock.v0.json
schema/ni.readiness-rules.v0.json
schema/ni.readiness-profiles.v0.json
schema/ni.feedback.v0.json
schema/ni.pressure.v0.json
schema/ni.amendment.v0.json
schema/ni.harness-candidate.v0.json
```

published schemas와 current `.ni` state files를 검증한다:

```bash
python3 scripts/check-schema.py
```

## Validation

Go code가 있을 때 실행한다:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

이 repository의 main quality entry point는 다음이다:

```bash
bash scripts/quality.sh
```

`scripts/quality.sh`는 formatting checks, Go tests, JSON checks, Markdown fence
checks, skill metadata checks, prompt budget checks, core-boundary self-tests,
smoke tests를 실행한다.

## License and Release Status

`ni`는 [MIT License](LICENSE)에 따라 licensed된다.

이 README는 package distribution이나 published binary release를 claim하지 않는다.
release process가 달리 말하지 않는 한 source, local build, local install mode를
사용한다.

Release readiness notes는
[docs/46_RELEASE_READINESS.ko.md](docs/46_RELEASE_READINESS.ko.md)에 있다. CI는
[.github/workflows/ci.yml](.github/workflows/ci.yml)에 정의되어 있다. Project
security policy는 `SECURITY.md`가 존재하지 않으므로 아직 publish되지 않았다.
