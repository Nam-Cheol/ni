# Conversation authoring fixture

This fixture shows the expected `ni-start` planning loop after `ni init`.

It is not a runnable harness and it does not execute downstream tools. The
transcript demonstrates how a model summarizes current planning state, asks
focused questions, persists the answer into `docs/plan/**` and
`.ni/contract.json`, then relies on `ni status` for readiness.

See [transcript.md](transcript.md).
