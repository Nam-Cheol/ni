# Model Pack Install Verification

이 문서는 NI Codex와 Claude model workspace pack의 검증된 설치 및 packaging
경로를 기록한다.

Model workspace pack은 UX layer다. Model과 user가 planning docs를 유지하고, CLI
proof를 review하고, CLI로 lock하고, CLI로 bounded handoff prompt를 compile하도록
돕는다. Pack 자체는 authority가 아니며 runtime execution을 추가하지 않는다.

## Current Status

| Pack | Repo-local usage | Manual copy usage | Zip package usage | Dry-run install | Global install claim |
| --- | --- | --- | --- | --- | --- |
| Codex skills | Available | Available | Available | Planned | Unverified |
| Claude skills | Repository files로 Available | Available | Available | User-provided target에 대해 Available | Unverified |

전체 model workspace pack status는 product path로는 **Experimental**이다. Global
host discovery와 global install location이 verify되지 않았기 때문이다. 위에 적은
source pack과 zip packaging path는 **Available**이다.

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

Manual copy는 두 pack 모두에서 Available이다. Matching package root의 skill
directories를 user가 해당 host에 대해 verify한 model-workspace skill folder로
copy한다.

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

## Verification Command

Run:

```bash
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

- **Available:** source packs, verified source paths에서 manual copy, Codex zip
  packaging, Claude zip packaging, Claude target-directory dry-run install.
- **Experimental:** broad product path로서의 model workspace packs. Global host
  discovery와 no-terminal operation이 external host behavior에 의존하기 때문이다.
- **Planned:** Codex dry-run installer와 future package-manager-like model pack
  installer.
