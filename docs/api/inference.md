# Inference engine

[`pkg/inference`](../../active-goference/pkg/inference/) supplies
`NewInferenceEngine(tolerance, maxIter)`. Its score methods compute:

- `ComputeVariationalFreeEnergy(pomdp, belief, policy)` from the belief's KL
  divergence against the prior and an expected log-likelihood term.
- `ComputeExpectedFreeEnergy(pomdp, policy)` from observation entropy and
  a policy KL term. It starts internal prediction from a uniform prior.
- `MinimizeFreeEnergy(pomdp, belief)` by scoring candidate single-action
  policies using the sum of VFE and EFE.

The methods are useful for inspecting the example model. See the
[mathematical boundaries](../math/foundations.md) before interpreting the scores.
