# NI Claude Skill Pack

이 package는 Claude Code 또는 Claude Skill-compatible workspace에서 사용할 수
있는 NI workflow instruction을 담고 있다.

이 skill들은 authoring, review, lock, prompt compilation UX를 돕는다. Authority는
계속 `ni` CLI에 있다. Readiness, lock 생성, lock hash verification, prompt
compilation은 CLI result를 기준으로 판단한다.

## Skills

| Skill | Purpose |
| --- | --- |
| `ni-start` | Conversation-driven planning을 계속하고 `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`을 함께 유지한다. |
| `ni-status-review` | `ni status --proof` output을 설명하고 다음 planning question을 찾되 second readiness engine이 되지 않는다. |
| `ni-end` | CLI readiness를 review하고 explicit confirmation 뒤에만 `ni end`로 lock한다. |
| `ni-run` | Valid lock에서 bounded handoff prompt를 compile하고 downstream work는 실행하지 않는다. |

## Authority Rules

- Skills are UX; CLI is authority.
- Readiness claim 전에는 `ni status`를 run 또는 request한다.
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

## Installation

Claude-compatible environment마다 skill folder location이 다를 수 있다. 이
repository는 global path를 assume하지 않는다.

사용 중인 environment가 document하고 사용자가 verify한 target directory를
지정한다:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

Installer는 target이 명시된 경우에만 skill directories를 copy한다. Existing skill
directory는 `--force` 없이는 preserve한다.

