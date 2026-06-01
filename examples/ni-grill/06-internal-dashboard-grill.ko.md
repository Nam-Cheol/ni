# 06. Internal Dashboard Benchmark Grill

이 transcript는 checked-in `internal-dashboard` benchmark case에 `ni-grill`을
dogfood한 기록이다.

Target workspace:

```text
examples/benchmark-report/cases/internal-dashboard/workspace/
```

Isolated benchmark workspace에 대해 실행한 command:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
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
않고, generated prompt를 실행하지 않고, dashboard product-readiness claim을
만들지 않는다.

## Grill findings

1. GRILL-001 — Critical — claim boundary
   Affected: `RISK-004` /
   `examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.md`
   Concern: Evidence file은 `READY`가 benchmark planning-meeting artifact
   readiness에만 적용된다고 말하지만, transition row의 "Safe to hand off as
   benchmark planning artifact" 문구는 같은 row에서 금지된 downstream
   interpretation을 직접 이름 붙이지 않는다.
   Why it matters: downstream reader가 뒤의 `not_measured` section 없이
   transition table만 인용하면 dashboard product work가 ready인 것처럼 들릴 수
   있다.
   Question: Transition table에 dashboard product readiness, dashboard
   implementation quality, downstream agent performance가 여전히
   `not_measured`라는 inline note를 넣어야 하는가?
   Answer shape: yes/no와 exact row wording, 또는 adjacent scope note가
   충분하다는 이유.
   Suggested action: clarify
   Blocks ni-end: maybe

2. GRILL-002 — High — acceptance evidence
   Affected: `EVAL-002` / `OQ-004` /
   `examples/benchmark-report/cases/internal-dashboard/workspace/docs/plan/07_evaluation_contract.md`
   Concern: Accepted evidence는 모든 required answer field와 testable
   pass/fail criteria를 요구하지만, specific planning meeting date는 unassigned
   상태다.
   Why it matters: benchmark artifact에 대해 `READY`는 유효할 수 있지만,
   future planning handoff에서는 "next scheduled planning meeting"이 placeholder인지
   intentionally deferred operational detail인지 알아야 할 수 있다.
   Question: Benchmark가 meeting date를 artifact readiness에는 non-blocking인
   unassigned detail로 명시해야 하는가?
   Answer shape: accepted non-blocking note, deferred operational follow-up, 또는
   planning owner role만으로 충분하다는 설명.
   Suggested action: clarify
   Blocks ni-end: maybe

3. GRILL-003 — High — risk and non-goal clarity
   Affected: `RISK-002` / `NG-002` /
   `examples/benchmark-report/cases/internal-dashboard/workspace/docs/plan/05_constraints.md`
   Concern: Privacy boundary는 private customer data와 sensitive source data를
   금지하지만, `OQ-003`은 "approved internal dashboard source material required
   to validate the benchmark case"도 허용한다.
   Why it matters: 방어 가능한 boundary지만, reader는 누가 source material을
   approve하는지 또는 evidence file에 어느 정도까지 복사할 수 있는지 모를 수
   있다.
   Question: Case가 internal dashboard source material의 approval role과 최대
   allowed source excerpt shape를 이름 붙여야 하는가?
   Answer shape: approval role과 reference-only, summary-only, short excerpt rule
   중 하나.
   Suggested action: clarify
   Blocks ni-end: maybe

4. GRILL-004 — Medium — handoff boundary
   Affected: `CAP-004` / `ART-007` /
   `examples/benchmark-report/cases/internal-dashboard/14-bounded-prompt-summary.md`
   Concern: Bounded prompt summary는 4000-character prompt가 compile되었음을
   확인하지만, prompt text 자체가 truncation 이후 product-readiness overclaim을
   피했는지는 말하지 않는다.
   Why it matters: Truncated prompt는 lock boundary를 보존하면서도 중요한 warning
   text를 잘라내거나 ambiguous closing instruction을 남길 수 있다.
   Question: Generated 4000-character prompt에 대해 manual prompt-boundary review
   note를 추가해야 하는가?
   Answer shape: reviewer role과 no execution instruction, no dashboard build
   claim, no product-readiness claim, preserved source-of-truth warning checklist.
   Suggested action: clarify
   Blocks ni-end: no

5. GRILL-005 — Low — docs/contract sync claim
   Affected: `.ni/contract.json` / `DEC-004` /
   `examples/benchmark-report/cases/internal-dashboard/workspace/docs/plan/11_decision_log.md`
   Concern: `DEC-004`는 named person이 unassigned여도 approval role을 explicit하게
   유지하지만, contract와 docs는 concrete owner가 아니라 role에 의존한다.
   Why it matters: Role-based acceptance는 benchmark에서는 acceptable하지만,
   downstream planning meeting은 unnamed owner를 real approval로 취급하지 않아야
   한다.
   Question: Approval이 person-completed가 아니라 role-defined라는 점을 설명할 때
   evidence summary에서 `DEC-004`를 cite해야 하는가?
   Answer shape: evidence summary에 `DEC-004` cite, non-blocking follow-up 추가,
   또는 현 상태 유지 rationale.
   Suggested action: keep as note
   Blocks ni-end: no

## Grill result

`ni-grill`은 CLI-readiness blocker를 찾지 못했다. 가장 강한 hardening question은
`GRILL-003`이다. "approved internal dashboard source material"이라는 문구는 더
명확한 approval role과 excerpt rule을 얻을 수 있다. 나머지 findings는 주로
`READY` 이후 dashboard product readiness와 bounded prompt interpretation 주변의
claim boundary를 보호한다.
