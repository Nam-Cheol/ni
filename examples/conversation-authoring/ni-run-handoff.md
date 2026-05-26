# ni-run handoff transcript

This fixture is illustrative. It shows the handoff shape for `ni-run`; it is
not exact output from a live model.

## Inferred target

**User**

Invoke `ni-run` for the locked refund triage plan.

**Model using ni-run**

I will use the locked project context to infer the target, then let the CLI
verify the lock while compiling the prompt.

Current locked context names `human-team` as the downstream handoff target.

```bash
ni run --dir . --target human-team --max-chars 4000
```

CLI result:

```text
Target: human-team

You are executing a locked NI project plan.
Authoritative sources:
- .ni/plan.lock.json
- .ni/contract.json
- docs/plan/

Goal:
Use the locked plan to prepare the next human-team implementation handoff.
```

The prompt is 4000 characters or less and references the locked authoritative
files.

`ni` compiled this prompt only. It did not execute implementation, Codex, shell
commands, adapters, queues, PR automation, or downstream runtime work.

## Target question

**User**

Invoke `ni-run` for this locked plan.

**Model using ni-run**

`.ni/plan.lock.json` exists, but I cannot infer a single downstream target from
the request or locked project context. Which target should I compile?

Options: `generic`, `codex`, `human-team`, `hyper-run`, `namba-ai`,
`ouroboros`, `spec-kit`.

**User**

Use `codex`.

**Model using ni-run**

```bash
ni run --dir . --target codex --max-chars 4000
```

CLI result:

```text
Target: codex

You are executing a locked NI project plan.
Authoritative sources:
- .ni/plan.lock.json
- .ni/contract.json
- docs/plan/

Goal:
Build the smallest valid next product increment from the locked plan.
```

The prompt is 4000 characters or less and references the locked authoritative
files.

`ni` compiled this prompt only. It did not execute implementation, Codex, shell
commands, adapters, queues, PR automation, or downstream runtime work.

## Blocked stale lock

**User**

Invoke `ni-run --target human-team`.

**Model using ni-run**

```bash
ni run --dir . --target human-team --max-chars 4000
```

CLI result:

```text
BLOCKED
lock hash mismatch for docs/plan/02_capabilities.md
```

`BLOCKED`. I will not synthesize a replacement prompt, run implementation, or
repair the lockfile. Return to planning authoring, resolve the mismatch, and
lock the plan again before running `ni-run`.
