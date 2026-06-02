# Conversation Proof Capture Reliability

## Current status

Conversation proof capture is documented as a planning audit trail for
conversation-driven authoring. `docs/83_CONVERSATION_PROOF_CAPTURE.md`,
`ni-start`, `ni-grill`, benchmark examples, and model workspace skill packs
already distinguish planning proof from execution evidence.

The current factual boundaries are:

- release binary: Available for verified release assets;
- curl installer: Available for verified release assets;
- Homebrew: Planned / v0.5 candidate only;
- model workspace packs: Experimental as a broad product path;
- no-terminal method: Experimental / assisted;
- `ni-kernel`: pre-runtime Project Intent Compiler;
- runtime execution, shell adapters, Codex exec adapters, queues, PR
  automation, release automation, and downstream execution layers: not included.

Conversation proof text can make planning state easier to inspect. It cannot
replace deterministic CLI authority. `ni status` decides readiness, `ni end`
writes `.ni/plan.lock.json`, and `ni run` verifies lock hashes before compiling
a bounded handoff prompt.

## Reliability goal

Reliable proof capture means a reader can trace a planning claim to the exact
conversation turn, changed planning artifacts, affected contract IDs, CLI
readiness output, lock state, and bounded handoff surface without confusing
that trail for implementation evidence.

Reliable proof capture must:

- name what changed and what did not change;
- keep tentative or inferred statements as assumptions or open questions;
- preserve the difference between conversation proof, acceptance evidence,
  benchmark evidence, readiness status, lock hash, locked plan, downstream
  handoff prompt, and real implementation evidence;
- show when proof text came from `ni status --proof --next-questions`;
- keep "Skills are UX; CLI is authority." visible near model workspace proof
  wording;
- keep no-terminal proof draft-only until a trusted CLI run validates the docs
  and contract.

## Proof capture lifecycle

```text
planning conversation
-> docs/plan and .ni/contract.json
-> ni status --proof --next-questions
-> readiness explanation
-> ni end lock
-> .ni/plan.lock.json
-> ni run bounded handoff prompt
```

The lifecycle is directional. A model may draft or explain proof-related
planning text, but the lifecycle advances only through CLI surfaces where CLI
authority is required.

## What proof capture can support

| Claim | Supported? | Evidence surface | Notes |
| --- | --- | --- | --- |
| The planning conversation captured user intent | Yes | Planning proof block; `docs/plan/**`; `.ni/contract.json` | Supported only for the named planning turn and changed records. |
| Required questions were asked | Yes | `ni status --proof --next-questions`; planning transcript or proof block | The highest-priority question group must come from CLI output when available. |
| Answers were preserved | Yes | Changed docs, contract IDs, session summary, planning proof | Tentative answers must remain assumptions or open questions until confirmed. |
| Readiness was evaluated by CLI surfaces | Yes | `ni status --proof --next-questions` output | The CLI status must be quoted or summarized faithfully. |
| A planning contract is `READY`, `BLOCKED`, or `READY_WITH_DEFERRALS` | Yes | `ni status` output | Proof text may explain the status but does not decide it. |
| Downstream handoff was compiled from a locked plan | Yes | `.ni/plan.lock.json`; `ni run` output; prompt character count | This proves prompt compilation only, not downstream work. |
| A lock exists and records locked hashes | Yes | `.ni/plan.lock.json` written by `ni end` | Models must not write or repair the lock by hand. |
| Acceptance evidence was named for planning records | Yes | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`; `docs/plan/**`; status proof | This is planning acceptance evidence, not implementation correctness. |
| Benchmark case scope stayed pre-runtime | Yes | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; benchmark examples; `demo-check` | Supported for checked case artifacts and `not_measured` boundaries. |

## What proof capture cannot support

| Claim | Supported? | Why not | Required future evidence |
| --- | --- | --- | --- |
| Implementation correctness | No | Proof capture does not run or test implementation. | Product-specific tests, reviews, runtime logs, or implementation evidence outside `ni-kernel`. |
| Downstream agent success | No | `ni run` compiles a prompt only. | Downstream-owned execution records and evaluation plan. |
| Product readiness | No | `READY` applies to the declared planning contract scope. | Product acceptance tests, release criteria, operator approval, or field evidence. |
| Benchmark effect size | No | Current benchmark cases are qualitative readiness drills. | Repeated trials, measurement protocol, and statistical analysis. |
| Adoption improvement | No | No external usage data is collected by proof capture. | User research, telemetry, or adoption study with consent and scope. |
| Cost reduction | No | No runtime cost is measured. | Cost baseline, repeated runs, and measured deltas. |
| Latency reduction | No | No runtime latency is measured. | Runtime benchmark with defined environment. |
| Real-world approval | No | Synthetic or planning approval is not external approval. | Named real reviewer, approval artifact, and applicable governance process. |
| Fieldwork authorization | No | Research fixture answers are synthetic and pre-runtime. | Institutional, legal, safety, and fieldwork authorization records. |
| Deterministic validation without CLI | No | Model-only proof is a draft audit trail. | Trusted `ni status`, `ni end`, and `ni run` CLI output. |
| Homebrew availability | No | Proof capture is unrelated to package-manager publication. | Tap/formula publication plus `brew install`, `ni --help`, and `ni version` evidence. |
| Broad model workspace availability | No | Model workspace packs remain Experimental as a broad product path. | Host-level install/discovery and provider-specific usage verification. |

## Reliability risks

| Risk | Failure mode | Guardrail |
| --- | --- | --- |
| Model-generated proof text overstates readiness | A summary says a plan is ready without exact CLI output. | Require `ni status` for `READY`, `READY_WITH_DEFERRALS`, and `BLOCKED`. |
| No-terminal workflow is mistaken for deterministic validation | Assisted draft proof is treated as trusted proof. | Label no-terminal proof draft-only until a trusted CLI run validates it. |
| Benchmark proof is mistaken for implementation quality | `READY` benchmark artifact is quoted as product quality. | Keep `artifact-readiness only`, `not_measured`, and no-execution labels near benchmark claims. |
| Model workspace skill output is mistaken for CLI authority | Skill wording sounds like a readiness engine. | Preserve "Skills are UX; CLI is authority." and require CLI commands for readiness, lock, and prompt claims. |
| Stale lock or changed intent is not clearly surfaced | A handoff prompt is trusted after docs or intent changed. | Stop on stale lock or hash mismatch and report `BLOCKED`. |
| Examples imply downstream execution success | Prompt compilation is read as downstream completion. | State that generated prompts are inert handoff material and are not executed by `ni`. |

## Required wording rules

| Say this | Do not say this |
| --- | --- |
| "`ni status` reports `READY` for the declared planning contract scope." | "The product is ready." |
| "The proof block summarizes changed planning artifacts and affected IDs." | "The model proved readiness." |
| "`ni end` wrote `.ni/plan.lock.json`." | "The model locked the plan." |
| "`ni run` compiled a bounded handoff prompt from a valid lock." | "`ni` ran the downstream work." |
| "No-terminal proof is draft-only until a trusted CLI run validates it." | "No-terminal mode deterministically validates the plan." |
| "Model workspace packs are Experimental as a broad product path." | "Model workspace packs are Available globally." |
| "Benchmark `READY` is artifact-readiness only." | "The benchmark proves implementation quality or product impact." |
| "The research-protocol fixture does not prove real approval or fieldwork authorization." | "The protocol is approved for fieldwork." |
| "Homebrew remains Planned until tap/formula install evidence exists." | "Homebrew is Available." |

## Documentation alignment checklist

Skills may help draft or explain proof-related planning text. Skills do not
determine readiness, lock plans, or replace `ni status`, `ni end`, or `ni run`.

| Surface | Required alignment |
| --- | --- |
| `README.md` and `README.ko.md` | Preserve `ni-kernel` boundary, status vocabulary, no-terminal assisted wording, and model workspace Experimental status. |
| `docs/83_CONVERSATION_PROOF_CAPTURE.md` | Define proof capture, lifecycle, CLI authority, no-terminal draft limits, and model workspace skill limits. |
| `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md` | Keep acceptance evidence as planning evidence, not execution evidence. |
| `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md` | Keep `not_measured`, artifact-readiness, synthetic-fixture, and no-execution boundaries visible. |
| `docs/99_MODEL_WORKSPACE_STATUS.md` | Preserve Experimental model workspace status, `not_verified` host/provider claims, and CLI authority. |
| `docs/100_V0_5_WORK_PACKET_COMPLETION_AUDIT.md` | Preserve the selection record for this reliability pass and the prior GRILL closure evidence. |
| Examples | Avoid implying downstream execution, product readiness, real research approval, fieldwork authorization, no-terminal deterministic validation, or benchmark effect size. |
| Skills | Say that skills may help draft or explain proof-related planning text, but they do not determine readiness, lock plans, or replace `ni status`, `ni end`, or `ni run`. |

## Validation surface

Current validation covers several proof-capture overclaim risks:

- `go run ./cmd/ni status --dir . --proof --next-questions` verifies current
  repository readiness through the CLI.
- `python3 scripts/check-install-docs.py` protects install and distribution
  status claims, including Homebrew and model workspace availability.
- `bash scripts/check-skill-packs.sh` protects model workspace skill metadata,
  Experimental status, and CLI authority wording. It also checks the durable
  proof-capture reliability markers in `docs/83`, this document, and skill pack
  README files.
- `bash scripts/demo-check.sh` checks examples for benchmark, no-terminal, and
  ni-grill boundary wording and avoids generated prompt execution.
- `bash scripts/quality.sh` runs the static documentation checks, Go tests when
  Go files exist, and smoke checks.
- `bash scripts/smoke.sh`, `bash scripts/install-check.sh`, and
  `bash scripts/release-check.sh` cover CLI smoke, local install, and release
  readiness surfaces without changing Homebrew or model workspace status.

Manual audit still remains necessary for broad prose changes, new examples,
new benchmark claims, new no-terminal copy, and any future amendment/relock
wording. Static checks should stay lightweight and protect durable boundary
phrases rather than freeze every sentence.

## Follow-up candidates

- stale-lock proof wording audit;
- amend/relock UX audit;
- third benchmark proof-capture case;
- model workspace proof wording verification;
- no-terminal proof capture examples.
