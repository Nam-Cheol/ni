# v0.5 Reliability Docs Release Readiness Sweep

## Current status

- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Sweep goal

This sweep checks the v0.5 reliability documentation set for consistency,
navigation, Korean companion drift, validation-surface coverage, and overclaim
boundaries before the project moves to the next v0.5 direction. It is a
documentation and claim-boundary audit only; it does not change lock semantics,
stale-lock hash semantics, CLI behavior, runtime boundaries, or implementation.

## Reliability doc inventory

| Doc | Purpose | Primary boundary | Korean companion? | Linked from roadmap? | Notes |
| --- | --- | --- | --- | --- | --- |
| `docs/101` | Records the conversation proof capture reliability pass. | Planning proof is not implementation proof; CLI decides readiness, lock, and prompt compilation. | Yes | Yes | Helpful overlap with `docs/83`, `docs/97`, and `docs/99`; overlap is clarifying because it places proof capture beside public status boundaries. |
| `docs/102` | Audits locked-plan change-control UX before the stale-lock diagnostic landed. | Changed intent must stop stale handoff; `ni run` remains bounded prompt compilation only. | Yes | Yes | Historical audit overlaps with `docs/103` and `docs/104`; overlap is useful because it explains why later diagnostics and examples exist. |
| `docs/103` | Documents the implemented `LOCK-STALE` diagnostic. | `LOCK-STALE` is existing-lock staleness, not product or implementation failure. | Yes | Yes | Clear handoff to `docs/104`, `docs/106`, `docs/107`, and `docs/108`. |
| `docs/104` | Gives practical amend/relock examples after `LOCK-STALE`. | Recovery order stays `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000`. | Yes | Yes | Overlap with no-terminal examples is helpful because user workflows differ by evidence level. |
| `docs/105` | Verifies model workspace stale-lock wording. | Model workspace packs remain Experimental; skills can draft/explain but cannot validate, lock, relock, or update `.ni/plan.lock.json`. | Yes | Yes, added in this sweep | Navigation gap found: roadmap linked surrounding reliability docs but not this one. |
| `docs/106` | Shows no-terminal stale-lock examples. | No-terminal remains Experimental / assisted; exact trusted CLI output is required for readiness, lock freshness, relock, hash verification, and bounded prompt compilation claims. | Yes | Yes | Works with `docs/107`; overlap is not confusing because `docs/106` is scenario-based. |
| `docs/107` | Defines no-terminal transcript quality levels and copy-paste fields. | Transcript claims are scoped to model-only, pasted, trusted runner, fixture, or target-workspace evidence. | Yes | Yes | Strongly protects fixture-versus-target and validation-script transcript boundaries. |
| `docs/108` | Records broader changed-intent fixture coverage. | Fixture relock is separate from project-root relock; `.ni/session.json` is documented as mutable continuity state under current lock semantics. | Yes | Yes | Aligns tests and docs without changing lock-input semantics. |
| `docs/83` | Defines conversation proof capture and no-terminal proof limits. | Proof capture is an audit trail, not readiness authority. | Yes | Related roadmap surface | Useful upstream reference for `docs/101`. |
| `docs/97` | Defines benchmark claim boundaries. | Benchmark `READY` and `not_measured` evidence are planning-artifact scope only. | Yes | README and v0.5 roadmap | Required guard for benchmark examples and demo checks. |
| `docs/99` | Defines model workspace status. | Model workspace packs are Experimental until host-level install or discovery is verified. | Yes | README and this sweep's roadmap edit references `docs/105` for stale-lock wording | Canonical status doc for model workspace claims. |

## Recommended reading order

1. `docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md` - start with proof
   capture because it frames what planning evidence can and cannot prove.
2. `docs/102_CHANGE_CONTROL_UX_AUDIT.md` - read the historical UX risk that
   motivated stale-lock diagnostics.
3. `docs/103_STALE_LOCK_DIAGNOSTIC.md` - read the implemented `LOCK-STALE`
   behavior and recovery wording.
4. `docs/104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md` - move from diagnostic to
   concrete recovery examples.
5. `docs/108_CHANGED_INTENT_FIXTURE_COVERAGE.md` - check fixture coverage and
   project-root safety after the workflow examples.
6. `docs/106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md` - read no-terminal stale-lock
   scenarios once the normal CLI recovery path is clear.
7. `docs/107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md` - then read transcript
   quality rules for trusted runner and fixture evidence.
8. `docs/105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.md` - read the
   model workspace-specific stale-lock and skill-boundary audit.
9. `docs/99_MODEL_WORKSPACE_STATUS.md` - finish the model workspace status
   baseline and host-verification boundary.
10. `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md` - close with benchmark and
    `not_measured` boundaries that keep evidence claims scoped.

This order moves from general proof, to changed-intent mechanics, to fixture and
no-terminal evidence, then to model workspace and benchmark status surfaces.

## Claim boundary audit

| Claim area | Expected wording/status | Observed status | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Release binary | Release binary: Available | README, install docs, roadmap, and reliability docs keep release binaries Available only for verified v0.4.0 assets. | Yes | Guarded by `check-install-docs.py` and `release-check.sh`. |
| Curl installer | Curl installer: Available | README and installer verification docs keep curl installer Available for verified v0.4.0 assets. | Yes | No package-manager availability is implied. |
| Homebrew | Homebrew: Planned / v0.5 candidate | README and roadmap keep Homebrew Planned; roadmap says v0.5 is earliest scheduled implementation point. | Yes | No Homebrew Available claim found in audited status surfaces. |
| Model workspace packs | Model workspace packs: Experimental | README, package READMEs, `docs/99`, and `docs/105` preserve Experimental broad product status. | Yes | Host-level/global install and provider runtime behavior remain `not_verified`. |
| No-terminal method | No-terminal method: Experimental / assisted | No-terminal docs, `docs/106`, and `docs/107` preserve assisted-only wording. | Yes | Exact trusted CLI output is required for CLI-state claims. |
| CLI authority | CLI is authority. | `ni status`, `ni end`, and `ni run` remain the authoritative gates in docs and skills. | Yes | No model-readiness authority claim found. |
| Skills role | Skills are UX. | Skill packages and reliability docs state "Skills are UX; CLI is authority." | Yes | Skills may draft/explain; they do not lock, relock, or replace CLI validation. |
| READY | READY is planning artifact readiness only. | Benchmark and reliability docs scope `READY` to declared planning contract or artifact-readiness scope. | Yes | Product readiness remains out of scope. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | `docs/103`, `docs/104`, `docs/106`, `docs/107`, and `docs/108` use this meaning. | Yes | Also states stale can coexist with current planning `READY`. |
| ni run | Bounded handoff prompt compilation only; no downstream execution. | README, command docs, reliability docs, skills, and examples preserve non-execution. | Yes | `ni run` refuses stale handoff and does not relock. |
| fixture relock | Fixture relock is not project-root relock. | `docs/107` and `docs/108` state fixture evidence supports only fixture claims. | Yes | Final reports must still check project-root lockfile diffs. |
| trusted runner transcript | Claims require exact CLI output from the shown workspace/command/time. | `docs/106` and `docs/107` preserve model-only, pasted, trusted runner, fixture, and target-workspace levels. | Yes | No no-terminal deterministic validation claim found. |
| benchmark evidence | `not_measured` boundaries preserved. | `docs/97`, benchmark examples, and `demo-check.sh` keep product, research, adoption, cost, latency, and statistical claims out of scope. | Yes | Benchmark remains qualitative planning-artifact evidence. |
| kernel non-execution boundary | ni is not an execution harness. | README, roadmap, reliability docs, skill docs, and scripts preserve no task runner, adapter, queue, PR automation, release automation, or downstream execution layer. | Yes | No runtime feature was added. |

## Korean companion audit

| Doc pair | Pass? | Notes |
| --- | --- | --- |
| `docs/83` / `.ko.md` | Yes | Companion preserves proof-capture lifecycle, CLI authority, and no-terminal draft-only boundaries. |
| `docs/97` / `.ko.md` | Yes | Companion keeps `READY`, `not_measured`, dashboard, research approval, and fieldwork limitations scoped. |
| `docs/99` / `.ko.md` | Yes | Companion preserves Experimental model workspace status and `not_verified` host/provider claims. |
| `docs/100` / `.ko.md` | Yes | Companion preserves status vocabulary, GRILL closure, and next-direction boundaries. |
| `docs/101` / `.ko.md` | Yes | Companion does not overpromise proof capture or no-terminal validation. |
| `docs/102` / `.ko.md` | Yes | Companion preserves change-control and authority boundaries. |
| `docs/103` / `.ko.md` | Yes | Companion preserves `LOCK-STALE`, recovery flow, and stale-does-not-prove limits. |
| `docs/104` / `.ko.md` | Yes | Companion preserves recovery order and skill/no-terminal limits. |
| `docs/105` / `.ko.md` | Yes | Companion preserves Experimental, `not_verified`, skill drafting-only, and no-relock wording. |
| `docs/106` / `.ko.md` | Yes | Companion preserves Experimental / assisted and exact trusted CLI output requirements. |
| `docs/107` / `.ko.md` | Yes | Companion preserves model-only, pasted CLI output, trusted runner, fixture, and target-workspace transcript categories. |
| `docs/108` / `.ko.md` | Yes | Companion preserves fixture relock versus project-root relock and current `.ni/session.json` hash behavior. |
| `docs/109` / `.ko.md` | Yes | Added in this sweep; companion preserves commands, paths, status constants, diagnostic labels, and exact boundary phrases. |

## Navigation and cross-link audit

| Surface | Expected link or navigation role | Pass? | Notes |
| --- | --- | --- | --- |
| README | Public entry should link install, no-terminal, model workspace, benchmark, Homebrew, and command surfaces without crowding reliability docs. | Yes | No README link added; the roadmap is the better reliability-doc index. |
| Roadmap v0.5 section | Should index reliability docs 101-109 and status docs without excessive cross-links. | Yes | Added narrow links for `docs/105` and this `docs/109` sweep. |
| `docs/103` | Should point from diagnostic to examples, no-terminal, transcript, and fixture coverage. | Yes | Existing links are sufficient. |
| `docs/104` | Should point from amend/relock examples to no-terminal, transcript, and fixture coverage. | Yes | Existing links are sufficient. |
| `docs/106` | Should point no-terminal examples to transcript quality. | Yes | Existing link to `docs/107` is sufficient. |
| `docs/107` | Should point transcript quality to changed-intent fixture coverage. | Yes | Existing link to `docs/108` is sufficient. |
| `docs/108` | Should identify the next release-readiness sweep. | Yes | Existing recommended next task names this sweep; roadmap now links this doc. |

## Validation surface

| Script/check | Current protection | Gap | Change made |
| --- | --- | --- | --- |
| `check-install-docs.py` | Enforces install/distribution status rows, release/curl/Homebrew/model workspace markers, and forbidden availability claims. | It checks selected stable phrases, not every semantic overclaim. | None. Current surface is stable and low-noise. |
| `check-skill-packs.sh` | Verifies skill pack files, Experimental status, CLI authority, `LOCK-STALE`, no-relock, no-lockfile-update, and recovery-order wording. | It does not parse every skill sentence. | None. Current exact phrase checks are appropriate. |
| `demo-check.sh` | Verifies benchmark docs, no-terminal docs, transcript quality, fixture relock wording, ni-grill docs-only boundaries, and seed-only exports. | It is broad but still example-driven; future docs can drift outside checked markers. | None. No new risky phrase justified another check. |
| `quality.sh` | Runs JSON/schema/markdown/formatting/readme/install/skill/prompt-budget/core-boundary/asset checks, `gofmt -w .`, `go test ./...`, and `smoke.sh`. | Broad wrapper can mechanically format Go if touched. | None. No Go files touched. |
| `smoke.sh` | Builds `ni`, exercises public commands, readiness, lock, prompt compilation, exports, amendment, relock, and seed boundaries in temporary workspaces. | Temporary fixture relocks must not be reported as project-root relock. | None. The reporting boundary is documented in this sweep. |
| `install-check.sh` | Verifies source, build, local binary, and temporary local install paths. | Does not verify Homebrew or global model workspace install. | None. Homebrew remains Planned; host-level model workspace behavior remains unverified. |
| `release-check.sh` | Verifies release-readiness docs, release facts, release pipeline markers, benchmark boundaries, examples, Go tests, smoke, demo, install, and status proof. | It does not make a v0.5 release-candidate decision. | None. This sweep selects a release-candidate readiness audit as the next task. |

## Findings

| Finding | Severity | Surface | Change made | Blocks v0.5? |
| --- | --- | --- | --- | --- |
| No material claim-boundary contradiction found across audited v0.5 reliability docs. | Low | `docs/101` through `docs/108`, related docs, README, packages, examples, scripts | Added this sweep document to record the audit. | No. |
| Roadmap did not link `docs/105` even though it linked surrounding v0.5 reliability docs. | Low | `docs/51_POST_RELEASE_ROADMAP.md`, `.ko.md` | Added a narrow roadmap pointer for model workspace stale-lock wording verification. | No. |
| Roadmap needed a pointer to the new release-readiness sweep. | Low | `docs/51_POST_RELEASE_ROADMAP.md`, `.ko.md` | Added a narrow roadmap pointer for this `docs/109` sweep. | No. |

There are no material findings that block v0.5. The findings are navigation and
audit-record improvements only.

## Changes made

- `docs/109_V0_5_RELIABILITY_DOCS_RELEASE_READINESS_SWEEP.md`: added this
  English release-readiness sweep.
- `docs/109_V0_5_RELIABILITY_DOCS_RELEASE_READINESS_SWEEP.ko.md`: added Korean
  companion with the same claim boundaries.
- `docs/51_POST_RELEASE_ROADMAP.md`: added narrow links to `docs/105` and this
  sweep.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: added matching Korean roadmap links.

No Go files, skill docs, static checks, root lockfiles, runtime behavior, or
release actions were changed.

## What this sweep proves

- v0.5 reliability docs are internally aligned for audited surfaces.
- current public status boundaries are preserved.
- Korean companions do not overpromise for audited pairs.
- validation scripts cover selected stable boundary phrases.

## What this sweep does not prove

- implementation correctness
- downstream execution success
- product readiness
- benchmark effect size
- adoption/cost/latency improvement
- global model workspace behavior
- host-level model workspace verification
- Homebrew availability
- no-terminal deterministic validation
- exhaustive changed-intent coverage beyond existing fixtures

## Remaining risks

- docs may drift after future edits
- static checks cover selected phrases, not all semantic overclaims
- provider host behavior remains unverified
- Homebrew remains Planned until tested
- external user validation remains limited
- v0.5 release readiness still needs a final release-candidate check if a
  release is planned

## Recommended next task

Selected next task: E. v0.5 release candidate readiness audit.

Why: the reliability docs now have an explicit sweep record and only narrow
navigation findings. Before starting another adoption or implementation lane,
the project should decide whether v0.5 is release-candidate ready, what remains
blocking, and which status claims can be carried into an RC plan without
overclaiming Homebrew, model workspace host behavior, no-terminal validation,
or benchmark impact.

## Next task prompt

```text
Proceed with the v0.5 release candidate readiness audit in /Users/namba/Documents/project/ni.

This is an audit and documentation task only. Do not publish, tag, release, relock the project root, run ni end on the project root, execute generated prompts, add downstream execution behavior, add adapters, add queues, add PR automation, add release automation, mark Homebrew Available, mark model workspace packs Available, or claim no-terminal deterministic validation.

Goal:
Decide whether ni v0.5 is ready to become a release candidate, based on current docs, status claims, validation scripts, examples, and v0.5 reliability evidence.

Context:
ni is ni-kernel: a pre-runtime Project Intent Compiler for AI Agents.
Core protocol:
conversation -> project contract -> readiness gate -> lock hash -> downstream handoff.
Core flow:
ni init -> planning conversation -> docs/plan/** + .ni/contract.json + .ni/session.json -> ni status --proof --next-questions -> ni end -> .ni/plan.lock.json -> ni run --max-chars 4000 -> bounded downstream handoff prompt.

Current status boundaries:
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni run compiles a bounded handoff prompt and does not execute downstream work.
- READY is planning artifact readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- Fixture relock is not project-root relock.
- Benchmark evidence preserves not_measured boundaries.

Scope:
Audit README.md, README.ko.md, docs/51_POST_RELEASE_ROADMAP.md, docs/51_POST_RELEASE_ROADMAP.ko.md, docs/95 through docs/109 and Korean companions, docs/97, docs/99, docs/no-terminal*, examples/benchmark-report/, examples/no-terminal-assisted/, examples/ni-grill/, packages/claude-skills/, packages/codex-skills/, .agents/skills/, and validation scripts.

Required deliverable:
Add docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md and docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md if Korean companions are maintained.

The audit must include:
- current factual status table;
- release-candidate readiness criteria;
- evidence inventory for v0.5 reliability docs and examples;
- claim-boundary audit for release binary, curl installer, Homebrew, model workspace packs, no-terminal, CLI authority, skills role, LOCK-STALE, ni run, fixture relock, benchmark evidence, and the kernel non-execution boundary;
- validation command results;
- blockers, deferrals, warnings, and non-blocking risks;
- explicit decision: RC-ready, RC-ready with deferrals, or not RC-ready;
- changed files and why;
- selected next task after the audit, choosing exactly one.

Run:
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- bash scripts/quality.sh
- bash scripts/smoke.sh
- bash scripts/install-check.sh
- bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

If Go files are touched, also run:
- gofmt -w .
- GOCACHE=/private/tmp/ni-go-cache go test ./...

Final report must confirm:
- Go files touched or not;
- docs touched or not;
- skill docs touched or not;
- static checks touched or not;
- Homebrew remains Planned / v0.5 candidate;
- Model workspace packs remain Experimental;
- No-terminal method remains Experimental / assisted;
- Skills are UX; CLI is authority is preserved;
- ni run remains bounded prompt compilation only;
- .ni/contract.json, .ni/session.json, and .ni/plan.lock.json were not modified on the project root;
- ni end was not run on the project root;
- no relock was run on the project root;
- no fixture relock was claimed as project-root relock;
- no downstream execution behavior, shell adapter, Codex exec adapter, queue, PR automation, release action, Homebrew Available claim, model workspace Available claim, no-terminal deterministic claim, or benchmark overclaim was added.
```
