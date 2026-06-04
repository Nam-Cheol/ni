# First-User Onboarding Smoke Test

## Current status

State:

- v0.5.0 publication: verified
- current-tree README two-path onboarding: implemented
- ni init . guided onboarding: implemented
- macOS current-tree smoke: 이번 task에서 verified
- Windows real-host execution: macOS-only development host에서는 deferred
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Smoke goal

이 smoke test는 first user가 current repository tree에서 README 형태의 경로를 따라
usable planning workspace에 도달할 수 있는지 확인한다. Release publication check,
Windows real-host transcript, downstream execution proof가 아니다.

## Smoke lanes

| Lane | Purpose | Binary source | Expected version | Result | Notes |
| --- | --- | --- | --- | --- | --- |
| current-tree onboarding smoke | 이 checkout에서 README first-user command-name onboarding을 verify한다. | Local `go build` output을 temporary user bin directory로 copy. | Current tree output; observed `0.0.0-dev`. | macOS에서 verified. | Temporary `HOME`, `PATH`, project directories를 사용했다. |
| published v0.5.0 install smoke | Released installer 또는 release artifact behavior를 verify한다. | Published GitHub Release assets 또는 curl installer. | `0.5.0`. | 이번 task에서는 run하지 않음. | `install.sh --dry-run --version 0.5.0`은 asset selection 확인용으로만 실행했다. Release asset을 download하거나 install하지 않았다. |
| Windows real-host smoke | Windows에서 PowerShell install, new-session command resolution, help/version, uninstall을 verify한다. | Real Windows host, VM, CI runner, external tester의 `install.ps1`. | Published lane은 `0.5.0`, current-tree Windows lane은 current-tree version. | Deferred. | Static Windows installer safety checks는 passed; real-host execution은 claim하지 않는다. |

## macOS current-tree transcript

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| build current tree | `GOCACHE=/private/tmp/ni-go-cache go build -o /private/tmp/ni-task202-bin/ni ./cmd/ni` | Local binary builds. | Output 없이 build됨. | Yes | Published v0.5.0 binary가 아니라 Lane A이다. |
| fresh shell command lookup | `command -v ni` | Command name으로 resolve된다. | `/private/tmp/ni-task202-smoke/bin/ni` | Yes | Temporary `HOME`과 `PATH`를 둔 `/bin/zsh -f`에서 실행. |
| help | `ni --help` | Absolute path 없이 help succeeds. | Output에 `ni is a project intent compiler.` 포함. | Yes | Full output은 `/private/tmp/ni-task202-smoke/help.txt`에 capture. |
| version | `ni version` | Absolute path 없이 version succeeds. | `0.0.0-dev` | Yes | Current-tree build behavior. |
| project directory | `mkdir first-user-project && cd first-user-project` | Fresh project directory exists. | `/private/tmp/ni-task202-smoke/projects` 아래 생성. | Yes | Project root files를 사용하지 않았다. |
| init | `ni init . --yes` | Init creates planning artifacts and no lockfile. | `.ni/contract.json`, `.ni/session.json`, 12개 `docs/plan` files 생성. | Yes | `--yes`는 supported unattended path이다. |
| status | `ni status --proof --next-questions` | Readiness gate가 init 뒤 run된다. | Command succeeded and reported `NI Intent Readiness: BLOCKED`. | Yes | Missing intent가 있는 minimal first-run workspace에서는 expected result. |
| interactive stdin | `printf ... | ni init . --interactive` | Interactive guided init이 stdin에서 동작한다. | `.ni/contract.json`과 12개 `docs/plan` files 생성. | Yes | `/private/tmp/ni-task202-safety` 아래 separate temporary fixture. |

## Generated artifact check

| Artifact | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `.ni/contract.json` | `ni init .` 뒤 exists. | Exists. | Yes | Temporary first-user project only. |
| `.ni/session.json` | `ni init .` 뒤 exists. | Exists. | Yes | Temporary first-user project only. |
| `docs/plan/**` | `ni init .` 뒤 exists. | 12 files. | Yes | `docs/plan/00_project_brief.md`부터 `11_decision_log.md`까지 포함. |
| `.ni/plan.lock.json` | Init이 만들면 안 된다. | Absent. | Yes | Lock command를 run하지 않았다. |
| downstream generated prompt | Init/status가 execute 또는 generate하면 안 된다. | `.ni/generated` absent. | Yes | Temporary project에서 `ni run`을 run하지 않았다. |

## Existing project safety

| Check | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| first init | Planning files를 만든다. | Expected planning files 생성. | Yes | Temporary project `/private/tmp/ni-task202-safety/projects/repeat-project`. |
| second init | Existing files를 silently overwrite하지 않는다. | `existing ni planning files found; ni init will not overwrite them.` 및 `adding missing files only.` 출력. | Yes | `.ni/contract.json`과 `.ni/session.json` SHA-256 값이 unchanged. |

## Lockfile safety

| Check | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| locked project init | Existing `.ni/plan.lock.json`은 수정하지 않는다. | Lockfile warning과 `No files changed by ni init; the lockfile was not modified.` 출력. | Yes | Sentinel lockfile SHA-256 unchanged. |
| root lockfile | Project-root `.ni/plan.lock.json`을 touch하지 않는다. | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` output 없음. | Yes | Project-root `ni end` 또는 relock을 run하지 않았다. |

## Installer smoke

| Check | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `install.sh --dry-run --version 0.5.0` | Install 없이 expected release asset을 선택한다. | `ni_0.5.0_darwin_arm64.tar.gz`와 v0.5.0 download URLs 출력. | Yes | Dry-run only; public v0.5.0 install smoke가 아니다. |
| `scripts/test-install-sh.sh` | Local fake release fixture가 dry-run, checksum, install, PATH resolution, uninstall을 test한다. | `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh`를 통해 passed. | Yes | Network 없이 installer behavior를 verify한다. |
| uninstall scope | Installed binary와 ni-managed PATH block만 제거한다. | install-check 내부의 `scripts/test-install-sh.sh`를 통해 passed. | Yes | Script가 binary removal과 ni-managed PATH block removal을 verify했다. |

## Windows status

| Area | Status | Evidence | Notes |
| --- | --- | --- | --- |
| Windows installer implementation | Present. | `install.ps1` exists. | `%LOCALAPPDATA%\ni\bin`에 install하고 User PATH만 update한다. |
| Windows static safety checks | Passed. | `python3 scripts/check-install-ps1.py`가 `Windows installer static safety checks passed` 출력. | Static proof only. |
| Windows real-host execution | Deferred. | Windows host, VM, CI runner, external tester transcript를 사용하지 않았다. | Windows global install verified라고 claim하지 않는다. |

## README alignment

| README claim | Smoke evidence | Pass? | Notes |
| --- | --- | --- | --- |
| macOS path starts with installer inspection, dry-run, install, new shell, `ni --help`, and `ni version`. | README에 path가 있고, current-tree smoke는 fresh shell context에서 command-name help/version을 별도로 verify했다. | Yes | README public installer path는 release-lane language로 유지. |
| First project uses `mkdir`, `cd`, `ni init .`, `ni status --proof --next-questions`, `ni end`, and `ni run --max-chars 4000`. | README의 "First project in 5 minutes" flow에 전체 flow가 있다. | Yes | 이 task는 project root에서 `ni end` 또는 `ni run`을 run하지 않았다. |
| Windows path uses PowerShell installer and a new PowerShell session. | Static installer checks passed; real-host execution deferred. | Yes | README는 deferred Windows claim을 유지한다. |
| `ni run` compiles but does not execute downstream work. | README가 이 boundary를 말하고, smoke는 generated prompt를 실행하지 않았다. | Yes | Runtime execution boundary preserved. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS global install | Current-tree smoke는 absolute path 없이 command-name resolution을 prove해야 한다. | Temporary `PATH`에서 `ni`를 resolve하고 help/version을 run했다. | Yes | Tested local binary lane을 prove하며, 모든 shell configuration을 prove하지 않는다. |
| Windows global install | Windows transcript 없이 verified claim 금지. | Static safety check only; real-host execution deferred. | Yes | Windows verified claim 추가 없음. |
| Homebrew | Homebrew remains Planned / v0.5 candidate. | README/docs가 Planned wording을 유지한다. | Yes | Homebrew Available claim 없음. |
| ni init . | Init은 planning artifacts만 만든다. | Temporary projects가 docs와 `.ni` planning files를 만들었다. | Yes | Init이 lockfile 또는 generated prompt를 만들지 않았다. |
| ni status | CLI readiness gate가 authority이다. | Status가 run되고 missing first-run intent 때문에 `BLOCKED`를 report했다. | Yes | Model judgment가 readiness를 선언하지 않았다. |
| ni end | 이 task에서 project root에 lock command를 run하면 안 된다. | Project root에서 run하지 않았다. | Yes | Root relock 없음. |
| ni run | Prompt compiler only; execution 없음. | Temporary smoke나 project root에서 run하지 않았다. | Yes | Generated prompt execution 없음. |
| READY | CLI-only readiness, not product readiness. | First-run temp project는 `BLOCKED`로 남았다. | Yes | Blockers가 answer되기 전에는 expected. |
| Runtime execution boundary | Task runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation, release automation, downstream execution layer 없음. | Runtime behavior 추가 없음. | Yes | Docs-only Task 202 pass. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | `M` | Existing change; Task 202 new edit expected 없음. | 이 task 전부터 README visual/onboarding change가 있었다. |
| README.ko.md | `M` | Existing change; Task 202 new edit expected 없음. | Korean companion은 이 task 전부터 modified. |
| cmd/ni/* | unchanged in this task | No | Read and built only. |
| internal/* | unchanged in this task | No | Read and tested only. |
| docs/121* | unchanged in this task | No | Prior onboarding evidence로 사용. |
| docs/123* | `??` | Yes | New smoke-test documentation and Korean companion. |
| install.sh | unchanged | No | Dry-run/read only. |
| install.ps1 | unchanged | No | Static safety check/read only. |
| scripts/* | unchanged in this task | No | Validation only. |
| temporary smoke directories | not tracked | No | `/private/tmp/ni-task202-*` 아래 사용. |
| `.ni/contract.json` | unchanged | No | Protected project-root planning state. |
| `.ni/session.json` | unchanged | No | Protected project-root planning state. |
| `.ni/plan.lock.json` | unchanged | No | Protected project-root lockfile. |
| unexpected files | Task 202에서 생성한 unexpected files는 현재 없음 | No | Final validation 뒤 recheck. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; final status는 pre-existing README/docs/asset changes와 new `docs/123*` files를 보여준다. |
| `gofmt -w .` | 별도로 run하지 않았다. Task 202는 Go files를 touch하지 않았고 `quality.sh`가 successfully completed. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /private/tmp/ni-task202-bin/ni ./cmd/ni` | Passed. |
| current-tree first-user smoke in temporary `HOME` and project directory | Passed. |
| command-name `ni --help` from fresh shell context | Passed. |
| command-name `ni version` from fresh shell context | Passed; `0.0.0-dev`. |
| `ni init . --yes` in temporary project | Passed. |
| `ni status --proof --next-questions` in temporary project | Passed; `BLOCKED` with first-run next questions. |
| repeated `ni init .` safety smoke | Passed; contract/session hashes unchanged. |
| lockfile safety smoke | Passed; sentinel `.ni/plan.lock.json` hash unchanged. |
| interactive stdin smoke | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `sh install.sh --dry-run --version 0.5.0` | Passed; dry-run only. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root는 `NI Intent Readiness: READY`와 no blockers, deferrals, warnings를 report했다. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; `0.0.0-dev` 출력. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed; fixture prompt compilation은 project-root execution과 분리됐다. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed; fixture `ni end`/`ni run` paths는 project-root lock state와 분리됐다. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed; source, build, temporary global command, installer, uninstall, Windows static checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; output 없음. |

## Changes made

- 이 Task 202 first-user onboarding smoke-test record를 추가했다.
- Claim boundary가 matching되는 Korean companion record를 추가했다.
- Smoke가 current-tree onboarding blocker를 발견하지 않았기 때문에 Task 202 code
  changes는 만들지 않았다.

## What this smoke proves

- current-tree macOS first-user command-name onboarding works for the tested
  path;
- `ni init .` creates expected planning artifacts in a temporary project;
- `ni status --proof --next-questions` runs after init;
- repeated init and lockfile safety behavior are bounded in temporary fixtures;
- project-root `.ni` files were not modified.

## What this smoke does not prove

- Windows real-host execution works;
- Homebrew is Available;
- every shell configuration works;
- external users succeed;
- downstream execution succeeds;
- no-terminal is deterministic;
- published v0.5.0 artifact가 current-tree `ni init .` behavior를 포함한다는 점은
  해당 artifact로 verify하기 전까지 증명하지 않는다.

## Recommended next task

Selected: B. external user validation plan

Current-tree macOS first-user smoke가 passed했고 다음 proof gap은 external user가
repository-local context 없이 clean machine에서 README path를 reproduce할 수 있는지
여부이기 때문에 이것을 선택한다.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: External first-user onboarding validation plan.

Goal:
Create a claim-safe plan for an external user or tester to validate the README
first-user path without relying on repository-local assumptions.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/22_INSTALL.md
- docs/install-curl.md
- docs/install-curl.ko.md
- docs/123_FIRST_USER_ONBOARDING_SMOKE_TEST.md
- docs/123_FIRST_USER_ONBOARDING_SMOKE_TEST.ko.md
- install.sh
- install.ps1
- scripts/check-install-docs.py

Rules:
- Do not publish, tag, release, upload, or push.
- Do not run ni end on the project root.
- Do not relock the project root.
- Do not edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json.
- Do not execute generated prompts.
- Do not claim Windows real-host verification without a Windows transcript.
- Do not mark Homebrew Available.
- Do not imply ni run executes downstream work.

Work:
1. Define a short external tester script for macOS that records shell, PATH,
   installer command, new-shell command-name checks, ni init ., and
   ni status --proof --next-questions.
2. Define the exact evidence transcript required before claiming external
   success.
3. Keep current-tree and published v0.5.0 release lanes separate.
4. Keep Windows as deferred unless a real Windows host transcript is supplied.
5. Add a concise docs note and Korean companion if the repo maintains companion
   docs.

Validation:
- git status --short
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, selected validation lane, exact external evidence needed,
validation results, protected .ni diff, and confirmation that no release action,
project-root relock, generated prompt execution, Homebrew Available claim, or
Windows verified claim was added.
```
