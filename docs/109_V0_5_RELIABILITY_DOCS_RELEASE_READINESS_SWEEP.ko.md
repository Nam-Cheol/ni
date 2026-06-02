# v0.5 Reliability Docs Release Readiness Sweep

## Current status

- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Sweep goal

이 sweep은 다음 v0.5 direction으로 이동하기 전에 v0.5 reliability documentation
set의 consistency, navigation, Korean companion drift, validation surface, and
overclaim boundaries를 확인한다. 이것은 documentation and claim-boundary audit
only이며 lock semantics, stale-lock hash semantics, CLI behavior, runtime
boundaries, implementation을 변경하지 않는다.

## Reliability doc inventory

| Doc | Purpose | Primary boundary | Korean companion? | Linked from roadmap? | Notes |
| --- | --- | --- | --- | --- | --- |
| `docs/101` | Conversation proof capture reliability pass를 기록한다. | Planning proof는 implementation proof가 아니며 CLI가 readiness, lock, prompt compilation을 결정한다. | Yes | Yes | `docs/83`, `docs/97`, `docs/99`와 겹치지만 proof capture와 public status boundaries를 함께 보여주므로 helpful하다. |
| `docs/102` | Stale-lock diagnostic 전 locked-plan change-control UX를 audit한다. | Changed intent는 stale handoff를 stop해야 하며 `ni run`은 bounded prompt compilation only로 남는다. | Yes | Yes | `docs/103`, `docs/104`와의 historical overlap은 later diagnostics와 examples의 이유를 설명하므로 useful하다. |
| `docs/103` | Implemented `LOCK-STALE` diagnostic을 document한다. | `LOCK-STALE`는 existing-lock staleness이지 product 또는 implementation failure가 아니다. | Yes | Yes | `docs/104`, `docs/106`, `docs/107`, `docs/108`로 clear handoff가 있다. |
| `docs/104` | `LOCK-STALE` 이후 practical amend/relock examples를 제공한다. | Recovery order는 `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000`로 유지된다. | Yes | Yes | No-terminal examples와의 overlap은 evidence level별 user workflow가 다르므로 confusing하지 않다. |
| `docs/105` | Model workspace stale-lock wording을 verify한다. | Model workspace packs remain Experimental; skills can draft/explain but cannot validate, lock, relock, or update `.ni/plan.lock.json`. | Yes | Yes, added in this sweep | Navigation gap 발견: roadmap이 주변 reliability docs를 link했지만 이 doc은 빠져 있었다. |
| `docs/106` | No-terminal stale-lock examples를 보여준다. | No-terminal remains Experimental / assisted; readiness, lock freshness, relock, hash verification, bounded prompt compilation claims에는 exact trusted CLI output이 필요하다. | Yes | Yes | `docs/107`과 함께 동작한다; `docs/106`은 scenario-based라 overlap이 helpful하다. |
| `docs/107` | No-terminal transcript quality levels와 copy-paste fields를 정의한다. | Transcript claims는 model-only, pasted, trusted runner, fixture, target-workspace evidence에 scope된다. | Yes | Yes | Fixture-versus-target과 validation-script transcript boundaries를 강하게 보호한다. |
| `docs/108` | Broader changed-intent fixture coverage를 기록한다. | Fixture relock is separate from project-root relock; `.ni/session.json`은 current lock semantics에서 mutable continuity state로 document된다. | Yes | Yes | Lock-input semantics를 바꾸지 않고 tests와 docs를 align한다. |
| `docs/83` | Conversation proof capture와 no-terminal proof limits를 정의한다. | Proof capture는 audit trail이지 readiness authority가 아니다. | Yes | Related roadmap surface | `docs/101`의 upstream reference로 useful하다. |
| `docs/97` | Benchmark claim boundaries를 정의한다. | Benchmark `READY`와 `not_measured` evidence는 planning-artifact scope only이다. | Yes | README and v0.5 roadmap | Benchmark examples와 demo checks의 required guard이다. |
| `docs/99` | Model workspace status를 정의한다. | Host-level install 또는 discovery가 verified되기 전까지 Model workspace packs are Experimental. | Yes | README and this sweep's roadmap edit references `docs/105` for stale-lock wording | Model workspace claims의 canonical status doc이다. |

## Recommended reading order

1. `docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md` - proof capture가
   planning evidence로 무엇을 증명하고 증명하지 않는지 먼저 frame한다.
2. `docs/102_CHANGE_CONTROL_UX_AUDIT.md` - stale-lock diagnostics를 만든
   historical UX risk를 읽는다.
3. `docs/103_STALE_LOCK_DIAGNOSTIC.md` - implemented `LOCK-STALE` behavior와
   recovery wording을 확인한다.
4. `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md` - diagnostic에서 concrete
   recovery examples로 이동한다.
5. `docs/108_CHANGED_INTENT_FIXTURE_COVERAGE.md` - workflow examples 이후
   fixture coverage와 project-root safety를 확인한다.
6. `docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md` - normal CLI recovery path가
   clear한 뒤 no-terminal stale-lock scenarios를 읽는다.
7. `docs/107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md` - trusted runner와
   fixture evidence의 transcript quality rules를 읽는다.
8. `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md` - model
   workspace-specific stale-lock and skill-boundary audit를 읽는다.
9. `docs/99_MODEL_WORKSPACE_STATUS.md` - model workspace status baseline과
   host-verification boundary를 확인한다.
10. `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md` - benchmark and `not_measured`
    boundaries로 evidence claims를 scope하며 마무리한다.

이 순서는 general proof에서 changed-intent mechanics, fixture and no-terminal
evidence, model workspace and benchmark status surfaces로 이동한다.

## Claim boundary audit

| Claim area | Expected wording/status | Observed status | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Release binary | Release binary: Available | README, install docs, roadmap, reliability docs가 verified v0.4.0 assets에 대해서만 release binaries Available을 유지한다. | Yes | `check-install-docs.py`와 `release-check.sh`가 guard한다. |
| Curl installer | Curl installer: Available | README와 installer verification docs가 verified v0.4.0 assets에 대해 curl installer Available을 유지한다. | Yes | Package-manager availability를 imply하지 않는다. |
| Homebrew | Homebrew: Planned / v0.5 candidate | README와 roadmap이 Homebrew Planned를 유지하며 roadmap은 v0.5가 earliest scheduled implementation point라고 말한다. | Yes | Audited status surfaces에서 Homebrew Available claim은 발견되지 않았다. |
| Model workspace packs | Model workspace packs: Experimental | README, package READMEs, `docs/99`, `docs/105`가 broad product status Experimental을 보존한다. | Yes | Host-level/global install과 provider runtime behavior는 `not_verified`로 남는다. |
| No-terminal method | No-terminal method: Experimental / assisted | No-terminal docs, `docs/106`, `docs/107`이 assisted-only wording을 보존한다. | Yes | CLI-state claims에는 exact trusted CLI output이 필요하다. |
| CLI authority | CLI is authority. | Docs and skills에서 `ni status`, `ni end`, `ni run`이 authoritative gates로 남는다. | Yes | Model-readiness authority claim은 발견되지 않았다. |
| Skills role | Skills are UX. | Skill packages와 reliability docs가 "Skills are UX; CLI is authority."를 유지한다. | Yes | Skills may draft/explain; they do not lock, relock, or replace CLI validation. |
| READY | READY is planning artifact readiness only. | Benchmark and reliability docs가 `READY`를 declared planning contract 또는 artifact-readiness scope로 제한한다. | Yes | Product readiness는 out of scope로 남는다. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | `docs/103`, `docs/104`, `docs/106`, `docs/107`, `docs/108`가 이 의미를 사용한다. | Yes | Stale이 current planning `READY`와 coexist할 수 있음도 말한다. |
| ni run | Bounded handoff prompt compilation only; no downstream execution. | README, command docs, reliability docs, skills, examples가 non-execution을 보존한다. | Yes | `ni run` refuses stale handoff and does not relock. |
| fixture relock | Fixture relock is not project-root relock. | `docs/107` and `docs/108` state fixture evidence supports only fixture claims. | Yes | Final reports는 project-root lockfile diffs를 계속 확인해야 한다. |
| trusted runner transcript | Claims require exact CLI output from the shown workspace/command/time. | `docs/106` and `docs/107` preserve model-only, pasted, trusted runner, fixture, and target-workspace levels. | Yes | No-terminal deterministic validation claim은 발견되지 않았다. |
| benchmark evidence | `not_measured` boundaries preserved. | `docs/97`, benchmark examples, `demo-check.sh`가 product, research, adoption, cost, latency, statistical claims를 out of scope로 유지한다. | Yes | Benchmark remains qualitative planning-artifact evidence. |
| kernel non-execution boundary | ni is not an execution harness. | README, roadmap, reliability docs, skill docs, scripts가 no task runner, adapter, queue, PR automation, release automation, downstream execution layer를 보존한다. | Yes | Runtime feature는 추가하지 않았다. |

## Korean companion audit

| Doc pair | Pass? | Notes |
| --- | --- | --- |
| `docs/83` / `.ko.md` | Yes | Companion은 proof-capture lifecycle, CLI authority, no-terminal draft-only boundaries를 보존한다. |
| `docs/97` / `.ko.md` | Yes | Companion은 `READY`, `not_measured`, dashboard, research approval, fieldwork limitations를 scope한다. |
| `docs/99` / `.ko.md` | Yes | Companion은 Experimental model workspace status와 `not_verified` host/provider claims를 보존한다. |
| `docs/100` / `.ko.md` | Yes | Companion은 status vocabulary, GRILL closure, next-direction boundaries를 보존한다. |
| `docs/101` / `.ko.md` | Yes | Companion은 proof capture나 no-terminal validation을 overpromise하지 않는다. |
| `docs/102` / `.ko.md` | Yes | Companion은 change-control and authority boundaries를 보존한다. |
| `docs/103` / `.ko.md` | Yes | Companion은 `LOCK-STALE`, recovery flow, stale-does-not-prove limits를 보존한다. |
| `docs/104` / `.ko.md` | Yes | Companion은 recovery order와 skill/no-terminal limits를 보존한다. |
| `docs/105` / `.ko.md` | Yes | Companion은 Experimental, `not_verified`, skill drafting-only, no-relock wording을 보존한다. |
| `docs/106` / `.ko.md` | Yes | Companion은 Experimental / assisted와 exact trusted CLI output requirements를 보존한다. |
| `docs/107` / `.ko.md` | Yes | Companion은 model-only, pasted CLI output, trusted runner, fixture, target-workspace transcript categories를 보존한다. |
| `docs/108` / `.ko.md` | Yes | Companion은 fixture relock versus project-root relock과 current `.ni/session.json` hash behavior를 보존한다. |
| `docs/109` / `.ko.md` | Yes | 이 sweep에서 추가했다; companion은 commands, paths, status constants, diagnostic labels, exact boundary phrases를 보존한다. |

## Navigation and cross-link audit

| Surface | Expected link or navigation role | Pass? | Notes |
| --- | --- | --- | --- |
| README | Public entry는 install, no-terminal, model workspace, benchmark, Homebrew, command surfaces를 crowding 없이 link해야 한다. | Yes | README link는 추가하지 않았다; roadmap이 더 좋은 reliability-doc index다. |
| Roadmap v0.5 section | Reliability docs 101-109와 status docs를 excessive cross-links 없이 index해야 한다. | Yes | `docs/105`와 this `docs/109` sweep에 대한 narrow links를 추가했다. |
| `docs/103` | Diagnostic에서 examples, no-terminal, transcript, fixture coverage로 point해야 한다. | Yes | Existing links are sufficient. |
| `docs/104` | Amend/relock examples에서 no-terminal, transcript, fixture coverage로 point해야 한다. | Yes | Existing links are sufficient. |
| `docs/106` | No-terminal examples에서 transcript quality로 point해야 한다. | Yes | Existing link to `docs/107` is sufficient. |
| `docs/107` | Transcript quality에서 changed-intent fixture coverage로 point해야 한다. | Yes | Existing link to `docs/108` is sufficient. |
| `docs/108` | Next release-readiness sweep를 identify해야 한다. | Yes | Existing recommended next task가 this sweep를 name하며 roadmap이 now links this doc. |

## Validation surface

| Script/check | Current protection | Gap | Change made |
| --- | --- | --- | --- |
| `check-install-docs.py` | Install/distribution status rows, release/curl/Homebrew/model workspace markers, forbidden availability claims를 enforce한다. | Selected stable phrases만 check하며 모든 semantic overclaim은 check하지 않는다. | None. Current surface is stable and low-noise. |
| `check-skill-packs.sh` | Skill pack files, Experimental status, CLI authority, `LOCK-STALE`, no-relock, no-lockfile-update, recovery-order wording을 verify한다. | 모든 skill sentence를 parse하지는 않는다. | None. Current exact phrase checks are appropriate. |
| `demo-check.sh` | Benchmark docs, no-terminal docs, transcript quality, fixture relock wording, ni-grill docs-only boundaries, seed-only exports를 verify한다. | Broad하지만 example-driven이다; future docs can drift outside checked markers. | None. No new risky phrase justified another check. |
| `quality.sh` | JSON/schema/markdown/formatting/readme/install/skill/prompt-budget/core-boundary/asset checks, `gofmt -w .`, `go test ./...`, `smoke.sh`를 run한다. | Go가 touched되면 broad wrapper가 mechanically format할 수 있다. | None. No Go files touched. |
| `smoke.sh` | `ni`를 build하고 public commands, readiness, lock, prompt compilation, exports, amendment, relock, seed boundaries를 temporary workspaces에서 exercise한다. | Temporary fixture relocks must not be reported as project-root relock. | None. Reporting boundary is documented in this sweep. |
| `install-check.sh` | Source, build, local binary, temporary local install paths를 verify한다. | Homebrew 또는 global model workspace install을 verify하지 않는다. | None. Homebrew remains Planned; host-level model workspace behavior remains unverified. |
| `release-check.sh` | Release-readiness docs, release facts, release pipeline markers, benchmark boundaries, examples, Go tests, smoke, demo, install, status proof를 verify한다. | v0.5 release-candidate decision을 만들지는 않는다. | None. This sweep selects a release-candidate readiness audit as the next task. |

## Findings

| Finding | Severity | Surface | Change made | Blocks v0.5? |
| --- | --- | --- | --- | --- |
| Audited v0.5 reliability docs 전체에서 material claim-boundary contradiction은 발견되지 않았다. | Low | `docs/101` through `docs/108`, related docs, README, packages, examples, scripts | 이 sweep document를 추가해 audit을 기록했다. | No. |
| Roadmap이 주변 v0.5 reliability docs는 link했지만 `docs/105`는 link하지 않았다. | Low | `docs/51_POST_RELEASE_ROADMAP.md`, `.ko.md` | Model workspace stale-lock wording verification에 대한 narrow roadmap pointer를 추가했다. | No. |
| Roadmap에 new release-readiness sweep pointer가 필요했다. | Low | `docs/51_POST_RELEASE_ROADMAP.md`, `.ko.md` | This `docs/109` sweep에 대한 narrow roadmap pointer를 추가했다. | No. |

v0.5를 block하는 material findings는 없다. Findings는 navigation and audit-record
improvements only이다.

## Changes made

- `docs/109_V0_5_RELIABILITY_DOCS_RELEASE_READINESS_SWEEP.md`: English
  release-readiness sweep를 추가했다.
- `docs/109_V0_5_RELIABILITY_DOCS_RELEASE_READINESS_SWEEP.ko.md`: same claim
  boundaries를 가진 Korean companion을 추가했다.
- `docs/51_POST_RELEASE_ROADMAP.md`: `docs/105`와 this sweep에 대한 narrow links를
  추가했다.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean roadmap links를
  추가했다.

Go files, skill docs, static checks, root lockfiles, runtime behavior, release
actions는 변경하지 않았다.

## What this sweep proves

- v0.5 reliability docs are internally aligned for audited surfaces.
- current public status boundaries are preserved.
- Korean companions do not overpromise for audited pairs.
- validation scripts cover selected stable boundary phrases.

## What this sweep does not prove

- implementation correctness
- downstream execution success
- product readiness
- benchmark effect size
- adoption/cost/latency improvement
- global model workspace behavior
- host-level model workspace verification
- Homebrew availability
- no-terminal deterministic validation
- exhaustive changed-intent coverage beyond existing fixtures

## Remaining risks

- docs may drift after future edits
- static checks cover selected phrases, not all semantic overclaims
- provider host behavior remains unverified
- Homebrew remains Planned until tested
- external user validation remains limited
- v0.5 release readiness still needs a final release-candidate check if a
  release is planned

## Recommended next task

Selected next task: E. v0.5 release candidate readiness audit.

Why: reliability docs now have an explicit sweep record and only narrow
navigation findings. 다른 adoption 또는 implementation lane을 시작하기 전에 v0.5가
release-candidate ready인지, 무엇이 blocking인지, 어떤 status claims를 RC plan에
overclaim 없이 가져갈 수 있는지 결정해야 한다. Homebrew, model workspace host
behavior, no-terminal validation, benchmark impact는 계속 scope를 지켜야 한다.

## Next task prompt

```text
Proceed with the v0.5 release candidate readiness audit in /Users/namba/Documents/project/ni.

This is an audit and documentation task only. Do not publish, tag, release, relock the project root, run ni end on the project root, execute generated prompts, add downstream execution behavior, add adapters, add queues, add PR automation, add release automation, mark Homebrew Available, mark model workspace packs Available, or claim no-terminal deterministic validation.

Goal:
Decide whether ni v0.5 is ready to become a release candidate, based on current docs, status claims, validation scripts, examples, and v0.5 reliability evidence.

Context:
ni is ni-kernel: a pre-runtime Project Intent Compiler for AI Agents.
Core protocol:
conversation -> project contract -> readiness gate -> lock hash -> downstream handoff.
Core flow:
ni init -> planning conversation -> docs/plan/** + .ni/contract.json + .ni/session.json -> ni status --proof --next-questions -> ni end -> .ni/plan.lock.json -> ni run --max-chars 4000 -> bounded downstream handoff prompt.

Current status boundaries:
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni run compiles a bounded handoff prompt and does not execute downstream work.
- READY is planning artifact readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- Fixture relock is not project-root relock.
- Benchmark evidence preserves not_measured boundaries.

Scope:
Audit README.md, README.ko.md, docs/51_POST_RELEASE_ROADMAP.md, docs/51_POST_RELEASE_ROADMAP.ko.md, docs/95 through docs/109 and Korean companions, docs/97, docs/99, docs/no-terminal*, examples/benchmark-report/, examples/no-terminal-assisted/, examples/ni-grill/, packages/claude-skills/, packages/codex-skills/, .agents/skills/, and validation scripts.

Required deliverable:
Add docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md and docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md if Korean companions are maintained.

The audit must include:
- current factual status table;
- release-candidate readiness criteria;
- evidence inventory for v0.5 reliability docs and examples;
- claim-boundary audit for release binary, curl installer, Homebrew, model workspace packs, no-terminal, CLI authority, skills role, LOCK-STALE, ni run, fixture relock, benchmark evidence, and the kernel non-execution boundary;
- validation command results;
- blockers, deferrals, warnings, and non-blocking risks;
- explicit decision: RC-ready, RC-ready with deferrals, or not RC-ready;
- changed files and why;
- selected next task after the audit, choosing exactly one.

Run:
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- bash scripts/quality.sh
- bash scripts/smoke.sh
- bash scripts/install-check.sh
- bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

If Go files are touched, also run:
- gofmt -w .
- GOCACHE=/private/tmp/ni-go-cache go test ./...

Final report must confirm:
- Go files touched or not;
- docs touched or not;
- skill docs touched or not;
- static checks touched or not;
- Homebrew remains Planned / v0.5 candidate;
- Model workspace packs remain Experimental;
- No-terminal method remains Experimental / assisted;
- Skills are UX; CLI is authority is preserved;
- ni run remains bounded prompt compilation only;
- .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified on the project root;
- ni end was not run on the project root;
- no relock was run on the project root;
- no fixture relock was claimed as project-root relock;
- no downstream execution behavior, shell adapter, Codex exec adapter, queue, PR automation, release action, Homebrew Available claim, model workspace Available claim, no-terminal deterministic claim, or benchmark overclaim was added.
```
