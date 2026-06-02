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
- This audit does not publish, tag, release, create a GitHub release, or upload
  assets.

## Audit goal

This audit checks artifact and install readiness using dry-run/check-only
evidence. It reviews the release scripts, install scripts, local build path,
version/help surfaces, GoReleaser configuration, GitHub release workflow,
release-note chain, Homebrew boundary, and protected `.ni` file safety before
any later v0.5 release process.

## Decision

Decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS.

Justification: local dry-run/check-only evidence passed for the source build,
temporary binary, help/version output, install checks, smoke checks, release
checks, demo checks, and quality checks. The decision is not a full pass because
no v0.5 tag, GitHub release, asset upload, hosted checksum, Homebrew formula,
cross-platform hosted install, external install verification, model workspace
host verification, or external user validation was performed.

## Artifact / release surface inventory

| Surface | Files or commands | Observed state | Risk | Notes |
| --- | --- | --- | --- | --- |
| release-check script | `scripts/release-check.sh` | Available and passed | Low | Check-only gate; it exercises tests, smoke, demos, install checks, release docs, and claim boundaries without publishing. |
| install-check script | `scripts/install-check.sh` | Available and passed | Low | Uses source, build, and temporary local install paths. |
| smoke script | `scripts/smoke.sh` | Available and passed | Low | Exercises temporary fixture `ni end` and `ni relock`; this is fixture-only evidence, not project-root relock. |
| quality script | `scripts/quality.sh` | Available and passed | Low | Runs formatting, docs/static checks, Go tests, and smoke checks. |
| install docs | `README.md`, `README.ko.md`, `docs/22_INSTALL.md`, `docs/install-curl.md`, `docs/53_DISTRIBUTION_STRATEGY.md` | Available and bounded | Low | Release binary and curl installer claims remain scoped to verified v0.4.0 assets. |
| release docs / release notes draft | `docs/110_*`, `docs/111_*`, `docs/112_*` | Available and bounded | Low | `RC_READY_WITH_DEFERRALS` and `RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES` are preserved. |
| GitHub workflows | `.github/workflows/release.yml`, `.github/workflows/ci.yml` | Present | Medium | Release workflow runs on `v*` tags and uses `goreleaser release --clean`; not run locally because it would be a real release path. |
| goreleaser or equivalent config | `.goreleaser.yaml` | Present | Medium | Config defines OS/arch archives and checksums. Local `goreleaser` binary is unavailable, so `goreleaser check` and snapshot archive build were not run in this audit. |
| local binary build | `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/ni-v0.5-artifact-dry-run/ni ./cmd/ni` | Passed | Low | Temporary binary was built outside the repo. |
| ni help output | `go run ./cmd/ni --help`; `/tmp/ni-v0.5-artifact-dry-run/ni --help` | Passed | Low | Output includes `ni is a project intent compiler.` and the public command list. |
| ni version output | `go run ./cmd/ni version`; `/tmp/ni-v0.5-artifact-dry-run/ni version` | Passed | Low | Output was `0.0.0-dev` for source/local dry-run builds; suitable for development checks, not a v0.5 release claim. |
| curl installer | `install.sh`, `docs/install-curl.md`, `docs/22_INSTALL.md` | Available and bounded | Medium | Verified availability remains v0.4.0-scoped; no v0.5 hosted install was attempted. |
| release binary docs | `README.md`, `docs/22_INSTALL.md`, `docs/53_DISTRIBUTION_STRATEGY.md` | Available and bounded | Medium | Availability remains verified v0.4.0 release assets, not v0.5 publication. |
| Homebrew docs/status | `docs/80_HOMEBREW_DECISION.md`, `docs/53_DISTRIBUTION_STRATEGY.md`, roadmap | Planned / v0.5 candidate | Medium | No tap/formula/audit/install evidence exists in this audit. |

## Dry-run commands

| Command | Ran? | Result | Mutation risk | Notes |
| --- | --- | --- | --- | --- |
| `git status --short` | Yes | Pass | None | Initial status was clean. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Yes | Pass | Low | Uses `/private/tmp/ni-go-cache` workaround for cache writes. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Yes | Pass | None | Reported `NI Intent Readiness: READY`; blockers, deferrals, warnings: `None`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Yes | Pass | None | Help output rendered successfully. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Yes | Pass | None | Output: `0.0.0-dev`. |
| `mkdir -p /tmp/ni-v0.5-artifact-dry-run` | Yes | Pass | Low | Created a temporary output directory outside the repo. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/ni-v0.5-artifact-dry-run/ni ./cmd/ni` | Yes | Pass | Low | Built a local temporary binary only. |
| `/tmp/ni-v0.5-artifact-dry-run/ni --help` | Yes | Pass | None | Temporary binary help output rendered successfully. |
| `/tmp/ni-v0.5-artifact-dry-run/ni version` | Yes | Pass | None | Output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Yes | Pass | None | Install/distribution claim markers passed. |
| `bash scripts/check-skill-packs.sh` | Yes | Pass | Low | Reports repo-local evidence and keeps global install unverified. |
| `bash scripts/demo-check.sh` | Yes | Pass | Low | Creates temporary prompt/export artifacts only; no downstream prompt was executed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Yes | Pass | Low | Exercises fixture `ni end` and `ni relock`; not project-root relock. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Yes | Pass | Low | Uses temporary install path plus `bin/ni` local build output. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Yes | Pass | Low | Check-only release readiness gate passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Yes | Pass | Low | Broad quality wrapper passed after edits. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Yes | Pass | None | No protected project-root `.ni` file diff. |
| `goreleaser check` | No | Not run | Low if installed | Local `goreleaser` was unavailable. |
| `goreleaser release --snapshot --clean` | No | Not run | Medium | Local `goreleaser` was unavailable; this would create local `dist/` artifacts if run. |
| `.github/workflows/release.yml` / `goreleaser release --clean` | No | Not run | High | Tag-triggered publish path; not appropriate for this audit. |
| `git tag`, `git push`, GitHub release creation, asset upload | No | Not run | High | Forbidden by task scope. |

## Version and help audit

| Surface | Observed output or behavior | Pass? | Notes |
| --- | --- | --- | --- |
| ni --help or equivalent | `ni is a project intent compiler.` plus public command list. | Yes | Both `go run` and temporary binary help passed. |
| ni version or equivalent | `0.0.0-dev`. | Yes | Correct for source/local dry-run; not a v0.5 release version claim. |
| README command references | README uses `ni --help`, `ni version`, `make build`, and curl installer dry-run examples. | Yes | No wrong command reference found. |
| install docs command references | `docs/22_INSTALL.md` uses `go run ./cmd/ni --help`, `go run ./cmd/ni version`, `make build`, `ni --help`, `ni version`, and `install.sh --dry-run`. | Yes | No fix needed. |
| release notes command references | docs/111 and docs/112 keep validation commands and release boundary wording. | Yes | No release/publish claim found. |

## Install surface audit

| Surface | Expected status | Observed status | Pass? | Notes |
| --- | --- | --- | --- | --- |
| release binary | Available | Available for verified v0.4.0 release assets | Yes | This audit does not claim v0.5 release binaries are published. |
| curl installer | Available | Available for verified v0.4.0 release assets | Yes | `install.sh` is present and install docs are bounded. |
| Homebrew | Planned / v0.5 candidate | Planned / v0.5 candidate | Yes | No Homebrew Available claim. |
| checksum / sha256 references | Required for release archives | Present in `.goreleaser.yaml`, install docs, and installer behavior | Yes | Hosted v0.5 checksum availability is deferred. |
| install-check behavior | Source/build/temp install checks pass | Passed | Yes | Temporary local install proof only. |
| release-check behavior | Release-readiness gate passes | Passed | Yes | Check-only gate; no release action. |

## Blockers

None.

| Blocker | Severity | Evidence | Required fix |
| --- | --- | --- | --- |
| None | n/a | Check-only artifact surfaces passed. | n/a |

## Deferrals

| Deferral | Reason | Why it does not block dry-run readiness | Required future evidence |
| --- | --- | --- | --- |
| actual v0.5 publication | This audit does not publish, tag, release, or create a GitHub release. | Dry-run readiness can pass without claiming publication. | Authorized release checklist, tag, workflow run, and release page verification. |
| GitHub release asset upload | No assets were uploaded. | Local build/check evidence is enough for this dry-run decision. | GitHub release assets and matching checksums visible on the release page. |
| Homebrew implementation / availability | No tap/formula/audit/install proof exists. | Homebrew is a distribution candidate, not required for dry-run artifact readiness. | Tap, formula, sha256, audit, `brew install`, `ni --help`, and `ni version` proof. |
| cross-platform binary verification | Local build checked only this environment. | GoReleaser config defines target matrix, but hosted cross-platform artifacts are future release evidence. | Matrix archive build, checksum file, and platform-specific `ni --help` / `ni version` proof. |
| external install verification | No external user or separate host ran the install path in this audit. | Local dry-run/install checks passed. | User-run or host-specific install transcript with checksum, help, and version output. |
| model workspace host verification | Model workspace packs remain Experimental. | Model workspace availability is separate from artifact dry-run readiness. | Host-specific install/discovery proof and provider behavior transcript. |
| external user validation | No external users were observed. | Dry-run script/doc readiness can pass before adoption proof. | Maintained external validation notes, user-run transcripts, or scoped user research. |

## Warnings

| Warning | Evidence | Mitigation | Next action |
| --- | --- | --- | --- |
| `goreleaser` was not installed locally. | `command -v goreleaser` returned no path. | Record GoReleaser config as present but unchecked locally. | Run `goreleaser check` and, if appropriate, `goreleaser release --snapshot --clean` in a later environment with GoReleaser installed. |
| `0.0.0-dev` is correct for dry-run source builds but not a release version. | `go run ./cmd/ni version` and temporary binary `version` output were `0.0.0-dev`. | Treat this as local development output only. | Verify linker-injected version in final release artifact process. |
| `scripts/install-check.sh` creates `bin/ni`. | Install-check passed and uses `make build`. | Keep generated build output unstaged and report it if visible. | Clean or ignore local build output as needed before commit. |

## Risks

| Risk | Impact | Mitigation | Follow-up |
| --- | --- | --- | --- |
| dry-run is not publication | Maintainers could overread local checks as hosted release evidence. | This document states no publish/tag/release/upload occurred. | Use a separate publication checklist before any release action. |
| local build is not cross-platform verification | Platform-specific archive or install bugs could remain. | Keep cross-platform verification deferred. | Verify GoReleaser matrix artifacts and per-platform help/version output. |
| release-check coverage may not prove hosted artifact availability | Checks can pass while no GitHub release exists. | Distinguish local check gates from hosted artifact proof. | Verify release page assets and checksums after publication. |
| Homebrew remains Planned | Package-manager users still lack `brew install`. | Preserve `Homebrew: Planned / v0.5 candidate`. | Run Homebrew implementation audit if selected. |
| external user validation remains limited | Adoption friction may be unknown. | Keep external validation as a deferral. | Prepare external user validation plan if adoption proof becomes priority. |
| provider host behavior remains unverified | Model workspace behavior may vary by host. | Keep model workspace packs Experimental. | Run model workspace host verification audit if selected. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Published/released status | Do not claim v0.5 is published or released. | No v0.5 publication claim added. | Yes | This audit explicitly says no release action occurred. |
| Artifact upload status | Do not claim assets were uploaded. | No upload claim added. | Yes | Local temporary binary only. |
| GitHub release status | Do not claim a v0.5 GitHub release exists. | No GitHub release claim added. | Yes | Release workflow was inspected, not run. |
| Homebrew | Keep Homebrew: Planned / v0.5 candidate. | Preserved. | Yes | No formula or `brew install` claim. |
| Model workspace packs | Keep Model workspace packs: Experimental. | Preserved. | Yes | Host/global install remains unverified. |
| No-terminal | Keep No-terminal method: Experimental / assisted. | Preserved. | Yes | No deterministic no-terminal claim. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | Demo/check prompt compilation is not downstream execution. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | Preserved. | Yes | No implementation/adoption/cost/latency/downstream-performance claim. |
| Fixture relock | Fixture relock is not project-root relock. | Preserved. | Yes | Smoke/release-check fixture `ni end` and `ni relock` are separated. |
| Runtime execution boundary | `ni` is not a task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. | Preserved. | Yes | No runtime feature added. |

## Git status / generated file audit

| Path or group | Status from git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| docs/110_* | no output | No | The prior task's docs/110 files are not currently visible in git status; likely already clean or committed. |
| docs/111_* | no output | No | The prior task's docs/111 files are not currently visible in git status; likely already clean or committed. |
| docs/112_* | `M docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`; `M docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md` | Yes | Narrow follow-up pointers to docs/113 added. |
| docs/113_* | `?? docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md`; `?? docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md` | Yes | New artifact dry-run audit docs. |
| docs/51* | `M docs/51_POST_RELEASE_ROADMAP.md`; `M docs/51_POST_RELEASE_ROADMAP.ko.md` | Yes | Narrow roadmap pointers to docs/113 added. |
| generated artifacts, if any | no tracked output expected | No | Temporary `/tmp/ni-v0.5-artifact-dry-run` output was removed before final reporting; `bin/ni` may be created by install checks but should not be staged. |
| .ni/contract.json | no output | No | Protected project-root file unchanged. |
| .ni/session.json | no output | No | Protected project-root file unchanged. |
| .ni/plan.lock.json | no output | No | Protected project-root file unchanged. |
| unexpected files | none expected | No | Final git status should show only docs/51, docs/112, and docs/113 changes. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, and warnings are `None`. |
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

- `docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md`: added this English artifact
  dry-run audit.
- `docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md`: added Korean companion.
- `docs/51_POST_RELEASE_ROADMAP.md`: added a narrow docs/113 pointer.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: added matching Korean pointer.
- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`: added a narrow docs/113
  follow-up pointer.
- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`: added matching Korean
  follow-up pointer.

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

Why: the artifact dry-run passed with explicit deferrals, so the next safest
step is a non-executing publication checklist that lists the exact gated actions
and evidence required before a maintainer tags, publishes, uploads assets, or
updates availability claims.

Follow-up publication checklist:
[`114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md`](114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md)
records that non-executing checklist and keeps publication actions future-gated.

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
