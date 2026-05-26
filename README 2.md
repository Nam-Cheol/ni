# ni

`ni` is a project intent compiler.

It turns ongoing planning conversations into a locked project contract, then derives a project-specific harness from that contract. The first product target is not a generic task runner. The first target is a small planning kernel with four authoritative operations:

1. maintain the planning contract,
2. check readiness with deterministic rules,
3. lock the accepted plan with hashes,
4. compile a short execution goal prompt from the locked plan.

Positioning: `ni` sits before execution runtimes. It is not a SPEC runner, multi-agent layer, task queue, or Hyper Run clone. It can produce downstream-compatible prompts or seed harness material after lock, but the kernel remains the authority for contract readiness, locking, and source-of-truth checks.

## Product boundary

`ni` has two layers.

```text
ni-kernel
  docs contract
  readiness gate
  lockfile
  prompt compiler
  source-of-truth rule

ni-generated-harness
  project-specific work graph
  project-specific evaluation plan
  project-specific evidence rules
  project-specific adapter choices
```

The kernel must stay small. The generated harness may vary by project.

## v0 goal

The v0 workflow is:

```text
ni init
  Create docs/plan templates and .ni/contract.json.
  Defaults to readiness profile prototype unless --profile is provided.

ni status
  Validate whether the planning contract is ready under the active profile.

ni end
  Refuse if blocked. If ready, create .ni/plan.lock.json.

ni run
  Read the locked plan and print a 4000-character-or-less goal prompt.
```

`ni run` does not directly execute Codex in v0. It only compiles the prompt. Execution adapters come later.

Readiness profiles are planning confidence profiles only: `concept`, `prototype`, `mvp`, `beta`, and `production`. They do not create implementation stages or execution lifecycle state.

## Source of truth

After `ni end`, the authority order is:

```text
.ni/plan.lock.json
.ni/contract.json
docs/plan/**
chat transcript
model inference
```

If hashes in `.ni/plan.lock.json` do not match current files, `ni run` must stop.

## Non-goals before v0 contract/readiness/lock/prompt works

Do not add these before the kernel is working:

- shell adapter,
- Codex exec adapter,
- queue,
- PR automation,
- release automation,
- hooks,
- web UI,
- multi-agent orchestration,
- generic SPEC runner.

## Development mode

Use the prompt files in `prompts/` sequentially. Treat each prompt as one small implementation unit and one commit.
