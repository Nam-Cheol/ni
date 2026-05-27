# Visual Assets

This repository keeps visual assets lightweight and source-first. The README
hero remains `assets/hero.svg`; it should not be replaced by a large image-only
surface.

## Current Assets

| Asset | Role | Source of truth | Notes |
| --- | --- | --- | --- |
| `assets/hero.svg` | README hero | SVG | Lightweight inline-shareable repository asset. |
| `assets/social-card.svg` | Optional social card experiment | SVG | Editable source for social previews and marketing experiments. |
| `assets/social-card.png` | Optional generated social card | `assets/social-card.svg` | May be generated locally when a platform needs PNG upload. Do not commit a large PNG unless necessary. |

## Rules

- Keep the README hero as SVG.
- Treat generated images as optional marketing assets, not kernel behavior.
- Use `assets/social-card.svg` as the editable source of truth for any social
  card export.
- Do not use remote images for repository visual assets.
- Do not add false product claims to image copy.
- Keep file sizes small; prefer SVG unless a platform requires PNG.
- Provide meaningful alt text whenever embedding a visual asset.

## Social Card Copy

Primary copy:

```text
Don't run the agent yet.
Compile the intent first.
```

Subcopy:

```text
Project Intent Compiler for AI Agents.
```

Suggested alt text:

```text
ni social card: Don't run the agent yet. Compile the intent first. Project Intent Compiler for AI Agents.
```

## PNG Export

PNG export is optional. If a social platform requires PNG, generate it locally
from `assets/social-card.svg` and check the resulting file size before
committing it.
