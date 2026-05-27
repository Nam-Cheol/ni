# Visual Design Spec

Task 100 defines the stable visual direction for README and social assets
before any hero, banner, or card asset is generated or edited.

This document is a design contract only. It does not redesign the README,
implement an asset generator, add screenshots, add GIFs, or change runtime
behavior.

## Product Feel

`ni` should feel precise, calm, deterministic, pre-runtime, and contract-first.
Visuals should make the reader feel that work is being slowed down just enough
to become explicit, accepted, and safe to hand off.

The primary brand phrase is:

```text
Don't run the agent yet. Compile the intent first.
```

The supporting product description is:

```text
Project Intent Compiler for AI Agents.
```

Important copy should live in Markdown whenever possible. Images may reinforce
the phrase, but they must not be the only place where essential meaning appears.

## Color Palette

Use neutral ink with a soft blue, purple, and green gradient. The palette should
stay quiet and product-like, not decorative or neon.

| Token | Hex | Use |
| --- | --- | --- |
| `ink` | `#111827` | Primary text and dark rules. |
| `muted-ink` | `#475569` | Secondary labels and subdued copy. |
| `paper` | `#F8FAFC` | Light base background. |
| `line` | `#D8E1EE` | Strokes, dividers, and grid lines. |
| `blue` | `#5B8DEF` | Primary gradient stop and active accents. |
| `purple` | `#8B7CF6` | Secondary gradient stop. |
| `green` | `#4FBF9F` | Readiness, lock, or verified accents. |
| `blue-soft` | `#DCEBFF` | Soft blue panels and chips. |
| `purple-soft` | `#EDE9FE` | Soft purple panels and chips. |
| `green-soft` | `#DCFCE7` | Soft green panels and chips. |

Gradients should be deterministic and reviewable. Prefer two or three explicit
linear gradient stops over complex filters or generated bitmap effects.

## Background Style

Use a clean light background with one controlled gradient field and optional
thin geometric structure:

- base fill: `paper`;
- gradient: soft blue to purple to green, low contrast, anchored diagonally;
- structure: subtle grid, contract lines, lock rings, or bounded panels;
- shadows: minimal, soft, and not required;
- no bokeh blobs, floating orbs, noisy textures, or remote background images.

Backgrounds should remain legible when GitHub renders the asset at reduced
width. The visual should still read as intentional if all text inside the SVG is
removed.

## Typography Fallback

Use Markdown text for important copy. If short SVG text is necessary, use a
system fallback stack only:

```text
-apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif
```

SVG text must not import external fonts. Do not depend on exact font metrics for
line wrapping, alignment, or clipping. Keep SVG text short enough that fallback
font differences do not change the meaning or break the layout.

## Badge And Chip Style

Badges and chips should look like quiet contract status markers:

- rounded rectangle with `6px` to `10px` radius;
- light fill from `blue-soft`, `purple-soft`, `green-soft`, or `paper`;
- `line` stroke or no stroke;
- `ink` or `muted-ink` text;
- short labels only, such as `readiness`, `lockfile`, `prompt <= 4000`;
- no emoji, icon-font dependency, or remote capsule service for core README art.

Remote status badges that already exist in the README may remain as external
status signals, but they are not the primary visual system.

## Asset Dimensions

Core assets must declare `width`, `height`, and `viewBox`.

| Asset class | Canonical size | Aspect ratio | Notes |
| --- | ---: | --- | --- |
| README hero | `1200 x 460` | `60:23` | Matches the current README hero surface. Keep important shapes inside an `80px` safe area. |
| README card | `520 x 320` | `13:8` | Use for small visual modules only. Prefer Markdown for card copy. |
| Social card | `1200 x 630` | `40:21` | Standard preview shape for Open Graph style exports. |

Do not hand-position multiline text near edges. If an asset includes live SVG
text, test it at the canonical size and at the reduced width used in the README.

## Image Alt Text Policy

Every embedded image must have meaningful fallback alt text.

Alt text should:

- state the product name or asset role;
- include the primary phrase when the phrase appears only in the visual;
- describe the visual function, not decorative implementation details;
- stay useful if the image fails to load.

Preferred hero alt text:

```text
ni: Don't run the agent yet. Compile the intent first.
```

Preferred social card alt text:

```text
ni social card: Don't run the agent yet. Compile the intent first. Project Intent Compiler for AI Agents.
```

If the same visual is used in Korean README content, the surrounding Markdown
should carry the localized explanation even when the image itself is
language-neutral.

## Local Asset Naming Rules

All core README assets must be local files under `assets/`.

Use lowercase kebab-case names with a stable asset role:

```text
assets/hero.svg
assets/card-why.svg
assets/card-start.svg
assets/card-docs.svg
assets/social-card.svg
assets/social-card.png
```

For future language-specific or generated variants, use explicit suffixes:

```text
assets/hero.ko.svg
assets/social-card.ko.svg
assets/social-card.png
assets/social-card.ko.png
```

Rules:

- SVG source is the editable source of truth.
- PNG output is optional and should be committed only when a target platform
  requires it.
- Do not overwrite a source SVG with unreviewable generated output.
- Do not add date-stamped or random filenames for core README assets.
- Do not store primary README visuals outside `assets/`.

## Remote Image Policy

The primary README hero and core README visual assets must not rely on remote
render services. They must render from repository-local files.

Allowed:

- local SVG assets under `assets/`;
- optional locally generated PNG exports when a platform requires PNG;
- existing remote status badges when they are factual and non-essential;
- remote capsule-style examples documented as optional inspiration only. See
  `docs/64_CAPSULE_STYLE_NOTES.md`.

Not allowed for primary README visuals:

- remote hero images;
- remote dynamic SVG renderers;
- remote screenshot services;
- remote badge or capsule services used as the main hero;
- external font, CSS, script, or image dependencies inside SVG.

## Banned SVG Patterns

Do not use these patterns in core README assets:

- emoji inside SVG;
- `foreignObject`;
- external font imports;
- external CSS, scripts, or images;
- long paragraphs inside SVG;
- important copy that exists only inside an image;
- Korean text inside SVG unless converted or tested carefully;
- hand-positioned multiline text without reduced-width rendering tests;
- font-specific layout that clips when fallback fonts render;
- filters or masks that make the source hard to review;
- generated SVG with unstable IDs, random coordinates, or non-deterministic
  metadata.

Short SVG labels are acceptable only when the layout is tested and the same
meaning is available in Markdown or alt text.

## Readiness For Asset Work

Before editing or generating README assets, check this spec and the asset audit:

- `docs/59_VISUAL_ASSET_AUDIT.md`
- `docs/60_VISUAL_DESIGN_SPEC.md`
- `docs/64_CAPSULE_STYLE_NOTES.md`

Asset work should preserve the kernel boundary: visual changes may improve
communication, but they must not imply runtime execution, downstream adapter
behavior, package distribution, or product claims that are not implemented.
