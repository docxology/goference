package inference

import (
	"math"

	"active-goference/pkg/core"
)

// InferenceEngineImpl implements the InferenceEngine interface
type InferenceEngineImpl struct {
	Tolerance float64
	MaxIter   int
}

// NewInferenceEngine creates a new inference engine
func NewInferenceEngine(tolerance float64, maxIter int) *InferenceEngineImpl {
	return &InferenceEngineImpl{
		Tolerance: tolerance,
		MaxIter:   maxIter,
	}
}

// ComputeVariationalFreeEnergy computes the variational free energy
// F = KL[q(s)||p(s)] - E_q[ln p(o|s)]
func (ie *InferenceEngineImpl) ComputeVariationalFreeEnergy(pomdp *core.POMDP, belief core.Belief, policy core.Policy) float64 {
	// Compute KL divergence between belief and prior
	kl := ie.computeKLDivergence(belief, ie.getPrior(pomdp))

	// Compute expected log likelihood
	expectedLogLikelihood := ie.computeExpectedLogLikelihood(pomdp, belief, policy)

	// VFE = KL - E[ln p(o|s)]
	return kl - expectedLogLikelihood
}

// ComputeExpectedFreeEnergy computes the expected free energy for a policy
// G(π) = E_q[H[p(o|s)]] + KL[q(s|π)||p(s)]
func (ie *InferenceEngineImpl) ComputeExpectedFreeEnergy(pomdp *core.POMDP, policy core.Policy) float64 {
	// Compute expected entropy of observations
	expectedEntropy := ie.computeExpectedEntropy(pomdp, policy)

	// Compute KL divergence between predicted and prior beliefs
	kl := ie.computePolicyKLDivergence(pomdp, policy)

	return expectedEntropy + kl
}

// MinimizeFreeEnergy performs free energy minimization to find optimal policy
func (ie *InferenceEngineImpl) MinimizeFreeEnergy(pomdp *core.POMDP, currentBelief core.Belief) (core.Policy, float64) {
	bestPolicy := core.NewPolicy()
	minFE := math.Inf(1)

	// Simple policy search (in practice, this would be more sophisticated)
	for _, action := range pomdp.Actions {
		policy := core.NewPolicy()
		policy.SetAction(currentBelief, action)
		fe := ie.ComputeVariationalFreeEnergy(pomdp, currentBelief, policy)
		efe := ie.ComputeExpectedFreeEnergy(pomdp, policy)

		totalFE := fe + efe

		if totalFE < minFE {
			minFE = totalFE
			bestPolicy = policy
		}
	}

	return bestPolicy, minFE
}

// Helper methods

// Public methods for testing
func (ie *InferenceEngineImpl) GetPrior(pomdp *core.POMDP) core.Belief {
	return ie.getPrior(pomdp)
}

func (ie *InferenceEngineImpl) ComputeKLDivergence(q, p core.Belief) float64 {
	return ie.computeKLDivergence(q, p)
}

func (ie *InferenceEngineImpl) ComputeExpectedLogLikelihood(pomdp *core.POMDP, belief core.Belief, policy core.Policy) float64 {
	return ie.computeExpectedLogLikelihood(pomdp, belief, policy)
}

func (ie *InferenceEngineImpl) ComputeExpectedEntropy(pomdp *core.POMDP, policy core.Policy) float64 {
	return ie.computeExpectedEntropy(pomdp, policy)
}

func (ie *InferenceEngineImpl) PredictNextBelief(pomdp *core.POMDP, currentBelief core.Belief, policy core.Policy) core.Belief {
	return ie.predictNextBelief(pomdp, currentBelief, policy)
}

// Private helper methods
func (ie *InferenceEngineImpl) getPrior(pomdp *core.POMDP) core.Belief {
	// Uniform prior for simplicity
	return core.NewBelief(pomdp.States)
}

func (ie *InferenceEngineImpl) computeKLDivergence(q, p core.Belief) float64 {
	kl := 0.0
	for i := range q {
		if q[i] > 0 && p[i] > 0 {
			kl += q[i] * math.Log(q[i]/p[i])
		}
	}
	return kl
}

func (ie *InferenceEngineImpl) computeExpectedLogLikelihood(pomdp *core.POMDP, belief core.Belief, policy core.Policy) float64 {
	expectedLL := 0.0

	// For each possible next state and observation
	for _, action := range pomdp.Actions {
		selectedAction := policy.GetAction(belief)
		if selectedAction != action {
			continue
		}

		for i, state := range pomdp.States {
			if pomdp.ObservationsProb[state] == nil {
				continue
			}
			// Expected log likelihood under current belief
			for _, obs := range pomdp.Observations {
				if pomdp.ObservationsProb[state][action] != nil {
					obsProb := pomdp.ObservationsProb[state][action][obs]
					if obsProb > 0 {
						expectedLL += belief[i] * obsProb * math.Log(obsProb)
					}
				}
			}
		}
	}

	return expectedLL
}

func (ie *InferenceEngineImpl) computeExpectedEntropy(pomdp *core.POMDP, policy core.Policy) float64 {
	expectedEntropy := 0.0

	// Compute entropy of predicted observations
	for _, action := range pomdp.Actions {
		selectedAction := policy.GetAction(core.NewBelief(pomdp.States)) // Simplified
		if selectedAction != action {
			continue
		}

		// Compute observation distribution
		obsDist := make(map[string]float64)
		totalProb := 0.0

		for _, state := range pomdp.States {
			if pomdp.ObservationsProb[state] != nil && pomdp.ObservationsProb[state][action] != nil {
				for obs, prob := range pomdp.ObservationsProb[state][action] {
					obsDist[obs] += prob // Simplified aggregation
					totalProb += prob
				}
			}
		}

		// Normalize and compute entropy
		entropy := 0.0
		for _, prob := range obsDist {
			if prob > 0 {
				entropy -= prob * math.Log(prob)
			}
		}
		expectedEntropy += entropy
	}

	return expectedEntropy
}

func (ie *InferenceEngineImpl) computePolicyKLDivergence(pomdp *core.POMDP, policy core.Policy) float64 {
	// Simplified KL divergence for predicted vs prior beliefs
	prior := ie.getPrior(pomdp)

	// Predict next belief based on policy
	predictedBelief := ie.predictNextBelief(pomdp, prior, policy)

	return ie.computeKLDivergence(predictedBelief, prior)
}

func (ie *InferenceEngineImpl) predictNextBelief(pomdp *core.POMDP, currentBelief core.Belief, policy core.Policy) core.Belief {
	predicted := make(core.Belief, len(currentBelief))

	action := policy.GetAction(currentBelief)

	// Simple prediction: transition based on action
	for i, fromState := range pomdp.States {
		if pomdp.Transitions[fromState] != nil && pomdp.Transitions[fromState][action] != nil {
			for j, toState := range pomdp.States {
				if prob, exists := pomdp.Transitions[fromState][action][toState]; exists {
					predicted[j] += currentBelief[i] * prob
				}
			}
		}
	}

	return predicted
}
