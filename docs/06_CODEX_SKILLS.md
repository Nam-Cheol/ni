# Codex skills

Codex skills are the first UX layer for `ni`, but they are not the source of authority.

## Skills to provide

```text
.agents/skills/ni-start/SKILL.md
.agents/skills/ni-end/SKILL.md
.agents/skills/ni-run/SKILL.md
```

## Skill boundaries

```text
ni-start
  Update planning docs and .ni/contract.json.
  Show readiness gaps.
  Do not lock.

ni-end
  Run or request ni status.
  If blocked, explain blockers.
  If ready, summarize the contract and ask for explicit confirmation.
  After confirmation, run or request ni end.
  Do not personally declare completion.

ni-run
  Require an existing valid lockfile.
  Run or request ni run.
  Do not implement directly before a generated harness/work packet is proposed.
```

## Authority rule

Skills are interaction protocols. The CLI validates readiness, lock state, and prompt constraints.

See [ni-end confirmation](35_NI_END_CONFIRMATION.md) for the required
confirmation flow before lock creation.
