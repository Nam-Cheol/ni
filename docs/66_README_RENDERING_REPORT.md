# README Rendering Report

Task 112 validates `README.md` and `README.ko.md` as public GitHub product
pamphlets. This report records rendering facts only. It does not add product
claims, install paths, release-binary claims, or runtime behavior.

## Hero Rendering

The top README surface now renders in this order in both languages:

1. Hero image: `assets/hero.svg`
2. Language chips: `assets/badge-english.svg`, `assets/badge-korean.svg`
3. Factual trust badges: MIT license, CI workflow, security policy, docs index
4. Slogan
5. One-line product description
6. Why ni
7. Start in 60 seconds

The duplicate visible trust-signal text that appeared between the badge row and
the slogan was removed. The slogan and one-line product description remain
Markdown/HTML text, not image-only content.

## Badge Rendering

| Badge group | Rendering check | Result |
| --- | --- | --- |
| Language chips | Local SVG images linked to `README.md` and `README.ko.md`. | Pass |
| Trust badges | Remote shields.io badge images linked to existing repository files. | Pass |
| Badge facts | Badges point to `LICENSE`, `.github/workflows/ci.yml`, `SECURITY.md`, and `docs/00_START_HERE.md`. | Pass |

## Table Rendering

Both README files contain two Markdown tables:

| README | Table | Columns | Result |
| --- | --- | ---: | --- |
| `README.md` | Choose your path | 4 | Pass |
| `README.md` | Read next | 2 | Pass |
| `README.ko.md` | Choose your path | 4 | Pass |
| `README.ko.md` | Read next | 2 | Pass |

The surface checker now validates body-row column counts, not only separator
rows.

## Code Block Rendering

Both README files use balanced fenced code blocks with info strings:

| README | Fence languages | Result |
| --- | --- | --- |
| `README.md` | `bash`, `bash`, `bash`, `text` | Pass |
| `README.ko.md` | `bash`, `bash`, `bash`, `text` | Pass |

## Link Check Summary

All local README asset references exist. All local README links point to
existing files or directories.

| Reference class | Result |
| --- | --- |
| Local assets | Pass |
| HTML `href` links | Pass |
| Markdown links | Pass |
| Root duplicate README files | Pass: none present |

## English/Korean Parity

The English and Korean README files share the same public structure:

1. Hero and badges
2. Why ni
3. What ni gives you
4. Start in 60 seconds
5. Choose your path
6. Demo
7. What ni is not
8. Read next

The Korean README keeps Korean companion links where maintained, including
`docs/product-story.ko.md`, `docs/no-terminal.ko.md`,
`docs/commands.ko.md`, and `docs/63_README_VISUAL_WIREFRAME.ko.md`.

## Product Claim Audit

The hero and sales pitch do not mention specific harness or runtime products.
`Codex` and `Claude` appear only in the `Model workspaces` usage path, where
they describe assisted planning and not kernel authority.

The README continues to state that `ni run` compiles a bounded prompt or seed
and does not execute shell commands, queues, agents, or downstream work.

## Remaining Visual Issues

No blocking rendering issues remain for the current Markdown surface.

Known non-blocking visual constraints remain:

- The trust badges depend on remote shields.io image rendering.
- The local SVGs contain live text and rely on renderer font fallback.
- The Korean README intentionally reuses English-text SVG assets for this
  version.

