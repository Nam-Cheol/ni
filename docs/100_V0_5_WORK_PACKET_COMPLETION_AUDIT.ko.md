# v0.5 Work Packet Completion Audit

## Current state

첫 세 개 v0.5 work packet은 documentation, example, static-check alignment
작업으로 complete이다:

- release binary: verified v0.4.0 assets 기준 Available;
- curl installer: verified v0.4.0 assets 기준 Available;
- Homebrew: Planned / v0.5 candidate only;
- model workspace packs: broad product path로 Experimental;
- no-terminal method: Experimental / assisted;
- runtime execution, shell adapters, Codex exec, queues, downstream agents, PR
  automation: `ni-kernel`에 포함되지 않음.

현재 authoritative CLI proof는 다음과 같다:

```text
NI Intent Readiness: READY
Blockers: None.
Deferrals: None.
Warnings: None.
```

이 audit은 `ni end`, relock, `.ni/plan.lock.json` edit, `.ni/contract.json`
edit, `.ni/session.json` edit, generated prompt execution, runtime behavior
추가를 하지 않았다.

## GRILL closure status

| Finding | Status | Evidence | Notes |
| --- | --- | --- | --- |
| GRILL-001 | addressed | `docs/51_POST_RELEASE_ROADMAP.md`; `docs/53_DISTRIBUTION_STRATEGY.md`; `docs/80_HOMEBREW_DECISION.md` | Roadmap과 distribution docs가 v0.5 evidence/adoption direction과 맞고 kernel boundary를 보존한다. |
| GRILL-002 | addressed | `docs/53_DISTRIBUTION_STRATEGY.md`; `docs/80_HOMEBREW_DECISION.md`; `README.md` | Release binary와 curl installer는 verified v0.4.0 assets 기준 Available이고 Homebrew는 Planned로 남는다. |
| GRILL-003 | addressed | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md` | Acceptance matrix가 v0.5 lane별 completion evidence와 `not_measured` boundary를 정의한다. |
| GRILL-004 | addressed | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `docs/77_BENCHMARK_CASE_STUDY.md`; `examples/benchmark-report/**`; `scripts/demo-check.sh` | Benchmark summaries가 `READY`, artifact-readiness, synthetic-fixture, non-execution, `not_measured` boundaries를 visible하게 유지한다. |
| GRILL-005 | addressed | `docs/99_MODEL_WORKSPACE_STATUS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/*-skills/README.md`; `scripts/check-skill-packs.sh` | Model workspace packs는 broad path로 Experimental이고 host/global install 및 provider behavior는 not verified로 남는다. Skills are UX; CLI is authority. |

GRILL-003부터 GRILL-005까지 open finding은 없다. Host-level model workspace
verification은 closure gap이 아니라 future work이다.

## Work packet completion

| Work packet | Evidence | GRILL finding addressed | Complete? | Notes |
| --- | --- | --- | --- | --- |
| Acceptance evidence clarity | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`; `docs/95_V0_5_ACCEPTANCE_EVIDENCE.ko.md` | GRILL-003 | yes | v0.5 lane별 evidence shape, status vocabulary, `not_measured` rules를 정의한다. |
| Benchmark claim-boundary pass | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`; benchmark examples; `scripts/demo-check.sh` | GRILL-004 | yes | Benchmark `READY` claims를 product/research outcome이 아니라 artifact readiness와 synthetic fixture로 제한한다. |
| Model workspace status preservation | `docs/99_MODEL_WORKSPACE_STATUS.md`; `docs/99_MODEL_WORKSPACE_STATUS.ko.md`; skill pack READMEs; `scripts/check-skill-packs.sh` | GRILL-005 | yes | broad model workspace status는 Experimental, host/provider claims는 not_verified, no-terminal은 assisted, CLI authority wording은 유지된다. |

## Claim/status audit

| Claim area | Expected status | Actual status | Pass? |
| --- | --- | --- | --- |
| Release binary | verified v0.4.0 release assets 기준 Available | README, distribution, Homebrew decision docs가 verified v0.4.0 assets 기준 Available이라고 말한다. | yes |
| Curl installer | verified v0.4.0 release assets 기준 Available | README, distribution, install docs, Homebrew decision docs가 verified v0.4.0 assets 기준 Available이라고 말한다. | yes |
| Homebrew | Planned; published/tested tap 또는 formula 없음 | README와 Homebrew docs가 Planned를 유지하고 tap/formula/install proof 전 Available wording을 금지한다. | yes |
| Model workspace packs | broad product path로 Experimental | README, distribution docs, `docs/99`, pack READMEs가 Experimental을 유지하고 not_verified host/provider claims를 보존한다. | yes |
| No-terminal method | Experimental / assisted; deterministic validation에는 CLI proof 필요 | README, distribution docs, no-terminal/example docs가 assisted-only wording을 보존한다. | yes |
| Internal-dashboard benchmark | benchmark planning-meeting artifact readiness에 대해서만 `READY` | Benchmark docs와 examples가 dashboard product readiness, implementation quality, downstream-agent performance를 `not_measured`로 보존한다. | yes |
| Research-protocol benchmark | synthetic benchmark protocol planning artifact readiness에 대해서만 `READY` | Benchmark docs와 examples가 real research approval, fieldwork authorization, research quality, intervention effectiveness를 `not_measured`로 보존한다. | yes |
| Runtime execution boundary | `ni-kernel`에 포함되지 않음 | README, roadmap, example coverage, grill docs, skills가 runtime execution, shell adapters, Codex exec, queues, downstream agents, PR automation을 제외한다. | yes |

## Validation surface

| Check area | Evidence | Status |
| --- | --- | --- |
| Benchmark claim boundaries | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `scripts/demo-check.sh` | present |
| Model workspace overclaim prevention | `docs/99_MODEL_WORKSPACE_STATUS.md`; `scripts/check-skill-packs.sh`; `scripts/check-install-docs.py` | present |
| Skill pack metadata and authority wording | `scripts/check-skill-packs.sh` | present |
| Install docs consistency | `scripts/check-install-docs.py`; `bash scripts/quality.sh` | present |
| Demo coverage | `scripts/demo-check.sh`; `docs/82_EXAMPLE_COVERAGE.md` | present |

## Next direction candidates

Scoring은 1-5이며 높을수록 좋다. `Cost`는 expected cost가 낮을수록 높은
점수다. `Boundary risk`는 `ni-kernel` boundary risk가 낮을수록 높은 점수다.

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| A. Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; lock/relock semantics에 더 가깝다. |
| B. Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Proof-capture consistency 이후의 strong follow-up. |
| C. Conversation proof capture reliability pass | 5 | 5 | 4 | 4 | 5 | 5 | 28 | Selected. |
| D. Model workspace host verification audit | 3 | 4 | 3 | 2 | 2 | 2 | 16 | Defer; host claims는 과장되기 쉽다. |
| E. Homebrew implementation audit | 3 | 4 | 3 | 2 | 2 | 3 | 17 | Defer; public availability-claim risk가 높다. |
| F. Landing page decision | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; product reliability보다 낮은 priority. |

## Selected next direction

Conversation proof capture reliability pass.

## Why this next

첫 세 개 v0.5 packet은 GRILL evidence와 claim-boundary notes를 닫았다. 다음으로
가치가 큰 작업은 sustained authoring loop를 더 일관되게 만드는 것이다: model이
user answer를 어떻게 기록하고, changed files와 IDs를 어떻게 이름 붙이고,
before/after `ni status` proof를 어떻게 보여주며, planning proof와 execution
evidence를 어떻게 분리하는지다.

이 방향은 third benchmark case보다 먼저 할 가치가 있다. 새 benchmark는 더 깨끗한
proof-capture pattern을 재사용해야 하기 때문이다. Change-control UX보다 boundary
risk가 낮고, 이후 amend/relock/stale-lock work를 설명하는 기반도 된다. Homebrew,
model host verification, landing page보다 public-claim 또는 adoption-message risk가
낮다.

## Next task prompt

```text
Goal:
Implement the v0.5 conversation proof capture reliability pass.

This is a documentation, examples, and lightweight-checks task. Do not add runtime behavior or implement any downstream work.

Context:
The first three v0.5 work packets are complete:
1. docs/95_V0_5_ACCEPTANCE_EVIDENCE.md addressed GRILL-003.
2. docs/97_BENCHMARK_CLAIM_BOUNDARIES.md addressed GRILL-004.
3. docs/99_MODEL_WORKSPACE_STATUS.md addressed GRILL-005.

docs/100_V0_5_WORK_PACKET_COMPLETION_AUDIT.md selected "Conversation proof capture reliability pass" as the next v0.5 direction.

Read first:
- AGENTS.md
- README.md
- README.ko.md
- docs/83_CONVERSATION_PROOF_CAPTURE.md
- docs/95_V0_5_ACCEPTANCE_EVIDENCE.md
- docs/97_BENCHMARK_CLAIM_BOUNDARIES.md
- docs/99_MODEL_WORKSPACE_STATUS.md
- docs/93_NI_GRILL_DOGFOOD.md
- docs/100_V0_5_WORK_PACKET_COMPLETION_AUDIT.md
- docs/31_NI_START_BEHAVIOR.md
- docs/91_NI_GRILL.md
- docs/92_NI_GRILL_OUTPUT_BUDGET.md
- packages/claude-skills/README.md
- packages/claude-skills/README.ko.md
- packages/codex-skills/README.md
- packages/codex-skills/README.ko.md
- packages/claude-skills/**
- packages/codex-skills/**
- examples/ni-start-dogfood/
- examples/conversation-authoring/
- examples/no-terminal-assisted/
- examples/ni-grill/
- scripts/check-skill-packs.sh
- scripts/demo-check.sh

Run before editing:
- go run ./cmd/ni status --dir . --proof --next-questions

Make these changes:
- Review conversation proof capture wording in docs, skill pack READMEs, checked-in skill files, and examples.
- Ensure proof capture consistently distinguishes:
  planning proof from execution evidence;
  model-authored summaries from CLI authority;
  before/after status from model judgment;
  changed docs/contract/session artifacts from unchanged artifacts;
  open blockers from advisory GRILL findings.
- Keep "Skills are UX; CLI is authority" visible where proof capture could be mistaken for readiness authority.
- Ensure no-terminal proof remains assisted/draft-only until trusted CLI proof exists.
- Update docs/83_CONVERSATION_PROOF_CAPTURE.md and Korean companion if maintained.
- Update package or repo-local skill wording only where it makes proof capture more consistent.
- Update examples only if they need clearer proof-capture references or boundary wording.
- Add lightweight static checks only if they protect durable proof-capture or authority wording without making prose brittle.

Rules:
- Do not run ni end or relock.
- Do not edit .ni/plan.lock.json manually.
- Do not update .ni/contract.json or .ni/session.json.
- Do not execute generated prompts.
- Do not run Codex exec.
- Do not call model APIs.
- Do not add runtime execution.
- Do not add shell adapters, Codex adapters, downstream agents, queues, PR automation, issue publishing, release automation, evidence runners, or task-runner behavior.
- Do not add user-facing contract add/list/set commands.
- Do not mark Homebrew Available.
- Do not upgrade model workspace packs from Experimental to Available as a broad product path.
- Do not claim no-terminal deterministic validation.
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
- Summarize proof-capture consistency changes.
- Include validation results.
- Confirm no implementation, runtime execution, generated prompt execution, Codex exec, release action, Homebrew availability claim, broad model-pack availability upgrade, no-terminal deterministic claim, relock, or lockfile edit was added.
```
