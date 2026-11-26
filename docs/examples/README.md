# Simulated examples

From the `active-goference/` module, invoke the standalone demo files individually:

```bash
go run examples/basic_example.go
go run examples/navigation_example.go
go run examples/exploration_example.go
go run comprehensive_demo.go
```

The basic demo illustrates belief updating with a small discrete state space.
The navigation demo uses simulated observations and transition probabilities.
The exploration demo supplies a different likelihood and reward table. The
[scoring limits](../math/foundations.md) apply to interpreting those runs.

The standalone demos write text and may produce PNG visualizations. The
[CLI scenarios](../guides/cli.md) use a separate runner with a known label-parser
limitation.
The [retained basic plot](../../active-goference/output/png_simple/basic_belief_evolution_2025-09-22_08-39-19.png)
and [dated step traces](../../active-goference/output/traces/) are example
artifacts. New runs use timestamped output paths. Observe the action and belief
sequence in the traces, and treat timings as local measurements.
