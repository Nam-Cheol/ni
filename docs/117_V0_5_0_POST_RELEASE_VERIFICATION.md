# v0.5.0 Post-Release Verification

## Current status

- Human approval choice: APPROVE_PUBLICATION_PREP_ONLY
- v0.5.0 publication: performed and verified in this document
- Release binary: verified in this document
- Curl installer: verified in this document
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Verification goal

This verifies the actual v0.5.0 release state without creating a new release,
tag, asset upload, workflow run, Homebrew formula, project-root relock, or
runtime execution behavior.

## Decision

Decision: V0_5_0_POST_RELEASE_VERIFIED_WITH_NOTES.

Justification: the Git tag, GitHub release, hosted assets, checksum file,
darwin/arm64 release binary, and isolated curl installer path all verified for
v0.5.0. Notes remain because Homebrew is still Planned / v0.5 candidate,
Windows execution was not run on a Windows host, cross-platform execution beyond
the current platform remains unverified, model workspace host behavior remains
Experimental, no-terminal deterministic validation is not claimed, and external
user validation remains deferred.

## Release identity

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Git tag | `v0.5.0` exists locally | `git tag --list v0.5.0` returned `v0.5.0` | Yes | Local tag present. |
| GitHub release URL | `https://github.com/Nam-Cheol/ni/releases/tag/v0.5.0` | `gh release view` returned the same URL | Yes | Network lookup used `gh`. |
| Release name | `v0.5.0` | `v0.5.0` | Yes | From GitHub release metadata. |
| Latest status | Latest release should be v0.5.0 | `gh release view --repo Nam-Cheol/ni` returned `tagName: v0.5.0` | Yes | The local `gh` JSON surface does not expose `isLatest`; this is inferred from the default latest release lookup. |
| Draft/prerelease status | Not draft, not prerelease | `isDraft: false`, `isPrerelease: false` | Yes | From GitHub release metadata. |
| Published time | Published release timestamp present | `2026-06-02T08:13:27Z` | Yes | From GitHub release metadata. |
| Commit | Tag resolves locally | `b8fec7fa9615a861acf4eba59733c800c70c6cca` | Yes | `gh` reports `targetCommitish: main`; local tag rev verified separately. |
| Asset count | Expected archives plus checksum | 6 assets | Yes | 5 platform archives plus `ni_0.5.0_checksums.txt`. |

## Asset inventory

| Asset | Purpose/platform | Downloaded? | Checksum verified? | Execution verified? | Notes |
| --- | --- | --- | --- | --- | --- |
| `ni_0.5.0_checksums.txt` | Checksum file | Yes | n/a | n/a | Downloaded to `/tmp/ni-v0.5.0-post-release-verify`. |
| `ni_0.5.0_darwin_amd64.tar.gz` | macOS amd64 | Yes | Yes | No | Asset checksum verified; not current platform. |
| `ni_0.5.0_darwin_arm64.tar.gz` | macOS arm64 | Yes | Yes | Yes | Current platform artifact. |
| `ni_0.5.0_linux_amd64.tar.gz` | Linux amd64 | Yes | Yes | No | Asset checksum verified; Linux execution not run locally. |
| `ni_0.5.0_linux_arm64.tar.gz` | Linux arm64 | Yes | Yes | No | Asset checksum verified; Linux execution not run locally. |
| `ni_0.5.0_windows_amd64.zip` | Windows amd64 | Yes | Yes | No | Asset checksum verified; Windows execution requires a Windows host. |

## Checksum verification

| Asset | Expected checksum | Actual checksum | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni_0.5.0_darwin_amd64.tar.gz` | `e9ba596ee3d7fde6fd3fbb3af2f20a4181c97be62582c20ce85b38cb18e3b64e` | matched | Yes | `shasum -a 256 -c ni_0.5.0_checksums.txt` |
| `ni_0.5.0_darwin_arm64.tar.gz` | `aad8c61b34429adda0065823a8434ed386dbbd328691670697a9b07114194fbb` | matched | Yes | Current platform artifact. |
| `ni_0.5.0_linux_amd64.tar.gz` | `3631b3cfa203cefc662f5c1d639fbad76197b052f0d2f7d723efe5f267d7170f` | matched | Yes | Checksum only. |
| `ni_0.5.0_linux_arm64.tar.gz` | `3a5cc8275f28d1fb0a37a57cfef67298f69cdc6c71781bb6b57fb6c6556ecb87` | matched | Yes | Checksum only. |
| `ni_0.5.0_windows_amd64.zip` | `f06d0beb03b761cb2612b167a2b34126e582486cdb46474400c9226229442144` | matched | Yes | Checksum only; Windows execution not run. |

## Binary verification

| Binary source | Command | Observed output | Pass? | Notes |
| --- | --- | --- | --- | --- |
| extracted current-platform artifact | `ni --help` | Starts with `ni is a project intent compiler.` and lists public commands | Yes | Extracted from `ni_0.5.0_darwin_arm64.tar.gz`. |
| extracted current-platform artifact | `ni version` | `0.5.0` | Yes | Release linker version is correct. |
| installed via curl installer | `ni --help` | Starts with `ni is a project intent compiler.` and lists public commands | Yes | Installed into isolated `/tmp` `BINDIR`. |
| installed via curl installer | `ni version` | `0.5.0` | Yes | Release linker version is correct. |

## Installer verification

| Installer path | Version | Destination | Verification | Uninstall check | Pass? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| macOS/Linux curl installer | `0.5.0` | `/tmp/ni-v0.5.0-post-release-verify/installed/bin/ni` | Dry-run, checksum verification, `ni --help`, and `ni version` passed | Remove the installed `ni` file from the chosen `BINDIR`; temporary directory is cleaned outside the repo | Yes | The installer printed `Verified checksum for ni_0.5.0_darwin_arm64.tar.gz`. |
| Windows manual zip path | `0.5.0` | User-chosen directory containing `ni.exe` | Asset exists and checksum verified | Remove the copied `ni.exe` from the chosen directory | Notes | Execution on real Windows was not run locally. |
| Homebrew | Planned / v0.5 candidate | n/a | Not run | n/a | Notes | No tap, formula, `brew install`, `ni --help`, or `ni version` Homebrew evidence. |

## README install-doc sync

| Section | Expected post-release wording | Change made | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS install | Use verified v0.5.0 curl installer after inspection | README and install docs now use `VERSION="0.5.0"` | Yes | Isolated installer check passed. |
| macOS uninstall | Remove the installed `ni` from the chosen `BINDIR` | Preserved | Yes | No package-manager uninstall claim. |
| Windows install | Use v0.5.0 `windows/amd64` zip with checksum comparison | README updated to v0.5.0 | Yes | Execution remains bounded to manual Windows verification. |
| Windows uninstall | Remove the copied `ni.exe` | Preserved | Yes | No MSI, winget, Chocolatey, Scoop, or Homebrew claim. |
| verification command | Run `ni --help` and `ni version` | Preserved and updated to v0.5.0 | Yes | Current-platform output verified. |
| Homebrew status | Homebrew remains Planned / v0.5 candidate | Preserved | Yes | No `brew install` instruction. |
| ni run boundary | `ni run` compiles only | Preserved | Yes | No downstream execution claim. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks post-release verification? |
| --- | --- | --- | --- |
| Homebrew availability | No tap/formula/install proof | Tap, formula, sha256, `brew install`, `ni --help`, and `ni version` | No |
| Windows execution on real Windows | This host is macOS darwin/arm64 | Windows transcript showing checksum, extraction, `ni.exe --help`, and `ni.exe version` | No |
| cross-platform install execution | Only darwin/arm64 execution was run | Per-platform execution transcripts | No |
| model workspace host verification | Broad product path remains Experimental | Host-specific install/discovery proof and provider behavior transcript | No |
| no-terminal deterministic validation not claimed | No trusted runner transcript for no-terminal proof | Exact trusted CLI output for target workspace | No |
| external user validation | No external user transcript was collected | Maintained external validation notes or user-run transcript | No |
| additional benchmark breadth if relevant | Post-release verification is not benchmark expansion | Broader benchmark design with `not_measured` boundaries | No |

## Blockers

None.

| Blocker | Evidence | Required fix |
| --- | --- | --- |
| None | Release, assets, checksums, current-platform binary, and curl installer passed. | n/a |

## Warnings

| Warning | Evidence | Mitigation |
| --- | --- | --- |
| Windows execution not run | Windows asset checksum passed, but this host cannot execute `ni.exe` as Windows | Keep README wording bounded to manual Windows verification. |
| Homebrew remains unavailable | No Homebrew verification evidence exists | Keep Homebrew Planned / v0.5 candidate. |
| `gh` JSON field mismatch | `isLatest` is not supported by this installed `gh`; latest status was checked by default `gh release view` | Document the inference instead of overclaiming unsupported metadata. |

## Risks

| Risk | Impact | Follow-up |
| --- | --- | --- |
| Platform-specific runtime issue outside darwin/arm64 | Users on Linux or Windows could hit unverified execution behavior | Run platform-specific verification where possible. |
| Installer docs could be overread as Homebrew availability | Package-manager users may expect `brew install` | Keep Homebrew deferral prominent. |
| Historical pre-publication docs can appear stale | docs/113 through docs/115 record earlier states | Link to this post-release verification from the roadmap and install docs. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.0 published status | Claim only after release evidence | Verified release URL, tag, published timestamp, assets | Yes | No new release action run. |
| Release binary | Available only after hosted asset and checksum verification | v0.5.0 assets and checksums verified | Yes | Current-platform execution passed. |
| Curl installer | Available only after isolated install verification | v0.5.0 isolated install passed | Yes | Installer does not install model skills or run downstream work. |
| Homebrew | Planned / v0.5 candidate | Preserved | Yes | No Homebrew Available claim. |
| Windows install | Asset/checksum verified, execution bounded unless run on Windows | README says Windows execution remains manual verification boundary | Yes | No Windows execution overclaim. |
| Model workspace packs | Experimental | Preserved | Yes | No global/host verification claim. |
| No-terminal | Experimental / assisted | Preserved | Yes | No deterministic validation claim. |
| ni run | Prompt compilation only | Preserved | Yes | No generated prompt executed. |
| Benchmark evidence | No implementation/adoption/cost/latency/downstream-quality proof | Preserved | Yes | `not_measured` boundary unchanged. |
| Runtime execution boundary | ni is not task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop | Preserved | Yes | No runtime behavior added. |

## Git status / inclusion check

| Path or group | git status --short | git ls-files / tracked check | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| README.md | modified | tracked | Yes | v0.5.0 install sync. |
| README.ko.md | modified | tracked | Yes | Korean companion sync. |
| docs/110_* | no new change expected | tracked | No | Historical RC audit remains unchanged. |
| docs/111_* | no new change expected | tracked | No | Historical release-note draft remains unchanged. |
| docs/112_* | no new change expected | tracked | No | Historical final preflight remains unchanged. |
| docs/113_* | no new change expected | tracked | No | Historical dry-run audit remains unchanged. |
| docs/114_* | no new change expected | tracked | No | Historical publication checklist remains unchanged. |
| docs/115_* | no new change expected | tracked | No | Historical human approval packet remains unchanged. |
| docs/116_* | no new change expected | tracked | No | README visual prompt pass remains unchanged. |
| docs/117_* | new | new until added | Yes | Post-release verification docs. |
| docs/51* | modified | tracked | Yes | Narrow pointer to docs/117. |
| downloaded temp artifacts | outside repo | not tracked | No | `/tmp/ni-v0.5.0-post-release-verify`; cleaned after verification. |
| generated artifacts | none expected | not tracked | No | No generated images added. |
| .ni/contract.json | no diff expected | tracked protected file | No | Must remain unchanged. |
| .ni/session.json | no diff expected | tracked protected file | No | Must remain unchanged. |
| .ni/plan.lock.json | no diff expected | tracked protected file | No | Must remain unchanged. |
| unexpected files | none expected | n/a | No | Final `git status --short` should be docs/scripts only. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean; final status reviewed after edits. |
| `git tag --list v0.5.0` | Pass | Returned `v0.5.0`. |
| `git rev-parse v0.5.0` | Pass | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git ls-files docs/110_* ... docs/116_*` | Pass | Required historical docs are tracked. |
| `gh release view v0.5.0 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets,targetCommitish` | Pass | Release metadata and 6 assets returned. |
| `gh release view --repo Nam-Cheol/ni --json tagName,url,publishedAt` | Pass | Default latest release lookup returned `v0.5.0`. |
| `gh release download v0.5.0 --repo Nam-Cheol/ni --dir /tmp/ni-v0.5.0-post-release-verify --pattern '*' --clobber` | Pass | Downloaded 6 release assets outside the repo. |
| `cd /tmp/ni-v0.5.0-post-release-verify && shasum -a 256 -c ni_0.5.0_checksums.txt` | Pass | All 5 platform archives verified. |
| extracted artifact `ni --help` | Pass | Help rendered. |
| extracted artifact `ni version` | Pass | `0.5.0`. |
| `env BINDIR=/tmp/ni-v0.5.0-post-release-verify/installed/bin sh install.sh --dry-run --version 0.5.0` | Pass | Correct darwin/arm64 URLs printed. |
| `env BINDIR=/tmp/ni-v0.5.0-post-release-verify/installed/bin sh install.sh --version 0.5.0` | Pass | Checksum verified and binary installed. |
| installed binary `ni --help` | Pass | Help rendered. |
| installed binary `ni version` | Pass | `0.5.0`. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | Final validation command. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass | Help rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass | Development source output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed after v0.5.0 marker update. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack boundary checks passed. |
| `bash scripts/demo-check.sh` | Pass | Demo checks passed; no generated prompt executed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Smoke passed; fixture `ni end` / relock is not project-root relock. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source/build/temp install checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Check-only release gate passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` diff. |

## Changes made

- `README.md`: updated install/uninstall wording to verified v0.5.0 facts.
- `README.ko.md`: mirrored README changes without adding broader claims.
- `docs/22_INSTALL.md`: updated release binary and curl installer details to
  verified v0.5.0 facts.
- `docs/install-curl.md`: updated curl installer verification wording to v0.5.0.
- `docs/install-curl.ko.md`: updated Korean curl installer companion wording.
- `docs/51_POST_RELEASE_ROADMAP.md`: added a narrow pointer to this document.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: added matching Korean pointer.
- `docs/117_V0_5_0_POST_RELEASE_VERIFICATION.md`: added this verification.
- `docs/117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md`: added Korean companion.
- `scripts/check-install-docs.py`: updated install-doc markers to v0.5.0.
- `scripts/release-check.sh`: updated README/install marker checks to the
  verified v0.5.0 post-release state.

## What this verification proves

- release existence
- release asset presence
- checksum verification for all hosted platform archives
- current-platform darwin/arm64 binary behavior
- curl installer behavior for v0.5.0 in an isolated temporary destination
- README install-doc sync to verified v0.5.0 facts

## What this verification does not prove

- Homebrew is Available
- untested platforms execute correctly
- model workspace host behavior is verified
- no-terminal is deterministic
- downstream execution succeeds
- benchmark effect size or causal impact
- external users succeed

## Recommended next task

Selected next task: C. Homebrew implementation audit.

Why: release verification passed with notes, and Homebrew is now the main
public install-path deferral. The next task should audit Homebrew requirements
without publishing a formula or claiming availability before tap, formula,
sha256, `brew install`, `ni --help`, and `ni version` evidence exists.

## Next task prompt

```text
Proceed with a Homebrew implementation audit in /Users/namba/Documents/project/ni.

Goal:
Audit what remains before Homebrew can move from Planned / v0.5 candidate to
Available after the verified v0.5.0 release. This is an audit and documentation
task only unless the user separately approves actual tap/formula publication.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/22_INSTALL.md
- docs/install-curl.md
- docs/install-curl.ko.md
- docs/51_POST_RELEASE_ROADMAP.md
- docs/51_POST_RELEASE_ROADMAP.ko.md
- docs/53_DISTRIBUTION_STRATEGY.md
- docs/53_DISTRIBUTION_STRATEGY.ko.md
- docs/54_HOMEBREW_DISTRIBUTION.md
- docs/54_HOMEBREW_DISTRIBUTION.ko.md
- docs/71_HOMEBREW_FORMULA_DRAFT.md
- docs/72_HOMEBREW_TAP_PLAN.md
- docs/72_HOMEBREW_TAP_PLAN.ko.md
- docs/80_HOMEBREW_DECISION.md
- docs/80_HOMEBREW_DECISION.ko.md
- docs/117_V0_5_0_POST_RELEASE_VERIFICATION.md
- docs/117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md
- .goreleaser.yaml
- .github/workflows/release.yml
- scripts/check-install-docs.py
- scripts/release-check.sh

Preserve:
- Homebrew: Planned / v0.5 candidate until tap, formula, sha256, brew install,
  ni --help, and ni version are all verified.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni run compiles a bounded handoff prompt only and does not execute downstream
  work.

Do not:
- publish a Homebrew formula
- create or push a tap repository
- run release workflows
- upload release assets
- create a new GitHub release or tag
- run ni end on the project root
- relock the project root
- execute generated prompts
- add runtime execution behavior
- claim Homebrew Available before full evidence exists

Required output:
- Add or update a narrow Homebrew audit document that lists current v0.5.0
  release facts, exact formula inputs, checksum requirements, tap requirements,
  local formula verification steps, published tap verification steps, abort
  criteria, and README wording gates.
- Keep README and README.ko unchanged unless a concrete overclaim is found.
- Keep .ni/contract.json, .ni/session.json, and .ni/plan.lock.json unchanged.

Validation:
- git status --short
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, Homebrew decision, blockers, deferrals, validation
results, and confirm no publication, no project-root relock, no generated
prompt execution, and no Homebrew Available claim.
```
