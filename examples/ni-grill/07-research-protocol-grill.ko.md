# 07. Research Protocol Benchmark Grill

이 transcript는 checked-in `research-protocol` benchmark case에 `ni-grill`을
dogfood한 기록이다.

Target workspace:

```text
examples/benchmark-report/cases/research-protocol/workspace/
```

Isolated benchmark workspace에 대해 실행한 command:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions
```

Observed status:

```text
NI Intent Readiness: READY

Blockers:
- None.

Deferrals:
- None.

Warnings:
- None.
```

Deterministic blocker나 deferral이 없었으므로 아래 findings는 pre-handoff
hardening questions이다. 이것들은 `ni status`를 대체하지 않고, lock을 approve하지
않고, generated prompt를 실행하지 않고, fieldwork를 수행하지 않고, real research
approval 또는 research-quality claim을 만들지 않는다.

## Grill findings

1. GRILL-001 — Critical — claim boundary
   Affected: `RISK-004` /
   `examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md`
   Concern: Case는 `READY`가 synthetic benchmark protocol planning artifact
   readiness에만 해당한다고 강하게 말하지만, "Safe to hand off as benchmark
   planning artifact" 문구는 adjacent non-approval context 없이 재사용될 수 있다.
   Why it matters: Research-protocol handoff language는 context 밖에서 인용되면
   fieldwork authorization 또는 ethics approval로 오해될 수 있다.
   Question: Readiness transition table에 real research approval, fieldwork
   authorization, research quality, intervention effectiveness가 여전히
   `not_measured`라는 inline wording을 넣어야 하는가?
   Answer shape: yes/no와 exact table wording, 또는 existing critical scope note가
   충분하다는 이유.
   Suggested action: clarify
   Blocks ni-end: maybe

2. GRILL-002 — High — acceptance evidence
   Affected: `EVAL-002` / `OQ-005` /
   `examples/benchmark-report/cases/research-protocol/workspace/docs/plan/07_evaluation_contract.md`
   Concern: Acceptance evidence는 planning owner와 privacy/safety reviewer를
   요구하지만, fixture는 synthetic review ownership과 real institutional approval을
   구분하지 않는다.
   Why it matters: Benchmark는 `READY`일 수 있지만, human reader에게 reviewer
   names가 real governance approval이 아니라 fixture roles라는 더 강한 cue가
   필요할 수 있다.
   Question: Protocol에 reviewer roles가 synthetic benchmark roles이며 IRB,
   ethics board, legal, city approval이 아니라는 fixture-reviewer note를 추가해야
   하는가?
   Answer shape: explicit non-approval note, reviewer role clarification, 또는
   existing `not_measured` sections가 충분하다는 rationale.
   Suggested action: clarify
   Blocks ni-end: maybe

3. GRILL-003 — High — risk and non-goal clarity
   Affected: `RISK-003` / `NG-001` /
   `examples/benchmark-report/cases/research-protocol/workspace/docs/plan/05_constraints.md`
   Concern: Field safety rules는 benchmark planning에는 충분히 구체적이지만,
   실제 현장에서 따를 수 있는 operational rules처럼 읽힐 수 있다.
   Why it matters: Specific safety rules는 ambiguity를 줄이지만, benchmark
   artifact에서는 field team deployment authorization이 아니라 planning
   constraints로 남아야 한다.
   Question: Safety section에 이 rules가 future protocol review의 checklist
   criteria이며 이 benchmark로 fieldwork를 시작하면 안 된다는 문장을 추가해야
   하는가?
   Answer shape: constraints 또는 delivery operation에 추가할 one sentence, 또는
   existing non-goals가 이미 충분하다는 이유.
   Suggested action: clarify
   Blocks ni-end: maybe

4. GRILL-004 — Medium — synthetic fixture label
   Affected: `DEC-004` / `OQ-001` through `OQ-005` /
   `examples/benchmark-report/cases/research-protocol/10-answer-packet.md`
   Concern: Answer packet은 상단에서 synthetic fixture answers라고 명확히
   label하지만, individual OQ sections에는 header 없이 복사될 수 있는 realistic
   operational details가 있다.
   Why it matters: 복사된 section이 synthetic-fixture boundary를 잃으면 real
   approved study instructions처럼 보일 수 있다.
   Question: 각 `OQ-*` section이 required answer block 또는 section lead에서
   "Synthetic benchmark fixture answer"를 반복해야 하는가?
   Answer shape: section마다 label 반복, copy-safety note 추가, 또는 top-level
   label만으로 충분하다는 설명.
   Suggested action: clarify
   Blocks ni-end: no

5. GRILL-005 — Medium — prompt boundary
   Affected: `CAP-004` / `ART-008` /
   `examples/benchmark-report/cases/research-protocol/14-bounded-prompt-summary.md`
   Concern: Bounded prompt summary는 character count와 non-execution을 증명하지만,
   truncation이 research non-approval warning을 보존했는지는 증명하지 않는다.
   Why it matters: 4000-character seed prompt는 kernel output으로 valid할 수
   있지만, external use 전에 overclaim을 피하기 위한 human review가 여전히 필요할
   수 있다.
   Question: Case가 no fieldwork authorization, no participant recruitment, no
   research approval, no model API call, preserved source-of-truth warning에 대한
   prompt-boundary review checklist를 기록해야 하는가?
   Answer shape: reviewer role과 checklist result, 또는 이 benchmark에서는 prompt
   text를 boundedness 이상으로 review하지 않는다는 explicit decision.
   Suggested action: clarify
   Blocks ni-end: no

## Grill result

`ni-grill`은 CLI-readiness blocker를 찾지 못했다. 가장 강한 hardening question은
`GRILL-003`이다. Detailed safety rules는 유용하지만 fieldwork authorization으로
혼동되면 안 된다. 나머지 findings는 `READY` 이후 synthetic fixture labeling,
acceptance evidence, `not_measured` claim boundary를 강화한다.
