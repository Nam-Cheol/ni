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

v0.5 is a release-candidate draft for `ni`, a Project Intent Compiler for AI
Agents. It strengthens the Intent Lock Protocol around conversation proof,
change control, stale-lock diagnostics, assisted no-terminal workflows, model
workspace boundaries, and benchmark claim boundaries before downstream agent
work begins.

This draft does not claim that v0.5 has been released. `ni run` still compiles a
bounded handoff prompt from a current locked plan and does not execute
downstream work.

### Highlights

- Project Intent Compiler positioning remains explicit: `ni` compiles intent
  before downstream agent work.
- Proof capture reliability work clarifies what planning evidence can and
  cannot prove.
- Change-control UX audit records the stale-lock and amended-intent safety
  model.
- `LOCK-STALE` diagnostic documents that an existing lock no longer matches
  current planning inputs.
- Amend/relock workflow examples explain the recovery order without making the
  model the authority.
- No-terminal stale-lock examples and transcript checklist distinguish
  model-only drafts, pasted CLI output, trusted runner transcripts, fixture
  evidence, and target-workspace evidence.
- Changed-intent fixture coverage broadens stale-lock coverage while keeping
  fixture relock separate from project-root relock.
- Model workspace wording verification preserves Experimental status and skill
  boundaries.
- Benchmark claim-boundary hardening keeps `not_measured` limits visible.
- Task 190 recorded the RC readiness decision:
  `RC_READY_WITH_DEFERRALS`.

### Reliability improvements

The v0.5 reliability set in `docs/101` through `docs/109` adds a connected
record for proof capture, change control, stale-lock diagnostics, amend/relock
examples, model workspace wording, no-terminal examples, transcript quality,
changed-intent fixtures, and release-readiness sweep coverage.

Together, these docs keep the core protocol conservative:

```text
conversation -> project contract -> readiness gate -> lock hash -> downstream handoff
```

The reliability work does not add an execution harness, task runner, shell
adapter, Codex exec adapter, queue, PR automation, release automation, or
execution evidence loop.

### Stale-lock and change-control

v0.5 improves the changed-intent path around locked plans:

- `LOCK-STALE` means the existing lock no longer matches current planning
  inputs.
- `ni status` can warn when the project is otherwise ready but the existing
  lock is stale.
- `ni end` remains the CLI-authoritative lock or relock step after changed
  intent has been reviewed.
- `ni run` refuses stale handoff and points back to the recovery order.
- Fixture coverage checks representative changed-intent cases, non-lockable
  false positives, fixture relock recovery, and project-root safety.
- Amend/relock examples show the user path without weakening readiness,
  mitigation, risk, or non-goal criteria.

Recovery order remains:

```text
review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000
```

### Proof capture and acceptance evidence

The proof capture and acceptance evidence work clarifies that planning proof is
an audit trail, not implementation proof. `ni status` remains the authority for
readiness, `ni end` remains the authority for locking, and `ni run` remains the
authority for bounded prompt compilation from a current lock.

The v0.5 acceptance evidence matrix defines the files, commands, status
vocabulary, and `not_measured` boundaries expected for each v0.5 lane. It helps
future maintainers avoid upgrading draft, fixture, or benchmark evidence into
product readiness claims.

### Benchmark claim boundaries

Benchmark evidence supports bounded planning-artifact claims only. The current
benchmark surfaces may show that a vague request can become an auditable
planning artifact, a `READY` isolated benchmark workspace, and a bounded
handoff prompt.

They do not prove implementation correctness, downstream agent performance,
rework reduction, adoption improvement, cost improvement, latency improvement,
statistical effect, dashboard product readiness, real research approval,
fieldwork authorization, research quality, or intervention effectiveness.
Those remain `not_measured` where relevant.

### Model workspace and skills

- Model workspace packs remain Experimental.
- Skills are UX; CLI is authority.
- Skills may draft or explain planning text.
- Skills may help explain `LOCK-STALE`.
- Skills do not determine readiness.
- Skills do not lock or relock.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- Skills do not update `.ni/plan.lock.json`.
- Host-level/global install, provider runtime behavior, and cross-machine
  install remain unverified unless later host-specific proof exists.

### No-terminal assisted workflow

- No-terminal method remains Experimental / assisted.
- Model-only draft is draft-only.
- Exact CLI output from a trusted runner is required before claiming readiness,
  lock freshness, relock, hash verification, or bounded handoff compilation.
- A fixture transcript supports fixture claims only.
- A target-workspace claim requires exact target-workspace command output.
- No-terminal assistance does not provide deterministic validation by itself.

### Validation and tests

The v0.5 RC readiness source records passing validation for CLI readiness, Go
tests, install claim checks, skill-pack checks, demo checks, smoke checks,
quality checks, install checks, release checks, and protected-file diff checks.

Relevant fixture coverage includes stale-lock warnings, stale `ni run` refusal,
fixture relock recovery, changed-intent lockable input coverage, non-lockable
false positives, benchmark claim-boundary checks, no-terminal transcript
boundaries, ni-grill docs-only boundaries, and seed-only export boundaries.

Validation scripts may exercise `ni end` inside temporary fixture workspaces.
Those fixture runs must not be described as project-root relock.

### Known deferrals

| Deferral | Status | Why deferred | Required future evidence |
| --- | --- | --- | --- |
| Homebrew implementation / availability | Planned / v0.5 candidate | No tap/formula availability is claimed. | Tap, formula, sha256, audit, `brew install` output, `ni --help`, and `ni version` verification. |
| Model workspace host verification | Experimental | Host-level/global install, provider runtime behavior, and cross-machine install are not verified. | Host-specific install or discovery proof plus provider behavior transcript. |
| External user validation | Limited | RC docs preserve boundaries but do not prove external adoption or user success. | User-run transcripts, maintained external validation notes, or scoped user research. |
| Additional benchmark breadth | Bounded | Current benchmark evidence remains qualitative and planning-artifact scoped. | Additional benchmark case or broader report that preserves `not_measured` boundaries. |
| No-terminal deterministic validation not claimed | Experimental / assisted | Model-only and pasted-output workflows cannot prove CLI state without exact trusted CLI output. | Trusted runner transcripts from the target workspace for `ni status`, `ni end`, hash verification, and `ni run`. |
| v0.5 release artifacts | Draft only | This task does not tag, publish, upload assets, or create a GitHub release. | Release artifact dry-run, final release preflight, actual tag/release action by an authorized maintainer. |

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

This section is a draft pending actual v0.5 release artifact verification. It
must not be copied into final release notes as a v0.5 availability claim until
the final release path is verified.

Currently true install paths:

- Source: Available from this checkout with Go.
- Local binary: Available from this checkout with `make build` and `./bin/ni`.
- Release binary: Available for verified v0.4.0 release assets.
- Curl installer: Available for verified v0.4.0 release assets.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Homebrew: Planned / v0.5 candidate.

No downloadable v0.5 artifacts are claimed by this draft.

### Maintainer validation checklist

Run before final release-note promotion or any release action:

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
| Homebrew | Homebrew: Planned / v0.5 candidate. | Homebrew Available; package-manager install works. | Yes | Availability requires tap, formula, sha256, `brew install`, `ni --help`, and `ni version` proof. |
| Model workspace packs | Model workspace packs: Experimental. | Model workspace packs Available; global install verified; provider runtime behavior verified. | Yes | Host-level/global behavior remains unverified. |
| No-terminal | No-terminal method: Experimental / assisted. | No-terminal deterministic validation; model-only draft as CLI state. | Yes | Exact trusted CLI output is required for state claims. |
| ni run | Compiles a bounded handoff prompt only. | Executes downstream work, shell commands, Codex, agents, or queues. | Yes | Prompt compilation remains pre-runtime. |
| READY | Planning contract readiness only. | Product readiness, implementation correctness, or downstream success. | Yes | Scope remains the planning artifact. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | Product failure, implementation failure, or benchmark failure. | Yes | Stale can coexist with current planning `READY`. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | Causal impact, adoption, cost, latency, dashboard readiness, research approval, fieldwork authorization. | Yes | Unsupported benchmark claims remain `not_measured`. |
| Fixture relock | Fixture relock is separate from project-root relock. | Validation fixture relock updated project-root `.ni/plan.lock.json`. | Yes | Fixture runs support fixture claims only. |
| Trusted runner transcript | Claims scoped to exact workspace, command, output, and time. | Global lock freshness or target-workspace state without exact output. | Yes | No-terminal claims need exact CLI output. |
| Runtime execution boundary | `ni` is not an execution harness. | Task runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, or execution evidence loop. | Yes | Kernel remains authoritative for planning contracts, readiness, locks, and prompt compilation. |
| RC decision | `RC_READY_WITH_DEFERRALS`. | v0.5 released; RC has no deferrals. | Yes | Task 190 decision is preserved conservatively. |

## Changes made

- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md`: adds this English RC polish
  and release-note draft.
- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md`: adds the Korean
  companion.
- `docs/51_POST_RELEASE_ROADMAP.md`: adds a narrow pointer to this draft.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: adds the matching Korean pointer.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`: adds a narrow follow-up
  pointer to this draft.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`: adds the matching
  Korean follow-up pointer.

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

Why: the release-note wording is now drafted and claim-audited, but one final
preflight should verify the draft, validation commands, protected-file safety,
and release/non-release boundary before any release action is considered.

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
