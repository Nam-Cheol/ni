# Model Workspace Status

## Current status

Model workspace packs는 broad product path로 **Experimental**이다.

Repo-local skill files, package source folders, zip packaging scripts, metadata
checks는 이 repository에서 verified되었다. Host-level/global install, provider
runtime behavior, cross-machine installation은 나중의 host-specific verification
document가 명시하지 않는 한 이 repository에서 verified된 것이 아니다.

## What is verified

| Evidence | Status | Verification |
| --- | --- | --- |
| Repo-local skill files | Verified | `.agents/skills/**/SKILL.md` |
| Claude skill package source | Verified | `packages/claude-skills/**` |
| Codex skill package source | Verified | `packages/codex-skills/**` |
| Skill metadata | Verified | `bash scripts/check-skill-packs.sh` |
| Claude zip package | Verified | `bash scripts/package-claude-skills.sh` |
| Codex zip package | Verified | `bash scripts/package-codex-skills.sh` |
| CLI authority wording | Verified | skill docs and README files |

## What is not verified

| Claim | Status | Reason |
| --- | --- | --- |
| Global Claude install | not_verified | host-level install was not tested |
| Global Codex install | not_verified | host-level install was not tested |
| Provider runtime behavior | not_verified | no provider API or host behavior was tested |
| Cross-machine install | not_verified | no multi-machine install matrix |
| Skills replace CLI validation | false | CLI remains authority |

## Status vocabulary

- **Experimental:** pack source와 packaging은 available이지만 host-level install
  또는 provider behavior는 fully verified되지 않았다.
- **Available:** 특정 host path에서 install과 usage verification을 마친 경우에만
  그 path에 대해 사용할 수 있다.
- **not_verified:** claim하면 안 된다.
- **UX layer:** authoring을 돕는 model instructions이며 readiness를 결정하지
  않는다.
- **CLI authority:** `ni status`, `ni end`, `ni run`, lock hashes, prompt
  compiler가 계속 authoritative하다.

## Rules for README/docs

- Host-level install이 verified되지 않은 한 model workspace packs를 broad
  product path로 "Available"이라고 말하지 않는다.
- Tested되지 않은 global install이 작동한다고 말하지 않는다.
- Skills가 CLI 없이 lock, validate, compile할 수 있다고 imply하지 않는다.
- Provider behavior가 guaranteed라고 imply하지 않는다.
- No-terminal mode가 deterministic이라고 imply하지 않는다.
- Experimental status를 marketing copy 뒤에 숨기지 않는다.

## Rules for skills

각 NI skill은 다음을 말하거나 보존해야 한다:

- Skills are UX; CLI is authority.
- Downstream work를 execute하지 않는다.
- `.ni/plan.lock.json`을 manually modify하지 않는다.
- Model judgment로 readiness를 approve하지 않는다.
- 관련될 때 `ni status --proof --next-questions`를 run 또는 request한다.
- User-facing questions는 사용자의 현재 언어를 쓰되 IDs, commands, file paths,
  status constants는 보존한다.

## How status may become Available later

1. 하나의 host environment를 고른다.
2. Documented method로 skill pack을 install한다.
3. Host가 skills를 discover하는지 확인한다.
4. `ni-start`, `ni-grill`, `ni-end`, `ni-run`을 invoke한다.
5. Skills가 CLI authority를 보존하는지 확인한다.
6. Downstream execution이 없는지 확인한다.
7. Verification doc을 기록한다.
8. 해당 host path만 Available로 update한다.

## Boundary

Model workspace packs는 runtime execution, model APIs, adapters, downstream
agents, queues, task running을 추가하지 않는다.
