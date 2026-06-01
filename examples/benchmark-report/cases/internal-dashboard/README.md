# Internal Dashboard Benchmark Case

This is a manual, qualitative benchmark case for the vague customer-dashboard
request in `testdata/benchmark/vague-requests/customer-dashboard/`.

It now includes an isolated ni planning workspace:

```text
examples/benchmark-report/cases/internal-dashboard/workspace/
```

Measured ni-path result after applying the user-provided answers for `OQ-001`
through `OQ-004`:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions`
- Readiness: `READY`
- Scope of readiness: benchmark planning-meeting artifact readiness only
- Lock created: yes, inside `workspace/.ni/plan.lock.json`
- Bounded prompt compiled: yes
- Prompt character count: `4000`

The provided answers resolve the benchmark artifact blockers, not the real
dashboard product requirements. This case does not implement a dashboard,
execute a downstream agent, call a model API, or run dashboard build commands.

Added blocker evidence:

- `08-blocker-analysis.md` explains why `OQ-001` through `OQ-004` block lock,
  what kind of user answer would resolve each blocker, and which unsafe
  assumptions ni avoided.
- `09-resolution-path.md` defines how a future resolved variant could update
  docs, contract records, risks, and evaluations before `ni status`, `ni end`,
  and `ni run` become eligible.
- `10-answer-packet.md` and `10-answer-packet.ko.md` provide fillable
  user-facing answer packets for `OQ-001` through `OQ-004` without resolving
  the blockers or changing benchmark readiness.
- `11-resolved-status-proof.md` records the measured `READY` status after the
  answers were applied to the isolated workspace.
- `12-resolved-next-questions.md` records that no next-question groups remained
  after the resolved status run.
- `13-lock-summary.md` records the isolated workspace lock.
- `14-bounded-prompt-summary.md` records the bounded prompt command and 4000
  character count.

The benchmark evidence is qualitative and auditable. It does not claim
statistical significance, adoption metrics, rework reduction, or downstream
implementation quality.
