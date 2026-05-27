# 02. Direct-to-agent risk

A direct implementation path may start by guessing:

- the team is customer support,
- the data source is a local JSON fixture,
- the dashboard needs charts, filters, and ticket assignment,
- "stay on top" means SLA breach monitoring,
- authentication is unnecessary,
- the implementation target is a web app,
- fake data is acceptable,
- success means the page renders.

Those assumptions might be reasonable, but they are still hidden product
decisions. Once code exists, the user has to review implementation details just
to discover that the intent was never clarified.

The failure is not that the agent is lazy. The failure is that execution starts
before the project contract exists.
