# Research Protocol Benchmark Case

This is the second v0.5 real benchmark evidence case. It measures a vague
non-software research-protocol request before downstream work begins, then
records a resolved synthetic-fixture variant inside the isolated workspace.

Draft vague request:

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

Initial ni-path result:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions`
- Readiness: `BLOCKED`
- Blockers: `OQ-001` through `OQ-005`
- Evidence: `06-ni-status-proof.md`, `07-ni-next-questions.md`

Resolved ni-path result after applying synthetic benchmark fixture answers:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions`
- Readiness: `READY`
- Blockers: none
- Lock created: yes, inside `workspace/.ni/plan.lock.json`
- Bounded prompt compiled: yes
- Prompt character count: `4000`

The synthetic answers are benchmark fixture data only. They are not real
fieldwork approval, not actual research authorization, not proof of research
quality, and not empirical evidence.

## Claim boundary

| Measured | Not measured |
| --- | --- |
| `BLOCKED` before synthetic answers; `READY` after synthetic protocol artifact answers; isolated lock; 4000-character prompt | real research approval; fieldwork authorization; research quality; intervention effectiveness; downstream agent performance; rework reduction; adoption; cost; latency; statistical effect |

`READY` means synthetic benchmark protocol planning artifact readiness only. It
keeps explicit no downstream execution, no implementation claim, and no
statistical claim boundaries.

## Files

- `01-original-request.md`: the vague research-protocol request.
- `02-direct-to-agent-risk.md`: assumptions a downstream actor might invent.
- `03-ni-path.md`: expected ni path and stop rule.
- `04-measurement-table.md`: before/after qualitative measurement table.
- `05-not-measured.md`: explicit non-execution and not_measured boundary.
- `06-ni-status-proof.md`: exact initial `ni status --proof --next-questions`
  output.
- `07-ni-next-questions.md`: initial next-question group and blocker names.
- `08-blocker-analysis.md`: why each blocker matters, what answer would
  resolve it later, and what would be unsafe to assume.
- `08-blocker-analysis.ko.md`: Korean companion for the blocker analysis.
- `09-resolution-path.md`: step-by-step path from `BLOCKED` to possible
  `READY` or `READY_WITH_DEFERRALS` without inventing answers.
- `09-resolution-path.ko.md`: Korean companion for the resolution path.
- `10-answer-packet.md`: answer packet for `OQ-001` through `OQ-005`.
- `10-answer-packet.ko.md`: Korean companion for the answer packet.
- `11-resolved-status-proof.md`: measured `READY` proof after synthetic fixture
  answers were applied.
- `12-resolved-next-questions.md`: records that no next questions remained.
- `13-lock-summary.md`: isolated workspace lock evidence.
- `14-bounded-prompt-summary.md`: bounded prompt command and 4000 character
  count.
- `15-before-after-evidence.md`: before/after proof from vague request
  `BLOCKED` state to synthetic fixture `READY` lock and bounded prompt.
- `15-before-after-evidence.ko.md`: Korean companion for the before/after
  evidence package.
- `16-lessons.md`: lessons from applying ni to a non-software research-planning
  case.
- `16-lessons.ko.md`: Korean companion for the lessons.
- `workspace/`: isolated ni workspace for this benchmark case.

## Readiness Transition

| Stage | Readiness | Lock | ni-run prompt | Prompt count | Meaning |
| --- | --- | --- | --- | --- | --- |
| Vague research request | BLOCKED | No | No | not_measured | Not safe to hand off |
| Synthetic answered protocol artifact | READY | Yes, isolated workspace only | Yes | 4000 | Safe to hand off as benchmark planning artifact |

## Non-execution boundary

This benchmark case does not claim research protocol quality, fieldwork
readiness, intervention decision readiness, research outcome validity,
adoption, cost, latency, rework reduction, or statistical effect size.

The resolved lock and prompt are seed evidence only. They do not execute
downstream agents, model APIs, generated prompts, fieldwork, participant
recruitment, data collection, dashboard or research implementation, issue
publishing, PR automation, shell adapters, or queues.
