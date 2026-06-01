# Risks and security

## RISK-001: A downstream actor could invent the study question, method, locations, or intervention decision criteria.

Severity: high

Mitigation: Keep OQ-001 and OQ-002 open blockers until the user accepts the
research question, supported decision, final artifact, participant or
observation scope, inclusion and exclusion criteria, and locations.

## RISK-002: Participant privacy, consent, translation, accessibility, or sensitive-data boundaries could be skipped.

Severity: high

Mitigation: Keep OQ-003 open as a blocker and require accepted consent,
privacy, data handling, translation, accessibility, and sensitive-data
boundaries before lock.

## RISK-003: Heat or weather field safety and vulnerable-group safeguards could be under-specified.

Severity: high

Mitigation: Keep OQ-004 open as a blocker and require accepted heat or weather
stop rules, field-team safety ownership, vulnerable-group safeguards, and
escalation path before lock.

## RISK-004: The benchmark could be misrepresented as fieldwork readiness or intervention decision readiness.

Severity: high

Mitigation: State in docs, contract, and case reports that this task measures
initial intent readiness only; no fieldwork, participant recruitment, data
collection, analysis, intervention placement, or research outcome is validated.

## RISK-005: A handoff prompt could be compiled before readiness, making unresolved research intent look actionable.

Severity: high

Mitigation: Do not run ni end or ni run while ni status reports BLOCKED; record
lock as no, bounded prompt as no, and prompt count as not_measured.
