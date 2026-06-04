# Public Install Parity and Patch Readiness

## Current status

State:
- v0.5.0 publication: verified
- current-tree README two-path onboarding: implemented
- current-tree ni init . Bubble Tea TUI: implemented
- current-tree first-user smoke after TUI: FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES
- published v0.5.0 install parity: 이 문서에서 mismatch로 verified
- Windows real-host execution: macOS-only development host에서는 deferred
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Audit goal

이 audit는 public v0.5.0 install path가 current README onboarding path와 맞는지
확인한다:

```bash
ni --help
ni version
ni init .
ni status --proof --next-questions
ni end
ni run --max-chars 4000
```

Lane은 분리한다. Current-tree behavior는 published v0.5.0 behavior의 proof로
사용하지 않는다.

## Decision

PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED

Justification: curl installer로 설치한 published v0.5.0 binary는 `0.5.0`을
report하고 command-name `ni --help` / `ni version`은 pass한다. 하지만 current
README-required `ni init .` step은 `unknown init option: .`로 fail한다.
Current-tree `ni init .`은 동작하고 Bubble Tea v2 / Lip Gloss v2 TUI는 v0.5.0
tag 이후 추가되었으므로, onboarding path를 숨기기보다 v0.5.1 patch release
plan을 준비하는 것이 맞다.

## Lane comparison

| Lane | Binary source | Version output | ni init . behavior | ni status behavior | Result | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| current-tree | Current checkout에서 build한 `/private/tmp/ni-task205-current/bin/ni` | `0.0.0-dev` | `ni init . --yes` passed; temp project에 planning artifacts 생성. | Passed and empty first-user intent라 정직하게 `BLOCKED` report. | Pass with expected `BLOCKED` status | Source build이며 release proof가 아니다. |
| published v0.5.0 | `https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh` curl installer가 release asset `ni_0.5.0_darwin_arm64.tar.gz`를 temp BINDIR에 install | `0.5.0` | Failed: `unknown init option: .` | Init 실패 때문에 run하지 않음. | Mismatch | Help/version parity는 있지만 README first-project parity는 없다. |
| Windows real-host, deferred unless run | Not run | Not observed | Not observed | Not observed | Deferred | Windows host, VM, CI runner, tester transcript를 사용하지 않았다. |

## Published v0.5.0 transcript

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| command-name ni --help | `HOME="$tmp/home" PATH="$tmp/bin:..." sh -lc "command -v ni; ni --help"` | `ni`가 temp BINDIR에서 resolve되고 help가 성공적으로 exit. | `command -v ni`가 `/private/tmp/ni-task205-public.yvsGkZ/bin/ni`를 출력했고 help는 성공적으로 `/tmp/ni-task205-public-help.txt`로 redirect됨. | Yes | Temp HOME/BINDIR를 쓰는 fresh shell context. |
| command-name ni version | `ni version` | `0.5.0` | `0.5.0` | Yes | Installed artifact의 release linker version 확인. |
| ni init . | `"$tmp/bin/ni" init .` in temp project | README-required current-directory init succeeds. | `unknown init option: .` | No | 이 단계가 parity break. |
| ni status --proof --next-questions, if init succeeds | `ni status --proof --next-questions` | Init 성공 후 run. | Not run. | n/a | Init이 실패해 status를 실행할 workspace가 없었다. |
| generated artifact check, if init succeeds | `find "$tmp/project" -maxdepth 3 -type f` | `.ni/contract.json`, `.ni/session.json`, `docs/plan/**` 존재. | Not run. | n/a | Artifact creation 전 init 실패. |

Installer가 만든 temporary binary는 `sh install.sh --uninstall`로 제거했다. Uninstall은
temp binary와 temporary `.zshrc`의 ni-managed PATH block을 제거했다.

## README parity audit

| README claim | Current-tree evidence | Published v0.5.0 evidence | Pass? | Required action |
| --- | --- | --- | --- | --- |
| Curl installer로 v0.5.0 설치 후 새 shell에서 `ni --help` 실행. | Current installer script는 command-name verification을 지원한다. | Published v0.5.0 lane에서 `ni` resolve와 help exit 성공. | Yes | 유지. |
| Install 후 `ni version` 실행. | Current-tree source build는 `0.0.0-dev`; release build는 linker version이어야 한다. | Published v0.5.0은 `0.5.0` report. | Yes | Release/source distinction visible 유지. |
| 첫 project를 `ni init .`로 시작. | Current-tree lane passed. | Published v0.5.0은 `unknown init option: .`로 fail. | No | v0.5.1 patch release plan 준비; README는 이 parity note로 link. |
| `ni init .`이 guided project intent wizard를 연다. | Current-tree docs/code는 Bubble Tea v2 / Lip Gloss v2 TUI가 v0.5.0 이후 추가됐음을 보여준다. | Published v0.5.0은 init까지 도달하지 못함. | No for v0.5.0 | v0.5.1 patch inputs에 TUI/positional init changes 포함. |
| Windows real-host global install verified. | README는 deferred라고 말한다. | 이 audit에서 Windows run 없음. | Yes | Deferred wording 유지. |
| Homebrew is Available. | README는 `Homebrew: Planned / v0.5 candidate`라고 말한다. | Homebrew test/publication 없음. | Yes | Planned 유지. |
| `ni run` executes downstream work. | README는 bounded prompt compilation only라고 말한다. | 여기서 run하지 않음; boundary audit item. | Yes | Non-execution wording 유지. |

## Patch-release assessment

| Surface | Current-tree behavior | Published v0.5.0 behavior | Patch release needed? | Notes |
| --- | --- | --- | --- | --- |
| ni init . | Current-directory init과 guided/fallback behavior 동작. | `unknown init option: .`로 fail. | Yes | README-required command가 v0.5.0에 없다. |
| Bubble Tea TUI | `internal/tui/init`에 v0.5.0 이후 구현. | v0.5.0 release artifact에는 없음. | Yes | Patch validation에 TUI와 non-interactive fallback tests 포함. |
| macOS global install path handling | Current `install.sh`는 `--update-path`와 reversible PATH block 지원. | v0.5.0 tag `install.sh`에는 `--update-path`가 없다; current public raw `main` installer는 v0.5.0을 `--update-path`로 install 가능. | Yes | Installer script와 binary는 서로 다른 public surface다. |
| install.sh --update-path | Current tree와 public raw `main` script에 있음. | v0.5.0 tag에는 없음. | Yes | Patch plan은 installer/script state와 v0.5.1 artifact를 명시적으로 validate해야 한다. |
| Windows install.ps1 | Current tree에 구현. | v0.5.0 이후 추가; Windows real-host execution 없음. | Yes, for release inclusion; real-host proof still deferred | Transcript 없이 Windows execution claim 금지. |
| README two-path onboarding | Current tree에 구현. | Published v0.5.0 binary는 current first-project path를 완료할 수 없음. | Yes | README note 추가; useful onboarding 제거하지 않음. |
| docs/120 acceptance | v0.5.0 이후 추가되었고 global install acceptance 설명. | v0.5.0 release tag에는 없음. | Yes | Patch release docs set에 포함. |
| docs/125 first-user smoke | Current-tree smoke passed with notes. | Published v0.5.0 behavior proof가 아니라고 명시. | Yes | 이 audit는 그 proof gap을 mismatch로 닫는다. |

## Patch-release inputs

If patch is needed, list:
- expected version: v0.5.1
- changes to include:
  - current-tree `ni init .` positional target support
  - Bubble Tea v2 / Lip Gloss v2 guided init TUI
  - non-interactive fallback and existing-file/lockfile safety behavior
  - current README two-path onboarding note and docs/120 through docs/126 parity evidence
  - current `install.sh --update-path` behavior and Windows `install.ps1` packaging, with Windows execution still deferred unless a transcript exists
- validation required:
  - release artifact help/version for v0.5.1
  - temporary HOME/BINDIR curl installer install for v0.5.1
  - command-name `ni --help`
  - command-name `ni version`
  - `ni init .` in a temporary project
  - `ni status --proof --next-questions` if init succeeds
  - generated artifact check if init succeeds
  - protected project-root `.ni` diff remains empty
- release action explicitly not performed in this task.

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| published v0.5.0 | Help/version install proof는 claim 가능하지만 current `ni init .` parity는 claim 금지. | Help/version passed; `ni init .` failed. | No for full README parity | Decision은 patch needed. |
| current-tree behavior | Current-tree evidence만 current-tree onboarding claim에 사용 가능. | Current lane passed with `0.0.0-dev` source build. | Yes | v0.5.0 lane과 분리. |
| macOS install | Temporary HOME/BINDIR를 쓰고 user install pollution을 피해야 한다. | Temporary install 사용; uninstall confirmed. | Yes | Project-root `.ni` files 변경 없음. |
| Windows install | Windows transcript 없으면 deferred. | Deferred. | Yes | Static checks가 나중에 pass해도 real-host claim은 아님. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | Planned wording preserved. | Yes | Homebrew Available claim 없음. |
| ni init . | Current-tree feature; release parity는 public claim 전 verify 필요. | v0.5.0 mismatch observed. | No for v0.5.0 | v0.5.1 patch release plan recommended. |
| ni run | Prompt compiler only. | Not run; docs preserve non-execution boundary. | Yes | Generated prompts 실행하지 않음. |
| READY | Model judgment가 아니라 `ni status`에서만 나와야 한다. | Current temp project status는 `BLOCKED`; release status는 init 실패로 run하지 않음. | Yes | Readiness overclaim 없음. |
| Runtime execution boundary | Task runner, execution harness, shell/Codex adapter, queue, PR/release automation, downstream execution layer 없음. | Runtime behavior 추가/실행 없음. | Yes | Audit only. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | `M` after this task | Yes | docs/126을 link하는 narrow parity note. |
| README.ko.md | `M` after this task | Yes | Korean companion parity note. |
| docs/120* | clean | No | v0.5.0 이후 추가; 이 task에서는 edit 없음. |
| docs/124* | clean | No | Current-tree TUI evidence로 사용. |
| docs/125* | clean | No | Current-tree smoke evidence로 사용. |
| docs/126* | `??` after this task | Yes | Public install parity audit and Korean companion. |
| install.sh | clean | No | Current tree에는 `--update-path`; edit 없음. |
| install.ps1 | clean | No | Windows real-host run 없음. |
| scripts/* | clean | No | Script update 필요 없음. |
| temporary smoke directories | outside repo | No | `/private/tmp/ni-task205-current/**`, `/private/tmp/ni-task205-public.*`. |
| .ni/contract.json | clean and empty diff | No | Protected project-root planning state. |
| .ni/session.json | clean and empty diff | No | Protected project-root session state. |
| .ni/plan.lock.json | clean and empty diff | No | Protected project-root lockfile. |
| unexpected files | none observed before edits | No | Final 전에 recheck. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Initially clean; after edits, expected README and docs/126 changes only. |
| `git log --oneline -5` | Top commits include `93dc241 Document first-user onboarding smoke after TUI` and `4ae4a85 Redesign README and add Bubble Tea init TUI`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Empty before edits; rechecked after validation. |
| published v0.5.0 install smoke in temporary HOME/BINDIR | Partial pass, then mismatch: installer installed `0.5.0`; `ni init .` failed. |
| command-name `ni --help` from fresh shell context for published v0.5.0 | Passed; help exited successfully from temp BINDIR. |
| command-name `ni version` from fresh shell context for published v0.5.0 | Passed; output `0.5.0`. |
| `ni init .` with published v0.5.0 in temporary project | Failed: `unknown init option: .`. |
| `ni status --proof --next-questions` with published v0.5.0 if init succeeds | Not run because init failed. |
| current-tree first-user smoke for comparison | Passed with local source build; status reported expected `BLOCKED`. |
| `gofmt -w .` | Passed; protected planning files 변경 없음. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root는 blockers, deferrals, warnings 없이 `NI Intent Readiness: READY` report. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help에 `ni init [.]` 포함. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed after next-task prompt line을 explicit `must not include` boundary phrase로 변경. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |

## Changes made

- 이 public install parity audit 추가.
- Korean companion document 추가.
- README / README.ko에 이 audit를 가리키는 narrow parity note 추가.

## What this audit proves

- Published v0.5.0은 current README first-user path와 맞지 않는다. `ni init .`이
  `unknown init option: .`로 fail한다.
- Recommended next task는 v0.5.1 patch release plan이다.
- Project-root `.ni` files는 untouched.

## What this audit does not prove

- Windows real-host execution works unless a Windows transcript exists.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

A. prepare v0.5.1 patch release plan

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: prepare a v0.5.1 patch release plan for public install parity.

Goal:
Create a release-plan document that packages the current-tree onboarding fixes
needed because published v0.5.0 passes help/version install proof but fails the
current README-required `ni init .` step with `unknown init option: .`.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/22_INSTALL.md
- docs/120_GLOBAL_INSTALL_ACCEPTANCE.md
- docs/121_README_TWO_PATH_INSTALL_AND_INIT_TUI.md
- docs/124_REPO_CLEANUP_README_REDESIGN_INIT_TUI.md
- docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- install.sh
- install.ps1
- cmd/ni/
- internal/
- scripts/release-check.sh

Rules:
- Do not publish, tag, create a GitHub release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  `ni end` on the project root, relock the project root, edit
  `.ni/plan.lock.json`, execute generated prompts, or add runtime execution
  behavior.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution without a Windows transcript.
- Keep `Skills are UX; CLI is authority.`
- Keep `ni run` as bounded prompt compilation only.

Plan content required:
- decision: V0_5_1_PATCH_RELEASE_PLAN_READY or BLOCKED
- exact release rationale tied to docs/126 evidence
- included changes: `ni init .` positional support, Bubble Tea v2 / Lip Gloss v2
  TUI, non-interactive fallback, init safety behavior, current installer PATH
  handling, README/docs parity notes
- excluded changes (patch must not include): task runner, SPEC runner, shell/Codex adapter, queue, PR
  automation, release automation, Homebrew Available claim, Windows verified
  claim, downstream execution
- validation matrix for v0.5.1 local preflight and post-publication hosted
  artifact smoke
- protected `.ni` file boundary
- rollback / docs correction path if v0.5.1 publication is not approved
- complete next executable prompt for the authorized v0.5.1 release preflight

Validation:
- git status --short
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, release-plan decision, validation results, protected
`.ni` diff, and confirmation that no publish/tag/release/upload/project-root
relock/generated prompt execution occurred.
```
