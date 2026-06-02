# Change-Control UX Audit

## Current status

이 audit은 stale locks, amended intent, relock expectations, downstream handoff
safety 주변의 ni locked-plan change-control user experience를 검토한다.

Follow-up: Task 183은 [`103_STALE_LOCK_DIAGNOSTIC.ko.md`](103_STALE_LOCK_DIAGNOSTIC.ko.md)에
문서화된 focused `LOCK-STALE` diagnostic을 구현했다. 이 audit은 해당 diagnostic의
동기가 된 historical finding으로 남는다.

현재 factual boundaries는 다음과 같다.

- release binary: verified v0.4.0 assets 기준 Available;
- curl installer: verified v0.4.0 assets 기준 Available;
- Homebrew: Planned / v0.5 candidate only;
- model workspace packs: broad product path로 Experimental;
- no-terminal method: Experimental / assisted;
- `ni-kernel`: pre-runtime Project Intent Compiler;
- runtime execution, shell adapters, Codex exec adapters, queues, PR
  automation, release automation, downstream execution layers: 포함하지 않음.

이 repository의 현재 CLI readiness proof는 다음과 같다.

```text
NI Intent Readiness: READY
Blockers: None.
Deferrals: None.
Warnings: None.
```

이 audit은 `ni end`, relock flow, `.ni/contract.json` edit, `.ni/session.json`
edit, `.ni/plan.lock.json` edit, generated prompt execution, runtime behavior
추가를 하지 않았다.

## Intended change-control model

`.ni/plan.lock.json`이 존재한 뒤 intended source-of-truth order는 다음과 같다.

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

Locked plan은 intent가 바뀌기 전까지 downstream handoff의 authority다. Lock은
`.ni/contract.json`과 required `docs/plan/**` files의 hashes를 기록한다.
`.ni/session.json`은 locked docs 아래의 mutable continuity state이므로 hash되지
않는다.

Lock 이후 `.ni/contract.json` 또는 required `docs/plan/**`가 바뀌면 old lock은 더
이상 current intent를 나타내지 않을 수 있다. Valid lock을 요구하는 handoff
surfaces는 stale hashes를 refuse하고 `BLOCKED`를 report해야 한다. 그 다음 user는
change를 review하고, amend 또는 re-plan하고, readiness를 다시 실행한 뒤, CLI를
통해 relock한 다음 new downstream handoff를 만들어야 한다.

Skills are UX; CLI is authority. Skills는 amended planning text draft를 도울 수
있지만 readiness를 determine하지 않고, plans를 lock/relock하지 않고, `ni status`,
`ni end`, `ni run`을 replace하지 않고, `.ni/plan.lock.json`을 update하지 않는다.

## User-facing lifecycle

```text
pre-lock planning
-> readiness check
-> lock
-> downstream handoff from locked plan
-> intent changes
-> stale/change risk is surfaced
-> user reviews change
-> amend or re-plan
-> readiness check again
-> relock
-> new downstream handoff
```

이 lifecycle에서 `ni status`는 readiness authority, `ni end`는 first-lock
authority, `ni relock`은 relock authority, `ni run`은 bounded prompt compiler다.
`ni run`은 downstream work를 execute하지 않는다.

## Current behavior audit

| Surface | Current behavior / wording | Pass? | Risk | Recommendation |
| --- | --- | --- | --- | --- |
| README | `ni status`, `ni end`, `ni run`을 gates로 설명하고, `ni run`이 valid locked plan에서 compile하며 execute하지 않는다고 말한다. Post-lock intent-change lifecycle은 깊게 설명하지 않는다. | Partial | Reader가 safe handoff는 이해하지만 changed intent 이후 amend/relock path를 모를 수 있다. | README는 간결하게 유지하고, diagnostic work가 landing된 뒤 deeper change-control guidance를 link한다. |
| docs/83 | Proof capture, CLI authority, no-terminal draft limits를 정의하고 stale side는 repair하거나 blocking으로 유지해야 한다고 말한다. `ni run`이 lock hashes를 verify한다고 말한다. | Yes | Low. Proof-capture 중심 문서이지 full amend/relock guide는 아니다. | As-is로 유지하고, stale-lock wording이 확장되면 이 audit 또는 future examples로 연결한다. |
| docs/101 | Planning proof와 execution evidence를 분리하고, stale lock risk와 validation surfaces를 문서화한다. | Yes | Low. Risk는 문서화하지만 detailed CLI examples는 없다. | As-is로 유지하고 이 audit을 change-control follow-up으로 사용한다. |
| ni status --proof --next-questions | Docs, contract, sync, decisions, risks, evaluations, blocker questions에서 readiness를 authority 있게 report한다. Current implementation은 existing `.ni/plan.lock.json` hashes를 verify하지 않는다. | Partial | User가 post-lock docs/contract edits 뒤에도 `READY`를 보고, lock-dependent command 전까지 existing lock이 stale임을 모를 수 있다. | Readiness semantics를 성급하게 바꾸지 말고 focused stale-lock diagnostic 또는 warning surface를 추가한다. |
| ni end | Readiness를 실행하고 `BLOCKED`가 아니면 `.ni/plan.lock.json`을 쓴다. Output은 lock path와 status로 짧다. Existing code는 first lock과 existing lock overwrite를 visible하게 구분하지 않는다. | Partial | User가 `ni end`를 accidental relock path로 사용해 amendment/relock expectations를 우회할 수 있다. | CLI authority를 보존하면서 existing-lock cases 주변의 follow-up diagnostic/refusal 또는 guidance를 추가한다. |
| ni run | `.ni/plan.lock.json`을 요구하고 locked hashes를 verify하며, stale hashes에서 `BLOCKED: lock hash mismatch for <path>`로 refuse하고 bounded prompt만 compile한다. | Yes | Stale error는 맞지만 terse하다. User에게 review, amend, readiness, relock recovery를 말하지 않는다. | Stale-lock refusal 근처에 clearer changed-intent guidance를 추가한다. |
| skill packs | Codex/Claude pack READMEs와 skills가 Experimental status, "Skills are UX; CLI is authority.", no manual lock editing, no readiness by model judgment, stale-lock `BLOCKED` behavior를 보존한다. | Yes | Low. Individual skill files are aligned. | Existing static checks를 유지하고, new amend/relock wording이 생길 때만 update한다. |
| no-terminal assisted workflow | Assisted drafting only로 문서화되어 있다. Draft proof는 deterministic validation이 아니며 readiness, lock, hash proof, prompt compile에는 trusted CLI proof가 필요하다고 말한다. | Yes | Next CLI step이 visible하지 않으면 user가 model summary를 relock처럼 취급할 수 있다. | Draft-only labels를 유지하고 future no-terminal relock guidance가 written된 뒤 examples를 추가한다. |
| examples | Benchmark, ni-grill, no-terminal, resolution-path examples는 대체로 no-execution, isolated-workspace, `not_measured`, amendment/relock stop wording을 보존한다. | Yes | Fixture relocks 또는 benchmark locks가 context 없이 인용되면 project-root relock으로 오해될 수 있다. | Examples 근처에 "isolated workspace only"와 "root not relocked" labels를 유지한다. |
| validation scripts | `check-install-docs.py`, `check-skill-packs.sh`, `demo-check.sh`, `quality.sh`, `smoke.sh`, `install-check.sh`, `release-check.sh`가 status, boundary, skill, install, demo, release surfaces를 보호한다. | Yes | 모든 stale-lock UX copy를 현재 enforce하지 않고 broad prose는 manual audit이 필요하다. | 이 audit에서는 broad brittle check를 추가하지 않는다. Target diagnostic wording이 stable해진 뒤 focused checks만 추가한다. |

## Stale-lock and changed-intent risks

| Risk | Severity | Current mitigation | Gap | Recommended next action |
| --- | --- | --- | --- | --- |
| User가 lock 이후 `docs/plan/**`를 edit하고 old lock이 current intent를 여전히 반영한다고 가정한다. | High | `ni run`, export, graph, harness, feedback, pressure paths는 valid lock이 필요할 때 lock hashes를 verify한다. | `ni status`는 existing-lock staleness를 surfacing하지 않고 readiness를 report할 수 있다. | Stale-lock warning/diagnostic 추가. |
| User가 lock 이후 `.ni/contract.json`을 바꾸고 mismatch를 눈치채지 못한 채 `ni run`을 실행한다. | High | `ni run`은 lock hash mismatch를 `BLOCKED`로 refuse한다. | Error text가 terse하고 amendment/relock recovery를 설명하지 않는다. | Clearer stale-lock recovery wording 추가. |
| Model workspace skill이 new planning text를 draft했는데 user가 CLI state가 바뀌었다고 가정한다. | Medium | Skill docs는 skills are UX, CLI is authority, skills do not lock or determine readiness라고 말한다. | Every turn에서 proof wording이 visible하지 않으면 user summaries가 overstate할 수 있다. | Static skill checks와 proof-capture language 유지. |
| No-terminal assisted workflow가 deterministic relock으로 오해된다. | Medium | No-terminal docs는 draft proof가 deterministic validation이 아니라고 말한다. | Detailed no-terminal amend/relock example이 없다. | CLI diagnostics가 더 clear해진 뒤 examples 검토. |
| `ni-run` handoff가 fresh validation으로 오해된다. | Medium | `ni run`은 hashes를 verify하고 prompt compilation only라고 말한다. | Readiness를 다시 실행하지 않으며 compilation이 fresh planning acceptance가 아님을 자세히 설명하지 않는다. | Future handoff wording에서 clarify. |
| Proof text가 implementation evidence로 오해된다. | Medium | docs/83과 docs/101이 planning proof와 execution evidence를 분리한다. | New examples에는 manual audit이 필요하다. | Demo/skill checks는 lightweight하게 유지. |
| Benchmark evidence가 downstream execution success로 오해된다. | Medium | docs/97과 benchmark examples가 `not_measured`, no-execution, artifact-readiness labels를 보존한다. | Long examples에서는 boundaries가 묻힐 수 있다. | `READY` rows 근처에 boundary labels 유지. |
| Root lockfile이 manually edited된다. | High | Skills와 docs는 manual `.ni/plan.lock.json` edit을 금지한다. Lock hash checks가 많은 downstream uses를 catch한다. | Lockfile 자체 manual edits는 command가 load/verify하기 전까지 obvious하지 않을 수 있다. | Future lock validation work에서 targeted diagnostics 추가. |
| Validation scripts가 fixture relocks를 exercise하고 project-root relock으로 오해된다. | Medium | Prior audits와 benchmark docs가 isolated workspaces와 root no-relock boundaries를 label한다. | Final reports는 distinction을 explicit하게 유지해야 한다. | Root lockfile diff와 command scope를 계속 report한다. |

## What must remain CLI-authoritative

- Readiness is determined by CLI.
- Locking is performed by CLI.
- Prompt compilation is performed by CLI.
- Skills do not validate readiness.
- Skills do not lock or relock.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- Skills do not update `.ni/plan.lock.json`.
- No-terminal workflows do not provide deterministic validation.
- `ni run` does not execute downstream work.

## Say this / do not say this

| Say this | Do not say this |
| --- | --- |
| "The locked plan is the authority for downstream handoff until intent changes." | "The model relocked the plan." |
| "If intent changes after lock, run the CLI readiness and lock flow again before generating a new handoff." | "The no-terminal workflow deterministically validated the amended plan." |
| "Skills can help draft amended planning text, but the CLI remains the authority." | "`ni run` verified the implementation." |
| "`ni run` compiles a bounded handoff prompt from the locked plan; it does not execute the work." | "The benchmark proves downstream execution quality." |
| "A stale lock or hash mismatch is `BLOCKED` until the plan is reviewed, amended if needed, checked for readiness, and relocked." | "The skill pack replaces CLI validation." |

## Audit findings

| Finding | Severity | Surface | Evidence | Recommendation | Blocks v0.5? |
| --- | --- | --- | --- | --- | --- |
| `ni status`가 stale existing-lock state를 surface하지 않는다. | High | `ni status --proof --next-questions` | `cmd/ni/main.go`는 readiness를 evaluate하지만 lock verification을 call하지 않는다. `docs/commands.md`도 `ni status`가 lock hashes를 verify하지 않는다고 말한다. | Users가 handoff 전에 무엇이 trustworthy한지 볼 수 있도록 focused stale-lock warning 또는 diagnostic을 추가한다. | No, but should be next. |
| Existing-lock `ni end` behavior가 relock expectations와 명확히 분리되지 않는다. | High | `ni end`; `ni relock`; docs | `runEnd`는 `lock.Create`로 write한다. `runRelock`은 amendment-gated relock을 implement한다. User-facing `ni end` output은 terse하다. | Existing lock이 있을 때 `ni end`가 warn/refuse해야 하는지 결정하고 recovery를 document한다. | No, but should be addressed before stronger change-control claims. |
| Stale `ni run` refusal은 correct하지만 terse하다. | Medium | `ni run` | `internal/core/prompt/prompt.go`는 `BLOCKED: lock hash mismatch for <path>`를 반환한다. | User-facing recovery wording을 추가한다: changed intent review, amend/re-plan, readiness, relock, new handoff. | No. |
| Docs와 skills는 authority boundaries를 보존한다. | Low | docs/83, docs/101, skill packs, no-terminal docs | Audited docs가 CLI authority, draft-only no-terminal proof, no downstream execution을 반복한다. | Current wording을 유지하고 next diagnostic text가 stable해질 때까지 broad static checks를 피한다. | No. |

Material findings는 있지만, stale hashes를 downstream handoff가 무시한다는 evidence는
아니다. 가장 강한 current protection은 lock-dependent handoff commands가 hash
mismatches를 refuse한다는 점이다. 가장 약한 user-facing surface는 `ni status`가
healthy해 보여도 existing lock은 stale일 수 있다는 점이다.

## Validation surface

Current scripts는 다음을 check한다.

- `go run ./cmd/ni status --dir . --proof --next-questions`: current repository
  readiness through CLI;
- `python3 scripts/check-install-docs.py`: install, distribution, Homebrew, model
  workspace status claim boundaries;
- `bash scripts/check-skill-packs.sh`: skill metadata, Experimental status, CLI
  authority wording, proof-capture markers, stale-lock boundary markers in skill
  packs;
- `bash scripts/quality.sh`: broad static docs checks, Go formatting/tests, smoke
  checks;
- `bash scripts/demo-check.sh`: benchmark, no-terminal, ni-grill, seed-only
  boundaries;
- `bash scripts/smoke.sh`: source CLI smoke behavior;
- `bash scripts/install-check.sh`: local install behavior;
- `bash scripts/release-check.sh`: release readiness surfaces.

Stale-lock prose, amend/relock examples, `ni status` versus lock-verification
wording, future change-control diagnostic text에는 manual audit이 여전히 필요하다.

이 task에서는 new static check를 추가하지 않았다. Current gap은 missing phrase만이
아니라 user-facing diagnostic design decision이기 때문이다.

## Recommended next task

Selected next action: A. implement stale-lock warning/diagnostic.

Why: 이 audit은 lock-dependent handoff commands가 stale hashes를 refuse한다는 점을
확인했다. 하지만 `ni status --proof --next-questions`는 existing
`.ni/plan.lock.json`이 current contract/docs와 여전히 match하는지 알려주지 않은
채 readiness를 report할 수 있다. Focused diagnostic이 downstream handoff 전에
무엇이 trustworthy하고, 무엇이 stale이고, 무엇을 amend/relock해야 하는지 직접
답하는 가장 작은 next step이다.

## Next task prompt

```text
Goal:
Implement a focused stale-lock warning/diagnostic for ni change-control UX.

This is a small v0.5 reliability implementation task. Do not add runtime execution, adapters, queues, PR automation, release automation, or downstream execution behavior.

Context:
ni is ni-kernel: a pre-runtime Project Intent Compiler. The locked plan is the authority for downstream handoff until intent changes. After .ni/plan.lock.json exists, source-of-truth order is:
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history

Task 182 added:
- docs/102_CHANGE_CONTROL_UX_AUDIT.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.ko.md

The audit found:
- lock-dependent handoff commands already refuse stale lock hashes;
- ni status evaluates readiness but does not surface whether an existing lock is stale;
- ni end output/behavior does not clearly distinguish first lock from relock expectations;
- ni run stale-lock refusal is correct but terse.

Read first:
- AGENTS.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.ko.md
- docs/42_INTENT_LOCK_PROTOCOL.md
- docs/04_LOCKFILE.md
- docs/36_NI_RUN_HANDOFF.md
- docs/commands.md
- cmd/ni/main.go
- internal/core/lock/lock.go
- internal/core/prompt/prompt.go
- internal/core/amendment/amendment.go
- internal/core/prompt/prompt_test.go
- internal/core/lock/lock_test.go

Before editing, run:
- go run ./cmd/ni status --dir . --proof --next-questions

Implement the smallest coherent diagnostic that tells users:
- whether an existing .ni/plan.lock.json matches current .ni/contract.json and required docs/plan/**;
- which path mismatches first when the lock is stale;
- that stale lock means downstream handoff must stop;
- that the user should review changed intent, amend or re-plan as needed, run readiness again, relock through the CLI, then generate a new handoff;
- that .ni/session.json remains below locked docs and is not hashed.

Preferred implementation direction:
- Add lock-state reporting to ni status --proof and/or --next-questions without changing readiness rule semantics unless tests show a deterministic blocker is required.
- Keep JSON output structured if lock diagnostics are added to status results.
- Improve stale ni run error/help text only if it stays narrow and testable.
- Decide whether ni end should warn or refuse when an existing lock is present; if semantics change, update docs and tests. Do not silently overwrite an existing project-root lock in tests.

Do not:
- run ni end on the project root;
- run relock on the project root;
- manually edit .ni/plan.lock.json;
- edit .ni/contract.json or .ni/session.json unless the implementation genuinely requires fixture updates outside the project root;
- execute generated prompts;
- add task-runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, execution evidence loop, TUI, or web UI behavior;
- mark Homebrew Available;
- mark model workspace packs Available as a broad product path;
- claim no-terminal deterministic validation;
- claim benchmark evidence proves downstream execution quality.

Validation:
- gofmt -w . if Go files are touched
- go test ./...
- go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/quality.sh
- bash scripts/demo-check.sh
- bash scripts/smoke.sh
- bash scripts/install-check.sh
- bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- whether Go files were touched
- readiness result
- stale-lock diagnostic behavior added
- validation results
- confirmation that .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified
- confirmation that ni end and relock were not run on the project root
- confirmation that no runtime execution, shell adapter, Codex exec adapter, queue, PR automation, release action, Homebrew Available claim, model workspace Available claim, no-terminal deterministic claim, or benchmark overclaim was added
```
