# External User Validation Plan

## Current status

State:
- v0.5.1 release: published and verified.
- v0.5.1 release decision: V0_5_1_RELEASE_EXECUTED_WITH_NOTES.
- Public install parity mismatch from v0.5.0: tested macOS arm64 path에서 closed.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: Windows transcript 전까지 deferred.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Plan goal

이 plan은 실제 external user가 maintainer 도움 없이 public v0.5.1 README install
및 first-project onboarding path를 따라갈 수 있는지 validate하는 방법을 정의한다.
이 문서는 protocol, transcript template, evidence grading model,
claim-boundary checklist다. Later audit에 real tester transcript가 첨부되기
전까지 external users가 succeeded했다고 증명하지 않는다.

Primary question: external user가 README를 열고 macOS 또는 Windows를 선택한 뒤
`ni`를 install하고, 새 shell 또는 PowerShell에서 global command를 verify하고,
`ni init .`으로 first project를 만들고, `ni status --proof --next-questions`를
실행하고, `ni run`이 bounded handoff prompt만 compile하며 downstream work를
실행하지 않는다는 점을 이해하고, 요청받으면 uninstall할 수 있는가?

## Decision

EXTERNAL_VALIDATION_PLAN_READY_WITH_NOTES

Justification: external validation protocol, command sheets, transcript
template, evidence grading, pilot packet, claim-boundary audit이 준비되었다.
Notes는 이 task에 external user transcript가 포함되지 않았고, Windows real-host
validation이 여전히 pending이며, tester recruitment가 open 상태이기 때문에 남는다.

## Validation goals

| Goal | Evidence required | Pass criteria | Notes |
| --- | --- | --- | --- |
| README clarity | Tester가 README만 보고 시작하고 confusing했던 점을 기록. | Maintainer 설명 없이 platform path를 선택할 수 있다. | Feedback은 polish task를 만들 수 있다. |
| macOS install success | Install command, new shell, command-name checks, uninstall if tested를 포함한 full macOS transcript. | `ni --help`가 success이고 `ni version`이 command name 기준 `0.5.1`을 report한다. | macOS external pilot에 필요. |
| Windows install success | Real Windows host, VM, CI runner, 또는 external tester의 full Windows transcript. | New PowerShell session에서 `Get-Command ni`, `ni --help`, `ni version`이 동작한다. | Windows real-host deferral closure에 필요. |
| Command-name global resolution | New shell 이후 `command -v ni` 또는 `Get-Command ni`. | Absolute binary path 없이 shell이 `ni`를 resolve한다. | Local artifact execution과 별개다. |
| `ni --help` | Command transcript. | Help가 success이고 implemented CLI surface를 보여준다. | Help success만으로 first-project proof가 아니다. |
| `ni version` | Command transcript. | Output이 `0.5.1`이다. | External release validation에서 source-tree `0.0.0-dev`는 acceptable하지 않다. |
| `ni init .` | Transcript와 created-file listing. | Guided init이 시작되고 complete 또는 safe exit하며 unexpected overwrite가 없다. | TUI usability feedback이 필요하다. |
| TUI usability | Tester notes, 안전한 경우 screenshot, stuck point. | Bubble Tea v2 / Lip Gloss v2 flow를 navigate, complete, 또는 intentional exit할 수 있다. | Non-interactive fallback은 relevant할 때만 테스트한다. |
| Generated planning files | `find` 또는 directory listing. | `.ni/contract.json`, `.ni/session.json`, `docs/plan/**` 존재; `.ni/plan.lock.json`은 init으로 생성되지 않는다. | Real repo root가 아닌 fixture project를 사용한다. |
| Status comprehension | `ni status --proof --next-questions` transcript와 tester answers. | Tester가 basic `BLOCKED` versus `READY`와 `READY`가 CLI planning readiness only임을 이해한다. | Blank/incomplete first project는 정직하게 `BLOCKED`일 수 있다. |
| `ni run` comprehension | Tester paraphrase; optional command only after understanding. | Tester가 `ni run`이 bounded prompt를 compile하고 downstream work를 실행하지 않는다고 이해한다. | Optional command는 `--max-chars 4000`을 사용한다. |
| Uninstall path | Uninstall transcript 또는 tester note locating docs. | Tester가 uninstall하거나 relevant uninstall command를 찾을 수 있다. | Uninstall은 installer-managed binary/PATH만 제거해야 한다. |
| Claim-boundary comprehension | Tester answers와 transcript review. | Tester가 Homebrew Available, transcript 없는 Windows verified, deterministic no-terminal, agent execution을 결론내리지 않는다. | Boundary failure는 product/doc failure다. |

## Validation cohorts

| Cohort | Why needed | Required? | Status | Notes |
| --- | --- | --- | --- | --- |
| Maintainer sanity check | External users에게 요청하기 전 release path를 확인. | 이 plan에는 No; 별도 task에서 already done. | Tested macOS arm64 path done. | docs/132 post-release verification 참고. |
| External macOS user | Separate user 또는 machine에서 public README path를 테스트. | macOS external validation에는 Yes. | Future validation. | Maintainer setup을 보지 않은 user가 좋다. |
| External Windows user | Windows real-host execution deferral을 닫는다. | Windows real-host execution claim에는 Yes. | Future validation. | Static PowerShell checks만으로는 부족하다. |
| Backend developer familiar with CLI tools | Technical users도 만나는 friction을 찾는다. | Optional. | Future validation. | Plan ready의 필수 조건은 아니다. |
| AI-agent user who is not a Go developer | Intended audience에게 product story가 전달되는지 테스트. | Optional. | Future validation. | README/onboarding language에 유용하다. |

## Success criteria

| Criterion | Required? | Evidence | Pass state |
| --- | --- | --- | --- |
| User installs without maintainer intervention. | Yes | Transcript and tester note. | PASS or PASS_WITH_NOTES. |
| User opens a new shell or PowerShell and runs `ni` by command name. | Yes | `command -v ni` or `Get-Command ni`. | PASS. |
| `ni --help` succeeds. | Yes | Command transcript. | PASS. |
| `ni version` reports `0.5.1`. | Yes | Command transcript. | PASS. |
| User runs `ni init .` in a fixture project. | Yes | Command transcript. | PASS or PASS_WITH_NOTES. |
| User can navigate or complete the Bubble Tea v2 / Lip Gloss v2 TUI. | Yes | Transcript and feedback. | PASS or PASS_WITH_NOTES. |
| User can identify created files. | Yes | File listing. | PASS. |
| User can run `ni status --proof --next-questions`. | Yes | Command transcript. | PASS. |
| User understands `BLOCKED` versus `READY` at a basic level. | Yes | Feedback answer. | PASS or PASS_WITH_NOTES. |
| User understands `ni run` does not execute downstream work. | Yes | Feedback answer; optional command. | PASS. |
| User can uninstall or knows where uninstall docs are. | Yes | Command transcript or note. | PASS or PASS_WITH_NOTES. |
| No protected `.ni` files in the real project root are modified. | Yes | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json`. | Empty diff. |

## Failure criteria

| Failure criterion | Severity | Evidence | Blocker condition |
| --- | --- | --- | --- |
| README install instructions are not enough. | High | Tester cannot choose or complete a path. | Maintainer explanation이 필요하면 FAIL_BLOCKER. |
| Install script fails. | High | Installer transcript. | README scope 밖의 local policy issue가 문서화되지 않는 한 FAIL_BLOCKER. |
| PATH/global command resolution fails. | High | `command -v ni` or `Get-Command ni`. | FAIL_BLOCKER. |
| `ni version` is not `0.5.1`. | High | Version transcript. | FAIL_BLOCKER. |
| `ni init .` fails. | High | Command transcript. | User가 intentionally exit했고 corruption이 없지 않은 한 FAIL_BLOCKER. |
| TUI is confusing or unusable. | Medium to high | Tester feedback. | Proceed 또는 safe exit가 불가능하면 FAIL_BLOCKER. |
| TUI overwrites files unexpectedly. | High | File hashes or listing. | FAIL_BLOCKER. |
| `.ni/plan.lock.json` is created by init. | High | File listing. | FAIL_BLOCKER. |
| `ni status` makes user think product readiness is proven. | High | Tester feedback. | Claim-boundary comprehension에 대해 FAIL_BLOCKER. |
| README implies Homebrew Available. | High | README audit. | Corrected 전까지 FAIL_BLOCKER. |
| Windows tester cannot complete install due to PATH or PowerShell policy issue. | Medium to high | PowerShell transcript. | Windows real-host claim에는 FAIL_BLOCKER; documented이면 general plan에는 PASS_WITH_NOTES 가능. |
| User thinks `ni` runs the agent or implements the project. | High | Tester feedback. | Product comprehension에 대해 FAIL_BLOCKER. |

## Transcript template

```text
External ni v0.5.1 validation transcript

Tester:
Date:
Platform:
OS version:
Shell / PowerShell version:
Architecture:
Install method:
README URL used:
Release URL used:

Privacy note:
- You may redact usernames, home paths, machine names, project names, and tokens.
- Do not include secrets, credentials, private repository URLs, API keys, or personal tokens.

Command transcript:

1. Install command
<paste command and output>

2. New shell / PowerShell command-name check
macOS:
$ command -v ni
$ ni --help
$ ni version

Windows:
PS> Get-Command ni
PS> ni --help
PS> ni version

<paste command and output>

3. Create project directory
$ mkdir ni-external-validation
$ cd ni-external-validation

or:
PS> mkdir ni-external-validation
PS> cd ni-external-validation

<paste command and output>

4. First project init
$ ni init .

or:
PS> ni init .

<paste command and output>

5. Observed files
Expected:
- .ni/contract.json
- .ni/session.json
- docs/plan/**
- .ni/plan.lock.json should not be created by init

<paste file listing>

6. Status proof
$ ni status --proof --next-questions

or:
PS> ni status --proof --next-questions

<paste command and output>

7. Optional prompt compilation
Only run this if you understand that it compiles a bounded handoff prompt and
does not execute agents, shell commands, or downstream work.

$ ni run --max-chars 4000

or:
PS> ni run --max-chars 4000

<paste command and output, or write NOT_RUN>

8. Uninstall, if tested
macOS:
$ BINDIR="$HOME/.local/bin" sh install.sh --uninstall

Windows:
PS> .\install.ps1 -Uninstall

<paste command and output, or write NOT_RUN>

User feedback:
- What was confusing?
- Did README explain enough?
- Did TUI explain enough?
- Did you understand what ni does not do?
- Did you understand `BLOCKED` versus `READY`?
- Did you understand that `ni run` does not execute downstream work?
- Did any command fail?
```

## macOS command sheet

```bash
VERSION="0.5.1"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

새 shell을 연다:

```bash
command -v ni
ni --help
ni version
```

Fixture project를 만든다:

```bash
mkdir ni-external-validation
cd ni-external-validation
ni init .
find . -maxdepth 3 -type f | sort
ni status --proof --next-questions
```

Optional prompt compilation은 user가 `ni run`이 prompt만 compile하고 downstream
work를 실행하지 않는다는 점을 이해한 뒤에만 실행한다:

```bash
ni run --max-chars 4000
```

Uninstall, if tested:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Windows command sheet

```powershell
$Version = "0.5.1"
Invoke-WebRequest "https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1" -OutFile "install.ps1"
Get-Content .\install.ps1
.\install.ps1 -DryRun -Version $Version
.\install.ps1 -Version $Version
```

새 PowerShell session을 연다:

```powershell
Get-Command ni
ni --help
ni version
```

Fixture project를 만든다:

```powershell
mkdir ni-external-validation
cd ni-external-validation
ni init .
Get-ChildItem -Recurse -File | Sort-Object FullName | Select-Object FullName
ni status --proof --next-questions
```

Optional prompt compilation은 user가 `ni run`이 prompt만 compile하고 downstream
work를 실행하지 않는다는 점을 이해한 뒤에만 실행한다:

```powershell
ni run --max-chars 4000
```

Uninstall, if tested:

```powershell
.\install.ps1 -Uninstall
```

Windows notes:
- Installer는 기본적으로 Machine PATH가 아니라 User PATH를 사용해야 한다.
- Windows transcript가 생기기 전까지 Windows real-host execution verified라고 claim하지 않는다.
- PowerShell execution policy issue가 나타나면 capture한다.

## Evidence grading

| Grade | Meaning | Allowed claim | Forbidden claim |
| --- | --- | --- | --- |
| PASS | Required evidence가 captured되고 criteria를 충족했다. | Shown platform, host, version, transcript의 tested path가 passed. | 모든 users, 모든 platforms, untested install paths가 succeeded. |
| PASS_WITH_NOTES | Path는 completed했지만 non-blocking friction, deferrals, comprehension notes가 있다. | Tested path completed with stated notes. | Notes가 resolved 또는 irrelevant. |
| FAIL_BLOCKER | Required step이 failed했거나 claim-boundary failure가 있다. | Scoped path에서 blocker를 발견했다. | Product가 scoped evidence 밖에서 broadly unusable. |
| NOT_RUN | Step을 attempt하지 않았다. | 해당 step에는 evidence가 없다. | 해당 step의 success 또는 failure claim. |

Evidence categories:
- install
- command resolution
- version
- init TUI
- generated files
- status comprehension
- uninstall
- claim-boundary comprehension

## Pilot validation packet

Testers에게 보낼 packet:

```text
Please validate ni v0.5.1 from the public README.

Start here:
- README: https://github.com/Nam-Cheol/ni
- Release: https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1

Use the command sheet for your platform:
- macOS: install.sh path from README
- Windows: install.ps1 path from README

Expected timebox: 20 to 30 minutes.

Please paste the transcript template results. You may redact usernames, home
paths, machine names, project names, and tokens. Do not include secrets.

Important product boundary:
- ni compiles project intent before execution.
- ni run compiles a bounded handoff prompt only.
- ni does not execute agents, shell commands, downstream work, PR automation,
  release automation, queues, or task-runner behavior.
- Homebrew remains Planned / v0.5 candidate.
- Windows real-host execution is not verified until a Windows transcript exists.
```

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| external validation | Transcript 전에는 external users succeeded claim 금지. | 이 plan에는 external transcript가 없다. | Yes | Decision은 WITH_NOTES. |
| Windows real-host execution | Windows transcript 없이는 deferred. | Deferred. | Yes | Static PowerShell checks는 real-host execution이 아니다. |
| Homebrew | Homebrew: Planned / v0.5 candidate 유지. | Preserved. | Yes | Homebrew Available claim 없음. |
| `ni run` | Bounded prompt compilation only. | Preserved. | Yes | Optional pilot command도 non-executing으로 명시. |
| READY | Product readiness가 아니라 CLI planning readiness only. | Preserved. | Yes | Tester는 `BLOCKED` versus `READY`를 설명해야 한다. |
| TUI | Downstream execution이 아니라 guided first-project intent setup. | Preserved. | Yes | TUI usability는 onboarding UX로 테스트한다. |
| runtime execution boundary | Task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation, release automation, downstream execution layer 없음. | Preserved. | Yes | 이 task는 docs only. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | clean | No | 이 task에서는 read only. |
| README.ko.md | clean | No | 이 task에서는 read only. |
| docs/132* | clean | No | Current status source evidence. |
| docs/133* | added | Yes | External validation plan and Korean companion. |
| docs/51* | modified | Yes | docs/133로 향하는 narrow roadmap pointer. |
| `.ni/contract.json` | clean | No | Protected. |
| `.ni/session.json` | clean | No | Protected. |
| `.ni/plan.lock.json` | clean | No | Protected. |
| unexpected files | none expected | No | Commit 전 recheck. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; docs/51 modifications와 docs/133 additions만 있었다. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; root는 blockers, deferrals, warnings 없이 `NI Intent Readiness: READY`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

## Changes made

| File | Why |
| --- | --- |
| `docs/133_EXTERNAL_USER_VALIDATION_PLAN.md` | External validation protocol, transcript template, platform command sheets, evidence grading, pilot packet, claim-boundary audit 추가. |
| `docs/133_EXTERNAL_USER_VALIDATION_PLAN.ko.md` | 동일한 claim boundaries를 가진 Korean companion. |
| `docs/51_POST_RELEASE_ROADMAP.md` | 이 external validation plan으로 향하는 narrow pointer 추가. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Korean roadmap companion pointer. |

## What this plan proves

- External validation protocol is ready.
- Claim boundaries are explicit.
- Platform-specific evidence requirements are defined.
- Transcript template and pilot packet are ready for testers.

## What this plan does not prove

- External users succeeded.
- Windows real-host execution works.
- Homebrew is Available.
- Downstream execution succeeds.
- No-terminal is deterministic.
- Benchmark effect size or causal impact.

## Recommended next task

Selected next task: A. run macOS external pilot validation.

Selection rationale: v0.5.1 publication과 tested macOS arm64 install parity는 이미
문서화되어 있지만 independent external user transcript는 없다. 가장 friction이 낮은
next evidence gap은 macOS external pilot이다. Windows pilot validation은 Windows
real-host execution을 닫기 위한 별도 required path로 남는다.

## Next task prompt

```text
Task: run one macOS external pilot validation for ni v0.5.1.

Do not ask the tester to use private maintainer knowledge. Send them only the
README link, v0.5.1 release link, macOS command sheet, transcript template, and
privacy note from docs/133_EXTERNAL_USER_VALIDATION_PLAN.md.

Tester packet:
- README: https://github.com/Nam-Cheol/ni
- Release: https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1
- Expected timebox: 20 to 30 minutes.
- The tester may redact usernames, home paths, machine names, project names,
  and tokens.
- The tester must not paste secrets.

Ask the tester to run:
1. Inspect and install with install.sh for VERSION=0.5.1.
2. Open a new shell.
3. Run command-name `command -v ni`, `ni --help`, and `ni version`.
4. Create a fixture project directory.
5. Run `ni init .`.
6. List created files and confirm `.ni/plan.lock.json` was not created by init.
7. Run `ni status --proof --next-questions`.
8. Explain in their own words that `ni run` compiles a bounded prompt and does
   not execute downstream work.
9. Optionally run `ni run --max-chars 4000` only after that explanation.
10. Test uninstall or identify the uninstall command.

After the transcript arrives, create a result audit. Grade install, command
resolution, version, init TUI, generated files, status comprehension, uninstall,
and claim-boundary comprehension as PASS, PASS_WITH_NOTES, FAIL_BLOCKER, or
NOT_RUN. Do not claim external validation succeeded unless the transcript
supports it. Do not claim Windows real-host execution, Homebrew Available,
no-terminal deterministic validation, downstream execution success, or product
readiness from this macOS transcript.
```
