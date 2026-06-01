# Domain and state model

## Core entities

```text
research_protocol
study_decision
research_question
study_artifact
participant_scope
observation_scope
location
consent_boundary
privacy_boundary
data_handling_rule
accessibility_requirement
field_safety_rule
weather_stop_rule
vulnerable_group_safeguard
review_owner
acceptance_evidence
open_question
lockfile
prompt
```

## Resolved fixture state

The benchmark is a research-protocol planning fixture for fictional Riverside
East and Oak Market corridors. Accepted observation units are public street
segments, public plazas, transit-adjacent waiting areas, and publicly
accessible pedestrian corridors.

Optional participant interaction is limited to adult community members who
voluntarily give non-identifying feedback. The fixture excludes minors,
health-status collection, targeted vulnerable-group recruitment, private
property, schools, care facilities, indoor spaces, locations requiring special
access, unsafe locations, and interactions involving minors.

Segment selection is limited to up to 12 public segments using a mix of shade
deficit, pedestrian activity, and planning relevance. Each segment must record
why it was included.

Raw field notes are stored in the project workspace for up to 30 days, then
reduced to a summarized planning memo. The final memo keeps only aggregated,
non-identifying evidence.
