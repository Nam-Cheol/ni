# v0.5 Release Publication Checklist

## Current status

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
- 이 checklist는 publish, tag, release, GitHub release creation, asset upload,
  v0.5 released claim을 수행하지 않는다.

## Checklist goal

이 문서는 future human-gated publication checklist다. Maintainer가 later v0.5
publication 전에 확인해야 하는 checks, manual gates, verification evidence,
abort criteria, forbidden claims를 문서화한다. 이것은 publication action, release
workflow, tag creation, asset upload, Homebrew publication, availability upgrade가
아니다.

## Checklist readiness decision

Decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES.

Justification: future publication path가 documented되었고, release-readiness
chain이 visible하며, required gates가 이 task에서 실행하지 않은 actions와 분리되어
있다. Notes는 publication이 future-gated이고, docs/114가 new이며, docs/113과
docs/51에 narrow pointers가 추가되고, Homebrew가 Planned / v0.5 candidate로 남고,
hosted artifacts, checksums, cross-platform installs, external user validation이
deferred로 남기 때문이다.

## Release-readiness chain

| Doc | Decision / role | Status | Notes |
| --- | --- | --- | --- |
| docs/110 | RC_READY_WITH_DEFERRALS | Present and tracked | v0.5 release-candidate readiness audit를 기록하며 release completion을 claim하지 않는다. |
| docs/111 | release notes draft | Present and tracked | Draft-only release notes; v0.5 publication, tag, upload, Homebrew availability claim 없음. |
| docs/112 | RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES | Present and tracked | Final preflight는 release-note and no-release boundaries를 preserve한다. |
| docs/113 | ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS | Present and tracked | Dry-run/check-only artifact readiness passed with explicit deferrals. |
| docs/114 | PUBLICATION_CHECKLIST_READY_WITH_NOTES | New in this task | Non-executing future publication checklist; publication actions were not run. |

## Pre-publication gate

| Check | Command or evidence | Required result | Status in this task | Notes |
| --- | --- | --- | --- | --- |
| git status review | `git status --short` | Expected docs-only changes visible; unexpected files 없음. | Run; pass | Initial status was clean; final status is documented below. |
| protected .ni diff | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Project-root protected `.ni` diff 없음. | Run; pass | 이 task는 root `.ni` files를 edit하지 않았다. |
| ni status proof | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | `NI Intent Readiness: READY`; blockers, deferrals, warnings 없음. | Run; pass | CLI remains authority. |
| Go tests | `GOCACHE=/private/tmp/ni-go-cache go test ./...` | All packages pass. | Run; pass | Temp Go cache workaround 사용. |
| install docs check | `python3 scripts/check-install-docs.py` | Install/distribution claim markers pass. | Run; pass | Homebrew and release/curl wording을 bounded로 유지한다. |
| skill pack check | `bash scripts/check-skill-packs.sh` | Experimental status and CLI-authority checks pass. | Run; pass | Skills are UX; CLI is authority. |
| demo check | `bash scripts/demo-check.sh` | Demo, benchmark, no-terminal, seed-only checks pass. | Run; pass | Temporary compiled prompts are not executed. |
| quality check | `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Broad static/test/smoke wrapper passes. | Run; pass | Fixture `ni end`를 exercise할 수 있으나 project-root relock이 아니다. |
| smoke check | `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Public CLI smoke paths pass. | Run; pass | Temporary fixture `ni end` and `ni relock` only. |
| install check | `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Source, build, temporary local install paths pass. | Run; pass | Public release install은 수행하지 않았다. |
| release check | `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Check-only release readiness gate passes. | Run; pass | Release action 없음. |
| ni --help | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Help output renders and names `ni` as a project intent compiler. | Run; pass | Downstream work를 execute하지 않는다. |
| ni version | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Development output is explicit. | Run; pass | Output: `0.0.0-dev`; v0.5 release version claim이 아니다. |
| temporary local build help/version | Covered by docs/113 artifact dry-run evidence. | Temporary binary help/version passed in docs/113. | Reviewed; not rerun here | 이 task는 release artifacts를 만들지 않는다. |
| release notes draft review | docs/111 and docs/112 | Draft-only and preflight boundaries preserved. | Reviewed; pass | v0.5 publication claim 없음. |
| no-overclaim review | docs/51, docs/112, docs/113, docs/114 | Forbidden claim introduced 없음. | Run; pass | See claim-boundary audit. |

## Publication actions - future gated steps, not run

| Action | Future command or UI action | Required precondition | Verification required after action | Run in this task? |
| --- | --- | --- | --- | --- |
| create release tag | Future manual `git tag -a v0.5.0 -m "..."` or maintainer-approved equivalent | Clean working tree, passing validation, human approval packet, final version decision | Tag points to intended commit and does not include unapproved changes | No |
| push release tag | Future manual `git push origin v0.5.0` | Correct local tag and explicit maintainer approval | Remote tag exists and points to intended commit | No |
| create GitHub release | Future GitHub UI or tag-triggered release workflow | Tag pushed, release notes approved, no protected-file drift | GitHub release page exists and matches intended tag | No |
| upload release assets | Future GoReleaser release workflow or maintainer-controlled upload | Release workflow approved; artifact matrix and config reviewed | Expected archives are present and downloadable | No |
| publish checksums | Future GoReleaser checksum asset or manual checksum upload | Assets generated from intended commit and version | `ni_<version>_checksums.txt` matches each hosted archive | No |
| verify hosted assets | Future hosted download checks | GitHub release exists and assets uploaded | Downloads succeed and checksums match | No |
| update release notes from draft | Future edit from docs/111 draft into release notes | Hosted assets and checksums verified | Release notes do not overclaim Homebrew, model workspace, no-terminal, or benchmark evidence | No |
| update public docs after publication | Future README/install/release-status update | Actual v0.5 publication and hosted verification complete | Docs reflect exactly what was published | No |
| optional Homebrew formula work, still deferred unless separately verified | Future Homebrew tap/formula task | Tap, formula, sha256, `brew install`, `ni --help`, and `ni version` verification plan | Homebrew is only marked Available after all Homebrew evidence passes | No |

## Artifact and checksum gate

| Artifact / checksum item | Required evidence before publication | Required evidence after publication | Notes |
| --- | --- | --- | --- |
| current-platform binary | Local build or release-candidate binary runs `ni --help` and `ni version` with intended version behavior. | Hosted current-platform asset downloads, extracts, runs `ni --help`, and runs `ni version`. | docs/113 only proves dry-run source/local behavior. |
| cross-platform binaries if applicable | GoReleaser config and matrix reviewed; snapshot check if available. | Each expected hosted archive is present and can be downloaded; platform-specific smoke verification where feasible. | Cross-platform install remains deferred in this task. |
| checksums / sha256 | Checksum generation path reviewed and expected filename known. | `ni_<version>_checksums.txt` matches all hosted archives. | Hosted checksums are not available until publication. |
| hosted download URLs | Expected GitHub release URL and asset names reviewed. | URLs are reachable and point to the intended tag/version. | No hosted v0.5 URLs were verified here. |
| curl installer target | Installer target version and asset naming reviewed before use. | `install.sh --version <version>` retrieves intended version and verifies checksum when available. | Future post-publication check only. |
| release notes asset list | Draft lists only assets that should exist. | Published release notes match actual hosted assets. | Do not publish asset claims before upload verification. |

## Install verification gate

| Install path | Current status | Future verification required | Can claim Available now? | Notes |
| --- | --- | --- | --- | --- |
| release binary | Available for verified v0.4.0 release assets | v0.5 release page, downloadable archives, checksums, extraction, `ni --help`, and `ni version` | Yes, only for verified v0.4.0 assets; No for v0.5 until publication evidence exists | This checklist does not publish v0.5 assets. |
| curl installer | Available for verified v0.4.0 release assets | `install.sh --dry-run --version <v0.5>`, real temporary install, checksum verification, `ni --help`, `ni version` | Yes, only for verified v0.4.0 assets; No for v0.5 until publication evidence exists | Future post-publication verification required. |
| Homebrew | Planned / v0.5 candidate | Tap, formula, sha256, `brew install` output, `ni --help`, and `ni version` | No | This checklist does not create or publish a Homebrew formula. |

## Homebrew deferral gate

- Homebrew remains Planned / v0.5 candidate.
- Do not claim Homebrew Available until tap, formula, sha256, brew install
  output, ni --help, and ni version are verified.
- This checklist does not create or publish a Homebrew formula.

## Post-publication verification gate

These checks are future post-publication checks and were not performed in this
task:

- GitHub release page exists.
- Assets are downloadable.
- Checksums match.
- Current-platform binary runs `ni --help`.
- Current-platform binary runs `ni version`.
- Curl installer retrieves intended version.
- Docs match published assets.
- No Homebrew Available claim unless Homebrew verification is complete.
- No benchmark, no-terminal, model workspace, or ni-run overclaim introduced.

## Rollback / abort criteria

| Condition | Abort or rollback action | Reason | Notes |
| --- | --- | --- | --- |
| validation failure | Abort before tag or publication. | Release state is not trustworthy. | Fix and rerun check-only validation. |
| tag mismatch | Delete or correct only under maintainer-approved release recovery policy. | Wrong commit could publish wrong intent. | Do not proceed with assets until tag is verified. |
| asset checksum mismatch | Abort publication or remove bad assets. | Users cannot verify downloads. | Regenerate from intended commit and version. |
| hosted artifact unavailable | Pause release announcement and fix hosted assets. | Docs would point to unavailable artifacts. | Do not update public docs to say v0.5 is available. |
| curl installer mismatch | Abort curl availability claim. | Installer could retrieve the wrong asset/version. | Keep curl installer v0.5 claim pending. |
| version/help mismatch | Abort artifact availability claim. | Binary may be wrong build or wrong version. | Verify linker flags and artifact source. |
| accidental overclaim in README/docs | Revert or correct docs before publication. | Public docs could promise unsupported status. | Recheck no-overclaim boundaries. |
| protected .ni file modified unexpectedly | Stop and investigate before any release step. | Root planning state may have changed. | Do not silently edit `.ni/contract.json`, `.ni/session.json`, or `.ni/plan.lock.json`. |
| Homebrew availability claim without verification | Remove claim and keep Homebrew Planned. | Homebrew evidence is absent. | Requires tap, formula, sha256, `brew install`, `ni --help`, and `ni version`. |
| generated prompt executed accidentally | Stop and document incident. | `ni run` is prompt compilation only. | Release process must not become downstream execution. |
| downstream execution behavior introduced accidentally | Abort release checklist and remove runtime behavior. | Violates `ni-kernel` boundary. | No task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. |

## Forbidden claims

- v0.5 has been released, unless actually released in a later task
- GitHub release exists, unless actually created
- assets were uploaded, unless actually uploaded
- Homebrew is Available, unless verified
- model workspace packs are Available
- no-terminal deterministic validation
- ni run executes downstream work
- benchmark proves implementation quality or downstream execution quality
- fixture relock is project-root relock
- validation-script fixture relock is project-root relock

## Git status / inclusion check

| Path or group | git status --short | git ls-files / tracked check | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| docs/110_* | no output | tracked | No new change | Present as expected. |
| docs/111_* | no output | tracked | No new change | Present as expected. |
| docs/112_* | no output | tracked | No new change | Present as expected; no docs/112 edit was needed. |
| docs/113_* | `M docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md`; `M docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md` | tracked | Yes | Narrow docs/114 follow-up pointer added. |
| docs/114_* | `?? docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md`; `?? docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md` | untracked until added | Yes | New publication checklist docs. |
| docs/51* | `M docs/51_POST_RELEASE_ROADMAP.md`; `M docs/51_POST_RELEASE_ROADMAP.ko.md` | tracked | Yes | Narrow roadmap pointers to docs/114 added. |
| generated artifacts | no output from `git status --short`; ignored `dist/` exists | ignored by `.gitignore`; tracked files 없음 | No | `dist/ni-codex-skills.zip` and `dist/ni-claude-skills.zip` are ignored skill-pack artifacts not created by this task. |
| .ni/contract.json | no output | tracked protected file | No | No diff. |
| .ni/session.json | no output | tracked protected file | No | No diff. |
| .ni/plan.lock.json | no output | tracked protected file | No | No diff. |
| unexpected files | none | n/a | No | Final status should show only docs/51, docs/113, and docs/114 changes. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Published/released status | Do not claim v0.5 is published or released. | Checklist says publication is future-gated and not run. | Yes | No v0.5 released claim. |
| Artifact upload status | Do not claim release assets were uploaded. | Upload is listed as a future gated step. | Yes | Run in this task: No. |
| GitHub release status | Do not claim a v0.5 GitHub release exists. | GitHub release creation is future-gated. | Yes | No release page verification was claimed. |
| Homebrew | Keep Homebrew: Planned / v0.5 candidate. | Preserved. | Yes | Homebrew Available remains forbidden without full evidence. |
| Model workspace packs | Keep Model workspace packs: Experimental. | Preserved. | Yes | No host/global provider verification claim. |
| No-terminal | Keep No-terminal method: Experimental / assisted. | Preserved. | Yes | No deterministic no-terminal validation claim. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | No generated prompt was executed. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | Preserved. | Yes | No implementation-quality or downstream-execution-quality claim. |
| Fixture relock | Fixture relock is separate from project-root relock. | Preserved. | Yes | Validation fixture `ni end` is not project-root relock. |
| Runtime execution boundary | `ni` is not a task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation, release automation, downstream execution layer, or execution evidence loop. | Preserved. | Yes | No runtime behavior added. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean; final status documented expected docs-only changes. |
| `git ls-files docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md` | Pass | docs/110 through docs/113 are tracked. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, warnings are `None`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass | Help output rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass | Output: `0.0.0-dev`; development output only. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack checks passed; global install remains unverified. |
| `bash scripts/demo-check.sh` | Pass | Public demos verified without downstream runtime execution. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Fixture `ni end` and `ni relock` paths exercised only in temporary workspaces. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source, build, and temporary install paths passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Release readiness gate passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` file diff. |

## Changes made

- `docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md`: English non-executing
  publication checklist 추가.
- `docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md`: Korean companion 추가.
- `docs/51_POST_RELEASE_ROADMAP.md`: narrow docs/114 pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer 추가.
- `docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md`: narrow docs/114 follow-up pointer
  추가.
- `docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md`: matching Korean follow-up
  pointer 추가.

## What this checklist proves

- a future publication path has been documented
- publication actions are gated and explicitly not run
- known deferrals remain explicit
- no release action was performed
- expected release-readiness docs are visible to git status / tracked checks as
  documented

## What this checklist does not prove

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

Selected next task: A. v0.5 publication human approval packet.

Why: checklist가 ready with notes 상태이므로, actual publication 전에 다음 safe step은
human approval packet이다. 이 packet은 tag, publish, upload assets, availability
claims update 없이 gates, required evidence, explicit deferrals, exact manual
approval questions를 다시 정리해야 한다.

## Next task prompt

```text
Proceed with a v0.5 publication human approval packet in /Users/namba/Documents/project/ni.

This is a documentation and approval-packet task only. Do not publish, tag, create a GitHub release, upload assets, run a release workflow, run goreleaser publish, create or publish a Homebrew formula, run ni end on the project root, relock the project root, execute generated prompts, add release automation, add runtime execution behavior, or mark v0.5 as released.

Use these docs as the current release-readiness chain:
- docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md
- docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md
- docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md
- docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md
- docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md

Preserve:
- RC decision: RC_READY_WITH_DEFERRALS
- Release notes preflight decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES
- Artifact dry-run decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS
- Publication checklist decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES
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
Create a human approval packet that a maintainer can read before deciding whether to perform a separate real v0.5 publication task. The packet must list exact approval questions, required evidence, manual release gates, abort criteria, known deferrals, and forbidden overclaims. It must not perform the release.

Required output:
- docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.md
- docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md if Korean companion docs are maintained
- narrow roadmap / docs/114 cross-links only if useful

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
