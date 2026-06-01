# Project brief

## Product type

Document product: benchmark planning artifact.

## Delivery surfaces

- Document.

## Purpose

Prepare the internal-dashboard benchmark case as a planning-meeting artifact so
a planning owner, product lead, internal operations lead, reviewer, and
stakeholders can decide whether the case has enough structure, evidence, and
acceptance criteria to use in a planning meeting.

## Problem

The original request asked for a dashboard for "the customer team" to see what
is going on with accounts and know who needs attention. The user-provided
answers for `OQ-001` through `OQ-004` clarify benchmark artifact readiness, not
dashboard product readiness. The workspace must therefore preserve the product
unknowns as non-goals while evaluating whether the benchmark planning packet is
ready for review.

## Success definition

This benchmark workspace succeeds if the provided answers are captured in
`docs/plan/**` and `.ni/contract.json`, `ni status --proof --next-questions`
honestly reports the resulting readiness, and any lock or bounded prompt is
created only inside this isolated workspace after the CLI allows it.
