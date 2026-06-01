# Research Protocol Benchmark Case

이 문서는 두 번째 v0.5 real benchmark evidence case다. Downstream work가 시작되기
전, vague non-software research-protocol request를 측정하고, isolated workspace에
resolved synthetic-fixture variant를 기록한다.

Draft vague request:

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

Initial ni-path result:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions`
- Readiness: `BLOCKED`
- Blockers: `OQ-001`부터 `OQ-005`
- Evidence: `06-ni-status-proof.md`, `07-ni-next-questions.md`

Synthetic benchmark fixture answer를 적용한 뒤 resolved ni-path result:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions`
- Readiness: `READY`
- Blockers: none
- Lock created: yes, `workspace/.ni/plan.lock.json` 안에서만
- Bounded prompt compiled: yes
- Prompt character count: `4000`

Synthetic answer는 benchmark fixture data일 뿐이다. Real fieldwork approval,
actual research authorization, proof of research quality, empirical evidence가
아니다.

## Files

- `01-original-request.md`: vague research-protocol request.
- `02-direct-to-agent-risk.md`: downstream actor가 발명할 수 있는 assumption.
- `03-ni-path.md`: expected ni path와 stop rule.
- `04-measurement-table.md`: before/after qualitative measurement table.
- `05-not-measured.md`: explicit non-execution 및 not_measured boundary.
- `06-ni-status-proof.md`: initial `ni status --proof --next-questions` output.
- `07-ni-next-questions.md`: initial next-question group과 blocker name.
- `08-blocker-analysis.md`: 각 blocker가 왜 중요한지, 나중에 어떤 answer가
  resolve할 수 있는지, 어떤 unsafe assumption을 피해야 하는지 설명한다.
- `08-blocker-analysis.ko.md`: blocker analysis의 Korean companion.
- `09-resolution-path.md`: answer를 발명하지 않고 `BLOCKED`에서 가능한 `READY`
  또는 `READY_WITH_DEFERRALS`로 가는 step-by-step path.
- `09-resolution-path.ko.md`: resolution path의 Korean companion.
- `10-answer-packet.md`: `OQ-001`부터 `OQ-005`까지의 answer packet.
- `10-answer-packet.ko.md`: answer packet의 Korean companion.
- `11-resolved-status-proof.md`: synthetic fixture answer 적용 뒤 측정한
  `READY` proof.
- `12-resolved-next-questions.md`: 남은 next question이 없음을 기록한다.
- `13-lock-summary.md`: isolated workspace lock evidence.
- `14-bounded-prompt-summary.md`: bounded prompt command와 4000 character count.
- `workspace/`: isolated ni workspace for this benchmark case.

## Non-execution boundary

이 benchmark case는 research protocol quality, fieldwork readiness,
intervention decision readiness, research outcome validity, adoption, cost,
latency, rework reduction, statistical effect size를 주장하지 않는다.

Resolved lock과 prompt는 seed evidence일 뿐이다. Downstream agents, model APIs,
generated prompts, fieldwork, participant recruitment, data collection,
dashboard 또는 research implementation, issue publishing, PR automation, shell
adapter, queue를 실행하지 않는다.
