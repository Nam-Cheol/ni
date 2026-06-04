# v0.5.1 Publication Checklist

## Current status

State:
- v0.5.0 publication: verified
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES
- v0.5.1 RC validation decision: V0_5_1_RC_VALIDATION_PASS_WITH_NOTES
- v0.5.1 artifact dry-run decision: V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES
- v0.5.1 release notes decision: V0_5_1_RELEASE_NOTES_READY_WITH_NOTES
- v0.5.1 release: not published
- Homebrew: Planned / v0.5 candidate
- Windows real-host execution: macOS-only development host에서는 deferred
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Checklist goal

이 checklist는 later human-approved v0.5.1 release를 gate한다. Maintainer가
publication 전, 중, 후에 무엇을 verify해야 하는지 정리하지만, 이 task는 publish,
tag, GitHub release creation, asset upload, release workflow 실행, GoReleaser
publish, Homebrew formula 생성/게시, project root의 `ni end`, project root relock,
generated prompt execution, runtime execution behavior 추가, v0.5.1 released
claim을 수행하지 않는다.

아래 모든 release action은 `Run in this task? No`로 표시한다.

## Decision

V0_5_1_PUBLICATION_CHECKLIST_READY_WITH_NOTES

Justification: publication checklist는 complete이고 future release actions가
check-only validation과 분리되어 있다. Notes는 v0.5.1이 아직 explicit human
approval을 필요로 하고, v0.5.1이 not published이며, hosted artifacts와 public
`install.sh` retrieval이 아직 없고, Windows real-host execution이 deferred이며,
Homebrew가 Planned / v0.5 candidate로 남고, external user validation이 수행되지
않았고, model workspace behavior가 Experimental이며, no-terminal method가
Experimental / assisted이고, GoReleaser full matrix dry-run은 GoReleaser가
available한 환경 전까지 deferred이기 때문에 남는다.

## Release-readiness chain

| Evidence | Decision / result | Pass? | Notes |
| --- | --- | --- | --- |
| docs/126 | PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED | Yes | Published v0.5.0은 `ni --help`와 `ni version`은 pass하지만 `ni init .`은 fail한다; v0.5.1 patch가 필요하다. |
| docs/127 | V0_5_1_PATCH_PLAN_READY_WITH_NOTES | Yes | Patch scope, exclusions, validation matrix, protected `.ni` boundary가 documented되어 있다. |
| docs/128 | V0_5_1_RC_VALIDATION_PASS_WITH_NOTES | Yes | Current-tree RC behavior와 checker suite가 passed했고 publication, Homebrew, Windows, external validation은 deferred다. |
| docs/129 | V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES | Yes | Local darwin/arm64 release-like artifact가 `0.5.1`을 report했고 local checksum이 verified되었다; full GoReleaser matrix와 hosted assets는 deferred다. |
| docs/130 | V0_5_1_RELEASE_NOTES_READY_WITH_NOTES | Yes | Release notes draft는 conservative하며 publication을 claim하지 않고 known deferrals를 preserve한다. |

## Pre-publication gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| repository state review | `git status --short`, recent log, `git diff --name-only v0.5.0..HEAD`, `git diff --stat v0.5.0..HEAD` review. | No | Yes | 이 checklist는 required gate를 기록한다; 아래 validation은 이 task의 check-only state를 기록한다. |
| protected `.ni` diff check | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json`가 empty. | No | Yes | Protected diff가 있으면 release execution은 blocked다. |
| validation commands | Go tests, install docs checks, skill pack checks, demo, quality, smoke, install, release checks pass. | No | Yes | Exact release commit에서 다시 실행한다. |
| release notes final review | docs/130 release notes draft를 actual release artifacts와 hosted asset inventory에 맞춰 review. | No | Yes | Unavailable artifacts를 mention하는 notes를 publish하지 않는다. |
| version injection confirmation | Release-built `ni version`이 `0.5.1`을 report한다. | No | Yes | Source `go run ./cmd/ni version`은 `0.0.0-dev`일 수 있으며 release proof가 아니다. |
| artifact dry-run confirmation | Archive names, `ni_0.5.1_checksums.txt`, current-platform binary smoke가 verified. | No | Yes | docs/129는 local darwin/arm64 fallback을 다룬다; future task는 available하면 GoReleaser를 사용한다. |
| install.sh future URL check | `sh install.sh --dry-run --version 0.5.1`이 expected v0.5.1 URLs를 print한다. | No | Yes | Dry-run URL construction은 public retrieval이 아니다. |
| claim-boundary audit | README/docs가 hosted v0.5.1, Homebrew Available, Windows real-host proof, `ni run` execution을 claim하지 않는다. | No | Yes | 이 문서의 audit table이 minimum boundary set이다. |
| human approval requirement | Maintainer가 이 checklist 뒤 listed approval option 중 하나를 explicit하게 선택한다. | No | Yes | 이 task는 어떤 option도 선택하지 않는다. |

## Tag gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| verify no existing v0.5.1 tag | `git tag --list v0.5.1` returns empty. | No | Yes | Tag가 존재하면 계속하기 전에 verify하거나 abort한다. |
| create v0.5.1 tag | Annotated tag가 intended release commit을 가리킨다. | No | Yes | Later explicitly approved release task에서만 create한다. |
| verify tag commit | `git rev-parse v0.5.1^{}` 또는 equivalent가 intended commit과 match한다. | No | Yes | Model judgment로 tag target을 assume하지 않는다. |
| push tag | Remote tag가 존재하고 intended commit을 가리킨다. | No | Yes | Later explicitly approved release task에서만 push한다. |

## Artifact gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| build release artifacts | GoReleaser 또는 approved release build가 expected v0.5.1 archives를 만든다. | No | Yes | Later release task에서만 build한다. |
| verify version output 0.5.1 | Extracted current-platform artifact가 `ni version`을 실행해 `0.5.1`을 print한다. | No | Yes | Publication claim 전 required proof다. |
| generate checksums | `ni_0.5.1_checksums.txt`가 intended artifacts에서 generated된다. | No | Yes | Checksum file은 exact uploaded assets와 match해야 한다. |
| verify checksums | `shasum -a 256 -c ni_0.5.1_checksums.txt` 또는 equivalent가 pass한다. | No | Yes | Mismatch는 release abort다. |
| verify current-platform artifact | Extracted binary가 temp workspace에서 `ni --help`, `ni version`, `ni init . --yes`, `ni status --proof --next-questions`를 pass한다. | No | Yes | Blank intent fixture status가 `BLOCKED`일 수 있으며 expected다. |
| verify archive names | Archives가 `.goreleaser.yaml` naming for v0.5.1와 match한다. | No | Yes | Expected names include `ni_0.5.1_darwin_arm64.tar.gz` and `ni_0.5.1_checksums.txt`. |

## GitHub release gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| create GitHub release | GitHub release가 intended tag의 `v0.5.1`로 존재한다. | No | Yes | Later explicitly approved release task에서만 create한다. |
| attach assets | Expected archives가 GitHub release에 attached되어 있다. | No | Yes | Visible하기 전 hosted artifacts를 claim하지 않는다. |
| attach checksums | `ni_0.5.1_checksums.txt`가 attached되어 있다. | No | Yes | Installer와 user verification에 required다. |
| verify draft status | Release가 의도된 draft가 아니면 not draft다. | No | Yes | Draft가 intentional이면 docs가 그렇게 말해야 한다. |
| verify prerelease status | Release가 의도된 prerelease가 아니면 not prerelease다. | No | Yes | Maintainer가 다르게 결정하지 않는 한 v0.5.1 patch는 normal release로 기대된다. |
| verify release URL | Release URL이 열리고 intended tag/version을 가리킨다. | No | Yes | Final release execution report에 URL을 기록한다. |

## Post-publication gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| download hosted artifact | Hosted v0.5.1 archive가 GitHub release에서 download된다. | No | Yes | Hosted asset availability는 publication 뒤에만 시작된다. |
| verify checksum | Hosted artifact checksum이 hosted `ni_0.5.1_checksums.txt`와 match한다. | No | Yes | Mismatch는 public availability claim을 abort한다. |
| run hosted artifact `ni --help` | Extracted hosted artifact가 help를 render한다. | No | Yes | Command name 또는 extracted binary path 중 무엇을 썼는지 기록한다. |
| run hosted artifact `ni version` | Extracted hosted artifact가 `0.5.1`을 print한다. | No | Yes | Required release proof다. |
| isolated `install.sh --version 0.5.1` | Temporary HOME/BINDIR install이 hosted v0.5.1을 retrieve하고 checksum을 verify하며 command-name `ni`를 install한다. | No | Yes | Hosted assets가 존재한 뒤에만 가능하다. |
| isolated curl installer verification | Fresh shell `ni --help`, `ni version`, `ni init .`, uninstall이 temp locations에서 pass한다. | No | Yes | User install state를 오염시키지 않는다. |
| README / docs final check | README, README.ko, install docs, release notes가 actual hosted state와 match한다. | No | Yes | v0.5.1 proof 전에는 필요한 경우 v0.5.0 parity note를 유지한다. |
| Homebrew boundary check | Separate Homebrew proof 없이는 Homebrew Available claim 없음. | No | Yes | Homebrew는 기본적으로 Planned / v0.5 candidate다. |

## Abort criteria

| Abort condition | Why it matters | Required action | Notes |
| --- | --- | --- | --- |
| validation failure | Release commit을 trust할 수 없다. | Stop, fix, rerun check-only validation. | Failing gate를 우회해서 tag하지 않는다. |
| version mismatch | Users가 wrong binary를 받을 수 있다. | Stop, fix linker/version injection, rebuild artifacts. | `0.0.0-dev`는 source run에는 acceptable하지만 release artifact에는 아니다. |
| checksum mismatch | Downloads를 안전하게 verify할 수 없다. | Stop, regenerate artifacts/checksums from intended commit. | Bad assets를 upload하거나 유지하지 않는다. |
| hosted asset unavailable | Installer와 docs가 missing files를 가리킬 수 있다. | Stop public availability claims and fix release assets. | `install.sh`는 missing assets를 retrieve할 수 없다. |
| install.sh v0.5.1 retrieval failure after publication | Public install path가 broken이다. | Stop release announcement or correct docs until install succeeds. | Proof 전 public retrieval claim을 하지 않는다. |
| protected `.ni` diff appears | Project-root planning state가 changed되었을 수 있다. | Stop and investigate; root planning files를 silently edit하거나 relock하지 않는다. | Protected files are `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`. |
| Homebrew overclaim | Package-manager proof가 없다. | Remove claim and keep Homebrew Planned / v0.5 candidate. | Homebrew는 separate tap/formula/install proof가 필요하다. |
| Windows real-host overclaim | Static checks는 Windows execution proof가 아니다. | Remove claim until Windows transcript exists. | macOS-only development host는 Windows real-host execution을 prove할 수 없다. |
| ni run execution overclaim | ni-kernel boundary 위반이다. | Remove claim and verify `ni run` remains bounded prompt compilation only. | `ni run` must not execute downstream work. |

## Human approval options

Actual v0.5.1 release execution은 이 checklist 뒤 explicit human approval이
필요하다. 이 task는 어떤 option도 선택하지 않는다.

- APPROVE_V0_5_1_RELEASE_EXECUTION
- DO_NOT_APPROVE_FIX_FIRST
- DEFER_V0_5_1_RELEASE

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks checklist? |
| --- | --- | --- | --- |
| v0.5.1 publication | 이 task는 non-publishing이고 human approval option이 선택되지 않았다. | Explicit approval, tag, GitHub release, hosted assets, post-publication checks. | No |
| hosted artifacts | GitHub release 또는 asset upload가 수행되지 않았다. | Release page inventory, archive download, checksum verification. | No |
| install.sh actual v0.5.1 retrieval | Hosted v0.5.1 assets가 아직 없다. | Isolated hosted install with checksum, `ni --help`, `ni version`, `ni init .`. | No |
| Windows real-host execution | Current development host는 macOS-only다. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | Tap/formula/install proof가 없다. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | 이 task에는 external user 또는 separate machine transcript가 없다. | External install/init/status transcript. | No |
| model workspace host behavior | Host-level/global install과 provider behavior는 unverified다. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal은 Experimental / assisted로 남는다. | Trusted CLI proof for a target workspace. | No |
| GoReleaser full matrix dry-run if unavailable | Local GoReleaser availability는 host마다 다를 수 있다. | GoReleaser installed environment에서 check/dry-run matrix 실행. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Preserved. | Yes | Publish/tag/release action 없음. |
| hosted artifacts | GitHub release upload 전에는 assets exist claim 금지. | Preserved. | Yes | docs/129 local artifact는 hosted proof가 아니다. |
| install.sh v0.5.1 retrieval | Hosted assets 존재와 install pass 전에는 public retrieval claim 금지. | Preserved. | Yes | Dry-run URL construction only. |
| Homebrew | Must remain Homebrew: Planned / v0.5 candidate. | Preserved. | Yes | Homebrew Available claim 없음. |
| Windows real-host execution | Windows transcript 없이는 deferred로 남아야 한다. | Preserved. | Yes | Static checks는 real-host execution이 아니다. |
| ni init . | Guided project intent setup only. | Preserved. | Yes | Planning artifacts를 만들지만 agents를 run하지 않는다. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | Generated prompt execution 없음. |
| READY | CLI readiness only, not product readiness. | Preserved. | Yes | `ni status` is authority. |
| runtime execution boundary | No task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | Runtime behavior 추가 없음. |
| human approval status | 이 task가 선택하면 안 된다. | Preserved. | Yes | Options are listed only. |

## Git status / inclusion check

| Path or group | git status --short | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | task start에는 clean; `v0.5.0..HEAD`에서 changed | Yes | Public onboarding and parity note. |
| README.ko.md | task start에는 clean; `v0.5.0..HEAD`에서 changed | Yes | Korean companion. |
| docs/126* | tracked | Yes | Public install parity evidence. |
| docs/127* | tracked | Yes | v0.5.1 patch plan. |
| docs/128* | tracked | Yes | v0.5.1 RC validation. |
| docs/129* | tracked | Yes | v0.5.1 artifact dry-run. |
| docs/130* | tracked | Yes | v0.5.1 release notes finalization. |
| docs/131* | added by this task | Yes | Publication checklist and Korean companion. |
| CHANGELOG.md | absent | No | Publication을 imply하지 않기 위해 추가하지 않는다. |
| RELEASE.md | absent | No | Publication을 imply하지 않기 위해 추가하지 않는다. |
| .ni/contract.json | no diff | No direct edit | Protected. |
| .ni/session.json | no diff | No direct edit | Protected. |
| .ni/plan.lock.json | no diff | No direct edit | Protected. |
| unexpected files | task start에는 none | No | Validation 뒤 recheck한다. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; `docs/51_POST_RELEASE_ROADMAP*` modifications와 new `docs/131*` files만 visible하다. |
| `git log --oneline --decorate -20` | Checked; 이 task 전 HEAD는 `36609df Add v0.5.1 artifact and release notes roadmap links`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; no v0.5.1 tag exists. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Checked; docs/126 through docs/130 are tracked in current patch delta. |
| `git diff --stat v0.5.0..HEAD` | Checked; docs/131 전 70 files changed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; edits 전과 validation 뒤 모두 empty. |
| Required ripgrep scans | Release, version, install, Homebrew, runtime boundary surfaces reviewed. |
| `CHANGELOG.md` / `RELEASE.md` presence check | Both files are absent. |
| `.github/workflows/` presence check | `.github/workflows/ci.yml` and `.github/workflows/release.yml` are present. |
| release scripts / installers check | `install.sh`, `install.ps1`, `scripts/release-check.sh`, `scripts/install-check.sh`, `scripts/release-dry-run.sh` are present. |
| `gofmt -w .` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root가 `NI Intent Readiness: READY`를 report했고 blockers, deferrals, warnings는 none. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output은 `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |

## Changes made

| File | Why |
| --- | --- |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.md` | English v0.5.1 publication checklist, human approval options, gates, deferrals, claim-boundary audit 추가. |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.ko.md` | 같은 boundaries를 가진 Korean companion 추가. |

## What this checklist proves

- publication checklist is ready with notes
- release actions are gated
- human approval is still required
- no release action was performed

## What this checklist does not prove

- v0.5.1 has been published
- v0.5.1 artifacts are hosted
- `install.sh` retrieves v0.5.1 publicly
- Windows real-host execution works
- Homebrew is Available
- external users succeed
- downstream execution succeeds
- no-terminal is deterministic

## Recommended next task

Selected next task: A. wait for human approval

Selection rationale: checklist는 ready with notes지만 actual v0.5.1 release
execution은 아직 explicit human approval이 필요하다. User가 이 task에서 release
execution을 approve하지 않았으므로 B는 selected가 아니다.

## Next task prompt

Human approval request template only:

```text
v0.5.1 publication checklist is ready with notes.

Please choose exactly one:
- APPROVE_V0_5_1_RELEASE_EXECUTION
- DO_NOT_APPROVE_FIX_FIRST
- DEFER_V0_5_1_RELEASE

No release action will run unless APPROVE_V0_5_1_RELEASE_EXECUTION is selected
explicitly in a later task.
```
