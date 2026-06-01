# Constraints

## Hard constraints

- Readiness must be rule-based, not model-feeling-based.
- OQ-001 through OQ-005 are resolved only by the clearly labeled synthetic
  benchmark fixture answers.
- Do not run `ni end` or `ni run` while `ni status` reports `BLOCKED`.
- Do not execute fieldwork, participant recruitment, data collection, analysis,
  intervention placement, downstream agents, generated prompts, model APIs,
  shell adapters, queues, telemetry paths, release automation, or runtime
  harnesses.
- Do not claim research protocol quality, fieldwork readiness, intervention
  decision readiness, participant outcomes, adoption, rework reduction, cost,
  latency, or statistical effect size.
- Prompt output from `ni run` must be 4000 characters or less if a later task
  reaches a valid lock.
- Do not collect personal, health, financial, immigration, employment, or
  household-level information.
- Do not record vulnerable-group status.
- Do not collect identifiers, photos of identifiable people, license plates,
  raw audio/video, precise resident addresses, credentials, tokens, or private
  customer data.
- Observation windows are weekdays between 09:00 and 16:00 during the summer
  study period, only when field safety conditions permit.
- Fieldwork must be canceled or paused if heat index exceeds 38 C, an official
  heat emergency is active, lightning occurs within 10 miles, air quality is
  unhealthy, or the field lead judges conditions unsafe.

## Non-goals

- NG-001: no fieldwork, participant recruitment, data collection, analysis, or
  intervention placement from this benchmark.
- NG-002: no ethics approval, research protocol quality, fieldwork readiness,
  intervention decision readiness, or real-world outcome claim.
- NG-003: no downstream agent, generated prompt execution, model API, shell
  adapter, queue, telemetry path, release automation, or runtime harness.
- NG-004: no adoption, rework, cost, latency, statistical effect, or downstream
  agent quality claim.
