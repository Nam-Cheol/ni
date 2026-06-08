# v0.6.0 Public Install Parity and Release Readiness

## Current status

State at readiness-audit time:
- v0.5.1 release: published and verified
- v0.6.0 release: pre-publication; docs/140 post-release verification이 supersede
- Namba Intent rename: implemented in current tree
- primary command: namba-intent
- deprecated ni shim: transition-only
- .ni/ compatibility: preserved
- public install retrieval of namba-intent: audit 시점 release-gated; 이후 docs/140에서 verified
- Homebrew: Planned / v0.5 candidate
- Windows real-host verification: pending
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

## Audit goal

이 audit는 pre-release current-tree rename이 v0.6.0 release preparation으로 진행할 준비가
되었는지, 그리고 public install docs가 안전하게 bounded되어 있는지 확인한다.
Publish, tag, asset upload, release workflow, GoReleaser publish, project-root
`namba-intent end`, project-root relock, generated prompt execution, Homebrew
material creation, downstream execution behavior는 수행하지 않는다.

## Decision

V0_6_0_RELEASE_READINESS_READY_WITH_NOTES

Rationale: current-tree command rename, installer configuration, release asset
naming, `.ni/` compatibility, deprecated `ni` shim behavior, non-execution
boundaries는 v0.6.0 release preparation으로 진행할 준비가 되어 있다. Audit
시점의 public v0.6.0 install retrieval과 hosted artifacts는 release-time
evidence gates였고 publication 이후 docs/140이 supersede한다. Windows
real-host execution, Homebrew availability, external user validation은 future
evidence gate로 남는다.

## Current-tree rename readiness

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| product name | Namba Intent | README, CLI help, docs가 current-tree surfaces에서 Namba Intent를 사용한다. | Yes | Publication proof는 이후 docs/140에 기록. |
| primary command | `namba-intent` | `cmd/namba-intent`가 있고 shared CLI logic에 delegate한다. | Yes | `go run ./cmd/namba-intent --help` passed. |
| cmd/namba-intent | primary entrypoint | `cmd/namba-intent/main.go`가 `internal/cli.Run`을 호출한다. | Yes | Duplicated command logic 없음. |
| deprecated ni shim | warn and delegate | `cmd/ni/main.go`가 `ni is deprecated; use namba-intent.`를 출력하고 delegate한다. | Yes | Transition-only. |
| internal CLI sharing | shared implementation | `internal/cli`가 command dispatch와 command-name option을 소유한다. | Yes | Shim은 compatibility output에만 `CommandName: "ni"`를 넘긴다. |
| .ni compatibility | preserve `.ni/` | Init/status/lock/session paths는 계속 `.ni/`를 사용한다. | Yes | `.namba-intent` directory 없음. |
| namba-ai distinction | do not claim `namba` | docs/135와 docs/136이 NambaAI `namba`와 Namba Intent `namba-intent`를 구분한다. | Yes | Repo는 `Nam-Cheol/ni` 유지. |
| no downstream execution | run compiles prompt only | README, skills, CLI prompt wording이 prompt-compiler-only boundary를 유지한다. | Yes | Shell/Codex adapter 없음. |

## Public install boundary

| Surface | Expected boundary | Observed state | Pass? | Required action |
| --- | --- | --- | --- | --- |
| README.md | audit 시점 current main과 latest public v0.5.1을 분리한다. | audit 시점에는 v0.6.0을 upcoming으로 두고 v0.5.1은 여전히 `ni`일 수 있다고 말했다. | Yes | v0.6.0 verification 뒤 docs/140이 supersede. |
| README.ko.md | Korean companion이 claim을 넓히면 안 된다. | README boundary와 aligned. | Yes | 계속 sync. |
| install.sh | Future primary binary는 `namba-intent`. | `namba-intent_<version>` assets를 선택하고 `namba-intent`를 install한다. | Yes | v0.6.0 전 public retrieval claim 금지. |
| install.ps1 | Future primary binary는 `namba-intent.exe`. | `namba-intent_<version>_windows_amd64.zip`를 선택하고 `%LOCALAPPDATA%\namba-intent\bin`에 install한다. | Yes | Real Windows transcript pending. |
| docs/22_INSTALL.md | Public v0.5.1과 current-main v0.6.0 paths를 분리한다. | Public `namba-intent` installer retrieval을 Release-gated로 update했다. | Yes | Publication 뒤 docs/140에서 reverified. |
| docs/install-curl.md | Public `namba-intent` retrieval claim 금지. | Historical v0.5.1 `ni` evidence와 future `namba-intent` assets를 분리했다. | Yes | Hosted assets 존재 뒤 reverify. |
| public latest release | audit 시점 latest public release는 v0.5.1일 수 있다. | `git tag --list v0.5.1`은 `v0.5.1`; `v0.6.0` tag absent. | Yes | Publication 뒤 docs/140이 supersede. |
| v0.6.0 release status | audit 시점 pre-publication. | Audit 중 local `v0.6.0` tag 없음. | Yes | Human approval 뒤 publication했고 docs/140에서 verified. |

## Installer readiness

| Installer | Expected future v0.6.0 behavior | Current-tree state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| install.sh | `namba-intent_<version>_<os>_<arch>` archive에서 `namba-intent` install. | Implemented. | Yes | Validation은 local fake release assets 사용. |
| install.ps1 | `namba-intent_<version>_windows_amd64.zip`에서 `namba-intent.exe` install. | Implemented. | Yes | macOS host에서는 static check only. |
| uninstall | Installer-managed primary binary와 PATH entry 제거. | Unix/PowerShell uninstall target이 `namba-intent` names다. | Yes | Public uninstall은 run하지 않음. |
| verification commands | `namba-intent --help`, `namba-intent version` 출력. | 양쪽 installer에 implemented. | Yes | README current-tree command와 일치. |
| legacy ni handling | Shim transition-only. | Source-tree shim exists; installer primary path는 `ni`에 의존하지 않는다. | Yes | Windows `ni` alias cleanup은 historical only. |
| Windows alias handling | PowerShell `ni`에 의존하지 않는다. | `install.ps1`는 alias cleanup이 `namba-intent.exe`에 필요 없다고 말한다. | Yes | Real-host Windows verification pending. |

## Release tooling readiness

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| version injection | `ni/internal/version.Version` injection. | `.goreleaser.yaml`이 `-X ni/internal/version.Version={{ .Version }}`를 사용한다. | Yes | Source `go run`은 `0.0.0-dev`. |
| asset names | Future artifacts use `namba-intent_...`. | `.goreleaser.yaml`, install scripts, release pipeline docs가 match한다. | Yes | Hosted assets는 아직 없음. |
| checksum names | Future checksum file은 `namba-intent_<version>_checksums.txt`. | Config와 installers가 match한다. | Yes | Published v0.6.0 checksum proof는 docs/140에 기록. |
| release-check | Check-only release gate. | tests, quality, smoke, install, docs, boundaries를 확인한다. | Yes | Publish하지 않음. |
| install-check | Source/build/temp install gate. | `namba-intent` command-name path와 installer tests를 확인한다. | Yes | Temp paths 사용. |
| GoReleaser config | Primary command plus shim. | `.goreleaser.yaml`이 `namba-intent`와 `ni` shim을 한 archive에 build한다. | Yes | GoReleaser publish not run. |
| GitHub workflow | Tag-only release workflow. | `.github/workflows/release.yml`은 `v*` tags에서 GoReleaser를 호출한다. | Yes | Workflow not run. |

## Docs and skills audit

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| README | Current-tree Namba Intent, public v0.5.1 distinction. | Bounded. | Yes | v0.6.0 publication claim 없음. |
| README.ko | Korean companion. | Bounded. | Yes | English보다 더 약속하지 않음. |
| current docs | current/future surfaces는 Namba Intent와 `namba-intent`. | docs/22, install-curl, docs/67, docs/120 updated. | Yes | Historical v0.5.x docs는 `ni` 유지. |
| historical docs | Actual past commands preserve. | v0.5.1 release evidence는 `ni`. | Yes | Intentional. |
| Claude skills | Skills are UX; CLI is authority. | Package README와 skill docs는 `namba-intent` authority를 사용하고 `ni-*` skill IDs를 유지한다. | Yes | Broad status remains Experimental. |
| Codex skills | Skills are UX; CLI is authority. | Package README와 `.agents/skills`가 `namba-intent` examples를 사용한다. | Yes | Skill은 CLI gates를 replace하지 않음. |
| .agents skills | Repo-local UX only. | `ni-start`, `ni-grill`, `ni-end`, `ni-run`이 CLI authority를 보존한다. | Yes | Downstream execution behavior 없음. |

## Current-tree evidence

| Evidence | Result | Notes |
| --- | --- | --- |
| docs/136 implementation | Pass | Rename implementation and validation record. |
| docs/137 smoke | Pass with notes | command-name help/version, init/status, repeated init, lockfile safety, shim delegation 기록. |
| go test | Pass | `GOCACHE=/private/tmp/ni-go-cache go test ./...`. |
| install docs check | Pass | `python3 scripts/check-install-docs.py`. |
| install ps1 check | Pass | `python3 scripts/check-install-ps1.py`. |
| quality | Pass | `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh`. |
| release-check | Pass | `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh`. |
| protected .ni diff | Pass | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` empty. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks v0.6.0 readiness? |
| --- | --- | --- | --- |
| v0.6.0 publication | 이 audit은 non-publishing. | Human-approved tag, release workflow, release metadata. | Superseded | 이후 docs/140에서 verified. |
| public install retrieval of namba-intent | Hosted v0.6.0 assets는 audit 시점 존재하지 않았다. | Published v0.6.0 assets에서 isolated install + help/version proof. | Superseded | 이후 docs/140에서 verified. |
| hosted artifacts | 이 audit에서는 release action 없음. | Asset inventory and checksum verification. | Superseded | 이후 docs/140에서 verified. |
| Windows real-host execution | macOS host가 증명할 수 없다. | Windows PowerShell install, new-session help/version, uninstall transcript. | No, Windows verified claim에는 필요. |
| Homebrew Available | Tap/formula/install proof 없음. | Tap/formula, checksums, audit, install, `namba-intent --help`, `namba-intent version`. | No, Homebrew remains Planned. |
| external user validation | External tester transcript 없음. | Tester transcript and comprehension review. | No, known note. |
| model workspace host behavior | Global host/provider behavior not verified. | Host-specific install/discovery/runtime proof. | No, Experimental preserved. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.6.0 publication status | Publication proof 전까지 release-gated로 유지. | Audit 시점 보존. | Yes | Tag/release action 이후 docs/140이 supersede. |
| public install | Proof 전 public `namba-intent` retrieval claim 금지. | docs/22와 curl docs가 Release-gated로 표시한다. | Yes | v0.6.0 proof는 docs/140에 기록. |
| Namba Intent identity | Current tree may use it. | README/help/docs에 사용. | Yes | Release claim은 future-gated. |
| namba-intent command | Primary current-tree command. | Implemented and validated. | Yes | Public install proof는 docs/140에 기록. |
| deprecated ni shim | Transition-only. | Warns and delegates. | Yes | Not primary path. |
| namba-ai distinction | Do not use `namba`. | docs/135와 docs/136에서 preserved. | Yes | Repo name은 `Nam-Cheol/ni`. |
| Homebrew | Planned / v0.5 candidate. | Preserved. | Yes | Available claim 없음. |
| Windows real-host verification | Pending. | Preserved. | Yes | Static checks only. |
| run behavior | Prompt compilation only. | Preserved. | Yes | Prompt execution 없음. |
| runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR/release automation, downstream execution layer. | Preserved. | Yes | Runtime behavior 추가 없음. |

## Git status / inclusion check

| Path or group | git status --short | Expected in v0.6.0? | Notes |
| --- | --- | --- | --- |
| README.md | clean | Yes | Bounded current-tree surface already tracked. |
| README.ko.md | clean | Yes | Korean companion already tracked. |
| cmd/namba-intent/* | clean | Yes | Primary command tracked. |
| cmd/ni/* | clean | Yes | Deprecated shim tracked. |
| internal/cli/* | clean | Yes | Shared implementation tracked. |
| install.sh | clean | Yes | Future `namba-intent` installer tracked. |
| install.ps1 | clean | Yes | Future Windows installer tracked. |
| docs/135* | clean | Yes | Rename plan tracked. |
| docs/136* | clean | Yes | Rename implementation tracked. |
| docs/137* | untracked before this task | Yes | First-user smoke evidence from prior current-tree smoke. |
| docs/138* | added in this task | Yes | This audit pair. |
| scripts/* | modified | Yes | Install docs checker updated for Release-gated current-main installer. |
| packages/* | modified | Yes | Claude package authority wording aligned. |
| .agents/* | clean | Yes | Existing repo-local skills remain. |
| .ni/contract.json | no diff | Yes | Protected. |
| .ni/session.json | no diff | Yes | Protected. |
| .ni/plan.lock.json | no diff | Yes | Protected. |
| unexpected files | docs/137 untracked 외 없음 | No | Generated prompts 실행하지 않음. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; docs/137은 이 task 전 untracked였고 docs/138 및 bounded docs/checker edits가 추가됨. |
| `git log --oneline --decorate -20` | Passed; HEAD `80acd80 Implement Namba Intent rename`. |
| `git tag --list v0.5.1` | Passed; `v0.5.1` exists. |
| `git tag --list v0.6.0` | Passed; empty. |
| `git rev-parse v0.5.1` | Passed; `b588f6b2e13111841081d186bd0e70d3c0bfbd6c`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/namba-intent --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/namba-intent version` | Passed; source output `0.0.0-dev`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; deprecation warning verified. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff --check` | Passed. |

## Changes made

- `docs/22_INSTALL.md`: verified v0.5.1 `ni` evidence와 Release-gated current-main
  `namba-intent` installer behavior를 분리했다.
- `docs/install-curl.md`, `docs/install-curl.ko.md`: 같은 public install boundary를
  추가했다.
- `docs/120_GLOBAL_INSTALL_ACCEPTANCE.md`, Korean companion: primary global install
  acceptance command를 `namba-intent`로 맞췄다.
- `docs/67_RELEASE_PIPELINE.md`, Korean companion: actual GoReleaser primary
  binary plus shim configuration과 맞췄다.
- `packages/claude-skills/README.md`: CLI authority를 `namba-intent`로 명명했다.
- `scripts/check-install-docs.py`: Release-gated current-main installer boundary를
  enforce하도록 update했다.
- 이 audit와 Korean companion을 추가했다.

## What this audit proves

State only:
- current-tree rename readiness is ready with notes;
- public install docs are bounded after this audit;
- next release preparation can proceed;
- no release action was performed.

## What this audit does not prove

State:
- v0.6.0 has been published;
- public install retrieves namba-intent;
- hosted v0.6.0 artifacts exist;
- Windows real-host execution works;
- Homebrew is Available;
- external users succeed;
- downstream execution succeeds.

## Recommended next task

A. v0.6.0 release notes draft

Selection rationale: readiness는 notes와 함께 통과했고 남은 proof gaps는 expected
release-time 또는 external verification gates다. Artifact dry-run이나 publication
전에 release notes draft가 가장 안전한 next preparation artifact다.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: draft v0.6.0 Namba Intent release notes without publishing.

Use docs/135_NAMBA_INTENT_RENAME_PLAN.md,
docs/136_NAMBA_INTENT_RENAME_IMPLEMENTATION.md,
docs/137_NAMBA_INTENT_FIRST_USER_SMOKE.md, and
docs/138_V0_6_0_PUBLIC_INSTALL_PARITY_AND_RELEASE_READINESS.md as source
evidence.

Goal:
Create v0.6.0 release notes that explain the Namba Intent rename, primary
`namba-intent` command, transition-only `ni` shim, preserved `.ni/`
compatibility, installer and release asset rename, and prompt-compiler-only
runtime boundary.

Required boundaries:
- Do not publish, tag, create a GitHub release, upload assets, run release
  workflows, run GoReleaser publish, create or publish Homebrew formula, mark
  Homebrew Available, claim Windows real-host verification, run ni end or
  namba-intent end on the project root, relock the project root, edit
  .ni/plan.lock.json, execute generated prompts, or add downstream execution
  behavior.
- Post-release work에서는 docs/140을 publication 및 public install verification
  record로 사용한다.
- Preserve Homebrew: Planned / v0.5 candidate, Model workspace packs:
  Experimental, No-terminal method: Experimental / assisted, and Skills are
  UX; CLI is authority.

Deliverables:
- Add docs/139_V0_6_0_RELEASE_NOTES_DRAFT.md.
- Add docs/139_V0_6_0_RELEASE_NOTES_DRAFT.ko.md.
- Include sections for summary, included changes, migration notes, install
  boundary, validation evidence, known deferrals, what this does not prove,
  and recommended next task.

Validation:
- git status --short
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
- git diff --check

Final response:
Report changed files, release-note decision, public install boundary, validation
results, protected .ni diff result, confirmation that no publication/tag/release
action occurred, and the selected next task.
```
