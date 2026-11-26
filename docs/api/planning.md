# Policy planner

[`pkg/planning`](../../active-goference/pkg/planning/) supplies
`NewPolicyPlanner(horizon, temperature, engine)`. `PlanPolicy` picks the
lowest-scoring action at each rollout step and propagates the belief through
that action's transition table. This is a greedy rollout over its configured
horizon.

`SoftmaxActionSelection` offers a separate temperature-based draw.
`PredictNextBelief`, `SelectBestAction`, `HierarchicalPlanning`, and
`MergePolicies` expose prediction, selection, and policy composition.
Score interpretation follows the [inference engine's limits](inference.md).
