# Active Goference API Reference

## Overview

Complete API reference for Active Goference, a production-ready active inference framework implemented in Go. This index provides comprehensive navigation through all API components, organized by package and functionality.

## 🏗️ Architecture Overview

Active Goference follows a modular architecture with clear separation of concerns:

```
active-goference/
├── pkg/core/           # Fundamental interfaces and data structures
├── pkg/belief/         # Belief representation and Bayesian updating
├── pkg/inference/      # Variational inference and free energy calculations
├── pkg/planning/       # Policy planning and decision-making
└── pkg/utils/          # Logging, visualization, and utilities
```

Each package provides genuine mathematical implementations with real numerical computations.

## 📦 Package Index

### Core Package - Foundation Layer
**File:** [`core.md`](core.md) - Fundamental interfaces and data structures

#### Core Interfaces
- **[BeliefUpdater](core.md#beliefupdater)** - Bayesian belief updating contract
- **[PolicyPlanner](core.md#policyplanner)** - Policy planning and decision-making contract
- **[InferenceEngine](core.md#inferenceengine)** - Variational inference contract
- **[ActiveInferenceAgent](core.md#activeinferenceagent)** - Unified active inference agent

#### Core Data Structures
- **[POMDP](core.md#pomdp)** - Partially Observable Markov Decision Process
- **[Belief](core.md#belief)** - Probability distributions over states
- **[Policy](core.md#policy)** - Action mappings from beliefs

#### Constructor Functions
- **[NewPOMDP](core.md#constructor)** - Initialize POMDP environments
- **[NewBelief](core.md#constructor-1)** - Create uniform belief distributions
- **[NewPolicy](core.md#constructor-2)** - Initialize policy mappings
- **[NewActiveInferenceAgent](core.md#constructor-3)** - Construct complete agents

### Belief Package - Bayesian Inference Layer
**File:** [`belief.md`](belief.md) - Belief representation and Bayesian updating

#### Belief Representations
- **[CategoricalBelief](belief.md#categoricalbelief)** - Discrete probability distributions over finite state spaces
- **[GaussianBelief](belief.md#gaussianbelief)** - Continuous multivariate Gaussian distributions with Kalman filtering
- **[DirichletBelief](belief.md#dirichletbelief)** - Parameter learning using Dirichlet distributions

#### Bayesian Updating
- **[BeliefUpdaterImpl](belief.md#beliefupdaterimpl)** - Core implementation of Bayesian belief updating for POMDPs

#### Constructor Functions
- **[NewCategoricalBelief](belief.md#constructor)** - Initialize categorical belief distributions
- **[NewGaussianBelief](belief.md#constructor-1)** - Create Gaussian belief representations
- **[NewDirichletBelief](belief.md#constructor-2)** - Initialize Dirichlet parameter learning
- **[NewBeliefUpdater](belief.md#constructor-3)** - Construct belief updater implementations

### Inference Package - Variational Inference Layer
**File:** [`inference.md`](inference.md) - Variational inference and free energy calculations

#### Inference Engine
- **[InferenceEngineImpl](inference.md#inferenceengineimpl)** - Core variational inference implementation with configurable precision

#### Free Energy Calculations
- **[ComputeVariationalFreeEnergy](inference.md#variational-free-energy)** - VFE: F[q] = KL[q||p] - E_q[ln p(o|s)]
- **[ComputeExpectedFreeEnergy](inference.md#expected-free-energy)** - EFE: G(π) = E_q[H[p(o|s)]] + KL[q(s|π)||p(s)]
- **[MinimizeFreeEnergy](inference.md#free-energy-minimization)** - Policy optimization through free energy minimization

#### Mathematical Operations
- **[computeKLDivergence](inference.md#kl-divergence)** - KL divergence with proper zero handling
- **[computeExpectedLogLikelihood](inference.md#expected-log-likelihood)** - Expected log-likelihood computation
- **[computeExpectedEntropy](inference.md#expected-entropy)** - Expected observation entropy
- **[predictNextBelief](inference.md#belief-prediction)** - Belief state prediction for planning

#### Constructor Functions
- **[NewInferenceEngine](inference.md#constructor)** - Configure tolerance and iteration limits

### Planning Package - Decision-Making Layer
**File:** [`planning.md`](planning.md) - Policy planning and decision-making algorithms

#### Policy Planning Engine
- **[PolicyPlannerImpl](planning.md#policyplannerimpl)** - Core policy planning with configurable horizons and temperature

#### Planning Algorithms
- **[PlanPolicy](planning.md#planpolicy)** - Multi-step policy construction with EFE minimization
- **[selectBestAction](planning.md#selectbestaction)** - Deterministic action selection via EFE minimization
- **[predictNextBelief](planning.md#predictnextbelief)** - Belief state prediction for lookahead planning

#### Advanced Planning Methods
- **[SoftmaxActionSelection](planning.md#softmaxactionselection)** - Probabilistic action selection with temperature control
- **[HierarchicalPlanning](planning.md#hierarchicalplanning)** - Multi-level hierarchical policy planning
- **[mergePolicies](planning.md#mergepolicies)** - Policy composition and conflict resolution

#### Constructor Functions
- **[NewPolicyPlanner](planning.md#constructor)** - Configure planning horizon and exploration parameters

### Utils Package - Infrastructure Layer
**File:** [`utils.md`](utils.md) - Logging, visualization, and utility functions

#### Logging Infrastructure
- **[Logger](utils.md#logger)** - Multi-level logging with organized file output streams
- **[Log](utils.md#global-logger-instance)** - Global logger instance for application-wide use

#### Output Management
- **[OutputManager](utils.md#outputmanager)** - Centralized file output with timestamp organization
- **[Output](utils.md#global-output-manager-instance)** - Global output manager instance

#### Visualization Functions
- **[VisualizeBelief](utils.md#visualizebelief)** - Text-based belief distribution visualization
- **[VisualizePolicy](utils.md#visualizepolicy)** - Policy action mapping visualization
- **[VisualizePOMDP](utils.md#visualizepomdp)** - POMDP structure and parameters overview
- **[CreateBeliefHeatmap](utils.md#createbeliefheatmap)** - Belief evolution heatmaps over time
- **[VisualizeFreeEnergy](utils.md#visualizefreeenergy)** - Free energy analysis and trends
- **[CreateDetailedTrace](utils.md#createdetailedtrace)** - Step-by-step execution traces
- **[CreateComprehensiveReport](utils.md#createcomprehensivereport)** - Session summary reports

#### Mathematical Utilities
- **[beliefToString](utils.md#helper-functions)** - Belief array to string conversion
- **[beliefEntropy](utils.md#beliefentropy)** - Shannon entropy calculation for beliefs
- **[log2](utils.md#log2)** - Base-2 logarithm for information theory
- **[policyKLDivergence](utils.md#policykldivergence)** - Policy comparison using KL divergence

## 🔄 Integration Patterns

### Complete Active Inference Workflow

```go
// 1. Setup environment (Core)
pomdp := core.NewPOMDP(states, actions, observations)
belief := core.NewBelief(pomdp.States)

// 2. Initialize components
updater := belief.NewBeliefUpdater()
engine := inference.NewInferenceEngine(1e-6, 100)
planner := planning.NewPolicyPlanner(3, 1.0, engine)

// 3. Create agent (Core)
agent := core.NewActiveInferenceAgent(pomdp, updater, engine, planner)

// 4. Active inference loop
for step := 0; step < maxSteps; step++ {
    // Get observation (from environment)
    observation := getObservation()

    // Perform active inference (Core)
    action, newBelief, policy := agent.Act(ctx, belief, observation)

    // Log progress (Utils)
    utils.Log.Info("Step", step, "Action:", action)

    // Visualize results (Utils)
    beliefViz := utils.VisualizeBelief(newBelief, pomdp.States)
    utils.Output.SaveTextToFile(beliefViz, fmt.Sprintf("belief_%d", step), "visualizations")

    // Update belief
    belief = newBelief

    // Execute action in environment
    executeAction(action)
}
```

### Component Interchangeability

```go
// Different belief representations
categoricalBelief := belief.NewCategoricalBelief(states)
gaussianBelief := belief.NewGaussianBelief(2)
dirichletBelief := belief.NewDirichletBelief(3, 2.0)

// Different inference engines
preciseEngine := inference.NewInferenceEngine(1e-9, 1000)
fastEngine := inference.NewInferenceEngine(1e-3, 50)

// Different planning strategies
reactivePlanner := planning.NewPolicyPlanner(1, 0.1, engine)     // Myopic
strategicPlanner := planning.NewPolicyPlanner(10, 1.0, engine)   // Far-sighted
exploratoryPlanner := planning.NewPolicyPlanner(5, 3.0, engine)  // Exploratory

// Compose different agent configurations
agents := []core.ActiveInferenceAgent{
    core.NewActiveInferenceAgent(pomdp, belief.NewBeliefUpdater(), preciseEngine, strategicPlanner),
    core.NewActiveInferenceAgent(pomdp, belief.NewBeliefUpdater(), fastEngine, reactivePlanner),
}
```

## 📊 Mathematical Operations Index

### Probability Theory
- **Bayesian Updating**: [BeliefUpdater.UpdateBelief](core.md#beliefupdater), [BeliefUpdaterImpl.UpdateBelief](belief.md#beliefupdaterimpl)
- **Normalization**: [Belief.Normalize](core.md#belief), [CategoricalBelief.Normalize](belief.md#normalize)
- **Entropy**: [CategoricalBelief.Entropy](belief.md#entropy), [GaussianBelief.Entropy](belief.md#entropy-1)

### Information Theory
- **KL Divergence**: [InferenceEngineImpl.computeKLDivergence](inference.md#kl-divergence)
- **Expected Entropy**: [InferenceEngineImpl.computeExpectedEntropy](inference.md#expected-entropy)
- **Free Energy**: [InferenceEngineImpl.ComputeVariationalFreeEnergy](inference.md#variational-free-energy)

### Optimization
- **Policy Selection**: [PolicyPlannerImpl.PlanPolicy](planning.md#planpolicy)
- **Free Energy Minimization**: [InferenceEngineImpl.MinimizeFreeEnergy](inference.md#free-energy-minimization)
- **Softmax Selection**: [PolicyPlannerImpl.SoftmaxActionSelection](planning.md#softmaxactionselection)

## 🧪 Testing & Validation

### Unit Testing
- **Belief Operations**: Test probability normalization, Bayesian updating
- **Free Energy Calculations**: Validate mathematical correctness of VFE/EFE
- **Policy Planning**: Test action selection and belief prediction accuracy

### Integration Testing
- **Complete Agent**: End-to-end active inference workflow validation
- **Component Interaction**: Interface compliance and data flow correctness
- **Performance Benchmarks**: Timing and memory usage validation

### Mathematical Validation
- **Probability Conservation**: Σ p_i = 1.0 within numerical precision
- **Information Theory**: KL ≥ 0, entropy bounds validation
- **Bayesian Consistency**: Proper likelihood × prior × posterior relationships

## 🚀 Quick Start Guide

### Basic Setup
```go
// 1. Define POMDP
pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1", "a2"}, []string{"o1", "o2"})

// 2. Setup components
updater := belief.NewBeliefUpdater()
engine := inference.NewInferenceEngine(1e-6, 100)
planner := planning.NewPolicyPlanner(3, 1.0, engine)

// 3. Create agent
agent := core.NewActiveInferenceAgent(pomdp, updater, engine, planner)

// 4. Initialize belief
belief := core.NewBelief(pomdp.States)

// 5. Run active inference
action, newBelief, policy := agent.Act(context.Background(), belief, "o1")
```

### Advanced Configuration
```go
// Custom belief representations
customBelief := belief.NewGaussianBelief(3)  // 3D continuous belief

// Precision inference
highPrecisionEngine := inference.NewInferenceEngine(1e-9, 1000)

// Long-term planning
strategicPlanner := planning.NewPolicyPlanner(20, 0.5, highPrecisionEngine)

// Specialized agent
specializedAgent := core.NewActiveInferenceAgent(pomdp, updater, highPrecisionEngine, strategicPlanner)
```

## 📚 Related Documentation

- **[Mathematical Foundations](../math/foundations.md)** - Theoretical background and derivations
- **[CLI Guide](../guides/cli.md)** - Command-line interface usage
- **[Setup Guide](../guides/setup.md)** - Development environment configuration
- **[Examples](../examples/README.md)** - Usage examples and scenarios
- **[Overview](../overview.md)** - High-level project overview

All API components use genuine mathematical implementations with real numerical computations and proper probabilistic reasoning.
