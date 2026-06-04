# README Visual Asset Pass

## Current status

- v0.5.0 publication: verified
- README two-path onboarding: completed
- ni init . guided onboarding: implemented
- Windows real-host execution: deferred on macOS-only development host
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Pass goal

This pass prepares and integrates README visual assets without overclaiming
validation, install availability, or downstream execution behavior. It keeps
exact commands and status claims in Markdown or deterministic SVG, avoids fake
terminal proof, and treats generated raster images as future work unless real
files are present.

## Decision

VISUAL_ASSETS_PREPARED_WITH_PLACEHOLDERS

This task adds one deterministic local SVG diagram and references it from both
README surfaces. No AI-generated raster image was created or claimed.

## README image inventory

| Image reference | Path or URL | Exists? | Purpose | Action | Notes |
| --- | --- | --- | --- | --- | --- |
| README.md hero, README.ko.md hero | `assets/hero.svg` | Yes | Top visual identity and product framing. | Remain as-is. | Deterministic SVG; exact product copy remains nearby in Markdown and SVG. |
| README.md language chip, README.ko.md language chip | `assets/badge-english.svg` | Yes | English README link. | Remain as-is. | Deterministic SVG chip. |
| README.md language chip, README.ko.md language chip | `assets/badge-korean.svg` | Yes | Korean README link. | Remain as-is. | Deterministic SVG chip. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/license-MIT-f4b860` | Remote | License signal. | Remain as-is. | Remote factual badge; not AI-generated. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/CI-workflow%20exists-25334a` | Remote | CI workflow existence signal. | Remain as-is. | Does not claim CI passed. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/security-policy%20exists-2d5a52` | Remote | Security policy existence signal. | Remain as-is. | Does not claim security audit results. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/docs-index%20exists-5b8def` | Remote | Docs index existence signal. | Remain as-is. | Does not claim docs completeness. |
| README.md intent lock flow, README.ko.md intent lock flow | `assets/intent-lock-flow.svg` | Yes | Conceptual flow from conversation to bounded handoff. | Added and referenced. | Deterministic SVG; not a terminal screenshot or CLI proof. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-vague-intent.svg` | Yes | Show hidden missing planning inputs. | Remain as-is. | docs/116 has a future text-light prompt. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-early-execution.svg` | Yes | Show unsafe early work start. | Remain as-is. | docs/116 has a future text-light prompt. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-rework.svg` | Yes | Show rework risk from hidden assumptions. | Remain as-is. | docs/116 has a future text-light prompt. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-capture-intent.svg` | Yes | Show conversation becoming docs and contract. | Remain as-is. | docs/116 has a future text-light prompt. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-lock-contract.svg` | Yes | Show readiness and lock gate. | Remain as-is. | docs/116 has a future text-light prompt. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-handoff-safely.svg` | Yes | Show bounded handoff prompt. | Remain as-is. | Must not imply downstream execution. |

Exact text should stay in Markdown or deterministic SVG instead of AI-generated
raster images.

## docs/116 prompt usage

| Prompt / visual | Intended placement | Asset strategy | Used in this task? | Notes |
| --- | --- | --- | --- | --- |
| `assets/hero.svg` | README top hero | Keep existing deterministic SVG; future raster concept only if generated and reviewed. | No | Existing file remains referenced. |
| language chips | README language navigation | Keep deterministic SVG. | No | Not suitable for raster generation. |
| factual trust badges | README badge row | Keep remote shields or future deterministic SVG replacement. | No | Raster AI is not appropriate for factual badges. |
| pain cards | Why ni card row | Keep current SVGs; future text-light raster or SVG variants can use docs/116 prompts. | No | No fake terminal proof or runtime execution scenes. |
| payoff cards | What ni gives you card row | Keep current SVGs; future text-light raster or SVG variants can use docs/116 prompts. | No | Handoff visual must stop before execution. |
| intent lock flow | README intro, after product definition | Deterministic SVG placeholder added locally. | Yes | Derived from docs/116 flow and claim-boundary guidance, not from AI image generation. |

docs/116 remains the reusable prompt pack for future image-generation work. This
document summarizes usage instead of duplicating every prompt.

## Assets added

| Asset path | Type | Purpose | Referenced by README? | Notes |
| --- | --- | --- | --- | --- |
| `assets/intent-lock-flow.svg` | Deterministic SVG | Conceptual Intent Lock Protocol flow. | Yes, by README.md and README.ko.md. | Text-light, local, inspectable, no terminal proof, no third-party logos. |

## README integration

| Surface | Change | Pass? | Notes |
| --- | --- | --- | --- |
| README.md | Added one local `assets/intent-lock-flow.svg` image after the product definition. | Yes | Existing install and status wording preserved. |
| README.ko.md | Mirrored the same local image reference after the Korean product definition. | Yes | Commands and status strings remain unchanged. |
| alt text | Added accurate conceptual alt text. | Yes | Does not claim validation evidence. |
| local image paths | All local README image paths exist. | Yes | Verified by `python3 scripts/check-readme-surface.py` and `python3 scripts/check-assets.py`. |
| captions | No caption added. | Yes | Avoids extra claim surface. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Homebrew: Planned / v0.5 candidate. | README still says Planned / v0.5 candidate. | Yes | No `brew install` path added. |
| Windows real-host execution | Deferred unless a Windows transcript exists. | README still says real-host execution is deferred on this macOS-only host. | Yes | New SVG is platform-neutral. |
| generated images | Do not claim generated images exist unless files exist. | Only a deterministic local SVG file is claimed. | Yes | No AI-generated raster image was created. |
| ni run | `ni run` compiles a bounded handoff prompt and does not execute downstream work. | README wording remains non-executing. | Yes | SVG labels stop at handoff. |
| model workspace packs | Model workspace packs: Experimental. | README remains Experimental. | Yes | Skills are UX; CLI is authority. |
| no-terminal | No-terminal method: Experimental / assisted. | README remains Experimental / assisted. | Yes | No deterministic no-terminal claim. |
| benchmark evidence | Benchmark evidence must not prove implementation quality. | No benchmark claim added. | Yes | Visual asset pass is not benchmark evidence. |
| runtime execution boundary | ni is not a task runner, SPEC runner, execution harness, adapter, queue, or downstream execution layer. | README non-goals remain intact. | Yes | No runtime behavior added. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | M | Yes | Adds existing local image reference. |
| README.ko.md | M | Yes | Mirrors README.md image reference. |
| docs/116* | unchanged | No | Prompt pack remains source for future raster generation. |
| docs/121* | unchanged | No | Prior two-path onboarding pass remains baseline. |
| docs/122* | A | Yes | This visual asset audit and Korean companion. |
| docs/assets/readme/* | none | No | No files created under this optional path. |
| assets/intent-lock-flow.svg | A | Yes | Adds deterministic SVG placeholder under the existing README asset directory. |
| generated artifacts | none | No | No raster export or fake screenshot added. |
| .ni/contract.json | unchanged | No | Protected project-root planning state. |
| .ni/session.json | unchanged | No | Protected project-root planning state. |
| .ni/plan.lock.json | unchanged | No | Protected project-root lockfile. |
| unexpected files | none observed | No | Recheck before commit. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; expected README, roadmap, docs/122, and `assets/intent-lock-flow.svg` changes only. |
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

- Added `assets/intent-lock-flow.svg` as a local deterministic SVG
  diagram.
- Updated README.md and README.ko.md to reference the SVG only after the file
  exists.
- Added this audit and the Korean companion.
- Linked this pass from the post-release roadmap.

## What this pass proves

- README visual asset references are audited.
- Assets were integrated only if files exist.
- The prompt pack remains reusable for future image generation.
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

This is selected because no AI-generated raster images were generated in this
pass. Image files must be generated outside Codex or provided before README
references are switched away from deterministic local SVGs.

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
