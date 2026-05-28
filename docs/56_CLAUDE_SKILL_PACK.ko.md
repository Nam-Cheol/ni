# Claude Skill Pack

Claude skill pack은 Claude Code 또는 Claude Skill-compatible workspace에서 `ni`를
model workspace workflow로 사용할 수 있게 한다.

이 pack은 UX distribution이다. `ni-kernel`을 바꾸지 않고, Claude APIs를 call하지
않으며, downstream work를 execute하지 않는다. Readiness, lock 생성, lock hash
verification, prompt compilation의 authority는 계속 `ni` CLI다.

```text
Claude skills -> conversation workflow
ni CLI -> readiness gate, lockfile, hash check, prompt compiler
```

## Pack Contents

```text
packages/claude-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

Downloadable archive는 다음 command로 만든다:

```bash
bash scripts/package-claude-skills.sh
```

Output path:

```text
dist/ni-claude-skills.zip
```

## Skill Behavior

| Skill | Role | Authority boundary |
| --- | --- | --- |
| `ni-start` | `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`을 conversation-driven authoring으로 유지한다. | `ni status --dir . --next-questions`를 run 또는 request해야 하며 lock하지 않는다. |
| `ni-status-review` | `ni status --proof` output과 next questions를 설명한다. | CLI status를 그대로 preserve하며 second readiness engine이 되면 안 된다. |
| `ni-end` | Readiness를 review하고 explicit lock confirmation을 받는다. | `ni end --dir .` 전 `ni status --dir .`를 run 또는 request해야 하며 `.ni/plan.lock.json`을 manually write하지 않는다. |
| `ni-run` | Locked plan에서 bounded handoff prompt를 compile한다. | `ni run --dir . --target <target> --max-chars 4000`을 run 또는 request해야 하며 downstream work를 execute하지 않는다. |

## Installation

Claude-compatible environment마다 skill folder location이 다를 수 있다. 이
repository는 global default를 assume하거나 document하지 않는다.

사용자의 environment가 document하고 사용자가 verify한 directory를 지정한다:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

Installer safety properties:

- `--target`이 required다.
- `--dry-run`은 file 변경 없이 copy operations를 출력한다.
- Existing skill directories는 `--force` 없이는 preserve한다.
- Script는 skill files만 copy한다. `ni` binary를 install하거나 invoke하지 않는다.
- Script는 Claude APIs를 call하지 않는다.

## Boundaries

Claude pack must not do the following:

- implementation execute;
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
package-script contents, guarded dry-run installer를 검증한다:

```bash
bash scripts/check-skill-packs.sh
```

Repository quality check도 skill-pack check를 실행한다:

```bash
bash scripts/quality.sh
```

Packaging check는 archive 생성 여부를 검증한다:

```bash
bash scripts/package-claude-skills.sh
```

Manual copy, zip package, dry-run install status는
[Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.ko.md)를
참고한다.
