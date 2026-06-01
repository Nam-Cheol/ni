# v0.5 work packet selection

## Current locked state

- Readiness: `go run ./cmd/ni status --dir . --proof --next-questions` reports `READY`, with no blockers, deferrals, or warnings.
- Lock state: `.ni/plan.lock.json` exists and records `READY`; it was locked at `2026-06-01T06:09:41Z`.
- Generated prompt path: `.ni/generated/v0.5-roadmap.prompt.txt` exists and is exactly 4000 characters by `wc -c`.
- v0.5 direction: CAP-019, REQ-019, EVAL-023, RISK-018, DEC-017, and the roadmap docs focus v0.5 on evidence quality, conversation-authoring reliability, ni-grill quality, change-control UX, broader non-software product examples, factual adoption hardening, and separate-package-only downstream integrations. The kernel remains pre-runtime and non-executing.

Scoring uses 1-5 where higher is better. `Cost` means lower expected cost scores higher. `Boundary risk` means lower risk to the ni-kernel boundary scores higher.

## Candidate comparison

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Acceptance evidence clarity pass | 5 | 5 | 5 | 4 | 5 | 5 | 29 | First |
| Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Defer until evidence criteria are clearer |
| Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; audit first before semantics change |
| ni-grill dogfood repair pass | 4 | 4 | 4 | 4 | 5 | 5 | 26 | Defer; partly covered by the selected packet |
| Model workspace host verification audit | 3 | 4 | 3 | 3 | 3 | 3 | 19 | Defer; avoid host-availability overclaim |
| Homebrew implementation plan | 3 | 3 | 2 | 2 | 3 | 2 | 15 | Defer; Planned until tested evidence exists |
| Landing page decision / implementation | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; lower priority than trust evidence |

## Selected first work packet

Acceptance evidence clarity pass.

Implementation note: this packet is implemented by
[`95_V0_5_ACCEPTANCE_EVIDENCE.md`](95_V0_5_ACCEPTANCE_EVIDENCE.md).

## Why this first

This packet directly addresses GRILL-003, the remaining medium acceptance-evidence note from the ni-grill dogfood report. It improves the trust surface that all other v0.5 work depends on: benchmark cases, conversation reliability, change-control UX, model-pack claims, Homebrew status, and downstream seed work all need visible evidence criteria before a later task can credibly call them complete.

It beats the third benchmark case because a new case should reuse a clear evidence checklist instead of inventing its own completion bar. It beats change-control UX because that work can drift into lock semantics if it starts before evidence expectations are explicit. It beats Homebrew, landing page, and model-pack availability work because those are adoption surfaces with higher public-claim risk and lower immediate evidence value.

## Work packet definition

Title: v0.5 acceptance evidence clarity pass.

Goal: Define the minimum evidence package that counts as completion for each v0.5 lane: benchmark evidence, conversation-authoring reliability, ni-grill quality, change-control UX, Homebrew readiness, model workspace pack availability, no-terminal proof, product-surface examples, downstream seed formats, and landing-page adoption claims.

Scope:

- Add a concise v0.5 acceptance-evidence checklist.
- Tie the checklist to CAP-019, REQ-019, EVAL-023, RISK-016, RISK-018, DEC-015, and DEC-017.
- Update roadmap and dogfood references so GRILL-003 is either addressed or explicitly retained with a narrower follow-up.
- Preserve `not_measured` boundaries for benchmark and adoption claims.
- Preserve Homebrew as `Planned` and model workspace packs as `Experimental` unless verified host-level evidence exists.
- Keep all downstream integration work as separate packages, seed formats, or notes outside `ni-kernel`.

Non-goals:

- Do not implement benchmark case 3.
- Do not change lock/relock/amendment semantics.
- Do not implement Homebrew, package-manager distribution, landing page publishing, model host installation, downstream integrations, or runtime behavior.
- Do not execute `.ni/generated/v0.5-roadmap.prompt.txt`.
- Do not run Codex exec or add adapters, queues, PR automation, issue publishing, release automation, shell execution, or user-facing contract `add` / `list` / `set` commands.
- Do not edit `.ni/plan.lock.json` manually or relock.

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

- A v0.5 acceptance-evidence checklist exists and names one evidence shape per roadmap lane.
- Each checklist lane states what is measured, what remains `not_measured`, and which file or command can prove completion.
- GRILL-003 is updated from an open medium note to addressed, or retained with a smaller named follow-up.
- The roadmap points future v0.5 tasks to the checklist before implementation.
- No public claim is strengthened beyond existing evidence.
- `ni status` remains `READY`, and validation commands pass.

Risks:

- Evidence checklist becomes too broad to use. Mitigation: keep each lane to one minimum proof shape and one boundary note.
- Benchmark or adoption claims overreach. Mitigation: keep `not_measured`, Homebrew `Planned`, and model-pack `Experimental` wording next to the relevant evidence row.
- Docs drift from locked source-of-truth. Mitigation: do not edit locked `docs/plan/**` or `.ni/contract.json`; use this as a post-lock planning-support doc unless a later amendment explicitly changes root planning state.

Follow-up task: Select or create the third benchmark case using the accepted v0.5 evidence checklist.

## Tasks deferred

- Homebrew: deferred because it remains `Planned` until tap, formula, checksums, audit, install, published tap install, `ni --help`, and `ni version` evidence all pass.
- Landing page: deferred because README remains the canonical quick entry and adoption-page work has lower evidence value than the trust checklist.
- Additional benchmark cases: deferred until the evidence checklist defines what every future benchmark must prove and keep `not_measured`.
- Model pack availability upgrade: deferred because host-level install or discovery is still unverified for broad availability claims.
- Downstream integrations: deferred because they must remain separate packages, target exports, seed formats, or downstream-owned notes rather than `ni-kernel` behavior.

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
