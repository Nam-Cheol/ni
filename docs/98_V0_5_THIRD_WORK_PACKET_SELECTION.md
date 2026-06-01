# v0.5 third work packet selection

## Current locked state

- Readiness: `go run ./cmd/ni status --dir . --proof --next-questions`
  reports `READY`, with no blockers, deferrals, or warnings.
- Lock state: `.ni/plan.lock.json` exists and records `READY`; it was locked at
  `2026-06-01T06:09:41Z`.
- Generated prompt path: `.ni/generated/v0.5-roadmap.prompt.txt` exists and is
  exactly 4000 characters by `wc -c`.
- v0.5 direction: evidence quality, conversation-authoring reliability,
  ni-grill challenge quality, change-control UX, broader product-surface
  examples, factual adoption hardening, optional Homebrew only after tested
  evidence, model workspace packs as `Experimental` unless host-level install
  is verified, no-terminal as assisted, and downstream integrations only as
  separate packages, target exports, seed formats, or downstream-owned notes.
  The kernel remains pre-runtime and non-executing.
- Completed first work packet: Task 175 added the v0.5 acceptance evidence
  matrix and addressed GRILL-003 as an evidence-criteria clarification.
- Completed second work packet: Task 177 added benchmark claim-boundary
  clarification and addressed GRILL-004 by keeping `not_measured`,
  artifact-readiness, synthetic-fixture, and non-execution boundaries visible.
- Remaining GRILL notes: GRILL-005 remains as a model workspace status
  preservation note.

Scoring uses 1-5 where higher is better. `Cost` means lower expected cost
scores higher. `Boundary risk` means lower risk to the ni-kernel boundary
scores higher.

## Candidate comparison

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Model workspace status preservation pass | 4 | 5 | 5 | 4 | 4 | 5 | 27 | Third |
| Conversation proof capture reliability pass | 4 | 5 | 4 | 4 | 5 | 4 | 26 | Strong follow-up after GRILL notes are closed |
| Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Defer; benchmark work just received a boundary pass |
| Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; higher lock-semantics risk |
| Model workspace host verification audit | 3 | 4 | 3 | 2 | 2 | 2 | 16 | Defer; honest host verification is not ready enough |
| Homebrew implementation plan | 3 | 3 | 2 | 2 | 3 | 2 | 15 | Defer; public availability risk remains high |
| Landing page decision / implementation | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; lower priority than status correctness |

## Selected third work packet

Model workspace status preservation pass.

## Why this third

The first two v0.5 packets made evidence expectations explicit and then applied
that discipline to benchmark claim boundaries. The remaining GRILL pressure is
GRILL-005: preserve model workspace wording so Codex and Claude skill packs are
not mistaken for CLI authority, global host availability, runtime adapters, or
deterministic no-terminal validation.

This beats the host verification audit because the repository already has
enough evidence to preserve truthful status, but not enough host-level evidence
to upgrade broad model workspace availability. It beats the third benchmark
case because benchmark boundaries were just clarified and another benchmark can
wait until the remaining status note is closed. It beats change-control UX
because preserving factual workspace status has a smaller semantic blast radius
than touching amendment, relock, stale-lock, or changed-intent behavior.

## Work packet definition

Title: v0.5 model workspace status preservation pass.

Goal: Make README surfaces, model workspace docs, Codex and Claude skill pack
READMEs, no-terminal guidance, and validation checks consistently preserve this
status: model workspace packs are `Experimental` as a broad product path unless
host-level install or discovery is verified; repo-local/source/zip/manual-copy
paths may be described only to the extent they are actually verified; skills are
UX and the CLI remains authority.

Scope:

- Review status wording in README, install/adoption docs, model workspace docs,
  no-terminal docs, skill pack READMEs, checked-in `.agents/skills/**`, and
  package docs.
- Keep `Available`, `Experimental`, `Planned`, and `Unverified` vocabulary
  consistent with existing evidence.
- Add or adjust lightweight static checks only for durable status and authority
  wording.
- Update `docs/93_NI_GRILL_DOGFOOD.md` and Korean companion so GRILL-005 is
  addressed only if the pass actually preserves the required status split.
- Preserve Korean companion docs where matching English docs are changed.

Non-goals:

- Do not verify or claim a new global Codex, Claude, or generic model host
  install path.
- Do not upgrade model workspace packs from `Experimental` to `Available` as a
  broad product path.
- Do not add runtime execution, model API calls, shell adapters, Codex adapters,
  downstream agents, queues, PR automation, release automation, or evidence
  runners.
- Do not make skill packs CLI authority or deterministic readiness gates.
- Do not claim no-terminal mode is deterministic.
- Do not mark Homebrew `Available`.
- Do not execute generated prompts, call `codex exec`, run `ni end`, relock,
  publish, tag, or release.
- Do not edit `.ni/plan.lock.json`, `.ni/contract.json`, or `.ni/session.json`.

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

- README and model workspace docs preserve the status split: source/repo-local
  and zip/manual-copy paths are described only where verified; broad model
  workspace product availability remains `Experimental`; global host discovery
  remains unverified unless evidence is added.
- Codex and Claude skill pack READMEs say skills are UX and the CLI is
  authority.
- No-terminal docs keep assisted drafting separate from deterministic CLI
  readiness, lock, hash, and prompt claims.
- Static checks cover durable authority/status wording without depending on
  brittle prose blocks.
- GRILL-005 is either marked addressed by this pass or retained as a narrower
  named follow-up.
- No public claim is strengthened beyond existing evidence.
- `ni status` remains `READY`, and validation commands pass.

Risks:

- Status wording becomes too cautious and hides verified source or zip paths.
  Mitigation: preserve the narrow verified paths while keeping broad host
  availability `Experimental` or `Unverified`.
- A model pack phrase implies CLI authority. Mitigation: keep "Skills are UX;
  CLI is authority" near user-facing pack guidance and in skill README files.
- Static checks become brittle. Mitigation: check durable concepts such as
  `Experimental`, host-level/global host verification, no-terminal assisted
  wording, and CLI authority rather than exact paragraphs.
- Host verification is tempted mid-pass. Mitigation: leave host verification as
  a deferred audit unless this task explicitly gathers reproducible host-level
  evidence.

Follow-up task: Choose between conversation proof capture reliability,
change-control UX audit, or the third benchmark case after GRILL-005 is closed.

## Tasks deferred

- Homebrew: deferred because it remains `Planned` until tap, formula, checksums,
  audit, install, published tap install, `ni --help`, and `ni version` evidence
  all pass.
- Landing page: deferred because status correctness and claim discipline are
  higher-priority trust work than marketing surface expansion.
- Additional benchmark cases: deferred because benchmark claim boundaries were
  just clarified and another case should reuse those patterns later.
- Change-control UX: deferred because it is important but has higher
  lock-semantics risk than preserving model workspace status wording.
- Model pack availability upgrade: deferred because broad availability requires
  host-level install or discovery evidence.
- Downstream integrations: deferred because integrations must remain separate
  packages, target exports, seed formats, or downstream-owned notes rather than
  `ni-kernel` behavior.

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
