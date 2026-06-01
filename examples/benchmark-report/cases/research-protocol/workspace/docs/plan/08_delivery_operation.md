# Delivery and operation

## Delivery surfaces

- document
- workflow
- human_service

## Initial delivery

Initial delivery is a checked-in benchmark case with planning docs, a contract,
status proof, next-question evidence, a measurement table, and boundary notes.
It is not a research protocol approved for fieldwork.

## Operating model

- Planning docs are committed to git.
- Contract JSON is committed to git.
- `ni status --proof --next-questions` is the authority for measured
  readiness.
- If readiness is `BLOCKED`, do not lock and do not compile a prompt.
- If readiness is `READY` or `READY_WITH_DEFERRALS`, `ni end` and `ni run` may
  run only inside this isolated benchmark workspace.
- The synthetic pre-fieldwork checklist is: research question accepted;
  locations selected from approved list; data types allowed/prohibited; consent
  script prepared; translation/accessibility needs addressed; field safety
  rules documented; reviewer assigned; risks mitigated; non-goals explicit;
  `ni status` not `BLOCKED`.
