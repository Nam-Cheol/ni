#!/usr/bin/env python3
"""Check install/distribution documentation availability claims."""

from __future__ import annotations

from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
ALLOWED_STATUSES = {"Available", "Experimental", "Release-gated", "Planned"}

README_EXPECTED = {
    "Source": "Available",
    "Local binary": "Available",
    "Model workspaces": "Experimental",
    "No-terminal method": "Experimental",
    "Release binary": "Available",
    "Curl installer": "Available",
    "Homebrew": "Planned",
}

DISTRIBUTION_EXPECTED = {
    "Source mode": "Available",
    "Local binary mode": "Available",
    "Release binary mode": "Available",
    "Curl installer mode": "Available",
    "Package manager mode": "Planned",
    "Model workspace mode": "Experimental",
    "No-terminal mode": "Experimental",
}

INSTALL_EXPECTED = README_EXPECTED

REQUIRED_MARKERS = {
    "README.md": [
        "README shows two primary first-success paths.",
        "### macOS",
        "### Windows",
        "ni init .",
        "Homebrew: Planned / v0.5 candidate",
        "Real-host Windows execution remains deferred",
    ],
    "README.ko.md": [
        "README는 첫 성공을 위한 두 가지 primary path만 보여줍니다.",
        "### macOS",
        "### Windows",
        "ni init .",
        "Homebrew: Planned / v0.5 candidate",
        "실제 Windows host execution은",
    ],
    "docs/22_INSTALL.md": [
        "README intentionally shows only two primary first-success paths:",
        "Every public install path has exactly one status:",
        "Release binary status: Available.",
        "Curl installer status: Available for verified v0.5.1 release assets.",
        "Package manager status: Planned.",
        "curl installer availability only for the verified v0.5.1 installer path",
    ],
    "docs/install-curl.md": [
        "Status: Available for the verified v0.5.1 GitHub Release assets.",
        "The v0.5.1 verification passed on 2026-06-08.",
        "Open a new shell after installation, then check the global command",
        "BINDIR=\"$HOME/.local/bin\" sh install.sh --uninstall",
    ],
    "docs/install-curl.ko.md": [
        "Status: verified v0.5.1 GitHub Release assets에 대해 Available이다.",
        "v0.5.1 verification은 2026-06-08에 통과했다.",
        "설치 후 새 shell을 열고 global command를 help 또는 version command로 확인한다",
        "BINDIR=\"$HOME/.local/bin\" sh install.sh --uninstall",
    ],
    "docs/120_GLOBAL_INSTALL_ACCEPTANCE.md": [
        "A macOS install is successful only when all of these are true:",
        "A Windows install is successful only when all of these are true:",
        "Do not claim Windows execution verified until a Windows PowerShell install",
        "Skills are UX; CLI is authority.",
    ],
    "docs/120_GLOBAL_INSTALL_ACCEPTANCE.ko.md": [
        "macOS install은 다음이 모두 true일 때만 성공이다:",
        "Windows install은 다음이 모두 true일 때만 성공이다:",
        "Windows execution verified라고 claim하지 않는다",
        "Skills are UX; CLI is authority.",
    ],
    "docs/54_HOMEBREW_DISTRIBUTION.md": [
        "Current status: Planned.",
        "There is no published Homebrew formula",
    ],
    "docs/54_HOMEBREW_DISTRIBUTION.ko.md": [
        "현재 상태: Planned.",
        "published Homebrew formula는 없고",
    ],
    "docs/70_RELEASE_VERIFICATION_v0.3.0.md": [
        "This release-asset verification alone does not",
        "v0.3.0 Curl Installer Verification",
        "Available for the verified v0.3.0 release assets",
        "Homebrew, Scoop, and package-manager distribution remain Planned or unavailable",
    ],
    "docs/70_RELEASE_VERIFICATION_v0.3.0.ko.md": [
        "이 release-asset verification만으로는 curl 설치",
        "v0.3.0 Curl Installer Verification",
        "verified v0.3.0 release assets에 대해 Available",
        "Homebrew, Scoop, 패키지 매니저 배포는 별도 publish와 verification 전까지",
    ],
    "docs/71_CURL_INSTALLER_VERIFICATION_v0.3.0.md": [
        "Curl installer status: Available for the verified v0.3.0 release assets.",
        "Homebrew status: Planned.",
        "No package-manager availability is claimed by this verification.",
    ],
    "docs/71_CURL_INSTALLER_VERIFICATION_v0.3.0.ko.md": [
        "Curl installer status: verified v0.3.0 release assets에 대해 Available.",
        "Homebrew status: Planned.",
        "이 검증은 package-manager availability를 claim하지 않는다.",
    ],
    "docs/87_CURL_INSTALLER_VERIFICATION_v0.4.0.md": [
        "Curl installer status: Available for the verified v0.4.0 release assets.",
        "Homebrew status: Planned.",
        "No package-manager availability is claimed by this verification.",
    ],
    "docs/87_CURL_INSTALLER_VERIFICATION_v0.4.0.ko.md": [
        "Curl installer status: verified v0.4.0 release assets에 대해 Available.",
        "Homebrew status: Planned.",
        "이 검증은 package-manager availability를 claim하지 않는다.",
    ],
    "docs/72_HOMEBREW_TAP_PLAN.md": [
        "Current Homebrew status: Planned.",
        "There is no owner-confirmed tap",
        "Homebrew remains Planned until all of these are true:",
    ],
    "docs/72_HOMEBREW_TAP_PLAN.ko.md": [
        "Current Homebrew status: Planned.",
        "Owner-confirmed tap",
        "Homebrew는 다음이 모두 true가 될 때까지 Planned로 남는다:",
    ],
    "docs/50_LAUNCH_CHECKLIST.md": [
        "only claims curl installer availability after verification",
        "verified release binary, and",
        "verified curl installer availability",
        "curl installer path that has been verified against those assets",
    ],
    "docs/50_LAUNCH_CHECKLIST.ko.md": [
        "curl installer availability는 verification 이후에만 claim",
        "verified release binary, verified curl",
        "installer availability만 claim",
        "verified된 curl installer path만 claim",
    ],
    "docs/51_POST_RELEASE_ROADMAP.md": [
        "after the verified v0.5.1 release",
        "117_V0_5_0_POST_RELEASE_VERIFICATION.md",
        "132_V0_5_1_POST_RELEASE_VERIFICATION.md",
    ],
    "docs/51_POST_RELEASE_ROADMAP.ko.md": [
        "verified v0.5.1 release",
        "117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md",
        "132_V0_5_1_POST_RELEASE_VERIFICATION.ko.md",
    ],
    "docs/53_DISTRIBUTION_STRATEGY.md": [
        "Model workspace packs | Experimental",
        "host-level/global install and provider behavior remain unverified unless documented",
        "Model Workspace Status",
    ],
    "docs/53_DISTRIBUTION_STRATEGY.ko.md": [
        "Model workspace packs | Experimental",
        "host-level/global install과 provider behavior는 documented되기 전까지 unverified",
        "Model Workspace Status",
    ],
    "docs/99_MODEL_WORKSPACE_STATUS.md": [
        "Model workspace packs are **Experimental** as a broad product path.",
        "Global Claude install | not_verified",
        "Global Codex install | not_verified",
        "Provider runtime behavior | not_verified",
        "Skills are UX; CLI is authority.",
    ],
    "docs/99_MODEL_WORKSPACE_STATUS.ko.md": [
        "Model workspace packs는 broad product path로 **Experimental**이다.",
        "Global Claude install | not_verified",
        "Global Codex install | not_verified",
        "Provider runtime behavior | not_verified",
        "Skills are UX; CLI is authority.",
    ],
    "docs/68_RELEASE_NOTES_v0.3.0.md": [
        "Tag suggestion: `v0.3.0`",
        "do not publish a release",
        "Source-first usage",
        "Release binary pipeline configuration for future GitHub Release assets",
    ],
    "docs/68_RELEASE_NOTES_v0.3.0.ko.md": [
        "Tag suggestion: `v0.3.0`",
        "Release를 publish하거나",
        "Source-first usage",
        "Release binary pipeline configuration for future GitHub Release assets",
    ],
}

FORBIDDEN_AFFIRMATIVE = [
    "Status: available. `install.sh`",
    "| Homebrew | Available |",
    "| Model workspaces | Available |",
    "| Model workspace packs | Available |",
    "| Model workspace mode | Available |",
    "| Package manager mode | Available |",
    "| Homebrew | 사용 가능 |",
    "| Package manager mode | 사용 가능 |",
    "Homebrew is available",
    "Package manager mode is available",
    "model workspace packs are available globally",
    "global Codex install is verified",
    "global Claude install is verified",
    "works in all Codex",
    "works in all Claude",
    "The curl installer, Homebrew, Scoop, and package-manager distribution remain\nnot available until separately verified.",
    "curl 설치 스크립트, Homebrew, Scoop, 패키지 매니저 배포는 별도 검증 전까지\nAvailable 상태가 아니다.",
]


def fail(message: str) -> None:
    raise SystemExit(message)


def read(path: str) -> str:
    target = ROOT / path
    if not target.exists():
        fail(f"{path} is missing")
    return target.read_text(encoding="utf-8")


def table_rows(text: str) -> dict[str, str]:
    rows: dict[str, str] = {}
    for line in text.splitlines():
        stripped = line.strip()
        if not stripped.startswith("|") or "---" in stripped:
            continue
        cells = [cell.strip() for cell in stripped.strip("|").split("|")]
        if len(cells) < 2 or cells[0] in {"Path", "Track"}:
            continue
        rows[cells[0]] = cells[1]
    return rows


def check_expected_rows(path: str, expected: dict[str, str]) -> None:
    rows = table_rows(read(path))
    for label, status in expected.items():
        actual = rows.get(label)
        if actual is None:
            fail(f"{path} is missing install/distribution row: {label}")
        if actual not in ALLOWED_STATUSES:
            fail(f"{path} row {label!r} has unsupported status: {actual}")
        if actual != status:
            fail(f"{path} row {label!r} has status {actual!r}, expected {status!r}")


def check_required_markers() -> None:
    for path, markers in REQUIRED_MARKERS.items():
        text = read(path)
        normalized = " ".join(text.split())
        missing = [marker for marker in markers if marker not in text and marker not in normalized]
        if missing:
            fail(f"{path} is missing install-doc markers: {missing}")


def check_forbidden_claims() -> None:
    paths = [
        "README.md",
        "README.ko.md",
        "docs/22_INSTALL.md",
        "docs/install-curl.md",
        "docs/install-curl.ko.md",
        "docs/53_DISTRIBUTION_STRATEGY.md",
        "docs/53_DISTRIBUTION_STRATEGY.ko.md",
        "docs/99_MODEL_WORKSPACE_STATUS.md",
        "docs/99_MODEL_WORKSPACE_STATUS.ko.md",
        "docs/54_HOMEBREW_DISTRIBUTION.md",
        "docs/54_HOMEBREW_DISTRIBUTION.ko.md",
        "docs/70_RELEASE_VERIFICATION_v0.3.0.md",
        "docs/70_RELEASE_VERIFICATION_v0.3.0.ko.md",
        "docs/71_CURL_INSTALLER_VERIFICATION_v0.3.0.md",
        "docs/71_CURL_INSTALLER_VERIFICATION_v0.3.0.ko.md",
        "docs/87_CURL_INSTALLER_VERIFICATION_v0.4.0.md",
        "docs/87_CURL_INSTALLER_VERIFICATION_v0.4.0.ko.md",
        "docs/72_HOMEBREW_TAP_PLAN.md",
        "docs/72_HOMEBREW_TAP_PLAN.ko.md",
        "docs/50_LAUNCH_CHECKLIST.md",
        "docs/50_LAUNCH_CHECKLIST.ko.md",
        "docs/51_POST_RELEASE_ROADMAP.md",
        "docs/51_POST_RELEASE_ROADMAP.ko.md",
        "docs/68_RELEASE_NOTES_v0.3.0.md",
        "docs/68_RELEASE_NOTES_v0.3.0.ko.md",
    ]
    for path in paths:
        text = read(path)
        for phrase in FORBIDDEN_AFFIRMATIVE:
            if phrase in text:
                fail(f"{path} contains forbidden affirmative install claim: {phrase}")


def main() -> None:
    check_expected_rows("docs/22_INSTALL.md", INSTALL_EXPECTED)
    for path in ["docs/53_DISTRIBUTION_STRATEGY.md", "docs/53_DISTRIBUTION_STRATEGY.ko.md"]:
        check_expected_rows(path, DISTRIBUTION_EXPECTED)
    check_required_markers()
    check_forbidden_claims()
    print("install docs checks passed")


if __name__ == "__main__":
    main()
