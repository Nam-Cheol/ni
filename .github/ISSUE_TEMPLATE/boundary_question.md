---
name: Boundary question
about: Ask whether a proposal fits ni-kernel or violates the non-execution boundary
title: "Boundary: "
labels: boundary
assignees: ""
---

## Proposed Change

<!-- Describe the feature, workflow, target, or integration you are considering. -->

## Where It Might Belong

<!-- Pick the closest current home. -->

- `ni-kernel` readiness validation
- `ni-kernel` docs/contract sync
- `ni-kernel` lock/hash correctness
- `ni-kernel` prompt compilation
- inert downstream target seed
- outside `ni-kernel`

## Non-Execution Boundary

<!-- ni-kernel must not become an execution runtime. -->

- Would this start, schedule, track, or complete downstream work?
- Would this add a task runner, SPEC runner, Codex exec, shell adapter, queue, agent team, PR automation, or release automation?
- Would this create kernel-owned runtime state for a downstream tool?
- If it produces target material, is that material derived from a valid locked plan?

## Why the Boundary Is Unclear

<!-- What part needs project maintainer judgment? -->
