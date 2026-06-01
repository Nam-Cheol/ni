# Research protocol benchmark answer packet

This packet records the synthetic fixture answers used for the
research-protocol benchmark blockers. These answers were created for ni
benchmark evidence only. They are not real fieldwork approval, actual research
authorization, proof of research quality, or empirical evidence.

Use answers only inside:

```text
examples/benchmark-report/cases/research-protocol/workspace/
```

Do not apply them to the repository-root planning lock or root `.ni/` state.

## Status

- Initial readiness: `BLOCKED`
- Resolved readiness after applying synthetic fixture answers: `READY`
- Lock created: yes, inside the isolated benchmark workspace only
- `ni-run` prompt compiled: yes, bounded to 4000 characters
- Required answers: `OQ-001` through `OQ-005`

## How to use this packet

1. Treat the answers below as synthetic benchmark fixture answers only.
2. Use the answers to update only the isolated research-protocol benchmark
   workspace.
3. Run `ni status --proof --next-questions` against that workspace.
4. Lock only if readiness becomes `READY` or `READY_WITH_DEFERRALS`.
5. Run `ni run` only after a valid benchmark workspace lock.

## OQ-001 - Research question, supported decision, and final study artifact

Prompt:

What research question should this study answer, what decision should it
support, and what final artifact is enough for planning?

Required answer fields:

- Research question: Which public outdoor blocks in the fictional Riverside
  East and Oak Market corridors appear to need further shade or cooling
  intervention review based on observed shade deficit, public-space heat
  exposure, and non-identifying community feedback?
- Supported decision: Select a short list of candidate blocks for a later
  design review. The study does not choose final intervention locations.
- Decision owner: City Resilience Program planning owner.
- Final study artifact: A planning memo containing the research question,
  candidate-block shortlist, evidence table, privacy and safety boundaries,
  limitations, and a pre-fieldwork readiness checklist.
- What should not be decided by this study: Final intervention placement,
  construction scope, procurement, budget approval, public-health claims,
  clinical heat-risk conclusions, or production deployment.

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

- Observation units: Public street segments, public plazas,
  transit-adjacent waiting areas, and publicly accessible pedestrian corridors.
- Participant groups, if any: Optional adult community members who voluntarily
  give non-identifying feedback. No minors, no health-status collection, and no
  targeted vulnerable-group recruitment.
- Inclusion criteria: Outdoor public locations in Riverside East or Oak Market
  with visible pedestrian use, low shade coverage, or prior public planning
  interest.
- Exclusion criteria: Private property, schools, care facilities, indoor
  spaces, locations requiring special access, locations where observation would
  create safety risk, and any interaction involving minors.
- Locations or neighborhood types: Riverside East commercial corridor, Oak
  Market transit corridor, and adjacent public waiting areas selected from a
  pre-approved location list.
- Observation window: Weekdays between 09:00 and 16:00 during the summer study
  period, only when field safety conditions permit.
- Sampling or selection rule: Select up to 12 public segments using a mix of
  shade deficit, pedestrian activity, and planning relevance. Record why each
  segment was included.
- Out-of-scope populations or locations: Private households, medical
  facilities, schools, minors, health-status groups, individual resident
  tracking, and non-public spaces.

Unsafe assumptions to avoid:

- Recruiting vulnerable participants without criteria.
- Treating all neighborhoods as interchangeable.
- Assuming any hot day or any location is valid.

## OQ-003 - Consent, privacy, data handling, translation, accessibility, and sensitive-data boundaries

Prompt:

What data may be collected or used, how is consent handled, what data is
sensitive, and what accessibility or translation support is required?

Required answer fields:

- Data types allowed: Public-space observation notes, segment-level shade
  observations, public infrastructure notes, non-identifying voluntary
  comments, public map references, and aggregated segment-level evidence.
- Data types prohibited: Names, contact information, photos of identifiable
  people, license plates, medical or health data, household-level data, precise
  resident addresses, credentials, tokens, private customer data, and raw
  audio/video.
- Consent approach: No identifiable participant data is collected. For
  voluntary comments, use a short plain-language consent script stating that
  comments are optional, non-identifying, and used only for planning evidence.
  Do not collect comments from minors.
- Privacy constraints: Keep all notes aggregated at segment level. Do not
  include direct quotes if they could identify a person. Do not publish raw
  notes. Do not collect or store identifiers.
- Storage/retention rule: Raw field notes are stored in the project workspace
  for up to 30 days, then reduced to a summarized planning memo. The final memo
  keeps only aggregated, non-identifying evidence.
- Translation needs: Provide the consent script and short study explanation in
  English and Spanish for this fictional benchmark case.
- Accessibility needs: Use plain-language materials, readable font size, and
  accessible observation routes. Do not require participants to use a digital
  form.
- Sensitive-data boundaries: Do not collect personal, health, financial,
  immigration, employment, or household-level information. Do not record
  vulnerable-group status.
- Data that must remain out of scope: Identifiable people, medical history,
  household-level heat exposure, private utility usage, financial data, precise
  home addresses, and any source requiring restricted access.

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

- Field team safety rules: Field teams work in pairs, carry water, have
  charged phones, check weather before departure, and stay in public areas
  only.
- Heat/weather stop conditions: Cancel or pause fieldwork if heat index
  exceeds 38 C, an official heat emergency is active, lightning occurs within
  10 miles, air quality is unhealthy, or the field lead judges conditions
  unsafe.
- Maximum exposure or shift limits: Maximum 30 minutes of continuous outdoor
  observation followed by at least 10 minutes of shaded or indoor rest. Maximum
  total outdoor observation time is 3 hours per field day.
- Hydration/rest requirements: Each team member carries water, takes scheduled
  shade breaks, and may stop immediately without penalty if heat stress
  symptoms appear.
- Vulnerable-group safeguards: Do not target minors, elderly residents,
  unhoused residents, or medically vulnerable people for comments. Do not ask
  health-status questions. If a vulnerable person volunteers feedback, keep it
  non-identifying and optional.
- Emergency escalation path: Field team contacts field lead first. For urgent
  danger, call local emergency services. Field lead records the incident and
  cancels the field session if needed.
- Conditions that cancel fieldwork: Extreme heat, unsafe air quality,
  lightning, civil safety concern, lack of two-person team, lack of water/rest
  access, or any site condition that makes observation unsafe.

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

- Reviewer or approval owner: City Resilience Program planning owner plus a
  privacy/safety reviewer.
- Review audience: Planning owner, field lead, privacy/safety reviewer, and
  community engagement lead.
- Minimum protocol artifact: A protocol memo including research question,
  location selection rule, observation scope, consent script, privacy
  boundaries, safety stop rules, non-goals, risk mitigations, and acceptance
  checklist.
- Acceptance evidence: All OQ fields answered, high risks mitigated,
  non-goals explicit, consent/privacy/safety sections reviewed, fieldwork stop
  rules documented, and `ni status` returns `READY` or
  `READY_WITH_DEFERRALS` in the isolated benchmark workspace.
- Pass/fail criteria: Pass if the protocol memo is complete, privacy and
  safety boundaries are explicit, fieldwork stop rules are testable, the
  reviewer is named, and no blocker is silently resolved. Fail if consent is
  vague, safety rules are missing, identifiable data may be collected,
  vulnerable-group safeguards are absent, or acceptance owner is unclear.
- Pre-fieldwork readiness checklist: Research question accepted; locations
  selected from approved list; data types allowed/prohibited; consent script
  prepared; translation/accessibility needs addressed; field safety rules
  documented; reviewer assigned; risks mitigated; non-goals explicit;
  `ni status` not `BLOCKED`.
- What is explicitly not required: Actual fieldwork, participant recruitment,
  data collection, statistical analysis, final intervention placement, budget
  approval, dashboard or software implementation, or downstream agent
  execution.
- Who can approve moving forward: City Resilience Program planning owner and
  privacy/safety reviewer jointly.

Unsafe assumptions to avoid:

- Treating a draft protocol as fieldwork-ready.
- Treating informal agreement as approval.
- Starting data collection before acceptance evidence exists.

## After answers are provided

Completed benchmark steps:

1. Updated only the isolated benchmark workspace planning state.
2. Updated `docs/plan`.
3. Updated `.ni/contract.json`.
4. Updated `.ni/session.json`.
5. Ran `ni status --proof --next-questions`.
6. Recorded `READY` in `11-resolved-status-proof.md`.
7. Locked only inside the benchmark workspace.
8. Compiled a bounded prompt from the benchmark workspace.
9. Updated measurement table honestly.

Rules:

- Do not edit root `.ni/plan.lock.json`.
- Do not run repository-root `ni end` or `ni relock`.
- Do not run `ni end` or `ni run` in the research-protocol benchmark workspace
  while it reports `BLOCKED`.
- Do not execute downstream agents.
- Do not run fieldwork.
- Do not recruit participants.
- Do not collect data.
- Do not call model APIs.
- Do not fake prompt or lock evidence.
