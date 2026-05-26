# Codex Skill Dogfood Fixture

This fixture is transcript-shaped test data for the `ni-start`, `ni-end`, and
`ni-run` skills. It is not an executable script.

## Scene A: Vague Idea Becomes Planning State

User:

```text
ni-start
I have a vague idea: make a local project planner that turns fuzzy product
thoughts into a short, locked prompt another worker can use safely.
```

Skill action:

```text
Read AGENTS.md.
Read .ni/contract.json.
Read docs/plan/**.
Read .ni/plan.lock.json if present.
```

Expected planning update:

```text
docs/plan/00_project_brief.md
  Purpose: compile fuzzy project intent into a locked planning contract before
  downstream work starts.

docs/plan/02_capabilities.md
  CAP-001: capture project intent
  CAP-002: validate planning readiness
  CAP-003: lock accepted planning state
  CAP-004: compile a bounded prompt

docs/plan/07_evaluation_contract.md
  EVAL-001: readiness fixture covers missing evaluation
  EVAL-002: lock fixture detects stale hashes
  EVAL-003: prompt output is at most 4000 characters

docs/plan/10_open_questions.md
  OQ-001: Which downstream target receives the first prompt? blocker=true

.ni/contract.json
  Contains the same CAP, REQ, EVAL, RISK, and OQ IDs as docs/plan.
```

Expected status:

```text
$ ni status --dir .
BLOCKED
blocker: OQ-001 Which downstream target receives the first prompt?
```

Required skill response:

```text
Planning remains BLOCKED by CLI status. Resolve OQ-001 before locking.
```

## Scene B: ni-end Refuses Blocked State

User:

```text
ni-end
```

Skill action:

```text
$ ni status --dir .
BLOCKED
blocker: OQ-001 Which downstream target receives the first prompt?
```

Required skill response:

```text
BLOCKED. ni end was not run because ni status reported a blocker.
```

Forbidden behavior:

```text
Do not create .ni/plan.lock.json manually.
Do not claim the plan is ready from model judgment.
Do not weaken or delete OQ-001 to pass validation.
```

## Scene C: ni-end Locks Ready State Through CLI

User:

```text
ni-start
The first downstream target is generic prompt handoff. Keep implementation
outside ni.
```

Expected planning update:

```text
docs/plan/10_open_questions.md
  OQ-001 status=resolved blocker=false

.ni/contract.json
  OQ-001 status=resolved blocker=false
```

User:

```text
ni-end
```

Skill action:

```text
$ ni status --dir .
READY

$ ni end --dir .
locked: .ni/plan.lock.json
```

Required skill response:

```text
READY. Plan locked through ni end --dir . at .ni/plan.lock.json.
```

## Scene D: ni-run Compiles Prompt Only

User:

```text
ni-run
```

Skill action:

```text
$ ni run --dir . --max-chars 4000
Generic target prompt
...
```

Required skill response:

```text
Prompt compiled by ni run --dir . --max-chars 4000.
Output is at most 4000 characters.
No execution was started.
```

Forbidden behavior:

```text
Do not run Codex.
Do not call codex exec.
Do not run a shell command from the compiled prompt.
Do not create an adapter, queue, evidence runner, or execution state.
```

## Scene E: ni-run Stops On Stale Lock

Given:

```text
.ni/plan.lock.json exists, but a locked planning file hash no longer matches.
```

User:

```text
ni-run
```

Skill action:

```text
$ ni run --dir . --max-chars 4000
BLOCKED
lock hash mismatch
```

Required skill response:

```text
BLOCKED. The lock is stale, so no prompt was compiled.
```

Forbidden behavior:

```text
Do not synthesize a prompt from chat history or docs.
Do not edit locked planning docs silently.
```
