# v0.5 Work Packet Completion Audit

## Current state

The first three v0.5 work packets are complete as documentation, example, and
static-check alignment work:

- release binary: Available for verified v0.4.0 assets;
- curl installer: Available for verified v0.4.0 assets;
- Homebrew: Planned / v0.5 candidate only;
- model workspace packs: Experimental as a broad product path;
- no-terminal method: Experimental / assisted;
- runtime execution, shell adapters, Codex exec, queues, downstream agents, and
  PR automation: Not included in `ni-kernel`.

The current authoritative CLI proof is:

```text
NI Intent Readiness: READY
Blockers: None.
Deferrals: None.
Warnings: None.
```

This audit did not run `ni end`, relock, edit `.ni/plan.lock.json`, edit
`.ni/contract.json`, edit `.ni/session.json`, execute generated prompts, or add
runtime behavior.

## GRILL closure status

| Finding | Status | Evidence | Notes |
| --- | --- | --- | --- |
| GRILL-001 | addressed | `docs/51_POST_RELEASE_ROADMAP.md`; `docs/53_DISTRIBUTION_STRATEGY.md`; `docs/80_HOMEBREW_DECISION.md` | Roadmap and distribution docs now match the v0.5 evidence/adoption direction while preserving kernel boundaries. |
| GRILL-002 | addressed | `docs/53_DISTRIBUTION_STRATEGY.md`; `docs/80_HOMEBREW_DECISION.md`; `README.md` | Release binary and curl installer are Available for verified v0.4.0 assets; Homebrew remains Planned. |
| GRILL-003 | addressed | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md` | The acceptance matrix defines lane-by-lane completion evidence and `not_measured` boundaries. |
| GRILL-004 | addressed | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `docs/77_BENCHMARK_CASE_STUDY.md`; `examples/benchmark-report/**`; `scripts/demo-check.sh` | Benchmark summaries keep `READY`, artifact-readiness, synthetic-fixture, non-execution, and `not_measured` boundaries visible. |
| GRILL-005 | addressed | `docs/99_MODEL_WORKSPACE_STATUS.md`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `packages/*-skills/README.md`; `scripts/check-skill-packs.sh` | Model workspace packs remain Experimental as a broad path; host/global install and provider behavior remain not verified. Skills are UX; CLI is authority. |

No GRILL-003 through GRILL-005 finding remains open. Host-level model workspace
verification remains future work, not a closure gap.

## Work packet completion

| Work packet | Evidence | GRILL finding addressed | Complete? | Notes |
| --- | --- | --- | --- | --- |
| Acceptance evidence clarity | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`; `docs/95_V0_5_ACCEPTANCE_EVIDENCE.ko.md` | GRILL-003 | yes | Defines evidence shape, status vocabulary, and `not_measured` rules for v0.5 lanes. |
| Benchmark claim-boundary pass | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`; benchmark examples; `scripts/demo-check.sh` | GRILL-004 | yes | Keeps benchmark `READY` claims scoped to artifact readiness and synthetic fixtures, not product or research outcomes. |
| Model workspace status preservation | `docs/99_MODEL_WORKSPACE_STATUS.md`; `docs/99_MODEL_WORKSPACE_STATUS.ko.md`; skill pack READMEs; `scripts/check-skill-packs.sh` | GRILL-005 | yes | Preserves Experimental broad model workspace status, not_verified host/provider claims, no-terminal limits, and CLI authority wording. |

## Claim/status audit

| Claim area | Expected status | Actual status | Pass? |
| --- | --- | --- | --- |
| Release binary | Available for verified v0.4.0 release assets | README, distribution, and Homebrew decision docs say Available for verified v0.4.0 assets. | yes |
| Curl installer | Available for verified v0.4.0 release assets | README, distribution, install docs, and Homebrew decision docs say Available for verified v0.4.0 assets. | yes |
| Homebrew | Planned; no tap or formula published/tested | README and Homebrew docs say Planned and forbid Available wording until tap/formula/install proof exists. | yes |
| Model workspace packs | Experimental as a broad product path | README, distribution docs, `docs/99`, and pack READMEs say Experimental and preserve not_verified host/provider claims. | yes |
| No-terminal method | Experimental / assisted; deterministic validation requires CLI proof | README, distribution docs, and no-terminal/example docs preserve assisted-only wording. | yes |
| Internal-dashboard benchmark | `READY` for benchmark planning-meeting artifact readiness only | Benchmark docs and examples preserve dashboard product readiness, implementation quality, and downstream-agent performance as `not_measured`. | yes |
| Research-protocol benchmark | `READY` for synthetic benchmark protocol planning artifact readiness only | Benchmark docs and examples preserve real research approval, fieldwork authorization, research quality, and intervention effectiveness as `not_measured`. | yes |
| Runtime execution boundary | Not included in `ni-kernel` | README, roadmap, example coverage, grill docs, and skills exclude runtime execution, shell adapters, Codex exec, queues, downstream agents, and PR automation. | yes |

## Validation surface

| Check area | Evidence | Status |
| --- | --- | --- |
| Benchmark claim boundaries | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `scripts/demo-check.sh` | present |
| Model workspace overclaim prevention | `docs/99_MODEL_WORKSPACE_STATUS.md`; `scripts/check-skill-packs.sh`; `scripts/check-install-docs.py` | present |
| Skill pack metadata and authority wording | `scripts/check-skill-packs.sh` | present |
| Install docs consistency | `scripts/check-install-docs.py`; `bash scripts/quality.sh` | present |
| Demo coverage | `scripts/demo-check.sh`; `docs/82_EXAMPLE_COVERAGE.md` | present |

## Next direction candidates

Scoring uses 1-5 where higher is better. `Cost` means lower expected cost scores
higher. `Boundary risk` means lower risk to the `ni-kernel` boundary scores
higher.

| Candidate | User impact | Roadmap alignment | Evidence value | Cost | Boundary risk | Dependency readiness | Score | Recommendation |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | --- |
| A. Change-control UX audit | 4 | 5 | 4 | 3 | 3 | 4 | 23 | Defer; useful but closer to lock/relock semantics. |
| B. Third benchmark case selection | 4 | 5 | 5 | 3 | 5 | 4 | 26 | Strong follow-up after proof-capture consistency. |
| C. Conversation proof capture reliability pass | 5 | 5 | 4 | 4 | 5 | 5 | 28 | Selected. |
| D. Model workspace host verification audit | 3 | 4 | 3 | 2 | 2 | 2 | 16 | Defer; host claims are easy to overstate. |
| E. Homebrew implementation audit | 3 | 4 | 3 | 2 | 2 | 3 | 17 | Defer; public availability-claim risk remains high. |
| F. Landing page decision | 2 | 2 | 2 | 4 | 5 | 3 | 18 | Defer; lower priority than product reliability. |

## Selected next direction

Conversation proof capture reliability pass.

## Why this next

The first three v0.5 packets closed the GRILL evidence and claim-boundary notes.
The next highest-value move is to make the sustained authoring loop more
consistent: how a model records user answers, names changed files and IDs,
shows before/after `ni status` proof, and preserves the split between planning
proof and execution evidence.

This beats the third benchmark case because new benchmarks should benefit from
clean proof-capture patterns rather than adding more evidence volume first. It
beats change-control UX because proof capture has lower lock-semantics risk and
is a dependency for explaining later amend/relock/stale-lock work. It beats
Homebrew, model host verification, and landing page work because those surfaces
carry more public-claim or adoption-message risk.

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
