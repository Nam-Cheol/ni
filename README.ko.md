<p align="center">
  <img src="assets/hero.svg" alt="ni hero banner: Project Intent Compiler for AI Agents" width="100%">
</p>

<p align="center">
  <a href="README.md" aria-label="Read in English"><img alt="English" src="assets/badge-english.svg" width="84" height="28"></a>
  <a href="README.ko.md" aria-label="한국어로 읽기"><img alt="Korean" src="assets/badge-korean.svg" width="84" height="28"></a>
</p>

<p align="center">
  <a href="LICENSE"><img alt="License MIT" src="https://img.shields.io/badge/license-MIT-f4b860"></a>
  <a href=".github/workflows/ci.yml"><img alt="CI workflow exists" src="https://img.shields.io/badge/CI-workflow%20exists-25334a"></a>
  <a href="SECURITY.md"><img alt="Security policy exists" src="https://img.shields.io/badge/security-policy%20exists-2d5a52"></a>
  <a href="docs/00_START_HERE.md"><img alt="Docs index exists" src="https://img.shields.io/badge/docs-index%20exists-5b8def"></a>
</p>

<h1 align="center">agent를 아직 실행하지 마세요. 먼저 의도를 컴파일하세요.</h1>

<p align="center"><strong>ni is a Project Intent Compiler for AI Agents.</strong></p>

`ni`는 planning conversation을 docs contract로 바꾸고, readiness를 확인하고,
accepted plan을 lock한 뒤 bounded downstream handoff prompt를 compile합니다.

<p align="center">
  <img src="assets/intent-lock-flow.svg" alt="Intent Lock Protocol flow: conversation, project contract, readiness gate, lock hash, bounded handoff." width="100%">
</p>

## 왜 ni인가

AI agents는 빠릅니다. `ni`는 느려야 하는 한 부분, 즉 implementation 전에
project intent가 무엇인지 결정하는 부분만 느리게 만듭니다.

- 빠진 users, acceptance criteria, risks, non-goals, blockers를 드러냅니다.
- Deterministic CLI rules로 planning readiness를 확인합니다.
- Accepted plan을 lock하고 trusted planning sources를 hash합니다.
- Downstream actors를 위한 짧은 prompt를 compile하지만 실행하지 않습니다.

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

새 shell을 연 뒤 global command를 verify합니다:

```bash
ni --help
ni version
```

Installer로 설치한 binary와, 추가했다면 ni-managed PATH block을 uninstall합니다:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

Homebrew: Planned / v0.5 candidate.

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

새 PowerShell session을 연 뒤 global command를 verify합니다:

```powershell
ni --help
ni version
```

Installer로 설치한 binary와 `ni`가 추가한 User PATH entry를 uninstall합니다:

```powershell
.\install.ps1 -Uninstall
```

Windows installer code와 static safety checks는 있습니다. 실제 Windows host
execution은 macOS-only development host에서는 deferred 상태이며 Windows install
transcript가 생기기 전까지 verified라고 claim하지 않습니다.

## 5분 첫 project

Public v0.5.0 install parity note: 아래 첫 project flow는 current-tree
`ni init .` onboarding을 사용합니다. Published v0.5.0 binary는 `ni --help`와
`ni version`을 verify하지만 positional `ni init .` form은 지원하지 않습니다.
자세한 증거는 [docs/126](docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.ko.md)에
있습니다.

```bash
mkdir my-project
cd my-project
ni init .
ni status --proof --next-questions
ni end
ni run --max-chars 4000
```

`ni init .`은 guided project intent wizard를 열고 `.ni/contract.json`,
`.ni/session.json`, `docs/plan/**`을 만듭니다.

`ni status --proof --next-questions`는 CLI-authoritative readiness gate입니다.
Model은 update를 draft할 수 있지만 readiness는 `ni status`가 결정합니다.

`ni end`는 CLI gate가 허용한 뒤 accepted plan을 lock하고
`.ni/plan.lock.json`을 씁니다.

`ni run --max-chars 4000`은 bounded downstream handoff prompt를 compile합니다.
Prompt, agents, shell commands, downstream work를 실행하지 않고 product readiness를
증명하지 않습니다.

## ni가 하는 일

| Command | Role |
| --- | --- |
| `ni init .` | Planning workspace와 guided intent draft를 만듭니다. |
| `ni status --proof --next-questions` | Readiness, blockers, next planning questions를 확인합니다. |
| `ni end` | CLI gate를 통해 accepted plan을 lock합니다. |
| `ni run --max-chars 4000` | Valid lock에서 bounded prompt를 compile합니다. |

## ni가 하지 않는 일

`ni`는 task runner, SPEC runner, multi-agent execution layer, queue, shell
adapter, PR automation system, release automation system, downstream execution
runtime이 아닙니다.

## Status

- v0.5.0 publication: verified.
- Release binary: Available.
- Curl installer: Available.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: Windows transcript 전까지 deferred.
- Model workspace packs: Experimental. Host-level/global install은 documented되기 전까지 unverified.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.

## 다음에 읽을 것

| Read | Why |
| --- | --- |
| [Install ni](docs/22_INSTALL.md) | 상세 install, release binary, curl installer, uninstall paths. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | Readiness, locking, hash trust, blocked handoff rules. |
| [터미널 없이 계획하기](docs/no-terminal.ko.md) | Assisted workflow boundaries; deterministic validation 아님. |
| [Model Workspace Status](docs/99_MODEL_WORKSPACE_STATUS.ko.md) | Experimental model workspace boundaries와 verification state. |
| [Benchmark Claim Boundaries](docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md) | Benchmark `READY`, `not_measured`, prompt evidence가 증명하는 것과 아닌 것. |
| [Command reference](docs/commands.ko.md) | Implemented CLI surface. |

License: `ni`는 [MIT License](LICENSE)로 배포됩니다.
