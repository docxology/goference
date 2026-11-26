# Abstract {#sec:abstract}

Active Goference demonstrates discrete belief updating and policy scoring in Go.
The library separates model tables, observation-based updates, free-energy
scores, and a greedy rollout planner. This note describes the code's simplified
model, relates it to active-inference literature, and identifies the evidence
available from the package tests and example traces. The examples illustrate
numerical behavior in small simulated environments.
