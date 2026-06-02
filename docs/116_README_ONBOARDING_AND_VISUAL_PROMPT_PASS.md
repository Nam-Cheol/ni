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
- This pass does not publish, tag, release, create a GitHub release, upload
  assets, or mark v0.5 as released.

Note: this pass preserves the DO_NOT_APPROVE_FIX_FIRST fix-first context from
the pasted task. Existing docs/115 content records a later publication-prep
decision in the repository baseline; this pass does not edit docs/115 and does
not perform release actions.

## Pass goal

This pass fixes README onboarding instability before publication approval is
revisited. It keeps README as a concise product pamphlet, adds bounded install
and uninstall guidance, adds a short first-use tutorial, and records
reproducible visual prompts without generating new images.

## README onboarding audit

| Area | Expected user answer | Current state | Change made | Notes |
| --- | --- | --- | --- | --- |
| What is ni? | ni is a Project Intent Compiler for AI Agents. | Present. | Preserved. | README keeps the compile-before-run framing. |
| Why compile intent before agent execution? | Vague intent should become an explicit plan before any handoff. | Present. | Preserved and reinforced through the tutorial. | Runtime execution is not implied. |
| macOS install | Use source, local binary, verified v0.4.0 release binary, or verified v0.4.0 curl installer. | Partially present. | Added a concise macOS install section using the verified curl installer path. | Default installer target is `$HOME/.local/bin/ni`. |
| macOS uninstall | Remove the installed `ni` binary from the chosen install directory. | Not explicit enough. | Added exact default removal command and BINDIR caveat. | PATH cleanup is manual only if the user added a PATH line. |
| Windows install | Use the verified v0.4.0 `windows/amd64` release zip; no package-manager installer is claimed. | Present in install docs, not README. | Added concise PowerShell release-zip flow. | Hash comparison remains manual and visible. |
| Windows uninstall | Remove the copied `ni.exe` from the chosen directory. | Not explicit enough. | Added bounded PowerShell removal example. | Uses the user's actual install path. |
| First-use tutorial | Show init, planning conversation, status proof, blocker resolution, lock, and bounded prompt compile. | Short 60-second path existed. | Added "First project in 5 minutes". | `ni run` is explicitly non-executing. |
| What ni does not do | ni is not a task runner, SPEC runner, runtime, adapter, queue, or release automation system. | Present. | Preserved. | Boundary remains prominent. |
| Links to deeper docs | README should route to install, protocol, benchmark, Homebrew, command, and visual docs. | Present. | Kept concise; docs/116 can be linked from roadmap. | README is not a full manual. |

## README image inventory

| Image reference | Path | Exists? | Current alt text | Purpose | Issue | Recommended action |
| --- | --- | --- | --- | --- | --- | --- |
| README.md hero, README.ko.md hero | `assets/hero.svg` | Yes | `ni hero banner: Project Intent Compiler for AI Agents` | Top visual identity and product framing. | Contains exact product text; raster generation could distort words. | Keep deterministic SVG; if replaced, use text-light art and keep exact copy in Markdown/SVG. |
| README.md language chip, README.ko.md language chip | `assets/badge-english.svg` | Yes | `English` | Link to English README. | Exact text belongs in SVG/Markdown, not raster AI. | Keep deterministic SVG chip. |
| README.md language chip, README.ko.md language chip | `assets/badge-korean.svg` | Yes | `한국어` | Link to Korean README. | Exact Korean text belongs in SVG/Markdown, not raster AI. | Keep deterministic SVG chip. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/license-MIT-f4b860` | Remote | `License MIT` | License signal. | Remote dependency; exact text badge. | Keep or replace with local deterministic SVG; do not AI-generate. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/CI-workflow%20exists-25334a` | Remote | `CI workflow exists` | CI workflow existence signal. | Remote dependency; exact factual claim. | Keep or replace with local deterministic SVG; do not imply CI passed. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/security-policy%20exists-2d5a52` | Remote | `Security policy exists` | Security policy existence signal. | Remote dependency; exact factual claim. | Keep or replace with local deterministic SVG; do not imply audit results. |
| README.md trust badge, README.ko.md trust badge | `https://img.shields.io/badge/docs-index%20exists-5b8def` | Remote | `Docs index exists` | Docs index existence signal. | Remote dependency; exact factual claim. | Keep or replace with local deterministic SVG; do not imply documentation completeness. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-vague-intent.svg` | Yes | Vague intent alt text in each README. | Show hidden missing planning inputs. | Visual contains text-like card framing; exact claims should stay in Markdown. | Keep SVG or replace with text-light abstract card. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-early-execution.svg` | Yes | Early execution alt text in each README. | Show unsafe early work start. | Same as above. | Keep SVG or replace with text-light abstract card. |
| README.md pain card, README.ko.md pain card | `assets/card-pain-rework.svg` | Yes | Rework alt text in each README. | Show cost of hidden assumptions. | Same as above. | Keep SVG or replace with text-light abstract card. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-capture-intent.svg` | Yes | Capture intent alt text in each README. | Show conversation becoming docs and contract. | Same as above. | Keep SVG or replace with text-light abstract card. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-lock-contract.svg` | Yes | Lock contract alt text in each README. | Show readiness and lock gate. | Same as above. | Keep SVG or replace with text-light abstract card. |
| README.md payoff card, README.ko.md payoff card | `assets/card-payoff-handoff-safely.svg` | Yes | Handoff safely alt text in each README. | Show bounded handoff prompt. | Same as above. | Keep SVG or replace with text-light abstract card that does not imply execution. |

## Image-generation prompt pack

No image generation was performed in this task. Prompts below are for a later
visual pass. Exact CLI text, status constants, badges, and product claims should
be rendered in Markdown or deterministic SVG, not AI-generated raster text.

### assets/hero.svg

Purpose: README hero visual identity for compile-before-run.
Recommended use: Optional replacement concept only; keep exact product copy in
Markdown or deterministic SVG.
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
Exact text policy: Put "Don't run the agent yet. Compile the intent first." in
Markdown or deterministic SVG only.
Alt text: Abstract banner showing conversation becoming a locked planning
contract before downstream handoff.
Claim-boundary warning: Must not imply ni executes agents or downstream work.

### assets/badge-english.svg

Purpose: Link to English README.
Recommended use: Keep deterministic SVG or HTML text link.
Suggested dimensions: 84 x 28.
Style: Small flat language chip.
Prompt: Create a simple local-language navigation chip background with no
readable text, suitable for overlaying deterministic SVG text later.
Negative prompt: No flags, no brand marks, no fake status badge, no tiny raster
letters, no national symbolism.
Exact text policy: Render "English" only in deterministic SVG or Markdown.
Alt text: English language link.
Claim-boundary warning: Language navigation only; no product status claim.

### assets/badge-korean.svg

Purpose: Link to Korean README.
Recommended use: Keep deterministic SVG or HTML text link.
Suggested dimensions: 84 x 28.
Style: Small flat language chip.
Prompt: Create a simple local-language navigation chip background with no
readable text, suitable for overlaying deterministic Korean SVG text later.
Negative prompt: No flags, no brand marks, no fake status badge, no tiny raster
letters, no national symbolism.
Exact text policy: Render "한국어" only in deterministic SVG or Markdown.
Alt text: Korean language link.
Claim-boundary warning: Language navigation only; no product status claim.

### license badge

Purpose: License signal.
Recommended use: Keep remote shields badge or replace with local deterministic
SVG; do not use raster AI.
Suggested dimensions: approximately 92 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape that can receive deterministic SVG text.
Negative prompt: No legal seal, no fake certification, no brand marks, no
readable raster text.
Exact text policy: Render "License MIT" only with shields.io, Markdown, or
deterministic SVG.
Alt text: License MIT.
Claim-boundary warning: Shows license existence only.

### CI workflow badge

Purpose: CI workflow existence signal.
Recommended use: Keep remote shields badge or replace with local deterministic
SVG; do not use raster AI.
Suggested dimensions: approximately 150 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape for deterministic status text.
Negative prompt: No green passing checkmark unless a passing run is being
claimed, no fake workflow screenshot, no readable raster text.
Exact text policy: Render "CI workflow exists" only with shields.io, Markdown,
or deterministic SVG.
Alt text: CI workflow exists.
Claim-boundary warning: Must not imply CI passed, only that workflow file exists.

### security policy badge

Purpose: Security policy existence signal.
Recommended use: Keep remote shields badge or replace with local deterministic
SVG; do not use raster AI.
Suggested dimensions: approximately 175 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape for deterministic status text.
Negative prompt: No security certification seal, no audit passed claim, no
shield logo from a third party, no readable raster text.
Exact text policy: Render "Security policy exists" only with shields.io,
Markdown, or deterministic SVG.
Alt text: Security policy exists.
Claim-boundary warning: Must not imply security audit or vulnerability status.

### docs index badge

Purpose: Docs index existence signal.
Recommended use: Keep remote shields badge or replace with local deterministic
SVG; do not use raster AI.
Suggested dimensions: approximately 150 x 20.
Style: Flat factual badge.
Prompt: Not recommended for image generation. If a visual background is needed,
create a text-free flat badge shape for deterministic status text.
Negative prompt: No completeness ribbon, no fake documentation score, no
readable raster text.
Exact text policy: Render "Docs index exists" only with shields.io, Markdown,
or deterministic SVG.
Alt text: Docs index exists.
Claim-boundary warning: Must not imply docs are complete or publication-ready.

### assets/card-pain-vague-intent.svg

Purpose: Visualize hidden planning gaps inside plausible requests.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Abstract document card, missing fields, muted warning accent.
Prompt: Create a text-free visual card showing a plausible request with hidden
empty fields for users, acceptance criteria, risks, non-goals, and blockers.
Use abstract document panels and subtle missing-field indicators. Keep it clean
and language-neutral.
Negative prompt: No readable text, no terminal logs, no fake user data, no
third-party UI, no runtime execution scene.
Exact text policy: Keep labels and explanations in Markdown or deterministic
SVG only.
Alt text: Vague intent can hide missing users, acceptance criteria, risks,
non-goals, or blockers.
Claim-boundary warning: Must not imply ni detects every possible ambiguity.

### assets/card-pain-early-execution.svg

Purpose: Visualize stopping before premature work starts.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Abstract gate before a work lane, restrained warning accent.
Prompt: Create a text-free visual card showing a work lane paused at a gate
before implementation begins, with planning materials waiting for validation.
Use abstract shapes, not literal brand tools.
Negative prompt: No robots executing work, no shell commands, no code editor
screenshot, no fake CI output, no readable text.
Exact text policy: Keep "Early execution" and explanation in Markdown or
deterministic SVG only.
Alt text: Work should not begin just because a request sounds plausible.
Claim-boundary warning: Must not imply ni can run or supervise implementation.

### assets/card-pain-rework.svg

Purpose: Visualize costly rework from hidden assumptions.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Abstract branching plan with one path crossed out and corrected.
Prompt: Create a text-free visual card showing a plan branch being corrected
before costly implementation rework spreads. Use document nodes, a gentle
correction mark, and clear before/after path structure.
Negative prompt: No blame imagery, no fake metrics, no implementation-quality
claim, no readable text.
Exact text policy: Keep "Rework" and explanation in Markdown or deterministic
SVG only.
Alt text: Hidden assumptions become expensive after work starts from the wrong
plan.
Claim-boundary warning: Must not imply benchmarked cost reduction.

### assets/card-payoff-capture-intent.svg

Purpose: Visualize conversation becoming docs and a contract draft.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Conversation bubbles turning into structured pages.
Prompt: Create a text-free visual card showing model-user planning conversation
being organized into docs and a contract draft. Use abstract bubbles, document
cards, and field outlines with no readable text.
Negative prompt: No chat-provider UI, no brand marks, no fake transcript, no
readable text, no claim that a model alone decides readiness.
Exact text policy: Keep exact file paths in Markdown or deterministic SVG only.
Alt text: Planning conversation becomes explicit docs and a contract draft.
Claim-boundary warning: Must preserve that models draft and the CLI validates.

### assets/card-payoff-lock-contract.svg

Purpose: Visualize readiness and lock hash trust.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Contract page, readiness gate, lock, hash line.
Prompt: Create a text-free visual card showing a contract passing through a
readiness gate into a lock with hash-like abstract marks. Keep it precise,
minimal, and language-neutral.
Negative prompt: No fake `READY` output, no real terminal, no fake signature,
no legal seal, no readable text.
Exact text policy: Put `ni status`, `ni end`, `READY`, and hash wording in
Markdown or deterministic SVG only.
Alt text: Readiness and lock commands gate the accepted plan, hashes, and
source of truth.
Claim-boundary warning: Must not imply readiness without `ni status`.

### assets/card-payoff-handoff-safely.svg

Purpose: Visualize bounded prompt compilation after a valid lock.
Recommended use: Text-light README card.
Suggested dimensions: 520 x 320, 13:8.
Style: Locked plan producing a sealed handoff packet, not an execution scene.
Prompt: Create a text-free visual card showing a locked contract producing a
bounded handoff packet. The packet should stop before any runtime or agent work
begins. Use calm developer-tool styling and clear separation between source
contract and downstream packet.
Negative prompt: No robot worker, no shell command, no queue, no PR automation,
no fake downstream success, no readable text.
Exact text policy: Put `ni run --max-chars 4000` in Markdown or deterministic
SVG only.
Alt text: A valid locked plan compiles into a bounded handoff prompt or derived
seed material.
Claim-boundary warning: Must state that `ni run` compiles only and does not
execute downstream work.

## macOS install / uninstall notes

README now recommends the verified v0.4.0 curl installer path for macOS after
script inspection. It documents `install.sh`'s default target as
`$HOME/.local/bin/ni`, notes that `BINDIR` can override the directory, and
verifies with `--help` and `version`.

| Install path | README status | Evidence boundary | Notes |
| --- | --- | --- | --- |
| Source | Available | Local checkout with Go. | Existing README path preserved. |
| Local binary | Available | `make build` / `make install-local`. | Detailed steps stay in docs/22. |
| Release binary | Available | Verified v0.4.0 release assets. | Manual archive path remains documented. |
| Curl installer | Available | Verified v0.4.0 installer path. | Default target is `$HOME/.local/bin/ni`. |
| Homebrew | Planned / v0.5 candidate | No tap/formula/install proof. | README forbids `brew install` for now. |

Uninstall wording removes the exact installed `ni` file from the chosen
directory and treats PATH cleanup as manual only when the user added a line.

## Windows install / uninstall notes

README now documents only the v0.4.0 `windows/amd64` release zip path for
Windows. It does not claim MSI, winget, Chocolatey, Scoop, Homebrew, or a
verified Windows installer.

| Install path | README status | Evidence boundary | Notes |
| --- | --- | --- | --- |
| Release zip `windows/amd64` | Available for v0.4.0 release assets | Archive and checksum file from the same release. | User compares `Get-FileHash` with checksum text. |
| Manual PATH placement | Manual / user-controlled | User chooses the directory. | README says to place `ni.exe` in a controlled directory. |
| Windows package manager | Not claimed | No verification. | winget, Chocolatey, Scoop, MSI are not claimed. |
| Windows arm64 | Not claimed | Installer says Windows arm64 release asset is not configured. | No availability claim. |

Uninstall wording removes the copied `ni.exe` from the actual install directory
and removes PATH only if the user added a ni-specific PATH entry.

## First-use tutorial notes

The README tutorial shows:

1. `ni init`
2. model-user planning conversation before execution
3. `docs/plan/**`, `.ni/contract.json`, and `.ni/session.json`
4. `ni status --proof --next-questions`
5. resolving `BLOCKED` questions or gaps
6. `ni end`
7. `ni run --max-chars 4000`

It states explicitly that `ni run` compiles a bounded handoff prompt from
`.ni/plan.lock.json` and does not execute prompts, agents, shell commands, or
downstream work.

## README.ko companion audit

| Section | English source | Korean companion status | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Hero and definition | README.md top section | Mirrored | Yes | Product definition preserved. |
| Start in 60 seconds | README.md | Mirrored | Yes | Commands unchanged. |
| First project tutorial | README.md | Added | Yes | Commands, paths, and constants preserved. |
| Choose your path | README.md | Already mirrored | Yes | Status rows unchanged. |
| macOS install / uninstall | README.md | Added | Yes | Same verified path and default target. |
| Windows install / uninstall | README.md | Added | Yes | Same no-package-manager boundary. |
| What ni is not | README.md | Preserved | Yes | Non-execution boundary preserved. |
| Read next | README.md | Preserved | Yes | Existing companion links remain. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Published/released status | Do not claim v0.5 has been published. | README/docs116 do not claim v0.5 release. | Yes | No tag, upload, or release action. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | README says Homebrew remains Planned / v0.5 candidate and forbids `brew install`. | Yes | No Homebrew Available claim. |
| Model workspace packs | Model workspace packs: Experimental. | README path table remains Experimental. | Yes | Skills are UX; CLI is authority. |
| No-terminal | No-terminal method: Experimental / assisted. | README path table remains Experimental. | Yes | No deterministic no-terminal claim. |
| ni run | Prompt compilation only. | Tutorial and payoff copy say no execution. | Yes | No downstream work execution. |
| READY | Planning contract readiness only. | Tutorial says CLI decides readiness; docs116 does not claim product readiness. | Yes | `READY` is not product readiness. |
| LOCK-STALE | Existing lock no longer matches current planning inputs. | No new stale-lock claim added. | Yes | Existing docs remain source for recovery. |
| Benchmark evidence | No implementation-quality or downstream-quality proof. | docs116 prompt pack avoids benchmark overclaim. | Yes | No empirical effect claim. |
| Runtime execution boundary | ni-kernel stops before runtime execution. | README "What ni is not" preserved. | Yes | No task runner or harness added. |
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
| docs/115_* | no final diff | tracked | No | Existing baseline records the later publication-prep decision; not edited here. |
| docs/116_* | Added | untracked until staged | Yes | Required output for this pass. |
| docs/51* | Modified | tracked | Yes | Narrow docs/116 roadmap pointer only. |
| image assets | no new image files | tracked existing assets | No | No generated images added. |
| generated artifacts | none expected | not applicable | No | Validation may create ignored temp/build output. |
| .ni/contract.json | no diff expected | tracked | No | Protected root file. |
| .ni/session.json | no diff expected | tracked | No | Protected root file. |
| .ni/plan.lock.json | no diff expected | tracked | No | Protected root lockfile. |
| unexpected files | none | final status reviewed | No | Final visible changes are README, docs/51, and docs/116. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Pass | Final visible changes: README.md, README.ko.md, docs/51*, and untracked docs/116*. |
| `git ls-files docs/110_... docs/115_...` | Pass | docs/110_* through docs/115_* are tracked. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | Uses temp Go cache workaround. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Pass | `NI Intent Readiness: READY`; blockers, deferrals, warnings: None. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass | Help output rendered. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass | Output: `0.0.0-dev`; source build only. |
| `python3 scripts/check-install-docs.py` | Pass | Install claim boundaries preserved. |
| `bash scripts/check-skill-packs.sh` | Pass | Model workspace packs remain Experimental; global install remains unverified. |
| `bash scripts/demo-check.sh` | Pass | Temporary prompts/exports compiled as seed or prompt artifacts only; not executed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Broad quality wrapper passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Exercises fixture `ni end` / `ni relock`; not project-root relock. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Source, build, and temporary install paths passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Check-only release readiness gate passed; no release action. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Pass | No protected project-root `.ni` diff. |

## Changes made

- `README.md`: added first-use tutorial, macOS install/uninstall, and Windows
  install/uninstall.
- `README.ko.md`: added matching Korean companion sections without expanding
  status claims.
- `docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.md`: added this audit and
  prompt pack.
- `docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.ko.md`: Korean companion
  maintained separately.
- `docs/51_POST_RELEASE_ROADMAP.md` / `.ko.md`: optional narrow pointer if
  included in final diff.

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

Why: README onboarding is improved and image prompts are complete enough for a
later visual pass, but the highest publication-facing risk is still public
install wording across README, docs/22, curl installer docs, and release status.
A focused install-doc verification pass can re-check macOS and Windows wording
before publication approval is revisited.

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
