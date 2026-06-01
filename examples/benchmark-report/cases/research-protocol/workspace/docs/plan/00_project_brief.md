# Project brief

## Product type

research_protocol

## Delivery surfaces

- document
- workflow
- human_service

## Purpose

Prepare an initial benchmark planning workspace for a vague summer neighborhood
cooling study request so ni can measure whether required research intent is
blocked before any fieldwork, participant recruitment, data collection,
intervention decision, downstream agent, or model API call begins.

## Problem

The vague request asks:

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

This sounds actionable, but it does not yet define the research question,
supported decision, participant or observation scope, consent and privacy
rules, data handling, accessibility, heat/weather safety rules, review owner,
or acceptance evidence. Those gaps should remain visible instead of being
invented by a downstream actor.

## Success definition

For this initial benchmark task, success means `ni status --proof
--next-questions` reports the unresolved research-protocol intent honestly. If
the result is `BLOCKED`, no lock is created, no prompt is compiled, and prompt
count remains `not_measured`.
