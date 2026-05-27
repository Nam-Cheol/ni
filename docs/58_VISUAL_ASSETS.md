# Visual Assets

This repository keeps visual assets lightweight and source-first. The README
hero remains `assets/hero.svg`; it should not be replaced by a large image-only
surface.

## Current Assets

| Asset | Role | Source of truth | Notes |
| --- | --- | --- | --- |
| `assets/hero.svg` | README hero | `assets/source/hero.template.svg` | Lightweight inline-shareable repository asset. |
| `assets/social-card.svg` | Optional social card | `assets/source/social-card.template.svg` | Generated source for social previews and marketing experiments. |
| `assets/social-card.png` | Optional generated social card | `assets/social-card.svg` | May be generated locally when a platform needs PNG upload. Do not commit a large PNG unless necessary. |

## Rules

- Keep the README hero as SVG.
- Treat generated images as optional marketing assets, not kernel behavior.
- Use `assets/source/social-card.template.svg` and `scripts/render-assets.py`
  as the editable source of truth for any social card SVG export.
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
ni — Project Intent Compiler for AI Agents.
```

Suggested alt text:

```text
ni social card: Don't run the agent yet. Compile the intent first. ni — Project Intent Compiler for AI Agents.
```

## PNG Export

PNG export is optional. If a social platform requires PNG, generate it locally
from `assets/social-card.svg` and check the resulting file size before
committing it.
