# v0.5 work packet selection

## Current locked state

- Readiness: `go run ./cmd/ni status --dir . --proof --next-questions`는 `READY`를 보고하며 blockers, deferrals, warnings는 모두 없습니다.
- Lock state: `.ni/plan.lock.json`가 존재하고 `READY`를 기록합니다. Lock 시각은 `2026-06-01T06:09:41Z`입니다.
- Generated prompt path: `.ni/generated/v0.5-roadmap.prompt.txt`가 존재하고 `wc -c` 기준 정확히 4000 characters입니다.
- v0.5 direction: CAP-019, REQ-019, EVAL-023, RISK-018, DEC-017, roadmap docs는 v0.5를 evidence quality, conversation-authoring reliability, ni-grill quality, change-control UX, broader non-software product examples, factual adoption hardening, separate-package-only downstream integrations에 맞춥니다. Kernel은 계속 pre-runtime이고 non-executing입니다.

Scoring은 1-5이며 높을수록 좋습니다. `Cost`는 expected cost가 낮을수록 높은 점수입니다. `Boundary risk`는 ni-kernel boundary risk가 낮을수록 높은 점수입니다.

## Candidate comparison

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Acceptance evidence clarity pass | 5 | 5 | 5 | 4 | 5 | 5 | 29 | First |
| Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Evidence criteria를 먼저 명확히 한 뒤 defer |
| Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; semantics 변경 전에 audit 우선 |
| ni-grill dogfood repair pass | 4 | 4 | 4 | 4 | 5 | 5 | 26 | Defer; selected packet이 일부 포함 |
| Model workspace host verification audit | 3 | 4 | 3 | 3 | 3 | 3 | 19 | Defer; host availability overclaim 방지 |
| Homebrew implementation plan | 3 | 3 | 2 | 2 | 3 | 2 | 15 | Defer; tested evidence 전까지 Planned |
| Landing page decision / implementation | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; trust evidence보다 낮은 우선순위 |

## Selected first work packet

Acceptance evidence clarity pass.

Implementation note: 이 packet은
[`95_V0_5_ACCEPTANCE_EVIDENCE.md`](95_V0_5_ACCEPTANCE_EVIDENCE.md)에서
implemented되었다.

## Why this first

이 packet은 ni-grill dogfood report에서 남은 GRILL-003 medium acceptance-evidence note를 직접 다룹니다. Benchmark cases, conversation reliability, change-control UX, model-pack claims, Homebrew status, downstream seed work는 모두 later task가 complete라고 말하기 전에 visible evidence criteria가 필요합니다.

Third benchmark case보다 먼저인 이유는 새 case가 자체 completion bar를 만들지 않고 clear evidence checklist를 재사용해야 하기 때문입니다. Change-control UX보다 먼저인 이유는 evidence expectation이 명확해지기 전 lock semantics로 drift할 수 있기 때문입니다. Homebrew, landing page, model-pack availability보다 먼저인 이유는 그 작업들이 public-claim risk가 더 크고 immediate evidence value가 낮기 때문입니다.

## Work packet definition

Title: v0.5 acceptance evidence clarity pass.

Goal: v0.5 lane별 minimum evidence package를 정의합니다: benchmark evidence, conversation-authoring reliability, ni-grill quality, change-control UX, Homebrew readiness, model workspace pack availability, no-terminal proof, product-surface examples, downstream seed formats, landing-page adoption claims.

Scope:

- Concise v0.5 acceptance-evidence checklist를 추가합니다.
- Checklist를 CAP-019, REQ-019, EVAL-023, RISK-016, RISK-018, DEC-015, DEC-017에 연결합니다.
- Roadmap과 dogfood reference를 업데이트해 GRILL-003을 addressed로 만들거나 더 좁은 follow-up으로 남깁니다.
- Benchmark와 adoption claims의 `not_measured` boundaries를 보존합니다.
- Verified host-level evidence가 없으면 Homebrew는 `Planned`, model workspace packs는 `Experimental`로 유지합니다.
- Downstream integration work는 `ni-kernel` 밖의 separate packages, seed formats, notes로 유지합니다.

Non-goals:

- Benchmark case 3를 구현하지 않습니다.
- Lock/relock/amendment semantics를 바꾸지 않습니다.
- Homebrew, package-manager distribution, landing page publishing, model host installation, downstream integrations, runtime behavior를 구현하지 않습니다.
- `.ni/generated/v0.5-roadmap.prompt.txt`를 실행하지 않습니다.
- Codex exec, adapters, queues, PR automation, issue publishing, release automation, shell execution, user-facing contract `add` / `list` / `set` commands를 추가하지 않습니다.
- `.ni/plan.lock.json`를 수동 편집하거나 relock하지 않습니다.

Expected changed files:

- `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`
- `docs/95_V0_5_ACCEPTANCE_EVIDENCE.ko.md` if Korean companion docs are maintained
- `docs/51_POST_RELEASE_ROADMAP.md`
- `docs/51_POST_RELEASE_ROADMAP.ko.md` if the roadmap companion is maintained
- `docs/93_NI_GRILL_DOGFOOD.md`
- `docs/93_NI_GRILL_DOGFOOD.ko.md` if the dogfood companion is maintained
- Possibly `docs/77_BENCHMARK_CASE_STUDY.md` and `docs/82_EXAMPLE_COVERAGE.md` only if cross-links are needed

Validation commands:

```bash
go run ./cmd/ni status --dir . --proof --next-questions
bash scripts/quality.sh
bash scripts/demo-check.sh
bash scripts/smoke.sh
bash scripts/check-skill-packs.sh
go test ./... # only if Go files are touched or quality.sh does not already cover the relevant Go change
```

Completion criteria:

- v0.5 acceptance-evidence checklist가 존재하고 roadmap lane별 evidence shape를 하나씩 명명합니다.
- 각 checklist lane은 무엇을 measured로 볼지, 무엇을 `not_measured`로 남길지, 어떤 file 또는 command가 completion을 증명할지 말합니다.
- GRILL-003은 open medium note에서 addressed가 되거나, 더 작은 named follow-up으로 남습니다.
- Roadmap은 future v0.5 tasks가 implementation 전에 checklist를 보도록 안내합니다.
- Public claim은 existing evidence보다 강해지지 않습니다.
- `ni status`는 `READY`로 유지되고 validation commands가 통과합니다.

Risks:

- Evidence checklist가 너무 넓어져 쓰기 어려워질 수 있습니다. Mitigation: lane마다 one minimum proof shape와 one boundary note로 제한합니다.
- Benchmark 또는 adoption claims가 overreach할 수 있습니다. Mitigation: relevant evidence row 옆에 `not_measured`, Homebrew `Planned`, model-pack `Experimental` wording을 유지합니다.
- Docs가 locked source-of-truth와 drift할 수 있습니다. Mitigation: later amendment가 root planning state를 명시적으로 바꾸기 전에는 locked `docs/plan/**` 또는 `.ni/contract.json`를 편집하지 않습니다.

Follow-up task: Accepted v0.5 evidence checklist를 사용해 third benchmark case를 선택하거나 작성합니다.

## Tasks deferred

- Homebrew: tap, formula, checksums, audit, install, published tap install, `ni --help`, `ni version` evidence가 모두 통과할 때까지 `Planned`입니다.
- Landing page: README가 canonical quick entry이고 adoption page work는 trust checklist보다 evidence value가 낮습니다.
- Additional benchmark cases: future benchmark가 무엇을 prove하고 무엇을 `not_measured`로 둘지 checklist가 먼저 필요합니다.
- Model pack availability upgrade: broad availability claim을 위한 host-level install 또는 discovery가 아직 unverified입니다.
- Downstream integrations: `ni-kernel` behavior가 아니라 separate packages, target exports, seed formats, downstream-owned notes로 남아야 합니다.

## Next executable Codex prompt

```text
Goal:
Implement the v0.5 acceptance evidence clarity pass.

This is a documentation and evidence-boundary task. Do not implement runtime behavior.

Context:
The first v0.5 work packet selected from the locked roadmap is "Acceptance evidence clarity pass". It addresses GRILL-003 from docs/93_NI_GRILL_DOGFOOD.md: v0.5 acceptance evidence is broad and spread across plan records, while future work needs a short checklist saying what artifact, CLI proof, review, or verification completes each roadmap lane.

Read first:
- AGENTS.md
- docs/94_V0_5_WORK_PACKET_SELECTION.md
- .ni/plan.lock.json
- .ni/contract.json
- .ni/session.json
- docs/plan/**
- docs/51_POST_RELEASE_ROADMAP.md
- docs/51_POST_RELEASE_ROADMAP.ko.md
- docs/77_BENCHMARK_CASE_STUDY.md
- docs/82_EXAMPLE_COVERAGE.md
- docs/83_CONVERSATION_PROOF_CAPTURE.md
- docs/90_ENGINEERING_SKILLS_APPLICABILITY.md
- docs/91_NI_GRILL.md
- docs/92_NI_GRILL_OUTPUT_BUDGET.md
- docs/93_NI_GRILL_DOGFOOD.md
- docs/93_NI_GRILL_DOGFOOD.ko.md
- README.md
- README.ko.md
- packages/claude-skills/
- packages/codex-skills/
- examples/benchmark-report/cases/internal-dashboard/
- examples/benchmark-report/cases/research-protocol/
- examples/ni-grill/

Run before editing:
- go run ./cmd/ni status --dir . --proof --next-questions

Make these changes:
- Add docs/95_V0_5_ACCEPTANCE_EVIDENCE.md.
- Add docs/95_V0_5_ACCEPTANCE_EVIDENCE.ko.md if Korean companion docs are maintained.
- Define one concise evidence row per v0.5 lane:
  benchmark evidence, conversation-authoring reliability, ni-grill quality, change-control UX, Homebrew readiness, model workspace pack availability, no-terminal proof, product-surface examples, downstream seed formats, and landing-page adoption claims.
- For each lane, state:
  what counts as completion evidence;
  what remains not_measured or unclaimed;
  which existing or expected file/command proves it;
  which boundary must not be crossed.
- Update docs/51_POST_RELEASE_ROADMAP.md and its Korean companion, if maintained, to point v0.5 tasks at the evidence checklist.
- Update docs/93_NI_GRILL_DOGFOOD.md and its Korean companion, if maintained, so GRILL-003 is marked addressed by the new checklist or retained only as a narrower follow-up.
- Add links from docs/77_BENCHMARK_CASE_STUDY.md or docs/82_EXAMPLE_COVERAGE.md only if needed for discoverability.

Rules:
- Do not implement benchmark case 3.
- Do not execute .ni/generated/v0.5-roadmap.prompt.txt.
- Do not run Codex exec.
- Do not call model APIs.
- Do not add runtime execution.
- Do not add shell adapters, Codex adapters, queues, PR automation, issue publishing, release automation, evidence runners, or downstream agents.
- Do not add user-facing contract add/list/set commands.
- Do not mark Homebrew Experimental or Available.
- Do not upgrade model workspace packs from Experimental unless host-level install or discovery is verified in this task.
- Do not edit .ni/plan.lock.json manually.
- Do not run ni end, ni relock, publish, tag, or release.
- Do not weaken accepted risks, mitigations, requirements, evaluations, non-goals, or benchmark boundaries.

Validation:
- go run ./cmd/ni status --dir . --proof --next-questions
- bash scripts/quality.sh
- bash scripts/demo-check.sh
- bash scripts/smoke.sh
- bash scripts/check-skill-packs.sh
- go test ./... if Go files are touched

Final response:
- List changed files.
- State readiness result.
- Summarize the evidence checklist.
- State whether GRILL-003 is addressed or remains a narrower follow-up.
- Include validation results.
- Confirm no implementation, runtime execution, release action, Homebrew availability claim, model-pack availability upgrade, relock, or lockfile edit was added.
```
