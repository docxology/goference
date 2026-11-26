# Belief representations

[`pkg/belief`](../../active-goference/pkg/belief/) exposes:

- `CategoricalBelief` with `Normalize`, likelihood-vector `Update`,
  `GetProbability`, and entropy in bits.
- `GaussianBelief` with mean and diagonal variance, a componentwise update,
  and differential entropy.
- `DirichletBelief` with concentration counts, count updates, sampling, and mean.

`NewBeliefUpdater` supplies the discrete updater used by the agent.
`UpdateBelief` accepts the prior, observation, previous action, and POMDP.
It copies the prior, multiplies by each state's likelihood for the supplied
observation and action, then normalizes. With no observation row for the action
it uses a unit likelihood for that state. A missing observation in an existing
row has zero likelihood. Transition propagation happens separately in the
planner.
