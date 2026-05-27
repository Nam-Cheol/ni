<p align="center">
  <img src="assets/hero.svg" alt="ni hero banner: Project Intent Compiler for AI Agents" width="100%">
</p>

<p align="center">
  <a href="README.md" aria-label="Read in English"><img alt="English" src="assets/badge-english.svg" width="84" height="28"></a>
  <a href="README.ko.md" aria-label="한국어로 읽기"><img alt="한국어" src="assets/badge-korean.svg" width="84" height="28"></a>
</p>

<p align="center">
  <a href="LICENSE"><img alt="License MIT" src="https://img.shields.io/badge/license-MIT-f4b860"></a>
  <a href=".github/workflows/ci.yml"><img alt="CI workflow exists" src="https://img.shields.io/badge/CI-workflow%20exists-25334a"></a>
  <a href="SECURITY.md"><img alt="Security policy exists" src="https://img.shields.io/badge/security-policy%20exists-2d5a52"></a>
  <a href="docs/00_START_HERE.md"><img alt="Docs index exists" src="https://img.shields.io/badge/docs-index%20exists-5b8def"></a>
</p>

<p align="center"><sub>신뢰 신호: MIT license; CI workflow; security policy; docs index.</sub></p>

<h1 align="center">agent를 아직 실행하지 마세요. 먼저 의도를 컴파일하세요.</h1>

<p align="center"><strong>ni는 implementation work가 시작되기 전에 planning conversation을 locked project contract로 바꿉니다.</strong></p>

`ni`는 AI Agents를 위한 Project Intent Compiler입니다. intent를 명시화하고,
plan readiness를 검사하고, accepted contract를 lock한 뒤 bounded handoff prompt
또는 derived seed material을 compile합니다.

## 왜 ni인가

<p align="center">
  <img src="assets/card-pain-vague-intent.svg" alt="Vague intent: plausible prompt 안에 users, acceptance criteria, risks, non-goals, blockers가 빠져 있을 수 있습니다." width="32%">
  <img src="assets/card-pain-early-execution.svg" alt="Early execution: request가 plausible하다는 이유만으로 work를 시작하면 안 됩니다." width="32%">
  <img src="assets/card-pain-rework.svg" alt="Rework: hidden assumptions는 people, models, tools가 wrong plan에서 시작한 뒤 더 비싸집니다." width="32%">
</p>

### 모호한 intent

Prompt가 바로 실행 가능해 보여도 users, acceptance criteria, risks, non-goals,
blocker questions가 빠져 있을 수 있습니다.

### 너무 이른 실행

request가 그럴듯하게 들린다는 이유만으로 work를 시작하면 안 됩니다.

### 재작업

Hidden assumptions는 people, models, tools가 wrong plan에서 시작한 뒤 더
비싸집니다.

## ni가 주는 것

<p align="center">
  <img src="assets/card-payoff-capture-intent.svg" alt="Capture intent: planning conversation은 explicit docs와 contract draft가 됩니다." width="32%">
  <img src="assets/card-payoff-lock-contract.svg" alt="Lock contract: readiness and lock commands가 accepted plan, hashes, source of truth를 gate합니다." width="32%">
  <img src="assets/card-payoff-handoff-safely.svg" alt="Handoff safely: valid locked plan이 bounded prompt 또는 derived seed material로 compile됩니다." width="32%">
</p>

### intent 포착

Planning conversation은 explicit docs와 contract draft가 됩니다.

### contract 잠금

`ni status`와 `ni end`가 readiness, hashes, lock creation을 gate합니다.

### 안전한 handoff

`ni run`은 valid locked plan에서 bounded prompt 또는 seed를 compile합니다.
Shell commands, queues, agents, downstream work는 실행하지 않습니다.

## 60초 시작

`ni`는 현재 source-first입니다:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
```

conversation으로 `./my-plan/docs/plan/**`과 `./my-plan/.ni/contract.json`을
채운 뒤, authoritative call은 CLI에 맡깁니다:

```bash
go run ./cmd/ni status --dir ./my-plan --next-questions
go run ./cmd/ni end --dir ./my-plan
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

## Choose your path

| Path | Status | Use it when | Boundary |
| --- | --- | --- | --- |
| Source | Available | `go run ./cmd/ni ...`를 실행할 수 있습니다. | Full deterministic `status`, `end`, `run`. |
| Local binary | Available after local build | 이 checkout에서 `./bin/ni` 또는 local install을 원할 때 사용합니다. | Local build/install만 말하며 hosted binary는 주장하지 않습니다. |
| Model workspaces | Available as assisted planning | Codex 또는 Claude가 docs와 contract records draft를 돕게 하고 싶을 때 사용합니다. | Skills are UX; readiness와 lock authority는 CLI입니다. |
| No-terminal method | Available as assisted method | CLI run 전 Intent Lock method를 배우거나 draft하고 싶을 때 사용합니다. | Useful drafting이며 deterministic validation은 아닙니다. |
| Release binary | Planned | Published release에서 Go 없이 `ni`를 받고 싶을 때 사용합니다. | Public release assets와 checksums를 기다립니다. |
| Curl installer | Release-gated | Release assets가 생긴 뒤 one-command install을 원할 때 사용합니다. | `install.sh`는 있지만 public install은 release assets를 기다립니다. |
| Homebrew | Planned | Package manager를 선호할 때 사용합니다. | Published tap이나 formula가 없습니다. |

자세한 내용은 [Install ni](docs/22_INSTALL.md),
[터미널 없이 계획하기](docs/no-terminal.ko.md),
[Model Workspace Packs](docs/55_MODEL_WORKSPACE_PACKS.md)를 참고하세요.

## Demo

가장 좋은 첫 demo는 blocked demo입니다:

```bash
go run ./cmd/ni status --dir examples/ambiguous-prompt-blocked/workspace
```

```text
BLOCKED
```

이 결과가 핵심입니다. vague request는 handoff 전에 멈춰야 합니다.
[Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) walkthrough를
참고하세요.

## ni가 아닌 것

`ni`는 task runner, spec runner, multi-agent execution layer, queue, shell
adapter, PR automation system, release automation system, downstream work
runtime이 아닙니다. kernel은 planning contracts, readiness, lockfiles, hash
checks, prompt compilation을 소유합니다.

## 다음에 읽을 것

| Read | Why |
| --- | --- |
| [Why ni exists](docs/product-story.ko.md) | Compile-before-run 뒤의 짧은 product story. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | Readiness, locking, hash trust, blocked handoff의 깊은 규칙. |
| [Install ni](docs/22_INSTALL.md) | Source, local build, release-gated install details. |
| [Command reference](docs/commands.ko.md) | Implemented CLI surface. |
| [README Visual Wireframe](docs/63_README_VISUAL_WIREFRAME.ko.md) | 이 README의 visual layout contract. |
