# Change-Control UX Audit

## Current status

This audit reviews ni's locked-plan change-control user experience around stale
locks, amended intent, relock expectations, and downstream handoff safety.

Follow-up: Task 183 implemented the focused `LOCK-STALE` diagnostic documented
in [`103_STALE_LOCK_DIAGNOSTIC.md`](103_STALE_LOCK_DIAGNOSTIC.md). This audit
remains the historical finding that motivated the diagnostic.

Current factual boundaries:

- release binary: Available for verified v0.4.0 assets;
- curl installer: Available for verified v0.4.0 assets;
- Homebrew: Planned / v0.5 candidate only;
- model workspace packs: Experimental as a broad product path;
- no-terminal method: Experimental / assisted;
- `ni-kernel`: pre-runtime Project Intent Compiler;
- runtime execution, shell adapters, Codex exec adapters, queues, PR
  automation, release automation, and downstream execution layers: not
  included.

The current CLI readiness proof for this repository is:

```text
NI Intent Readiness: READY
Blockers: None.
Deferrals: None.
Warnings: None.
```

This audit did not run `ni end`, run a relock flow, edit `.ni/contract.json`,
edit `.ni/session.json`, edit `.ni/plan.lock.json`, execute generated prompts,
or add runtime behavior.

## Intended change-control model

After `.ni/plan.lock.json` exists, the intended source-of-truth order is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

The locked plan is the authority for downstream handoff until intent changes.
The lock records hashes for `.ni/contract.json` and required `docs/plan/**`
files. `.ni/session.json` is not hashed because it is mutable continuity state
below locked docs.

When `.ni/contract.json` or required `docs/plan/**` changes after lock, the
old lock may no longer represent current intent. Handoff surfaces that require
a valid lock must refuse stale hashes and report `BLOCKED`. A user should then
review the change, amend or re-plan, run readiness again, and relock through
the CLI before producing a new downstream handoff.

Skills are UX; CLI is authority. Skills may help draft amended planning text,
but they do not determine readiness, lock or relock plans, replace `ni status`,
`ni end`, or `ni run`, or update `.ni/plan.lock.json`.

## User-facing lifecycle

```text
pre-lock planning
-> readiness check
-> lock
-> downstream handoff from locked plan
-> intent changes
-> stale/change risk is surfaced
-> user reviews change
-> amend or re-plan
-> readiness check again
-> relock
-> new downstream handoff
```

In this lifecycle, `ni status` is the readiness authority, `ni end` is the
first-lock authority, `ni relock` is the relock authority, and `ni run` is the
bounded prompt compiler. `ni run` does not execute downstream work.

## Current behavior audit

| Surface | Current behavior / wording | Pass? | Risk | Recommendation |
| --- | --- | --- | --- | --- |
| README | Explains `ni status`, `ni end`, and `ni run` as gates and says `ni run` compiles from a valid locked plan without execution. It does not deeply explain the post-lock intent-change lifecycle. | Partial | A reader may understand safe handoff but not know the amend/relock path after changed intent. | Keep README concise; link deeper change-control guidance after the diagnostic work lands. |
| docs/83 | Defines proof capture, CLI authority, no-terminal draft limits, and states stale sides should be repaired or kept blocking. It says `ni run` verifies lock hashes. | Yes | Low. It is proof-capture focused, not a full amend/relock guide. | Keep as-is; point to this audit or future examples if stale-lock wording expands. |
| docs/101 | Separates planning proof from execution evidence, names stale lock risk, and documents validation surfaces. | Yes | Low. It documents the risk but does not prescribe detailed CLI examples. | Keep as-is; use this audit as the change-control follow-up. |
| ni status --proof --next-questions | Authoritatively reports readiness from docs, contract, sync, decisions, risks, evaluations, and blocker questions. Current implementation does not verify existing `.ni/plan.lock.json` hashes. | Partial | A user can see `READY` after post-lock docs/contract edits and not learn that the existing lock is stale until a lock-dependent command runs. | Add a focused stale-lock diagnostic or warning surface in a follow-up without changing readiness semantics casually. |
| ni end | Runs readiness and writes `.ni/plan.lock.json` when readiness is not `BLOCKED`. Output is short: lock path and status. Existing code does not visibly distinguish first lock from overwriting an existing lock. | Partial | Users may use `ni end` as an accidental relock path, bypassing amendment/relock expectations. | Add a follow-up diagnostic/refusal or guidance around existing-lock cases, preserving CLI authority. |
| ni run | Requires `.ni/plan.lock.json`, verifies locked hashes, refuses stale hashes with `BLOCKED: lock hash mismatch for <path>`, and compiles a bounded prompt only. | Yes | The stale error is correct but terse; it does not tell the user to review, amend, rerun readiness, and relock. | Add clearer changed-intent guidance near stale-lock refusal. |
| skill packs | Codex and Claude pack READMEs and skills preserve Experimental status, "Skills are UX; CLI is authority.", no manual lock editing, no readiness by model judgment, and stale-lock `BLOCKED` behavior. | Yes | Low. Individual skill files are aligned. | Keep existing static checks; update only if new amend/relock wording is added. |
| no-terminal assisted workflow | Documents assisted drafting only. It says draft proof is not deterministic validation and trusted CLI proof is required for readiness, lock, hash proof, and prompt compile. | Yes | Users may still socially treat a model summary as a relock if the next CLI step is not visible. | Keep draft-only labels; add examples only if future no-terminal relock guidance is written. |
| examples | Benchmark, ni-grill, no-terminal, and resolution-path examples generally preserve no-execution, isolated-workspace, `not_measured`, and amendment/relock stop wording. | Yes | Fixture relocks or benchmark locks can be mistaken for project-root relock if evidence is quoted without context. | Keep "isolated workspace only" and "root not relocked" labels near examples. |
| validation scripts | `check-install-docs.py`, `check-skill-packs.sh`, `demo-check.sh`, `quality.sh`, `smoke.sh`, `install-check.sh`, and `release-check.sh` protect status, boundary, skill, install, demo, and release surfaces. | Yes | They do not currently enforce all stale-lock UX copy, and broad prose remains manually audited. | Do not add a broad brittle check in this audit; add focused checks only after the target diagnostic wording is stable. |

## Stale-lock and changed-intent risks

| Risk | Severity | Current mitigation | Gap | Recommended next action |
| --- | --- | --- | --- | --- |
| User edits `docs/plan/**` after lock and assumes the old lock still reflects current intent. | High | `ni run`, export, graph, harness, feedback, and pressure paths verify lock hashes when they require a valid lock. | `ni status` can still report readiness without surfacing existing-lock staleness. | Add stale-lock warning/diagnostic. |
| User changes `.ni/contract.json` after lock and runs `ni run` without noticing mismatch. | High | `ni run` refuses lock hash mismatch with `BLOCKED`. | Error text is terse and does not explain amendment/relock recovery. | Add clearer stale-lock recovery wording. |
| Model workspace skill drafts new planning text but user assumes CLI state changed. | Medium | Skill docs say skills are UX, CLI is authority, and skills do not lock or determine readiness. | User summaries can still overstate unless every turn keeps proof wording visible. | Preserve static skill checks and proof-capture language. |
| No-terminal assisted workflow is mistaken for deterministic relock. | Medium | No-terminal docs say draft proof is not deterministic validation. | No detailed no-terminal amend/relock example exists. | Consider examples after CLI diagnostics are clearer. |
| `ni-run` handoff is mistaken for fresh validation. | Medium | `ni run` verifies hashes and says prompt compilation only. | It does not rerun readiness or explain that compilation is not fresh planning acceptance. | Clarify in future handoff wording. |
| Proof text is mistaken for implementation evidence. | Medium | docs/83 and docs/101 separate planning proof from execution evidence. | Manual audit remains needed for new examples. | Keep demo and skill checks lightweight. |
| Benchmark evidence is mistaken for downstream execution success. | Medium | docs/97 and benchmark examples preserve `not_measured`, no-execution, and artifact-readiness labels. | Long examples can bury boundaries. | Keep boundary labels near `READY` rows. |
| Root lockfile is manually edited. | High | Skills and docs forbid manual `.ni/plan.lock.json` edits. Lock hash checks catch many downstream uses. | Manual edits to lockfile itself may not be obvious until a command loads or verifies it. | Add targeted diagnostics in future lock validation work. |
| Validation scripts exercise fixture relocks and are mistaken for project-root relock. | Medium | Prior audits and benchmark docs label isolated workspaces and root no-relock boundaries. | Final reports must keep the distinction explicit. | Continue reporting root lockfile diff and command scope. |

## What must remain CLI-authoritative

- Readiness is determined by CLI.
- Locking is performed by CLI.
- Prompt compilation is performed by CLI.
- Skills do not validate readiness.
- Skills do not lock or relock.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- Skills do not update `.ni/plan.lock.json`.
- No-terminal workflows do not provide deterministic validation.
- `ni run` does not execute downstream work.

## Say this / do not say this

| Say this | Do not say this |
| --- | --- |
| "The locked plan is the authority for downstream handoff until intent changes." | "The model relocked the plan." |
| "If intent changes after lock, run the CLI readiness and lock flow again before generating a new handoff." | "The no-terminal workflow deterministically validated the amended plan." |
| "Skills can help draft amended planning text, but the CLI remains the authority." | "`ni run` verified the implementation." |
| "`ni run` compiles a bounded handoff prompt from the locked plan; it does not execute the work." | "The benchmark proves downstream execution quality." |
| "A stale lock or hash mismatch is `BLOCKED` until the plan is reviewed, amended if needed, checked for readiness, and relocked." | "The skill pack replaces CLI validation." |

## Audit findings

| Finding | Severity | Surface | Evidence | Recommendation | Blocks v0.5? |
| --- | --- | --- | --- | --- | --- |
| `ni status` does not surface stale existing-lock state. | High | `ni status --proof --next-questions` | `cmd/ni/main.go` evaluates readiness but does not call lock verification; `docs/commands.md` says `ni status` does not verify lock hashes. | Add a focused stale-lock warning or diagnostic so users can see what is still trustworthy before handoff. | No, but should be next. |
| Existing-lock `ni end` behavior is not clearly separated from relock expectations. | High | `ni end`; `ni relock`; docs | `runEnd` writes via `lock.Create`; `runRelock` implements amendment-gated relock. User-facing `ni end` output is terse. | Decide whether `ni end` should warn/refuse when an existing lock is present, and document recovery. | No, but should be addressed before stronger change-control claims. |
| Stale `ni run` refusal is correct but terse. | Medium | `ni run` | `internal/core/prompt/prompt.go` returns `BLOCKED: lock hash mismatch for <path>`. | Add user-facing recovery wording: review changed intent, amend or re-plan, run readiness, relock, then compile a new handoff. | No. |
| Docs and skills preserve authority boundaries. | Low | docs/83, docs/101, skill packs, no-terminal docs | The audited docs repeatedly state CLI authority, draft-only no-terminal proof, and no downstream execution. | Keep current wording; avoid broad static checks until the next diagnostic text is stable. | No. |

There are material findings, but they are UX and validation-alignment gaps
rather than evidence that downstream handoff currently ignores stale hashes.
The strongest current protection is that lock-dependent handoff commands refuse
hash mismatches. The weakest user-facing surface is that `ni status` can look
healthy while an existing lock may be stale.

## Validation surface

Current scripts check:

- `go run ./cmd/ni status --dir . --proof --next-questions`: repository
  readiness through the CLI;
- `python3 scripts/check-install-docs.py`: install, distribution, Homebrew, and
  model workspace status claim boundaries;
- `bash scripts/check-skill-packs.sh`: skill metadata, Experimental status,
  CLI authority wording, proof-capture markers, and stale-lock boundary
  markers in skill packs;
- `bash scripts/quality.sh`: broad static docs checks, Go formatting/tests, and
  smoke checks;
- `bash scripts/demo-check.sh`: benchmark, no-terminal, ni-grill, and seed-only
  boundaries;
- `bash scripts/smoke.sh`: source CLI smoke behavior;
- `bash scripts/install-check.sh`: local install behavior;
- `bash scripts/release-check.sh`: release readiness surfaces.

Manual audit remains necessary for stale-lock prose, amend/relock examples,
`ni status` versus lock-verification wording, and any future change-control
diagnostic text.

No new static check was added in this task. The current gap is not a missing
phrase alone; it is a user-facing diagnostic design decision.

## Recommended next task

Selected next action: A. implement stale-lock warning/diagnostic.

Why: this audit found that lock-dependent handoff commands refuse stale hashes,
but `ni status --proof --next-questions` can report readiness without telling
the user whether an existing `.ni/plan.lock.json` still matches current
contract and docs. A focused diagnostic is the smallest next step that directly
answers what remains trustworthy, what is stale, and what must be amended or
relocked before downstream handoff.

## Next task prompt

```text
Goal:
Implement a focused stale-lock warning/diagnostic for ni change-control UX.

This is a small v0.5 reliability implementation task. Do not add runtime execution, adapters, queues, PR automation, release automation, or downstream execution behavior.

Context:
ni is ni-kernel: a pre-runtime Project Intent Compiler. The locked plan is the authority for downstream handoff until intent changes. After .ni/plan.lock.json exists, source-of-truth order is:
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history

Task 182 added:
- docs/102_CHANGE_CONTROL_UX_AUDIT.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.ko.md

The audit found:
- lock-dependent handoff commands already refuse stale lock hashes;
- ni status evaluates readiness but does not surface whether an existing lock is stale;
- ni end output/behavior does not clearly distinguish first lock from relock expectations;
- ni run stale-lock refusal is correct but terse.

Read first:
- AGENTS.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.md
- docs/102_CHANGE_CONTROL_UX_AUDIT.ko.md
- docs/42_INTENT_LOCK_PROTOCOL.md
- docs/04_LOCKFILE.md
- docs/36_NI_RUN_HANDOFF.md
- docs/commands.md
- cmd/ni/main.go
- internal/core/lock/lock.go
- internal/core/prompt/prompt.go
- internal/core/amendment/amendment.go
- internal/core/prompt/prompt_test.go
- internal/core/lock/lock_test.go

Before editing, run:
- go run ./cmd/ni status --dir . --proof --next-questions

Implement the smallest coherent diagnostic that tells users:
- whether an existing .ni/plan.lock.json matches current .ni/contract.json and required docs/plan/**;
- which path mismatches first when the lock is stale;
- that stale lock means downstream handoff must stop;
- that the user should review changed intent, amend or re-plan as needed, run readiness again, relock through the CLI, then generate a new handoff;
- that .ni/session.json remains below locked docs and is not hashed.

Preferred implementation direction:
- Add lock-state reporting to ni status --proof and/or --next-questions without changing readiness rule semantics unless tests show a deterministic blocker is required.
- Keep JSON output structured if lock diagnostics are added to status results.
- Improve stale ni run error/help text only if it stays narrow and testable.
- Decide whether ni end should warn or refuse when an existing lock is present; if semantics change, update docs and tests. Do not silently overwrite an existing project-root lock in tests.

Do not:
- run ni end on the project root;
- run relock on the project root;
- manually edit .ni/plan.lock.json;
- edit .ni/contract.json or .ni/session.json unless the implementation genuinely requires fixture updates outside the project root;
- execute generated prompts;
- add task-runner, SPEC runner, shell adapter, Codex exec adapter, queue, PR automation, release automation, execution evidence loop, TUI, or web UI behavior;
- mark Homebrew Available;
- mark model workspace packs Available as a broad product path;
- claim no-terminal deterministic validation;
- claim benchmark evidence proves downstream execution quality.

Validation:
- gofmt -w . if Go files are touched
- go test ./...
- go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/quality.sh
- bash scripts/demo-check.sh
- bash scripts/smoke.sh
- bash scripts/install-check.sh
- bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- whether Go files were touched
- readiness result
- stale-lock diagnostic behavior added
- validation results
- confirmation that .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified
- confirmation that ni end and relock were not run on the project root
- confirmation that no runtime execution, shell adapter, Codex exec adapter, queue, PR automation, release action, Homebrew Available claim, model workspace Available claim, no-terminal deterministic claim, or benchmark overclaim was added
```
