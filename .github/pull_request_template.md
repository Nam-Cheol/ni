## Summary

<!-- What changed, and why? -->

## Boundary Check

<!-- ni is a Project Intent Compiler, not an execution runtime. -->

- [ ] This keeps task runners, SPEC runners, Codex exec, queues, agent teams, and PR/release automation out of `ni-kernel`.
- [ ] This change validates, locks, compiles, documents, or derives inert seed material from project intent.
- [ ] This does not create kernel-owned downstream execution state.

## Validation

<!-- Include command output or explain why a command was not run. -->

- [ ] `bash scripts/quality.sh`
- [ ] `go test ./...` if Go files were touched

## Notes

<!-- Risks, open questions, or follow-up work. -->
