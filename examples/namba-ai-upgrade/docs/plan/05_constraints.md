# Constraints

## Hard constraints

- Do not modify namba-ai in this example.
- Do not run namba-ai.
- Do not call Codex exec.
- Do not create a SPEC runner.
- Do not create execution tasks.
- Do not add shell, Codex, queue, PR, release, plugin, TUI, or web execution machinery.
- Readiness must come from `ni status`, not model judgment.
- Locking must come from `ni end`, not manual lockfile edits.
- Prompt output from `ni run` must be 4000 characters or less.

## Current limitation constraints

The upgrade plan treats namba-ai's current limitations as planning inputs, not proof that the namba-ai implementation is broken. Downstream implementers must verify repository facts before changing namba-ai.

## Pre-runtime compiler relation

NI sits before runtime orchestration:

```text
human intent -> docs/plan + .ni/contract.json -> ni status -> ni end -> ni run/export -> downstream tool
```

namba-ai may later consume seed material, but NI must not become namba-ai's runtime scheduler or state store.
