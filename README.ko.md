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

<h1 align="center">agent를 아직 실행하지 마세요. 먼저 의도를 컴파일하세요.</h1>

<p align="center"><strong>ni는 implementation work가 시작되기 전에 planning conversation을 locked project contract로 바꿉니다.</strong></p>

`ni`는 AI Agents를 위한 Project Intent Compiler입니다. intent를 명시화하고,
plan readiness를 검사하고, accepted contract를 lock한 뒤 bounded handoff prompt
또는 derived seed material을 compile합니다.

<p align="center">
  <img src="assets/intent-lock-flow.svg" alt="Intent Lock Protocol flow: conversation, project contract, readiness gate, lock hash, bounded handoff." width="100%">
</p>

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

## Install

README는 첫 성공을 위한 두 가지 primary path만 보여줍니다. Source, local build,
Linux, release archive, advanced uninstall details는
[Install ni](docs/22_INSTALL.md)에 있습니다.

### macOS

Verified v0.5.0 release binary를 curl installer로 설치합니다. 설치 전 script를
먼저 inspect합니다:

```bash
VERSION="0.5.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

새 shell을 연 뒤 global command를 verify하고 project를 시작합니다:

```bash
ni --help
ni version
mkdir my-project
cd my-project
ni init .
```

Installer로 설치한 binary와, 추가했다면 ni-managed PATH block을 uninstall합니다:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

Homebrew: Planned / v0.5 candidate. Planned package-manager work는 docs를
참고하세요.

### Windows

PowerShell installer는 기본적으로 `%LOCALAPPDATA%\ni\bin`에 설치하고 User PATH만
업데이트합니다:

```powershell
$Version = "0.5.0"
Invoke-WebRequest "https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1" -OutFile "install.ps1"
Get-Content .\install.ps1
.\install.ps1 -DryRun -Version $Version
.\install.ps1 -Version $Version
```

새 PowerShell session을 연 뒤 global command를 verify하고 project를 시작합니다:

```powershell
ni --help
ni version
mkdir my-project
cd my-project
ni init .
```

Installer로 설치한 binary와 `ni`가 추가한 User PATH entry를 uninstall합니다:

```powershell
.\install.ps1 -Uninstall
```

Windows installer code와 static safety checks는 있습니다. 실제 Windows host
execution은 macOS-only development host에서는 deferred 상태이며 Windows install
transcript가 생기기 전까지 verified라고 claim하지 않습니다.

## 5분 첫 project

`ni`가 무엇을 구현하게 하지 않고 안전하게 시험하려면 이 flow를 사용하세요:

```bash
mkdir my-project
cd my-project
ni init .
ni status --proof --next-questions
ni end
ni run --max-chars 4000
```

`ni init .`은 current directory에서 guided project intent setup을 시작합니다.
Agents나 downstream work를 실행하지 않고 `.ni/contract.json`, `.ni/session.json`,
`docs/plan/**`을 초기화합니다.

`ni status --proof --next-questions`는 authoritative readiness check입니다.
`BLOCKED` questions나 gaps가 있으면 planning conversation에서 해결합니다. Model은
update를 draft할 수 있지만 readiness는 `ni status`가 결정합니다.

`ni end`는 CLI가 ready를 보고한 뒤 accepted plan을 lock하고
`.ni/plan.lock.json`을 씁니다.

`ni run --max-chars 4000`은 fresh lock에서 bounded downstream handoff prompt를
compile합니다. Prompt, agents, shell commands, downstream work를 실행하지 않고
product readiness를 증명하지 않습니다.

자세한 내용은 [Install ni](docs/22_INSTALL.md),
[터미널 없이 계획하기](docs/no-terminal.ko.md),
[Model Workspace Status](docs/99_MODEL_WORKSPACE_STATUS.ko.md),
[Model Workspace Packs](docs/55_MODEL_WORKSPACE_PACKS.md),
[Model Pack Install Verification](docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md)를
참고하세요.

Model workspace packs: Experimental. Host-level/global install은 documented되기 전까지 unverified입니다. No-terminal method: Experimental / assisted. Skills are UX; CLI is authority.

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

License: `ni`는 [MIT License](LICENSE)로 배포됩니다.

## 다음에 읽을 것

| Read | Why |
| --- | --- |
| [Why ni exists](docs/product-story.ko.md) | Compile-before-run 뒤의 짧은 product story. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | Readiness, locking, hash trust, blocked handoff의 깊은 규칙. |
| [Install ni](docs/22_INSTALL.md) | Source, local build, release binary, curl installer details. |
| [Benchmark Claim Boundaries](docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md) | Benchmark `READY`, `not_measured`, 4000-character prompt evidence가 무엇을 증명하고 무엇을 증명하지 않는지. |
| [Homebrew Decision](docs/80_HOMEBREW_DECISION.ko.md) | Homebrew는 Planned로 유지하며 tap implementation은 v0.5로 defer. |
| [Homebrew Tap Plan](docs/72_HOMEBREW_TAP_PLAN.ko.md) | Planned Homebrew route; package-manager availability claim 없음. |
| [Command reference](docs/commands.ko.md) | Implemented CLI surface. |
| [README Visual Wireframe](docs/63_README_VISUAL_WIREFRAME.ko.md) | 이 README의 visual layout contract. |
