# Contract summary: Neighborhood Cooling Study Protocol

## Identity

- Product type: `research_protocol`
- Delivery surface: `document`
- Interaction mode: `human_to_human`
- Readiness profile: `prototype`
- Lock status: `READY`

## Purpose

Plan a documented field research protocol for comparing street-level cooling
interventions before any study execution begins.

## Accepted capabilities

- `CAP-001`: Define the study question, hypothesis, sampling frame, and
  inclusion criteria.
- `CAP-002`: Specify data handling, ethics, and participant-contact boundaries.
- `CAP-003`: Define reproducible analysis and evidence requirements for the
  protocol.

## Evaluation model

- `EVAL-001`: Independent sampling reproducibility review.
- `EVAL-002`: Ethics boundary review for data items and participant contact.
- `EVAL-003`: Analysis plan reproducibility review.

The evaluations are reviewer and protocol based. They are not automated
software tests.

## Non-goal boundary

ni does not run fieldwork, collect data, deploy sensors, perform analysis, or
grant ethics approval. `ni run` only compiles the locked plan into a downstream
handoff prompt.
