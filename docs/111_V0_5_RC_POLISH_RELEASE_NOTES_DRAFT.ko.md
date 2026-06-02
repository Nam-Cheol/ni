# v0.5 RC Polish / Release Notes Draft

## Current status

- RC decision: RC_READY_WITH_DEFERRALS
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- This document is a draft and does not publish, tag, or release v0.5.

## Draft release notes

### Summary

v0.5는 AI Agents를 위한 Project Intent Compiler인 `ni`의 release-candidate
draft다. Downstream agent work가 시작되기 전에 conversation proof,
change control, stale-lock diagnostics, assisted no-terminal workflows, model
workspace boundaries, benchmark claim boundaries를 더 단단하게 만든다.

이 draft는 v0.5가 release되었다고 claim하지 않는다. `ni run`은 여전히 current
locked plan에서 bounded handoff prompt를 compile하며 downstream work를 execute하지
않는다.

### Highlights

- Project Intent Compiler positioning을 유지한다: `ni`는 downstream agent work
  전에 intent를 compile한다.
- Proof capture reliability work는 planning evidence가 무엇을 증명하고 무엇을
  증명하지 않는지 clarify한다.
- Change-control UX audit은 stale-lock과 amended-intent safety model을 기록한다.
- `LOCK-STALE` diagnostic은 existing lock이 current planning inputs와 더 이상
  match하지 않는다는 뜻을 문서화한다.
- Amend/relock workflow examples는 model을 authority로 만들지 않고 recovery
  order를 설명한다.
- No-terminal stale-lock examples와 transcript checklist는 model-only drafts,
  pasted CLI output, trusted runner transcripts, fixture evidence,
  target-workspace evidence를 구분한다.
- Changed-intent fixture coverage는 fixture relock과 project-root relock을
  분리하면서 stale-lock coverage를 넓힌다.
- Model workspace wording verification은 Experimental status와 skill boundaries를
  보존한다.
- Benchmark claim-boundary hardening은 `not_measured` limits를 visible하게
  유지한다.
- Task 190은 RC readiness decision을 기록했다:
  `RC_READY_WITH_DEFERRALS`.

### Reliability improvements

`docs/101` through `docs/109`의 v0.5 reliability set은 proof capture,
change control, stale-lock diagnostics, amend/relock examples, model workspace
wording, no-terminal examples, transcript quality, changed-intent fixtures,
release-readiness sweep coverage를 연결된 기록으로 추가한다.

이 docs는 core protocol을 conservative하게 유지한다:

```text
conversation -> project contract -> readiness gate -> lock hash -> downstream handoff
```

Reliability work는 execution harness, task runner, shell adapter, Codex exec
adapter, queue, PR automation, release automation, execution evidence loop를
추가하지 않는다.

### Stale-lock and change-control

v0.5는 locked plan 주변의 changed-intent path를 개선한다:

- `LOCK-STALE` means the existing lock no longer matches current planning inputs.
- `ni status`는 project가 otherwise ready라도 existing lock이 stale이면 warning을
  보여줄 수 있다.
- `ni end`는 changed intent review 뒤 lock 또는 relock을 수행하는
  CLI-authoritative step으로 남는다.
- `ni run`은 stale handoff를 refuse하고 recovery order로 되돌린다.
- Fixture coverage는 representative changed-intent cases, non-lockable false
  positives, fixture relock recovery, project-root safety를 check한다.
- Amend/relock examples는 readiness, mitigation, risk, non-goal criteria를
  약화하지 않고 user path를 보여준다.

Recovery order는 유지된다:

```text
review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000
```

### Proof capture and acceptance evidence

Proof capture와 acceptance evidence work는 planning proof가 audit trail이지
implementation proof가 아님을 clarify한다. `ni status`는 readiness authority,
`ni end`는 locking authority, `ni run`은 current lock에서 bounded prompt
compilation authority로 남는다.

v0.5 acceptance evidence matrix는 각 v0.5 lane에 필요한 files, commands, status
vocabulary, `not_measured` boundaries를 정의한다. Future maintainers가 draft,
fixture, benchmark evidence를 product readiness claim으로 upgrade하지 않도록 돕는다.

### Benchmark claim boundaries

Benchmark evidence는 bounded planning-artifact claims만 support한다. Current
benchmark surfaces는 vague request가 auditable planning artifact, `READY` isolated
benchmark workspace, bounded handoff prompt로 바뀔 수 있음을 보여줄 수 있다.

그 evidence는 implementation correctness, downstream agent performance,
rework reduction, adoption improvement, cost improvement, latency improvement,
statistical effect, dashboard product readiness, real research approval,
fieldwork authorization, research quality, intervention effectiveness를 증명하지
않는다. Relevant한 곳에서는 이 항목들이 `not_measured`로 남는다.

### Model workspace and skills

- Model workspace packs remain Experimental.
- Skills are UX; CLI is authority.
- Skills may draft or explain planning text.
- Skills may help explain `LOCK-STALE`.
- Skills do not determine readiness.
- Skills do not lock or relock.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- Skills do not update `.ni/plan.lock.json`.
- Host-level/global install, provider runtime behavior, cross-machine install은
  later host-specific proof가 있기 전까지 unverified로 남는다.

### No-terminal assisted workflow

- No-terminal method remains Experimental / assisted.
- Model-only draft is draft-only.
- Readiness, lock freshness, relock, hash verification, bounded handoff
  compilation을 claim하려면 trusted runner의 exact CLI output이 필요하다.
- Fixture transcript는 fixture claims만 support한다.
- Target-workspace claim은 exact target-workspace command output을 요구한다.
- No-terminal assistance itself는 deterministic validation을 제공하지 않는다.

### Validation and tests

v0.5 RC readiness source는 CLI readiness, Go tests, install claim checks,
skill-pack checks, demo checks, smoke checks, quality checks, install checks,
release checks, protected-file diff checks가 passing이었다고 기록한다.

Relevant fixture coverage에는 stale-lock warnings, stale `ni run` refusal,
fixture relock recovery, changed-intent lockable input coverage, non-lockable
false positives, benchmark claim-boundary checks, no-terminal transcript
boundaries, ni-grill docs-only boundaries, seed-only export boundaries가 포함된다.

Validation scripts는 temporary fixture workspaces에서 `ni end`를 exercise할 수
있다. Those fixture runs must not be described as project-root relock.

### Known deferrals

| Deferral | Status | Why deferred | Required future evidence |
| --- | --- | --- | --- |
| Homebrew implementation / availability | Planned / v0.5 candidate | Tap/formula availability를 claim하지 않는다. | Tap, formula, sha256, audit, `brew install` output, `ni --help`, `ni version` verification. |
| Model workspace host verification | Experimental | Host-level/global install, provider runtime behavior, cross-machine install이 verified되지 않았다. | Host-specific install 또는 discovery proof와 provider behavior transcript. |
| External user validation | Limited | RC docs는 boundaries를 보존하지만 external adoption 또는 user success를 증명하지 않는다. | User-run transcripts, maintained external validation notes, 또는 scoped user research. |
| Additional benchmark breadth | Bounded | Current benchmark evidence는 qualitative and planning-artifact scoped 상태다. | `not_measured` boundaries를 보존하는 additional benchmark case 또는 broader report. |
| No-terminal deterministic validation not claimed | Experimental / assisted | Model-only와 pasted-output workflows는 exact trusted CLI output 없이는 CLI state를 prove하지 못한다. | Target workspace의 `ni status`, `ni end`, hash verification, `ni run` trusted runner transcripts. |
| v0.5 release artifacts | Draft only | 이 task는 tag, publish, upload assets, GitHub release creation을 하지 않는다. | Release artifact dry-run, final release preflight, authorized maintainer의 actual tag/release action. |

### What v0.5 still does not do

- does not execute downstream work
- does not run agents
- does not provide a task runner
- does not provide a SPEC runner
- does not provide a shell adapter
- does not provide a Codex exec adapter
- does not provide queues
- does not provide PR automation
- does not provide release automation
- does not provide an execution evidence loop
- does not make Homebrew Available
- does not make model workspace packs Available
- does not make no-terminal deterministic
- does not prove implementation correctness or downstream execution quality

### Install / upgrade notes

이 section은 actual v0.5 release artifact verification 전의 draft다. Final release
path가 verified되기 전에는 v0.5 availability claim으로 final release notes에 옮기면
안 된다.

Currently true install paths:

- Source: Go가 있으면 this checkout에서 Available.
- Local binary: `make build`와 `./bin/ni`로 this checkout에서 Available.
- Release binary: verified v0.4.0 release assets에 대해 Available.
- Curl installer: verified v0.4.0 release assets에 대해 Available.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Homebrew: Planned / v0.5 candidate.

이 draft는 downloadable v0.5 artifacts가 있다고 claim하지 않는다.

### Maintainer validation checklist

Final release-note promotion 또는 release action 전에 실행한다:

```bash
GOCACHE=/private/tmp/ni-go-cache go test ./...
GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
python3 scripts/check-install-docs.py
bash scripts/check-skill-packs.sh
bash scripts/demo-check.sh
GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh
GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh
GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
```

### Release-note claim audit

| Claim area | Allowed wording | Forbidden wording | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Homebrew: Planned / v0.5 candidate. | Homebrew Available; package-manager install works. | Yes | Availability는 tap, formula, sha256, `brew install`, `ni --help`, `ni version` proof를 요구한다. |
| Model workspace packs | Model workspace packs: Experimental. | Model workspace packs Available; global install verified; provider runtime behavior verified. | Yes | Host-level/global behavior remains unverified. |
| No-terminal | No-terminal method: Experimental / assisted. | No-terminal deterministic validation; model-only draft as CLI state. | Yes | State claims에는 exact trusted CLI output이 필요하다. |
| ni run | Bounded handoff prompt만 compile한다. | Downstream work, shell commands, Codex, agents, queues를 execute한다. | Yes | Prompt compilation은 pre-runtime으로 남는다. |
| READY | Planning contract readiness only. | Product readiness, implementation correctness, downstream success. | Yes | Scope는 planning artifact로 유지된다. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | Product failure, implementation failure, benchmark failure. | Yes | Stale은 current planning `READY`와 coexist할 수 있다. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | Causal impact, adoption, cost, latency, dashboard readiness, research approval, fieldwork authorization. | Yes | Unsupported benchmark claims remain `not_measured`. |
| Fixture relock | Fixture relock is separate from project-root relock. | Validation fixture relock updated project-root `.ni/plan.lock.json`. | Yes | Fixture runs support fixture claims only. |
| Trusted runner transcript | Exact workspace, command, output, time에 scoped된 claims. | Exact output 없는 global lock freshness 또는 target-workspace state. | Yes | No-terminal claims need exact CLI output. |
| Runtime execution boundary | `ni` is not an execution harness. | Task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, execution evidence loop. | Yes | Kernel은 planning contracts, readiness, locks, prompt compilation의 authority로 남는다. |
| RC decision | `RC_READY_WITH_DEFERRALS`. | v0.5 released; RC has no deferrals. | Yes | Task 190 decision을 conservative하게 보존한다. |

## Changes made

- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md`: English RC polish and
  release-note draft 추가.
- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md`: Korean companion 추가.
- `docs/51_POST_RELEASE_ROADMAP.md`: 이 draft에 대한 narrow pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer 추가.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`: 이 draft에 대한 narrow
  follow-up pointer 추가.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`: matching Korean
  follow-up pointer 추가.

## What this draft proves

- release-note wording is aligned with audited v0.5 boundaries
- RC_READY_WITH_DEFERRALS is represented conservatively
- known deferrals are explicit
- no release action was performed

## What this draft does not prove

- v0.5 has been published
- artifacts were uploaded
- Homebrew is Available
- model workspace host behavior is verified
- no-terminal is deterministic
- downstream execution succeeds
- external users succeed
- benchmark effect size or causal impact

## Recommended next task

Selected next task: E. v0.5 release notes final preflight.

Why: release-note wording은 이제 drafted and claim-audited 상태지만, 어떤 release
action도 고려하기 전에 draft, validation commands, protected-file safety,
release/non-release boundary를 final preflight로 확인해야 한다.

Follow-up preflight: `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`는 그
final check를 기록하고 draft-only, no-release boundary를 보존한다.

## Next task prompt

```text
Proceed with v0.5 release notes final preflight in /Users/namba/Documents/project/ni.

This is a documentation and validation preflight task only. Do not publish, tag, create a GitHub release, upload assets, run goreleaser publish, run ni end on the project root, relock the project root, execute generated prompts, create Homebrew formula files, change Go implementation, add release automation, add runtime execution behavior, or mark v0.5 as released.

Use docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md and docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md as the draft release-note sources. Preserve:
- RC decision: RC_READY_WITH_DEFERRALS
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- ni run compiles a bounded handoff prompt only.
- READY is planning contract readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- fixture relock is separate from project-root relock.
- benchmark evidence keeps not_measured boundaries.

Audit the English and Korean drafts for drift, missing required sections, overclaims, and release-like wording. Make only narrow wording fixes or cross-links if needed. Do not broadly rewrite README, CHANGELOG, release history, skill docs, scripts, or runtime code.

Run:
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final report must confirm no project-root relock, no protected .ni file changes, no generated prompt execution, no release/tag/publish/upload action, no Homebrew Available claim, no model workspace Available claim, no no-terminal deterministic claim, no benchmark overclaim, and whether validation scripts exercised fixture ni end while keeping fixture runs distinct from project-root relock.
```
