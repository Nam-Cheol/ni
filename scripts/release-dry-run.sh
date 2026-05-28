#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
RELEASE_VERSION="v0.3.0"

cd "$ROOT"

run_step() {
  local label="$1"
  shift
  echo "release-dry-run: $label" >&2
  "$@"
}

echo "release-dry-run: planned release version $RELEASE_VERSION" >&2

run_step "go tests" go test ./...
run_step "quality checks" bash scripts/quality.sh
run_step "smoke checks" bash scripts/smoke.sh
run_step "public demo checks" bash scripts/demo-check.sh
run_step "source, build, and local install checks" bash scripts/install-check.sh
run_step "release readiness gate" bash scripts/release-check.sh

run_step "verify release notes and manual steps exist" python3 - <<'PY'
from pathlib import Path

version = "v0.3.0"
required = [
    Path("docs/68_RELEASE_NOTES_v0.3.0.md"),
    Path("docs/68_RELEASE_NOTES_v0.3.0.ko.md"),
    Path("docs/69_MANUAL_RELEASE_STEPS.md"),
    Path("docs/69_MANUAL_RELEASE_STEPS.ko.md"),
]

for path in required:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

notes = Path("docs/68_RELEASE_NOTES_v0.3.0.md").read_text(encoding="utf-8")
manual = Path("docs/69_MANUAL_RELEASE_STEPS.md").read_text(encoding="utf-8")

required_note_markers = [
    f"# ni {version} - Project Intent Compiler for AI Agents",
    f"Tag suggestion: `{version}`",
    "Project Intent Compiler positioning",
    "README as product pamphlet",
    "Local deterministic SVG visual system",
    "Intent Lock Protocol",
    "Ambiguous prompt blocked demo",
    "Non-software demos",
    "Status proof",
    "Model workspace packs",
    "Source-first usage",
    "Release binaries and curl installer availability",
    "available only after",
    "Task runner",
    "SPEC runner",
    "Multi-agent execution layer",
    "Codex exec adapter",
    "Shell adapter",
    "Queue",
    "PR automation",
    "Release automation inside `ni-kernel`",
    "Downstream execution runtime",
    "`ni run` compiles a prompt only",
]
missing = [marker for marker in required_note_markers if marker not in notes]
if missing:
    raise SystemExit(f"release notes are missing markers: {missing}")

required_manual_markers = [
    "git status --short",
    "bash scripts/release-dry-run.sh",
    "git tag -a v0.3.0 -m",
    "git push origin v0.3.0",
    "Wait for the GitHub Actions release workflow",
    "Confirm checksums match",
    "Only after assets and checksums exist",
    "Do not mark curl installer availability as available",
]
missing = [marker for marker in required_manual_markers if marker not in manual]
if missing:
    raise SystemExit(f"manual release steps are missing markers: {missing}")

forbidden_claims = [
    "published binary packages are available",
    "curl installer is available",
    "brew install ni",
]
for claim in forbidden_claims:
    if claim in notes:
        raise SystemExit(f"release notes appear to claim availability: {claim}")
PY

if command -v goreleaser >/dev/null 2>&1; then
  run_step "GoReleaser config check" goreleaser check
  run_step "GoReleaser snapshot release" goreleaser release --snapshot --clean
else
  cat >&2 <<'EOF'
release-dry-run: GoReleaser is not installed, so the GoReleaser check and
release-dry-run: snapshot archive build were NOT RUN.
release-dry-run: install GoReleaser, then rerun:
release-dry-run:   brew install goreleaser
release-dry-run:   bash scripts/release-dry-run.sh
EOF
fi

echo "release-dry-run: completed without creating tags, pushing, or publishing"
