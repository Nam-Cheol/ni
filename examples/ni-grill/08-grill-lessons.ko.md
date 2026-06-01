# 08. Grill Lessons

Benchmark dogfood는 `ni status`가 `READY`라고 말한 뒤에도 `ni-grill`이 왜
유용한지 보여준다.

`READY`는 deterministic readiness gates가 통과했다는 뜻이다. 모든 claim이
자동으로 강해졌거나, 모든 table을 안전하게 인용할 수 있거나, 모든 prompt가
truncation 뒤에도 최고의 warning을 보존하거나, 모든 future reader가 planning
evidence와 downstream execution의 boundary를 이해한다는 뜻은 아니다.

`ni-grill`은 planning quality, claim boundary, handoff safety를 pressure-test한다:

- Advisory pressure에 `Critical`, `High`, `Medium`, `Low`, `Note`를 붙여
  사용자가 무엇을 먼저 볼지 알 수 있게 한다.
- `READY`가 더 큰 product 또는 research claim이 아니라 accepted artifact에
  scoped되어 있는지 확인한다.
- Acceptance evidence가 다음 human reviewer에게 충분히 구체적인지 묻는다.
- 주변 `not_measured` boundary 없이 인용될 수 있는 문구를 잡아낸다.
- Realistic benchmark fixture details가 real approval 또는 execution instructions로
  오해될 수 있는지 검토한다.
- CLI gate가 통과한 뒤에도 non-goals와 risks를 계속 visible하게 둔다.

Plan이 중요하거나, risky하거나, public-facing이면 `ni-end` 전에 `ni-grill`을
실행해야 한다. Benchmark case에서는 특히 유용하다. Benchmark evidence는 쉽게
overclaim할 수 있다. Lock, bounded prompt, `READY` status는 intent-readiness를
증명할 수 있지만 product quality, downstream agent performance, fieldwork
readiness, research approval, empirical effect를 증명하지 않는다.

Internal-dashboard grill은 `READY`가 benchmark planning-meeting artifact readiness에만
적용됨을 확인했다. 남은 유용한 질문은 inline claim boundary, source-material
approval, role-based acceptance, prompt-boundary review에 관한 것이었다.

Research-protocol grill은 `READY`가 synthetic benchmark protocol planning artifact
readiness에만 적용됨을 확인했다. 남은 유용한 질문은 real research approval
boundary, fixture reviewer roles, checklist criteria로서의 safety rules, repeated
synthetic labels, prompt-boundary review에 관한 것이었다.

Pattern은 단순하다. 먼저 `ni status --proof --next-questions`를 실행하고,
deterministic blocker가 있으면 그것을 존중한다. 그 다음 `ni-grill`로 accepted
또는 nearly accepted planning content를 challenge한다. Output은 budgeted로 유지한다.
기본적으로 최대 5 findings, `Critical`/`High` findings가 있으면 먼저 최대 3개,
그리고 생략한 lower-priority findings는 summary로 처리한다.

`ni-grill`은 new empirical claims를 만들거나, readiness gates를 약화하거나, lock을
approve하거나, generated prompts를 실행하거나, model APIs를 호출하거나, products를
implement해서는 안 된다.
