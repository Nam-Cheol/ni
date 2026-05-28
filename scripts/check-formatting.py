#!/usr/bin/env python3
"""Formatting regression checks for public packaging files."""

from __future__ import annotations

import os
import re
import shutil
import subprocess
from pathlib import Path


ROOT = Path(os.environ.get("NI_FORMAT_ROOT", Path(__file__).resolve().parents[1]))

MARKDOWN_TABLES = {
    "README.md": ["## Choose your path", "## Read next"],
    "README.ko.md": ["## Choose your path", "## 다음에 읽을 것"],
    "packages/codex-skills/README.md": ["## Skills"],
    "packages/codex-skills/README.ko.md": ["## Skills"],
    "packages/claude-skills/README.md": ["## Skills"],
    "packages/claude-skills/README.ko.md": ["## Skills"],
}

YAML_FILES = [
    *[
        str(path.relative_to(ROOT))
        for path in sorted((ROOT / ".github" / "workflows").glob("*.yml"))
    ],
    ".goreleaser.yaml",
]

SHELL_FILES = [
    "install.sh",
    *[str(path.relative_to(ROOT)) for path in sorted((ROOT / "scripts").glob("*.sh"))],
]

MARKDOWN_FILES = [
    "README.md",
    "README.ko.md",
    *[
        str(path.relative_to(ROOT))
        for base in ("docs", "packages")
        for path in sorted((ROOT / base).rglob("*.md"))
    ],
]

MIN_LINE_COUNTS = {
    "README.md": 80,
    "README.ko.md": 80,
    "scripts/quality.sh": 20,
    ".github/workflows/release.yml": 30,
    ".goreleaser.yaml": 35,
    "install.sh": 120,
}

SHELL_SHEBANGS = {
    "install.sh": "#!/usr/bin/env sh",
}

YAML_COLLAPSE_KEYS = {
    ".github/workflows/release.yml": ("name", "on", "permissions", "jobs", "steps"),
    ".github/workflows/ci.yml": ("name", "on", "jobs", "steps"),
    ".goreleaser.yaml": ("version", "project_name", "builds", "archives", "checksum"),
}


def fail(message: str) -> None:
    raise SystemExit(message)


def read_lines(path: str) -> list[str]:
    file_path = ROOT / path
    if not file_path.exists():
        fail(f"{path} is missing")
    return file_path.read_text(encoding="utf-8").splitlines()


def check_duplicate_readmes() -> None:
    duplicates = ["README 2.md", "README.ko 2.md"]
    present = [path for path in duplicates if (ROOT / path).exists()]
    if present:
        fail("unexpected duplicate root README files: " + ", ".join(present))


def cell_count(line: str) -> int:
    stripped = line.strip()
    if stripped.startswith("|"):
        stripped = stripped[1:]
    if stripped.endswith("|"):
        stripped = stripped[:-1]
    return len(stripped.split("|"))


def is_separator(line: str) -> bool:
    parts = line.strip().strip("|").split("|")
    return bool(parts) and all(re.fullmatch(r"\s*:?-{3,}:?\s*", part) for part in parts)


def find_table(lines: list[str], heading: str, path: str) -> tuple[int, list[str]]:
    try:
        start = lines.index(heading)
    except ValueError:
        fail(f"{path} is missing required heading: {heading}")

    table_start = None
    for index in range(start + 1, len(lines)):
        line = lines[index]
        if line.startswith("## "):
            break
        if line.strip().startswith("|"):
            table_start = index
            break

    if table_start is None:
        fail(f"{path} has no Markdown table after {heading}")

    table = []
    for line in lines[table_start:]:
        if not line.strip().startswith("|"):
            break
        table.append(line)

    return table_start + 1, table


def check_table(path: str, heading: str) -> None:
    lines = read_lines(path)
    line_number, table = find_table(lines, heading, path)

    if len(table) < 3:
        fail(f"{path}:{line_number} table after {heading} needs header, separator, and rows")
    if not is_separator(table[1]):
        fail(f"{path}:{line_number + 1} table separator is invalid after {heading}")

    expected_cells = cell_count(table[0])
    if expected_cells < 2:
        fail(f"{path}:{line_number} table after {heading} has too few columns")

    for offset, row in enumerate(table, start=0):
        cells = cell_count(row)
        if cells != expected_cells:
            fail(
                f"{path}:{line_number + offset} table after {heading} has "
                f"{cells} columns; expected {expected_cells}"
            )


def check_markdown_tables() -> None:
    for path, headings in MARKDOWN_TABLES.items():
        for heading in headings:
            check_table(path, heading)


def check_min_line_counts() -> None:
    for path, minimum in MIN_LINE_COUNTS.items():
        line_count = len(read_lines(path))
        if line_count < minimum:
            fail(f"{path} has only {line_count} lines; possible line-collapsed formatting")


def check_markdown_file_shape(path: str) -> None:
    in_fence = False
    for line_number, line in enumerate(read_lines(path), start=1):
        stripped = line.strip()

        if "```" in line:
            if not stripped.startswith("```"):
                fail(f"{path}:{line_number} contains an inline fenced code marker")
            if stripped.count("```") > 1:
                fail(f"{path}:{line_number} opens and closes a fenced code block inline")
            if not re.fullmatch(r"```[A-Za-z0-9_.+-]*", stripped):
                fail(f"{path}:{line_number} has code fence content on the fence line")
            in_fence = not in_fence
            continue

        if in_fence:
            continue

        if stripped.startswith("#"):
            if re.search(r"\S\s+#{1,6}\s+\S", stripped):
                fail(f"{path}:{line_number} contains more than one heading on one line")
            if "|" in stripped:
                fail(f"{path}:{line_number} appears to contain a collapsed heading and table row")
            continue

        if re.search(r"\S\s+#{1,6}\s+\S", line):
            fail(f"{path}:{line_number} contains a heading after prose on the same line")

        if stripped.startswith("|"):
            if re.search(r"\|\s*\|\s*\|", line):
                fail(f"{path}:{line_number} appears to contain collapsed Markdown table rows")
            if re.search(r"\|\s*(?:#{1,6}\s|```)", line):
                fail(f"{path}:{line_number} appears to contain collapsed Markdown after a table row")


def check_no_collapsed_markdown_blocks() -> None:
    for path in MARKDOWN_FILES:
        check_markdown_file_shape(path)


def check_quality_script_header() -> None:
    lines = read_lines("scripts/quality.sh")
    if lines[:2] != ["#!/usr/bin/env bash", "set -euo pipefail"]:
        fail("scripts/quality.sh must start with '#!/usr/bin/env bash' and 'set -euo pipefail'")


def check_install_script_header() -> None:
    lines = read_lines("install.sh")
    if lines[:2] != ["#!/usr/bin/env sh", "set -eu"]:
        fail("install.sh must start with '#!/usr/bin/env sh' and 'set -eu'")


def check_shell_shebang_lines() -> None:
    for path in SHELL_FILES:
        first_line = read_lines(path)[0]
        expected = SHELL_SHEBANGS.get(path, "#!/usr/bin/env bash")
        if first_line != expected:
            fail(f"{path}:1 must be exactly {expected!r}; possible collapsed shell script")


def run_check(command: list[str], label: str) -> None:
    result = subprocess.run(command, cwd=ROOT, text=True, capture_output=True)
    if result.returncode != 0:
        details = result.stderr.strip() or result.stdout.strip()
        fail(f"{label} failed: {details}")


def check_yaml_not_collapsed() -> None:
    for path in YAML_FILES:
        lines = read_lines(path)
        if len(lines) < 10:
            fail(f"{path} has only {len(lines)} lines; possible line-collapsed YAML")

        expected_keys = YAML_COLLAPSE_KEYS.get(path, ())
        for line_number, line in enumerate(lines, start=1):
            key_hits = [
                key for key in expected_keys if re.search(rf"(^|\s){re.escape(key)}:", line)
            ]
            if len(key_hits) >= 3:
                fail(
                    f"{path}:{line_number} contains multiple YAML sections on one line: "
                    + ", ".join(key_hits)
                )

            if len(line) > 240 and len(re.findall(r"(^|\s)[A-Za-z0-9_-]+:", line)) >= 3:
                fail(f"{path}:{line_number} is suspiciously long YAML with multiple keys")


def check_yaml() -> None:
    check_yaml_not_collapsed()

    if shutil.which("ruby"):
        run_check(
            [
                "ruby",
                "-e",
                "require 'yaml'; ARGV.each { |path| YAML.load_file(path) }",
                *YAML_FILES,
            ],
            "YAML parse check",
        )
        return

    for path in YAML_FILES:
        lines = read_lines(path)
        if any("\t" in line for line in lines):
            fail(f"{path} contains tab indentation")
        if len(lines) < 10:
            fail(f"{path} is too short to be the expected multi-line YAML")


def check_shell_syntax() -> None:
    for path in SHELL_FILES:
        first_line = read_lines(path)[0]
        shell = "sh"
        if first_line == "#!/usr/bin/env bash":
            shell = "bash"
        run_check([shell, "-n", path], f"shell syntax check for {path}")


def main() -> None:
    check_duplicate_readmes()
    check_markdown_tables()
    check_min_line_counts()
    check_no_collapsed_markdown_blocks()
    check_quality_script_header()
    check_install_script_header()
    check_shell_shebang_lines()
    check_yaml()
    check_shell_syntax()
    print("formatting checks passed")


if __name__ == "__main__":
    main()
