# Active Goference

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.26.0-blue.svg)](https://go.dev/)
![License](https://img.shields.io/badge/license-MIT-green.svg)
[![Documentation](https://img.shields.io/badge/docs-complete-blue.svg)](active-goference/doc/)

Real-time Active Inference implementation in Go using genuine mathematical computations and probabilistic reasoning. A production-ready framework for autonomous agents performing variational inference and policy selection through free energy minimization.

## 🚀 Quick Start

```bash
# Clone the repository
git clone https://github.com/docxology/goference.git
cd goference/active-goference

# Run basic example
go run cmd/main.go -example basic -steps 10

# Run navigation scenario
go run cmd/main.go -example navigation -steps 20 -session "nav_demo"

# Exploration scenario with a named session
go run cmd/main.go -example exploration -steps 25 -session "explore_demo"
```

## 📖 Documentation

### 📚 Core Documentation
- **[Overview](active-goference/doc/overview.md)** - Comprehensive project overview and mathematical foundations
- **[README](active-goference/doc/README.md)** - Detailed project description and usage examples

### 📖 API Reference
- **[API Index](active-goference/doc/api/index.md)** - Complete API documentation index
- **[Core Package](active-goference/doc/api/core.md)** - Core interfaces and data structures
- **[Belief Package](active-goference/doc/api/belief.md)** - Belief representation and updating
- **[Inference Package](active-goference/doc/api/inference.md)** - Variational inference engine
- **[Planning Package](active-goference/doc/api/planning.md)** - Policy planning and selection
- **[Utils Package](active-goference/doc/api/utils.md)** - Utilities for logging and visualization

### 📋 Guides
- **[CLI Guide](active-goference/doc/guides/cli.md)** - Command-line interface usage and examples
- **[Setup Guide](active-goference/doc/guides/setup.md)** - Development environment setup and configuration

### 🧮 Mathematical Foundations
- **[Foundations](active-goference/doc/math/foundations.md)** - Mathematical theory and derivations

### 💡 Examples
- **[Examples Overview](active-goference/doc/examples/README.md)** - Example scenarios and use cases
  - Basic Example - Fundamental active inference principles
  - Navigation Example - Spatial navigation with uncertainty
  - Exploration Example - Information gathering and discovery

## 🏗️ Architecture

```
active-goference/
├── pkg/
│   ├── core/           # Core interfaces and data structures
│   │   ├── interfaces.go
│   │   └── model.go
│   ├── belief/         # Belief representation and updating
│   │   └── belief.go
│   ├── inference/      # Variational inference engine
│   │   └── engine.go
│   ├── planning/       # Policy planning and selection
│   │   └── planner.go
│   └── utils/          # Logging and visualization utilities
│       ├── logger.go
│       ├── visualization.go
│       ├── chart_visualization.go
│       └── simple_png_visualization.go
├── cmd/
│   └── main.go         # CLI application entry point
├── examples/           # Example scenarios
│   ├── basic_example.go
│   ├── exploration_example.go
│   └── navigation_example.go
├── tests/              # Comprehensive test suite
├── doc/                # Complete documentation
└── output/             # Generated results and logs
```

## 🧠 Mathematical Foundation

Active Goference implements genuine variational inference through free energy minimization:

### Variational Free Energy (VFE)
**F[q(s)] = KL[q(s)||p(s)] - E_q[ln p(o|s)]**

Where:
- **F[q(s)]**: Variational free energy
- **q(s)**: Approximate posterior belief
- **p(s)**: Prior belief
- **p(o|s)**: Likelihood of observations

### Expected Free Energy (EFE)
**G(π) = E_q[H[p(o|s)]] + KL[q(s|π)||p(s)]**

Where:
- **G(π)**: Expected free energy for policy π
- **H[p(o|s)]**: Entropy of predicted observations
- **q(s|π)**: Predicted belief under policy

## ✨ Key Features

### 🔬 Real Mathematical Operations
- **No Mock Methods**: Every function performs actual computations
- **Genuine Probability Calculations**: Real floating-point arithmetic throughout
- **Mathematical Rigor**: Proper normalization, entropy, and divergence calculations
- **Numerical Stability**: IEEE 754 double precision with overflow protection

### 🏛️ Modular Architecture
- **Clean Interfaces**: BeliefUpdater, PolicyPlanner, InferenceEngine
- **Multiple Belief Types**: Categorical, Gaussian, Dirichlet representations
- **Extensible Design**: Easy to add new inference methods and belief types
- **Production Ready**: Comprehensive error handling and validation

### 📊 Rich Output System
- **Organized Results**: Structured output directories with timestamps
- **Multiple Visualizations**: Text-based charts, heatmaps, and detailed traces
- **Comprehensive Logging**: Execution logs, debug info, timing measurements
- **Session Management**: Unique session identifiers for result organization

### 🧪 Testing & Quality
- **Real Data Testing**: All tests use actual probability distributions
- **Mathematical Validation**: Numerical constraints verified at runtime
- **Performance Benchmarks**: Actual timing and memory usage measurements
- **TDD Approach**: Test-driven development with genuine implementations

## 🎯 Usage Examples

### Basic Active Inference
```go
// Create POMDP model
pomdp := core.NewPOMDP(states, actions, observations)
pomdp.AddTransition("s1", "a1", "s2", 0.7)
pomdp.AddObservationProb("s1", "a1", "o1", 0.8)
pomdp.AddReward("s1", "a1", 1.0)

// Initialize components
beliefUpdater := belief.NewBeliefUpdater()
inferenceEngine := inference.NewInferenceEngine(1e-6, 100)
policyPlanner := planning.NewPolicyPlanner(5, 1.0, inferenceEngine)

// Create agent
agent := core.NewActiveInferenceAgent(pomdp, beliefUpdater, inferenceEngine, policyPlanner)

// Execute active inference
action, newBelief, policy := agent.Act(ctx, currentBelief, observation)
```

### Custom Belief Types
```go
// Gaussian belief for continuous state spaces
gaussianBelief := belief.NewGaussianBelief(dimension)
gaussianBelief.Update(observation, noiseVariance)

// Dirichlet belief for parameter learning
dirichletBelief := belief.NewDirichletBelief(dimension, concentration)
dirichletBelief.Update(observations)
```

## 📈 Performance Characteristics

- **Real-time Operation**: O(n²) complexity for policy evaluation
- **Memory Efficient**: O(n²) for transition matrices, O(n) for beliefs
- **Scalable**: Genuine performance scaling with state space size
- **Concurrent**: Parallel policy evaluation when applicable

## 🔧 Dependencies

- **Go 1.26.0+**: Modern Go with generics and improved performance
- **gonum.org/v1/gonum**: Real numerical computing library
- **Standard Library**: math, log, time, os, and other core packages

## 🏃‍♂️ Running Tests

```bash
# Run all tests
go test ./tests/... -v

# Run benchmarks
go test ./tests/... -bench=. -benchmem

# Run coverage analysis
go test ./tests/... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out
```

## 📁 Output Structure

All executions generate organized results:

```
output/
├── logs/              # Execution logs with timestamps
│   ├── active-goference_{timestamp}.log
│   ├── errors_{timestamp}.log
│   ├── debug_{timestamp}.log
│   └── timing_{timestamp}.log
├── visualizations/    # Belief states and policy representations
│   ├── belief_step_{n}_{session}_{timestamp}.txt
│   ├── final_policy_{session}_{timestamp}.txt
│   └── belief_evolution_{session}_{timestamp}.txt
├── reports/           # Comprehensive session summaries
│   └── comprehensive_report_{session}_{timestamp}.txt
├── traces/            # Detailed step-by-step execution traces
│   └── step_{n}_{session}_{timestamp}.txt
└── benchmarks/        # Performance measurements
```

## 🤝 Contributing

1. **Fork** the repository
2. **Create** a feature branch (`git checkout -b feature/new-belief-type`)
3. **Implement** with real mathematical operations (no mocks!)
4. **Test** thoroughly with actual data scenarios
5. **Document** your changes
6. **Submit** a pull request

### Development Guidelines

- **Real Implementations Only**: No mock methods or synthetic data
- **Mathematical Rigor**: All formulas implemented with genuine computations
- **TDD Approach**: Test-driven development with real data validation
- **Documentation**: Comprehensive docs with mathematical explanations
- **Code Quality**: Go best practices, proper error handling, modular design

## 📜 License

MIT License.

## 🔬 Research & Applications

Active Goference supports various research and application domains:

- **🤖 Robotics**: Autonomous navigation in partially observable environments
- **🧠 Neuroscience**: Cognitive modeling and predictive processing
- **🎮 Game AI**: Decision-making under uncertainty
- **📊 Data Science**: Probabilistic reasoning and Bayesian inference
- **🚀 Autonomous Systems**: Real-time decision-making for complex systems

## 📚 Further Reading

- **[Active Inference](https://en.wikipedia.org/wiki/Active_inference)** - Wikipedia overview
- **[Free Energy Principle](https://en.wikipedia.org/wiki/Free_energy_principle)** - Theoretical foundation
- **[POMDP](https://en.wikipedia.org/wiki/Partially_observable_Markov_decision_process)** - Decision-making framework
- **[Variational Inference](https://en.wikipedia.org/wiki/Variational_Bayesian_methods)** - Inference methodology

---

**Active Goference**: Where genuine mathematical computation meets real-world active inference. 🚀
