# v0.5 Release Notes Final Preflight

## Current status

- RC decision: RC_READY_WITH_DEFERRALS
- Release notes draft: docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- This preflight does not publish, tag, release, or upload assets.

## Preflight goal

이 preflight는 later release process 전에 v0.5 RC release-note wording, adjacent
RC docs, roadmap navigation, Korean companion consistency, git visibility,
validation commands, forbidden release or availability claims를 확인한다. 이것은
documentation and claim-boundary audit only다.

## Decision

Decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES.

Justification: release-note draft는 conservative하며 later release use에 적합하다.
`RC_READY_WITH_DEFERRALS`가 preserved되었고 known deferrals가 explicit하게 남아
있으며 release, asset-upload, Homebrew Available, model workspace Available,
no-terminal deterministic, benchmark overclaim, ni-run execution,
fixture-as-root relock claim은 발견되지 않았다. Notes는 새 final preflight docs와
narrow cross-link edits를 next commit에 포함해야 한다는 commit-inclusion notes뿐이다.

## Release-note claim audit

| Claim area | Expected boundary | Observed wording | Pass? | Notes |
| --- | --- | --- | --- | --- |
| RC decision | `RC_READY_WITH_DEFERRALS`; release completion이 아님. | docs/111은 `RC_READY_WITH_DEFERRALS`를 state하고 document가 draft-only라고 말한다. | Yes | Decision은 deferrals를 upgrade하지 않고 preserved된다. |
| Published/released status | v0.5가 published 또는 released되었다고 claim하지 않는다. | docs/111은 v0.5를 publish, tag, release하지 않는다고 말한다. | Yes | Draft에서 release-like wording은 발견되지 않았다. |
| Artifact upload status | Uploaded v0.5 artifacts claim 없음. | docs/111은 downloadable v0.5 artifacts를 claim하지 않는다고 말한다. | Yes | Release binary availability는 verified v0.4.0 assets에 scoped된다. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | docs/111은 Homebrew Planned를 유지하고 required future evidence를 listed한다. | Yes | Homebrew Available claim 없음. |
| Model workspace packs | Model workspace packs: Experimental. | docs/111은 Experimental과 unverified host/provider boundaries를 유지한다. | Yes | Broad Available 또는 host-verification claim 없음. |
| No-terminal | No-terminal method: Experimental / assisted. | docs/111은 CLI-state claims에 exact trusted CLI output을 요구한다. | Yes | Deterministic no-terminal validation claim 없음. |
| ni run | Bounded handoff prompt compilation only. | docs/111은 `ni run`이 bounded prompt를 compile하며 downstream work를 execute하지 않는다고 말한다. | Yes | Executor claim 없음. |
| READY | Planning contract readiness only. | docs/111은 `READY`를 planning artifact readiness로 scope한다. | Yes | Product-readiness claim 없음. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | docs/111은 expected meaning과 recovery order를 반복한다. | Yes | Product 또는 implementation failure claim 없음. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | docs/111은 unsupported claims를 listed하고 `not_measured` boundaries를 유지한다. | Yes | Causal, adoption, cost, latency, dashboard, research, fieldwork overclaim 없음. |
| Fixture relock | Fixture relock is separate from project-root relock. | docs/111은 fixture runs를 project-root relock으로 describe하면 안 된다고 말한다. | Yes | Boundary remains visible. |
| Trusted runner transcript | Claims scoped to exact workspace, command, output, and time. | docs/111은 no-terminal readiness, lock freshness, relock, hash verification, bounded handoff compilation claims에 exact CLI output을 요구한다. | Yes | Model-only CLI-state claim 없음. |
| Runtime execution boundary | `ni` is not an execution harness. | docs/111은 execution harness, task runner, adapters, queues, PR automation, release automation, execution evidence loop를 exclude한다. | Yes | Kernel boundary preserved. |

## Adjacent-doc consistency audit

| Doc | Expected relationship to docs/111 | Pass? | Notes |
| --- | --- | --- | --- |
| docs/110 | Readiness audit는 docs/111로 이어지고 `RC_READY_WITH_DEFERRALS`를 preserve해야 한다. | Yes | docs/110은 docs/111로 link한다; 이 preflight는 narrow docs/112 follow-up pointer를 추가한다. |
| docs/109 | Release-readiness sweep은 prior reliability-doc audit으로 남아야 한다. | Yes | docs/109는 RC decision을 docs/110에 defer하고 status boundaries를 aligned하게 유지한다. |
| docs/51 roadmap | Roadmap은 excessive links 없이 docs/109, docs/110, docs/111, now docs/112를 index해야 한다. | Yes | Narrow docs/112 pointer 추가. |
| README | Public install/status claims는 unchanged로 남아야 한다. | Yes | README는 release binary/curl Available을 v0.4.0에 scope하고 Homebrew Planned를 유지한다. |
| README.ko | Korean public install/status claims는 unchanged로 남아야 한다. | Yes | README.ko도 same status split을 유지한다. |
| docs/97 benchmark boundaries | Benchmark limits는 docs/111 `not_measured` language와 match해야 한다. | Yes | docs/111은 planning-artifact and unsupported-claim boundaries를 preserve한다. |
| docs/99 model workspace status | Model workspace broad path는 Experimental로 남아야 한다. | Yes | docs/111은 Experimental, `not_verified`, CLI-authority boundaries를 preserve한다. |
| docs/103 stale lock diagnostic | `LOCK-STALE` meaning은 implemented diagnostic과 match해야 한다. | Yes | docs/111은 same existing-lock/current-input mismatch meaning을 사용한다. |
| docs/107 transcript checklist | No-terminal transcript categories는 scoped로 남아야 한다. | Yes | docs/111은 model-only, trusted runner, fixture, target-workspace limits를 mirror한다. |
| docs/108 changed-intent fixture coverage | Fixture relock은 project-root relock과 separate로 남아야 한다. | Yes | docs/111과 이 preflight는 distinction을 preserve한다. |

## Korean companion audit

| Doc pair | Pass? | Notes |
| --- | --- | --- |
| docs/111 English/Korean | Yes | Korean companion은 status constants, draft-only wording, `not_measured`, no-terminal, Homebrew, model workspace boundaries를 preserve한다. |
| docs/112 English/Korean | Yes | 같은 decision, checks, deferrals, next task를 가진 Korean companion을 추가했다. |
| docs/110 English/Korean | Yes | Pair는 `RC_READY_WITH_DEFERRALS`를 preserve한다; matching docs/112 follow-up pointers 추가. |
| docs/51 English/Korean if touched | Yes | Matching roadmap pointers for docs/112 추가. |

## Known deferrals check

| Deferral | Still explicit? | Blocks release-note preflight? | Required future evidence |
| --- | --- | --- | --- |
| Homebrew implementation / availability | Yes | No | Tap, formula, sha256, audit, `brew install` output, `ni --help`, `ni version` verification. |
| Model workspace host verification | Yes | No | Host-specific install/discovery proof and provider behavior transcript. |
| External user validation | Yes | No | User-run transcripts, maintained external validation notes, or scoped user research. |
| Additional benchmark breadth if relevant | Yes | No | Additional benchmark case or broader report preserving `not_measured` boundaries. |
| No-terminal deterministic validation not claimed | Yes | No | Exact trusted CLI output from the target workspace for readiness, lock freshness, relock, hash verification, bounded handoff compilation claims. |

## Git status / inclusion check

| Path or group | Status from git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| docs/110_* | `M docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`; `M docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md` | Yes | Tracked files; narrow docs/112 follow-up pointers only. |
| docs/111_* | `M docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md`; `M docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md` | Yes | Tracked files; narrow docs/112 follow-up pointers only. |
| docs/112_* | `?? docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`; `?? docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md` | Yes | New final preflight docs from this task. |
| docs/51* | `M docs/51_POST_RELEASE_ROADMAP.md`; `M docs/51_POST_RELEASE_ROADMAP.ko.md` | Yes | Tracked roadmap pointers to docs/112. |
| .ni/contract.json | no output | No | Protected project-root file unchanged. |
| .ni/session.json | no output | No | Protected project-root file unchanged. |
| .ni/plan.lock.json | no output | No | Protected project-root file unchanged. |
| unexpected files | none | No | 이 task 후 unexpected modified or untracked files는 없다. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Pre-edit status was clean; post-edit status shows expected docs/51, docs/110, docs/111 modifications and untracked docs/112 files. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, warnings are `None`. |
| `python3 scripts/check-install-docs.py` | Pass | Install/distribution claim markers pass. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack status, CLI authority, stale-lock, no-relock boundaries pass. |
| `bash scripts/demo-check.sh` | Pass | Demo, benchmark, no-terminal, ni-grill, seed-only boundaries pass. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passes. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Smoke checks pass; fixture `ni end` and fixture relock paths are temporary-workspace checks only. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source, build, local binary, temporary local install paths pass. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Release-readiness and claim-boundary checks pass. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` file diff. |

## Findings

| Finding | Severity | Surface | Change made | Blocks preflight? |
| --- | --- | --- | --- | --- |
| RC release-note wording is conservative and aligned with audited boundaries. | Info | docs/111 and adjacent docs | Added final preflight record. | No |
| New docs/112 files and narrow cross-link edits must be included with the RC audit/release-note chain in the next commit. | Low | git status | Recorded inclusion note in this preflight. | No |

## Changes made

- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`: English final preflight
  record 추가.
- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`: Korean companion 추가.
- `docs/51_POST_RELEASE_ROADMAP.md`: narrow docs/112 pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer 추가.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`: narrow docs/112
  follow-up pointer 추가.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`: matching Korean
  follow-up pointer 추가.
- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md`: narrow docs/112 follow-up
  pointer 추가.
- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md`: matching Korean
  follow-up pointer 추가.

## What this preflight proves

- release-note draft wording is aligned with audited v0.5 boundaries
- RC_READY_WITH_DEFERRALS is preserved
- known deferrals remain explicit
- no release action was performed
- expected files for the RC audit / release-note chain are visible to git status

## What this preflight does not prove

- v0.5 has been published
- release artifacts were uploaded
- GitHub release exists
- Homebrew is Available
- model workspace host behavior is verified
- no-terminal is deterministic
- downstream execution succeeds
- external users succeed
- benchmark effect size or causal impact

## Recommended next task

Selected next task: A. v0.5 artifact dry-run audit.

Why: release-note wording is clean enough for later use, and the next uncertainty
before any release action is whether v0.5 artifacts can be packaged and checked
in a dry-run path without publishing, tagging, uploading assets, or upgrading
deferred distribution claims.

## Next task prompt

```text
Proceed with a v0.5 artifact dry-run audit in /Users/namba/Documents/project/ni.

This is a dry-run audit and documentation task only. Do not publish, tag, create a GitHub release, upload assets, run goreleaser publish, create Homebrew formula files, run ni end on the project root, relock the project root, execute generated prompts, add release automation, add runtime execution behavior, or mark v0.5 as released.

Use these docs as the release-note and RC-readiness sources:
- docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md
- docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md
- docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md

Preserve:
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

Goal:
Audit what a v0.5 artifact dry-run would need to verify before any release action. Confirm package/build commands, expected archive/checksum names, version/help behavior, release-note inputs, protected-file safety, and claim boundaries. Do not claim artifacts exist unless the dry run actually creates local temporary artifacts, and do not claim they were published.

Required checks:
- git status --short
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
