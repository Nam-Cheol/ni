# Related work

This document positions `ni` against adjacent agent, specification, and harness projects.

The point is not to rank those projects. The point is to keep `ni` inside its product boundary:

```text
conversation -> planning docs -> contract -> readiness gate -> lockfile -> prompt or seed material
```

`ni` is a pre-runtime project intent compiler. It should make project intent explicit, validate whether planning is structurally ready, lock the accepted plan, and compile a short downstream prompt or seed artifact from the locked source of truth.

`ni` is not an execution runtime, SPEC runner, task queue, agent team, evidence loop, or project growth system. Those can exist downstream, but they must be derived from a locked plan rather than absorbed into the kernel.

## Positioning matrix

| Category | Examples | What they optimize for | `ni` boundary |
| --- | --- | --- | --- |
| Host enhancer | `oh-my-codex`, `oh-my-claude` | Better ergonomics inside a specific agent host | `ni` may emit host-compatible instructions after lock, but the kernel must stay host-neutral. |
| Spec-first coding OS | `Ouroboros` | A larger operating model around specs, tasks, code, and execution | `ni` owns readiness and locking, not the whole coding operating system. |
| SDD workflow toolkit | `Spec Kit`, `spec-first` | Spec-driven documents, tasks, and implementation workflow | `ni` may use spec-shaped planning inputs, but must not become a generic SPEC runner. |
| Evidence-first execution growth runtime | `orange-hyper-run` | Running work, collecting evidence, completing loops, and growing a project through runtime feedback | `ni` must not copy run or complete behavior into core. It may generate downstream-compatible seed material only after lock. |
| Heavy harness or guardrail template | `ai-harness-template`, `MoAI-ADK` | Guardrails, agents, adapters, checks, and execution scaffolding | `ni` may derive a project-specific harness proposal, but the kernel must not become that harness. |
| Design-time compiler | `agent-compiler` | Compile design-time intent into agent-facing artifacts | `ni` aligns with compile-before-run, but its source of truth is the locked project contract. |

## Host enhancers

Host enhancers such as `oh-my-codex` and `oh-my-claude` improve the user experience inside a particular model host. They can package prompts, commands, conventions, aliases, or workflow affordances for that host.

`ni` should not compete at this layer first. Its kernel should stay independent of any one host. A host integration may be useful later, but it should consume locked contract data rather than define the product model.

Boundary rules:

- Do not make Codex, Claude, or any other host the conceptual center of `ni`.
- Do not store project truth in host prompts when it belongs in `.ni/contract.json` or `docs/plan/**`.
- Only generate host-compatible seed material after `.ni/plan.lock.json` exists and hashes are valid.

## Spec-first coding OS

Spec-first coding systems such as `Ouroboros` describe a broader operating model around specifications, implementation work, and project progress. That can be useful for teams that want one system to organize both planning and execution.

`ni` is narrower. It stops at the pre-runtime boundary. Its authority is the readiness gate and lockfile, not the ongoing execution operating system.

Boundary rules:

- Do not turn `ni` into a project growth runtime.
- Do not add a permanent task lifecycle to the kernel.
- Keep project-specific implementation strategy in generated harnesses or downstream tools.

## SDD workflow toolkits

SDD workflow toolkits such as `Spec Kit` and `spec-first` organize work around explicit specifications. They can provide templates, task breakdowns, review steps, or conventions for moving from spec to implementation.

`ni` overlaps with the spec-first instinct but has a different authority model. Specs are inputs and derived artifacts. The locked planning contract is the source of truth after `ni end`.

Boundary rules:

- Do not make `ni` a generic SPEC runner.
- Do not treat a SPEC file as ready unless `ni status` accepts the contract.
- Do not weaken acceptance criteria or open-question handling to fit a spec workflow.

## Evidence-first execution growth runtimes

Evidence-first execution runtimes such as `orange-hyper-run` optimize for running work, observing evidence, and using completion loops to grow or advance a project. That is intentionally downstream from `ni`.

`ni` can make those runtimes safer by compiling a locked, explicit plan before they run. It should not import their runtime loop into the kernel.

Boundary rules:

- Do not add `run`, `complete`, retry, queue, or evidence-loop behavior to the kernel.
- Do not copy Hyper Run behavior into core commands.
- Do not let downstream evidence mutate locked planning docs silently.
- Generate downstream-compatible seed material only from a valid lock.

## Heavy harness and guardrail templates

Harness templates such as `ai-harness-template` and `MoAI-ADK` can define adapters, agents, guardrails, evidence expectations, test commands, and execution scaffolding.

`ni` should be able to describe what a project-specific harness needs, but that harness is derived and mutable. The kernel remains authoritative only for the contract, readiness gate, lockfile, prompt compiler, and source-of-truth rule.

Boundary rules:

- Do not add agent teams to the kernel.
- Do not add adapter execution to the kernel.
- Do not add evidence runners or task queues to the kernel.
- Keep generated harness material traceable to `.ni/plan.lock.json`.

## Design-time compilers

Design-time compiler projects such as `agent-compiler` share the idea that agent-facing artifacts should be compiled from higher-level intent before execution begins.

`ni` fits this family, with a stricter project-planning source of truth. It compiles from planning docs and `.ni/contract.json`, then refuses stale or incomplete plans through deterministic gates.

Boundary rules:

- Treat compilation as pre-runtime preparation.
- Keep validation deterministic where possible.
- Make every downstream artifact traceable to the locked contract.

## Practical rule

When a proposed feature sounds like execution, orchestration, project growth, agent management, or completion tracking, it belongs outside `ni-kernel` unless it is a read-only or compile-only artifact derived from a valid lock.
