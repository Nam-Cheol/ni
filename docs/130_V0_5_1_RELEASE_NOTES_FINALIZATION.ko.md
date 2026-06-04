# v0.5.1 Release Notes Finalization

## Current status

State:
- v0.5.0 publication: verified.
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES.
- v0.5.1 RC validation decision: V0_5_1_RC_VALIDATION_PASS_WITH_NOTES.
- v0.5.1 artifact dry-run decision: V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES.
- v0.5.1 release: not published.
- v0.5.1 tag: 이 checkout의 finalization 시점에는 absent.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: macOS-only development host에서는 deferred.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Finalization goal

이 문서는 publication checklist 또는 release action 전에 conservative한 v0.5.1
release notes draft를 finalize한다. 이 문서는 v0.5.1을 publish하지 않고, tag를
create하지 않고, GitHub Release를 create하지 않고, assets를 upload하지 않고,
release workflows를 run하지 않고, GoReleaser publish를 run하지 않고, Homebrew
formula를 create/publish하지 않고, project root에서 `ni end`를 run하지 않고,
project root를 relock하지 않고, generated prompts를 execute하지 않고, runtime
execution behavior를 add하지 않는다.

## Decision

V0_5_1_RELEASE_NOTES_READY_WITH_NOTES

Justification: release notes는 patch를 정확하고 보수적으로 설명할 준비가
되었다. Notes는 남아 있다. v0.5.1은 published되지 않았고, hosted v0.5.1
assets는 없고, public `install.sh` retrieval은 publication 전에는 성공을
claim할 수 없고, Windows real-host execution은 deferred이고, Homebrew는
Planned / v0.5 candidate이며, external user validation은 수행되지 않았고,
local GoReleaser full matrix dry-run은 GoReleaser가 local에서 unavailable이라
deferred이다.

## Release notes draft

### Summary

v0.5.1은 public install parity를 위한 patch release다. Published v0.5.0은
`ni --help`와 `ni version`을 통과했지만 README-required first project step인
`ni init .`에서 실패했으므로, current-tree onboarding fixes를 release delta로
묶는다.

### Why this patch exists

현재 README는 existing project directory에서 `ni init .`으로 첫 프로젝트를
시작하는 flow를 안내한다. Published v0.5.0은 install되고 version `0.5.0`을
report하지만 `ni init .`은 다음 오류로 실패한다:

```text
unknown init option: .
```

v0.5.1은 release artifact behavior를 current README onboarding path와
align한다.

### Fixes

- Public v0.5.0 install parity mismatch for `ni init .`를 fix한다.
- `ni init .`을 current-directory guided setup entry point로 동작하게 한다.
- Init은 planning workspace creation에 bounded된다. Agents 또는 downstream
  work를 run하지 않는다.

### Added

- `ni init .` positional target support.
- Bubble Tea v2 / Lip Gloss v2 guided init TUI.
- CI, scripts, non-TTY contexts를 위한 non-interactive fallback.
- Init 중 existing file protection.
- Init 중 `.ni/plan.lock.json` protection.
- Created/unchanged files를 보여주는 post-TUI plain text summary.
- Windows User PATH installer code와 static safety checks.

### Changed

- README onboarding은 macOS와 Windows primary install paths 중심으로 align된다.
- Install docs와 checkers는 `ni --help`, `ni version`, `ni init .`을 통한
  command-name verification에 맞춰진다.
- Domain init logic은 TUI rendering에서 분리되어 file planning을 testable하고
  deterministic하게 유지한다.

### Validation

- Current-tree tests와 release checks는 documented validation commands에서 pass한다.
- `ni/internal/version.Version=0.5.1`로 build한 local release-like darwin/arm64
  artifact는 `0.5.1`을 report한다.
- Local current-platform archive checksum은 generated and verified 되었다.
- Local release-like artifact는 temporary workspaces에서 `ni --help`,
  `ni version`, `ni init . --yes`, `ni status --proof --next-questions` smoke를
  통과했다.
- Audit evidence는 docs/126, docs/127, docs/128, docs/129에 기록되어 있다.

### Known deferrals

- v0.5.1은 아직 published되지 않았다.
- Hosted v0.5.1 artifacts and checksums는 아직 존재하지 않는다.
- `install.sh` public retrieval of v0.5.1은 hosted assets가 존재하고
  post-publication install check가 pass하기 전에는 claim할 수 없다.
- Windows real-host execution은 Windows transcript가 생길 때까지 deferred다.
- Homebrew remains Planned / v0.5 candidate.
- External user validation은 수행되지 않았다.
- Model workspace packs remain Experimental.
- No-terminal method remains Experimental / assisted.
- Local GoReleaser full matrix dry-run은 GoReleaser가 local에서 unavailable이라
  deferred다.

### What this patch does not do

- 이 patch 자체는 v0.5.1을 publish하지 않는다.
- Task runner, SPEC runner, execution harness, shell adapter, Codex exec
  adapter, queue, PR automation, release automation, downstream execution layer를
  add하지 않는다.
- `ni run`이 downstream work를 execute하게 만들지 않는다.
- Homebrew Available을 mark하지 않는다.
- Windows real-host execution을 verify하지 않는다.
- External user success를 prove하지 않는다.
- No-terminal을 deterministic하게 만들지 않는다.
- Benchmark evidence가 implementation correctness 또는 downstream execution
  quality를 prove한다고 claim하지 않는다.

### Upgrade/install note

v0.5.1이 실제로 published되고 hosted assets가 verified된 뒤, users는 documented
installer path로 published release를 install하고 다음을 verify해야 한다:

```bash
ni --help
ni version
ni init .
```

Publication 전까지 README와 install guidance는 honest해야 한다. v0.5.0은
verified public release이고, v0.5.1은 prepared patch delta다.

### Maintainer verification checklist

- Intended release commit과 clean working tree를 confirm한다.
- Authorized release tag를 create하기 전에 `v0.5.1` tag absence를 confirm한다.
- Full validation gate를 run한다.
- Release build version injection이 `0.5.1`을 report하는지 verify한다.
- Publication 후 hosted archive inventory와 checksums를 verify한다.
- Current-platform command-name `ni --help`와 `ni version`을 verify한다.
- Hosted `install.sh --version 0.5.1` install을 temporary HOME/BINDIR에서 verify한다.
- Installed `ni init .`와 `ni status --proof --next-questions`를 verify한다.
- Windows real-host execution, Homebrew, external validation, model workspace,
  no-terminal, downstream execution deferrals는 separate evidence가 생길 때까지
  explicit하게 유지한다.

## Patch rationale

| Issue | Evidence | v0.5.1 response | Notes |
| --- | --- | --- | --- |
| v0.5.0 `ni init .` mismatch | docs/126은 published v0.5.0이 `ni --help`와 `ni version`을 pass했지만 `ni init .`에서 `unknown init option: .`로 fail했다고 기록한다. | `ni init .` positional target support를 release delta에 포함한다. | Main patch reason. |
| README onboarding parity | Current README는 first project path로 `ni init .`을 사용한다. | Release artifact를 README onboarding과 align한다. | v0.5.0을 complete로 보이게 하려고 onboarding path를 숨기지 않는다. |
| First-user TUI | docs/124와 docs/128은 Bubble Tea v2 / Lip Gloss v2 guided init behavior를 기록한다. | Guided init TUI를 v0.5.1에 포함한다. | TUI는 intent를 collect할 뿐 readiness authority가 아니다. |
| Non-interactive fallback | docs/128과 docs/129는 `ni init . --yes` smoke checks를 기록한다. | CI와 non-TTY contexts를 위한 fallback behavior를 preserve한다. | Automation에 TUI requirement 없음. |
| Install docs/checkers | README, install docs, scripts가 command-name verification에 align된다. | Docs/checkers를 release notes와 synchronized 상태로 유지한다. | Windows real-host execution은 deferred. |

## Release notes claim audit

| Claim area | Expected boundary | Observed wording | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Draft says v0.5.1 has not been published. | Yes | Release action 없음. |
| v0.5.1 artifacts | Hosted artifacts exist claim 금지. | Draft says hosted v0.5.1 assets do not exist yet. | Yes | Local artifact dry-run과 별개. |
| `install.sh` v0.5.1 retrieval | Publication 전 public retrieval claim 금지. | Draft says retrieval cannot be claimed until hosted assets exist. | Yes | Dry-run URL construction은 public retrieval이 아니다. |
| Homebrew | Must remain Homebrew: Planned / v0.5 candidate. | Draft preserves Planned / v0.5 candidate and must not claim Homebrew Available. | Yes | Tap/formula/install proof 없음. |
| Windows real-host execution | Transcript 없으면 deferred 유지. | Draft says Windows real-host execution remains deferred. | Yes | Static checks는 real-host proof가 아니다. |
| `ni init .` | v0.5.1 patch support는 claim 가능하지만 v0.5.0 parity는 claim 금지. | Draft ties support to the v0.5.1 patch delta. | Yes | v0.5.0 mismatch remains explicit. |
| `ni run` | Bounded prompt compilation only. | Draft says patch does not make `ni run` execute downstream work. | Yes | Downstream execution claim 없음. |
| READY | `ni status`에서만 나와야 하며 model judgment가 아니다. | Draft does not use READY as product readiness. | Yes | CLI remains authority. |
| Benchmark evidence | Implementation correctness 또는 downstream execution quality claim 금지. | Draft explicitly avoids this claim. | Yes | Benchmark overclaim 없음. |
| Runtime execution boundary | Task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, downstream execution layer 없음. | Draft lists these as excluded. | Yes | Kernel boundary preserved. |

## Validation evidence

| Evidence | Result | Notes |
| --- | --- | --- |
| docs/126 public install parity | PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED. | v0.5.0 help/version passed, `ni init .` failed. |
| docs/127 patch plan | V0_5_1_PATCH_PLAN_READY_WITH_NOTES. | Patch scope와 exclusions가 documented. |
| docs/128 RC validation | V0_5_1_RC_VALIDATION_PASS_WITH_NOTES. | Current-tree behavior와 docs/checkers passed. |
| docs/129 artifact dry-run | V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES. | Local current-platform artifact version/checksum/init smoke passed; GoReleaser matrix deferred. |
| current validation commands | 이 문서의 validation results를 참고한다. | Go tests, docs checks, smoke, install, release, quality gates 포함. |
| protected `.ni` diff | Empty. | `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` unchanged. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks release notes? |
| --- | --- | --- | --- |
| v0.5.1 publication | 이 task는 non-publishing. | Authorized tag, release workflow, GitHub Release, hosted asset proof. | No |
| Hosted artifacts | v0.5.1 release page/assets가 아직 없다. | Release page inventory and checksum verification. | No |
| `install.sh` actual v0.5.1 retrieval | Hosted assets는 publication 전에는 없다. | Isolated hosted install, checksum verification, `ni --help`, `ni version`, `ni init .`. | No |
| Windows real-host execution | Current development host is macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | Tap/formula/install proof 없음. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| External user validation | Separate user/machine transcript 없음. | External install/init/status transcript. | No |
| Model workspace host behavior | Host-level/global install과 provider behavior는 unverified. | Host-specific discovery/install/provider transcript. | No |
| No-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Trusted CLI proof for a target workspace. | No |
| GoReleaser full matrix dry-run | GoReleaser is unavailable locally. | GoReleaser installed environment에서 check/dry-run matrix. | No |

## Git status / inclusion check

| Path or group | `git status --short` | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | task start 기준 HEAD relative clean; `v0.5.0..HEAD`에서 changed. | Yes | Onboarding and public parity note. |
| README.ko.md | task start 기준 HEAD relative clean; `v0.5.0..HEAD`에서 changed. | Yes | Korean companion onboarding and parity note. |
| docs/126* | tracked. | Yes | Public install parity evidence. |
| docs/127* | tracked. | Yes | Patch release plan. |
| docs/128* | tracked. | Yes | Release-candidate validation. |
| docs/129* | tracked. | Yes | Artifact dry-run evidence. |
| docs/130* | added by this task. | Yes | Release notes finalization and Korean companion. |
| CHANGELOG.md | absent. | No | Release-history confusion을 피하려고 add하지 않음. |
| RELEASE.md | absent. | No | Publication imply를 피하려고 add하지 않음. |
| `.ni/contract.json` | no diff. | No direct edit | Protected. |
| `.ni/session.json` | no diff. | No direct edit | Protected. |
| `.ni/plan.lock.json` | no diff. | No direct edit | Protected. |
| unexpected files | none expected. | No | Validation 후 recheck. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | `docs/51_POST_RELEASE_ROADMAP*` modifications와 new `docs/130*` files만 표시한다. |
| `git log --oneline --decorate -20` | Edit 전 checked; current HEAD was `e23653b Add v0.5.1 artifact dry-run docs`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; no v0.5.1 tag exists. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Checked; docs/126 through docs/129 are tracked in the current patch delta. |
| `git diff --stat v0.5.0..HEAD` | Checked; docs/130 전 68 files changed. |
| Required ripgrep scans | Release, version, install, Homebrew, runtime boundary surfaces reviewed; CHANGELOG.md and RELEASE.md는 존재하지 않아 ripgrep이 missing file을 report했다. |
| `gofmt -w .` | Passed; 이 task는 Go source change를 introduce하지 않았다. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root는 `NI Intent Readiness: READY`, blockers/deferrals/warnings none을 report했다. |
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
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md` | English release notes finalization, release notes draft, claim-boundary audit, deferrals, validation plan, next task prompt 추가. |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.ko.md` | 같은 boundaries를 가진 Korean companion 추가. |
| `docs/51_POST_RELEASE_ROADMAP.md` | docs/129와 docs/130에 대한 narrow pointer 추가. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Matching Korean pointers 추가. |

## What this finalization proves

- v0.5.1 release notes는 audited boundaries 아래 ready다.
- Notes는 publication을 claim하지 않고 patch를 describe한다.
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
