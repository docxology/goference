# Development setup

Install [Go](https://go.dev/dl/) 1.26 or later. The Go module and its dependency
checksums are in [active-goference](../../active-goference/).

```bash
git clone https://github.com/docxology/goference.git
cd goference/active-goference
go mod verify
go build ./...
go test -count=1 ./...
go vet ./...
```

The standalone files in `examples/` have the `ignore` build constraint.
Recursive module builds and tests skip those files. Run a standalone example
by file as shown in the [examples guide](../examples/README.md).

For package coverage and benchmark runs from the module directory:

```bash
go test -count=1 -coverpkg=./pkg/... -coverprofile=coverage.out ./tests/...
go tool cover -func=coverage.out
go test ./tests/... -run '^$' -bench=. -benchmem
```

Benchmark timings depend on the machine and workload. The example and utility
tests write logs and outputs relative to their working directories. Use a
disposable copy of the module if you need to keep your working tree clean.
