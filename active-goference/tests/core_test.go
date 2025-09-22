package tests

import (
	"context"
	"math"
	"testing"

	"active-goference/pkg/belief"
	"active-goference/pkg/core"
	"active-goference/pkg/inference"
	"active-goference/pkg/planning"
)

func TestNewPOMDP(t *testing.T) {
	states := []string{"s1", "s2"}
	actions := []string{"a1", "a2"}
	observations := []string{"o1", "o2"}

	// Note: Active inference doesn't use discount factors
	pomdp := core.NewPOMDP(states, actions, observations)

	if len(pomdp.States) != 2 || pomdp.States[0] != "s1" {
		t.Errorf("Expected states ['s1', 's2'], got %v", pomdp.States)
	}
	if len(pomdp.Actions) != 2 || pomdp.Actions[0] != "a1" {
		t.Errorf("Expected actions ['a1', 'a2'], got %v", pomdp.Actions)
	}
	// Note: Active inference doesn't use discount factors
}

func TestAddTransition(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	err := pomdp.AddTransition("s1", "a1", "s2", 0.8)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if pomdp.Transitions["s1"]["a1"]["s2"] != 0.8 {
		t.Errorf("Expected transition prob 0.8, got %f", pomdp.Transitions["s1"]["a1"]["s2"])
	}
}

func TestAddTransitionInvalidState(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	err := pomdp.AddTransition("s3", "a1", "s2", 0.8)
	if err == nil {
		t.Error("Expected error for invalid state")
	}
}

func TestAddObservationProb(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1", "o2"})

	err := pomdp.AddObservationProb("s1", "a1", "o1", 0.7)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if pomdp.ObservationsProb["s1"]["a1"]["o1"] != 0.7 {
		t.Errorf("Expected observation prob 0.7, got %f", pomdp.ObservationsProb["s1"]["a1"]["o1"])
	}
}

func TestAddObservationProbInvalidObservation(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	err := pomdp.AddObservationProb("s1", "a1", "o2", 0.7)
	if err == nil {
		t.Error("Expected error for invalid observation")
	}
}

func TestAddReward(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	err := pomdp.AddReward("s1", "a1", 10.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if pomdp.Rewards["s1"]["a1"] != 10.0 {
		t.Errorf("Expected reward 10.0, got %f", pomdp.Rewards["s1"]["a1"])
	}
}

func TestAddRewardInvalidAction(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	err := pomdp.AddReward("s1", "a2", 10.0)
	if err == nil {
		t.Error("Expected error for invalid action")
	}
}

func TestNewBelief(t *testing.T) {
	states := []string{"s1", "s2", "s3"}
	belief := core.NewBelief(states)

	expected := 1.0 / 3.0
	for i, p := range belief {
		if p != expected {
			t.Errorf("Expected belief[%d] = %f, got %f", i, expected, p)
		}
	}
}

func TestBeliefNormalize(t *testing.T) {
	belief := core.Belief{0.3, 0.4, 0.2}
	belief.Normalize()

	sum := 0.0
	for _, p := range belief {
		sum += p
	}

	if math.Abs(sum-1.0) > 1e-10 {
		t.Errorf("Expected sum 1.0, got %f", sum)
	}
}

func TestBeliefNormalizeZeroSum(t *testing.T) {
	belief := core.Belief{0.0, 0.0, 0.0}
	belief.Normalize()

	// Should remain unchanged for zero sum
	for _, p := range belief {
		if p != 0.0 {
			t.Errorf("Expected belief to remain 0.0, got %f", p)
		}
	}
}

func TestNewPolicy(t *testing.T) {
	policy := core.NewPolicy()

	if policy.Actions == nil {
		t.Error("Expected non-nil policy actions")
	}

	if len(policy.Actions) != 0 {
		t.Errorf("Expected empty policy, got length %d", len(policy.Actions))
	}
}

func TestPolicyGetAction(t *testing.T) {
	belief := core.Belief{0.5, 0.5}
	policy := core.NewPolicy()
	policy.SetAction(belief, "a1")

	action := policy.GetAction(belief)
	if action != "a1" {
		t.Errorf("Expected action 'a1', got '%s'", action)
	}
}

func TestPolicyGetActionNonExistentBelief(t *testing.T) {
	belief := core.Belief{0.5, 0.5}
	policy := core.NewPolicy()

	action := policy.GetAction(belief)
	if action != "" {
		t.Errorf("Expected empty action for non-existent belief, got '%s'", action)
	}
}

func TestPolicySetAction(t *testing.T) {
	belief := core.Belief{0.5, 0.5}
	policy := core.NewPolicy()
	policy.SetAction(belief, "a1")

	if policy.Actions["0.500,0.500"] != "a1" {
		t.Errorf("Expected action 'a1' for belief, got '%s'", policy.Actions["0.500,0.500"])
	}
}

func TestNewActiveInferenceAgent(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1"}, []string{"a1"}, []string{"o1"})
	// Using real implementations from their respective packages
	updater := &belief.BeliefUpdaterImpl{}
	engine := &inference.InferenceEngineImpl{}
	planner := &planning.PolicyPlannerImpl{}

	agent := core.NewActiveInferenceAgent(pomdp, updater, engine, planner)

	if agent.POMDP != pomdp {
		t.Error("Expected agent POMDP to match")
	}
}

func TestActiveInferenceAgentAct(t *testing.T) {
	// Test with real implementations
	pomdp := core.NewPOMDP([]string{"s1"}, []string{"a1"}, []string{"o1"})
	updater := belief.NewBeliefUpdater()
	engine := inference.NewInferenceEngine(1e-6, 100)
	planner := planning.NewPolicyPlanner(3, 1.0, engine)

	agent := core.NewActiveInferenceAgent(pomdp, updater, engine, planner)

	// Real context
	ctx := context.Background()

	belief := core.NewBelief(pomdp.States)
	action, newBelief, policy := agent.Act(ctx, belief, "o1")

	if action == "" {
		t.Error("Expected non-empty action")
	}
	if len(newBelief) == 0 {
		t.Error("Expected non-empty belief")
	}
	if len(policy.Actions) == 0 {
		t.Error("Expected non-empty policy")
	}
}

func TestPOMDPWithEmptyStates(t *testing.T) {
	pomdp := core.NewPOMDP([]string{}, []string{"a1"}, []string{"o1"})

	if len(pomdp.States) != 0 {
		t.Errorf("Expected empty states, got %v", pomdp.States)
	}
}

func TestBeliefWithZeroStates(t *testing.T) {
	belief := core.NewBelief([]string{})

	if len(belief) != 0 {
		t.Errorf("Expected empty belief, got length %d", len(belief))
	}
}

func TestPolicyWithMultipleBeliefs(t *testing.T) {
	policy := core.NewPolicy()

	belief1 := core.Belief{0.5, 0.5}
	belief2 := core.Belief{0.3, 0.7}

	policy.SetAction(belief1, "a1")
	policy.SetAction(belief2, "a2")

	if len(policy.Actions) != 2 {
		t.Errorf("Expected 2 actions, got %d", len(policy.Actions))
	}

	if policy.GetAction(belief1) != "a1" {
		t.Error("Expected action 'a1' for belief1")
	}

	if policy.GetAction(belief2) != "a2" {
		t.Error("Expected action 'a2' for belief2")
	}
}
