# Active Goference

Active Goference is an early-stage Go implementation of discrete active
inference for small, simulated partially observable environments. It provides
categorical belief updates, free-energy scores, and a greedy policy planner.
Gaussian and Dirichlet belief utilities are also available.

## Quick start

Install Go 1.26 or later and run from the Go module directory:

```bash
cd active-goference
go build ./...
go test ./...
go run examples/basic_example.go
```

The standalone basic example writes PNG figures. The separate
[CLI guide](docs/guides/cli.md) explains the limits of the CLI scenario flags.

Start at the [documentation index](docs/README.md) for setup, API guides, and
model limitations. Read the [implementation note](docs/manuscript/README.md)
for the mathematical context, a dated example plot, and the pinned render recipe.

Use [CITATION.cff](CITATION.cff) to cite the software. The project is provided
under the [MIT License](LICENSE).
