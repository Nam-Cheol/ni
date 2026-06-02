# README Onboarding and Visual Prompt Pass

## Current status

- Human approval choice: DO_NOT_APPROVE_FIX_FIRST
- RC decision: RC_READY_WITH_DEFERRALS
- Release notes preflight decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES
- Artifact dry-run decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS
- Publication checklist decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES
- Human approval packet decision: HUMAN_APPROVAL_PACKET_READY_WITH_NOTES
- Release binary: Available
- Curl installer: Available
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.
- 이 pass는 publish, tag, release, GitHub release creation, asset upload,
  v0.5 released marking을 수행하지 않는다.

Note: 이 pass는 pasted task의 DO_NOT_APPROVE_FIX_FIRST fix-first context를
보존한다. Existing docs/115 content는 repository baseline에서 later
publication-prep decision을 기록하고 있다. 이 pass는 docs/115를 edit하지 않고
release action을 수행하지 않는다.

## Pass goal

이 pass는 publication approval을 다시 논의하기 전에 README onboarding instability를
고친다. README는 concise product pamphlet로 유지하고, bounded install/uninstall
guidance, 짧은 first-use tutorial, 그리고 image generation 없이 reproducible
visual prompts를 추가한다.

## README onboarding audit

| Area | Expected user answer | Current state | Change made | Notes |
| --- | --- | --- | --- | --- |
| What is ni? | ni is a Project Intent Compiler for AI Agents. | Present. | Preserved. | README는 compile-before-run framing을 유지한다. |
| Why compile intent before agent execution? | Vague intent는 handoff 전에 explicit plan이 되어야 한다. | Present. | Tutorial에서 보강. | Runtime execution을 imply하지 않는다. |
| macOS install | Source, local binary, verified v0.4.0 release binary, verified v0.4.0 curl installer를 사용한다. | Partially present. | Verified curl installer path를 쓰는 concise macOS install section 추가. | Default installer target은 `$HOME/.local/bin/ni`. |
| macOS uninstall | 선택한 install directory에서 installed `ni` binary를 제거한다. | 충분히 explicit하지 않았다. | Exact default removal command와 BINDIR caveat 추가. | PATH cleanup은 사용자가 PATH line을 추가한 경우에만 manual. |
| Windows install | Verified v0.4.0 `windows/amd64` release zip을 사용한다; package-manager installer claim 없음. | Install docs에는 있었고 README에는 부족했다. | Concise PowerShell release-zip flow 추가. | Hash comparison은 manual and visible. |
| Windows uninstall | 선택한 directory에서 copied `ni.exe`를 제거한다. | 충분히 explicit하지 않았다. | Bounded PowerShell removal example 추가. | 실제 install path를 사용해야 한다. |
| First-use tutorial | init, planning conversation, status proof, blocker resolution, lock, bounded prompt compile을 보여준다. | Short 60-second path가 있었다. | "5분 첫 project" 추가. | `ni run`은 explicitly non-executing. |
| What ni does not do | ni는 task runner, SPEC runner, runtime, adapter, queue, release automation system이 아니다. | Present. | Preserved. | Boundary remains prominent. |
| Links to deeper docs | README는 install, protocol, benchmark, Homebrew, command, visual docs로 route한다. | Present. | Concise 유지; docs/116은 roadmap에서 link할 수 있다. | README를 full manual로 만들지 않는다. |

## README image inventory

| Image reference | Path | Exists? | Current alt text | Purpose | Issue | Recommended action |
| --- | --- | --- | --- | --- | --- | --- |
| README.md hero, README.ko.md hero | `assets/hero.svg` | Yes | `ni hero banner: Project Intent Compiler for AI Agents` | Top visual identity and product framing. | Exact product text가 있어 raster generation에서 왜곡될 수 있다. | Deterministic SVG 유지; replacement 시 exact copy는 Markdown/SVG에 둔다. |
| README.md language chip, README.ko.md language chip | `assets/badge-english.svg` | Yes | `English` | English README link. | Exact text는 raster AI가 아니라 SVG/Markdown에 속한다. | Deterministic SVG chip 유지. |
| README.md language chip, README.ko.md language chip | `assets/badge-korean.svg` | Yes | `한국어` | Korean README link. | Exact Korean text는 raster AI가 아니라 SVG/Markdown에 속한다. | Deterministic SVG chip 유지. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/license-MIT-f4b860` | Remote | `License MIT` | License signal. | Remote dependency; exact text badge. | 유지하거나 local deterministic SVG로 교체; AI-generate 금지. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/CI-workflow%20exists-25334a` | Remote | `CI workflow exists` | CI workflow existence signal. | Remote dependency; exact factual claim. | 유지하거나 local deterministic SVG로 교체; CI passed를 imply하지 않는다. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/security-policy%20exists-2d5a52` | Remote | `Security policy exists` | Security policy existence signal. | Remote dependency; exact factual claim. | 유지하거나 local deterministic SVG로 교체; audit result를 imply하지 않는다. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/docs-index%20exists-5b8def` | Remote | `Docs index exists` | Docs index existence signal. | Remote dependency; exact factual claim. | 유지하거나 local deterministic SVG로 교체; docs completeness를 imply하지 않는다. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-vague-intent.svg` | Yes | 각 README의 Vague intent alt text. | Hidden planning inputs를 보여준다. | Visual은 text-like card framing을 포함한다; exact claims는 Markdown에 있어야 한다. | SVG 유지 또는 text-light abstract card로 교체. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-early-execution.svg` | Yes | 각 README의 Early execution alt text. | Unsafe early work start를 보여준다. | Same as above. | SVG 유지 또는 text-light abstract card로 교체. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-rework.svg` | Yes | 각 README의 Rework alt text. | Hidden assumptions의 cost를 보여준다. | Same as above. | SVG 유지 또는 text-light abstract card로 교체. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-capture-intent.svg` | Yes | 각 README의 Capture intent alt text. | Conversation이 docs와 contract가 되는 것을 보여준다. | Same as above. | SVG 유지 또는 text-light abstract card로 교체. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-lock-contract.svg` | Yes | 각 README의 Lock contract alt text. | Readiness and lock gate를 보여준다. | Same as above. | SVG 유지 또는 text-light abstract card로 교체. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-handoff-safely.svg` | Yes | 각 README의 Handoff safely alt text. | Bounded handoff prompt를 보여준다. | Same as above. | Execution을 imply하지 않는 text-light abstract card로 교체 가능. |

## Image-generation prompt pack

이 task에서는 image generation을 수행하지 않았다. 아래 prompts는 later visual pass용이다.
Exact CLI text, status constants, badges, product claims는 AI-generated raster
text가 아니라 Markdown 또는 deterministic SVG로 render해야 한다.

### assets/hero.svg

Purpose: README hero visual identity for compile-before-run.
Recommended use: Optional replacement concept only; exact product copy는
Markdown 또는 deterministic SVG에 둔다.
Suggested dimensions: 1200 x 460, 60:23.
Style: Crisp abstract editorial illustration, local-first developer tool,
soft grid, contract page, lock shape, handoff path, restrained colors.
Prompt: Create a text-free hero banner for a pre-runtime project intent compiler
for AI agents. Show a planning conversation transforming into a structured
contract, passing through a readiness gate and lock, then becoming a bounded
handoff packet. Use abstract interface shapes, document panels, hash lines, and
a clear left-to-right flow. Keep it clean, professional, high contrast, and
readable at repository README width. Do not include logos, screenshots, terminal
output, real brand marks, or readable text.
Negative prompt: No GitHub UI screenshot, no terminal output, no fake validation
logs, no release badge, no Homebrew badge, no robots running tasks, no shell
execution, no tiny text, no copyrighted logos, no claim that v0.5 is released.
Exact text policy: "Don't run the agent yet. Compile the intent first."는
Markdown 또는 deterministic SVG에만 둔다.
Alt text: Abstract banner showing conversation becoming a locked planning
contract before downstream handoff.
Claim-boundary warning: ni가 agents 또는 downstream work를 execute한다고 imply하면 안 된다.

### assets/badge-english.svg

Purpose: English README link.
Recommended use: Deterministic SVG 또는 HTML text link 유지.
Suggested dimensions: 84 x 28.
Style: Small flat language chip.
Prompt: Create a simple local-language navigation chip background with no
readable text, suitable for overlaying deterministic SVG text later.
Negative prompt: No flags, no brand marks, no fake status badge, no tiny raster
letters, no national symbolism.
Exact text policy: "English"는 deterministic SVG 또는 Markdown으로만 render한다.
Alt text: English language link.
Claim-boundary warning: Language navigation only; no product status claim.

### assets/badge-korean.svg

Purpose: Korean README link.
Recommended use: Deterministic SVG 또는 HTML text link 유지.
Suggested dimensions: 84 x 28.
Style: Small flat language chip.
Prompt: Create a simple local-language navigation chip background with no
readable text, suitable for overlaying deterministic Korean SVG text later.
Negative prompt: No flags, no brand marks, no fake status badge, no tiny raster
letters, no national symbolism.
Exact text policy: "한국어"는 deterministic SVG 또는 Markdown으로만 render한다.
Alt text: Korean language link.
Claim-boundary warning: Language navigation only; no product status claim.

### license badge

Purpose: License signal.
Recommended use: Remote shields badge 유지 또는 local deterministic SVG 교체; raster AI 금지.
Suggested dimensions: approximately 92 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape that can receive deterministic SVG text.
Negative prompt: No legal seal, no fake certification, no brand marks, no
readable raster text.
Exact text policy: "License MIT"는 shields.io, Markdown, deterministic SVG로만 render한다.
Alt text: License MIT.
Claim-boundary warning: License existence only.

### CI workflow badge

Purpose: CI workflow existence signal.
Recommended use: Remote shields badge 유지 또는 local deterministic SVG 교체; raster AI 금지.
Suggested dimensions: approximately 150 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape for deterministic status text.
Negative prompt: No green passing checkmark unless a passing run is being
claimed, no fake workflow screenshot, no readable raster text.
Exact text policy: "CI workflow exists"는 shields.io, Markdown, deterministic SVG로만 render한다.
Alt text: CI workflow exists.
Claim-boundary warning: Workflow file exists만 의미한다; CI passed를 의미하지 않는다.

### security policy badge

Purpose: Security policy existence signal.
Recommended use: Remote shields badge 유지 또는 local deterministic SVG 교체; raster AI 금지.
Suggested dimensions: approximately 175 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape for deterministic status text.
Negative prompt: No security certification seal, no audit passed claim, no
shield logo from a third party, no readable raster text.
Exact text policy: "Security policy exists"는 shields.io, Markdown, deterministic SVG로만 render한다.
Alt text: Security policy exists.
Claim-boundary warning: Security audit 또는 vulnerability status를 imply하지 않는다.

### docs index badge

Purpose: Docs index existence signal.
Recommended use: Remote shields badge 유지 또는 local deterministic SVG 교체; raster AI 금지.
Suggested dimensions: approximately 150 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape for deterministic status text.
Negative prompt: No completeness ribbon, no fake documentation score, no
readable raster text.
Exact text policy: "Docs index exists"는 shields.io, Markdown, deterministic SVG로만 render한다.
Alt text: Docs index exists.
Claim-boundary warning: Docs complete 또는 publication-ready를 imply하지 않는다.

### assets/card-pain-vague-intent.svg

Purpose: Plausible requests 안의 hidden planning gaps를 visualize한다.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Abstract document card, missing fields, muted warning accent.
Prompt: Create a text-free visual card showing a plausible request with hidden
empty fields for users, acceptance criteria, risks, non-goals, and blockers.
Use abstract document panels and subtle missing-field indicators. Keep it clean
and language-neutral.
Negative prompt: No readable text, no terminal logs, no fake user data, no
third-party UI, no runtime execution scene.
Exact text policy: Labels and explanations는 Markdown 또는 deterministic SVG에 둔다.
Alt text: Vague intent can hide missing users, acceptance criteria, risks,
non-goals, or blockers.
Claim-boundary warning: ni가 모든 ambiguity를 탐지한다고 imply하지 않는다.

### assets/card-pain-early-execution.svg

Purpose: Premature work start 전에 stop하는 것을 visualize한다.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Abstract gate before a work lane, restrained warning accent.
Prompt: Create a text-free visual card showing a work lane paused at a gate
before implementation begins, with planning materials waiting for validation.
Use abstract shapes, not literal brand tools.
Negative prompt: No robots executing work, no shell commands, no code editor
screenshot, no fake CI output, no readable text.
Exact text policy: "Early execution"과 explanation은 Markdown 또는 deterministic SVG에 둔다.
Alt text: Work should not begin just because a request sounds plausible.
Claim-boundary warning: ni가 implementation을 run 또는 supervise한다고 imply하지 않는다.

### assets/card-pain-rework.svg

Purpose: Hidden assumptions에서 생기는 costly rework를 visualize한다.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Abstract branching plan with one path crossed out and corrected.
Prompt: Create a text-free visual card showing a plan branch being corrected
before costly implementation rework spreads. Use document nodes, a gentle
correction mark, and clear before/after path structure.
Negative prompt: No blame imagery, no fake metrics, no implementation-quality
claim, no readable text.
Exact text policy: "Rework"와 explanation은 Markdown 또는 deterministic SVG에 둔다.
Alt text: Hidden assumptions become expensive after work starts from the wrong plan.
Claim-boundary warning: Benchmarked cost reduction을 imply하지 않는다.

### assets/card-payoff-capture-intent.svg

Purpose: Conversation이 docs와 contract draft가 되는 것을 visualize한다.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Conversation bubbles turning into structured pages.
Prompt: Create a text-free visual card showing model-user planning conversation
being organized into docs and a contract draft. Use abstract bubbles, document
cards, and field outlines with no readable text.
Negative prompt: No chat-provider UI, no brand marks, no fake transcript, no
readable text, no claim that a model alone decides readiness.
Exact text policy: Exact file paths는 Markdown 또는 deterministic SVG에 둔다.
Alt text: Planning conversation becomes explicit docs and a contract draft.
Claim-boundary warning: Models draft and CLI validates boundary를 보존한다.

### assets/card-payoff-lock-contract.svg

Purpose: Readiness and lock hash trust를 visualize한다.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Contract page, readiness gate, lock, hash line.
Prompt: Create a text-free visual card showing a contract passing through a
readiness gate into a lock with hash-like abstract marks. Keep it precise,
minimal, and language-neutral.
Negative prompt: No fake `READY` output, no real terminal, no fake signature,
no legal seal, no readable text.
Exact text policy: `ni status`, `ni end`, `READY`, hash wording은 Markdown 또는
deterministic SVG에 둔다.
Alt text: Readiness and lock commands gate the accepted plan, hashes, and
source of truth.
Claim-boundary warning: `ni status` 없이 readiness를 imply하지 않는다.

### assets/card-payoff-handoff-safely.svg

Purpose: Valid lock 이후 bounded prompt compilation을 visualize한다.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Locked plan producing a sealed handoff packet, not an execution scene.
Prompt: Create a text-free visual card showing a locked contract producing a
bounded handoff packet. The packet should stop before any runtime or agent work
begins. Use calm developer-tool styling and clear separation between source
contract and downstream packet.
Negative prompt: No robot worker, no shell command, no queue, no PR automation,
no fake downstream success, no readable text.
Exact text policy: `ni run --max-chars 4000`은 Markdown 또는 deterministic SVG에 둔다.
Alt text: A valid locked plan compiles into a bounded handoff prompt or derived
seed material.
Claim-boundary warning: `ni run`은 compile만 하며 downstream work를 execute하지 않는다.

## macOS install / uninstall notes

README는 script inspection 후 verified v0.4.0 curl installer path를 macOS 권장
경로로 설명한다. `install.sh` default target은 `$HOME/.local/bin/ni`이고,
`BINDIR`로 directory를 override할 수 있으며, verification은 `--help`와 `version`이다.

| Install path | README status | Evidence boundary | Notes |
| --- | --- | --- | --- |
| Source | Available | Go가 있는 local checkout. | Existing README path preserved. |
| Local binary | Available | `make build` / `make install-local`. | Detail은 docs/22에 둔다. |
| Release binary | Available | Verified v0.4.0 release assets. | Manual archive path preserved. |
| Curl installer | Available | Verified v0.4.0 installer path. | Default target은 `$HOME/.local/bin/ni`. |
| Homebrew | Planned / v0.5 candidate | Tap/formula/install proof 없음. | README는 현재 `brew install`을 금지한다. |

Uninstall wording은 선택한 directory의 exact installed `ni` file을 제거하고,
PATH cleanup은 사용자가 직접 line을 추가한 경우에만 manual로 둔다.

## Windows install / uninstall notes

README는 Windows에 대해 v0.4.0 `windows/amd64` release zip path만 문서화한다.
MSI, winget, Chocolatey, Scoop, Homebrew, verified Windows installer는 claim하지 않는다.

| Install path | README status | Evidence boundary | Notes |
| --- | --- | --- | --- |
| Release zip `windows/amd64` | Available for v0.4.0 release assets | Same release의 archive와 checksum file. | User가 `Get-FileHash`와 checksum text를 비교한다. |
| Manual PATH placement | Manual / user-controlled | User가 directory를 선택한다. | README는 controlled directory에 `ni.exe`를 두라고 말한다. |
| Windows package manager | Not claimed | Verification 없음. | winget, Chocolatey, Scoop, MSI claim 없음. |
| Windows arm64 | Not claimed | Installer says Windows arm64 release asset is not configured. | Availability claim 없음. |

Uninstall wording은 actual install directory에서 copied `ni.exe`를 제거하고,
PATH entry는 ni-specific으로 추가한 경우에만 제거한다.

## First-use tutorial notes

README tutorial은 다음을 보여준다:

1. `ni init`
2. execution 전 model-user planning conversation
3. `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`
4. `ni status --proof --next-questions`
5. `BLOCKED` questions 또는 gaps 해결
6. `ni end`
7. `ni run --max-chars 4000`

Tutorial은 `ni run`이 `.ni/plan.lock.json`에서 bounded handoff prompt를 compile할
뿐이며 prompts, agents, shell commands, downstream work를 execute하지 않는다고
명시한다.

## README.ko companion audit

| Section | English source | Korean companion status | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Hero and definition | README.md top section | Mirrored | Yes | Product definition preserved. |
| Start in 60 seconds | README.md | Mirrored | Yes | Commands unchanged. |
| First project tutorial | README.md | Added | Yes | Commands, paths, constants preserved. |
| Choose your path | README.md | Already mirrored | Yes | Status rows unchanged. |
| macOS install / uninstall | README.md | Added | Yes | Same verified path and default target. |
| Windows install / uninstall | README.md | Added | Yes | Same no-package-manager boundary. |
| What ni is not | README.md | Preserved | Yes | Non-execution boundary preserved. |
| Read next | README.md | Preserved | Yes | Existing companion links remain. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Published/released status | v0.5 published claim 금지. | README/docs116은 v0.5 release를 claim하지 않는다. | Yes | No tag, upload, release action. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | README는 Homebrew remains Planned / v0.5 candidate와 `brew install` 금지를 말한다. | Yes | No Homebrew Available claim. |
| Model workspace packs | Model workspace packs: Experimental. | README path table remains Experimental. | Yes | Skills are UX; CLI is authority. |
| No-terminal | No-terminal method: Experimental / assisted. | README path table remains Experimental. | Yes | No deterministic no-terminal claim. |
| ni run | Prompt compilation only. | Tutorial and payoff copy say no execution. | Yes | No downstream work execution. |
| READY | Planning contract readiness only. | Tutorial says CLI decides readiness; docs116 does not claim product readiness. | Yes | `READY` is not product readiness. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | New stale-lock claim 없음. | Yes | Existing docs remain source for recovery. |
| Benchmark evidence | Implementation-quality 또는 downstream-quality proof 금지. | docs116 prompt pack avoids benchmark overclaim. | Yes | No empirical effect claim. |
| Runtime execution boundary | ni-kernel stops before runtime execution. | README "ni가 아닌 것" preserved. | Yes | No task runner or harness added. |
| Image prompts | Conceptual prompts only, no generated assets. | Prompt pack forbids fake validation output and release claims. | Yes | Exact text stays Markdown/SVG. |

## Git status / inclusion check

| Path or group | git status --short | git ls-files / tracked check | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| README.md | Modified | tracked | Yes | Onboarding improvements. |
| README.ko.md | Modified | tracked | Yes | Companion improvements. |
| docs/110_* | no new edit in this pass | tracked | No new change | Existing release-candidate docs. |
| docs/111_* | no new edit in this pass | tracked | No new change | Existing release-note draft docs. |
| docs/112_* | no new edit in this pass | tracked | No new change | Existing preflight docs. |
| docs/113_* | no new edit in this pass | tracked | No new change | Existing artifact dry-run docs. |
| docs/114_* | no new edit in this pass | tracked | No new change | Existing publication checklist docs. |
| docs/115_* | no final diff | tracked | No | Existing baseline은 later publication-prep decision을 기록한다; 여기서는 edit하지 않았다. |
| docs/116_* | Added | untracked until staged | Yes | Required output for this pass. |
| docs/51* | Modified | tracked | Yes | Narrow docs/116 roadmap pointer only. |
| image assets | no new image files | tracked existing assets | No | Generated images added 없음. |
| generated artifacts | none expected | not applicable | No | Validation이 ignored temp/build output을 만들 수 있다. |
| .ni/contract.json | no diff expected | tracked | No | Protected root file. |
| .ni/session.json | no diff expected | tracked | No | Protected root file. |
| .ni/plan.lock.json | no diff expected | tracked | No | Protected root lockfile. |
| unexpected files | none | final status reviewed | No | Final visible changes는 README, docs/51, docs/116이다. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Final visible changes: README.md, README.ko.md, docs/51*, untracked docs/116*. |
| `git ls-files docs/110_... docs/115_...` | Pass | docs/110_* through docs/115_* are tracked. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | Temp Go cache workaround 사용. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, warnings: None. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass | Help output rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass | Output: `0.0.0-dev`; source build only. |
| `python3 scripts/check-install-docs.py` | Pass | Install claim boundaries preserved. |
| `bash scripts/check-skill-packs.sh` | Pass | Model workspace packs remain Experimental; global install remains unverified. |
| `bash scripts/demo-check.sh` | Pass | Temporary prompts/exports compiled as seed or prompt artifacts only; not executed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Fixture `ni end` / `ni relock`를 exercise한다; project-root relock이 아니다. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source, build, temporary install paths passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Check-only release readiness gate passed; release action 없음. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | Protected project-root `.ni` diff 없음. |

## Changes made

- `README.md`: first-use tutorial, macOS install/uninstall, Windows install/uninstall 추가.
- `README.ko.md`: same promise level의 Korean companion sections 추가.
- `docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.md`: audit and prompt pack 추가.
- `docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.ko.md`: 이 Korean companion 추가.
- `docs/51_POST_RELEASE_ROADMAP.md` / `.ko.md`: final diff에 포함되는 경우 narrow pointer만 추가.

## What this pass proves

- README onboarding was audited and improved for the touched surfaces.
- README image prompts are now reproducible for later image generation.
- macOS / Windows install and uninstall wording is bounded by current evidence.
- first-use tutorial preserves CLI authority and non-execution boundaries.
- no release action was performed.

## What this pass does not prove

- v0.5 has been published.
- GitHub release exists.
- assets were uploaded.
- Homebrew is Available.
- Windows installer exists unless actually verified.
- generated images have been produced.
- generated images are final assets.
- model workspace host behavior is verified.
- no-terminal is deterministic.
- downstream execution succeeds.
- benchmark effect size or causal impact.

## Recommended next task

Selected next task: B. README install-doc verification pass.

Why: README onboarding은 개선되었고 image prompts는 later visual pass에 충분하지만,
publication-facing risk가 가장 큰 영역은 README, docs/22, curl installer docs,
release status 전반의 public install wording이다. Focused install-doc verification
pass가 publication approval을 다시 논의하기 전에 macOS와 Windows wording을 재검증할
수 있다.

## Next task prompt

```text
Proceed with a README install-doc verification pass in /Users/namba/Documents/project/ni.

Goal:
Verify README.md, README.ko.md, docs/22_INSTALL.md, docs/install-curl.md, and docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.md for install and uninstall accuracy before publication approval is revisited.

Scope:
- documentation and static checks only
- verify macOS default installer target from install.sh
- verify Windows release zip wording against documented v0.4.0 windows/amd64 asset naming
- keep Homebrew: Planned / v0.5 candidate
- keep Model workspace packs: Experimental
- keep No-terminal method: Experimental / assisted
- preserve Skills are UX; CLI is authority.
- do not publish, tag, create a GitHub release, upload assets, run goreleaser publish, create or publish Homebrew formula, execute generated prompts, run ni end on the project root, relock the project root, or edit protected root .ni files

Validation:
- git status --short
- python3 scripts/check-install-docs.py
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Report changed files, exact install/uninstall claim boundaries, any remaining blockers, and confirmation that no release action or project-root relock was performed.
```
