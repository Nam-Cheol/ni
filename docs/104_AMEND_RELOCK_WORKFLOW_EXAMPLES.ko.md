# Amend / Relock Workflow Examples

## Current status

ni는 stale existing locks를 `LOCK-STALE` diagnostic으로 surface한다.
개념적으로 `LOCK-STALE`는 existing lock is stale, 즉 current planning inputs가
`.ni/plan.lock.json`과 다르다는 뜻이다.

이 문서는 lock semantics를 바꾸지 않는다. `ni status --proof --next-questions`는
readiness와 stale-lock diagnostics를 surface하고, `ni end`는 이 recovery flow에서
CLI-authoritative relock step이며, `ni run`은 lock이 current인 뒤에만 bounded
handoff prompt를 compile한다. Skills are UX; CLI is authority.

## Core recovery flow

```text
review changed intent
-> ni status --proof --next-questions
-> ni end
-> ni run --max-chars 4000
```

| Step | Purpose | Boundary |
| --- | --- | --- |
| Review changed intent | `.ni/contract.json`, required `docs/plan/**`, change를 만든 conversation을 inspect한다. | Ambiguous 또는 conflicting edits를 accepted decisions로 취급하지 않는다. |
| `ni status --proof --next-questions` | CLI가 readiness, blockers, deferrals, warnings, proof, next questions를 report하게 한다. | Model judgment는 readiness를 determine하지 않는다. |
| `ni end` | Readiness가 허용한 뒤 CLI가 current lock을 write하게 한다. | `.ni/plan.lock.json`을 manually edit하지 않고, skill이 plan을 relock했다고 말하지 않는다. |
| `ni run --max-chars 4000` | Current lock에서 bounded downstream handoff prompt를 compile한다. | `ni run`은 downstream work를 execute하지 않는다. |

Relock 전에는 `ni run`이 stale handoff를 refuse해야 한다. Relock 뒤에는 `ni run`이
current lock에서 compile해야 하며 prompt compiler only로 남아야 한다.

## Amend vs re-plan vs new project

| Change type | Meaning | Example | Expected action | Can reuse existing project? |
| --- | --- | --- | --- | --- |
| Amend | 같은 accepted project intent 안에서 review하고 relock할 수 있는 scoped change. | 더 엄격한 wording constraint를 추가하거나 acceptance phrase를 clarify한다. | Changed intent를 review하고 planning state를 update한 뒤 core recovery flow를 실행한다. | Yes. |
| Re-plan | Relock 전에 planning questions를 다시 열어야 하는 material change. | New primary capability, actor, requirement, risk, delivery constraint를 추가한다. | Questions를 다시 열고 blockers를 resolve한 뒤 readiness가 허용할 때만 core recovery flow를 실행한다. | Usually yes, planning이 update된 뒤. |
| New project | 같은 locked intent로 취급하면 안 되는 change. | Planning compiler를 downstream work를 실행하는 tool로 바꾸거나 unrelated product로 이동한다. | Old intent를 relock하지 말고 new project intent를 시작한다. | No. |

## Example 1: Minor scope clarification after lock

| Field | Example |
| --- | --- |
| Scenario | Locked plan이 "write install docs"라고 말하고, user가 source, release binary, curl installer paths를 분리해야 한다고 clarify한다. |
| What changed | New capability를 추가하지 않고 scope wording이 더 precise해졌다. |
| Why the existing lock may be stale | Required `docs/plan/**` 또는 `.ni/contract.json`이 이제 `.ni/plan.lock.json`과 다르게 hash될 수 있다. |
| What the user should inspect | Changed requirement text, linked evaluations, non-goals, affected install-status claims. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Existing lock이 current planning inputs와 더 이상 match하지 않으면 stale handoff를 refuse한다. |
| What `ni run` should do after relock | Current lock에서 bounded handoff prompt를 compile하고 execute하지 않는다. |
| What skills may help with | Clearer planning text를 draft하고 `LOCK-STALE` warning을 explain한다. |
| What skills must not claim | Skills must not determine readiness, lock or relock, replace `ni status`, `ni end`, or `ni run`, or update `.ni/plan.lock.json`. |

## Example 2: Requirement or capability change after lock

| Field | Example |
| --- | --- |
| Scenario | Prompt compilation용 locked plan에 target seed conformance notes capability가 추가된다. |
| What changed | Capability 또는 requirement가 바뀌었으므로 evaluations, risks, artifacts traceability update가 필요할 수 있다. |
| Why the existing lock may be stale | Lock은 changed capability graph가 아니라 old contract와 required planning docs를 기록했다. |
| What the user should inspect | Capability IDs, requirement links, evaluation coverage, high-severity risks, non-goals, downstream seed boundaries. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Stale handoff를 refuse한다. Stale lock은 trusted downstream instructions를 만들면 안 된다. |
| What `ni run` should do after relock | Accepted current lock에서만 bounded prompt를 compile한다. |
| What skills may help with | Amended capability wording을 draft하고 missing evaluation links를 identify한다. |
| What skills must not claim | Skills must not weaken acceptance criteria to reach readiness or say the model relocked the plan. |

## Example 3: Deferral changes after lock

| Field | Example |
| --- | --- |
| Scenario | Deferred Homebrew task가 audit 대상으로 reprioritized되지만, Homebrew는 implementation과 install proof 전까지 `Planned`로 남아야 한다. |
| What changed | Deferral, follow-up, sequencing decision이 바뀌었다. |
| Why the existing lock may be stale | Accepted plan의 deferred work, roadmap pointers, decision log가 lock과 더 이상 match하지 않을 수 있다. |
| What the user should inspect | Deferral records, roadmap links, accepted decisions, `Available` / `Experimental` / `Planned` status words, non-goals. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Deferral change가 lockable inputs를 바꾸었다면 stale handoff를 refuse한다. |
| What `ni run` should do after relock | Deferred item이 implemented되었다고 claim하지 않고 handoff prompt를 compile한다. |
| What skills may help with | 어떤 deferral wording이 바뀌었는지 explain하고 claim boundaries를 preserve한다. |
| What skills must not claim | Skills must not mark Homebrew `Available`, mark model workspace packs `Available`, or replace CLI readiness. |

## Example 4: `docs/plan/**` changed after lock

| Field | Example |
| --- | --- |
| Scenario | `.ni/plan.lock.json`이 존재한 뒤 user가 `docs/plan/07_evaluation_contract.md`를 edit한다. |
| What changed | Required planning doc이 lock 이후 바뀌었다. |
| Why the existing lock may be stale | Required planning docs는 lockable inputs이므로 doc hash가 recorded lock hash와 다를 수 있다. |
| What the user should inspect | Exact doc diff, matching `.ni/contract.json` updates, change가 accepted/draft/deferred/conflicting 중 무엇인지. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Stale handoff를 refuse하고 stop-before-execution boundary를 preserve한다. |
| What `ni run` should do after relock | CLI가 plan state를 accepted한 경우에만 new current lock에서 compile한다. |
| What skills may help with | Diff를 summarize하고 synchronized docs/contract updates를 draft한다. |
| What skills must not claim | Skills must not silently edit locked planning docs or say proof text proves implementation correctness. |

## Example 5: Model workspace skill drafts amendment after lock

| Field | Example |
| --- | --- |
| Scenario | Lock이 존재한 뒤 `ni-start`가 model workspace에서 revised planning text를 draft한다. |
| What changed | Model이 amended planning text를 draft했지만 CLI로 check되기 전에는 state가 authoritative하지 않다. |
| Why the existing lock may be stale | Draft가 `.ni/contract.json` 또는 required `docs/plan/**`에 accepted되면 current planning inputs가 `.ni/plan.lock.json`과 달라질 수 있다. |
| What the user should inspect | Draft vs accepted records, assumptions, open questions, docs/contract sync, any `LOCK-STALE` warning. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Stale handoff를 refuse한다. Skill draft는 lock이 아니다. |
| What `ni run` should do after relock | Current lock에서 bounded prompt를 compile하고 downstream work를 execute하지 않는다. |
| What skills may help with | Amended planning text를 draft하고 `LOCK-STALE`를 explain한다. |
| What skills must not claim | Skills do not determine readiness, do not lock or relock, do not replace `ni status`, `ni end`, or `ni run`, and do not update `.ni/plan.lock.json`. |

## Example 6: No-terminal assisted workflow hits stale lock

| Field | Example |
| --- | --- |
| Scenario | No-terminal assisted user가 planning changed라는 model summary를 받고, trusted runner가 나중에 `LOCK-STALE`를 report한다. |
| What changed | Assisted drafting이 planning text를 바꾸거나 changed intent를 identify했다. |
| Why the existing lock may be stale | Current lockable planning inputs가 `.ni/plan.lock.json`과 다르다. |
| What the user should inspect | Assisted draft, trusted runner의 CLI output, changed docs/contract records, blocker questions. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Stale handoff를 refuse한다. No-terminal workflows do not provide deterministic validation. |
| What `ni run` should do after relock | CLI lock이 current인 뒤에만 bounded handoff prompt를 compile한다. |
| What skills may help with | Warning을 explain하고 review용 draft planning amendments를 prepare한다. |
| What skills must not claim | No-terminal assistance must not claim deterministic validation, lock freshness, or implementation proof. |

## Say this / do not say this

| Say this | Do not say this |
| --- | --- |
| "LOCK-STALE means the existing lock no longer matches current planning inputs." | "The model relocked the plan." |
| "Review the changed intent, run `ni status --proof --next-questions`, then run `ni end` before generating a new `ni run` handoff." | "The skill updated `.ni/plan.lock.json`." |
| "Skills can help draft amended planning text, but the CLI remains the authority." | "The no-terminal workflow deterministically validated the amended plan." |
| "`ni run` compiles a bounded handoff prompt from the current lock; it does not execute downstream work." | "`ni run` verified the implementation." |
| "`LOCK-STALE` does not prove implementation failure, benchmark failure, or product readiness." | "The benchmark proves downstream execution quality." |

## Validation and current coverage

Current tests/checks cover:

- stale status warning;
- current lock no warning;
- no lock no warning;
- `ni run` stale refusal;
- fixture relock recovery.

다음은 manual audit으로 남는다:

- user interpretation of amend vs re-plan;
- example clarity;
- skill host behavior;
- proof wording that could be mistaken for implementation correctness;
- benchmark wording that could be mistaken for downstream execution quality.

## Follow-up candidates

| Candidate | Why it is plausible |
| --- | --- |
| Broader changed-intent fixture coverage | Current planning inputs가 recorded lock과 diverge하는 더 많은 방식을 exercise한다. |
| Model workspace stale-lock wording verification | Skill-host summaries가 readiness 또는 relock authority를 overstate하지 않게 한다. |
| Homebrew implementation audit | Tap/formula evidence가 생긴 뒤에만 유용하다. Homebrew를 early `Available`로 mark하면 안 된다. |
| Third benchmark case selection | `not_measured`와 non-execution boundaries를 preserve하면서 evidence를 확장한다. |

## Recommended next task

Selected next task: model workspace stale-lock wording verification.

Why: 이 examples는 이제 human workflow를 설명하지만, model workspace summaries는
skill이 plan을 relock했다거나 no-terminal changes를 validated했다거나 CLI proof 없이
`ni run`이 safe해졌다고 overclaim할 위험이 가장 큰 surface다. Narrow wording
verification pass는 lock semantics를 바꾸거나 runtime execution을 추가하지 않고 그
surface를 harden할 수 있다.

## Next task prompt

```text
Goal:
Verify model workspace stale-lock wording after the amend/relock workflow examples pass.

Context:
ni is ni-kernel: a Project Intent Compiler for AI Agents. It is not a task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation system, release automation system, or downstream execution layer.

Current status boundaries:
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- LOCK-STALE means the existing lock is stale because current planning inputs differ from .ni/plan.lock.json.

Read first:
- AGENTS.md
- README.md
- README.ko.md
- docs/99_MODEL_WORKSPACE_STATUS.md
- docs/99_MODEL_WORKSPACE_STATUS.ko.md
- docs/103_STALE_LOCK_DIAGNOSTIC.md
- docs/103_STALE_LOCK_DIAGNOSTIC.ko.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md
- packages/claude-skills/README.md
- packages/claude-skills/README.ko.md
- packages/codex-skills/README.md
- packages/codex-skills/README.ko.md
- packages/claude-skills/*/SKILL.md
- packages/codex-skills/*/SKILL.md
- .agents/skills/*/SKILL.md
- scripts/check-skill-packs.sh

Task:
Audit model workspace and skill-pack wording for stale-lock recovery. Add only narrow wording updates if needed so these surfaces consistently say:
- Skills may help draft amended planning text.
- Skills may help explain a LOCK-STALE warning.
- Skills do not determine readiness.
- Skills do not lock or relock.
- Skills do not replace ni status, ni end, or ni run.
- Skills do not update .ni/plan.lock.json.
- ni run compiles a bounded handoff prompt only after the lock is current and does not execute downstream work.
- no-terminal assisted workflows do not provide deterministic validation.

Do not:
- change Go implementation or lock semantics;
- add a new ni amend command;
- run ni end or relock on the project root;
- edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json on the project root;
- mark Homebrew Available;
- mark model workspace packs Available;
- claim no-terminal deterministic validation;
- claim benchmark evidence proves downstream execution quality;
- add runtime execution, shell adapters, Codex exec adapters, queues, PR automation, or release automation.

Validation:
- go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/quality.sh
- bash scripts/demo-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- whether Go files were touched
- whether skill docs were touched
- whether static checks were touched
- readiness result
- validation results
- confirmation that .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified
- confirmation that ni end/relock were not run on the project root
- complete next executable Codex prompt for the selected next task
```
