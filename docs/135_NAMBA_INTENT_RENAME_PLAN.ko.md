# Namba Intent Rename Plan

## Current status

State:
- v0.5.1 release: published and verified.
- current product/CLI name: ni.
- observed Windows conflict: PowerShell `ni -> New-Item`.
- existing namba-ai command: namba.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host verification: pending.
- Skills are UX; CLI is authority.
- ni is currently a pre-runtime Project Intent Compiler for AI Agents.

## Rename goal

이번 rename의 목표는 PowerShell built-in `ni` alias conflict를 피하고, 기존
namba-ai project와 이 project를 명확히 구분하는 것이다.

Future product identity:

- Product name: Namba Intent.
- CLI command: `namba-intent`.
- Legacy/internal name: `ni`, transitional only.
- Config directory: compatibility를 위해 `.ni/` 유지.
- Repository: 지금은 `Nam-Cheol/ni` 유지. Repo rename은 deferred.

Tagline:

```text
Don't run the agent yet.
Compile the intent first.
```

## Decision

RENAME_TO_NAMBA_INTENT

Justification: 현재 `ni` command에는 실제 PowerShell alias conflict가 있고,
`namba` command는 이미 namba-ai project가 사용한다. 다음 product-facing name은
Namba Intent, primary CLI command는 `namba-intent`로 가는 것이 맞다.

## Name decision

| Area | Decision | Reason | Notes |
| --- | --- | --- | --- |
| product name | Namba Intent | Namba product family를 유지하면서 pre-execution intent compiler임을 드러낸다. | v0.6.0에서 product-facing copy는 `ni`에서 Namba Intent로 이동한다. |
| CLI command | `namba-intent` | PowerShell `ni -> New-Item`과 기존 namba-ai `namba` command를 모두 피한다. | v0.6.0의 primary command로 사용한다. |
| repo name | 지금은 `Nam-Cheol/ni` 유지. | Repo rename은 release, install, URL churn을 추가한다. | Command migration이 안정된 뒤 deferred task로 다룬다. |
| config directory | `.ni/` 유지. | Existing locks, contract paths, docs, tests, user workspaces가 의존한다. | `.ni/` rename은 명시적으로 deferred. |
| old command compatibility | One transition release 동안 compatibility shim을 선호. | Current macOS/Linux users의 breakage를 줄이되 Windows는 unsafe short command를 피한다. | Shim warning: `ni is deprecated; use namba-intent`. |
| release version | v0.6.0. | Primary command 변경은 user-facing change라 tiny patch로 처리하면 안 된다. | Release notes에서 migration을 명확히 써야 한다. |

## Difference from namba-ai

| Surface | NambaAI | Namba Intent | Notes |
| --- | --- | --- | --- |
| command | `namba` | `namba-intent` | 이 project의 CLI command로 `namba`를 쓰지 않는다. |
| purpose | Codex workflow, SPEC execution, queue, sync, PR, land flows. | Pre-execution project intent compile, readiness, lock, handoff prompt. | Product family branding은 공유할 수 있지만 command scope는 분리한다. |
| execution behavior | Implementation workflows를 실행할 수 있다. | Implementation을 실행하지 않는다. | Namba Intent는 pre-runtime control layer로 남는다. |
| planning behavior | SPEC/workflow-oriented project execution planning. | Intent Lock Protocol: docs contract, readiness gate, lockfile, prompt compiler, source-of-truth rule. | Kernel is authoritative. |
| queue / PR / release automation | namba-ai workflow scope에 들어갈 수 있다. | Out of scope. | Namba Intent는 queue, PR automation, release automation, downstream execution layer가 되면 안 된다. |
| config directory | namba-ai에서는 `.namba/`. | `.ni/` compatibility 유지. | v0.6.0에서 `.ni/`를 rename하지 않는다. |
| target user | Codex/SPEC workflow operations를 실행하는 사용자. | Downstream work 전에 intent를 compile하고 lock해야 하는 사용자와 agent. | Boundary는 execution 전이다. |

## Name usage audit

Required scan:

```bash
rg -n "\bni\b|ni.exe|ni init|ni status|ni end|ni run|Nam-Cheol/ni|NambaAI|namba-ai|\bnamba\b|\.ni" README.md README.ko.md docs install.sh install.ps1 scripts cmd internal packages .agents
```

Observed surfaces:

| Surface | Current usage | Migration implication |
| --- | --- | --- |
| README and README.ko | Product name, installer URLs, Windows notes, command examples, flow examples가 `ni`를 사용한다. | Product copy는 Namba Intent로, command examples는 `namba-intent`로 바꾸되 `.ni/` compatibility note는 유지한다. |
| docs | Install, release, Homebrew, Windows alias, model workspace, readiness docs가 `ni`를 넓게 사용한다. | Current/future docs는 조심스럽게 update하고 historical release docs는 old name을 version context와 함께 유지할 수 있다. |
| install.sh | `ni`를 install하고 `Nam-Cheol/ni`, `ni --help`, `ni version`을 출력한다. | Primary installed binary를 `namba-intent`로 바꾸고 safe platform에서 deprecated `ni` shim 여부를 결정한다. |
| install.ps1 | `%LOCALAPPDATA%\ni\bin\ni.exe`를 install하고 PowerShell alias workaround를 관리한다. | `namba-intent.exe` install을 선호한다. Alias behavior 때문에 Windows `ni` shim에 의존하지 않는다. |
| scripts | Validation, release, install, demo scripts가 `ni` archive와 command names를 가정한다. | Checkers와 release scripts를 한 migration pass에서 함께 update한다. |
| cmd/internal tests | Help text, init text, protected `.ni` paths, stale-lock messages, run prompts가 `ni`를 사용한다. | Kernel semantics나 `.ni` paths를 바꾸지 않고 command/product text를 update한다. |
| packages and .agents | Skill names, examples, zip names, CLI-authority wording이 `ni`를 사용한다. | Skill pack examples를 update하고 authority text를 renamed CLI에 맞춘다. |

이 audit 결과 rename은 blind text replacement가 아니어야 한다. Historical v0.5.x
evidence는 실제 과거 behavior를 기록하므로 old command를 보존할 수 있고,
v0.6.0 user-facing instructions는 Namba Intent와 `namba-intent`를 primary로
사용해야 한다.

## Migration scope

| Surface | Change | Include in v0.6.0? | Notes |
| --- | --- | --- | --- |
| binary name | `namba-intent`를 build/distribute한다. | Yes | 모든 new docs와 install의 primary command. |
| install.sh | `namba-intent`를 install하고, safe Unix-like hosts에서는 deprecated `ni` shim을 optional로 install할 수 있다. | Yes | Uninstall은 reversible하고 installer-managed files에만 scope를 둔다. |
| install.ps1 | `namba-intent.exe`를 install한다. | Yes | Real-host transcript 전에는 Windows `ni` shim에 의존하지 않는다. |
| README | Product-facing identity와 command examples를 rename한다. | Yes | Release truth와 Homebrew boundaries를 정직하게 유지한다. |
| README.ko | Korean companion update. | Yes | `Skills are UX; CLI is authority.` 같은 exact boundary phrase를 보존한다. |
| docs | Forward-looking docs와 migration docs를 update한다. | Yes | Historical evidence docs는 past-state proof로 old name을 유지할 수 있다. |
| package skills | Skill pack examples와 authority language를 update한다. | Yes | Skill names는 별도 compatibility decision이 필요할 수 있다. |
| release assets | Archives를 `ni_...`에서 `namba-intent_...`로 rename한다. | Yes | Checksums는 new asset names와 일치해야 한다. |
| checksums | New asset names 기준으로 checksums를 생성한다. | Yes | Old checksum files를 재사용하지 않는다. |
| CI / release scripts | Artifact, command, checker expectations를 update한다. | Yes | `go test`, install checks, release checks가 new primary command를 사용해야 한다. |
| `.ni/` | 변경하지 않는다. | No | Compatibility path. Explicit deferred rename. |
| Homebrew | 이 rename task에서 formula를 publish하지 않는다. | No | Formula naming은 command migration 후 별도 평가한다. |
| repo rename | `Nam-Cheol/ni` 유지. | No | Command migration에서 URL/install churn을 피하기 위해 deferred. |

## Compatibility policy

Recommendation: B. compatibility shim.

`namba-intent`가 v0.6.0 primary command가 되어야 한다. Deprecated `ni` shim은
safe하고 maintainable한 platform에서 one transition release 동안 유지할 수 있다.
Shim은 delegation 전에 다음 warning을 출력해야 한다:

```text
ni is deprecated; use namba-intent.
```

Windows는 PowerShell이 이미 `ni -> New-Item`을 정의하므로 `ni`에 의존하면 안
된다. Windows compatibility shim을 고려하려면 user-facing docs 전에 real Windows
PowerShell host validation transcript가 필요하다.

Compatibility rules:

- New docs는 `namba-intent`를 사용한다.
- Existing `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`,
  `docs/plan/**` behavior는 compatible해야 한다.
- `ni run` behavior는 rename 후 `namba-intent run`이 되며, bounded handoff
  prompt compilation only이고 downstream execution이 아니다.
- Deprecated alias는 kernel-owned execution state가 되면 안 된다.

## Risks

| Risk | Impact | Mitigation |
| --- | --- | --- |
| breaking existing users | Current users가 `ni` script나 습관을 갖고 있을 수 있다. | Safe platform에서 one transition release warning shim과 clear release notes를 제공한다. |
| confusion with existing namba-ai | 사용자가 `namba-intent`와 `namba`를 같은 것으로 오해할 수 있다. | README, docs, release notes에 distinction table을 유지한다. |
| docs drift | 일부 docs는 historical v0.5.x behavior를, 일부 docs는 v0.6.0을 설명할 수 있다. | Historical evidence와 current instructions를 분리하고 explicit version context를 쓴다. |
| install script migration | Old binaries, PATH blocks, profile blocks가 남을 수 있다. | Fresh install, update install, uninstall, old-binary cleanup migration tests를 추가한다. |
| release asset naming | Installers/checksums가 old `ni_...` assets를 가리킬 수 있다. | GoReleaser, install scripts, release checks를 함께 update한다. |
| Windows PowerShell behavior | Alias가 남으면 `ni`가 `New-Item`을 실행한다. | Windows primary는 `namba-intent.exe`로 하고, Windows `ni` compatibility는 real-host proof 전까지 deferred한다. |
| Homebrew formula naming | Homebrew core에 unrelated `ni` formula가 있다. | Homebrew는 deferred하고 `namba-intent` formula naming을 별도 평가한다. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| namba-ai distinction | Namba Intent에 `namba`를 쓰면 안 된다. | 이 plan은 `namba`를 namba-ai에 남기고 `namba-intent`를 선택한다. | Yes | CLI collision을 피한다. |
| Namba Intent identity | Product rename plan only; implementation deferred. | 이 문서는 future identity를 결정하지만 code rename은 하지 않는다. | Yes | v0.6.0 implementation은 later task. |
| `run` behavior | Prompt compilation only. | Bounded handoff prompt compilation으로 유지한다. | Yes | Downstream execution을 추가하지 않았다. |
| Homebrew | Planned / v0.5 candidate only. | 유지. | Yes | Formula publication 또는 Available claim 없음. |
| Windows verification | Transcript 전까지 pending. | 유지. | Yes | Static installer checks는 real-host proof가 아니다. |
| runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, or downstream execution layer. | 유지. | Yes | Rename은 kernel scope를 바꾸지 않는다. |

## Recommended next task

A. implement Namba Intent rename

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: Implement the v0.6.0 Namba Intent rename.

Use docs/135_NAMBA_INTENT_RENAME_PLAN.md and
docs/135_NAMBA_INTENT_RENAME_PLAN.ko.md as the authoritative rename plan.

Decision:
- Product name: Namba Intent.
- Primary CLI command: namba-intent.
- Do not use namba as this project's CLI command.
- Keep .ni/ for compatibility.
- Keep repository Nam-Cheol/ni for now.
- Prefer a compatibility shim for one transition release only where safe.
- Windows should use namba-intent.exe as the primary command and must not rely
  on ni unless real-host PowerShell behavior is proven safe.

Scope:
- Rename product-facing current instructions from ni to Namba Intent where they
  describe the current v0.6.0+ product.
- Rename primary command examples from ni to namba-intent.
- Update Go command metadata, help text, tests, installers, release asset names,
  checksums expectations, docs, README, README.ko, scripts, package skills, and
  .agents examples needed for the new primary command.
- Preserve historical v0.5.x evidence docs where they describe past verified
  ni behavior, but add version context if needed.
- Preserve .ni/contract.json, .ni/session.json, .ni/plan.lock.json, schema
  paths, lockfile path, and docs/plan/** behavior.

Compatibility policy:
- namba-intent is primary.
- A deprecated ni shim may be kept for a transition release on platforms where
  it is safe and maintainable.
- The shim must warn: ni is deprecated; use namba-intent.
- Do not make Windows ni compatibility a claim without a real Windows
  PowerShell transcript.

Forbidden:
- Do not rename .ni/.
- Do not rename the repository.
- Do not publish, tag, create a GitHub release, upload assets, run release
  workflows, create or publish a Homebrew formula, or mark Homebrew Available.
- Do not run ni end on the project root.
- Do not relock the project root.
- Do not execute generated prompts.
- Do not add runtime execution behavior.
- Do not make run execute downstream work.
- Do not add task runner, SPEC runner, queue, PR automation, release automation,
  shell adapter, or Codex exec adapter behavior.

Validation:
- git status --short
- gofmt -w .
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- primary rename behavior
- compatibility shim behavior
- validation results
- protected .ni diff result
- confirmation that no publication, tag, release, upload, root relock, prompt
  execution, Homebrew publication, or downstream execution behavior occurred
```
