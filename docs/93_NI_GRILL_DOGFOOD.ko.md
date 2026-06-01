# ni-grill dogfood on ni

## Status

- CLI readiness: `go run ./cmd/ni status --dir . --proof --next-questions` 결과는 `READY`이며 blockers, deferrals, warnings는 모두 `None`이다.
- ni-grill role: advisory planning-quality pressure일 뿐이다. CLI readiness를 바꾸거나, lock을 approve하거나, `ni status`를 대체하지 않았다.
- Scope: `AGENTS.md`, README surfaces, `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`, `.ni/generated/v0.4.0-post-release.prompt.txt`, roadmap, benchmark, Homebrew, conversation-proof, release, model-pack, ni-grill docs와 checked-in benchmark / skill-pack examples를 읽었다.
- Non-execution boundary: 이 dogfood report는 `ni end`, `ni relock`, generated prompts, downstream agents, model APIs, release actions, shell adapters, queues, PR automation, runtime execution을 실행하지 않았다.

## Findings summary

| ID | Severity | Category | Affected | Suggested action | Blocks ni-end |
| --- | --- | --- | --- | --- | --- |
| GRILL-001 | High | roadmap specificity | `docs/51_POST_RELEASE_ROADMAP.md`; CAP-019 / DEC-015 | clarify | maybe |
| GRILL-002 | High | distribution status | `docs/80_HOMEBREW_DECISION.md`; CAP-020 / DEC-016 | clarify | maybe |
| GRILL-003 | Medium | acceptance evidence | CAP-019 / EVAL-023; `docs/51_POST_RELEASE_ROADMAP.md` | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`로 addressed | no |
| GRILL-004 | Low | benchmark claim boundary | `docs/77_BENCHMARK_CASE_STUDY.md`; `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`; `examples/benchmark-report/**` | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`로 addressed | no |
| GRILL-005 | Note | model workspace status | `docs/55_MODEL_WORKSPACE_PACKS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/*-skills/**` | keep as note | no |

## Full findings

Grill findings:
1. GRILL-001 — High — roadmap specificity
   Affected: `docs/51_POST_RELEASE_ROADMAP.md`; CAP-019 / DEC-015
   Concern: Locked plan은 v0.5가 real benchmark data, broader product surfaces, conversation authoring reliability, lock/change-control UX, optional tested Homebrew, factual landing-page work, separate-package downstream integrations에 집중한다고 말한다. 하지만 `docs/51_POST_RELEASE_ROADMAP.md`는 아직 v0.5를 target seed quality and conformance 중심으로 설명하고, benchmark data를 v0.6으로 밀어둔다.
   Why it matters: downstream handoff가 특히 benchmark evidence와 conversation reliability의 시점에서 locked v0.5 direction이 아니라 오래된 public roadmap을 따를 수 있다.
   Question: `docs/51_POST_RELEASE_ROADMAP.md`를 CAP-019 / DEC-015와 맞게 고칠까, 아니면 locked plan이 기존 v0.5/v0.6 split을 유지한다고 명시할까?
   Answer shape: accepted roadmap source 하나와 exact phase wording, 또는 mismatch를 documented follow-up으로 defer한다는 decision
   Suggested action: clarify
   Blocks ni-end: maybe

2. GRILL-002 — High — distribution status
   Affected: `docs/80_HOMEBREW_DECISION.md`; CAP-020 / DEC-016
   Concern: Homebrew decision은 `Homebrew: Planned`를 올바르게 유지한다. 다만 current distribution state가 아직 verified v0.3.0 release binaries와 curl assets를 말한다. Locked post-release state와 install docs는 verified v0.4.0 assets를 current로 말한다.
   Why it matters: Distribution decision page 하나가 Homebrew status는 맞게 유지하면서도 stale release/curl evidence를 가리키면 public trust가 약해질 수 있다.
   Question: Homebrew decision page를 release binary와 curl installer는 verified v0.4.0 assets 기준 Available이고 Homebrew는 Planned라고 refresh할까?
   Answer shape: yes/no plus exact replacement status bullets, 또는 dated v0.3.0 decision을 historical context로 남기는 rationale
   Suggested action: clarify
   Blocks ni-end: maybe

3. GRILL-003 — Medium — acceptance evidence
   Affected: CAP-019 / EVAL-023; `docs/51_POST_RELEASE_ROADMAP.md`
   Concern: v0.5 acceptance evidence는 broad하고 plan records에 흩어져 있다. EVAL-023은 scripts와 themes를 말하지만 public roadmap은 real benchmark data, conversation reliability, lock/change-control UX, Homebrew readiness별 minimum evidence package를 아직 정의하지 않는다.
   Why it matters: downstream work가 각 v0.5 lane에 필요한 evidence를 증명하지 않고도 complete처럼 보이는 docs나 examples를 만들 수 있다.
   Question: v0.5 roadmap lane마다 어떤 artifact, CLI proof, review, verification이 completion evidence인지 짧은 checklist가 필요할까?
   Answer shape: lane별 evidence shape 하나씩 있는 checklist, 또는 EVAL-023으로 지금은 충분하다는 explicit decision
   Suggested action: clarify
   Blocks ni-end: no

   Resolution note: `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`로 addressed되었다. 이
   문서는 v0.5 work를 위한 lane별 completion evidence, status vocabulary,
   verification references, `not_measured` boundaries를 정의한다. 이것은 어떤
   v0.5 lane이 complete라는 뜻이 아니다.

4. GRILL-004 — Low — benchmark claim boundary
   Affected: `docs/77_BENCHMARK_CASE_STUDY.md`; `examples/benchmark-report/**`
   Concern: Benchmark claims는 대체로 잘 scoped되어 있다. Report는 manual qualitative work라고 말하고, `not_measured` boundaries를 label하며, dashboard `READY`는 artifact readiness로, research `READY`는 synthetic fixture readiness로 제한한다. 남은 pressure는 visibility다. 긴 case-study page에서는 detailed status와 prompt excerpt 아래로 boundary가 묻힐 수 있다.
   Why it matters: boundary가 다른 곳에 있어도 reader가 `READY` transition만 인용할 수 있다.
   Question: Future benchmark summaries에서는 모든 `READY` table 또는 transition row 옆에 `not_measured` boundary를 유지해야 할까?
   Answer shape: yes/no plus future benchmark pages를 위한 compact summary rule
   Suggested action: keep as note
   Blocks ni-end: no

   Resolution note: `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`로 addressed되었다.
   이 문서는 required claim-boundary label, status vocabulary, dashboard/research
   case-specific limit, review checklist를 정의한다.
   `docs/77_BENCHMARK_CASE_STUDY.ko.md`, `docs/82_EXAMPLE_COVERAGE.ko.md`,
   benchmark report examples도 이 boundary를 가리키고 `READY` transition 옆에
   `not_measured` marker를 유지한다.

5. GRILL-005 — Note — model workspace status
   Affected: `docs/55_MODEL_WORKSPACE_PACKS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/claude-skills/**`; `packages/codex-skills/**`
   Concern: Model workspace pack statuses는 현재 factual하다. Source와 zip paths는 verified된 범위에서 Available이고, global host discovery는 unverified이며, no-terminal은 assisted이고, skills는 반복해서 "Skills are UX; CLI is authority."라고 말한다.
   Why it matters: 이 wording은 model packs가 CLI authority나 execution adapters로 오해되는 일을 막는다.
   Question: Future model-pack edits에서도 현재 status split, 즉 Available source/zip paths, Experimental broad product path, Unverified global host discovery, CLI authority를 유지할까?
   Answer shape: yes/no, 또는 product language가 나중에 바뀔 경우 더 좁은 status vocabulary
   Suggested action: keep as note
   Blocks ni-end: no

## Lower-priority findings not shown

Minor editorial parity checks는 roadmap과 distribution-status findings가 이미 top pressure를 덮기 때문에 promoted하지 않았다.

## Follow-up

GRILL-001과 GRILL-002는 Task 172의 roadmap 및 distribution documentation
alignment로 addressed되었다. GRILL-003은 evidence-criteria clarification으로
`docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`에서 addressed되었다. GRILL-004는
`docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`로 addressed되었고, GRILL-005는 model
workspace status preservation note로 남는다.

## What ni-grill did not do

- did not change readiness
- did not edit lockfile
- did not execute generated prompt
- did not publish release
- did not call downstream agents
- did not make model judgment authoritative

## Recommended next task

Improve roadmap specificity
