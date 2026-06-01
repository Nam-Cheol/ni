# ni-grill dogfood on ni

## Status

- CLI readiness: `READY` from `go run ./cmd/ni status --dir . --proof --next-questions`; blockers, deferrals, and warnings were all `None`.
- ni-grill role: advisory planning-quality pressure only. It did not change CLI readiness, approve lock, or replace `ni status`.
- Scope: read `AGENTS.md`, README surfaces, `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`, `.ni/generated/v0.4.0-post-release.prompt.txt`, roadmap, benchmark, Homebrew, conversation-proof, release, model-pack, and ni-grill docs plus the checked-in benchmark and skill-pack examples.
- Non-execution boundary: this dogfood report did not run `ni end`, `ni relock`, generated prompts, downstream agents, model APIs, release actions, shell adapters, queues, PR automation, or runtime execution.

## Findings summary

| ID | Severity | Category | Affected | Suggested action | Blocks ni-end |
| --- | --- | --- | --- | --- | --- |
| GRILL-001 | High | roadmap specificity | `docs/51_POST_RELEASE_ROADMAP.md`; CAP-019 / DEC-015 | clarify | maybe |
| GRILL-002 | High | distribution status | `docs/80_HOMEBREW_DECISION.md`; CAP-020 / DEC-016 | clarify | maybe |
| GRILL-003 | Medium | acceptance evidence | CAP-019 / EVAL-023; `docs/51_POST_RELEASE_ROADMAP.md` | clarify | no |
| GRILL-004 | Low | benchmark claim boundary | `docs/77_BENCHMARK_CASE_STUDY.md`; `examples/benchmark-report/**` | keep as note | no |
| GRILL-005 | Note | model workspace status | `docs/55_MODEL_WORKSPACE_PACKS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/*-skills/**` | keep as note | no |

## Full findings

Grill findings:
1. GRILL-001 — High — roadmap specificity
   Affected: `docs/51_POST_RELEASE_ROADMAP.md`; CAP-019 / DEC-015
   Concern: The locked plan says v0.5 is focused on real benchmark data, broader product surfaces, conversation authoring reliability, lock/change-control UX, optional tested Homebrew, factual landing-page work, and separate-package downstream integrations. `docs/51_POST_RELEASE_ROADMAP.md` still frames v0.5 primarily as target seed quality and conformance, with benchmark data pushed to v0.6.
   Why it matters: downstream handoff could follow the older public roadmap instead of the locked v0.5 direction, especially around when benchmark evidence and conversation reliability should be improved.
   Question: Should `docs/51_POST_RELEASE_ROADMAP.md` be updated so v0.5 matches CAP-019 / DEC-015, or should the locked plan explicitly keep the older v0.5/v0.6 split?
   Answer shape: one accepted roadmap source plus exact phase wording, or a decision to defer the mismatch as a documented follow-up
   Suggested action: clarify
   Blocks ni-end: maybe

2. GRILL-002 — High — distribution status
   Affected: `docs/80_HOMEBREW_DECISION.md`; CAP-020 / DEC-016
   Concern: The Homebrew decision correctly keeps `Homebrew: Planned`, but its current distribution state still names verified v0.3.0 release binaries and curl assets while the locked post-release state and install docs name verified v0.4.0 assets as current.
   Why it matters: public trust can suffer when one distribution decision page preserves the correct Homebrew status but points readers at stale release/curl evidence.
   Question: Should the Homebrew decision page be refreshed to say release binary and curl installer are Available for verified v0.4.0 assets while Homebrew remains Planned?
   Answer shape: yes/no plus exact replacement status bullets, or rationale for leaving the dated v0.3.0 decision as historical context
   Suggested action: clarify
   Blocks ni-end: maybe

3. GRILL-003 — Medium — acceptance evidence
   Affected: CAP-019 / EVAL-023; `docs/51_POST_RELEASE_ROADMAP.md`
   Concern: v0.5 acceptance evidence is broad and spread across plan records. EVAL-023 names scripts and themes, but the public roadmap does not yet define minimum evidence packages for real benchmark data, conversation reliability, lock/change-control UX, and Homebrew readiness.
   Why it matters: downstream work may produce documents or examples that look complete without proving the specific evidence needed for each v0.5 lane.
   Question: Should v0.5 have a short acceptance-evidence checklist that says what artifact, CLI proof, review, or verification completes each roadmap lane?
   Answer shape: checklist with one evidence shape per lane, or explicit decision that EVAL-023 is sufficient for now
   Suggested action: clarify
   Blocks ni-end: no

4. GRILL-004 — Low — benchmark claim boundary
   Affected: `docs/77_BENCHMARK_CASE_STUDY.md`; `examples/benchmark-report/**`
   Concern: Benchmark claims are generally well-scoped: the report calls the work manual and qualitative, labels `not_measured` boundaries, limits dashboard `READY` to artifact readiness, and limits research `READY` to synthetic fixture readiness. The remaining pressure is visibility: the long case-study page can bury those boundaries below detailed status and prompt excerpts.
   Why it matters: a reader may quote a `READY` transition without the nearby `not_measured` limits, even though those limits are present elsewhere.
   Question: Should future benchmark summaries keep the `not_measured` boundary next to every `READY` table or transition row?
   Answer shape: yes/no plus a compact summary rule for future benchmark pages
   Suggested action: keep as note
   Blocks ni-end: no

5. GRILL-005 — Note — model workspace status
   Affected: `docs/55_MODEL_WORKSPACE_PACKS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/claude-skills/**`; `packages/codex-skills/**`
   Concern: Model workspace pack statuses are currently factual: source and zip paths are available where verified, global host discovery remains unverified, no-terminal is assisted, and the skills repeatedly say "Skills are UX; CLI is authority."
   Why it matters: preserving this wording prevents model packs from being mistaken for CLI authority or execution adapters.
   Question: Should future model-pack edits preserve the current status split: Available source/zip paths, Experimental broad product path, Unverified global host discovery, and CLI authority?
   Answer shape: yes/no, or a narrower status vocabulary if product language changes later
   Suggested action: keep as note
   Blocks ni-end: no

## Lower-priority findings not shown

Minor editorial parity checks were not promoted because the top pressure is already covered by the roadmap and distribution-status findings.

## Follow-up

GRILL-001 and GRILL-002 are addressed by the Task 172 roadmap and distribution
documentation alignment. GRILL-003 remains a medium acceptance-evidence note,
GRILL-004 remains a low benchmark claim-boundary note, and GRILL-005 remains a
model workspace status preservation note.

## What ni-grill did not do

- did not change readiness
- did not edit lockfile
- did not execute generated prompt
- did not publish release
- did not call downstream agents
- did not make model judgment authoritative

## Recommended next task

Improve roadmap specificity
