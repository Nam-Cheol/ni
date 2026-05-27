#!/usr/bin/env python3
"""Validate repository SVG assets and generated asset drift."""

from __future__ import annotations

import re
import subprocess
import sys
import tempfile
import xml.etree.ElementTree as ET
from html.parser import HTMLParser
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
ASSETS = ROOT / "assets"
MAX_SVG_BYTES = 30_000
README_FILES = [
    ROOT / "README.md",
    ROOT / "README.ko.md",
]
REQUIRED_ASSETS = [
    ASSETS / "hero.svg",
    ASSETS / "badge-english.svg",
    ASSETS / "badge-korean.svg",
    ASSETS / "card-pain-vague-intent.svg",
    ASSETS / "card-pain-early-execution.svg",
    ASSETS / "card-pain-rework.svg",
    ASSETS / "card-payoff-capture-intent.svg",
    ASSETS / "card-payoff-lock-contract.svg",
    ASSETS / "card-payoff-handoff-safely.svg",
    ASSETS / "card-start.svg",
    ASSETS / "card-contract.svg",
    ASSETS / "card-handoff.svg",
]
GENERATED_NAMES = [path.name for path in REQUIRED_ASSETS]
EMOJI_RE = re.compile(
    "["
    "\U0001F1E6-\U0001F1FF"
    "\U0001F300-\U0001FAFF"
    "\u2600-\u27BF"
    "\ufe0f"
    "]"
)
REMOTE_RESOURCE_RE = re.compile(
    r"(?:url\(\s*['\"]?https?://|(?:href|src)\s*=\s*['\"]https?://)",
    re.IGNORECASE,
)


class LocalAssetParser(HTMLParser):
    def __init__(self) -> None:
        super().__init__()
        self.references: list[str] = []

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        for name, value in attrs:
            if name.lower() not in {"src", "href"} or value is None:
                continue
            if value.startswith("assets/"):
                self.references.append(value)


def fail(message: str) -> None:
    raise SystemExit(message)


def local_name(tag: str) -> str:
    return tag.rsplit("}", 1)[-1]


def parse_svg(path: Path) -> ET.Element:
    try:
        tree = ET.parse(path)
    except ET.ParseError as error:
        fail(f"{path.relative_to(ROOT)} is not parseable XML: {error}")
    return tree.getroot()


def has_external_href(root: ET.Element) -> bool:
    for element in root.iter():
        for key, value in element.attrib.items():
            name = local_name(key).lower()
            lowered = value.lower().strip()
            if name == "href" and lowered.startswith(("http://", "https://")):
                return True
            if "url(http://" in lowered or "url(https://" in lowered:
                return True
    return False


def check_svg(path: Path) -> None:
    text = path.read_text(encoding="utf-8")
    if path.stat().st_size > MAX_SVG_BYTES:
        fail(f"{path.relative_to(ROOT)} exceeds {MAX_SVG_BYTES} bytes")
    if REMOTE_RESOURCE_RE.search(text):
        fail(f"{path.relative_to(ROOT)} contains a remote HTTP resource reference")
    if EMOJI_RE.search(text):
        fail(f"{path.relative_to(ROOT)} contains emoji or emoji-style codepoints")

    root = parse_svg(path)
    if local_name(root.tag) != "svg":
        fail(f"{path.relative_to(ROOT)} root element is not svg")
    if not root.attrib.get("viewBox"):
        fail(f"{path.relative_to(ROOT)} is missing viewBox")
    if not (root.attrib.get("width") and root.attrib.get("height")):
        fail(f"{path.relative_to(ROOT)} is missing width/height sizing")
    if any(local_name(element.tag) == "foreignObject" for element in root.iter()):
        fail(f"{path.relative_to(ROOT)} contains foreignObject")
    if has_external_href(root):
        fail(f"{path.relative_to(ROOT)} contains an external href or url reference")


def check_required_assets() -> None:
    missing = [path.relative_to(ROOT) for path in REQUIRED_ASSETS if not path.exists()]
    if missing:
        fail("missing required assets: " + ", ".join(str(path) for path in missing))


def check_readme_asset_references() -> None:
    for readme in README_FILES:
        if not readme.exists():
            continue
        parser = LocalAssetParser()
        parser.feed(readme.read_text(encoding="utf-8"))
        missing = []
        for reference in parser.references:
            target = ROOT / reference
            if not target.exists():
                missing.append(reference)
        if missing:
            fail(
                f"{readme.relative_to(ROOT)} references missing local assets: "
                + ", ".join(sorted(set(missing)))
            )


def check_generated_assets_current() -> None:
    with tempfile.TemporaryDirectory() as tmp:
        tmp_path = Path(tmp)
        result = subprocess.run(
            [sys.executable, "scripts/render-assets.py", "--out-dir", str(tmp_path)],
            cwd=ROOT,
            text=True,
            capture_output=True,
        )
        if result.returncode != 0:
            details = result.stderr.strip() or result.stdout.strip()
            fail(f"asset render check failed: {details}")

        stale = []
        for name in GENERATED_NAMES:
            expected = (tmp_path / name).read_text(encoding="utf-8")
            actual = (ASSETS / name).read_text(encoding="utf-8")
            if actual != expected:
                stale.append(f"assets/{name}")
        if stale:
            fail("generated assets are stale; run python3 scripts/render-assets.py: " + ", ".join(stale))


def main() -> None:
    check_required_assets()
    check_readme_asset_references()
    for path in sorted(ASSETS.glob("*.svg")):
        check_svg(path)
    check_generated_assets_current()
    print("asset checks passed")


if __name__ == "__main__":
    main()
