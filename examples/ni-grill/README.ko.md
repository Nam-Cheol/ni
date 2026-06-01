# ni-grill Planning Challenge

이 docs-only example은 `ni-end` 전에 `ni-grill`이 draft NI plan을 어떻게
challenge하는지 보여준다.

Expected `ni status`: not claimed. 이 example은 transcript fixture이며 trusted
CLI workspace가 아니다.

## 무엇을 보여주는가

- `ni-grill`은 `docs/plan/**`, `.ni/contract.json`, `ni status --proof --next-questions`를 읽는다.
- Deterministic blockers가 있으면 새로운 critique 전에 그 blocker를 먼저 사용한다.
- Vague decisions, acceptance evidence, risk, non-goals, handoff,
  docs/contract sync에 대해 focused `GRILL-*` questions를 묻는다.
- Findings에 `Critical`, `High`, `Medium`, `Low`, `Note`를 붙이고 기본적으로
  최대 5 findings만 보여준다.
- `ni status`가 `READY`라고 말한 benchmark case에서도 claim boundaries,
  `not_measured` sections, handoff wording을 pressure-test할 수 있다.
- Downstream work를 execute하지 않고, product를 implement하지 않고, model
  judgment로 lock을 approve하지 않는다.
- User answers는 `ni-start`와 같은 discipline으로 docs, contract, session에
  함께 기록하고 status proof를 다시 확인한다.

## Files

- `01-draft-plan.md`: weak accepted content가 있는 draft plan excerpt.
- `02-grill-questions.md`: status를 읽고 grill pressure를 적용한 model output.
- `03-user-answers.md`: focused grill questions에 대한 user answers.
- `04-docs-contract-delta.md`: answers 뒤 planning updates와 proof shape.
- `05-status-after-grill.md`: grill update 뒤 status proof summary.
- `06-internal-dashboard-grill.md`: isolated `READY` proof 뒤
  `internal-dashboard` case에 적용한 benchmark grill.
- `07-research-protocol-grill.md`: isolated `READY` proof 뒤
  `research-protocol` case에 적용한 benchmark grill.
- `08-grill-lessons.md`: benchmark evidence에 `ni-grill`을 dogfood하며 얻은
  lessons.
- `09-ni-project-grill.md`: ni 자체의 current planning state에 `ni-grill`을
  적용한 read-only dogfood report.

## Boundary

`ni-grill` challenges planning quality before lock. It does not execute work.
Skills are UX; CLI is authority. Benchmark grill은 claim-boundary review이며
new empirical claims를 만들거나, generated prompts를 실행하거나, model APIs를
호출하거나, products를 implement하거나, fieldwork를 수행하거나, `ni-grill`을
CLI보다 authoritative하게 만들지 않는다.

Default grill output은 budgeted이다. 최대 5 findings, 그리고 `Critical`/`High`
findings가 있으면 최대 3개까지만 먼저 보여준다. Lower-priority findings는 전부
나열하지 않고 summarize해야 한다.

Documentation review와 `bash scripts/quality.sh`가 이 example을 확인한다.
