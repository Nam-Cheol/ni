# NI Codex Skill Pack

이 package는 repo-local usage를 위한 Codex-style NI workflow skill을 담고 있다.

이 skill들은 Codex workspace에서 planning contract를 author, review, lock,
compile하는 UX를 돕는다. Authority는 계속 `ni` CLI에 있다. Readiness, lock
생성, lock hash verification, prompt compilation은 CLI result를 기준으로 판단한다.

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
- `codex exec`를 call하지 않는다.
- Downstream work를 execute하지 않는다.
- Shell/Codex adapters, evidence runners, queues, model orchestration, PR
  automation, release automation, TUI, web UI behavior를 추가하지 않는다.

## Packaging

Repository root에서 실행한다:

```bash
bash scripts/package-codex-skills.sh
```

Archive output:

```text
dist/ni-codex-skills.zip
```

## Repo-Local Usage

이 repository는 repo-local skill usage만 verify한다. Skill directories를
workspace-local `.agents/skills/` directory로 copy한다:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

Skill이 authority를 요구하면 project workspace에서 relevant `ni` CLI command를
실행한다.

Global Codex skill installation과 discovery path는 이 package에서 claim하지
않는다. 특정 Codex environment에서 verify되기 전까지 global usage는 experimental
또는 planned로 취급한다.
