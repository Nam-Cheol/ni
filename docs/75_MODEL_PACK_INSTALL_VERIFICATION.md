# Model Pack Install Verification

This document records the verified installation and packaging paths for the NI
Codex and Claude model workspace packs.

Model workspace packs are UX layers. They help a model and user maintain
planning docs, review CLI proof, lock through the CLI, and compile a bounded
handoff prompt through the CLI. They are not authoritative and they do not add
runtime execution.

## Current Status

| Pack | Repo-local usage | Manual copy usage | Zip package usage | Dry-run install | Global install claim |
| --- | --- | --- | --- | --- | --- |
| Codex skills | Available | Available | Available | Planned | Unverified |
| Claude skills | Available as repository files | Available | Available | Available with user-provided target | Unverified |

Overall model workspace pack status: **Experimental** as a product path,
because global host discovery and global install locations are not verified.
The source packs and zip packaging paths listed above are **Available**.

## Verified Source Layout

The source pack roots are:

```text
packages/codex-skills/
packages/claude-skills/
```

Each pack must contain:

```text
README.md
README.ko.md
ni-start/SKILL.md
ni-status-review/SKILL.md
ni-end/SKILL.md
ni-run/SKILL.md
```

Each `SKILL.md` must include frontmatter-style `name:` and `description:`
metadata, an explicit authority boundary, and visible instructions to preserve
the CLI as the readiness, lock, hash, and prompt compilation authority.

## Installation Paths

### Repo-local Codex usage

Repo-local Codex-style usage is verified by copying the Codex skill directories
into a workspace-local `.agents/skills/` directory:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

This is a file-copy workflow only. It does not install `ni`, call `codex exec`,
or execute downstream work. The skills must run or request `ni status`,
`ni end`, and `ni run` when those authoritative gates are needed.

### Manual Copy Usage

Manual copy is available for both packs. Copy the skill directories from the
matching package root into a model-workspace skill folder that the user has
verified for that host.

This repository does not claim a global Codex or Claude skill path. If the host
environment cannot load skills from the copied files, the user may still read
the instructions manually and paste exact CLI proof back into the conversation.
That manual proof workflow is experimental and does not replace the CLI.

### Zip Package Usage

Create the Codex archive:

```bash
bash scripts/package-codex-skills.sh
```

Expected output:

```text
dist/ni-codex-skills.zip
```

Create the Claude archive:

```bash
bash scripts/package-claude-skills.sh
```

Expected output:

```text
dist/ni-claude-skills.zip
```

The zip archives are portable skill bundles. They do not install the `ni`
binary, run model APIs, invoke `codex exec`, or execute downstream work.

### Claude Dry-Run Install

The Claude pack has a guarded copy script. It requires a target directory that
the user has already verified:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

The dry run prints copy operations without changing files. The install script
does not assume a global Claude skill path and preserves existing skill
directories unless `--force` is passed.

Codex dry-run install support is planned. Codex global installation remains
unverified in this repository.

## Verification Command

Run:

```bash
bash scripts/check-skill-packs.sh
```

The checker verifies:

- both package roots exist;
- every expected skill has `SKILL.md`;
- every skill has `name:` and `description:` metadata;
- each pack README preserves the CLI authority boundary;
- package scripts include the required files and create zip archives;
- the Claude installer supports `--dry-run` and `--target`;
- the Claude dry-run installer completes without changing files.

The checker does not call Codex APIs, Claude APIs, or downstream execution
systems.

## Boundary Checklist

Every model workspace pack must keep these rules visible:

- Skills are UX; the CLI is authority.
- Readiness claims require `ni status`.
- Lock creation requires `ni end`.
- Prompt compilation requires `ni run`.
- Lock hash mismatches stop the workflow with `BLOCKED`.
- Skills must not edit `.ni/plan.lock.json` manually.
- Skills must not call model APIs.
- Skills must not call `codex exec`.
- Skills must not execute implementation, adapters, queues, PR automation,
  release automation, or downstream runtime work.

## Release Status Language

Use this status language until more host-specific verification exists:

- **Available:** source packs, manual copy from verified source paths, Codex zip
  packaging, Claude zip packaging, and Claude target-directory dry-run install.
- **Experimental:** model workspace packs as a broad product path, because
  global host discovery and no-terminal operation still depend on external
  host behavior.
- **Planned:** Codex dry-run installer and any future package-manager-like
  model pack installer.
