# v0.5.1 Post-Release Verification

## Current status

State:
- v0.5.1 release: 이 문서에서 published and verified.
- v0.5.0 publication: verified.
- Public install parity mismatch: tested macOS arm64 path에서는 v0.5.1로 addressed.
- v0.5.1 release notes decision: V0_5_1_RELEASE_NOTES_READY_WITH_NOTES.
- v0.5.1 publication checklist decision: V0_5_1_PUBLICATION_CHECKLIST_READY_WITH_NOTES.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: macOS-only development host에서는 deferred.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Verification goal

이 문서는 approved release execution 이후 실제 v0.5.1 publication과 public install
behavior를 verify한다. Hosted release evidence와 source-tree validation을
분리하고 remaining deferrals를 명시적으로 유지한다.

## Decision

V0_5_1_RELEASE_EXECUTED_WITH_NOTES

Justification: tag, GitHub Release, hosted assets, checksums, current-platform
hosted artifact, isolated `install.sh --version 0.5.1` path가 모두 verify되었다.
Notes는 Windows real-host execution, Homebrew availability, external user
validation, model workspace host behavior, no-terminal deterministic validation,
local GoReleaser full-matrix dry-run을 아직 claim하지 않기 때문에 남는다.

## Release identity

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| tag | `v0.5.1` | Local tag exists; remote annotated tag는 `945bd46fdf94872ba921c63d4a2fc43ea54de3be`로 dereference된다. | Yes | Remote tag object는 `b588f6b2e13111841081d186bd0e70d3c0bfbd6c`. |
| release URL | `https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1` | `https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1` | Yes | `gh release view`로 verify. |
| release title | `v0.5.1` | `v0.5.1` | Yes | Normal GitHub Release. |
| draft status | Not draft | `false` | Yes | GitHub release metadata. |
| prerelease status | Not prerelease | `false` | Yes | GitHub release metadata. |
| latest status | v0.5.1 if intended | Default `gh release view`가 `v0.5.1`을 반환. | Yes | Latest lookup이 v0.5.1 URL을 반환. |
| published time | Present | `2026-06-08T00:50:40Z` | Yes | GitHub absolute timestamp. |
| release commit | Intended release commit | `945bd46fdf94872ba921c63d4a2fc43ea54de3be` | Yes | Release execution 당시 `main` HEAD. |
| release body | Concise docs/130-based notes | GoReleaser default changelog 이후 GitHub Release body를 업데이트. | Yes | Homebrew, Windows, model workspace, no-terminal, runtime boundaries preserved. |

## Asset inventory

| Asset | Platform | Size | Downloaded? | Checksum verified? | Execution verified? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `ni_0.5.1_checksums.txt` | checksums | 471 | Yes | n/a | n/a | Hosted checksum file. |
| `ni_0.5.1_darwin_amd64.tar.gz` | macOS amd64 | 1971271 | Yes | Yes | Not run | 이 host의 current platform이 아님. |
| `ni_0.5.1_darwin_arm64.tar.gz` | macOS arm64 | 1845991 | Yes | Yes | Yes | Current platform artifact. |
| `ni_0.5.1_linux_amd64.tar.gz` | Linux amd64 | 1980921 | Yes | Yes | Not run | Cross-platform execution claim 없음. |
| `ni_0.5.1_linux_arm64.tar.gz` | Linux arm64 | 1812172 | Yes | Yes | Not run | Cross-platform execution claim 없음. |
| `ni_0.5.1_windows_amd64.zip` | Windows amd64 | 2075537 | Yes | Yes | Not run | Windows real-host execution은 deferred. |

## Checksum verification

| File | Expected checksum | Actual checksum | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni_0.5.1_darwin_amd64.tar.gz` | `19459e19888a5043a1ee70e76696ad04990c35e987a35dbe5fab1472f67f1958` | `shasum -a 256 -c` 결과 동일 | Yes | Hosted asset matched checksum file. |
| `ni_0.5.1_darwin_arm64.tar.gz` | `bb97c7831cb4c906b6140a4f05e01823b23111c9f917996d6e0874c23f158c73` | `shasum -a 256 -c` 결과 동일 | Yes | Current platform archive. |
| `ni_0.5.1_linux_amd64.tar.gz` | `cdb4aca68beb5ae2397215c898bd135546f57ff140046e900632339bc53b35d1` | `shasum -a 256 -c` 결과 동일 | Yes | Execution claim 없음. |
| `ni_0.5.1_linux_arm64.tar.gz` | `29c3030e18152f0a4877c9d665d833e049d5b0498e9dbd3550a28777ff4bc1c4` | `shasum -a 256 -c` 결과 동일 | Yes | Execution claim 없음. |
| `ni_0.5.1_windows_amd64.zip` | `c44803197b8c0ae5401de71794f78b58708d639d0d25ca2b41f2d6e414ad435a` | `shasum -a 256 -c` 결과 동일 | Yes | Windows execution claim 없음. |

## Hosted artifact verification

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| `ni --help` | `/tmp/ni-v0.5.1-post-release-verify/extracted-darwin-arm64/ni --help` | Help exits 0 and includes `ni init [.]`. | Help exited 0 and listed `ni init [.]`. | Yes | Hosted darwin/arm64 artifact. |
| `ni version` | `/tmp/ni-v0.5.1-post-release-verify/extracted-darwin-arm64/ni version` | `0.5.1` | `0.5.1` | Yes | Release linker version verified. |
| `ni init . --yes` | Hosted binary in temp project | Planning files created. | `.ni/contract.json`, `.ni/session.json`, `.ni/readiness.*`, `.ni/pressure.json`, `.ni/harness.candidates.json`, `docs/plan/**` 생성. | Yes | Downstream execution 없음. |
| `ni status --proof --next-questions` | Hosted binary in initialized temp project | Empty intent에 대해 honest first-run `BLOCKED`. | `NI Intent Readiness: BLOCKED` with R014, OQ-001, R015, R016 blockers. | Yes | Expected first-user state. |

## install.sh verification

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| `install.sh --version 0.5.1` | `HOME=/tmp/... BINDIR=/tmp/... sh install.sh --version 0.5.1` | Hosted darwin/arm64 asset download, checksum verify, temp `ni` install. | `ni_0.5.1_darwin_arm64.tar.gz` download, checksum verify, temp binary install. | Yes | User install state를 건드리지 않음. |
| command-name `ni --help` | Fresh `zsh -f` with temp BINDIR first on PATH | `ni` resolves from temp BINDIR and help exits 0. | `command -v ni`가 temp path 출력; help exit 0. | Yes | Command-name verification. |
| command-name `ni version` | Fresh `zsh -f` with temp BINDIR first on PATH | `0.5.1` | `0.5.1` | Yes | Installed artifact verified. |
| installed `ni init . --yes` | Temp project with temp BINDIR on PATH | Planning files created. | Init created expected planning files. | Yes | Tested host의 public first-project path 확인. |
| installed `ni status --proof --next-questions` | Temp installed binary in temp project | Honest first-run `BLOCKED`. | `BLOCKED` with first-run blockers and next questions. | Yes | Blank-intent expected result. |
| uninstall | Same temp HOME/BINDIR with `sh install.sh --uninstall` | Remove temp binary. | `/tmp/ni-v0.5.1-post-release-verify/install-bin/ni` removed. | Yes | User binary 제거 없음. |
| cleanup | `test -e temp/bin/ni`; fresh `command -v ni` | Binary absent and temp command not found. | 두 command 모두 expected non-zero. | Yes | Temp verification directory는 audit trail로 남김. |

## Public install parity closure

| Surface | v0.5.0 behavior | v0.5.1 behavior | Closed? | Notes |
| --- | --- | --- | --- | --- |
| `ni init .` | `unknown init option: .`로 fail. | Hosted and installed v0.5.1 pass `ni init . --yes`. | Yes for tested macOS arm64 path | Non-TTY verification이라 interactive TUI 자체는 exercise하지 않음. |
| README first-user path | v0.5.0은 current README first-project command와 맞지 않음. | v0.5.1은 `ni init .`과 init 이후 status proof 지원. | Yes for tested macOS arm64 path | Blank intent는 정직하게 `BLOCKED`. |
| `install.sh` v0.5.1 retrieval | Publication 전에는 존재할 수 없었음. | Hosted v0.5.1 retrieve, checksum verify, temp command install 성공. | Yes | Isolated HOME/BINDIR. |
| current-platform artifact | v0.5.0에는 current `ni init .` support 없음. | Hosted darwin/arm64 artifact가 `0.5.1` report, init/status smoke pass. | Yes | 다른 platforms는 checksum-only. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks v0.5.1 verification? |
| --- | --- | --- | --- |
| Windows real-host execution | Current host는 macOS. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | Tap/formula/install proof 없음. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | 이 task에는 separate user/machine transcript 없음. | External install/init/status transcript. | No |
| model workspace host behavior | Host-level/global install과 provider behavior는 unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Target workspace에 대한 trusted CLI proof. | No |
| GoReleaser local full matrix dry-run | Release workflow는 GoReleaser 성공; local `goreleaser` binary는 unavailable. | 필요시 local GoReleaser check/dry-run matrix. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Release evidence 이후에만 claim. | `2026-06-08T00:50:40Z`에 publish. | Yes | Release URL verified. |
| hosted artifacts | Asset inventory와 checksums 이후에만 claim. | 6개 assets visible; all checksums passed. | Yes | 5 archives plus checksum file. |
| `install.sh` v0.5.1 retrieval | Isolated install proof 이후에만 claim. | Isolated install passed. | Yes | Temp HOME/BINDIR only. |
| Homebrew | Planned / v0.5 candidate 유지. | Preserved. | Yes | Homebrew Available claim 없음. |
| Windows real-host execution | Transcript 없으면 deferred. | Preserved. | Yes | Windows asset checksum only. |
| `ni init .` | Tested v0.5.1 macOS arm64 path에 대해 claim 가능. | Hosted and installed paths passed. | Yes | Cross-platform execution claim 없음. |
| `ni run` | Bounded prompt compilation only. | Release body and docs에서 preserved. | Yes | Generated prompt execution 없음. |
| READY | CLI readiness only, product readiness 아님. | Blank temp workspaces는 올바르게 `BLOCKED`. | Yes | Product readiness claim 없음. |
| runtime execution boundary | Task runner, SPEC runner, execution harness, shell/Codex adapter, queue, PR automation, release automation system, downstream execution layer 없음. | Preserved. | Yes | Release process는 distribution infrastructure. |

## Git status / inclusion check

| Path or group | git status --short | Expected after release? | Notes |
| --- | --- | --- | --- |
| README.md | post-release docs update로 modified | Yes | Version and parity note updated to v0.5.1. |
| README.ko.md | post-release docs update로 modified | Yes | Korean companion. |
| docs/126* | post-release docs update로 modified | Yes | Narrow addressed-by-v0.5.1 note. |
| docs/130* | post-release docs update로 modified | Yes | Narrow post-release pointer. |
| docs/131* | post-release docs update로 modified | Yes | Narrow post-release pointer. |
| docs/132* | Added | Yes | Post-release verification docs. |
| temporary artifact directory | Repo 밖 `/tmp/ni-v0.5.1-post-release-verify` | No commit | Audit trail로 repo 밖에 남김. |
| `.ni/contract.json` | No diff | No direct edit | Protected. |
| `.ni/session.json` | No diff | No direct edit | Protected. |
| `.ni/plan.lock.json` | No diff | No direct edit | Protected. |
| unexpected files | None expected | No | Final validation에서 recheck. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Release execution 전 clean; verification 이후 post-release docs changed. |
| `git branch --show-current` | `main`. |
| `git remote -v` | `origin`은 `https://github.com/Nam-Cheol/ni.git`. |
| `git log --oneline --decorate -20` | Release HEAD는 `945bd46 Add v0.5.1 publication checklist roadmap links`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` before release | Empty. |
| `git ls-remote --tags origin v0.5.1` before release | Empty. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | README/TUI/install changes and docs/126 through docs/131 포함. |
| `git diff --stat v0.5.0..HEAD` | Release execution 전 72 files changed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Release 전과 verification 후 모두 empty. |
| `gofmt -w .` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; root는 blockers/deferrals/warnings 없이 `NI Intent Readiness: READY`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Release execution 전 passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| current-tree first-user smoke | Temp source binary로 passed; repeated init은 files preserved; synthetic lockfile was not modified. |
| `git push origin v0.5.1` | Passed; tag pushed. |
| `gh run watch 27109991213 --repo Nam-Cheol/ni --exit-status` | Passed; workflow completed successfully in 1m19s. |
| `gh release view v0.5.1 --repo Nam-Cheol/ni` | Passed; release metadata and assets verified. |
| `shasum -a 256 -c ni_0.5.1_checksums.txt` | Downloaded hosted archives 전체 passed. |
| hosted artifact smoke | Darwin/arm64 help, version, init, status passed. |
| isolated `install.sh --version 0.5.1` | Passed; checksum verified, temp command installed, version `0.5.1`, init/status passed, uninstall removed temp binary. |

## Changes made

| File | Why |
| --- | --- |
| `README.md` | Verified install version and first-project parity note를 v0.5.1로 update. |
| `README.ko.md` | Korean companion update. |
| `docs/22_INSTALL.md` | Release binary and curl installer evidence를 v0.5.1로 update. |
| `docs/install-curl.md` | Installer verification wording을 v0.5.1로 update. |
| `docs/install-curl.ko.md` | Korean companion update. |
| `docs/51_POST_RELEASE_ROADMAP.md` | docs/132 post-release verification pointer 추가. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Korean companion update. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md` | v0.5.1이 tested path의 v0.5.0 mismatch를 addressed했다는 narrow note 추가. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.ko.md` | Korean companion update. |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md` | Post-release pointer 추가. |
| `docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.ko.md` | Korean companion update. |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.md` | Post-release pointer 추가. |
| `docs/131_V0_5_1_PUBLICATION_CHECKLIST.ko.md` | Korean companion update. |
| `docs/132_V0_5_1_POST_RELEASE_VERIFICATION.md` | English post-release verification 추가. |
| `docs/132_V0_5_1_POST_RELEASE_VERIFICATION.ko.md` | Korean companion 추가. |
| `scripts/check-install-docs.py` | Install-doc markers를 v0.5.1로 update. |
| `scripts/release-check.sh` | Current release marker와 v0.5.1 post-release guard update. |

## What this verification proves

- v0.5.1 was published.
- Hosted assets exist.
- Hosted checksums verify.
- Current-platform hosted artifact reports `0.5.1`.
- `install.sh` retrieves and installs v0.5.1 in an isolated temp path.
- Public install parity mismatch는 tested macOS arm64 path에서 closed.

## What this verification does not prove

- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.
- Every platform executes correctly.

## Recommended next task

Selected next task: A. external user validation plan

Selection rationale: release publication과 tested macOS arm64 install parity는
verify되었다. 가장 값진 remaining proof gap은 external user 또는 separate
machine transcript다. Windows real-host execution과 Homebrew는 별도 future lane으로
남긴다.

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
