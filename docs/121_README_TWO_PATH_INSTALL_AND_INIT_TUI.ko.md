# README Two-Path Install and Init TUI

## Current status

- v0.5.0 publication: verified
- macOS/Linux global install path handling: implemented
- Windows User PATH installer: implemented
- Windows real-host execution: macOS-only development host에서는 deferred
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Pass goal

이번 pass는 README onboarding을 단순화하고 `ni init .`을 개선한다. New user는
두 primary install path를 보고, global command name으로 `ni`를 verify한 뒤,
current directory에서 guided project intent setup을 시작할 수 있어야 한다. 이
flow는 downstream execution과 혼동되면 안 된다.

## README install simplification

| Area | Before | After | Notes |
| --- | --- | --- | --- |
| macOS path | README가 source, local binary, release binary, curl installer, model workspace, no-terminal, Homebrew rows를 첫 성공 path 전에 함께 보여줬다. | README는 macOS primary path 하나를 보여준다: `install.sh` inspect, dry-run, install, 새 shell, `ni --help`, `ni version`, `ni init .`. | Source/release/local build alternatives는 `docs/22_INSTALL.md`로 이동. |
| Windows path | README가 broad path matrix 뒤에 긴 installer section을 포함했다. | README는 Windows primary path 하나를 보여준다: `install.ps1` inspect, dry-run, install, 새 PowerShell, `ni --help`, `ni version`, `ni init .`. | Windows real-host execution은 transcript 전까지 deferred. |
| Homebrew | Homebrew가 README path matrix에 Planned로 있었다. | README는 `Homebrew: Planned / v0.5 candidate`만 언급하고 package-manager details는 docs로 보낸다. | Homebrew Available claim 없음. |
| advanced install docs | README가 install manual 상당 부분을 포함했다. | README는 source, local build, Linux, release archive, advanced uninstall details를 `docs/22_INSTALL.md`로 링크한다. | README는 product pamphlet로 유지. |
| uninstall | macOS/Windows uninstall step이 긴 install section 안에 있었다. | 각 primary path가 uninstall command를 포함한다. | Uninstall은 installer-managed binary/PATH entry로 제한. |
| verification | Verification이 여러 section에 흩어져 있었다. | 두 primary path 모두 새 shell/session 뒤 `ni --help`, `ni version`, `ni init .`을 요구한다. | Global command verification이 install success criterion으로 유지된다. |

## First-use tutorial

README는 첫 project를 다음 flow 중심으로 둔다:

```bash
mkdir my-project
cd my-project
ni init .
ni status --proof --next-questions
ni end
ni run --max-chars 4000
```

이 flow는 non-execution boundary를 보존한다. `ni init .`은 initial intent
artifacts를 만들고, `ni status --proof --next-questions`는 readiness를 검사하며,
`ni end`는 accepted plan을 lock하고, `ni run --max-chars 4000`은 bounded
downstream handoff prompt를 compile한다. `ni run`은 downstream work를 실행하지
않는다.

## ni init . TUI behavior

| Behavior | Expected | Implemented | Notes |
| --- | --- | --- | --- |
| current-directory target | `ni init .`이 existing project directory에서 동작한다. | Implemented. | Positional target support를 추가했고 `--dir`는 유지했다. |
| interactive prompts | Interactive stdin에서 minimal project-intent setup을 guide한다. | Implemented. | Project name, goal, audience, eventual downstream task, constraints/non-goals, success criteria, blockers, deferrals를 묻는다. |
| non-interactive behavior | Existing scripts와 non-interactive runs는 stable해야 한다. | Implemented. | `--dir`는 기본적으로 non-interactive이고, `--interactive`로 opt in한다. |
| generated files | Init이 `.ni/contract.json`, `.ni/session.json`, `docs/plan/**`을 쓴다. | Existing `docstore.InitWithOptions`를 통해 구현. | Guided fields는 parallel contract writer가 아니라 existing writer로 전달된다. |
| existing file handling | Existing planning files를 silently overwrite하지 않는다. | Implemented. | Existing paths를 보고하고 non-interactive init은 missing files only를 추가한다. Interactive mode는 keep existing, add missing files only, abort를 제공한다. |
| lockfile safety | Existing `.ni/plan.lock.json`은 수정하지 않는다. | Implemented. | Lockfile이 있으면 warning을 출력하고 status/amend/relock flow를 안내한 뒤 write 없이 종료한다. |
| abort behavior | User가 write 전에 abort할 수 있다. | Implemented. | Guided artifacts를 쓰기 전에 interactive confirmation을 한다. |

## Tests added

- `TestInitCurrentDirectoryTarget`
- `TestInitNonInteractiveBehaviorIsStable`
- `TestInitInteractiveFlowFromStdin`
- `TestInitDoesNotSilentlyOverwriteExistingFiles`
- `TestInitDoesNotModifyExistingLockfile`
- `TestInitDoesNotCreateDownstreamExecutionArtifacts`

## What this pass proves

- README now presents two primary install paths
- `ni init .` onboarding was implemented or improved for tested behavior
- global command verification remains the install success criterion
- `ni run` remains bounded prompt compilation only

## What this pass does not prove

- Windows real-host execution works unless a Windows transcript exists
- Homebrew is Available
- downstream execution succeeds
- no-terminal is deterministic
- external users succeed

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS install | Available path는 global command name verification을 요구해야 한다. | README와 install docs가 새 shell 뒤 `ni --help`, `ni version`을 요구한다. | Yes | Local validation은 macOS host와 temporary paths를 사용한다. |
| Windows install | Installer code는 문서화할 수 있지만 real-host execution은 claim하면 안 된다. | README는 real-host execution이 macOS-only host에서 deferred라고 말한다. | Yes | `scripts/check-install-ps1.py`는 static safety proof only. |
| Homebrew | Planned / v0.5 candidate only. | README와 docs는 Homebrew Planned를 유지한다. | Yes | `brew install` path 없음. |
| ni init . | Guided intent setup only. | CLI는 planning artifacts를 쓰고 agents를 실행하지 않는다. | Yes | Existing writer를 재사용한다. |
| ni run | Bounded prompt compilation only. | README는 downstream work를 실행하지 않는다고 말한다. | Yes | Runtime behavior 추가 없음. |
| READY | CLI readiness only, not product readiness. | README는 readiness를 `ni status`에 묶어 둔다. | Yes | Model judgment는 readiness가 아니다. |
| Model workspace packs | Experimental. | README는 model workspace docs로 링크하고 packs를 Available로 표시하지 않는다. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | README는 deterministic validation claim 없이 no-terminal docs로 링크한다. | Yes | Trusted CLI proof가 계속 필요하다. |
| Runtime execution boundary | Task runner, shell adapter, Codex exec adapter, queue, downstream execution 금지. | README와 code가 non-execution boundary를 유지한다. | Yes | Execution adapter 추가 없음. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | M | Yes | Two-path install and first project tutorial. |
| README.ko.md | M | Yes | Korean companion mirrors README.md. |
| cmd/ni/* | M | Yes | `ni init .` target, guided prompts, tests. |
| internal/* | M | Yes | Guided init fields in existing docstore writer. |
| docs/22_INSTALL.md | M | Yes | Detailed install matrix and primary path handoff. |
| docs/install-curl* | M | Yes | Post-install `ni init .` step. |
| docs/120* | unchanged | No | Existing global install acceptance remains current. |
| docs/121* | A | Yes | This audit and Korean companion. |
| install.sh | unchanged | No | Installer shell behavior 변경 없음. |
| install.ps1 | unchanged | No | PowerShell installer behavior 변경 없음. |
| scripts/* | M | Yes | README/install/release checkers를 two-path README에 맞췄고, `demo-check.sh`는 writable default Go cache를 사용한다. |
| .ni/contract.json | unchanged | No | Protected project-root planning state. |
| .ni/session.json | unchanged | No | Protected project-root planning state. |
| .ni/plan.lock.json | unchanged | No | Protected project-root lockfile. |
| unexpected files | none observed | No | Commit 전 재확인. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; expected README, docs, Go, scripts changes only. |
| `gofmt -w .` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; `NI Intent Readiness: READY`, blockers/deferrals/warnings none. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build는 `0.0.0-dev`를 출력. |
| Temp cwd `ni init . --yes` via a built temporary binary | Passed; `.ni/contract.json`, `.ni/session.json`, `docs/plan/00_project_brief.md` 생성. |
| `python3 scripts/check-install-docs.py` | Passed |
| `python3 scripts/check-install-ps1.py` | Passed |
| `bash scripts/check-skill-packs.sh` | Passed |
| `bash scripts/demo-check.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; output 없음. |

## Changes made

- README와 README.ko.md를 macOS/Windows primary install paths 중심으로 단순화했다.
- `ni init .` positional target handling과 small guided interactive init flow를 추가했다.
- Script용 non-interactive `--dir` initialization은 stable하게 유지했다.
- Init의 existing-file과 lockfile safety handling을 추가했다.
- Install docs에 first-use `ni init .` handoff를 추가했다.
- README가 detailed path matrix를 요구하지 않도록 install docs checker를 업데이트했다.

## Recommended next task

Selected: A. generate README visual assets from docs/116 prompts

README/TUI onboarding이 다음 product-facing pass를 진행할 만큼 stable하고, Windows
real-host verification은 이 macOS-only development host에서 unavailable이기 때문에
이 task를 선택한다.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: Generate and verify README visual assets from docs/116 prompts.

Read AGENTS.md, README.md, README.ko.md, docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.md, docs/121_README_TWO_PATH_INSTALL_AND_INIT_TUI.md, and the current assets/ files.

Goal:
Improve README visuals without changing install availability claims or execution boundaries.

Rules:
- Do not publish, tag, release, upload, or push.
- Do not edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json.
- Do not run ni end or relock the project root.
- Do not imply Homebrew is Available.
- Do not imply Windows real-host execution is verified.
- Do not imply ni run executes downstream work.
- Keep exact product/status text in Markdown or deterministic SVG, not AI-generated raster text.

Work:
1. Audit existing README image references and confirm which files exist.
2. Generate or update only visual assets that can be verified locally.
3. Keep README.md and README.ko.md alt text accurate and claim-safe.
4. Add a short docs note with assets changed, verification, and claim-boundary audit.
5. Run README/rendering and quality checks available in the repo.

Validation:
- git status --short
- python3 scripts/check-install-docs.py
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, visual assets changed, README/README.ko status, claim-boundary audit, validation results, protected .ni diff, and confirmation that no publish/tag/release/upload/project-root relock/generated prompt execution occurred.
```
