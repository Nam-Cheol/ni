#!/usr/bin/env python3
"""Render deterministic README SVG assets from local templates."""

from __future__ import annotations

import argparse
import html
from pathlib import Path
from string import Template


ROOT = Path(__file__).resolve().parents[1]
SOURCE = ROOT / "assets" / "source"
DEFAULT_OUT = ROOT / "assets"
FONT_FAMILY = '-apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif'


def xml_text(value: str) -> str:
    return html.escape(value, quote=False)


def xml_attr(value: str) -> str:
    return html.escape(value, quote=True)


def load_template(name: str) -> Template:
    return Template((SOURCE / name).read_text(encoding="utf-8"))


def render_rows(labels: list[str], accent_color: str) -> str:
    rows = []
    for index, label in enumerate(labels):
        y = index * 42
        rows.append(
            "\n".join(
                [
                    f'      <g transform="translate(0 {y})">',
                    f'        <circle cx="10" cy="10" r="6" fill="{xml_attr(accent_color)}"/>',
                    f'        <path d="M28 10 H396" stroke="#D8E1EE" stroke-width="2" stroke-linecap="round"/>',
                    f'        <text x="28" y="16" fill="#475569">{xml_text(label)}</text>',
                    "      </g>",
                ]
            )
        )
    return "\n".join(rows)


CARD_DATA = [
    {
        "filename": "card-start.svg",
        "title": "Start path",
        "desc": "Initialize a planning workspace, check readiness, lock intent, and compile a prompt.",
        "heading": "Start",
        "labels": ["init workspace", "status gate", "end lock", "run prompt"],
        "footer": "prompt compilation only",
        "footer_width": "214",
        "soft_color": "#DCEBFF",
        "accent_color": "#4FBF9F",
        "icon_path": "M15 24 H33 M26 15 L35 24 L26 33",
        "icon_stroke": "#111827",
    },
    {
        "filename": "card-contract.svg",
        "title": "Contract path",
        "desc": "Planning records become an explicit project contract before execution.",
        "heading": "Contract",
        "labels": ["requirements", "risks", "evaluations", "open questions"],
        "footer": "docs and contract stay paired",
        "footer_width": "246",
        "soft_color": "#EDE9FE",
        "accent_color": "#8B7CF6",
        "icon_path": "M17 14 H31 L36 19 V34 H17 Z M31 14 V20 H36 M22 25 H31 M22 31 H28",
        "icon_stroke": "#FFFFFF",
    },
    {
        "filename": "card-handoff.svg",
        "title": "Handoff path",
        "desc": "A locked plan compiles into a bounded handoff prompt or mutable seed material.",
        "heading": "Handoff",
        "labels": ["lock hash", "source of truth", "bounded prompt", "stop on drift"],
        "footer": "downstream seeds are derived",
        "footer_width": "238",
        "soft_color": "#DCFCE7",
        "accent_color": "#5B8DEF",
        "icon_path": "M16 18 H28 M24 14 L30 18 L24 22 M32 30 H20 M24 26 L18 30 L24 34",
        "icon_stroke": "#FFFFFF",
    },
    {
        "filename": "card-pain-vague-intent.svg",
        "title": "Vague intent",
        "desc": "A plausible prompt can still hide missing users, acceptance criteria, risks, and blockers.",
        "heading": "Vague intent",
        "labels": ["missing users", "unclear acceptance", "open blockers"],
        "footer": "sounds ready, is not",
        "footer_width": "184",
        "soft_color": "#DCEBFF",
        "accent_color": "#5B8DEF",
        "icon_path": "M16 18 H32 M16 26 H28 M16 34 H24",
        "icon_stroke": "#FFFFFF",
    },
    {
        "filename": "card-pain-early-execution.svg",
        "title": "Early execution",
        "desc": "Work should not begin just because a request sounds plausible.",
        "heading": "Early run",
        "labels": ["before agreement", "model momentum", "unchecked assumptions"],
        "footer": "pause before handoff",
        "footer_width": "186",
        "soft_color": "#EDE9FE",
        "accent_color": "#8B7CF6",
        "icon_path": "M16 16 L34 24 L16 32 Z M36 16 V32",
        "icon_stroke": "#FFFFFF",
    },
    {
        "filename": "card-pain-rework.svg",
        "title": "Rework",
        "desc": "Hidden assumptions get expensive after people, models, or tools start from the wrong plan.",
        "heading": "Rework",
        "labels": ["wrong plan", "late blockers", "lost context"],
        "footer": "cost moves downstream",
        "footer_width": "200",
        "soft_color": "#DCFCE7",
        "accent_color": "#4FBF9F",
        "icon_path": "M33 17 A12 12 0 1 0 36 28 M36 17 V28 H25",
        "icon_stroke": "#111827",
    },
    {
        "filename": "card-payoff-capture-intent.svg",
        "title": "Capture intent",
        "desc": "Planning conversation becomes explicit docs and a contract draft.",
        "heading": "Capture",
        "labels": ["purpose", "requirements", "open questions"],
        "footer": "conversation becomes records",
        "footer_width": "250",
        "soft_color": "#DCEBFF",
        "accent_color": "#4FBF9F",
        "icon_path": "M17 14 H31 L36 19 V34 H17 Z M31 14 V20 H36 M22 25 H31 M22 31 H28",
        "icon_stroke": "#111827",
    },
    {
        "filename": "card-payoff-lock-contract.svg",
        "title": "Lock contract",
        "desc": "Readiness and lock commands gate the accepted plan, hashes, and source of truth.",
        "heading": "Lock",
        "labels": ["status gate", "hash proof", "source of truth"],
        "footer": "ready means checked",
        "footer_width": "180",
        "soft_color": "#EDE9FE",
        "accent_color": "#8B7CF6",
        "icon_path": "M18 24 H34 V36 H18 Z M22 24 V19 A8 8 0 0 1 30 19 V24",
        "icon_stroke": "#FFFFFF",
    },
    {
        "filename": "card-payoff-handoff-safely.svg",
        "title": "Handoff safely",
        "desc": "A valid locked plan compiles into a bounded prompt or derived seed material.",
        "heading": "Handoff",
        "labels": ["bounded prompt", "stop on drift", "no execution"],
        "footer": "compiled, not run",
        "footer_width": "158",
        "soft_color": "#DCFCE7",
        "accent_color": "#5B8DEF",
        "icon_path": "M16 18 H28 M24 14 L30 18 L24 22 M32 30 H20 M24 26 L18 30 L24 34",
        "icon_stroke": "#FFFFFF",
    },
]

BADGE_DATA = [
    {
        "filename": "badge-english.svg",
        "title": "English language chip",
        "desc": "Link chip for the English README.",
        "label": "English",
        "width": "84",
        "fill": "#F8FAFC",
        "stroke": "#94A3B8",
        "text_fill": "#172033",
    },
    {
        "filename": "badge-korean.svg",
        "title": "Korean language chip",
        "desc": "Link chip for the Korean README.",
        "label": "한국어",
        "width": "84",
        "fill": "#F8FAFC",
        "stroke": "#94A3B8",
        "text_fill": "#172033",
    },
]


def render_hero(out_dir: Path) -> Path:
    template = load_template("hero.template.svg")
    content = template.substitute(
        title=xml_text("ni hero"),
        desc=xml_text("Project Intent Compiler for AI Agents."),
        font_family=xml_attr(FONT_FAMILY),
    )
    path = out_dir / "hero.svg"
    path.write_text(content + "\n", encoding="utf-8")
    return path


def render_card(out_dir: Path, data: dict[str, str | list[str]], index: int) -> Path:
    template = load_template("card.template.svg")
    labels = data["labels"]
    if not isinstance(labels, list):
        raise TypeError("card labels must be a list")
    content = template.substitute(
        title=xml_text(str(data["title"])),
        desc=xml_text(str(data["desc"])),
        gradient_id=f"card-bg-{index}",
        soft_color=xml_attr(str(data["soft_color"])),
        accent_color=xml_attr(str(data["accent_color"])),
        font_family=xml_attr(FONT_FAMILY),
        icon_path=xml_attr(str(data["icon_path"])),
        icon_stroke=xml_attr(str(data["icon_stroke"])),
        heading=xml_text(str(data["heading"])),
        rows=render_rows([str(label) for label in labels], str(data["accent_color"])),
        footer_width=xml_attr(str(data["footer_width"])),
        footer_fill=xml_attr(str(data["soft_color"])),
        footer=xml_text(str(data["footer"])),
    )
    path = out_dir / str(data["filename"])
    path.write_text(content + "\n", encoding="utf-8")
    return path


def render_badge(out_dir: Path, data: dict[str, str], index: int) -> Path:
    template = load_template("badge.template.svg")
    width = int(data["width"])
    content = template.substitute(
        width=xml_attr(data["width"]),
        inner_width=xml_attr(str(width - 1)),
        title_id=xml_attr(f"badge-title-{index}"),
        desc_id=xml_attr(f"badge-desc-{index}"),
        title=xml_text(data["title"]),
        desc=xml_text(data["desc"]),
        label=xml_text(data["label"]),
        text_x=xml_attr(str(width // 2)),
        fill=xml_attr(data["fill"]),
        stroke=xml_attr(data["stroke"]),
        text_fill=xml_attr(data["text_fill"]),
        font_family=xml_attr(FONT_FAMILY),
    )
    path = out_dir / data["filename"]
    path.write_text(content + "\n", encoding="utf-8")
    return path


def render_assets(out_dir: Path) -> list[Path]:
    out_dir.mkdir(parents=True, exist_ok=True)
    rendered = [render_hero(out_dir)]
    for index, data in enumerate(CARD_DATA, start=1):
        rendered.append(render_card(out_dir, data, index))
    for index, data in enumerate(BADGE_DATA, start=1):
        rendered.append(render_badge(out_dir, data, index))
    return rendered


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument(
        "--out-dir",
        type=Path,
        default=DEFAULT_OUT,
        help="Directory for generated SVG output. Defaults to assets/.",
    )
    args = parser.parse_args()

    for path in render_assets(args.out_dir):
        print(path.relative_to(ROOT) if path.is_relative_to(ROOT) else path)


if __name__ == "__main__":
    main()
