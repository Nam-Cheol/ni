# External User Validation Plan

## Current status

State:
- v0.5.1 release: published and verified.
- v0.5.1 release decision: V0_5_1_RELEASE_EXECUTED_WITH_NOTES.
- Public install parity mismatch from v0.5.0: closed for tested macOS arm64 path.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: deferred until Windows transcript exists.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Plan goal

This plan defines how to validate whether a real external user can follow the
public v0.5.1 README install and first-project onboarding path without
maintainer assistance. It is a protocol, transcript template, evidence grading
model, and claim-boundary checklist. It does not itself prove that external
users succeeded unless real tester transcripts are attached in a later audit.

Primary question: can an external user open the README, choose macOS or
Windows, install `ni`, verify the global command in a new shell or PowerShell,
create a first project with `ni init .`, run `ni status --proof
--next-questions`, understand that `ni run` compiles a bounded handoff prompt
only, and uninstall if asked?

## Decision

EXTERNAL_VALIDATION_PLAN_READY_WITH_NOTES

Justification: the external validation protocol, command sheets, transcript
template, evidence grading, pilot packet, and claim-boundary audit are ready.
Notes remain because no external user transcript is included in this task,
Windows real-host validation is still pending, and tester recruitment remains
open.

## Validation goals

| Goal | Evidence required | Pass criteria | Notes |
| --- | --- | --- | --- |
| README clarity | Tester starts from README only and records what was confusing. | Tester can choose a platform path without maintainer explanation. | Feedback may still produce polish tasks. |
| macOS install success | Full macOS transcript with install command, new shell, command-name checks, and uninstall if tested. | `ni --help` exits successfully and `ni version` reports `0.5.1` from command name. | Required for macOS external pilot. |
| Windows install success | Full Windows transcript from a real Windows host, VM, CI runner, or external tester. | `Get-Command ni`, `ni --help`, and `ni version` work in a new PowerShell session. | Required to close Windows real-host deferral. |
| Command-name global resolution | `command -v ni` or `Get-Command ni` after a new shell. | The shell resolves `ni` without an absolute binary path. | This is distinct from local artifact execution. |
| `ni --help` | Command transcript. | Help exits successfully and shows the implemented CLI surface. | Do not use help success as first-project proof. |
| `ni version` | Command transcript. | Output is `0.5.1`. | Source-tree `0.0.0-dev` is not acceptable for external release validation. |
| `ni init .` | Transcript plus created-file listing. | Guided init starts and completes or exits safely without unexpected overwrite. | TUI usability feedback is required. |
| TUI usability | Tester notes, screenshots only if safe, and any stuck point. | Tester can navigate, complete, or intentionally exit the Bubble Tea v2 / Lip Gloss v2 flow. | Non-interactive fallback may be tested only if relevant. |
| Generated planning files | `find` or directory listing. | `.ni/contract.json`, `.ni/session.json`, and `docs/plan/**` exist; `.ni/plan.lock.json` is not created by init. | Use a fixture project, not the real repo root. |
| Status comprehension | `ni status --proof --next-questions` transcript and tester answers. | Tester understands basic `BLOCKED` versus `READY` and that `READY` is CLI planning readiness only. | Blank or incomplete first projects may honestly be `BLOCKED`. |
| `ni run` comprehension | Tester paraphrase; optional command only after understanding. | Tester understands `ni run` compiles a bounded prompt and does not execute downstream work. | Optional command must use `--max-chars 4000`. |
| Uninstall path | Uninstall transcript or tester note locating docs. | Tester can uninstall, or can identify the relevant uninstall command. | Uninstall should remove installer-managed binary/PATH only. |
| Claim-boundary comprehension | Tester answers and transcript review. | Tester does not conclude Homebrew is Available, Windows is verified without transcript, no-terminal is deterministic, or `ni` runs agents. | Boundary failures are product/doc failures. |

## Validation cohorts

| Cohort | Why needed | Required? | Status | Notes |
| --- | --- | --- | --- | --- |
| Maintainer sanity check | Confirms release path before asking external users. | No for this plan; already done separately. | Done for tested macOS arm64 path. | See docs/132 post-release verification. |
| External macOS user | Tests the public README path on a separate user or machine. | Yes for macOS external validation. | Future validation. | Prefer a user who has not watched maintainer setup. |
| External Windows user | Closes Windows real-host execution deferral. | Yes to claim Windows real-host execution. | Future validation. | Static PowerShell checks are not enough. |
| Backend developer familiar with CLI tools | Finds friction that technical users still hit. | Optional. | Future validation. | Useful but not required before the plan is ready. |
| AI-agent user who is not a Go developer | Tests whether the product story lands with the intended audience. | Optional. | Future validation. | Useful for README and onboarding language. |

## Success criteria

| Criterion | Required? | Evidence | Pass state |
| --- | --- | --- | --- |
| User installs without maintainer intervention. | Yes | Transcript and tester note. | PASS or PASS_WITH_NOTES. |
| User opens a new shell or PowerShell and runs `ni` by command name. | Yes | `command -v ni` or `Get-Command ni`. | PASS. |
| `ni --help` succeeds. | Yes | Command transcript. | PASS. |
| `ni version` reports `0.5.1`. | Yes | Command transcript. | PASS. |
| User runs `ni init .` in a fixture project. | Yes | Command transcript. | PASS or PASS_WITH_NOTES. |
| User can navigate or complete the Bubble Tea v2 / Lip Gloss v2 TUI. | Yes | Transcript and feedback. | PASS or PASS_WITH_NOTES. |
| User can identify created files. | Yes | File listing. | PASS. |
| User can run `ni status --proof --next-questions`. | Yes | Command transcript. | PASS. |
| User understands `BLOCKED` versus `READY` at a basic level. | Yes | Feedback answer. | PASS or PASS_WITH_NOTES. |
| User understands `ni run` does not execute downstream work. | Yes | Feedback answer; optional command. | PASS. |
| User can uninstall or knows where uninstall docs are. | Yes | Command transcript or note. | PASS or PASS_WITH_NOTES. |
| No protected `.ni` files in the real project root are modified. | Yes | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json`. | Empty diff. |

## Failure criteria

| Failure criterion | Severity | Evidence | Blocker condition |
| --- | --- | --- | --- |
| README install instructions are not enough. | High | Tester cannot choose or complete a path. | FAIL_BLOCKER if maintainer explanation is required. |
| Install script fails. | High | Installer transcript. | FAIL_BLOCKER unless caused by local policy outside README scope and documented. |
| PATH/global command resolution fails. | High | `command -v ni` or `Get-Command ni`. | FAIL_BLOCKER. |
| `ni version` is not `0.5.1`. | High | Version transcript. | FAIL_BLOCKER. |
| `ni init .` fails. | High | Command transcript. | FAIL_BLOCKER unless user exited intentionally and no corruption occurred. |
| TUI is confusing or unusable. | Medium to high | Tester feedback. | FAIL_BLOCKER if user cannot proceed or exit safely. |
| TUI overwrites files unexpectedly. | High | File hashes or listing. | FAIL_BLOCKER. |
| `.ni/plan.lock.json` is created by init. | High | File listing. | FAIL_BLOCKER. |
| `ni status` makes user think product readiness is proven. | High | Tester feedback. | FAIL_BLOCKER for claim-boundary comprehension. |
| README implies Homebrew Available. | High | README audit. | FAIL_BLOCKER until corrected. |
| Windows tester cannot complete install due to PATH or PowerShell policy issue. | Medium to high | PowerShell transcript. | FAIL_BLOCKER for Windows real-host claim; may be PASS_WITH_NOTES for general plan if documented. |
| User thinks `ni` runs the agent or implements the project. | High | Tester feedback. | FAIL_BLOCKER for product comprehension. |

## Transcript template

```text
External ni v0.5.1 validation transcript

Tester:
Date:
Platform:
OS version:
Shell / PowerShell version:
Architecture:
Install method:
README URL used:
Release URL used:

Privacy note:
- You may redact usernames, home paths, machine names, project names, and tokens.
- Do not include secrets, credentials, private repository URLs, API keys, or personal tokens.

Command transcript:

1. Install command
<paste command and output>

2. New shell / PowerShell command-name check
macOS:
$ command -v ni
$ ni --help
$ ni version

Windows:
PS> Get-Command ni
PS> ni --help
PS> ni version

<paste command and output>

3. Create project directory
$ mkdir ni-external-validation
$ cd ni-external-validation

or:
PS> mkdir ni-external-validation
PS> cd ni-external-validation

<paste command and output>

4. First project init
$ ni init .

or:
PS> ni init .

<paste command and output>

5. Observed files
Expected:
- .ni/contract.json
- .ni/session.json
- docs/plan/**
- .ni/plan.lock.json should not be created by init

<paste file listing>

6. Status proof
$ ni status --proof --next-questions

or:
PS> ni status --proof --next-questions

<paste command and output>

7. Optional prompt compilation
Only run this if you understand that it compiles a bounded handoff prompt and
does not execute agents, shell commands, or downstream work.

$ ni run --max-chars 4000

or:
PS> ni run --max-chars 4000

<paste command and output, or write NOT_RUN>

8. Uninstall, if tested
macOS:
$ BINDIR="$HOME/.local/bin" sh install.sh --uninstall

Windows:
PS> .\install.ps1 -Uninstall

<paste command and output, or write NOT_RUN>

User feedback:
- What was confusing?
- Did README explain enough?
- Did TUI explain enough?
- Did you understand what ni does not do?
- Did you understand `BLOCKED` versus `READY`?
- Did you understand that `ni run` does not execute downstream work?
- Did any command fail?
```

## macOS command sheet

```bash
VERSION="0.5.1"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

Open a new shell:

```bash
command -v ni
ni --help
ni version
```

Create a fixture project:

```bash
mkdir ni-external-validation
cd ni-external-validation
ni init .
find . -maxdepth 3 -type f | sort
ni status --proof --next-questions
```

Optional prompt compilation, only after the user understands that `ni run`
compiles a prompt and does not execute downstream work:

```bash
ni run --max-chars 4000
```

Uninstall, if tested:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Windows command sheet

```powershell
$Version = "0.5.1"
Invoke-WebRequest "https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1" -OutFile "install.ps1"
Get-Content .\install.ps1
.\install.ps1 -DryRun -Version $Version
.\install.ps1 -Version $Version
```

Open a new PowerShell session:

```powershell
Get-Command ni
ni --help
ni version
```

Create a fixture project:

```powershell
mkdir ni-external-validation
cd ni-external-validation
ni init .
Get-ChildItem -Recurse -File | Sort-Object FullName | Select-Object FullName
ni status --proof --next-questions
```

Optional prompt compilation, only after the user understands that `ni run`
compiles a prompt and does not execute downstream work:

```powershell
ni run --max-chars 4000
```

Uninstall, if tested:

```powershell
.\install.ps1 -Uninstall
```

Windows notes:
- The installer should use User PATH by default, not Machine PATH.
- Do not claim Windows real-host execution is verified until this transcript exists.
- Capture PowerShell execution policy issues if they appear.

## Evidence grading

| Grade | Meaning | Allowed claim | Forbidden claim |
| --- | --- | --- | --- |
| PASS | Required evidence was captured and met the criteria. | The tested path passed for the shown platform, host, version, and transcript. | All users, all platforms, or untested install paths succeeded. |
| PASS_WITH_NOTES | The path completed but had non-blocking friction, deferrals, or comprehension notes. | The tested path completed with stated notes. | The notes are resolved or irrelevant. |
| FAIL_BLOCKER | A required step failed or produced a claim-boundary failure. | The validation found a blocker in the scoped path. | The product is broadly unusable beyond the scoped evidence. |
| NOT_RUN | The step was not attempted. | No evidence was collected for this step. | Any success or failure claim for that step. |

Evidence categories:
- install
- command resolution
- version
- init TUI
- generated files
- status comprehension
- uninstall
- claim-boundary comprehension

## Pilot validation packet

Send this packet to testers:

```text
Please validate ni v0.5.1 from the public README.

Start here:
- README: https://github.com/Nam-Cheol/ni
- Release: https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1

Use the command sheet for your platform:
- macOS: install.sh path from README
- Windows: install.ps1 path from README

Expected timebox: 20 to 30 minutes.

Please paste the transcript template results. You may redact usernames, home
paths, machine names, project names, and tokens. Do not include secrets.

Important product boundary:
- ni compiles project intent before execution.
- ni run compiles a bounded handoff prompt only.
- ni does not execute agents, shell commands, downstream work, PR automation,
  release automation, queues, or task-runner behavior.
- Homebrew remains Planned / v0.5 candidate.
- Windows real-host execution is not verified until a Windows transcript exists.
```

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| external validation | Do not claim external users succeeded before transcripts exist. | This plan contains no external transcript. | Yes | Decision remains WITH_NOTES. |
| Windows real-host execution | Must remain deferred without Windows transcript. | Deferred. | Yes | Static PowerShell checks are not real-host execution. |
| Homebrew | Must remain Homebrew: Planned / v0.5 candidate. | Preserved. | Yes | No Homebrew Available claim. |
| `ni run` | Bounded prompt compilation only. | Preserved. | Yes | Optional pilot command is explicitly non-executing. |
| READY | CLI planning readiness only, not product readiness. | Preserved. | Yes | Tester must explain `BLOCKED` versus `READY`. |
| TUI | Guided first-project intent setup, not downstream execution. | Preserved. | Yes | TUI usability is tested as onboarding UX. |
| runtime execution boundary | No task runner, SPEC runner, execution harness, shell adapter, Codex exec adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | This task adds docs only. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | clean | No | Read only in this task. |
| README.ko.md | clean | No | Read only in this task. |
| docs/132* | clean | No | Source evidence for current status. |
| docs/133* | added | Yes | External validation plan and Korean companion. |
| docs/51* | modified | Yes | Narrow roadmap pointer to docs/133. |
| `.ni/contract.json` | clean | No | Protected. |
| `.ni/session.json` | clean | No | Protected. |
| `.ni/plan.lock.json` | clean | No | Protected. |
| unexpected files | none expected | No | Recheck before commit. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; only docs/51 modifications and docs/133 additions were present. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; root reported `NI Intent Readiness: READY` with no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

## Changes made

| File | Why |
| --- | --- |
| `docs/133_EXTERNAL_USER_VALIDATION_PLAN.md` | Adds the external validation protocol, transcript template, platform command sheets, evidence grading, pilot packet, and claim-boundary audit. |
| `docs/133_EXTERNAL_USER_VALIDATION_PLAN.ko.md` | Korean companion with the same claim boundaries. |
| `docs/51_POST_RELEASE_ROADMAP.md` | Adds a narrow pointer to this external validation plan. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Korean roadmap companion pointer. |

## What this plan proves

- The external validation protocol is ready.
- Claim boundaries are explicit.
- Platform-specific evidence requirements are defined.
- The transcript template and pilot packet are ready for testers.

## What this plan does not prove

- External users succeeded.
- Windows real-host execution works.
- Homebrew is Available.
- Downstream execution succeeds.
- No-terminal is deterministic.
- Benchmark effect size or causal impact.

## Recommended next task

Selected next task: A. run macOS external pilot validation.

Selection rationale: v0.5.1 publication and tested macOS arm64 install parity
are already documented, but no independent external user transcript exists.
The lowest-friction next evidence gap is a macOS external pilot. Windows pilot
validation remains the separate required path for closing Windows real-host
execution.

## Next task prompt

```text
Task: run one macOS external pilot validation for ni v0.5.1.

Do not ask the tester to use private maintainer knowledge. Send them only the
README link, v0.5.1 release link, macOS command sheet, transcript template, and
privacy note from docs/133_EXTERNAL_USER_VALIDATION_PLAN.md.

Tester packet:
- README: https://github.com/Nam-Cheol/ni
- Release: https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1
- Expected timebox: 20 to 30 minutes.
- The tester may redact usernames, home paths, machine names, project names,
  and tokens.
- The tester must not paste secrets.

Ask the tester to run:
1. Inspect and install with install.sh for VERSION=0.5.1.
2. Open a new shell.
3. Run command-name `command -v ni`, `ni --help`, and `ni version`.
4. Create a fixture project directory.
5. Run `ni init .`.
6. List created files and confirm `.ni/plan.lock.json` was not created by init.
7. Run `ni status --proof --next-questions`.
8. Explain in their own words that `ni run` compiles a bounded prompt and does
   not execute downstream work.
9. Optionally run `ni run --max-chars 4000` only after that explanation.
10. Test uninstall or identify the uninstall command.

After the transcript arrives, create a result audit. Grade install, command
resolution, version, init TUI, generated files, status comprehension, uninstall,
and claim-boundary comprehension as PASS, PASS_WITH_NOTES, FAIL_BLOCKER, or
NOT_RUN. Do not claim external validation succeeded unless the transcript
supports it. Do not claim Windows real-host execution, Homebrew Available,
no-terminal deterministic validation, downstream execution success, or product
readiness from this macOS transcript.
```
