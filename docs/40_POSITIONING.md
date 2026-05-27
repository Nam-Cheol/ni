# Positioning

`ni` is the Project Intent Compiler for AI Agents.

Don't run the agent yet. Compile the intent first.

`ni` turns planning conversations into a locked, versioned, verifiable project
contract before Codex, Claude, Spec Kit, Hyper Run, namba-ai, a generated
harness, or a human team starts execution.

Lock intent before any harness runs.

## Category

`ni` is a pre-runtime intent contract layer. It sits upstream of execution
systems and compiles accepted planning state into locked source-of-truth
material that downstream systems may consume.

## Distinctions

`ni` is not:

- a task runner,
- a SPEC runner,
- a multi-agent execution layer,
- a Codex adapter,
- Hyper Run,
- Spec Kit,
- Ouroboros,
- namba-ai.

Those systems may consume locked `ni` output. They do not define `ni`'s kernel
authority.

## Kernel Boundary

`ni-kernel` owns the planning contract, readiness gate, lockfile, prompt
compiler, and source-of-truth rule.

Downstream prompts, seed packages, work graphs, evaluation-plan proposals,
evidence-rule notes, and harness seed material are derived and mutable. They
must be traceable to `.ni/plan.lock.json`, and they must not become
kernel-owned execution state.
