# Capabilities

## CAP-001: Capture the benchmark artifact readiness scope as planning state

Record the original dashboard request, the user-provided `OQ-001` through
`OQ-004` answers, visible assumptions, non-goals, risks, and readiness evidence
without turning the benchmark case into dashboard implementation instructions.

## CAP-002: Validate the answer packet against planning-meeting readiness criteria

Check that the benchmark answer packet has all required fields filled, a clear
supported decision, testable pass/fail criteria, explicit privacy boundaries,
and no unsupported assumptions treated as resolved blockers.

## CAP-003: Preserve the non-execution benchmark boundary

Keep final product direction, implementation scope, production release
readiness, dashboard UI quality, downstream agent performance, model API calls,
and sensitive source data outside the benchmark case.

## CAP-004: Compile a bounded handoff prompt only after a valid isolated lock

If and only if the isolated workspace reaches `READY` or
`READY_WITH_DEFERRALS`, use `ni end` inside the workspace and then `ni run
--max-chars 4000` inside the workspace to compile a bounded prompt without
executing downstream work.
