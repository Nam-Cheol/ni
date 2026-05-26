# Readiness profiles

Readiness profiles are planning profiles. They tune how strict `ni status` is before a plan can be locked. They do not describe implementation stages, runtime behavior, task queues, agent execution, or sprint management.

Profiles must work for software and non-software products. The profile names are maturity labels for planning confidence only.

## Supported profiles

```text
concept
prototype
mvp
beta
production
```

`prototype` is the default profile.

## Contract field

The selected profile lives in `.ni/contract.json`:

```json
{
  "schema": "ni.contract.v0",
  "readiness_profile": "prototype"
}
```

## Profile definitions

Profile definitions live in `.ni/readiness.profiles.json`. Each profile maps readiness issue IDs to deterministic severities:

```text
blocker   prevents locking
deferral  allows READY_WITH_DEFERRALS
```

The kernel decides readiness from those severities. A model may explain the result, but it may not override it.

## Profile intent

```text
concept
  Early intent, audience, and uncertainty discovery. Traceability gaps may be explicit deferrals, while invalid contracts, invalid decisions, blocker questions, and unmitigated high-severity risks still block.

prototype
  Default planning readiness. Capability, evaluation, risk, artifact or requirement, and non-goal traceability must be present before lock.

mvp
  A usable first-release planning profile. It keeps the same blocking traceability as prototype readiness and treats unresolved non-blocking questions as explicit deferrals.

beta
  A broader validation planning profile. Core traceability blocks lock; non-blocking unresolved questions remain visible as deferrals.

production
  High-confidence planning readiness. Deferred decisions and open questions block lock until resolved.
```

## CLI selection

```bash
ni init --profile concept
ni init --profile prototype
ni init --profile mvp
ni init --profile beta
ni init --profile production
```

`ni status` prints the active readiness profile with the readiness result.
