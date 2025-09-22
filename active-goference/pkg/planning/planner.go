package planning

import (
	"math"
	"math/rand"

	"active-goference/pkg/core"
	"active-goference/pkg/inference"
)

// PolicyPlannerImpl implements the PolicyPlanner interface
type PolicyPlannerImpl struct {
	Horizon         int
	Temperature     float64
	InferenceEngine *inference.InferenceEngineImpl
}

// NewPolicyPlanner creates a new policy planner
func NewPolicyPlanner(horizon int, temperature float64, engine *inference.InferenceEngineImpl) *PolicyPlannerImpl {
	return &PolicyPlannerImpl{
		Horizon:         horizon,
		Temperature:     temperature,
		InferenceEngine: engine,
	}
}

// PlanPolicy plans a policy based on current belief
func (pp *PolicyPlannerImpl) PlanPolicy(pomdp *core.POMDP, belief core.Belief) core.Policy {
	policy := core.NewPolicy()

	// Multi-step planning
	for step := 0; step < pp.Horizon; step++ {
		// Evaluate each possible action
		bestAction := pp.selectBestAction(pomdp, belief)
		policy.SetAction(belief, bestAction)

		// Update belief for next step (simplified)
		belief = pp.predictNextBelief(pomdp, belief, bestAction)
	}

	return policy
}

// selectBestAction selects the best action using expected free energy
// Public methods for testing
func (pp *PolicyPlannerImpl) SelectBestAction(pomdp *core.POMDP, belief core.Belief) string {
	return pp.selectBestAction(pomdp, belief)
}

func (pp *PolicyPlannerImpl) PredictNextBelief(pomdp *core.POMDP, currentBelief core.Belief, action string) core.Belief {
	return pp.predictNextBelief(pomdp, currentBelief, action)
}

func (pp *PolicyPlannerImpl) MergePolicies(policy1, policy2 core.Policy) core.Policy {
	return pp.mergePolicies(policy1, policy2)
}

// Private helper methods
func (pp *PolicyPlannerImpl) selectBestAction(pomdp *core.POMDP, belief core.Belief) string {
	bestAction := ""
	minEFE := math.Inf(1)

	for _, action := range pomdp.Actions {
		// Create policy for this action
		actionPolicy := core.NewPolicy()
		actionPolicy.SetAction(belief, action)

		// Compute expected free energy
		efe := pp.InferenceEngine.ComputeExpectedFreeEnergy(pomdp, actionPolicy)

		if efe < minEFE {
			minEFE = efe
			bestAction = action
		}
	}

	return bestAction
}

// predictNextBelief predicts the next belief state
func (pp *PolicyPlannerImpl) predictNextBelief(pomdp *core.POMDP, currentBelief core.Belief, action string) core.Belief {
	predicted := make(core.Belief, len(currentBelief))

	// Simple prediction based on transitions
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

// SoftmaxActionSelection performs softmax action selection
func (pp *PolicyPlannerImpl) SoftmaxActionSelection(pomdp *core.POMDP, belief core.Belief) string {
	actionScores := make(map[string]float64)

	// Compute scores for each action
	for _, action := range pomdp.Actions {
		actionPolicy := core.NewPolicy()
		actionPolicy.SetAction(belief, action)
		efe := pp.InferenceEngine.ComputeExpectedFreeEnergy(pomdp, actionPolicy)
		// Lower EFE is better, so negate for softmax
		actionScores[action] = -efe
	}

	// Compute softmax probabilities
	expSum := 0.0
	for _, score := range actionScores {
		expSum += math.Exp(score / pp.Temperature)
	}

	probabilities := make(map[string]float64)
	for action, score := range actionScores {
		probabilities[action] = math.Exp(score/pp.Temperature) / expSum
	}

	// Select action based on probabilities
	r := rand.Float64()
	cumulative := 0.0
	for action, prob := range probabilities {
		cumulative += prob
		if r <= cumulative {
			return action
		}
	}

	// Fallback to first action
	return pomdp.Actions[0]
}

// HierarchicalPlanning performs hierarchical planning
func (pp *PolicyPlannerImpl) HierarchicalPlanning(pomdp *core.POMDP, belief core.Belief, levels int) core.Policy {
	policy := core.NewPolicy()

	for level := 0; level < levels; level++ {
		// Plan at this level
		levelPolicy := pp.PlanPolicy(pomdp, belief)
		policy = pp.mergePolicies(policy, levelPolicy)
	}

	return policy
}

// mergePolicies merges two policies
func (pp *PolicyPlannerImpl) mergePolicies(policy1, policy2 core.Policy) core.Policy {
	merged := core.NewPolicy()

	// Simple merge: prioritize policy2
	for beliefStr, action := range policy1.Actions {
		merged.Actions[beliefStr] = action
	}
	for beliefStr, action := range policy2.Actions {
		merged.Actions[beliefStr] = action
	}

	return merged
}
