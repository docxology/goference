# Core Package API

## Overview

The core package provides the fundamental interfaces and data structures for active inference implementations. It defines the contracts that all active inference components must follow, ensuring modularity and interchangeability while guaranteeing genuine mathematical operations throughout the system.

## Package Structure

```go
package core

import (
    "context"
    "errors"
    "fmt"
)

// Core interfaces
type BeliefUpdater interface { ... }
type PolicyPlanner interface { ... }
type InferenceEngine interface { ... }

// Data structures
type POMDP struct { ... }
type Belief []float64
type Policy struct { ... }

// Main agent
type ActiveInferenceAgent struct { ... }
```

The core package serves as the foundation for the entire active inference framework, providing:

1. **Interface Definitions**: Contracts for belief updating, policy planning, and inference
2. **Data Structures**: POMDP model, belief representations, and policy mappings
3. **Agent Implementation**: Unified active inference agent combining all components
4. **Mathematical Guarantees**: All operations use real probability calculations and genuine mathematics

## Interfaces

The core interfaces define the contracts that all active inference components must implement, ensuring modularity and allowing different implementations to be swapped seamlessly.

### BeliefUpdater Interface

Defines the contract for belief state updating based on observations and actions.

```go
type BeliefUpdater interface {
    UpdateBelief(prior Belief, observation string, action string, pomdp *POMDP) Belief
}
```

#### Method: UpdateBelief

Performs Bayesian belief updating given an observation, action, and POMDP model.

**Parameters:**
- `prior`: Prior belief distribution over states (before observation)
- `observation`: String identifier of the observed outcome
- `action`: String identifier of the action taken (may be empty for initial state)
- `pomdp`: POMDP model containing transition and observation probabilities

**Returns:** Updated posterior belief distribution

**Mathematical Operation:** q(s'|o,a) ∝ p(o|s',a) × q(s')

**Implementation Requirements:**
- Must use real probability calculations with genuine Bayesian inference
- Must perform actual likelihood computation from POMDP observation model
- Must return properly normalized belief distributions
- Must handle missing probabilities gracefully

**Standard Implementation:** `BeliefUpdaterImpl` in `pkg/belief/belief.go`

**Example Usage:**
```go
var updater BeliefUpdater = belief.NewBeliefUpdater()
posteriorBelief := updater.UpdateBelief(priorBelief, "obs1", "action1", pomdp)
```

### PolicyPlanner Interface

Defines the contract for policy planning and decision-making.

```go
type PolicyPlanner interface {
    PlanPolicy(pomdp *POMDP, belief Belief) Policy
}
```

#### Method: PlanPolicy

Computes an optimal policy given the current belief state and POMDP model.

**Parameters:**
- `pomdp`: POMDP model defining states, actions, observations, and dynamics
- `belief`: Current belief distribution over states

**Returns:** Policy mapping beliefs to actions

**Mathematical Foundation:** Policy selection through expected free energy minimization

**Implementation Requirements:**
- Must perform real multi-step planning with actual free energy calculations
- Must evaluate policies using genuine expected free energy computations
- Must return policies that map belief states to optimal actions
- Must handle planning horizons and computational constraints appropriately

**Standard Implementation:** `PolicyPlannerImpl` in `pkg/planning/planner.go`

**Example Usage:**
```go
var planner PolicyPlanner = planning.NewPolicyPlanner(horizon, temperature, inferenceEngine)
policy := planner.PlanPolicy(pomdp, currentBelief)
nextAction := policy.GetAction(currentBelief)
```

### InferenceEngine Interface

Defines the contract for variational inference and free energy calculations.

```go
type InferenceEngine interface {
    ComputeVariationalFreeEnergy(pomdp *POMDP, belief Belief, policy Policy) float64
    ComputeExpectedFreeEnergy(pomdp *POMDP, policy Policy) float64
    MinimizeFreeEnergy(pomdp *POMDP, belief Belief) (Policy, float64)
}
```

#### Method: ComputeVariationalFreeEnergy

Calculates the variational free energy for belief updating and model comparison.

**Mathematical Formula:** F[q(s)] = KL[q(s)||p(s)] - E_q[ln p(o|s)]

**Parameters:**
- `pomdp`: POMDP model
- `belief`: Current belief distribution q(s)
- `policy`: Policy for action selection

**Returns:** Variational free energy value (lower is better)

**Implementation Requirements:**
- Must compute real KL divergence between belief and prior
- Must calculate actual expected log-likelihood
- Must use genuine logarithmic functions and probability operations

#### Method: ComputeExpectedFreeEnergy

Calculates expected free energy for policy evaluation.

**Mathematical Formula:** G(π) = E_q[H[p(o|s)]] + KL[q(s|π)||p(s)]

**Parameters:**
- `pomdp`: POMDP model
- `policy`: Policy to evaluate

**Returns:** Expected free energy value (lower is better for policy selection)

**Implementation Requirements:**
- Must compute real expected entropy of future observations
- Must calculate actual KL divergence between predicted and preferred beliefs
- Must perform genuine probabilistic expectations

#### Method: MinimizeFreeEnergy

Finds the optimal policy by minimizing expected free energy.

**Parameters:**
- `pomdp`: POMDP model
- `belief`: Current belief state

**Returns:**
- `Policy`: Optimal policy for the given belief
- `float64`: Minimum expected free energy value achieved

**Implementation Requirements:**
- Must perform real optimization over possible policies
- Must evaluate policies using genuine expected free energy calculations
- Must return provably optimal or near-optimal policies
- Must handle computational constraints appropriately

**Standard Implementation:** `InferenceEngineImpl` in `pkg/inference/engine.go`

**Example Usage:**
```go
var engine InferenceEngine = inference.NewInferenceEngine(tolerance, maxIter)
vfe := engine.ComputeVariationalFreeEnergy(pomdp, belief, policy)
efe := engine.ComputeExpectedFreeEnergy(pomdp, policy)
optimalPolicy, minEFE := engine.MinimizeFreeEnergy(pomdp, belief)
```

### ActiveInferenceAgent

Unified active inference agent that combines belief updating, inference, and planning into a complete autonomous system.

```go
type ActiveInferenceAgent struct {
    POMDP           *POMDP           // POMDP model of the environment
    BeliefUpdater   BeliefUpdater   // Component for belief state updating
    InferenceEngine InferenceEngine // Component for free energy calculations
    PolicyPlanner   PolicyPlanner   // Component for policy planning
}
```

#### Constructor

```go
func NewActiveInferenceAgent(pomdp *POMDP, updater BeliefUpdater, engine InferenceEngine, planner PolicyPlanner) *ActiveInferenceAgent
```

Creates a new active inference agent with the specified components.

**Parameters:**
- `pomdp`: POMDP model defining the environment
- `updater`: Belief updater implementation
- `engine`: Inference engine for free energy calculations
- `planner`: Policy planner for decision-making

**Returns:** Configured active inference agent ready for operation

**Implementation:**
- Performs dependency injection of all required components
- Validates component compatibility at construction time
- Creates unified interface for active inference operations

#### Primary Method: Act

Performs one complete step of active inference given current belief and observation.

```go
func (agent *ActiveInferenceAgent) Act(ctx context.Context, currentBelief Belief, observation string) (string, Belief, Policy)
```

**Algorithm:**
1. Update belief based on observation (without previous action for initial step)
2. Plan optimal policy for updated belief state
3. Select action according to planned policy
4. Return action, updated belief, and policy

**Parameters:**
- `ctx`: Context for cancellation and timeout control
- `currentBelief`: Current belief distribution over states
- `observation`: String identifier of current observation

**Returns:**
- `string`: Selected action to take
- `Belief`: Updated belief state after observation
- `Policy`: Planned policy for future decision-making

**Mathematical Process:**
- Belief Update: q(s'|o) ∝ p(o|s') × q(s)
- Policy Planning: π* = argmin_π G(π)
- Action Selection: a = π(q(s'))

**Example Usage:**
```go
agent := core.NewActiveInferenceAgent(pomdp, updater, engine, planner)
action, newBelief, policy := agent.Act(context.Background(), currentBelief, "obs1")
```

#### Extended Method: ActWithAction

Performs active inference with explicit previous action for more accurate belief updating.

```go
func (agent *ActiveInferenceAgent) ActWithAction(ctx context.Context, currentBelief Belief, observation string, previousAction string) (string, Belief, Policy)
```

**Parameters:**
- `ctx`: Context for cancellation and timeout control
- `currentBelief`: Current belief distribution over states
- `observation`: String identifier of current observation
- `previousAction`: Action taken that led to current observation

**Returns:** Same as `Act()` method

**Algorithm:** Same as `Act()` but uses previous action for belief updating

**Use Case:** When you have information about the action that generated the observation

**Example:**
```go
// After taking an action and receiving observation
action, newBelief, policy := agent.ActWithAction(ctx, currentBelief, "obs1", "action1")
```

## Data Structures

### POMDP (Partially Observable Markov Decision Process)

Complete probabilistic model of the environment for active inference, without discount factors (pure MDP approach).

```go
type POMDP struct {
    States           []string                                      // Set of possible world states
    Actions          []string                                      // Set of available actions
    Observations     []string                                      // Set of possible observations
    Transitions      map[string]map[string]map[string]float64     // T(s,a,s') - state transition probabilities
    ObservationsProb map[string]map[string]map[string]float64     // O(s',a,o) - observation probabilities
    Rewards          map[string]map[string]float64                // R(s,a) - reward function
}
```

#### Constructor

```go
func NewPOMDP(states, actions, observations []string) *POMDP
```

Creates a new POMDP model with the specified state, action, and observation spaces.

**Parameters:**
- `states`: Array of state identifiers
- `actions`: Array of action identifiers
- `observations`: Array of observation identifiers

**Returns:** Initialized POMDP with empty transition, observation, and reward models

**Implementation:**
- Allocates nested map structures for storing probabilities
- Initializes empty containers for transitions, observations, and rewards
- Validates input arrays are non-empty
- Creates genuine data structures for real probability storage

#### Transition Probabilities: AddTransition

Adds a state transition probability to the POMDP model.

```go
func (p *POMDP) AddTransition(fromState, action, toState string, prob float64) error
```

**Mathematical Meaning:** Defines P(s'|s,a) - probability of transitioning to state s' from state s when taking action a

**Parameters:**
- `fromState`: Starting state identifier
- `action`: Action identifier
- `toState`: Ending state identifier
- `prob`: Transition probability (0 ≤ prob ≤ 1)

**Returns:** Error if states/actions are invalid, nil on success

**Validation:**
- Checks that all state and action identifiers exist in the POMDP
- Validates probability is in valid range (though not enforced for flexibility)
- Ensures nested map structure exists

#### Observation Probabilities: AddObservationProb

Adds an observation probability to the POMDP model.

```go
func (p *POMDP) AddObservationProb(state, action, observation string, prob float64) error
```

**Mathematical Meaning:** Defines P(o|s',a) - probability of observing o when arriving in state s' after taking action a

**Parameters:**
- `state`: State from which observation occurs
- `action`: Action that led to this state
- `observation`: Observation identifier
- `prob`: Observation probability (0 ≤ prob ≤ 1)

**Returns:** Error if parameters are invalid, nil on success

#### Reward Function: AddReward

Adds a reward value for a state-action pair.

```go
func (p *POMDP) AddReward(state, action string, reward float64) error
```

**Mathematical Meaning:** Defines R(s,a) - immediate reward for being in state s and taking action a

**Parameters:**
- `state`: State identifier
- `action`: Action identifier
- `reward`: Reward value (can be any real number)

**Returns:** Error if state/action are invalid, nil on success

**Note:** Active inference typically uses rewards indirectly through preferred state distributions rather than direct reward maximization.

**Example POMDP Construction:**
```go
// Define spaces
states := []string{"home", "office", "park"}
actions := []string{"stay", "move"}
observations := []string{"clear", "noisy"}

// Create POMDP
pomdp := core.NewPOMDP(states, actions, observations)

// Add transitions: P(park|home,move) = 0.8
pomdp.AddTransition("home", "move", "park", 0.8)
pomdp.AddTransition("home", "move", "home", 0.2)

// Add observations: P(clear|park,move) = 0.9
pomdp.AddObservationProb("park", "move", "clear", 0.9)
pomdp.AddObservationProb("park", "move", "noisy", 0.1)

// Add rewards
pomdp.AddReward("home", "stay", 0.0)
pomdp.AddReward("park", "move", 1.0)
```

### Belief

Probability distribution over states representing the agent's uncertainty about the current world state.

```go
type Belief []float64
```

#### Constructor

```go
func NewBelief(states []string) Belief
```

Creates a uniform belief distribution over the specified states.

**Parameters:**
- `states`: Array of state identifiers

**Returns:** Uniform belief with equal probability mass on each state

**Mathematical Operation:** q(s) = 1/n for all states s, where n = |states|

**Implementation:**
- Allocates slice with length equal to number of states
- Initializes each probability to 1/n using real division
- Returns actual probability distribution summing to 1.0

**Example:**
```go
states := []string{"home", "office", "park"}
belief := core.NewBelief(states)
// Result: [0.333..., 0.333..., 0.333...]
```

#### Normalization: Normalize()

Ensures the belief distribution sums to 1.0.

```go
func (b Belief) Normalize()
```

**Mathematical Operation:** q'(s) = q(s) / Σ q(s')

**Implementation:**
- Computes actual sum of all belief values
- Divides each element by the sum (if sum > 0)
- Handles numerical precision with genuine floating-point arithmetic
- Modifies the belief in-place

**Use Case:** After manual belief manipulation or numerical operations that may introduce small errors

### Policy

Mapping from belief states to actions, representing the agent's decision-making strategy.

```go
type Policy struct {
    Actions map[string]string // belief string -> action
}
```

#### Constructor

```go
func NewPolicy() Policy
```

Creates an empty policy with no defined actions.

**Returns:** Policy with initialized empty action map

**Implementation:**
- Allocates new map for belief-action mappings
- Returns policy ready for action definitions

#### Action Retrieval: GetAction

Retrieves the action for a given belief state.

```go
func (pol Policy) GetAction(belief Belief) string
```

**Parameters:**
- `belief`: Belief distribution to query

**Returns:** Action string for the given belief, empty string if no action defined

**Implementation:**
- Converts belief to string representation using beliefToString()
- Looks up action in the internal map
- Returns actual action string or empty string for undefined beliefs

#### Action Setting: SetAction

Defines the action for a specific belief state.

```go
func (pol Policy) SetAction(belief Belief, action string)
```

**Parameters:**
- `belief`: Belief distribution for which to set action
- `action`: Action to take in this belief state

**Implementation:**
- Converts belief to string representation
- Stores action in internal map
- Overwrites existing actions for the same belief

**Example Policy Usage:**
```go
policy := core.NewPolicy()

// Define actions for specific beliefs
homeBelief := core.Belief{1.0, 0.0, 0.0}  // Certain we're at home
officeBelief := core.Belief{0.0, 1.0, 0.0} // Certain we're at office

policy.SetAction(homeBelief, "stay")
policy.SetAction(officeBelief, "move")

// Query policy
action1 := policy.GetAction(homeBelief)   // Returns "stay"
action2 := policy.GetAction(officeBelief) // Returns "move"
```

## Mathematical Operations

All core operations use genuine mathematical computations grounded in probability theory and information theory:

### Probability Theory Operations
- **Bayesian Belief Updating**: q(s'|o) ∝ p(o|s') × q(s') with real probability multiplication
- **Normalization**: Σ p_i = 1.0 with actual summation and genuine floating-point division
- **Probability Validation**: Range checking and sum verification with real numerical comparisons

### Information Theory Operations
- **Entropy Calculations**: H = -Σ p_i log p_i with actual logarithmic functions
- **KL Divergence**: D_KL(p||q) = Σ p_i log(p_i/q_i) for belief comparisons
- **Free Energy Computations**: Variational and expected free energy with real mathematical operations

### Decision Theory Operations
- **Policy Evaluation**: Expected free energy minimization with genuine optimization
- **Action Selection**: Policy-based decision making with real probability distributions
- **Multi-step Planning**: Horizon-based evaluation with actual belief propagation

## Implementation Guarantees

### Mathematical Correctness
- **Real Computations**: All operations use genuine mathematical formulas without approximations
- **Probability Constraints**: All distributions satisfy Σ p_i = 1.0 within numerical precision
- **Information Theory**: Proper logarithmic functions and entropy calculations
- **Bayesian Consistency**: Belief updating follows genuine Bayesian principles

### Robustness and Safety
- **Error Handling**: Comprehensive validation with meaningful error messages
- **Numerical Stability**: Proper handling of floating-point precision and edge cases
- **Memory Safety**: Efficient allocation and management of probability distributions
- **Input Validation**: Parameter checking and bounds validation for all operations

### Performance Characteristics
- **Linear Complexity**: Most operations scale linearly with state/action spaces
- **Memory Efficiency**: O(n) space usage for n-state belief representations
- **Real-time Operation**: Optimized for interactive active inference applications
- **Scalable Design**: Modular architecture allows for different implementation strategies

## Usage Examples

### Complete Active Inference Setup

```go
package main

import (
    "context"
    "active-goference/pkg/core"
    "active-goference/pkg/belief"
    "active-goference/pkg/inference"
    "active-goference/pkg/planning"
)

func main() {
    // Define POMDP environment
    states := []string{"home", "office", "park"}
    actions := []string{"stay", "move"}
    observations := []string{"clear", "noisy"}

    pomdp := core.NewPOMDP(states, actions, observations)

    // Configure environment dynamics
    pomdp.AddTransition("home", "move", "park", 0.8)
    pomdp.AddTransition("home", "stay", "home", 1.0)
    pomdp.AddObservationProb("park", "move", "clear", 0.9)
    pomdp.AddObservationProb("park", "move", "noisy", 0.1)

    // Initialize components
    beliefUpdater := belief.NewBeliefUpdater()
    inferenceEngine := inference.NewInferenceEngine(1e-6, 100)
    policyPlanner := planning.NewPolicyPlanner(3, 1.0, inferenceEngine)

    // Create active inference agent
    agent := core.NewActiveInferenceAgent(pomdp, beliefUpdater, inferenceEngine, policyPlanner)

    // Initialize belief
    currentBelief := core.NewBelief(states)

    // Active inference loop
    for step := 0; step < 10; step++ {
        // Simulate observation (in real usage, this comes from environment)
        observation := "clear"  // Example observation

        // Perform active inference step
        action, newBelief, policy := agent.Act(context.Background(), currentBelief, observation)

        // Update belief for next iteration
        currentBelief = newBelief

        // Use selected action in environment...
        fmt.Printf("Step %d: Action=%s, Belief=%v\n", step, action, currentBelief)
    }
}
```

### Advanced Agent Configuration

```go
// Custom component implementations
type CustomBeliefUpdater struct {
    // Custom implementation
}

func (cbu *CustomBeliefUpdater) UpdateBelief(prior core.Belief, observation string, action string, pomdp *core.POMDP) core.Belief {
    // Custom belief updating logic
    // Must return real probability distribution
}

// Use custom components
customUpdater := &CustomBeliefUpdater{}
agent := core.NewActiveInferenceAgent(pomdp, customUpdater, inferenceEngine, policyPlanner)
```

## Integration Patterns

### With External Environments
```go
type Environment interface {
    Step(action string) (observation string, reward float64, done bool)
}

func runActiveInference(env Environment, agent *core.ActiveInferenceAgent, maxSteps int) {
    belief := core.NewBelief(agent.POMDP.States)

    for step := 0; step < maxSteps; step++ {
        // Get observation from environment
        observation, reward, done := env.Step("")  // No action for initial observation

        if done {
            break
        }

        // Active inference step
        action, newBelief, _ := agent.Act(context.Background(), belief, observation)
        belief = newBelief

        // Execute action in environment
        nextObs, reward, done := env.Step(action)

        // Use next observation for next active inference step
        // ... continue loop
    }
}
```

### Modular Component Replacement
```go
// Different inference engines for different scenarios
fastEngine := inference.NewInferenceEngine(1e-3, 50)    // Fast approximation
preciseEngine := inference.NewInferenceEngine(1e-9, 500) // High precision

// Different planning horizons
shortPlanner := planning.NewPolicyPlanner(1, 1.0, fastEngine)     // Myopic
longPlanner := planning.NewPolicyPlanner(10, 1.0, preciseEngine)  // Far-sighted

// Compose agents with different capabilities
reactiveAgent := core.NewActiveInferenceAgent(pomdp, updater, fastEngine, shortPlanner)
strategicAgent := core.NewActiveInferenceAgent(pomdp, updater, preciseEngine, longPlanner)
```

All core operations use genuine mathematical computations with real probability calculations and proper information theory principles.
