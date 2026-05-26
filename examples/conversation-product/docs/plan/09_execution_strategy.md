# Execution strategy

## v0 execution strategy

Do not execute implementation automatically. Use `ni run` to compile a short prompt after the plan is locked.

Downstream work may turn the locked conversation protocol into training material, transcript fixtures, or a service playbook. That work is outside ni-kernel and must not become kernel-owned runtime state.
