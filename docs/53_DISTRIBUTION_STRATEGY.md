# Distribution Strategy

This document defines how users can adopt `ni` before they know Go, and later
before they use a terminal directly.

The strategy is directional. It must not be read as a claim that all paths are
available today. Current availability is source-first. Future distribution work
belongs to repository infrastructure, packaging, and documentation. It is not
`ni-kernel` runtime execution behavior.

`ni-kernel` remains the authority for:

- docs contract creation and validation;
- readiness through `ni status`;
- lock creation through `ni end`;
- prompt compilation through `ni run`.

Distribution automation may help users obtain or invoke the CLI. It must not
turn `ni` into a task runner, shell adapter, package release runtime, hosted
execution service, or multi-agent execution layer.

## Distribution Matrix

| Track | Status | User type | Required dependency | Trust model | Implementation work needed |
| --- | --- | --- | --- | --- | --- |
| Source mode | Available | Developers, early evaluators, contributors, Go-comfortable users | Go 1.22 or newer; Git optional for version metadata | Trust the checked-out source, local Go toolchain, repository tests, and quality checks | Keep `go run`, `make build`, `make test`, and `make quality` documented and working |
| Release binary mode | Next | Users comfortable with a terminal but not with Go | Terminal; OS-specific downloaded `ni` binary | Trust GitHub Releases assets plus published checksums or signatures once implemented | Define manual release asset process, build per OS/arch, publish checksums, document verification and rollback |
| Curl installer mode | Script added, release-gated | Terminal users who want one command and no Go setup | `curl` or equivalent downloader; POSIX shell on supported platforms | Trust the installer script only after it verifies downloaded release assets against published checksums or signatures | Keep `install.sh` small, auditable, and covered by installer checks; do not describe it as usable until release assets exist |
| Package manager mode | Planned | Users who prefer platform package managers | Homebrew first; Scoop later if Windows demand appears | Trust package manager metadata, formulas/manifests, and checksums that point to official release assets | Create a Homebrew tap or formula after release binaries stabilize; consider Scoop later; keep publishing outside `ni-kernel` |
| Model workspace mode | Available in repo-local form; portable packs planned | Codex/Claude users who author plans through a model workspace | A model workspace that can read the repository docs and invoke the `ni` CLI as authority | Trust the CLI gates, not the model; skills are UX over docs and `.ni/contract.json` | Package and document portable skill packs; keep skill behavior aligned with `ni status`, `ni end`, and `ni run` |
| No-terminal mode | Planned | Non-technical users, product leads, researchers, and teams who want docs-first planning without direct terminal use | Downloadable model pack and docs-first workflow; some trusted runner still invokes `ni` gates behind the scenes | Trust visible docs, lockfile hashes, and CLI-generated status/lock/run outputs; the model may not declare readiness on its own | Design downloadable model pack, guided docs workflow, and proof display; do not add a hosted service or terminal-less web runtime in this task |

## Track Details

### 1. Source mode

Source mode is available now and is the only fully supported distribution path
for first-time local use.

Users clone the repository and run:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
```

They may also build a local binary:

```bash
make build
./bin/ni --help
```

This path is best for contributors, evaluators, and users who already have Go
installed. Its trust model is transparent source plus local validation.

### 2. Release binary mode

Release binary mode is next, not available yet.

The goal is to let users download `ni` from GitHub Releases without installing
Go. Release assets should be plain OS/architecture binaries or archives with
checksums. Verification should be documented before the path is presented as
supported.

This is repository distribution infrastructure. It must not add release
automation to `ni-kernel`, and it must not change `ni run` into an executor.

### 3. Curl installer mode

Curl installer mode is script-ready, but public use is still release-gated.

The installer is a small `install.sh` that downloads a verified release asset.
It does not build from source by default, execute downstream work, or hide the
trust boundary. The script explains what it downloads, where it installs `ni`,
and how it verifies the asset.

This track depends on release binary mode. The script is tested with local fake
release assets, but it must not be presented as a working public install path
until real release assets and verification metadata exist.

### 4. Package manager mode

Package manager mode is planned, not available yet.

Homebrew is the first likely package manager because it matches the initial
developer audience and local macOS usage. Scoop can be considered later if
Windows demand appears. Package definitions should point to official release
assets and checksums.

Package publishing is external repository infrastructure. It is not part of the
Intent Lock Protocol and must not become kernel-owned execution state.

### 5. Model workspace mode

Model workspace mode is available today only in repo-local form. The repository
contains model-facing skill material for planning, locking, and prompt
compilation, but the CLI remains the authority.

The intended user is someone working in Codex, Claude, or a similar model
workspace. The model helps author `docs/plan/**` and `.ni/contract.json`, then
uses:

```bash
ni status
ni end
ni run
```

or source equivalents as the deterministic gates.

Portable skill packs are planned distribution work. They should package UX and
instructions, not bypass readiness, lock, or hash verification.

### 6. No-terminal mode

No-terminal mode is planned, not available yet.

The intended shape is a downloadable model pack and docs-first workflow where a
non-terminal user can inspect the plan, status proof, lock proof, and compiled
handoff without typing commands directly. A trusted local or workspace runner
may invoke `ni`, but the user-facing experience should remain explicit about
which outputs came from the CLI.

This is not a terminal-less web service, hosted execution service, or hidden
agent runner. The core contract remains:

```text
planning conversation -> docs contract -> readiness gate -> lockfile -> prompt
```

## Availability Rules

- Only source mode, local binary build, local install, and repo-local model
  workspace usage may be described as available today.
- Release binaries must not be described as available until GitHub Releases
  assets exist for supported platforms.
- Curl install must not be described as available until real release assets
  exist and `install.sh` verifies them.
- Package manager install must not be described as available until packages or
  formulas are published.
- No-terminal mode must not be described as available until a downloadable model
  pack and proof-oriented workflow exist.

## Boundary Rules

Distribution work may add scripts, release checklists, package metadata,
checksums, installer tests, docs, and downloadable model packs.

Distribution work must not add:

- shell adapter behavior to `ni run`;
- Codex or Claude execution adapters inside `ni-kernel`;
- release automation as kernel behavior;
- queues, PR automation, or execution evidence loops;
- a terminal-less hosted service as part of this strategy.

The repository may later automate builds and publishing as infrastructure. That
automation must remain outside the runtime behavior of `ni`.
