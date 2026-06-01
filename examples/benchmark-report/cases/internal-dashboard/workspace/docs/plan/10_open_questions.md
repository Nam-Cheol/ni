# Open questions

## OQ-001: Who is the primary benchmark artifact user, and what decision should the artifact support?

Blocker: false

Status: resolved

Resolution: The primary user is the planning meeting owner, product lead, or
internal operations lead. Secondary users are the engineering lead, QA or
reviewer, documentation maintainer, and stakeholders. The supported decision is
whether the internal-dashboard benchmark case has enough structure, evidence,
and acceptance criteria to be used in a planning meeting. This does not decide
final product direction, implementation scope, production release readiness, or
unresolved blocker resolution.

## OQ-002: What artifact signals need attention, and how should they be ranked?

Blocker: false

Status: resolved

Resolution: Meeting-readiness signals are completeness of required fields,
clarity of decision boundaries, presence of acceptance evidence, separation of
required answers from optional notes, and absence of unsupported assumptions.
Ranking prioritizes acceptance-critical items first: supported decision,
pass/fail criteria, minimum artifact, approval owner, then source and privacy
constraints. Subjective confidence, visual polish alone, implementation
preference, speculative stakeholder intent, undocumented assumptions, and
unrelated metrics are excluded.

## OQ-003: Which source systems, fields, freshness rules, privacy constraints, and access roles may the benchmark packet use?

Blocker: false

Status: resolved

Resolution: Allowed sources are project documentation, benchmark case files,
review notes, planning artifacts, issue or PR references, and approved internal
dashboard source material required to validate the benchmark case. Allowed
fields are case name, artifact path, benchmark objective, required answer
fields, optional notes, acceptance criteria, review status, timestamps, and
non-sensitive operational context. Prohibited fields include personal data,
credentials, tokens, private customer data, confidential business metrics,
production secrets, sensitive raw logs, and unrelated production telemetry.
Read access is limited to project members who need the benchmark case for
planning or review.

## OQ-004: What acceptance evidence must be ready for the planning meeting?

Blocker: false

Status: resolved

Resolution: The audience is the planning owner, engineering or review lead, and
relevant stakeholders. Timing is the next scheduled planning meeting, with the
specific date currently unassigned. The minimum artifact is a completed answer
packet with all required `OQ` fields, clear pass/fail criteria, explicit
non-goals, and enough evidence references for review. The packet passes if all
required fields are complete, the supported decision is clear, acceptance
criteria are testable, privacy boundaries are explicit, and unresolved blockers
are marked instead of answered implicitly. It fails if required fields are
blank, criteria are vague, sensitive data is included, source freshness is
unknown, or blockers are resolved without approved evidence. Final
implementation, production deployment, full technical design, complete
dashboard UI, exhaustive metrics analysis, and resolution of every blocker are
not required.
