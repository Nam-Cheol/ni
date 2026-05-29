# Demo Verification

`scripts/demo-check.sh` verifies that the public demos shown in the README still
work from source. It is a source-level proof check for `ni-kernel`, not a
downstream execution harness.

Run it from the repository root:

```bash
bash scripts/demo-check.sh
```

The script checks:

- `examples/ambiguous-prompt-blocked/workspace` reports `BLOCKED`.
- `--next-questions` renders blocker questions for the ambiguous prompt demo.
- `examples/research-protocol` reports the status documented in its README.
- the research protocol `human-team` prompt compiles to a temporary file when
  the example is locked.
- `examples/conversation-product` reports the status documented in its README.
- the conversation product `human-team` prompt compiles to a temporary file when
  the example is locked.
- `examples/ni-start-dogfood/workspace` reports the status documented in its
  README.
- the ni-start dogfood `human-team` prompt compiles to a temporary file when
  the example is locked.
- `examples/conversation-authoring` reports the status documented in its
  README.
- the conversation-authoring `human-team` prompt compiles from its existing
  lock to a temporary file.
- `examples/namba-ai-upgrade` reports the status documented in its README.
- the namba-ai upgrade `codex` prompt compiles from its existing lock to a
  temporary file without invoking Codex.
- `examples/benchmark-report` remains a docs-only benchmark report template
  with `not_measured` placeholders and benchmark non-execution markers.

The script must stay deterministic. It must not call Codex, model APIs, shell
adapters, Hyper Run, namba-ai, Spec Kit, Ouroboros, or any downstream agent
runtime. Prompt outputs are written under a temporary directory and removed when
the script exits.

`scripts/release-check.sh` runs `scripts/demo-check.sh` because the README demos
are release proof assets. `scripts/quality.sh` does not run it directly; use the
demo check when changing README demos, example workspaces, lock validation,
status output, or prompt compilation behavior.
