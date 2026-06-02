# v0.5 Release Candidate Readiness Audit

## Current status

- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Audit goal

이 audit은 현재 repo를 v0.5 release-candidate-ready state로 취급할 수 있는지
결정한다. Evidence, claim boundaries, deferrals, residual risks, next task를
기록한다. Publish, tag, release, project root에서 `ni end`, project-root relock,
Go implementation, lock semantics, skill packages, static checks, release
automation은 변경하지 않는다.

## Decision

Decision: `RC_READY_WITH_DEFERRALS`.

Justification: 현재 repo는 required CLI, Go, install, skill-pack, demo, smoke,
quality, release checks를 pass한다. `ni status`는 `READY`를 report한다.
Project-root lockfiles는 unchanged이며 release, installer, model workspace,
no-terminal, benchmark, stale-lock, skill, kernel boundary overclaim은 발견되지
않았다. 남은 gap은 RC blocker가 아니라 explicit deferrals다. Homebrew는
`Planned / v0.5 candidate`, model workspace host-level behavior는
`Experimental`, no-terminal은 `Experimental / assisted`, external user
validation은 limited, benchmark breadth는 bounded로 남는다.

## RC readiness criteria

| Criterion | Result | Evidence | Notes |
| --- | --- | --- | --- |
| Positioning and README accuracy | Pass | README, README.ko.md, docs/40, docs/53, docs/109 | Public wording은 ni를 pre-runtime Project Intent Compiler로 보존한다. |
| CLI validation | Pass | `ni status --proof --next-questions` | Project root는 blockers, deferrals, warnings 없이 `READY`를 report한다. |
| Go tests | Pass | `go test ./...` | 이 audit은 Go files를 변경하지 않았다. |
| Release/install docs | Pass | README, install docs, `check-install-docs.py` | Release binary와 curl installer는 Available이며 Homebrew는 overclaim하지 않는다. |
| Release binary status | Pass | README, docs/53, install docs | Release binary: Available. |
| Curl installer status | Pass | README, docs/install-curl*, installer docs | Curl installer: Available. |
| Homebrew status | Deferral | README, docs/54, docs/80, docs/109 | Homebrew: Planned / v0.5 candidate. |
| v0.5 reliability docs | Pass | docs/101 through docs/109 | Reliability docs are linked and boundary-audited. |
| Stale-lock/change-control reliability | Pass | docs/103, docs/104, docs/108, tests | `LOCK-STALE`와 recovery order가 documented and tested 상태다. |
| Benchmark claim boundaries | Pass | docs/97, examples/benchmark-report, `demo-check.sh` | Benchmark evidence는 bounded and not causal 상태다. |
| Model workspace status | Deferral | docs/99, docs/105, package READMEs | Model workspace packs: Experimental. |
| No-terminal status | Deferral | docs/no-terminal*, docs/106, docs/107 | No-terminal method: Experimental / assisted. |
| Skill-pack boundaries | Pass | packages/*-skills, .agents/skills, `check-skill-packs.sh` | Skills are UX; CLI is authority. |
| Korean companions | Pass | `.ko.md` companions and Korean roadmap | Exact status terms are preserved. |
| Project-root lockfile safety | Pass | protected-file diff check | `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` unchanged. |
| Runtime boundary | Pass | README, docs/53, docs/109, command docs | Kernel boundary는 task runner, adapter, queue, downstream execution behavior를 exclude한다. |

## Evidence inventory

| Evidence | Result | Notes |
| --- | --- | --- |
| ni status | Pass | Project root prints `NI Intent Readiness: READY`; blockers, deferrals, warnings는 `None`. |
| go test ./... | Pass | All packages pass. |
| check-install-docs.py | Pass | Install/distribution claim markers pass. |
| check-skill-packs.sh | Pass | Skill-pack status, CLI authority, stale-lock, no-relock boundaries pass. |
| demo-check.sh | Pass | Benchmark, no-terminal, transcript, ni-grill, seed-only boundaries pass. |
| quality.sh | Pass | Broad quality wrapper passes. |
| smoke.sh | Pass | Fixture workflows, locks, stale refusal, prompt compilation, seed boundaries pass. |
| install-check.sh | Pass | Source, build, local binary, temporary local install paths pass. |
| release-check.sh | Pass | Release-readiness and claim-boundary checks pass. |
| git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json | Pass | No protected-file diff. |
| docs/101 through docs/109 | Pass | v0.5 reliability set is present and linked from roadmap. |
| examples/benchmark-report | Pass | `not_measured` and planning-artifact boundaries remain visible. |
| examples/ni-grill | Pass | Grill examples remain docs/review evidence, not readiness authority. |
| skill packages | Pass | Codex, Claude, repo-local skills preserve UX-only boundaries. |

## Blockers

None.

## Deferrals

| Deferral | Status | RC impact | Required future proof |
| --- | --- | --- | --- |
| Homebrew implementation/availability | Planned / v0.5 candidate | Availability를 claim하지 않으므로 non-blocking. | Tap/formula, checksums, audit, local formula install, published tap install, `ni --help`, `ni version`. |
| Model workspace host verification | Experimental | Broad status가 Experimental이므로 non-blocking. | Host-specific install or discovery proof and provider behavior transcript. |
| External user validation | Limited | RC documentation readiness에는 non-blocking이지만 adoption risk다. | User-run transcripts or maintained external validation notes. |
| Additional benchmark breadth | Bounded | Current claims가 qualitative and scoped라서 non-blocking. | Third benchmark case 또는 broader benchmark report without causal overclaims. |

## Warnings

None.

## Risks

| Risk | Residual impact | Mitigation |
| --- | --- | --- |
| docs may drift | Future edits could overclaim status or authority. | RC path에서 `check-install-docs.py`, `check-skill-packs.sh`, `demo-check.sh`, release checks를 유지한다. |
| static checks cover selected phrases not all semantic overclaims | Passing scripts는 exhaustive semantic proof가 아니다. | Release notes 또는 public launch copy 전 human claim-boundary audit을 유지한다. |
| provider host behavior unverified | Model workspace UX may behave differently in real hosts. | Host proof 전까지 Model workspace packs: Experimental 유지. |
| Homebrew Planned until tested | Users may expect package-manager availability. | Formula/tap validation 전까지 Homebrew: Planned / v0.5 candidate 유지. |
| external user validation limited | Adoption issues may appear after RC. | External validation을 post-RC 또는 release-polish task로 다룬다. |
| benchmark evidence bounded/not causal | Case studies can be misread as product or performance proof. | `not_measured` and planning-artifact boundaries 보존. |
| no-terminal assisted/not deterministic | Users may confuse assisted drafting with CLI proof. | No-terminal docs에서 exact trusted CLI output requirements 유지. |

## Claim-boundary audit

| Claim area | Expected boundary | Result |
| --- | --- | --- |
| READY | Planning contract readiness only; product readiness가 아니다. | Pass |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | Pass |
| ni run | Bounded handoff prompt만 compile하고 stale handoff를 refuse한다. | Pass |
| Homebrew | Planned / v0.5 candidate; not Available. | Pass |
| Model workspace packs | Experimental; host/global behavior is not verified. | Pass |
| No-terminal | Experimental / assisted; CLI-state claims require exact CLI output. | Pass |
| Skills | Skills are UX; CLI is authority. | Pass |
| Benchmark evidence | Planning-artifact evidence only; impact, adoption, cost, latency are `not_measured`. | Pass |
| Fixture relock | Fixture relock is separate from project-root relock. | Pass |
| Trusted runner transcript | Shown workspace, command, output, time only를 support한다. | Pass |
| Runtime execution | `ni-kernel`에 포함되지 않으며 downstream work는 core 밖에 남는다. | Pass |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; no blockers, deferrals, warnings. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack checks passed. |
| `bash scripts/demo-check.sh` | Pass | Initial sandboxed run hit default Go build cache permissions; approved exact-command rerun passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Quality checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Smoke checks passed in temporary fixtures. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Install checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Release checks passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected-file diff. |

Validation scripts는 temporary workspaces에서 fixture `ni end` flows를 exercise한다.
이 fixture relocks는 project-root relock과 별개다. 이 audit은 project root에서
`ni end`를 run하지 않았다.

## Changed files

- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`: this RC readiness audit 추가.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`: Korean companion 추가.
- `docs/51_POST_RELEASE_ROADMAP.md`: this RC audit에 대한 narrow pointer 추가.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer 추가.

## What this audit proves

- Current repo can be treated as `RC_READY_WITH_DEFERRALS`.
- Required validation commands pass.
- Current public status claims remain bounded.
- Root `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` were not changed.
- Homebrew, model workspace packs, no-terminal, benchmark, skill, kernel
  boundary claims are preserved.

## What this audit does not prove

- Homebrew availability.
- Provider host behavior for model workspace packs.
- Deterministic no-terminal validation without exact CLI output.
- External user adoption or satisfaction.
- Benchmark causal effect, statistical significance, adoption, cost, latency, or implementation quality.
- Downstream execution quality.

## Recommended next task

Selected next task: E. v0.5 RC polish / release notes draft.

Why: RC decision은 `RC_READY_WITH_DEFERRALS`이며, 다음 useful step은 publish, tag,
release 또는 deferred status upgrade 없이 release-note wording과 public RC polish를
준비하는 것이다.

Follow-up draft: `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md`는
`RC_READY_WITH_DEFERRALS`, known deferrals, draft-only wording,
release/non-release boundaries를 보존하면서 RC polish and release-note draft를
기록한다.

Final preflight: `docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`는 release
action 전에 release-note draft, adjacent RC docs, navigation, git visibility,
forbidden release or availability claims를 check한다.

## Next task prompt

```text
Proceed with v0.5 RC polish and a release notes draft in /Users/namba/Documents/project/ni.

This is a documentation-only polish task. Do not publish, tag, release, run ni end on the project root, relock the project root, edit .ni/contract.json, edit .ni/session.json, edit .ni/plan.lock.json, change Go implementation, change skill packages, add release automation, add Homebrew availability claims, mark model workspace packs Available, claim deterministic no-terminal validation, or claim benchmark causal impact.

Use docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md as the readiness source. Preserve the decision RC_READY_WITH_DEFERRALS.

Create or update the smallest appropriate release-note draft document for v0.5 RC polish. The draft must preserve:
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- READY is planning contract readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- ni run compiles a bounded handoff prompt only.
- Fixture relock is separate from project-root relock.
- Benchmark evidence keeps not_measured boundaries.

Include:
- concise RC summary;
- what changed in v0.5;
- what remains deferred;
- user-facing install/status caveats;
- validation evidence summary;
- explicit non-goals and claim boundaries;
- changed files;
- recommended next task, choosing exactly one.

Run:
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final report must confirm no project-root relock, no protected .ni file changes, no release/tag/publish action, no Homebrew Available claim, no model workspace Available claim, no deterministic no-terminal validation claim, and no benchmark overclaim.
```
