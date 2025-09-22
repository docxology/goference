package tests

import (
	"testing"

	"active-goference/pkg/core"
	"active-goference/pkg/inference"
	"active-goference/pkg/planning"
)

func TestNewPolicyPlanner(t *testing.T) {
	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(5, 1.0, engine)

	if planner.Horizon != 5 {
		t.Errorf("Expected horizon 5, got %d", planner.Horizon)
	}

	if planner.Temperature != 1.0 {
		t.Errorf("Expected temperature 1.0, got %f", planner.Temperature)
	}
}

func TestPlanPolicy(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1", "a2"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, engine)

	belief := core.Belief{0.5, 0.5}

	policy := planner.PlanPolicy(pomdp, belief)

	if len(policy.Actions) == 0 {
		t.Error("Expected non-empty policy")
	}
}

func TestSelectBestAction(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1", "a2"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, engine)

	belief := core.Belief{0.5, 0.5}

	action := planner.SelectBestAction(pomdp, belief)

	if action == "" {
		t.Error("Expected non-empty action")
	}
}

func TestPredictNextBelief(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	pomdp.AddTransition("s1", "a1", "s2", 0.8)

	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, engine)

	currentBelief := core.Belief{0.5, 0.5}
	action := "a1"

	predicted := planner.PredictNextBelief(pomdp, currentBelief, action)

	if len(predicted) != 2 {
		t.Errorf("Expected 2 states, got %d", len(predicted))
	}

	if predicted[0] < 0 || predicted[1] < 0 {
		t.Error("Predicted beliefs should be non-negative")
	}
}

func TestSoftmaxActionSelection(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1", "a2"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, engine)

	belief := core.Belief{0.5, 0.5}

	action := planner.SoftmaxActionSelection(pomdp, belief)

	if action == "" {
		t.Error("Expected non-empty action")
	}

	if action != "a1" && action != "a2" {
		t.Errorf("Expected action 'a1' or 'a2', got '%s'", action)
	}
}

func TestHierarchicalPlanning(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1", "a2"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, engine)

	belief := core.Belief{0.5, 0.5}

	policy := planner.HierarchicalPlanning(pomdp, belief, 2)

	if len(policy.Actions) == 0 {
		t.Error("Expected non-empty policy")
	}
}

func TestMergePolicies(t *testing.T) {
	_ = inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, nil)

	belief1 := core.Belief{0.5, 0.5}
	belief2 := core.Belief{0.3, 0.7}

	policy1 := core.NewPolicy()
	policy1.SetAction(belief1, "a1")
	policy2 := core.NewPolicy()
	policy2.SetAction(belief2, "a2")

	merged := planner.MergePolicies(policy1, policy2)

	if len(merged.Actions) != 2 {
		t.Errorf("Expected 2 policies, got %d", len(merged.Actions))
	}

	if merged.Actions["0.500,0.500"] != "a1" {
		t.Errorf("Expected action 'a1' for belief1, got '%s'", merged.Actions["0.500,0.500"])
	}

	if merged.Actions["0.300,0.700"] != "a2" {
		t.Errorf("Expected action 'a2' for belief2, got '%s'", merged.Actions["0.300,0.700"])
	}
}
