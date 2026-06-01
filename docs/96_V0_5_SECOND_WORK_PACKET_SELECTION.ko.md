# v0.5 second work packet selection

## Current locked state

- Readiness: `go run ./cmd/ni status --dir . --proof --next-questions`는 `READY`를 보고하며 blockers, deferrals, warnings는 모두 없습니다.
- Lock state: `.ni/plan.lock.json`가 존재하고 `READY`를 기록합니다. Lock 시각은 `2026-06-01T06:09:41Z`입니다.
- Generated prompt path: `.ni/generated/v0.5-roadmap.prompt.txt`가 존재하고 `wc -c` 기준 정확히 4000 characters입니다.
- v0.5 direction: real benchmark evidence, conversation-authoring reliability, ni-grill quality, change-control UX, broader non-software product examples, model workspace pack verification, Homebrew as `Planned` / v0.5 candidate, separate-package-only downstream integrations, pre-runtime non-executing `ni-kernel`.
- First work packet completion: Task 175는 `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`와 `docs/95_V0_5_ACCEPTANCE_EVIDENCE.ko.md`를 추가했고, roadmap, benchmark/example docs, ni-grill dogfood, work-packet selection docs에서 link했으며, GRILL-003을 evidence-criteria clarification으로 addressed했다.
- Remaining GRILL notes: GRILL-004는 benchmark claim-boundary note로 남고, GRILL-005는 model workspace status preservation note로 남는다.

Scoring은 1-5이며 높을수록 좋습니다. `Cost`는 expected cost가 낮을수록 높은 점수입니다. `Boundary risk`는 ni-kernel boundary risk가 낮을수록 높은 점수입니다.

## Candidate comparison

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Benchmark claim-boundary pass | 5 | 5 | 5 | 4 | 5 | 5 | 29 | First |
| Model workspace status preservation pass | 4 | 4 | 4 | 4 | 3 | 5 | 24 | Defer; important but overclaim risk가 더 큼 |
| Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; benchmark boundary cleanup 이후 audit |
| Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Defer; benchmark boundary rules를 더 visible하게 만든 뒤 |
| Conversation proof capture reliability pass | 4 | 5 | 4 | 4 | 5 | 4 | 26 | Defer; benchmark claim boundaries 이후 strong follow-up |
| Model workspace host verification audit | 3 | 4 | 3 | 3 | 3 | 3 | 19 | Defer; global host availability overclaim 방지 |
| Homebrew implementation plan | 3 | 3 | 2 | 2 | 3 | 2 | 15 | Defer; tested evidence 전까지 `Planned` |
| Landing page decision / implementation | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; benchmark/reliability trust work보다 낮은 우선순위 |

## Selected second work packet

Benchmark claim-boundary pass.

## Why this second

첫 v0.5 packet은 어떤 evidence가 count되는지 정의했다. 다음으로 유용한 move는 그 evidence discipline을 가장 public하고 trust-sensitive한 v0.5 surface인 benchmark evidence에 적용하는 것이다. GRILL-004는 benchmark claims가 대체로 well-scoped되어 있지만 긴 page에서 `not_measured` boundaries가 status와 prompt excerpts 아래에 묻힐 수 있다고 말한다.

Third benchmark case보다 먼저인 이유는 새 case를 늘리기 전에 prominent claim-boundary pattern을 먼저 상속해야 하기 때문이다. Model workspace verification보다 먼저인 이유는 model-pack status는 이미 `Experimental` wording과 checks로 보호되고 있지만 benchmark evidence는 v0.5 credibility lane의 중심이기 때문이다. Change-control UX보다 먼저인 이유는 lock/relock semantics에 닿기 전에 low boundary-risk로 기존 GRILL note를 닫을 수 있기 때문이다.

## Work packet definition

Title: v0.5 benchmark claim-boundary pass.

Goal: 새 benchmark case나 더 강한 empirical claims를 추가하지 않고, benchmark docs, benchmark examples, example coverage, demo checks 전반에서 measured / `not_measured` benchmark boundaries를 더 prominent하고 consistent하게 만든다.

Scope:

- `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`의 real benchmark evidence lane을 사용한다.
- Benchmark summary, case README files, before/after evidence files, ni-grill benchmark examples에서 boundary visibility를 검토한다.
- Reader가 readiness result만 따로 인용할 수 있는 `READY` transition summaries 옆에 `not_measured`와 non-execution boundaries를 둔다.
- Checks를 추가하거나 조정할 경우 boundary wording과 required files만 확인하고, 새 empirical outcomes는 확인하지 않는다.
- 실제로 boundary를 더 visible하게 만들었을 때만 `docs/93_NI_GRILL_DOGFOOD.md`와 Korean companion에서 GRILL-004를 addressed로 표시한다.

Non-goals:

- Third benchmark case를 만들지 않는다.
- Generated prompts를 rerun하거나 execute하지 않는다.
- Model APIs, downstream agents, shell adapters, queues, product implementations를 실행하지 않는다.
- Benchmark data가 implementation quality, downstream-agent performance, adoption, cost, latency, rework reduction, statistical effect size를 증명한다고 claim하지 않는다.
- Lock/relock semantics, `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`를 바꾸지 않는다.
- Homebrew를 `Available`로 표시하거나, model workspace packs를 global availability로 upgrade하거나, no-terminal deterministic validation을 claim하지 않는다.

Expected changed files:

- `docs/77_BENCHMARK_CASE_STUDY.md`
- `docs/82_EXAMPLE_COVERAGE.md`
- `docs/93_NI_GRILL_DOGFOOD.md`
- `docs/93_NI_GRILL_DOGFOOD.ko.md`
- `examples/benchmark-report/README.md`
- `examples/benchmark-report/README.ko.md`
- `examples/benchmark-report/cases/internal-dashboard/README.md`
- `examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.md`
- `examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.ko.md`
- `examples/benchmark-report/cases/research-protocol/README.md`
- `examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md`
- `examples/benchmark-report/cases/research-protocol/15-before-after-evidence.ko.md`
- `examples/ni-grill/06-internal-dashboard-grill.md`
- `examples/ni-grill/07-research-protocol-grill.md`
- `scripts/demo-check.sh` only if a lightweight static boundary check is useful

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

- Benchmark summary docs는 각 benchmark `READY` 또는 before/after transition claim 가까이에 `not_measured` boundaries를 유지한다.
- Internal-dashboard와 research-protocol case docs는 main readiness summaries 가까이에 non-execution과 artifact-readiness-only wording을 보존한다.
- Example coverage는 future benchmark work를 v0.5 evidence matrix와 claim-boundary pattern으로 안내한다.
- GRILL-004는 이 pass로 addressed되거나, 더 좁은 named follow-up으로 남는다.
- Benchmark/adoption claim은 existing evidence보다 강해지지 않는다.
- `ni status`는 `READY`로 유지되고 validation commands가 통과한다.

Risks:

- Boundary wording이 반복적으로 보일 수 있다. Mitigation: main claim 가까이에 boundary를 두고, 모든 detail을 반복하기보다 detailed `not_measured` sections로 link한다.
- Check가 brittle해질 수 있다. Mitigation: exact prose block이 아니라 durable boundary phrases와 required files를 확인한다.
- Benchmark evidence가 실제보다 강하게 보일 수 있다. Mitigation: `not_measured`, manual qualitative, artifact-readiness-only, synthetic-fixture-only, no-downstream-execution wording을 보존한다.

Follow-up task: Benchmark claim-boundary patterns가 stable해진 뒤 third benchmark case를 선택한다.

## Tasks deferred

- Homebrew: tap, formula, checksums, audit, install, published tap install, `ni --help`, `ni version` evidence가 모두 통과할 때까지 `Planned`다.
- Landing page: README가 canonical quick entry이며 trust/evidence work가 더 높은 우선순위다.
- Additional benchmark cases: next case를 추가하기 전에 boundary wording을 더 visible하고 reusable하게 만든다.
- Model pack availability upgrade: broad availability에는 host-level install 또는 discovery evidence가 필요하다.
- Downstream integrations: separate packages, target exports, seed formats, downstream-owned notes로 남아야 한다.
- Change-control UX: 중요하지만 benchmark boundary pass보다 lock-semantics risk가 높으므로 defer한다.

## Next executable Codex prompt

```text
Goal:
Implement the v0.5 benchmark claim-boundary pass.

This is a documentation, examples, and lightweight-checks task. Do not add runtime behavior or new benchmark cases.

Context:
Task 175 added docs/95_V0_5_ACCEPTANCE_EVIDENCE.md and addressed GRILL-003. Task 176 selected "Benchmark claim-boundary pass" as the second v0.5 work packet. GRILL-004 remains: benchmark claims are generally scoped, but long benchmark pages can bury not_measured boundaries below status and prompt excerpts.

Read first:
- AGENTS.md
- docs/96_V0_5_SECOND_WORK_PACKET_SELECTION.md
- docs/95_V0_5_ACCEPTANCE_EVIDENCE.md
- docs/77_BENCHMARK_CASE_STUDY.md
- docs/82_EXAMPLE_COVERAGE.md
- docs/93_NI_GRILL_DOGFOOD.md
- docs/93_NI_GRILL_DOGFOOD.ko.md
- examples/benchmark-report/README.md
- examples/benchmark-report/README.ko.md
- examples/benchmark-report/cases/internal-dashboard/README.md
- examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.md
- examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.ko.md
- examples/benchmark-report/cases/research-protocol/README.md
- examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md
- examples/benchmark-report/cases/research-protocol/15-before-after-evidence.ko.md
- examples/ni-grill/06-internal-dashboard-grill.md
- examples/ni-grill/07-research-protocol-grill.md
- scripts/demo-check.sh

Run before editing:
- go run ./cmd/ni status --dir . --proof --next-questions

Make these changes:
- Make the benchmark report summary and case summaries keep not_measured boundaries close to each READY or before/after readiness transition claim.
- Keep internal-dashboard READY scoped to benchmark planning-meeting artifact readiness only.
- Keep research-protocol READY scoped to synthetic benchmark fixture readiness only.
- Keep dashboard product readiness, research approval, fieldwork authorization, implementation quality, downstream-agent performance, adoption, cost, latency, rework reduction, and statistical effect size explicitly not_measured where relevant.
- Update docs/82_EXAMPLE_COVERAGE.md only if it needs a short benchmark boundary rule or link.
- Update examples/ni-grill benchmark files only if they need to point to the clearer boundary rule.
- Update docs/93_NI_GRILL_DOGFOOD.md and .ko so GRILL-004 is marked addressed only if the boundary visibility was actually improved; otherwise keep it as a narrower follow-up.
- Add a lightweight demo-check assertion only if it protects durable boundary wording without making prose brittle.

Rules:
- Do not implement or select benchmark case 3.
- Do not execute generated prompts.
- Do not run Codex exec.
- Do not call model APIs.
- Do not add runtime execution.
- Do not add shell adapters, downstream agents, queues, PR automation, issue publishing, release automation, or evidence runners.
- Do not add user-facing contract add/list/set commands.
- Do not mark Homebrew Available.
- Do not claim model workspace packs are globally verified unless host-level install is actually verified.
- Do not claim no-terminal deterministic validation.
- Do not run ni end or ni relock.
- Do not edit .ni/plan.lock.json manually.
- Do not update .ni/contract.json or .ni/session.json.
- Do not fake benchmark/adoption claims.

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
- Summarize boundary visibility improvements.
- State whether GRILL-004 is addressed or remains a narrower follow-up.
- Include validation results.
- Confirm no implementation, runtime execution, release action, Homebrew availability claim, global model-pack claim, no-terminal deterministic claim, generated prompt execution, relock, or lockfile edit was added.
```
