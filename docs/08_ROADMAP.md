# Roadmap

The v0.1 release-candidate shape is complete: `ni-kernel` can initialize a
planning workspace, validate readiness deterministically, lock a ready plan,
compile a bounded prompt, and export locked-plan seed material. The kernel
remains pre-runtime; downstream harnesses and adapters may consume derived
material, but they are not kernel-owned execution state.

## Completed: v0.1 RC

### Phase 0: reframe

```text
SPEC-000 reframe docs to Project Intent Compiler for AI Agents
```

### Phase 1: kernel CLI

The contract/readiness/lock/prompt kernel is implemented:

```text
SPEC-001 bootstrap CLI
SPEC-002 project docs template
SPEC-003 contract model
SPEC-004 readiness status
SPEC-005 lockfile end
SPEC-006 prompt compiler run
```

### Phase 2: Codex UX and generated harness proposals

```text
SPEC-007 Codex skills
SPEC-008 work graph proposal
SPEC-009 generated harness contract
```

### Phase 3: dogfood and v0.1 release candidate

```text
SPEC-010 ni plans ni
```

The v0.1 RC also includes target registry entries, seed exports, readiness
profiles, product-shape fields, inert feedback and pressure tracking, amendment
and relock flow, and collaboration checks.

## Next: v0.1.0 stabilization

The next phase should make the release candidate shippable without expanding
the kernel boundary:

- keep README, release notes, roadmap, prompt archive, and docs aligned;
- tighten examples and target documentation around locked-plan seed material;
- verify `ni status`, `ni end`, `ni run`, and `ni export` behavior against the
  documented source-of-truth rule;
- preserve the rule that `ni run` compiles prompts only;
- do not add shell adapters, Codex adapters, queues, evidence runners, PR automation,
  release automation, plugin systems, and UI work.

## Next: v0.2 planning

v0.2 may improve downstream seed quality while keeping downstream state outside
the kernel:

- richer target-specific seed packages;
- clearer generated harness candidate review flows;
- better feedback-to-amendment workflows;
- additional deterministic checks for collaboration, pressure, and target
  export consistency.

## Historical experiment notes

```text
SPEC-011 Codex exec experiment
SPEC-012 shell adapter experiment
SPEC-013 evidence runner experiment
```

These entries were early execution-experiment placeholders. They are not active
kernel roadmap items for v0.1.0 or v0.2. The retained Codex exec material lives
only as a local experiment under `docs/experiments/` and `prompts/012-*`; it
must not become `ni run` behavior or a Codex adapter. Shell adapters and
evidence runners remain out of scope for the kernel.
