# Changed-Intent Fixture Coverage

## Current status

- `LOCK-STALE` is implemented as a CLI diagnostic.
- `ni status` can surface stale existing locks.
- `ni run` refuses stale handoff.
- Fixture relock is separate from project-root relock.
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.

## Coverage goal

The goal is to test representative changed-intent cases and avoid false
positives. The fixture coverage shows which current lockable inputs stale an
existing `.ni/plan.lock.json`, which non-lockable changes do not stale it, and
how `ni status`, `ni run`, and fixture relock behave without touching the
project-root lock.

## Lock-input notes

Current implementation treats the lockable planning input as:

- `.ni/contract.json`;
- the required `docs/plan/**` files returned by `readiness.RequiredDocs(root)`
  and recorded in `.ni/plan.lock.json`.

`lock.Create` records those files in `.ni/plan.lock.json`. `lock.Verify` checks
the recorded files and reports mismatches. `.ni/session.json` is listed below
locked docs in the source-of-truth order, but it is not currently part of the
file hash list. This document does not change that behavior.

New `docs/plan/**` files that are not in the lockfile's recorded required docs
do not stale the existing lock under current semantics. That is tested as a
false-positive guard, not as a claim that every future planning file should be
ignored.

## Fixture matrix

| Case | Changed input | Expected status behavior | Expected `ni run` behavior | Relock behavior | Covered by test? |
| --- | --- | --- | --- | --- | --- |
| no lock | none; ready fixture has no `.ni/plan.lock.json` yet | `READY`, no `LOCK-STALE` warning | `ni run` remains unavailable until a lock exists | no relock | Yes: `TestStatusStaleLockDiagnostics` |
| current lock | `.ni/plan.lock.json` matches recorded inputs | `READY`, no `LOCK-STALE` warning | compiles bounded handoff prompt | no relock needed | Yes: `TestStatusStaleLockDiagnostics` |
| docs/plan content change | `docs/plan/02_capabilities.md` content changes after lock | `READY` with `LOCK-STALE` warning and first mismatch | refuses stale handoff with `BLOCKED: lock hash mismatch` | fixture `ni end` clears warning | Yes: `TestChangedIntentFixtureStaleLockMatrix`, `TestStaleLockRunRefusalAndFixtureRecovery` |
| new docs/plan file | `docs/plan/99_fixture_note.md` is added after lock | `READY`, no `LOCK-STALE` warning under current recorded-file semantics | compiles bounded handoff prompt | no relock needed for this non-recorded file | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| contract planning-input change | `.ni/contract.json` purpose changes with matching project brief | `READY` with `LOCK-STALE` warning and first mismatch at `.ni/contract.json` | refuses stale handoff | fixture relock would update the fixture lock only | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| deferral or requirement change | deferral-style contract and docs change after lock | `READY_WITH_DEFERRALS` with `LOCK-STALE` warning | refuses stale handoff | fixture relock would update the fixture lock only | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| session change | `.ni/session.json` changes after lock | `READY`, no `LOCK-STALE` warning under current semantics | compiles bounded handoff prompt | no relock needed for session-only change | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| non-lockable file change | fixture `README.md` changes after lock | `READY`, no `LOCK-STALE` warning | compiles bounded handoff prompt | no relock needed | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| fixture relock recovery | lockable planning input changes, then `ni end` runs in the fixture | stale warning disappears in the fixture | `ni run --max-chars 4000` compiles again in the fixture | fixture relock only | Yes: `TestStaleLockRunRefusalAndFixtureRecovery` |
| validation-script fixture ni end | fixture `ni end` runs while project-root lock is read-only evidence | project-root state is not inferred from fixture output | no project-root handoff claim | project-root lockfile bytes remain unchanged | Yes: `TestFixtureRelockDoesNotModifyProjectRootLock` |

## Tests added

Added or expanded:

- `TestChangedIntentFixtureStaleLockMatrix`
- `TestFixtureRelockDoesNotModifyProjectRootLock`

Existing coverage retained:

- `TestStatusStaleLockDiagnostics`
- `TestStaleLockRunRefusalAndFixtureRecovery`

## What this proves

- `ni` detects representative stale-lock cases covered by fixtures.
- `ni` avoids at least one non-lockable false positive.
- `ni run` refuses stale handoff in covered cases.
- Fixture relock clears the stale diagnostic inside the fixture.

## What this does not prove

- implementation correctness;
- downstream execution success;
- product readiness;
- benchmark effect size;
- adoption/cost/latency improvement;
- global model workspace behavior;
- Homebrew availability;
- no-terminal deterministic validation;
- project-root relock when only fixture relock was run.

## Project-root safety

- Tests use temporary fixtures.
- Fixture `ni end` does not mean project root was relocked.
- `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` on the
  project root must remain unchanged.
- The project-root lockfile test reads the root lockfile before and after a
  fixture relock and verifies that fixture relock did not change it.

## Validation surface

Validation commands used for this pass:

```bash
gofmt -w cmd/ni/main_test.go
GOCACHE=/private/tmp/ni-go-cache go test ./cmd/ni -run 'Test(StatusStaleLockDiagnostics|StaleLockRunRefusalAndFixtureRecovery|ChangedIntentFixtureStaleLockMatrix|FixtureRelockDoesNotModifyProjectRootLock)'
GOCACHE=/private/tmp/ni-go-cache go test ./...
go run ./cmd/ni status --dir . --proof --next-questions
python3 scripts/check-install-docs.py
bash scripts/check-skill-packs.sh
bash scripts/demo-check.sh
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
```

`GOCACHE=/private/tmp/ni-go-cache` is an environment-specific workaround for a
local Go build cache permission issue.

## Remaining risks

- Not every possible planning-input mutation is covered.
- Future lock-input changes need corresponding fixture updates.
- Fixture paths can differ from user workspaces.
- User interpretation still depends on docs and transcript quality.
- Provider host behavior remains unverified.

## Recommended next task

Selected next task: A. release readiness sweep for v0.5 reliability docs.

Why: changed-intent fixture coverage now has tests and documentation for the
main stale-lock path, non-lockable false positives, fixture relock recovery,
and project-root safety. The next useful step is to review the v0.5 reliability
docs as a connected set before any release-facing claim hardening.

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
- Fixture relock must not be claimed as project-root relock.

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
- docs/108_CHANGED_INTENT_FIXTURE_COVERAGE.md
- docs/108_CHANGED_INTENT_FIXTURE_COVERAGE.ko.md
- scripts/demo-check.sh
- scripts/check-skill-packs.sh
- scripts/check-install-docs.py
- scripts/quality.sh

Tasks:
1. Audit the v0.5 reliability docs as a connected set for stale or missing cross-links.
2. Confirm English and Korean companion docs preserve the same boundaries without translating control words away.
3. Confirm docs keep Release binary and Curl installer Available, Homebrew Planned, model workspace packs Experimental, and no-terminal Experimental / assisted.
4. Confirm docs do not claim deterministic no-terminal validation, model workspace readiness authority, fixture relock as project-root relock, generated prompt execution, downstream execution success, implementation quality, benchmark effect size, adoption, cost, or latency improvement.
5. Add only narrow wording or cross-link fixes if evidence shows drift.
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
- Confirmation that no runtime execution, generated prompt execution, Homebrew Available claim, model workspace Available claim, deterministic no-terminal validation claim, fixture relock as project-root relock claim, or benchmark overclaim was added.
- Exactly one recommended next task with a complete executable Codex prompt.
```
