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

이 checkout에서 평가하거나 개발할 때는 source로 시작하세요:

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

## 5분 첫 project

`ni`가 무엇을 구현하게 하지 않고 안전하게 시험하려면 이 flow를 사용하세요:

1. Planning workspace를 만듭니다.

```bash
go run ./cmd/ni init --dir ./my-plan --profile prototype
```

2. 실행 전에 model-user planning conversation을 진행합니다. Planning workflow는
   purpose, actors, requirements, risks, evaluations, non-goals, decisions,
   artifacts, open questions를 `docs/plan/**`, `.ni/contract.json`,
   `.ni/session.json`에 기록하거나 업데이트합니다.

3. CLI에 authoritative readiness proof를 요청합니다.

```bash
go run ./cmd/ni status --dir ./my-plan --proof --next-questions
```

4. `BLOCKED` questions나 gaps가 있으면 planning conversation에서 해결합니다.
   Model은 update를 draft할 수 있지만 readiness는 `ni status`가 결정합니다.

5. CLI가 plan ready를 보고한 뒤에만 lock합니다.

```bash
go run ./cmd/ni end --dir ./my-plan
```

6. Bounded downstream handoff prompt를 compile합니다.

```bash
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

`ni run`은 `.ni/plan.lock.json`에서 prompt를 compile합니다. Prompt를 실행하거나,
agents, shell commands, downstream work를 실행하지 않으며 product readiness를
증명하지 않습니다.

## Choose your path

| Path | Status | Start with | Use it when |
| --- | --- | --- | --- |
| Source | Available | `go run ./cmd/ni --help` | Go가 있고 development 또는 evaluation을 가장 투명하게 시작하고 싶을 때. |
| Local binary | Available | `make build && ./bin/ni --help` | 이 checkout에서 `./bin/ni` 또는 local install을 원할 때. |
| Release binary | Available | [v0.4.0 release](https://github.com/Nam-Cheol/ni/releases/tag/v0.4.0) | Go 없이 설치하고 checksum을 직접 검증하고 싶을 때. |
| Curl installer | Available | `sh install.sh --dry-run --version 0.4.0` | Script를 먼저 inspect한 뒤 작은 shell installer를 쓰고 싶을 때. |
| Model workspaces | Experimental | [Model Workspace Status](docs/99_MODEL_WORKSPACE_STATUS.ko.md) | Supported model workspace 안에서 `ni-start`, `ni-grill`, `ni-end`, `ni-run` guidance를 사용한다. Skills are UX; the CLI is authority. Host-level/global install은 documented되기 전까지 unverified이다. |
| No-terminal method | Experimental | [터미널 없이 계획하기](docs/no-terminal.ko.md) | Trusted runner가 CLI proof를 만들기 전 docs와 contract assisted drafting을 하고 싶을 때; model judgment는 lock이 아닙니다. |
| Homebrew | Planned | [Homebrew Decision](docs/80_HOMEBREW_DECISION.ko.md) | Package manager를 선호할 때; implementation은 v0.5로 defer되었고 published 또는 tested tap/formula는 없습니다. |

### Which path should I choose?

Go가 있으면 Source를 고르세요. 이 checkout에서 반복 가능한 binary를 원하면
Local binary가 맞습니다. Go 없이 직접 checksum을 검증하고 싶으면 Release
binary를, shell script를 먼저 inspect할 수 있으면 Curl installer를 고르세요.
Model-assisted planning에는 Model workspaces를 사용하고, No-terminal method는
CLI proof가 생기기 전 assisted drafting 용도로만 사용하세요. Package-manager
install이 필수라면 Homebrew를 기다려야 합니다.

Minimal curl installer check:

```bash
VERSION="0.4.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
```

Source, local binary, release binary, curl installer의 전체 절차는
[Install ni](docs/22_INSTALL.md)를 참고하세요. Manual release path는 같은
v0.4.0 release에서 matching archive와 `ni_0.4.0_checksums.txt`를 download하고,
archive checksum을 verify하고, 압축을 해제한 뒤 `ni --help`와 `ni version`을
실행하는 것이다.

Release status: v0.4.0 release binaries는 asset과 checksum 검증 후 Available입니다.
Curl installer는 실제 v0.4.0 release assets에 대해 검증된 뒤 Available입니다.
Homebrew를 포함한 package-manager distribution은 아직 Available이 아닙니다.

## macOS install / uninstall

권장 verified path는 v0.4.0 curl installer를 inspect한 뒤 사용하는 것입니다.
`install.sh`는 기본적으로 `ni` binary만 `$HOME/.local/bin/ni`에 설치합니다.
다른 directory를 쓰려면 `BINDIR`를 지정합니다.

```bash
VERSION="0.4.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --version "$VERSION"
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

`ni`를 command name으로 찾지 못하면 선택한 `BINDIR`를 `PATH`에 추가하세요.

Installer로 설치된 binary는 정확한 설치 파일을 삭제해 uninstall합니다:

```bash
rm -f "$HOME/.local/bin/ni"
```

다른 `BINDIR`에 설치했다면 그 directory의 `ni`를 삭제하세요. `ni`만을 위해
추가한 `PATH` line이 있다면 shell profile에서 직접 제거하세요. Homebrew:
Planned / v0.5 candidate 상태이므로 아직 `brew install`을 사용하지 않습니다.

## Windows install / uninstall

Verified public Windows package-manager installer는 아직 문서화되어 있지 않습니다.
현재 문서화된 Windows path는 `windows/amd64`용 v0.4.0 release binary archive입니다.
MSI, winget, Chocolatey, Scoop, Homebrew path는 claim하지 않습니다.

```powershell
$Version = "0.4.0"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_windows_amd64.zip" -OutFile "ni_$($Version)_windows_amd64.zip"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_checksums.txt" -OutFile "ni_$($Version)_checksums.txt"
Get-FileHash "ni_$($Version)_windows_amd64.zip" -Algorithm SHA256
Select-String "ni_$($Version)_windows_amd64.zip" "ni_$($Version)_checksums.txt"
Expand-Archive "ni_$($Version)_windows_amd64.zip" -DestinationPath "ni_$($Version)_windows_amd64"
.\ni_$($Version)_windows_amd64\ni.exe --help
.\ni_$($Version)_windows_amd64\ni.exe version
```

Extracted binary를 trust하기 전에 hash output과 checksum line을 비교하세요.
반복 사용하려면 직접 관리하는 directory에 `ni.exe`를 두고 그 directory를 user
`PATH`에 추가합니다.

Uninstall은 복사해 둔 executable을 그 directory에서 삭제하는 것입니다:

```powershell
Remove-Item "$HOME\bin\ni.exe"
```

실제로 `ni.exe`를 둔 path를 사용하세요. User `PATH` entry는 `ni`만을 위해
추가한 경우에만 제거하세요.

License: `ni`는 [MIT License](LICENSE)로 배포됩니다.

자세한 내용은 [Install ni](docs/22_INSTALL.md),
[터미널 없이 계획하기](docs/no-terminal.ko.md),
[Model Workspace Status](docs/99_MODEL_WORKSPACE_STATUS.ko.md),
[Model Workspace Packs](docs/55_MODEL_WORKSPACE_PACKS.md),
[Model Pack Install Verification](docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md)를
참고하세요.

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
| [Install ni](docs/22_INSTALL.md) | Source, local build, release binary, curl installer details. |
| [Benchmark Claim Boundaries](docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md) | Benchmark `READY`, `not_measured`, 4000-character prompt evidence가 무엇을 증명하고 무엇을 증명하지 않는지. |
| [Homebrew Decision](docs/80_HOMEBREW_DECISION.ko.md) | Homebrew는 Planned로 유지하며 tap implementation은 v0.5로 defer. |
| [Homebrew Tap Plan](docs/72_HOMEBREW_TAP_PLAN.ko.md) | Planned Homebrew route; package-manager availability claim 없음. |
| [Command reference](docs/commands.ko.md) | Implemented CLI surface. |
| [README Visual Wireframe](docs/63_README_VISUAL_WIREFRAME.ko.md) | 이 README의 visual layout contract. |
