# Capsule Style Notes

Task 108은 `ni`가 capsule-style README banner를 visual inspiration으로 참고하되
primary README surface에 remote dependency를 추가하지 않는 방식을 기록한다.

이 문서는 visual policy note only다. README assets, product behavior, asset
generator, kernel boundary를 바꾸지 않는다.

## Why Capsule-Style Banners Are Useful

Capsule-style banners는 README first impression을 잡는 데 좋은 inspiration이 될
수 있다. GitHub reader에게 익숙하고, product name, short phrase, quiet visual
mood를 하나의 scan-friendly image에 담기 좋다.

`ni`에서는 이 style이 기존 visual direction과 잘 맞는다:

- centered product identity;
- restrained gradients;
- short text;
- single banner-shaped surface;
- polished first-screen README rhythm.

이 style은 local template composition, spacing, contrast를 잡는 데 참고할 수
있지만, `ni` README assets의 source of truth는 아니다.

## Remote Render Convenience

Remote query-rendered SVG services는 README author가 asset file을 commit하지
않고 URL parameters로 banner를 표현할 수 있어 편리하다. Reference only로 보는
example pattern은 다음과 같다:

```text
https://capsule-render.vercel.app/api?type=waving&text=ni&fontSize=64
```

실제 service URL에서는 query text와 special characters를 escape해야 하고,
rendered result는 service의 parameter model에 좌우된다. 이 repository는 그런
pattern을 documentation의 comparison 또는 inspiration으로 참고할 수 있지만,
primary README hero로 사용하지 않는다.

## Local Asset Policy

`ni`의 primary README visuals는 local, deterministic, reviewable해야 한다.

Core README art에 local assets를 쓰는 이유는 repository가 다음을 직접 control할
수 있기 때문이다:

- repo에 commit되는 exact SVG source;
- text escaping과 line breaks;
- font fallback behavior;
- dimensions, safe areas, reduced-width rendering;
- GitHub, offline readers, mirrors가 README를 render할 때의 availability;
- visual changes에 대한 review diffs.

Remote query-rendered SVGs는 quick experiments에는 유용할 수 있지만, text
escaping, font fallback, service parameters, service availability가 rendered
output에 영향을 줄 수 있다. 이 repository에는 local SVG template pipeline이 더
controllable하다.

## ni Decision

`ni`는 core README assets에 local SVG templates를 사용한다.

Editable source of truth는 `assets/source/`와 `scripts/render-assets.py`에 있다.
Generated SVG files는 `assets/` 아래에 commit되어 README가 remote hero, remote
capsule service, external font, external CSS, script, remote image dependency
없이 render될 수 있게 한다.

Allowed:

- local SVG templates와 generated local SVG files;
- platform이 PNG를 요구할 때 optional locally generated PNG exports;
- non-essential factual remote status badges;
- optional inspiration으로 문서화된 remote capsule-style URL patterns.

Primary README visuals에서 금지:

- `assets/hero.svg`를 remote capsule-render URL로 대체하는 것;
- remote dynamic SVG renderer를 main hero로 추가하는 것;
- SVG assets 안에서 external fonts, CSS, scripts, images에 의존하는 것;
- essential README meaning을 rendered image 안에서만 제공하는 것.

See also:

- `docs/60_VISUAL_DESIGN_SPEC.ko.md`
- `docs/61_ASSET_PIPELINE.ko.md`
- `docs/63_README_VISUAL_WIREFRAME.ko.md`
