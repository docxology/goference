package tests

import (
	"fmt"
	"testing"

	"active-goference/pkg/core"
)

func BenchmarkNewPOMDP(b *testing.B) {
	states := []string{"s1", "s2", "s3", "s4", "s5"}
	actions := []string{"a1", "a2", "a3"}
	observations := []string{"o1", "o2", "o3", "o4"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = core.NewPOMDP(states, actions, observations)
	}
}

func BenchmarkAddTransition(b *testing.B) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pomdp.AddTransition("s1", "a1", "s2", 0.8)
	}
}

func BenchmarkBeliefNormalize(b *testing.B) {
	belief := core.Belief{0.3, 0.4, 0.2, 0.1}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		belief.Normalize()
	}
}

func BenchmarkPolicySetAction(b *testing.B) {
	policy := core.NewPolicy()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		belief := core.Belief{0.2, 0.3, 0.5}
		policy.SetAction(belief, "a1")
	}
}

func BenchmarkPolicyGetAction(b *testing.B) {
	policy := core.NewPolicy()
	belief := core.Belief{0.2, 0.3, 0.5}
	policy.SetAction(belief, "a1")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = policy.GetAction(belief)
	}
}

func BenchmarkNewBelief(b *testing.B) {
	states := []string{"s1", "s2", "s3", "s4", "s5", "s6", "s7", "s8", "s9", "s10"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = core.NewBelief(states)
	}
}

func BenchmarkPOMDPWithLargeStateSpace(b *testing.B) {
	// Create a larger POMDP for stress testing
	numStates := 100
	states := make([]string, numStates)
	actions := []string{"a1", "a2", "a3", "a4", "a5"}
	observations := make([]string, numStates)

	for i := 0; i < numStates; i++ {
		states[i] = fmt.Sprintf("s%d", i)
		observations[i] = fmt.Sprintf("o%d", i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pomdp := core.NewPOMDP(states, actions, observations)
		// Add some transitions
		for j := 0; j < 10; j++ {
			pomdp.AddTransition(states[j], actions[0], states[(j+1)%numStates], 0.1)
		}
	}
}

func BenchmarkBeliefOperations(b *testing.B) {
	belief := core.NewBelief([]string{"s1", "s2", "s3", "s4", "s5"})

	b.Run("Normalize", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			testBelief := make(core.Belief, len(belief))
			copy(testBelief, belief)
			testBelief.Normalize()
		}
	})

	b.Run("Copy", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			newBelief := make(core.Belief, len(belief))
			copy(newBelief, belief)
		}
	})
}

func BenchmarkPolicyOperations(b *testing.B) {
	policy := core.NewPolicy()

	b.Run("SetMultipleActions", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			for j := 0; j < 10; j++ {
				belief := core.Belief{0.1 * float64(j+1), 0.9 / float64(j+1)}
				policy.SetAction(belief, fmt.Sprintf("a%d", j))
			}
		}
	})

	b.Run("GetMultipleActions", func(b *testing.B) {
		// First populate the policy
		for j := 0; j < 10; j++ {
			belief := core.Belief{0.1 * float64(j+1), 0.9 / float64(j+1)}
			policy.SetAction(belief, fmt.Sprintf("a%d", j))
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for j := 0; j < 10; j++ {
				belief := core.Belief{0.1 * float64(j+1), 0.9 / float64(j+1)}
				_ = policy.GetAction(belief)
			}
		}
	})
}
