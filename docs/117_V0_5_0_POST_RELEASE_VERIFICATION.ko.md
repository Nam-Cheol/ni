# v0.5.0 Post-Release Verification

## Current status

- Human approval choice: APPROVE_PUBLICATION_PREP_ONLY
- v0.5.0 publication: 이 문서에서 performed and verified
- Release binary: 이 문서에서 verify
- Curl installer: 이 문서에서 verify
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Verification goal

이 문서는 새 release, tag, asset upload, workflow run, Homebrew formula,
project-root relock, runtime execution behavior를 만들지 않고 실제 v0.5.0
release state를 verify한다.

## Decision

Decision: V0_5_0_POST_RELEASE_VERIFIED_WITH_NOTES.

Justification: Git tag, GitHub release, hosted assets, checksum file,
darwin/arm64 release binary, isolated curl installer path가 v0.5.0에 대해
검증되었다. Notes는 Homebrew가 여전히 Planned / v0.5 candidate이고, Windows
execution은 Windows host에서 실행하지 않았으며, current platform 밖의
cross-platform execution, model workspace host behavior, no-terminal
deterministic validation, external user validation이 deferred이기 때문에 남는다.

## Release identity

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Git tag | `v0.5.0` exists locally | `git tag --list v0.5.0` returned `v0.5.0` | Yes | Local tag present. |
| GitHub release URL | `https://github.com/Nam-Cheol/ni/releases/tag/v0.5.0` | `gh release view` returned same URL | Yes | Network lookup used `gh`. |
| Release name | `v0.5.0` | `v0.5.0` | Yes | GitHub release metadata. |
| Latest status | Latest release should be v0.5.0 | `gh release view --repo Nam-Cheol/ni` returned `tagName: v0.5.0` | Yes | Installed `gh` JSON에는 `isLatest` field가 없어 default latest lookup에서 infer했다. |
| Draft/prerelease status | Not draft, not prerelease | `isDraft: false`, `isPrerelease: false` | Yes | GitHub release metadata. |
| Published time | Published release timestamp present | `2026-06-02T08:13:27Z` | Yes | GitHub release metadata. |
| Commit | Tag resolves locally | `b8fec7fa9615a861acf4eba59733c800c70c6cca` | Yes | `gh` reports `targetCommitish: main`; local tag rev 별도 확인. |
| Asset count | Expected archives plus checksum | 6 assets | Yes | 5 platform archives plus `ni_0.5.0_checksums.txt`. |

## Asset inventory

| Asset | Purpose/platform | Downloaded? | Checksum verified? | Execution verified? | Notes |
| --- | --- | --- | --- | --- | --- |
| `ni_0.5.0_checksums.txt` | Checksum file | Yes | n/a | n/a | `/tmp/ni-v0.5.0-post-release-verify`에 download. |
| `ni_0.5.0_darwin_amd64.tar.gz` | macOS amd64 | Yes | Yes | No | Checksum only. |
| `ni_0.5.0_darwin_arm64.tar.gz` | macOS arm64 | Yes | Yes | Yes | Current platform artifact. |
| `ni_0.5.0_linux_amd64.tar.gz` | Linux amd64 | Yes | Yes | No | Linux execution은 local에서 run하지 않음. |
| `ni_0.5.0_linux_arm64.tar.gz` | Linux arm64 | Yes | Yes | No | Linux execution은 local에서 run하지 않음. |
| `ni_0.5.0_windows_amd64.zip` | Windows amd64 | Yes | Yes | No | Windows execution은 Windows host가 필요. |

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
| extracted current-platform artifact | `ni --help` | `ni is a project intent compiler.`로 시작하고 public commands를 list | Yes | `ni_0.5.0_darwin_arm64.tar.gz`에서 extract. |
| extracted current-platform artifact | `ni version` | `0.5.0` | Yes | Release linker version correct. |
| installed via curl installer | `ni --help` | `ni is a project intent compiler.`로 시작하고 public commands를 list | Yes | Isolated `/tmp` `BINDIR`에 install. |
| installed via curl installer | `ni version` | `0.5.0` | Yes | Release linker version correct. |

## Installer verification

| Installer path | Version | Destination | Verification | Uninstall check | Pass? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| macOS/Linux curl installer | `0.5.0` | `/tmp/ni-v0.5.0-post-release-verify/installed/bin/ni` | Dry-run, checksum verification, `ni --help`, `ni version` passed | 선택한 `BINDIR`에서 installed `ni` file을 삭제; temp dir은 repo 밖에서 cleanup | Yes | Installer printed `Verified checksum for ni_0.5.0_darwin_arm64.tar.gz`. |
| Windows manual zip path | `0.5.0` | User-chosen directory containing `ni.exe` | Asset exists and checksum verified | 선택한 directory에서 copied `ni.exe` 삭제 | Notes | Real Windows execution은 local에서 run하지 않음. |
| Homebrew | Planned / v0.5 candidate | n/a | Not run | n/a | Notes | Tap, formula, `brew install`, `ni --help`, `ni version` evidence 없음. |

## README install-doc sync

| Section | Expected post-release wording | Change made | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS install | Verified v0.5.0 curl installer after inspection | README and install docs now use `VERSION="0.5.0"` | Yes | Isolated installer check passed. |
| macOS uninstall | 선택한 `BINDIR`의 installed `ni` 삭제 | Preserved | Yes | Package-manager uninstall claim 없음. |
| Windows install | v0.5.0 `windows/amd64` zip과 checksum comparison | README updated to v0.5.0 | Yes | Windows execution remains manual verification boundary. |
| Windows uninstall | Copied `ni.exe` 삭제 | Preserved | Yes | MSI, winget, Chocolatey, Scoop, Homebrew claim 없음. |
| verification command | `ni --help` and `ni version` | Preserved and updated to v0.5.0 | Yes | Current-platform output verified. |
| Homebrew status | Homebrew remains Planned / v0.5 candidate | Preserved | Yes | `brew install` instruction 없음. |
| ni run boundary | `ni run` compiles only | Preserved | Yes | Downstream execution claim 없음. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks post-release verification? |
| --- | --- | --- | --- |
| Homebrew availability | Tap/formula/install proof 없음 | Tap, formula, sha256, `brew install`, `ni --help`, `ni version` | No |
| Windows execution on real Windows | 이 host는 macOS darwin/arm64 | Windows transcript with checksum, extraction, `ni.exe --help`, `ni.exe version` | No |
| cross-platform install execution | darwin/arm64 execution만 run | Per-platform execution transcripts | No |
| model workspace host verification | Broad product path remains Experimental | Host-specific install/discovery proof and provider behavior transcript | No |
| no-terminal deterministic validation not claimed | Trusted runner transcript 없음 | Exact trusted CLI output for target workspace | No |
| external user validation | External user transcript 없음 | Maintained external validation notes or user-run transcript | No |
| additional benchmark breadth if relevant | Post-release verification은 benchmark expansion이 아님 | Broader benchmark design with `not_measured` boundaries | No |

## Blockers

None.

| Blocker | Evidence | Required fix |
| --- | --- | --- |
| None | Release, assets, checksums, current-platform binary, curl installer passed. | n/a |

## Warnings

| Warning | Evidence | Mitigation |
| --- | --- | --- |
| Windows execution not run | Windows asset checksum passed, but this host cannot execute `ni.exe` as Windows | README wording을 manual Windows verification으로 bound. |
| Homebrew remains unavailable | Homebrew verification evidence 없음 | Homebrew Planned / v0.5 candidate 유지. |
| `gh` JSON field mismatch | `isLatest` not supported; default `gh release view`로 latest 확인 | Unsupported metadata를 overclaim하지 않고 inference로 기록. |

## Risks

| Risk | Impact | Follow-up |
| --- | --- | --- |
| darwin/arm64 밖 platform-specific runtime issue | Linux/Windows users may hit unverified execution behavior | Platform-specific verification run. |
| Installer docs overread as Homebrew availability | Package-manager users may expect `brew install` | Homebrew deferral prominent 유지. |
| Historical pre-publication docs can appear stale | docs/113 through docs/115 record earlier states | Roadmap and install docs link to this post-release verification. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.0 published status | Release evidence 후에만 claim | Release URL, tag, published timestamp, assets verified | Yes | New release action run 없음. |
| Release binary | Hosted asset and checksum verification 후 Available | v0.5.0 assets and checksums verified | Yes | Current-platform execution passed. |
| Curl installer | Isolated install verification 후 Available | v0.5.0 isolated install passed | Yes | Installer does not install model skills or run downstream work. |
| Homebrew | Planned / v0.5 candidate | Preserved | Yes | Homebrew Available claim 없음. |
| Windows install | Asset/checksum verified, execution bounded unless run on Windows | README says Windows execution remains manual verification boundary | Yes | Windows execution overclaim 없음. |
| Model workspace packs | Experimental | Preserved | Yes | Global/host verification claim 없음. |
| No-terminal | Experimental / assisted | Preserved | Yes | Deterministic validation claim 없음. |
| ni run | Prompt compilation only | Preserved | Yes | Generated prompt executed 없음. |
| Benchmark evidence | No implementation/adoption/cost/latency/downstream-quality proof | Preserved | Yes | `not_measured` boundary unchanged. |
| Runtime execution boundary | ni is not task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop | Preserved | Yes | Runtime behavior added 없음. |

## Git status / inclusion check

| Path or group | git status --short | git ls-files / tracked check | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| README.md | modified | tracked | Yes | v0.5.0 install sync. |
| README.ko.md | modified | tracked | Yes | Korean companion sync. |
| docs/110_* | no new change expected | tracked | No | Historical RC audit unchanged. |
| docs/111_* | no new change expected | tracked | No | Historical draft unchanged. |
| docs/112_* | no new change expected | tracked | No | Historical preflight unchanged. |
| docs/113_* | no new change expected | tracked | No | Historical dry-run audit unchanged. |
| docs/114_* | no new change expected | tracked | No | Historical publication checklist unchanged. |
| docs/115_* | no new change expected | tracked | No | Historical human approval packet unchanged. |
| docs/116_* | no new change expected | tracked | No | README visual prompt pass unchanged. |
| docs/117_* | new | new until added | Yes | Post-release verification docs. |
| docs/51* | modified | tracked | Yes | Narrow pointer to docs/117. |
| downloaded temp artifacts | outside repo | not tracked | No | `/tmp/ni-v0.5.0-post-release-verify`; cleanup after verification. |
| generated artifacts | none expected | not tracked | No | Generated images added 없음. |
| .ni/contract.json | no diff expected | tracked protected file | No | Must remain unchanged. |
| .ni/session.json | no diff expected | tracked protected file | No | Must remain unchanged. |
| .ni/plan.lock.json | no diff expected | tracked protected file | No | Must remain unchanged. |
| unexpected files | none expected | n/a | No | Final status should be docs/scripts only. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean; final status reviewed after edits. |
| `git tag --list v0.5.0` | Pass | Returned `v0.5.0`. |
| `git rev-parse v0.5.0` | Pass | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git ls-files docs/110_* ... docs/116_*` | Pass | Required historical docs are tracked. |
| `gh release view v0.5.0 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets,targetCommitish` | Pass | Release metadata and 6 assets returned. |
| `gh release view --repo Nam-Cheol/ni --json tagName,url,publishedAt` | Pass | Default latest release lookup returned `v0.5.0`. |
| `gh release download v0.5.0 --repo Nam-Cheol/ni --dir /tmp/ni-v0.5.0-post-release-verify --pattern '*' --clobber` | Pass | Downloaded 6 release assets outside repo. |
| `cd /tmp/ni-v0.5.0-post-release-verify && shasum -a 256 -c ni_0.5.0_checksums.txt` | Pass | All 5 platform archives verified. |
| extracted artifact `ni --help` | Pass | Help rendered. |
| extracted artifact `ni version` | Pass | `0.5.0`. |
| `env BINDIR=/tmp/ni-v0.5.0-post-release-verify/installed/bin sh install.sh --dry-run --version 0.5.0` | Pass | Correct darwin/arm64 URLs printed. |
| `env BINDIR=/tmp/ni-v0.5.0-post-release-verify/installed/bin sh install.sh --version 0.5.0` | Pass | Checksum verified and binary installed. |
| installed binary `ni --help` | Pass | Help rendered. |
| installed binary `ni version` | Pass | `0.5.0`. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | Final validation command. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; no blockers, deferrals, warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass | Help rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass | Development source output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed after v0.5.0 marker update. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack boundary checks passed. |
| `bash scripts/demo-check.sh` | Pass | Demo checks passed; generated prompt executed 없음. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Smoke passed; fixture `ni end` / relock is not project-root relock. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source/build/temp install checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Check-only release gate passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | Protected project-root `.ni` diff 없음. |

## Changes made

- `README.md`: verified v0.5.0 facts로 install/uninstall wording update.
- `README.ko.md`: broader claim 없이 README changes mirror.
- `docs/22_INSTALL.md`: release binary and curl installer details를 verified
  v0.5.0 facts로 update.
- `docs/install-curl.md`: curl installer verification wording을 v0.5.0으로 update.
- `docs/install-curl.ko.md`: Korean curl installer companion wording update.
- `docs/51_POST_RELEASE_ROADMAP.md`: 이 문서로 narrow pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer 추가.
- `docs/117_V0_5_0_POST_RELEASE_VERIFICATION.md`: English verification 추가.
- `docs/117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md`: Korean companion 추가.
- `scripts/check-install-docs.py`: install-doc markers를 v0.5.0으로 update.
- `scripts/release-check.sh`: README/install marker checks를 verified v0.5.0
  post-release state로 update.

## What this verification proves

- release existence
- release asset presence
- all hosted platform archives checksum verification
- current-platform darwin/arm64 binary behavior
- v0.5.0 curl installer behavior in isolated temporary destination
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

Why: release verification passed with notes이고, 이제 Homebrew가 주요 public
install-path deferral이다. 다음 task는 tap, formula, sha256, `brew install`,
`ni --help`, `ni version` evidence 전에는 publish하거나 Available을 claim하지
않는 Homebrew audit이어야 한다.

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
