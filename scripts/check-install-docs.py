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

REQUIRED_MARKERS = {
    "README.md": [
        "v0.3.0 release binaries are available",
        "The curl installer is available after verification against the",
        "including Homebrew, is not available yet",
    ],
    "README.ko.md": [
        "v0.3.0 release binaries는 asset과 checksum 검증 후 Available입니다",
        "Curl installer는 실제 v0.3.0 release assets에 대해 검증된 뒤 Available입니다",
        "Homebrew를 포함한 package-manager distribution은 아직 Available이 아닙니다",
    ],
    "docs/22_INSTALL.md": [
        "Release binary status: Available.",
        "Curl installer status: Available for verified v0.3.0 release assets.",
        "Package manager status: Planned.",
        "curl installer availability only for the verified v0.3.0 installer path",
    ],
    "docs/install-curl.md": [
        "Status: Available for the verified v0.3.0 GitHub Release assets.",
        "The v0.3.0 verification passed on 2026-05-28.",
    ],
    "docs/install-curl.ko.md": [
        "Status: verified v0.3.0 GitHub Release assets에 대해 Available이다.",
        "v0.3.0 verification은 2026-05-28에 통과했다.",
    ],
    "docs/54_HOMEBREW_DISTRIBUTION.md": [
        "Current status: Planned.",
        "There is no published Homebrew formula",
    ],
    "docs/54_HOMEBREW_DISTRIBUTION.ko.md": [
        "현재 상태: Planned.",
        "published Homebrew formula는 없고",
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
    "| Package manager mode | Available |",
    "Homebrew is available",
    "Package manager mode is available",
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
        "docs/68_RELEASE_NOTES_v0.3.0.md",
        "docs/68_RELEASE_NOTES_v0.3.0.ko.md",
    ]
    for path in paths:
        text = read(path)
        for phrase in FORBIDDEN_AFFIRMATIVE:
            if phrase in text:
                fail(f"{path} contains forbidden affirmative install claim: {phrase}")


def main() -> None:
    for path in ["README.md", "README.ko.md"]:
        check_expected_rows(path, README_EXPECTED)
    for path in ["docs/53_DISTRIBUTION_STRATEGY.md", "docs/53_DISTRIBUTION_STRATEGY.ko.md"]:
        check_expected_rows(path, DISTRIBUTION_EXPECTED)
    check_required_markers()
    check_forbidden_claims()
    print("install docs checks passed")


if __name__ == "__main__":
    main()
