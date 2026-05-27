# Visual Asset Audit

Task 099 audits the visual asset surface before any README redesign. It does
not redesign assets, add raster exports, add remote image dependencies, change
product behavior, or claim new install paths.

## Scope Checked

- `assets/`
- `README.md`
- `README.ko.md`
- Markdown files under `docs/`
- root README duplicates such as `README 2.md` and `README.ko 2.md`

No root duplicate README files were present during this audit.

## README Rendering Surface

| Surface | Image references | Notes |
| --- | --- | --- |
| `README.md` | `assets/hero.svg`, `assets/card-why.svg`, `assets/card-start.svg`, `assets/card-docs.svg`, four shields.io badges | Local SVGs are rendered through HTML `<img>` tags. Badges are remote dependencies already present before this audit. |
| `README.ko.md` | `assets/hero.svg`, `assets/card-why.svg`, `assets/card-start.svg`, `assets/card-docs.svg`, four shields.io badges | Same local and remote image surface as English README. Local SVG text remains English while surrounding README copy is Korean. |
| `docs/58_VISUAL_ASSETS.md` | Mentions `assets/hero.svg`, `assets/social-card.svg`, and optional `assets/social-card.png` | No embedded image. `assets/social-card.png` is documented as optional generated output and is not present in the repository. |
| Other `docs/**/*.md` files | None found | No Markdown image embeds or direct SVG/PNG/JPG/GIF/WebP references found outside `docs/58_VISUAL_ASSETS.md`. |

## Asset Inventory

| Asset | Size | Used by | Contains text | Emoji | `foreignObject` | Dimensions | GitHub README rendering safety | Recommended action |
| --- | ---: | --- | --- | --- | --- | --- | --- | --- |
| `assets/hero.svg` | 4,425 bytes | `README.md`, `README.ko.md`, `docs/58_VISUAL_ASSETS.md` mention | Yes: headline, subcopy, labels | No | No | `width="1200"`, `height="460"`, `viewBox="0 0 1200 460"` | Caution: structurally safe as an SVG image, but live text depends on font fallback and fixed coordinates. | Keep for now. Before README relaunch, either remove text from the SVG, convert approved text to paths, or replace with a controlled asset generated from a checked source. |
| `assets/card-why.svg` | 1,584 bytes | `README.md`, `README.ko.md` | Yes: card title, body, callout | No | No | `width="520"`, `height="320"`, `viewBox="0 0 520 320"` | Caution: structurally safe, but body copy is fixed live text and may wrap or clip differently across renderers. | Keep for now. Before README relaunch, make the card text-free or convert approved final copy to paths. |
| `assets/card-start.svg` | 1,804 bytes | `README.md`, `README.ko.md` | Yes: title, command labels, note | No | No | `width="520"`, `height="320"`, `viewBox="0 0 520 320"` | Caution: structurally safe, but command labels are live text and tied to fixed coordinates. | Keep for now. Before README relaunch, decide whether command labels belong in README text instead of the image. |
| `assets/card-docs.svg` | 1,436 bytes | `README.md`, `README.ko.md` | Yes: title, document labels, callout | No | No | `width="520"`, `height="320"`, `viewBox="0 0 520 320"` | Caution: structurally safe, but documentation labels are live text and English-only in the Korean README. | Keep for now. Before README relaunch, make language-specific variants or move labels out of the image. |
| `assets/social-card.svg` | 4,365 bytes | `docs/58_VISUAL_ASSETS.md` mention only | Yes: headline, tagline, labels | No | No | `width="1200"`, `height="630"`, `viewBox="0 0 1200 630"` | Not currently in the README. Structurally safe for SVG consumers, but social platforms often require PNG export and live text remains font-dependent. | Keep as editable source only. Do not commit a PNG until a platform requires it and the export is checked. |

## Findings

- Every committed visual asset is SVG.
- Every committed SVG contains live `<text>` nodes.
- No committed SVG uses emoji.
- No committed SVG uses `foreignObject`.
- Every committed SVG has `width`, `height`, and `viewBox`.
- README and README.ko use the same English-text SVG assets.
- README and README.ko include pre-existing remote shields.io badge images.
- No committed PNG, JPG, JPEG, GIF, or WebP files were found.
- `assets/social-card.png` is documented as optional generated output, but it is
  not present.

## Visual Risks Before README Relaunch

1. Live SVG text can render differently on GitHub, local browsers, social
   previews, and high-DPI screenshots because the assets request `Inter` but
   rely on fallback fonts.
2. Fixed-coordinate SVG text has no wrapping model. A copy edit, font fallback,
   or translation can cause clipping or awkward spacing.
3. The Korean README embeds English SVG copy, so a redesigned bilingual README
   needs either language-neutral visuals or explicit language-specific assets.
4. Remote shields.io badges are network dependencies in the README rendering
   surface. They are existing badges, but a controlled visual system should
   decide whether they remain acceptable.
5. `docs/58_VISUAL_ASSETS.md` names an optional PNG export path that does not
   exist. That is acceptable today, but the README relaunch should avoid
   implying a committed raster asset unless one is actually added.

## Repair Plan

1. Define asset classes before redesign: README hero, README navigation cards,
   badges, and social preview.
2. Decide whether README visuals may contain text. Prefer text-free visuals with
   accessible README copy nearby; if text is required, freeze final copy and
   convert it to paths or generate from a controlled source.
3. Create bilingual rules for local assets before changing README. Do not reuse
   English text images in Korean-first surfaces unless the English is
   intentional product vocabulary.
4. Keep `foreignObject`, emoji-dependent rendering, remote README art, and
   unreviewed generated raster images out of the controlled asset set.
5. Add a lightweight asset check later that reports SVG text nodes,
   `foreignObject`, missing dimensions, and remote README image references.
