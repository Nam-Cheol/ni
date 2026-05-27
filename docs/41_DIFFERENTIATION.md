# Differentiation Map

`ni` is the pre-runtime Project Intent Compiler for AI Agents.

It does not replace Codex, Spec Kit, Hyper Run, Ouroboros, namba-ai,
oh-my-codex, or oh-my-claude. It gives downstream tools a locked intent
contract to obey before they plan, orchestrate, execute, evaluate, or refine
work.

The category boundary is:

```text
conversation -> docs/plan + .ni/contract.json -> ni status -> ni end -> locked intent -> downstream use
```

`ni-kernel` owns the planning contract, readiness gate, lockfile, prompt
compiler, and source-of-truth rule. Downstream seed material is derived from a
valid lock and remains mutable outside the kernel.

## Comparison Table

| Row | oh-my-codex / oh-my-claude | Spec Kit / spec-first | Ouroboros | orange-hyper-run | namba-ai | ni |
| --- | --- | --- | --- | --- | --- | --- |
| When does it operate? | Inside or around a model host while a user works in that host. | During software specification and implementation workflow. | Across a spec-first coding lifecycle. | During execution and evidence collection loops after work starts. | Around Codex-oriented prompt refinement and workflow orchestration. | Before runtime, while project intent is still being compiled, validated, and locked. |
| Primary artifact | Host-specific prompts, commands, conventions, or orchestration affordances. | Specs, tasks, implementation plans, and related SDD workflow materials. | Specs plus coding-agent operating state and lifecycle records. | Runtime work packets, evidence records, completion state, and growth feedback. | Refined prompts, workflow instructions, and Codex-oriented orchestration material. | `docs/plan/**`, `.ni/contract.json`, `.ni/plan.lock.json`, and bounded downstream prompts or seed notes. |
| Does it execute? | It may orchestrate or enhance actions inside a host, depending on the host layer. | It may drive or organize implementation workflow, depending on the toolkit. | Yes, it is oriented around a coding-agent operating model that can include execution phases. | Yes, execution and evidence loops are central. | It may orchestrate Codex-facing workflow, but execution belongs to the host or downstream system. | No. `ni run` compiles a prompt only, and exports write inert seed material only. |
| Product scope | Host enhancer or orchestration layer. | Software SDD workflow toolkit. | Spec-first coding Agent OS. | Evidence-first execution growth runtime. | Codex-oriented prompt refinement and workflow orchestration. | Pre-runtime Project Intent Compiler for AI Agents. |
| Authority model | Host conventions and user workflow shape behavior. | Spec and workflow conventions guide downstream implementation. | The Agent OS governs the spec-to-code lifecycle. | Runtime evidence and completion loops guide project growth. | Prompt and workflow orchestration guides Codex-facing work. | The CLI readiness gate and lockfile are authoritative; skills and models are UX. |
| Main failure it prevents | Repetitive or inconsistent host usage. | Ambiguous software implementation without explicit specs. | Unstructured coding-agent progress across a larger lifecycle. | Unverified execution, weak evidence, or incomplete runtime loops. | Low-quality prompts and scattered Codex workflow state. | Starting any downstream agent, harness, toolkit, or runtime from ambiguous, stale, or unlocked intent. |
| How downstream tools consume it | They may receive host-compatible prompt seeds or instructions after lock. | They may receive locked acceptance criteria, risks, constraints, and spec seed notes. | It may receive locked planning context and seed notes without making `ni` own its lifecycle. | It may receive readiness expectations, evidence requirements, and first-run focus from a valid lock. | It may receive Codex-oriented prompt seed material and suggested spec boundaries from a valid lock. | `ni` is the upstream producer: downstream tools consume locked contracts, compiled prompts, or seed packages. |

## Boundary Notes

### Host enhancers

`oh-my-codex` and `oh-my-claude` improve or orchestrate work inside a specific
agent host. `ni` should stay host-neutral in the kernel. Host-specific output
is allowed only as downstream-compatible prompt or seed material derived from a
valid lock.

### SDD workflow toolkits

`Spec Kit` and `spec-first` help teams move from software specification to
implementation. `ni` can produce seed notes for such tools, but it must not
become a generic SPEC runner or weaken readiness rules to fit a spec workflow.

### Coding-agent operating systems

`Ouroboros` is a broader spec-first coding Agent OS. `ni` is deliberately
narrower: it validates and locks project intent, then stops before owning the
coding lifecycle.

### Execution growth runtimes

`orange-hyper-run` operates after work starts, where execution, evidence, and
completion loops matter. `ni` must not copy that runtime behavior into core.
Its value is to give an execution runtime a clear locked plan before it runs.

### Prompt refinement and orchestration

`namba-ai` can refine prompts and orchestrate Codex-oriented workflows. `ni`
can hand it locked intent and seed boundaries, but `ni` should not become a
Codex-only prompt workflow or namba runtime.

## Practical Rule

If a feature would run tasks, manage agents, maintain queues, collect runtime
evidence, adapt to a target runtime, or complete downstream work, it is outside
`ni-kernel`.

If a feature validates planning state, locks accepted intent, checks locked
hashes, compiles a bounded prompt, or writes inert seed material from a valid
lock, it can fit inside `ni-kernel`.
