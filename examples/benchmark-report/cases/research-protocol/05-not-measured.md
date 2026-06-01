# What was not measured

This resolved research-protocol benchmark case still did not measure:

- fieldwork;
- participant recruitment;
- data collection;
- analysis;
- intervention placement decisions;
- research protocol quality;
- fieldwork readiness in the real world;
- research approval or ethics approval;
- downstream agent performance;
- research outcome validity;
- adoption;
- cost;
- latency;
- rework reduction;
- statistical effect size.

No fieldwork was done. No participant recruitment was done. No data collection
was done. No analysis was done. No intervention placement was decided. No
downstream agent was run. No research outcome was validated. No adoption, cost,
latency, rework, or statistical effect was measured.

No model API was called. No generated prompt was executed. No shell adapter,
queue, telemetry path, runtime harness, release automation, issue publishing,
PR automation, or implementation work was added.

`ni status --proof --next-questions` reported `READY` only after synthetic
benchmark fixture answers were applied to the isolated workspace. `ni end`
created a lock inside that workspace, and `ni run --max-chars 4000` compiled a
bounded prompt seed. Those artifacts prove bounded intent handoff only; they
do not prove real research approval or real fieldwork readiness.
