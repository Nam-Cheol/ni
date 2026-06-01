# v0.5 Acceptance Evidence

## Purpose

이 문서는 v0.5 roadmap lanes에서 무엇을 completion evidence로 볼지 정의한다.
나중 작업이 어떤 lane을 done이라고 claim하기 전에, broad v0.5 direction을
inspectable evidence expectations로 바꾸기 위한 문서다.

이것은 planning acceptance evidence다. Execution evidence가 아니며,
downstream implementation quality를 증명하지 않고, runtime execution을
authorize하지 않는다. NI plan의 CLI readiness와 lock state는 계속 CLI가
authority다. `ni status`는 readiness를 결정하고, `ni end`는 lock을 쓰며,
`ni run`은 valid lock에서 bounded prompt를 compile한다.

## Evidence principles

1. Evidence는 inspectable해야 한다.
2. Evidence는 files, commands, documented decisions를 reference해야 한다.
3. Evidence는 `not_measured` boundaries를 보존해야 한다.
4. Evidence는 `Available`, `Experimental`, `Planned`를 구분해야 한다.
5. Evidence는 model judgment를 authority로 만들면 안 된다.
6. Evidence는 readiness gates를 약화하면 안 된다.
7. Evidence는 downstream execution을 암시하면 안 된다.

## Evidence matrix

| Lane | Completion claim | Required evidence | Verification command or file | Status vocabulary | Not measured / not claimed |
| --- | --- | --- | --- | --- | --- |
| Real benchmark evidence | v0.5 benchmark cases는 NI가 handoff 전에 readiness gaps를 어떻게 드러내는지 보여준다. | Case workspace 또는 docs; status proof; next questions; resolved case라면 before/after evidence; `not_measured` table; no downstream execution statement; visible claim-boundary markers. | `docs/77_BENCHMARK_CASE_STUDY.md`; `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`; `examples/benchmark-report/**`; `bash scripts/demo-check.sh` | `measured`; `not_measured`; `BLOCKED`; artifact-readiness에 한정된 `READY` | Implementation quality; downstream agent performance; adoption; cost; latency; statistical effect size |
| Conversation-authoring reliability | `ni-start` / `ni-grill` / `ni-end` flow가 더 따라가기 쉽고 안전하다. | First-run card docs; grouped next questions; planning proof capture; examples; skill pack wording. | `docs/31_NI_START_BEHAVIOR.md`; `docs/83_CONVERSATION_PROOF_CAPTURE.md`; `examples/ni-start-dogfood/**`; `bash scripts/check-skill-packs.sh` | `documented`; `verified by examples`; `Experimental` | Real user success rate; providers 전반의 model quality |
| ni-grill quality | `ni-grill`은 readiness gate가 되지 않으면서 planning quality를 challenge한다. | `docs/91_NI_GRILL.md`; `docs/92_NI_GRILL_OUTPUT_BUDGET.md`; ni-grill examples; severity/budget rules; CLI authority boundary. | `examples/ni-grill/**`; `bash scripts/demo-check.sh`; `bash scripts/check-skill-packs.sh` | `advisory`; `severity-labeled`; `CLI decides` | Planning defects의 실제 감소; user satisfaction |
| Change-control UX | Locked plan changes를 더 이해하고 audit하기 쉽다. | Amended docs; amendment record; status proof; relock behavior; stale lock 또는 hash mismatch examples. | Amendment/relock과 stale-lock proof를 보여주는 future docs, tests, examples | `Planned`; `audited`; `verified` | Team-scale merge success; production workflow adoption |
| Homebrew | Homebrew는 tested tap/formula evidence가 있을 때만 available이다. | Tap repository; formula; sha256; `brew install` output; `ni --help`; `ni version`. | `docs/80_HOMEBREW_DECISION.md`; future Homebrew verification doc | Current status: `Planned` / v0.5 candidate; evidence가 있을 때만 later `Experimental` 또는 `Available` | Package-manager adoption; long-term formula maintenance |
| Model workspace packs | Skill packs가 supported model workspaces에서 설치하거나 사용하기 더 쉽다. | Package scripts; zip contents; `SKILL.md` files; status-preservation doc; `Available` claim에는 host-level install verification 필요. | `bash scripts/check-skill-packs.sh`; `bash scripts/package-claude-skills.sh`; `bash scripts/package-codex-skills.sh`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `docs/99_MODEL_WORKSPACE_STATUS.ko.md` | `Experimental`; specific host path의 host-level install과 usage verification 뒤에만 `Available` | Model provider behavior; tested되지 않은 global host compatibility |
| No-terminal assisted workflow | Users는 CLI 없이 drafting을 시작할 수 있지만, trusted validation에는 여전히 CLI가 필요하다. | `docs/no-terminal.md`; `examples/no-terminal-assisted/**`; clear not-deterministic warning. | `docs/no-terminal.md`; `bash scripts/demo-check.sh` | `Experimental`; `assisted`; `not deterministic` | CLI 없는 trusted readiness; lock; hash verification; prompt compilation |
| Product surface expansion | NI는 software planning만 지원하는 것이 아니다. | Research protocol, operations process, education program, document product, physical product planning 또는 유사 examples; status proof 또는 docs-only boundary. | `examples/**`; `docs/82_EXAMPLE_COVERAGE.md` | `example-backed`; `docs-only`; `measured` | Real-world adoption; downstream implementation quality |
| Downstream integrations | Downstream integrations는 separate packages, seed exports, target formats로 남는다. | Separate package 또는 docs; no `ni-kernel` runtime execution; target/export conformance. | `docs/45_TARGET_STORY.md`; target conformance docs/tests when present; `bash scripts/smoke.sh` | `seed`; `export`; `separate package`; `not ni-kernel runtime` | Downstream execution success |

## Claim boundary rules

- Implementation을 측정하지 않았다면 "improves implementation quality"라고 말하지 않는다.
- Repeated trials를 측정하지 않았다면 "reduces rework"라고 말하지 않는다.
- Install 또는 verification evidence가 없으면 `Available`이라고 말하지 않는다.
- Readiness scope가 그렇게 말하지 않는 한 `READY`가 product ready를 뜻한다고 말하지 않는다.
- Host-level install이 verified되지 않았다면 model workspace packs가 global이라고 말하지 않는다.
- No-terminal을 deterministic이라고 말하지 않는다.

## GRILL-003 closure

GRILL-003은 이 acceptance evidence matrix로 addressed되었다. 이것은 모든 v0.5
lane이 complete라는 뜻이 아니다. Future tasks가 completion을 claim하기 전에
충족해야 하는 lane별 completion evidence requirements가 더 명확해졌다는 뜻이다.

GRILL-004는 `docs/97_BENCHMARK_CLAIM_BOUNDARIES.ko.md`로 addressed되었다. 이
문서는 benchmark claim limit을 case-study와 example surface 옆에서 보이게
유지한다. GRILL-005는 `docs/99_MODEL_WORKSPACE_STATUS.ko.md`로 addressed되었다.
이 문서는 broad product path로서의 Experimental status를 보존하고, verified
repository evidence를 명명하며, host-level/global install과 provider behavior는
documented되기 전까지 unverified로 둔다.

## How to use this document

v0.5 task를 시작하기 전에:

1. Evidence matrix에서 task lane을 식별한다.
2. Editing 전에 required evidence를 확인한다.
3. Claim 옆에 `not_measured` boundaries를 보이게 유지한다.
4. Completion을 claim하기 전에 examples, docs, tests를 업데이트한다.
5. Evidence가 충분한지 `ni-grill`로 challenge한다.
6. Deterministic readiness에는 `ni status`를 사용한다. Model judgment로 대체하지 않는다.
