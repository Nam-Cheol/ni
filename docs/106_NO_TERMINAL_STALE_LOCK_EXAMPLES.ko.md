# No-Terminal Stale-Lock Examples

## Current status

- No-terminal method는 Experimental / assisted이다.
- `LOCK-STALE`는 CLI diagnostic이다.
- Model-only no-terminal text는 draft-only이다.
- Readiness, lock freshness, relock, hash verification, bounded handoff
  compilation을 claim하기 전에는 trusted runner의 exact CLI output이 필요하다.

## What no-terminal can and cannot prove

| Evidence level | Can support | Cannot support | Required caveat |
| --- | --- | --- | --- |
| Model-only draft | Draft planning text, changed-intent interpretation, assumptions, open questions, conceptual `LOCK-STALE` explanation. | Readiness, lock freshness, relock, hash verification, bounded handoff compilation. | "This is a draft interpretation until exact CLI output is available." |
| Pasted CLI output | `ni status`에 보이는 `LOCK-STALE` warning 설명처럼 pasted command output에 한정된 statement. | Shown command 밖의 claims, global workspace state, relock success, later edits 이후 hash freshness, implementation correctness. | Exact command result만 설명하고 missing CLI lines를 invent하지 않는다. |
| Trusted runner transcript | Target workspace 또는 explicitly isolated fixture에서의 실제 `ni status`, `ni end`, `ni run` behavior. | Product implementation quality, downstream execution quality, host/provider behavior, 다른 workspace evidence. | Workspace 또는 fixture를 name하고 claims를 해당 transcript에 scope한다. |

## Core stale-lock recovery flow

```text
review changed intent
-> ni status --proof --next-questions
-> ni end
-> ni run --max-chars 4000
```

No-terminal assisted terms에서는 다음과 같다.

- Model은 changed intent review를 draft할 수 있다.
- User 또는 trusted runner가 `ni status --proof --next-questions`를 run해야 한다.
- Relock이 필요하고 readiness가 허용하면 user 또는 trusted runner가 `ni end`를
  run해야 한다.
- Bounded handoff를 compile하려면 user 또는 trusted runner가
  `ni run --max-chars 4000`을 run해야 한다.
- Model은 exact pasted CLI output을 explain할 수 있지만 CLI results를 invent하면
  안 된다.

## Example 1: Model-only planning draft after suspected changed intent

| Field | Example |
| --- | --- |
| Scenario | No-terminal user가 lock 이후 accepted planning intent가 바뀌었을 수 있다고 말하지만 CLI output은 제공하지 않는다. |
| Available evidence | Model-only draft와 user-described changed intent. |
| What the model may say | "This is a draft interpretation until exact CLI output is available. Changed intent가 planning text에 영향을 준 것 같으므로 trusted runner에게 `ni status --proof --next-questions` 실행을 요청하기 전에 review하세요." |
| What the model must not say | "The plan is READY", "the lock is stale", "the lock is current", "the model relocked the plan." |
| Required next CLI action | Trusted runner가 target workspace에서 `ni status --proof --next-questions`를 run한다. |
| Whether `ni run` is safe before relock | No. Exact CLI output과 current lock 없이 no-terminal은 `ni run`이 safe하다고 claim할 수 없다. |
| Whether `ni run` is safe after relock | `ni end`가 succeed하고 exact `ni run --max-chars 4000` output이 available일 때만 가능하다. |
| Boundary note | Model-only text는 changed intent를 draft할 수 있지만 deterministic validation을 제공하지 않는다. |

## Example 2: User pastes LOCK-STALE from ni status

| Field | Example |
| --- | --- |
| Scenario | User가 `ni status --proof --next-questions`에서 나온 `WARNING: LOCK-STALE existing lock is stale. Current planning inputs differ from .ni/plan.lock.json.`를 paste한다. |
| Available evidence | `ni status`의 pasted CLI output; pasted `ni end` 또는 `ni run` output은 없다. |
| What the model may say | "`LOCK-STALE` means the existing lock no longer matches current planning inputs. Changed intent를 review한 뒤 CLI recovery flow를 사용하세요." |
| What the model must not say | "The model verified hashes", "the plan was relocked", "`ni run` can compile now." |
| Required next CLI action | Changed intent를 review한 뒤 readiness가 허용하고 user가 accepted plan을 confirm하면 `ni end`를 run한다. |
| Whether `ni run` is safe before relock | No. Stale lock은 trusted handoff에 사용하면 안 된다. |
| Whether `ni run` is safe after relock | Trusted `ni end` output이 current lock written을 보여준 뒤 `ni run --max-chars 4000`이 그 lock에 대해 run될 때만 가능하다. |
| Boundary note | Pasted `ni status` output은 shown status만 support하며 relock 또는 bounded handoff compilation은 support하지 않는다. |

## Example 3: Trusted runner provides exact ni status output

| Field | Example |
| --- | --- |
| Scenario | User는 terminal을 local로 run할 수 없지만 teammate가 target workspace의 exact `ni status --proof --next-questions` transcript를 제공한다. |
| Available evidence | `ni status`에 대한 trusted runner transcript; workspace path와 timestamp가 포함되어 있다. |
| What the model may say | "이 transcript는 해당 workspace와 시점의 readiness와 warning state를 support할 수 있습니다. `LOCK-STALE`가 보이면 changed intent review 뒤 next step은 여전히 `ni end`입니다." |
| What the model must not say | "The provider host is verified", "the implementation is ready", "later edits are covered by this old transcript." |
| Required next CLI action | Status가 허용하고 relock이 필요하면 trusted runner가 user confirmation 뒤 `ni end`를 run한다. |
| Whether `ni run` is safe before relock | Transcript가 `LOCK-STALE`를 보여주거나 current lock proof가 없으면 No. |
| Whether `ni run` is safe after relock | Trusted `ni end`와 trusted `ni run --max-chars 4000` output 뒤 prompt compilation으로만 Yes. |
| Boundary note | Trusted runner transcript는 model-only text보다 강하지만, 여전히 shown workspace와 command에 scope된다. |

## Example 4: Model workspace skill drafts amended planning text

| Field | Example |
| --- | --- |
| Scenario | Stale-lock warning이 pasted된 뒤 model workspace의 `ni-start`가 amended `docs/plan/**` text를 draft한다. |
| Available evidence | Pasted `LOCK-STALE` output과 model-drafted amendment text. |
| What the model may say | "이 draft는 changed intent review에 도움을 줄 수 있습니다. Accepted planning edits가 적용된 뒤 trusted runner가 `ni status --proof --next-questions`를 run해야 합니다." |
| What the model must not say | "The skill updated `.ni/plan.lock.json`", "the skill relocked the plan", "the no-terminal workflow deterministically validated the amended plan." |
| Required next CLI action | Accepted planning changes를 적용한 뒤 `ni status --proof --next-questions`를 run하고, 허용되면 `ni end`를 run한다. |
| Whether `ni run` is safe before relock | No. Skill drafting은 lock freshness 또는 hash verification이 아니다. |
| Whether `ni run` is safe after relock | CLI lock이 current이고 exact `ni run --max-chars 4000` output이 있을 때만 가능하다. |
| Boundary note | Skills are UX; CLI is authority. Skills는 amendments를 draft하고 pasted CLI output을 explain할 수 있지만 lock 또는 relock하지 않는다. |

## Example 5: No CLI output is available

| Field | Example |
| --- | --- |
| Scenario | User가 lock이 current인지 묻지만 CLI output 또는 trusted runner transcript가 없다. |
| Available evidence | Model-only conversation과 repository text, current CLI output은 없음. |
| What the model may say | "No-terminal cannot claim lock freshness without exact CLI output from a trusted runner." |
| What the model must not say | "The pasted draft proves the lock is current", "the no-terminal workflow verified hashes", "`ni run` is safe." |
| Required next CLI action | Run 또는 request `ni status --proof --next-questions`. |
| Whether `ni run` is safe before relock | No. Deterministic current-lock proof가 없다. |
| Whether `ni run` is safe after relock | Trusted `ni end`와 `ni run --max-chars 4000` output이 available일 때만 가능하다. |
| Boundary note | CLI output이 없으면 model은 draft guidance만 제공할 수 있고 CLI-state claims는 할 수 없다. |

## Example 6: Trusted runner relocks and provides ni run output

| Field | Example |
| --- | --- |
| Scenario | Trusted runner가 changed intent를 review하고 `ni status --proof --next-questions`, `ni end`를 run한 뒤 exact `ni run --max-chars 4000` output을 제공한다. |
| Available evidence | Status, relock, bounded prompt compilation에 대한 trusted runner transcript. |
| What the model may say | "`ni run` compiled a bounded handoff prompt from the current lock shown in this transcript; it did not execute downstream work." |
| What the model must not say | "`ni run` verified the implementation", "the benchmark proves downstream execution quality", "the generated prompt was executed by ni." |
| Required next CLI action | Pasted transcript 설명에는 추가 action이 없다. 이후 changed intent가 생기면 recovery flow를 다시 시작한다. |
| Whether `ni run` is safe before relock | No. Transcript의 safe compilation claim은 recorded relock 이후에만 적용된다. |
| Whether `ni run` is safe after relock | Exact `ni run --max-chars 4000` output이 있을 때 bounded prompt compilation으로만 Yes. |
| Boundary note | `ni run` compiles a bounded handoff prompt from the current lock; it does not execute downstream work. |

## Say this / do not say this

| Say this | Do not say this |
| --- | --- |
| "This is a draft interpretation until exact CLI output is available." | "The no-terminal workflow deterministically validated the amended plan." |
| "LOCK-STALE means the existing lock no longer matches current planning inputs." | "The model relocked the plan." |
| "A trusted runner must run `ni status --proof --next-questions` before claiming the current readiness state." | "The skill updated `.ni/plan.lock.json`." |
| "`ni end` is the CLI-authoritative relock step after changed intent has been reviewed." | "The pasted draft proves the lock is current." |
| "`ni run` compiles a bounded handoff prompt from the current lock; it does not execute downstream work." | "`ni run` verified the implementation." |
| "A trusted runner transcript can support only the workspace and commands shown." | "The benchmark proves downstream execution quality." |

## Documentation alignment

Updated:

- `docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md`
- `docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md`
- `docs/no-terminal.md`
- `docs/no-terminal.ko.md`
- `docs/51_POST_RELEASE_ROADMAP.md`
- `docs/51_POST_RELEASE_ROADMAP.ko.md`
- `docs/103_STALE_LOCK_DIAGNOSTIC.md`
- `docs/103_STALE_LOCK_DIAGNOSTIC.ko.md`
- `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md`
- `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md`
- `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md`
- `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md`
- `scripts/demo-check.sh`

Skill files는 변경하지 않았다. 기존 package 및 repo-local skill wording은 충분했다.
이전 model workspace stale-lock verification pass가 이미 `LOCK-STALE`, amended
drafting, no-relock, no-lockfile-update, recovery order를 explicit하게 만들고
statically checked하기 때문이다.

## Validation and current coverage

Current checks cover:

- `scripts/check-skill-packs.sh`를 통한 skill pack boundary wording;
- `scripts/check-install-docs.py`를 통한 install docs consistency;
- `scripts/demo-check.sh`를 통한 demo coverage;
- `scripts/quality.sh`를 통한 quality checks;
- status warnings, current locks, stale `ni run` refusal, fixture relock
  recovery에 대한 prior implementation stale-lock tests.

다음은 manually audited로 남는다.

- no-terminal user interpretation;
- trusted runner transcript quality;
- provider host behavior;
- users가 draft text와 CLI output을 구분하는지 여부.

## Remaining risks

- Users may paste incomplete CLI output.
- Users may treat model-drafted text as CLI state.
- Trusted runner context may differ from the target workspace.
- Provider host behavior remains unverified.
- No-terminal remains assisted, not deterministic.

## Recommended next task

Selected next task: F. no-terminal transcript quality checklist.

Why: examples가 이제 stale-lock path를 설명하지만, 다음 weak point는 pasted 또는
delegated transcripts의 품질이다. Checklist는 users와 trusted runners가 command,
workspace, timestamp, output, scope를 포함하도록 도울 수 있으며 CLI behavior를
바꾸거나 deterministic no-terminal validation을 claim하지 않는다.

## Next task prompt

```text
Goal:
Create a no-terminal transcript quality checklist for trusted runner handoffs.

This is a documentation and example-boundary task. Do not change Go implementation, lock semantics, stale-lock hash semantics, CLI command behavior, validation semantics, or prompt compilation behavior.

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
- No-terminal assisted workflows do not provide deterministic validation, lock freshness, hash verification, relock, or bounded handoff compilation without exact CLI output from a trusted runner.

Read first:
- AGENTS.md
- README.md
- README.ko.md
- docs/no-terminal.md
- docs/no-terminal.ko.md
- docs/103_STALE_LOCK_DIAGNOSTIC.md
- docs/103_STALE_LOCK_DIAGNOSTIC.ko.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md
- docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md
- docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md
- docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md
- docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md
- examples/no-terminal-assisted/README.md
- examples/no-terminal-assisted/README.ko.md
- scripts/demo-check.sh

Task:
Add a narrow no-terminal transcript quality checklist that helps users distinguish model-only drafts, pasted CLI output, and trusted runner transcripts.

The checklist should require at least:
- exact command;
- workspace path or fixture path;
- timestamp or run context;
- full relevant CLI output;
- whether output came from target workspace or isolated fixture;
- whether any planning files changed after the transcript;
- what claims the transcript can and cannot support;
- next required CLI action in the sequence review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000.

Preserve these boundaries:
- model-only text is draft-only;
- pasted CLI output supports only the command shown;
- trusted runner transcripts support only the workspace or fixture shown;
- no-terminal does not determine readiness, lock, relock, verify hashes, or compile bounded handoff without exact CLI output;
- ni run compiles a bounded handoff prompt and does not execute downstream work;
- LOCK-STALE does not prove implementation failure, product unreadiness, benchmark failure, or downstream execution failure.

Do not:
- edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json on the project root;
- run ni end or relock on the project root;
- execute generated prompts;
- add runtime execution, shell adapters, Codex exec adapters, queues, PR automation, release automation, or execution evidence loops;
- mark Homebrew Available;
- mark model workspace packs Available;
- claim no-terminal deterministic validation;
- claim host-level model workspace verification;
- claim benchmark evidence proves implementation quality or downstream execution quality.

Validation:
- go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- bash scripts/quality.sh
- bash scripts/smoke.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- whether Go files were touched
- whether skill docs were touched
- whether static checks were touched
- checklist summary
- how the checklist preserves CLI authority
- validation results
- confirmation that .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified
- confirmation that ni end/relock were not run on the project root
- confirmation that no runtime execution, generated prompt execution, Homebrew Available claim, model workspace Available claim, no-terminal deterministic validation claim, host-level model workspace verification claim, or benchmark overclaim was added
- selected next task and complete next executable Codex prompt
```
