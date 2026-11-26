# Core model and agent

[`pkg/core`](../../active-goference/pkg/core/) defines `POMDP`, `Belief`,
and `Policy`. `NewPOMDP(states, actions, observations)` initializes the model.
Populate its tables with `AddTransition`, `AddObservationProb`, and
`AddReward`. Check each returned error. Model setters validate labels, and
callers own the bounds and normalization of the numeric probability rows.

`NewBelief` initializes a uniform distribution over the supplied states.
`Policy.SetAction` and `GetAction` use keys rounded to three decimal places
per belief component, which can merge close beliefs.

`NewActiveInferenceAgent` accepts a model, a belief updater, an inference
engine, and a planner. `Act` passes an empty previous action to the updater.
Use `ActWithAction` when supplying the previous action. Both return the
selected action, updated belief, and policy.
