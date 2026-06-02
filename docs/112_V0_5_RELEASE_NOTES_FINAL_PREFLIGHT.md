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

This preflight checks the v0.5 RC release-note wording, adjacent RC docs,
roadmap navigation, Korean companion consistency, git visibility, validation
commands, and forbidden release or availability claims before any later release
process. It is a documentation and claim-boundary audit only.

## Decision

Decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES.

Justification: the release-note draft is conservative and suitable for later
release use; `RC_READY_WITH_DEFERRALS` is preserved; known deferrals remain
explicit; and no release, asset-upload, Homebrew Available, model workspace
Available, no-terminal deterministic, benchmark overclaim, ni-run execution, or
fixture-as-root relock claim was found. The only notes are commit-inclusion
notes for the newly added final preflight docs and narrow cross-link edits.

## Release-note claim audit

| Claim area | Expected boundary | Observed wording | Pass? | Notes |
| --- | --- | --- | --- | --- |
| RC decision | `RC_READY_WITH_DEFERRALS`; not release completion. | docs/111 states `RC_READY_WITH_DEFERRALS` and says the document is draft-only. | Yes | Decision is preserved without upgrading deferrals. |
| Published/released status | v0.5 is not claimed as published or released. | docs/111 says it does not publish, tag, or release v0.5. | Yes | No release-like wording found in the draft. |
| Artifact upload status | No v0.5 artifacts are claimed uploaded. | docs/111 says no downloadable v0.5 artifacts are claimed. | Yes | Release binary availability remains scoped to verified v0.4.0 assets. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | docs/111 keeps Homebrew Planned and lists required future evidence. | Yes | No Homebrew Available claim. |
| Model workspace packs | Model workspace packs: Experimental. | docs/111 keeps Experimental and unverified host/provider boundaries. | Yes | No broad Available or host-verification claim. |
| No-terminal | No-terminal method: Experimental / assisted. | docs/111 requires exact trusted CLI output for CLI-state claims. | Yes | No deterministic no-terminal validation claim. |
| ni run | Bounded handoff prompt compilation only. | docs/111 says `ni run` compiles a bounded prompt and does not execute downstream work. | Yes | No executor claim. |
| READY | Planning contract readiness only. | docs/111 scopes `READY` to planning artifact readiness. | Yes | No product-readiness claim. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | docs/111 repeats the expected meaning and recovery order. | Yes | No product or implementation failure claim. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits. | docs/111 lists unsupported claims and keeps `not_measured` boundaries. | Yes | No causal, adoption, cost, latency, dashboard, research, or fieldwork overclaim. |
| Fixture relock | Fixture relock is separate from project-root relock. | docs/111 says fixture runs must not be described as project-root relock. | Yes | Boundary remains visible. |
| Trusted runner transcript | Claims scoped to exact workspace, command, output, and time. | docs/111 requires exact CLI output for no-terminal readiness, lock freshness, relock, hash verification, and bounded handoff compilation claims. | Yes | No model-only CLI-state claim. |
| Runtime execution boundary | `ni` is not an execution harness. | docs/111 excludes execution harness, task runner, adapters, queues, PR automation, release automation, and execution evidence loop. | Yes | Kernel boundary preserved. |

## Adjacent-doc consistency audit

| Doc | Expected relationship to docs/111 | Pass? | Notes |
| --- | --- | --- | --- |
| docs/110 | Readiness audit should feed docs/111 and preserve `RC_READY_WITH_DEFERRALS`. | Yes | docs/110 links to docs/111; this preflight adds a narrow docs/112 follow-up pointer. |
| docs/109 | Release-readiness sweep should remain the prior reliability-doc audit. | Yes | docs/109 defers RC decision to docs/110 and keeps status boundaries aligned. |
| docs/51 roadmap | Roadmap should index docs/109, docs/110, docs/111, and now docs/112 without excessive links. | Yes | Narrow docs/112 pointer added. |
| README | Public install/status claims should stay unchanged. | Yes | README keeps release binary/curl Available for v0.4.0 and Homebrew Planned. |
| README.ko | Korean public install/status claims should stay unchanged. | Yes | README.ko keeps the same status split. |
| docs/97 benchmark boundaries | Benchmark limits should match docs/111 `not_measured` language. | Yes | docs/111 preserves planning-artifact and unsupported-claim boundaries. |
| docs/99 model workspace status | Model workspace broad path should remain Experimental. | Yes | docs/111 preserves Experimental, `not_verified`, and CLI-authority boundaries. |
| docs/103 stale lock diagnostic | `LOCK-STALE` meaning should match the implemented diagnostic. | Yes | docs/111 uses the same existing-lock/current-input mismatch meaning. |
| docs/107 transcript checklist | No-terminal transcript categories should remain scoped. | Yes | docs/111 mirrors model-only, trusted runner, fixture, and target-workspace limits. |
| docs/108 changed-intent fixture coverage | Fixture relock must remain separate from project-root relock. | Yes | docs/111 and this preflight preserve the distinction. |

## Korean companion audit

| Doc pair | Pass? | Notes |
| --- | --- | --- |
| docs/111 English/Korean | Yes | Korean companion preserves status constants, draft-only wording, `not_measured`, no-terminal, Homebrew, and model workspace boundaries. |
| docs/112 English/Korean | Yes | Korean companion added with the same decision, checks, deferrals, and next task. |
| docs/110 English/Korean | Yes | Pair preserves `RC_READY_WITH_DEFERRALS`; matching docs/112 follow-up pointers added. |
| docs/51 English/Korean if touched | Yes | Matching roadmap pointers added for docs/112. |

## Known deferrals check

| Deferral | Still explicit? | Blocks release-note preflight? | Required future evidence |
| --- | --- | --- | --- |
| Homebrew implementation / availability | Yes | No | Tap, formula, sha256, audit, `brew install` output, `ni --help`, and `ni version` verification. |
| Model workspace host verification | Yes | No | Host-specific install/discovery proof and provider behavior transcript. |
| External user validation | Yes | No | User-run transcripts, maintained external validation notes, or scoped user research. |
| Additional benchmark breadth if relevant | Yes | No | Additional benchmark case or broader report preserving `not_measured` boundaries. |
| No-terminal deterministic validation not claimed | Yes | No | Exact trusted CLI output from the target workspace for readiness, lock freshness, relock, hash verification, and bounded handoff compilation claims. |

## Git status / inclusion check

| Path or group | Status from git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| docs/110_* | `M docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`; `M docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md` | Yes | Tracked files; only narrow docs/112 follow-up pointers added. |
| docs/111_* | `M docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md`; `M docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md` | Yes | Tracked files; only narrow docs/112 follow-up pointers added. |
| docs/112_* | `?? docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`; `?? docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md` | Yes | New final preflight docs from this task. |
| docs/51* | `M docs/51_POST_RELEASE_ROADMAP.md`; `M docs/51_POST_RELEASE_ROADMAP.ko.md` | Yes | Tracked roadmap pointers to docs/112. |
| .ni/contract.json | no output | No | Protected project-root file unchanged. |
| .ni/session.json | no output | No | Protected project-root file unchanged. |
| .ni/plan.lock.json | no output | No | Protected project-root file unchanged. |
| unexpected files | none | No | No unexpected modified or untracked files observed after this task. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Pre-edit status was clean; post-edit status shows expected docs/51, docs/110, docs/111 modifications and untracked docs/112 files. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, and warnings are `None`. |
| `python3 scripts/check-install-docs.py` | Pass | Install/distribution claim markers pass. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack status, CLI authority, stale-lock, and no-relock boundaries pass. |
| `bash scripts/demo-check.sh` | Pass | Demo, benchmark, no-terminal, ni-grill, and seed-only boundaries pass. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passes. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Smoke checks pass; fixture `ni end` and fixture relock paths are temporary-workspace checks only. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source, build, local binary, and temporary local install paths pass. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Release-readiness and claim-boundary checks pass. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` file diff. |

## Findings

| Finding | Severity | Surface | Change made | Blocks preflight? |
| --- | --- | --- | --- | --- |
| RC release-note wording is conservative and aligned with audited boundaries. | Info | docs/111 and adjacent docs | Added final preflight record. | No |
| New docs/112 files and narrow cross-link edits must be included with the RC audit/release-note chain in the next commit. | Low | git status | Recorded inclusion note in this preflight. | No |

## Changes made

- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md`: added this English final
  preflight record.
- `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`: added Korean companion.
- `docs/51_POST_RELEASE_ROADMAP.md`: added narrow docs/112 pointer.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: added matching Korean pointer.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`: added narrow docs/112
  follow-up pointer.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`: added matching
  Korean follow-up pointer.
- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md`: added narrow docs/112
  follow-up pointer.
- `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md`: added matching Korean
  follow-up pointer.

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

Why: the release-note wording is clean enough for later use, and the next
uncertainty before any release action is whether v0.5 artifacts can be packaged
and checked in a dry-run path without publishing, tagging, uploading assets, or
upgrading deferred distribution claims.

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
