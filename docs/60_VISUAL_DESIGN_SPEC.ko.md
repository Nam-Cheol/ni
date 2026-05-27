# Visual Design Spec

Task 100은 README와 social assets의 hero, banner, card를 생성하거나 수정하기
전에 안정적인 visual direction을 정의한다.

이 문서는 design contract only다. README를 redesign하지 않고, asset generator를
구현하지 않고, screenshots나 GIFs를 추가하지 않고, runtime behavior를 바꾸지
않는다.

## Product Feel

`ni`는 precise, calm, deterministic, pre-runtime, contract-first 느낌이어야
한다. Visual은 work를 바로 실행하기 전에 잠시 멈추고, intent를 explicit,
accepted, handoff-safe 상태로 만든다는 인상을 줘야 한다.

Primary brand phrase:

```text
Don't run the agent yet. Compile the intent first.
```

Supporting product description:

```text
Project Intent Compiler for AI Agents.
```

Important copy는 가능한 한 Markdown에 둔다. Image는 phrase를 reinforce할 수
있지만, essential meaning이 image 안에만 존재하면 안 된다.

## Color Palette

Neutral ink와 soft blue, purple, green gradient를 사용한다. Palette는 quiet
product feel이어야 하며 decorative 또는 neon 느낌이면 안 된다.

| Token | Hex | Use |
| --- | --- | --- |
| `ink` | `#111827` | Primary text와 dark rules. |
| `muted-ink` | `#475569` | Secondary labels와 subdued copy. |
| `paper` | `#F8FAFC` | Light base background. |
| `line` | `#D8E1EE` | Strokes, dividers, grid lines. |
| `blue` | `#5B8DEF` | Primary gradient stop과 active accents. |
| `purple` | `#8B7CF6` | Secondary gradient stop. |
| `green` | `#4FBF9F` | Readiness, lock, verified accents. |
| `blue-soft` | `#DCEBFF` | Soft blue panels와 chips. |
| `purple-soft` | `#EDE9FE` | Soft purple panels와 chips. |
| `green-soft` | `#DCFCE7` | Soft green panels와 chips. |

Gradients는 deterministic and reviewable해야 한다. Complex filters나 generated
bitmap effects보다 명시적인 two or three stop linear gradient를 선호한다.

## Background Style

Clean light background에 controlled gradient field 하나와 optional thin
geometric structure를 사용한다:

- base fill: `paper`;
- gradient: soft blue to purple to green, low contrast, diagonal anchor;
- structure: subtle grid, contract lines, lock rings, bounded panels;
- shadows: minimal, soft, and not required;
- bokeh blobs, floating orbs, noisy textures, remote background images 금지.

GitHub가 asset을 작은 width로 render해도 background는 읽혀야 한다. SVG 안의
text를 모두 제거해도 visual이 intentional하게 보여야 한다.

## Typography Fallback

Important copy는 Markdown text로 둔다. Short SVG text가 필요하다면 system
fallback stack만 사용한다:

```text
-apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif
```

SVG text는 external fonts를 import하면 안 된다. Line wrapping, alignment,
clipping이 exact font metrics에 의존하면 안 된다. Fallback font 차이가 meaning
또는 layout을 깨지 않을 만큼 짧은 SVG text만 허용한다.

## Badge And Chip Style

Badges와 chips는 quiet contract status markers처럼 보여야 한다:

- `6px` to `10px` radius의 rounded rectangle;
- `blue-soft`, `purple-soft`, `green-soft`, 또는 `paper` light fill;
- `line` stroke 또는 no stroke;
- `ink` 또는 `muted-ink` text;
- `readiness`, `lockfile`, `prompt <= 4000` 같은 short labels only;
- core README art에는 emoji, icon-font dependency, remote capsule service 금지.

README에 이미 존재하는 remote status badges는 external status signals로 남을 수
있지만 primary visual system은 아니다.

## Asset Dimensions

Core assets는 `width`, `height`, `viewBox`를 선언해야 한다.

| Asset class | Canonical size | Aspect ratio | Notes |
| --- | ---: | --- | --- |
| README hero | `1200 x 460` | `60:23` | Current README hero surface와 맞춘다. Important shapes는 `80px` safe area 안에 둔다. |
| README card | `520 x 320` | `13:8` | Small visual modules에만 사용한다. Card copy는 Markdown을 선호한다. |
| Social card | `1200 x 630` | `40:21` | Open Graph style exports의 standard preview shape. |

Multiline text를 edge 근처에 손으로 배치하지 않는다. Asset이 live SVG text를
포함한다면 canonical size와 README reduced width에서 모두 test한다.

## Image Alt Text Policy

모든 embedded image는 meaningful fallback alt text를 가져야 한다.

Alt text는 다음을 지켜야 한다:

- product name 또는 asset role을 말한다;
- phrase가 visual 안에만 나타난다면 primary phrase를 포함한다;
- decorative implementation details가 아니라 visual function을 설명한다;
- image load가 실패해도 유용해야 한다.

Preferred hero alt text:

```text
ni: Don't run the agent yet. Compile the intent first.
```

Preferred social card alt text:

```text
ni social card: Don't run the agent yet. Compile the intent first. Project Intent Compiler for AI Agents.
```

같은 visual을 Korean README content에서 사용한다면 image가 language-neutral이어도
surrounding Markdown이 localized explanation을 제공해야 한다.

## Local Asset Naming Rules

Core README assets는 모두 `assets/` 아래 local files여야 한다.

Stable asset role을 포함한 lowercase kebab-case names를 사용한다:

```text
assets/hero.svg
assets/card-why.svg
assets/card-start.svg
assets/card-docs.svg
assets/social-card.svg
assets/social-card.png
```

Future language-specific 또는 generated variants에는 explicit suffix를 사용한다:

```text
assets/hero.ko.svg
assets/social-card.ko.svg
assets/social-card.png
assets/social-card.ko.png
```

Rules:

- SVG source가 editable source of truth다.
- PNG output은 optional이며 target platform이 PNG를 요구할 때만 commit한다.
- Source SVG를 unreviewable generated output으로 덮어쓰지 않는다.
- Core README assets에는 date-stamped 또는 random filenames를 추가하지 않는다.
- Primary README visuals를 `assets/` 밖에 두지 않는다.

## Remote Image Policy

Primary README hero와 core README visual assets는 remote render services에
의존하면 안 된다. Repository-local files에서 render되어야 한다.

Allowed:

- `assets/` 아래 local SVG assets;
- platform이 PNG를 요구할 때 optional locally generated PNG exports;
- factual and non-essential 기존 remote status badges;
- optional inspiration으로 문서화된 remote capsule-style examples only.

Primary README visuals에서 금지:

- remote hero images;
- remote dynamic SVG renderers;
- remote screenshot services;
- main hero로 사용하는 remote badge 또는 capsule services;
- SVG 안의 external font, CSS, script, image dependencies.

## Banned SVG Patterns

Core README assets에서 다음 patterns를 사용하지 않는다:

- emoji inside SVG;
- `foreignObject`;
- external font imports;
- external CSS, scripts, images;
- long paragraphs inside SVG;
- important copy that exists only inside an image;
- Korean text inside SVG unless converted or tested carefully;
- hand-positioned multiline text without reduced-width rendering tests;
- font-specific layout that clips when fallback fonts render;
- filters or masks that make the source hard to review;
- generated SVG with unstable IDs, random coordinates, or non-deterministic
  metadata.

Short SVG labels는 layout이 tested이고 같은 meaning이 Markdown 또는 alt text에
있을 때만 허용한다.

## Readiness For Asset Work

README assets를 edit하거나 generate하기 전에 이 spec과 asset audit를 확인한다:

- `docs/59_VISUAL_ASSET_AUDIT.md`
- `docs/60_VISUAL_DESIGN_SPEC.md`

Asset work는 kernel boundary를 보존해야 한다. Visual changes는 communication을
개선할 수 있지만 runtime execution, downstream adapter behavior, package
distribution, 구현되지 않은 product claims를 암시하면 안 된다.
