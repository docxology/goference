# 🖥️ Active Goference CLI Guide

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.24.6-blue.svg)](https://golang.org/)
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

### **Custom Environment**
```bash
# Define your own POMDP environment
go run cmd/main.go \
  -states "home,office,park,cafe" \
  -actions "walk,drive,bike,stay" \
  -observations "clear,rainy,traffic,quiet" \
  -steps 30 \
  -session "custom_city"
```

## 📋 **Command Line Flags Reference**

### **Core Configuration Flags**

#### **Environment Definition**
| Flag | Type | Default | Description | Example |
|------|------|---------|-------------|---------|
| `-states` | string | `"s1,s2"` | Comma-separated list of POMDP states | `-states "home,office,park"` |
| `-actions` | string | `"a1,a2"` | Comma-separated list of available actions | `-actions "walk,drive,stay"` |
| `-observations` | string | `"o1,o2"` | Comma-separated list of possible observations | `-observations "clear,rainy,traffic"` |

#### **Execution Control**
| Flag | Type | Default | Description | Example |
|------|------|---------|-------------|---------|
| `-steps` | int | `10` | Number of simulation steps to execute | `-steps 50` |
| `-example` | string | `"basic"` | Pre-configured scenario to run | `-example navigation` |
| `-session` | string | `"default"` | Unique session identifier for output organization | `-session "experiment_1"` |

### **Advanced Configuration Flags**

#### **Planning Parameters**
| Flag | Type | Default | Description | Use Case |
|------|------|---------|-------------|----------|
| `-horizon` | int | `3` | Planning horizon (steps ahead to look) | Strategic vs reactive planning |
| `-temperature` | float | `1.0` | Softmax temperature for action selection | Exploration vs exploitation |
| `-tolerance` | float | `1e-6` | Convergence tolerance for optimization | Precision vs speed trade-off |

#### **Inference Parameters**
| Flag | Type | Default | Description | Use Case |
|------|------|---------|-------------|----------|
| `-inference-tol` | float | `1e-6` | Free energy minimization tolerance | High precision inference |
| `-max-iter` | int | `100` | Maximum iterations for optimization | Prevent infinite loops |

### **Output Control Flags**
| Flag | Type | Default | Description | Purpose |
|------|------|---------|-------------|---------|
| `-verbose` | bool | `false` | Enable detailed console output | Development debugging |
| `-quiet` | bool | `false` | Suppress console output | Batch processing |
| `-output-dir` | string | `"output"` | Base directory for all outputs | Custom output location |

### **Development Flags**
| Flag | Type | Default | Description | Use Case |
|------|------|---------|-------------|----------|
| `-debug` | bool | `false` | Enable debug logging and checks | Development troubleshooting |
| `-profile` | bool | `false` | Enable performance profiling | Performance analysis |
| `-validate` | bool | `false` | Run mathematical validation checks | Testing and verification |

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
go run cmd/main.go -example navigation -steps 50 -horizon 10 -temperature 0.5 -session "strategic_nav"

# Exploratory navigation
go run cmd/main.go -example navigation -steps 30 -temperature 2.0 -session "explore_nav"
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

**Exploration Parameters:**
```bash
# High exploration incentive
go run cmd/main.go -example exploration -steps 40 -temperature 3.0 -horizon 5 -session "high_explore"

# Balanced exploration-exploitation
go run cmd/main.go -example exploration -steps 35 -temperature 1.5 -horizon 8 -session "balanced_explore"
```

## 🔧 **Advanced Usage Patterns**

### **Custom POMDP Configuration**

#### **Large-Scale Environments**
```bash
# Complex urban navigation scenario
go run cmd/main.go \
  -states "home,office,park,cafe,store,school,hospital,library" \
  -actions "walk,bike,bus,car,taxi,stay" \
  -observations "clear,rainy,sunny,cloudy,traffic,quiet,crowded" \
  -steps 100 \
  -horizon 5 \
  -temperature 1.2 \
  -session "urban_navigation_study"
```

#### **Scientific Research Scenarios**
```bash
# Cognitive modeling experiment
go run cmd/main.go \
  -states "low_confidence,medium_confidence,high_confidence" \
  -actions "gather_info,make_decision,postpone" \
  -observations "supporting,contradicting,neutral,unclear" \
  -steps 200 \
  -horizon 3 \
  -temperature 0.8 \
  -inference-tol 1e-8 \
  -session "cognitive_modeling_experiment"
```

#### **Real-time Decision Making**
```bash
# Low-latency financial trading simulation
go run cmd/main.go \
  -states "bull_market,bear_market,sideways_market" \
  -actions "buy,sell,hold" \
  -observations "positive_news,negative_news,neutral_news" \
  -steps 1000 \
  -horizon 1 \
  -temperature 0.1 \
  -quiet \
  -session "trading_bot_simulation"
```

### **Performance Optimization**

#### **High-Precision Computation**
```bash
# Maximum accuracy for research validation
go run cmd/main.go \
  -example navigation \
  -steps 50 \
  -inference-tol 1e-12 \
  -max-iter 1000 \
  -temperature 0.01 \
  -validate \
  -session "high_precision_research"
```

#### **Fast Approximation**
```bash
# Speed-optimized for rapid prototyping
go run cmd/main.go \
  -example basic \
  -steps 100 \
  -inference-tol 1e-3 \
  -max-iter 10 \
  -temperature 2.0 \
  -quiet \
  -session "fast_prototype"
```

### **Batch Processing and Automation**

#### **Parameter Sweep Studies**
```bash
#!/bin/bash
# Temperature parameter sweep for exploration-exploitation analysis
echo "Running temperature parameter sweep..."
for temp in 0.1 0.5 1.0 2.0 5.0; do
  echo "Running with temperature = $temp"
  go run cmd/main.go \
    -example exploration \
    -steps 50 \
    -temperature $temp \
    -session "temp_sweep_${temp}" \
    -quiet
done
echo "Temperature sweep completed."
```

#### **Planning Horizon Analysis**
```bash
#!/bin/bash
# Planning horizon impact study
echo "Analyzing planning horizon effects..."
for horizon in 1 2 3 5 8 13; do
  echo "Testing horizon = $horizon"
  go run cmd/main.go \
    -example navigation \
    -steps 30 \
    -horizon $horizon \
    -session "horizon_study_${horizon}" \
    -quiet
done
echo "Horizon analysis completed."
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
    -session "comparison_${scenario}" \
    -quiet
done
echo "Scenario comparison completed."
```

#### **Scalability Testing**
```bash
#!/bin/bash
# Test performance scaling with problem size
echo "Testing scalability..."
sizes=(5 10 15 20)
for n in "${sizes[@]}"; do
  # Generate states dynamically (simplified example)
  states="s1"
  for ((i=2; i<=n; i++)); do
    states="${states},s${i}"
  done

  echo "Testing with $n states..."
  go run cmd/main.go \
    -states "$states" \
    -actions "a1,a2" \
    -observations "o1,o2" \
    -steps 20 \
    -session "scalability_${n}_states" \
    -quiet
done
echo "Scalability testing completed."
```

## 📁 **Output Organization**

### **Directory Structure**
```
output/
├── logs/                    # 📝 Runtime logging
│   ├── active-goference_*.log     # Main application logs
│   ├── debug_*.log               # Debug-level information
│   ├── errors_*.log              # Error messages only
│   ├── timing_*.log              # Performance timing data
│   └── trace_*.log               # Detailed execution traces
├── visualizations/         # 📊 Text-based visualizations
│   ├── belief_evolution_*.txt    # Belief state progression
│   ├── policy_*.txt              # Policy visualizations
│   ├── free_energy_*.txt         # Energy landscape
│   └── final_belief_*.txt        # Terminal belief state
├── traces/                 # 🔍 Detailed execution traces
│   ├── step_*.txt                # Per-step detailed analysis
│   └── comprehensive_trace_*.txt # Full session trace
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

#### **Trace Logs** (`trace_*.log`)
- Complete execution flow
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
# Enable detailed timing logs
go run cmd/main.go -example navigation -timing -session "performance_test"

# View timing results
cat output/logs/timing_performance_test.log
```

### **Memory Usage Tracking**
```bash
# Monitor memory allocation patterns
go run cmd/main.go -example exploration -benchmark -session "memory_analysis"

# Analyze memory usage in reports
grep "memory" output/reports/comprehensive_report_memory_analysis.txt
```

### **Scalability Analysis**
```bash
#!/bin/bash
# Performance scaling study
echo "Performance scaling analysis..."
for states in 2 4 8 16 32; do
  start_time=$(date +%s%N)

  go run cmd/main.go \
    -states "$(seq -s, 1 $states | sed 's/,/,s/g')" \
    -actions "a1,a2" \
    -observations "o1,o2" \
    -steps 10 \
    -quiet \
    -session "scaling_${states}_states"

  end_time=$(date +%s%N)
  duration=$(( ($end_time - $start_time) / 1000000 )) # Convert to milliseconds
  echo "States: $states, Time: ${duration}ms"
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
# Enable debug logging for troubleshooting
go run cmd/main.go -debug -session "debug_session"

# Check debug logs for detailed information
tail -f output/logs/debug_debug_session.log
```

### **Validation Mode**
```bash
# Run with mathematical validation checks
go run cmd/main.go -validate -session "validation_run"

# Check for validation errors in logs
grep "VALIDATION" output/logs/active-goference_validation_run.log
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
# Run with development configuration
go run cmd/main.go \
  -debug \
  -validate \
  -timing \
  -session "development_run"
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
go run cmd/main.go -example basic -steps 5 -quiet
go run cmd/main.go -example navigation -steps 5 -quiet

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
- Reduce planning horizon (`-horizon 2`)
- Decrease maximum iterations (`-max-iter 50`)
- Use larger convergence tolerance (`-inference-tol 1e-4`)

#### **Slow Convergence**
**Symptoms**: Long execution times, many iterations required
**Solutions**:
- Adjust temperature parameter (`-temperature 1.5`)
- Reduce inference tolerance (`-inference-tol 1e-3`)
- Use shorter planning horizon (`-horizon 3`)

#### **Incorrect Results**
**Symptoms**: Unexpected behavior, non-convergent beliefs
**Solutions**:
- Enable validation (`-validate`)
- Check debug logs (`-debug`)
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
# 1. Enable all debugging features
go run cmd/main.go \
  -debug \
  -validate \
  -timing \
  -example basic \
  -steps 3 \
  -session "debug_workflow"

# 2. Check error logs first
cat output/logs/errors_debug_workflow.log

# 3. Review main logs for execution flow
cat output/logs/active-goference_debug_workflow.log

# 4. Examine debug logs for internal state
cat output/logs/debug_debug_workflow.log

# 5. Check timing for performance bottlenecks
cat output/logs/timing_debug_workflow.log

# 6. Review comprehensive report
cat output/reports/comprehensive_report_debug_workflow.txt
```

## 🔗 **Integration Examples**

### **Research Pipeline Integration**
```bash
#!/bin/bash
# Automated research study pipeline
echo "Starting research pipeline..."

# Parameter combinations for study
temperatures=(0.5 1.0 2.0)
horizons=(1 3 5)

for temp in "${temperatures[@]}"; do
  for horizon in "${horizons[@]}"; do
    echo "Running T=$temp, H=$horizon"

    go run cmd/main.go \
      -example navigation \
      -steps 50 \
      -temperature $temp \
      -horizon $horizon \
      -quiet \
      -session "study_T${temp}_H${horizon}"

    # Extract key metrics for analysis
    grep "final_free_energy" output/reports/comprehensive_report_study_T${temp}_H${horizon}.txt >> results.txt
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
    -validate \
    -session "monitor_${timestamp}" \
    -quiet

  # Check for errors
  if [ -s "output/logs/errors_monitor_${timestamp}.log" ]; then
    echo "ERROR detected at ${timestamp}"
    cat output/logs/errors_monitor_${timestamp}.log
  fi

  # Brief pause between runs
  sleep 60
done
```

### **Docker Integration**
```dockerfile
# Dockerfile for goference CLI
FROM golang:1.21-alpine AS builder

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
