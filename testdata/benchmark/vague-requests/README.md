# Vague Requests Benchmark Corpus

This directory contains seed requests for the manual intent readiness benchmark.
Each request is intentionally plausible but underspecified.

Use these files as inputs to the protocol in
[`docs/43_BENCHMARK_PROTOCOL.md`](../../../docs/43_BENCHMARK_PROTOCOL.md).
Do not execute downstream agents from these prompts. Score them as planning
inputs only.

## Fixture Format

Each fixture is a directory with four files:

- `request.md`: the direct-to-agent prompt text and category label.
- `expected-hidden-assumptions.md`: reviewer seed notes for assumptions a
  downstream actor would need to invent before execution.
- `expected-readiness-gaps.md`: reviewer seed notes aligned to the benchmark
  metrics.
- `suggested-ni-questions.md`: clarification questions a ni-start conversation
  could ask before readiness.

These expectation files are not empirical results. They make the manual review
surface concrete and auditable.

## Fixtures

- [`customer-dashboard/`](customer-dashboard/): software dashboard with unclear
  users, data, privacy, and success criteria.
- [`conversation-memory-companion/`](conversation-memory-companion/):
  conversation product with unclear memory, consent, and autonomy boundaries.
- [`community-heat-field-study/`](community-heat-field-study/): research
  protocol with unclear participants, safety, and outputs.
- [`support-handoff-process/`](support-handoff-process/): internal operations
  process with unclear issue scope, ownership, and escalation rules.
- [`manager-training-program/`](manager-training-program/): education program
  with unclear learning outcomes, policy review, and measurement.
- [`board-update-brief/`](board-update-brief/): document product with unclear
  evidence, audience, review path, and claims.
- [`partner-sync-cli/`](partner-sync-cli/): API or CLI tool with unclear sync
  semantics, safety controls, and data contracts.
- [`desk-lamp-prototype/`](desk-lamp-prototype/): physical product planning
  with unclear safety, prototype fidelity, and customer-test criteria.
- [`namba-ai-upgrade/`](namba-ai-upgrade/): Namba AI upgrade style project with
  unclear memory, interaction quality, and evaluation boundaries.
- [`weekly-status-automation/`](weekly-status-automation/): ambiguous automation
  request with unclear source access, review flow, and distribution rules.
