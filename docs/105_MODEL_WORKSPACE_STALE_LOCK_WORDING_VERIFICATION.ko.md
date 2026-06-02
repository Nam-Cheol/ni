# Model Workspace Stale-Lock Wording Verification

## Current status

- Model workspace packs는 broad product path로 Experimental이다.
- Skills are UX; CLI is authority.
- `LOCK-STALE`는 existing lock이 current planning inputs와 더 이상 match하지
  않을 때의 CLI diagnostic이다.
- Skills는 `LOCK-STALE` 주변의 설명이나 draft를 도울 수 있지만 validate, lock,
  relock, `ni status`, `ni end`, `ni run` replace, `.ni/plan.lock.json` update는
  하지 않는다.
- `ni run`은 lock이 current인 뒤 bounded handoff prompt만 compile한다. Downstream
  work를 execute하지 않는다.

## Surfaces audited

| Surface | Files | Result | Notes |
| --- | --- | --- | --- |
| Claude skill package README | `packages/claude-skills/README.md`, `packages/claude-skills/README.ko.md` | Updated | `LOCK-STALE`, amended draft, no-relock, no-lockfile-update, recovery order wording을 추가했고 Experimental status를 보존했다. |
| Codex skill package README | `packages/codex-skills/README.md`, `packages/codex-skills/README.ko.md` | Updated | Claude package와 같은 narrow wording update를 적용했다. |
| Claude `SKILL.md` files | `packages/claude-skills/*/SKILL.md` | Updated | 모든 Claude skill files에 shared stale-lock boundary block을 추가했다. |
| Codex `SKILL.md` files | `packages/codex-skills/*/SKILL.md` | Updated | 모든 Codex skill files에 같은 boundary block을 추가했다. |
| repo-local `.agents` skills if present | `.agents/skills/*/SKILL.md` | Updated | repo-local Codex-style skills에도 같은 boundary block을 추가했다. |
| docs/99 model workspace status | `docs/99_MODEL_WORKSPACE_STATUS.md`, `docs/99_MODEL_WORKSPACE_STATUS.ko.md` | Updated | Host verification claims를 바꾸지 않고 model-workspace stale-lock rules를 추가했다. |
| docs/103 stale lock diagnostic | `docs/103_STALE_LOCK_DIAGNOSTIC.md`, `docs/103_STALE_LOCK_DIAGNOSTIC.ko.md` | Audited, no change | 이미 `LOCK-STALE`, stale이 증명하는 것/증명하지 않는 것, recovery flow, authority boundaries를 설명한다. |
| docs/104 amend/relock workflow examples | `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md`, `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md` | Audited, no change | 이미 model workspace와 no-terminal stale-lock examples, 올바른 CLI recovery order를 포함한다. |
| scripts/check-skill-packs.sh | `scripts/check-skill-packs.sh` | Updated | Skill files와 package READMEs의 durable stale-lock boundary phrases를 확인하는 low-noise checks를 추가했다. |
| roadmap and proof-capture docs | `docs/51_*`, `docs/83_*`, `docs/101_*`, `docs/102_*` | Audited, no change | Existing wording이 CLI authority, no-terminal assisted limits, model workspace Experimental status, non-execution boundaries를 보존한다. |

## Required wording boundary

| Boundary | Correct wording | Forbidden wording | Status |
| --- | --- | --- | --- |
| CLI authority | Skills are UX; CLI is authority. | Skills가 readiness를 decide하거나 CLI gates를 replace한다. | Preserved and statically checked. |
| skill drafting only | Skills may help draft amended planning text. | Skills가 changed intent를 스스로 accept한다. | Added to skill surfaces. |
| readiness determination | `ni status`가 `READY`, `BLOCKED`, `READY_WITH_DEFERRALS`를 determine한다. | Model workspace가 readiness를 determine한다. | Preserved. |
| lock/relock authority | `ni end`는 CLI readiness가 허용한 뒤 current lock을 write한다. | Skill이 plan을 lock 또는 relock한다. | Added explicit no-relock wording. |
| `.ni/plan.lock.json` mutation | Skills do not update `.ni/plan.lock.json`. | Skill이 `.ni/plan.lock.json`을 updated 또는 repaired했다. | Added and statically checked. |
| `LOCK-STALE` explanation | `LOCK-STALE` means the existing lock no longer matches current planning inputs. | `LOCK-STALE`가 implementation failure, product unreadiness, benchmark failure, downstream execution failure를 증명한다. | Preserved in docs/103 and added to skill surfaces. |
| no-terminal assisted workflow | No-terminal은 Experimental / assisted이며 deterministic validation이 아니다. | No-terminal이 amended plans 또는 lock freshness를 deterministically validate한다. | Preserved. |
| model workspace Experimental status | Host-level install/discovery가 verified되기 전까지 model workspace packs는 Experimental이다. | Model workspace packs are Available globally. | Preserved. |
| `ni run` non-execution boundary | `ni run`은 current lock에서 bounded handoff prompt를 compile하며 downstream work를 execute하지 않는다. | `ni run`이 implementation, Codex, shell commands, downstream agents를 execute한다. | Preserved. |

## Findings

| Finding | Severity | Surface | Evidence | Change made | Blocks v0.5? |
| --- | --- | --- | --- | --- | --- |
| Skill package surfaces의 stale-lock wording은 correct했지만 implicit했다. | Low | Skill package READMEs and `SKILL.md` files | Existing wording은 stale locks/hash mismatches에서 `BLOCKED`로 stop한다고 말했지만 `LOCK-STALE`와 full recovery order를 일관되게 name하지 않았다. | Explicit `LOCK-STALE`, draft-only, no-relock, no-lockfile-update, recovery-order wording을 추가했다. | No. |
| Model workspace status docs가 stale-lock skill boundary를 아직 name하지 않았다. | Low | docs/99 | docs/99는 Experimental status와 CLI authority를 보존했지만 model workspace rules를 `LOCK-STALE`와 연결하지 않았다. | Narrow `LOCK-STALE` skill rules와 recovery order를 추가했다. | No. |
| Durable stale-lock phrases가 skill packs에서 statically checked되지 않았다. | Low | `scripts/check-skill-packs.sh` | Script는 이미 Experimental status와 CLI authority wording을 확인했다. | Low-noise stale-lock boundary phrases에 대한 exact checks를 추가했다. | No. |

No material finding blocks v0.5. 이번 변경은 wording and guardrail hardening only다.

## Changes made

- `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md`: 이 audit과
  next-task selection을 기록한다.
- `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md`: 같은
  boundaries를 가진 Korean companion.
- `docs/99_MODEL_WORKSPACE_STATUS.md` and `.ko.md`: narrow stale-lock model
  workspace rules 추가.
- `packages/claude-skills/README.md` and `.ko.md`: stale-lock authority bullets
  추가.
- `packages/codex-skills/README.md` and `.ko.md`: stale-lock authority bullets
  추가.
- `packages/claude-skills/*/SKILL.md`, `packages/codex-skills/*/SKILL.md`,
  `.agents/skills/*/SKILL.md`: shared stale-lock boundary block 추가.
- `scripts/check-skill-packs.sh`: skill files와 package READMEs의
  `LOCK-STALE` wording을 위한 stable static checks 추가.

Skill files는 stale-lock behavior가 존재하지만 model workspace surfaces에서 충분히
explicit하지 않았기 때문에 변경했다.

## Validation surface

`scripts/check-skill-packs.sh`는 현재 다음을 enforce한다.

- skill metadata 존재;
- package README files가 English/Korean 모두 존재;
- package READMEs가 `Status: Experimental.`을 보존;
- host/global install은 `not_verified`로 유지;
- "Skills are UX; CLI is authority."가 visible함;
- skills가 `ni status`, `ni end`, `ni run`을 replace하지 않음;
- skills가 `.ni/plan.lock.json`을 manually edit하지 않음;
- downstream execution, provider API, adapter, automation phrases는 기존 script
  checks가 막는 범위에서 blocked.

이번 pass는 risky phrases가 durable, exact, low-noise이기 때문에 static check를
추가했다.

- `LOCK-STALE`;
- `Skills may help draft amended planning text.`;
- `Skills may help explain \`LOCK-STALE\`.`;
- `Skills do not lock or relock.`;
- `Skills do not update \`.ni/plan.lock.json\`.`;
- `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000`.

이 check는 의도적으로 narrow하다. 모든 stale-lock paragraph를 parse하거나 broad
prose style을 enforce하지 않는다.

## Remaining risks

- Provider host behavior remains unverified.
- Global install behavior remains unverified.
- Users may still misread model-drafted planning text as CLI state.
- No-terminal workflow remains assisted, not deterministic.
- Future prose changes, examples, provider-specific install claims에는 manual
  audit이 여전히 필요하다.

Follow-up note: no-terminal stale-lock examples는
[`106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md`](106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md)에
기록되어 있다.

## Recommended next task

Selected next task: E. no-terminal stale-lock example pass.

Why: model workspace wording은 이제 skills가 `LOCK-STALE` 주변에서 무엇을 할 수
있고 무엇을 하면 안 되는지 말한다. 가장 가까운 remaining user-confusion path는
no-terminal assisted user가 model-drafted stale-lock explanation을 deterministic
validation 또는 relock proof처럼 취급하는 경우다. Focused example pass는 lock
semantics를 바꾸거나 execution behavior를 추가하지 않고 이 surface를 harden한다.

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
