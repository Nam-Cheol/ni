# Next Steps

`ni` v0.1 RC is a pre-runtime project intent compiler. The next phase should
stabilize that boundary before adding any downstream execution work.

## v0.1.0

- Keep README, release notes, roadmap, prompt archive, and examples consistent
  with the completed v0.1 RC.
- Re-run `bash scripts/quality.sh` before release cuts.
- Verify that locked plans remain the source of truth for prompt compilation
  and seed exports.
- Improve documentation only where it clarifies the kernel boundary or release
  use.

## v0.2

- Improve target-specific seed exports for downstream systems.
- Refine generated harness candidate review without turning candidates into a
  kernel-owned queue.
- Make feedback, pressure, amendment, and collaboration workflows easier to
  audit.
- Add deterministic consistency checks where documentation or seed output can
  drift.

## Do Not Add To Kernel

- runtime execution,
- Codex or shell adapters,
- evidence runners,
- queues,
- PR automation,
- release automation,
- plugin systems,
- TUI or web UI.
