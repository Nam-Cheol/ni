# Distribution Strategy

This document defines how users can adopt `ni` before they know Go, and later
before they use a terminal directly.

The strategy is directional. It must not be read as a claim that all paths are
available today. Current availability includes source, local builds, local
installs, verified v0.4.0 release archives, the verified v0.4.0 curl installer
path, and repo-local model workspace assistance. Package managers remain
planned until their external assets exist and are verified. Future
package-manager and installer work belongs to repository infrastructure,
packaging, and documentation. It is not `ni-kernel` runtime execution behavior.

`ni-kernel` remains the authority for:

- docs contract creation and validation;
- readiness through `ni status`;
- lock creation through `ni end`;
- prompt compilation through `ni run`.

Distribution automation may help users obtain or invoke the CLI. It must not
turn `ni` into a task runner, shell adapter, package release runtime, hosted
execution service, or multi-agent execution layer.

## Current Factual Status

| Path | Status | Notes |
| --- | --- | --- |
| Source / Go | Available | Developer path. |
| Local binary | Available | Built from this checkout; local install path. |
| Release binary | Available | Verified v0.4.0 assets. |
| Curl installer | Available | Verified against v0.4.0 release assets. |
| Model workspace packs | Experimental | UX layer; CLI remains authority; host-level/global install and provider behavior remain unverified unless documented. |
| No-terminal method | Experimental / assisted | Drafting only; deterministic validation requires CLI proof from a trusted runner. |
| Homebrew | Planned / v0.5 candidate | Deferred, not guaranteed; not Available until a tap/formula exists and `brew install`, `ni --help`, and `ni version` are tested. |
| Runtime execution, shell adapters, Codex exec, queues, PR automation | Not included | Not part of `ni-kernel`; future downstream integration must be separate packages, target exports, or seed formats. |

## Distribution Matrix

| Track | Status | User type | Required dependency | Trust model | Implementation work needed |
| --- | --- | --- | --- | --- | --- |
| Source mode | Available | Developers, early evaluators, contributors, Go-comfortable users | Go 1.22 or newer; Git optional for version metadata | Trust the checked-out source, local Go toolchain, repository tests, and quality checks | Keep `go run`, `make build`, `make test`, and `make quality` documented and working |
| Local binary mode | Available | Users who want `./bin/ni` or a local install from this checkout | Go 1.22 or newer; local shell | Trust the checked-out source, local build, and temporary install checks | Keep `make build`, `make install-local`, and `bash scripts/install-check.sh` working |
| Release binary mode | Available | Users comfortable with a terminal but not with Go | Terminal; OS-specific downloaded `ni` binary from the verified v0.4.0 release | Trust GitHub Release assets plus published checksums after verification | Keep manual release asset process, OS/arch builds, checksums, verification docs, and rollback notes current |
| Curl installer mode | Available | Terminal users who want one command and no Go setup | `curl` or equivalent downloader; POSIX shell on supported platforms | Trust the installer script after inspecting it; it verifies real release assets with published checksums when available | Keep `install.sh` small, auditable, covered by installer checks, and reverify it against real release assets before new availability claims |
| Package manager mode | Planned | Users who prefer platform package managers | Homebrew first; Scoop later if Windows demand appears | Trust package manager metadata, formulas/manifests, and checksums that point to official release assets | Create a Homebrew tap or formula after release binaries stabilize; consider Scoop later; keep publishing outside `ni-kernel` |
| Model workspace mode | Experimental | Codex/Claude users who author plans through a model workspace | A model workspace that can read the repository docs and invoke the `ni` CLI as authority | Trust the CLI gates, not the model; skills are UX over docs and `.ni/contract.json` | Package and document portable skill packs; keep skill behavior aligned with `ni status`, `ni end`, and `ni run` |
| No-terminal mode | Experimental | Non-technical users, product leads, researchers, and teams who want docs-first planning without direct terminal use | Assisted docs-first workflow; a trusted runner still invokes `ni` gates for deterministic validation | Trust visible drafts only as planning inputs; trust readiness, lock, hashes, and compiled prompts only after CLI-generated proof | Keep assisted no-terminal docs factual; do not claim deterministic validation without CLI proof |

## Track Details

### 1. Source mode

Source mode is available now and is the primary supported distribution path for
first-time local use.

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

### 2. Local binary mode

Local binary mode is available.

Users can build and install `ni` from the checked-out source:

```bash
make build
make install-local
```

This path does not depend on release assets.

### 3. Release binary mode

Release binary mode is available for the verified v0.4.0 GitHub Release assets.

Users can download `ni` from GitHub Releases without installing Go. The
documented trust path is: choose the OS/arch asset, download the matching
checksum file from the same release, verify the checksum, unpack the binary, and
run `ni --help` and `ni version`.

This is repository distribution infrastructure. It must not add release
automation to `ni-kernel`, and it must not change `ni run` into an executor.

### 4. Curl installer mode

Curl installer mode is available for the verified v0.4.0 release assets.

The installer is a small `install.sh` that downloads a verified release asset.
It does not build from source by default, execute downstream work, or hide the
trust boundary. The script explains what it downloads, where it installs `ni`,
and how it verifies the asset.

This track depends on release binary mode. The script is tested with local fake
release assets, including checksum mismatch behavior, and was verified against
the real v0.4.0 darwin/arm64 release archive and checksum file on 2026-05-29.

### 5. Package manager mode

Package manager mode is planned and deferred, not available yet.

Homebrew is the first likely package manager because it matches the initial
developer audience and local macOS usage. Scoop can be considered later if
Windows demand appears. Package definitions should point to official release
assets and checksums.

Package publishing is external repository infrastructure. It is not part of the
Intent Lock Protocol and must not become kernel-owned execution state.
Homebrew can become Available only after a tap/formula exists and `brew
install`, `ni --help`, and `ni version` have been tested against that package
path.

### 6. Model workspace mode

Model workspace mode is experimental today. The repository contains
model-facing skill material for planning, locking, and prompt compilation, but
the CLI remains the authority.

The intended user is someone working in Codex, Claude, or a similar model
workspace. The model helps author `docs/plan/**` and `.ni/contract.json`, then
uses:

```bash
ni status
ni end
ni run
```

or source equivalents as the deterministic gates.

Repo-local skill files, package source folders, zip packaging scripts, and
metadata checks are verified repository evidence. Host-level/global install,
provider runtime behavior, and cross-machine installation remain unverified
unless a host-specific verification document records otherwise. Portable skill
packs package UX and instructions; they must not bypass readiness, lock, or hash
verification. See [Model Workspace Status](99_MODEL_WORKSPACE_STATUS.md) for
the status vocabulary.

### 7. No-terminal mode

No-terminal mode is experimental as an assisted docs-first drafting method.

The intended shape is a downloadable model pack and docs-first workflow where a
non-terminal user can start a plan without typing commands directly. Today it
is only an assisted planning method: a trusted local or workspace runner must
invoke `ni` for deterministic validation, lock creation, hash verification, and
prompt compilation. The user-facing experience should remain explicit about
which outputs are drafts and which outputs came from the CLI.

This is not a terminal-less web service, hosted execution service, or hidden
agent runner. The core contract remains:

```text
planning conversation -> docs contract -> readiness gate -> lockfile -> prompt
```

## Availability Rules

- Source mode, local binary build, local install, and repo-local model
  workspace assistance may be described as available or experimental today.
- Release binaries may be described as available for verified GitHub Release
  assets and checksums, currently v0.4.0.
- Curl install may be described as available for real release assets and
  checksums that have been verified with `install.sh`, currently v0.4.0.
- Package manager install must not be described as available until packages or
  formulas are published.
- No-terminal mode may be described only as assisted, drafting, or
  experimental until a downloadable model pack and proof-oriented workflow
  exist; it must not be described as deterministic validation without CLI
  proof.
- Model workspace packs must remain Experimental as a broad product path until
  host-level install and usage verification exists for a specific host path.
  Do not claim global Codex or Claude install, provider behavior, or
  cross-machine compatibility without recorded evidence.

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
