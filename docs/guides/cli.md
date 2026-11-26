# CLI guide

From `active-goference/`, list the flags or run a simulated scenario:

```bash
go run ./cmd -h
go run ./cmd -example basic -steps 5 -session basic_demo
go run ./cmd -example navigation -steps 5 -session nav_demo
go run ./cmd -example exploration -steps 5 -session explore_demo
```

`-example` selects the supplied model, `-steps` controls loop length, and
`-session` labels the report. The CLI uses shared, timestamped output directories.
Use separate working directories to isolate concurrent runs.

The CLI currently parses every nonempty state, action, and observation flag as
`s1,s2`. The scenario tables expect different action and observation labels,
so their model setup is incomplete. The bundled CLI runs exercise argument
handling and reporting. Use the [standalone examples](../examples/README.md)
to inspect the scenario models with their own labels, or use `core.NewPOMDP`
in Go. The [open model-configuration issue](../TODO.md) records the parser work.

Runs write logs, step traces, text visualizations, and a report under `output/`
relative to the module working directory. See the [examples guide](../examples/README.md)
for the separate PNG-producing demos.
