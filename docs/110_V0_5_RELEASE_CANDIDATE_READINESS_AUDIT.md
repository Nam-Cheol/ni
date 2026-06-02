# v0.5 Release Candidate Readiness Audit

## Current status

- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Audit goal

This audit decides whether the current repository can be treated as a v0.5
release-candidate-ready state. It records evidence, claim boundaries,
deferrals, residual risks, and the next task. It does not publish, tag, release,
run `ni end` on the project root, relock the project root, or change Go
implementation, lock semantics, skill packages, static checks, or release
automation.

## Decision

Decision: `RC_READY_WITH_DEFERRALS`.

Justification: the current repository passes the required CLI, Go, install,
skill-pack, demo, smoke, quality, and release checks; `ni status` reports
`READY`; root lockfiles are unchanged; and no release, installer, model
workspace, no-terminal, benchmark, stale-lock, skill, or kernel boundary
overclaim was found. The remaining gaps are explicit deferrals rather than RC
blockers: Homebrew remains `Planned / v0.5 candidate`, model workspace
host-level behavior remains `Experimental`, no-terminal remains `Experimental /
assisted`, external user validation remains limited, and benchmark breadth is
bounded.

## RC readiness criteria

| Criterion | Result | Evidence | Notes |
| --- | --- | --- | --- |
| Positioning and README accuracy | Pass | README, README.ko.md, docs/40, docs/53, docs/109 | Public wording preserves ni as a pre-runtime Project Intent Compiler. |
| CLI validation | Pass | `ni status --proof --next-questions` | Project root reports `READY` with no blockers, deferrals, or warnings. |
| Go tests | Pass | `go test ./...` | No Go files were changed by this audit. |
| Release/install docs | Pass | README, install docs, `check-install-docs.py` | Release binary and curl installer are Available; Homebrew is not overclaimed. |
| Release binary status | Pass | README, docs/53, install docs | Release binary: Available. |
| Curl installer status | Pass | README, docs/install-curl*, installer docs | Curl installer: Available. |
| Homebrew status | Deferral | README, docs/54, docs/80, docs/109 | Homebrew: Planned / v0.5 candidate. |
| v0.5 reliability docs | Pass | docs/101 through docs/109 | Reliability docs are linked and boundary-audited. |
| Stale-lock/change-control reliability | Pass | docs/103, docs/104, docs/108, tests | `LOCK-STALE` and recovery order are documented and tested. |
| Benchmark claim boundaries | Pass | docs/97, examples/benchmark-report, `demo-check.sh` | Benchmark evidence remains bounded and not causal. |
| Model workspace status | Deferral | docs/99, docs/105, package READMEs | Model workspace packs: Experimental. |
| No-terminal status | Deferral | docs/no-terminal*, docs/106, docs/107 | No-terminal method: Experimental / assisted. |
| Skill-pack boundaries | Pass | packages/*-skills, .agents/skills, `check-skill-packs.sh` | Skills are UX; CLI is authority. |
| Korean companions | Pass | `.ko.md` companions and Korean roadmap | Exact status terms are preserved. |
| Project-root lockfile safety | Pass | protected-file diff check | `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` were not changed. |
| Runtime boundary | Pass | README, docs/53, docs/109, command docs | Kernel boundary excludes task runner, adapter, queue, and downstream execution behavior. |

## Evidence inventory

| Evidence | Result | Notes |
| --- | --- | --- |
| ni status | Pass | Project root prints `NI Intent Readiness: READY`; blockers, deferrals, and warnings are `None`. |
| go test ./... | Pass | All packages pass. |
| check-install-docs.py | Pass | Install/distribution claim markers pass. |
| check-skill-packs.sh | Pass | Skill-pack status, CLI authority, stale-lock, and no-relock boundaries pass. |
| demo-check.sh | Pass | Benchmark, no-terminal, transcript, ni-grill, and seed-only boundaries pass. |
| quality.sh | Pass | Broad quality wrapper passes. |
| smoke.sh | Pass | Fixture workflows, locks, stale refusal, prompt compilation, and seed boundaries pass. |
| install-check.sh | Pass | Source, build, local binary, and temporary local install paths pass. |
| release-check.sh | Pass | Release-readiness and claim-boundary checks pass. |
| git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json | Pass | No protected-file diff. |
| docs/101 through docs/109 | Pass | v0.5 reliability set is present and linked from roadmap. |
| examples/benchmark-report | Pass | `not_measured` and planning-artifact boundaries remain visible. |
| examples/ni-grill | Pass | Grill examples remain docs/review evidence, not readiness authority. |
| skill packages | Pass | Codex, Claude, and repo-local skills preserve UX-only boundaries. |

## Blockers

None.

## Deferrals

| Deferral | Status | RC impact | Required future proof |
| --- | --- | --- | --- |
| Homebrew implementation/availability | Planned / v0.5 candidate | Non-blocking because availability is not claimed. | Tap/formula, checksums, audit, local formula install, published tap install, `ni --help`, and `ni version`. |
| Model workspace host verification | Experimental | Non-blocking because broad status remains Experimental. | Host-specific install or discovery proof and provider behavior transcript. |
| External user validation | Limited | Non-blocking for RC documentation readiness; still an adoption risk. | User-run transcripts or maintained external validation notes. |
| Additional benchmark breadth | Bounded | Non-blocking because current claims remain qualitative and scoped. | A third benchmark case or broader benchmark report without causal overclaims. |

## Warnings

None.

## Risks

| Risk | Residual impact | Mitigation |
| --- | --- | --- |
| docs may drift | Future edits could overclaim status or authority. | Keep `check-install-docs.py`, `check-skill-packs.sh`, `demo-check.sh`, and release checks in the RC path. |
| static checks cover selected phrases not all semantic overclaims | Passing scripts is not exhaustive semantic proof. | Keep human claim-boundary audits before release notes or public launch copy. |
| provider host behavior unverified | Model workspace UX may behave differently in real hosts. | Keep Model workspace packs: Experimental until host proof exists. |
| Homebrew Planned until tested | Users may expect package-manager availability. | Keep Homebrew: Planned / v0.5 candidate until formula/tap validation passes. |
| external user validation limited | Adoption issues may appear after RC. | Treat external validation as a post-RC or release-polish task. |
| benchmark evidence bounded/not causal | Case studies can be misread as product or performance proof. | Preserve `not_measured` and planning-artifact boundaries. |
| no-terminal assisted/not deterministic | Users may confuse assisted drafting with CLI proof. | Keep exact trusted CLI output requirements in no-terminal docs. |

## Claim-boundary audit

| Claim area | Expected boundary | Result |
| --- | --- | --- |
| READY | Planning contract readiness only; not product readiness. | Pass |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | Pass |
| ni run | Compiles a bounded handoff prompt only and refuses stale handoff. | Pass |
| Homebrew | Planned / v0.5 candidate; not Available. | Pass |
| Model workspace packs | Experimental; host/global behavior is not verified. | Pass |
| No-terminal | Experimental / assisted; exact CLI output is required for CLI-state claims. | Pass |
| Skills | Skills are UX; CLI is authority. | Pass |
| Benchmark evidence | Planning-artifact evidence only; impact, adoption, cost, and latency are `not_measured`. | Pass |
| Fixture relock | Fixture relock is separate from project-root relock. | Pass |
| Trusted runner transcript | Supports only the shown workspace, command, output, and time. | Pass |
| Runtime execution | Not included in `ni-kernel`; downstream work remains outside core. | Pass |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | All Go packages passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; no blockers, deferrals, or warnings. |
| `python3 scripts/check-install-docs.py` | Pass | Install docs checks passed. |
| `bash scripts/check-skill-packs.sh` | Pass | Skill-pack checks passed. |
| `bash scripts/demo-check.sh` | Pass | Initial sandboxed run hit default Go build cache permissions; approved exact-command rerun passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Quality checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Smoke checks passed in temporary fixtures. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Install checks passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Release checks passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected-file diff. |

Validation scripts exercise fixture `ni end` flows in temporary workspaces.
Those fixture relocks are distinct from project-root relock. This audit did not
run `ni end` on the project root.

## Changed files

- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md`: added this RC readiness audit.
- `docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`: added Korean companion.
- `docs/51_POST_RELEASE_ROADMAP.md`: added a narrow pointer to this RC audit.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: added matching Korean pointer.

## What this audit proves

- The current repo can be treated as `RC_READY_WITH_DEFERRALS`.
- Required validation commands pass.
- Current public status claims remain bounded.
- Root `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` were not changed.
- Homebrew, model workspace packs, no-terminal, benchmark, skill, and kernel
  boundary claims are preserved.

## What this audit does not prove

- Homebrew availability.
- Provider host behavior for model workspace packs.
- Deterministic no-terminal validation without exact CLI output.
- External user adoption or satisfaction.
- Benchmark causal effect, statistical significance, adoption, cost, latency, or implementation quality.
- Downstream execution quality.

## Recommended next task

Selected next task: E. v0.5 RC polish / release notes draft.

Why: the RC decision is `RC_READY_WITH_DEFERRALS`, and the next useful step is
to prepare release-note wording and public RC polish without publishing,
tagging, releasing, or upgrading deferred status claims.

Follow-up draft: `docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md` records the
RC polish and release-note draft while preserving `RC_READY_WITH_DEFERRALS`,
known deferrals, draft-only wording, and release/non-release boundaries.

## Next task prompt

```text
Proceed with v0.5 RC polish and a release notes draft in /Users/namba/Documents/project/ni.

This is a documentation-only polish task. Do not publish, tag, release, run ni end on the project root, relock the project root, edit .ni/contract.json, edit .ni/session.json, edit .ni/plan.lock.json, change Go implementation, change skill packages, add release automation, add Homebrew availability claims, mark model workspace packs Available, claim deterministic no-terminal validation, or claim benchmark causal impact.

Use docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md as the readiness source. Preserve the decision RC_READY_WITH_DEFERRALS.

Create or update the smallest appropriate release-note draft document for v0.5 RC polish. The draft must preserve:
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- READY is planning contract readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- ni run compiles a bounded handoff prompt only.
- Fixture relock is separate from project-root relock.
- Benchmark evidence keeps not_measured boundaries.

Include:
- concise RC summary;
- what changed in v0.5;
- what remains deferred;
- user-facing install/status caveats;
- validation evidence summary;
- explicit non-goals and claim boundaries;
- changed files;
- recommended next task, choosing exactly one.

Run:
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final report must confirm no project-root relock, no protected .ni file changes, no release/tag/publish action, no Homebrew Available claim, no model workspace Available claim, no deterministic no-terminal validation claim, and no benchmark overclaim.
```
