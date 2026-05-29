#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
RELEASE_VERSION="v0.4.0"

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

run_step "verify release plan, preflight, and manual steps exist" python3 - <<'PY'
from pathlib import Path

version = "v0.4.0"
required = [
    Path("docs/84_RELEASE_PLAN_v0.4.0.md"),
    Path("docs/84_RELEASE_PLAN_v0.4.0.ko.md"),
    Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.md"),
    Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.ko.md"),
]

for path in required:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

plan = Path("docs/84_RELEASE_PLAN_v0.4.0.md").read_text(encoding="utf-8")
preflight = Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.md").read_text(encoding="utf-8")

required_plan_markers = [
    f"# ni {version}",
    "Conversation Authoring Hardening",
    "First-run `ni-start` conversation card",
    "`SYNC-014`",
    "`SYNC-015`",
    "`SYNC-016`",
    "Grouped question output",
    "Model workspace packs",
    "No-terminal",
    "Homebrew remains Planned",
    "No task runner.",
    "SPEC runner",
    "Codex exec adapter",
    "No shell adapter.",
    "No queue.",
    "PR automation",
    "No release automation inside `ni-kernel`.",
    "`ni run` remains a prompt compiler only",
]
missing = [marker for marker in required_plan_markers if marker not in plan]
if missing:
    raise SystemExit(f"release plan is missing markers: {missing}")

required_manual_markers = [
    "git status --short",
    "bash scripts/release-dry-run.sh",
    "git tag -a v0.4.0 -m",
    "git push origin v0.4.0",
    "Wait for the GitHub Actions release workflow",
    "Verify release assets",
    "Verify the current-platform binary",
    "Verify the curl installer",
]
missing = [marker for marker in required_manual_markers if marker not in preflight]
if missing:
    raise SystemExit(f"preflight manual release steps are missing markers: {missing}")

forbidden_claims = [
    "Homebrew: Available",
    "published binary packages are available",
    "curl installer is available",
    "brew install ni",
]
for claim in forbidden_claims:
    if claim in plan or claim in preflight:
        raise SystemExit(f"v0.4.0 docs appear to claim availability: {claim}")
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
