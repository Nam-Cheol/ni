# v0.4.1 안정화 릴리스 계획

Date: 2026-05-29

Status: draft release plan only. 이 문서는 release를 publish하지 않고, tag를
create하지 않고, assets를 upload하지 않고, Homebrew를 Available로 표시하지
않고, runtime execution을 추가하지 않는다.

## Goal

v0.4 adoption hardening 이후 작은 안정화 릴리스를 준비한다. 이번 릴리스는
기존 pre-runtime authoring path를 더 신뢰하기 쉽고, 검증하기 쉽고, 설명하기
쉽게 만드는 데 집중한다. `ni-kernel`을 runner로 확장하지 않는다.

Release story는 다음과 같다:

```text
ni init -> ni-start conversation -> ni status proof -> ni end -> ni run prompt
```

`ni run`은 prompt compiler로만 남는다. Codex, shell, model API, queue,
adapter, SPEC workflow, downstream work를 실행하면 안 된다.

## Included

| Area | v0.4.1 stabilization scope | Evidence or source |
| --- | --- | --- |
| `ni-start` conversation UX hardening | First-run과 resume behavior를 더 따라가기 쉽게 만든다. Current state를 summarize하고, stable IDs를 보존하고, focused questions를 1개에서 3개 묻고, planning edits 이후 changed files를 보여준다. | [Conversation Authoring UX Audit](79_CONVERSATION_AUTHORING_UX_AUDIT.ko.md), [ni-start behavior](31_NI_START_BEHAVIOR.ko.md) |
| Readiness proof clarity | `ni status --proof --next-questions`를 visible readiness explanation으로 유지한다. Blockers, deferrals, warnings, passed checks와 model judgment가 readiness를 선언할 수 없다는 규칙을 더 명확히 한다. | [Status proof](44_STATUS_PROOF.md), [Readiness interview](34_READINESS_INTERVIEW.md) |
| Docs/contract sync diagnostics | Docs와 `.ni/contract.json` drift에 대한 user-facing `R012` repair language를 안정화한다: affected ID, location, problem, why it matters, suggested repair, `ni-end` blocking 여부. | [Docs-contract sync](33_DOCS_CONTRACT_SYNC.md) |
| Next-question improvements | `BLOCKED` 이후 deterministic `next_questions`를 normal planning loop로 포장한다. Questions는 ID를 보존하고 incomplete decisions를 accept하도록 압박하지 않는다. | [Readiness interview](34_READINESS_INTERVIEW.md), [Status proof](44_STATUS_PROOF.md) |
| Model workspace pack UX | Codex와 Claude packs를 UX layers로 유지한다. Project 열기, `ni-start` invoke, CLI proof 보존, lock mismatch에서 stop하는 task-first usage language를 개선한다. | [Model Workspace Packs](55_MODEL_WORKSPACE_PACKS.md), [Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.md) |
| Example coverage | Blocked ambiguity, conversation authoring, research protocol planning, conversation products, bounded handoff prompts를 보여주는 checked examples를 유지한다. Downstream execution은 하지 않는다. | [Examples](18_EXAMPLES.md), [Demo Verification](48_DEMO_VERIFICATION.md) |
| No-terminal assisted workflow | Terminal을 쓸 수 없는 사용자의 proof-capture story를 개선한다. Model과 docs/contract를 draft한 뒤, readiness, lock, prompt claims 전에 trusted runner가 만든 exact CLI output을 paste한다. | [No-Terminal Planning](no-terminal.ko.md) |
| Benchmark case study expansion | Benchmark evidence는 qualitative하고 transparent하게 유지한다. Case-study coverage는 static documented readiness comparison으로만 확장한다. External model APIs를 호출하거나 downstream work를 실행하지 않는다. | [Benchmark Case Studies](77_BENCHMARK_CASE_STUDY.ko.md), [Benchmark Protocol](43_BENCHMARK_PROTOCOL.md) |
| Homebrew decision | Tap, formula, checksums, audit, install, clean-environment verification이 release 전에 구현되지 않으면 Homebrew는 v0.4.1에서 Planned로 유지한다. | [Homebrew Decision](80_HOMEBREW_DECISION.ko.md) |

## Not Included

- Execution runtime.
- Codex exec adapter.
- Shell adapter.
- Task runner.
- SPEC runner.
- Queue.
- Multi-agent orchestration.
- Homebrew support if the tap and verification work are not implemented.

이 exclusions는 launch chores 누락이 아니라 release boundary다. 이 boundary는
`ni-kernel`을 Intent Lock Protocol에 집중하게 한다: docs contract, readiness
gate, lockfile, prompt compiler, source-of-truth rule.

## Release Candidate Validation Checklist

`v0.4.1` release candidate를 tag하거나 publish하기 전에 다음 checks를 실행한다:

- [ ] `go test ./...`
- [ ] quality: `bash scripts/quality.sh`
- [ ] smoke: `bash scripts/smoke.sh`
- [ ] demo-check: `bash scripts/demo-check.sh`
- [ ] install-check: `bash scripts/install-check.sh`
- [ ] release-check: `bash scripts/release-check.sh`
- [ ] fresh-install-check: `bash scripts/fresh-install-check.sh`
- [ ] skill-pack check: `bash scripts/check-skill-packs.sh`

이 checklist는 evidence collection일 뿐이다. Release assets를 publish하거나,
tags를 push하거나, Homebrew tap을 update하거나, unverified distribution path를
Available로 표시하지 않는다.

## Availability Rules

- Matching implementation과 verification이 repository 또는 published release
  assets에 있을 때만 해당 path를 Available로 표시한다.
- v0.4.1 release에 verified tap formula와 clean install proof가 포함되지 않으면
  Homebrew는 Planned로 유지한다.
- Model workspace packs는 UX로 유지한다. Planning을 guide할 수는 있지만
  `ni status`, `ni end`, `ni run`이 authority다.
- No-terminal assisted workflow는 trusted runner의 exact CLI proof가 제공되지
  않으면 Experimental로 유지한다.
- Benchmark results는 measured case studies에 한정한다. Statistical
  significance나 downstream implementation quality를 claim하지 않는다.

## Release Summary Draft

`v0.4.1`은 adoption-hardening surface를 위한 stabilization release다.
`ni init`에서 `ni-start`, readiness proof review, lock confirmation, bounded
prompt compilation까지 이어지는 경로를 다듬는다. Focus는 clarity다:
readiness proofs를 설명하기 쉽게 만들고, docs/contract drift를 repair하기 쉽게
만들고, next questions가 자연스러운 다음 planning turn처럼 느껴지게 만들고,
model workspace packs가 CLI authority를 보존하도록 돕는다.

이 릴리스는 product category를 바꾸지 않는다. `ni`는 계속 AI Agents를 위한
Project Intent Compiler이며, not a task runner, not an execution harness,
not a SPEC runner, not an adapter layer, not a queue, not an orchestration
system이다.
