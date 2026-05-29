# Direct-to-Agent Risk Analysis

The direct prompt sounds actionable enough that a downstream actor might start
building a dashboard immediately. The readiness risk is that the actor would
need to invent intent before implementation.

Visible gaps:

- primary users are described only as "the customer team";
- the outcome behind "who needs attention" is not defined;
- account signals, source systems, required fields, and freshness rules are
  missing;
- privacy and access-control boundaries for customer data are not stated;
- "easy to use" has no acceptance test;
- "ready for the next planning meeting" does not name the meeting date,
  audience, or minimum useful scope;
- non-goals are absent.

Material hidden assumptions:

- customer success managers are the primary users;
- attention means renewal risk, support escalation, product usage drop, or
  unpaid invoice status;
- trusted source data already exists;
- customer health signals can be exposed without additional review;
- a table or chart is enough;
- historical trends are either required or out of scope.

Direct-to-agent boundary:

This analysis does not run an agent, call a model API, scaffold a dashboard, or
claim implementation results.
