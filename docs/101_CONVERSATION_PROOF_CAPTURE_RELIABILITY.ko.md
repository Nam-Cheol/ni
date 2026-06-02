# Conversation Proof Capture Reliability

## Current status

Conversation proof capture는 conversation-driven authoring을 위한 planning audit
trail로 문서화되어 있다. `docs/83_CONVERSATION_PROOF_CAPTURE.md`, `ni-start`,
`ni-grill`, benchmark examples, model workspace skill packs는 이미 planning proof와
execution evidence를 구분한다.

현재 factual boundaries는 다음과 같다.

- release binary: verified release assets에 대해 Available;
- curl installer: verified release assets에 대해 Available;
- Homebrew: Planned / v0.5 candidate only;
- model workspace packs: broad product path로 Experimental;
- no-terminal method: Experimental / assisted;
- `ni-kernel`: pre-runtime Project Intent Compiler;
- runtime execution, shell adapters, Codex exec adapters, queues, PR
  automation, release automation, downstream execution layers: 포함하지 않음.

Conversation proof text는 planning state를 더 inspectable하게 만들 수 있다. 하지만
deterministic CLI authority를 대체할 수 없다. `ni status`가 readiness를 결정하고,
`ni end`가 `.ni/plan.lock.json`을 쓰며, `ni run`이 lock hashes를 검증한 뒤 bounded
handoff prompt를 compile한다.

## Reliability goal

Reliable proof capture란 reader가 planning claim을 exact conversation turn, changed
planning artifacts, affected contract IDs, CLI readiness output, lock state,
bounded handoff surface로 trace할 수 있고, 이것을 implementation evidence로
혼동하지 않는 상태를 뜻한다.

Reliable proof capture는 다음을 지켜야 한다.

- 무엇이 바뀌었고 무엇이 바뀌지 않았는지 이름 붙인다.
- tentative 또는 inferred statements를 assumptions 또는 open questions로 유지한다.
- conversation proof, acceptance evidence, benchmark evidence, readiness status,
  lock hash, locked plan, downstream handoff prompt, real implementation
  evidence의 차이를 보존한다.
- proof text가 `ni status --proof --next-questions`에서 온 경우를 보여준다.
- model workspace proof wording 근처에 "Skills are UX; CLI is authority."를
  유지한다.
- trusted CLI run이 docs와 contract를 validate하기 전까지 no-terminal proof를
  draft-only로 유지한다.

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

이 lifecycle은 directional하다. Model은 proof-related planning text를 draft하거나
explain할 수 있지만, CLI authority가 필요한 단계는 CLI surface를 통해서만 진행된다.

## What proof capture can support

| Claim | Supported? | Evidence surface | Notes |
| --- | --- | --- | --- |
| Planning conversation이 user intent를 captured했다 | Yes | Planning proof block; `docs/plan/**`; `.ni/contract.json` | Named planning turn과 changed records에 대해서만 supported. |
| Required questions가 asked되었다 | Yes | `ni status --proof --next-questions`; planning transcript 또는 proof block | Highest-priority question group은 가능하면 CLI output에서 와야 한다. |
| Answers가 preserved되었다 | Yes | Changed docs, contract IDs, session summary, planning proof | Tentative answers는 confirmed 전까지 assumptions 또는 open questions로 남아야 한다. |
| Readiness가 CLI surfaces로 evaluated되었다 | Yes | `ni status --proof --next-questions` output | CLI status는 정확히 quote하거나 faithful하게 summarize해야 한다. |
| Planning contract가 `READY`, `BLOCKED`, 또는 `READY_WITH_DEFERRALS`이다 | Yes | `ni status` output | Proof text는 status를 explain할 수 있지만 decide하지 않는다. |
| Downstream handoff가 locked plan에서 compiled되었다 | Yes | `.ni/plan.lock.json`; `ni run` output; prompt character count | Prompt compilation만 증명하며 downstream work는 증명하지 않는다. |
| Lock이 존재하고 locked hashes를 기록한다 | Yes | `ni end`가 쓴 `.ni/plan.lock.json` | Model은 lock을 hand-write하거나 repair하면 안 된다. |
| Planning records에 acceptance evidence가 named되었다 | Yes | `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md`; `docs/plan/**`; status proof | Planning acceptance evidence이지 implementation correctness가 아니다. |
| Benchmark case scope가 pre-runtime으로 유지되었다 | Yes | `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; benchmark examples; `demo-check` | Checked case artifacts와 `not_measured` boundaries에 대해서만 supported. |

## What proof capture cannot support

| Claim | Supported? | Why not | Required future evidence |
| --- | --- | --- | --- |
| Implementation correctness | No | Proof capture는 implementation을 run하거나 test하지 않는다. | `ni-kernel` 밖의 product-specific tests, reviews, runtime logs, implementation evidence. |
| Downstream agent success | No | `ni run`은 prompt만 compile한다. | Downstream-owned execution records와 evaluation plan. |
| Product readiness | No | `READY`는 declared planning contract scope에 적용된다. | Product acceptance tests, release criteria, operator approval, field evidence. |
| Benchmark effect size | No | Current benchmark cases는 qualitative readiness drills이다. | Repeated trials, measurement protocol, statistical analysis. |
| Adoption improvement | No | Proof capture는 external usage data를 collect하지 않는다. | Consent와 scope가 있는 user research, telemetry, adoption study. |
| Cost reduction | No | Runtime cost를 measure하지 않는다. | Cost baseline, repeated runs, measured deltas. |
| Latency reduction | No | Runtime latency를 measure하지 않는다. | Defined environment의 runtime benchmark. |
| Real-world approval | No | Synthetic 또는 planning approval은 external approval이 아니다. | Named real reviewer, approval artifact, applicable governance process. |
| Fieldwork authorization | No | Research fixture answers는 synthetic이며 pre-runtime이다. | Institutional, legal, safety, fieldwork authorization records. |
| Deterministic validation without CLI | No | Model-only proof는 draft audit trail이다. | Trusted `ni status`, `ni end`, `ni run` CLI output. |
| Homebrew availability | No | Proof capture는 package-manager publication과 무관하다. | Tap/formula publication plus `brew install`, `ni --help`, `ni version` evidence. |
| Broad model workspace availability | No | Model workspace packs는 broad product path로 Experimental이다. | Host-level install/discovery와 provider-specific usage verification. |

## Reliability risks

| Risk | Failure mode | Guardrail |
| --- | --- | --- |
| Model-generated proof text overstates readiness | Summary가 exact CLI output 없이 plan이 ready라고 말한다. | `READY`, `READY_WITH_DEFERRALS`, `BLOCKED`에는 `ni status`를 요구한다. |
| No-terminal workflow is mistaken for deterministic validation | Assisted draft proof가 trusted proof처럼 취급된다. | Trusted CLI run이 validate하기 전까지 no-terminal proof를 draft-only로 label한다. |
| Benchmark proof is mistaken for implementation quality | `READY` benchmark artifact가 product quality로 인용된다. | Benchmark claims 근처에 `artifact-readiness only`, `not_measured`, no-execution labels를 둔다. |
| Model workspace skill output is mistaken for CLI authority | Skill wording이 readiness engine처럼 보인다. | "Skills are UX; CLI is authority."를 보존하고 readiness, lock, prompt claims에는 CLI commands를 요구한다. |
| Stale lock or changed intent is not clearly surfaced | Docs 또는 intent 변경 뒤에도 handoff prompt를 trusted로 취급한다. | Stale lock 또는 hash mismatch에서 stop하고 `BLOCKED`를 report한다. |
| Examples imply downstream execution success | Prompt compilation이 downstream completion처럼 읽힌다. | Generated prompts는 inert handoff material이며 `ni`가 실행하지 않는다고 말한다. |

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
| `README.md` and `README.ko.md` | `ni-kernel` boundary, status vocabulary, no-terminal assisted wording, model workspace Experimental status를 preserve한다. |
| `docs/83_CONVERSATION_PROOF_CAPTURE.md` | Proof capture, lifecycle, CLI authority, no-terminal draft limits, model workspace skill limits를 define한다. |
| `docs/95_V0_5_ACCEPTANCE_EVIDENCE.md` | Acceptance evidence를 planning evidence로 유지하고 execution evidence로 만들지 않는다. |
| `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md` | `not_measured`, artifact-readiness, synthetic-fixture, no-execution boundaries를 visible하게 유지한다. |
| `docs/99_MODEL_WORKSPACE_STATUS.md` | Experimental model workspace status, `not_verified` host/provider claims, CLI authority를 preserve한다. |
| `docs/100_V0_5_WORK_PACKET_COMPLETION_AUDIT.md` | 이 reliability pass의 selection record와 prior GRILL closure evidence를 preserve한다. |
| Examples | Downstream execution, product readiness, real research approval, fieldwork authorization, no-terminal deterministic validation, benchmark effect size를 imply하지 않는다. |
| Skills | Skills may help draft or explain proof-related planning text, but they do not determine readiness, lock plans, or replace `ni status`, `ni end`, or `ni run`. |

## Validation surface

Current validation은 여러 proof-capture overclaim risks를 cover한다.

- `go run ./cmd/ni status --dir . --proof --next-questions`는 현재 repository
  readiness를 CLI로 verify한다.
- `python3 scripts/check-install-docs.py`는 Homebrew와 model workspace
  availability를 포함한 install/distribution status claims를 보호한다.
- `bash scripts/check-skill-packs.sh`는 model workspace skill metadata,
  Experimental status, CLI authority wording을 보호한다. 또한 `docs/83`, 이 문서,
  skill pack README files의 durable proof-capture reliability markers를 확인한다.
- `bash scripts/demo-check.sh`는 benchmark, no-terminal, ni-grill boundary
  wording을 확인하고 generated prompt execution을 피한다.
- `bash scripts/quality.sh`는 static documentation checks, Go files가 있을 때 Go
  tests, smoke checks를 실행한다.
- `bash scripts/smoke.sh`, `bash scripts/install-check.sh`,
  `bash scripts/release-check.sh`는 Homebrew 또는 model workspace status를 바꾸지
  않으면서 CLI smoke, local install, release readiness surfaces를 cover한다.

Broad prose changes, new examples, new benchmark claims, new no-terminal copy,
future amendment/relock wording에는 여전히 manual audit이 필요하다. Static checks는
모든 문장을 고정하기보다 durable boundary phrases를 가볍게 보호해야 한다.

## Follow-up candidates

- stale-lock proof wording audit;
- amend/relock UX audit;
- third benchmark proof-capture case;
- model workspace proof wording verification;
- no-terminal proof capture examples.
