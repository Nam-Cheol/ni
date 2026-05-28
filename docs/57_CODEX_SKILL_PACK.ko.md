# Codex Skill Pack

Codex skill pack은 repo-local Codex workspace에서 `ni`를 model workspace
workflow로 사용할 수 있게 한다.

이 pack은 UX distribution이다. `ni-kernel`을 바꾸지 않고, `codex exec`를 call하지
않으며, downstream work를 execute하지 않는다. Readiness, lock 생성, lock hash
verification, prompt compilation의 authority는 계속 `ni` CLI다.

```text
Codex skills -> conversation workflow
ni CLI -> readiness gate, lockfile, hash check, prompt compiler
```

## Pack Contents

```text
packages/codex-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

Downloadable archive는 다음 command로 만든다:

```bash
bash scripts/package-codex-skills.sh
```

Output path:

```text
dist/ni-codex-skills.zip
```

## Skill Behavior

| Skill | Role | Authority boundary |
| --- | --- | --- |
| `ni-start` | `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`을 conversation-driven authoring으로 유지한다. | `ni status --dir . --next-questions`를 run 또는 request해야 하며 lock하지 않는다. |
| `ni-status-review` | `ni status --proof` output과 next questions를 설명한다. | CLI status를 그대로 preserve하며 second readiness engine이 되면 안 된다. |
| `ni-end` | Readiness를 review하고 explicit lock confirmation을 받는다. | `ni end --dir .` 전 `ni status --dir .`를 run 또는 request해야 하며 `.ni/plan.lock.json`을 manually write하지 않는다. |
| `ni-run` | Locked plan에서 bounded handoff prompt를 compile한다. | `ni run --dir . --target <target> --max-chars 4000`을 run 또는 request해야 하며 downstream work를 execute하지 않는다. |

## Repo-Local Installation

Repo-local usage가 verified path다. Skill directories를 workspace-local
`.agents/skills/` directory로 copy한다:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

이 package는 `ni` binary를 install하거나 invoke하지 않는다. Skill instruction만
제공한다. Authority가 필요한 경우 skill은 `ni` CLI command를 run 또는 request해야
한다.

Global Codex skill installation과 discovery path는 여기서 claim하지 않는다. 특정
Codex environment에서 verify되기 전까지 global usage는 experimental 또는 planned로
취급한다.

## Boundaries

Codex pack must not do the following:

- implementation execute;
- `codex exec` call;
- shell 또는 Codex adapters 추가;
- evidence runners, queues, PR automation, release automation, model
  orchestration 추가;
- `.ni/plan.lock.json` manually create, edit, repair;
- ready state를 만들기 위해 readiness blockers 약화;
- downstream-owned execution state 생성.

CLI가 stale lock 또는 hash mismatch를 report하면 모든 skill은 stop하고 `BLOCKED`를
report해야 한다.

## Validation

Dedicated skill-pack check는 source layout, metadata, authority boundaries,
package-script contents를 검증한다:

```bash
bash scripts/check-skill-packs.sh
```

Repository quality check도 skill-pack check를 실행한다:

```bash
bash scripts/quality.sh
```

Packaging check는 archive 생성 여부를 검증한다:

```bash
bash scripts/package-codex-skills.sh
```

Codex dry-run install support는 Planned다. Repo-local usage, manual copy, zip
packaging이 verified paths다. 자세한 내용은
[Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.ko.md)를
참고한다.
