# Homebrew Formula Local Verification

## Current status

- v0.5.0 publication: verified
- Post-release verification decision: V0_5_0_POST_RELEASE_VERIFIED_WITH_NOTES
- Homebrew implementation audit decision: HOMEBREW_IMPLEMENTATION_READY_WITH_DEFERRALS
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- This verification did not publish a Homebrew formula or mark Homebrew Available.

## Verification goal

This pass checks whether the local Homebrew draft formula at
`packaging/homebrew-draft/ni.rb` can be validated locally before any tap
publication plan. It prioritizes non-mutating checks, records the pre-existing
Homebrew and PATH state, attempts the safest available Homebrew checks, and
stops before creating a tap or changing user-facing install claims.

## Decision

Decision: HOMEBREW_FORMULA_LOCAL_BLOCKED.

Justification: the draft formula has valid Ruby syntax and the local machine
has Homebrew 5.1.14, but this Homebrew version rejects both path-based audit and
path-based install for the draft formula. `brew install --build-from-source
--formula ./packaging/homebrew-draft/ni.rb` did not install anything; Homebrew
requires formulae to live in a tap. Because no install occurred, installed
`ni --help`, installed `ni version`, formula `brew test`, and uninstall cleanup
could not be verified in this task.

Homebrew remains Planned / v0.5 candidate.

## Draft formula inspection

| Item | Observed state | Pass? | Notes |
| --- | --- | --- | --- |
| formula path | `packaging/homebrew-draft/ni.rb` | Yes | Draft-only; not a tap formula. |
| class name | `Ni < Formula` | Yes | Matches intended formula name, but Homebrew core already has a different `ni` formula. |
| desc | `Project Intent Compiler for AI Agents` | Yes | Matches the current product definition. |
| homepage | `https://github.com/Nam-Cheol/ni` | Yes | Project homepage. |
| url | `https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz` | Yes | Same source archive recorded in docs/118. |
| sha256 | `67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02` | Yes | Previously verified in docs/118. |
| license | `MIT` | Yes | Present. |
| dependencies | `depends_on "go" => :build` | Yes | Source-build formula. |
| install block | `go build -trimpath ... -o bin/"ni" ./cmd/ni` | Yes | Injects `#{version}` into `ni/internal/version.Version`. |
| test block | `ni --help` and `ni version` | Partial | Good intended checks, but not executed because local install was blocked. |
| version inference | `0.5.0` inferred from tag URL | Partial | Expected by formula URL, but installed binary output was not verified. |

## Pre-existing local state

| Check | Result | Notes |
| --- | --- | --- |
| `brew --version` | Pass: `Homebrew 5.1.14` | Homebrew is available locally. |
| `brew developer` | Pass: developer mode disabled | The audit attempt briefly enabled developer mode; it was turned off. |
| `brew list --formula ni` | Not installed | `Error: No such keg: /opt/homebrew/Cellar/ni`. |
| `brew info ni` | Pass with warning context | Homebrew core has a different `ni` formula: `stable 30.1.0`, "Selects the right Node package manager based on lockfiles", not installed. |
| `which ni` | Not found | No pre-existing PATH `ni` binary was found. |
| pre-existing `ni version` | Not runnable | `zsh:1: command not found: ni`. |
| `brew --prefix` | `/opt/homebrew` | Used only to confirm local prefix. |

## Formula checks

| Command | Result | Mutation risk | Notes |
| --- | --- | --- | --- |
| `ruby -c packaging/homebrew-draft/ni.rb` | Pass: `Syntax OK` | None | Syntax-only check. |
| `brew audit --strict --new --online --formula ./packaging/homebrew-draft/ni.rb` | Blocked | Low | Homebrew 5.1.14 reports `Calling brew audit [path ...] is disabled! Use brew audit [name ...] instead.` |
| `brew audit --strict --new --online ./packaging/homebrew-draft/ni.rb` | Blocked | Low | Same path-audit rejection. |
| `brew developer off` | Pass | None | Cleaned up developer mode after audit attempt. |
| `brew install --build-from-source --formula ./packaging/homebrew-draft/ni.rb` | Blocked before install | Medium if allowed | Homebrew rejected the path formula: `Homebrew requires formulae to be in a tap`. |
| installed `ni --help` | Not run | n/a | No installed formula binary existed. |
| installed `ni version` | Not run | n/a | No installed formula binary existed. |
| `brew test ni` | Not run | n/a | Formula was not installed. |
| `brew uninstall ni` | Not run | n/a | Nothing was installed by this task. |

## Cleanup and environment preservation

| Check | Result | Notes |
| --- | --- | --- |
| post-attempt `brew list --formula ni` | Not installed | Confirms the failed install did not leave a `ni` keg. |
| post-attempt `which ni` | Not found | PATH was not changed to point at a Homebrew-installed `ni`. |
| post-attempt `brew developer` | Developer mode disabled | The environment was returned to non-developer mode. |

No user-installed `ni` was overwritten. No Homebrew `ni` was uninstalled. No tap
repository was created or pushed.

## Blockers

| Blocker | Evidence | Required future fix |
| --- | --- | --- |
| Path-based audit is disabled | Homebrew 5.1.14 rejects `brew audit [path ...]`. | Audit from a real tap or another Homebrew-supported named formula workflow. |
| Path-based install is rejected | Homebrew says formulae must be in a tap. | Create or use an owner-approved tap, then install from that tap. |
| Name collision risk | Homebrew core already has a different formula named `ni`. | Use fully qualified tap install wording and reassess formula naming before user-facing docs promotion. |

## What this proves

- The local draft formula has valid Ruby syntax.
- The local Homebrew environment is available.
- No pre-existing Homebrew-installed `ni` or PATH `ni` was present.
- The current local Homebrew does not support validating this draft by path.
- The failed install attempt did not leave a Homebrew `ni` keg.

## What this does not prove

- Homebrew is Available.
- A public tap exists.
- The formula can be audited from the intended tap.
- The formula can be installed from the intended tap.
- Installed `ni --help` works.
- Installed `ni version` reports `0.5.0`.
- Formula `brew test` passes.
- Uninstall cleanup works after a successful install.
- Cross-machine Homebrew behavior.

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Planned / v0.5 candidate until full proof exists. | Preserved. | Yes | No Available claim added. |
| Release binary | Available after verified v0.5.0 assets/checksums. | Preserved. | Yes | No change. |
| Curl installer | Available after v0.5.0 installer verification. | Preserved. | Yes | No change. |
| Model workspace packs | Experimental unless host-level install/discovery is verified. | Preserved. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | Preserved. | Yes | No deterministic validation claim. |
| ni run | Bounded prompt compilation only. | Preserved. | Yes | No generated prompt executed. |
| Runtime execution boundary | `ni` is not a task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. | Preserved. | Yes | No runtime behavior added. |

## Recommended next task

Selected next task: owner-approved tap-based Homebrew validation.

Next executable prompt:

```text
Proceed with owner-approved tap-based Homebrew validation for ni.

This task may not mark Homebrew Available until all evidence is complete. Do not
publish, push, or create a public tap unless the user explicitly approves that
publication action.

Goal:
Validate the reviewed `packaging/homebrew-draft/ni.rb` through a Homebrew tap
workflow because Homebrew 5.1.14 rejects local path-based audit and install.

Required evidence before any Available claim:
- owner-approved tap path exists
- formula name and possible conflict with Homebrew core `ni` are resolved
- formula URL is final
- formula sha256 is recomputed and verified
- `brew audit` passes through a supported named-formula or tap workflow
- `brew install` works from the intended formula path
- installed `ni --help` works
- installed `ni version` reports `0.5.0`
- formula `brew test` passes
- uninstall cleanup works
- README.md, README.ko.md, and docs/22_INSTALL.md are updated only after the proof exists
```
