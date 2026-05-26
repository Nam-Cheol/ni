#!/usr/bin/env python3
from __future__ import annotations

import argparse
import re
import sys
from dataclasses import dataclass
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]

RISKY_PATTERNS = [
    (re.compile(r"\bagent teams?\b", re.IGNORECASE), "agent team"),
    (re.compile(r"\bparallel execution\b", re.IGNORECASE), "parallel execution"),
    (re.compile(r"\btask queues?\b", re.IGNORECASE), "task queue"),
    (re.compile(r"\bSPEC runner\b", re.IGNORECASE), "SPEC runner"),
    (re.compile(r"\bautomatic implementation\b", re.IGNORECASE), "automatic implementation"),
    (re.compile(r"\bPR automation\b", re.IGNORECASE), "PR automation"),
    (re.compile(r"\brelease automation\b", re.IGNORECASE), "release automation"),
    (
        re.compile(r"\b(?:execution[- ]?)?evidence[- ]loop\b", re.IGNORECASE),
        "execution evidence loop",
    ),
    (re.compile(r"\bself-evolving harness\b", re.IGNORECASE), "self-evolving harness"),
    (re.compile(r"\bhyper complete\b", re.IGNORECASE), "hyper complete"),
    (re.compile(r"\bhyper run clone\b", re.IGNORECASE), "hyper run clone"),
]

EXCLUDED_DIRS = {
    ".git",
    ".ni",
    "docs/examples",
    "docs/experiments",
    "docs/plan",
    "scripts/testdata",
}

SOURCE_ROOTS = [
    Path("cmd/ni"),
    Path("internal/core"),
]

DOC_PATHS = [
    Path("AGENTS.md"),
    Path("MANIFEST.md"),
    Path("README.md"),
]

DOC_ROOTS = [
    Path("docs"),
]

TEXT_SUFFIXES = {".go", ".md", ".json"}

ALLOW_COMMENT = "ni-boundary-allow"

BOUNDARY_LANGUAGE = re.compile(
    r"\b("
    r"not|"
    r"no|"
    r"never|"
    r"forbidden|"
    r"non-goals?|"
    r"do not|"
    r"must not|"
    r"should not|"
    r"does not|"
    r"is not|"
    r"must stay|"
    r"must remain|"
    r"outside|"
    r"downstream"
    r")\b",
    re.IGNORECASE,
)

RELATED_WORK_LANGUAGE = re.compile(
    r"\b("
    r"related work|"
    r"positioning?|"
    r"adjacent|"
    r"comparison|"
    r"compare|"
    r"boundary|"
    r"examples?|"
    r"optimize|"
    r"downstream|"
    r"not|"
    r"do not|"
    r"must not|"
    r"should not|"
    r"outside"
    r")\b",
    re.IGNORECASE,
)


@dataclass(frozen=True)
class Finding:
    path: Path
    line_no: int
    phrase: str
    line: str


def is_excluded(path: Path) -> bool:
    path_parts = path.parts
    for excluded in EXCLUDED_DIRS:
        excluded_parts = Path(excluded).parts
        if path_parts[: len(excluded_parts)] == excluded_parts:
            return True
    return False


def candidate_files(root: Path) -> list[Path]:
    candidates: set[Path] = set()

    for rel in DOC_PATHS:
        path = root / rel
        if path.exists() and path.is_file():
            candidates.add(path)

    for rel_root in DOC_ROOTS + SOURCE_ROOTS:
        path_root = root / rel_root
        if not path_root.exists():
            continue
        for path in path_root.rglob("*"):
            if not path.is_file():
                continue
            rel = path.relative_to(root)
            if is_excluded(rel):
                continue
            if path.suffix in TEXT_SUFFIXES:
                candidates.add(path)

    return sorted(candidates)


def has_allow_comment(lines: list[str], index: int) -> bool:
    if ALLOW_COMMENT in lines[index]:
        return True
    if index > 0 and ALLOW_COMMENT in lines[index - 1]:
        return True
    return False


def is_boundary_context(lines: list[str], index: int) -> bool:
    start = max(0, index - 10)
    context = "\n".join(lines[start : index + 1])
    return bool(BOUNDARY_LANGUAGE.search(context))


def is_related_work_context(rel: Path, lines: list[str], index: int) -> bool:
    if rel != Path("docs/11_RELATED_WORK.md"):
        return False
    start = max(0, index - 8)
    context = "\n".join(lines[start : index + 1])
    return bool(RELATED_WORK_LANGUAGE.search(context))


def scan_file(root: Path, path: Path) -> list[Finding]:
    rel = path.relative_to(root)

    try:
        lines = path.read_text(encoding="utf-8").splitlines()
    except UnicodeDecodeError:
        return []

    findings: list[Finding] = []
    for index, line in enumerate(lines):
        for pattern, phrase in RISKY_PATTERNS:
            if not pattern.search(line):
                continue
            if (
                has_allow_comment(lines, index)
                or is_boundary_context(lines, index)
                or is_related_work_context(rel, lines, index)
            ):
                continue
            findings.append(
                Finding(
                    path=rel,
                    line_no=index + 1,
                    phrase=phrase,
                    line=line.strip(),
                )
            )
    return findings


def scan_root(root: Path) -> list[Finding]:
    findings: list[Finding] = []
    for path in candidate_files(root):
        findings.extend(scan_file(root, path))
    return findings


def print_findings(findings: list[Finding]) -> None:
    for finding in findings:
        print(
            f"Core boundary drift: {finding.path}:{finding.line_no}: "
            f"{finding.phrase}: {finding.line}"
        )
    if findings:
        print(
            "Core boundary guard failed. Keep runtime execution claims outside "
            "ni-kernel docs/source, or add a local ni-boundary-allow comment "
            "only for an explicit boundary/comparison use."
        )


def run_self_test() -> bool:
    fixture_root = ROOT / "scripts/testdata/core-boundary"
    allowed_root = fixture_root / "allowed"
    blocked_root = fixture_root / "blocked"

    allowed_findings = scan_root(allowed_root)
    if allowed_findings:
        print("Core boundary self-test failed: related-work fixture was flagged.")
        print_findings(allowed_findings)
        return False

    blocked_findings = scan_root(blocked_root)
    if not blocked_findings:
        print("Core boundary self-test failed: blocked fixture was not flagged.")
        return False

    expected = {"task queue", "parallel execution"}
    observed = {finding.phrase for finding in blocked_findings}
    missing = expected - observed
    if missing:
        print(
            "Core boundary self-test failed: blocked fixture missed "
            + ", ".join(sorted(missing))
            + "."
        )
        print_findings(blocked_findings)
        return False

    print("Core boundary self-test passed")
    return True


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--root", type=Path, default=ROOT)
    parser.add_argument("--self-test", action="store_true")
    args = parser.parse_args()

    root = args.root.resolve()
    findings = scan_root(root)
    if findings:
        print_findings(findings)
        return 1

    if args.self_test and not run_self_test():
        return 1

    print("Core boundary check passed")
    return 0


if __name__ == "__main__":
    sys.exit(main())
