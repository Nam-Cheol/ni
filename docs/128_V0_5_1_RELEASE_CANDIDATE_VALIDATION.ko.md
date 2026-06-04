# v0.5.1 Release Candidate Validation

## Current status

State:
- v0.5.0 publication: verified.
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES.
- current-tree first-user smoke after TUI: FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES.
- v0.5.1 release: not published.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: macOS-only development host에서는 deferred.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

이 문서는 current-tree가 v0.5.1 release candidate가 될 준비가 되었는지 검증한다.
하지만 publish, tag, asset upload, GitHub Release creation, release workflow 실행,
Homebrew formula 생성 또는 publish, project root `ni end`, project root relock,
generated prompt execution, runtime execution behavior 추가는 수행하지 않는다.

## Decision

V0_5_1_RC_VALIDATION_PASS_WITH_NOTES

Justification: audited current tree는 v0.5.1 install-parity patch scope를 cover하고
local validation gates가 passed했다. Notes는 남는다. v0.5.1은 아직 published되지
않았고, release artifacts는 아직 없으며, `install.sh`는 아직 v0.5.1을 retrieve할
수 없고, Windows real-host execution은 deferred이며, Homebrew는 Planned / v0.5
candidate이고, external user validation은 수행되지 않았다.

## Patch scope validation

| Scope item | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni init .` positional target | Current-directory init works. | `cmd/ni/main.go` help와 temp binary smoke가 `ni init . --yes`를 accept했다. | Yes | Published v0.5.0에는 이 path가 없다. |
| Bubble Tea v2 TUI | Interactive guided init이 Bubble Tea v2를 사용한다. | `internal/tui/init/model.go`가 `charm.land/bubbletea/v2`를 import한다. | Yes | Safe한 경우에만 interactive launch. |
| Lip Gloss v2 styling | TUI view가 Lip Gloss v2 styling을 사용한다. | `internal/tui/init/model.go`가 `charm.land/lipgloss/v2`를 import하고 styled views를 render한다. | Yes | Styling은 render-only다. |
| non-interactive fallback | CI/non-TTY에서도 동작한다. | `ni init . --yes`가 temporary project에서 passed. | Yes | Automation은 TUI를 요구하지 않는다. |
| domain/render separation | File writing은 TUI rendering 밖에 있어야 한다. | `internal/core/docstore`가 init file planning/writes를 맡고 `internal/tui/init`은 model/view만 맡는다. | Yes | Deterministic tests를 지원한다. |
| existing file protection | Init은 planning files를 silently overwrite하면 안 된다. | Repeated `ni init . --yes`가 existing/unchanged files와 `created files: none`을 report했다. | Yes | Additive-only behavior. |
| lockfile protection | Init은 `.ni/plan.lock.json`을 수정하면 안 된다. | Sentinel lockfile SHA-256이 `1ede15b645ba3dab7c56783753a6420477a7ec58847c3e80acb2d0fa792d5960`로 유지됐다. | Yes | Command가 lockfile warning과 함께 stopped. |
| post-TUI plain summary | Guided setup 뒤 plain text summary가 있어야 한다. | `cmd/ni/main.go`가 plain init summaries를 출력하고 tests가 summary text를 assert한다. | Yes | TUI 밖의 output이다. |
| README onboarding | macOS/Windows first paths가 prominent하고 bounded해야 한다. | README와 README.ko가 two primary install paths와 `ni init .`을 보여준다. | Yes | Public v0.5.0 parity note 보존. |
| install docs/checkers | Docs와 validators가 current behavior와 align해야 한다. | `python3 scripts/check-install-docs.py`와 `python3 scripts/check-install-ps1.py` passed. | Yes | Windows checker는 static safety only. |
| no downstream execution | Patch는 `ni run`을 downstream execution으로 만들면 안 된다. | README/docs/help가 `ni run`을 bounded prompt compilation으로 유지한다. | Yes | 이 validation은 runtime execution behavior를 추가하지 않았다. |

## Repository state

| Surface | Observed state | Notes |
| --- | --- | --- |
| git status | `README.md`, `README.ko.md` modified; docs/126, docs/127, docs/128 untracked after this task. | Staging/commit 없음. |
| v0.5.0 tag | Present; `git rev-parse v0.5.0` returned `b8fec7fa9615a861acf4eba59733c800c70c6cca`. | Public mismatch baseline. |
| v0.5.1 tag | Absent. | RC validation에는 올바른 상태; tag 생성하지 않음. |
| diff `v0.5.0..HEAD` | docs/128 전 기준 60 files changed, 7302 insertions, 476 deletions. | Current-tree init, installer, docs, checker work 포함. |
| protected `.ni` diff | Empty. | Project root `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` unchanged. |
| generated artifacts | Repo 안에는 unexpected artifact 없음. | Temporary binary/project는 `/private/tmp`에 있음. |

docs/126과 docs/127은 현재 untracked release-planning evidence다. v0.5.1 RC
planning bundle에는 expected이지만, 이 task에서는 stage하지 않았다.

## Current-tree first-user validation

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| command-name `ni --help` | `PATH=/private/tmp/ni-v051-rc-bin.3tTS9H:... zsh -f -c 'ni --help'` | PATH-resolved command가 동작하고 `ni init [.]`를 list한다. | Passed. | Yes | `command -v ni`는 temporary binary를 가리켰다. |
| command-name `ni version` | `PATH=/private/tmp/ni-v051-rc-bin.3tTS9H:... zsh -f -c 'ni version'` | Source/current-tree binary가 version을 출력한다. | `0.0.0-dev`. | Yes | Release build는 `0.5.1`을 inject해야 한다. |
| `ni init . --yes` | Temp project `ni init . --yes`. | Planning workspace를 만든다. | `.ni/contract.json`, `.ni/session.json`, `docs/plan/**` created; `.ni/plan.lock.json` not created. | Yes | Downstream prompt execution 없음. |
| `ni status --proof --next-questions` | Temp project status. | Actual first-run readiness를 report한다. | `NI Intent Readiness: BLOCKED` with R014, OQ-001, R015, R016. | Yes | Empty first-run scaffold에서는 BLOCKED가 expected. |
| repeated init | Temp project `ni init . --yes` again. | Silently overwrite하지 않는다. | Existing/unchanged files와 `created files: none`. | Yes | Existing planning docs preserved. |
| lockfile safety | Sentinel `.ni/plan.lock.json` 생성 후 `ni init . --yes`. | Init은 lockfile을 수정하면 안 된다. | Warning printed; no files changed; sentinel hash unchanged. | Yes | Repo root가 아니라 temporary project에서 수행. |

## TUI model coverage

| Behavior | Evidence | Pass? | Notes |
| --- | --- | --- | --- |
| model | `internal/tui/init/model.go`; `internal/tui/init/model_test.go`. | Yes | Model state directly unit-tested. |
| Init | `TestModelInitialStateUsesAltScreen`. | Yes | Initial state exposes alt-screen view. |
| Update | `updateForTest`가 key messages를 model에 전달한다. | Yes | Navigation/confirmation tests로 cover. |
| View | `m.View()` rendered content checked. | Yes | View output은 readiness authority가 아니다. |
| AltScreen | `view.AltScreen` asserted in tests. | Yes | TUI behavior only. |
| Lip Gloss v2 | `charm.land/lipgloss/v2` import와 styled view code. | Yes | Code-level evidence plus package tests. |
| Up/Down | Key tests cover `tea.KeyUp` and `tea.KeyDown`. | Yes | Cursor movement covered. |
| Left/Right | Key tests cover `tea.KeyLeft` and `tea.KeyRight`. | Yes | Horizontal movement covered. |
| Enter | Key tests cover `tea.KeyEnter`. | Yes | Select/confirm paths covered. |
| Esc | Key tests cover `tea.KeyEsc`. | Yes | Back/cancel path covered. |
| q | Key tests cover `tea.KeyRunes` with `q`. | Yes | Cancel path covered. |
| cancel | `TestModelCanCancelAtConfirm` and `q` path. | Yes | Canceled result asserted. |
| confirm | `TestModelCanConfirmIntent`. | Yes | Confirmed result asserted. |
| post-TUI summary | `cmd/ni/main_test.go` summary assertions and `confirmGuidedIntent` summary. | Yes | Plain text summary exists outside TUI rendering. |

## README/install docs audit

| Surface | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| README.md | macOS/Windows primary paths, Homebrew not Available, `ni run` compile-only, READY not product readiness, TUI not authority. | Matches; v0.5.0 parity note와 status list 포함. | Yes | `ni init .`은 guided project intent setup. |
| README.ko.md | Korean companion은 English보다 더 많이 promise하면 안 된다. | English boundaries와 match. | Yes | Required control words preserved. |
| docs/22_INSTALL.md | Install behavior와 uninstall guidance aligned. | Check passed. | Yes | Windows real-host execution remains deferred. |
| docs/install-curl.md | Curl installer guidance는 available release assets와 command-name verification으로 bounded. | Check passed. | Yes | Release 전 v0.5.1 retrieval claim 없음. |
| docs/install-curl.ko.md | Korean companion bounded like English. | Check passed. | Yes | Homebrew는 curl path 밖에 있음. |

## Version/release gate

| Gate | Required evidence before release | Current state | Blocks RC? | Notes |
| --- | --- | --- | --- | --- |
| release version injection | Release binary가 `0.5.1`을 report해야 한다. | `.goreleaser.yaml`이 `ni/internal/version.Version={{ .Version }}`를 inject한다. | No | Artifact dry-run/release build에서 verify 필요. |
| v0.5.1 tag | Intended release commit에 tag 존재. | Absent. | No | Pre-publication RC validation에는 맞다. |
| artifact build | v0.5.1 archives/checksums generated. | Not run. | No | 다음 uncertainty. |
| artifact version output | Extracted artifact `ni version`이 `0.5.1` report. | Not available yet. | No | Publication 전 required. |
| `install.sh` v0.5.1 retrieval | Release 후 installer가 v0.5.1을 retrieve. | Publication 전에는 impossible. | No | 지금 claim하면 안 된다. |
| curl installer isolated install | Temporary install에서 `ni --help`, `ni version` pass. | Current-tree binary smoke only; hosted v0.5.1 install 없음. | No | Artifact 존재 후 required. |
| checksum verification | `ni_0.5.1_checksums.txt`가 artifacts와 match. | Not available yet. | No | Artifact dry-run에서 verify해야 한다. |

## v0.5.1 release notes delta

Draft delta:
- Fix: public install parity mismatch from v0.5.0 where installed `ni --help`
  and `ni version` passed but `ni init .` failed.
- Add: `ni init .` positional target support.
- Add: Bubble Tea v2 / Lip Gloss v2 guided init TUI.
- Add: non-interactive fallback preservation for CI and non-TTY contexts.
- Add: existing file and `.ni/plan.lock.json` protection during init.
- Improve: README macOS / Windows two-path onboarding.
- Improve: install docs and checker alignment for command-name verification.
- Does not add downstream execution.
- Does not make Homebrew Available.
- Does not verify Windows real-host execution.
- Does not make no-terminal deterministic.

이 delta는 draft-only다. Actual v0.5.1 artifacts와 installer path가 verified되기 전
published release notes로 복사하면 안 된다.

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks v0.5.1 RC? |
| --- | --- | --- | --- |
| Windows real-host execution | Current development host가 macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | Published tap/formula/install proof 없음. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | 이 task에는 external user transcript 없음. | User 또는 separate-host install/init transcript. | No |
| model workspace host behavior | Host-level/global install과 provider behavior unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Target workspace에 대한 trusted CLI transcript. | No |
| README raster images | Raster image generation은 out of scope. | 필요하면 separate asset task. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Preserved. | Yes | Release action 없음. |
| published v0.5.0 behavior | Installed help/version pass와 `ni init .` failure를 구분해야 한다. | README/docs/126/docs/127에 preserved. | Yes | Public mismatch가 v0.5.1 이유. |
| current-tree behavior | Local current-tree support만 claim할 수 있다. | Current-tree binary smoke passed. | Yes | Hosted artifact claim 아님. |
| Homebrew | Planned / v0.5 candidate 유지. | Preserved. | Yes | Homebrew Available claim 없음. |
| Windows real-host execution | Transcript 없으면 deferred. | Preserved. | Yes | Static installer checks only. |
| `ni init .` | Guided setup only. | Planning docs와 `.ni` skeleton을 만든다. | Yes | Agents 실행 없음. |
| `ni run` | Bounded prompt compilation only. | README/help/docs preserve compile-only wording. | Yes | Downstream work execute하면 안 된다. |
| READY | CLI readiness only, not product readiness. | Root `ni status` reports READY; docs keep CLI authority. | Yes | Model judgment로 readiness 선언하지 않음. |
| TUI readiness boundary | TUI는 intent를 모을 수 있지만 readiness를 결정하지 않는다. | TUI model은 readiness result를 쓰지 않는다. | Yes | `ni status`가 gate. |
| runtime execution boundary | `ni-kernel` must not include task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | Runtime behavior 추가 없음. |

## Git status / inclusion check

| Path or group | `git status --short` | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | `M README.md` | Yes | Public parity note and onboarding. |
| README.ko.md | `M README.ko.md` | Yes | Korean companion. |
| cmd/ni/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Positional init and tests. |
| internal/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Docstore/TUI support. |
| go.mod | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Bubble Tea v2 / Lip Gloss v2 dependencies. |
| go.sum | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Dependency checksums. |
| install.sh | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Current install path handling. |
| install.ps1 | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Windows User PATH installer. |
| docs/124* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Init TUI evidence. |
| docs/125* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | First-user smoke evidence. |
| docs/126* | `?? docs/126_*` | Yes | Public v0.5.0 mismatch evidence. |
| docs/127* | `?? docs/127_*` | Yes | v0.5.1 patch plan. |
| docs/128* | `?? docs/128_*` | Yes | 이 RC validation과 English companion. |
| scripts/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Install/release/checker alignment. |
| `.ni/contract.json` | no diff | No direct edit | Protected. |
| `.ni/session.json` | no diff | No direct edit | Protected. |
| `.ni/plan.lock.json` | no diff | No direct edit | Protected. |
| unexpected files | none observed in repo status | No | Temporary files are outside repo. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; expected modified/untracked docs only. |
| `git log --oneline --decorate -20` | Passed; post-v0.5.0 onboarding/install commits visible. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; v0.5.1 tag 없음. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Passed; candidate files listed. |
| `git diff --stat v0.5.0..HEAD` | Passed; docs/128 전 기준 60 files, 7302 insertions, 476 deletions. |
| required ripgrep scans | Passed; RC validation을 block하는 scope-expanding runtime/release claim 없음. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root가 blockers, deferrals, warnings 없이 `NI Intent Readiness: READY` report. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help includes `ni init [.]`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build reports `0.0.0-dev`. |
| temporary current-tree binary build | Passed; `/private/tmp/ni-v051-rc-bin.3tTS9H/ni`. |
| command-name temp `ni --help` | Passed. |
| command-name temp `ni version` | Passed; `0.0.0-dev`. |
| temp `ni init . --yes` | Passed. |
| temp `ni status --proof --next-questions` | Passed; blank first-run intent에 대해 command가 correctly BLOCKED report. |
| repeated temp `ni init . --yes` | Passed; no overwrites. |
| temp lockfile safety | Passed; sentinel hash unchanged. |
| TUI model behavior tests | `go test ./...`를 통해 `ni/internal/tui/init` 포함 passed. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | docs/128 addition 후 Passed. |
| `bash scripts/demo-check.sh` | docs/128 addition 후 Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | docs/128 addition 후 Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | docs/128 addition 후 Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | docs/128 addition 후 Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | docs/128 addition 후 Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

이 RC validation doc task는 Go files를 수정하지 않았으므로 final edited set에는
`gofmt -w .`가 필요하지 않았다.

## Changes made

| File | Why |
| --- | --- |
| `docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.md` | English RC validation, decision, evidence, release notes delta, next task prompt 추가. |
| `docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.ko.md` | 같은 boundaries의 Korean companion 추가. |
| `docs/51_POST_RELEASE_ROADMAP.md` | docs/128로 가는 narrow pointer 추가. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Matching Korean pointer 추가. |

## What this validation proves

- Current tree는 audited criteria 기준 v0.5.1 RC로 보인다.
- First-user current-tree path는 notes와 함께 works: blank scaffold는 real intent가
  공급될 때까지 correctly BLOCKED다.
- Planned v0.5.1 patch scope는 current tree에 covered.
- Release gates are explicit and still pending for actual artifacts.
- Release action은 수행하지 않았다.

## What this validation does not prove

- v0.5.1 has been published.
- v0.5.1 artifacts exist.
- `install.sh` retrieves v0.5.1.
- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

Selected next task: A. v0.5.1 artifact dry-run.

Selection rationale: RC validation은 notes와 함께 passes이고, 다음 uncertainty는
publication 전에 release artifact path가 `0.5.1`을 inject하고 consistent checksums를
만들며 install parity를 보존하는지다.

## Next executable Codex prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Goal:
Run a v0.5.1 artifact dry-run for the current release-candidate tree. This is
pre-publication validation only.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- docs/127_V0_5_1_PATCH_RELEASE_PLAN.md
- docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.md
- docs/22_INSTALL.md
- docs/install-curl.md
- install.sh
- install.ps1
- .goreleaser.yaml
- .github/workflows/release.yml
- cmd/ni/
- internal/
- scripts/release-check.sh

Required checks:
- git status --short
- git tag --list v0.5.1
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- Check whether goreleaser is installed.
- If goreleaser is installed, run a local snapshot/dry-run artifact build only,
  inspect generated archive names, checksums, and extracted current-platform
  `ni --help` / `ni version` output.
- If goreleaser is not installed, do not install it automatically; document the
  unavailable command and run the closest local `go build` linker-flag check
  that proves `ni version` can report `0.5.1`.
- Build a temporary current-platform binary outside the repo with
  `-ldflags "-X ni/internal/version.Version=0.5.1"` and verify command-name
  `ni --help`, `ni version`, `ni init . --yes`, repeated init, and lockfile
  protection in a temporary project.

Decision:
Use exactly one:
- V0_5_1_ARTIFACT_DRY_RUN_PASS
- V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES
- V0_5_1_ARTIFACT_DRY_RUN_BLOCKED

Rules:
- Do not publish, tag, create a GitHub Release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  ni end on the project root, relock the project root, execute generated
  prompts, add runtime execution behavior, or mark v0.5.1 as released.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution verified unless a Windows transcript
  exists.
- Do not claim `install.sh` retrieves v0.5.1 unless a real hosted v0.5.1 asset
  exists and the installer retrieves it.
- Keep Skills are UX; CLI is authority.
- Keep `ni run` as bounded prompt compilation only.

Final response:
Report changed files, artifact dry-run decision, version-injection result,
checksum/archive result if available, current-platform binary smoke, known
deferrals, validation results, protected .ni diff, and confirmation that no
publish/tag/release/upload/project-root relock/generated prompt execution
occurred.
```
