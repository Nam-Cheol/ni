# Conversation authoring fixture

This fixture shows the expected conversational planning loop after `ni init`.

It is not a runnable harness and it does not execute downstream tools. The
`transcript.md` demonstrates how a model summarizes current planning state, asks
focused questions, persists the answer into `docs/plan/**` and
`.ni/contract.json`, then relies on `ni status` for readiness.

`ni-end-confirmation.md` demonstrates how `ni-end` summarizes a CLI-ready plan,
asks for explicit confirmation, and only then lets `ni end` write the lock.

See [transcript.md](transcript.md) and
[ni-end-confirmation.md](ni-end-confirmation.md).
