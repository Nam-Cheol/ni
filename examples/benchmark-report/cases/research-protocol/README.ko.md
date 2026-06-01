# Research Protocol Benchmark Case

이 문서는 두 번째 v0.5 real benchmark evidence case다. Downstream work가 시작되기
전, vague non-software research-protocol request의 initial readiness를 측정한다.

Draft vague request:

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

Isolated workspace에서 측정한 ni-path result:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions`
- Readiness: `BLOCKED`
- Blockers: `OQ-001`부터 `OQ-005`
- Lock created: no
- Bounded prompt compiled: no
- Prompt character count: `not_measured`

이 case는 ni가 unclear research intent를 visible하게 만들고 fieldwork,
participant recruitment, data collection, analysis, intervention placement,
downstream agents, model APIs, generated prompt execution 전에 handoff를 멈출 수
있음을 보여준다.

## Files

- `01-original-request.md`: vague research-protocol request.
- `02-direct-to-agent-risk.md`: downstream actor가 발명할 수 있는 assumption.
- `03-ni-path.md`: expected ni path와 stop rule.
- `04-measurement-table.md`: initial-state qualitative measurement table.
- `05-not-measured.md`: explicit non-execution 및 not_measured boundary.
- `06-ni-status-proof.md`: exact `ni status --proof --next-questions` output.
- `07-ni-next-questions.md`: next-question group과 blocker name.
- `workspace/`: isolated ni workspace for this benchmark case.

## Non-execution boundary

이 benchmark case는 research protocol quality, fieldwork readiness,
intervention decision readiness, research outcome validity, adoption, cost,
latency, rework reduction, statistical effect size를 주장하지 않는다.

Status가 `BLOCKED`인 동안 workspace에는 lock이 없다. Later task가 blocker를
명시적으로 answer 또는 defer하고 status를 다시 실행하기 전까지 이 case에서
`ni end`나 `ni run`을 실행하지 않는다.
