# Research Protocol Benchmark Case

This is the second v0.5 real benchmark evidence case. It measures the initial
readiness of a vague non-software research-protocol request before any
downstream work begins.

Draft vague request:

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

Measured ni-path result for the isolated workspace:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions`
- Readiness: `BLOCKED`
- Blockers: `OQ-001` through `OQ-005`
- Lock created: no
- Bounded prompt compiled: no
- Prompt character count: `not_measured`

This case proves only that ni can make unclear research intent visible and stop
handoff before fieldwork, participant recruitment, data collection, analysis,
intervention placement, downstream agents, model APIs, or generated prompt
execution.

## Files

- `01-original-request.md`: the vague research-protocol request.
- `02-direct-to-agent-risk.md`: assumptions a downstream actor might invent.
- `03-ni-path.md`: expected ni path and stop rule.
- `04-measurement-table.md`: initial-state qualitative measurement table.
- `05-not-measured.md`: explicit non-execution and not_measured boundary.
- `06-ni-status-proof.md`: exact `ni status --proof --next-questions` output.
- `07-ni-next-questions.md`: next-question group and blocker names.
- `08-blocker-analysis.md`: why each blocker matters, what answer would
  resolve it later, and what would be unsafe to assume.
- `08-blocker-analysis.ko.md`: Korean companion for the blocker analysis.
- `09-resolution-path.md`: step-by-step path from `BLOCKED` to possible
  `READY` or `READY_WITH_DEFERRALS` without inventing answers.
- `09-resolution-path.ko.md`: Korean companion for the resolution path.
- `10-answer-packet.md`: human-fillable answer packet for `OQ-001` through
  `OQ-005`.
- `10-answer-packet.ko.md`: Korean companion for the answer packet.
- `workspace/`: isolated ni workspace for this benchmark case.

## Non-execution boundary

This benchmark case does not claim research protocol quality, fieldwork
readiness, intervention decision readiness, research outcome validity,
adoption, cost, latency, rework reduction, or statistical effect size.

No lock exists for the workspace while status is `BLOCKED`. Do not run
`ni end` or `ni run` for this case until a later task explicitly answers or
defers the blockers and reruns status.
