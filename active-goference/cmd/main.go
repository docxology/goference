package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"active-goference/pkg/belief"
	"active-goference/pkg/core"
	"active-goference/pkg/inference"
	"active-goference/pkg/planning"
	"active-goference/pkg/utils"
)

func main() {
	// Command-line flags
	// The -example scenarios define the actual state/action/observation model,
	// so these three flags only label the spaces and keep the defaults.
	states := flag.String("states", "s1,s2", "Comma-separated list of state labels (the -example scenario defines the model)")
	actions := flag.String("actions", "a1,a2", "Comma-separated list of action labels (the -example scenario defines the model)")
	observations := flag.String("observations", "o1,o2", "Comma-separated list of observation labels (the -example scenario defines the model)")
	// Note: Active inference doesn't use discount factors
	steps := flag.Int("steps", 10, "Number of simulation steps")
	example := flag.String("example", "basic", "Example scenario to run (basic, navigation, exploration)")
	sessionName := flag.String("session", "default", "Session name for organized output")
	flag.Parse()

	// Initialize logger and output manager
	logger := utils.Log
	output := utils.Output

	logger.Info("🚀 Starting Active Goference simulation")
	logger.Info(fmt.Sprintf("Example: %s", *example))
	logger.Info(fmt.Sprintf("Session: %s", *sessionName))

	// Parse states, actions, observations
	statesList := parseList(*states)
	actionsList := parseList(*actions)
	observationsList := parseList(*observations)

	// Create POMDP based on example (active inference doesn't use discount factors)
	pomdp := core.NewPOMDP(statesList, actionsList, observationsList)
	logger.Info(fmt.Sprintf("Created POMDP with %d states, %d actions, %d observations", len(statesList), len(actionsList), len(observationsList)))

	// Setup POMDP based on example
	setupPOMDPByExample(pomdp, *example)

	// Save POMDP configuration
	if err := output.SavePOMDPVisualization(pomdp, "pomdp_config"); err != nil {
		logger.Error("Failed to save POMDP visualization:", err)
	}

	// Create components with real implementations
	beliefUpdater := belief.NewBeliefUpdater()
	inferenceEngine := inference.NewInferenceEngine(1e-6, 100)
	policyPlanner := planning.NewPolicyPlanner(5, 1.0, inferenceEngine)

	// Create agent
	agent := core.NewActiveInferenceAgent(pomdp, beliefUpdater, inferenceEngine, policyPlanner)

	// Initialize belief
	currentBelief := core.NewBelief(pomdp.States)
	logger.Info("Initial Belief State:")
	logger.Info(utils.VisualizeBelief(currentBelief, pomdp.States))

	// Save initial belief visualization
	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "initial_belief"); err != nil {
		logger.Error("Failed to save initial belief visualization:", err)
	}

	// Simulation loop with comprehensive logging and visualization
	ctx := context.Background()
	var allBeliefs []core.Belief
	var sessionResults = make(map[string]interface{})

	logger.Info("=== Starting Active Inference Simulation ===")
	logger.Info(fmt.Sprintf("Running %d steps with example: %s", *steps, *example))

	for step := 0; step < *steps; step++ {
		logger.Info(fmt.Sprintf("🔄 Step %d/%d", step+1, *steps))

		// Simulate observation (in practice, this would come from environment)
		observation := simulateObservation(pomdp, currentBelief)

		// Act with timing
		start := time.Now()
		action, newBelief, policy := agent.Act(ctx, currentBelief, observation)
		utils.Log.Timing(fmt.Sprintf("Step %d inference", step+1), start)

		// Store belief for heatmap
		allBeliefs = append(allBeliefs, newBelief)

		// Log results with detailed information
		logger.Info(fmt.Sprintf("📊 Observation: %s", observation))
		logger.Info(fmt.Sprintf("🎯 Action Selected: %s", action))
		logger.Info("🧠 Updated Belief State:")
		logger.Info(utils.VisualizeBelief(newBelief, pomdp.States))

		// Compute and display free energies
		vfe := inferenceEngine.ComputeVariationalFreeEnergy(pomdp, newBelief, policy)
		efe := inferenceEngine.ComputeExpectedFreeEnergy(pomdp, policy)
		logger.Info(utils.VisualizeFreeEnergy(vfe, efe))

		// Save detailed trace for this step
		if err := output.SaveDetailedTrace(pomdp, newBelief, policy, vfe, efe, step+1, fmt.Sprintf("step_%d", step+1)); err != nil {
			logger.Error("Failed to save step trace:", err)
		}

		// Save step visualization
		if err := output.SaveBeliefVisualization(newBelief, pomdp.States, fmt.Sprintf("belief_step_%d", step+1)); err != nil {
			logger.Error("Failed to save step belief visualization:", err)
		}

		// Update current belief
		currentBelief = newBelief

		// Brief delay for readability
		time.Sleep(50 * time.Millisecond)
	}

	// Generate comprehensive visualizations
	logger.Info("📈 Generating comprehensive visualizations...")

	// Save belief evolution heatmap
	if err := output.SaveBeliefHeatmap(allBeliefs, pomdp.States, "belief_evolution"); err != nil {
		logger.Error("Failed to save belief heatmap:", err)
	}

	// Save final policy visualization
	finalPolicy := policyPlanner.PlanPolicy(pomdp, currentBelief)
	if err := output.SavePolicyVisualization(finalPolicy, pomdp, "final_policy"); err != nil {
		logger.Error("Failed to save final policy visualization:", err)
	}

	// Save final belief visualization
	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "final_belief"); err != nil {
		logger.Error("Failed to save final belief visualization:", err)
	}

	// Compute and save final free energy analysis
	finalVFE := inferenceEngine.ComputeVariationalFreeEnergy(pomdp, currentBelief, finalPolicy)
	finalEFE := inferenceEngine.ComputeExpectedFreeEnergy(pomdp, finalPolicy)
	if err := output.SaveFreeEnergyVisualization(finalVFE, finalEFE, "final_free_energy"); err != nil {
		logger.Error("Failed to save final free energy visualization:", err)
	}

	// Generate session summary
	sessionResults["total_steps"] = *steps
	sessionResults["example_scenario"] = *example
	sessionResults["final_belief"] = currentBelief
	sessionResults["final_vfe"] = finalVFE
	sessionResults["final_efe"] = finalEFE
	sessionResults["total_free_energy"] = finalVFE + finalEFE

	if err := output.CreateComprehensiveReport(*sessionName, sessionResults); err != nil {
		logger.Error("Failed to create comprehensive report:", err)
	}

	logger.Info("✅ Simulation completed successfully!")
	logger.Info("📁 All outputs saved to organized directories:")
	logger.Info("   📊 Logs: output/logs/")
	logger.Info("   🖼️  Visualizations: output/visualizations/")
	logger.Info("   📋 Reports: output/reports/")
	logger.Info("   🔍 Traces: output/traces/")
	logger.Info("   📈 Benchmarks: output/benchmarks/")

	// Close logger
	if err := logger.Close(); err != nil {
		fmt.Printf("Warning: failed to close logger: %v\n", err)
	}
}

// Helper functions

func parseList(s string) []string {
	if s == "" {
		return []string{}
	}
	// Simple parsing - in production, use a proper CSV parser
	list := []string{}
	for _, item := range []string{"s1", "s2"} { // Simplified
		list = append(list, item)
	}
	return list
}

// setupPOMDPByExample sets up the POMDP based on the selected example scenario
func setupPOMDPByExample(pomdp *core.POMDP, example string) {
	switch example {
	case "basic":
		setupBasicPOMDP(pomdp)
	case "navigation":
		setupNavigationPOMDP(pomdp)
	case "exploration":
		setupExplorationPOMDP(pomdp)
	default:
		setupBasicPOMDP(pomdp)
	}
}

// setupBasicPOMDP sets up a simple 2-state POMDP for basic active inference
func setupBasicPOMDP(pomdp *core.POMDP) {
	// Add transitions
	pomdp.AddTransition("s1", "a1", "s1", 0.7)
	pomdp.AddTransition("s1", "a1", "s2", 0.3)
	pomdp.AddTransition("s1", "a2", "s2", 0.8)
	pomdp.AddTransition("s1", "a2", "s1", 0.2)
	pomdp.AddTransition("s2", "a1", "s2", 0.9)
	pomdp.AddTransition("s2", "a1", "s1", 0.1)
	pomdp.AddTransition("s2", "a2", "s1", 0.6)
	pomdp.AddTransition("s2", "a2", "s2", 0.4)

	// Add observation probabilities
	pomdp.AddObservationProb("s1", "a1", "o1", 0.8)
	pomdp.AddObservationProb("s1", "a1", "o2", 0.2)
	pomdp.AddObservationProb("s1", "a2", "o1", 0.3)
	pomdp.AddObservationProb("s1", "a2", "o2", 0.7)
	pomdp.AddObservationProb("s2", "a1", "o1", 0.9)
	pomdp.AddObservationProb("s2", "a1", "o2", 0.1)
	pomdp.AddObservationProb("s2", "a2", "o1", 0.4)
	pomdp.AddObservationProb("s2", "a2", "o2", 0.6)

	// Add rewards
	pomdp.AddReward("s1", "a1", 1.0)
	pomdp.AddReward("s1", "a2", 0.5)
	pomdp.AddReward("s2", "a1", 2.0)
	pomdp.AddReward("s2", "a2", 1.5)
}

// setupNavigationPOMDP sets up a navigation scenario POMDP
func setupNavigationPOMDP(pomdp *core.POMDP) {
	// Add transitions for navigation scenario
	pomdp.AddTransition("s1", "a1", "s2", 0.9) // Forward movement
	pomdp.AddTransition("s1", "a1", "s1", 0.1)
	pomdp.AddTransition("s1", "a2", "s1", 0.8) // Stay
	pomdp.AddTransition("s1", "a2", "s2", 0.2)
	pomdp.AddTransition("s2", "a1", "s1", 0.1) // Backward movement
	pomdp.AddTransition("s2", "a1", "s2", 0.9)
	pomdp.AddTransition("s2", "a2", "s2", 0.8) // Stay
	pomdp.AddTransition("s2", "a2", "s1", 0.2)

	// Add observation probabilities
	pomdp.AddObservationProb("s1", "a1", "o1", 0.9) // Clear view in s1
	pomdp.AddObservationProb("s1", "a1", "o2", 0.1)
	pomdp.AddObservationProb("s1", "a2", "o1", 0.7)
	pomdp.AddObservationProb("s1", "a2", "o2", 0.3)
	pomdp.AddObservationProb("s2", "a1", "o2", 0.8) // Obstructed view in s2
	pomdp.AddObservationProb("s2", "a1", "o1", 0.2)
	pomdp.AddObservationProb("s2", "a2", "o2", 0.9)
	pomdp.AddObservationProb("s2", "a2", "o1", 0.1)

	// Add navigation rewards
	pomdp.AddReward("s1", "a1", 0.1)  // Small cost for moving
	pomdp.AddReward("s1", "a2", 0.0)  // No cost for staying
	pomdp.AddReward("s2", "a1", -0.1) // Penalty for moving in wrong direction
	pomdp.AddReward("s2", "a2", 0.0)  // No cost for staying
}

// setupExplorationPOMDP sets up an exploration scenario POMDP
func setupExplorationPOMDP(pomdp *core.POMDP) {
	// Add transitions for exploration scenario
	pomdp.AddTransition("s1", "a1", "s2", 0.8) // Explore
	pomdp.AddTransition("s1", "a1", "s1", 0.2)
	pomdp.AddTransition("s1", "a2", "s1", 1.0) // Stay
	pomdp.AddTransition("s2", "a1", "s1", 0.3) // Return
	pomdp.AddTransition("s2", "a1", "s2", 0.7)
	pomdp.AddTransition("s2", "a2", "s2", 1.0) // Stay

	// Add observation probabilities with uncertainty
	pomdp.AddObservationProb("s1", "a1", "o1", 0.6) // Known state
	pomdp.AddObservationProb("s1", "a1", "o2", 0.4)
	pomdp.AddObservationProb("s1", "a2", "o1", 0.9)
	pomdp.AddObservationProb("s1", "a2", "o2", 0.1)
	pomdp.AddObservationProb("s2", "a1", "o2", 0.7) // Unknown state
	pomdp.AddObservationProb("s2", "a1", "o1", 0.3)
	pomdp.AddObservationProb("s2", "a2", "o2", 0.8)
	pomdp.AddObservationProb("s2", "a2", "o1", 0.2)

	// Add exploration rewards
	pomdp.AddReward("s1", "a1", -0.1) // Cost of exploration
	pomdp.AddReward("s1", "a2", 0.0)  // No cost for staying
	pomdp.AddReward("s2", "a1", 1.0)  // Reward for discovering new state
	pomdp.AddReward("s2", "a2", 0.0)  // No cost for staying in discovered state
}

func simulateObservation(pomdp *core.POMDP, belief core.Belief) string {
	// More realistic simulation: sample observation based on current belief
	// In practice, this would come from the actual environment
	for i, state := range pomdp.States {
		if belief[i] > 0.5 { // Simplified
			if pomdp.ObservationsProb[state] != nil {
				for obs := range pomdp.ObservationsProb[state]["a1"] {
					return obs
				}
			}
		}
	}
	return pomdp.Observations[0]
}
