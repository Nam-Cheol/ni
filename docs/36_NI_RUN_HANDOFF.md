# ni-run handoff

`ni-run` is the conversational handoff step after a plan is locked. It helps a
model compile the locked plan into a target prompt, but it does not execute the
prompt or start downstream implementation.

The CLI remains authoritative:

```text
.ni/plan.lock.json exists -> ni run verifies lock hashes -> prompt or path is shown
```

## Required Flow

1. Read `AGENTS.md` and confirm `.ni/plan.lock.json` exists.
2. Infer the handoff target from the user request or locked project context.
3. Ask for the target only when it is not inferable.
4. Run or request `ni run --dir . --target <target> --max-chars 4000`.
5. Use `--out <path>` when the user asks for a file or the prompt is too long
   to display comfortably.
6. Treat the `ni run` result as the lock-hash verification. If it reports a
   missing or stale lock, report `BLOCKED` and stop.
7. Show the generated prompt or the output path.
8. State that `ni` compiled a prompt only and did not execute implementation.

The model must not synthesize a prompt from memory when the CLI refuses to
compile one.

## Target Inference

Use the target named by the user when one is explicit.

If no target is named, infer it from locked project context, especially:

- `docs/plan/09_execution_strategy.md`,
- `docs/plan/11_decision_log.md`,
- `.ni/contract.json`,
- recent conversation attached to the current locked plan.

Default to `codex` only when the project context says Codex is the current
experiment or selected downstream target. Do not select `codex` simply because
the assistant is Codex.

If the context points to multiple plausible targets, ask one short target
question before running `ni run`.

## Handoff Output

For stdout output, show the command and the prompt:

```text
$ ni run --dir . --target human-team --max-chars 4000

<compiled prompt>

`ni` compiled this prompt only. It did not execute implementation.
```

For file output, show the command and path:

```text
$ ni run --dir . --target codex --max-chars 4000 --out .ni/generated/codex.prompt.txt

Prompt written to .ni/generated/codex.prompt.txt.
`ni` compiled this prompt only. It did not execute implementation.
```

## Blocked Behavior

If `.ni/plan.lock.json` is missing, `ni-run` is blocked before prompt
compilation. The user should return to `ni-end` after the plan is ready.

If `ni run` reports a stale lock or hash mismatch, report `BLOCKED` and stop.
Do not repair the lockfile, edit locked planning docs, or produce a hand-written
substitute prompt.

## Boundary

`ni-run` does not call `codex exec`, shells, adapters, queues, PR automation,
test runners for downstream work, or any target runtime. It runs or requests
only NI validation and prompt compilation commands needed to prove the lock and
produce the bounded prompt.
