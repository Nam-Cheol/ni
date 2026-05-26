#!/usr/bin/env python3
"""Validate that target exports stay seed material.

The checker is intentionally path-oriented. Target-specific export contents may
evolve, but these downstream runtime packet shapes must stay outside NI-owned
export directories.
"""

from __future__ import annotations

import argparse
import os
import sys
from pathlib import Path


RULES = {
    "hyper-run": {
        "required": {
            "plan.md",
            "ni-context.md",
            "readiness-expectations.md",
            "evidence-requirements.md",
            "first-run-focus.md",
        },
        "forbidden": {
            ".hyper",
            ".hyper/goals",
            ".hyper/goals/GOAL-0001",
            "tasks.md",
            "evidence.md",
            "review.md",
            "next.md",
        },
    },
    "namba-ai": {
        "required": {
            "planning.md",
            "ni-lock-summary.md",
            "capability-map.md",
            "evaluation-map.md",
            "risk-map.md",
            "suggested-spec-boundaries.md",
        },
        "forbidden": {
            ".namba",
            ".namba/specs",
            "SPEC-001.md",
            "SPEC-002.md",
            "SPEC_SEQUENCE.md",
            "specs",
            "tasks.md",
            "run.md",
            "sync.md",
            "pr.md",
            "land.md",
        },
    },
    "ouroboros": {
        "required": {"ouroboros-seed-notes.md"},
        "forbidden": {
            ".ouroboros",
            ".ouroboros/runtime",
            "execute",
            "execute.md",
            "evaluate",
            "evaluate.md",
            "evolve",
            "evolve.md",
            "interview.md",
            "crystallize.md",
            "runtime",
        },
    },
    "spec-kit": {
        "required": {"spec-kit-seed-notes.md"},
        "forbidden": {
            ".specify",
            ".specify/specs",
            ".specify/memory",
            ".github/prompts",
            ".claude/commands",
            ".codex/commands",
            "slash-commands.md",
            "commands",
            "specify.md",
            "plan.md",
            "tasks.md",
        },
    },
}


def rel_paths(root: Path) -> set[str]:
    paths: set[str] = set()
    for dirpath, dirnames, filenames in os.walk(root):
        current = Path(dirpath)
        for name in dirnames:
            paths.add((current / name).relative_to(root).as_posix())
        for name in filenames:
            paths.add((current / name).relative_to(root).as_posix())
    return paths


def validate(target: str, export_dir: Path) -> list[str]:
    if target not in RULES:
        return [f"unsupported target {target!r}"]
    if not export_dir.is_dir():
        return [f"export dir does not exist: {export_dir}"]

    rules = RULES[target]
    paths = rel_paths(export_dir)
    errors: list[str] = []

    missing = sorted(rules["required"] - paths)
    if missing:
        errors.append(f"{target}: missing required seed files: {', '.join(missing)}")

    unexpected = sorted(paths - rules["required"])
    if unexpected:
        errors.append(f"{target}: created non-seed paths: {', '.join(unexpected)}")

    forbidden = sorted(rules["forbidden"] & paths)
    if forbidden:
        errors.append(f"{target}: created runtime packet paths: {', '.join(forbidden)}")

    return errors


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Check that an ni export target produced seed material only."
    )
    parser.add_argument("--target", required=True, choices=sorted(RULES))
    parser.add_argument("--dir", required=True, type=Path)
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    errors = validate(args.target, args.dir)
    if errors:
        for error in errors:
            print(error, file=sys.stderr)
        return 1
    print(f"{args.target} export conforms to seed-only boundary")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
