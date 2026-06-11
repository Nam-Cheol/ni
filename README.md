![namba-intent hero](assets/namba-intent-hero.png)

# namba-intent

<p align="center">
  <a href="README.md" aria-label="Read in English"><img alt="English" src="assets/badge-english.svg" width="84" height="28"></a>
  <a href="README.ko.md" aria-label="Read in Korean"><img alt="Korean" src="assets/badge-korean.svg" width="84" height="28"></a>
</p>

<p align="center">
  <a href="LICENSE"><img alt="License MIT" src="https://img.shields.io/badge/license-MIT-f4b860"></a>
  <a href=".github/workflows/ci.yml"><img alt="CI workflow exists" src="https://img.shields.io/badge/CI-workflow%20exists-25334a"></a>
  <a href="SECURITY.md"><img alt="Security policy exists" src="https://img.shields.io/badge/security-policy%20exists-2d5a52"></a>
  <a href="docs/00_START_HERE.md"><img alt="Docs index exists" src="https://img.shields.io/badge/docs-index%20exists-5b8def"></a>
</p>

**Don't run the agent yet. Compile the intent first.**

namba-intent is a Project Intent Compiler for AI Agents.

namba-intent helps turn rough project intent into a reviewed, structured,
locked intent before downstream work begins.

v0.6.4 is the current published namba-intent release. The primary command is
`namba-intent`; the legacy `ni` command is a deprecated shim only.

## Workflow

![namba-intent flow](assets/namba-intent-flow.png)

- Capture Intent: collect the rough goal, users, constraints, risks, and blockers.
- Clarify Context: turn vague or missing context into explicit planning questions.
- Draft Structure: write the intent into docs and a machine-readable contract.
- Review & Refine: use deterministic status proof and user review to improve the plan.
- Lock Intent: lock the accepted intent before downstream work begins.

## Why namba-intent

AI agents move fast. namba-intent slows down only the part that should be slow:
deciding what the project actually means before implementation starts.

- Capture missing users, acceptance criteria, risks, non-goals, and blockers.
- Check planning readiness with deterministic CLI rules.
- Lock the accepted plan and hash the trusted planning sources.
- Compile a short prompt for downstream actors without executing it.

## Install

README shows two primary first-success paths for the current tree. Source,
local build, release archive, pinned installs, dry-run, inspect-first,
`BINDIR`, uninstall details, and the v0.5.1 public-release distinction live in
[Install namba-intent](docs/22_INSTALL.md).

### macOS

Install the latest namba-intent release with the curl installer:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --update-path
```

Open a new shell after installation, then verify the command:

```bash
namba-intent --help
namba-intent version
```

Homebrew: Planned / v0.5 candidate.

### Windows

The PowerShell installer is configured to install `namba-intent.exe` to
`%LOCALAPPDATA%\namba-intent\bin` by default and update User PATH only. Windows
real-host verification remains pending until a Windows transcript exists.

```powershell
$Installer = Join-Path $env:TEMP "namba-intent-install.ps1"
irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer
powershell -NoProfile -ExecutionPolicy Bypass -File $Installer
```

Open a new PowerShell session and verify:

```powershell
namba-intent --help
namba-intent version
```

PowerShell alias cleanup for `ni -> New-Item` is legacy v0.5.x guidance and is
not required for `namba-intent.exe`. Real-host Windows execution remains
deferred until a Windows transcript exists.

## First project in 5 minutes

```bash
mkdir my-project
cd my-project
namba-intent init .
namba-intent status --proof --next-questions
namba-intent end
namba-intent run --max-chars 4000
```

`namba-intent init .` opens a guided setup and creates `.ni/contract.json`,
`.ni/session.json`, and `docs/plan/**`. Your first answers can be rough or
incomplete; use `status` to find the missing pieces. namba-intent keeps
`.ni/` for compatibility.

`namba-intent status --proof --next-questions` is the CLI-authoritative
readiness gate. A model or skill can help draft docs and answer next questions,
but the CLI decides readiness.

`namba-intent end` locks the accepted plan and writes `.ni/plan.lock.json` only
after the CLI gate reports READY and you agree with the plan.

`namba-intent run --max-chars 4000` compiles a bounded downstream handoff
prompt. It does not execute the prompt, run agents, run shell commands, or
prove product readiness.

Model assistance is optional: ask an AI assistant to read `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json`, then help refine the plan. Skills
are UX; CLI is authority.

## What namba-intent Does

| Command | Role |
| --- | --- |
| `namba-intent init .` | Create a planning workspace and guided intent draft. |
| `namba-intent status --proof --next-questions` | Check readiness, blockers, and next planning questions. |
| `namba-intent end` | Lock the accepted plan through the CLI gate. |
| `namba-intent run --max-chars 4000` | Compile a bounded prompt from a valid lock. |

## What namba-intent Does Not Do

namba-intent is not a task runner, SPEC runner, multi-agent execution layer,
queue, shell adapter, PR automation system, release automation system, or
downstream execution runtime.

## Status

- v0.5.1 publication: verified for the historical `ni` command.
- v0.6.4 release: published and verified for `namba-intent` on macOS darwin/arm64.
- Primary command in current tree: `namba-intent`.
- Deprecated transition shim: `ni` warns `ni is deprecated; use namba-intent.`
- Repository: `Nam-Cheol/ni` retained.
- Config directory: `.ni/` retained.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: deferred until a Windows transcript exists.
- Model workspace packs: Experimental. Host-level/global install remains unverified unless documented.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.

## Read next

| Read | Why |
| --- | --- |
| [Install namba-intent](docs/22_INSTALL.md) | Detailed install, release binary, curl installer, and uninstall paths. |
| [Rename implementation](docs/136_NAMBA_INTENT_RENAME_IMPLEMENTATION.md) | v0.6.0 command rename surfaces and claim boundaries. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | Readiness, locking, hash trust, and blocked handoff rules. |
| [No-Terminal Planning](docs/no-terminal.md) | Assisted workflow boundaries; not deterministic validation. |
| [Model Workspace Status](docs/99_MODEL_WORKSPACE_STATUS.md) | Experimental model workspace boundaries and verification state. |

License: namba-intent is licensed under the [MIT License](LICENSE).
