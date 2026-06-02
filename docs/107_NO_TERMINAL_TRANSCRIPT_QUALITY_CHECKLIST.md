# No-Terminal Transcript Quality Checklist

## Current status

- No-terminal method is Experimental / assisted.
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- Model-only text is draft-only.
- Exact CLI output from a trusted runner is required before claiming readiness,
  lock freshness, relock, hash verification, or bounded handoff compilation.

No-terminal can help draft planning amendments, prepare commands for a trusted
runner, and explain exact pasted CLI output. It cannot provide deterministic
validation by itself, determine readiness without `ni status` output, lock or
relock without `ni end` output, verify lock freshness or hashes without exact
CLI output from a trusted runner, compile the authoritative bounded handoff
without `ni run` output, or execute downstream work.

`ni status` determines readiness and surfaces stale-lock diagnostics. `ni end`
is the CLI-authoritative lock/relock step. `ni run` compiles a bounded handoff
prompt from the current lock. `ni run` does not execute downstream work.

## Transcript quality levels

| Level | Description | Can support | Cannot support | Required caveat |
| --- | --- | --- | --- | --- |
| model-only draft | Text drafted by a model without actual CLI output. | Draft planning text, changed-intent review, assumptions, open questions, and conceptual explanations. | CLI-state claims, readiness, lock freshness, relock, hash verification, bounded handoff compilation, or downstream execution claims. | "This is a draft interpretation until exact CLI output is available." |
| unusable transcript | Ambiguous, invented, stale, wrong-workspace, or too incomplete to identify the command and output. | At most a request for a better transcript. | Readiness, `LOCK-STALE`, relock, prompt compilation, target workspace state, implementation correctness, or product readiness. | Treat as no CLI evidence. |
| partial transcript | Some pasted output is present, but workspace, command, status, timestamp, or relevant surrounding lines are missing. | Limited interpretation of only the visible lines. | Whole-command result, lock freshness, relock success, later hash state, or target workspace state if workspace identity is missing. | Name the missing fields before making any claim. |
| pasted CLI output | User-pasted exact output from a command, with enough command context to know what was run. | Claims directly shown by that output, such as the readiness or warning state printed by `ni status`. | Claims about commands not shown, files changed after the command, other workspaces, implementation quality, or downstream execution success. | Scope every claim to the pasted command and time. |
| trusted runner transcript | Command, workspace context, and exact CLI output from a trusted runner, with exit status or freshness notes when available. | CLI-state claims for the shown command in the shown workspace, if complete enough. | Product readiness, benchmark effect size, provider host behavior, or claims outside the transcript. | Name whether the workspace is target or fixture. |
| fixture transcript | A transcript from a temporary or isolated test/demo workspace. | fixture transcript supports only fixture claims. | Target workspace readiness, target relock, target lock freshness, or target `ni run` compilation. | Say "fixture-only" near any claim. |
| target-workspace transcript | A transcript from the actual project workspace being discussed. | Target workspace CLI-state claims shown by the transcript, if complete enough. | Later state after edits, unrelated workspaces, implementation correctness, downstream execution success, product readiness, benchmark effect size, adoption, cost, or latency improvement. | It proves only the workspace, command, output, and moment shown. |

In this checklist, usable pasted CLI output means the `pasted CLI output` row
when command context is complete enough. Usable trusted runner transcript means
the `trusted runner transcript` row when the minimum fields below are present.
Fixture-only transcript and target-workspace transcript are scope labels that
must stay visible in summaries.

## Minimum trusted-runner transcript fields

A usable trusted runner transcript should include:

- Workspace identity or path, with privacy-safe redaction allowed.
- Whether the workspace is target or fixture.
- Command run.
- Exact command output.
- Exit status if available.
- Timestamp or freshness note if available.
- Whether files were changed before or after the command.
- Whether `ni end` was run on the target workspace or only inside a fixture.
- Whether `.ni/contract.json`, `.ni/session.json`, or `.ni/plan.lock.json`
  changed.
- If `ni run` output is included, confirmation that the prompt was compiled but
  not executed.

Do not require sensitive absolute paths. Privacy-safe redaction is acceptable
as long as target versus fixture identity remains clear.

## Claim boundary matrix

| Evidence | Allowed claim | Forbidden claim | Next action |
| --- | --- | --- | --- |
| no CLI output | The model can draft planning text or explain the protocol conceptually. | "The model verified readiness without CLI output." | Ask for `ni status --proof --next-questions` output from the target workspace. |
| partial `ni status` output | Only the visible lines can be interpreted. | "The partial transcript proves the lock is current." | Request the full command, workspace type, exact output, and freshness note. |
| complete `ni status` output | The readiness and warnings shown by that output for that workspace and time. | Relock success, current hash state after later edits, or `ni run` safety unless shown. | If `READY` and no warnings are shown, ask for `ni end` only after accepted intent is confirmed. |
| `ni status` output with `LOCK-STALE` | `LOCK-STALE` means the existing lock no longer matches current planning inputs. | Implementation failure, product unreadiness, benchmark failure, downstream execution failure, or completed relock. | Review changed intent, then run `ni end` if readiness allows. |
| `ni end` output | The lock or relock behavior shown by that output. | Later lock freshness, downstream handoff compilation, or implementation verification. | If lock is current, run `ni run --max-chars 4000` to compile the bounded handoff. |
| `ni run` output | Bounded prompt compilation shown in that output. | "`ni run` verified the implementation" or "the generated prompt was executed by ni." | Confirm whether the prompt was only compiled and not executed. |
| fixture transcript | Fixture-only CLI behavior. | "The fixture relock proves the project root was relocked." | Run or request the same needed command in the target workspace. |
| target-workspace transcript | Target workspace CLI-state claims shown by the complete transcript. | Claims about later edits, unrelated workspaces, implementation correctness, product readiness, benchmark effect size, adoption, cost, or latency. | Continue only with claims scoped to the shown command and time. |
| validation script transcript | Script behavior and any fixture commands it clearly ran. | Project-root relock unless the transcript explicitly shows `ni end` on the project root. | Separate fixture runs from root workspace state in the summary. |

No transcript proves implementation correctness, downstream execution success,
product readiness, benchmark effect size, adoption, cost, or latency
improvement.

## Core stale-lock recovery flow

```text
review changed intent
-> ni status --proof --next-questions
-> ni end
-> ni run --max-chars 4000
```

In no-terminal assisted terms:

- The model may draft the changed intent review.
- The model may prepare commands for a trusted runner.
- The user or trusted runner must run `ni status`.
- The user or trusted runner must run `ni end` if relock is needed.
- The user or trusted runner must run `ni run` to compile the bounded handoff.
- The model may explain exact pasted CLI output.
- The model must not invent CLI results.

## Transcript examples

### Example 1: Unusable model-invented transcript

| Field | Example |
| --- | --- |
| Transcript quality | model-only draft; unusable transcript for CLI-state claims. |
| Available evidence | "The model says ni status was READY" with no command output. |
| Claim allowed | The model drafted an interpretation. |
| Claim not allowed | Readiness, lock freshness, relock, hash verification, or bounded handoff compilation. |
| Missing information | Workspace, command, exact output, exit status, freshness, file-change context. |
| Required next action | Ask a trusted runner to run `ni status --proof --next-questions` in the target workspace. |
| Boundary note | Model-only draft cannot support CLI-state claims. |

### Example 2: Partial pasted ni status output with missing workspace context

| Field | Example |
| --- | --- |
| Transcript quality | partial transcript. |
| Available evidence | Pasted lines show `NI Intent Readiness: READY`, but not workspace, command, warnings, or timestamp. |
| Claim allowed | The visible pasted line says `READY`. |
| Claim not allowed | Target workspace readiness, absence of `LOCK-STALE`, current lock freshness, or safe `ni run`. |
| Missing information | Target versus fixture identity, full command, full output, warnings section, file changes, freshness. |
| Required next action | Request a complete target-workspace transcript with the minimum fields. |
| Boundary note | Incomplete pasted output can support only limited interpretation. |

### Example 3: Usable pasted ni status output showing READY and no warnings

| Field | Example |
| --- | --- |
| Transcript quality | usable pasted CLI output. |
| Available evidence | Target workspace, command `ni status --proof --next-questions`, exact output showing `NI Intent Readiness: READY`, no blockers, no deferrals, no warnings. |
| Claim allowed | The pasted `ni status` output supports the readiness and warning state shown for that workspace at that time. |
| Claim not allowed | Lock/relock happened, `ni run` compiled a prompt, or later edits remain covered. |
| Missing information | `ni end` output and `ni run` output, if those claims are needed. |
| Required next action | If the accepted plan is confirmed and lock/relock is needed, run `ni end`. |
| Boundary note | Pasted `ni status` output can support only the readiness/stale-lock state shown in that output. |

### Example 4: Usable pasted ni status output showing LOCK-STALE

| Field | Example |
| --- | --- |
| Transcript quality | usable pasted CLI output. |
| Available evidence | Target workspace `ni status --proof --next-questions` output includes `LOCK-STALE`. |
| Claim allowed | The existing lock is stale because current planning inputs differ from `.ni/plan.lock.json`. |
| Claim not allowed | Relock success, implementation failure, product unreadiness, benchmark failure, downstream execution failure, or safe `ni run`. |
| Missing information | Changed-intent review, `ni end` output, and later `ni run --max-chars 4000` output. |
| Required next action | Review changed intent, then run `ni end` if readiness allows and the user confirms. |
| Boundary note | `LOCK-STALE` is a CLI diagnostic, not an implementation verdict. |

### Example 5: Fixture transcript that must not be claimed as target workspace state

| Field | Example |
| --- | --- |
| Transcript quality | fixture-only transcript. |
| Available evidence | A validation script creates a temporary workspace, runs `ni end`, then runs `ni run`. |
| Claim allowed | The fixture flow exercised lock and bounded prompt compilation inside that fixture. |
| Claim not allowed | The project root was relocked, the target lock is current, or target `ni run` is safe. |
| Missing information | Target workspace transcript for the same command sequence. |
| Required next action | Run or request target-workspace `ni status`, then `ni end` and `ni run` only if needed. |
| Boundary note | A fixture transcript supports fixture claims, not target workspace claims. |

### Example 6: Target workspace transcript showing relock and subsequent ni run output

| Field | Example |
| --- | --- |
| Transcript quality | usable trusted runner transcript; target-workspace transcript. |
| Available evidence | Target workspace transcript shows changed-intent review note, `ni status --proof --next-questions`, `ni end`, and `ni run --max-chars 4000` output. |
| Claim allowed | `ni end` was the CLI-authoritative relock step shown in the target workspace, and `ni run` compiled a bounded handoff prompt from the current lock. |
| Claim not allowed | The generated prompt was executed, implementation quality was verified, or downstream execution succeeded. |
| Missing information | None for explaining the shown transcript; later edits need fresh proof. |
| Required next action | Do not execute the generated prompt as part of `ni`; any downstream execution is separate. |
| Boundary note | `ni run` compiles a bounded handoff prompt from the current lock; it does not execute downstream work. |

### Example 7: ni run output pasted without preceding lock freshness evidence

| Field | Example |
| --- | --- |
| Transcript quality | partial transcript for lock-freshness claims; pasted CLI output for the visible `ni run` result. |
| Available evidence | User pastes `ni run --max-chars 4000` output but no current `ni status` or `ni end` transcript. |
| Claim allowed | The pasted `ni run` output can support only bounded prompt compilation shown in that output. |
| Claim not allowed | Current readiness before the run, why the lock was current, or that later edits are covered. |
| Missing information | Workspace type, prior status, lock/relock proof, file-change timing, exit status. |
| Required next action | Request the preceding `ni status` and, if relevant, `ni end` transcript for the same workspace. |
| Boundary note | Pasted `ni run` output can support only bounded prompt compilation shown in that output. |

### Example 8: Validation script transcript with fixture ni end runs

| Field | Example |
| --- | --- |
| Transcript quality | validation script transcript with fixture-only subclaims unless root commands are explicit. |
| Available evidence | `bash scripts/quality.sh` or `bash scripts/demo-check.sh` output shows tests passed and temporary fixture commands ran. |
| Claim allowed | The validation script passed and any fixture `ni end` runs happened only where the script says they happened. |
| Claim not allowed | Project-root `.ni/plan.lock.json` changed, project-root relock succeeded, or target workspace handoff was compiled unless explicitly shown. |
| Missing information | Root lockfile diff and root command transcript if root relock is claimed. |
| Required next action | State fixture/root distinction, then inspect git diff or run target workspace commands if target claims are needed. |
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

Privacy-safe redaction is allowed. Do not paste secrets. If an absolute path is
sensitive, replace private segments with `[redacted]` while keeping target
versus fixture identity clear.

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
- `scripts/demo-check.sh`

No skill files were changed. Existing skill wording was sufficient because the
current skill pack checks already preserve `LOCK-STALE`, amended drafting,
no-relock, no-lockfile-update, and the recovery order. This task adds transcript
quality guidance for users and trusted runners rather than changing model
workspace skill behavior.

## Validation and current coverage

Current checks cover:

- skill pack boundary wording through `scripts/check-skill-packs.sh`;
- no-terminal stale-lock docs through `scripts/demo-check.sh`;
- this transcript-quality checklist through `scripts/demo-check.sh`;
- install docs consistency through `scripts/check-install-docs.py`;
- demo coverage through `scripts/demo-check.sh`;
- quality checks through `scripts/quality.sh`;
- stale-lock tests from the prior implementation.

The following remain manually audited:

- transcript completeness;
- whether pasted output is exact;
- whether trusted runner workspace matches target workspace;
- whether users distinguish fixture from target;
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

Why: the recent reliability docs now cover proof capture, change control,
`LOCK-STALE`, amend/relock, model workspace wording, no-terminal stale-lock
examples, and transcript quality. A release readiness sweep can verify the
cross-links, bilingual parity, static checks, and status boundaries as a set
without changing kernel behavior or starting downstream execution work.

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
