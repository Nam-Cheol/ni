# Product principles

## 1. Plans are external state

Do not rely on a model's memory of the conversation. Important intent must be written into `docs/plan` and `.ni/contract.json`.

## 2. Readiness is not a feeling

Readiness is a deterministic gate. It is not the model's impression that the docs are good enough.

## 3. Lock before execution

Execution prompts should be compiled only from a locked project contract.

## 4. Prompts should be short

`ni run` must emit a 4000-character-or-less prompt. The prompt should point to authoritative files rather than embed the entire plan.

## 5. Harnesses are derived

A project-specific harness should be generated from the locked contract. It should not become the kernel.

## 6. Provider-specific behavior is outside the kernel

Codex is the first experimental execution environment, not the conceptual center of `ni`.
