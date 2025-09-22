package belief

import (
	"math"

	"active-goference/pkg/core"

	"gonum.org/v1/gonum/stat/distuv"
)

// CategoricalBelief represents a categorical distribution over discrete states
type CategoricalBelief struct {
	Probabilities []float64
	States        []string
}

// NewCategoricalBelief creates a new categorical belief
func NewCategoricalBelief(states []string) *CategoricalBelief {
	n := len(states)
	probs := make([]float64, n)
	for i := range probs {
		probs[i] = 1.0 / float64(n)
	}
	return &CategoricalBelief{
		Probabilities: probs,
		States:        states,
	}
}

// Normalize normalizes the probabilities
func (cb *CategoricalBelief) Normalize() {
	sum := 0.0
	for _, p := range cb.Probabilities {
		sum += p
	}
	if sum > 0 {
		for i := range cb.Probabilities {
			cb.Probabilities[i] /= sum
		}
	}
}

// Update updates the belief based on likelihood
func (cb *CategoricalBelief) Update(likelihood []float64) {
	if len(likelihood) != len(cb.Probabilities) {
		return
	}
	for i := range cb.Probabilities {
		cb.Probabilities[i] *= likelihood[i]
	}
	cb.Normalize()
}

// GetProbability returns the probability of a state
func (cb *CategoricalBelief) GetProbability(state string) float64 {
	for i, s := range cb.States {
		if s == state {
			return cb.Probabilities[i]
		}
	}
	return 0.0
}

// Entropy calculates the entropy of the belief
func (cb *CategoricalBelief) Entropy() float64 {
	entropy := 0.0
	for _, p := range cb.Probabilities {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

// GaussianBelief represents a Gaussian distribution over continuous states
type GaussianBelief struct {
	Mean     []float64
	Variance []float64
	Dim      int
}

// NewGaussianBelief creates a new Gaussian belief
func NewGaussianBelief(dim int) *GaussianBelief {
	mean := make([]float64, dim)
	variance := make([]float64, dim)
	for i := range variance {
		variance[i] = 1.0 // Default variance
	}
	return &GaussianBelief{
		Mean:     mean,
		Variance: variance,
		Dim:      dim,
	}
}

// Update updates the Gaussian belief using Kalman filter-like update
func (gb *GaussianBelief) Update(observation []float64, observationNoise float64) {
	if len(observation) != gb.Dim {
		return
	}
	// Simple update assuming identity observation matrix
	for i := range gb.Mean {
		k := gb.Variance[i] / (gb.Variance[i] + observationNoise)
		gb.Mean[i] = gb.Mean[i] + k*(observation[i]-gb.Mean[i])
		gb.Variance[i] = (1 - k) * gb.Variance[i]
	}
}

// Entropy calculates the differential entropy
func (gb *GaussianBelief) Entropy() float64 {
	entropy := 0.0
	for _, v := range gb.Variance {
		entropy += 0.5 * math.Log2(2*math.Pi*math.E*v)
	}
	return entropy
}

// DirichletBelief represents a Dirichlet distribution for parameter learning
type DirichletBelief struct {
	Concentration []float64
	Dim           int
}

// NewDirichletBelief creates a new Dirichlet belief
func NewDirichletBelief(dim int, concentration float64) *DirichletBelief {
	conc := make([]float64, dim)
	for i := range conc {
		conc[i] = concentration
	}
	return &DirichletBelief{
		Concentration: conc,
		Dim:           dim,
	}
}

// Update updates the Dirichlet belief with new observations
func (db *DirichletBelief) Update(observations []int) {
	for _, obs := range observations {
		if obs >= 0 && obs < db.Dim {
			db.Concentration[obs] += 1.0
		}
	}
}

// Sample samples from the Dirichlet distribution
func (db *DirichletBelief) Sample() []float64 {
	sample := make([]float64, db.Dim)
	for i := range sample {
		gamma := distuv.Gamma{Alpha: db.Concentration[i], Beta: 1.0}
		sample[i] = gamma.Rand()
	}
	// Normalize
	sum := 0.0
	for _, s := range sample {
		sum += s
	}
	for i := range sample {
		sample[i] /= sum
	}
	return sample
}

// Mean returns the mean of the Dirichlet distribution
func (db *DirichletBelief) Mean() []float64 {
	mean := make([]float64, db.Dim)
	sumConc := 0.0
	for _, c := range db.Concentration {
		sumConc += c
	}
	for i, c := range db.Concentration {
		mean[i] = c / sumConc
	}
	return mean
}

// BeliefUpdaterImpl implements the BeliefUpdater interface
type BeliefUpdaterImpl struct{}

// NewBeliefUpdater creates a new belief updater
func NewBeliefUpdater() *BeliefUpdaterImpl {
	return &BeliefUpdaterImpl{}
}

// UpdateBelief updates the belief based on observation
func (bu *BeliefUpdaterImpl) UpdateBelief(prior core.Belief, observation string, action string, pomdp *core.POMDP) core.Belief {
	newBelief := make(core.Belief, len(prior))
	copy(newBelief, prior)

	// Compute likelihood based on observation model
	for i, state := range pomdp.States {
		likelihood := 1.0
		if pomdp.ObservationsProb[state] != nil && pomdp.ObservationsProb[state][action] != nil {
			likelihood = pomdp.ObservationsProb[state][action][observation]
		}
		newBelief[i] *= likelihood
	}

	// Normalize
	newBelief.Normalize()
	return newBelief
}
