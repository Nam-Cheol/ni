---
name: Feature request
about: Propose a ni-kernel improvement or downstream seed format
title: "Feature: "
labels: enhancement
assignees: ""
---

## Proposal

<!-- What should change? -->

## Kernel or Downstream Seed?

<!-- ni-kernel validates, locks, and compiles project intent. Downstream seed material is derived and inert. -->

- Does this belong in `ni-kernel`?
- Or should it be downstream target seed material?
- What locked-plan data would this feature consume or validate?

## Boundary Check

<!-- Core contributions must not add execution runtime behavior. -->

- Does this start, schedule, track, or complete downstream work?
- Does this add a task runner, SPEC runner, Codex exec, shell adapter, queue, agent team, PR automation, or release automation?
- If yes, why is this not better handled outside `ni-kernel`?

## Expected User Value

<!-- What planning, validation, locking, prompt compilation, example, fixture, or documentation problem does this solve? -->

## Acceptance Notes

<!-- What evidence would prove this feature works without weakening ni's boundary? -->
