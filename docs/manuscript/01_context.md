# Context and scope {#sec:context}

Active inference connects perception and action in a generative model and offers
a framework for reasoning about uncertainty and preferences
[@friston2017active; @parr2022active]. The Goference examples represent a
partially observable process with finite state, action, and observation labels.
They are runnable illustrations of selected components of that framework.

The model stores transition and observation tables. The belief updater and
planner operate through Go interfaces. Categorical beliefs drive the demos.
The package also has separate Gaussian and Dirichlet belief utilities for
programmatic use.
