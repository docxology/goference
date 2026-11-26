# Mathematical foundations

The discrete updater weights the supplied prior by the observation likelihood
for each state, then normalizes. The planner propagates beliefs separately
through the transition table. See [belief updates](../api/belief.md) for the
fallbacks in the current code.

The implemented variational free-energy score subtracts an expected log
likelihood from a KL divergence against the model prior. The expected
free-energy score adds an observation-entropy term to a policy KL term.
[Inference details](../api/inference.md) describe the inputs to these functions.

The current entropy helper can accumulate observation masses without
normalizing the mixture. The EFE API also lacks a current-belief parameter
and starts internal prediction from a uniform belief. Interpret reported
scores as diagnostics from this implementation. The [implementation
note](../manuscript/README.md) gives the literature context and a dated plot.
