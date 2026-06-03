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

This audit decides whether Homebrew implementation can proceed safely after the
verified v0.5.0 release. It checks the current tap/formula state, chooses the
smallest formula path, records exact sha256 evidence for the draft source URL,
and keeps the public availability gate separate from local review.

## Decision

Decision: HOMEBREW_IMPLEMENTATION_READY_WITH_DEFERRALS.

Justification: v0.5.0 release evidence is sufficient to draft a source-build
formula, and this task created a local non-published draft at
`packaging/homebrew-draft/ni.rb` with a verified source archive sha256. Homebrew
cannot become Available yet because the public tap repository is absent, no
formula is published, no real tap install was run, and user-facing install docs
must remain Planned until full evidence exists.

Follow-up: [`119_HOMEBREW_FORMULA_LOCAL_VERIFICATION.md`](119_HOMEBREW_FORMULA_LOCAL_VERIFICATION.md)
records the later local verification attempt. That pass kept Homebrew Planned
and selected `HOMEBREW_FORMULA_LOCAL_BLOCKED` because Homebrew 5.1.14 rejected
path-based audit and path-based install before any formula binary was installed.

## Homebrew status audit

| Surface | Observed state | Pass? | Notes |
| --- | --- | --- | --- |
| tap repository | `Nam-Cheol/homebrew-tap` was not found by `git ls-remote` and `gh repo view`. | Yes | This confirms Homebrew is not Available and publication remains future work. |
| formula file | No published tap formula was found; local draft added at `packaging/homebrew-draft/ni.rb`. | Yes | Draft only; not a tap. |
| README status | `README.md` says Homebrew is Planned and no tap or formula is published or tested. | Yes | No `brew install` availability claim. |
| README.ko status | `README.ko.md` mirrors Planned status and says not to use `brew install` yet. | Yes | Korean companion remains bounded. |
| docs/22 install status | `docs/22_INSTALL.md` says package manager status is Planned and Homebrew is not published. | Yes | No package-manager instructions are presented as usable. |
| roadmap status | `docs/51_POST_RELEASE_ROADMAP.md` keeps Homebrew as a v0.5 candidate gated by tap/formula/install proof. | Yes | This audit adds a pointer to docs/118. |
| post-release verification status | `docs/117_V0_5_0_POST_RELEASE_VERIFICATION.md` records Homebrew as Planned / v0.5 candidate. | Yes | No Homebrew evidence was claimed there. |
| current user-facing install command | None. README and install docs do not present `brew install` as ready for users. | Yes | Intended command shape remains future-gated. |

## Recommended implementation path

| Decision area | Recommendation | Rationale | Risk | Notes |
| --- | --- | --- | --- | --- |
| tap repository name | Use `Nam-Cheol/homebrew-tap`. | Matches existing docs/54, docs/72, and Homebrew short tap convention. | Owner may choose another tap name. | If changed, update command wording before publication. |
| formula name | Use `ni`. | Matches CLI binary and project name. | Possible core/tap name collision is unverified. | Use fully qualified `Nam-Cheol/tap/ni` in docs after publication. |
| source choice | Use GitHub source archive for tag `v0.5.0`. | Homebrew can build from source with one URL and one sha256. | GitHub-generated archives should be pinned and rechecked before publication. | URL recorded below. |
| sha256 source | Compute from exact source archive URL. | Avoids inventing checksum and avoids per-platform binary formula complexity. | sha must be recomputed if URL changes. | Verified in this audit. |
| build strategy | Build `./cmd/ni` with Go. | Simpler and more Homebrew-compatible than selecting prebuilt archives per platform. | Requires Go as a build dependency. | Formula declares `depends_on "go" => :build`. |
| install command shape | Future: `brew install Nam-Cheol/tap/ni`. | Fully qualified command avoids ambiguity. | Not valid until tap/formula are published and tested. | Do not put this in README as available yet. |
| uninstall command shape | Future: `brew uninstall ni` after install; fully qualified cleanup may be documented if needed. | Standard Homebrew uninstall path. | Not verified until install is run. | Keep unavailable until real install proof exists. |
| test block | Run `ni --help` and `ni version`. | Confirms CLI shape and version injection. | Local test may not catch all tap publication issues. | Draft formula includes both checks. |
| publication gate | Publish only after tap exists, formula audit passes, local install passes, and docs are updated in English/Korean. | Prevents false availability claims. | Requires a separate owner-controlled task. | This task does not publish. |

## Formula draft status

| Item | Status | Evidence | Notes |
| --- | --- | --- | --- |
| draft formula path | Created | `packaging/homebrew-draft/ni.rb` | Local draft only; not a tap. |
| URL | Verified | `https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz` | Source archive for tag `v0.5.0`. |
| sha256 | Verified | `67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02` | Computed from downloaded archive. |
| dependency declaration | Present | `depends_on "go" => :build` | Build-from-source formula. |
| install block | Present | `go build -trimpath ... ./cmd/ni` | Injects `0.5.0` into `ni version`. |
| test block | Present | Checks `ni --help` and `ni version`. | Formula-level test only; not yet run through install. |
| local audit | Deferred | `brew audit` path-based local audit is disabled in Homebrew 5.1.14. | See validation section. |
| local install | Not run | Deferred due user-environment mutation risk. | `brew install` would affect local Homebrew state. |
| uninstall cleanup | Not run | No Homebrew install was performed. | Future proof required. |

## SHA256 evidence

| Source URL | SHA256 | Verified? | Command | Notes |
| --- | --- | --- | --- | --- |
| `https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz` | `67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02` | Yes | `curl -L ... -o /private/tmp/ni-v0.5.0-source.tar.gz`; `shasum -a 256 /private/tmp/ni-v0.5.0-source.tar.gz` | Source archive used by draft formula. |

## Local Homebrew validation

| Command | Ran? | Result | Mutation risk | Notes |
| --- | --- | --- | --- | --- |
| `brew --version` | Yes | Pass: `Homebrew 5.1.14` | None | Confirms local brew is available. |
| `brew audit --strict --online --new-formula packaging/homebrew-draft/ni.rb` | Yes | Failed before audit: current Homebrew uses `--new`, not `--new-formula`; sandbox also could not create Homebrew cache directories. | Low | Command corrected below. |
| `brew audit --strict --online --new packaging/homebrew-draft/ni.rb` | Yes | Deferred: Homebrew 5.1.14 reports path-based `brew audit [path ...]` is disabled. | Low | This did not install the formula; it enabled developer mode, which was cleaned up with `brew developer off`. |
| `ruby -c packaging/homebrew-draft/ni.rb` | Yes | Pass: `Syntax OK`. | None | Syntax-only check, not a Homebrew formula audit. |
| `brew install --build-from-source --formula ./packaging/homebrew-draft/ni.rb` | No | Deferred | Medium | Would mutate the user's Homebrew installation. |
| installed `ni --help` | No | Deferred | Medium | Requires local Homebrew install. |
| installed `ni version` | No | Deferred | Medium | Requires local Homebrew install. |
| `brew uninstall` cleanup | No | Deferred | Medium | No Homebrew install was performed. |

## Homebrew availability gate

Homebrew can become Available only after all of these future evidence items
exist:

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
| tap publication | `Nam-Cheol/homebrew-tap` does not exist. | Public tap repository and owner confirmation. | No |
| real user-facing brew install | No published tap/formula exists. | `brew install Nam-Cheol/tap/ni` output after publication. | No |
| cross-machine Homebrew verification | Only this macOS host was audited. | Clean host transcript with audit/install/help/version/uninstall. | No |
| Homebrew docs promotion to Available | Availability evidence is incomplete. | README/README.ko and install-doc update after full proof. | No |
| Homebrew uninstall verification | No Homebrew install was performed. | `brew uninstall ni` or equivalent cleanup transcript. | No |

## Warnings

| Warning | Evidence | Mitigation |
| --- | --- | --- |
| Local draft is not a tap formula | `packaging/homebrew-draft/ni.rb` is inside `ni-kernel`, not `Nam-Cheol/homebrew-tap`. | Keep draft clearly marked and do not document it as a user install path. |
| Local `brew audit` cannot audit a path | Homebrew 5.1.14 reports path-based `brew audit [path ...]` is disabled. | Run audit from the external tap using the formula name or supported tap workflow before publication. |

## Risks

| Risk | Impact | Follow-up |
| --- | --- | --- |
| formula may work locally but fail in published tap | Users could hit install failures after publication. | Test from the actual public tap before docs promotion. |
| source archive sha may change if wrong URL type is used | Formula checksum mismatch would break install. | Pin the final URL and recompute sha256 during publication. |
| user-facing install command may change depending on tap name | README could point to the wrong command. | Update docs only after owner confirms tap name. |
| Homebrew analytics / formula naming conflicts are unverified | `ni` may collide or need qualified install wording. | Check Homebrew formula naming before public docs update. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Planned / v0.5 candidate until full tap/formula/install proof exists. | Preserved; local draft only. | Yes | No Available claim added. |
| Release binary | Available after verified v0.5.0 assets/checksums. | Preserved. | Yes | Existing docs/117 evidence used. |
| Curl installer | Available after isolated v0.5.0 installer verification. | Preserved. | Yes | No change to installer claims. |
| Model workspace packs | Experimental unless host-level install/discovery is verified. | Preserved. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | Preserved. | Yes | No deterministic validation claim. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | No generated prompt executed. |
| Benchmark evidence | No implementation-quality, adoption, cost, latency, or downstream-quality proof. | Preserved. | Yes | `not_measured` boundary unchanged. |
| Runtime execution boundary | ni is not a task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. | Preserved. | Yes | Formula is distribution-only draft. |

## Git status / inclusion check

| Path or group | git status --short | git ls-files / tracked check | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| README.md | clean | tracked | No | No overclaim found. |
| README.ko.md | clean | tracked | No | No overclaim found. |
| docs/22_INSTALL.md | clean | tracked | No | Package manager remains Planned. |
| docs/install-curl* | clean | tracked | No | No curl wording change needed. |
| docs/51* | modified | tracked | Yes | Adds docs/118 pointer. |
| docs/117* | clean | tracked | No | Existing recommendation already points to Homebrew audit. |
| docs/118* | added | untracked until added | Yes | New audit docs. |
| packaging/homebrew-draft/* | added | untracked until added | Yes | Local draft formula only. |
| .ni/contract.json | clean | tracked protected file | No | Must remain unchanged. |
| .ni/session.json | clean | tracked protected file | No | Must remain unchanged. |
| .ni/plan.lock.json | clean | tracked protected file | No | Must remain unchanged. |
| unexpected files | none expected | n/a | No | Final `git status --short` should show only docs/51, docs/118, and packaging draft. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Initial status was clean; final status reviewed after edits. |
| `go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; no blockers, deferrals, or warnings. |
| `brew --version` | Pass | `Homebrew 5.1.14`. |
| `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git` | Pass for audit | Returned repository not found, proving the intended tap is absent. |
| `gh repo view Nam-Cheol/homebrew-tap --json nameWithOwner,url,isPrivate` | Pass for audit | Returned repository-not-found. |
| `gh release view v0.5.0 --repo Nam-Cheol/ni --json ...` | Pass | Release metadata and 6 assets returned. |
| `curl -L https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz -o /private/tmp/ni-v0.5.0-source.tar.gz` | Pass | Downloaded source archive outside the repo. |
| `shasum -a 256 /private/tmp/ni-v0.5.0-source.tar.gz` | Pass | Produced draft formula sha256. |
| `brew audit --strict --online --new-formula packaging/homebrew-draft/ni.rb` | Fail before audit | Invalid option on Homebrew 5.1.14; also sandbox cache write failed. |
| `brew audit --strict --online --new packaging/homebrew-draft/ni.rb` | Deferred | Homebrew 5.1.14 disables path-based `brew audit [path ...]`; no formula install occurred. |
| `brew developer off` | Pass | Cleaned up developer mode automatically enabled by audit attempt. |
| `ruby -c packaging/homebrew-draft/ni.rb` | Pass | `Syntax OK`. |
| `brew install --build-from-source --formula ./packaging/homebrew-draft/ni.rb` | Not run | Deferred to avoid mutating user Homebrew state. |
| `go test ./...` | Pass | All Go packages passed with `GOCACHE=/private/tmp/ni-go-cache`. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack checks passed; global install remains unverified. |
| `bash scripts/demo-check.sh` | Pass | Public demos verified without downstream runtime execution. |
| `rg -n "Homebrew is Available|..." ...` | Reviewed | Matches are negated, future-gated, or required "does not prove" wording; no touched doc claims Homebrew Available. |
| `bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `bash scripts/release-check.sh` | Pass | Release readiness gate passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` diff. |

## Changes made

- `packaging/homebrew-draft/ni.rb`: added a v0.5.0 source-build Homebrew draft.
- `docs/118_HOMEBREW_IMPLEMENTATION_AUDIT.md`: added this audit.
- `docs/118_HOMEBREW_IMPLEMENTATION_AUDIT.ko.md`: added Korean companion audit.
- `docs/51_POST_RELEASE_ROADMAP.md`: added a narrow pointer to docs/118.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: added matching Korean pointer.

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
