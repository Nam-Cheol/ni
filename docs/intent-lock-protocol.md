# Intent Lock Protocol

The Intent Lock Protocol is the core mechanism of `ni-kernel`: a deterministic
pre-runtime control layer that defines how planning conversations become a
project contract, when that contract is ready to lock, how the accepted plan is
hashed, what downstream actors may trust, and when handoff must stop because
intent changed.

In short:

```text
conversation -> accepted contract -> readiness gate -> lock hash -> bounded handoff
```

The public rule is simple: downstream execution must not start from ambiguous
or stale intent.

For the full protocol details, including inputs, gates, source-of-truth
precedence, outputs, stale-lock behavior, and non-goals, read
[Protocol Specification](42_INTENT_LOCK_PROTOCOL.md).
