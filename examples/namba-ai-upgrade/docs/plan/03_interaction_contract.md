# Interaction contract

## Interaction mode

Human to system. A human uses NI commands to initialize, review, lock, and compile planning artifacts.

## Product interaction

The interaction sequence for this example is:

```text
ni init -> edit docs/plan and .ni/contract.json -> ni status -> ni end -> ni run --target codex --out generated/codex.goal.prompt.txt
```

Only NI commands are part of the example. The compiled prompt is text. It is not an instruction for NI to start Codex, namba-ai, a shell adapter, a queue, or a PR automation loop.

## User control

The user controls whether any downstream tool receives the prompt. If downstream work discovers a contract issue, it should report the issue back as planning pressure or an amendment candidate rather than editing locked files silently.

## Codex-limited boundary

Codex may be a target audience for the generated prompt. Codex is not an authority for readiness, lock state, acceptance criteria, or source-of-truth order. If Codex is later used manually, it must verify `.ni/plan.lock.json` and stop with `BLOCKED` on mismatch.
