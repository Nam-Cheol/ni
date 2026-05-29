# Execution strategy

## v0 execution strategy

Do not execute implementation automatically. This benchmark workspace is
expected to remain `BLOCKED` until the dashboard intent questions are answered.
If a future update resolves the blockers and `ni status` reports `READY` or
`READY_WITH_DEFERRALS`, `ni end` may lock only this workspace and `ni run` may
compile a bounded handoff prompt. The compiled prompt must remain inert seed
material and must not execute downstream work.
