# Amend / Relock Workflow Examples

## Current status

ni surfaces stale existing locks with the `LOCK-STALE` diagnostic. Conceptually,
`LOCK-STALE` means the existing lock is stale: current planning inputs differ
from `.ni/plan.lock.json`.

This does not change lock semantics. `ni status --proof --next-questions`
surfaces readiness and stale-lock diagnostics, `ni end` is the
CLI-authoritative relock step for this recovery flow, and `ni run` compiles a
bounded handoff prompt only after the lock is current. Skills are UX; CLI is
authority.

## Core recovery flow

```text
review changed intent
-> ni status --proof --next-questions
-> ni end
-> ni run --max-chars 4000
```

| Step | Purpose | Boundary |
| --- | --- | --- |
| Review changed intent | Inspect what changed in `.ni/contract.json`, required `docs/plan/**`, and the conversation that motivated the change. | Do not treat ambiguous or conflicting edits as accepted decisions. |
| `ni status --proof --next-questions` | Let the CLI report readiness, blockers, deferrals, warnings, proof, and next questions. | Model judgment does not determine readiness. |
| `ni end` | Let the CLI write the current lock after readiness allows it. | Do not manually edit `.ni/plan.lock.json`; do not say a skill relocked the plan. |
| `ni run --max-chars 4000` | Compile a bounded downstream handoff prompt from the current lock. | `ni run` does not execute downstream work. |

Before relock, `ni run` should refuse stale handoff. After relock, `ni run`
should compile from the current lock and remain a prompt compiler only.

## Amend vs re-plan vs new project

| Change type | Meaning | Example | Expected action | Can reuse existing project? |
| --- | --- | --- | --- | --- |
| Amend | Scoped change that can be reviewed within the same accepted project intent. | Add a tighter wording constraint or clarify an acceptance phrase. | Review the changed intent, update planning state, run the core recovery flow. | Yes. |
| Re-plan | Material change that reopens planning questions before relock. | Add a new primary capability, actor, requirement, risk, or delivery constraint. | Reopen questions, resolve blockers, then run the core recovery flow only after readiness allows it. | Usually yes, after planning is updated. |
| New project | Change that should not be treated as the same locked intent. | Turn a planning compiler into a tool that runs downstream work or move to an unrelated product. | Start a new project intent instead of relocking the old one. | No. |

## Example 1: Minor scope clarification after lock

| Field | Example |
| --- | --- |
| Scenario | A locked plan says "write install docs"; the user clarifies that source, release binary, and curl installer paths must stay separate. |
| What changed | Scope wording becomes more precise without adding a new capability. |
| Why the existing lock may be stale | Required `docs/plan/**` or `.ni/contract.json` may now hash differently from `.ni/plan.lock.json`. |
| What the user should inspect | The changed requirement text, linked evaluations, non-goals, and any affected install-status claims. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Refuse stale handoff if the existing lock no longer matches current planning inputs. |
| What `ni run` should do after relock | Compile a bounded handoff prompt from the current lock and not execute it. |
| What skills may help with | Draft clearer planning text and explain a `LOCK-STALE` warning. |
| What skills must not claim | Skills must not determine readiness, lock or relock, replace `ni status`, `ni end`, or `ni run`, or update `.ni/plan.lock.json`. |

## Example 2: Requirement or capability change after lock

| Field | Example |
| --- | --- |
| Scenario | A locked plan for prompt compilation gains a new capability for target seed conformance notes. |
| What changed | A capability or requirement changed, so traceability to evaluations, risks, and artifacts may need updates. |
| Why the existing lock may be stale | The lock recorded the old contract and required planning docs, not the changed capability graph. |
| What the user should inspect | Capability IDs, requirement links, evaluation coverage, high-severity risks, non-goals, and downstream seed boundaries. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Refuse stale handoff; a stale lock must not produce trusted downstream instructions. |
| What `ni run` should do after relock | Compile a bounded prompt from the accepted current lock only. |
| What skills may help with | Draft amended capability wording and identify missing evaluation links. |
| What skills must not claim | Skills must not weaken acceptance criteria to reach readiness or say the model relocked the plan. |

## Example 3: Deferral changes after lock

| Field | Example |
| --- | --- |
| Scenario | A deferred Homebrew task is reprioritized for audit, but Homebrew must remain `Planned` until implementation and install proof exist. |
| What changed | A deferral, follow-up, or sequencing decision changed. |
| Why the existing lock may be stale | The accepted plan's deferred work, roadmap pointers, or decision log may no longer match the lock. |
| What the user should inspect | Deferral records, roadmap links, accepted decisions, `Available` / `Experimental` / `Planned` status words, and non-goals. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Refuse stale handoff if the deferral change changed lockable inputs. |
| What `ni run` should do after relock | Compile the handoff prompt without claiming the deferred item is implemented. |
| What skills may help with | Explain which deferral wording changed and preserve claim boundaries. |
| What skills must not claim | Skills must not mark Homebrew `Available`, mark model workspace packs `Available`, or replace CLI readiness. |

## Example 4: `docs/plan/**` changed after lock

| Field | Example |
| --- | --- |
| Scenario | A user edits `docs/plan/07_evaluation_contract.md` after `.ni/plan.lock.json` exists. |
| What changed | A required planning doc changed after lock. |
| Why the existing lock may be stale | Required planning docs are lockable inputs, so the doc hash can differ from the recorded lock hash. |
| What the user should inspect | The exact doc diff, any matching `.ni/contract.json` updates, and whether the change is accepted, draft, deferred, or conflicting. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Refuse stale handoff and preserve the stop-before-execution boundary. |
| What `ni run` should do after relock | Compile from the new current lock only if the CLI accepted the plan state. |
| What skills may help with | Summarize the diff and draft synchronized docs/contract updates. |
| What skills must not claim | Skills must not silently edit locked planning docs or say proof text proves implementation correctness. |

## Example 5: Model workspace skill drafts amendment after lock

| Field | Example |
| --- | --- |
| Scenario | `ni-start` drafts revised planning text in a model workspace after a lock exists. |
| What changed | The model drafted amended planning text, but the CLI state is not authoritative until checked. |
| Why the existing lock may be stale | If the draft is accepted into `.ni/contract.json` or required `docs/plan/**`, current planning inputs can differ from `.ni/plan.lock.json`. |
| What the user should inspect | Draft versus accepted records, assumptions, open questions, docs/contract sync, and any `LOCK-STALE` warning. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Refuse stale handoff; a skill draft is not a lock. |
| What `ni run` should do after relock | Compile a bounded prompt from the current lock and not execute downstream work. |
| What skills may help with | Draft amended planning text and explain `LOCK-STALE`. |
| What skills must not claim | Skills do not determine readiness, do not lock or relock, do not replace `ni status`, `ni end`, or `ni run`, and do not update `.ni/plan.lock.json`. |

## Example 6: No-terminal assisted workflow hits stale lock

| Field | Example |
| --- | --- |
| Scenario | A no-terminal assisted user receives a model summary that says planning changed, and a trusted runner later reports `LOCK-STALE`. |
| What changed | Assisted drafting changed planning text or identified changed intent after lock. |
| Why the existing lock may be stale | The current lockable planning inputs differ from `.ni/plan.lock.json`. |
| What the user should inspect | The assisted draft, the trusted runner's CLI output, changed docs/contract records, and any blocker questions. |
| Required CLI recovery flow | `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000` |
| What `ni run` should do before relock | Refuse stale handoff; no-terminal workflows do not provide deterministic validation. |
| What `ni run` should do after relock | Compile a bounded handoff prompt only after the CLI lock is current. |
| What skills may help with | Explain the warning and prepare draft planning amendments for review. |
| What skills must not claim | No-terminal assistance must not claim deterministic validation, lock freshness, or implementation proof. |

For more no-terminal stale-lock cases, evidence levels, and say/do-not-say
examples, see
[`106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md`](106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md).
For a copy-paste checklist that helps trusted runners preserve workspace,
command, output, file-change, and fixture-versus-target context, see
[`107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md`](107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md).

## Say this / do not say this

| Say this | Do not say this |
| --- | --- |
| "LOCK-STALE means the existing lock no longer matches current planning inputs." | "The model relocked the plan." |
| "Review the changed intent, run `ni status --proof --next-questions`, then run `ni end` before generating a new `ni run` handoff." | "The skill updated `.ni/plan.lock.json`." |
| "Skills can help draft amended planning text, but the CLI remains the authority." | "The no-terminal workflow deterministically validated the amended plan." |
| "`ni run` compiles a bounded handoff prompt from the current lock; it does not execute downstream work." | "`ni run` verified the implementation." |
| "`LOCK-STALE` does not prove implementation failure, benchmark failure, or product readiness." | "The benchmark proves downstream execution quality." |

## Validation and current coverage

Current tests and checks cover:

- stale status warning;
- current lock no warning;
- no lock no warning;
- `ni run` stale refusal;
- fixture relock recovery.

The following remain manually audited:

- user interpretation of amend vs re-plan;
- example clarity;
- skill host behavior;
- proof wording that could be mistaken for implementation correctness;
- benchmark wording that could be mistaken for downstream execution quality.

## Follow-up candidates

| Candidate | Why it is plausible |
| --- | --- |
| Broader changed-intent fixture coverage | Exercises more ways current planning inputs can diverge from the recorded lock. |
| Model workspace stale-lock wording verification | Keeps skill-host summaries from overstating readiness or relock authority. |
| Homebrew implementation audit | Useful only after tap/formula evidence exists; should not mark Homebrew `Available` early. |
| Third benchmark case selection | Expands evidence while preserving `not_measured` and non-execution boundaries. |

## Recommended next task

Selected next task: model workspace stale-lock wording verification.

Why: these examples now explain the human workflow, but model workspace summaries
are the highest-risk place for overclaiming that a skill relocked a plan,
validated no-terminal changes, or made `ni run` safe without CLI proof. A narrow
wording verification pass can harden that surface without changing lock
semantics or adding runtime execution.

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
