# Public Announcement Kit

This kit collects copy that can be pasted into public launch surfaces for `ni`.
It does not post to social channels, publish a release, create tags, upload
assets, run agents, or automate distribution.

## Core Copy

One-liner:

> Don't run the agent yet. Compile the intent first.

Short product description:

> ni turns planning conversations into locked project contracts before AI agents
> or teams start work.

Category:

`ni` is a Project Intent Compiler for AI Agents. It is a pre-runtime intent
contract layer: it helps create, validate, lock, and compile planning contracts
before downstream execution begins.

## Boundary Warnings

Use these warnings in public launch material:

- `ni` does not execute downstream work.
- `ni` is not a task runner.
- `ni` is not an agent orchestration runtime.
- `ni run` compiles a bounded handoff prompt only.
- Model workspace packs are UX layers; the CLI remains the authority for
  readiness, lock creation, hash validation, and prompt compilation.

## Install Paths

Current public install paths are documented in
[Install ni](22_INSTALL.md) and [Curl Installer](install-curl.md).

### Release Binary

Status: Available for the verified v0.3.0 GitHub Release archives and
checksums.

Download the matching archive and `ni_0.3.0_checksums.txt` from:

<https://github.com/Nam-Cheol/ni/releases/tag/v0.3.0>

Then verify the checksum, unpack the binary, and run:

```bash
ni --help
ni version
```

### Curl Installer

Status: Available for the verified v0.3.0 release assets.

Inspect the installer before running it:

```bash
VERSION="0.3.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --version "$VERSION"
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

### Source

Status: Available.

Use source mode when developing `ni` or trying the CLI from a checkout:

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
go run ./cmd/ni status --dir .
```

Source and local build paths require Go 1.22 or newer.

### Model Workspace Packs

Status: Available for source packs, manual copy from verified source paths,
Codex zip packaging, Claude zip packaging, and Claude target-directory dry-run
install. Experimental as a broad product path because global host discovery and
global install locations are not verified.

See [Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.md).
The packs do not install `ni`, call model APIs, invoke downstream execution, or
replace CLI proof.

## Announcement Drafts

### GitHub Release Note Summary

`ni` v0.3.0 introduces the Project Intent Compiler for AI Agents: a small CLI
for turning planning conversations into checked, locked project contracts before
agents or teams begin work.

Tagline: Don't run the agent yet. Compile the intent first.

This release includes `ni init`, `ni status`, `ni end`, and `ni run`, plus
lockfile hash validation, status proof output, install documentation, verified
release binary and curl installer paths, and experimental model workspace
packs.

Boundary: `ni` does not execute downstream work. It is not a task runner or
agent orchestration runtime. `ni run` compiles a bounded handoff prompt only.

### X/Twitter Short Post

Don't run the agent yet. Compile the intent first.

`ni` turns planning conversations into locked project contracts before AI agents
or teams start work.

v0.3.0 includes the core CLI, verified release binary and curl install paths,
and experimental model workspace packs.

Not a task runner. Not an orchestration runtime.

### LinkedIn Post

I am sharing `ni`, a Project Intent Compiler for AI Agents.

The idea is simple: before a coding agent, workflow, or team starts work, turn
the planning conversation into a checked and locked project contract.

`ni` gives that workflow a deterministic pre-runtime layer: initialize a
planning workspace, review readiness, lock accepted intent, validate hashes, and
compile a bounded handoff prompt.

One-liner: Don't run the agent yet. Compile the intent first.

`ni` v0.3.0 has verified release binary and curl installer paths, source usage,
and experimental model workspace packs. It does not execute downstream work and
it is not a task runner or agent orchestration runtime.

### Hacker News / Reddit Style Post

I built `ni`, a small CLI that treats planning as something to compile before
execution starts.

It turns a planning conversation into a project contract, checks whether the
contract is ready, locks accepted intent with hashes, and compiles a short
handoff prompt for downstream agents or teams.

The product boundary is intentionally narrow: `ni` does not run the downstream
work. It is not a task runner and not an orchestration runtime. The first
release focuses on the kernel: `init`, `status`, `end`, and `run`.

One-liner: Don't run the agent yet. Compile the intent first.

### Korean Community Post

`ni`를 공개합니다.

`ni`는 AI agent나 팀이 작업을 시작하기 전에 planning conversation을 locked project
contract로 바꾸는 Project Intent Compiler입니다.

핵심 문장은 이겁니다:

> Don't run the agent yet. Compile the intent first.

`ni` v0.3.0은 `init`, `status`, `end`, `run` 중심의 작은 CLI입니다. Readiness를
확인하고, accepted intent를 lock하고, hash를 검증하고, downstream agent나 팀에
넘길 짧은 handoff prompt를 compile합니다.

주의할 점도 명확합니다. `ni`는 downstream work를 실행하지 않습니다. Task runner가
아니고, agent orchestration runtime도 아닙니다.

## FAQ

### Is ni a replacement for coding agents?

No. `ni` sits before coding agents or human teams. It creates, validates, locks,
and compiles planning contracts so downstream actors know what intent they may
trust.

### Can I use it without Go?

Yes, if you use the verified v0.3.0 release binary or the verified curl
installer path. Source usage and local builds require Go 1.22 or newer.

### Can I use it without terminal?

Partially. Model workspace packs can help draft planning docs in a model
workspace, but deterministic readiness, lock creation, hash validation, and
prompt compilation still require CLI proof today.

### What does ni lock?

`ni end` locks the accepted planning state into `.ni/plan.lock.json`. The lock
records hashes for the contract, planning docs, and accepted plan material so
downstream actors can detect intent changes.

### What happens after ni-run?

`ni run` compiles a 4000-character-or-less handoff prompt from the locked plan.
After that, a downstream agent, team, or workflow may choose to use the prompt,
but `ni` does not execute that work.
