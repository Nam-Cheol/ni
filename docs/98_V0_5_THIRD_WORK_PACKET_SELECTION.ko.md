# v0.5 third work packet selection

## Current locked state

- Readiness: `go run ./cmd/ni status --dir . --proof --next-questions`는
  `READY`를 보고하며 blockers, deferrals, warnings는 모두 없습니다.
- Lock state: `.ni/plan.lock.json`가 존재하고 `READY`를 기록합니다. Lock
  시각은 `2026-06-01T06:09:41Z`입니다.
- Generated prompt path: `.ni/generated/v0.5-roadmap.prompt.txt`가 존재하고
  `wc -c` 기준 정확히 4000 characters입니다.
- v0.5 direction: evidence quality, conversation-authoring reliability,
  ni-grill challenge quality, change-control UX, broader product-surface
  examples, factual adoption hardening, tested evidence 이후에만 optional
  Homebrew, host-level install이 verified되기 전까지 model workspace packs는
  `Experimental`, no-terminal은 assisted, downstream integrations는 separate
  packages, target exports, seed formats, downstream-owned notes로만 유지합니다.
  Kernel은 pre-runtime, non-executing 상태를 유지합니다.
- Completed first work packet: Task 175는 v0.5 acceptance evidence matrix를
  추가하고 GRILL-003을 evidence-criteria clarification으로 addressed했습니다.
- Completed second work packet: Task 177은 benchmark claim-boundary
  clarification을 추가하고 `not_measured`, artifact-readiness, synthetic-fixture,
  non-execution boundaries를 visible하게 유지해 GRILL-004를 addressed했습니다.
- Remaining GRILL notes: GRILL-005만 model workspace status preservation note로
  남아 있습니다.

Scoring은 1-5이며 높을수록 좋습니다. `Cost`는 expected cost가 낮을수록 높은
점수입니다. `Boundary risk`는 ni-kernel boundary risk가 낮을수록 높은 점수입니다.

## Candidate comparison

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Model workspace status preservation pass | 4 | 5 | 5 | 4 | 4 | 5 | 27 | Third |
| Conversation proof capture reliability pass | 4 | 5 | 4 | 4 | 5 | 4 | 26 | GRILL notes를 닫은 뒤 strong follow-up |
| Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Defer; benchmark work는 방금 boundary pass를 받음 |
| Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; lock-semantics risk가 더 큼 |
| Model workspace host verification audit | 3 | 4 | 3 | 2 | 2 | 2 | 16 | Defer; honest host verification 준비도가 낮음 |
| Homebrew implementation plan | 3 | 3 | 2 | 2 | 3 | 2 | 15 | Defer; public availability risk가 아직 큼 |
| Landing page decision / implementation | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; status correctness보다 낮은 우선순위 |

## Selected third work packet

Model workspace status preservation pass.

## Why this third

첫 두 v0.5 packet은 evidence expectation을 명시하고 그 discipline을 benchmark
claim boundaries에 적용했다. 남은 GRILL pressure는 GRILL-005다. Codex와
Claude skill packs가 CLI authority, global host availability, runtime adapters,
deterministic no-terminal validation으로 오해되지 않게 model workspace wording을
보존해야 한다.

Host verification audit보다 먼저인 이유는 repository가 truthful status를
보존할 evidence는 이미 갖고 있지만, broad model workspace availability를 upgrade할
host-level evidence는 아직 충분하지 않기 때문이다. Third benchmark case보다 먼저인
이유는 benchmark boundary가 방금 clarified되었고, 남은 status note를 먼저 닫을 수
있기 때문이다. Change-control UX보다 먼저인 이유는 factual workspace status 보존이
amendment, relock, stale-lock, changed-intent behavior보다 semantic blast radius가
작기 때문이다.

## Work packet definition

Title: v0.5 model workspace status preservation pass.

Goal: README surfaces, model workspace docs, Codex and Claude skill pack READMEs,
no-terminal guidance, validation checks가 다음 status를 일관되게 보존하게 만든다:
model workspace packs는 host-level install 또는 discovery가 verified되기 전까지
broad product path로는 `Experimental`이다. Repo-local/source/zip/manual-copy
paths는 실제로 verified된 범위 안에서만 설명한다. Skills are UX and the CLI
remains authority.

Scope:

- README, install/adoption docs, model workspace docs, no-terminal docs, skill
  pack READMEs, checked-in `.agents/skills/**`, package docs의 status wording을
  검토한다.
- `Available`, `Experimental`, `Planned`, `Unverified` vocabulary를 existing
  evidence와 맞춘다.
- Durable status와 authority wording을 위한 lightweight static checks만
  추가하거나 조정한다.
- 실제로 required status split을 보존하게 되었을 때만
  `docs/93_NI_GRILL_DOGFOOD.md`와 Korean companion에서 GRILL-005를 addressed로
  표시한다.
- 변경된 English docs에 maintained Korean companion이 있으면 같이 보존한다.

Non-goals:

- 새 global Codex, Claude, generic model host install path를 verify하거나 claim하지
  않는다.
- Model workspace packs를 broad product path로 `Experimental`에서 `Available`로
  upgrade하지 않는다.
- Runtime execution, model API calls, shell adapters, Codex adapters, downstream
  agents, queues, PR automation, release automation, evidence runners를 추가하지
  않는다.
- Skill packs를 CLI authority 또는 deterministic readiness gates로 만들지 않는다.
- No-terminal mode가 deterministic이라고 claim하지 않는다.
- Homebrew를 `Available`로 표시하지 않는다.
- Generated prompts 실행, `codex exec`, `ni end`, relock, publish, tag, release를
  하지 않는다.
- `.ni/plan.lock.json`, `.ni/contract.json`, `.ni/session.json`를 편집하지 않는다.

Expected changed files:

- `README.md`
- `README.ko.md`
- `docs/55_MODEL_WORKSPACE_PACKS.md`
- `docs/55_MODEL_WORKSPACE_PACKS.ko.md` if maintained
- `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`
- `docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md`
- `docs/no-terminal.md`
- `docs/no-terminal.ko.md`
- `docs/93_NI_GRILL_DOGFOOD.md`
- `docs/93_NI_GRILL_DOGFOOD.ko.md`
- `packages/claude-skills/README.md`
- `packages/claude-skills/README.ko.md`
- `packages/codex-skills/README.md`
- `packages/codex-skills/README.ko.md`
- `.agents/skills/**/SKILL.md` only if authority/status wording is inconsistent
- `scripts/check-skill-packs.sh` only if a durable static check is useful

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

- README와 model workspace docs가 status split을 보존한다: source/repo-local 및
  zip/manual-copy paths는 verified된 범위 안에서만 설명하고, broad model workspace
  product availability는 `Experimental`이며, global host discovery는 evidence가
  추가되기 전까지 unverified로 남는다.
- Codex와 Claude skill pack READMEs는 skills are UX and the CLI is authority라고
  말한다.
- No-terminal docs는 assisted drafting을 deterministic CLI readiness, lock, hash,
  prompt claims와 분리한다.
- Static checks는 brittle prose blocks에 의존하지 않고 durable authority/status
  wording을 보호한다.
- GRILL-005는 이 pass로 addressed되거나 더 좁은 named follow-up으로 남는다.
- Public claim은 existing evidence보다 강해지지 않는다.
- `ni status`는 `READY`로 유지되고 validation commands가 통과한다.

Risks:

- Status wording이 너무 조심스러워 verified source 또는 zip paths를 숨길 수 있다.
  Mitigation: narrow verified paths는 보존하되 broad host availability는
  `Experimental` 또는 `Unverified`로 유지한다.
- Model pack phrase가 CLI authority를 암시할 수 있다. Mitigation: user-facing pack
  guidance와 skill README files 가까이에 "Skills are UX; CLI is authority"를
  유지한다.
- Static checks가 brittle해질 수 있다. Mitigation: exact paragraph가 아니라
  `Experimental`, host-level/global host verification, no-terminal assisted
  wording, CLI authority 같은 durable concepts를 확인한다.
- Pass 도중 host verification을 하고 싶어질 수 있다. Mitigation: 이 task가
  reproducible host-level evidence를 명시적으로 수집하지 않는 한 host verification은
  deferred audit로 남긴다.

Follow-up task: GRILL-005를 닫은 뒤 conversation proof capture reliability,
change-control UX audit, third benchmark case 중 다음 작업을 선택한다.

## Tasks deferred

- Homebrew: tap, formula, checksums, audit, install, published tap install,
  `ni --help`, `ni version` evidence가 모두 통과할 때까지 `Planned`다.
- Landing page: status correctness와 claim discipline이 marketing surface
  expansion보다 우선순위가 높다.
- Additional benchmark cases: benchmark claim boundaries가 방금 clarified되었고,
  다음 case는 그 pattern을 재사용해야 하므로 defer한다.
- Change-control UX: 중요하지만 model workspace status wording 보존보다
  lock-semantics risk가 높으므로 defer한다.
- Model pack availability upgrade: broad availability에는 host-level install 또는
  discovery evidence가 필요하다.
- Downstream integrations: separate packages, target exports, seed formats,
  downstream-owned notes로 남아야 하며 `ni-kernel` behavior가 되면 안 된다.

## Next executable Codex prompt

```text
Goal:
Implement the v0.5 model workspace status preservation pass.

This is a documentation and lightweight-checks task. Do not add runtime behavior or upgrade model workspace availability claims.

Context:
Task 175 completed the v0.5 acceptance evidence criteria. Task 177 completed the benchmark claim-boundary clarification and addressed GRILL-004. GRILL-005 remains as a model workspace status preservation note: model workspace pack docs and skill pack READMEs must keep saying that model workspace packs are Experimental as a broad product path unless host-level install or discovery is verified, and that skills are UX while the CLI remains authority.

Read first:
- AGENTS.md
- README.md
- README.ko.md
- .ni/plan.lock.json
- .ni/contract.json
- .ni/session.json
- docs/93_NI_GRILL_DOGFOOD.md
- docs/93_NI_GRILL_DOGFOOD.ko.md
- docs/95_V0_5_ACCEPTANCE_EVIDENCE.md
- docs/97_BENCHMARK_CLAIM_BOUNDARIES.md
- docs/55_MODEL_WORKSPACE_PACKS.md
- docs/75_MODEL_PACK_INSTALL_VERIFICATION.md
- docs/no-terminal.md
- docs/no-terminal.ko.md
- packages/claude-skills/README.md
- packages/claude-skills/README.ko.md
- packages/codex-skills/README.md
- packages/codex-skills/README.ko.md
- packages/claude-skills/**
- packages/codex-skills/**
- .agents/skills/**
- scripts/check-skill-packs.sh

Run before editing:
- go run ./cmd/ni status --dir . --proof --next-questions

Make these changes:
- Review README, README.ko, model workspace docs, no-terminal docs, skill pack READMEs, and checked-in skill files for model workspace status wording.
- Preserve this status split everywhere it matters:
  source/repo-local skill files may be available where they exist;
  zip/manual-copy/package scripts may be available only where checked by scripts;
  broad model workspace product availability remains Experimental;
  global host install or discovery remains unverified unless this task produces reproducible host-level proof.
- Keep "Skills are UX; CLI is authority" or equivalent wording visible in skill pack README and skill guidance.
- Keep no-terminal mode described as assisted drafting only until trusted CLI proof exists.
- Add or adjust static checks in scripts/check-skill-packs.sh only if they protect durable status/authority wording without making prose brittle.
- Update docs/93_NI_GRILL_DOGFOOD.md and .ko so GRILL-005 is marked addressed only if the status split is now consistently preserved; otherwise retain it as a narrower follow-up.
- Maintain Korean companion docs for any changed English docs that already have a maintained Korean companion.

Rules:
- Do not verify or claim a new global Codex, Claude, or generic model host install path unless you have reproducible host-level evidence in this task.
- Do not upgrade model workspace packs from Experimental to Available as a broad product path.
- Do not claim no-terminal deterministic validation.
- Do not make skills CLI authority.
- Do not add runtime execution.
- Do not add shell adapters, Codex adapters, downstream agents, queues, PR automation, issue publishing, release automation, evidence runners, model API calls, or task-runner behavior.
- Do not mark Homebrew Available.
- Do not execute generated prompts.
- Do not run Codex exec.
- Do not run ni end or relock.
- Do not publish, tag, or release.
- Do not edit .ni/plan.lock.json manually.
- Do not update .ni/contract.json or .ni/session.json.
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
- Summarize status-preservation changes.
- State whether GRILL-005 is addressed or remains a narrower follow-up.
- Include validation results.
- Confirm no implementation, runtime execution, generated prompt execution, Codex exec, release action, Homebrew availability claim, broad model-pack availability upgrade, no-terminal deterministic claim, relock, or lockfile edit was added.
```
