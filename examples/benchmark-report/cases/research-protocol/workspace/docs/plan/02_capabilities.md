# Capabilities

## CAP-001: Capture the vague neighborhood cooling study request as an initial research-protocol benchmark workspace.

This capability records the request as a research_protocol planning case with
document, workflow, and human_service delivery surfaces. It traces to REQ-001
and REQ-002 and produces ART-001.

## CAP-002: Make missing research questions, participant scope, consent, data, safety, and acceptance evidence explicit as blocker questions.

This capability kept OQ-001 through OQ-005 open during the initial benchmark
measurement, then resolves them only with the clearly labeled synthetic
benchmark fixture answers. It traces to REQ-003, REQ-004, and REQ-007 and
produces ART-002, ART-003, ART-005, and ART-006.

## CAP-003: Preserve the pre-runtime benchmark boundary for research-protocol evidence.

This capability prevents the benchmark from becoming fieldwork, participant
recruitment, data collection, analysis, intervention placement, downstream
agent execution, model API execution, or runtime work. It traces to REQ-005 and
REQ-006 and produces ART-004.

## CAP-004: Compile a bounded handoff prompt only after a valid isolated lock.

This capability allows `ni end` and `ni run --max-chars 4000` only when
`ni status` reports `READY` or `READY_WITH_DEFERRALS` for the isolated
benchmark workspace. It traces to REQ-005 and REQ-007 and produces ART-007 and
ART-008.
