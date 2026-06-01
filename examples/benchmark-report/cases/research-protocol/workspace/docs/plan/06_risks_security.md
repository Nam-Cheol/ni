# Risks and security

## RISK-001: A downstream actor could invent the study question, method, locations, or intervention decision criteria.

Severity: high

Mitigation: Use only the synthetic benchmark fixture answers for the accepted
research question, supported decision, final artifact, participant or
observation scope, inclusion and exclusion criteria, and locations. Record
that the fixture selects candidate blocks for later design review only and
does not choose final intervention locations.

## RISK-002: Participant privacy, consent, translation, accessibility, or sensitive-data boundaries could be skipped.

Severity: high

Mitigation: Record allowed and prohibited data, non-identifying voluntary
comment consent language, segment-level aggregation, 30-day raw note reduction,
English and Spanish materials, accessibility expectations, and explicit
sensitive-data exclusions before lock.

## RISK-003: Heat or weather field safety and vulnerable-group safeguards could be under-specified.

Severity: high

Mitigation: Record paired field teams, water, charged phones, public-area-only
work, heat/weather stop rules, 30-minute continuous exposure limits, 10-minute
rest breaks, maximum 3 outdoor observation hours per field day, vulnerable
group safeguards, and emergency escalation path before lock.

## RISK-004: The benchmark could be misrepresented as fieldwork readiness or intervention decision readiness.

Severity: high

Mitigation: State in docs, contract, evidence files, and case reports that
synthetic answers are benchmark fixture data only. No fieldwork, participant
recruitment, data collection, analysis, intervention placement, or research
outcome is validated.

## RISK-005: A handoff prompt could be compiled before readiness, making unresolved research intent look actionable.

Severity: high

Mitigation: Do not run ni end or ni run while ni status reports BLOCKED. If the
resolved fixture reports READY or READY_WITH_DEFERRALS, run ni end and ni run
only inside the isolated benchmark workspace and record the evidence without
executing downstream work.
