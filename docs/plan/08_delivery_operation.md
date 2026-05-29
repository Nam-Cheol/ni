# Delivery and operation

## Delivery

The v0.2 kernel ships as a local Go CLI with repository-local Codex skills and Markdown planning docs. The primary authoring surface after `ni init` is model-user conversation, not manual contract editing commands.

The v0.2 differentiation surface is documentation and proof-asset driven: README.md, README.ko.md, positioning docs, demos, benchmark protocol, status proof, target story, and release readiness checklist explain and verify the Intent Lock Protocol without starting downstream execution.

The v0.3 packaging surface adds README pamphlet strategy, SVG-first visual
identity, source-first plus release-binary distribution planning, and Codex- and
Claude-style model workspace packs. These surfaces help users understand and
adopt the kernel; they do not add `ni-kernel` runtime execution.

The v0.4 adoption surface builds on verified v0.3.0 release binaries, the
verified curl installer, and model workspace pack work. It prioritizes
fresh-user install verification, Homebrew as a planned package-manager path,
model workspace pack verification, README use-anywhere clarity, announcement
copy, one qualitative benchmark case study, and an optional lightweight static
landing page. These are adoption and proof surfaces, not kernel runtime state.

The v0.5 roadmap surface builds on v0.4 adoption hardening. It prioritizes real
benchmark data, more product surfaces and examples, more reliable conversation
authoring, stronger lock/relock/amendment and change-control UX, optional
Homebrew only if tested, optional factual landing page work, and optional
downstream integrations only as separate packages. These are evidence,
documentation, validation, distribution, and seed-quality surfaces, not
`ni-kernel` runtime state.

The v0.4.0 post-release surface records verified distribution state after the
published release assets, checksums, current-platform binary, curl installer,
and install docs were checked. Release binary and curl installer paths are
Available for verified v0.4.0 assets; Homebrew remains Planned / deferred; and
model workspace packs remain UX over docs and CLI proof.

## Delivery surfaces

- cli

## Accepted delivery surface details

The accepted delivery surface is `cli`. Users interact with `ni-kernel` through
the local command line for deterministic readiness, lock/relock, and bounded
prompt compilation. Documentation, model workspace packs, and generated prompts
support that CLI flow, but they do not replace CLI authority or add runtime
execution.

Tracked planning state:

- `docs/plan/**`
- `.ni/contract.json`
- `.ni/readiness.rules.json`
- `.ni/readiness.profiles.json`
- `.ni/plan.lock.json`

Generated or inert state:

- `.ni/generated/**`
- `.ni/feedback.jsonl`
- `.ni/pressure.json`
- `.ni/harness.candidates.json`
- `.ni/amendments/**`
- `.ni/locks/**`

Packaging and distribution references:

- `README.md`
- `README.ko.md`
- `docs/52_README_PAMPHLET_STRATEGY.md`
- `docs/53_DISTRIBUTION_STRATEGY.md`
- `docs/54_HOMEBREW_DISTRIBUTION.md`
- `docs/22_INSTALL.md`
- `docs/69_MANUAL_RELEASE_STEPS.md`
- `docs/55_MODEL_WORKSPACE_PACKS.md`
- `docs/56_CLAUDE_SKILL_PACK.md`
- `docs/57_CODEX_SKILL_PACK.md`
- `docs/58_VISUAL_ASSETS.md`
- `docs/70_RELEASE_VERIFICATION_v0.3.0.md`
- `docs/71_CURL_INSTALLER_VERIFICATION_v0.3.0.md`
- `docs/72_HOMEBREW_TAP_PLAN.md`
- `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`
- `docs/76_ANNOUNCEMENT_KIT.md`
- `docs/77_BENCHMARK_CASE_STUDY.md`
- `docs/78_LANDING_PAGE_PLAN.md`
- `docs/79_CONVERSATION_AUTHORING_UX_AUDIT.md`
- `docs/80_HOMEBREW_DECISION.md`
- `docs/81_RELEASE_PLAN_v0.4.1.md`
- `docs/84_RELEASE_PLAN_v0.4.0.md`
- `docs/85_RELEASE_PREFLIGHT_v0.4.0.md`
- `docs/86_RELEASE_VERIFICATION_v0.4.0.md`
- `docs/87_CURL_INSTALLER_VERIFICATION_v0.4.0.md`

## Operating model

1. Run `ni init` to create the initial planning structure.
2. Use model-user conversation with `ni-start` to update `docs/plan/**`, `.ni/contract.json`, and bounded `.ni/session.json` continuity state together.
3. Run `ni status` to get deterministic readiness gaps.
4. Resolve blockers without weakening accepted criteria.
5. Run `ni end` for a first lock or `ni relock` after an applied amendment.
6. Run `ni run --target <target> --max-chars 4000` to compile a bounded handoff prompt.
7. Treat downstream feedback as inert until it becomes an explicit amendment.
8. Use the differentiation proof assets to show why locked intent should precede downstream agent, human-team, or harness work.
9. Keep public packaging claims factual: release binaries and curl installer
   availability stay tied to verified v0.4.0 release evidence, package-manager
   availability stays planned until published and tested, model workspace packs
   remain UX, and no-terminal mode remains assisted unless CLI proof is
   supplied.
10. For v0.4 adoption, compile a locked Codex handoff prompt only after
    readiness passes and the amended plan is relocked. Do not execute the
    generated prompt or add downstream runtime behavior.
11. For v0.5 roadmap work, compile a locked Codex handoff prompt only after
    readiness passes and the amended plan is relocked. Do not execute the
    generated prompt, call `codex exec`, add adapters to `ni-kernel`, or mark
    Homebrew Available without tested Homebrew evidence.

## Validation

When Go code exists, run:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

For this v0.2 planning contract, also verify:

```bash
ni status --dir .
ni run --dir . --target codex --max-chars 4000
```

For the v0.5 roadmap lock, also verify:

```bash
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/demo-check.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
bash scripts/fresh-install-check.sh
```

Run `bash scripts/fresh-install-check.sh` only when the script is present.

For the v0.4.0 post-release state lock, also verify:

```bash
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/demo-check.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
go run ./cmd/ni run --dir . --target codex --max-chars 4000
```
