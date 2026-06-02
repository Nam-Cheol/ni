# Model Workspace Stale-Lock Wording Verification

## Current status

- Model workspace packs are Experimental as a broad product path.
- Skills are UX; CLI is authority.
- `LOCK-STALE` exists as a CLI diagnostic for an existing lock that no longer
  matches current planning inputs.
- Skills may explain or draft around `LOCK-STALE`, but they do not validate,
  lock, relock, replace `ni status`, `ni end`, or `ni run`, or update
  `.ni/plan.lock.json`.
- `ni run` compiles a bounded handoff prompt only after the lock is current; it
  does not execute downstream work.

## Surfaces audited

| Surface | Files | Result | Notes |
| --- | --- | --- | --- |
| Claude skill package README | `packages/claude-skills/README.md`, `packages/claude-skills/README.ko.md` | Updated | Added explicit `LOCK-STALE`, amended-draft, no-relock, no-lockfile-update, and recovery-order wording while preserving Experimental status. |
| Codex skill package README | `packages/codex-skills/README.md`, `packages/codex-skills/README.ko.md` | Updated | Same narrow wording update as the Claude package. |
| Claude `SKILL.md` files | `packages/claude-skills/*/SKILL.md` | Updated | Added a shared stale-lock boundary block to all Claude skill files. |
| Codex `SKILL.md` files | `packages/codex-skills/*/SKILL.md` | Updated | Added the same boundary block to all Codex skill files. |
| repo-local `.agents` skills if present | `.agents/skills/*/SKILL.md` | Updated | Added the same boundary block to the repo-local Codex-style skills. |
| docs/99 model workspace status | `docs/99_MODEL_WORKSPACE_STATUS.md`, `docs/99_MODEL_WORKSPACE_STATUS.ko.md` | Updated | Added model-workspace stale-lock rules without changing host verification claims. |
| docs/103 stale lock diagnostic | `docs/103_STALE_LOCK_DIAGNOSTIC.md`, `docs/103_STALE_LOCK_DIAGNOSTIC.ko.md` | Audited, no change | Already explains `LOCK-STALE`, what stale does and does not prove, recovery flow, and authority boundaries. |
| docs/104 amend/relock workflow examples | `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md`, `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md` | Audited, no change | Already includes model workspace and no-terminal stale-lock examples with the correct CLI recovery order. |
| scripts/check-skill-packs.sh | `scripts/check-skill-packs.sh` | Updated | Added low-noise checks for durable stale-lock boundary phrases in skill files and package READMEs. |
| roadmap and proof-capture docs | `docs/51_*`, `docs/83_*`, `docs/101_*`, `docs/102_*` | Audited, no change | Existing wording preserves CLI authority, no-terminal assisted limits, model workspace Experimental status, and non-execution boundaries. |

## Required wording boundary

| Boundary | Correct wording | Forbidden wording | Status |
| --- | --- | --- | --- |
| CLI authority | Skills are UX; CLI is authority. | Skills decide readiness or replace CLI gates. | Preserved and statically checked. |
| skill drafting only | Skills may help draft amended planning text. | Skills accept changed intent by themselves. | Added to skill surfaces. |
| readiness determination | `ni status` determines `READY`, `BLOCKED`, and `READY_WITH_DEFERRALS`. | Model workspace determines readiness. | Preserved. |
| lock/relock authority | `ni end` writes the current lock after CLI readiness allows it. | A skill locks or relocks the plan. | Added explicit no-relock wording. |
| `.ni/plan.lock.json` mutation | Skills do not update `.ni/plan.lock.json`. | A skill updated or repaired `.ni/plan.lock.json`. | Added and statically checked. |
| `LOCK-STALE` explanation | `LOCK-STALE` means the existing lock no longer matches current planning inputs. | `LOCK-STALE` proves implementation failure, product unreadiness, benchmark failure, or downstream execution failure. | Preserved in docs/103 and added to skill surfaces. |
| no-terminal assisted workflow | No-terminal is Experimental / assisted and not deterministic validation. | No-terminal deterministically validates amended plans or lock freshness. | Preserved. |
| model workspace Experimental status | Model workspace packs are Experimental unless host-level install/discovery is verified. | Model workspace packs are Available globally. | Preserved. |
| `ni run` non-execution boundary | `ni run` compiles a bounded handoff prompt from a current lock and does not execute downstream work. | `ni run` executes implementation, Codex, shell commands, or downstream agents. | Preserved. |

## Findings

| Finding | Severity | Surface | Evidence | Change made | Blocks v0.5? |
| --- | --- | --- | --- | --- | --- |
| Stale-lock wording was correct but implicit on skill package surfaces. | Low | Skill package READMEs and `SKILL.md` files | Existing wording said stale locks/hash mismatches stop as `BLOCKED`, but did not consistently name `LOCK-STALE` or the full recovery order. | Added explicit `LOCK-STALE`, draft-only, no-relock, no-lockfile-update, and recovery-order wording. | No. |
| Model workspace status docs did not yet name the stale-lock skill boundary. | Low | docs/99 | docs/99 preserved Experimental status and CLI authority, but did not connect model workspace rules to `LOCK-STALE`. | Added narrow `LOCK-STALE` skill rules and recovery order. | No. |
| Durable stale-lock phrases were not statically checked in skill packs. | Low | `scripts/check-skill-packs.sh` | The script already checked Experimental status and CLI authority wording. | Added exact checks for low-noise stale-lock boundary phrases. | No. |

No material finding blocks v0.5. The changes are wording and guardrail
hardening only.

## Changes made

- `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md`: records this
  audit and next-task selection.
- `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md`: Korean
  companion with the same boundaries.
- `docs/99_MODEL_WORKSPACE_STATUS.md` and `.ko.md`: add narrow stale-lock
  model workspace rules.
- `packages/claude-skills/README.md` and `.ko.md`: add stale-lock authority
  bullets.
- `packages/codex-skills/README.md` and `.ko.md`: add stale-lock authority
  bullets.
- `packages/claude-skills/*/SKILL.md`, `packages/codex-skills/*/SKILL.md`,
  `.agents/skills/*/SKILL.md`: add a shared stale-lock boundary block.
- `scripts/check-skill-packs.sh`: add stable static checks for `LOCK-STALE`
  wording in skill files and package READMEs.

Skill files were changed because the audit found that the stale-lock behavior
was present but not explicit enough on the model workspace surfaces.

## Validation surface

`scripts/check-skill-packs.sh` currently enforces:

- skill metadata exists;
- package README files exist in English and Korean;
- package READMEs preserve `Status: Experimental.`;
- host/global install remains `not_verified`;
- "Skills are UX; CLI is authority." remains visible;
- skills do not replace `ni status`, `ni end`, or `ni run`;
- skills do not manually edit `.ni/plan.lock.json`;
- downstream execution, provider API, adapter, and automation phrases remain
  blocked where the script already checks them.

This pass added a static check because the risky phrases are durable, exact,
and low-noise:

- `LOCK-STALE`;
- `Skills may help draft amended planning text.`;
- `Skills may help explain \`LOCK-STALE\`.`;
- `Skills do not lock or relock.`;
- `Skills do not update \`.ni/plan.lock.json\`.`;
- `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000`.

The check is intentionally narrow. It does not try to parse every stale-lock
paragraph or enforce broad prose style.

## Remaining risks

- Provider host behavior remains unverified.
- Global install behavior remains unverified.
- Users may still misread model-drafted planning text as CLI state.
- No-terminal workflow remains assisted, not deterministic.
- Manual audit is still needed for future prose changes, examples, and
  provider-specific install claims.

## Recommended next task

Selected next task: E. no-terminal stale-lock example pass.

Why: model workspace wording now says what skills may and may not do around
`LOCK-STALE`. The nearest remaining user-confusion path is a no-terminal
assisted user treating a model-drafted stale-lock explanation as deterministic
validation or relock proof. A focused example pass can harden that surface
without changing lock semantics or adding execution behavior.

## Next task prompt

```text
Goal:
Run a no-terminal stale-lock example pass.

This is a documentation and example-boundary task. Do not change Go implementation, lock semantics, stale-lock hash semantics, or CLI command behavior.

Context:
ni is ni-kernel: a pre-runtime Project Intent Compiler for AI Agents. It is not a task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation system, release automation system, or downstream execution layer.

Current status boundaries:
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- No-terminal assisted workflows do not provide deterministic validation, lock freshness, hash verification, or prompt compilation without exact CLI output from a trusted runner.

Read first:
- AGENTS.md
- README.md
- README.ko.md
- docs/no-terminal.md
- docs/no-terminal.ko.md
- examples/no-terminal-assisted/README.md
- examples/no-terminal-assisted/README.ko.md
- docs/83_CONVERSATION_PROOF_CAPTURE.md
- docs/83_CONVERSATION_PROOF_CAPTURE.ko.md
- docs/103_STALE_LOCK_DIAGNOSTIC.md
- docs/103_STALE_LOCK_DIAGNOSTIC.ko.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md
- docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md
- docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md
- scripts/demo-check.sh

Task:
Audit and narrowly update no-terminal assisted docs/examples so stale-lock guidance consistently says:
- no-terminal is Experimental / assisted;
- model text may draft or explain changed intent and LOCK-STALE;
- model-only proof is draft-only and not deterministic validation;
- trusted CLI output is required for ni status, ni end, lock hash verification, and ni run;
- recovery order is review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000;
- no-terminal assistance does not lock, relock, update .ni/plan.lock.json, execute downstream work, prove implementation failure, prove product unreadiness, prove benchmark failure, or prove downstream execution failure.

Add or update a small no-terminal stale-lock example only if it clarifies the user path without creating a fake root relock or generated prompt execution claim.

Do not:
- run ni end or relock on the project root;
- edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json on the project root;
- execute generated prompts;
- mark Homebrew Available;
- mark model workspace packs Available;
- claim host-level model workspace verification;
- claim no-terminal deterministic validation;
- claim benchmark evidence proves implementation quality or downstream execution;
- add runtime execution, shell adapters, Codex exec adapters, queues, PR automation, release automation, or an execution evidence loop.

Validation:
- go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- whether Go files were touched
- whether no-terminal docs/examples were touched
- readiness result
- validation results
- confirmation that .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified
- confirmation that ni end/relock were not run on the project root
- confirmation that no runtime execution, generated prompt execution, Homebrew Available claim, model workspace Available claim, no-terminal deterministic validation claim, or benchmark overclaim was added
- selected next task and complete next executable Codex prompt
```
