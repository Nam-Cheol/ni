# Domain and state model

## Core entities

```text
benchmark case
answer packet
required answer field
optional note
acceptance criterion
planning meeting owner
reviewer
stakeholder
evidence reference
review status
freshness marker
privacy/access boundary
bounded handoff prompt
```

## Attention and ranking rules

For this benchmark artifact, "needs attention" means a planning item that
directly affects meeting readiness. The highest-priority items are the
supported decision, pass/fail criteria, minimum artifact, approval owner, and
source/privacy constraints. Optional notes and explanatory context rank below
acceptance-critical fields.

A case passes review only if every required field is filled, the supported
decision is explicit, pass/fail criteria are testable, and no blocker is
silently resolved without evidence. Subjective confidence, visual polish alone,
implementation preference, speculative stakeholder intent, undocumented
assumptions, and metrics not tied to planning-meeting acceptance criteria are
excluded signals.

## Source and freshness rules

Allowed sources are project documentation, benchmark case files, review notes,
planning artifacts, issue or PR references, and approved internal dashboard
source material required to validate the benchmark case. The packet may include
case name, artifact path, benchmark objective, required answer fields, optional
notes, acceptance criteria, review status, timestamps, and non-sensitive
operational context.

The packet must use the most recent approved project artifacts available before
the planning meeting. If artifacts conflict, the newer approved artifact takes
priority unless it is marked draft, deprecated, or incomplete. Stale or
uncertain source status must be marked stale or TBD rather than treated as
current.
