# README Rendering Report

Task 112는 `README.md`와 `README.ko.md`가 GitHub에서 public product pamphlet로
깨끗하게 보이는지 검증한다. 이 문서는 rendering fact만 기록하며, product claim,
install path, release binary claim, runtime behavior를 추가하지 않는다.

## Hero Rendering

두 README의 top surface는 이제 같은 순서로 rendering된다:

1. Hero image: `assets/hero.svg`
2. Language chips: `assets/badge-english.svg`, `assets/badge-korean.svg`
3. Factual trust badges: MIT license, CI workflow, security policy, docs index
4. Slogan
5. One-line product description
6. Why ni
7. Start in 60 seconds

Badge row와 slogan 사이에 보이던 duplicate trust-signal text는 제거했다.
Slogan과 one-line product description은 image-only content가 아니라
Markdown/HTML text로 남아 있다.

## Badge Rendering

| Badge group | Rendering check | Result |
| --- | --- | --- |
| Language chips | Local SVG images가 `README.md`, `README.ko.md`로 link된다. | Pass |
| Trust badges | Remote shields.io badge images가 existing repository files로 link된다. | Pass |
| Badge facts | Badges는 `LICENSE`, `.github/workflows/ci.yml`, `SECURITY.md`, `docs/00_START_HERE.md`를 가리킨다. | Pass |

## Table Rendering

두 README는 각각 Markdown table 두 개를 가진다:

| README | Table | Columns | Result |
| --- | --- | ---: | --- |
| `README.md` | Choose your path | 4 | Pass |
| `README.md` | Read next | 2 | Pass |
| `README.ko.md` | Choose your path | 4 | Pass |
| `README.ko.md` | Read next | 2 | Pass |

Surface checker는 이제 separator row뿐 아니라 body row의 column count도
검증한다.

## Code Block Rendering

두 README 모두 info string이 있는 balanced fenced code block을 사용한다:

| README | Fence languages | Result |
| --- | --- | --- |
| `README.md` | `bash`, `bash`, `bash`, `text` | Pass |
| `README.ko.md` | `bash`, `bash`, `bash`, `text` | Pass |

## Link Check Summary

모든 local README asset reference가 존재한다. 모든 local README link는 존재하는
file 또는 directory를 가리킨다.

| Reference class | Result |
| --- | --- |
| Local assets | Pass |
| HTML `href` links | Pass |
| Markdown links | Pass |
| Root duplicate README files | Pass: none present |

## English/Korean Parity

English README와 Korean README는 같은 public structure를 유지한다:

1. Hero and badges
2. Why ni
3. What ni gives you
4. Start in 60 seconds
5. Choose your path
6. Demo
7. What ni is not
8. Read next

Korean README는 companion Korean docs가 maintained되는 곳에서 Korean link를
유지한다. 대상은 `docs/product-story.ko.md`, `docs/no-terminal.ko.md`,
`docs/commands.ko.md`, `docs/63_README_VISUAL_WIREFRAME.ko.md`다.

## Product Claim Audit

Hero와 sales pitch는 specific harness 또는 runtime product를 언급하지 않는다.
`Codex`와 `Claude`는 `Model workspaces` usage path에서만 등장하며, 이는 assisted
planning을 설명할 뿐 kernel authority를 뜻하지 않는다.

README는 계속 `ni run`이 bounded prompt 또는 seed를 compile하며 shell commands,
queues, agents, downstream work를 실행하지 않는다고 말한다.

## Remaining Visual Issues

현재 Markdown surface에는 blocking rendering issue가 남아 있지 않다.

Known non-blocking visual constraints는 남아 있다:

- Trust badges는 remote shields.io image rendering에 의존한다.
- Local SVGs는 live text를 포함하고 renderer font fallback에 의존한다.
- Korean README는 이번 version에서 English-text SVG assets를 의도적으로 재사용한다.

