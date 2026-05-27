# Social Card

`assets/social-card.svg` is an optional share card for repository previews,
release posts, and other lightweight marketing surfaces. It does not replace
the README hero, and it does not change `ni-kernel` behavior.

## Asset

| File | Role | Source |
| --- | --- | --- |
| `assets/social-card.svg` | Optional social preview card | `assets/source/social-card.template.svg` rendered by `scripts/render-assets.py`. |
| `assets/social-card.png` | Optional platform upload export | Not committed for v1. Generate only when a platform requires PNG and the file size is reviewed. |

## Copy

The v1 card copy is English-only:

```text
Don't run the agent yet.
Compile the intent first.
ni — Project Intent Compiler for AI Agents.
```

Do not add Korean text to `assets/social-card.svg` for v1 unless a localized
variant is separately designed and validated.

## Generation

Run the deterministic SVG renderer from the repository root:

```bash
python3 scripts/render-assets.py
```

Then run the asset checker:

```bash
python3 scripts/check-assets.py
```

The checker validates the SVG structure, size limit, local-only references,
emoji exclusion, and deterministic drift against a fresh render.

## README Boundary

The README hero remains `assets/hero.svg`. The social card is an optional
asset for external preview surfaces, not a replacement hero and not a product
claim beyond the Project Intent Compiler positioning.
