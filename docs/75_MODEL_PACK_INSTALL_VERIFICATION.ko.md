# Model Pack Install Verification

이 문서는 NI Codex와 Claude model workspace pack의 검증된 설치 및 packaging
경로를 기록한다.

Model workspace pack은 UX layer다. Model과 user가 planning docs를 유지하고, CLI
proof를 review하고, CLI로 lock하고, CLI로 bounded handoff prompt를 compile하도록
돕는다. Pack 자체는 authority가 아니며 실행 기능을 추가하지 않는다.

## Current Status

전체 model workspace pack status는 product path로는 **Experimental**이다.

아래의 더 구체적인 status evidence를 사용한다. Source files, package roots,
metadata checks, zip packaging은 verified되었다. Host-level/global install,
provider runtime behavior, cross-machine installation은 나중의 host-specific
verification document가 명시하지 않는 한 이 repository에서 verified된 것이 아니다.

| Pack | Repo-local source | Manual copy workflow | Zip package | Dry-run install | Global install claim |
| --- | --- | --- | --- | --- | --- |
| Codex skills | Verified | Host target이 verified되기 전까지 Experimental | Verified | Planned | not_verified |
| Claude skills | Verified | Host target이 verified되기 전까지 Experimental | Verified | User-provided target dry run에 대해서만 Verified | not_verified |

Model workspace packs를 broad product path로 **Available**이라고 설명하지 않는다.
**Available**은 특정 host path에서 install과 usage verification을 마친 경우에만 그
path에 대해 사용한다. Status vocabulary는
[Model Workspace Status](99_MODEL_WORKSPACE_STATUS.ko.md)를 참고한다.

## Verified Source Layout

Source pack roots:

```text
packages/codex-skills/
packages/claude-skills/
```

각 pack은 다음 files를 포함해야 한다:

```text
README.md
README.ko.md
ni-start/SKILL.md
ni-status-review/SKILL.md
ni-end/SKILL.md
ni-run/SKILL.md
```

각 `SKILL.md`에는 frontmatter-style `name:` 및 `description:` metadata, explicit
authority boundary, CLI가 readiness, lock, hash, prompt compilation authority라는
visible instruction이 있어야 한다.

## Installation Paths

### Copy This Folder Quick Guides

Codex repo-local usage는 Codex skill folder를 workspace의 `.agents/skills/`
directory로 copy한다:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

Claude-compatible host는 해당 host에서 document되고 verify된 target skill
directory를 고른 뒤 Claude skill folder를 copy한다:

```bash
TARGET=/path/to/verified/claude-skills
mkdir -p "$TARGET"
cp -R packages/claude-skills/ni-start "$TARGET/"
cp -R packages/claude-skills/ni-status-review "$TARGET/"
cp -R packages/claude-skills/ni-end "$TARGET/"
cp -R packages/claude-skills/ni-run "$TARGET/"
```

이 절차는 file-copy workflow일 뿐이다. Global host install 또는 global skill
discovery를 prove하지 않는다.

### Repo-local Codex usage

Repo-local Codex-style usage는 Codex skill directories를 workspace-local
`.agents/skills/` directory로 copy하는 방식으로 verify된다:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

이것은 file-copy workflow일 뿐이다. `ni`를 install하지 않고, `codex exec`를
call하지 않으며, downstream work를 execute하지 않는다. Skills는 authoritative
gate가 필요할 때 `ni status`, `ni end`, `ni run`을 run 또는 request해야 한다.

### Manual Copy Usage

Manual copy는 specific target folder와 host loading behavior를 user가 verify하기
전까지 experimental workflow다. Matching package root의 skill directories를
verified model-workspace skill folder로 copy한다.

이 repository는 global Codex 또는 Claude skill path를 claim하지 않는다. Host
environment가 copied files를 skill로 load하지 못해도 user는 instructions를 직접
읽고 exact CLI proof를 conversation에 paste할 수 있다. 그 manual proof workflow는
Experimental이며 CLI를 대체하지 않는다.

### Zip Package Usage

Codex archive 생성:

```bash
bash scripts/package-codex-skills.sh
```

Expected output:

```text
dist/ni-codex-skills.zip
```

Claude archive 생성:

```bash
bash scripts/package-claude-skills.sh
```

Expected output:

```text
dist/ni-claude-skills.zip
```

Zip archive는 portable skill bundle이다. `ni` binary를 install하지 않고, model
APIs를 run하지 않고, `codex exec`를 invoke하지 않으며, downstream work를 execute하지
않는다.

### Claude Dry-Run Install

Claude pack에는 guarded copy script가 있다. User가 이미 verify한 target directory가
필요하다:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

Dry run은 file 변경 없이 copy operations를 출력한다. Install script는 global Claude
skill path를 assume하지 않고, `--force`가 없으면 existing skill directories를
preserve한다.

Codex dry-run install support는 Planned다. Codex global installation은 이
repository에서 Unverified 상태다.

## Verify The Pack

Skills 목록을 확인한다:

```bash
find packages/codex-skills -mindepth 1 -maxdepth 1 -type d -name 'ni-*' -print | sort
find packages/claude-skills -mindepth 1 -maxdepth 1 -type d -name 'ni-*' -print | sort
```

`SKILL.md` files와 README boundary text를 확인한다:

```bash
find packages/codex-skills packages/claude-skills -path '*/SKILL.md' -print | sort
bash scripts/check-skill-packs.sh
```

Checker가 verify하는 항목:

- 두 package root가 존재한다;
- 모든 expected skill에 `SKILL.md`가 있다;
- 모든 skill에 `name:` 및 `description:` metadata가 있다;
- 각 pack README가 CLI authority boundary를 유지한다;
- package scripts가 required files를 포함하고 zip archive를 만든다;
- Claude installer가 `--dry-run`과 `--target`을 지원한다;
- Claude dry-run installer가 file 변경 없이 완료된다.

Checker는 Codex APIs, Claude APIs, downstream execution systems를 call하지 않는다.

Zip archives를 package한다:

```bash
bash scripts/package-codex-skills.sh
bash scripts/package-claude-skills.sh
```

Archives를 inspect한다:

```bash
unzip -l dist/ni-codex-skills.zip
unzip -l dist/ni-claude-skills.zip
```

Archive inspection은 pack root, pack README files, `SKILL.md` files가 있는 네 개의
expected skill folders만 보여야 한다.

## What This Does Not Do

- Codex APIs, Claude APIs, `codex exec`를 run하지 않는다.
- Implementation 또는 downstream work를 execute하지 않는다.
- Readiness, locking, hash checks, prompt compilation에 대한 `ni` CLI validation을
  replace하지 않는다.

## Boundary Checklist

모든 model workspace pack은 다음 rules를 visible하게 유지해야 한다:

- Skills are UX; the CLI is authority.
- Readiness claim에는 `ni status`가 필요하다.
- Lock creation에는 `ni end`가 필요하다.
- Prompt compilation에는 `ni run`이 필요하다.
- Lock hash mismatch에서는 workflow가 `BLOCKED`로 stop한다.
- Skills는 `.ni/plan.lock.json`을 manually edit하면 안 된다.
- Skills는 model APIs를 call하면 안 된다.
- Skills는 `codex exec`를 call하면 안 된다.
- Do not execute implementation, adapters, queues, PR automation, release
  automation, downstream runtime work.

## Release Status Language

Host-specific verification이 더 생기기 전까지 다음 status language를 사용한다:

- **Experimental:** broad product path로서의 model workspace packs. Host-level
  install, global discovery, no-terminal operation, provider behavior가 external
  host behavior에 의존하기 때문이다.
- **Verified:** 이 repository에서 통과하는 repo-local skill files, package source
  roots, metadata checks, zip package scripts.
- **Available:** 특정 host path에서 install과 usage verification을 마친 경우에만
  그 path에 대해 사용한다.
- **not_verified:** global Codex install, global Claude install, provider runtime
  behavior, cross-machine installation.
- **Planned:** Codex dry-run installer와 future package-manager-like model pack
  installer.
