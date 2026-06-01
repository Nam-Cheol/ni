# v0.5 second work packet selection

## Current locked state

- Readiness: `go run ./cmd/ni status --dir . --proof --next-questions` reports `READY`, with no blockers, deferrals, or warnings.
- Lock state: `.ni/plan.lock.json` exists and records `READY`; it was locked at `2026-06-01T06:09:41Z`.
- Generated prompt path: `.ni/generated/v0.5-roadmap.prompt.txt` exists and is exactly 4000 characters by `wc -c`.
- v0.5 direction: real benchmark evidence, conversation-authoring reliability, ni-grill quality, change-control UX, broader non-software product examples, model workspace pack verification, Homebrew as `Planned` / v0.5 candidate, separate-package-only downstream integrations, and a pre-runtime non-executing `ni-kernel`.
- First work packet completion: Task 175 added `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md` and `docs/95_V0_5_ACCEPTANCE_EVIDENCE.ko.md`, linked them from roadmap, benchmark/example docs, ni-grill dogfood, and work-packet selection docs, and addressed GRILL-003 as an evidence-criteria clarification.
- Remaining GRILL notes: GRILL-004 remains a benchmark claim-boundary note, and GRILL-005 remains a model workspace status preservation note.

Scoring uses 1-5 where higher is better. `Cost` means lower expected cost scores higher. `Boundary risk` means lower risk to the ni-kernel boundary scores higher.

## Candidate comparison

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| Benchmark claim-boundary pass | 5 | 5 | 5 | 4 | 5 | 5 | 29 | First |
| Model workspace status preservation pass | 4 | 4 | 4 | 4 | 3 | 5 | 24 | Defer; important but higher overclaim risk |
| Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; audit after benchmark boundary cleanup |
| Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Defer until benchmark boundary rules are more visible |
| Conversation proof capture reliability pass | 4 | 5 | 4 | 4 | 5 | 4 | 26 | Defer; strong follow-up after benchmark claim boundaries |
| Model workspace host verification audit | 3 | 4 | 3 | 3 | 3 | 3 | 19 | Defer; avoid global host availability overclaim |
| Homebrew implementation plan | 3 | 3 | 2 | 2 | 3 | 2 | 15 | Defer; `Planned` until tested evidence exists |
| Landing page decision / implementation | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; lower priority than benchmark/reliability trust work |

## Selected second work packet

Benchmark claim-boundary pass.

## Why this second

The first v0.5 packet defined what evidence should count. The next useful move is
to apply that evidence discipline to the most public and trust-sensitive v0.5
surface: benchmark evidence. GRILL-004 already says the benchmark claims are
generally well-scoped but that visibility can degrade when long pages bury
`not_measured` boundaries below status and prompt excerpts.

This beats a third benchmark case because new cases should inherit prominent
claim-boundary patterns before the repository adds more evidence volume. It
beats model workspace verification because model-pack status is already guarded
by `Experimental` wording and checks, while benchmark evidence is the central
credibility lane for v0.5. It beats change-control UX because this pass can
close an existing GRILL note with low kernel-boundary risk before touching
lock/relock semantics.

## Work packet definition

Title: v0.5 benchmark claim-boundary pass.

Goal: Make measured / `not_measured` benchmark boundaries more prominent and
consistent across benchmark docs, benchmark examples, example coverage, and
demo checks without adding new benchmark cases or stronger empirical claims.

Scope:

- Use the real benchmark evidence lane in `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`.
- Review benchmark summary, case README files, before/after evidence files, and ni-grill benchmark examples for boundary visibility.
- Put `not_measured` and non-execution boundaries next to `READY` transition summaries where a reader might quote a readiness result alone.
- Add or adjust checks only for boundary wording and required files, not for new empirical outcomes.
- Update `docs/93_NI_GRILL_DOGFOOD.md` and Korean companion to mark GRILL-004 addressed only if the pass actually makes the boundary more visible.

Non-goals:

- Do not create the third benchmark case.
- Do not rerun or execute generated prompts.
- Do not run model APIs, downstream agents, shell adapters, queues, or product implementations.
- Do not claim benchmark data proves implementation quality, downstream-agent performance, adoption, cost, latency, rework reduction, or statistical effect size.
- Do not change lock/relock semantics, `.ni/contract.json`, `.ni/session.json`, or `.ni/plan.lock.json`.
- Do not mark Homebrew `Available`, upgrade model workspace packs to global availability, or claim no-terminal deterministic validation.

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

- Benchmark summary docs keep `not_measured` boundaries near each benchmark `READY` or before/after transition claim.
- Internal-dashboard and research-protocol case docs preserve non-execution and artifact-readiness-only wording near their main readiness summaries.
- Example coverage points future benchmark work to the v0.5 evidence matrix and claim-boundary pattern.
- GRILL-004 is either marked addressed by this pass or retained with a narrower named follow-up.
- No benchmark/adoption claim is strengthened beyond existing evidence.
- `ni status` remains `READY`, and validation commands pass.

Risks:

- Boundary wording becomes repetitive. Mitigation: keep the boundary near the main claim, then link to detailed `not_measured` sections instead of duplicating every detail.
- A check becomes too brittle. Mitigation: check for durable boundary phrases and required files, not exact prose blocks.
- Benchmark evidence appears stronger than it is. Mitigation: preserve `not_measured`, manual qualitative, artifact-readiness-only, synthetic-fixture-only, and no-downstream-execution wording.

Follow-up task: Select the third benchmark case after benchmark claim-boundary patterns are stable.

## Tasks deferred

- Homebrew: deferred because it remains `Planned` until tap, formula, checksums, audit, install, published tap install, `ni --help`, and `ni version` evidence all pass.
- Landing page: deferred because README remains the canonical quick entry and trust/evidence work is higher priority.
- Additional benchmark cases: deferred until boundary wording is more visible and reusable for the next case.
- Model pack availability upgrade: deferred because broad availability requires host-level install or discovery evidence.
- Downstream integrations: deferred because they must remain separate packages, target exports, seed formats, or downstream-owned notes.
- Change-control UX: deferred because it is important but has higher lock-semantics risk than a benchmark boundary pass.

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
