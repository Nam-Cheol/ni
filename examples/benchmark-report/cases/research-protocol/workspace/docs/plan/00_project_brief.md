# Project brief

## Product type

research_protocol

## Delivery surfaces

- document
- workflow
- human_service

## Purpose

Prepare a benchmark planning workspace for a vague summer neighborhood cooling
study request so ni can measure the transition from unresolved research intent
to an explicit, locked planning contract before any fieldwork, participant
recruitment, data collection, intervention decision, downstream agent, or model
API call begins.

## Problem

The vague request asks:

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

This sounds actionable, but it did not initially define the research question,
supported decision, participant or observation scope, consent and privacy
rules, data handling, accessibility, heat/weather safety rules, review owner,
or acceptance evidence. The initial benchmark kept those gaps visible instead
of allowing a downstream actor to invent them.

This resolved benchmark variant fills those gaps with synthetic benchmark
fixture answers. The answers are not real fieldwork approval, not actual
research authorization, not proof of research quality, and not empirical
evidence. They exist only to test ni's pre-runtime intent-readiness flow.

## Synthetic fixture scope

Research question:

Which public outdoor blocks in the fictional Riverside East and Oak Market
corridors appear to need further shade or cooling intervention review based on
observed shade deficit, public-space heat exposure, and non-identifying
community feedback?

Supported decision:

Select a short list of candidate blocks for a later design review. The study
does not choose final intervention locations.

Decision owner:

City Resilience Program planning owner.

Final study artifact:

A planning memo containing the research question, candidate-block shortlist,
evidence table, privacy and safety boundaries, limitations, and a pre-fieldwork
readiness checklist.

What this study must not decide:

Final intervention placement, construction scope, procurement, budget approval,
public-health claims, clinical heat-risk conclusions, or production deployment.

## Success definition

For this resolved benchmark task, success means the isolated workspace records
the synthetic fixture answers, `ni status --proof --next-questions` reports
the measured readiness honestly, and any lock or bounded prompt is created only
inside this benchmark workspace after the CLI allows it. Success does not mean
real fieldwork readiness or research approval.
