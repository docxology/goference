# Worked example and verification {#sec:example}

The standalone basic demo starts with a uniform discrete belief and records
successive beliefs, observations, actions, and free-energy scores. The archived
line plot ([@fig:basic-belief]) shows how the two belief series diverge over
the simulated steps.

![Line plot of the archived basic demo. Belief in the model state labeled dangerous rises toward one, while belief in safe falls toward zero.](../../output/figures/basic_belief_evolution.png){#fig:basic-belief width=85%}

The timestamp on the archived run is September 22, 2025. The plot is an
example output from that run. The `go test ./tests/...` suite exercises
probability updates, model types, scoring, planning, and utilities. The CLI smoke checks exercise the basic, navigation, and exploration flags.
Their shared label-parser limitation leaves the scenario tables only partly
configured. The plotted data come from the standalone basic demo, which
constructs its own model.
