# No-Terminal Stale-Lock Examples

## Current status

- No-terminal method is Experimental / assisted.
- `LOCK-STALE` is a CLI diagnostic.
- Model-only no-terminal text is draft-only.
- Exact CLI output from a trusted runner is required before claiming readiness,
  lock freshness, relock, hash verification, or bounded handoff compilation.

For transcript quality levels, minimum trusted-runner fields, and a copy-paste
handoff template, see
[`107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md`](107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md).

## What no-terminal can and cannot prove

| Evidence level | Can support | Cannot support | Required caveat |
| --- | --- | --- | --- |
| Model-only draft | Draft planning text, changed-intent interpretation, assumptions, open questions, and conceptual `LOCK-STALE` explanation. | Readiness, lock freshness, relock, hash verification, or bounded handoff compilation. | "This is a draft interpretation until exact CLI output is available." |
| Pasted CLI output | Statements limited to the pasted command output, such as explaining a visible `LOCK-STALE` warning from `ni status`. | Claims beyond the shown command, global workspace state, relock success, hash freshness after later edits, or implementation correctness. | Explain only the exact command result shown; do not invent missing CLI lines. |
| Trusted runner transcript | Actual `ni status`, `ni end`, or `ni run` behavior in the target workspace or an explicitly isolated fixture. | Product implementation quality, downstream execution quality, host/provider behavior, or evidence from another workspace. | Name the workspace or fixture and keep claims scoped to that transcript. |

## Core stale-lock recovery flow

```text
review changed intent
-> ni status --proof --next-questions
-> ni end
-> ni run --max-chars 4000
```

In no-terminal assisted terms:

- The model may draft the changed intent review.
- The user or trusted runner must run `ni status --proof --next-questions`.
- The user or trusted runner must run `ni end` if relock is needed and readiness
  allows it.
- The user or trusted runner must run `ni run --max-chars 4000` to compile the
  bounded handoff.
- The model may explain exact pasted CLI output, but must not invent CLI
  results.

## Example 1: Model-only planning draft after suspected changed intent

| Field | Example |
| --- | --- |
| Scenario | A no-terminal user says the accepted planning intent may have changed after lock, but provides no CLI output. |
| Available evidence | Model-only draft and user-described changed intent. |
| What the model may say | "This is a draft interpretation until exact CLI output is available. The changed intent appears to affect planning text, so review it before asking a trusted runner to run `ni status --proof --next-questions`." |
| What the model must not say | "The plan is READY", "the lock is stale", "the lock is current", or "the model relocked the plan." |
| Required next CLI action | A trusted runner runs `ni status --proof --next-questions` in the target workspace. |
| Whether `ni run` is safe before relock | No. Without exact CLI output and a current lock, no-terminal cannot claim `ni run` is safe. |
| Whether `ni run` is safe after relock | Only after `ni end` succeeds and exact `ni run --max-chars 4000` output is available. |
| Boundary note | Model-only text may draft changed intent, but it does not provide deterministic validation. |

## Example 2: User pastes LOCK-STALE from ni status

| Field | Example |
| --- | --- |
| Scenario | A user pastes `WARNING: LOCK-STALE existing lock is stale. Current planning inputs differ from .ni/plan.lock.json.` from `ni status --proof --next-questions`. |
| Available evidence | Pasted CLI output from `ni status`; no pasted `ni end` or `ni run` output. |
| What the model may say | "`LOCK-STALE` means the existing lock no longer matches current planning inputs. Review changed intent, then use the CLI recovery flow." |
| What the model must not say | "The model verified hashes", "the plan was relocked", or "`ni run` can compile now." |
| Required next CLI action | Review changed intent, then run `ni end` if readiness allows and the user confirms the accepted plan. |
| Whether `ni run` is safe before relock | No. A stale lock must not be used for trusted handoff. |
| Whether `ni run` is safe after relock | Only after trusted `ni end` output shows the current lock was written, then `ni run --max-chars 4000` runs against that lock. |
| Boundary note | Pasted `ni status` output supports only the status shown, not relock or bounded handoff compilation. |

## Example 3: Trusted runner provides exact ni status output

| Field | Example |
| --- | --- |
| Scenario | The user cannot run a terminal locally, but a teammate provides an exact transcript of `ni status --proof --next-questions` from the target workspace. |
| Available evidence | Trusted runner transcript for `ni status`; workspace path and timestamp are included. |
| What the model may say | "The transcript can support the readiness and warning state for that workspace at that time. If it shows `LOCK-STALE`, the next step remains `ni end` after changed intent review." |
| What the model must not say | "The provider host is verified", "the implementation is ready", or "later edits are covered by this old transcript." |
| Required next CLI action | If status allows and relock is needed, the trusted runner runs `ni end` after user confirmation. |
| Whether `ni run` is safe before relock | No, if the transcript shows `LOCK-STALE` or no current lock proof exists. |
| Whether `ni run` is safe after relock | Yes only as prompt compilation, after trusted `ni end` and trusted `ni run --max-chars 4000` output. |
| Boundary note | A trusted runner transcript is stronger than model-only text, but it is still scoped to the workspace and command shown. |

## Example 4: Model workspace skill drafts amended planning text

| Field | Example |
| --- | --- |
| Scenario | `ni-start` in a model workspace drafts amended `docs/plan/**` text after a stale-lock warning is pasted. |
| Available evidence | Pasted `LOCK-STALE` output plus model-drafted amendment text. |
| What the model may say | "This draft may help review changed intent. A trusted runner must run `ni status --proof --next-questions` after accepted planning edits are applied." |
| What the model must not say | "The skill updated `.ni/plan.lock.json`", "the skill relocked the plan", or "the no-terminal workflow deterministically validated the amended plan." |
| Required next CLI action | Apply accepted planning changes, then run `ni status --proof --next-questions`; if allowed, run `ni end`. |
| Whether `ni run` is safe before relock | No. Skill drafting is not lock freshness or hash verification. |
| Whether `ni run` is safe after relock | Only after the CLI lock is current and exact `ni run --max-chars 4000` output exists. |
| Boundary note | Skills are UX; CLI is authority. Skills may draft amendments and explain pasted CLI output, but they do not lock or relock. |

## Example 5: No CLI output is available

| Field | Example |
| --- | --- |
| Scenario | The user asks whether the lock is current, but no CLI output or trusted runner transcript is available. |
| Available evidence | Model-only conversation and repository text, not current CLI output. |
| What the model may say | "No-terminal cannot claim lock freshness without exact CLI output from a trusted runner." |
| What the model must not say | "The pasted draft proves the lock is current", "the no-terminal workflow verified hashes", or "`ni run` is safe." |
| Required next CLI action | Run or request `ni status --proof --next-questions`. |
| Whether `ni run` is safe before relock | No. There is no deterministic current-lock proof. |
| Whether `ni run` is safe after relock | Only after trusted `ni end` and `ni run --max-chars 4000` output are available. |
| Boundary note | Absence of CLI output means the model can only provide draft guidance, not CLI-state claims. |

## Example 6: Trusted runner relocks and provides ni run output

| Field | Example |
| --- | --- |
| Scenario | A trusted runner reviews changed intent, runs `ni status --proof --next-questions`, runs `ni end`, and then provides exact `ni run --max-chars 4000` output. |
| Available evidence | Trusted runner transcript for status, relock, and bounded prompt compilation. |
| What the model may say | "`ni run` compiled a bounded handoff prompt from the current lock shown in this transcript; it did not execute downstream work." |
| What the model must not say | "`ni run` verified the implementation", "the benchmark proves downstream execution quality", or "the generated prompt was executed by ni." |
| Required next CLI action | None for explaining the pasted transcript; any later changed intent restarts the recovery flow. |
| Whether `ni run` is safe before relock | No. The transcript's safe compilation claim only applies after the recorded relock. |
| Whether `ni run` is safe after relock | Yes, as bounded prompt compilation only, when exact `ni run --max-chars 4000` output is present. |
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

No skill files were changed. Existing package and repo-local skill wording was
sufficient because the prior model workspace stale-lock verification pass
already made `LOCK-STALE`, amended drafting, no-relock, no-lockfile-update, and
the recovery order explicit and statically checked.

## Validation and current coverage

Current checks cover:

- skill pack boundary wording through `scripts/check-skill-packs.sh`;
- install docs consistency through `scripts/check-install-docs.py`;
- demo coverage through `scripts/demo-check.sh`;
- quality checks through `scripts/quality.sh`;
- stale-lock tests from the prior implementation for status warnings, current
  locks, stale `ni run` refusal, and fixture relock recovery.

The following remain manually audited:

- no-terminal user interpretation;
- trusted runner transcript quality;
- provider host behavior;
- whether users distinguish draft text from CLI output.

## Remaining risks

- Users may paste incomplete CLI output.
- Users may treat model-drafted text as CLI state.
- Trusted runner context may differ from the target workspace.
- Provider host behavior remains unverified.
- No-terminal remains assisted, not deterministic.

## Recommended next task

Selected next task: F. no-terminal transcript quality checklist.

Why: the examples now explain the stale-lock path, but the next weak point is
the quality of pasted or delegated transcripts. A checklist can help users and
trusted runners include command, workspace, timestamp, output, and scope without
changing CLI behavior or claiming deterministic no-terminal validation.

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
