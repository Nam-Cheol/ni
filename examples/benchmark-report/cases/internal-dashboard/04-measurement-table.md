# Measurement Table

This is one manual qualitative readiness assessment for the vague dashboard
request. It is not repeated benchmark data, not statistical evidence, and not a
claim about implementation outcomes.

| Criterion | Direct-to-agent risk | ni path evidence | Improved? |
| --- | --- | --- | --- |
| Missing acceptance criteria | Missing pass/fail checks for account health, priority ranking, freshness, performance, usability, and meeting acceptance. | Required before lock as REQ/EVAL records linked to dashboard capabilities. | Yes, for readiness visibility. |
| Unmitigated high-risk items | Customer data exposure, incorrect prioritization, and stale account signals are visible but unmitigated. | High-severity RISK records must include mitigation or explicit accepted rationale before readiness can pass. | Yes, as a gate expectation. |
| Unresolved blockers | Primary users, source systems, required fields, meeting date, and launch surface are unknown. | ni-start questions turn these into blocker questions; `ni status` should remain `BLOCKED` until answered or explicitly deferred. | Yes. |
| Hidden assumptions | Users, metrics, source systems, privacy review, deadline, and visualization format would be invented by the downstream actor. | Assumptions are expected to become open questions, accepted decisions, requirements, risks, or non-goals. | Yes. |
| Non-goal coverage | Missing; request does not exclude CRM replacement, workflow automation, forecasting, or write-back behavior. | Non-goals are expected contract records before lock. | Yes. |
| Delivery surface clarity | Assumed web dashboard, but prototype, report, embedded CRM view, or planning document are not distinguished. | Readiness interview guidance asks for delivery surface; docs and contract must agree. | Yes. |
| Actor/outcome clarity | "Customer team" and "who needs attention" are too broad to guide implementation. | Actor/outcome records must name who uses the dashboard and which decision it supports. | Yes. |
| Evaluation evidence clarity | No evidence is named for correctness, freshness, access, or meeting readiness. | Evaluation records are expected for data checks, prioritization review, usability review, and planning acceptance. | Yes. |
| Bounded handoff prompt availability | Unavailable; the direct prompt has no lock-verified compiled target prompt. | `not_measured`; no actual lock or `ni run` output exists for this case. | Not measured. |
