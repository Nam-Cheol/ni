# 08. Grill Lessons

The benchmark dogfood shows why `ni-grill` is useful even after `ni status`
says `READY`.

`READY` means deterministic readiness gates passed. It does not automatically
mean every claim is strong, every table can be quoted safely, every prompt
preserves the best warning after truncation, or every future reader will
understand the boundary between planning evidence and downstream execution.

`ni-grill` helps pressure-test planning quality, claim boundaries, and handoff
safety:

- It labels advisory pressure with `Critical`, `High`, `Medium`, `Low`, and
  `Note` so users can see what deserves attention first.
- It checks whether `READY` is scoped to the accepted artifact instead of a
  larger product or research claim.
- It asks whether acceptance evidence is specific enough for the next human
  reviewer.
- It catches phrases that could be quoted without nearby `not_measured`
  boundaries.
- It reviews whether realistic benchmark fixture details could be mistaken for
  real approval or execution instructions.
- It keeps non-goals and risks visible after the CLI gate has passed.

`ni-grill` should run before `ni-end` when the plan is important, risky, or
public-facing. It is especially useful for benchmark cases because benchmark
evidence can accidentally overclaim: a lock, a bounded prompt, or a `READY`
status can prove intent-readiness without proving product quality, downstream
agent performance, fieldwork readiness, research approval, or empirical
effect.

The internal-dashboard grill confirmed that `READY` applies only to benchmark
planning-meeting artifact readiness. The useful remaining questions were about
inline claim boundaries, source-material approval, role-based acceptance, and
prompt-boundary review.

The research-protocol grill confirmed that `READY` applies only to synthetic
benchmark protocol planning artifact readiness. The useful remaining questions
were about real research approval boundaries, fixture reviewer roles, safety
rules as checklist criteria, repeated synthetic labels, and prompt-boundary
review.

The pattern is simple: run `ni status --proof --next-questions` first, honor
any deterministic blockers, then use `ni-grill` to challenge accepted or nearly
accepted planning content. Keep the output budgeted: at most 5 findings by
default, at most 3 `Critical`/`High` findings first when they exist, and a
summary for omitted lower-priority findings.

`ni-grill` must not create new empirical claims, weaken readiness gates,
approve lock, execute generated prompts, call model APIs, or implement
products.
