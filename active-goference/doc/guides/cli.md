# 🖥️ Active Goference CLI Guide

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26.0-blue.svg)](https://go.dev/)
[![CLI](https://img.shields.io/badge/cli-complete-blue.svg)](../README.md)

**Comprehensive command-line interface for running active inference simulations with genuine mathematical computations and organized output management.**

The Active Goference CLI provides both simple quick-start commands and advanced configuration options for running sophisticated active inference experiments. All mathematical operations use real computations with proper probabilistic reasoning.

## 🎯 **Quick Start**

### **Basic Execution**
```bash
# Run with default 2-state, 2-action, 2-observation scenario
go run cmd/main.go

# Run for 10 steps with default parameters
go run cmd/main.go -steps 10
```

### **Built-in Scenarios**
```bash
# Navigation scenario - movement planning with uncertainty
go run cmd/main.go -example navigation -steps 20 -session "nav_demo"

# Exploration scenario - information gathering and uncertainty reduction
go run cmd/main.go -example exploration -steps 15 -session "explore_demo"

# Basic scenario - fundamental active inference principles
go run cmd/main.go -example basic -steps 25 -session "basic_analysis"
```

### **Custom Run Length and Session**
```bash
# The -example scenario defines the state/action/observation model, so custom
# runs vary the step count and the output session instead of the model itself.
go run cmd/main.go -example basic -steps 30 -session "custom_city"
```

## 📋 **Command Line Flags Reference**

### **Core Configuration Flags**

#### **Environment Definition**
| Flag | Type | Default | Description | Example |
|------|------|---------|-------------|---------|
| `-states` | string | `"s1,s2"` | Comma-separated state labels; the `-example` scenario defines the model | `-states "s1,s2"` |
| `-actions` | string | `"a1,a2"` | Comma-separated action labels; the `-example` scenario defines the model | `-actions "a1,a2"` |
| `-observations` | string | `"o1,o2"` | Comma-separated observation labels; the `-example` scenario defines the model | `-observations "o1,o2"` |

#### **Execution Control**
| Flag | Type | Default | Description | Example |
|------|------|---------|-------------|---------|
| `-steps` | int | `10` | Number of simulation steps to execute | `-steps 50` |
| `-example` | string | `"basic"` | Pre-configured scenario to run | `-example navigation` |
| `-session` | string | `"default"` | Unique session identifier for output organization | `-session "experiment_1"` |

### **No Advanced Flags**

The CLI defines only the six flags above. Planning horizon, softmax temperature,
convergence tolerance, iteration limits, output directory, and logging verbosity
are set in code rather than on the command line:

- `planning.NewPolicyPlanner(horizon, temperature, engine)` - planning horizon and softmax temperature
- `inference.NewInferenceEngine(tolerance, maxIterations)` - inference tolerance and iteration cap
- `utils.Log` - the shared logger (`Log.Debug`, `Log.Info`, `Log.Warn`, `Log.Error`)
- `utils.Output` - the shared output manager that writes `output/{logs,visualizations,reports,traces}`

## 🎮 **Built-in Example Scenarios**

### **Basic Scenario**
Fundamental active inference demonstration with simple 2-state, 2-action, 2-observation dynamics.

```bash
go run cmd/main.go -example basic -steps 15 -session "basic_demo"
```

**Configuration:**
- **States**: s1, s2 (two locations)
- **Actions**: a1, a2 (two possible actions)
- **Observations**: o1, o2 (two observation types)
- **Dynamics**: Simple transition and observation probabilities
- **Purpose**: Demonstrate core active inference principles

### **Navigation Scenario**
Spatial navigation with movement uncertainty and observation noise.

```bash
go run cmd/main.go -example navigation -steps 20 -session "nav_study"
```

**Configuration:**
- **States**: Two spatial locations with realistic movement dynamics
- **Actions**: Forward movement, staying in place
- **Observations**: Clear view, obstructed view with different noise levels
- **Dynamics**: Probabilistic movement with occasional failures
- **Purpose**: Study path planning under uncertainty

**Advanced Navigation Options:**
```bash
# Long-term navigation planning
go run cmd/main.go -example navigation -steps 50 -session "strategic_nav"

# Longer exploratory navigation
go run cmd/main.go -example navigation -steps 30 -session "explore_nav"
```

### **Exploration Scenario**
Information gathering and epistemic behavior demonstration.

```bash
go run cmd/main.go -example exploration -steps 25 -session "epistemic_study"
```

**Configuration:**
- **States**: Known and unknown regions
- **Actions**: Exploration actions, staying actions
- **Observations**: Information-rich vs information-poor observations
- **Dynamics**: Discovery mechanics with information value
- **Purpose**: Demonstrate information-seeking behavior

**Exploration Runs:**
```bash
# Longer exploration run
go run cmd/main.go -example exploration -steps 40 -session "long_explore"

# Shorter exploration run
go run cmd/main.go -example exploration -steps 35 -session "short_explore"
```

Exploration/exploitation balance is a code-level setting: the softmax
temperature passed to `planning.NewPolicyPlanner(horizon, temperature, engine)`
in `cmd/main.go` (default `5, 1.0`).

## 🔧 **Advanced Usage Patterns**

### **Custom POMDP Configuration**

#### **Larger Model Spaces (code level)**

Environments larger than the built-in 2-state scenarios are defined in Go, not
on the command line: build the POMDP with `core.NewPOMDP(states, actions,
observations)`, add transitions/observations/rewards, then hand it to
`core.NewActiveInferenceAgent`. `cmd/main.go`'s `setupBasicPOMDP`,
`setupNavigationPOMDP`, and `setupExplorationPOMDP` are the reference examples.
Longer runs only need `-steps`:

```bash
# Long run of a built-in scenario
go run cmd/main.go -example basic -steps 100 -session "long_run"
```

### **Performance Optimization**

#### **Precision vs. Speed (code level)**

Inference precision is fixed in `cmd/main.go` when the engine and planner are
constructed:

```go
// Maximum accuracy: tight tolerance, many iterations
inferenceEngine := inference.NewInferenceEngine(1e-12, 1000)

// Fast approximation: loose tolerance, few iterations
inferenceEngine := inference.NewInferenceEngine(1e-3, 10)
```

The same file shows the defaults used by the CLI (`1e-6, 100` for inference and
`planning.NewPolicyPlanner(5, 1.0, inferenceEngine)` for planning).

### **Batch Processing and Automation**

#### **Step-Count Sweep Study**
```bash
#!/bin/bash
# Run-length sweep for the exploration scenario
echo "Running step-count sweep..."
for steps in 10 25 50 100; do
  echo "Running with steps = $steps"
  go run cmd/main.go \
    -example exploration \
    -steps $steps \
    -session "steps_sweep_${steps}"
done
echo "Sweep completed."
```

#### **Comparative Scenario Analysis**
```bash
#!/bin/bash
# Compare different built-in scenarios
echo "Comparing active inference scenarios..."
scenarios=("basic" "navigation" "exploration")
for scenario in "${scenarios[@]}"; do
  echo "Running $scenario scenario..."
  go run cmd/main.go \
    -example $scenario \
    -steps 25 \
    -session "comparison_${scenario}"
done
echo "Scenario comparison completed."
```

#### **Scalability Testing**

Problem size is compiled in, not passed on the command line: `-states`,
`-actions`, and `-observations` do not resize the model. To measure scaling with
larger state spaces, add a scenario builder in `cmd/main.go` (mirroring
`setupBasicPOMDP`) and then use the existing benchmarks or a profile:

```bash
# Benchmarks in tests/core_bench_test.go
go test ./tests/... -bench=. -benchmem

# CPU profile of a run
go run -cpuprofile=cpu.prof cmd/main.go -example basic -steps 50
go tool pprof cpu.prof
```

## 📁 **Output Organization**

### **Directory Structure**
```
output/
├── logs/                    # 📝 Runtime logging
│   ├── active-goference_*.log     # Main application logs
│   ├── debug_*.log               # Debug-level information
│   ├── errors_*.log              # Error messages only
│   └── timing_*.log              # Performance timing data
├── visualizations/         # 📊 Text-based visualizations
│   ├── belief_evolution_*.txt    # Belief state progression
│   ├── policy_*.txt              # Policy visualizations
│   ├── free_energy_*.txt         # Energy landscape
│   └── final_belief_*.txt        # Terminal belief state
├── traces/                 # 🔍 Detailed execution traces
│   └── step_*.txt                # Per-step detailed analysis
├── reports/                # 📋 Comprehensive reports
│   └── comprehensive_report_*.txt # Session summary
├── charts/                 # 📈 Chart data (if PNG visualization enabled)
│   ├── belief_evolution_*.png   # Belief evolution charts
│   ├── vfe_curve_*.png          # Variational free energy curves
│   └── efe_curve_*.png          # Expected free energy curves
└── png_simple/             # 🖼️ Simple PNG visualizations
    ├── basic_dashboard_*.png    # Simple dashboard plots
    └── basic_belief_*.png       # Basic belief plots
```

### **Visualization Types**

#### **Belief Evolution**
- **Format**: Text-based probability distributions over time
- **Purpose**: Track belief updates through active inference steps
- **Example**: Shows convergence to true state with uncertainty reduction

#### **Policy Visualization**
- **Format**: Action selection probabilities for different belief states
- **Purpose**: Understand decision-making strategy
- **Example**: Illustrates exploration vs exploitation trade-offs

#### **Free Energy Curves**
- **Format**: VFE and EFE values over planning horizon
- **Purpose**: Monitor optimization progress and convergence
- **Example**: Shows energy minimization during inference

### **Log Categories**

#### **Main Application Logs** (`active-goference_*.log`)
```
2025-09-22 08:39:19 INFO Starting Active Inference Application
2025-09-22 08:39:19 INFO Initial belief: [0.5 0.5]
2025-09-22 08:39:19 INFO Step 0 - Action: move
2025-09-22 08:39:19 INFO Step 1 - Action: stay
2025-09-22 08:39:25 INFO Active Inference session completed
```

#### **Debug Logs** (`debug_*.log`)
- Detailed internal state information
- Mathematical computation intermediate values
- Belief update calculations
- Policy planning decisions

#### **Error Logs** (`errors_*.log`)
- Runtime errors and exceptions
- Validation failures
- File I/O errors
- Mathematical computation errors

#### **Timing Logs** (`timing_*.log`)
```
step_0: 45.2ms
step_1: 42.8ms
step_2: 51.1ms
belief_update: 12.3ms
policy_planning: 28.9ms
inference_computation: 8.7ms
```

#### **Trace Output** (`debug_*.log`, `output/traces/`)
- Complete execution flow (`Logger.Trace` writes into the debug log)
- Function entry/exit points
- Parameter values
- Return values

### **Comprehensive Reports**
- **Session Summary**: Total steps, final belief, performance metrics
- **Mathematical Validation**: Free energy values, probability sums
- **Configuration Details**: All parameters used in simulation
- **Performance Statistics**: Timing breakdowns, memory usage

## 🔍 **Performance Monitoring**

### **Built-in Timing Analysis**
```bash
# Every run logs step timings via Logger.Timing
go run cmd/main.go -example navigation -steps 20 -session "performance_test"

# View timing results
cat "$(ls -t output/logs/timing_*.log | head -1)"
```

### **Memory Usage Tracking**
```bash
# Allocation measurements come from the benchmark suite
go test ./tests/... -bench=. -benchmem

# Session metrics are summarised in the newest report
cat "$(ls -t output/reports/comprehensive_report_*.txt | head -1)"
```

### **Scalability Analysis**
```bash
#!/bin/bash
# Run-length timing study
echo "Run-length timing analysis..."
for steps in 10 25 50 100; do
  start_time=$(date +%s%N)

  go run cmd/main.go \
    -example basic \
    -steps $steps \
    -session "timing_${steps}_steps"

  end_time=$(date +%s%N)
  duration=$(( ($end_time - $start_time) / 1000000 )) # Convert to milliseconds
  echo "Steps: $steps, Time: ${duration}ms"
done
```

## 🚨 **Error Handling**

### **Common Error Patterns**

#### **Configuration Errors**
```
Error: Invalid state transition probability: -0.1
Solution: Ensure all transition probabilities are in [0,1] range
```

#### **Mathematical Computation Errors**
```
Error: NaN detected in free energy calculation
Solution: Check belief state normalization and numerical stability
```

#### **File I/O Errors**
```
Error: Cannot create output directory
Solution: Ensure write permissions for ./output/ directory
```

### **Debug Mode**
```bash
# Every run writes main, debug, error, and timing logs named by timestamp
go run cmd/main.go -example basic -steps 5 -session "debug_session"

# Follow the newest debug log for detailed information
tail -f "$(ls -t output/logs/debug_*.log | head -1)"
```

### **Validation Mode**
```bash
# Mathematical validation lives in the test suite
go test ./tests/... -v
```

## 🧪 **Development and Testing**

### **Testing Commands**
```bash
# Run all tests with coverage
go test ./... -v -cover

# Run specific package tests
go test ./pkg/core -v
go test ./pkg/inference -v
go test ./pkg/belief -v

# Run benchmarks
go test -bench=. ./tests/

# Run integration tests
go test -tags=integration ./tests/
```

### **Development Mode**
```bash
# Development run: logs plus the full test suite
go run cmd/main.go -example basic -steps 5 -session "development_run"
go test ./tests/... -v
```

### **Continuous Integration**
```bash
#!/bin/bash
# CI pipeline script
echo "Running CI pipeline..."

# Build
go build ./cmd/main.go

# Test
go test ./... -coverprofile=coverage.out

# Lint
golangci-lint run

# Run example scenarios
go run cmd/main.go -example basic -steps 5
go run cmd/main.go -example navigation -steps 5

echo "CI pipeline completed successfully"
```

## 🔧 **Implementation Details**

### **CLI Architecture**
- **Flag Parsing**: Standard Go `flag` package with custom validation
- **Configuration Management**: Centralized config struct with validation
- **Error Handling**: Structured error types with context
- **Logging Integration**: Unified logging with multiple output streams
- **Output Management**: Hierarchical output directory structure

### **Mathematical Validation**
- **Probability Normalization**: Automatic checking of probability distributions
- **Free Energy Convergence**: Monitoring of optimization convergence
- **Numerical Stability**: Detection of NaN/inf values in computations
- **Boundary Conditions**: Validation of parameter ranges and constraints

### **Performance Optimizations**
- **Memory Pooling**: Reuse of belief state objects
- **Lazy Evaluation**: Deferred computation of expensive operations
- **Parallel Processing**: Concurrent belief updates where applicable
- **Caching**: Memoization of repeated calculations

## 🐛 **Troubleshooting**

### **Common Issues**

#### **High Memory Usage**
**Symptoms**: Application consumes excessive RAM, slow performance
**Solutions**:
- Reduce the planning horizon in `planning.NewPolicyPlanner(horizon, ...)`
- Decrease the iteration cap in `inference.NewInferenceEngine(tol, maxIter)`
- Use a larger convergence tolerance in `inference.NewInferenceEngine`

#### **Slow Convergence**
**Symptoms**: Long execution times, many iterations required
**Solutions**:
- Adjust the softmax temperature in `planning.NewPolicyPlanner`
- Loosen the inference tolerance in `inference.NewInferenceEngine`

#### **Incorrect Results**
**Symptoms**: Unexpected behavior, non-convergent beliefs
**Solutions**:
- Run the test suite (`go test ./tests/... -v`)
- Check the debug logs in `output/logs/`
- Verify POMDP configuration parameters
- Validate probability distributions

#### **File Permission Errors**
**Symptoms**: Cannot create output files or directories
**Solutions**:
- Ensure write permissions for current directory
- Check available disk space
- Run in directory with proper permissions

### **Debugging Workflow**
```bash
# 1. Run a short session; all four log streams are written every run
go run cmd/main.go \
  -example basic \
  -steps 3 \
  -session "debug_workflow"

# 2. Check error logs first
cat "$(ls -t output/logs/errors_*.log | head -1)"

# 3. Review main logs for execution flow
cat "$(ls -t output/logs/active-goference_*.log | head -1)"

# 4. Examine debug logs for internal state
cat "$(ls -t output/logs/debug_*.log | head -1)"

# 5. Check timing for performance bottlenecks
cat "$(ls -t output/logs/timing_*.log | head -1)"

# 6. Review comprehensive report
cat "$(ls -t output/reports/comprehensive_report_*.txt | head -1)"
```

## 🔗 **Integration Examples**

### **Research Pipeline Integration**
```bash
#!/bin/bash
# Automated research study pipeline
echo "Starting research pipeline..."

# Parameter combinations for study
scenarios=(basic navigation exploration)
step_counts=(25 50)

for scenario in "${scenarios[@]}"; do
  for steps in "${step_counts[@]}"; do
    echo "Running $scenario with $steps steps"

    go run cmd/main.go \
      -example $scenario \
      -steps $steps \
      -session "study_${scenario}_${steps}"

    # Extract key metrics for analysis (report files are timestamped, and the
    # newest one belongs to the run that just finished)
    grep -h "final" "$(ls -t output/reports/comprehensive_report_*.txt | head -1)" >> results.txt
  done
done

echo "Research pipeline completed. Results in results.txt"
```

### **Continuous Monitoring Setup**
```bash
#!/bin/bash
# Continuous monitoring script
echo "Starting continuous monitoring..."

while true; do
  timestamp=$(date +%Y%m%d_%H%M%S)

  go run cmd/main.go \
    -example exploration \
    -steps 10 \
    -session "monitor_${timestamp}"

  # Check the newest error log
  error_log=$(ls -t output/logs/errors_*.log | head -1)
  if [ -s "${error_log}" ]; then
    echo "ERROR detected at ${timestamp}"
    cat "${error_log}"
  fi

  # Brief pause between runs
  sleep 60
done
```

### **Docker Integration**
```dockerfile
# Dockerfile for goference CLI
FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o goference ./cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/goference .

# Default command with common parameters
CMD ["./goference", "-example", "basic", "-steps", "20", "-session", "docker_run"]
```

```bash
# Run in Docker with custom parameters
docker run --rm -v $(pwd)/output:/app/output goference \
  -example navigation \
  -steps 100 \
  -session "docker_analysis"
```

---

## 📚 **Additional Resources**

- **[API Documentation](../api/)**: Complete API reference for all packages
- **[Mathematical Foundations](../math/foundations.md)**: Theoretical background and derivations
- **[Setup Guide](setup.md)**: Development environment configuration
- **[Examples](../examples/)**: Code examples for common use cases
- **[Overview](../overview.md)**: High-level project overview

---

*For questions or issues, please check the troubleshooting section above or refer to the [main documentation](../README.md).*

**File**: `timing_{timestamp}.log`
```
INFO: Step 1 inference completed in 12.5ms
INFO: Step 2 inference completed in 11.8ms
INFO: Belief update completed in 3.2ms
INFO: Policy planning completed in 8.9ms
```

## Advanced Usage

### Custom POMDP Configuration

Real programmatic POMDP setup through example functions:

```bash
# Run with custom states but use basic example setup
go run cmd/main.go \
  -states "location1,location2,location3,location4" \
  -actions "north,south,east,west,stay" \
  -observations "landmark_a,landmark_b,landmark_c,nothing" \
  -example basic \
  -steps 30 \
  -session "custom_4x5x4"
```

### Long-Running Simulations

Extended simulations with real progress tracking:

```bash
go run cmd/main.go \
  -example exploration \
  -steps 100 \
  -session "long_exploration_study"
```

### Comparative Analysis

Multiple runs with different configurations:

```bash
# Run basic with different step counts
for steps in 10 25 50 100; do
  go run cmd/main.go \
    -example basic \
    -steps $steps \
    -session "basic_steps_${steps}"
done
```

## Error Handling

### Invalid Configurations

Real error handling with actual validation:

```
ERROR: Invalid state count: states must be non-empty
ERROR: Invalid action specification: actions cannot be empty
ERROR: Invalid observation model: observations must match state space
```

### File System Issues

Real file system error handling:

```
ERROR: Failed to create output directory: permission denied
ERROR: Failed to write log file: disk full
ERROR: Failed to save visualization: I/O error
```

### Numerical Issues

Real numerical validation and error reporting:

```
WARN: Numerical precision warning: probabilities sum to 0.999999 instead of 1.0
WARN: Matrix inversion warning: condition number is high
ERROR: Free energy calculation failed: invalid probability distribution
```

## Development and Testing

### Test Execution

Real test execution with actual validation:

```bash
go test ./tests/ -v
```

### Benchmark Execution

Real performance benchmarking:

```bash
go test ./tests/ -bench=. -benchmem
```

### Coverage Analysis

Real test coverage measurement:

```bash
go test ./tests/ -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## Implementation Details

### Real Mathematical Operations

All CLI executions use genuine computations:

- **Probability calculations**: Real floating-point arithmetic
- **Belief updating**: Actual Bayesian inference operations
- **Free energy minimization**: Genuine variational inference
- **Policy planning**: Real expected free energy optimization

### Memory Management

Actual memory allocation with real data structures:

- **Belief arrays**: Real probability vectors with proper sizing
- **Transition matrices**: Actual sparse matrix representations
- **Policy storage**: Genuine action mapping structures
- **Output buffers**: Real file I/O with proper buffering

### Concurrent Execution

Real concurrent file operations with proper synchronization:

- **Log file writing**: Actual thread-safe logging operations
- **Visualization generation**: Genuine concurrent output generation
- **Report creation**: Real concurrent data aggregation
- **Performance monitoring**: Actual timing with proper concurrency

## Troubleshooting

### Common Issues

**High memory usage**:
- Reduce state space size
- Use smaller planning horizons
- Check for memory leaks in custom configurations

**Slow performance**:
- Decrease number of states/actions
- Reduce planning horizon
- Profile with `go tool pprof`

**Numerical instability**:
- Check probability normalization
- Validate transition matrices
- Verify observation model consistency

**Output file issues**:
- Check write permissions
- Ensure sufficient disk space
- Verify output directory creation

### Debug Mode

Enable detailed debugging output:

```bash
# Debug logs contain real internal state information
# Check output/logs/debug_{timestamp}.log for detailed execution traces
```

### Performance Profiling

Real performance analysis:

```bash
go run -cpuprofile=cpu.prof cmd/main.go -example basic -steps 50
go tool pprof cpu.prof
```

## Integration

### Script Integration

Real shell script integration for batch processing:

```bash
#!/bin/bash
# Batch processing script
for example in basic navigation exploration; do
  for steps in 10 25 50; do
    go run cmd/main.go \
      -example $example \
      -steps $steps \
      -session "${example}_${steps}"
  done
done
```

### Configuration Files

Future extension for configuration file support:

```json
{
  "states": ["s1", "s2", "s3"],
  "actions": ["a1", "a2"],
  "observations": ["o1", "o2"],
  "transitions": {
    "s1": {"a1": {"s2": 0.8, "s1": 0.2}}
  },
  "simulation": {
    "steps": 20,
    "session": "config_based_run"
  }
}
```

All CLI operations use real mathematical computations with genuine data processing and actual file system interactions.
