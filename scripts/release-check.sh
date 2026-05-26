#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
QUICKSTART_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-release-check.XXXXXX")"

trap 'rm -rf "$QUICKSTART_TMP"' EXIT

cd "$ROOT"

run_step() {
  local label="$1"
  shift
  echo "release-check: $label" >&2
  "$@"
}

require_output() {
  local expected="$1"
  local file="$2"
  if ! grep -Fq "$expected" "$file"; then
    echo "release-check failed: expected output to contain: $expected" >&2
    sed -n '1,120p' "$file" >&2
    return 1
  fi
}
export -f require_output

run_step "release checklist is complete" python3 - <<'PY'
from pathlib import Path

path = Path("docs/23_RELEASE_CHECKLIST_v0.1.0.md")
text = path.read_text(encoding="utf-8")

required = [
    "CI passes",
    "Smoke passes",
    "Golden tests pass",
    "Schemas validate",
    "README quickstart verified",
    "No stale roadmap references",
    "No core-boundary violations",
    "No release automation claims",
    "No runtime execution claims",
    "All public commands listed in README have smoke coverage",
    "bash scripts/release-check.sh",
]

missing = [item for item in required if item not in text]
if missing:
    raise SystemExit(f"{path} is missing required checklist items: {missing}")
PY

run_step "CI workflow covers release validation" python3 - <<'PY'
from pathlib import Path

path = Path(".github/workflows/ci.yml")
text = path.read_text(encoding="utf-8")

required = [
    "push:",
    "pull_request:",
    "go-version: \"1.22.x\"",
    "go test ./...",
    "bash scripts/quality.sh",
    "bash scripts/smoke.sh",
]

missing = [item for item in required if item not in text]
if missing:
    raise SystemExit(f"{path} is missing required CI entries: {missing}")
PY

run_step "schemas validate" python3 scripts/check-schema.py
run_step "core boundary has no violations" python3 scripts/check-core-boundary.py --self-test
run_step "Go tests pass" go test ./...
run_step "golden tests pass" go test ./cmd/ni -run Golden -count=1
run_step "smoke passes" bash scripts/smoke.sh

run_step "README quickstart works in go run mode" bash -c '
  go run ./cmd/ni --help >"$1/go-run-help.out"
  require_output "ni is a project intent compiler" "$1/go-run-help.out"
  go run ./cmd/ni version >"$1/go-run-version.out"
  require_output "0.0.0-dev" "$1/go-run-version.out"
  go run ./cmd/ni init --dir "$1/plan" --profile prototype >"$1/init.out"
  require_output "initialized ni planning workspace" "$1/init.out"
  go run ./cmd/ni status --dir "$1/plan" >"$1/status.out"
  require_output "BLOCKED" "$1/status.out"
' bash "$QUICKSTART_TMP"

run_step "README quickstart works in built binary mode" bash -c '
  make build
  ./bin/ni --help >"$1/bin-help.out"
  require_output "ni is a project intent compiler" "$1/bin-help.out"
  ./bin/ni version >"$1/bin-version.out"
  if [[ ! -s "$1/bin-version.out" ]]; then
    echo "release-check failed: built binary version output is empty" >&2
    exit 1
  fi
' bash "$QUICKSTART_TMP"

run_step "README quickstart works in local install mode" bash -c '
  make install-local BINDIR="$1/bin"
  "$1/bin/ni" version >"$1/install-version.out"
  if [[ ! -s "$1/install-version.out" ]]; then
    echo "release-check failed: installed binary version output is empty" >&2
    exit 1
  fi
' bash "$QUICKSTART_TMP"

run_step "roadmap has no stale release references" python3 - <<'PY'
from pathlib import Path
import re

paths = [
    Path("docs/08_ROADMAP.md"),
    Path("docs/20_NEXT_STEPS.md"),
    Path("docs/19_RELEASE_NOTES_v0.1.md"),
]

required_markers = {
    Path("docs/08_ROADMAP.md"): [
        "v0.1 release-candidate shape is complete",
        "Next: v0.1.0 stabilization",
    ],
    Path("docs/20_NEXT_STEPS.md"): [
        "## v0.1.0",
    ],
    Path("docs/19_RELEASE_NOTES_v0.1.md"): [
        "Release Notes: v0.1 Release Candidate",
    ],
}

stale_patterns = [
    re.compile(r"\bTODO\b"),
    re.compile(r"\bTBD\b"),
    re.compile(r"\bcoming soon\b", re.IGNORECASE),
    re.compile(r"\bnot yet implemented\b", re.IGNORECASE),
    re.compile(r"\bplanned for v0\.1\b", re.IGNORECASE),
]

for path in paths:
    text = path.read_text(encoding="utf-8")
    missing = [marker for marker in required_markers[path] if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing release roadmap markers: {missing}")
    for pattern in stale_patterns:
        match = pattern.search(text)
        if match:
            raise SystemExit(f"{path} contains stale roadmap language: {match.group(0)}")
PY

run_step "no release automation or runtime execution claims" python3 - <<'PY'
from pathlib import Path

phrases = ["release automation", "runtime execution"]
boundary_markers = [
    "not",
    "no ",
    "do not",
    "must not",
    "does not",
    "should not",
    "without",
    "out of scope",
    "excluded",
    "not included",
    "outside",
    "downstream",
    "do not add",
    "not add",
]

paths = [Path("README.md")]
paths.extend(sorted(Path("docs").glob("*.md")))

for path in paths:
    lines = path.read_text(encoding="utf-8").splitlines()
    for index, line in enumerate(lines):
        lowered = line.lower()
        matched = [phrase for phrase in phrases if phrase in lowered]
        if not matched:
            continue
        start = max(0, index - 10)
        context = "\n".join(lines[start : index + 1]).lower()
        if not any(marker in context for marker in boundary_markers):
            raise SystemExit(
                f"{path}:{index + 1} contains an affirmative claim about {matched}"
            )
PY

run_step "README public commands have smoke coverage" python3 - <<'PY'
from pathlib import Path
import re

readme = Path("README.md").read_text(encoding="utf-8")
smoke = Path("scripts/smoke.sh").read_text(encoding="utf-8")

readme_commands = set()
for match in re.finditer(
    r"(?:go run \./cmd/ni|(?:\./bin/ni|~/.local/bin/ni)|`ni)\s+([a-z][a-z-]*)",
    readme,
):
    readme_commands.add(match.group(1))

smoke_commands = set(re.findall(r'run_cmd\s+"ni\s+([a-z][a-z-]*)', smoke))

missing = sorted(readme_commands - smoke_commands)
if missing:
    raise SystemExit(
        "README lists public commands without smoke coverage: " + ", ".join(missing)
    )
PY

echo "release-check: v0.1.0 release gate passed"
