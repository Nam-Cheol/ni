# NI Codex Skill Pack

이 package는 repo-local usage를 위한 Codex-style NI workflow skill을 담고 있다.

이 skill들은 Codex workspace에서 planning contract를 author, review, lock,
compile하는 UX를 돕는다. Authority는 계속 `ni` CLI에 있다. Readiness, lock
생성, lock hash verification, prompt compilation은 CLI result를 기준으로 판단한다.

## Status

Status: Experimental.

Verified: source files, package zip, metadata checks.

Not verified: global host install, provider behavior, cross-machine install,
global Codex skill discovery.

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

## Copy This Folder

이 repository는 repo-local skill usage만 verify한다. Skill directories를
workspace-local `.agents/skills/` directory로 copy한다:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-grill .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

Zip archive에서 사용할 때는 먼저 unpack한 뒤 같은 skill folder를 copy한다:

```bash
unzip -q dist/ni-codex-skills.zip -d /tmp/ni-codex-skills-unpacked
mkdir -p .agents/skills
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-start .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-grill .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-status-review .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-end .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-run .agents/skills/
```

Skill이 authority를 요구하면 project workspace에서 relevant `ni` CLI command를
실행한다.

Global Codex skill installation과 discovery path는 이 package에서 claim하지
않는다. 특정 Codex environment에서 verify되기 전까지 global usage는 experimental
또는 planned로 취급한다.

## Manual Copy And Zip Usage

Manual copy는 이 source tree 또는 unpacked zip archive에서 Available이다. 현재
model host에 대해 user가 verify한 target folder로 skill directories만 copy한다.
Host-specific path와 loading behavior가 verify되지 않았다면 그 target을 global
Codex install path로 설명하지 않는다.

Codex dry-run install support는 Planned다.

## Verify The Pack

Skills 목록을 확인한다:

```bash
find packages/codex-skills -mindepth 1 -maxdepth 1 -type d -name 'ni-*' -print | sort
```

`SKILL.md` files를 확인한다:

```bash
find packages/codex-skills -path '*/SKILL.md' -print | sort
bash scripts/check-skill-packs.sh
```

Zip package를 만든다:

```bash
bash scripts/package-codex-skills.sh
```

Archive contents를 inspect한다:

```bash
unzip -l dist/ni-codex-skills.zip
```

Full installation과 verification status는
`docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md`를 참고한다.

Broad product path의 Experimental status와 not_verified host/provider boundaries는
`docs/99_MODEL_WORKSPACE_STATUS.ko.md`를 참고한다.

## What This Does Not Do

- Codex APIs 또는 `codex exec`를 run하지 않는다.
- Implementation 또는 downstream work를 execute하지 않는다.
- Readiness, locking, hash checks, prompt compilation에 대한 `ni` CLI validation을
  replace하지 않는다.
