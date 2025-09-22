package tests

import (
	"math"
	"testing"

	"active-goference/pkg/belief"
	"active-goference/pkg/core"
)

func TestNewCategoricalBelief(t *testing.T) {
	states := []string{"s1", "s2", "s3"}
	cb := belief.NewCategoricalBelief(states)

	if len(cb.Probabilities) != 3 {
		t.Errorf("Expected 3 probabilities, got %d", len(cb.Probabilities))
	}

	expected := 1.0 / 3.0
	for i, p := range cb.Probabilities {
		if math.Abs(p-expected) > 1e-6 {
			t.Errorf("Expected probability[%d] = %f, got %f", i, expected, p)
		}
	}
}

func TestCategoricalBeliefNormalize(t *testing.T) {
	cb := &belief.CategoricalBelief{
		Probabilities: []float64{0.3, 0.4, 0.2},
		States:        []string{"s1", "s2", "s3"},
	}
	cb.Normalize()

	sum := 0.0
	for _, p := range cb.Probabilities {
		sum += p
	}

	if math.Abs(sum-1.0) > 1e-6 {
		t.Errorf("Expected sum 1.0, got %f", sum)
	}
}

func TestCategoricalBeliefUpdate(t *testing.T) {
	cb := belief.NewCategoricalBelief([]string{"s1", "s2", "s3"})
	likelihood := []float64{0.5, 0.3, 0.2}

	cb.Update(likelihood)

	// Check that probabilities are updated and normalized
	if math.Abs(cb.Probabilities[0]-0.5) > 1e-6 {
		t.Errorf("Expected updated probability 0.5, got %f", cb.Probabilities[0])
	}
}

func TestCategoricalBeliefEntropy(t *testing.T) {
	cb := &belief.CategoricalBelief{
		Probabilities: []float64{0.5, 0.5},
		States:        []string{"s1", "s2"},
	}

	entropy := cb.Entropy()
	expected := -0.5*math.Log2(0.5) - 0.5*math.Log2(0.5)

	if math.Abs(entropy-expected) > 1e-6 {
		t.Errorf("Expected entropy %f, got %f", expected, entropy)
	}
}

func TestNewGaussianBelief(t *testing.T) {
	gb := belief.NewGaussianBelief(2)

	if len(gb.Mean) != 2 {
		t.Errorf("Expected 2 dimensions, got %d", len(gb.Mean))
	}

	if len(gb.Variance) != 2 {
		t.Errorf("Expected 2 variances, got %d", len(gb.Variance))
	}
}

func TestGaussianBeliefUpdate(t *testing.T) {
	gb := belief.NewGaussianBelief(1)
	observation := []float64{1.0}
	observationNoise := 0.1

	gb.Update(observation, observationNoise)

	// Check that mean is updated according to Kalman filter
	// k = 1.0 / (1.0 + 0.1) = 0.909
	// new_mean = 0.0 + 0.909 * (1.0 - 0.0) = 0.909
	expectedMean := 0.9090909090909091
	if math.Abs(gb.Mean[0]-expectedMean) > 1e-6 {
		t.Errorf("Expected mean %.6f, got %f", expectedMean, gb.Mean[0])
	}
}

func TestGaussianBeliefEntropy(t *testing.T) {
	gb := &belief.GaussianBelief{
		Mean:     []float64{0.0},
		Variance: []float64{1.0},
		Dim:      1,
	}

	entropy := gb.Entropy()
	expected := 0.5 * math.Log2(2*math.Pi*math.E*1.0)

	if math.Abs(entropy-expected) > 1e-6 {
		t.Errorf("Expected entropy %f, got %f", expected, entropy)
	}
}

func TestNewDirichletBelief(t *testing.T) {
	db := belief.NewDirichletBelief(3, 1.0)

	if len(db.Concentration) != 3 {
		t.Errorf("Expected 3 concentrations, got %d", len(db.Concentration))
	}

	for i, c := range db.Concentration {
		if c != 1.0 {
			t.Errorf("Expected concentration[%d] = 1.0, got %f", i, c)
		}
	}
}

func TestDirichletBeliefUpdate(t *testing.T) {
	db := belief.NewDirichletBelief(3, 1.0)
	observations := []int{0, 1, 0}

	db.Update(observations)

	if db.Concentration[0] != 3.0 {
		t.Errorf("Expected concentration[0] = 3.0, got %f", db.Concentration[0])
	}

	if db.Concentration[1] != 2.0 {
		t.Errorf("Expected concentration[1] = 2.0, got %f", db.Concentration[1])
	}
}

func TestDirichletBeliefMean(t *testing.T) {
	db := &belief.DirichletBelief{
		Concentration: []float64{2.0, 3.0},
		Dim:           2,
	}

	mean := db.Mean()
	sumConc := 5.0

	if math.Abs(mean[0]-2.0/sumConc) > 1e-6 {
		t.Errorf("Expected mean[0] = %f, got %f", 2.0/sumConc, mean[0])
	}

	if math.Abs(mean[1]-3.0/sumConc) > 1e-6 {
		t.Errorf("Expected mean[1] = %f, got %f", 3.0/sumConc, mean[1])
	}
}

func TestNewBeliefUpdater(t *testing.T) {
	updater := belief.NewBeliefUpdater()

	if updater == nil {
		t.Error("Expected non-nil updater")
	}
}

func TestBeliefUpdaterUpdateBelief(t *testing.T) {
	updater := belief.NewBeliefUpdater()
	pomdp := &core.POMDP{
		States:       []string{"s1", "s2"},
		Observations: []string{"o1"},
		ObservationsProb: map[string]map[string]map[string]float64{
			"s1": {"a1": {"o1": 0.8}},
			"s2": {"a1": {"o1": 0.2}},
		},
	}

	prior := core.Belief{0.5, 0.5}
	observation := "o1"
	action := "a1"

	newBelief := updater.UpdateBelief(prior, observation, action, pomdp)

	if len(newBelief) != 2 {
		t.Errorf("Expected 2 states, got %d", len(newBelief))
	}

	// Check that belief is updated based on observation probabilities
	expected1 := 0.5 * 0.8 / (0.5*0.8 + 0.5*0.2)
	expected2 := 0.5 * 0.2 / (0.5*0.8 + 0.5*0.2)

	if math.Abs(newBelief[0]-expected1) > 1e-6 {
		t.Errorf("Expected newBelief[0] = %f, got %f", expected1, newBelief[0])
	}

	if math.Abs(newBelief[1]-expected2) > 1e-6 {
		t.Errorf("Expected newBelief[1] = %f, got %f", expected2, newBelief[1])
	}
}
