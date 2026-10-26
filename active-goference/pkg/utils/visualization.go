package utils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"active-goference/pkg/core"
)

// OutputManager handles all output operations to organized directories
type OutputManager struct {
	baseDir string
}

// NewOutputManager creates a new output manager
func NewOutputManager() *OutputManager {
	return &OutputManager{
		baseDir: "output",
	}
}

// EnsureDirectory ensures a directory exists
func (om *OutputManager) EnsureDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}

// SaveTextToFile saves text content to a file in the organized structure
func (om *OutputManager) SaveTextToFile(content, filename, category string) error {
	if err := om.EnsureDirectory(filepath.Join(om.baseDir, category)); err != nil {
		return err
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(om.baseDir, category, fmt.Sprintf("%s_%s.txt", filename, timestamp))

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}

// VisualizeBelief creates a text-based visualization of a belief state
func VisualizeBelief(belief core.Belief, states []string) string {
	var sb strings.Builder
	sb.WriteString("Belief State:\n")
	sb.WriteString("=============\n")

	maxProb := 0.0
	for _, p := range belief {
		if p > maxProb {
			maxProb = p
		}
	}

	for i, state := range states {
		prob := belief[i]
		barLength := int(prob / maxProb * 50)
		bar := strings.Repeat("█", barLength)
		sb.WriteString(fmt.Sprintf("%-15s: %6.4f %s\n", state, prob, bar))
	}

	return sb.String()
}

// SaveBeliefVisualization saves belief visualization to file
func (om *OutputManager) SaveBeliefVisualization(belief core.Belief, states []string, filename string) error {
	content := VisualizeBelief(belief, states)
	return om.SaveTextToFile(content, filename, "visualizations")
}

// VisualizePolicy creates a text-based visualization of a policy
func VisualizePolicy(policy core.Policy, pomdp *core.POMDP) string {
	var sb strings.Builder
	sb.WriteString("Policy:\n")
	sb.WriteString("=======\n")

	for beliefStr, action := range policy.Actions {
		sb.WriteString(fmt.Sprintf("Belief %s -> Action: %s\n", beliefStr, action))
	}

	return sb.String()
}

// SavePolicyVisualization saves policy visualization to file
func (om *OutputManager) SavePolicyVisualization(policy core.Policy, pomdp *core.POMDP, filename string) error {
	content := VisualizePolicy(policy, pomdp)
	return om.SaveTextToFile(content, filename, "visualizations")
}

// VisualizePOMDP creates a text-based overview of the POMDP
func VisualizePOMDP(pomdp *core.POMDP) string {
	var sb strings.Builder
	sb.WriteString("POMDP Overview (Active Inference):\n")
	sb.WriteString("==================================\n")
	sb.WriteString(fmt.Sprintf("States: %v\n", pomdp.States))
	sb.WriteString(fmt.Sprintf("Actions: %v\n", pomdp.Actions))
	sb.WriteString(fmt.Sprintf("Observations: %v\n", pomdp.Observations))
	sb.WriteString("Note: Active inference doesn't use discount factors\n")

	return sb.String()
}

// SavePOMDPVisualization saves POMDP overview to file
func (om *OutputManager) SavePOMDPVisualization(pomdp *core.POMDP, filename string) error {
	content := VisualizePOMDP(pomdp)
	return om.SaveTextToFile(content, filename, "visualizations")
}

// beliefToString converts a belief to a string representation
func beliefToString(belief core.Belief) string {
	str := ""
	for i, p := range belief {
		if i > 0 {
			str += ","
		}
		str += fmt.Sprintf("%.3f", p)
	}
	return str
}

// CreateBeliefHeatmap creates a heatmap representation of beliefs over time
func CreateBeliefHeatmap(beliefs []core.Belief, states []string) string {
	var sb strings.Builder
	sb.WriteString("Belief Evolution Heatmap:\n")
	sb.WriteString("========================\n")

	for t, belief := range beliefs {
		sb.WriteString(fmt.Sprintf("t=%d: ", t))
		for _, p := range belief {
			if p > 0.5 {
				sb.WriteString("██")
			} else if p > 0.2 {
				sb.WriteString("▓▓")
			} else {
				sb.WriteString("░░")
			}
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

// SaveBeliefHeatmap saves belief heatmap to file
func (om *OutputManager) SaveBeliefHeatmap(beliefs []core.Belief, states []string, filename string) error {
	content := CreateBeliefHeatmap(beliefs, states)
	return om.SaveTextToFile(content, filename, "visualizations")
}

// VisualizeFreeEnergy creates a visualization of free energy values
func VisualizeFreeEnergy(vfe, efe float64) string {
	var sb strings.Builder
	sb.WriteString("Active Inference Free Energy Analysis:\n")
	sb.WriteString("=====================================\n")
	sb.WriteString(fmt.Sprintf("Variational Free Energy (VFE): %.6f\n", vfe))
	sb.WriteString("  - Used for belief updating from observations\n")
	sb.WriteString(fmt.Sprintf("Expected Free Energy (EFE): %.6f\n", efe))
	sb.WriteString("  - Used for policy planning and decision making\n")
	sb.WriteString("Note: VFE and EFE serve different purposes in active inference\n")
	sb.WriteString("      They are not summed as they represent different processes\n")

	return sb.String()
}

// SaveFreeEnergyVisualization saves free energy analysis to file
func (om *OutputManager) SaveFreeEnergyVisualization(vfe, efe float64, filename string) error {
	content := VisualizeFreeEnergy(vfe, efe)
	return om.SaveTextToFile(content, filename, "visualizations")
}

// CreateDetailedTrace creates a detailed trace of active inference execution
func CreateDetailedTrace(pomdp *core.POMDP, belief core.Belief, policy core.Policy, vfe, efe float64, step int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("=== Active Inference Step %d ===\n", step))
	sb.WriteString(fmt.Sprintf("Timestamp: %s\n", time.Now().Format("2006-01-02 15:04:05")))
	sb.WriteString("\n--- POMDP Configuration ---\n")
	sb.WriteString(VisualizePOMDP(pomdp))
	sb.WriteString("\n--- Current Belief ---\n")
	sb.WriteString(VisualizeBelief(belief, pomdp.States))
	sb.WriteString("\n--- Policy ---\n")
	sb.WriteString(VisualizePolicy(policy, pomdp))
	sb.WriteString("\n--- Free Energy Analysis ---\n")
	sb.WriteString(VisualizeFreeEnergy(vfe, efe))
	sb.WriteString("\n--- Mathematical Details ---\n")
	sb.WriteString(fmt.Sprintf("Belief Entropy: %.6f\n", beliefEntropy(belief)))
	sb.WriteString("\n==========================\n")

	return sb.String()
}

// SaveDetailedTrace saves detailed trace to file
func (om *OutputManager) SaveDetailedTrace(pomdp *core.POMDP, belief core.Belief, policy core.Policy, vfe, efe float64, step int, filename string) error {
	content := CreateDetailedTrace(pomdp, belief, policy, vfe, efe, step)
	return om.SaveTextToFile(content, filename, "traces")
}

// Helper functions for trace calculations
func beliefEntropy(belief core.Belief) float64 {
	entropy := 0.0
	for _, p := range belief {
		if p > 0 {
			entropy -= p * log2(p)
		}
	}
	return entropy
}

func log2(x float64) float64 {
	return math.Log2(x)
}

// CreateComprehensiveReport creates a comprehensive execution report
func (om *OutputManager) CreateComprehensiveReport(sessionName string, results map[string]interface{}) error {
	var sb strings.Builder
	sb.WriteString("Active Goference Comprehensive Report\n")
	sb.WriteString("=====================================\n")
	sb.WriteString(fmt.Sprintf("Session: %s\n", sessionName))
	sb.WriteString(fmt.Sprintf("Generated: %s\n", time.Now().Format("2006-01-02 15:04:05")))
	sb.WriteString("\n--- Session Summary ---\n")

	for key, value := range results {
		sb.WriteString(fmt.Sprintf("%s: %v\n", key, value))
	}

	return om.SaveTextToFile(sb.String(), "comprehensive_report", "reports")
}

// Global output manager instance
var Output = NewOutputManager()
