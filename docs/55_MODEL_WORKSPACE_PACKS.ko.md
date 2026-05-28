# Model Workspace Packs

이 문서는 사용자가 `ni`를 terminal-only Go CLI로만 보지 않고 Codex, Claude,
generic model workspace 안에서 사용할 수 있게 하는 pack strategy를 정의한다.

이 strategy는 보수적이다. Model pack은 UX distribution이다. Model과 user가
planning state를 author, review, lock, compile하도록 돕지만 authority가 되지는
않는다. Readiness, lock 생성, lock hash verification, prompt compilation의
deterministic authority는 계속 CLI에 있다.

```text
skills and instructions -> user/model planning UX
ni CLI -> readiness gate, lockfile, hash check, prompt compiler
```

어떤 pack도 model API calls, execution runtime behavior, downstream adapters,
hidden implementation steps를 추가하면 안 된다. Valid CLI result를 얻지 못하면
pack은 model judgment로 대체하지 말고 `BLOCKED`를 report해야 한다.

## Status Legend

| Status | Meaning |
| --- | --- |
| Available | 이 repository에 matching files가 있고 해당 environment에서 오늘 사용할 수 있는 workflow다. |
| Experimental | Existing docs나 copied instructions로 시도할 수 있지만 packaged 또는 fully verified distribution path는 아니다. |
| Planned | Strategy는 accepted되었지만 matching files, packaging, installer support가 아직 없다. |
| Unverified | Environment가 비슷한 shape을 지원할 수 있지만 이 repository가 required host-specific file structure 또는 install path를 verify하지 않았다. |

## Pack Types

| Pack type | Purpose | Current status | Boundary |
| --- | --- | --- | --- |
| Repo-local `ni` skills | Current project workspace에 checked-in된 skill files를 사용한다. | Codex-style repo-local skills에 한해 Available. | Model은 planning docs와 `.ni/contract.json` edits를 도울 수 있지만 `ni status`, `ni end`, `ni run`이 authority다. |
| User-global skill pack | 같은 `ni` UX를 user의 global model workspace skill folder에 설치한다. | Claude pack은 user-provided, verified target directory에 한해 Available; global Codex installation은 여기서 verify되지 않았으므로 claimed하지 않는다. | CLI가 bundled, installed, bypassed된다고 imply하면 안 된다. |
| Downloadable zip skill pack | Skills, prompts, README instructions를 portable archive로 제공한다. | `scripts/check-skill-packs.sh`와 package scripts가 통과한 뒤 Codex 및 Claude skill packs에서 Available. | Archive는 UX만 package한다; CLI gates를 호출할 trusted way가 여전히 필요하다. |
| Manual copy-paste workflow | No-terminal users가 planning text를 model workspace에 붙여넣고 CLI-produced proofs를 다시 conversation에 붙여넣게 한다. | Workflow pattern으로는 Experimental; supported no-terminal product path는 아니다. | Copied result가 CLI에서 나온 것이 아니면 model은 readiness, lock, handoff를 선언할 수 없다. |
| Future package installer | Future package mechanism으로 model workspace packs를 install 또는 update한다. | Planned. | Installer work는 distribution infrastructure이며 `ni-kernel` execution behavior가 아니다. |

## Environment Matrix

| Environment | Repo-local skills | User-global pack | Downloadable zip pack | Manual copy-paste | Future package installer |
| --- | --- | --- | --- | --- | --- |
| Codex-style skill folders | Available: `.agents/skills/ni-start`, `.agents/skills/ni-end`, `.agents/skills/ni-run`이 이 repo에 있고 `packages/codex-skills`에는 `ni-status-review`도 있다. | Unverified: global Codex install path는 verified 또는 documented available 상태가 아니다. | Available: `scripts/package-codex-skills.sh`가 `dist/ni-codex-skills.zip`을 만든다. | Experimental: users가 docs와 CLI outputs를 paste할 수 있지만 현재 preferred path는 repo-local skills다. | Planned. |
| Claude Skills / slash commands | `packages/claude-skills` 아래 repository files로 Available; slash-command behavior는 claim하지 않는다. | `scripts/install-claude-skills.sh --target <verified-dir>`에 한해 Available; global Claude path는 assume하지 않는다. | Available: `scripts/package-claude-skills.sh`가 `dist/ni-claude-skills.zip`을 만든다. | Host가 skills를 load하지 않는 경우 generic instructions copy는 여전히 Experimental이다. | Planned. |
| Generic model instruction packs | Experimental: model이 repository docs를 읽을 수 있지만 `.agents/skills`는 Codex-style convention이다. | Planned. | Planned. | Experimental: visible instructions와 pasted CLI proofs를 사용한다. | Planned. |

이 matrix는 global Codex installation이나 Claude Skills distribution을
available로 표시하지 않는다. 그런 claim에는 이 repository에서 아직 수행하지 않은
host-specific verification이 필요하다.

## Skill Surface

Model workspace pack은 다음 UX actions를 cover해야 한다.

| UX action | Role | Current repository state | Status |
| --- | --- | --- | --- |
| `ni-start` | Planning conversation을 계속하고 user intent에서 `docs/plan/**`, `.ni/contract.json`, session continuity를 update한다. | `.agents/skills/ni-start/SKILL.md`가 있다. | Repo-local Codex-style skills에서 Available. |
| `ni-end` | CLI readiness를 review하고 explicit user confirmation을 받은 뒤 CLI가 `.ni/plan.lock.json`을 쓰게 한다. | `.agents/skills/ni-end/SKILL.md`가 있다. | Repo-local Codex-style skills에서 Available. |
| `ni-run` | Valid lock에서 4000 characters 이하의 handoff prompt를 compile한다. | `.agents/skills/ni-run/SKILL.md`가 있다. | Repo-local Codex-style skills에서 Available. |
| `ni-status-review` | `ni status` 또는 `ni status --proof` output을 설명하고 blockers와 next planning question을 제안한다. | `packages/claude-skills/ni-status-review/SKILL.md`와 `packages/codex-skills/ni-status-review/SKILL.md`가 있다. | Claude 및 Codex skill packs에서 Available. |
| `ni-readme-pamphlet-review` | `docs/52_README_PAMPHLET_STRATEGY.md` 기준으로 README changes를 review한다. | Standalone skill file은 아직 없다. | Planned and optional. |

`ni-status-review`는 model이 status output을 reinterpret하지 않고 preserve해야 하는
main proof를 다루므로 유용하다. 그러나 second readiness engine이 아니라 review
skill이어야 한다.

`ni-readme-pamphlet-review`는 release와 documentation polish에 유용할 수 있지만
Intent Lock Protocol의 일부가 아니다. Optional로 유지해야 하며 locking이나 prompt
compilation에 required되어서는 안 된다.

## Required Pack Contents

### Repo-local Codex-style pack

Repository-local pack이 현재 available shape이다:

```text
.agents/skills/ni-start/SKILL.md
.agents/skills/ni-end/SKILL.md
.agents/skills/ni-run/SKILL.md
docs/06_CODEX_SKILLS.md
docs/31_NI_START_BEHAVIOR.md
docs/35_NI_END_CONFIRMATION.md
docs/36_NI_RUN_HANDOFF.md
```

이 pack은 model workspace가 repository files를 읽고, CLI를 직접 실행하거나 user에게
exact CLI output을 요청할 수 있다고 가정한다.

Packaged Codex skill source는 다음과 같다:

```text
packages/codex-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

### User-global skill pack

User-global pack은 repo-local behavior를 mirror하지만 특정 project checkout 밖에
있다. Claude-compatible environment의 경우 이 repository는 user-provided,
verified target directory를 요구하는 safe copy script를 제공한다. Global Claude
path를 claim하거나 assume하지 않는다.

포함해야 할 내용:

- repo-local skills와 같은 behavioral instructions;
- project root를 locate한 뒤 edit해야 한다는 visible requirement;
- readiness claim 전 `ni status`를 run 또는 request해야 한다는 visible requirement;
- lock hash mismatch에서는 stop해야 한다는 visible requirement;
- update와 uninstall story.

### Downloadable zip skill pack

Zip pack은 portable archive여야 하며 kernel behavior를 바꾸는 installer가 아니다.
Codex archive shape:

```text
ni-codex-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

다음 command로 생성한다:

```bash
bash scripts/package-codex-skills.sh
```

Output path:

```text
dist/ni-codex-skills.zip
```

Claude archive shape은 다음과 같다:

```text
ni-claude-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

다음 command로 생성한다:

```bash
bash scripts/package-claude-skills.sh
```

Output path:

```text
dist/ni-claude-skills.zip
```

### Manual copy-paste workflow

Manual copy-paste는 model-guided planning에는 참여할 수 있지만 terminal commands를
직접 실행할 수 없는 users를 위한 것이다.

Workflow는 다음과 같다:

1. User가 project goal과 current planning docs를 model workspace에 copy한다.
2. Model은 `docs/plan/**`과 `.ni/contract.json` edits를 visible patches 또는
   replacement snippets로 propose한다.
3. Trusted local actor가 edits를 적용하고 `ni status`를 실행한다.
4. User가 exact status proof를 model workspace에 paste한다.
5. Status가 `BLOCKED`이면 model은 next planning question을 묻는다.
6. Status가 `READY` 또는 `READY_WITH_DEFERRALS`이면 model은 `ni-end` flow를
   guide할 수 있지만 lock은 CLI만 만들 수 있다.
7. Lock creation 이후 user가 lock result 또는 `ni run` output을 workspace에
   paste한다.

Manual workflow는 `ni`가 CLI 없이 작동한다는 claim이 아니다. No-terminal users를
loop 안에 두면서 trusted runner가 authoritative proof를 만들게 하는 방법이다.

### Future package installer

Future package installer는 model workspace packs를 install 또는 update할 수 있다.
Kernel functioning에 required되지 않는 distribution infrastructure로 다뤄야 한다.

Installer가 할 수 있는 일:

- verified host locations로 skill 또는 instruction files를 copy한다;
- 무엇을 어디에 install했는지 report한다;
- existing user edits를 default로 preserve한다;
- pack version metadata를 verify한다.

Installer가 하면 안 되는 일:

- downstream implementation 실행;
- model APIs 호출;
- readiness checks 약화;
- model judgment로 lockfiles 생성 또는 repair;
- CLI failures를 friendly model summaries 뒤에 숨기기.

## Authority Rules For Every Pack

- Skills are UX; the CLI is authority.
- Model은 docs를 draft하고 gaps를 detect하고 edits를 propose할 수 있다.
- Model은 `ni status` 없이 readiness를 declare할 수 없다.
- Model은 `ni end` 없이 lock할 수 없다.
- `ni run`이 missing 또는 stale lock을 report하면 model은 replacement prompt를
  compile할 수 없다.
- Open blocker questions는 locking을 막는다.
- High-severity risks에는 mitigation이 필요하다.
- `ni run` prompt output은 4000 characters 이하이어야 한다.
- `.ni/plan.lock.json`이 존재한 뒤 source-of-truth precedence는 다음과 같다:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

Lock hash mismatch가 있으면 모든 pack은 stop하고 `BLOCKED`를 report해야 한다.

## Availability Rules

- Repo-local Codex-style skills는 이 repository에서 available하다고 설명할 수 있다.
- Global Codex installation은 real install location과 loading behavior가 verify될
  때까지 available하다고 설명하면 안 된다.
- Claude skill installation은 user-provided, verified target directory를 요구해야
  하며 repository가 global Claude path를 assume하면 안 된다.
- Claude slash-command behavior는 available하다고 설명하면 안 된다. 이 pack은
  slash-command integration이 아니라 skills를 제공한다.
- Downloadable zip packs는 archive가 produced, inspected, documented된 경우에만
  available하다고 설명할 수 있다.
- Manual copy-paste는 experimental workflow로 설명할 수 있지만 complete
  no-terminal product로 설명하면 안 된다.
- Future package installer work는 real implementation과 validation path가 있기
  전까지 planned로 남아야 한다.

Current verification command, install paths, status language는
[Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.ko.md)를
참고한다.
