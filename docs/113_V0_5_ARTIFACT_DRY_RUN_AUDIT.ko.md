# v0.5 Artifact Dry-Run Audit

## Current status

- RC decision: RC_READY_WITH_DEFERRALS
- Release notes preflight decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- 이 audit는 publish, tag, release, GitHub release creation, asset upload를
  수행하지 않는다.

## Audit goal

이 audit는 dry-run/check-only evidence로 artifact와 install readiness를
확인한다. Later v0.5 release process 전에 release scripts, install scripts,
local build path, version/help surfaces, GoReleaser configuration, GitHub
release workflow, release-note chain, Homebrew boundary, protected `.ni` file
safety를 검토한다.

## Decision

Decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS.

Justification: source build, temporary binary, help/version output, install
checks, smoke checks, release checks, demo checks, quality checks에 대한 local
dry-run/check-only evidence가 통과했다. Full pass는 아니다. 이 audit에서는
v0.5 tag, GitHub release, asset upload, hosted checksum, Homebrew formula,
cross-platform hosted install, external install verification, model workspace
host verification, external user validation을 수행하지 않았다.

## Artifact / release surface inventory

| Surface | Files or commands | Observed state | Risk | Notes |
| --- | --- | --- | --- | --- |
| release-check script | `scripts/release-check.sh` | Available and passed | Low | Publish 없이 tests, smoke, demos, install checks, release docs, claim boundaries를 확인하는 check-only gate. |
| install-check script | `scripts/install-check.sh` | Available and passed | Low | Source, build, temporary local install paths를 사용한다. |
| smoke script | `scripts/smoke.sh` | Available and passed | Low | Temporary fixture `ni end`와 `ni relock`을 exercise한다; project-root relock evidence가 아니다. |
| quality script | `scripts/quality.sh` | Available and passed | Low | Formatting, docs/static checks, Go tests, smoke checks를 실행한다. |
| install docs | `README.md`, `README.ko.md`, `docs/22_INSTALL.md`, `docs/install-curl.md`, `docs/53_DISTRIBUTION_STRATEGY.md` | Available and bounded | Low | Release binary와 curl installer claims는 verified v0.4.0 assets에 scoped된다. |
| release docs / release notes draft | `docs/110_*`, `docs/111_*`, `docs/112_*` | Available and bounded | Low | `RC_READY_WITH_DEFERRALS`와 `RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES`가 preserved된다. |
| GitHub workflows | `.github/workflows/release.yml`, `.github/workflows/ci.yml` | Present | Medium | Release workflow는 `v*` tags에서 실행되고 `goreleaser release --clean`을 사용한다; real release path이므로 local에서 실행하지 않았다. |
| goreleaser or equivalent config | `.goreleaser.yaml` | Present | Medium | Config는 OS/arch archives와 checksums를 정의한다. Local `goreleaser` binary가 없어 `goreleaser check`와 snapshot archive build는 실행하지 않았다. |
| local binary build | `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/ni-v0.5-artifact-dry-run/ni ./cmd/ni` | Passed | Low | Temporary binary는 repo 밖에 build되었다. |
| ni help output | `go run ./cmd/ni --help`; `/tmp/ni-v0.5-artifact-dry-run/ni --help` | Passed | Low | Output includes `ni is a project intent compiler.` and the public command list. |
| ni version output | `go run ./cmd/ni version`; `/tmp/ni-v0.5-artifact-dry-run/ni version` | Passed | Low | Output은 source/local dry-run builds 기준 `0.0.0-dev`; v0.5 release claim이 아니다. |
| curl installer | `install.sh`, `docs/install-curl.md`, `docs/22_INSTALL.md` | Available and bounded | Medium | Verified availability는 v0.4.0-scoped로 유지된다; v0.5 hosted install은 시도하지 않았다. |
| release binary docs | `README.md`, `docs/22_INSTALL.md`, `docs/53_DISTRIBUTION_STRATEGY.md` | Available and bounded | Medium | Availability는 verified v0.4.0 release assets이며 v0.5 publication이 아니다. |
| Homebrew docs/status | `docs/80_HOMEBREW_DECISION.md`, `docs/53_DISTRIBUTION_STRATEGY.md`, roadmap | Planned / v0.5 candidate | Medium | 이 audit에는 tap/formula/audit/install evidence가 없다. |

## Dry-run commands

| Command | Ran? | Result | Mutation risk | Notes |
| --- | --- | --- | --- | --- |
| `git status --short` | Yes | Pass | None | Initial status는 clean이었다. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Yes | Pass | Low | Cache writes에는 `/private/tmp/ni-go-cache` workaround를 사용했다. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Yes | Pass | None | `NI Intent Readiness: READY`; blockers, deferrals, warnings: `None`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Yes | Pass | None | Help output rendered successfully. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Yes | Pass | None | Output: `0.0.0-dev`. |
| `mkdir -p /tmp/ni-v0.5-artifact-dry-run` | Yes | Pass | Low | Repo 밖 temporary output directory를 만들었다. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/ni-v0.5-artifact-dry-run/ni ./cmd/ni` | Yes | Pass | Low | Local temporary binary only. |
| `/tmp/ni-v0.5-artifact-dry-run/ni --help` | Yes | Pass | None | Temporary binary help output rendered successfully. |
| `/tmp/ni-v0.5-artifact-dry-run/ni version` | Yes | Pass | None | Output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Yes | Pass | None | Install/distribution claim markers passed. |
| `bash scripts/check-skill-packs.sh` | Yes | Pass | Low | Repo-local evidence를 report하고 global install은 unverified로 유지한다. |
| `bash scripts/demo-check.sh` | Yes | Pass | Low | Temporary prompt/export artifacts only; downstream prompt는 실행하지 않았다. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Yes | Pass | Low | Fixture `ni end`와 `ni relock`을 exercise한다; project-root relock이 아니다. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Yes | Pass | Low | Temporary install path와 `bin/ni` local build output을 사용한다. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Yes | Pass | Low | Check-only release readiness gate passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Yes | Pass | Low | Broad quality wrapper passed after edits. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Yes | Pass | None | Protected project-root `.ni` file diff 없음. |
| `goreleaser check` | No | Not run | Low if installed | Local `goreleaser` unavailable. |
| `goreleaser release --snapshot --clean` | No | Not run | Medium | Local `goreleaser` unavailable; 실행하면 local `dist/` artifacts를 만들 수 있다. |
| `.github/workflows/release.yml` / `goreleaser release --clean` | No | Not run | High | Tag-triggered publish path; 이 audit에 적합하지 않다. |
| `git tag`, `git push`, GitHub release creation, asset upload | No | Not run | High | Task scope에서 forbidden. |

## Version and help audit

| Surface | Observed output or behavior | Pass? | Notes |
| --- | --- | --- | --- |
| ni --help or equivalent | `ni is a project intent compiler.` plus public command list. | Yes | `go run`과 temporary binary help 모두 passed. |
| ni version or equivalent | `0.0.0-dev`. | Yes | Source/local dry-run에는 맞지만 v0.5 release version claim이 아니다. |
| README command references | README는 `ni --help`, `ni version`, `make build`, curl installer dry-run examples를 사용한다. | Yes | Wrong command reference 발견 없음. |
| install docs command references | `docs/22_INSTALL.md`는 `go run ./cmd/ni --help`, `go run ./cmd/ni version`, `make build`, `ni --help`, `ni version`, `install.sh --dry-run`을 사용한다. | Yes | Fix 필요 없음. |
| release notes command references | docs/111과 docs/112는 validation commands와 release boundary wording을 유지한다. | Yes | Release/publish claim 발견 없음. |

## Install surface audit

| Surface | Expected status | Observed status | Pass? | Notes |
| --- | --- | --- | --- | --- |
| release binary | Available | Available for verified v0.4.0 release assets | Yes | 이 audit는 v0.5 release binaries가 published되었다고 claim하지 않는다. |
| curl installer | Available | Available for verified v0.4.0 release assets | Yes | `install.sh`가 있고 install docs are bounded. |
| Homebrew | Planned / v0.5 candidate | Planned / v0.5 candidate | Yes | Homebrew Available claim 없음. |
| checksum / sha256 references | Required for release archives | `.goreleaser.yaml`, install docs, installer behavior에 present | Yes | Hosted v0.5 checksum availability는 deferred. |
| install-check behavior | Source/build/temp install checks pass | Passed | Yes | Temporary local install proof only. |
| release-check behavior | Release-readiness gate passes | Passed | Yes | Check-only gate; release action 없음. |

## Blockers

None.

| Blocker | Severity | Evidence | Required fix |
| --- | --- | --- | --- |
| None | n/a | Check-only artifact surfaces passed. | n/a |

## Deferrals

| Deferral | Reason | Why it does not block dry-run readiness | Required future evidence |
| --- | --- | --- | --- |
| actual v0.5 publication | 이 audit는 publish, tag, release, GitHub release creation을 하지 않는다. | Dry-run readiness는 publication claim 없이 pass할 수 있다. | Authorized release checklist, tag, workflow run, release page verification. |
| GitHub release asset upload | Assets were not uploaded. | Local build/check evidence만으로 dry-run decision에는 충분하다. | GitHub release assets and matching checksums visible on the release page. |
| Homebrew implementation / availability | Tap/formula/audit/install proof 없음. | Homebrew는 distribution candidate이며 dry-run artifact readiness의 필수 조건이 아니다. | Tap, formula, sha256, audit, `brew install`, `ni --help`, `ni version` proof. |
| cross-platform binary verification | Local build는 이 environment만 확인했다. | GoReleaser config는 target matrix를 정의하지만 hosted cross-platform artifacts는 future release evidence다. | Matrix archive build, checksum file, platform-specific `ni --help` / `ni version` proof. |
| external install verification | External user 또는 separate host가 이 audit에서 install path를 실행하지 않았다. | Local dry-run/install checks passed. | User-run or host-specific install transcript with checksum, help, version output. |
| model workspace host verification | Model workspace packs remain Experimental. | Model workspace availability는 artifact dry-run readiness와 별개다. | Host-specific install/discovery proof and provider behavior transcript. |
| external user validation | External users were not observed. | Dry-run script/doc readiness는 adoption proof 전에 pass할 수 있다. | Maintained external validation notes, user-run transcripts, or scoped user research. |

## Warnings

| Warning | Evidence | Mitigation | Next action |
| --- | --- | --- | --- |
| `goreleaser` was not installed locally. | `command -v goreleaser` returned no path. | GoReleaser config는 present로 기록하되 locally unchecked로 둔다. | Later environment with GoReleaser installed에서 `goreleaser check`와 필요시 `goreleaser release --snapshot --clean` 실행. |
| `0.0.0-dev` is correct for dry-run source builds but not a release version. | `go run ./cmd/ni version` and temporary binary `version` output were `0.0.0-dev`. | Local development output only로 취급한다. | Final release artifact process에서 linker-injected version을 verify한다. |
| `scripts/install-check.sh` creates `bin/ni`. | Install-check passed and uses `make build`. | Generated build output은 unstaged로 유지하고 visible하면 report한다. | Commit 전에 local build output을 clean 또는 ignore 상태로 둔다. |

## Risks

| Risk | Impact | Mitigation | Follow-up |
| --- | --- | --- | --- |
| dry-run is not publication | Maintainers가 local checks를 hosted release evidence로 오해할 수 있다. | 이 문서는 publish/tag/release/upload가 없었다고 명시한다. | Release action 전 separate publication checklist 사용. |
| local build is not cross-platform verification | Platform-specific archive 또는 install bug가 남을 수 있다. | Cross-platform verification을 deferred로 유지한다. | GoReleaser matrix artifacts와 per-platform help/version output verify. |
| release-check coverage may not prove hosted artifact availability | Checks가 통과해도 GitHub release는 없을 수 있다. | Local check gates와 hosted artifact proof를 구분한다. | Publication 후 release page assets/checksums verify. |
| Homebrew remains Planned | Package-manager users는 아직 `brew install`이 없다. | `Homebrew: Planned / v0.5 candidate`를 preserve한다. | 필요하면 Homebrew implementation audit 실행. |
| external user validation remains limited | Adoption friction이 unknown일 수 있다. | External validation을 deferral로 유지한다. | Priority가 되면 external user validation plan 준비. |
| provider host behavior remains unverified | Model workspace behavior가 host마다 다를 수 있다. | Model workspace packs remain Experimental. | 필요하면 model workspace host verification audit 실행. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Published/released status | v0.5 published/released claim 금지. | v0.5 publication claim 추가 없음. | Yes | 이 audit는 no release action을 명시한다. |
| Artifact upload status | Assets uploaded claim 금지. | Upload claim 추가 없음. | Yes | Local temporary binary only. |
| GitHub release status | v0.5 GitHub release exists claim 금지. | GitHub release claim 추가 없음. | Yes | Release workflow was inspected, not run. |
| Homebrew | Keep Homebrew: Planned / v0.5 candidate. | Preserved. | Yes | Formula 또는 `brew install` claim 없음. |
| Model workspace packs | Keep Model workspace packs: Experimental. | Preserved. | Yes | Host/global install remains unverified. |
| No-terminal | Keep No-terminal method: Experimental / assisted. | Preserved. | Yes | Deterministic no-terminal claim 없음. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | Demo/check prompt compilation은 downstream execution이 아니다. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | Preserved. | Yes | Implementation/adoption/cost/latency/downstream-performance claim 없음. |
| Fixture relock | Fixture relock is not project-root relock. | Preserved. | Yes | Smoke/release-check fixture `ni end`와 `ni relock`은 separated. |
| Runtime execution boundary | `ni` is not a task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. | Preserved. | Yes | Runtime feature added 없음. |

## Git status / generated file audit

| Path or group | Status from git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| docs/110_* | no output | No | Prior task의 docs/110 files는 현재 git status에 visible하지 않다; already clean 또는 committed로 보인다. |
| docs/111_* | no output | No | Prior task의 docs/111 files는 현재 git status에 visible하지 않다; already clean 또는 committed로 보인다. |
| docs/112_* | `M docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`; `M docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md` | Yes | docs/113으로 이어지는 narrow follow-up pointers 추가. |
| docs/113_* | `?? docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md`; `?? docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md` | Yes | New artifact dry-run audit docs. |
| docs/51* | `M docs/51_POST_RELEASE_ROADMAP.md`; `M docs/51_POST_RELEASE_ROADMAP.ko.md` | Yes | Narrow roadmap pointers to docs/113 추가. |
| generated artifacts, if any | no tracked output expected | No | Temporary `/tmp/ni-v0.5-artifact-dry-run` output은 final reporting 전에 제거했다; `bin/ni`는 install checks가 만들 수 있지만 staged되면 안 된다. |
| .ni/contract.json | no output | No | Protected project-root file unchanged. |
| .ni/session.json | no output | No | Protected project-root file unchanged. |
| .ni/plan.lock.json | no output | No | Protected project-root file unchanged. |
| unexpected files | none expected | No | Final git status should show only docs/51, docs/112, and docs/113 changes. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, warnings are `None`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass | Help output rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass | Output: `0.0.0-dev`. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/ni-v0.5-artifact-dry-run/ni ./cmd/ni` | Pass | Temporary local binary built outside repo. |
| `/tmp/ni-v0.5-artifact-dry-run/ni --help` | Pass | Help output rendered. |
| `/tmp/ni-v0.5-artifact-dry-run/ni version` | Pass | Output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack checks passed. |
| `bash scripts/demo-check.sh` | Pass | Public demos verified without downstream runtime execution. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Fixture `ni end` and `ni relock` paths exercised only in temporary workspaces. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source, build, and temporary install paths passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Release readiness gate passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passed after edits. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` file diff. |

## Changes made

- `docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md`: English artifact dry-run audit
  추가.
- `docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md`: Korean companion 추가.
- `docs/51_POST_RELEASE_ROADMAP.md`: narrow docs/113 pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer 추가.
- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`: narrow docs/113 follow-up
  pointer 추가.
- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`: matching Korean
  follow-up pointer 추가.

## What this audit proves

- local dry-run/check-only artifact surfaces passed as documented
- release/install docs remain bounded by audited status
- no release action was performed
- known artifact deferrals remain explicit

## What this audit does not prove

- v0.5 has been published
- GitHub release exists
- assets were uploaded
- hosted checksums are available
- Homebrew is Available
- cross-platform install works
- external users succeed
- model workspace host behavior is verified
- no-terminal is deterministic
- downstream execution succeeds
- benchmark effect size or causal impact

## Recommended next task

Selected next task: F. v0.5 release publication checklist.

Why: artifact dry-run이 explicit deferrals와 함께 passed했으므로, 다음으로 가장
safe한 단계는 maintainer가 tag, publish, upload assets, availability claims update
전에 필요한 exact gated actions와 evidence를 정리하는 non-executing publication
checklist다.

## Next task prompt

```text
Proceed with a v0.5 release publication checklist in /Users/namba/Documents/project/ni.

This is a checklist and documentation task only. Do not publish, tag, create a GitHub release, upload assets, run a release workflow, run goreleaser publish, create or publish a Homebrew formula, run ni end on the project root, relock the project root, execute generated prompts, add release automation, add runtime execution behavior, or mark v0.5 as released.

Use these docs as the current release-readiness chain:
- docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md
- docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md
- docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md
- docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md

Preserve:
- RC decision: RC_READY_WITH_DEFERRALS
- Release notes preflight decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES
- Artifact dry-run decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- ni run compiles a bounded handoff prompt only.
- READY is planning contract readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- fixture relock is separate from project-root relock.
- benchmark evidence keeps not_measured boundaries.

Goal:
Create a non-executing v0.5 publication checklist that tells a maintainer exactly what must be verified before any release action. The checklist should separate pre-publication checks, tag/release workflow actions, post-publication hosted artifact checks, checksum checks, current-platform binary checks, curl installer checks, rollback notes, Homebrew deferral notes, documentation update gates, and forbidden overclaims.

Required output:
- docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md
- docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md if Korean companion docs are maintained
- narrow roadmap / docs/113 cross-links only if useful

Run check-only validation:
- git status --short
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final report must confirm no project-root relock, no protected .ni file changes, no generated prompt execution, no release/tag/publish/upload action, no Homebrew Available claim, no model workspace Available claim, no no-terminal deterministic claim, no benchmark overclaim, and whether validation scripts exercised fixture ni end while keeping fixture runs distinct from project-root relock.
```
