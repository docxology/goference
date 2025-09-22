package tests

import (
	"math"
	"testing"

	"active-goference/pkg/core"
	"active-goference/pkg/inference"
)

func TestNewInferenceEngine(t *testing.T) {
	engine := inference.NewInferenceEngine(1e-6, 100)

	if engine.Tolerance != 1e-6 {
		t.Errorf("Expected tolerance 1e-6, got %f", engine.Tolerance)
	}

	if engine.MaxIter != 100 {
		t.Errorf("Expected maxIter 100, got %d", engine.MaxIter)
	}
}

func TestComputeVariationalFreeEnergy(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)

	belief := core.Belief{0.5, 0.5}
	policy := core.NewPolicy()
	policy.SetAction(belief, "a1")

	vfe := engine.ComputeVariationalFreeEnergy(pomdp, belief, policy)

	if math.IsInf(vfe, 0) {
		t.Error("VFE should be finite")
	}
}

func TestComputeExpectedFreeEnergy(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)

	policy := core.NewPolicy()
	policy.SetAction(core.Belief{0.5, 0.5}, "a1")

	efe := engine.ComputeExpectedFreeEnergy(pomdp, policy)

	if math.IsInf(efe, 0) {
		t.Error("EFE should be finite")
	}
}

func TestMinimizeFreeEnergy(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	engine := inference.NewInferenceEngine(1e-6, 100)

	belief := core.Belief{0.5, 0.5}

	policy, fe := engine.MinimizeFreeEnergy(pomdp, belief)

	if len(policy.Actions) == 0 {
		t.Error("Expected non-empty policy")
	}

	if math.IsInf(fe, 0) {
		t.Error("Free energy should be finite")
	}
}

func TestKLDivergence(t *testing.T) {
	engine := inference.NewInferenceEngine(1e-6, 100)

	q := core.Belief{0.5, 0.5}
	p := core.Belief{0.3, 0.7}

	kl := engine.ComputeKLDivergence(q, p)

	if kl < 0 {
		t.Errorf("KL divergence should be non-negative, got %f", kl)
	}
}

func TestExpectedLogLikelihood(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	pomdp.AddObservationProb("s1", "a1", "o1", 0.8)
	pomdp.AddObservationProb("s2", "a1", "o1", 0.2)

	engine := inference.NewInferenceEngine(1e-6, 100)
	belief := core.Belief{0.5, 0.5}
	policy := core.NewPolicy()
	policy.SetAction(belief, "a1")

	ell := engine.ComputeExpectedLogLikelihood(pomdp, belief, policy)

	if math.IsInf(ell, -1) { // Should not be -inf
		t.Error("Expected log likelihood should not be -inf")
	}
}

func TestExpectedEntropy(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	pomdp.AddObservationProb("s1", "a1", "o1", 0.8)
	pomdp.AddObservationProb("s2", "a1", "o1", 0.2)

	engine := inference.NewInferenceEngine(1e-6, 100)
	policy := core.NewPolicy()
	policy.SetAction(core.Belief{0.5, 0.5}, "a1")

	entropy := engine.ComputeExpectedEntropy(pomdp, policy)

	if entropy < 0 {
		t.Errorf("Expected entropy should be non-negative, got %f", entropy)
	}
}

func TestInferencePredictNextBelief(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})
	pomdp.AddTransition("s1", "a1", "s2", 0.8)

	engine := inference.NewInferenceEngine(1e-6, 100)
	currentBelief := core.Belief{0.5, 0.5}
	policy := core.NewPolicy()
	policy.SetAction(currentBelief, "a1")

	predicted := engine.PredictNextBelief(pomdp, currentBelief, policy)

	if len(predicted) != 2 {
		t.Errorf("Expected 2 states, got %d", len(predicted))
	}

	// Check that prediction is reasonable
	if predicted[0] < 0 || predicted[1] < 0 {
		t.Error("Predicted beliefs should be non-negative")
	}
}
