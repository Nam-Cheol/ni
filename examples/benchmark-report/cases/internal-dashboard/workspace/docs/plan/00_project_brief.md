# Project brief

## Product type

Software dashboard.

## Delivery surfaces

- Web.

## Purpose

Plan an internal customer-account dashboard before implementation so the
customer team can decide which accounts need attention without downstream actors
inventing users, account signals, data boundaries, or meeting acceptance
criteria.

## Problem

The original request asks for a dashboard for "the customer team" to see what
is going on with accounts and know who needs attention. That is plausible
implementation input, but the users, source systems, attention criteria,
privacy boundary, review evidence, and next-meeting scope are still unresolved.

## Success definition

This benchmark workspace succeeds when the vague request is captured as
pre-runtime planning evidence and `ni status --proof --next-questions` reports
the unresolved intent as `BLOCKED` before any dashboard work starts.
