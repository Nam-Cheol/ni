#!/usr/bin/env python3
"""Static safety checks for the Windows installer."""

from __future__ import annotations

from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
INSTALLER = ROOT / "install.ps1"


def fail(message: str) -> None:
    raise SystemExit(message)


def main() -> None:
    if not INSTALLER.exists():
        fail("install.ps1 is missing")

    text = INSTALLER.read_text(encoding="utf-8")

    required = [
        '$env:LOCALAPPDATA',
        '[Environment]::GetEnvironmentVariable("Path", "User")',
        '[Environment]::SetEnvironmentVariable("Path", ($Entries -join ";"), "User")',
        "Add-UserPathEntry",
        "Remove-UserPathEntry",
        "Remove-Item $Target -Force",
        "Remove-Item $InstallDir -Force",
        "Open a new PowerShell session",
        "ni --help",
        "ni version",
        "does not install model skills or run downstream work",
    ]
    missing = [marker for marker in required if marker not in text]
    if missing:
        fail(f"install.ps1 is missing required safety markers: {missing}")

    forbidden = [
        "setx",
        '"Machine"',
        "'Machine'",
        "System PATH",
        "Start-Process -Verb RunAs",
        "-Scope Machine",
    ]
    present = [marker for marker in forbidden if marker in text]
    if present:
        fail(f"install.ps1 contains forbidden Windows install behavior: {present}")

    print("Windows installer static safety checks passed")


if __name__ == "__main__":
    main()
