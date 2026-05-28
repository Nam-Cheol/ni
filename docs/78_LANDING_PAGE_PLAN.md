# Landing Page Plan

Task 137 decides whether `ni` needs a GitHub Pages landing page beyond the
canonical README.

This document is a planning note only. It does not implement a site, add a
frontend stack, add analytics, create a hosted service, change install paths, or
make a website required for `ni` usage.

## Recommendation

Recommendation: GitHub Pages static HTML.

Keep `README.md` as the canonical quick entry. Add a minimal GitHub Pages page
only if the next public-facing task needs a cleaner share surface than the
GitHub repository view. The page should be a thin, static pamphlet that links
back to README and docs instead of becoming a second product manual.

Do not build a docs site now. The current product is still best served by a
small source-first landing surface and explicit protocol docs in the repository.

## Decision Summary

| Option | Fit | Tradeoff | Decision |
| --- | --- | --- | --- |
| README only | Strong default because README is already the canonical quick entry and product pamphlet. | GitHub's repository frame is less focused for sharing, screenshots, and first-impression copy. | Keep as source of truth. |
| GitHub Pages static HTML | Useful as a lightweight public doorway for launches, demos, and links from social posts. | Creates a second surface that must not drift from README. | Recommended as a follow-up. |
| Docs site later | Useful only when the docs grow into a navigation problem. | Too heavy for the current kernel stage and risks premature site maintenance. | Defer. |

## Why Not README Only

README-only is enough for usage. It should remain enough forever for a user to
understand what `ni` is, install it, and find the deeper docs.

A tiny landing page can still be useful because it gives public links a cleaner
first screen:

- the hero can focus on the core promise without repository chrome;
- the first demo can be easier to scan;
- install paths can be grouped without expanding the README further;
- screenshots, release posts, and social cards can point to one concise URL.

That benefit is presentation only. It must not introduce new product behavior or
new authority.

## Minimal Page Shape

Use a single static HTML page, optionally with one small CSS file. No bundler,
framework, analytics, server component, cookie banner, tracking script, or
external runtime dependency is needed.

Recommended sections:

1. Hero
2. Why ni
3. Start in 60 seconds
4. Install paths
5. Demo
6. Docs links

### Hero

Use the existing public line:

```text
Don't run the agent yet. Compile the intent first.
```

Support it with the category:

```text
Project Intent Compiler for AI Agents.
```

The hero should say that `ni` turns planning conversations into locked project
contracts before implementation work starts. Do not name specific downstream
harnesses in the hero.

### Why ni

Keep this section short and pain-first:

- prompts can look actionable while users, acceptance criteria, risks,
  non-goals, and blocker questions are still missing;
- early execution makes hidden assumptions expensive;
- `ni` makes intent explicit, checks readiness, locks accepted state, and stops
  handoff when intent changes.

Link to the Intent Lock Protocol for the deeper mechanism.

### Start In 60 Seconds

Show only implemented commands:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
go run ./cmd/ni end --dir ./my-plan
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

Make clear that authoring happens through conversation over `docs/plan/**` and
`.ni/contract.json`, while the CLI remains the authority for readiness, locking,
hash validation, and prompt compilation.

### Install Paths

List only paths already documented as available or planned:

- Source
- Local binary
- Release binary
- Curl installer
- Model workspaces as experimental planning assistance
- No-terminal method as experimental assisted drafting
- Homebrew as planned, not available

Each path should link back to `README.md`, `docs/22_INSTALL.md`, or the relevant
source doc. Do not claim package-manager distribution, hosted binaries, or
services that are not already documented.

### Demo

Use the blocked ambiguous-prompt demo:

```bash
go run ./cmd/ni status --dir examples/ambiguous-prompt-blocked/workspace
```

```text
BLOCKED
```

Explain that `BLOCKED` is the expected success state when intent is not ready.
Do not add downstream agent execution, shell adapters, queues, or runtime
automation.

### Docs Links

Keep the link map small:

- README
- Install ni
- Intent Lock Protocol
- Command reference
- Why ni exists
- README Visual Wireframe
- Release readiness

README remains the canonical quick entry; the landing page is only a public
doorway.

## Guardrails

- Do not add a heavy frontend stack.
- Do not add analytics.
- Do not add a hosted service.
- Do not make the website required for `ni` usage.
- Do not add runtime execution behavior.
- Do not add contract authoring CLI commands.
- Do not make GitHub Pages the source of truth for commands, install status, or
  protocol rules.
- Do not duplicate long documentation from README or docs.

## Follow-Up Task

Create a separate follow-up task only if the project wants the public doorway:

```text
Task: Implement minimal GitHub Pages landing page.

Scope:
- add one static HTML page and one small CSS file, or the smallest GitHub Pages
  layout the repository already supports;
- reuse README copy and existing visual assets where appropriate;
- link back to README as canonical quick entry;
- include hero, why ni, start in 60 seconds, install paths, demo, and docs
  links;
- do not add a framework, analytics, hosted service, or runtime execution
  behavior.

Validation:
- bash scripts/quality.sh
```

## Success Test

The landing page is worth building only if a new reader can understand `ni` in
one scan, then immediately choose README or install docs for the authoritative
next step. If the page starts to become a second README, protocol reference, or
docs portal, stop and keep README-only until a later docs-site task is justified.
