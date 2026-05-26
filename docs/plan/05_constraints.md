# Constraints

## Hard constraints

- `ni run` prompt output must be 4000 characters or less.
- Readiness must be rule-based, not model-feeling-based.
- Lockfile hash mismatch must block prompt compilation.
- Codex is an adapter/UX target, not the kernel.

## Initial technical constraints

- Prefer a small Go CLI.
- Avoid shell execution until after prompt compilation works.
- Keep JSON contract validation deterministic.
