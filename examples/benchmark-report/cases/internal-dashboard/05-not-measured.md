# What Was Not Measured

This dashboard case did not measure:

- actual `ni end` output;
- actual `ni run` output or prompt character count;
- dashboard implementation quality;
- downstream agent performance;
- user satisfaction;
- adoption;
- runtime safety;
- reduced rework;
- cost, latency, or statistical effect size.

No downstream agent was executed. No model API was called. No dashboard was
implemented. No queue, shell adapter, telemetry path, or runtime harness was
added.

The case now includes actual `ni status --proof --next-questions` output in
`06-ni-status-proof.md`. Because that output is `BLOCKED`, lock creation,
prompt compilation, and prompt character count remain `not_measured`.
