# v0.5.1 Artifact Dry-Run

## Current status

State:
- v0.5.0 publication: verified.
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES.
- v0.5.1 RC validation decision: V0_5_1_RC_VALIDATION_PASS_WITH_NOTES.
- v0.5.1 release: not published.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: macOS-only development host에서는 deferred.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Dry-run goal

이 dry-run은 current release-candidate tree가 release-like v0.5.1
current-platform artifact를 만들 수 있는지 확인한다. 핵심 확인 항목은 binary가
`0.5.1`을 report하는지, checksum을 local에서 generate/verify할 수 있는지,
first-user artifact behavior가 kernel boundary를 유지하는지다. 이 task는 publish,
tag, upload, GitHub Release creation, release workflow 실행, Homebrew formula 생성
또는 publish, project root `ni end`, project root relock, generated prompt
execution, `ni run` downstream execution 변경을 수행하지 않는다.

## Decision

V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES

Justification: release-like current-platform `darwin/arm64` artifact를 repo 밖에서
`ni/internal/version.Version=0.5.1`로 build했고, binary는 `0.5.1`을 report했으며,
help/init/status smoke checks가 passed했고, SHA-256 checksum을 local에서 생성하고
verify했다. Notes는 남는다. Local GoReleaser는 installed되어 있지 않고, full release
matrix는 GoReleaser로 생성하지 않았으며, v0.5.1은 not published이고, hosted v0.5.1
assets는 없고, `install.sh` public retrieval은 publication 전에는 성공할 수 없고,
Windows real-host execution은 deferred이고, Homebrew는 Planned / v0.5 candidate이며,
external user validation은 수행하지 않았다.

## Repository state

| Surface | Observed state | Notes |
| --- | --- | --- |
| git status | Task 시작 시 clean; 이 task는 docs/129 documentation만 uncommitted로 추가한다. | Staging/commit 없음. |
| v0.5.0 tag | Present; `git rev-parse v0.5.0` returned `b8fec7fa9615a861acf4eba59733c800c70c6cca`. | Verified baseline. |
| v0.5.1 tag | Absent. | Tag 생성하지 않음. |
| diff `v0.5.0..HEAD` | docs/129 전 기준 66 files changed, 9184 insertions, 474 deletions. | Current tree는 v0.5.1 patch/RC work를 포함한다. |
| protected `.ni` diff | Empty. | `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` unchanged. |
| generated artifacts | Release-like artifacts는 `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK` 아래에 있다. | Repository 밖. |

docs/126, docs/127, docs/128은 current tree에서 tracked 상태다.

## Release tooling audit

| Surface | Observed state | Pass? | Notes |
| --- | --- | --- | --- |
| version command | `cmd/ni/main.go`가 `version.Version`을 출력한다. | Yes | Source default는 `0.0.0-dev`로 유지. |
| version injection | `.goreleaser.yaml`이 `-X ni/internal/version.Version={{ .Version }}`를 사용한다. | Yes | Local fallback은 `-X ni/internal/version.Version=0.5.1`을 사용했다. |
| release-check script | `scripts/release-check.sh` present. | Yes | Publishing command가 아니라 check-only release surface gate. |
| install-check script | `scripts/install-check.sh` present. | Yes | Source/local install validation surface. |
| GoReleaser config | `.goreleaser.yaml` present. | Yes | `linux/amd64`, `linux/arm64`, `darwin/amd64`, `darwin/arm64`, `windows/amd64`를 정의하고 `windows/arm64`는 ignore. |
| GitHub workflow | `.github/workflows/release.yml` present. | Yes | Tag-triggered release path이며 이 dry-run에서는 실행하지 않음. |
| archive naming | `ni_{{ .Version }}_{{ .Os }}_{{ .Arch }}` with `tar.gz`, Windows `zip`. | Yes | Current-platform local archive는 `ni_0.5.1_darwin_arm64.tar.gz`. |
| checksum generation | GoReleaser config name은 `ni_{{ .Version }}_checksums.txt`. | Yes | Local fallback은 `shasum -a 256`으로 `ni_0.5.1_checksums.txt` 생성. |
| GoReleaser availability | `goreleaser --version` failed with command not found. | No | Full GoReleaser dry-run은 deferred; local Go fallback 사용. |

Release config 기준 expected v0.5.1 artifact names:
- `ni_0.5.1_linux_amd64.tar.gz`
- `ni_0.5.1_linux_arm64.tar.gz`
- `ni_0.5.1_darwin_amd64.tar.gz`
- `ni_0.5.1_darwin_arm64.tar.gz`
- `ni_0.5.1_windows_amd64.zip`
- `ni_0.5.1_checksums.txt`

## Artifact inventory

| Artifact | Path | Platform | Version output | Checksum generated? | Checksum verified? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| raw binary | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni` | darwin/arm64 | `0.5.1` | n/a | n/a | Local Go fallback과 version ldflags로 build. |
| archive | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_darwin_arm64.tar.gz` | darwin/arm64 | Archived binary source에서 `0.5.1` | Yes | Yes | `ni` 포함. |
| checksums | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_checksums.txt` | n/a | n/a | Yes | Yes | Local checksum file only; hosted 아님. |

## Current-platform artifact verification

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| `ni --help` | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni --help` | Help renders and includes `ni init [.]`. | Passed. | Yes | Absolute artifact path. |
| `ni version` | `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni version` | `0.5.1`. | `0.5.1`. | Yes | Version injection verified. |
| command-name `ni --help` | `PATH=/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK:$PATH zsh -f -c 'ni --help'` | PATH-resolved help works. | Passed. | Yes | `command -v ni`가 temp artifact를 가리킴. |
| command-name `ni version` | `PATH=/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK:$PATH zsh -f -c 'ni version'` | `0.5.1`. | `0.5.1`. | Yes | Fresh shell command-name check. |
| `ni init . --yes` | Artifact binary in `/private/tmp/ni-v0.5.1-artifact-project.xCKnV8`. | Planning workspace 생성. | `.ni/contract.json`, `.ni/session.json`, `.ni/project.json`, readiness files, harness/pressure seed files, `docs/plan/**` created. | Yes | Lockfile 또는 prompt 생성 없음. |
| command-name `ni init . --yes` | PATH-resolved artifact in `/private/tmp/ni-v0.5.1-artifact-path-project.a47CFC`. | Planning workspace 생성. | Passed. | Yes | Command-name init smoke도 실행. |
| `ni status --proof --next-questions` | Artifact binary in blank temp project. | Status proof를 실행하고 incomplete intent를 block. | `NI Intent Readiness: BLOCKED` with R014, OQ-001, R015, R016. | Yes | Expected first-run scaffold behavior. |

## Generated artifact check

| Artifact | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `.ni/contract.json` | Init 후 exists. | Exists. | Yes | Temp project only. |
| `.ni/session.json` | Init 후 exists. | Exists. | Yes | Temp project only. |
| `docs/plan/**` | Init 후 exists. | Twelve planning docs created. | Yes | Temp project only. |
| `.ni/plan.lock.json` | Init으로 생성되면 안 된다. | Not present. | Yes | Lock action 없음. |
| downstream generated prompt | Init/status가 execute하거나 create하면 안 된다. | `.ni` 아래 `*.prompt*` file 없음. | Yes | `ni run` 실행하지 않음. |

## Checksum verification

| File | SHA256 | Verified? | Command | Notes |
| --- | --- | --- | --- | --- |
| `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_darwin_arm64.tar.gz` | `593adc8f8643e6b89f3400eef2ee11e5eabd3e472847d2efd9897dc6117977f1` | Yes | `shasum -a 256 -c /private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni_0.5.1_checksums.txt` | Local checksum verification returned `OK`. |

Full release archive checksum verification은 GoReleaser 또는 publication task로 deferred다.
이번 dry-run은 current-platform local archive만 생성했다.

## install.sh future v0.5.1 gate

| Check | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `install.sh --dry-run --version 0.5.1` | Install 없이 v0.5.1 asset/checksum names를 구성. | Passed. | Yes | Dry-run only. |
| expected URL | `https://github.com/Nam-Cheol/ni/releases/download/v0.5.1/...` 사용. | `ni_0.5.1_darwin_arm64.tar.gz`와 `ni_0.5.1_checksums.txt` URLs printed. | Yes | URLs는 future-publication gates. |
| public URL availability | Release publication 전 claim 금지. | Available로 check하거나 claim하지 않음. | Yes | v0.5.1 release는 not published. |
| post-publication install gate | Hosted asset 설치 후 command-name help/version verify 필요. | Deferred. | Yes | Publication 후 required. |

## Release notes delta

v0.5.1 release notes delta는 docs/127과 docs/128에 기록된 patch scope를 유지한다:
- Fix: installed `ni --help`와 `ni version`은 passed했지만 `ni init .`은 failed한
  public v0.5.0 install parity mismatch.
- Add: `ni init .` positional target support.
- Add: Bubble Tea v2 / Lip Gloss v2 guided init TUI.
- Preserve: CI/non-TTY contexts를 위한 non-interactive fallback.
- Preserve: init 중 existing-file 및 `.ni/plan.lock.json` protection.
- Improve: README macOS / Windows two-path onboarding and command-name
  verification guidance.
- Exclude: downstream execution, Homebrew Available, Windows real-host proof,
  deterministic no-terminal claims.

이 delta는 actual v0.5.1 artifacts와 installer path가 publication 뒤 verified되기
전까지 draft-only다.

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks artifact dry-run? |
| --- | --- | --- | --- |
| v0.5.1 publication | 이 task는 non-publishing. | Tag, release workflow, hosted assets, post-publication checks. | No |
| hosted assets | GitHub Release를 만들지 않았다. | Release page asset inventory. | No |
| `install.sh` actual v0.5.1 retrieval | Public assets가 아직 없다. | Hosted v0.5.1 asset isolated install plus `ni --help` and `ni version`. | No |
| Windows real-host execution | Current development host가 macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | Published tap/formula/install proof 없음. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | Separate user 또는 machine transcript 없음. | External install/init/status transcript. | No |
| model workspace host behavior | Host-level/global install과 provider behavior unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Target workspace의 trusted CLI proof. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Preserved. | Yes | Release action 없음. |
| artifact dry-run | Local current-platform release-like artifact만 claim 가능. | Preserved. | Yes | Full GoReleaser matrix claim 없음. |
| `install.sh` v0.5.1 retrieval | Publication 전 public retrieval claim 금지. | Preserved. | Yes | Dry-run URL construction only. |
| current-tree behavior | Local artifact behavior만 claim 가능. | Help/version/init/status passed. | Yes | Hosted artifact claim 아님. |
| Homebrew | Planned / v0.5 candidate 유지. | Preserved. | Yes | Homebrew Available claim 없음. |
| Windows real-host execution | Transcript 없으면 deferred. | Preserved. | Yes | macOS host only. |
| `ni init .` | Guided setup only; no execution. | Planning workspace를 만든다. | Yes | Downstream prompt 없음. |
| `ni run` | Bounded prompt compilation only. | 이 dry-run에서는 실행하지 않음. | Yes | Execution behavior 추가 없음. |
| READY | CLI authority only. | Model-only readiness claim 없음. | Yes | Temp scaffold는 correctly BLOCKED. |
| runtime execution boundary | Task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, downstream execution layer 없음. | Preserved. | Yes | Runtime behavior 추가 없음. |

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
| unexpected files | docs/129 전 repo status에서 none observed. | No | Dry-run files stayed outside repo. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; task 시작 시 clean. |
| `git log --oneline --decorate -20` | Passed; current HEAD `842e565`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; v0.5.1 tag absent. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Passed; docs/126, docs/127, docs/128 tracked 확인. |
| `git diff --stat v0.5.0..HEAD` | Passed; docs/129 전 기준 66 files changed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |
| required ripgrep scans | Release/version/install/Homebrew/runtime boundary surfaces reviewed. |
| `goreleaser --version` | Failed with `command not found`; local Go fallback 사용. |
| `env GOCACHE=/private/tmp/ni-go-cache CGO_ENABLED=0 go build -ldflags='-s -w -X ni/internal/version.Version=0.5.1' -o /private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni ./cmd/ni` | Passed. |
| `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni --help` | Passed. |
| `/private/tmp/ni-v0.5.1-artifact-dry-run.QaUQlK/ni version` | Passed; `0.5.1`. |
| command-name `ni --help` | Temp artifact를 PATH 앞에 두고 Passed. |
| command-name `ni version` | Passed; `0.5.1`. |
| artifact `ni init . --yes` | Temp project에서 Passed. |
| artifact `ni status --proof --next-questions` | Passed; blank intent에 대해 correctly `BLOCKED`. |
| command-name `ni init . --yes` | Second temp project에서 Passed. |
| `tar -czf .../ni_0.5.1_darwin_arm64.tar.gz -C ... ni` | Passed. |
| `shasum -a 256 .../ni_0.5.1_darwin_arm64.tar.gz > .../ni_0.5.1_checksums.txt` | Passed. |
| `shasum -a 256 -c .../ni_0.5.1_checksums.txt` | Passed; returned `OK`. |
| `sh install.sh --dry-run --version 0.5.1` | Passed; future v0.5.1 asset/checksum URLs printed. |

## Changes made

| File | Why |
| --- | --- |
| `docs/129_V0_5_1_ARTIFACT_DRY_RUN.md` | English v0.5.1 artifact dry-run evidence, decision, deferrals, next task 추가. |
| `docs/129_V0_5_1_ARTIFACT_DRY_RUN.ko.md` | 같은 claim boundaries를 가진 Korean companion 추가. |

## What this dry-run proves

- Local release-like v0.5.1 current-platform artifact behavior passed.
- Artifact `ni version` reports `0.5.1`.
- Current-platform archive checksum을 local에서 generate/verify했다.
- Current-platform artifact는 first-user init/status path를 support한다.
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

Selection rationale: artifact dry-run은 notes와 함께 passed다. 따라서 다음
pre-publication step은 release notes finalization이며, publication, hosted installer
retrieval, Windows real-host execution, Homebrew, external validation은 계속 explicit
gates로 유지해야 한다.
