# v0.5.1 Release Notes Finalization

## Current status

State:
- v0.5.0 publication: verified.
- v0.5.1 post-release verification: see docs/132.
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES.
- v0.5.1 RC validation decision: V0_5_1_RC_VALIDATION_PASS_WITH_NOTES.
- v0.5.1 artifact dry-run decision: V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES.
- v0.5.1 release: not published.
- v0.5.1 tag: absent in this checkout at finalization time.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: deferred on the macOS-only development host.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Finalization goal

This document finalizes a conservative v0.5.1 release notes draft before any
publication checklist or release action. It does not publish v0.5.1, create a
tag, create a GitHub Release, upload assets, run release workflows, run
GoReleaser publish, create or publish a Homebrew formula, run `ni end` on the
project root, relock the project root, execute generated prompts, or add
runtime execution behavior.

## Decision

V0_5_1_RELEASE_NOTES_READY_WITH_NOTES

Post-release note: v0.5.1 was later published and verified in
[`132_V0_5_1_POST_RELEASE_VERIFICATION.md`](132_V0_5_1_POST_RELEASE_VERIFICATION.md).
The GitHub Release body was updated to a concise version of this notes draft.

Justification: the release notes are ready to describe the patch accurately and
conservatively. Notes remain because v0.5.1 is not published, hosted v0.5.1
assets do not exist, public `install.sh` retrieval cannot succeed until
publication, Windows real-host execution remains deferred, Homebrew remains
Planned / v0.5 candidate, external user validation has not been performed, and
the local GoReleaser full matrix dry-run remains deferred because GoReleaser is
unavailable locally.

## Release notes draft

### Summary

v0.5.1 is a patch release for public install parity. It packages the
current-tree onboarding fixes needed after published v0.5.0 passed
`ni --help` and `ni version` but failed the README-required `ni init .` first
project step.

### Why this patch exists

The current README teaches a first-user flow that starts an existing project
directory with `ni init .`. Published v0.5.0 can be installed and reports
version `0.5.0`, but `ni init .` fails with:

```text
unknown init option: .
```

v0.5.1 aligns the release artifact behavior with the current README onboarding
path.

### Fixes

- Fixes the public v0.5.0 install parity mismatch for `ni init .`.
- Makes `ni init .` work as the current-directory guided setup entry point.
- Keeps init bounded to planning workspace creation; it does not run agents or
  downstream work.

### Added

- `ni init .` positional target support.
- Bubble Tea v2 / Lip Gloss v2 guided init TUI.
- Non-interactive fallback for CI, scripts, and non-TTY contexts.
- Existing file protection during init.
- `.ni/plan.lock.json` protection during init.
- Post-TUI plain text summary so users can see created and unchanged files.
- Windows User PATH installer code and static safety checks.

### Changed

- README onboarding is aligned around macOS and Windows primary install paths.
- Install docs and checkers are aligned with command-name verification using
  `ni --help`, `ni version`, and `ni init .`.
- Domain init logic is separated from TUI rendering so file planning remains
  testable and deterministic.

### Validation

- Current-tree tests and release checks pass under the documented validation
  commands.
- A local release-like darwin/arm64 artifact built with
  `ni/internal/version.Version=0.5.1` reports `0.5.1`.
- The local current-platform archive checksum was generated and verified.
- The local release-like artifact passed `ni --help`, `ni version`,
  `ni init . --yes`, and `ni status --proof --next-questions` smoke checks in
  temporary workspaces.
- Audit evidence is recorded in docs/126, docs/127, docs/128, and docs/129.

### Known deferrals

- v0.5.1 has not been published.
- Hosted v0.5.1 artifacts and checksums do not exist yet.
- `install.sh` public retrieval of v0.5.1 cannot be claimed until hosted
  assets exist and a post-publication install check passes.
- Windows real-host execution remains deferred until a Windows transcript
  exists.
- Homebrew remains Planned / v0.5 candidate.
- External user validation has not been performed.
- Model workspace packs remain Experimental.
- No-terminal method remains Experimental / assisted.
- Local GoReleaser full matrix dry-run remains deferred because GoReleaser is
  unavailable locally.

### What this patch does not do

- It does not publish v0.5.1 by itself.
- It does not add a task runner, SPEC runner, execution harness, shell adapter,
  Codex exec adapter, queue, PR automation, release automation, or downstream
  execution layer.
- It does not make `ni run` execute downstream work.
- It does not mark Homebrew Available.
- It does not verify Windows real-host execution.
- It does not prove external user success.
- It does not make no-terminal deterministic.
- It does not claim benchmark evidence proves implementation correctness or
  downstream execution quality.

### Upgrade/install note

After v0.5.1 is actually published and hosted assets are verified, users should
install the published release through the documented installer path and verify:

```bash
ni --help
ni version
ni init .
```

Until publication, keep README and install guidance honest: v0.5.0 is the
verified public release, and v0.5.1 remains a prepared patch delta.

### Maintainer verification checklist

- Confirm the intended release commit and clean working tree.
- Confirm `v0.5.1` tag absence before creating any authorized release tag.
- Run the full validation gate.
- Verify release build version injection reports `0.5.1`.
- Verify hosted archive inventory and checksums after publication.
- Verify current-platform command-name `ni --help` and `ni version`.
- Verify hosted `install.sh --version 0.5.1` install in temporary HOME/BINDIR.
- Verify installed `ni init .` and `ni status --proof --next-questions`.
- Keep Windows real-host execution, Homebrew, external validation, model
  workspace, no-terminal, and downstream execution deferrals explicit unless
  separate evidence exists.

## Patch rationale

| Issue | Evidence | v0.5.1 response | Notes |
| --- | --- | --- | --- |
| v0.5.0 `ni init .` mismatch | docs/126 records published v0.5.0 passing `ni --help` and `ni version` but failing `ni init .` with `unknown init option: .`. | Add `ni init .` positional target support to the release delta. | This is the main patch reason. |
| README onboarding parity | Current README uses `ni init .` as the first project path. | Align the release artifact with README onboarding. | Do not hide the onboarding path to make v0.5.0 look complete. |
| First-user TUI | docs/124 and docs/128 record Bubble Tea v2 / Lip Gloss v2 guided init behavior. | Include the guided init TUI in v0.5.1. | TUI collects intent; it is not readiness authority. |
| Non-interactive fallback | docs/128 and docs/129 record `ni init . --yes` smoke checks. | Preserve fallback behavior for CI and non-TTY contexts. | No TUI requirement for automation. |
| Install docs/checkers | README, install docs, and scripts align command-name verification. | Keep docs/checkers synchronized with the release notes. | Windows real-host execution stays deferred. |

## Release notes claim audit

| Claim area | Expected boundary | Observed wording | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Draft says v0.5.1 has not been published. | Yes | No release action occurred. |
| v0.5.1 artifacts | Must not claim hosted artifacts exist. | Draft says hosted v0.5.1 assets do not exist yet. | Yes | Local artifact dry-run is separate. |
| `install.sh` v0.5.1 retrieval | Must not claim public retrieval before publication. | Draft says retrieval cannot be claimed until hosted assets exist. | Yes | Dry-run URL construction is not public retrieval. |
| Homebrew | Must remain Homebrew: Planned / v0.5 candidate. | Draft preserves Planned / v0.5 candidate and must not claim Homebrew Available. | Yes | No tap/formula/install proof. |
| Windows real-host execution | Must remain deferred without transcript. | Draft says Windows real-host execution remains deferred. | Yes | Static checks are not real-host proof. |
| `ni init .` | May claim v0.5.1 patch support, not v0.5.0 parity. | Draft ties support to the v0.5.1 patch delta. | Yes | v0.5.0 mismatch remains explicit. |
| `ni run` | Bounded prompt compilation only. | Draft says the patch does not make `ni run` execute downstream work. | Yes | No downstream execution claim. |
| READY | Must come from `ni status`, not model judgment. | Draft does not use READY as product readiness. | Yes | CLI remains authority. |
| Benchmark evidence | Must not claim implementation correctness or downstream execution quality. | Draft explicitly avoids this claim. | Yes | Benchmark overclaim avoided. |
| Runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, or downstream execution layer. | Draft lists these as excluded. | Yes | Kernel boundary preserved. |

## Validation evidence

| Evidence | Result | Notes |
| --- | --- | --- |
| docs/126 public install parity | PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED. | v0.5.0 help/version passed, `ni init .` failed. |
| docs/127 patch plan | V0_5_1_PATCH_PLAN_READY_WITH_NOTES. | Patch scope and exclusions are documented. |
| docs/128 RC validation | V0_5_1_RC_VALIDATION_PASS_WITH_NOTES. | Current-tree behavior and docs/checkers passed. |
| docs/129 artifact dry-run | V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES. | Local current-platform artifact version/checksum/init smoke passed; GoReleaser matrix deferred. |
| current validation commands | See validation results in this document. | Includes Go tests, docs checks, smoke, install, release, and quality gates. |
| protected `.ni` diff | Empty. | `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` unchanged. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks release notes? |
| --- | --- | --- | --- |
| v0.5.1 publication | This task is non-publishing. | Authorized tag, release workflow, GitHub Release, hosted asset proof. | No |
| Hosted artifacts | No release page/assets exist for v0.5.1 yet. | Release page inventory and checksum verification. | No |
| `install.sh` actual v0.5.1 retrieval | Hosted assets do not exist before publication. | Isolated hosted install, checksum verification, `ni --help`, `ni version`, `ni init .`. | No |
| Windows real-host execution | Current development host is macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | No tap/formula/install proof. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| External user validation | No separate user/machine transcript in this task. | External install/init/status transcript. | No |
| Model workspace host behavior | Host-level/global install and provider behavior remain unverified. | Host-specific discovery/install/provider transcript. | No |
| No-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Trusted CLI proof for a target workspace. | No |
| GoReleaser full matrix dry-run | GoReleaser is unavailable locally. | Environment with GoReleaser installed running check/dry-run matrix. | No |

## Git status / inclusion check

| Path or group | `git status --short` | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | clean relative to HEAD at task start; changed in `v0.5.0..HEAD`. | Yes | Onboarding and public parity note. |
| README.ko.md | clean relative to HEAD at task start; changed in `v0.5.0..HEAD`. | Yes | Korean companion onboarding and parity note. |
| docs/126* | tracked. | Yes | Public install parity evidence. |
| docs/127* | tracked. | Yes | Patch release plan. |
| docs/128* | tracked. | Yes | Release-candidate validation. |
| docs/129* | tracked. | Yes | Artifact dry-run evidence. |
| docs/130* | added by this task. | Yes | Release notes finalization and Korean companion. |
| CHANGELOG.md | absent. | No | Not added to avoid release-history confusion. |
| RELEASE.md | absent. | No | Not added to avoid implying publication. |
| `.ni/contract.json` | no diff. | No direct edit | Protected. |
| `.ni/session.json` | no diff. | No direct edit | Protected. |
| `.ni/plan.lock.json` | no diff. | No direct edit | Protected. |
| unexpected files | none expected. | No | Rechecked after validation. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Shows only `docs/51_POST_RELEASE_ROADMAP*` modifications and new `docs/130*` files. |
| `git log --oneline --decorate -20` | Checked before editing; current HEAD was `e23653b Add v0.5.1 artifact dry-run docs`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; no v0.5.1 tag exists. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Checked; docs/126 through docs/129 are tracked in the current patch delta. |
| `git diff --stat v0.5.0..HEAD` | Checked; 68 files changed before docs/130. |
| Required ripgrep scans | Reviewed release, version, install, Homebrew, and runtime boundary surfaces; missing CHANGELOG.md and RELEASE.md were reported by ripgrep because they do not exist. |
| `gofmt -w .` | Passed; no Go source change was introduced by this task. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root reported `NI Intent Readiness: READY` with no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help includes `ni init [.]`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

## Changes made

| File | Why |
| --- | --- |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md` | Added English release notes finalization, release notes draft, claim-boundary audit, deferrals, validation plan, and next task prompt. |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.ko.md` | Added Korean companion with matching boundaries. |
| `docs/51_POST_RELEASE_ROADMAP.md` | Added a narrow pointer to docs/129 and docs/130. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Added matching Korean pointers. |

## What this finalization proves

- v0.5.1 release notes are ready under audited boundaries.
- The notes describe the patch without claiming publication.
- Known deferrals remain explicit.
- No release action was performed.

## What this finalization does not prove

- v0.5.1 has been published.
- v0.5.1 artifacts are hosted.
- `install.sh` retrieves v0.5.1 publicly.
- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

A. v0.5.1 publication checklist

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Goal:
Create a non-publishing v0.5.1 publication checklist now that the release notes
are finalized with notes.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- docs/127_V0_5_1_PATCH_RELEASE_PLAN.md
- docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.md
- docs/129_V0_5_1_ARTIFACT_DRY_RUN.md
- docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md
- docs/22_INSTALL.md
- docs/install-curl.md
- install.sh
- install.ps1
- .goreleaser.yaml
- .github/workflows/release.yml
- scripts/release-check.sh

Task:
Create a concise v0.5.1 publication checklist that separates:
- pre-publication validation gates
- human approval gates
- tag and GitHub Release actions, described as future maintainer actions only
- artifact and checksum verification
- post-publication install.sh hosted retrieval verification
- current-platform command-name help/version/init smoke
- rollback and docs correction paths
- known deferrals for Windows real-host execution, Homebrew, external user
  validation, model workspace packs, no-terminal, and GoReleaser matrix proof

Decision:
Use exactly one:
- V0_5_1_PUBLICATION_CHECKLIST_READY
- V0_5_1_PUBLICATION_CHECKLIST_READY_WITH_NOTES
- V0_5_1_PUBLICATION_CHECKLIST_BLOCKED

Rules:
- Do not publish, tag, create a GitHub Release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  ni end on the project root, relock the project root, execute generated
  prompts, add runtime execution behavior, or mark v0.5.1 as released.
- Do not mark Homebrew Available.
- Do not claim hosted v0.5.1 assets exist before publication.
- Do not claim install.sh retrieves v0.5.1 publicly before hosted assets exist.
- Do not claim Windows real-host execution verified unless a Windows transcript
  exists.
- Keep Skills are UX; CLI is authority.
- Keep ni run as bounded prompt compilation only.

Validation:
- git status --short
- git tag --list v0.5.1
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh

Final response:
Report changed files, checklist decision, validation results, protected .ni
diff, known deferrals, and confirmation that no publish/tag/release/upload/root
relock/generated prompt execution occurred.
```
