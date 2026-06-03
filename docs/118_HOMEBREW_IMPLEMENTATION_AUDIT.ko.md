# Homebrew Implementation Audit

## Current status

- v0.5.0 publication: verified
- Post-release verification decision: V0_5_0_POST_RELEASE_VERIFIED_WITH_NOTES
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- This audit does not publish a Homebrew formula or mark Homebrew Available.

## Audit goal

이 audit는 verified v0.5.0 release 이후 Homebrew implementation을 안전하게
진행할 수 있는지 결정한다. Current tap/formula state, recommended formula
source, verified sha256, local validation boundary, future availability gate를
분리해 기록한다.

## Decision

Decision: HOMEBREW_IMPLEMENTATION_READY_WITH_DEFERRALS.

Justification: v0.5.0 release evidence는 source-build formula draft를 만들기에
충분하고, `packaging/homebrew-draft/ni.rb`에 verified source archive sha256을
포함한 local non-published draft를 추가했다. 하지만 public tap repository가
없고, formula가 published되지 않았고, real tap install이 실행되지 않았으므로
Homebrew는 아직 Available이 아니다.

Follow-up: [`119_HOMEBREW_FORMULA_LOCAL_VERIFICATION.ko.md`](119_HOMEBREW_FORMULA_LOCAL_VERIFICATION.ko.md)는
later local verification attempt를 기록한다. 그 pass는 Homebrew를 Planned로
유지했고, Homebrew 5.1.14가 path-based audit와 path-based install을 formula
binary 설치 전에 거부했기 때문에 `HOMEBREW_FORMULA_LOCAL_BLOCKED`를 선택했다.

## Homebrew status audit

| Surface | Observed state | Pass? | Notes |
| --- | --- | --- | --- |
| tap repository | `Nam-Cheol/homebrew-tap`은 `git ls-remote`와 `gh repo view`에서 not found. | Yes | Publication remains future work. |
| formula file | Published formula 없음; local draft는 `packaging/homebrew-draft/ni.rb`. | Yes | Draft only; tap 아님. |
| README status | `README.md`는 Homebrew를 Planned로 유지한다. | Yes | `brew install` availability claim 없음. |
| README.ko status | `README.ko.md`도 Planned이며 아직 `brew install`을 사용하지 말라고 말한다. | Yes | Korean companion bounded. |
| docs/22 install status | Package manager status: Planned. | Yes | Package-manager instruction을 usable로 제시하지 않음. |
| roadmap status | Homebrew는 tap/formula/install proof로 gated된 v0.5 candidate. | Yes | docs/118 pointer 추가. |
| post-release verification status | docs/117은 Homebrew를 Planned / v0.5 candidate로 기록. | Yes | No Homebrew proof claimed. |
| current user-facing install command | 없음. | Yes | Future-gated. |

## Recommended implementation path

| Decision area | Recommendation | Rationale | Risk | Notes |
| --- | --- | --- | --- | --- |
| tap repository name | `Nam-Cheol/homebrew-tap` | Existing docs and Homebrew short tap convention. | Owner may choose another name. | Command wording must follow final tap. |
| formula name | `ni` | Matches CLI binary. | Naming conflict unverified. | Prefer fully qualified `Nam-Cheol/tap/ni` after publication. |
| source choice | GitHub tag source archive for `v0.5.0` | One URL and one sha256; source build is Homebrew-friendly. | Recheck before publication. | URL below. |
| sha256 source | Exact source archive URL | No invented checksum. | Recompute if URL changes. | Verified in this audit. |
| build strategy | Go source build | Simpler than per-platform prebuilt archive selection. | Requires Go build dependency. | `depends_on "go" => :build`. |
| install command shape | Future `brew install Nam-Cheol/tap/ni` | Fully qualified command avoids ambiguity. | Not valid until published and tested. | Do not add as Available wording. |
| uninstall command shape | Future `brew uninstall ni` | Standard Homebrew cleanup. | Not verified until install runs. | Future proof required. |
| test block | `ni --help` and `ni version` | Checks CLI and version. | Tap publication issues may remain. | Draft includes both. |
| publication gate | Tap, audit, install, help/version, uninstall, docs update | Prevents overclaim. | Separate owner-controlled task. | This task does not publish. |

## Formula draft status

| Item | Status | Evidence | Notes |
| --- | --- | --- | --- |
| draft formula path | Created | `packaging/homebrew-draft/ni.rb` | Local draft only. |
| URL | Verified | `https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz` | Source archive. |
| sha256 | Verified | `67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02` | Downloaded archive에서 computed. |
| dependency declaration | Present | `depends_on "go" => :build` | Build-from-source. |
| install block | Present | `go build -trimpath ... ./cmd/ni` | Injects `0.5.0`. |
| test block | Present | `ni --help`, `ni version` | Install validation not yet run. |
| local audit | Deferred | Homebrew 5.1.14에서 path-based local audit가 disabled. | See validation section. |
| local install | Not run | Mutation risk. | Deferred. |
| uninstall cleanup | Not run | No install. | Future proof required. |

## SHA256 evidence

| Source URL | SHA256 | Verified? | Command | Notes |
| --- | --- | --- | --- | --- |
| `https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz` | `67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02` | Yes | `curl -L ... -o /private/tmp/ni-v0.5.0-source.tar.gz`; `shasum -a 256 /private/tmp/ni-v0.5.0-source.tar.gz` | Draft formula source. |

## Local Homebrew validation

| Command | Ran? | Result | Mutation risk | Notes |
| --- | --- | --- | --- | --- |
| `brew --version` | Yes | Pass: `Homebrew 5.1.14` | None | Local brew available. |
| `brew audit --strict --online --new-formula packaging/homebrew-draft/ni.rb` | Yes | Failed before audit: 현재 Homebrew는 `--new-formula`가 아니라 `--new`를 사용하며, sandbox에서는 Homebrew cache directory 생성도 막혔다. | Low | Corrected command는 아래 기록. |
| `brew audit --strict --online --new packaging/homebrew-draft/ni.rb` | Yes | Deferred: Homebrew 5.1.14가 path-based `brew audit [path ...]`를 disabled로 보고했다. | Low | Formula install은 하지 않았고, audit가 켠 developer mode는 `brew developer off`로 cleanup했다. |
| `ruby -c packaging/homebrew-draft/ni.rb` | Yes | Pass: `Syntax OK`. | None | Syntax-only check; Homebrew formula audit는 아니다. |
| `brew install --build-from-source --formula ./packaging/homebrew-draft/ni.rb` | No | Deferred | Medium | Would mutate user Homebrew state. |
| installed `ni --help` | No | Deferred | Medium | Requires install. |
| installed `ni version` | No | Deferred | Medium | Requires install. |
| `brew uninstall` cleanup | No | Deferred | Medium | No install performed. |

## Homebrew availability gate

Homebrew can become Available only after all of these future evidence items exist:

- public tap repository
- committed formula
- verified URL
- verified sha256
- brew audit result
- brew install output
- installed ni --help
- installed ni version
- brew uninstall cleanup
- README / README.ko update
- install docs update

## Blockers

None.

| Blocker | Evidence | Required fix |
| --- | --- | --- |
| None | v0.5.0 release assets exist, source archive sha256 was verified, and a local draft formula exists. | n/a |

## Deferrals

| Deferral | Reason | Required future evidence | Blocks current audit? |
| --- | --- | --- | --- |
| tap publication | `Nam-Cheol/homebrew-tap` does not exist. | Public tap and owner confirmation. | No |
| real user-facing brew install | No published tap/formula. | `brew install Nam-Cheol/tap/ni` output. | No |
| cross-machine Homebrew verification | Only this macOS host audited. | Clean host transcript. | No |
| Homebrew docs promotion to Available | Evidence incomplete. | Docs update after full proof. | No |
| Homebrew uninstall verification | No Homebrew install performed. | `brew uninstall ni` transcript. | No |

## Warnings

| Warning | Evidence | Mitigation |
| --- | --- | --- |
| Local draft is not a tap formula | Stored inside `ni-kernel`, not `Nam-Cheol/homebrew-tap`. | Keep draft marked and do not document it as a user install path. |
| Local `brew audit` cannot audit a path | Homebrew 5.1.14 reports path-based `brew audit [path ...]` is disabled. | Publication 전 external tap에서 formula name 또는 supported tap workflow로 audit한다. |

## Risks

| Risk | Impact | Follow-up |
| --- | --- | --- |
| formula may work locally but fail in published tap | Users could hit install failures. | Test actual public tap before docs promotion. |
| source archive sha may change if wrong URL type is used | Formula checksum mismatch. | Pin final URL and recompute during publication. |
| user-facing install command may change depending on tap name | README may point to wrong command. | Confirm tap name before docs update. |
| Homebrew analytics / formula naming conflicts are unverified | Qualified wording may be needed. | Check before public docs update. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Planned / v0.5 candidate until full proof. | Preserved; local draft only. | Yes | No Available claim. |
| Release binary | Available after verified v0.5.0 assets/checksums. | Preserved. | Yes | docs/117 evidence. |
| Curl installer | Available after isolated v0.5.0 installer verification. | Preserved. | Yes | No change. |
| Model workspace packs | Experimental unless host-level install/discovery is verified. | Preserved. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | Preserved. | Yes | No deterministic claim. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | No generated prompt executed. |
| Benchmark evidence | No implementation-quality or downstream-quality proof. | Preserved. | Yes | `not_measured` boundary unchanged. |
| Runtime execution boundary | Not task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. | Preserved. | Yes | Formula is distribution-only draft. |

## Git status / inclusion check

| Path or group | git status --short | git ls-files / tracked check | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| README.md | clean | tracked | No | No overclaim found. |
| README.ko.md | clean | tracked | No | No overclaim found. |
| docs/22_INSTALL.md | clean | tracked | No | Package manager remains Planned. |
| docs/install-curl* | clean | tracked | No | No change needed. |
| docs/51* | modified | tracked | Yes | Adds docs/118 pointer. |
| docs/117* | clean | tracked | No | Existing recommendation points here. |
| docs/118* | added | untracked until added | Yes | New audit docs. |
| packaging/homebrew-draft/* | added | untracked until added | Yes | Local draft formula. |
| .ni/contract.json | clean | tracked protected file | No | Must remain unchanged. |
| .ni/session.json | clean | tracked protected file | No | Must remain unchanged. |
| .ni/plan.lock.json | clean | tracked protected file | No | Must remain unchanged. |
| unexpected files | none expected | n/a | No | Final status should show only scoped files. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean; final status reviewed after edits. |
| `go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; no blockers, deferrals, or warnings. |
| `brew --version` | Pass | `Homebrew 5.1.14`. |
| `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git` | Pass for audit | Repository not found. |
| `gh repo view Nam-Cheol/homebrew-tap --json nameWithOwner,url,isPrivate` | Pass for audit | Repository not found. |
| `gh release view v0.5.0 --repo Nam-Cheol/ni --json ...` | Pass | Release metadata and 6 assets returned. |
| `curl -L https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz -o /private/tmp/ni-v0.5.0-source.tar.gz` | Pass | Downloaded source archive outside repo. |
| `shasum -a 256 /private/tmp/ni-v0.5.0-source.tar.gz` | Pass | Produced draft formula sha256. |
| `brew audit --strict --online --new-formula packaging/homebrew-draft/ni.rb` | Fail before audit | Homebrew 5.1.14에서 invalid option; sandbox cache write도 failed. |
| `brew audit --strict --online --new packaging/homebrew-draft/ni.rb` | Deferred | Homebrew 5.1.14 disables path-based `brew audit [path ...]`; formula install은 발생하지 않았다. |
| `brew developer off` | Pass | Audit attempt가 켠 developer mode cleanup. |
| `ruby -c packaging/homebrew-draft/ni.rb` | Pass | `Syntax OK`. |
| `brew install --build-from-source --formula ./packaging/homebrew-draft/ni.rb` | Not run | Deferred to avoid mutating user Homebrew state. |
| `go test ./...` | Pass | All Go packages passed with `GOCACHE=/private/tmp/ni-go-cache`. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack checks passed; global install remains unverified. |
| `bash scripts/demo-check.sh` | Pass | Public demos verified without downstream runtime execution. |
| `rg -n "Homebrew is Available|..." ...` | Reviewed | Matches are negated, future-gated, or required "does not prove" wording; touched doc가 Homebrew Available을 claim하지 않는다. |
| `bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `bash scripts/release-check.sh` | Pass | Release readiness gate passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | Protected project-root `.ni` diff 없음. |

## Changes made

- `packaging/homebrew-draft/ni.rb`: v0.5.0 source-build Homebrew draft 추가.
- `docs/118_HOMEBREW_IMPLEMENTATION_AUDIT.md`: audit 추가.
- `docs/118_HOMEBREW_IMPLEMENTATION_AUDIT.ko.md`: Korean companion audit 추가.
- `docs/51_POST_RELEASE_ROADMAP.md`: docs/118 pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: Korean pointer 추가.

## What this audit proves

- Homebrew implementation is ready to proceed with deferrals.
- A local formula draft was created.
- sha256 was verified for the v0.5.0 source archive.
- Local `brew audit` for the draft remains deferred because Homebrew 5.1.14
  disables path-based audit.
- Homebrew remains Planned until full tap, formula, audit, install, help,
  version, uninstall, and docs-update evidence exists.

## What this audit does not prove

- Homebrew is Available
- public tap exists, unless verified
- formula has been published, unless actually published in another task
- users can install via brew, unless real brew install from tap was verified
- cross-machine Homebrew behavior
- model workspace host behavior
- no-terminal deterministic validation
- downstream execution succeeds

## Recommended next task

Selected next task: Create and validate the external Homebrew tap in a separate
owner-controlled publication task.

Next executable prompt:

```text
Proceed with external Homebrew tap creation and validation for ni.

This task must run outside ni-kernel unless the user explicitly approves each
publication action. Do not change Homebrew from Planned to Available until all
evidence is complete.

Goal:
Create or confirm `Nam-Cheol/homebrew-tap`, add `Formula/ni.rb` from the
reviewed `packaging/homebrew-draft/ni.rb` source-build draft, validate it with
Homebrew, and collect exact evidence for future docs promotion.

Required evidence before any Available claim:
- public tap repository exists
- `Formula/ni.rb` is committed in the tap
- formula URL is final
- formula sha256 is recomputed and verified
- `brew audit --strict --online --new-formula Formula/ni.rb` passes or warnings are documented and fixed
- `brew install Nam-Cheol/tap/ni` works from the published tap
- installed `ni --help` works
- installed `ni version` reports `0.5.0`
- `brew uninstall ni` or equivalent cleanup works
- README.md, README.ko.md, and docs/22_INSTALL.md are updated only after the above evidence exists

Do not:
- run `ni end` on the project root
- relock the project root
- execute generated prompts
- add release automation
- add runtime execution behavior
- mark model workspace packs Available
- claim no-terminal deterministic validation
```
