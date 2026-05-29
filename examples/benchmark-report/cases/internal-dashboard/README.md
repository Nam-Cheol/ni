# Internal Dashboard Benchmark Case

This is a manual, qualitative benchmark case for the vague customer-dashboard
request in `testdata/benchmark/vague-requests/customer-dashboard/`.

It now includes an isolated ni planning workspace:

```text
examples/benchmark-report/cases/internal-dashboard/workspace/
```

Measured ni-path result:

- Command: `go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions`
- Readiness: `BLOCKED`
- Lock created: no
- Bounded prompt compiled: no
- Prompt character count: `not_measured`

Because readiness is `BLOCKED`, this case does not run `ni end` or `ni run`.
It does not implement a dashboard, execute a downstream agent, call a model API,
or run dashboard build commands.

Added blocker evidence:

- `08-blocker-analysis.md` explains why `OQ-001` through `OQ-004` block lock,
  what kind of user answer would resolve each blocker, and which unsafe
  assumptions ni avoided.
- `09-resolution-path.md` defines how a future resolved variant could update
  docs, contract records, risks, and evaluations before `ni status`, `ni end`,
  and `ni run` become eligible.

The benchmark evidence is qualitative and auditable. It does not claim
statistical significance, adoption metrics, rework reduction, or downstream
implementation quality.
