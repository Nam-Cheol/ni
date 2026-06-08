# Namba Intent Rename Implementation

## Current Status

State:
- v0.5.1 release: published and verified.
- rename decision: RENAME_TO_NAMBA_INTENT.
- product name: Namba Intent.
- primary command: `namba-intent`.
- repository: `Nam-Cheol/ni` retained.
- config directory: `.ni/` retained.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host verification: pending.
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

이 문서는 v0.6.0 publish 없이 current tree에서 rename surface를 구현한 결과를
기록합니다.

## Decision Summary

| Area | Decision | Implemented? | Notes |
| --- | --- | --- | --- |
| product name | Namba Intent | Yes | README, CLI help, current-tree docs가 새 identity를 사용합니다. |
| primary command | `namba-intent` | Yes | `cmd/namba-intent`가 shared CLI implementation을 호출합니다. |
| legacy `ni` shim | Safe한 곳에서 one transition release 동안 유지. | Yes | `cmd/ni`는 warning 후 delegate합니다. |
| repository name | `Nam-Cheol/ni` 유지. | Yes | Repo rename은 deferred. |
| config directory | `.ni/` 유지. | Yes | Contract, session, lock, docs/plan paths unchanged. |
| release version | Future v0.6.0 release. | Not released | Tag, publish, release workflow, asset upload 없음. |
| Homebrew | Planned / v0.5 candidate. | Deferred | Formula 생성 또는 Available claim 없음. |
| Windows | `namba-intent.exe` primary. | Installer에 구현. | Real-host verification은 pending. |

## Command Migration

| Old command | New command | Compatibility behavior | Notes |
| --- | --- | --- | --- |
| `ni init .` | `namba-intent init .` | `ni` shim이 warning 후 delegate. | `.ni/` workspace directory 유지. |
| `ni status --proof --next-questions` | `namba-intent status --proof --next-questions` | `ni` shim이 warning 후 delegate. | CLI remains readiness authority. |
| `ni end` | `namba-intent end` | `ni` shim이 warning 후 delegate. | 명시 승인 없이 project root에서 실행하지 않습니다. |
| `ni run --max-chars 4000` | `namba-intent run --max-chars 4000` | `ni` shim이 warning 후 delegate. | Prompt compilation only. |
| `ni version` | `namba-intent version` | `ni` shim이 warning 후 delegate. | Source build는 `0.0.0-dev`일 수 있습니다. |

## Difference From namba-ai

| Surface | NambaAI | Namba Intent | Notes |
| --- | --- | --- | --- |
| command | `namba` | `namba-intent` | 이 project는 `namba` command를 claim하지 않습니다. |
| purpose | Codex workflow, SPEC execution, queue, sync, PR, land. | Pre-execution intent compile, readiness, lock, handoff prompt. | Product family는 같아도 behavior는 다릅니다. |
| execution behavior | Implementation workflow를 실행할 수 있음. | Downstream work를 실행하지 않음. | Runtime 전에서 멈춥니다. |
| planning behavior | SPEC/workflow execution planning. | Intent Lock Protocol planning contract. | Kernel is authoritative. |
| queue / PR / release automation | namba-ai workflows scope. | Out of scope. | 여기서 추가하지 않았습니다. |
| config directory | `.namba/` | `.ni/` | Compatibility 유지. |
| target user | Workflow operators. | Handoff 전 intent를 compile하는 users/agents. | Boundary는 execution 전입니다. |

## Implementation Surfaces

| Surface | Change made | Notes |
| --- | --- | --- |
| `cmd/namba-intent` | Primary CLI entrypoint 추가. | `internal/cli.Run` 호출. |
| `cmd/ni` | Deprecated shim으로 전환. | stderr warning 후 delegate. |
| `internal/cli` | Shared CLI implementation으로 이동. | Help/init copy는 Namba Intent 기준. |
| `install.sh` | `namba-intent` install 및 verification. | Archive prefix `namba-intent_`. |
| `install.ps1` | `%LOCALAPPDATA%\namba-intent\bin`에 `namba-intent.exe` install. | Primary path에 `ni` alias block 불필요. |
| `README` / `README.ko` | Current-tree product copy updated. | v0.5.1 public release distinction 유지. |
| scripts | Install/release checks updated. | Broad validation passed. |
| skills / `.agents` | Skill IDs는 transition 동안 `ni-*` 유지, command examples는 `namba-intent` 사용. | Skills are UX; CLI is authority. |

## Compatibility Policy

`namba-intent`가 primary입니다. `ni` command는 safe하고 maintainable한 곳에서
one transition release 동안 deprecated shim으로 남을 수 있습니다.

```text
ni is deprecated; use namba-intent.
```

Windows primary install은 PowerShell `ni -> New-Item` alias 때문에 `ni`에
의존하지 않습니다. Legacy v0.5.x alias cleanup은 historical guidance이며
`namba-intent.exe`에는 필요하지 않습니다.

## Historical Docs Policy

Historical v0.5.x verification docs는 실제 과거 command와 public-release
evidence를 설명할 때 `ni`를 유지할 수 있습니다. Current user-facing
instructions는 Namba Intent와 `namba-intent`를 사용하고, latest published
v0.5.1과 current main이 다르면 explicit note를 둡니다.

## Claim-Boundary Audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| NambaAI distinction | 이 project에 `namba`를 쓰지 않음. | Primary command는 `namba-intent`. | Yes | namba-ai가 `namba` 유지. |
| Namba Intent identity | Current tree rename only. | Help/README가 Namba Intent 사용. | Yes | v0.6.0 published 아님. |
| legacy `ni` shim | Warning and delegate only. | stderr warning 후 delegate. | Yes | Windows primary path 아님. |
| `.ni/` compatibility | State directory rename 금지. | Protected `.ni` paths unchanged. | Yes | Root protected files 미수정. |
| Homebrew | Planned / v0.5 candidate only. | Available claim 없음. | Yes | Formula publication 없음. |
| Windows verification | Transcript 전까지 pending. | Static checks only. | Yes | Real-host claim 없음. |
| run behavior | Prompt compilation only. | `run`은 bounded prompt compilation 유지. | Yes | Agent/shell execution 없음. |

## Validation Results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Passed | 작업 전 dirty tree 확인. |
| `git tag --list v0.6.0` | Passed | `v0.6.0` tag 없음. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed | Protected root `.ni` diff 없음. |
| `gofmt -w cmd internal` | Passed | Go format 완료. |
| `go test ./...` | Environment failed | 기본 Go cache가 sandbox 밖이라 실패. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed | 모든 Go package 통과. |
| `python3 scripts/check-install-docs.py` | Passed | Install doc marker 통과. |
| `python3 scripts/check-install-ps1.py` | Passed | Windows installer static safety check 통과. |
| `bash scripts/check-skill-packs.sh` | Passed | Package와 repo-local skill packs 통과. |
| `bash scripts/test-install-sh.sh` | Passed | Unix installer dry-run/checksum/install/uninstall 통과. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed | Source, build, temp install, command-name check 통과. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed | Current-tree `namba-intent` CLI smoke 통과. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/demo-check.sh` | Passed | Demo prompts/exports seed-only boundary 통과. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed | Broad quality wrapper 통과. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed | Check-only release readiness gate 통과; release action 없음. |
| Fresh-shell temp binary smoke | Passed | `/private/tmp`에서 `namba-intent --help`, `version`, `init . --yes`, `status --proof --next-questions`, repeated init, lockfile safety, `ni` shim delegation 통과. |

## What This Implementation Proves

State only:
- current tree implements Namba Intent rename surfaces;
- `namba-intent` command works in current tree if validated;
- `.ni/` compatibility remains intact;
- no release action was performed.

## What This Implementation Does Not Prove

State:
- v0.6.0 has been published;
- public install retrieves `namba-intent`;
- Windows real-host execution works;
- Homebrew is Available;
- external users succeed;
- downstream execution succeeds.

## Recommended Next Task

A. Namba Intent first-user smoke.
