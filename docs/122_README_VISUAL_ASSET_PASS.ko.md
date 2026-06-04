# README Visual Asset Pass

## Current status

- v0.5.0 publication: verified
- README two-path onboarding: completed
- ni init . guided onboarding: implemented
- Windows real-host execution: macOS-only development host에서는 deferred
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Pass goal

이 pass는 validation, install availability, downstream execution behavior를
overclaim하지 않으면서 README visual assets를 준비하고 통합한다. Exact commands와
status claims는 Markdown 또는 deterministic SVG에 남기고, fake terminal proof를
피하며, 실제 파일이 없는 generated raster images는 future work로 둔다.

## Decision

VISUAL_ASSETS_PREPARED_WITH_PLACEHOLDERS

이번 task는 local deterministic SVG diagram 하나를 추가하고 두 README surface에서
참조한다. AI-generated raster image는 만들거나 claim하지 않았다.

## README image inventory

| Image reference | Path or URL | Exists? | Purpose | Action | Notes |
| --- | --- | --- | --- | --- | --- |
| README.md hero, README.ko.md hero | `assets/hero.svg` | Yes | Top visual identity and product framing. | Remain as-is. | Deterministic SVG; exact product copy는 nearby Markdown과 SVG에 남아 있다. |
| README.md language chip, README.ko.md language chip | `assets/badge-english.svg` | Yes | English README link. | Remain as-is. | Deterministic SVG chip. |
| README.md language chip, README.ko.md language chip | `assets/badge-korean.svg` | Yes | Korean README link. | Remain as-is. | Deterministic SVG chip. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/license-MIT-f4b860` | Remote | License signal. | Remain as-is. | Remote factual badge; AI-generated 아님. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/CI-workflow%20exists-25334a` | Remote | CI workflow existence signal. | Remain as-is. | CI passed를 claim하지 않는다. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/security-policy%20exists-2d5a52` | Remote | Security policy existence signal. | Remain as-is. | Security audit result를 claim하지 않는다. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/docs-index%20exists-5b8def` | Remote | Docs index existence signal. | Remain as-is. | Docs completeness를 claim하지 않는다. |
| README.md intent lock flow, README.ko.md intent lock flow | `assets/intent-lock-flow.svg` | Yes | Conversation에서 bounded handoff까지의 conceptual flow. | Added and referenced. | Deterministic SVG; terminal screenshot 또는 CLI proof가 아니다. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-vague-intent.svg` | Yes | Hidden missing planning inputs를 보여준다. | Remain as-is. | docs/116에 future text-light prompt가 있다. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-early-execution.svg` | Yes | Unsafe early work start를 보여준다. | Remain as-is. | docs/116에 future text-light prompt가 있다. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-rework.svg` | Yes | Hidden assumptions의 rework risk를 보여준다. | Remain as-is. | docs/116에 future text-light prompt가 있다. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-capture-intent.svg` | Yes | Conversation이 docs와 contract가 되는 것을 보여준다. | Remain as-is. | docs/116에 future text-light prompt가 있다. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-lock-contract.svg` | Yes | Readiness and lock gate를 보여준다. | Remain as-is. | docs/116에 future text-light prompt가 있다. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-handoff-safely.svg` | Yes | Bounded handoff prompt를 보여준다. | Remain as-is. | Downstream execution을 imply하면 안 된다. |

Exact text는 AI-generated raster images 대신 Markdown 또는 deterministic SVG에 둔다.

## docs/116 prompt usage

| Prompt / visual | Intended placement | Asset strategy | Used in this task? | Notes |
| --- | --- | --- | --- | --- |
| `assets/hero.svg` | README top hero | Existing deterministic SVG 유지; future raster concept는 생성과 review 이후에만 사용. | No | Existing file remains referenced. |
| language chips | README language navigation | Deterministic SVG 유지. | No | Raster generation에 적합하지 않다. |
| factual trust badges | README badge row | Remote shields 또는 future deterministic SVG replacement 유지. | No | Factual badge에는 raster AI가 적합하지 않다. |
| pain cards | Why ni card row | Current SVG 유지; future text-light raster 또는 SVG variants는 docs/116 prompts를 사용할 수 있다. | No | Fake terminal proof 또는 runtime execution scene 금지. |
| payoff cards | What ni gives you card row | Current SVG 유지; future text-light raster 또는 SVG variants는 docs/116 prompts를 사용할 수 있다. | No | Handoff visual은 execution 전에 멈춰야 한다. |
| intent lock flow | README intro, product definition 뒤 | Deterministic SVG placeholder를 local로 추가. | Yes | docs/116 flow와 claim-boundary guidance에서 파생했으며 AI image generation은 아니다. |

docs/116은 future image-generation work를 위한 reusable prompt pack으로 남는다.
이 문서는 모든 prompt를 중복하지 않고 usage만 요약한다.

## Assets added

| Asset path | Type | Purpose | Referenced by README? | Notes |
| --- | --- | --- | --- | --- |
| `assets/intent-lock-flow.svg` | Deterministic SVG | Conceptual Intent Lock Protocol flow. | Yes, README.md and README.ko.md. | Text-light, local, inspectable, no terminal proof, no third-party logos. |

## README integration

| Surface | Change | Pass? | Notes |
| --- | --- | --- | --- |
| README.md | Product definition 뒤에 local `assets/intent-lock-flow.svg` image 하나를 추가. | Yes | Existing install and status wording preserved. |
| README.ko.md | Korean product definition 뒤에 같은 local image reference를 mirror. | Yes | Commands and status strings remain unchanged. |
| alt text | Accurate conceptual alt text 추가. | Yes | Validation evidence를 claim하지 않는다. |
| local image paths | 모든 local README image path가 존재한다. | Yes | `python3 scripts/check-readme-surface.py`와 `python3 scripts/check-assets.py`로 검증. |
| captions | Caption 추가 없음. | Yes | Extra claim surface를 피한다. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Homebrew: Planned / v0.5 candidate. | README는 여전히 Planned / v0.5 candidate라고 말한다. | Yes | `brew install` path 추가 없음. |
| Windows real-host execution | Windows transcript 전까지 deferred. | README는 real-host execution이 macOS-only host에서 deferred라고 유지한다. | Yes | New SVG is platform-neutral. |
| generated images | File이 존재하지 않으면 generated images 존재를 claim하지 않는다. | Deterministic local SVG file만 claim한다. | Yes | AI-generated raster image 생성 없음. |
| ni run | `ni run`은 bounded handoff prompt를 compile하며 downstream work를 execute하지 않는다. | README wording remains non-executing. | Yes | SVG labels stop at handoff. |
| model workspace packs | Model workspace packs: Experimental. | README remains Experimental. | Yes | Skills are UX; CLI is authority. |
| no-terminal | No-terminal method: Experimental / assisted. | README remains Experimental / assisted. | Yes | Deterministic no-terminal claim 없음. |
| benchmark evidence | Benchmark evidence는 implementation quality를 prove하면 안 된다. | Benchmark claim 추가 없음. | Yes | Visual asset pass는 benchmark evidence가 아니다. |
| runtime execution boundary | ni는 task runner, SPEC runner, execution harness, adapter, queue, downstream execution layer가 아니다. | README non-goals remain intact. | Yes | Runtime behavior 추가 없음. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | M | Yes | Existing local image reference 추가. |
| README.ko.md | M | Yes | README.md image reference mirror. |
| docs/116* | unchanged | No | Prompt pack remains source for future raster generation. |
| docs/121* | unchanged | No | Prior two-path onboarding pass remains baseline. |
| docs/122* | A | Yes | This visual asset audit and Korean companion. |
| docs/assets/readme/* | none | No | 이 optional path에는 file을 만들지 않았다. |
| assets/intent-lock-flow.svg | A | Yes | Existing README asset directory 아래 deterministic SVG placeholder 추가. |
| generated artifacts | none | No | Raster export 또는 fake screenshot 추가 없음. |
| .ni/contract.json | unchanged | No | Protected project-root planning state. |
| .ni/session.json | unchanged | No | Protected project-root planning state. |
| .ni/plan.lock.json | unchanged | No | Protected project-root lockfile. |
| unexpected files | none observed | No | Commit 전 재확인. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; expected README, roadmap, docs/122, `assets/intent-lock-flow.svg` changes only. |
| Verify all README.md local image paths exist | Passed via `python3 scripts/check-readme-surface.py` and `python3 scripts/check-assets.py`. |
| Verify all README.ko.md local image paths exist | Passed via `python3 scripts/check-readme-surface.py` and `python3 scripts/check-assets.py`. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; `NI Intent Readiness: READY`, blockers/deferrals/warnings none. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build printed `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; no output. |

## Changes made

- `assets/intent-lock-flow.svg` local deterministic SVG diagram을 추가했다.
- File이 존재한 뒤에만 README.md와 README.ko.md에서 SVG를 참조했다.
- 이 audit과 Korean companion을 추가했다.
- Post-release roadmap에서 이 pass로 연결했다.

## What this pass proves

- README visual asset references are audited.
- Assets were integrated only if files exist.
- Prompt pack remains reusable for future image generation.
- README claim boundaries remain safe.

## What this pass does not prove

- New AI-generated images were created unless files are present.
- Generated images are validation proof.
- Windows real-host execution works.
- Homebrew is Available.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

Selected: A. generate actual README raster images with image-generation tool

이번 pass에서는 AI-generated raster images를 만들지 않았기 때문에 이 task를
선택한다. README references를 deterministic local SVG에서 바꾸기 전에는 image
files가 Codex 밖에서 생성되거나 제공되어야 한다.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: Generate actual README raster images from docs/116 prompts.

Goal:
Create real local README raster image files from the docs/116 prompt pack, then
integrate only the images that are actually present, claim-safe, text-light, and
locally verified.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.md
- docs/122_README_VISUAL_ASSET_PASS.md
- assets/intent-lock-flow.svg
- current assets/ files

Rules:
- Do not publish, tag, create a GitHub release, upload release assets, push, or run release workflows.
- Do not run ni end or relock the project root.
- Do not edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json.
- Do not execute generated prompts.
- Do not create fake terminal proof.
- Do not include third-party logos or copyrighted marks.
- Do not imply Homebrew is Available.
- Do not imply Windows real-host execution is verified.
- Do not imply ni run executes downstream work.
- Keep exact commands, statuses, and product claims in Markdown or deterministic SVG, not AI-generated raster text.

Work:
1. Generate or receive actual raster files for the text-light visuals selected from docs/116.
2. Place accepted files under assets/ or another stable documented local path.
3. Verify every local image path exists before editing README references.
4. Update README.md only for image files that exist and are claim-safe.
5. Mirror README.ko.md only when README.md changes.
6. Update docs/122 with final asset paths, prompt usage, claim-boundary audit, and validation results.

Validation:
- git status --short
- verify all README.md local image paths exist
- verify all README.ko.md local image paths exist
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, generated assets, README image reference changes,
README.ko image reference changes, local path verification, claim-boundary audit,
validation results, protected .ni diff, and confirmation that no
publish/tag/release/upload/project-root relock/generated prompt execution occurred.
```
