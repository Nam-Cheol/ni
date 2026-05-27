# Visual Asset Audit

Task 099는 README redesign 전에 visual asset surface를 audit한다. 이 문서는
asset을 redesign하지 않고, raster export를 추가하지 않고, remote image
dependency를 추가하지 않고, product behavior를 바꾸지 않고, 새로운 install path를
claim하지 않는다.

## 확인 범위

- `assets/`
- `README.md`
- `README.ko.md`
- `docs/` 아래 Markdown files
- `README 2.md`, `README.ko 2.md` 같은 root README duplicate files

이번 audit 시점에 root duplicate README files는 없었다.

## README Rendering Surface

| Surface | Image references | Notes |
| --- | --- | --- |
| `README.md` | `assets/hero.svg`, `assets/card-why.svg`, `assets/card-start.svg`, `assets/card-docs.svg`, shields.io badges 4개 | Local SVGs는 HTML `<img>` tags로 rendering된다. Badges는 이번 audit 전부터 있던 remote dependencies다. |
| `README.ko.md` | `assets/hero.svg`, `assets/card-why.svg`, `assets/card-start.svg`, `assets/card-docs.svg`, shields.io badges 4개 | English README와 같은 local 및 remote image surface다. Local SVG text는 주변 README copy가 Korean이어도 English로 남아 있다. |
| `docs/58_VISUAL_ASSETS.md` | `assets/hero.svg`, `assets/social-card.svg`, optional `assets/social-card.png` mention | Embedded image는 없다. `assets/social-card.png`는 optional generated output으로만 문서화되어 있고 repository에는 없다. |
| Other `docs/**/*.md` files | None found | `docs/58_VISUAL_ASSETS.md` 외에는 Markdown image embed나 direct SVG/PNG/JPG/GIF/WebP reference가 없었다. |

## Asset Inventory

| Asset | Size | Used by | Contains text | Emoji | `foreignObject` | Dimensions | GitHub README rendering safety | Recommended action |
| --- | ---: | --- | --- | --- | --- | --- | --- | --- |
| `assets/hero.svg` | 4,425 bytes | `README.md`, `README.ko.md`, `docs/58_VISUAL_ASSETS.md` mention | Yes: headline, subcopy, labels | No | No | `width="1200"`, `height="460"`, `viewBox="0 0 1200 460"` | Caution: SVG image로서 구조적으로는 안전하지만, live text가 font fallback과 fixed coordinates에 의존한다. | 지금은 유지한다. README relaunch 전에는 SVG에서 text를 제거하거나, 승인된 copy를 path로 변환하거나, checked source에서 controlled asset을 생성한다. |
| `assets/card-why.svg` | 1,584 bytes | `README.md`, `README.ko.md` | Yes: card title, body, callout | No | No | `width="520"`, `height="320"`, `viewBox="0 0 520 320"` | Caution: 구조적으로는 안전하지만, body copy가 fixed live text라 renderer마다 wrap 또는 clipping risk가 있다. | 지금은 유지한다. README relaunch 전에는 card를 text-free로 만들거나 승인된 final copy를 path로 변환한다. |
| `assets/card-start.svg` | 1,804 bytes | `README.md`, `README.ko.md` | Yes: title, command labels, note | No | No | `width="520"`, `height="320"`, `viewBox="0 0 520 320"` | Caution: 구조적으로는 안전하지만, command labels가 live text이고 fixed coordinates에 묶여 있다. | 지금은 유지한다. README relaunch 전에는 command labels를 image가 아니라 README text에 둘지 결정한다. |
| `assets/card-docs.svg` | 1,436 bytes | `README.md`, `README.ko.md` | Yes: title, document labels, callout | No | No | `width="520"`, `height="320"`, `viewBox="0 0 520 320"` | Caution: 구조적으로는 안전하지만, documentation labels가 live text이고 Korean README에서도 English-only다. | 지금은 유지한다. README relaunch 전에는 language-specific variants를 만들거나 labels를 image 밖으로 옮긴다. |
| `assets/social-card.svg` | 4,365 bytes | `docs/58_VISUAL_ASSETS.md` mention only | Yes: headline, tagline, labels | No | No | `width="1200"`, `height="630"`, `viewBox="0 0 1200 630"` | README에서는 현재 사용하지 않는다. SVG consumer에는 구조적으로 안전하지만, social platforms는 PNG export를 요구할 수 있고 live text는 font-dependent다. | Editable source로만 유지한다. Platform이 요구하고 export를 검토하기 전에는 PNG를 commit하지 않는다. |

## Findings

- Committed visual asset은 모두 SVG다.
- Committed SVG는 모두 live `<text>` nodes를 포함한다.
- Committed SVG는 emoji를 사용하지 않는다.
- Committed SVG는 `foreignObject`를 사용하지 않는다.
- Committed SVG는 모두 `width`, `height`, `viewBox`를 가진다.
- README와 README.ko는 같은 English-text SVG assets를 사용한다.
- README와 README.ko에는 기존 shields.io remote badge images가 있다.
- Committed PNG, JPG, JPEG, GIF, WebP files는 없었다.
- `assets/social-card.png`는 optional generated output으로 문서화되어 있지만
  현재 존재하지 않는다.

## README Relaunch 전 Visual Risks

1. Live SVG text는 `Inter`를 요청하지만 fallback fonts에 의존하므로 GitHub,
   local browsers, social previews, high-DPI screenshots에서 다르게 보일 수
   있다.
2. Fixed-coordinate SVG text에는 wrapping model이 없다. Copy edit, font
   fallback, translation 때문에 clipping이나 어색한 spacing이 생길 수 있다.
3. Korean README가 English SVG copy를 embed한다. Bilingual README redesign에는
   language-neutral visuals 또는 explicit language-specific assets가 필요하다.
4. Remote shields.io badges는 README rendering surface의 network dependencies다.
   기존 badge이지만 controlled visual system에서 계속 허용할지 결정해야 한다.
5. `docs/58_VISUAL_ASSETS.md`는 존재하지 않는 optional PNG export path를
   언급한다. 지금은 괜찮지만 README relaunch에서는 committed raster asset이
   있는 것처럼 보이지 않게 해야 한다.

## Repair Plan

1. Redesign 전에 asset classes를 정의한다: README hero, README navigation
   cards, badges, social preview.
2. README visuals가 text를 포함할 수 있는지 결정한다. 기본값은 nearby
   accessible README copy를 둔 text-free visuals다. Text가 필요하면 final copy를
   freeze한 뒤 path로 변환하거나 controlled source에서 생성한다.
3. README 변경 전에 bilingual local asset rule을 만든다. English가 intentional
   product vocabulary가 아니라면 Korean-first surface에서 English text image를
   재사용하지 않는다.
4. Controlled asset set에서 `foreignObject`, emoji-dependent rendering, remote
   README art, unreviewed generated raster images를 제외한다.
5. 이후 lightweight asset check를 추가해 SVG text nodes, `foreignObject`,
   missing dimensions, remote README image references를 보고하게 한다.
