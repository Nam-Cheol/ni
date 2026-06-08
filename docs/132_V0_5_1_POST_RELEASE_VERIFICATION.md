# v0.5.1 Post-Release Verification

## Current status

State:
- v0.5.1 release: published and verified in this document.
- v0.5.0 publication: verified.
- Public install parity mismatch: addressed by v0.5.1 for the tested macOS arm64 path.
- v0.5.1 release notes decision: V0_5_1_RELEASE_NOTES_READY_WITH_NOTES.
- v0.5.1 publication checklist decision: V0_5_1_PUBLICATION_CHECKLIST_READY_WITH_NOTES.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: deferred on macOS-only development host.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Verification goal

This document verifies the actual v0.5.1 publication and public install
behavior after the approved release execution. It separates hosted release
evidence from source-tree validation and keeps remaining deferrals explicit.

## Decision

V0_5_1_RELEASE_EXECUTED_WITH_NOTES

Justification: the tag, GitHub Release, hosted assets, checksums, current
platform hosted artifact, and isolated `install.sh --version 0.5.1` path all
verified. Notes remain because Windows real-host execution, Homebrew
availability, external user validation, model workspace host behavior,
no-terminal deterministic validation, and local GoReleaser full-matrix dry-run
are still not claimed.

## Release identity

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| tag | `v0.5.1` | Local tag exists; remote annotated tag dereferences to `945bd46fdf94872ba921c63d4a2fc43ea54de3be`. | Yes | Remote tag object was `b588f6b2e13111841081d186bd0e70d3c0bfbd6c`. |
| release URL | `https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1` | `https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1` | Yes | Verified with `gh release view`. |
| release title | `v0.5.1` | `v0.5.1` | Yes | Normal GitHub Release. |
| draft status | Not draft | `false` | Yes | Verified with GitHub release metadata. |
| prerelease status | Not prerelease | `false` | Yes | Verified with GitHub release metadata. |
| latest status | v0.5.1 if intended | Default `gh release view` returned `v0.5.1`. | Yes | Latest lookup returned the v0.5.1 URL. |
| published time | Present | `2026-06-08T00:50:40Z` | Yes | Absolute timestamp from GitHub. |
| release commit | Intended release commit | `945bd46fdf94872ba921c63d4a2fc43ea54de3be` | Yes | `main` HEAD at release execution. |
| release body | Concise docs/130-based notes | GitHub Release body updated after GoReleaser default changelog. | Yes | Body preserves Homebrew, Windows, model workspace, no-terminal, and runtime boundaries. |

## Asset inventory

| Asset | Platform | Size | Downloaded? | Checksum verified? | Execution verified? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `ni_0.5.1_checksums.txt` | checksums | 471 | Yes | n/a | n/a | Hosted checksum file. |
| `ni_0.5.1_darwin_amd64.tar.gz` | macOS amd64 | 1971271 | Yes | Yes | Not run | Non-current platform on this host. |
| `ni_0.5.1_darwin_arm64.tar.gz` | macOS arm64 | 1845991 | Yes | Yes | Yes | Current platform artifact on this host. |
| `ni_0.5.1_linux_amd64.tar.gz` | Linux amd64 | 1980921 | Yes | Yes | Not run | Cross-platform execution not claimed. |
| `ni_0.5.1_linux_arm64.tar.gz` | Linux arm64 | 1812172 | Yes | Yes | Not run | Cross-platform execution not claimed. |
| `ni_0.5.1_windows_amd64.zip` | Windows amd64 | 2075537 | Yes | Yes | Not run | Windows real-host execution remains deferred. |

## Checksum verification

| File | Expected checksum | Actual checksum | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni_0.5.1_darwin_amd64.tar.gz` | `19459e19888a5043a1ee70e76696ad04990c35e987a35dbe5fab1472f67f1958` | Same via `shasum -a 256 -c` | Yes | Hosted asset matched checksum file. |
| `ni_0.5.1_darwin_arm64.tar.gz` | `bb97c7831cb4c906b6140a4f05e01823b23111c9f917996d6e0874c23f158c73` | Same via `shasum -a 256 -c` | Yes | Current platform archive. |
| `ni_0.5.1_linux_amd64.tar.gz` | `cdb4aca68beb5ae2397215c898bd135546f57ff140046e900632339bc53b35d1` | Same via `shasum -a 256 -c` | Yes | Execution not claimed. |
| `ni_0.5.1_linux_arm64.tar.gz` | `29c3030e18152f0a4877c9d665d833e049d5b0498e9dbd3550a28777ff4bc1c4` | Same via `shasum -a 256 -c` | Yes | Execution not claimed. |
| `ni_0.5.1_windows_amd64.zip` | `c44803197b8c0ae5401de71794f78b58708d639d0d25ca2b41f2d6e414ad435a` | Same via `shasum -a 256 -c` | Yes | Windows execution not claimed. |

## Hosted artifact verification

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| `ni --help` | `/tmp/ni-v0.5.1-post-release-verify/extracted-darwin-arm64/ni --help` | Help exits 0 and includes `ni init [.]`. | Help exited 0 and listed `ni init [.]`. | Yes | Hosted darwin/arm64 artifact. |
| `ni version` | `/tmp/ni-v0.5.1-post-release-verify/extracted-darwin-arm64/ni version` | `0.5.1` | `0.5.1` | Yes | Release linker version verified. |
| `ni init . --yes` | Hosted binary in temp project | Planning files created. | Created `.ni/contract.json`, `.ni/session.json`, `.ni/readiness.*`, `.ni/pressure.json`, `.ni/harness.candidates.json`, and `docs/plan/**`. | Yes | No downstream execution. |
| `ni status --proof --next-questions` | Hosted binary in initialized temp project | Honest first-run `BLOCKED` for empty intent. | `NI Intent Readiness: BLOCKED` with R014, OQ-001, R015, and R016 blockers. | Yes | Expected first-user state. |

## install.sh verification

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| `install.sh --version 0.5.1` | `HOME=/tmp/... BINDIR=/tmp/... sh install.sh --version 0.5.1` | Download hosted darwin/arm64 asset, verify checksum, install temp `ni`. | Downloaded `ni_0.5.1_darwin_arm64.tar.gz`, verified checksum, installed temp binary. | Yes | User install state was not touched. |
| command-name `ni --help` | Fresh `zsh -f` with temp BINDIR first on PATH | `ni` resolves from temp BINDIR and help exits 0. | `command -v ni` printed temp path; help exited 0. | Yes | Command-name verification. |
| command-name `ni version` | Fresh `zsh -f` with temp BINDIR first on PATH | `0.5.1` | `0.5.1` | Yes | Installed artifact verified. |
| installed `ni init . --yes` | Temp project with temp BINDIR on PATH | Planning files created. | Init created expected planning files. | Yes | Confirms public first-project path for tested host. |
| installed `ni status --proof --next-questions` | Temp installed binary in temp project | Honest first-run `BLOCKED`. | `BLOCKED` with first-run blockers and next questions. | Yes | Expected blank-intent result. |
| uninstall | Same temp HOME/BINDIR with `sh install.sh --uninstall` | Remove temp binary. | Removed `/tmp/ni-v0.5.1-post-release-verify/install-bin/ni`. | Yes | No user binary removed. |
| cleanup | `test -e temp/bin/ni`; fresh `command -v ni` | Binary absent and temp command not found. | Both checks exited non-zero as expected. | Yes | Temp verification directory remains documented below. |

## Public install parity closure

| Surface | v0.5.0 behavior | v0.5.1 behavior | Closed? | Notes |
| --- | --- | --- | --- | --- |
| `ni init .` | Failed with `unknown init option: .`. | Hosted and installed v0.5.1 pass `ni init . --yes`. | Yes for tested macOS arm64 path | Interactive TUI was not exercised in this non-TTY verification. |
| README first-user path | v0.5.0 did not match the current README first-project command. | v0.5.1 supports `ni init .` and status proof after init. | Yes for tested macOS arm64 path | Blank intent correctly remains `BLOCKED`. |
| `install.sh` v0.5.1 retrieval | Could not exist before publication. | Retrieved hosted v0.5.1, verified checksum, installed temp command. | Yes | Isolated HOME/BINDIR. |
| current-platform artifact | v0.5.0 lacked current `ni init .` support. | Hosted darwin/arm64 artifact reports `0.5.1` and passes init/status smoke. | Yes | Other platforms checksum-verified only. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks v0.5.1 verification? |
| --- | --- | --- | --- |
| Windows real-host execution | Current host is macOS. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | No tap/formula/install proof. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | No separate user or machine transcript in this task. | External install/init/status transcript. | No |
| model workspace host behavior | Host-level/global install and provider behavior remain unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Trusted CLI proof for a target workspace. | No |
| GoReleaser local full matrix dry-run | Release workflow used GoReleaser successfully; local `goreleaser` binary is unavailable. | Local environment with GoReleaser running check/dry-run matrix if needed. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Claim only after release evidence. | Published at `2026-06-08T00:50:40Z`. | Yes | Release URL verified. |
| hosted artifacts | Claim only after asset inventory and checksums. | Six assets visible; all checksums passed. | Yes | Five archives plus checksum file. |
| `install.sh` v0.5.1 retrieval | Claim only after isolated install proof. | Isolated install passed. | Yes | Temp HOME/BINDIR only. |
| Homebrew | Keep Planned / v0.5 candidate. | Preserved. | Yes | No Homebrew Available claim. |
| Windows real-host execution | Keep deferred without transcript. | Preserved. | Yes | Windows asset checksum only. |
| `ni init .` | May claim for tested v0.5.1 macOS arm64 path. | Hosted and installed paths passed. | Yes | Cross-platform execution not claimed. |
| `ni run` | Bounded prompt compilation only. | Preserved in release body and docs. | Yes | No generated prompt execution. |
| READY | CLI readiness only, not product readiness. | Blank temp workspaces were correctly `BLOCKED`. | Yes | No product readiness claim. |
| runtime execution boundary | No task runner, SPEC runner, execution harness, shell/Codex adapter, queue, PR automation, release automation system, or downstream execution layer. | Preserved. | Yes | Release process is distribution infrastructure, not kernel behavior. |

## Git status / inclusion check

| Path or group | git status --short | Expected after release? | Notes |
| --- | --- | --- | --- |
| README.md | Modified by post-release docs update | Yes | Version and parity note updated to v0.5.1. |
| README.ko.md | Modified by post-release docs update | Yes | Korean companion updated. |
| docs/126* | Modified by post-release docs update | Yes | Narrow addressed-by-v0.5.1 note. |
| docs/130* | Modified by post-release docs update | Yes | Narrow post-release pointer. |
| docs/131* | Modified by post-release docs update | Yes | Narrow post-release pointer. |
| docs/132* | Added | Yes | Post-release verification docs. |
| temporary artifact directory | Outside repo at `/tmp/ni-v0.5.1-post-release-verify` | No commit | Kept outside repo for audit trail. |
| `.ni/contract.json` | No diff | No direct edit | Protected. |
| `.ni/session.json` | No diff | No direct edit | Protected. |
| `.ni/plan.lock.json` | No diff | No direct edit | Protected. |
| unexpected files | None expected | No | Rechecked in final validation. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Clean before release execution; post-release docs changed after verification. |
| `git branch --show-current` | `main`. |
| `git remote -v` | `origin` points to `https://github.com/Nam-Cheol/ni.git`. |
| `git log --oneline --decorate -20` | HEAD at release was `945bd46 Add v0.5.1 publication checklist roadmap links`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` before release | Empty. |
| `git ls-remote --tags origin v0.5.1` before release | Empty. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Included README/TUI/install changes and docs/126 through docs/131. |
| `git diff --stat v0.5.0..HEAD` | 72 files changed before release execution. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Empty before release and after verification. |
| `gofmt -w .` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; root reported `NI Intent Readiness: READY` with no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed before release execution. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| current-tree first-user smoke | Passed with temp source binary; repeated init preserved files; synthetic lockfile was not modified. |
| `git push origin v0.5.1` | Passed; tag pushed. |
| `gh run watch 27109991213 --repo Nam-Cheol/ni --exit-status` | Passed; workflow completed successfully in 1m19s. |
| `gh release view v0.5.1 --repo Nam-Cheol/ni` | Passed; release metadata and assets verified. |
| `shasum -a 256 -c ni_0.5.1_checksums.txt` | Passed for all downloaded hosted archives. |
| hosted artifact smoke | Passed for darwin/arm64 help, version, init, and status. |
| isolated `install.sh --version 0.5.1` | Passed; checksum verified, temp command installed, version `0.5.1`, init/status passed, uninstall removed temp binary. |

## Changes made

| File | Why |
| --- | --- |
| `README.md` | Updated verified install version and first-project parity note to v0.5.1. |
| `README.ko.md` | Korean companion update. |
| `docs/22_INSTALL.md` | Updated release binary and curl installer evidence to v0.5.1. |
| `docs/install-curl.md` | Updated installer verification wording to v0.5.1. |
| `docs/install-curl.ko.md` | Korean companion update. |
| `docs/51_POST_RELEASE_ROADMAP.md` | Added docs/132 post-release verification pointer. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Korean companion update. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md` | Added narrow note that v0.5.1 addressed the v0.5.0 mismatch for tested path. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.ko.md` | Korean companion update. |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md` | Added post-release pointer. |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.ko.md` | Korean companion update. |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.md` | Added post-release pointer. |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.ko.md` | Korean companion update. |
| `docs/132_V0_5_1_POST_RELEASE_VERIFICATION.md` | Added this post-release verification. |
| `docs/132_V0_5_1_POST_RELEASE_VERIFICATION.ko.md` | Korean companion. |
| `scripts/check-install-docs.py` | Updated install-doc markers to v0.5.1. |
| `scripts/release-check.sh` | Updated current release marker and v0.5.1 post-release guard. |

## What this verification proves

- v0.5.1 was published.
- Hosted assets exist.
- Hosted checksums verify.
- Current-platform hosted artifact reports `0.5.1`.
- `install.sh` retrieves and installs v0.5.1 in an isolated temp path.
- The public install parity mismatch is closed for the tested macOS arm64 path.

## What this verification does not prove

- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.
- Every platform executes correctly.

## Recommended next task

Selected next task: A. external user validation plan

Selection rationale: release publication and tested macOS arm64 install parity
verified. The highest-value remaining proof gap is an external user or separate
machine transcript. Windows real-host execution and Homebrew remain separate
future lanes.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: create an external user validation plan for ni v0.5.1.

Read AGENTS.md, README.md, README.ko.md, docs/132_V0_5_1_POST_RELEASE_VERIFICATION.md, docs/22_INSTALL.md, docs/install-curl.md, docs/install-curl.ko.md, docs/120_GLOBAL_INSTALL_ACCEPTANCE.md, docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md, docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md, and docs/131_V0_5_1_PUBLICATION_CHECKLIST.md.

Do not publish, tag, create a release, upload assets, run ni end on the project root, relock the project root, edit .ni/contract.json, edit .ni/session.json, edit .ni/plan.lock.json, execute generated prompts, mark Homebrew Available, claim Windows real-host execution without a Windows transcript, or add runtime execution behavior.

Create docs/133_V0_5_1_EXTERNAL_USER_VALIDATION_PLAN.md and a Korean companion if companion docs are maintained. The plan should define:
- validation goal and decision options;
- target tester profiles and platforms;
- exact install transcript requirements;
- required commands: inspect installer, install v0.5.1, fresh shell command-name ni --help, ni version, ni init ., ni status --proof --next-questions, uninstall;
- evidence template with timestamp, OS/arch, shell, command, output, and cleanup proof;
- claim boundaries for Homebrew, Windows, model workspace packs, no-terminal, ni run, READY, and runtime execution;
- failure triage and rollback-or-doc-correction criteria;
- complete next executable prompt for collecting one external transcript.

Run validation:
- git status --short
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final report must include changed files, validation results, protected .ni diff, and confirmation that Homebrew remains Planned / v0.5 candidate, Windows real-host execution remains deferred unless actually tested, and ni run remains bounded prompt compilation only.
```
