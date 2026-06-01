# Execution strategy

## v0 execution strategy

Do not execute implementation automatically. This benchmark workspace measures
whether the answer packet makes the planning-meeting artifact ready, not whether
the dashboard product is ready. If `ni status` reports `READY` or
`READY_WITH_DEFERRALS`, `ni end` may lock only this workspace and `ni run
--max-chars 4000` may compile a bounded handoff prompt. The compiled prompt
must remain inert seed material and must not execute downstream work.
