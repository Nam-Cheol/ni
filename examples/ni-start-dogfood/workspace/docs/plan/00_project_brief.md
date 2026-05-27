# Project brief

## Product type

conversation_product

## Delivery surfaces

- conversation
- document

## Purpose

Refund Triage Assistant Plan defines a support-agent assistant that drafts
refund recommendations from ticket facts and the internal refund policy page,
escalates ambiguity, and excludes refund approval, customer contact, and
runtimes.

## Problem

The starting idea was vague: "a refund triage assistant for support agents."
The planning conversation turned that into accepted requirements, non-goals,
risks, transcript evaluations, and a machine-readable contract without the user
typing contract authoring commands.

## Success definition

The example succeeds when `ni status` validates the model-maintained
`docs/plan/**` and `.ni/contract.json`, `ni end` writes the lock after explicit
confirmation, and `ni run` compiles a target prompt without executing
implementation.
