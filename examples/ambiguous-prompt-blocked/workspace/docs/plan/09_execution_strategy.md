# Execution strategy

## v0 execution strategy

Do not execute implementation automatically. Use `ni status` to expose blocker
questions. After the user answers, use `ni end` to lock and `ni run` to compile
a short target prompt.

## Stop condition

If the plan is locked and later intent changes, downstream work must stop until
the plan is amended or relocked.
