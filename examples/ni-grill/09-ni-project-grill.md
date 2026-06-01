# ni project grill

This example dogfoods `ni-grill` on the ni project itself. It is a read-only
planning-quality report, not a repair task.

## Status

- CLI readiness: `READY`
- Command: `go run ./cmd/ni status --dir . --proof --next-questions`
- ni-grill role: advisory pressure only; CLI readiness remains authoritative.
- Non-execution boundary: did not run `ni end`, `ni relock`, generated prompts,
  downstream agents, model APIs, release actions, shell adapters, queues, PR
  automation, or runtime execution.

## Findings summary

| ID | Severity | Category | Affected | Suggested action | Blocks ni-end |
| --- | --- | --- | --- | --- | --- |
| GRILL-001 | High | roadmap specificity | `docs/51_POST_RELEASE_ROADMAP.md`; CAP-019 / DEC-015 | clarify | maybe |
| GRILL-002 | High | distribution status | `docs/80_HOMEBREW_DECISION.md`; CAP-020 / DEC-016 | clarify | maybe |
| GRILL-003 | Medium | acceptance evidence | CAP-019 / EVAL-023; `docs/51_POST_RELEASE_ROADMAP.md` | clarify | no |
| GRILL-004 | Low | benchmark claim boundary | `docs/77_BENCHMARK_CASE_STUDY.md`; `examples/benchmark-report/**` | keep as note | no |
| GRILL-005 | Note | model workspace status | `docs/55_MODEL_WORKSPACE_PACKS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/*-skills/**` | keep as note | no |

## One full finding

Grill findings:
1. GRILL-001 — High — roadmap specificity
   Affected: `docs/51_POST_RELEASE_ROADMAP.md`; CAP-019 / DEC-015
   Concern: The locked plan says v0.5 is focused on real benchmark data, broader product surfaces, conversation authoring reliability, lock/change-control UX, optional tested Homebrew, factual landing-page work, and separate-package downstream integrations. `docs/51_POST_RELEASE_ROADMAP.md` still frames v0.5 primarily as target seed quality and conformance, with benchmark data pushed to v0.6.
   Why it matters: downstream handoff could follow the older public roadmap instead of the locked v0.5 direction.
   Question: Should `docs/51_POST_RELEASE_ROADMAP.md` be updated so v0.5 matches CAP-019 / DEC-015, or should the locked plan explicitly keep the older v0.5/v0.6 split?
   Answer shape: one accepted roadmap source plus exact phase wording, or a decision to defer the mismatch as a documented follow-up
   Suggested action: clarify
   Blocks ni-end: maybe

## Full report

See [`../../docs/93_NI_GRILL_DOGFOOD.md`](../../docs/93_NI_GRILL_DOGFOOD.md).

## Recommended next task

Improve roadmap specificity
