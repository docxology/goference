package tests

import (
	"testing"

	"active-goference/pkg/core"
	"active-goference/pkg/utils"
)

func TestNewLogger(t *testing.T) {
	logger := utils.NewLogger()

	if logger == nil {
		t.Error("Expected non-nil logger")
	}
}

func TestLoggerInfo(t *testing.T) {
	logger := utils.NewLogger()

	logger.Info("Test info message")

	// Cannot easily test log output, but ensure no panic
}

func TestLoggerError(t *testing.T) {
	logger := utils.NewLogger()

	logger.Error("Test error message")

	// Cannot easily test log output, but ensure no panic
}

func TestLoggerDebug(t *testing.T) {
	logger := utils.NewLogger()

	logger.Debug("Test debug message")

	// Cannot easily test log output, but ensure no panic
}

func TestLoggerWarn(t *testing.T) {
	logger := utils.NewLogger()

	logger.Warn("Test warning message")

	// Cannot easily test log output, but ensure no panic
}

func TestVisualizeBelief(t *testing.T) {
	belief := core.Belief{0.6, 0.4}
	states := []string{"s1", "s2"}

	visualization := utils.VisualizeBelief(belief, states)

	if len(visualization) == 0 {
		t.Error("Expected non-empty visualization")
	}

	if !containsString(visualization, "Belief State") {
		t.Error("Expected 'Belief State' in visualization")
	}
}

func TestVisualizePolicy(t *testing.T) {
	policy := core.NewPolicy()
	policy.SetAction(core.Belief{0.5, 0.5}, "a1")
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	visualization := utils.VisualizePolicy(policy, pomdp)

	if len(visualization) == 0 {
		t.Error("Expected non-empty visualization")
	}

	if !containsString(visualization, "Policy") {
		t.Error("Expected 'Policy' in visualization")
	}
}

func TestVisualizePOMDP(t *testing.T) {
	pomdp := core.NewPOMDP([]string{"s1", "s2"}, []string{"a1"}, []string{"o1"})

	visualization := utils.VisualizePOMDP(pomdp)

	if len(visualization) == 0 {
		t.Error("Expected non-empty visualization")
	}

	if !containsString(visualization, "POMDP Overview") {
		t.Error("Expected 'POMDP Overview' in visualization")
	}
}

func TestCreateBeliefHeatmap(t *testing.T) {
	beliefs := []core.Belief{
		{0.8, 0.2},
		{0.5, 0.5},
		{0.3, 0.7},
	}
	states := []string{"s1", "s2"}

	heatmap := utils.CreateBeliefHeatmap(beliefs, states)

	if len(heatmap) == 0 {
		t.Error("Expected non-empty heatmap")
	}

	if !containsString(heatmap, "Belief Evolution") {
		t.Error("Expected 'Belief Evolution' in heatmap")
	}
}

func TestVisualizeFreeEnergy(t *testing.T) {
	vfe := 0.5
	efe := 0.3

	visualization := utils.VisualizeFreeEnergy(vfe, efe)

	if len(visualization) == 0 {
		t.Error("Expected non-empty visualization")
	}

	if !containsString(visualization, "Free Energy") {
		t.Error("Expected 'Free Energy' in visualization")
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && s[:len(substr)] == substr ||
		len(s) > len(substr) && containsStringHelper(s, substr)
}

func containsStringHelper(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
