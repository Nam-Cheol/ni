# Badge System

Task 103은 README 상단에 language links와 factual trust signals를 위한 작은
badge system을 추가한다. Badge area는 factual해야 하며, 가능하면 local asset을
사용하고, 실제 존재하는 files 또는 public repository state와 일치해야 한다.

## Badge Area

README top section은 두 badge rows를 가진다:

| Row | Source | Purpose |
| --- | --- | --- |
| Language chips | `assets/badge-english.svg`, `assets/badge-korean.svg` | English/Korean README 사이를 연결한다. |
| Trust badges | Local files plus shields-style image URLs | 현재 true인 repository facts만 보여준다. |

Language chips는 generated local SVG files다. Text는 짧게 유지하고 `alt` text와
accessible link labels에도 있으므로 language choice가 rendered pixels 안에만
존재하지 않는다.

## Factual Badge Rules

Factual badge는 backing evidence가 있을 때만 추가한다:

| Badge | Evidence |
| --- | --- |
| License | `LICENSE`가 있고 표시하는 license를 말한다. |
| CI | `.github/workflows/` 아래 workflow가 있다. |
| Security | `SECURITY.md`가 있다. |
| Release | Public GitHub Release가 있다. |
| Docs | Docs index가 있다. 현재는 `docs/00_START_HERE.md`다. |

Tags, drafts, local release notes, planned release work만으로 release badge를
추가하지 않는다. Release badge는 public GitHub Release가 있다는 뜻이다.

## Current README Badges

현재 포함된 factual badges:

| Badge | Evidence |
| --- | --- |
| License MIT | `LICENSE`가 있고 MIT License를 포함한다. |
| CI workflow exists | `.github/workflows/ci.yml`가 있다. |
| Security policy exists | `SECURITY.md`가 있다. |
| Docs index exists | `docs/00_START_HERE.md`가 있다. |

현재 생략된 factual badges:

| Badge | Reason |
| --- | --- |
| Release | Public GitHub releases page가 현재 releases 없음으로 표시한다. |

## Asset Checks

Language chip SVGs는 `scripts/render-assets.py`가 generate하고
`scripts/check-assets.py`가 required assets로 검증한다. Badge 변경 뒤에는 다음을
실행한다:

```bash
python3 scripts/render-assets.py
python3 scripts/check-assets.py
bash scripts/quality.sh
```

Manual README inspection에서는 다음을 확인한다:

- language chips가 보이고 올바른 README files로 link된다;
- Korean text가 `assets/badge-korean.svg`에서 acceptably render된다;
- factual badges가 실제 repository evidence와 일치한다;
- 중요한 trust claim이 badge images 안에만 있지 않고 nearby text 또는
  accessible labels에도 반복된다.
