# README Pamphlet Strategy

This strategy defines the next job for `README.md`: make it a product landing
page for `ni`, not a technical specification.

`README.md` remains the canonical public README. `README.ko.md` remains the
Korean companion while companion docs are maintained. Deeper concept,
comparison, protocol, benchmark, roadmap, and command material should live in
docs and be linked from the README instead of explained inline.

## Job

The README should answer five first-screen questions in order:

1. What problem does `ni` solve?
2. What payoff does a locked intent contract give me?
3. What is the shortest path to try it?
4. What signals prove this is bounded and trustworthy?
5. Where do I read the deeper version?

The first screen should sell the product idea:

- agents are often started before project intent is accepted and testable;
- `ni` moves control before execution by compiling accepted planning state;
- the result is a locked, versioned, verifiable project contract;
- downstream actors receive bounded prompts or inert seed material only after
  readiness and lock checks pass.

The hero must not name specific downstream harness products. It may mention
models or workspaces as examples only when needed, such as Codex or Claude, but
the primary hero claim should be product-level: compile intent before any agent
or team starts work.

## Proposed README Section Order

1. Hero
2. Badges
3. Why ni
4. Three-step use path
5. Short demo
6. Install and use options
7. Trust signals
8. What ni is not
9. Read next
10. Development and release status

## Section Intent

### Hero

Keep the hero short and memorable:

- product name;
- language toggle;
- tagline: `Project Intent Compiler for AI Agents`;
- one sentence that explains the problem;
- one sentence that explains the payoff;
- a tiny flow showing conversation, readiness, lock, and prompt compilation.

Avoid product-by-product downstream lists in the hero. The hero should not feel
like target documentation.

### Badges

Use badges only as trust signals that are already true:

- CI status if available;
- license;
- source-first status if represented plainly.

Do not imply package-manager distribution, hosted service availability, or a
published binary release unless those paths exist.

### Why ni

Explain the user pain before the mechanism:

- prompts can sound actionable while hiding users, acceptance criteria, risks,
  non-goals, and blocker questions;
- hidden assumptions become expensive after an agent starts;
- `ni` makes intent explicit, checks it deterministically, locks it, and stops
  handoff when it changes.

Keep protocol terms light here. Link to the Intent Lock Protocol for the full
mechanism.

### Three-Step Use Path

Show the shortest source-first path:

```text
1. init a planning workspace
2. check readiness and resolve blockers
3. lock intent and compile a prompt
```

Use only implemented commands:

- `ni init`
- `ni status`
- `ni end`
- `ni run`

Do not add contract authoring CLI commands. Authoring remains model-user
conversation over docs and `.ni/contract.json`, with the CLI as authority.

### Short Demo

Keep one short demo near the top. The best default is the ambiguous prompt
blocked demo because it shows the product value quickly:

- start from a vague request;
- run `ni status`;
- show `BLOCKED`;
- explain that this is success because execution should not begin yet.

The demo must not run downstream agents or shells beyond local validation and
CLI commands already covered by smoke checks.

### Install and Use Options

Claim only implemented paths:

- source-first `go run`;
- local build;
- local install if the repository supports it.

Do not claim Homebrew, package-manager distribution, hosted binaries, or
release automation.

### Trust Signals

Move trust signals above deep references:

- deterministic readiness gate;
- lockfile hash verification;
- prompt output budget;
- `ni run` compiles prompts and does not execute them;
- CI, smoke checks, quality checks, demo checks, install checks;
- source-first release status;
- MIT license and security policy.

These signals should reassure the reader that `ni` is a pre-runtime kernel, not
an execution runtime.

### What ni Is Not

Keep this as one short boundary block. It should say that `ni` is not an
execution harness, task runner, multi-agent orchestration layer, queue, shell
adapter, or release system.

The README should not carry target-by-target boundary prose. Link to deeper
target docs instead.

### Read Next

End the pamphlet with a small map:

- install details;
- Intent Lock Protocol;
- command reference;
- target story;
- benchmark protocol;
- release readiness;
- launch checklist;
- post-release roadmap.

## Move Out Of README

The following material should move to docs and be linked from README:

- detailed protocol explanation;
- related-work comparison;
- target-by-target boundaries;
- full command reference;
- benchmark methodology;
- post-release roadmap.

README may keep one-sentence summaries of these topics, but the body should not
become the specification.

## Keep In README

The README should keep:

- hero;
- badges;
- why ni;
- three-step use path;
- short demo;
- install and use options;
- trust signals;
- one short `What ni is not` block;
- read-next links.

## Rewrite Guardrails

- Keep `README.md` canonical and `README.ko.md` as companion.
- Do not name specific downstream harness products in the README hero.
- Do not claim install paths that are not implemented.
- Do not add execution runtime behavior.
- Do not add contract authoring CLI commands.
- Do not add model execution adapters, shell adapters, queues, or agent
  orchestration.
- Keep `ni run` documented as prompt compilation only.

## Success Test

After the rewrite, a new reader should be able to answer:

- why `ni` exists;
- why they should care before starting an agent;
- how to try it from source;
- why the kernel boundary is trustworthy;
- which deeper doc to read next.

If the README first screen feels like a protocol reference, command manual,
target matrix, benchmark paper, or roadmap, the pamphlet rewrite has failed.
