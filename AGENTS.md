# AGENTS.md — goference

Guidance for agents working in this repository.

## What this repository is

Active Inference in Go: variational free energy belief updating and expected free
energy policy selection for POMDP agents, with categorical, Gaussian, and
Dirichlet belief representations built on gonum. Intended for robotics, game AI,
and cognitive modelling.

**The Go module is in `active-goference/`, not at the repository root.** Every
`go` command runs from that directory.

```
goference/
└── active-goference/     ← module root (module active-goference, go 1.26.0)
    ├── cmd/main.go
    ├── pkg/{belief,core,inference,planning,utils}
    ├── tests/
    ├── examples/
    ├── doc/
    └── output/
```

## Running

From `active-goference/`:

```bash
go run cmd/main.go -example basic -steps 10
go run cmd/main.go -example navigation -steps 20 -session nav_demo
go run cmd/main.go -example exploration -steps 25 -session explore_demo
```

`-session` names a run and keeps its output separable under `output/`.

## Tests

```bash
go test ./tests/... -v
go test ./tests/... -bench=. -benchmem
go test ./tests/... -cover -coverprofile=coverage.out
```

Six test files cover beliefs, inference, planning, core behaviour, utilities, and
benchmarks. There is no CI workflow in this repository, so the test suite runs
only when someone runs it — run it before claiming a change is safe.

## Layout contract

| Package | Responsibility |
|---|---|
| `pkg/core` | Generative model and shared types |
| `pkg/belief` | Categorical, Gaussian, and Dirichlet belief representations |
| `pkg/inference` | Variational free energy updates |
| `pkg/planning` | Expected free energy policy selection |
| `pkg/utils` | Numeric helpers (gonum-backed) |
| `cmd/main.go` | Example runner and flag parsing |

Keep the mathematical core in `pkg/` and the flag parsing in `cmd/`. A new
scenario belongs in `examples/` plus an `-example` value, not as a fourth
concern inside `cmd/main.go`.

## Conventions

- `go.mod` declares `go 1.26.0`. Do not lower the directive to accommodate an
  older toolchain; raise the local toolchain instead.
- Confidence and precision matrices are part of the model, not tuning knobs.
  Changing a default changes every reported result — say which one moved.
- `output/` and `active-goference.log` are run artifacts. Regenerate rather than
  hand-edit them.
- There is no packaging step: this is a module consumed by `go get` or run from
  source, not a released binary.

## Verification

Run `go build ./...` and `go test ./tests/...` from `active-goference/`, and
report the exact commands and their result. Benchmark claims need a matching
`-benchmem` run on the machine you describe; do not cite numbers from `doc/`
without re-running them.
