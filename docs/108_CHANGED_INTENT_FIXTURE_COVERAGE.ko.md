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

Goal은 representative changed-intent cases를 test하고 false positives를 피하는
것이다. Fixture coverage는 current lockable inputs 중 어떤 변경이 existing
`.ni/plan.lock.json`을 stale하게 만드는지, 어떤 non-lockable changes가 stale하게
만들지 않는지, 그리고 project-root lock을 건드리지 않고 `ni status`, `ni run`,
fixture relock이 어떻게 동작하는지 보여준다.

## Lock-input notes

Current implementation은 lockable planning input을 다음으로 취급한다:

- `.ni/contract.json`;
- `readiness.RequiredDocs(root)`가 반환하고 `.ni/plan.lock.json`에 recorded된
  required `docs/plan/**` files.

`lock.Create`는 이 files를 `.ni/plan.lock.json`에 record한다. `lock.Verify`는
recorded files를 check하고 mismatches를 report한다. `.ni/session.json`은
source-of-truth order에서 locked docs 아래에 listed되어 있지만, 현재 file hash list의
일부는 아니다. 이 문서는 그 behavior를 바꾸지 않는다.

Lockfile의 recorded required docs에 없는 새 `docs/plan/**` files는 current semantics
아래에서 existing lock을 stale하게 만들지 않는다. 이것은 false-positive guard로
tested된 것이며, future planning file이 모두 ignored되어야 한다는 claim이 아니다.

## Fixture matrix

| Case | Changed input | Expected status behavior | Expected `ni run` behavior | Relock behavior | Covered by test? |
| --- | --- | --- | --- | --- | --- |
| no lock | none; ready fixture에는 아직 `.ni/plan.lock.json`이 없음 | `READY`, no `LOCK-STALE` warning | `ni run`은 lock이 생길 때까지 unavailable | no relock | Yes: `TestStatusStaleLockDiagnostics` |
| current lock | `.ni/plan.lock.json`이 recorded inputs와 match | `READY`, no `LOCK-STALE` warning | bounded handoff prompt를 compile | no relock needed | Yes: `TestStatusStaleLockDiagnostics` |
| docs/plan content change | Lock 뒤 `docs/plan/02_capabilities.md` content changes | `READY` with `LOCK-STALE` warning and first mismatch | `BLOCKED: lock hash mismatch`로 stale handoff refuse | fixture `ni end` clears warning | Yes: `TestChangedIntentFixtureStaleLockMatrix`, `TestStaleLockRunRefusalAndFixtureRecovery` |
| new docs/plan file | Lock 뒤 `docs/plan/99_fixture_note.md` added | Current recorded-file semantics에서는 `READY`, no `LOCK-STALE` warning | bounded handoff prompt를 compile | non-recorded file에는 no relock needed | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| contract planning-input change | `.ni/contract.json` purpose changes with matching project brief | `READY` with `LOCK-STALE` warning and first mismatch at `.ni/contract.json` | stale handoff refuse | fixture relock would update the fixture lock only | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| deferral or requirement change | Lock 뒤 deferral-style contract and docs change | `READY_WITH_DEFERRALS` with `LOCK-STALE` warning | stale handoff refuse | fixture relock would update the fixture lock only | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| session change | Lock 뒤 `.ni/session.json` changes | Current semantics에서는 `READY`, no `LOCK-STALE` warning | bounded handoff prompt를 compile | session-only change에는 no relock needed | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| non-lockable file change | Fixture `README.md` changes after lock | `READY`, no `LOCK-STALE` warning | bounded handoff prompt를 compile | no relock needed | Yes: `TestChangedIntentFixtureStaleLockMatrix` |
| fixture relock recovery | Lockable planning input changes, then `ni end` runs in fixture | stale warning disappears in fixture | `ni run --max-chars 4000` compiles again in fixture | fixture relock only | Yes: `TestStaleLockRunRefusalAndFixtureRecovery` |
| validation-script fixture ni end | Fixture `ni end` runs while project-root lock is read-only evidence | Project-root state는 fixture output에서 inferred되지 않음 | no project-root handoff claim | project-root lockfile bytes remain unchanged | Yes: `TestFixtureRelockDoesNotModifyProjectRootLock` |

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
- Project root의 `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`은
  unchanged 상태로 남아야 한다.
- Project-root lockfile test는 fixture relock 전후 root lockfile을 read하고 fixture
  relock이 그것을 바꾸지 않았는지 verify한다.

## Validation surface

This pass에서 사용한 validation commands:

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

`GOCACHE=/private/tmp/ni-go-cache`는 local Go build cache permission issue에 대한
environment-specific workaround다.

## Remaining risks

- Not every possible planning-input mutation is covered.
- Future lock-input changes need corresponding fixture updates.
- Fixture paths can differ from user workspaces.
- User interpretation still depends on docs and transcript quality.
- Provider host behavior remains unverified.

## Recommended next task

Selected next task: A. release readiness sweep for v0.5 reliability docs.

Why: changed-intent fixture coverage가 이제 main stale-lock path, non-lockable false
positives, fixture relock recovery, project-root safety에 대한 tests와 docs를 갖췄다.
다음 유용한 step은 release-facing claim hardening 전에 v0.5 reliability docs를 연결된
set으로 review하는 것이다.

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
