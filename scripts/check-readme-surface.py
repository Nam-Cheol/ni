#!/usr/bin/env python3
"""Lightweight README surface regression checks."""

from __future__ import annotations

import re
import subprocess
import sys
from pathlib import Path
from typing import Any


ROOT = Path(__file__).resolve().parents[1]
README_FILES = ["README.md", "README.ko.md"]
DUPLICATE_READMES = ["README 2.md", "README.ko 2.md"]
TOP_LINK_LINE_LIMIT = 20
INSTALL_SCRIPT = "install.sh"
YAML_FILES = [".github/workflows/release.yml", ".goreleaser.yaml"]
TABLE_SEPARATOR_RE = re.compile(r"\s*:?-{3,}:?\s*")


def fail(message: str) -> None:
    raise SystemExit(message)


def read_lines(path: str) -> list[str]:
    file_path = ROOT / path
    if not file_path.exists():
        fail(f"{path} is missing")
    return file_path.read_text(encoding="utf-8").splitlines()


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
    check_markdown_tables()
    check_install_shell_syntax()
    check_yaml_files()
    print("README surface checks passed")


if __name__ == "__main__":
    main()
