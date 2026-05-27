# Asset Pipeline

Task 101은 core README assets를 위한 deterministic local SVG pipeline을
추가한다. 목표는 visual assets를 simple templates에서 생성해 작고 review 가능한
상태로 유지하고, model-authored freeform SVG drift를 막는 것이다.

## Source Files

Asset source of truth는 `assets/source/` 아래에 둔다:

| File | Role |
| --- | --- |
| `assets/source/hero.template.svg` | README hero template. |
| `assets/source/card.template.svg` | Small README cards shared template. |
| `scripts/render-assets.py` | Generated SVG output을 만드는 deterministic renderer. |
| `scripts/check-assets.py` | SVG structural checks와 drift checks. |

Renderer는 local files와 Python standard library만 사용한다. Remote assets를
fetch하지 않고, external fonts를 import하지 않고, render service를 호출하지
않고, PNG output을 만들지 않는다.

## Generated Files

Repository root에서 renderer를 실행한다:

```bash
python3 scripts/render-assets.py
```

Renderer는 다음 generated files를 쓴다:

```text
assets/hero.svg
assets/card-start.svg
assets/card-contract.svg
assets/card-handoff.svg
```

Generated files는 GitHub README rendering이 `assets/`의 SVG를 직접 소비하기
때문에 commit한다. 이 파일들의 editable source of truth는 templates와
renderer다.

## Checks

Repository root에서 asset checks를 실행한다:

```bash
python3 scripts/check-assets.py
```

Checker는 다음을 검증한다:

- required generated assets가 존재한다;
- top-level `assets/*.svg` files가 XML로 parse된다;
- checked SVG마다 `viewBox`가 있다;
- checked SVG마다 `width`와 `height`가 있다;
- `foreignObject`가 없다;
- emoji 및 emoji-style codepoints가 없다;
- external `href` 및 `url(http...)` references가 없다;
- file size가 repository limit 아래에 있다;
- generated outputs가 fresh deterministic render와 일치한다.

`scripts/quality.sh`는 `scripts/check-assets.py`를 실행하므로 stale generated
SVG는 normal quality gate에서 실패한다.

## Authoring Rules

- `assets/source/*.template.svg` 또는 `scripts/render-assets.py`를 수정한 뒤
  renderer를 실행한다.
- SVG 안의 text는 짧게 유지하고 nearby Markdown이나 alt text로 의미를 반복한다.
- `rect`, `path`, `circle`, `linearGradient`, tested short `text` 같은 simple
  SVG primitives를 사용한다.
- `foreignObject`, emoji, remote image references, external CSS, external font
  imports를 추가하지 않는다.
- 아직 이 pipeline에 binary PNG output을 추가하지 않는다.
- 이 pipeline으로 runtime execution, adapters, queues, downstream-owned state를
  암시하지 않는다.

## README Boundary

이 pipeline은 assets만 관리한다. README layout changes는 별도 task로 유지해야
한다. README가 이 문서로 link할 수는 있지만, asset generation 자체가 product
claims나 kernel boundary를 바꾸면 안 된다.
