#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

run_step() {
  local label="$1"
  shift
  echo "launch-check: $label" >&2
  "$@"
}

run_script_if_present() {
  local script="$1"
  if [[ -f "$script" ]]; then
    run_step "$script" bash "$script"
  else
    echo "launch-check: skipping missing $script" >&2
  fi
}

run_step "launch checklist docs are complete" python3 - <<'PY'
from pathlib import Path

required_sections = [
    "## Positioning",
    "## README Rendering",
    "## README.ko Parity",
    "## License",
    "## Security Policy",
    "## CI",
    "## Demo Verification",
    "## Benchmark Protocol",
    "## Release Draft",
    "## Install/Build Verification",
    "## Issue Templates",
    "## No False Claims",
    "## No Execution-Runtime Drift",
    "## Local Gate",
]

docs = [Path("docs/50_LAUNCH_CHECKLIST.md")]
if Path("README.ko.md").exists():
    docs.append(Path("docs/50_LAUNCH_CHECKLIST.ko.md"))

for path in docs:
    if not path.exists():
        raise SystemExit(f"{path} is missing")
    text = path.read_text(encoding="utf-8")
    missing = [section for section in required_sections if section not in text]
    if missing:
        raise SystemExit(f"{path} is missing launch checklist sections: {missing}")

    required_markers = [
        "Project Intent Compiler for AI Agents",
        "ni-kernel",
        "README.md",
        "README.ko.md",
        "LICENSE",
        "SECURITY.md",
        ".github/workflows/ci.yml",
        "scripts/demo-check.sh",
        "docs/43_BENCHMARK_PROTOCOL.md",
        "docs/47_RELEASE_DRAFT_v0.2.0.md",
        "scripts/install-check.sh",
        ".github/ISSUE_TEMPLATE/bug_report.md",
        ".github/ISSUE_TEMPLATE/feature_request.md",
        ".github/ISSUE_TEMPLATE/boundary_question.md",
        "No False Claims",
        "No Execution-Runtime Drift",
        "bash scripts/launch-check.sh",
    ]
    missing = [marker for marker in required_markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing launch checklist markers: {missing}")

    boundary_markers = [
        "does not publish",
        "does not publish packages",
        "하지 않는다",
        "절대 하지 않는다",
    ]
    if not any(marker in text for marker in boundary_markers):
        raise SystemExit(f"{path} does not state the non-publishing boundary")
PY

run_step "public launch resources exist" python3 - <<'PY'
from pathlib import Path

required_paths = [
    Path("README.md"),
    Path("README.ko.md"),
    Path("LICENSE"),
    Path("SECURITY.md"),
    Path("SECURITY.ko.md"),
    Path(".github/workflows/ci.yml"),
    Path(".github/ISSUE_TEMPLATE/bug_report.md"),
    Path(".github/ISSUE_TEMPLATE/feature_request.md"),
    Path(".github/ISSUE_TEMPLATE/boundary_question.md"),
    Path("docs/22_INSTALL.md"),
    Path("docs/40_POSITIONING.md"),
    Path("docs/42_INTENT_LOCK_PROTOCOL.md"),
    Path("docs/43_BENCHMARK_PROTOCOL.md"),
    Path("docs/46_RELEASE_READINESS.md"),
    Path("docs/46_RELEASE_READINESS.ko.md"),
    Path("docs/47_RELEASE_DRAFT_v0.2.0.md"),
    Path("docs/47_RELEASE_DRAFT_v0.2.0.ko.md"),
    Path("docs/48_DEMO_VERIFICATION.md"),
    Path("docs/50_LAUNCH_CHECKLIST.md"),
    Path("docs/50_LAUNCH_CHECKLIST.ko.md"),
]

missing = [str(path) for path in required_paths if not path.exists()]
if missing:
    raise SystemExit("launch resources are missing: " + ", ".join(missing))
PY

run_step "README launch links and parity markers exist" python3 - <<'PY'
from pathlib import Path

expectations = {
    "README.md": [
        "docs/40_POSITIONING.md",
        "docs/42_INTENT_LOCK_PROTOCOL.md",
        "docs/43_BENCHMARK_PROTOCOL.md",
        "docs/45_TARGET_STORY.md",
        "docs/46_RELEASE_READINESS.md",
        "docs/50_LAUNCH_CHECKLIST.md",
        "SECURITY.md",
        "MIT License",
        "does not claim package distribution or a published binary release",
        "Public demo verification",
        "Source/build/install verification",
    ],
    "README.ko.md": [
        "docs/40_POSITIONING.md",
        "docs/42_INTENT_LOCK_PROTOCOL.md",
        "docs/43_BENCHMARK_PROTOCOL.md",
        "docs/45_TARGET_STORY.md",
        "docs/46_RELEASE_READINESS.ko.md",
        "docs/50_LAUNCH_CHECKLIST.ko.md",
        "SECURITY.md",
        "MIT License",
        "package distribution이나 published binary release를 claim하지 않는다",
        "Public demo verification",
        "Source/build/install verification",
    ],
}

for path, markers in expectations.items():
    text = Path(path).read_text(encoding="utf-8")
    missing = [marker for marker in markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing public launch markers: {missing}")
PY

run_step "CI and issue templates preserve public boundary" python3 - <<'PY'
from pathlib import Path

ci = Path(".github/workflows/ci.yml").read_text(encoding="utf-8")
required_ci = [
    "push:",
    "pull_request:",
    "go test ./...",
    "bash scripts/quality.sh",
    "bash scripts/smoke.sh",
]
missing_ci = [marker for marker in required_ci if marker not in ci]
if missing_ci:
    raise SystemExit(f".github/workflows/ci.yml is missing: {missing_ci}")

template_markers = {
    ".github/ISSUE_TEMPLATE/bug_report.md": [
        "ni-kernel",
        "downstream runtime",
        "task runner",
        "queue",
        "PR automation",
        "release automation",
    ],
    ".github/ISSUE_TEMPLATE/feature_request.md": [
        "ni-kernel",
        "execution runtime",
        "Codex exec",
        "shell adapter",
        "PR automation",
        "release automation",
    ],
    ".github/ISSUE_TEMPLATE/boundary_question.md": [
        "ni-kernel",
        "task runner",
        "queue",
        "Codex exec",
        "PR automation",
        "release automation",
    ],
}

for path, markers in template_markers.items():
    text = Path(path).read_text(encoding="utf-8")
    missing = [marker for marker in markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing boundary markers: {missing}")
PY

run_step "public docs avoid false launch claims" python3 - <<'PY'
from pathlib import Path

forbidden_affirmative = [
    "brew install ni",
    "go install github.com",
    "published binary release is available",
    "published binary packages are available",
    "download the binary",
    "automatically publishes",
    "creates a GitHub release",
    "runs downstream agents",
    "executes Codex",
    "executes shell commands",
    "starts a queue",
]

negation_markers = [
    "does not",
    "do not",
    "must not",
    "not claim",
    "not included",
    "outside",
    "without",
    "no ",
    "claim하지",
    "하지 않는다",
    "실행하지",
    "생성하지",
    "않는다",
    "아니다",
]

paths = [
    Path("README.md"),
    Path("README.ko.md"),
    Path("docs/22_INSTALL.md"),
    Path("docs/46_RELEASE_READINESS.md"),
    Path("docs/46_RELEASE_READINESS.ko.md"),
    Path("docs/47_RELEASE_DRAFT_v0.2.0.md"),
    Path("docs/47_RELEASE_DRAFT_v0.2.0.ko.md"),
    Path("docs/50_LAUNCH_CHECKLIST.md"),
    Path("docs/50_LAUNCH_CHECKLIST.ko.md"),
]

for path in paths:
    text = path.read_text(encoding="utf-8")
    lowered = text.lower()
    for claim in forbidden_affirmative:
        needle = claim.lower()
        start = 0
        while True:
            index = lowered.find(needle, start)
            if index == -1:
                break
            context = lowered[max(0, index - 160) : index + len(needle) + 160]
            if not any(marker in context for marker in negation_markers):
                raise SystemExit(f"{path} appears to make a forbidden launch claim: {claim}")
            start = index + len(needle)
PY

run_script_if_present scripts/quality.sh
run_script_if_present scripts/smoke.sh
run_script_if_present scripts/demo-check.sh
run_script_if_present scripts/release-check.sh
run_script_if_present scripts/install-check.sh

echo "launch-check: public launch gate passed without publishing or downstream runtime execution"
