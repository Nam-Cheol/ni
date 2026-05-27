# Why ni exists

AI agents는 점점 더 잘 일한다. 하지만 더 어려운 문제는 여전히 같다. 그 work가
정말 시작되어도 되는지 어떻게 알 수 있는가?

많은 agent failure는 첫 file이 바뀌기 전에 시작된다. User가 충분히 actionable해
보이는 goal을 주고, agent가 그럴듯한 path를 찾고, system은 ambiguity를
implementation으로 바꾸기 시작한다. 빠진 users, 말하지 않은 acceptance
criteria, hidden risks, non-goals, unresolved questions, stale planning context가
모두 work 안으로 들어간다.

`ni`는 바로 그 순간을 위해 만들어졌다.

## Execution problem

Agents는 ambiguous intent를 자주 실행한다. Interface가 그렇게 유도하기 때문이다.
Prompt는 momentum을 요청한다. Model은 helpful하게 행동하도록 보상받는다. User는
work가 이미 움직이기 전까지 어떤 detail이 structurally important한지 모를 수
있다.

더 나은 prompt는 도움이 된다. 하지만 prompt refinement만으로는 충분하지 않다.
잘 다듬어진 prompt도 같은 unresolved intent를 숨길 수 있다. Accepted
constraints, evaluated capabilities, explicit non-goals, mitigated high-severity
risks, agreement 이후 plan이 바뀌었는지에 대한 clear answer가 없어도 decisive해
보일 수 있다.

Failure는 wording만의 문제가 아니다. Control의 문제다.

## Boundary를 앞으로 옮기기

`ni`는 control boundary를 execution 이전으로 옮긴다.

Agent가 일하면서 plan을 추론하게 두는 대신, `ni`는 planning conversation이 먼저
project contract가 되도록 요구한다. Contract는 더 긴 prompt가 아니다. Project가
누구를 위한 것인지, 무엇이 true여야 하는지, 무엇이 out of scope인지, 어떤 risk가
중요한지, 어떤 evidence가 충분한지, 어떤 open question이 handoff를 막는지를 담은
shared record다.

Contract가 ready가 아니면 execution은 시작되지 않아야 한다. 그 refusal 자체가
feature다. Uncertainty가 아직 해결하기 저렴할 때 visible하게 남겨두기 때문이다.

## Conversation에서 handoff까지

Product shape는 의도적으로 작다:

```text
conversation -> contract -> lock -> bounded handoff
```

Conversation은 planning record가 validate될 만큼 explicit해질 때 contract가
된다.

Contract는 accepted plan이 readiness gate를 통과하고 source of truth로 hash될 때
lock이 된다.

Lock은 `ni`가 locked plan에서 short prompt나 inert seed material을 compile할 때
bounded handoff가 된다.

Downstream work는 intent가 trustworthy해진 뒤에만 시작된다. Lock 이후 plan이
바뀌거나 current files가 locked hashes와 더 이상 맞지 않으면, handoff는 오래된
agreement가 아직 유효한 척하지 않고 멈춘다.

## 왜 중요한가

AI work는 충분히 빠르기 때문에 unclear intent의 비용도 빠르게 드러난다. 원하지
않은 scope, confident wrong turns, brittle acceptance, 그리고 무엇이 "really"
requested였는지에 대한 논쟁이 생긴다.

`ni`는 downstream agents에 또 다른 execution loop를 붙여서 더 똑똑하게 만들려는
제품이 아니다. Starting line을 더 안전하게 만든다. 목표는 더 많은 ceremony가
아니다. Yes, no, not yet을 말하기에 더 좋은 순간을 만드는 것이다.

Intent가 아직 ambiguous하다면 `ni`는 그것을 obvious하게 만들어야 한다. Intent가
accepted and locked되면 downstream actors는 trust할 수 있는 bounded artifact를
받는다. Product promise는 이것이다. 먼저 intent를 compile하고, 그 다음 work를
시작하라.

## Related work

Adjacent agent, specification, execution-harness projects와의 comparison은
[Related work](11_RELATED_WORK.md)와
[differentiation map](41_DIFFERENTIATION.md)에 있다. 이 story 뒤의 technical
protocol은 [Intent Lock Protocol](42_INTENT_LOCK_PROTOCOL.md)을 참고하라.
