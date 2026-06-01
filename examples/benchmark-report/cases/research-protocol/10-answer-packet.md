# Research protocol benchmark answer packet

This packet is a human-fillable form for the research-protocol benchmark
blockers. It does not fill in answers, resolve blockers, force readiness,
create a lock, compile a prompt, run research work, collect data, or authorize
fieldwork.

Use answers only inside:

```text
examples/benchmark-report/cases/research-protocol/workspace/
```

Do not apply them to the repository-root planning lock or root `.ni/` state.

## Status

- Current readiness: `BLOCKED`
- Lock created: no
- `ni-run` prompt compiled: no
- Required answers: `OQ-001` through `OQ-005`

## How to use this packet

1. Fill in the answers below.
2. Use the answers to update only the isolated research-protocol benchmark
   workspace.
3. Run `ni status --proof --next-questions` against that workspace.
4. Lock only if readiness becomes `READY` or `READY_WITH_DEFERRALS` and user
   confirmation exists.
5. Run `ni run` only after a valid benchmark workspace lock.

## OQ-001 - Research question, supported decision, and final study artifact

Prompt:

What research question should this study answer, what decision should it
support, and what final artifact is enough for planning?

Required answer fields:

- Research question:
- Supported decision:
- Decision owner:
- Final study artifact:
- What should not be decided by this study:

Unsafe assumptions to avoid:

- Assuming the study chooses intervention locations directly.
- Assuming the final artifact is a full statistical report.
- Assuming the research question is about all heat risk rather than a specific
  planning decision.

## OQ-002 - Participant/observation scope, criteria, locations, and windows

Prompt:

Who or what is observed, where, during what time window, and under what
inclusion/exclusion criteria?

Required answer fields:

- Observation units:
- Participant groups, if any:
- Inclusion criteria:
- Exclusion criteria:
- Locations or neighborhood types:
- Observation window:
- Sampling or selection rule:
- Out-of-scope populations or locations:

Unsafe assumptions to avoid:

- Recruiting vulnerable participants without criteria.
- Treating all neighborhoods as interchangeable.
- Assuming any hot day or any location is valid.

## OQ-003 - Consent, privacy, data handling, translation, accessibility, and sensitive-data boundaries

Prompt:

What data may be collected or used, how is consent handled, what data is
sensitive, and what accessibility or translation support is required?

Required answer fields:

- Data types allowed:
- Data types prohibited:
- Consent approach:
- Privacy constraints:
- Storage/retention rule:
- Translation needs:
- Accessibility needs:
- Sensitive-data boundaries:
- Data that must remain out of scope:

Unsafe assumptions to avoid:

- Collecting identifiable data unnecessarily.
- Treating public observation as consent-free in all cases.
- Ignoring translation or accessibility needs.
- Storing sensitive notes without retention boundaries.

## OQ-004 - Field safety, heat/weather stop rules, vulnerable-group safeguards, and escalation path

Prompt:

What safety rules protect field teams and participants during hot weather or
unsafe conditions?

Required answer fields:

- Field team safety rules:
- Heat/weather stop conditions:
- Maximum exposure or shift limits:
- Hydration/rest requirements:
- Vulnerable-group safeguards:
- Emergency escalation path:
- Conditions that cancel fieldwork:

Unsafe assumptions to avoid:

- Sending field teams out during dangerous heat.
- Continuing observations despite unsafe conditions.
- Treating vulnerable groups as ordinary observation targets without
  safeguards.

## OQ-005 - Reviewer, acceptance evidence, and pre-fieldwork readiness criteria

Prompt:

Who reviews this protocol, and what evidence is enough before any fieldwork
begins?

Required answer fields:

- Reviewer or approval owner:
- Review audience:
- Minimum protocol artifact:
- Acceptance evidence:
- Pass/fail criteria:
- Pre-fieldwork readiness checklist:
- What is explicitly not required:
- Who can approve moving forward:

Unsafe assumptions to avoid:

- Treating a draft protocol as fieldwork-ready.
- Treating informal agreement as approval.
- Starting data collection before acceptance evidence exists.

## After answers are provided

Expected next steps:

1. Update only the isolated benchmark workspace.
2. Update `docs/plan`.
3. Update `.ni/contract.json`.
4. Update `.ni/session.json` if present.
5. Run `ni status --proof --next-questions`.
6. If `BLOCKED`, document remaining blockers.
7. If `READY` or `READY_WITH_DEFERRALS`, lock only inside the benchmark
   workspace.
8. If locked, compile a bounded prompt from the benchmark workspace.
9. Update measurement table honestly.

Rules:

- Do not edit root `.ni/plan.lock.json`.
- Do not run repository-root `ni end` or `ni relock`.
- Do not run `ni end` or `ni run` in the research-protocol benchmark workspace
  while it remains `BLOCKED`.
- Do not execute downstream agents.
- Do not run fieldwork.
- Do not recruit participants.
- Do not collect data.
- Do not call model APIs.
- Do not fake prompt or lock evidence.
