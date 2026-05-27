# README Visual Wireframe

Task 104 defines the visual README layout before the README is rewritten again.

This document is a layout contract only. It does not rewrite `README.md`, add
new product claims, generate visual assets, or change runtime behavior.

## Design Intent

The README should read as a lightweight product pamphlet backed by deterministic
kernel proof. The first screen should make the value clear without turning the
README into protocol documentation.

Use this core message, already established in the README and visual spec:

```text
Don't run the agent yet. Compile the intent first.
```

Keep the technical depth in linked docs. README visuals may guide scanning, but
important meaning must also exist as Markdown text.

## README Section Order

Use this order for the next README rewrite:

1. Hero banner
2. Language and trust badges
3. Slogan
4. Three pain cards
5. Three ni payoff cards
6. Start in 60 seconds
7. Choose your path
8. Demo
9. What ni is not
10. Read next

## Wireframe

```text
[hero.svg]

Language chips        Trust badges
English | 한국어       MIT | CI workflow | Security | Docs

Don't run the agent yet. Compile the intent first.
Project Intent Compiler for AI Agents.

[pain card]           [pain card]           [pain card]
Vague intent          Early execution       Rework
Markdown fallback     Markdown fallback     Markdown fallback

[payoff card]         [payoff card]         [payoff card]
Capture intent        Lock contract         Handoff safely
Markdown fallback     Markdown fallback     Markdown fallback

## Start in 60 seconds
Markdown commands and source-first path.

## Choose your path
Markdown table for source, verified install paths, and assisted planning
paths. Codex and Claude may appear here only as usage path labels.

## Demo
Markdown command transcript showing a vague prompt blocked before execution.

## What ni is not
Short Markdown boundary block.

## Read next
Markdown link map to deeper docs.
```

## Section Rules

| Section | Format | Asset role | Markdown fallback |
| --- | --- | --- | --- |
| Hero banner | SVG image plus Markdown text | Use `assets/hero.svg`. | The slogan and product description must appear below the image as Markdown. |
| Language badges | SVG chips | Use local language chip SVGs. | Link labels and alt text must identify the languages. |
| Trust badges | Badges | Existing factual remote badges may remain non-essential. | A short Markdown trust-signal sentence must repeat the facts. |
| Slogan | Markdown | No separate SVG card. | The slogan is primary text and must not be image-only. |
| Three pain cards | SVG cards plus Markdown headings | New local SVG cards may reinforce the three pains. | Each card needs a nearby Markdown heading and one-sentence explanation. |
| Three ni payoff cards | SVG cards plus Markdown headings | New local SVG cards may reinforce the three payoffs. | Each card needs a nearby Markdown heading and one-sentence explanation. |
| Start in 60 seconds | Markdown | No SVG card. | Commands must remain copyable text. |
| Choose your path | Markdown table | No SVG card. | Usage boundaries must remain text. |
| Demo | Markdown transcript | No SVG card. | The blocked outcome must be visible as text. |
| What ni is not | Markdown | No SVG card. | Product boundaries must remain text. |
| Read next | Markdown table or list | No SVG card. | Links must remain normal Markdown links. |

## Visual Card Copy

Card text must stay short. Do not put long explanations inside SVGs.

### Pain Cards

| Card | SVG label | Markdown fallback sentence |
| --- | --- | --- |
| Vague intent | `Vague intent` | A prompt can sound actionable while users, acceptance criteria, risks, non-goals, or blocker questions are still missing. |
| Early execution | `Early execution` | Work should not begin just because a request sounds plausible. |
| Rework | `Rework` | Hidden assumptions become expensive after humans, models, or tools start from the wrong plan. |

### ni Payoff Cards

| Card | SVG label | Markdown fallback sentence |
| --- | --- | --- |
| Capture intent | `Capture intent` | Planning conversation becomes explicit docs and a contract draft. |
| Lock contract | `Lock contract` | `ni status` and `ni end` gate readiness, hashes, and lock creation. |
| Handoff safely | `Handoff safely` | `ni run` compiles a bounded prompt or seed from a valid locked plan. |

## Mobile-Friendly Fallback

The README must still work if SVG cards do not render, wrap poorly, or are read
by a narrow mobile viewport.

Use this pattern for every visual card group:

```markdown
<p align="center">
  <img src="assets/card-pain-vague-intent.svg" alt="Vague intent: missing users, acceptance criteria, risks, non-goals, or blockers can hide inside a plausible prompt." width="32%">
</p>

### Vague intent

A prompt can sound actionable while users, acceptance criteria, risks,
non-goals, or blocker questions are still missing.
```

Rules:

- Card alt text must explain the card, not just repeat the file name.
- The same meaning must appear in Markdown directly near the card.
- Headings must remain plain Markdown headings.
- Avoid side-by-side layouts that become unreadable on mobile.
- If three cards are shown in one HTML row, the following Markdown headings and
  sentences must still make sense when the images fail.

## Visual Sales Guardrails

- Do not mention specific harness products in the hero, pain cards, or payoff
  cards.
- Codex and Claude may appear only in usage path sections, where they describe
  assisted planning UX and not kernel authority.
- Do not claim package distribution, hosted service availability, or published
  release assets unless those facts exist.
- Do not imply that `ni run` executes shell commands, queues, agents, or
  downstream work.
- Do not introduce contract authoring CLI commands.
- Keep protocol details in docs and link to them from the README.

## Required Visual Assets

Required existing assets:

| Asset | Status | Role |
| --- | --- | --- |
| `assets/hero.svg` | Existing | Hero banner. |
| `assets/badge-english.svg` | Existing | English language chip. |
| `assets/badge-korean.svg` | Existing | Korean language chip. |

Required new card assets for the next README rewrite:

| Asset | Role |
| --- | --- |
| `assets/card-pain-vague-intent.svg` | Pain card for vague intent. |
| `assets/card-pain-early-execution.svg` | Pain card for early execution. |
| `assets/card-pain-rework.svg` | Pain card for rework. |
| `assets/card-payoff-capture-intent.svg` | Payoff card for capture intent. |
| `assets/card-payoff-lock-contract.svg` | Payoff card for lock contract. |
| `assets/card-payoff-handoff-safely.svg` | Payoff card for safe handoff. |

Optional legacy or transitional assets:

| Asset | Rule |
| --- | --- |
| `assets/card-why.svg` | May remain until the pain and payoff card set replaces it. |
| `assets/card-start.svg` | Should not replace the Markdown `Start in 60 seconds` section. |
| `assets/card-docs.svg` | Should not replace the Markdown `Read next` section. |

All card assets must follow `docs/60_VISUAL_DESIGN_SPEC.md`: local SVG, short
labels only, meaningful alt text, no remote dependencies, and no essential copy
that appears only inside the image.
