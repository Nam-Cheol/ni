# Second benchmark case selection

## Current benchmark state

The first v0.5 real benchmark evidence case is the internal-dashboard case in
`examples/benchmark-report/cases/internal-dashboard/`.

- Vague request: `BLOCKED`
- Answered artifact-readiness: `READY`
- Lock: isolated workspace only
- Prompt: 4000 characters
- Dashboard product readiness: `not_measured`

The first case is valuable because it shows hidden product assumptions becoming
visible blockers before handoff. Its claim boundary is intentionally narrow:
`READY` means benchmark planning-meeting artifact readiness only, not dashboard
product readiness.

## Candidate comparison

| Candidate | Differentiation value | Benchmark clarity | Lockability | Evidence quality | Cost | Boundary safety | Recommendation |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Research protocol | Very high: clearly non-software, human research, document and field-workflow surface. | High: a short vague study request exposes obvious participant, method, consent, safety, and artifact gaps. | High: user answers can reasonably move from `BLOCKED` to `READY` for a protocol artifact while leaving fieldwork out of scope. | Very high: likely to produce meaningful blockers, high-severity risk mitigations, status proof, isolated lock, and bounded human-team handoff. | Low: existing fixture and locked example provide vocabulary, but the benchmark should use a fresh isolated workspace. | Very high: fieldwork, participant data collection, ethics approval, analysis, and policy claims can stay explicit non-goals. | Select |
| Conversation product | Medium-high: expands beyond web/CLI into conversation surface, but existing examples already cover conversation products. | Medium: "personal but not creepy" is clear enough, but memory policy can sprawl. | Medium: lockable if memory scope, consent, deletion, and failure behavior are answered; could also remain valuable as `BLOCKED`. | High: strong blockers around memory, consent, trust, staleness, and cross-project boundaries. | Medium: understandable, but safety and memory controls may require careful scoping. | Medium-high: safe if no chatbot/runtime is built, but easy for readers to infer implementation quality. | Defer |
| Operations process | High: proves workflow/process planning and human-service boundaries. | High: support handoff problems are easy to understand and score. | High: likely lockable with issue scope, owners, escalation rules, accepted tools, and evidence criteria. | High: good role, approval, severity, and evidence blockers. | Low-medium: small workspace, but process examples may feel adjacent to software support operations. | High: can avoid tooling, SLA, staffing, or automation claims. | Strong later case |
| Education program | High: proves curriculum planning and assessment boundaries. | Medium-high: simple request, but readiness gaps may be less visibly blocker-grade to technical readers. | Medium-high: lockable with learner profile, outcomes, format, review, accessibility, and assessment answers. | Medium-high: good outcomes and review evidence, weaker privacy/safety pressure than research. | Low: easy to explain and keep small. | High: can avoid training delivery, HR policy changes, and behavioral impact claims. | Later case |

## Recommended second case

Choose the **Research protocol** case.

## Why this case

Research protocol is the best second v0.5 benchmark case because it expands the
proof surface farthest from software dashboards while staying easy to audit.
The request can be vague in one sentence, yet readiness depends on concrete
answers about participants, consent, safety, data handling, study artifacts,
review authority, and non-goals. That makes it ideal for proving that `ni`
blocks unclear intent before downstream work and can later lock a non-code
project contract for a document/human-team handoff.

It also has strong boundary safety. A successful benchmark can claim protocol
artifact readiness, lock validity, and bounded prompt compilation without
claiming that fieldwork happened, participant data was collected, an ethics
review was granted, analysis was run, or cooling interventions improved real
outcomes.

Existing repository material should be treated as context only:

- `testdata/benchmark/vague-requests/community-heat-field-study/` provides a
  suitable vague fixture and reviewer seed notes.
- `examples/research-protocol/` proves that a research protocol can be locked,
  but the v0.5 benchmark should create fresh isolated evidence instead of
  reusing old lock/run output as a new measurement.

## Draft vague request

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

## Expected first-run blockers

Likely first-run blockers should stay concrete and auditable. The exact IDs
will be assigned in the later benchmark workspace, but the expected blocker
shape is:

- `OQ-001`: Research question, decision, and final study artifact are not
  accepted.
- `OQ-002`: Participant or observation scope, inclusion/exclusion criteria, and
  locations are not accepted.
- `OQ-003`: Consent, privacy, data handling, translation, accessibility, and
  sensitive-data boundaries are not accepted.
- `OQ-004`: Field-team safety, heat/weather stop rules, and vulnerable-group
  safeguards are not accepted.
- `OQ-005`: Review owner, acceptance evidence, and readiness criteria before
  fieldwork are not accepted.
- Possible sync blocker: docs and `.ni/contract.json` must remain synchronized
  if the later planning conversation updates one without the other.

## Expected user-answer packet

Do not fill these in during selection. The later answer packet should collect:

- Study decision: what decision the protocol should help make and which final
  artifact is required.
- Research questions: primary question, secondary questions, and explicit
  non-claims.
- Scope: neighborhoods, observation windows, participant or non-participant
  boundaries, and inclusion/exclusion criteria.
- Method: interviews, observations, surveys, sensor review, or document review,
  plus why the chosen method is sufficient for the benchmark artifact.
- Consent and ethics: what requires ethics review before fieldwork, who reviews
  it, and what cannot proceed without approval.
- Data handling: allowed data, prohibited data, retention expectations,
  de-identification, and access limits.
- Safety: heat/weather stop rules, field-team safety owner, vulnerable-group
  safeguards, and escalation path.
- Accessibility: language, translation, disability access, and participant
  burden constraints.
- Non-goals: no fieldwork, no participant data collection, no medical advice,
  no policy commitment, no intervention deployment, no statistical outcome
  claim from `ni`.
- Acceptance evidence: required status proof, blocker disposition, risk
  mitigation review, isolated lock proof, and bounded prompt count.

## Expected measurement table

The later benchmark report should measure the same rubric as the existing
protocol without faking results:

| Criterion | Direct-to-agent prompt | ni intent-lock path | Evidence to record |
| --- | --- | --- | --- |
| Missing acceptance criteria | `not_measured` until reviewed | `not_measured` until status proof exists | Research questions, method, consent, safety, data handling, artifact acceptance |
| Unmitigated high-risk items | `not_measured` until reviewed | `not_measured` until status proof exists | Participant safety, vulnerable groups, privacy, heat exposure, false policy claims |
| Unresolved blockers | `not_measured` until reviewed | `not_measured` until CLI status exists | `OQ-*` blockers and grouped next-question output |
| Hidden assumptions | `not_measured` until reviewed | `not_measured` until docs/contract review | Recruitment, ethics, method, output, timeline, data authority |
| Non-goal coverage | `not_measured` until reviewed | `not_measured` until docs/contract review | Fieldwork, data collection, intervention deployment, medical advice, policy promises |
| Stale plan detection | `not_measured` until lock/amendment scenario exists | `not_measured` unless intentionally tested | Lock hash or changed-plan proof, only if included |
| Target prompt boundedness | `not_measured` before lock | `not_measured` until prompt compilation after lock | Character count at or below 4000 |
| Readiness before execution | `not_measured` for direct path | `not_measured` until `ni status` runs | Authoritative status output |

## Non-execution boundary

This selected case will not measure:

- implementation quality;
- downstream agent performance;
- real user impact;
- adoption;
- rework reduction;
- cost;
- latency;
- statistical effect size.

It must also avoid claiming fieldwork quality, ethics approval, participant
outcomes, intervention effectiveness, policy readiness, or analysis validity.
No downstream agent, model API, generated prompt execution, shell adapter,
queue, telemetry path, or runtime harness should be added.

## Next task

Task 164: Create isolated research-protocol benchmark workspace and measure
initial BLOCKED readiness
