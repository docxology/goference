# Limits {#sec:limits}

The CLI uses a simple observation-selection helper and hard-coded scenario
tables. Its parser reuses the state labels for actions and observations, so
the tables cannot be set up as intended. The standalone demos supply their
own labels and model tables. Step traces show what each program computed
for that run.

The expected free-energy interface receives a model and a policy without a
current-belief argument. The entropy helper uses a policy lookup at a uniform
belief and can aggregate observation masses without normalization. Model
setters check labels. Input probability bounds and row normalization remain
the caller's responsibility. The policy representation rounds belief keys
before lookup. These boundaries matter when comparing the demo outputs with
an independently specified active-inference model.
