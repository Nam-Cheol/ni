# Asset Pipeline

Task 101 adds a deterministic local SVG pipeline for core README assets. The
goal is to keep visual assets small, reviewable, and generated from simple
templates instead of freeform model-authored SVG.

## Source Files

The asset source of truth is under `assets/source/`:

| File | Role |
| --- | --- |
| `assets/source/hero.template.svg` | Template for the README hero. |
| `assets/source/card.template.svg` | Shared template for small README cards. |
| `assets/source/badge.template.svg` | Shared template for local README language chips. |
| `scripts/render-assets.py` | Deterministic renderer for generated SVG output. |
| `scripts/check-assets.py` | Structural and drift checks for SVG assets. |

The renderer uses only local files and Python standard library modules. It does
not fetch remote assets, import external fonts, call a render service, or create
PNG output.

## Generated Files

Run the renderer from the repository root:

```bash
python3 scripts/render-assets.py
```

It writes these generated files:

```text
assets/hero.svg
assets/badge-english.svg
assets/badge-korean.svg
assets/card-start.svg
assets/card-contract.svg
assets/card-handoff.svg
```

The generated files are committed because GitHub README rendering consumes SVG
directly from `assets/`. Treat the templates and renderer as the editable source
of truth for these files.

## Checks

Run the asset checks from the repository root:

```bash
python3 scripts/check-assets.py
```

The checker verifies:

- required generated assets exist;
- every top-level `assets/*.svg` file parses as XML;
- every checked SVG has a `viewBox`;
- every checked SVG has `width` and `height`;
- `foreignObject` is absent;
- emoji and emoji-style codepoints are absent;
- external `href` and `url(http...)` references are absent;
- file size stays under the repository limit;
- generated outputs match a fresh deterministic render.

`scripts/quality.sh` runs `scripts/check-assets.py`, so stale generated SVGs
fail the normal quality gate.

## Authoring Rules

- Edit `assets/source/*.template.svg` or `scripts/render-assets.py`, then run
  the renderer.
- Keep text inside SVG short and duplicated by nearby Markdown or alt text.
- Use simple SVG primitives such as `rect`, `path`, `circle`,
  `linearGradient`, and tested short `text`.
- Do not add `foreignObject`, emoji, remote image references, external CSS, or
  external font imports.
- Do not add binary PNG output in this pipeline yet.
- Do not use this pipeline to imply runtime execution, adapters, queues, or
  downstream-owned state.

## README Boundary

This pipeline controls assets only. README layout changes should remain a
separate task. A README may link to this document, but asset generation itself
must not change product claims or the kernel boundary.
