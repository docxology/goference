package utils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/wcharczuk/go-chart/v2"

	"active-goference/pkg/core"
)

// ChartVisualizer handles PNG chart generation for active inference visualizations
type ChartVisualizer struct {
	outputDir string
	width     int
	height    int
}

// NewChartVisualizer creates a new chart visualizer
func NewChartVisualizer(outputDir string, width, height int) *ChartVisualizer {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		Log.Error("Failed to create chart output directory:", err)
		return nil
	}

	return &ChartVisualizer{
		outputDir: outputDir,
		width:     width,
		height:    height,
	}
}

// GenerateVFECurve generates a PNG chart of VFE over time
func (cv *ChartVisualizer) GenerateVFECurve(vfeValues []float64, filename string) error {
	if len(vfeValues) == 0 {
		return fmt.Errorf("no VFE values to visualize")
	}

	graph := chart.Chart{
		Width:  cv.width,
		Height: cv.height,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			},
		},
		XAxis: chart.XAxis{
			Name: "Step",
			Range: &chart.ContinuousRange{
				Min: 1,
				Max: float64(len(vfeValues)),
			},
		},
		YAxis: chart.YAxis{
			Name: "Variational Free Energy (VFE)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "VFE",
				Style: chart.Style{
					StrokeColor: chart.ColorBlue,
					FillColor:   chart.ColorBlue.WithAlpha(64),
					StrokeWidth: 2,
				},
				XValues: func() []float64 {
					x := make([]float64, len(vfeValues))
					for i := 0; i < len(x); i++ {
						x[i] = float64(i + 1)
					}
					return x
				}(),
				YValues: vfeValues,
			},
		},
	}

	// Add title
	graph.Title = "Variational Free Energy Evolution"
	graph.TitleStyle = chart.Style{
		FontSize: 16,
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(cv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// GenerateEFECurve generates a PNG chart of EFE over time
func (cv *ChartVisualizer) GenerateEFECurve(efeValues []float64, filename string) error {
	if len(efeValues) == 0 {
		return fmt.Errorf("no EFE values to visualize")
	}

	graph := chart.Chart{
		Width:  cv.width,
		Height: cv.height,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			},
		},
		XAxis: chart.XAxis{
			Name: "Step",
			Range: &chart.ContinuousRange{
				Min: 1,
				Max: float64(len(efeValues)),
			},
		},
		YAxis: chart.YAxis{
			Name: "Expected Free Energy (EFE)",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "EFE",
				Style: chart.Style{
					StrokeColor: chart.ColorRed,
					FillColor:   chart.ColorRed.WithAlpha(64),
					StrokeWidth: 2,
				},
				XValues: func() []float64 {
					x := make([]float64, len(efeValues))
					for i := 0; i < len(x); i++ {
						x[i] = float64(i + 1)
					}
					return x
				}(),
				YValues: efeValues,
			},
		},
	}

	// Add title
	graph.Title = "Expected Free Energy Evolution"
	graph.TitleStyle = chart.Style{
		FontSize: 16,
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(cv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// GenerateBeliefEvolution generates a PNG chart of belief evolution over time
func (cv *ChartVisualizer) GenerateBeliefEvolution(beliefs []core.Belief, states []string, filename string) error {
	if len(beliefs) == 0 || len(states) == 0 {
		return fmt.Errorf("no beliefs or states to visualize")
	}

	graph := chart.Chart{
		Width:  cv.width,
		Height: cv.height,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			},
		},
		XAxis: chart.XAxis{
			Name: "Step",
			Range: &chart.ContinuousRange{
				Min: 1,
				Max: float64(len(beliefs)),
			},
		},
		YAxis: chart.YAxis{
			Name: "Belief Probability",
		},
	}

	// Create series for each state
	var series []chart.Series

	for stateIdx, state := range states {
		if stateIdx >= len(chart.DefaultColors) {
			break
		}

		var xValues []float64
		var yValues []float64

		for step, belief := range beliefs {
			if stateIdx < len(belief) {
				xValues = append(xValues, float64(step+1))
				yValues = append(yValues, belief[stateIdx])
			}
		}

		if len(xValues) > 0 {
			series = append(series, chart.ContinuousSeries{
				Name: state,
				Style: chart.Style{
					StrokeColor: chart.DefaultColors[stateIdx],
					StrokeWidth: 2,
				},
				XValues: xValues,
				YValues: yValues,
			})
		}
	}

	graph.Series = series
	graph.Title = "Belief Evolution Over Time"
	graph.TitleStyle = chart.Style{
		FontSize: 16,
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(cv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// GenerateMultiPanelDashboard generates a comprehensive multi-panel PNG dashboard
func (cv *ChartVisualizer) GenerateMultiPanelDashboard(
	pomdp *core.POMDP,
	beliefs []core.Belief,
	vfeValues []float64,
	efeValues []float64,
	observations []string,
	actions []string,
	filename string,
) error {
	// Create subplots - this would be more complex with the chart library
	// For now, create separate panels as individual charts

	// Generate VFE chart
	if len(vfeValues) > 0 {
		if err := cv.GenerateVFECurve(vfeValues, filename+"_vfe"); err != nil {
			Log.Error("Failed to generate VFE chart:", err)
		}
	}

	// Generate EFE chart
	if len(efeValues) > 0 {
		if err := cv.GenerateEFECurve(efeValues, filename+"_efe"); err != nil {
			Log.Error("Failed to generate EFE chart:", err)
		}
	}

	// Generate belief evolution chart
	if len(beliefs) > 0 {
		if err := cv.GenerateBeliefEvolution(beliefs, pomdp.States, filename+"_beliefs"); err != nil {
			Log.Error("Failed to generate belief evolution chart:", err)
		}
	}

	// Generate observations chart
	if len(observations) > 0 {
		if err := cv.GenerateObservationsChart(observations, filename+"_observations"); err != nil {
			Log.Error("Failed to generate observations chart:", err)
		}
	}

	// Generate actions chart
	if len(actions) > 0 {
		if err := cv.GenerateActionsChart(actions, filename+"_actions"); err != nil {
			Log.Error("Failed to generate actions chart:", err)
		}
	}

	return nil
}

// GenerateObservationsChart generates a chart of observation sequence
func (cv *ChartVisualizer) GenerateObservationsChart(observations []string, filename string) error {
	// Convert observations to numerical values for plotting
	obsMap := make(map[string]int)
	obsCount := 0
	var yValues []float64

	for _, obs := range observations {
		if _, exists := obsMap[obs]; !exists {
			obsMap[obs] = obsCount
			obsCount++
		}
		yValues = append(yValues, float64(obsMap[obs]))
	}

	graph := chart.Chart{
		Width:  cv.width,
		Height: cv.height,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			},
		},
		XAxis: chart.XAxis{
			Name: "Step",
			Range: &chart.ContinuousRange{
				Min: 1,
				Max: float64(len(observations)),
			},
		},
		YAxis: chart.YAxis{
			Name: "Observation Type",
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: float64(len(obsMap) - 1),
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "Observations",
				Style: chart.Style{
					StrokeColor: chart.ColorGreen,
					FillColor:   chart.ColorGreen.WithAlpha(64),
					StrokeWidth: 2,
				},
				XValues: func() []float64 {
					x := make([]float64, len(observations))
					for i := 0; i < len(x); i++ {
						x[i] = float64(i + 1)
					}
					return x
				}(),
				YValues: yValues,
			},
		},
	}

	graph.Title = "Observation Sequence"
	graph.TitleStyle = chart.Style{
		FontSize: 16,
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(cv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// GenerateActionsChart generates a chart of action sequence
func (cv *ChartVisualizer) GenerateActionsChart(actions []string, filename string) error {
	// Convert actions to numerical values for plotting
	actionMap := make(map[string]int)
	actionCount := 0
	var yValues []float64

	for _, action := range actions {
		if _, exists := actionMap[action]; !exists {
			actionMap[action] = actionCount
			actionCount++
		}
		yValues = append(yValues, float64(actionMap[action]))
	}

	graph := chart.Chart{
		Width:  cv.width,
		Height: cv.height,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			},
		},
		XAxis: chart.XAxis{
			Name: "Step",
			Range: &chart.ContinuousRange{
				Min: 1,
				Max: float64(len(actions)),
			},
		},
		YAxis: chart.YAxis{
			Name: "Action Type",
			Range: &chart.ContinuousRange{
				Min: 0,
				Max: float64(len(actionMap) - 1),
			},
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				Name: "Actions",
				Style: chart.Style{
					StrokeColor: chart.ColorOrange,
					FillColor:   chart.ColorOrange.WithAlpha(64),
					StrokeWidth: 2,
				},
				XValues: func() []float64 {
					x := make([]float64, len(actions))
					for i := 0; i < len(x); i++ {
						x[i] = float64(i + 1)
					}
					return x
				}(),
				YValues: yValues,
			},
		},
	}

	graph.Title = "Action Sequence"
	graph.TitleStyle = chart.Style{
		FontSize: 16,
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(cv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// GenerateStatisticalSummary generates a statistical summary as PNG
func (cv *ChartVisualizer) GenerateStatisticalSummary(
	beliefs []core.Belief,
	vfeValues []float64,
	efeValues []float64,
	observations []string,
	actions []string,
	filename string,
) error {
	// Create a simple bar chart with statistics
	graph := chart.Chart{
		Width:  cv.width,
		Height: cv.height,
		Background: chart.Style{
			Padding: chart.Box{
				Top:    20,
				Left:   20,
				Right:  20,
				Bottom: 20,
			},
		},
		XAxis: chart.XAxis{
			Name: "Metric",
		},
		YAxis: chart.YAxis{
			Name: "Value",
		},
	}

	// Create statistics
	var xValues []float64
	var yValues []float64
	var labels []string

	// Belief entropy statistics
	if len(beliefs) > 0 {
		avgEntropy := 0.0
		for _, belief := range beliefs {
			avgEntropy += calculateBeliefEntropy(belief)
		}
		avgEntropy /= float64(len(beliefs))

		xValues = append(xValues, 1.0)
		yValues = append(yValues, avgEntropy)
		labels = append(labels, "Avg Belief Entropy")
	}

	// Free energy statistics
	if len(vfeValues) > 0 {
		avgVFE := 0.0
		for _, vfe := range vfeValues {
			avgVFE += vfe
		}
		avgVFE /= float64(len(vfeValues))

		xValues = append(xValues, 2.0)
		yValues = append(yValues, avgVFE)
		labels = append(labels, "Avg VFE")
	}

	if len(efeValues) > 0 {
		avgEFE := 0.0
		for _, efe := range efeValues {
			avgEFE += efe
		}
		avgEFE /= float64(len(efeValues))

		xValues = append(xValues, 3.0)
		yValues = append(yValues, avgEFE)
		labels = append(labels, "Avg EFE")
	}

	// Unique observations/actions
	obsSet := make(map[string]bool)
	actionSet := make(map[string]bool)

	for _, obs := range observations {
		obsSet[obs] = true
	}
	for _, action := range actions {
		actionSet[action] = true
	}

	if len(obsSet) > 0 {
		xValues = append(xValues, 4.0)
		yValues = append(yValues, float64(len(obsSet)))
		labels = append(labels, "Unique Observations")
	}

	if len(actionSet) > 0 {
		xValues = append(xValues, 5.0)
		yValues = append(yValues, float64(len(actionSet)))
		labels = append(labels, "Unique Actions")
	}

	// Create simple chart
	if len(xValues) > 0 {
		graph.Series = []chart.Series{
			chart.ContinuousSeries{
				Name: "Statistics",
				Style: chart.Style{
					StrokeColor: chart.ColorBlue,
					FillColor:   chart.ColorBlue.WithAlpha(64),
					StrokeWidth: 2,
				},
				XValues: xValues,
				YValues: yValues,
			},
		}

		// Set discrete x-axis for labels
		graph.XAxis.ValueFormatter = func(v interface{}) string {
			if idx, ok := v.(float64); ok {
				idxInt := int(idx) - 1
				if idxInt >= 0 && idxInt < len(labels) {
					return labels[idxInt]
				}
			}
			return ""
		}
	}

	graph.Title = "Statistical Summary"
	graph.TitleStyle = chart.Style{
		FontSize: 16,
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(cv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer f.Close()

	return graph.Render(chart.PNG, f)
}

// Helper function to calculate belief entropy (chart version)
func calculateBeliefEntropy(belief core.Belief) float64 {
	entropy := 0.0
	for _, p := range belief {
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

// Global chart visualizer instance
var ChartVisualizerInstance = NewChartVisualizer("output/charts", 1200, 800)
