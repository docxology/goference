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

	logger.Info("🎯 Running Basic Active Inference Example")
	logger.Info("This demonstrates fundamental active inference principles")

	// Create a simple 2-state POMDP (active inference doesn't use discount factors)
	pomdp := core.NewPOMDP(
		[]string{"safe", "dangerous"},
		[]string{"explore", "stay"},
		[]string{"comfortable", "uncomfortable"},
	)

	// Setup basic scenario
	pomdp.AddTransition("safe", "explore", "safe", 0.8)
	pomdp.AddTransition("safe", "explore", "dangerous", 0.2)
	pomdp.AddTransition("safe", "stay", "safe", 1.0)
	pomdp.AddTransition("dangerous", "explore", "dangerous", 0.7)
	pomdp.AddTransition("dangerous", "explore", "safe", 0.3)
	pomdp.AddTransition("dangerous", "stay", "dangerous", 1.0)

	pomdp.AddObservationProb("safe", "explore", "comfortable", 0.9)
	pomdp.AddObservationProb("safe", "explore", "uncomfortable", 0.1)
	pomdp.AddObservationProb("safe", "stay", "comfortable", 1.0)
	pomdp.AddObservationProb("dangerous", "explore", "uncomfortable", 0.8)
	pomdp.AddObservationProb("dangerous", "explore", "comfortable", 0.2)
	pomdp.AddObservationProb("dangerous", "stay", "uncomfortable", 1.0)

	pomdp.AddReward("safe", "explore", 0.1)
	pomdp.AddReward("safe", "stay", 0.0)
	pomdp.AddReward("dangerous", "explore", -0.2)
	pomdp.AddReward("dangerous", "stay", 0.0)

	// Save configuration
	if err := output.SavePOMDPVisualization(pomdp, "basic_example_config"); err != nil {
		logger.Error("Failed to save POMDP config:", err)
	}

	// Create agent
	beliefUpdater := belief.NewBeliefUpdater()
	inferenceEngine := inference.NewInferenceEngine(1e-6, 100)
	policyPlanner := planning.NewPolicyPlanner(3, 1.0, inferenceEngine)
	agent := core.NewActiveInferenceAgent(pomdp, beliefUpdater, inferenceEngine, policyPlanner)

	// Initialize belief
	currentBelief := core.NewBelief(pomdp.States)
	logger.Info("Initial belief state:")
	logger.Info(utils.VisualizeBelief(currentBelief, pomdp.States))

	// Save initial visualization
	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "basic_initial_belief"); err != nil {
		logger.Error("Failed to save initial belief:", err)
	}

	// Run simulation with data collection for PNG visualizations
	logger.Info("Running basic active inference simulation...")
	steps := 5

	var vfeValues []float64
	var efeValues []float64
	var allBeliefs []core.Belief
	var observations []string
	var actions []string

	// Initialize with previous action for belief updating
	previousAction := "explore" // Default action for first observation

	for step := 0; step < steps; step++ {
		logger.Info(fmt.Sprintf("Step %d:", step+1))

		// Simulate observation
		observation := simulateBasicObservation(pomdp, currentBelief)
		observations = append(observations, observation)

		// Act with previous action for belief updating
		start := time.Now()
		action, newBelief, policy := agent.ActWithAction(nil, currentBelief, observation, previousAction)
		actions = append(actions, action)
		utils.Log.Timing(fmt.Sprintf("Basic step %d", step+1), start)

		// Update previous action for next iteration
		previousAction = action

		// Log results
		logger.Info(fmt.Sprintf("Observation: %s", observation))
		logger.Info(fmt.Sprintf("Action: %s", action))
		logger.Info("New belief:")
		logger.Info(utils.VisualizeBelief(newBelief, pomdp.States))

		// Compute free energies and collect for PNG
		vfe := inferenceEngine.ComputeVariationalFreeEnergy(pomdp, newBelief, policy)
		efe := inferenceEngine.ComputeExpectedFreeEnergy(pomdp, policy)
		vfeValues = append(vfeValues, vfe)
		efeValues = append(efeValues, efe)
		logger.Info(utils.VisualizeFreeEnergy(vfe, efe))

		// Save step data
		if err := output.SaveDetailedTrace(pomdp, newBelief, policy, vfe, efe, step+1, fmt.Sprintf("basic_step_%d", step+1)); err != nil {
			logger.Error("Failed to save trace:", err)
		}

		// Collect belief for evolution visualization
		allBeliefs = append(allBeliefs, newBelief)

		currentBelief = newBelief
	}

	// Generate final visualizations
	finalPolicy := policyPlanner.PlanPolicy(pomdp, currentBelief)
	if err := output.SavePolicyVisualization(finalPolicy, pomdp, "basic_final_policy"); err != nil {
		logger.Error("Failed to save policy:", err)
	}

	if err := output.SaveBeliefVisualization(currentBelief, pomdp.States, "basic_final_belief"); err != nil {
		logger.Error("Failed to save final belief:", err)
	}

	// Generate simple PNG visualizations
	logger.Info("📊 Generating simple PNG visualizations...")

	// Generate VFE curve
	if err := utils.SimplePNGVisualizerInstance.GenerateVFECurve(vfeValues, "basic_vfe_curve"); err != nil {
		logger.Error("Failed to generate VFE curve PNG:", err)
	} else {
		logger.Info("✅ VFE curve PNG generated")
	}

	// Generate EFE curve
	if err := utils.SimplePNGVisualizerInstance.GenerateEFECurve(efeValues, "basic_efe_curve"); err != nil {
		logger.Error("Failed to generate EFE curve PNG:", err)
	} else {
		logger.Info("✅ EFE curve PNG generated")
	}

	// Generate belief evolution
	if err := utils.SimplePNGVisualizerInstance.GenerateBeliefEvolution(allBeliefs, pomdp.States, "basic_belief_evolution"); err != nil {
		logger.Error("Failed to generate belief evolution PNG:", err)
	} else {
		logger.Info("✅ Belief evolution PNG generated")
	}

	// Generate multi-panel dashboard
	if err := utils.SimplePNGVisualizerInstance.GenerateMultiPanel(
		"Basic Active Inference Dashboard",
		allBeliefs, vfeValues, efeValues, observations, "basic_dashboard"); err != nil {
		logger.Error("Failed to generate dashboard PNG:", err)
	} else {
		logger.Info("✅ Multi-panel dashboard PNG generated")
	}

	logger.Info("✅ Basic example completed successfully!")
	logger.Info("📁 Check output directories for detailed results:")
	logger.Info("   📊 Text outputs: output/{logs,traces,visualizations,reports}/")
	logger.Info("   🖼️  PNG visualizations: output/png_visualizations/")
}

func simulateBasicObservation(pomdp *core.POMDP, belief core.Belief) string {
	// Simple observation simulation
	if belief[0] > belief[1] { // More belief in safe state
		return "comfortable"
	}
	return "uncomfortable"
}
