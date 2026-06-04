# v0.5.1 Artifact Dry-Run

## Current status

State:
- v0.5.0 publication: verified.
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES.
- v0.5.1 RC validation decision: V0_5_1_RC_VALIDATION_PASS_WITH_NOTES.
- v0.5.1 release: not published.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: deferred on the macOS-only development host.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Dry-run goal

This dry-run checks whether the current release-candidate tree can produce a
release-like v0.5.1 current-platform artifact whose binary reports `0.5.1`,
whose checksum can be generated and verified locally, and whose first-user
artifact behavior still respects the kernel boundary. It does not publish, tag,
upload, create a GitHub Release, run a release workflow, create or publish a
Homebrew formula, run `ni end` on the project root, relock the project root,
execute generated prompts, or make `ni run` execute downstream work.

## Decision

V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES

Justification: a release-like current-platform `darwin/arm64` artifact was built
outside the repository with `ni/internal/version.Version=0.5.1`, the binary
reported `0.5.1`, help/init/status smoke checks passed, and a SHA-256 checksum
was generated and verified locally. Notes remain because local GoReleaser is not
installed, the full release matrix was not produced by GoReleaser, v0.5.1 is not
published, hosted v0.5.1 assets do not exist, `install.sh` public retrieval
cannot succeed until publication, Windows real-host execution is deferred,
Homebrew remains Planned / v0.5 candidate, and external user validation has not
been performed.

## Repository state

| Surface | Observed state | Notes |
| --- | --- | --- |
| git status | Clean at task start; this task adds docs/129 as uncommitted documentation. | No staging or commit performed. |
| v0.5.0 tag | Present; `git rev-parse v0.5.0` returned `b8fec7fa9615a861acf4eba59733c800c70c6cca`. | Verified baseline. |
| v0.5.1 tag | Absent. | No tag was created. |
| diff `v0.5.0..HEAD` | 66 files changed, 9184 insertions, 474 deletions before docs/129. | Current tree contains the v0.5.1 patch/RC work. |
| protected `.ni` diff | Empty. | `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` unchanged. |
| generated artifacts | Release-like artifacts are under `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK`. | Outside the repository. |

docs/126, docs/127, and docs/128 are tracked in the current tree.

## Release tooling audit

| Surface | Observed state | Pass? | Notes |
| --- | --- | --- | --- |
| version command | `cmd/ni/main.go` prints `version.Version`. | Yes | Source default remains `0.0.0-dev`. |
| version injection | `.goreleaser.yaml` uses `-X ni/internal/version.Version={{ .Version }}`. | Yes | Local fallback used `-X ni/internal/version.Version=0.5.1`. |
| release-check script | `scripts/release-check.sh` is present. | Yes | Check-only release surface gate; not a publishing command. |
| install-check script | `scripts/install-check.sh` is present. | Yes | Source/local install validation surface. |
| GoReleaser config | `.goreleaser.yaml` is present. | Yes | Defines `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, and `windows/amd64`; ignores `windows/arm64`. |
| GitHub workflow | `.github/workflows/release.yml` is present. | Yes | Tag-triggered release path; not run in this dry-run. |
| archive naming | `ni_{{ .Version }}_{{ .Os }}_{{ .Arch }}` with `tar.gz`, Windows `zip`. | Yes | Current-platform local archive used `ni_0.5.1_darwin_arm64.tar.gz`. |
| checksum generation | GoReleaser config names `ni_{{ .Version }}_checksums.txt`. | Yes | Local fallback generated `ni_0.5.1_checksums.txt` with `shasum -a 256`. |
| GoReleaser availability | `goreleaser --version` failed with command not found. | No | Full GoReleaser dry-run remains deferred; local Go fallback was used. |

Expected v0.5.1 release artifact names from the release config:
- `ni_0.5.1_linux_amd64.tar.gz`
- `ni_0.5.1_linux_arm64.tar.gz`
- `ni_0.5.1_darwin_amd64.tar.gz`
- `ni_0.5.1_darwin_arm64.tar.gz`
- `ni_0.5.1_windows_amd64.zip`
- `ni_0.5.1_checksums.txt`

## Artifact inventory

| Artifact | Path | Platform | Version output | Checksum generated? | Checksum verified? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| raw binary | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni` | darwin/arm64 | `0.5.1` | n/a | n/a | Built with local Go fallback and version ldflags. |
| archive | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_darwin_arm64.tar.gz` | darwin/arm64 | `0.5.1` from archived binary source | Yes | Yes | Contains `ni`. |
| checksums | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_checksums.txt` | n/a | n/a | Yes | Yes | Local checksum file only; not hosted. |

## Current-platform artifact verification

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| `ni --help` | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni --help` | Help renders and includes `ni init [.]`. | Passed. | Yes | Absolute artifact path. |
| `ni version` | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni version` | `0.5.1`. | `0.5.1`. | Yes | Version injection verified. |
| command-name `ni --help` | `PATH=/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK:$PATH zsh -f -c 'ni --help'` | PATH-resolved help works. | Passed. | Yes | `command -v ni` resolved to the temp artifact. |
| command-name `ni version` | `PATH=/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK:$PATH zsh -f -c 'ni version'` | `0.5.1`. | `0.5.1`. | Yes | Fresh shell command-name check. |
| `ni init . --yes` | Artifact binary in `/private/tmp/ni-v0.5.1-artifact-project.xCKnV8`. | Creates planning workspace. | Created `.ni/contract.json`, `.ni/session.json`, `.ni/project.json`, readiness files, harness/pressure seed files, and `docs/plan/**`. | Yes | No lockfile or prompt created. |
| command-name `ni init . --yes` | PATH-resolved artifact in `/private/tmp/ni-v0.5.1-artifact-path-project.a47CFC`. | Creates planning workspace. | Passed. | Yes | Command-name init smoke was practical and run. |
| `ni status --proof --next-questions` | Artifact binary in blank temp project. | Runs status proof and blocks incomplete intent. | `NI Intent Readiness: BLOCKED` with R014, OQ-001, R015, and R016. | Yes | Expected first-run scaffold behavior. |

## Generated artifact check

| Artifact | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `.ni/contract.json` | Exists after init. | Exists. | Yes | Temp project only. |
| `.ni/session.json` | Exists after init. | Exists. | Yes | Temp project only. |
| `docs/plan/**` | Exists after init. | Twelve planning docs created. | Yes | Temp project only. |
| `.ni/plan.lock.json` | Must not be created by init. | Not present. | Yes | No lock action. |
| downstream generated prompt | Must not be executed or created by init/status. | No `*.prompt*` file under `.ni`. | Yes | `ni run` was not executed. |

## Checksum verification

| File | SHA256 | Verified? | Command | Notes |
| --- | --- | --- | --- | --- |
| `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_darwin_arm64.tar.gz` | `593adc8f8643e6b89f3400eef2ee11e5eabd3e472847d2efd9897dc6117977f1` | Yes | `shasum -a 256 -c /private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_checksums.txt` | Local checksum verification returned `OK`. |

Full release archive checksum verification remains deferred to a GoReleaser or
publication task because this dry-run generated only the current-platform local
archive.

## install.sh future v0.5.1 gate

| Check | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `install.sh --dry-run --version 0.5.1` | Constructs v0.5.1 asset/checksum names without installing. | Passed. | Yes | Dry-run only. |
| expected URL | Uses `https://github.com/Nam-Cheol/ni/releases/download/v0.5.1/...`. | Printed `ni_0.5.1_darwin_arm64.tar.gz` and `ni_0.5.1_checksums.txt` URLs. | Yes | URLs are future-publication gates. |
| public URL availability | Must not be claimed before release publication. | Not checked as available and not claimed. | Yes | v0.5.1 release is not published. |
| post-publication install gate | Must later install hosted asset and verify command-name help/version. | Deferred. | Yes | Required after publication. |

## Release notes delta

The v0.5.1 release notes delta remains the patch scope recorded in docs/127 and
docs/128:
- Fix: public v0.5.0 install parity mismatch where installed `ni --help` and
  `ni version` passed but `ni init .` failed.
- Add: `ni init .` positional target support.
- Add: Bubble Tea v2 / Lip Gloss v2 guided init TUI.
- Preserve: non-interactive fallback for CI and non-TTY contexts.
- Preserve: existing-file and `.ni/plan.lock.json` protection during init.
- Improve: README macOS / Windows two-path onboarding and command-name
  verification guidance.
- Exclude: downstream execution, Homebrew Available, Windows real-host proof,
  and deterministic no-terminal claims.

This delta is still draft-only until actual v0.5.1 artifacts and the installer
path are verified after publication.

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks artifact dry-run? |
| --- | --- | --- | --- |
| v0.5.1 publication | This task is non-publishing. | Tag, release workflow, hosted assets, and post-publication checks. | No |
| hosted assets | No GitHub Release was created. | Release page asset inventory. | No |
| `install.sh` actual v0.5.1 retrieval | Public assets do not exist yet. | Isolated install from hosted v0.5.1 asset plus `ni --help` and `ni version`. | No |
| Windows real-host execution | Current development host is macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | No published tap/formula/install proof. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | No separate user or machine transcript. | External install/init/status transcript. | No |
| model workspace host behavior | Host-level/global install and provider behavior remain unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Trusted CLI proof for a target workspace. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Preserved. | Yes | No release action occurred. |
| artifact dry-run | May claim local current-platform release-like artifact only. | Preserved. | Yes | Full GoReleaser matrix not claimed. |
| `install.sh` v0.5.1 retrieval | Must not claim public retrieval before publication. | Preserved. | Yes | Dry-run URL construction only. |
| current-tree behavior | May claim local artifact behavior. | Help/version/init/status passed. | Yes | Not a hosted artifact claim. |
| Homebrew | Must remain Planned / v0.5 candidate. | Preserved. | Yes | No Homebrew Available claim. |
| Windows real-host execution | Must remain deferred without transcript. | Preserved. | Yes | macOS host only. |
| `ni init .` | Guided setup only; no execution. | Creates planning workspace. | Yes | No downstream prompt. |
| `ni run` | Bounded prompt compilation only. | Not run in this dry-run. | Yes | No execution behavior added. |
| READY | CLI authority only. | No model-only readiness claim. | Yes | Temp scaffold correctly reported BLOCKED. |
| runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | No runtime behavior added. |

## Git status / inclusion check

| Path or group | `git status --short` | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Onboarding surface. |
| README.ko.md | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Korean companion. |
| cmd/ni/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | CLI patch surface. |
| internal/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Docstore/TUI support. |
| go.mod | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Bubble Tea / Lip Gloss dependencies. |
| go.sum | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Dependency checksums. |
| install.sh | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Installer path handling. |
| install.ps1 | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Windows User PATH installer. |
| docs/126* | tracked | Yes | v0.5.0 parity evidence. |
| docs/127* | tracked | Yes | v0.5.1 patch plan. |
| docs/128* | tracked | Yes | v0.5.1 RC validation. |
| docs/129* | added by this task | Yes | Artifact dry-run evidence. |
| temporary artifact directory | outside repo | No | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK`. |
| `.ni/contract.json` | no diff | No direct edit | Protected. |
| `.ni/session.json` | no diff | No direct edit | Protected. |
| `.ni/plan.lock.json` | no diff | No direct edit | Protected. |
| unexpected files | none observed in repo status before docs/129 | No | Dry-run files stayed outside repo. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; clean at task start. |
| `git log --oneline --decorate -20` | Passed; current HEAD `842e565`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; v0.5.1 tag absent. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Passed; docs/126, docs/127, and docs/128 are tracked. |
| `git diff --stat v0.5.0..HEAD` | Passed; 66 files changed before docs/129. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |
| required ripgrep scans | Reviewed release/version/install/Homebrew/runtime boundary surfaces. |
| `goreleaser --version` | Failed with `command not found`; local Go fallback used. |
| `env GOCACHE=/private/tmp/ni-go-cache CGO_ENABLED=0 go build -ldflags='-s -w -X ni/internal/version.Version=0.5.1' -o /private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni ./cmd/ni` | Passed. |
| `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni --help` | Passed. |
| `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni version` | Passed; `0.5.1`. |
| command-name `ni --help` | Passed with temp artifact first in PATH. |
| command-name `ni version` | Passed; `0.5.1`. |
| artifact `ni init . --yes` | Passed in temp project. |
| artifact `ni status --proof --next-questions` | Passed; correctly reported `BLOCKED` for blank intent. |
| command-name `ni init . --yes` | Passed in a second temp project. |
| `tar -czf .../ni_0.5.1_darwin_arm64.tar.gz -C ... ni` | Passed. |
| `shasum -a 256 .../ni_0.5.1_darwin_arm64.tar.gz > .../ni_0.5.1_checksums.txt` | Passed. |
| `shasum -a 256 -c .../ni_0.5.1_checksums.txt` | Passed; returned `OK`. |
| `sh install.sh --dry-run --version 0.5.1` | Passed; printed future v0.5.1 asset and checksum URLs. |

## Changes made

| File | Why |
| --- | --- |
| `docs/129_V0_5_1_ARTIFACT_DRY_RUN.md` | Added English v0.5.1 artifact dry-run evidence, decision, deferrals, and next task. |
| `docs/129_V0_5_1_ARTIFACT_DRY_RUN.ko.md` | Added Korean companion with matching claim boundaries. |

## What this dry-run proves

- Local release-like v0.5.1 current-platform artifact behavior passed.
- Artifact `ni version` reports `0.5.1`.
- A current-platform archive checksum was generated and verified locally.
- The current-platform artifact supports the first-user init/status path.
- Publication gates remain explicit.

## What this dry-run does not prove

- v0.5.1 has been published.
- Hosted artifacts exist.
- `install.sh` retrieves v0.5.1 from a public release.
- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.
- GoReleaser locally produced the full release matrix.

## Recommended next task

Selected next task: A. v0.5.1 release notes finalization.

Selection rationale: the artifact dry-run passed with notes, so the next
pre-publication step is to finalize release notes while keeping publication,
hosted installer retrieval, Windows real-host execution, Homebrew, and external
validation as explicit gates.
