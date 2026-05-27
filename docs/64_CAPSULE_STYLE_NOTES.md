# Capsule Style Notes

Task 108 records how `ni` treats capsule-style README banners as visual
inspiration without adding a remote dependency to the primary README surface.

This document is a visual policy note only. It does not change README assets,
product behavior, the asset generator, or the kernel boundary.

## Why Capsule-Style Banners Are Useful

Capsule-style banners can be useful inspiration for README first impressions.
They are compact, familiar to GitHub readers, and good at presenting a product
name, short phrase, and quiet visual mood in one scan-friendly image.

For `ni`, that style maps well to the existing visual direction:

- a centered product identity;
- restrained gradients;
- short text;
- a single banner-shaped surface;
- a polished first-screen README rhythm.

The style can inform local template composition, spacing, and contrast, but it
is not the source of truth for `ni` README assets.

## Remote Render Convenience

Remote query-rendered SVG services are convenient because a README author can
describe a banner with URL parameters instead of committing an asset file. An
example pattern, provided only as a reference, looks like this:

```text
https://capsule-render.vercel.app/api?type=waving&text=ni&fontSize=64
```

In a real service URL, query text and special characters need to be escaped, and
the rendered result depends on the service's parameter model. This repository
may use such patterns for comparison or inspiration in documentation, but not as
the primary README hero.

## Local Asset Policy

Primary README visuals for `ni` must be local, deterministic, and reviewable.

Use local assets for core README art because they give this repository direct
control over:

- the exact SVG source committed to the repo;
- text escaping and line breaks;
- font fallback behavior;
- dimensions, safe areas, and reduced-width rendering;
- availability when GitHub, offline readers, or mirrors render the README;
- review diffs for visual changes.

Remote query-rendered SVGs can be useful for quick experiments, but text
escaping, font fallback, service parameters, and service availability can affect
the rendered output. The local SVG template pipeline is more controllable for
this repository.

## ni Decision

`ni` uses local SVG templates for core README assets.

The editable source of truth lives in `assets/source/` and
`scripts/render-assets.py`. Generated SVG files under `assets/` are committed so
the README can render without fetching a remote hero, remote capsule service,
external font, external CSS, script, or remote image dependency.

Allowed:

- local SVG templates and generated local SVG files;
- optional locally generated PNG exports when a platform requires PNG;
- non-essential factual remote status badges;
- remote capsule-style URL patterns documented as optional inspiration.

Not allowed for primary README visuals:

- replacing `assets/hero.svg` with a remote capsule-render URL;
- adding a remote dynamic SVG renderer as the main hero;
- depending on external fonts, CSS, scripts, or images inside SVG assets;
- making essential README meaning available only through a rendered image.

See also:

- `docs/60_VISUAL_DESIGN_SPEC.md`
- `docs/61_ASSET_PIPELINE.md`
- `docs/63_README_VISUAL_WIREFRAME.md`
