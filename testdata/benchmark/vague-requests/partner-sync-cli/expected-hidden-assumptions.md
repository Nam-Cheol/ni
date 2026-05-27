# Expected Hidden Assumptions: Partner Sync CLI

- Partner records have stable identifiers that can prevent duplicates.
- Ops users can safely run a CLI from their environment.
- The system has an API or data export that supports syncing.
- "Sync" means create and update, not delete, merge, or reconcile conflicts.
- Authentication, audit logs, retries, and dry-run behavior can be decided by
  the implementer.
- Data privacy and partner contractual constraints are not blockers.
