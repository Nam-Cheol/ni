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

이 pass는 tap publication plan 전에 local Homebrew draft formula
`packaging/homebrew-draft/ni.rb`를 로컬에서 검증할 수 있는지 확인한다.
Non-mutating check를 우선하고, pre-existing Homebrew/PATH state를 기록하며,
가장 안전한 Homebrew check를 시도한 뒤 tap 생성이나 user-facing install claim
변경 전에서 멈춘다.

## Decision

Decision: HOMEBREW_FORMULA_LOCAL_BLOCKED.

Justification: draft formula는 valid Ruby syntax이고 local machine에는 Homebrew
5.1.14가 있지만, 이 Homebrew version은 draft formula의 path-based audit와
path-based install을 모두 거부한다. `brew install --build-from-source --formula
./packaging/homebrew-draft/ni.rb`는 아무것도 설치하지 않았고, Homebrew는
formula가 tap 안에 있어야 한다고 보고했다. Install이 발생하지 않았으므로
installed `ni --help`, installed `ni version`, formula `brew test`, uninstall
cleanup은 이 task에서 verify할 수 없었다.

Homebrew remains Planned / v0.5 candidate.

## Draft formula inspection

| Item | Observed state | Pass? | Notes |
| --- | --- | --- | --- |
| formula path | `packaging/homebrew-draft/ni.rb` | Yes | Draft-only; tap formula 아님. |
| class name | `Ni < Formula` | Yes | Intended formula name과 맞지만 Homebrew core에 다른 `ni` formula가 이미 있다. |
| desc | `Project Intent Compiler for AI Agents` | Yes | Current product definition과 일치. |
| homepage | `https://github.com/Nam-Cheol/ni` | Yes | Project homepage. |
| url | `https://github.com/Nam-Cheol/ni/archive/refs/tags/v0.5.0.tar.gz` | Yes | docs/118에 기록된 source archive와 동일. |
| sha256 | `67a694ff9e9e076b2cfc731c96575604e18abea03b1bb1f818e95b9aee54bb02` | Yes | docs/118에서 verified. |
| license | `MIT` | Yes | Present. |
| dependencies | `depends_on "go" => :build` | Yes | Source-build formula. |
| install block | `go build -trimpath ... -o bin/"ni" ./cmd/ni` | Yes | `#{version}`을 `ni/internal/version.Version`에 inject한다. |
| test block | `ni --help` and `ni version` | Partial | Intended check는 적절하지만 local install blocked 때문에 실행되지 않았다. |
| version inference | tag URL에서 `0.5.0` inferred | Partial | Formula URL 기준 expected이지만 installed binary output은 verify하지 못했다. |

## Pre-existing local state

| Check | Result | Notes |
| --- | --- | --- |
| `brew --version` | Pass: `Homebrew 5.1.14` | Local Homebrew available. |
| `brew developer` | Pass: developer mode disabled | Audit attempt가 잠깐 developer mode를 켰고 이후 껐다. |
| `brew list --formula ni` | Not installed | `Error: No such keg: /opt/homebrew/Cellar/ni`. |
| `brew info ni` | Pass with warning context | Homebrew core에는 다른 `ni` formula가 있다: `stable 30.1.0`, "Selects the right Node package manager based on lockfiles", not installed. |
| `which ni` | Not found | Pre-existing PATH `ni` binary 없음. |
| pre-existing `ni version` | Not runnable | `zsh:1: command not found: ni`. |
| `brew --prefix` | `/opt/homebrew` | Local prefix 확인에만 사용. |

## Formula checks

| Command | Result | Mutation risk | Notes |
| --- | --- | --- | --- |
| `ruby -c packaging/homebrew-draft/ni.rb` | Pass: `Syntax OK` | None | Syntax-only check. |
| `brew audit --strict --new --online --formula ./packaging/homebrew-draft/ni.rb` | Blocked | Low | Homebrew 5.1.14 reports `Calling brew audit [path ...] is disabled! Use brew audit [name ...] instead.` |
| `brew audit --strict --new --online ./packaging/homebrew-draft/ni.rb` | Blocked | Low | Same path-audit rejection. |
| `brew developer off` | Pass | None | Audit attempt 뒤 developer mode cleanup. |
| `brew install --build-from-source --formula ./packaging/homebrew-draft/ni.rb` | Blocked before install | Medium if allowed | Homebrew rejected the path formula: `Homebrew requires formulae to be in a tap`. |
| installed `ni --help` | Not run | n/a | Installed formula binary 없음. |
| installed `ni version` | Not run | n/a | Installed formula binary 없음. |
| `brew test ni` | Not run | n/a | Formula가 installed되지 않음. |
| `brew uninstall ni` | Not run | n/a | 이 task가 설치한 것이 없음. |

## Cleanup and environment preservation

| Check | Result | Notes |
| --- | --- | --- |
| post-attempt `brew list --formula ni` | Not installed | Failed install이 `ni` keg를 남기지 않았음을 확인. |
| post-attempt `which ni` | Not found | PATH가 Homebrew-installed `ni`를 가리키도록 바뀌지 않음. |
| post-attempt `brew developer` | Developer mode disabled | Environment를 non-developer mode로 되돌림. |

User-installed `ni`를 overwrite하지 않았다. Homebrew `ni`를 uninstall하지 않았다.
Tap repository를 create하거나 push하지 않았다.

## Blockers

| Blocker | Evidence | Required future fix |
| --- | --- | --- |
| Path-based audit is disabled | Homebrew 5.1.14 rejects `brew audit [path ...]`. | Real tap 또는 Homebrew-supported named formula workflow에서 audit한다. |
| Path-based install is rejected | Homebrew says formulae must be in a tap. | Owner-approved tap을 만들거나 사용한 뒤 그 tap에서 install한다. |
| Name collision risk | Homebrew core already has a different formula named `ni`. | User-facing docs promotion 전에 fully qualified tap install wording과 formula naming을 재검토한다. |

## What this proves

- Local draft formula는 valid Ruby syntax다.
- Local Homebrew environment는 available이다.
- Pre-existing Homebrew-installed `ni` 또는 PATH `ni`는 없었다.
- Current local Homebrew는 이 draft를 path로 validate하지 못한다.
- Failed install attempt는 Homebrew `ni` keg를 남기지 않았다.

## What this does not prove

- Homebrew is Available.
- Public tap exists.
- Formula can be audited from the intended tap.
- Formula can be installed from the intended tap.
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
| Runtime execution boundary | `ni`는 task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, execution evidence loop가 아니다. | Preserved. | Yes | Runtime behavior 추가 없음. |

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
