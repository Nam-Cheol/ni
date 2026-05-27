# Command Reference

이 문서는 `ni`의 detailed source-first command reference다.

`ni`는 AI Agents를 위한 Project Intent Compiler다. Commands는 planning
contract를 만들고 검증하며, accepted intent를 lock하고, bounded downstream
prompts 또는 inert seed material을 컴파일한다. Downstream agents, shell
commands, queues, adapters, PR automation, release automation을 실행하지 않는다.

## Boundary

kernel이 authority를 가지는 것은 다음이다:

- `docs/plan/**`
- `.ni/contract.json`
- deterministic readiness validation
- `.ni/plan.lock.json`
- lock hash verification
- bounded prompt compilation
- inert downstream seed exports and proposals

`.ni/plan.lock.json`이 존재한 뒤 source-of-truth precedence는 다음과 같다:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

locked hash가 더 이상 일치하지 않으면 target handoff commands는 `BLOCKED`로
멈춘다.

## Source Usage

source에서 CLI를 실행한다:

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
상세 정보는 [docs/22_INSTALL.md](22_INSTALL.md)를 참고하라.

## Command Reference Table

이 표의 어떤 command도 downstream work를 실행하지 않는다. Codex, shell
commands, agents, adapters, queues, PR automation, release automation, target
runtimes, work graphs, harness evidence를 실행하지 않는다.

| Command | Purpose | Example | Authority boundary | Lock behavior | Mutates kernel state? | Requires valid lock? | Can return `BLOCKED`? |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `ni --help` | implemented command usage를 출력한다. | `ni --help` | Informational CLI surface only. | lock을 읽거나 만들거나 검증하지 않는다. | No. | No. | No. |
| `ni version` | source version을 출력한다. | `ni version` | Informational CLI surface only. | lock을 읽거나 만들거나 검증하지 않는다. | No. | No. | No. |
| `ni init` | `docs/plan/**`, `.ni/contract.json`, readiness config, bounded session state가 있는 planning workspace를 만든다. | `ni init --dir ./my-plan --profile prototype` | kernel planning state를 시작한다. interactive contract editor가 아니며 contract `add`, `list`, `set` authoring commands를 추가하지 않는다. | `.ni/plan.lock.json`을 요구하거나 만들지 않는다. | Yes, kernel planning files를 만들거나 보존한다. | No. | No. |
| `ni status` | docs, contract, sync, decisions, risks, evaluations, blocker questions에서 deterministic readiness를 평가한다. | `ni status --dir .` | authoritative readiness gate다. model judgment가 override할 수 없다. | lock hashes를 만들거나 검증하지 않는다. | No. | No. | Yes, readiness result로 반환될 수 있다. |
| `ni status --proof` | 같은 readiness result와 rule-level proof evidence를 출력한다. | `ni status --dir . --proof` | planning conversations와 release checks를 위한 authoritative readiness proof다. | lock hashes를 만들거나 검증하지 않는다. | No. | No. | Yes, readiness result로 반환될 수 있다. |
| `ni status --next-questions` | readiness failures에서 focused planning questions를 도출한다. | `ni status --dir . --next-questions` | model-user planning을 위한 interview aid다. questions는 그 자체로 readiness를 바꾸지 않는다. | lock hashes를 만들거나 검증하지 않는다. | No. | No. | Yes, readiness result로 반환될 수 있다. |
| `ni end` | ready accepted plan을 lock한다. | `ni end --dir .` | first lock에 대한 CLI authority다. readiness를 실행하고 model-only readiness claims를 거부한다. | readiness가 `BLOCKED`가 아닐 때만 `.ni/plan.lock.json`을 쓴다. `.ni/contract.json`과 required `docs/plan/**`를 hash한다. | Yes, lockfile을 쓴다. | No existing lock required. | Yes, readiness가 `BLOCKED`일 때. |
| `ni run` | locked plan에서 bounded downstream prompt를 컴파일한다. | `ni run --dir . --target codex --max-chars 4000` | prompt compilation only. output은 handoff text이지 execution이 아니다. | `.ni/plan.lock.json`을 검증하고 stale hashes를 거부한다. | No kernel mutation. `--out`은 prompt artifact만 쓴다. | Yes. | Yes, lock hash mismatch 때. |
| `ni targets` | supported prompt/export targets를 나열한다. | `ni targets --json` | target registry only. targets는 consumption shapes이지 runtime adapters가 아니다. | lock을 읽거나 만들거나 검증하지 않는다. | No. | No. | No. |
| `ni export` | locked plan에서 downstream seed Markdown을 쓴다. | `ni export --dir . --target hyper-run --out ./seed` | seed export only. exported files는 derived and mutable downstream artifacts다. | `.ni/plan.lock.json`을 검증하고 stale hashes를 거부한다. | No kernel mutation. `--out`에 seed Markdown을 쓴다. | Yes. | Yes, lock hash mismatch 때. |
| `ni feedback` | inert downstream observations를 기록하거나 나열한다. | `ni feedback add --dir . --file ./feedback.json` | future planning cycles를 위한 evidence다. contract나 lock을 바꾸지 않는다. | lock이 있으면 feedback read/write 전에 검증한다. | `add`는 `.ni/feedback.jsonl`과 observed pressure를 mutate한다. `list`는 mutate하지 않는다. | Only when a lock exists. | Yes, existing stale lock 때. |
| `ni pressure` | observed planning pressure와 explicit promotion state를 추적한다. | `ni pressure status --dir .` | pressure는 human planning decision이 docs와 contract를 바꾸기 전까지 advisory다. | lock이 있으면 pressure read/write 전에 검증한다. | `promote`와 `retire`는 `.ni/pressure.json`을 mutate한다. `status`는 mutate하지 않는다. | Only when a lock exists. | Yes, existing stale lock 때. |
| `ni amend` | locked-plan changes를 위한 explicit amendment records를 만들고 조회하고 apply한다. | `ni amend create --dir . --title "Clarify acceptance criteria"` | amendment records는 intent changes를 설명한다. contract, docs, lock을 직접 edit하지 않는다. | `create`는 lock이 있으면 current lock hash를 기록한다. `apply`는 다른 source lock에 묶인 amendment를 거부한다. | `create`와 `apply`는 `.ni/amendments/**`를 mutate한다. `list`와 `show`는 mutate하지 않는다. | No. | Yes, 다른 source lock에 대해 amendment를 apply할 때. |
| `ni relock` | explicitly amended plan 뒤에 새 lock을 만든다. | `ni relock --dir .` | relocking에 대한 CLI authority다. amendment gate와 readiness gate를 보존한다. | existing lock이 필요하다. previous lock을 archive한다. current lock에 대한 applied amendment가 없으면 stale docs를 거부하고 blocked readiness도 거부한다. | Yes, previous lock을 archive하고 새 `.ni/plan.lock.json`을 쓴다. | Existing lock required. current hashes는 applied amendment가 있을 때만 stale일 수 있다. | Yes, blocked readiness 또는 applied amendment 없는 stale lock 때. |
| `ni diff` | 두 planning states 사이의 contract-level changes를 보여준다. | `ni diff --base ./base --head ./head --json` | informational collaboration check only. changes를 resolve하거나 apply하지 않는다. | inputs를 orient하기 위해 lockfiles를 읽을 수 있지만 valid locks를 요구하거나 lock gates를 enforce하지 않는다. | No. | No. | No. |
| `ni conflicts` | 두 planning states 사이의 semantic planning conflicts를 보고한다. | `ni conflicts --base ./base --head ./head` | collaboration guardrail only. conflicts를 보고하지만 merge, resolve, mutate하지 않는다. | lockfiles가 있으면 읽고 lock hash mismatches를 semantic conflicts로 보고한다. | No. | No. | Yes, blocking conflicts가 있을 때 conflict severity와 nonzero exit로. |
| `ni graph` | contract capabilities와 artifacts에서 read-only work graph proposal을 출력한다. | `ni graph --dir . --json` | inert proposal material only. task runner, queue, scheduler, execution graph가 아니다. | lock이 있으면 graph proposal 전에 검증한다. | No. | Only when a lock exists. | Yes, existing stale lock 때. |
| `ni harness` | generated-harness proposal records를 출력하거나 관리한다. | `ni harness plan --dir .` | inert proposal material only. evidence runner, adapter, queue, kernel-owned execution state가 아니다. | `plan`, `candidates`, `propose`, `validate`, `accept`, `retire`는 valid lock을 요구하고 검증한다. | `propose`, `validate`, `accept`, `retire`는 `.ni/harness.candidates.json`을 mutate한다. `plan`과 `candidates`는 mutate하지 않는다. | Yes. | Yes, lock hash mismatch 때. |

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

status가 `BLOCKED`면 execution은 시작하면 안 된다.

readiness가 통과한 뒤에만 lock한다:

```bash
go run ./cmd/ni end --dir <path>
```

valid lock에서 bounded downstream prompt를 컴파일한다:

```bash
go run ./cmd/ni run --dir <path> --target codex --max-chars 4000
```

`ni run`은 prompt를 출력하거나 쓴다. 그 prompt를 실행하지 않는다.

## Help and Version

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
```

`ni --help`는 implemented commands와 options를 나열한다. `ni version`은 source
version을 출력한다.

## init

`ni init`은 planning docs와 `.ni` skeleton을 만든다. contract editing session을
시작하지 않는다. 의도된 authoring flow는 workspace creation 이후의
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

`--interaction-mode`는 `human_to_system` 또는 `human_to_human` 같은 lowercase
identifier를 받는다.

이 fields는 planning과 status output을 안내한다. runtime stages나 execution
behavior를 만들지 않는다.

## status

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

## end

`ni end`는 ready plan을 lock한다.

```bash
go run ./cmd/ni end --dir <path>
```

readiness gate를 실행하고, `BLOCKED`를 거부하며, `.ni/contract.json`과 required
`docs/plan/**` files에 대한 hashes를 포함해 `.ni/plan.lock.json`을 쓴다.
`.ni/session.json`은 locked docs보다 아래에 있는 mutable planning aid이므로
hash되지 않는다.

## run

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

## targets

Targets는 locked plan을 소비하는 shapes다. `ni`가 실행하는 integrations도,
`ni`가 소유하는 runtime adapters도, `ni-kernel`의 일부가 되는 lifecycle state도
아니다.

target별 boundary는 [Target Story](45_TARGET_STORY.md)를 참고하라.

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

## export

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

## feedback

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

## pressure

`ni pressure`는 readiness rules를 그 자체로 바꾸지 않으면서 recurring planning
pressure를 추적한다.

```bash
go run ./cmd/ni pressure status --dir <path>
go run ./cmd/ni pressure status --dir <path> --json
go run ./cmd/ni pressure promote P-001 --dir <path>
go run ./cmd/ni pressure retire P-001 --dir <path>
```

Promotion은 explicit하고 staged된다:

```text
observed -> repeated -> promotable -> accepted
```

Accepted pressure도 locked contract를 바꾸기 전에 human planning decision을
필요로 한다.

## amend and relock

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

## diff and conflicts

`ni diff`와 `ni conflicts`는 planning states를 resolve하거나 mutate하지 않고
비교한다.

```bash
go run ./cmd/ni diff --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni diff --base <path-or-lock> --head <path-or-lock> --json
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock> --json
```

Inputs는 project directory, `.ni/contract.json`, `.ni/plan.lock.json` 중 하나일
수 있다. `ni conflicts`는 stale locks, conflicting decisions, weakened accepted
requirements, mitigation context 없는 risk severity reductions를 포함한 blocking
semantic conflicts를 발견하면 nonzero로 종료한다.

## graph and harness

`ni graph`와 `ni harness`는 optional downstream work를 inert seed/proposal
material로 설명한다. `ni graph`는 draft contract state에서도 proposal을 만들 수
있고 lock이 있으면 hashes를 검증한다. `ni harness`는 proposal과 candidate
lifecycle commands에 valid lock을 요구한다. command names와 달리, 이 outputs는
task runner, evidence runner, queue, adapter, kernel-owned execution state가
아니다.

```bash
go run ./cmd/ni graph --dir <path>
go run ./cmd/ni graph --dir <path> --json
go run ./cmd/ni harness plan --dir <path>
go run ./cmd/ni harness plan --dir <path> --json
go run ./cmd/ni harness candidates --dir <path>
go run ./cmd/ni harness candidates --dir <path> --json
go run ./cmd/ni harness propose --dir <path> --from-pressure P-001
go run ./cmd/ni harness validate CAND-001 --dir <path> --evidence <path>
go run ./cmd/ni harness accept CAND-001 --dir <path>
go run ./cmd/ni harness retire CAND-001 --dir <path>
```

kernel은 valid lock에서 work graphs, evaluation-plan proposals, evidence-rule
notes, downstream handoff material을 컴파일할 수 있다. 그것들을 실행해서는 안
된다.

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

이 repository의 main quality entry point는 다음이다:

```bash
bash scripts/quality.sh
```

`scripts/quality.sh`는 formatting checks, Go tests, JSON checks, Markdown fence
checks, skill metadata checks, prompt budget checks, core-boundary self-tests,
smoke tests를 실행한다.

Public demo verification은 별도의 release proof check다:

```bash
bash scripts/demo-check.sh
```

README demos, example workspaces, status output, prompt compilation behavior를
변경할 때 실행한다. 자세한 내용은 [docs/48_DEMO_VERIFICATION.md](48_DEMO_VERIFICATION.md)에
있다.
