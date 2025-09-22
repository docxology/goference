package main

import (
	"fmt"
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

	logger.Info("🧭 Running Navigation Active Inference Example")
	logger.Info("This demonstrates goal-directed navigation behavior")

	// Create a navigation POMDP with 4 states representing positions (active inference doesn't use discount factors)
	pomdp := core.NewPOMDP(
		[]string{"position_1", "position_2", "position_3", "goal"},
		[]string{"move_forward", "move_backward", "stay"},
		[]string{"clear", "obstructed", "goal_visible"},
	)

	// Setup navigation scenario
	setupNavigationTransitions(pomdp)
	setupNavigationObservations(pomdp)
	setupNavigationRewards(pomdp)

	// Save configuration
	if err := output.SavePOMDPVisualization(pomdp, "navigation_config"); err != nil {
		logger.Error("Failed to save POMDP config:", err)
	}

	// Create agent with longer planning horizon for navigation
	beliefUpdater := belief.NewBeliefUpdater()
	inferenceEngine := inference.NewInferenceEngine(1e-8, 200)
	policyPlanner := planning.NewPolicyPlanner(8, 0.8, inferenceEngine) // Longer horizon, lower temperature
	agent := core.NewActiveInferenceAgent(pomdp, beliefUpdater, inferenceEngine, policyPlanner)

	// Initialize belief - start at position 1 with uncertainty
	currentBelief := core.Belief{0.7, 0.2, 0.05, 0.05} // Mostly at position 1
	logger.Info("Initial belief state (starting at position 1):")
	logger.Info(utils.VisualizeBelief(currentBelief, pomdp.States))

	// Save initial visualization
	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "navigation_initial_belief"); err != nil {
		logger.Error("Failed to save initial belief:", err)
	}

	// Run navigation simulation
	logger.Info("Running navigation active inference simulation...")
	steps := 15

	for step := 0; step < steps; step++ {
		logger.Info(fmt.Sprintf("🔄 Navigation Step %d/%d", step+1, steps))

		// Simulate observation based on current position
		observation := simulateNavigationObservation(pomdp, currentBelief)

		// Act with timing
		start := time.Now()
		action, newBelief, policy := agent.Act(nil, currentBelief, observation)
		utils.Log.Timing(fmt.Sprintf("Navigation step %d", step+1), start)

		// Log detailed results
		logger.Info(fmt.Sprintf("📍 Current Position: %s", getCurrentPosition(currentBelief, pomdp.States)))
		logger.Info(fmt.Sprintf("👁️  Observation: %s", observation))
		logger.Info(fmt.Sprintf("🎯 Action Selected: %s", action))
		logger.Info("🧠 Updated Belief State:")
		logger.Info(utils.VisualizeBelief(newBelief, pomdp.States))

		// Compute and display free energies
		vfe := inferenceEngine.ComputeVariationalFreeEnergy(pomdp, newBelief, policy)
		efe := inferenceEngine.ComputeExpectedFreeEnergy(pomdp, policy)
		logger.Info(utils.VisualizeFreeEnergy(vfe, efe))

		// Save detailed trace
		if err := output.SaveDetailedTrace(pomdp, newBelief, policy, vfe, efe, step+1, fmt.Sprintf("nav_step_%d", step+1)); err != nil {
			logger.Error("Failed to save step trace:", err)
		}

		// Save step visualization
		if err := output.SaveBeliefVisualization(newBelief, pomdp.States, fmt.Sprintf("nav_belief_step_%d", step+1)); err != nil {
			logger.Error("Failed to save step belief:", err)
		}

		// Check if goal reached
		if newBelief[3] > 0.8 { // High confidence in goal state
			logger.Info("🎉 Goal reached! Active inference successfully navigated to the target.")
			break
		}

		currentBelief = newBelief
	}

	// Generate comprehensive visualizations
	finalPolicy := policyPlanner.PlanPolicy(pomdp, currentBelief)
	if err := output.SavePolicyVisualization(finalPolicy, pomdp, "navigation_final_policy"); err != nil {
		logger.Error("Failed to save policy:", err)
	}

	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "navigation_final_belief"); err != nil {
		logger.Error("Failed to save final belief:", err)
	}

	// Save navigation summary
	navigationSummary := fmt.Sprintf("Navigation Summary:\n"+
		"Starting Position: position_1\n"+
		"Goal Position: goal\n"+
		"Final Position: %s\n"+
		"Goal Reached: %t\n"+
		"Active Inference demonstrated goal-directed behavior through free energy minimization",
		getCurrentPosition(currentBelief, pomdp.States),
		currentBelief[3] > 0.8)

	if err := output.SaveTextToFile(navigationSummary, "navigation_summary", "reports"); err != nil {
		logger.Error("Failed to save navigation summary:", err)
	}

	logger.Info("✅ Navigation example completed successfully!")
	logger.Info("📁 Navigation traces and visualizations saved to organized output directories")
}

func setupNavigationTransitions(pomdp *core.POMDP) {
	// Forward movement transitions
	pomdp.AddTransition("position_1", "move_forward", "position_2", 0.9)
	pomdp.AddTransition("position_1", "move_forward", "position_1", 0.1)
	pomdp.AddTransition("position_1", "stay", "position_1", 1.0)

	pomdp.AddTransition("position_2", "move_forward", "position_3", 0.8)
	pomdp.AddTransition("position_2", "move_forward", "position_2", 0.2)
	pomdp.AddTransition("position_2", "stay", "position_2", 1.0)
	pomdp.AddTransition("position_2", "move_backward", "position_1", 0.9)
	pomdp.AddTransition("position_2", "move_backward", "position_2", 0.1)

	pomdp.AddTransition("position_3", "move_forward", "goal", 0.7)
	pomdp.AddTransition("position_3", "move_forward", "position_3", 0.3)
	pomdp.AddTransition("position_3", "stay", "position_3", 1.0)
	pomdp.AddTransition("position_3", "move_backward", "position_2", 0.8)
	pomdp.AddTransition("position_3", "move_backward", "position_3", 0.2)

	pomdp.AddTransition("goal", "stay", "goal", 1.0)
	pomdp.AddTransition("goal", "move_backward", "position_3", 0.1)
	pomdp.AddTransition("goal", "move_backward", "goal", 0.9)
}

func setupNavigationObservations(pomdp *core.POMDP) {
	// Clear view in early positions
	pomdp.AddObservationProb("position_1", "move_forward", "clear", 0.9)
	pomdp.AddObservationProb("position_1", "move_forward", "obstructed", 0.1)
	pomdp.AddObservationProb("position_1", "stay", "clear", 0.95)
	pomdp.AddObservationProb("position_1", "stay", "obstructed", 0.05)

	// Some uncertainty in middle positions
	pomdp.AddObservationProb("position_2", "move_forward", "clear", 0.7)
	pomdp.AddObservationProb("position_2", "move_forward", "obstructed", 0.3)
	pomdp.AddObservationProb("position_2", "stay", "clear", 0.8)
	pomdp.AddObservationProb("position_2", "stay", "obstructed", 0.2)

	// Goal becomes visible as we approach
	pomdp.AddObservationProb("position_3", "move_forward", "goal_visible", 0.6)
	pomdp.AddObservationProb("position_3", "move_forward", "clear", 0.4)
	pomdp.AddObservationProb("position_3", "stay", "goal_visible", 0.8)
	pomdp.AddObservationProb("position_3", "stay", "clear", 0.2)

	// Goal is clearly visible when reached
	pomdp.AddObservationProb("goal", "stay", "goal_visible", 1.0)
}

func setupNavigationRewards(pomdp *core.POMDP) {
	// Small cost for movement, reward for progress
	pomdp.AddReward("position_1", "move_forward", 0.1)
	pomdp.AddReward("position_1", "stay", 0.0)
	pomdp.AddReward("position_2", "move_forward", 0.2)
	pomdp.AddReward("position_2", "move_backward", -0.1)
	pomdp.AddReward("position_2", "stay", 0.0)
	pomdp.AddReward("position_3", "move_forward", 0.5)
	pomdp.AddReward("position_3", "move_backward", -0.2)
	pomdp.AddReward("position_3", "stay", 0.0)
	pomdp.AddReward("goal", "stay", 10.0) // High reward for reaching goal
}

func simulateNavigationObservation(pomdp *core.POMDP, belief core.Belief) string {
	// Simulate realistic observations based on current position belief
	maxIdx := 0
	maxProb := 0.0

	for i, prob := range belief {
		if prob > maxProb {
			maxProb = prob
			maxIdx = i
		}
	}

	currentState := pomdp.States[maxIdx]

	// Return observation based on current state
	switch currentState {
	case "position_1":
		return "clear"
	case "position_2":
		return "clear" // Sometimes obstructed
	case "position_3":
		return "goal_visible"
	case "goal":
		return "goal_visible"
	default:
		return "clear"
	}
}

func getCurrentPosition(belief core.Belief, states []string) string {
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
