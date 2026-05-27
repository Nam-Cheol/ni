# Security Policy

`ni` is an early, source-first project. The current product is `ni-kernel`, a
Project Intent Compiler that validates, locks, and compiles planning contracts.
It does not execute downstream agents, shell commands, queues, or runtime
adapters.

## Supported Versions

There is no stable supported release channel yet. Source tags may exist, but
they are not a stable security support channel unless release documentation says
otherwise.

## Reporting Security Issues

No private vulnerability reporting channel has been published yet.

Do not post secrets, credentials, private prompts, proprietary planning
contracts, or sensitive vulnerability details in public issues. Use GitHub
issues only for non-sensitive security reports, documentation corrections, or
questions about this policy.

A private security contact may be added in a future policy update.

## Scope

This policy applies to `ni-kernel` project files and behavior, including:

- planning docs and contract schemas,
- `.ni/contract.json` validation,
- `.ni/plan.lock.json` lock validation,
- source-of-truth and stale-lock checks,
- prompt and export compilation.

## Out of Scope

This policy does not cover downstream tools or execution environments,
including:

- Codex,
- Hyper Run,
- namba-ai,
- Spec Kit,
- Ouroboros,
- shell commands,
- generated prompts after they are executed outside `ni`.

## Secret Handling

Do not put secrets in `docs/plan/**`, `.ni/contract.json`, generated prompts,
examples, issues, or documentation. Planning contracts and prompts should be
treated as source-visible project artifacts unless a separate private workflow
is explicitly provided by the project owner.

## Runtime Boundary

`ni-kernel` compiles and validates intent before downstream execution. It does
not run downstream runtimes, agents, shell commands, or generated prompts.
