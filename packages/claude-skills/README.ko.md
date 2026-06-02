# NI Claude Skill Pack

이 package는 Claude Code 또는 Claude Skill-compatible workspace에서 사용할 수
있는 NI workflow instruction을 담고 있다.

이 skill들은 authoring, review, lock, prompt compilation UX를 돕는다. Authority는
계속 `ni` CLI에 있다. Readiness, lock 생성, lock hash verification, prompt
compilation은 CLI result를 기준으로 판단한다.

## Status

Status: Experimental.

Verified: source files, package zip, metadata checks, user-provided target에
대한 guarded dry-run copy.

Not verified: global host install, provider behavior, cross-machine install,
global Claude skill discovery.

Boundary: Skills are UX; CLI is authority.

## Skills

| Skill | Purpose |
| --- | --- |
| `ni-start` | Conversation-driven planning을 계속하고 `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`을 함께 유지한다. |
| `ni-grill` | Lock 전에 accepted 또는 nearly accepted planning content를 challenge한다; downstream work를 execute하거나 model judgment로 readiness를 approve하지 않는다. |
| `ni-status-review` | `ni status --proof` output을 설명하고 다음 planning question을 찾되 second readiness engine이 되지 않는다. |
| `ni-end` | CLI readiness를 review하고 explicit confirmation 뒤에만 `ni end`로 lock한다. |
| `ni-run` | Valid lock에서 bounded handoff prompt를 compile하고 downstream work는 실행하지 않는다. |

## Authority Rules

- Skills are UX; CLI is authority.
- Readiness claim 전에는 `ni status`를 run 또는 request한다.
- `ni-start`는 grouped `ni status --proof --next-questions` output이 있으면
  primary planning interview로 사용해야 한다.
- `ni-start`는 의미 있는 authoring update 뒤 changed files, affected IDs,
  before/after CLI status, remaining blockers, next question group을 담은 concise
  planning proof block을 보여야 한다.
- Skills may help draft or explain proof-related planning text.
- Skills do not determine readiness.
- Skills do not lock plans.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- `ni-grill` challenges planning quality before lock. It does not execute work.
- `ni status`가 `BLOCKED`이면 `ni-grill`은 새로운 critique를 만들기 전에
  deterministic blockers를 먼저 사용해야 한다.
- `ni-grill`은 model judgment로 lock을 approve하지 않는다.
- Lock claim 전에는 `ni end`를 run 또는 request한다.
- Handoff prompt claim 전에는 `ni run`을 run 또는 request한다.
- `.ni/plan.lock.json`을 manually edit하지 않는다.
- Stale lock 또는 hash mismatch가 있으면 stop하고 `BLOCKED`를 report한다.
- Claude APIs를 call하지 않는다.
- Downstream work를 execute하지 않는다.
- Shell/Codex adapters, evidence runners, queues, model orchestration, PR
  automation, release automation, TUI, web UI behavior를 추가하지 않는다.

## Packaging

Repository root에서 실행한다:

```bash
bash scripts/package-claude-skills.sh
```

Archive output:

```text
dist/ni-claude-skills.zip
```

## Copy This Folder

Claude-compatible environment마다 skill folder location이 다를 수 있다. 이
repository는 global path를 assume하지 않는다.

사용 중인 environment가 document하고 사용자가 verify한 target directory를
지정한다. 그 target으로 skill folder만 copy한다:

```bash
TARGET=/path/to/verified/claude-skills
mkdir -p "$TARGET"
cp -R packages/claude-skills/ni-start "$TARGET/"
cp -R packages/claude-skills/ni-grill "$TARGET/"
cp -R packages/claude-skills/ni-status-review "$TARGET/"
cp -R packages/claude-skills/ni-end "$TARGET/"
cp -R packages/claude-skills/ni-run "$TARGET/"
```

Zip archive에서 사용할 때는 먼저 unpack한 뒤 같은 skill folder를 copy한다:

```bash
unzip -q dist/ni-claude-skills.zip -d /tmp/ni-claude-skills-unpacked
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-start "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-grill "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-status-review "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-end "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-run "$TARGET/"
```

이 절차는 file-copy workflow일 뿐이다. Specific host path와 loading behavior가
verify되지 않았다면 target을 global Claude install path로 설명하지 않는다.

## Guarded Install Script

Claude pack에는 guarded copy script도 있다. 사용 중인 environment가 document하고
사용자가 verify한 target directory를 지정한다:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

Installer는 target이 명시된 경우에만 skill directories를 copy한다. Existing skill
directory는 `--force` 없이는 preserve한다.

Manual copy는 이 source tree 또는 unpacked zip archive에서 Available이다. 현재
Claude-compatible host에 대해 user가 verify한 target folder로 skill directories만
copy한다. 그 target을 global Claude install path로 설명하지 않는다.

## Verify The Pack

Skills 목록을 확인한다:

```bash
find packages/claude-skills -mindepth 1 -maxdepth 1 -type d -name 'ni-*' -print | sort
```

`SKILL.md` files를 확인한다:

```bash
find packages/claude-skills -path '*/SKILL.md' -print | sort
bash scripts/check-skill-packs.sh
```

Zip package를 만든다:

```bash
bash scripts/package-claude-skills.sh
```

Archive contents를 inspect한다:

```bash
unzip -l dist/ni-claude-skills.zip
```

Full installation과 verification status는
`docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md`를 참고한다.

Broad product path의 Experimental status와 not_verified host/provider boundaries는
`docs/99_MODEL_WORKSPACE_STATUS.ko.md`를 참고한다.

## What This Does Not Do

- Claude APIs를 run하지 않는다.
- Implementation 또는 downstream work를 execute하지 않는다.
- Readiness, locking, hash checks, prompt compilation에 대한 `ni` CLI validation을
  replace하지 않는다.
