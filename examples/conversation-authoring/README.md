# Conversation authoring fixture

This fixture shows the expected conversational planning loop after `ni init`.

It is not a runnable harness and it does not execute downstream tools.
`transcript.md` demonstrates how a model summarizes current planning state,
asks focused questions, persists the answer into `docs/plan/**` and
`.ni/contract.json`, then relies on `ni status` for readiness.

The completed example state is checked in under:

- `docs/plan/**`: generated human-facing plan docs after conversation resolved
  the blocker question.
- `.ni/contract.json`: matching machine-readable planning contract.
- `.ni/session.json`: bounded resume state, below docs and contract in
  authority.
- `.ni/plan.lock.json`: CLI-written lockfile for the completed plan.
- `generated/human-team.prompt.txt`: CLI-compiled handoff prompt.

The user never types contract authoring commands such as `contract add`,
`contract set`, or `contract list`; the model maintains the docs and contract
while the CLI validates, locks, and compiles.

`ni-end-confirmation.md` demonstrates how `ni-end` summarizes a CLI-ready plan,
asks for explicit confirmation, and only then lets `ni end` write the lock.

`ni-run-handoff.md` demonstrates how `ni-run` infers or asks for a target,
lets `ni run` verify the lock while compiling the prompt, and states that NI
did not execute implementation.

`session-resume.md` demonstrates how a later `ni-start` session resumes from
persisted docs, `.ni/contract.json`, and bounded `.ni/session.json` state
instead of hidden chat memory.

See [transcript.md](transcript.md),
[ni-end-confirmation.md](ni-end-confirmation.md),
[ni-run-handoff.md](ni-run-handoff.md), and
[session-resume.md](session-resume.md).
