# v0.5.1 Patch Release Plan

## Current status

State:
- v0.5.0 publication: verified.
- v0.5.0 release binary: Available.
- v0.5.0 curl installer: Available.
- published v0.5.0 install lane: `ni --help`와 `ni version` passed.
- published v0.5.0 `ni version`: `0.5.0`.
- published v0.5.0 `ni init .`: `unknown init option: .`로 failed.
- docs/126의 public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- current-tree `ni init .` Bubble Tea v2 / Lip Gloss v2 TUI: implemented.
- current-tree first-user onboarding smoke after TUI: FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES.
- Windows real-host execution: macOS-only development host에서는 deferred.
- Homebrew: Planned / v0.5 candidate.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- `ni run`은 bounded handoff prompt를 compile하며 downstream work를 실행하지 않는다.

이 문서는 release plan only이다. Publish, tag, GitHub Release creation, asset
upload, release workflow 실행, GoReleaser publish, Homebrew formula 생성 또는
publish, project root `ni end`, project root relock, generated prompt execution,
runtime execution behavior 추가를 수행하지 않는다.

## Decision

V0_5_1_PATCH_PLAN_READY_WITH_NOTES

Justification: patch scope는 명확하고 current-tree validation으로 public v0.5.0에
없는 behavior를 증명할 수 있다. 다만 Windows real-host execution, Homebrew
availability, external user validation, post-publication hosted artifact parity는
각각의 transcript가 생길 때까지 deferred로 남기므로 notes가 붙는다.

## Release rationale

현재 public onboarding docs는 `ni init .`을 포함한 first-success path를 안내한다.
Task 205는 published v0.5.0이 install 후 `ni --help`와 `ni version`은 통과하지만,
current-directory init step에서 `unknown init option: .`로 실패한다는 것을
보였다. `v0.5.0` tag 이후 current tree에는 positional `ni init .`, guided TUI,
non-interactive fallback, install-path handling이 추가되었다.

따라서 적절한 release action은 이미 구현된 onboarding/install-parity fixes를
v0.5.1 patch로 package하는 것이다. 이 patch는 execution layer나 Homebrew
availability release가 되면 안 된다.

## Patch scope

Required v0.5.1 contents:

- `ni init .` positional target support.
- Bubble Tea v2 / Lip Gloss v2 interactive init TUI.
- CI와 non-TTY context를 위한 non-interactive fallback.
- TUI rendering과 분리된 domain init logic.
- Existing file protection; init은 planning files를 silently overwrite하지 않는다.
- `.ni/plan.lock.json` protection; init은 existing lockfile을 수정하지 않는다.
- Post-TUI plain text summary.
- README macOS / Windows two-path onboarding.
- `ni --help`와 `ni version` command-name verification에 맞춘 install docs.
- Install docs, Windows installer static safety, README surface, demo, smoke,
  install, release checks의 current checker alignment.
- No downstream execution behavior.

Optional current-tree contents that may be included because they are already
present and bounded:

- Accidental duplicate `examples/namba-ai-upgrade/docs/plan/* 2.md` cleanup.
- Deterministic README SVG assets and asset drift checks.
- `install.sh`의 macOS/Linux global install path handling updates.
- `install.sh --update-path`와 `install.sh --uninstall` improvements.
- `install.ps1` User PATH safety and uninstall behavior.

Explicit exclusions:

The patch must not include these kernel-boundary items:

- Task runner, SPEC runner, execution harness, shell adapter, Codex exec
  adapter, queue, PR automation, `ni-kernel` 내부 release automation, downstream
  execution layer.
- Homebrew Available claim, tap publication, formula publication.
- Windows transcript 없는 Windows real-host verified claim.
- Model workspace packs의 broad Available status.
- No-terminal deterministic validation claim.
- `ni run`이 downstream work를 execute한다는 claim.

## Current-tree comparison

`v0.5.0` 이후 v0.5.1 patch scope에 속하는 commits:

| Commit | Patch에서의 역할 |
| --- | --- |
| `34a613a` Simplify README install paths and add init TUI onboarding | Two-path README와 initial `ni init .` onboarding implementation. |
| `0745621` Implement global install path handling | Command-name verification에 필요한 macOS/Linux installer PATH handling. |
| `9e587d8` Update README install docs for v0.5.0 release | Current public path의 release/install docs baseline. |
| `c943a53` Add README visual asset pass references | Deterministic README visual support; optional이지만 current tree에 있음. |
| `4ae4a85` Redesign README and add Bubble Tea init TUI | Main TUI, domain separation, tests, duplicate cleanup. |
| `93dc241` Document first-user onboarding smoke after TUI | Current-tree onboarding smoke evidence. |

Release-planning bundle에 속하는 current uncommitted documentation:

| Path | Role |
| --- | --- |
| `README.md` | docs/126으로 연결되는 public parity note. |
| `README.ko.md` | Korean companion parity note. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md` | Public v0.5.0 mismatch evidence. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.ko.md` | Korean companion mismatch evidence. |
| `docs/127_V0_5_1_PATCH_RELEASE_PLAN.md` | English release plan. |
| `docs/127_V0_5_1_PATCH_RELEASE_PLAN.ko.md` | 이 Korean companion release plan. |

Patch를 support하는 current-tree surfaces:

| Surface | Evidence |
| --- | --- |
| CLI init routing | `cmd/ni/main.go`는 `ni init [.]`을 accept하고 safe할 때 interactive positional init을 TUI로 route한다. |
| Init domain logic | `internal/core/docstore`는 file plan을 만들고, existing files를 skip하며, `.ni/plan.lock.json`을 protect한다. |
| Init TUI | `internal/tui/init`은 Bubble Tea v2와 Lip Gloss v2를 사용해 model/view를 render한다. |
| Tests | `cmd/ni/main_test.go`, `internal/core/docstore/docstore_test.go`, `internal/tui/init/model_test.go`가 positional init, fallback, existing files, lockfile safety, TUI model behavior를 cover한다. |
| Installers | `install.sh`는 `--update-path` / `--uninstall`을 support하고, `install.ps1`은 User PATH만 update한다. |
| Release versioning | `.goreleaser.yaml`은 release build에 `ni/internal/version.Version={{ .Version }}`를 inject한다. |
| Release workflow | `.github/workflows/release.yml`은 `v*` tags에서 tests, quality, release-check, GoReleaser를 run한다. |

v0.5.1 scope를 넓히면 안 되는 surfaces:

| Surface | Treatment |
| --- | --- |
| Homebrew draft formula | Planned / candidate evidence로만 유지; Available claim 또는 tap publication 없음. |
| Windows installer real execution | Real Windows install/new-session/uninstall transcript 전까지 deferred. |
| Model workspace packs | Experimental 및 CLI-authority boundary 유지. |
| No-terminal method | Assisted only; deterministic validation claim 없음. |
| Downstream seed material | Derived/mutable documentation일 수 있으나 kernel-owned execution state가 아니다. |

## Release criteria

v0.5.1 publication 전 모든 required criteria가 pass해야 한다:

| Criterion | Required result |
| --- | --- |
| `go test ./...` | Current tree에서 pass. |
| Project root `ni status --proof --next-questions` | Actual current state를 report; readiness wording은 CLI output만 근거가 된다. |
| Local `ni --help` | Works and includes `ni init [.]`. |
| Local `ni version` | Works; source/local build는 linker flags 없으면 `0.0.0-dev`일 수 있다. |
| First-user current-tree smoke | docs/125 또는 더 최신 equivalent 기준 pass with notes. |
| Temporary `ni init . --yes` | `.ni/contract.json`, `.ni/session.json`, `docs/plan/**` 생성. |
| Repeated `ni init . --yes` | Existing planning files를 silently overwrite하지 않음. |
| Lockfile safety | Existing `.ni/plan.lock.json`이 init으로 수정되지 않음. |
| Install docs checks | `scripts/check-install-docs.py` pass. |
| Windows static installer checks | `scripts/check-install-ps1.py` pass; real-host proof 아님. |
| Release check | `scripts/release-check.sh` pass. |
| Protected `.ni` files | 명시적으로 authorized amendment flow가 아닌 한 project-root `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` unchanged. |
| Release notes delta | v0.5.1 patch delta를 정확히 말하고 overclaim 없음. |
| Homebrew wording | Homebrew Available claim 없음. |
| Windows wording | Transcript 없는 Windows real-host verified claim 없음. |
| `ni run` wording | Prompt compilation only; downstream execution claim 없음. |

## Publication readiness gates

Maintainer는 actual v0.5.1 publication 전에 다음 gate를 완료해야 한다:

1. Release commit이 approved patch scope만 포함하는지 확인한다.
2. Project release process가 요구하면 release version을 bump 또는 set한다.
3. Release build linker flags로 `ni version`이 `0.5.1`을 report하는지 확인한다.
4. Full validation: `gofmt -w .`, `go test ./...`, `bash scripts/quality.sh`.
5. `bash scripts/release-check.sh`.
6. Public install parity 중심 release notes delta 작성.
7. Tooling이 available하면 artifact dry-run 실행.
8. Generated archives checksums 확인.
9. `install.sh`가 stale v0.5.0 artifacts가 아니라 v0.5.1을 retrieve하는지 확인.
10. Current-platform release binary가 command-name `ni --help`, `ni version`을
    통과하는지 확인.
11. Temporary `HOME` / `BINDIR`에서 isolated curl installer install for v0.5.1.
12. Installed v0.5.1 binary로 isolated `ni init .`과
    `ni status --proof --next-questions` 확인.
13. Public install instructions를 honest하게 유지하기 위해 필요한 경우에만 README update.
14. Windows transcript가 없으면 Windows real-host verification은 deferred 유지.
15. Separate task에서 full tap/formula/checksum/audit/install/help/version/uninstall
    proof가 pass하기 전까지 Homebrew Planned / v0.5 candidate 유지.

## Validation matrix

| Phase | Command or check | Expected |
| --- | --- | --- |
| planning task | `git status --short` | Expected docs/README planning changes only; staging 없음. |
| planning task | `git log --oneline --decorate -20` | Post-`v0.5.0` onboarding/install commits visible. |
| planning task | `git diff --name-only v0.5.0..HEAD` | Candidate patch files identified. |
| planning task | `git diff --stat v0.5.0..HEAD` | Patch size and surfaces confirmed. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Actual project-root status report. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass with source version, normally `0.0.0-dev`. |
| planning task | `python3 scripts/check-install-docs.py` | Pass. |
| planning task | `python3 scripts/check-install-ps1.py` | Pass as static Windows safety only. |
| planning task | `bash scripts/check-skill-packs.sh` | Pass; CLI authority wording preserved. |
| planning task | `bash scripts/demo-check.sh` | Pass. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass. |
| planning task | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Empty. |
| pre-publication | Release build `ni version` | Reports `0.5.1`. |
| pre-publication | Artifact dry-run and checksums | Archives and `ni_0.5.1_checksums.txt` internally consistent. |
| post-publication | Isolated installed `ni --help` / `ni version` | Command-name checks pass; version reports `0.5.1`. |
| post-publication | Isolated installed `ni init .` | Positional init works with release binary. |
| post-publication | Isolated installed `ni status --proof --next-questions` | Runs after init and reports actual readiness state. |

## Rollback or docs correction path

v0.5.1 publication이 approve되지 않거나 validation이 fail하면:

- Publish 또는 tag하지 않는다.
- docs/126을 public v0.5.0 parity mismatch record로 유지한다.
- v0.5.0에 positional `ni init .`이 없다는 README parity wording을 유지하거나,
  README first-project path가 더 이상 v0.5.0 parity를 imply하지 않도록 조정한다.
- Failed gate와 exact command output을 follow-up audit에 기록한다.
- Blocker 해결 후에만 이 plan을 다시 실행한다.

v0.5.1이 published 되었지만 hosted parity가 fail하면:

- Public install parity for `ni init .` claim을 중단한다.
- Broad install claim을 넓히기 전에 post-publication mismatch note를 추가한다.
- Kernel boundary를 숨기거나 acceptance criteria를 약화하지 말고 docs correction을 우선한다.

## Protected files

이 planning task는 다음을 edit, relock, silently regenerate하면 안 된다:

- `.ni/contract.json`
- `.ni/session.json`
- `.ni/plan.lock.json`

Explicitly authorized amendment 또는 relock flow 밖에서 protected file diff가
나오면 release-plan decision은 `V0_5_1_PATCH_PLAN_BLOCKED`가 된다.

## Next executable prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Goal:
Run the authorized v0.5.1 release preflight for public install parity. This is
still pre-publication unless the user separately authorizes tagging and release.

Read:
- AGENTS.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- docs/127_V0_5_1_PATCH_RELEASE_PLAN.md
- README.md
- README.ko.md
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
- git diff --name-only v0.5.0..HEAD
- git diff --stat v0.5.0..HEAD
- gofmt -w .
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- temporary current-tree binary smoke: ni init . --yes, repeated init, lockfile protection
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Release-preflight output:
- exact release delta for v0.5.1
- validation transcript summary
- release notes draft or update if needed
- artifact dry-run result if GoReleaser is available
- explicit decision: V0_5_1_PREFLIGHT_READY,
  V0_5_1_PREFLIGHT_READY_WITH_NOTES, or V0_5_1_PREFLIGHT_BLOCKED

Rules:
- Do not publish, tag, create a GitHub Release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  ni end on the project root, relock the project root, execute generated
  prompts, or add runtime execution behavior unless the user separately
  authorizes that exact release action.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution verified unless a Windows transcript
  exists.
- Keep Skills are UX; CLI is authority.
- Keep ni run as bounded prompt compilation only.

Final response:
Report changed files, preflight decision, validation results, protected .ni
diff, remaining notes, and confirmation that no publish/tag/release/upload/root
relock/generated prompt execution occurred.
```

## Validation run in this task

| Command | Result |
| --- | --- |
| `git status --short` | Validation 후 recheck 예정. |
| `git log --oneline --decorate -20` | Checked; post-`v0.5.0` onboarding/install commits visible. |
| `git diff --name-only v0.5.0..HEAD` | Checked; candidate patch files identified. |
| `git diff --stat v0.5.0..HEAD` | Checked; candidate patch size identified. |
| Required ripgrep scans from the task prompt | Checked; release-scope expansion 필요 없음. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root는 blocker, deferral, warning 없이 `NI Intent Readiness: READY`를 report했다. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help includes `ni init [.]`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Excluded-scope list에 explicit `must not include` boundary context를 추가한 뒤 Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | 같은 boundary-context fix 뒤 Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

Note: 첫 `quality.sh`와 `release-check.sh` attempt는 새 excluded-scope list가
`scripts/check-core-boundary.py`에 충분한 local `must not` context를 주지 못해
failed했다. Docs를 clarify했고 release scope는 broaden하지 않았다.
