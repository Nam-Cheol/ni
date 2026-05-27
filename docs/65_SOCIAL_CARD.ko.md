# Social Card

`assets/social-card.svg`는 repository previews, release posts, lightweight
marketing surfaces를 위한 optional share card다. 이 asset은 README hero를
대체하지 않으며 `ni-kernel` behavior도 바꾸지 않는다.

## Asset

| File | Role | Source |
| --- | --- | --- |
| `assets/social-card.svg` | Optional social preview card | `assets/source/social-card.template.svg`를 `scripts/render-assets.py`로 render한다. |
| `assets/social-card.png` | Optional platform upload export | v1에서는 commit하지 않는다. Platform이 PNG를 요구하고 file size를 검토한 경우에만 생성한다. |

## Copy

v1 card copy는 English-only다:

```text
Don't run the agent yet.
Compile the intent first.
ni — Project Intent Compiler for AI Agents.
```

Localized variant를 별도로 design하고 validate하기 전까지
`assets/social-card.svg`에는 Korean text를 넣지 않는다.

## Generation

Repository root에서 deterministic SVG renderer를 실행한다:

```bash
python3 scripts/render-assets.py
```

그다음 asset checker를 실행한다:

```bash
python3 scripts/check-assets.py
```

Checker는 SVG structure, size limit, local-only references, emoji exclusion,
fresh render 대비 deterministic drift를 검증한다.

## README Boundary

README hero는 `assets/hero.svg`로 유지한다. Social card는 external preview
surfaces를 위한 optional asset이며, replacement hero도 아니고 Project Intent
Compiler positioning을 넘어서는 product claim도 아니다.
