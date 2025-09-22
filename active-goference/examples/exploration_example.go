package main

import (
	"fmt"
	"math/rand"
	"time"

	"active-goference/pkg/belief"
	"active-goference/pkg/core"
	"active-goference/pkg/inference"
	"active-goference/pkg/planning"
	"active-goference/pkg/utils"
)

func main() {
	logger := utils.Log
	output := utils.Output

	logger.Info("🔍 Running Exploration Active Inference Example")
	logger.Info("This demonstrates information-seeking and curiosity-driven behavior")

	// Create an exploration POMDP with unknown states (active inference doesn't use discount factors)
	pomdp := core.NewPOMDP(
		[]string{"known_state", "unknown_state_1", "unknown_state_2", "unknown_state_3"},
		[]string{"explore", "stay", "investigate"},
		[]string{"familiar", "novel", "confusing", "rewarding"},
	)

	// Setup exploration scenario with high uncertainty
	setupExplorationTransitions(pomdp)
	setupExplorationObservations(pomdp)
	setupExplorationRewards(pomdp)

	// Save configuration
	if err := output.SavePOMDPVisualization(pomdp, "exploration_config"); err != nil {
		logger.Error("Failed to save POMDP config:", err)
	}

	// Create agent optimized for exploration (high temperature for more random behavior)
	beliefUpdater := belief.NewBeliefUpdater()
	inferenceEngine := inference.NewInferenceEngine(1e-6, 150)
	policyPlanner := planning.NewPolicyPlanner(5, 2.0, inferenceEngine) // High temperature for exploration
	agent := core.NewActiveInferenceAgent(pomdp, beliefUpdater, inferenceEngine, policyPlanner)

	// Initialize belief - start with high uncertainty
	currentBelief := core.Belief{0.8, 0.07, 0.07, 0.06} // Mostly in known state but with uncertainty
	logger.Info("Initial belief state (high uncertainty in unknown states):")
	logger.Info(utils.VisualizeBelief(currentBelief, pomdp.States))

	// Save initial visualization
	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "exploration_initial_belief"); err != nil {
		logger.Error("Failed to save initial belief:", err)
	}

	// Run exploration simulation
	logger.Info("Running exploration active inference simulation...")
	steps := 20
	exploredStates := make(map[string]bool)

	for step := 0; step < steps; step++ {
		logger.Info(fmt.Sprintf("🔍 Exploration Step %d/%d", step+1, steps))

		// Simulate observation with some randomness
		observation := simulateExplorationObservation(pomdp, currentBelief)

		// Act with timing
		start := time.Now()
		action, newBelief, policy := agent.Act(nil, currentBelief, observation)
		utils.Log.Timing(fmt.Sprintf("Exploration step %d", step+1), start)

		// Track explored states
		currentPos := getExplorationPosition(newBelief, pomdp.States)
		if !exploredStates[currentPos] {
			exploredStates[currentPos] = true
			logger.Info(fmt.Sprintf("🗺️  Discovered new state: %s", currentPos))
		}

		// Log results with exploration focus
		logger.Info(fmt.Sprintf("📍 Current Position: %s", currentPos))
		logger.Info(fmt.Sprintf("👁️  Observation: %s", observation))
		logger.Info(fmt.Sprintf("🎯 Action Selected: %s", action))
		logger.Info("🧠 Updated Belief State:")
		logger.Info(utils.VisualizeBelief(newBelief, pomdp.States))

		// Compute free energies (exploration should show high epistemic value)
		vfe := inferenceEngine.ComputeVariationalFreeEnergy(pomdp, newBelief, policy)
		efe := inferenceEngine.ComputeExpectedFreeEnergy(pomdp, policy)
		logger.Info(utils.VisualizeFreeEnergy(vfe, efe))

		// Save detailed trace
		if err := output.SaveDetailedTrace(pomdp, newBelief, policy, vfe, efe, step+1, fmt.Sprintf("explore_step_%d", step+1)); err != nil {
			logger.Error("Failed to save step trace:", err)
		}

		// Save step visualization
		if err := output.SaveBeliefVisualization(newBelief, pomdp.States, fmt.Sprintf("explore_belief_step_%d", step+1)); err != nil {
			logger.Error("Failed to save step belief:", err)
		}

		currentBelief = newBelief

		// Check if we've explored enough states
		if len(exploredStates) >= 3 {
			logger.Info("🎯 Exploration goal achieved! Agent has discovered multiple states.")
			break
		}
	}

	// Generate comprehensive visualizations
	finalPolicy := policyPlanner.PlanPolicy(pomdp, currentBelief)
	if err := output.SavePolicyVisualization(finalPolicy, pomdp, "exploration_final_policy"); err != nil {
		logger.Error("Failed to save policy:", err)
	}

	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "exploration_final_belief"); err != nil {
		logger.Error("Failed to save final belief:", err)
	}

	// Save exploration summary
	explorationSummary := fmt.Sprintf("Exploration Summary:\n"+
		"Starting State: known_state\n"+
		"States Discovered: %d/%d\n"+
		"Final Position: %s\n"+
		"Active Inference demonstrated curiosity-driven exploration\n"+
		"Information-seeking behavior through expected free energy minimization",
		len(exploredStates), len(pomdp.States),
		getExplorationPosition(currentBelief, pomdp.States))

	if err := output.SaveTextToFile(explorationSummary, "exploration_summary", "reports"); err != nil {
		logger.Error("Failed to save exploration summary:", err)
	}

	logger.Info("✅ Exploration example completed successfully!")
	logger.Info(fmt.Sprintf("🗺️  Discovered %d out of %d states", len(exploredStates), len(pomdp.States)))
	logger.Info("📁 Exploration traces and visualizations saved to organized output directories")
}

func setupExplorationTransitions(pomdp *core.POMDP) {
	// High uncertainty transitions - exploration leads to unknown states
	pomdp.AddTransition("known_state", "explore", "unknown_state_1", 0.4)
	pomdp.AddTransition("known_state", "explore", "unknown_state_2", 0.4)
	pomdp.AddTransition("known_state", "explore", "known_state", 0.2)
	pomdp.AddTransition("known_state", "stay", "known_state", 1.0)
	pomdp.AddTransition("known_state", "investigate", "known_state", 0.9)
	pomdp.AddTransition("known_state", "investigate", "unknown_state_1", 0.1)

	// Unknown states have uncertain transitions
	pomdp.AddTransition("unknown_state_1", "explore", "unknown_state_2", 0.5)
	pomdp.AddTransition("unknown_state_1", "explore", "unknown_state_3", 0.3)
	pomdp.AddTransition("unknown_state_1", "explore", "known_state", 0.2)
	pomdp.AddTransition("unknown_state_1", "investigate", "unknown_state_1", 0.8)
	pomdp.AddTransition("unknown_state_1", "investigate", "unknown_state_2", 0.2)

	pomdp.AddTransition("unknown_state_2", "explore", "unknown_state_3", 0.6)
	pomdp.AddTransition("unknown_state_2", "explore", "unknown_state_1", 0.3)
	pomdp.AddTransition("unknown_state_2", "explore", "known_state", 0.1)
	pomdp.AddTransition("unknown_state_2", "investigate", "unknown_state_2", 0.9)

	pomdp.AddTransition("unknown_state_3", "explore", "unknown_state_1", 0.4)
	pomdp.AddTransition("unknown_state_3", "explore", "unknown_state_2", 0.4)
	pomdp.AddTransition("unknown_state_3", "explore", "known_state", 0.2)
	pomdp.AddTransition("unknown_state_3", "investigate", "unknown_state_3", 0.7)
	pomdp.AddTransition("unknown_state_3", "investigate", "known_state", 0.3)
}

func setupExplorationObservations(pomdp *core.POMDP) {
	// Known state has familiar observations
	pomdp.AddObservationProb("known_state", "explore", "familiar", 0.9)
	pomdp.AddObservationProb("known_state", "explore", "novel", 0.1)
	pomdp.AddObservationProb("known_state", "stay", "familiar", 1.0)
	pomdp.AddObservationProb("known_state", "investigate", "familiar", 0.8)
	pomdp.AddObservationProb("known_state", "investigate", "confusing", 0.2)

	// Unknown states have novel/confusing observations
	pomdp.AddObservationProb("unknown_state_1", "explore", "novel", 0.7)
	pomdp.AddObservationProb("unknown_state_1", "explore", "confusing", 0.3)
	pomdp.AddObservationProb("unknown_state_1", "investigate", "novel", 0.4)
	pomdp.AddObservationProb("unknown_state_1", "investigate", "confusing", 0.6)

	pomdp.AddObservationProb("unknown_state_2", "explore", "novel", 0.6)
	pomdp.AddObservationProb("unknown_state_2", "explore", "rewarding", 0.4)
	pomdp.AddObservationProb("unknown_state_2", "investigate", "novel", 0.3)
	pomdp.AddObservationProb("unknown_state_2", "investigate", "rewarding", 0.7)

	pomdp.AddObservationProb("unknown_state_3", "explore", "confusing", 0.5)
	pomdp.AddObservationProb("unknown_state_3", "explore", "rewarding", 0.5)
	pomdp.AddObservationProb("unknown_state_3", "investigate", "confusing", 0.2)
	pomdp.AddObservationProb("unknown_state_3", "investigate", "rewarding", 0.8)
}

func setupExplorationRewards(pomdp *core.POMDP) {
	// Exploration has costs but discovering new states has rewards
	pomdp.AddReward("known_state", "explore", -0.1) // Cost of exploration
	pomdp.AddReward("known_state", "stay", 0.0)
	pomdp.AddReward("known_state", "investigate", -0.05) // Small cost for investigation

	pomdp.AddReward("unknown_state_1", "explore", -0.1)
	pomdp.AddReward("unknown_state_1", "investigate", 0.5) // Reward for investigating unknown

	pomdp.AddReward("unknown_state_2", "explore", -0.1)
	pomdp.AddReward("unknown_state_2", "investigate", 1.0) // Higher reward for this state

	pomdp.AddReward("unknown_state_3", "explore", -0.1)
	pomdp.AddReward("unknown_state_3", "investigate", 0.8) // Good reward for exploration
}

func simulateExplorationObservation(pomdp *core.POMDP, belief core.Belief) string {
	// Simulate observations with some randomness for exploration
	maxIdx := 0
	maxProb := 0.0

	for i, prob := range belief {
		if prob > maxProb {
			maxProb = prob
			maxIdx = i
		}
	}

	currentState := pomdp.States[maxIdx]

	// Add some randomness to observations to simulate real-world uncertainty
	rand.Seed(time.Now().UnixNano() + int64(maxIdx))
	randomness := rand.Float64()

	switch currentState {
	case "known_state":
		if randomness < 0.9 {
			return "familiar"
		}
		return "novel"
	case "unknown_state_1":
		if randomness < 0.6 {
			return "novel"
		}
		return "confusing"
	case "unknown_state_2":
		if randomness < 0.5 {
			return "novel"
		}
		return "rewarding"
	case "unknown_state_3":
		if randomness < 0.4 {
			return "confusing"
		}
		return "rewarding"
	default:
		return "confusing"
	}
}

func getExplorationPosition(belief core.Belief, states []string) string {
	maxIdx := 0
	maxProb := 0.0

	for i, prob := range belief {
		if prob > maxProb {
			maxProb = prob
			maxIdx = i
		}
	}

	return states[maxIdx]
}
