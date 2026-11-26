# Architecture and limitations

The [Go module](../active-goference/) separates the POMDP model, belief updates,
inference scores, policy selection, and presentation of run results.

| Package | Role |
| --- | --- |
| [Core](api/core.md) | State, action, and observation labels, transition and observation tables, agent interfaces |
| [Belief](api/belief.md) | Discrete updates and separate Gaussian and Dirichlet utilities |
| [Inference](api/inference.md) | Variational and expected free-energy scores |
| [Planning](api/planning.md) | Greedy rollout and separate softmax selection |
| [Utilities](api/utils.md) | Logging, text reports, and PNG visualizations |

The CLI supplies its own observations through a simple selection helper.
The standalone examples exercise the discrete model in simulation. The
[CLI guide](guides/cli.md) explains the current label-parser limitation. The [open work](TODO.md) tracks model validation,
expected free-energy normalization, belief keys, and logging initialization.
