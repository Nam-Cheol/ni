# Project brief

## Purpose

Draft a lightweight onboarding plan for a small team evaluating `ni` before
installing the CLI.

## Current status

This plan is no-terminal assisted. It is useful for learning and drafting, but
it is not deterministically validated and must not be treated as locked.

## Draft assumptions

- The team wants to understand the Intent Lock Protocol before adopting the
  CLI.
- A teammate or CI runner can later validate the draft with full `ni`.

## Open blocker questions

- Who will run `ni status`, `ni end`, and `ni run` before any implementation
  handoff?
- Which project will be used as the first real validation target?

## Non-goals

- Do not execute downstream work from this draft.
- Do not call a model API or web service from this example.
- Do not treat model-authored skills or model judgment as authoritative.
