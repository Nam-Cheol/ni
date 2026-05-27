# Vague Requests Benchmark Corpus

This directory contains seed requests for the manual intent readiness benchmark.
Each request is intentionally plausible but underspecified.

Use these files as inputs to the protocol in
[`docs/43_BENCHMARK_PROTOCOL.md`](../../../docs/43_BENCHMARK_PROTOCOL.md).
Do not execute downstream agents from these prompts. Score them as planning
inputs only.

## Files

- [`customer-dashboard.md`](customer-dashboard.md): an internal dashboard
  request with unclear users, data, and success criteria.
- [`onboarding-refresh.md`](onboarding-refresh.md): a lifecycle improvement
  request with unclear audience, scope, and measurement.
- [`field-study.md`](field-study.md): a research protocol request with unclear
  participants, safety constraints, and outputs.
