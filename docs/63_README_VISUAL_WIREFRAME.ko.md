# README Visual Wireframe

Task 104는 README를 다시 rewrite하기 전에 visual README layout을 정의한다.

이 문서는 layout contract only다. `README.md`를 rewrite하지 않고, new product
claims를 추가하지 않고, visual assets를 생성하지 않고, runtime behavior를
바꾸지 않는다.

## Design Intent

README는 deterministic kernel proof가 받치는 lightweight product pamphlet처럼
읽혀야 한다. First screen은 README를 protocol documentation으로 만들지 않고도
value를 분명히 보여줘야 한다.

이미 README와 visual spec에서 정한 core message를 사용한다:

```text
Don't run the agent yet. Compile the intent first.
```

Technical depth는 linked docs에 둔다. README visuals는 scanning을 도울 수
있지만 important meaning은 Markdown text에도 존재해야 한다.

## README Section Order

다음 README rewrite는 이 순서를 따른다:

1. Hero banner
2. Language and trust badges
3. Slogan
4. Three pain cards
5. Three ni payoff cards
6. Start in 60 seconds
7. Choose your path
8. Demo
9. What ni is not
10. Read next

## Wireframe

```text
[hero.svg]

Language chips        Trust badges
English | 한국어       MIT | CI workflow | Security | Docs

Don't run the agent yet. Compile the intent first.
Project Intent Compiler for AI Agents.

[pain card]           [pain card]           [pain card]
Vague intent          Early execution       Rework
Markdown fallback     Markdown fallback     Markdown fallback

[payoff card]         [payoff card]         [payoff card]
Capture intent        Lock contract         Handoff safely
Markdown fallback     Markdown fallback     Markdown fallback

## Start in 60 seconds
Markdown commands and source-first path.

## Choose your path
Markdown table for source, verified install paths, and assisted planning
paths. Codex and Claude may appear here only as usage path labels.

## Demo
Markdown command transcript showing a vague prompt blocked before execution.

## What ni is not
Short Markdown boundary block.

## Read next
Markdown link map to deeper docs.
```

## Section Rules

| Section | Format | Asset role | Markdown fallback |
| --- | --- | --- | --- |
| Hero banner | SVG image plus Markdown text | `assets/hero.svg`를 사용한다. | Slogan과 product description은 image 아래 Markdown으로도 있어야 한다. |
| Language badges | SVG chips | Local language chip SVGs를 사용한다. | Link labels와 alt text가 language를 식별해야 한다. |
| Trust badges | Badges | Existing factual remote badges는 non-essential로 남을 수 있다. | Short Markdown trust-signal sentence가 facts를 반복해야 한다. |
| Slogan | Markdown | 별도 SVG card 없음. | Slogan은 primary text이며 image-only이면 안 된다. |
| Three pain cards | SVG cards plus Markdown headings | New local SVG cards가 three pains를 reinforce할 수 있다. | 각 card에는 가까운 Markdown heading과 one-sentence explanation이 필요하다. |
| Three ni payoff cards | SVG cards plus Markdown headings | New local SVG cards가 three payoffs를 reinforce할 수 있다. | 각 card에는 가까운 Markdown heading과 one-sentence explanation이 필요하다. |
| Start in 60 seconds | Markdown | SVG card 없음. | Commands는 copyable text로 남아야 한다. |
| Choose your path | Markdown table | SVG card 없음. | Usage boundaries는 text로 남아야 한다. |
| Demo | Markdown transcript | SVG card 없음. | Blocked outcome은 text로 보여야 한다. |
| What ni is not | Markdown | SVG card 없음. | Product boundaries는 text로 남아야 한다. |
| Read next | Markdown table or list | SVG card 없음. | Links는 normal Markdown links로 남아야 한다. |

## Visual Card Copy

Card text는 짧게 유지한다. SVG 안에 long explanations를 넣지 않는다.

### Pain Cards

| Card | SVG label | Markdown fallback sentence |
| --- | --- | --- |
| Vague intent | `Vague intent` | Prompt가 actionable해 보여도 users, acceptance criteria, risks, non-goals, blocker questions가 빠져 있을 수 있다. |
| Early execution | `Early execution` | Request가 plausible하게 들린다는 이유만으로 work가 시작되면 안 된다. |
| Rework | `Rework` | Hidden assumptions는 humans, models, tools가 wrong plan에서 시작한 뒤 expensive해진다. |

### ni Payoff Cards

| Card | SVG label | Markdown fallback sentence |
| --- | --- | --- |
| Capture intent | `Capture intent` | Planning conversation이 explicit docs와 contract draft가 된다. |
| Lock contract | `Lock contract` | `ni status`와 `ni end`가 readiness, hashes, lock creation을 gate한다. |
| Handoff safely | `Handoff safely` | `ni run`은 valid locked plan에서 bounded prompt 또는 seed를 compile한다. |

## Mobile-Friendly Fallback

SVG cards가 render되지 않거나, 좁은 mobile viewport에서 wrap이 어색하거나,
screen reader로 읽히는 경우에도 README는 의미가 통해야 한다.

모든 visual card group은 이 pattern을 사용한다:

```markdown
<p align="center">
  <img src="assets/card-pain-vague-intent.svg" alt="Vague intent: missing users, acceptance criteria, risks, non-goals, or blockers can hide inside a plausible prompt." width="32%">
</p>

### Vague intent

A prompt can sound actionable while users, acceptance criteria, risks,
non-goals, or blocker questions are still missing.
```

Rules:

- Card alt text는 file name만 반복하지 않고 card를 설명해야 한다.
- 같은 meaning이 card 가까운 Markdown에 직접 있어야 한다.
- Headings는 plain Markdown headings로 남아야 한다.
- Mobile에서 unreadable해지는 side-by-side layout을 피한다.
- Three cards를 HTML row 하나로 보여주더라도 images가 실패했을 때 이어지는
  Markdown headings와 sentences만으로 의미가 통해야 한다.

## Visual Sales Guardrails

- Hero, pain cards, payoff cards에서는 specific harness products를 언급하지
  않는다.
- Codex와 Claude는 usage path sections에서만 assisted planning UX label로
  나타날 수 있으며 kernel authority처럼 설명하면 안 된다.
- Package distribution, hosted service availability, published release assets는
  실제 사실이 있을 때만 claim한다.
- `ni run`이 shell commands, queues, agents, downstream work를 execute한다고
  암시하지 않는다.
- Contract authoring CLI commands를 추가하지 않는다.
- Protocol details는 docs에 두고 README에서는 링크한다.

## Required Visual Assets

Required existing assets:

| Asset | Status | Role |
| --- | --- | --- |
| `assets/hero.svg` | Existing | Hero banner. |
| `assets/badge-english.svg` | Existing | English language chip. |
| `assets/badge-korean.svg` | Existing | Korean language chip. |

Required new card assets for the next README rewrite:

| Asset | Role |
| --- | --- |
| `assets/card-pain-vague-intent.svg` | Pain card for vague intent. |
| `assets/card-pain-early-execution.svg` | Pain card for early execution. |
| `assets/card-pain-rework.svg` | Pain card for rework. |
| `assets/card-payoff-capture-intent.svg` | Payoff card for capture intent. |
| `assets/card-payoff-lock-contract.svg` | Payoff card for lock contract. |
| `assets/card-payoff-handoff-safely.svg` | Payoff card for safe handoff. |

Optional legacy or transitional assets:

| Asset | Rule |
| --- | --- |
| `assets/card-why.svg` | Pain and payoff card set이 대체할 때까지 남을 수 있다. |
| `assets/card-start.svg` | Markdown `Start in 60 seconds` section을 대체하면 안 된다. |
| `assets/card-docs.svg` | Markdown `Read next` section을 대체하면 안 된다. |

모든 card assets는 `docs/60_VISUAL_DESIGN_SPEC.md`를 따라야 한다: local SVG,
short labels only, meaningful alt text, no remote dependencies, essential copy
that appears only inside the image 금지.
