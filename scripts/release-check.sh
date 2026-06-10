#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
QUICKSTART_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-release-check.XXXXXX")"
CURRENT_RELEASE_VERSION="v0.6.3"
PLANNED_RELEASE_VERSION="post-release follow-up"

trap 'rm -rf "$QUICKSTART_TMP"' EXIT

cd "$ROOT"

echo "release-check: current release version $CURRENT_RELEASE_VERSION" >&2
echo "release-check: planned release version $PLANNED_RELEASE_VERSION" >&2

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

run_step "release readiness checklist is complete" python3 - <<'PY'
from pathlib import Path

path = Path("docs/46_RELEASE_READINESS.md")
text = path.read_text(encoding="utf-8")

required = [
    "quality passes",
    "tests pass",
    "install-check passes",
    "README and README.ko are in sync",
    "examples exist",
    "status proof works",
    "benchmark protocol exists",
    "v0.3.0 release notes exist",
    "no runtime execution claims",
    "no false release/license/CI/security claims",
    "bash scripts/release-check.sh",
    "bash scripts/release-dry-run.sh",
    "bash scripts/install-check.sh",
]

missing = [item for item in required if item not in text]
if missing:
    raise SystemExit(f"{path} is missing required checklist items: {missing}")

if Path("README.ko.md").exists() and not Path("docs/46_RELEASE_READINESS.ko.md").exists():
    raise SystemExit("README.ko.md exists, but docs/46_RELEASE_READINESS.ko.md is missing")
PY

run_step "v0.3.0 release notes are factual and release-binary prepared" python3 - <<'PY'
from pathlib import Path

drafts = [
    Path("docs/68_RELEASE_NOTES_v0.3.0.md"),
    Path("docs/68_RELEASE_NOTES_v0.3.0.ko.md"),
]

for path in drafts:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

required_markers = [
    "# ni v0.3.0 - Project Intent Compiler for AI Agents",
    "Tag suggestion: `v0.3.0`",
    "Summary: Don't run the agent yet. Compile the intent first.",
    "## Why v0.3.0",
    "## Included",
    "Project Intent Compiler positioning",
    "README as product pamphlet",
    "Local deterministic SVG visual system",
    "Intent Lock Protocol",
    "`ni init`, `ni status`, `ni end`, and `ni run`",
    "Status proof report",
    "Locked plan hash validation",
    "Ambiguous prompt blocked demo",
    "Non-software demos",
    "Model workspace packs",
    "Source-first usage",
    "## Not Included",
    "Task runner",
    "SPEC runner",
    "Multi-agent execution layer",
    "Codex exec adapter",
    "Shell adapter",
    "Queue",
    "PR automation",
    "Release automation inside `ni-kernel`",
    "Downstream execution runtime",
    "Release binary pipeline",
    "Package manager distribution",
    "bash scripts/release-dry-run.sh",
    "go test ./...",
    "bash scripts/quality.sh",
    "ni_<version>_linux_amd64.tar.gz",
    "ni_<version>_linux_arm64.tar.gz",
    "ni_<version>_darwin_amd64.tar.gz",
    "ni_<version>_darwin_arm64.tar.gz",
    "ni_<version>_windows_amd64.zip",
    "ni_<version>_checksums.txt",
]

boundary_markers = [
    "do not publish a release",
    "do not claim hosted release assets",
    "not part of `ni` runtime behavior",
    "does not execute Codex",
    "Release를 publish하거나",
    "claim하지 않는다",
    "release automation을 추가하지 않는다",
    "Codex, shells, model APIs",
]

for path in drafts:
    text = path.read_text(encoding="utf-8")
    missing = [marker for marker in required_markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing release draft markers: {missing}")

    if not any(marker in text for marker in boundary_markers):
        raise SystemExit(f"{path} is missing source-first/non-execution boundary markers")

    forbidden_claims = [
        "Homebrew install",
        "brew install",
        "published binary packages are available",
        "curl installer is available",
        "automatically publishes",
    ]
    for claim in forbidden_claims:
        context = text[max(0, text.find(claim) - 120): text.find(claim) + len(claim) + 120]
        allowed_negations = [
            "does not",
            "do not",
            "not claim",
            "claim하지",
            "추가하지",
            "실행하지",
            "않는다",
        ]
        if claim in text and not any(marker in context for marker in allowed_negations):
            raise SystemExit(f"{path} appears to make a forbidden release claim: {claim}")
PY

run_step "v0.4.0 release plan and preflight are factual" python3 - <<'PY'
from pathlib import Path

required_paths = [
    Path("docs/84_RELEASE_PLAN_v0.4.0.md"),
    Path("docs/84_RELEASE_PLAN_v0.4.0.ko.md"),
    Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.md"),
    Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.ko.md"),
]

for path in required_paths:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

required_markers = {
    Path("docs/84_RELEASE_PLAN_v0.4.0.md"): [
        "# ni v0.4.0",
        "Conversation Authoring Hardening",
        "No `v0.4.0` tag is",
        "`ni run` remains a prompt compiler only",
        "Homebrew remains Planned",
        "No task runner.",
        "No Codex exec adapter.",
        "No shell adapter.",
        "No downstream agents.",
        "No queue.",
        "Create an annotated tag for `v0.4.0`.",
    ],
    Path("docs/84_RELEASE_PLAN_v0.4.0.ko.md"): [
        "# ni v0.4.0",
        "Conversation Authoring Hardening",
        "`v0.4.0` tag가 없으므로",
        "`ni run`은 계속 prompt compiler only",
        "Homebrew",
        "Planned",
        "No task runner.",
        "No Codex exec adapter.",
        "No shell adapter.",
        "No downstream agents.",
        "No queue.",
        "`v0.4.0` annotated tag를 만든다.",
    ],
    Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.md"): [
        "# ni v0.4.0 Release Preflight",
        "HEAD == origin/main",
        "`v0.4.0` tag absent before release",
        "GoReleaser injects",
        "binary release should report `0.4.0`",
        "No runtime execution behavior",
        "goreleaser check",
        "goreleaser release --snapshot --clean",
        "git tag -a v0.4.0",
        "git push origin v0.4.0",
    ],
    Path("docs/85_RELEASE_PREFLIGHT_v0.4.0.ko.md"): [
        "# ni v0.4.0 Release Preflight",
        "HEAD == origin/main",
        "`v0.4.0` tag absent before release",
        "GoReleaser",
        "binary release는 `0.4.0`",
        "runtime execution behavior 없음",
        "goreleaser check",
        "goreleaser release --snapshot --clean",
        "git tag -a v0.4.0",
        "git push origin v0.4.0",
    ],
}

for path, markers in required_markers.items():
    text = path.read_text(encoding="utf-8")
    missing = [marker for marker in markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing v0.4.0 preflight markers: {missing}")

for path in required_paths:
    text = path.read_text(encoding="utf-8")
    forbidden_claims = [
        "Homebrew: Available",
        "brew install Nam-Cheol",
        "runtime execution is included",
        "Codex exec adapter is included",
        "shell adapter is included",
        "tag was pushed",
        "release was published",
    ]
    for claim in forbidden_claims:
        if claim in text:
            raise SystemExit(f"{path} appears to make a forbidden v0.4.0 claim: {claim}")
PY

run_step "v0.5.0 release candidate and publication docs are factual" python3 - <<'PY'
from pathlib import Path

required_paths = [
    Path("docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md"),
    Path("docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md"),
    Path("docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md"),
    Path("docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md"),
    Path("docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md"),
    Path("docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md"),
    Path("docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md"),
    Path("docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md"),
    Path("docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md"),
    Path("docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md"),
    Path("docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.md"),
    Path("docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md"),
    Path("docs/117_V0_5_0_POST_RELEASE_VERIFICATION.md"),
    Path("docs/117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md"),
]

for path in required_paths:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

required_markers = {
    Path("docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.md"): [
        "Decision: `RC_READY_WITH_DEFERRALS`.",
        "Homebrew: Planned / v0.5 candidate",
        "Model workspace packs: Experimental",
        "No-terminal method: Experimental / assisted",
        "Skills are UX; CLI is authority.",
    ],
    Path("docs/110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md"): [
        "Decision: `RC_READY_WITH_DEFERRALS`.",
        "Homebrew: Planned / v0.5 candidate",
        "Model workspace packs: Experimental",
        "No-terminal method: Experimental / assisted",
        "Skills are UX; CLI is authority.",
    ],
    Path("docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.md"): [
        "# v0.5 RC Polish / Release Notes Draft",
        "This document is a draft and does not publish, tag, or release v0.5.",
        "Release binary: Available for verified v0.4.0 release assets.",
        "Curl installer: Available for verified v0.4.0 release assets.",
        "Homebrew: Planned / v0.5 candidate.",
        "No downloadable v0.5 artifacts are claimed by this draft.",
    ],
    Path("docs/111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md"): [
        "# v0.5 RC Polish / Release Notes Draft",
        "This document is a draft and does not publish, tag, or release v0.5.",
        "Release binary: verified v0.4.0 release assets에 대해 Available.",
        "Curl installer: verified v0.4.0 release assets에 대해 Available.",
        "Homebrew: Planned / v0.5 candidate.",
        "이 draft는 downloadable v0.5 artifacts가 있다고 claim하지 않는다.",
    ],
    Path("docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.md"): [
        "Decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES.",
        "v0.5 is not claimed as published or released.",
        "No v0.5 artifacts are claimed uploaded.",
        "Homebrew: Planned / v0.5 candidate.",
    ],
    Path("docs/112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md"): [
        "Decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES.",
        "v0.5가 published 또는 released되었다고 claim하지 않는다.",
        "Uploaded v0.5 artifacts claim 없음.",
        "Homebrew: Planned / v0.5 candidate.",
    ],
    Path("docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.md"): [
        "Decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS.",
        "no v0.5 tag, GitHub release, asset upload, hosted checksum",
        "Output: `0.0.0-dev`.",
        "not a v0.5 release claim",
    ],
    Path("docs/113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md"): [
        "Decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS.",
        "v0.5 tag, GitHub release, asset upload, hosted checksum",
        "Output: `0.0.0-dev`.",
        "v0.5 release claim이 아니다",
    ],
    Path("docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md"): [
        "Decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES.",
        'Future manual `git tag -a v0.5.0 -m "..."`',
        "`git push origin v0.5.0`",
        "Homebrew remains Planned / v0.5 candidate.",
    ],
    Path("docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md"): [
        "Decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES.",
        'Future manual `git tag -a v0.5.0 -m "..."`',
        "`git push origin v0.5.0`",
        "Homebrew remains Planned / v0.5 candidate.",
    ],
    Path("docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.md"): [
        "Decision: HUMAN_APPROVAL_PACKET_READY_WITH_NOTES.",
        "Decision: APPROVE_PUBLICATION_PREP_ONLY.",
        "Previous decision: DO_NOT_APPROVE_FIX_FIRST.",
        "Release tag target: v0.5.0.",
    ],
    Path("docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md"): [
        "Decision: HUMAN_APPROVAL_PACKET_READY_WITH_NOTES.",
        "Decision: APPROVE_PUBLICATION_PREP_ONLY.",
        "Previous decision: DO_NOT_APPROVE_FIX_FIRST.",
        "Release tag target: v0.5.0.",
    ],
    Path("docs/117_V0_5_0_POST_RELEASE_VERIFICATION.md"): [
        "Decision: V0_5_0_POST_RELEASE_VERIFIED_WITH_NOTES.",
        "v0.5.0 publication: performed and verified in this document",
        "Release binary: verified in this document",
        "Curl installer: verified in this document",
        "Homebrew: Planned / v0.5 candidate",
        "Model workspace packs: Experimental",
        "No-terminal method: Experimental / assisted",
        "Windows execution was not run on a Windows host",
    ],
    Path("docs/117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md"): [
        "Decision: V0_5_0_POST_RELEASE_VERIFIED_WITH_NOTES.",
        "v0.5.0 publication: 이 문서에서 performed and verified",
        "Release binary: 이 문서에서 verify",
        "Curl installer: 이 문서에서 verify",
        "Homebrew: Planned / v0.5 candidate",
        "Model workspace packs: Experimental",
        "No-terminal method: Experimental / assisted",
        "Real Windows execution은 local에서 run하지 않음.",
    ],
}

for path, markers in required_markers.items():
    text = path.read_text(encoding="utf-8")
    missing = [marker for marker in markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing v0.5.0 publication markers: {missing}")

for path in required_paths:
    text = path.read_text(encoding="utf-8")
    forbidden_claims = [
        "Homebrew: Available",
        "Model workspace packs: Available",
        "No-terminal method: Available",
        "no-terminal deterministic validation passed",
        "ni run executes downstream work",
        "benchmark evidence proves implementation quality",
    ]
    for claim in forbidden_claims:
        context = text[max(0, text.find(claim) - 500): text.find(claim) + len(claim) + 160]
        allowed_negations = [
            "Do not claim",
            "Do not",
            "does not",
            "not claim",
            "forbidden",
            "forbidden claims",
            "Forbidden",
            "claim 없음",
            "claim하지",
            "금지",
            "없음",
            "downstream agent performance",
        ]
        if claim in text and not any(marker in context for marker in allowed_negations):
            raise SystemExit(f"{path} appears to make a forbidden v0.5.0 claim: {claim}")
PY

run_step "v0.5.1 post-release docs are factual" python3 - <<'PY'
from pathlib import Path

required_paths = [
    Path("docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md"),
    Path("docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.ko.md"),
    Path("docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.md"),
    Path("docs/130_V0_5_1_RELEASE_NOTES_FINALIZATION.ko.md"),
    Path("docs/131_V0_5_1_PUBLICATION_CHECKLIST.md"),
    Path("docs/131_V0_5_1_PUBLICATION_CHECKLIST.ko.md"),
    Path("docs/132_V0_5_1_POST_RELEASE_VERIFICATION.md"),
    Path("docs/132_V0_5_1_POST_RELEASE_VERIFICATION.ko.md"),
]

for path in required_paths:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

required_markers = {
    Path("docs/132_V0_5_1_POST_RELEASE_VERIFICATION.md"): [
        "V0_5_1_RELEASE_EXECUTED_WITH_NOTES",
        "v0.5.1 release: published and verified in this document.",
        "Public install parity mismatch: addressed by v0.5.1 for the tested macOS arm64 path.",
        "Homebrew: Planned / v0.5 candidate.",
        "Windows real-host execution: deferred on macOS-only development host.",
        "Model workspace packs: Experimental.",
        "No-terminal method: Experimental / assisted.",
        "Skills are UX; CLI is authority.",
        "Hosted checksums verify.",
        "`install.sh` retrieves and installs v0.5.1 in an isolated temp path.",
        "Homebrew is Available.",
        "Windows real-host execution works.",
    ],
    Path("docs/132_V0_5_1_POST_RELEASE_VERIFICATION.ko.md"): [
        "V0_5_1_RELEASE_EXECUTED_WITH_NOTES",
        "v0.5.1 release: 이 문서에서 published and verified.",
        "Public install parity mismatch: tested macOS arm64 path에서는 v0.5.1로 addressed.",
        "Homebrew: Planned / v0.5 candidate.",
        "Windows real-host execution: macOS-only development host에서는 deferred.",
        "Model workspace packs: Experimental.",
        "No-terminal method: Experimental / assisted.",
        "Skills are UX; CLI is authority.",
        "Hosted checksums verify.",
        "`install.sh` retrieves and installs v0.5.1 in an isolated temp path.",
        "Homebrew is Available.",
        "Windows real-host execution works.",
    ],
}

for path, markers in required_markers.items():
    text = path.read_text(encoding="utf-8")
    missing = [marker for marker in markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing v0.5.1 post-release markers: {missing}")

for path in required_paths:
    text = path.read_text(encoding="utf-8")
    forbidden_claims = [
        "Homebrew: Available",
        "Model workspace packs: Available",
        "No-terminal method: Available",
        "no-terminal deterministic validation passed",
        "ni run executes downstream work",
    ]
    for claim in forbidden_claims:
        context = text[max(0, text.find(claim) - 500): text.find(claim) + len(claim) + 180]
        allowed_negations = [
            "Do not",
            "does not",
            "not claim",
            "What this verification does not prove",
            "claim 없음",
            "claim하지",
            "금지",
            "없음",
        ]
        if claim in text and not any(marker in context for marker in allowed_negations):
            raise SystemExit(f"{path} appears to make a forbidden v0.5.1 claim: {claim}")
PY

run_step "v0.6.3 post-release docs are factual" python3 - <<'PY'
from pathlib import Path

required_paths = [
    Path("docs/142_V0_6_3_POST_RELEASE_VERIFICATION.md"),
    Path("docs/142_V0_6_3_POST_RELEASE_VERIFICATION.ko.md"),
]

for path in required_paths:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

required_markers = {
    Path("docs/142_V0_6_3_POST_RELEASE_VERIFICATION.md"): [
        "V0_6_3_RELEASE_EXECUTED_AND_VERIFIED",
        "d048158a91f64888a71304ee1547ff6c4bbebe0e",
        "https://github.com/Nam-Cheol/ni/releases/tag/v0.6.3",
        "27244128964",
        "2026-06-10T00:11:42Z",
        "namba-intent_0.6.3_checksums.txt",
        "namba-intent_0.6.3_darwin_arm64.tar.gz",
        "namba-intent_0.6.3_windows_amd64.zip",
        "Default `gh release view --repo Nam-Cheol/ni` returned `v0.6.3`",
        "Installed binary version | `0.6.3`",
        "Returned expected first-run `BLOCKED` status",
        "Printed guidance only; no download or install",
        "Windows real-host install execution remains deferred",
        "no task runner, SPEC runner,",
        "downstream execution layer was added",
    ],
    Path("docs/142_V0_6_3_POST_RELEASE_VERIFICATION.ko.md"): [
        "V0_6_3_RELEASE_EXECUTED_AND_VERIFIED",
        "d048158a91f64888a71304ee1547ff6c4bbebe0e",
        "https://github.com/Nam-Cheol/ni/releases/tag/v0.6.3",
        "27244128964",
        "2026-06-10T00:11:42Z",
        "namba-intent_0.6.3_checksums.txt",
        "namba-intent_0.6.3_darwin_arm64.tar.gz",
        "namba-intent_0.6.3_windows_amd64.zip",
        "Default `gh release view --repo Nam-Cheol/ni`는 `v0.6.3`를 반환",
        "Installed binary version | `0.6.3`",
        "Expected first-run `BLOCKED` status 반환",
        "Guidance만 출력; download/install 없음",
        "Windows real-host\n  install execution은 Windows transcript가 있을 때까지 deferred",
        "task runner, SPEC runner, shell adapter, Codex",
        "downstream execution layer는",
        "추가하지 않았다",
    ],
}

for path, markers in required_markers.items():
    text = path.read_text(encoding="utf-8")
    missing = [marker for marker in markers if marker not in text]
    if missing:
        raise SystemExit(f"{path} is missing v0.6.3 post-release markers: {missing}")
PY

run_step "release facts match repository resources" python3 - <<'PY'
from pathlib import Path

readme = Path("README.md").read_text(encoding="utf-8")
readme_ko = Path("README.ko.md").read_text(encoding="utf-8")
install = Path("docs/22_INSTALL.md").read_text(encoding="utf-8")
readiness = Path("docs/46_RELEASE_READINESS.md").read_text(encoding="utf-8")
readiness_ko = Path("docs/46_RELEASE_READINESS.ko.md").read_text(encoding="utf-8")

if Path("LICENSE").exists():
    texts = {
        "README.md": readme,
        "README.ko.md": readme_ko,
        "docs/22_INSTALL.md": install,
        "docs/46_RELEASE_READINESS.md": readiness,
        "docs/46_RELEASE_READINESS.ko.md": readiness_ko,
    }
    required_license_markers = {
        "README.md": ("MIT License", "[MIT License](LICENSE)"),
        "README.ko.md": ("MIT License", "[MIT License](LICENSE)"),
        "docs/22_INSTALL.md": ("MIT License", "[MIT License](../LICENSE)"),
        "docs/46_RELEASE_READINESS.md": ("MIT License", "[MIT License](../LICENSE)"),
        "docs/46_RELEASE_READINESS.ko.md": ("MIT License", "[MIT License](../LICENSE)"),
    }
    for label, markers in required_license_markers.items():
        missing = [marker for marker in markers if marker not in texts[label]]
        if missing:
            raise SystemExit(f"{label} is missing factual license markers: {missing}")
else:
    if "does not claim a license" not in readiness:
        raise SystemExit("LICENSE is absent, but release readiness does not say so")

ci_path = Path(".github/workflows/ci.yml")
if ci_path.exists():
    ci_text = ci_path.read_text(encoding="utf-8")
    required_ci = [
        "push:",
        "pull_request:",
        "go test ./...",
        "bash scripts/quality.sh",
        "bash scripts/smoke.sh",
    ]
    missing = [item for item in required_ci if item not in ci_text]
    if missing:
        raise SystemExit(f"{ci_path} is missing required CI entries: {missing}")
    for label, text in {
        "README.md": readme,
        "README.ko.md": readme_ko,
        "docs/46_RELEASE_READINESS.md": readiness,
        "docs/46_RELEASE_READINESS.ko.md": readiness_ko,
    }.items():
        if ".github/workflows/ci.yml" not in text:
            raise SystemExit(f"{label} does not document the existing CI workflow")
else:
    for label, text in {"README.md": readme, "README.ko.md": readme_ko}.items():
        if "workflows/ci" in text or "badge.svg" in text:
            raise SystemExit(f"{label} links CI resources, but no CI workflow exists")

if Path("SECURITY.md").exists():
    for label, text in {"README.md": readme, "README.ko.md": readme_ko}.items():
        if "SECURITY.md" not in text:
            raise SystemExit(f"{label} should link SECURITY.md because it exists")
else:
    required_security = [
        "SECURITY.md` does not exist",
        "security policy",
        "TODO",
    ]
    missing = [item for item in required_security if item not in readiness]
    if missing:
        raise SystemExit(f"docs/46_RELEASE_READINESS.md is missing security-status markers: {missing}")

release_claim_markers = {
    "README.md": [
        "README shows two primary first-success paths for the current tree.",
        "Homebrew: Planned / v0.5 candidate.",
        "v0.6.3 release: published and verified for `namba-intent` on macOS darwin/arm64.",
        "Primary command in current tree: `namba-intent`.",
    ],
    "README.ko.md": [
        "README는 current tree의 첫 성공을 위한 두 가지 primary path만 보여줍니다.",
        "Homebrew: Planned / v0.5 candidate.",
        "v0.6.3 release: macOS darwin/arm64에서 `namba-intent` 기준 published and verified.",
        "Primary command: `namba-intent`.",
    ],
}
for label, text in {"README.md": readme, "README.ko.md": readme_ko}.items():
    missing = [marker for marker in release_claim_markers[label] if marker not in text]
    if missing:
        raise SystemExit(f"{label} is missing release-status markers: {missing}")

release_config = Path(".goreleaser.yaml")
release_workflow = Path(".github/workflows/release.yml")
release_pipeline_docs = [
    Path("docs/67_RELEASE_PIPELINE.md"),
    Path("docs/67_RELEASE_PIPELINE.ko.md"),
]
if not release_config.exists():
    raise SystemExit(".goreleaser.yaml is missing")
if not release_workflow.exists():
    raise SystemExit(".github/workflows/release.yml is missing")
for path in release_pipeline_docs:
    if not path.exists():
        raise SystemExit(f"{path} is missing")

config = release_config.read_text(encoding="utf-8")
workflow = release_workflow.read_text(encoding="utf-8")

required_config = [
    "version: 2",
    "project_name: namba-intent",
    "main: ./cmd/namba-intent",
    "binary: namba-intent",
    "CGO_ENABLED=0",
    "linux",
    "darwin",
    "windows",
    "amd64",
    "arm64",
    "goos: windows",
    "goarch: arm64",
    "formats:",
    "tar.gz",
    "zip",
    "checksum:",
    "ni/internal/version.Version={{ .Version }}",
]
missing = [item for item in required_config if item not in config]
if missing:
    raise SystemExit(f".goreleaser.yaml is missing release config markers: {missing}")

required_workflow = [
    "name: Release",
    "tags:",
    '"v*"',
    "contents: write",
    "actions/checkout@v4",
    "actions/setup-go@v5",
    "go test ./...",
    "bash scripts/quality.sh",
    "bash scripts/release-check.sh",
    "goreleaser/goreleaser-action@v6",
    "release --clean",
]
missing = [item for item in required_workflow if item not in workflow]
if missing:
    raise SystemExit(f".github/workflows/release.yml is missing workflow markers: {missing}")

for forbidden_trigger in ["pull_request:", "workflow_dispatch:", "schedule:", "branches:"]:
    if forbidden_trigger in workflow:
        raise SystemExit(
            f".github/workflows/release.yml must run only on v* tags, "
            f"but found {forbidden_trigger}"
        )

pipeline_required = [
    "bash scripts/release-dry-run.sh",
    "go test ./...",
    "bash scripts/quality.sh",
    "bash scripts/smoke.sh",
    "bash scripts/demo-check.sh",
    "bash scripts/install-check.sh",
    "bash scripts/release-check.sh",
    "goreleaser check",
    "goreleaser release --snapshot --clean",
    "GoReleaser Archive Matrix",
    "namba-intent_<version>_linux_amd64.tar.gz",
    "namba-intent_<version>_linux_arm64.tar.gz",
    "namba-intent_<version>_darwin_amd64.tar.gz",
    "namba-intent_<version>_darwin_arm64.tar.gz",
    "namba-intent_<version>_windows_amd64.zip",
    "namba-intent_<version>_checksums.txt",
]

availability_guards = [
    ("Release binary availability",),
    ("Curl installer availability",),
    ("Homebrew availability",),
    ("not available", "available하지"),
]

for path in release_pipeline_docs:
    text = path.read_text(encoding="utf-8")
    missing = [item for item in pipeline_required if item not in text]
    if missing:
        raise SystemExit(f"{path} is missing release pipeline markers: {missing}")
    missing_guards = [
        " or ".join(options)
        for options in availability_guards
        if not any(option in text for option in options)
    ]
    if missing_guards:
        raise SystemExit(f"{path} is missing availability guard markers: {missing_guards}")
PY

run_step "examples and benchmark protocol exist" python3 - <<'PY'
from pathlib import Path

examples = Path("examples")
if not examples.is_dir():
    raise SystemExit("examples/ is missing")

required_examples = [
    examples / "ambiguous-prompt-blocked" / "README.md",
    examples / "conversation-product" / "README.md",
    examples / "research-protocol" / "README.md",
    examples / "benchmark-report" / "README.md",
]
missing = [str(path) for path in required_examples if not path.exists()]
if missing:
    raise SystemExit(f"required examples are missing: {missing}")

benchmark = Path("docs/43_BENCHMARK_PROTOCOL.md")
text = benchmark.read_text(encoding="utf-8")
required = [
    "It is not an execution benchmark",
    "must not execute downstream agents",
    "Target prompt boundedness",
]
missing = [item for item in required if item not in text]
if missing:
    raise SystemExit(f"{benchmark} is missing benchmark protocol markers: {missing}")
PY

run_step "schemas validate" python3 scripts/check-schema.py
run_step "core boundary has no violations" python3 scripts/check-core-boundary.py --self-test
run_step "Go tests pass" go test ./...
run_step "golden tests pass" go test ./internal/cli -run Golden -count=1
run_step "smoke passes" bash scripts/smoke.sh
run_step "public demos verify" bash scripts/demo-check.sh
run_step "install and build paths pass" bash scripts/install-check.sh

run_step "status proof works" bash -c '
  go run ./cmd/namba-intent status --dir examples/conversation-product --proof >"$1/status-proof.out"
  require_output "NI Intent Readiness: READY" "$1/status-proof.out"
  require_output "Passed checks:" "$1/status-proof.out"
  require_output "Docs and contract are synchronized." "$1/status-proof.out"
' bash "$QUICKSTART_TMP"

run_step "README quickstart works in go run mode" bash -c '
  go run ./cmd/namba-intent init --dir "$1/plan" --profile prototype >"$1/init.out"
  require_output "initialized Namba Intent planning workspace" "$1/init.out"
  go run ./cmd/namba-intent status --dir "$1/plan" >"$1/status.out"
  require_output "BLOCKED" "$1/status.out"
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
    "아니다",
    "안 된다",
    "않고",
    "않는다",
    "없음",
    "실행하지",
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
        context = "\n".join(lines[start : min(len(lines), index + 2)]).lower()
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
    r"(?:go run \./cmd/namba-intent|(?:\./bin/namba-intent|~/.local/bin/namba-intent)|`namba-intent)\s+([a-z][a-z-]*)",
    readme,
):
    readme_commands.add(match.group(1))

smoke_commands = set(re.findall(r'run_cmd\s+"namba-intent\s+([a-z][a-z-]*)', smoke))

missing = sorted(readme_commands - smoke_commands)
if missing:
    raise SystemExit(
        "README lists public commands without smoke coverage: " + ", ".join(missing)
    )
PY

echo "release-check: release readiness gate passed"
