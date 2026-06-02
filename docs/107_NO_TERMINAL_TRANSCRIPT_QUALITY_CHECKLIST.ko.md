# No-Terminal Transcript Quality Checklist

## Current status

- No-terminal method is Experimental / assisted.
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- Model-only text는 draft-only이다.
- Readiness, lock freshness, relock, hash verification, bounded handoff
  compilation을 claim하기 전에는 trusted runner의 exact CLI output이 필요하다.

No-terminal은 planning amendments를 draft하고, trusted runner용 commands를
준비하고, exact pasted CLI output을 explain할 수 있다. 하지만 그 자체로
deterministic validation을 제공하거나, `ni status` output 없이 readiness를
determine하거나, `ni end` output 없이 lock 또는 relock하거나, trusted runner의 exact
CLI output 없이 lock freshness 또는 hashes를 verify하거나, `ni run` output 없이
authoritative bounded handoff를 compile하거나, downstream work를 execute할 수 없다.

`ni status` determines readiness and surfaces stale-lock diagnostics. `ni end`
is the CLI-authoritative lock/relock step. `ni run` compiles a bounded handoff
prompt from the current lock. `ni run` does not execute downstream work.

## Transcript quality levels

| Level | Description | Can support | Cannot support | Required caveat |
| --- | --- | --- | --- | --- |
| model-only draft | Actual CLI output 없이 model이 drafted한 text. | Draft planning text, changed-intent review, assumptions, open questions, conceptual explanation. | CLI-state claims, readiness, lock freshness, relock, hash verification, bounded handoff compilation, downstream execution claims. | "This is a draft interpretation until exact CLI output is available." |
| unusable transcript | Command와 output을 식별하기에 ambiguous, invented, stale, wrong-workspace, 또는 too incomplete한 transcript. | Better transcript 요청 정도. | Readiness, `LOCK-STALE`, relock, prompt compilation, target workspace state, implementation correctness, product readiness. | CLI evidence가 없는 것으로 취급한다. |
| partial transcript | 일부 pasted output은 있지만 workspace, command, status, timestamp, relevant surrounding lines가 빠진 transcript. | Visible lines에 대한 limited interpretation. | Whole-command result, lock freshness, relock success, later hash state, workspace identity가 없을 때 target workspace state. | Claim 전에 missing fields를 name한다. |
| pasted CLI output | 무엇을 run했는지 알 수 있을 만큼 command context가 있는 user-pasted exact command output. | `ni status`가 printed한 readiness 또는 warning state처럼 output이 직접 보여주는 claim. | Shown command 밖의 claims, command 이후 file changes, other workspaces, implementation quality, downstream execution success. | 모든 claim을 pasted command와 time에 scope한다. |
| trusted runner transcript | Trusted runner가 제공한 command, workspace context, exact CLI output, 가능하면 exit status 또는 freshness note. | Complete enough라면 shown command와 shown workspace에 대한 CLI-state claims. | Product readiness, benchmark effect size, provider host behavior, transcript 밖의 claims. | Workspace가 target인지 fixture인지 name한다. |
| fixture transcript | Temporary 또는 isolated test/demo workspace transcript. | Fixture transcript supports only fixture claims. | Target workspace readiness, target relock, target lock freshness, target `ni run` compilation. | 모든 claim 근처에 "fixture-only"라고 말한다. |
| target-workspace transcript | 논의 중인 actual project workspace에서 나온 transcript. | Complete enough라면 transcript가 보여주는 target workspace CLI-state claims. | Later state after edits, unrelated workspaces, implementation correctness, downstream execution success, product readiness, benchmark effect size, adoption, cost, latency improvement. | Shown workspace, command, output, moment만 증명한다. |

이 checklist에서 usable pasted CLI output은 command context가 충분히 complete한
`pasted CLI output` row를 뜻한다. Usable trusted runner transcript는 아래 minimum
fields가 present한 `trusted runner transcript` row를 뜻한다. Fixture-only transcript와
target-workspace transcript는 summaries에서 visible하게 유지해야 하는 scope label이다.

## Minimum trusted-runner transcript fields

Usable trusted runner transcript는 다음을 포함하는 것이 좋다:

- Workspace identity 또는 path, privacy-safe redaction 허용.
- Workspace가 target인지 fixture인지.
- Command run.
- Exact command output.
- Exit status if available.
- Timestamp 또는 freshness note if available.
- Command 전후 files changed 여부.
- `ni end`가 target workspace에서 run되었는지 fixture 안에서만 run되었는지.
- `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` changed 여부.
- `ni run` output이 포함된 경우 prompt가 compiled되었지만 executed되지 않았다는
  confirmation.

Sensitive absolute paths를 요구하지 않는다. Target versus fixture identity가 clear한
한 privacy-safe redaction은 허용된다.

## Claim boundary matrix

| Evidence | Allowed claim | Forbidden claim | Next action |
| --- | --- | --- | --- |
| no CLI output | Model은 planning text를 draft하거나 protocol을 conceptually explain할 수 있다. | "The model verified readiness without CLI output." | Target workspace의 `ni status --proof --next-questions` output을 요청한다. |
| partial `ni status` output | Visible lines만 interpret할 수 있다. | "The partial transcript proves the lock is current." | Full command, workspace type, exact output, freshness note를 요청한다. |
| complete `ni status` output | 해당 workspace와 time에 대해 output이 보여주는 readiness와 warnings. | Relock success, later edits 이후 current hash state, shown되지 않은 `ni run` safety. | `READY`이고 warnings가 없으면 accepted intent confirmed 뒤에만 `ni end`를 요청한다. |
| `ni status` output with `LOCK-STALE` | `LOCK-STALE` means the existing lock no longer matches current planning inputs. | Implementation failure, product unreadiness, benchmark failure, downstream execution failure, completed relock. | Changed intent를 review한 뒤 readiness가 허용하면 `ni end`를 run한다. |
| `ni end` output | 해당 output이 보여주는 lock 또는 relock behavior. | Later lock freshness, downstream handoff compilation, implementation verification. | Lock이 current이면 `ni run --max-chars 4000`으로 bounded handoff를 compile한다. |
| `ni run` output | 해당 output에 보이는 bounded prompt compilation. | "`ni run` verified the implementation" 또는 "the generated prompt was executed by ni." | Prompt가 compiled만 되었고 executed되지 않았는지 confirm한다. |
| fixture transcript | Fixture-only CLI behavior. | "The fixture relock proves the project root was relocked." | Target workspace에서 필요한 같은 command를 run 또는 request한다. |
| target-workspace transcript | Complete transcript가 보여주는 target workspace CLI-state claims. | Later edits, unrelated workspaces, implementation correctness, product readiness, benchmark effect size, adoption, cost, latency claims. | Shown command와 time에 scoped된 claim만 계속한다. |
| validation script transcript | Script behavior와 script가 clearly run한 fixture commands. | Transcript가 project root에서 `ni end`를 explicit하게 보여주지 않는 project-root relock claim. | Summary에서 fixture runs와 root workspace state를 분리한다. |

어떤 transcript도 implementation correctness, downstream execution success,
product readiness, benchmark effect size, adoption, cost, latency improvement를
증명하지 않는다.

## Core stale-lock recovery flow

```text
review changed intent
-> ni status --proof --next-questions
-> ni end
-> ni run --max-chars 4000
```

No-terminal assisted terms에서는 다음과 같다.

- Model은 changed intent review를 draft할 수 있다.
- Model은 trusted runner용 commands를 prepare할 수 있다.
- User 또는 trusted runner가 `ni status`를 run해야 한다.
- Relock이 필요하면 user 또는 trusted runner가 `ni end`를 run해야 한다.
- Bounded handoff를 compile하려면 user 또는 trusted runner가 `ni run`을 run해야
  한다.
- Model은 exact pasted CLI output을 explain할 수 있다.
- Model은 CLI results를 invent하면 안 된다.

## Transcript examples

### Example 1: Unusable model-invented transcript

| Field | Example |
| --- | --- |
| Transcript quality | model-only draft; CLI-state claims에는 unusable transcript. |
| Available evidence | Command output 없이 "The model says ni status was READY"라고 적힌 text. |
| Claim allowed | Model이 interpretation을 draft했다. |
| Claim not allowed | Readiness, lock freshness, relock, hash verification, bounded handoff compilation. |
| Missing information | Workspace, command, exact output, exit status, freshness, file-change context. |
| Required next action | Trusted runner에게 target workspace에서 `ni status --proof --next-questions`를 run하게 한다. |
| Boundary note | Model-only draft cannot support CLI-state claims. |

### Example 2: Partial pasted ni status output with missing workspace context

| Field | Example |
| --- | --- |
| Transcript quality | partial transcript. |
| Available evidence | Pasted lines에 `NI Intent Readiness: READY`가 보이지만 workspace, command, warnings, timestamp가 없다. |
| Claim allowed | Visible pasted line이 `READY`라고 말한다. |
| Claim not allowed | Target workspace readiness, `LOCK-STALE` absence, current lock freshness, safe `ni run`. |
| Missing information | Target versus fixture identity, full command, full output, warnings section, file changes, freshness. |
| Required next action | Minimum fields를 포함한 complete target-workspace transcript를 요청한다. |
| Boundary note | Incomplete pasted output can support only limited interpretation. |

### Example 3: Usable pasted ni status output showing READY and no warnings

| Field | Example |
| --- | --- |
| Transcript quality | usable pasted CLI output. |
| Available evidence | Target workspace, command `ni status --proof --next-questions`, exact output showing `NI Intent Readiness: READY`, no blockers, no deferrals, no warnings. |
| Claim allowed | Pasted `ni status` output은 해당 workspace와 time의 readiness와 warning state를 support한다. |
| Claim not allowed | Lock/relock happened, `ni run` compiled a prompt, later edits remain covered. |
| Missing information | 해당 claims가 필요하면 `ni end` output과 `ni run` output. |
| Required next action | Accepted plan이 confirmed되고 lock/relock이 필요하면 `ni end`를 run한다. |
| Boundary note | Pasted `ni status` output can support only the readiness/stale-lock state shown in that output. |

### Example 4: Usable pasted ni status output showing LOCK-STALE

| Field | Example |
| --- | --- |
| Transcript quality | usable pasted CLI output. |
| Available evidence | Target workspace `ni status --proof --next-questions` output에 `LOCK-STALE`가 포함되어 있다. |
| Claim allowed | Existing lock is stale because current planning inputs differ from `.ni/plan.lock.json`. |
| Claim not allowed | Relock success, implementation failure, product unreadiness, benchmark failure, downstream execution failure, safe `ni run`. |
| Missing information | Changed-intent review, `ni end` output, later `ni run --max-chars 4000` output. |
| Required next action | Changed intent를 review한 뒤 readiness가 허용하고 user가 confirm하면 `ni end`를 run한다. |
| Boundary note | `LOCK-STALE`는 CLI diagnostic이지 implementation verdict가 아니다. |

### Example 5: Fixture transcript that must not be claimed as target workspace state

| Field | Example |
| --- | --- |
| Transcript quality | fixture-only transcript. |
| Available evidence | Validation script가 temporary workspace를 만들고 `ni end`, `ni run`을 실행한다. |
| Claim allowed | Fixture flow가 해당 fixture 안에서 lock과 bounded prompt compilation을 exercised했다. |
| Claim not allowed | Project root was relocked, target lock is current, target `ni run` is safe. |
| Missing information | 같은 command sequence에 대한 target workspace transcript. |
| Required next action | Target-workspace `ni status`를 run 또는 request하고, 필요할 때만 `ni end`와 `ni run`을 run한다. |
| Boundary note | A fixture transcript supports fixture claims, not target workspace claims. |

### Example 6: Target workspace transcript showing relock and subsequent ni run output

| Field | Example |
| --- | --- |
| Transcript quality | usable trusted runner transcript; target-workspace transcript. |
| Available evidence | Target workspace transcript가 changed-intent review note, `ni status --proof --next-questions`, `ni end`, `ni run --max-chars 4000` output을 보여준다. |
| Claim allowed | `ni end`가 target workspace에 표시된 CLI-authoritative relock step이었고, `ni run`이 current lock에서 bounded handoff prompt를 compiled했다. |
| Claim not allowed | Generated prompt was executed, implementation quality was verified, downstream execution succeeded. |
| Missing information | Shown transcript 설명에는 없음; later edits에는 fresh proof가 필요하다. |
| Required next action | `ni`의 일부로 generated prompt를 execute하지 않는다. Downstream execution은 separate이다. |
| Boundary note | `ni run` compiles a bounded handoff prompt from the current lock; it does not execute downstream work. |

### Example 7: ni run output pasted without preceding lock freshness evidence

| Field | Example |
| --- | --- |
| Transcript quality | Lock-freshness claims에는 partial transcript; visible `ni run` result에는 pasted CLI output. |
| Available evidence | User가 `ni run --max-chars 4000` output을 paste하지만 current `ni status` 또는 `ni end` transcript는 없다. |
| Claim allowed | Pasted `ni run` output은 그 output에 보이는 bounded prompt compilation만 support한다. |
| Claim not allowed | Run 전 current readiness, why the lock was current, later edits are covered. |
| Missing information | Workspace type, prior status, lock/relock proof, file-change timing, exit status. |
| Required next action | 같은 workspace의 preceding `ni status`와 관련 있으면 `ni end` transcript를 요청한다. |
| Boundary note | Pasted `ni run` output can support only bounded prompt compilation shown in that output. |

### Example 8: Validation script transcript with fixture ni end runs

| Field | Example |
| --- | --- |
| Transcript quality | Root commands가 explicit하지 않으면 fixture-only subclaims를 가진 validation script transcript. |
| Available evidence | `bash scripts/quality.sh` 또는 `bash scripts/demo-check.sh` output이 tests passed와 temporary fixture commands를 보여준다. |
| Claim allowed | Validation script passed이고 fixture `ni end` runs는 script가 말한 곳에서만 happened. |
| Claim not allowed | Project-root `.ni/plan.lock.json` changed, project-root relock succeeded, target workspace handoff was compiled unless explicitly shown. |
| Missing information | Root relock을 claim하려면 root lockfile diff와 root command transcript. |
| Required next action | Fixture/root distinction을 state하고, target claims가 필요하면 git diff를 inspect하거나 target workspace commands를 run한다. |
| Boundary note | Validation script transcript must not describe fixture `ni end` runs as project-root relock. |

## Trusted runner copy-paste template

```text
Workspace:
Workspace type: target / fixture
Commands run:
Exact output:
Exit status:
Files changed:
Lockfile diff:
Was ni end run on project root? yes/no
Was relock run on project root? yes/no
Was ni run output compiled? yes/no
Was the generated prompt executed? no
```

Privacy-safe redaction은 허용된다. Secrets를 paste하지 않는다. Absolute path가
sensitive하면 private segments를 `[redacted]`로 바꾸되 target versus fixture
identity는 clear하게 유지한다.

## Say this / do not say this

| Say this | Do not say this |
| --- | --- |
| "This is a draft interpretation until exact CLI output is available." | "The model verified readiness without CLI output." |
| "This transcript supports only the workspace and command shown." | "The fixture relock proves the project root was relocked." |
| "A fixture transcript supports fixture claims, not target workspace claims." | "The partial transcript proves the lock is current." |
| "LOCK-STALE means the existing lock no longer matches current planning inputs." | "The no-terminal workflow deterministically validated the amended plan." |
| "ni end is the CLI-authoritative relock step after changed intent has been reviewed." | "The skill updated `.ni/plan.lock.json`." |
| "ni run compiles a bounded handoff prompt from the current lock; it does not execute downstream work." | "ni run verified the implementation." |
| "A validation script transcript may include fixture runs that do not change the target workspace." | "The benchmark proves downstream execution quality." |

## Documentation alignment

Updated:

- `docs/107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md`
- `docs/107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.ko.md`
- `docs/no-terminal.md`
- `docs/no-terminal.ko.md`
- `docs/51_POST_RELEASE_ROADMAP.md`
- `docs/51_POST_RELEASE_ROADMAP.ko.md`
- `docs/103_STALE_LOCK_DIAGNOSTIC.md`
- `docs/103_STALE_LOCK_DIAGNOSTIC.ko.md`
- `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md`
- `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md`
- `docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md`
- `docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md`
- `docs/108_CHANGED_INTENT_FIXTURE_COVERAGE.md`
- `docs/108_CHANGED_INTENT_FIXTURE_COVERAGE.ko.md`
- `scripts/demo-check.sh`

Skill files는 변경하지 않았다. Existing skill wording은 충분했다. Current skill
pack checks가 이미 `LOCK-STALE`, amended drafting, no-relock, no-lockfile-update,
recovery order를 보존하기 때문이다. 이 task는 model workspace skill behavior를
바꾸지 않고 users와 trusted runners를 위한 transcript quality guidance를 추가한다.

Representative changed-intent stale-lock cases와 non-lockable false-positive
guards에 대한 fixture-level proof는
[`108_CHANGED_INTENT_FIXTURE_COVERAGE.ko.md`](108_CHANGED_INTENT_FIXTURE_COVERAGE.ko.md)를
참고한다.

## Validation and current coverage

Current checks cover:

- `scripts/check-skill-packs.sh`를 통한 skill pack boundary wording;
- `scripts/demo-check.sh`를 통한 no-terminal stale-lock docs;
- `scripts/demo-check.sh`를 통한 this transcript-quality checklist;
- `scripts/check-install-docs.py`를 통한 install docs consistency;
- `scripts/demo-check.sh`를 통한 demo coverage;
- `scripts/quality.sh`를 통한 quality checks;
- prior implementation의 stale-lock tests.

다음은 manually audited로 남는다:

- transcript completeness;
- pasted output이 exact인지 여부;
- trusted runner workspace가 target workspace와 match하는지 여부;
- users가 fixture와 target을 구분하는지 여부;
- provider host behavior.

## Remaining risks

- Users may paste incomplete output.
- Users may omit workspace context.
- Users may confuse fixture transcripts with target workspace state.
- Users may treat model-drafted text as CLI state.
- Trusted runner context may differ from target workspace.
- Provider host behavior remains unverified.
- No-terminal remains assisted, not deterministic.

## Recommended next task

Selected next task: F. release readiness sweep for v0.5 reliability docs.

Why: 최근 reliability docs는 proof capture, change control, `LOCK-STALE`,
amend/relock, model workspace wording, no-terminal stale-lock examples,
transcript quality를 다룬다. Release readiness sweep은 kernel behavior를 바꾸거나
downstream execution work를 시작하지 않고 cross-links, bilingual parity, static
checks, status boundaries를 한 묶음으로 verify할 수 있다.

## Next task prompt

```text
Goal:
Run a release readiness sweep for v0.5 reliability docs.

This is a documentation and validation-boundary task. Do not change Go implementation, lock semantics, stale-lock hash semantics, CLI command behavior, validation semantics, prompt compilation behavior, release automation, Homebrew availability, or downstream execution behavior.

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
- ni run compiles a bounded handoff prompt from the current lock; it does not execute downstream work.

Read first:
- AGENTS.md
- README.md
- README.ko.md
- docs/51_POST_RELEASE_ROADMAP.md
- docs/51_POST_RELEASE_ROADMAP.ko.md
- docs/83_CONVERSATION_PROOF_CAPTURE.md
- docs/83_CONVERSATION_PROOF_CAPTURE.ko.md
- docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md
- docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.ko.md
- docs/103_STALE_LOCK_DIAGNOSTIC.md
- docs/103_STALE_LOCK_DIAGNOSTIC.ko.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md
- docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md
- docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md
- docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md
- docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md
- docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md
- docs/107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md
- docs/107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.ko.md
- scripts/demo-check.sh
- scripts/check-skill-packs.sh
- scripts/check-install-docs.py
- scripts/quality.sh

Tasks:
1. Audit the v0.5 reliability docs as a connected set for stale or missing cross-links.
2. Confirm English and Korean companion docs preserve the same boundaries without translating control words away.
3. Confirm docs keep Release binary and Curl installer Available, Homebrew Planned, model workspace packs Experimental, and no-terminal Experimental / assisted.
4. Confirm docs do not claim deterministic no-terminal validation, model workspace readiness authority, fixture relock as project-root relock, generated prompt execution, downstream execution success, implementation quality, benchmark effect size, adoption, cost, or latency improvement.
5. Add only narrow wording or cross-link fixes if evidence shows a drift.
6. Add or adjust a low-noise static check only if a durable risky phrase is missing and the check fits the existing script style.

Do not update:
- .ni/contract.json
- .ni/session.json
- .ni/plan.lock.json

Do not run on the project root:
- ni end
- relock flows
- generated downstream prompts

Validation:
- go run ./cmd/ni status --dir . --proof --next-questions
- bash scripts/check-skill-packs.sh
- python3 scripts/check-install-docs.py
- bash scripts/demo-check.sh
- bash scripts/quality.sh

Required output:
- A concise audit summary naming files changed, if any.
- Confirmation that root lock/session/contract files were not edited.
- Confirmation that no runtime execution, generated prompt execution, Homebrew Available claim, model workspace Available claim, deterministic no-terminal validation claim, or benchmark overclaim was added.
- Exactly one recommended next task with a complete executable Codex prompt.
```
