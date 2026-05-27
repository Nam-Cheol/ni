#!/usr/bin/env python3
"""Lightweight README surface regression checks."""

from __future__ import annotations

import re
import subprocess
import sys
from html.parser import HTMLParser
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parents[1]
README_FILES = ["README.md", "README.ko.md"]
DUPLICATE_READMES = ["README 2.md", "README.ko 2.md"]
TOP_LINK_LINE_LIMIT = 20
INSTALL_SCRIPT = "install.sh"
YAML_FILES = [".github/workflows/release.yml", ".goreleaser.yaml"]
TABLE_SEPARATOR_RE = re.compile(r"\s*:?-{3,}:?\s*")
MARKDOWN_LINK_RE = re.compile(r"(?<!!)\[[^\]]+\]\(([^)]+)\)")
REMOTE_LINK_RE = re.compile(r"^[a-z][a-z0-9+.-]*:", re.IGNORECASE)
SPECIFIC_RUNTIME_PRODUCTS = ["Codex", "Claude", "Hyper Run", "Ouroboros", "Namba AI", "SPEC Kit"]
TOP_SECTIONS = {
    "README.md": [
        ("hero image", "assets/hero.svg"),
        ("language chips", "assets/badge-english.svg"),
        ("factual trust badges", "img.shields.io/badge/license-MIT"),
        ("slogan", "<h1 align=\"center\">"),
        ("one-line product description", "<p align=\"center\"><strong>"),
        ("why ni", "## Why ni"),
        ("start in 60 seconds", "## Start in 60 seconds"),
    ],
    "README.ko.md": [
        ("hero image", "assets/hero.svg"),
        ("language chips", "assets/badge-english.svg"),
        ("factual trust badges", "img.shields.io/badge/license-MIT"),
        ("slogan", "<h1 align=\"center\">"),
        ("one-line product description", "<p align=\"center\"><strong>"),
        ("why ni", "## 왜 ni인가"),
        ("start in 60 seconds", "## 60초 시작"),
    ],
}


def fail(message: str) -> None:
    raise SystemExit(message)


def read_lines(path: str) -> list[str]:
    file_path = ROOT / path
    if not file_path.exists():
        fail(f"{path} is missing")
    return file_path.read_text(encoding="utf-8").splitlines()


class ReadmeHTMLParser(HTMLParser):
    def __init__(self) -> None:
        super().__init__()
        self.links: list[str] = []
        self.images: list[str] = []

    def handle_starttag(self, tag: str, attrs: list[tuple[str, str | None]]) -> None:
        for name, value in attrs:
            if value is None:
                continue
            lowered = name.lower()
            if lowered == "href":
                self.links.append(value)
            if lowered == "src":
                self.images.append(value)


def check_required_readmes() -> None:
    missing = [path for path in README_FILES if not (ROOT / path).exists()]
    if missing:
        fail("missing required README files: " + ", ".join(missing))


def check_duplicate_readmes() -> None:
    present = [path for path in DUPLICATE_READMES if (ROOT / path).exists()]
    if present:
        fail("unexpected duplicate root README files: " + ", ".join(present))


def check_top_language_links() -> None:
    checks = [
        ("README.md", "README.ko.md"),
        ("README.ko.md", "README.md"),
    ]
    for path, target in checks:
        top = "\n".join(read_lines(path)[:TOP_LINK_LINE_LIMIT])
        if target not in top:
            fail(f"{path} must link to {target} in the first {TOP_LINK_LINE_LIMIT} lines")


def find_line(lines: list[str], marker: str) -> int:
    for index, line in enumerate(lines, start=1):
        if marker in line:
            return index
    return 0


def check_top_section_order() -> None:
    for path, markers in TOP_SECTIONS.items():
        lines = read_lines(path)
        last_line = 0
        for label, marker in markers:
            line_number = find_line(lines, marker)
            if line_number == 0:
                fail(f"{path} is missing top-section marker for {label}: {marker}")
            if line_number <= last_line:
                fail(f"{path} top section renders out of order at {label}")
            last_line = line_number

        slogan_line = find_line(lines, "<h1 align=\"center\">")
        for line_number, line in enumerate(lines[: slogan_line - 1], start=1):
            if "Trust signals:" in line or "신뢰 신호:" in line:
                fail(f"{path}:{line_number} has visible trust-signal text before the slogan")


def is_table_separator(line: str) -> bool:
    parts = line.strip().strip("|").split("|")
    return bool(parts) and all(TABLE_SEPARATOR_RE.fullmatch(part) for part in parts)


def table_cell_count(line: str) -> int:
    stripped = line.strip()
    if stripped.startswith("|"):
        stripped = stripped[1:]
    if stripped.endswith("|"):
        stripped = stripped[:-1]
    return len(stripped.split("|"))


def iter_markdown_tables(lines: list[str]) -> list[tuple[int, list[str]]]:
    tables: list[tuple[int, list[str]]] = []
    in_fence = False
    index = 0
    while index < len(lines):
        line = lines[index]
        if line.strip().startswith("```"):
            in_fence = not in_fence
            index += 1
            continue
        if in_fence or not line.strip().startswith("|"):
            index += 1
            continue

        start = index
        table: list[str] = []
        while index < len(lines) and lines[index].strip().startswith("|"):
            table.append(lines[index])
            index += 1
        tables.append((start + 1, table))

    return tables


def check_markdown_tables() -> None:
    for path in README_FILES:
        lines = read_lines(path)
        for line_number, table in iter_markdown_tables(lines):
            if len(table) < 2:
                fail(f"{path}:{line_number} Markdown table is missing a separator row")
            if not is_table_separator(table[1]):
                fail(f"{path}:{line_number + 1} Markdown table separator row is invalid")
            expected_cells = table_cell_count(table[0])
            if table_cell_count(table[1]) != expected_cells:
                fail(f"{path}:{line_number + 1} Markdown table separator column count differs from header")
            for offset, row in enumerate(table[2:], start=2):
                if table_cell_count(row) != expected_cells:
                    fail(f"{path}:{line_number + offset} Markdown table row column count differs from header")


def check_code_blocks() -> None:
    for path in README_FILES:
        lines = read_lines(path)
        in_fence = False
        start = 0
        for line_number, line in enumerate(lines, start=1):
            stripped = line.strip()
            if not stripped.startswith("```"):
                continue
            if not in_fence:
                language = stripped[3:].strip()
                if not language:
                    fail(f"{path}:{line_number} Markdown code fence is missing an info string")
                in_fence = True
                start = line_number
            else:
                in_fence = False
        if in_fence:
            fail(f"{path}:{start} Markdown code fence is not closed")


def normalize_local_target(target: str) -> str | None:
    target = target.strip()
    if not target or target.startswith("#") or REMOTE_LINK_RE.match(target):
        return None
    if "#" in target:
        target = target.split("#", 1)[0]
    if "?" in target:
        target = target.split("?", 1)[0]
    return target or None


def check_local_references() -> None:
    for path in README_FILES:
        text = (ROOT / path).read_text(encoding="utf-8")
        parser = ReadmeHTMLParser()
        parser.feed(text)

        image_targets = [normalize_local_target(src) for src in parser.images]
        missing_images = [
            target for target in image_targets if target is not None and not (ROOT / target).exists()
        ]
        if missing_images:
            fail(f"{path} references missing local assets: " + ", ".join(sorted(set(missing_images))))

        markdown_targets = [match.group(1).split()[0] for match in MARKDOWN_LINK_RE.finditer(text)]
        link_targets = parser.links + markdown_targets
        missing_links = []
        for target in link_targets:
            normalized = normalize_local_target(target)
            if normalized is None:
                continue
            if not (ROOT / normalized).exists():
                missing_links.append(normalized)
        if missing_links:
            fail(f"{path} references missing local files: " + ", ".join(sorted(set(missing_links))))


def section_text(lines: list[str], start_marker: str, end_marker: str | None = None) -> str:
    start = find_line(lines, start_marker)
    if start == 0:
        return ""
    end = len(lines) + 1
    if end_marker is not None:
        found = find_line(lines, end_marker)
        if found != 0:
            end = found
    return "\n".join(lines[start - 1 : end - 1])


def check_product_claim_surface() -> None:
    for path in README_FILES:
        lines = read_lines(path)
        hero_sales = "\n".join(lines[: find_line(lines, "## Start in 60 seconds") or len(lines)])
        if path == "README.ko.md":
            hero_sales = "\n".join(lines[: find_line(lines, "## 60초 시작") or len(lines)])
        for product in SPECIFIC_RUNTIME_PRODUCTS:
            if product in hero_sales:
                fail(f"{path} mentions {product} in the hero or sales pitch")

        model_path = section_text(lines, "## Choose your path", "## Demo")
        full_text = "\n".join(lines)
        outside_model_path = full_text.replace(model_path, "")
        for product in ["Codex", "Claude"]:
            if product in outside_model_path:
                fail(f"{path} mentions {product} outside the model-workspace usage path")


def check_install_shell_syntax() -> None:
    script = ROOT / INSTALL_SCRIPT
    if not script.exists():
        return
    result = subprocess.run(["sh", "-n", INSTALL_SCRIPT], cwd=ROOT, text=True, capture_output=True)
    if result.returncode != 0:
        details = result.stderr.strip() or result.stdout.strip()
        fail(f"{INSTALL_SCRIPT} shell syntax check failed: {details}")


def check_yaml_with_pyyaml(paths: list[str]) -> bool:
    try:
        import yaml  # type: ignore[import-not-found]
    except ImportError:
        return False

    for path in paths:
        file_path = ROOT / path
        if not file_path.exists():
            continue
        try:
            loaded: Any = yaml.safe_load(file_path.read_text(encoding="utf-8"))
        except yaml.YAMLError as error:
            fail(f"{path} YAML parse failed: {error}")
        if not isinstance(loaded, dict):
            fail(f"{path} must parse to a YAML mapping")
    return True


def simple_yaml_check(path: str) -> None:
    file_path = ROOT / path
    if not file_path.exists():
        return
    lines = file_path.read_text(encoding="utf-8").splitlines()
    if not lines:
        fail(f"{path} is empty")
    if any("\t" in line for line in lines):
        fail(f"{path} contains tab indentation")
    meaningful = [
        line
        for line in lines
        if line.strip() and not line.lstrip().startswith("#") and not line.startswith("---")
    ]
    if not meaningful:
        fail(f"{path} has no YAML content")
    if not any(":" in line for line in meaningful):
        fail(f"{path} has no mapping-like YAML keys")
    for number, line in enumerate(lines, start=1):
        if not line.strip() or line.lstrip().startswith("#"):
            continue
        indent = len(line) - len(line.lstrip(" "))
        if indent % 2 != 0:
            fail(f"{path}:{number} uses odd indentation")


def check_yaml_files() -> None:
    if check_yaml_with_pyyaml(YAML_FILES):
        return
    for path in YAML_FILES:
        simple_yaml_check(path)


def main() -> None:
    check_required_readmes()
    check_duplicate_readmes()
    check_top_language_links()
    check_top_section_order()
    check_markdown_tables()
    check_code_blocks()
    check_local_references()
    check_product_claim_surface()
    check_install_shell_syntax()
    check_yaml_files()
    print("README surface checks passed")


if __name__ == "__main__":
    main()
