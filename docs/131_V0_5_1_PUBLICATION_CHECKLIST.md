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
- Windows real-host execution: deferred on macOS-only development host
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Checklist goal

This checklist gates a later human-approved v0.5.1 release. It tells a
maintainer what to verify before, during, and after publication, but this task
does not publish, tag, create a GitHub release, upload assets, run release
workflows, run GoReleaser publish, create or publish a Homebrew formula, run
`ni end` on the project root, relock the project root, execute generated
prompts, add runtime execution behavior, or mark v0.5.1 as released.

Every release action below is marked `Run in this task? No`.

## Decision

V0_5_1_PUBLICATION_CHECKLIST_READY_WITH_NOTES

Justification: the publication checklist is complete and future release actions
are separated from check-only validation. Notes remain because v0.5.1 still
requires explicit human approval, v0.5.1 is not published, hosted artifacts and
public `install.sh` retrieval do not exist yet, Windows real-host execution is
deferred, Homebrew remains Planned / v0.5 candidate, external user validation
has not been performed, model workspace behavior remains Experimental, the
no-terminal method remains Experimental / assisted, and GoReleaser full matrix
dry-run remains deferred unless GoReleaser is available.

## Release-readiness chain

| Evidence | Decision / result | Pass? | Notes |
| --- | --- | --- | --- |
| docs/126 | PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED | Yes | Published v0.5.0 passes `ni --help` and `ni version`, but fails `ni init .`; v0.5.1 patch is needed. |
| docs/127 | V0_5_1_PATCH_PLAN_READY_WITH_NOTES | Yes | Patch scope, exclusions, validation matrix, and protected `.ni` boundary are documented. |
| docs/128 | V0_5_1_RC_VALIDATION_PASS_WITH_NOTES | Yes | Current-tree RC behavior and checker suite passed with publication, Homebrew, Windows, and external validation deferred. |
| docs/129 | V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES | Yes | Local darwin/arm64 release-like artifact reports `0.5.1`; local checksum verified; full GoReleaser matrix and hosted assets deferred. |
| docs/130 | V0_5_1_RELEASE_NOTES_READY_WITH_NOTES | Yes | Release notes draft is conservative, does not claim publication, and preserves known deferrals. |

## Pre-publication gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| repository state review | `git status --short`, recent log, `git diff --name-only v0.5.0..HEAD`, and `git diff --stat v0.5.0..HEAD` reviewed. | No | Yes | This checklist records the required gate; validation below records this task's check-only state. |
| protected `.ni` diff check | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` is empty. | No | Yes | Any protected diff blocks release execution. |
| validation commands | Go tests, install docs checks, skill pack checks, demo, quality, smoke, install, and release checks pass. | No | Yes | Re-run on the exact release commit. |
| release notes final review | docs/130 release notes draft reviewed against actual release artifacts and hosted asset inventory. | No | Yes | Do not publish notes that mention unavailable artifacts. |
| version injection confirmation | Release-built `ni version` reports `0.5.1`. | No | Yes | Source `go run ./cmd/ni version` may report `0.0.0-dev`; that is not release proof. |
| artifact dry-run confirmation | Archive names, `ni_0.5.1_checksums.txt`, and current-platform binary smoke are verified. | No | Yes | docs/129 covers local darwin/arm64 fallback; future task should use GoReleaser when available. |
| install.sh future URL check | `sh install.sh --dry-run --version 0.5.1` prints expected v0.5.1 URLs. | No | Yes | Dry-run URL construction is not public retrieval. |
| claim-boundary audit | README/docs do not claim hosted v0.5.1, Homebrew Available, Windows real-host proof, or `ni run` execution. | No | Yes | Use this document's audit table as the minimum boundary set. |
| human approval requirement | Maintainer explicitly chooses one listed approval option after this checklist. | No | Yes | This task does not select any option. |

## Tag gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| verify no existing v0.5.1 tag | `git tag --list v0.5.1` returns empty. | No | Yes | If the tag exists, verify it before continuing or abort. |
| create v0.5.1 tag | Annotated tag points to intended release commit. | No | Yes | Create only in a later explicitly approved release task. |
| verify tag commit | `git rev-parse v0.5.1^{}` or equivalent matches intended commit. | No | Yes | Do not assume the tag target from model judgment. |
| push tag | Remote tag exists and points to intended commit. | No | Yes | Push only in a later explicitly approved release task. |

## Artifact gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| build release artifacts | GoReleaser or approved release build creates expected v0.5.1 archives. | No | Yes | Build only in a later release task. |
| verify version output 0.5.1 | Extracted current-platform artifact runs `ni version` and prints `0.5.1`. | No | Yes | Required before publication claim. |
| generate checksums | `ni_0.5.1_checksums.txt` generated from intended artifacts. | No | Yes | Checksum file must match exact uploaded assets. |
| verify checksums | `shasum -a 256 -c ni_0.5.1_checksums.txt` or equivalent passes. | No | Yes | Mismatch aborts release. |
| verify current-platform artifact | Extracted binary runs `ni --help`, `ni version`, `ni init . --yes`, and `ni status --proof --next-questions` in temp workspaces. | No | Yes | Temp fixture status may be `BLOCKED`; that is expected for blank intent. |
| verify archive names | Archives match `.goreleaser.yaml` naming for v0.5.1. | No | Yes | Expected names include `ni_0.5.1_darwin_arm64.tar.gz` and `ni_0.5.1_checksums.txt`. |

## GitHub release gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| create GitHub release | GitHub release exists for `v0.5.1` at the intended tag. | No | Yes | Create only in a later explicitly approved release task. |
| attach assets | Expected archives are attached to the GitHub release. | No | Yes | Do not claim hosted artifacts before this is visible. |
| attach checksums | `ni_0.5.1_checksums.txt` is attached. | No | Yes | Required for installer and user verification. |
| verify draft status | Release is not draft unless intentionally kept draft. | No | Yes | If draft is intentional, docs must say so. |
| verify prerelease status | Release is not prerelease unless intentionally marked prerelease. | No | Yes | v0.5.1 patch is expected as a normal release unless maintainer decides otherwise. |
| verify release URL | Release URL opens and points to the intended tag/version. | No | Yes | Record URL in final release execution report. |

## Post-publication gate

| Check | Required evidence | Run in this task? | Future release task required? | Notes |
| --- | --- | --- | --- | --- |
| download hosted artifact | Hosted v0.5.1 archive downloads from the GitHub release. | No | Yes | Hosted asset availability starts only after publication. |
| verify checksum | Hosted artifact checksum matches hosted `ni_0.5.1_checksums.txt`. | No | Yes | Mismatch aborts public availability claim. |
| run hosted artifact `ni --help` | Extracted hosted artifact renders help. | No | Yes | Use command name or explicit extracted binary path and record which was used. |
| run hosted artifact `ni version` | Extracted hosted artifact prints `0.5.1`. | No | Yes | Required release proof. |
| isolated `install.sh --version 0.5.1` | Temporary HOME/BINDIR install retrieves hosted v0.5.1, verifies checksum, and installs a command-name `ni`. | No | Yes | Only after hosted assets exist. |
| isolated curl installer verification | Fresh shell `ni --help`, `ni version`, `ni init .`, and uninstall pass in temp locations. | No | Yes | Do not pollute user install state. |
| README / docs final check | README, README.ko, install docs, and release notes match actual hosted state. | No | Yes | Keep v0.5.0 parity note if needed until v0.5.1 proof exists. |
| Homebrew boundary check | No Homebrew Available claim unless separate Homebrew proof exists. | No | Yes | Homebrew remains Planned / v0.5 candidate by default. |

## Abort criteria

| Abort condition | Why it matters | Required action | Notes |
| --- | --- | --- | --- |
| validation failure | Release commit cannot be trusted. | Stop, fix, rerun check-only validation. | Do not tag around a failing gate. |
| version mismatch | Users may receive the wrong binary. | Stop, fix linker/version injection, rebuild artifacts. | `0.0.0-dev` is acceptable for source runs, not release artifacts. |
| checksum mismatch | Downloads cannot be safely verified. | Stop, regenerate artifacts/checksums from intended commit. | Do not upload or keep bad assets. |
| hosted asset unavailable | Installer and docs would point at missing files. | Stop public availability claims and fix release assets. | `install.sh` cannot retrieve missing assets. |
| install.sh v0.5.1 retrieval failure after publication | Public install path is broken. | Stop release announcement or correct docs until install succeeds. | Keep public retrieval unclaimed before proof. |
| protected `.ni` diff appears | Project-root planning state may have changed. | Stop and investigate; do not silently edit or relock root planning files. | Protected files are `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`. |
| Homebrew overclaim | Package-manager proof is absent. | Remove claim and keep Homebrew Planned / v0.5 candidate. | Homebrew requires separate tap/formula/install proof. |
| Windows real-host overclaim | Static checks are not Windows execution proof. | Remove claim until Windows transcript exists. | macOS-only development host cannot prove Windows real-host execution. |
| ni run execution overclaim | Violates ni-kernel boundary. | Remove claim and verify `ni run` remains bounded prompt compilation only. | `ni run` must not execute downstream work. |

## Human approval options

Actual v0.5.1 release execution requires explicit human approval after this
checklist. This task does not select any option.

- APPROVE_V0_5_1_RELEASE_EXECUTION
- DO_NOT_APPROVE_FIX_FIRST
- DEFER_V0_5_1_RELEASE

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks checklist? |
| --- | --- | --- | --- |
| v0.5.1 publication | This task is non-publishing and no human approval option is selected. | Explicit approval, tag, GitHub release, hosted assets, and post-publication checks. | No |
| hosted artifacts | No GitHub release or asset upload was performed. | Release page inventory, archive download, checksum verification. | No |
| install.sh actual v0.5.1 retrieval | Hosted v0.5.1 assets do not exist yet. | Isolated hosted install with checksum, `ni --help`, `ni version`, and `ni init .`. | No |
| Windows real-host execution | Current development host is macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | No tap/formula/install proof exists. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | No external user or separate machine transcript in this task. | External install/init/status transcript. | No |
| model workspace host behavior | Host-level/global install and provider behavior remain unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Trusted CLI proof for a target workspace. | No |
| GoReleaser full matrix dry-run if unavailable | Local GoReleaser availability can vary by host. | Environment with GoReleaser installed running check/dry-run matrix. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Preserved. | Yes | No publish/tag/release action occurred. |
| hosted artifacts | Must not claim assets exist before GitHub release upload. | Preserved. | Yes | docs/129 local artifact is not hosted proof. |
| install.sh v0.5.1 retrieval | Must not claim public retrieval before hosted assets exist and install passes. | Preserved. | Yes | Dry-run URL construction only. |
| Homebrew | Must remain Homebrew: Planned / v0.5 candidate. | Preserved. | Yes | No Homebrew Available claim. |
| Windows real-host execution | Must remain deferred without Windows transcript. | Preserved. | Yes | Static checks are not real-host execution. |
| ni init . | Guided project intent setup only. | Preserved. | Yes | It creates planning artifacts; it does not run agents. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | No generated prompt execution. |
| READY | CLI readiness only, not product readiness. | Preserved. | Yes | `ni status` is authority. |
| runtime execution boundary | No task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | No runtime behavior added. |
| human approval status | Must not be selected by this task. | Preserved. | Yes | Options are listed only. |

## Git status / inclusion check

| Path or group | git status --short | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | clean at task start; changed in `v0.5.0..HEAD` | Yes | Public onboarding and parity note. |
| README.ko.md | clean at task start; changed in `v0.5.0..HEAD` | Yes | Korean companion. |
| docs/126* | tracked | Yes | Public install parity evidence. |
| docs/127* | tracked | Yes | v0.5.1 patch plan. |
| docs/128* | tracked | Yes | v0.5.1 RC validation. |
| docs/129* | tracked | Yes | v0.5.1 artifact dry-run. |
| docs/130* | tracked | Yes | v0.5.1 release notes finalization. |
| docs/131* | added by this task | Yes | Publication checklist and Korean companion. |
| CHANGELOG.md | absent | No | Not added, to avoid implying publication. |
| RELEASE.md | absent | No | Not added, to avoid implying publication. |
| .ni/contract.json | no diff | No direct edit | Protected. |
| .ni/session.json | no diff | No direct edit | Protected. |
| .ni/plan.lock.json | no diff | No direct edit | Protected. |
| unexpected files | none at task start | No | Recheck after validation. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; only `docs/51_POST_RELEASE_ROADMAP*` modifications and new `docs/131*` files are visible. |
| `git log --oneline --decorate -20` | Checked; HEAD was `36609df Add v0.5.1 artifact and release notes roadmap links` before this task. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; no v0.5.1 tag exists. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Checked; docs/126 through docs/130 are tracked in current patch delta. |
| `git diff --stat v0.5.0..HEAD` | Checked; 70 files changed before docs/131. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty before edits and after validation. |
| Required ripgrep scans | Reviewed release, version, install, Homebrew, and runtime boundary surfaces. |
| `CHANGELOG.md` / `RELEASE.md` presence check | Both files are absent. |
| `.github/workflows/` presence check | `.github/workflows/ci.yml` and `.github/workflows/release.yml` are present. |
| release scripts / installers check | `install.sh`, `install.ps1`, `scripts/release-check.sh`, `scripts/install-check.sh`, and `scripts/release-dry-run.sh` are present. |
| `gofmt -w .` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root reported `NI Intent Readiness: READY` with no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output `0.0.0-dev`. |
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
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.md` | Added English v0.5.1 publication checklist, human approval options, gates, deferrals, and claim-boundary audit. |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.ko.md` | Added Korean companion with the same boundaries. |

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

Selection rationale: the checklist is ready with notes, but actual v0.5.1
release execution still requires explicit human approval. B is not selected
because the user has not approved release execution in this task.

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
