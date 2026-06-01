# ni project grill

이 example은 `ni-grill`을 ni project 자체에 dogfood한 기록이다. Repair task가
아니라 read-only planning-quality report다.

## Status

- CLI readiness: `READY`
- Command: `go run ./cmd/ni status --dir . --proof --next-questions`
- ni-grill role: advisory pressure only; CLI readiness가 authoritative하다.
- Non-execution boundary: `ni end`, `ni relock`, generated prompts, downstream
  agents, model APIs, release actions, shell adapters, queues, PR automation,
  runtime execution을 실행하지 않았다.

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
   Concern: Locked plan은 v0.5가 real benchmark data, broader product surfaces, conversation authoring reliability, lock/change-control UX, optional tested Homebrew, factual landing-page work, separate-package downstream integrations에 집중한다고 말한다. 하지만 `docs/51_POST_RELEASE_ROADMAP.md`는 아직 v0.5를 target seed quality and conformance 중심으로 설명하고, benchmark data를 v0.6으로 밀어둔다.
   Why it matters: downstream handoff가 locked v0.5 direction이 아니라 older public roadmap을 따를 수 있다.
   Question: `docs/51_POST_RELEASE_ROADMAP.md`를 CAP-019 / DEC-015와 맞게 고칠까, 아니면 locked plan이 기존 v0.5/v0.6 split을 유지한다고 명시할까?
   Answer shape: accepted roadmap source 하나와 exact phase wording, 또는 mismatch를 documented follow-up으로 defer한다는 decision
   Suggested action: clarify
   Blocks ni-end: maybe

## Full report

[`../../docs/93_NI_GRILL_DOGFOOD.ko.md`](../../docs/93_NI_GRILL_DOGFOOD.ko.md)를
본다.

## Recommended next task

Improve roadmap specificity
