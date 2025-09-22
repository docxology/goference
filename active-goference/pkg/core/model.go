package core

import (
	"errors"
	"fmt"
)

// POMDP represents a Partially Observable Markov Decision Process
// Note: Active inference doesn't use discount factors - this is a pure MDP concept
type POMDP struct {
	States           []string
	Actions          []string
	Observations     []string
	Transitions      map[string]map[string]map[string]float64 // s -> a -> s' -> prob
	ObservationsProb map[string]map[string]map[string]float64 // s' -> a -> o -> prob
	Rewards          map[string]map[string]float64            // s -> a -> reward
}

// NewPOMDP creates a new POMDP instance for active inference
func NewPOMDP(states, actions, observations []string) *POMDP {
	return &POMDP{
		States:           states,
		Actions:          actions,
		Observations:     observations,
		Transitions:      make(map[string]map[string]map[string]float64),
		ObservationsProb: make(map[string]map[string]map[string]float64),
		Rewards:          make(map[string]map[string]float64),
	}
}

// AddTransition adds a transition probability
func (p *POMDP) AddTransition(fromState, action, toState string, prob float64) error {
	if !contains(p.States, fromState) || !contains(p.Actions, action) || !contains(p.States, toState) {
		return errors.New("invalid state or action")
	}
	if p.Transitions[fromState] == nil {
		p.Transitions[fromState] = make(map[string]map[string]float64)
	}
	if p.Transitions[fromState][action] == nil {
		p.Transitions[fromState][action] = make(map[string]float64)
	}
	p.Transitions[fromState][action][toState] = prob
	return nil
}

// AddObservationProb adds an observation probability
func (p *POMDP) AddObservationProb(state, action, observation string, prob float64) error {
	if !contains(p.States, state) || !contains(p.Actions, action) || !contains(p.Observations, observation) {
		return errors.New("invalid state, action, or observation")
	}
	if p.ObservationsProb[state] == nil {
		p.ObservationsProb[state] = make(map[string]map[string]float64)
	}
	if p.ObservationsProb[state][action] == nil {
		p.ObservationsProb[state][action] = make(map[string]float64)
	}
	p.ObservationsProb[state][action][observation] = prob
	return nil
}

// AddReward adds a reward for a state-action pair
func (p *POMDP) AddReward(state, action string, reward float64) error {
	if !contains(p.States, state) || !contains(p.Actions, action) {
		return errors.New("invalid state or action")
	}
	if p.Rewards[state] == nil {
		p.Rewards[state] = make(map[string]float64)
	}
	p.Rewards[state][action] = reward
	return nil
}

// Belief represents a probability distribution over states
type Belief []float64

// NewBelief creates a new uniform belief over states
func NewBelief(states []string) Belief {
	n := len(states)
	belief := make(Belief, n)
	for i := range belief {
		belief[i] = 1.0 / float64(n)
	}
	return belief
}

// Normalize normalizes the belief distribution
func (b Belief) Normalize() {
	sum := 0.0
	for _, p := range b {
		sum += p
	}
	if sum > 0 {
		for i := range b {
			b[i] /= sum
		}
	}
}

// Policy represents a mapping from beliefs to actions
type Policy struct {
	Actions map[string]string // belief string -> action
}

// NewPolicy creates a new policy
func NewPolicy() Policy {
	return Policy{Actions: make(map[string]string)}
}

// GetAction returns the action for a given belief
func (pol Policy) GetAction(belief Belief) string {
	beliefStr := beliefToString(belief)
	return pol.Actions[beliefStr]
}

// SetAction sets the action for a given belief
func (pol Policy) SetAction(belief Belief, action string) {
	beliefStr := beliefToString(belief)
	pol.Actions[beliefStr] = action
}

// beliefToString converts a belief to a string representation
func beliefToString(belief Belief) string {
	str := ""
	for i, p := range belief {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprintf("%.3f", p)
	}
	return str
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func equalBeliefs(a, b Belief) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
