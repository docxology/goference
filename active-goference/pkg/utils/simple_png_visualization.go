package utils

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/fogleman/gg"

	"active-goference/pkg/core"
)

// SimplePNGVisualizer provides basic PNG visualizations for active inference
type SimplePNGVisualizer struct {
	outputDir string
	width     int
	height    int
}

// NewSimplePNGVisualizer creates a new simple PNG visualizer
func NewSimplePNGVisualizer(outputDir string, width, height int) *SimplePNGVisualizer {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		Log.Error("Failed to create PNG output directory:", err)
		return nil
	}

	return &SimplePNGVisualizer{
		outputDir: outputDir,
		width:     width,
		height:    height,
	}
}

// GenerateVFECurve generates a simple line chart of VFE over time
func (spv *SimplePNGVisualizer) GenerateVFECurve(vfeValues []float64, filename string) error {
	if len(vfeValues) == 0 {
		return fmt.Errorf("no VFE values to visualize")
	}

	dc := gg.NewContext(spv.width, spv.height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	margin := 80
	plotWidth := spv.width - 2*margin
	plotHeight := spv.height - 2*margin

	// Find data range for proper scaling
	minVal, maxVal := vfeValues[0], vfeValues[0]
	for _, v := range vfeValues {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	rangeVal := maxVal - minVal
	if rangeVal == 0 {
		rangeVal = 1
	}
	minVal -= rangeVal * 0.1
	maxVal += rangeVal * 0.1

	// Draw grid lines
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.SetLineWidth(1)
	for i := 0; i <= 10; i++ {
		x := float64(margin + i*plotWidth/10)
		y := float64(margin + i*plotHeight/10)
		dc.DrawLine(x, float64(margin), x, float64(margin+plotHeight))
		dc.DrawLine(float64(margin), y, float64(margin+plotWidth), y)
	}
	dc.Stroke()

	// Draw axes
	dc.SetRGB(0.3, 0.3, 0.3)
	dc.SetLineWidth(2)
	dc.DrawLine(float64(margin), float64(margin), float64(margin), float64(margin+plotHeight))
	dc.DrawLine(float64(margin), float64(margin+plotHeight), float64(margin+plotWidth), float64(margin+plotHeight))
	dc.Stroke()

	// Draw axis labels and ticks
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)

	// Y-axis ticks and labels
	for i := 0; i <= 5; i++ {
		y := float64(margin + plotHeight - i*plotHeight/5)
		val := minVal + float64(i)*(maxVal-minVal)/5

		// Tick mark
		dc.SetLineWidth(1)
		dc.DrawLine(float64(margin-3), y, float64(margin+3), y)
		dc.Stroke()

		// Label
		dc.DrawStringAnchored(fmt.Sprintf("%.3f", val), float64(margin-8), y, 1, 0.5)
	}

	// X-axis ticks and labels
	for i := 0; i <= len(vfeValues)-1; i++ {
		if i%1 == 0 { // Label every step
			x := float64(margin + i*plotWidth/(len(vfeValues)-1))

			// Tick mark
			dc.SetLineWidth(1)
			dc.DrawLine(x, float64(margin+plotHeight-3), x, float64(margin+plotHeight+3))
			dc.Stroke()

			// Label
			dc.DrawStringAnchored(fmt.Sprintf("%d", i+1), x, float64(margin+plotHeight+12), 0.5, 0)
		}
	}

	// Axis titles
	dc.LoadFontFace("Arial", 12)
	dc.DrawStringAnchored("Simulation Step", float64(margin+plotWidth/2), float64(margin+plotHeight+35), 0.5, 0.5)
	dc.DrawStringAnchored("Variational Free Energy (VFE)", float64(margin-45), float64(margin+plotHeight/2), 0.5, 0.5)
	dc.RotateAbout(-math.Pi/2, float64(margin-45), float64(margin+plotHeight/2))
	dc.DrawStringAnchored("Variational Free Energy (VFE)", 0, 0, 0.5, 0.5)
	dc.RotateAbout(math.Pi/2, float64(margin-45), float64(margin+plotHeight/2))

	// Plot VFE curve with area fill
	if len(vfeValues) > 0 {
		// Fill area under curve
		dc.SetRGB(1, 0.8, 0.8) // Light red fill
		dc.MoveTo(float64(margin), float64(margin+plotHeight))
		for i, vfe := range vfeValues {
			x := float64(margin + i*plotWidth/(len(vfeValues)-1))
			y := float64(margin + plotHeight - int((vfe-minVal)/(maxVal-minVal)*float64(plotHeight)))
			dc.LineTo(x, y)
		}
		dc.LineTo(float64(margin+plotWidth), float64(margin+plotHeight))
		dc.ClosePath()
		dc.Fill()

		// Draw line
		dc.SetRGB(1, 0, 0) // Red for VFE line
		dc.SetLineWidth(3)
		for i := 0; i < len(vfeValues)-1; i++ {
			x1 := float64(margin + i*plotWidth/(len(vfeValues)-1))
			y1 := float64(margin + plotHeight - int((vfeValues[i]-minVal)/(maxVal-minVal)*float64(plotHeight)))
			x2 := float64(margin + (i+1)*plotWidth/(len(vfeValues)-1))
			y2 := float64(margin + plotHeight - int((vfeValues[i+1]-minVal)/(maxVal-minVal)*float64(plotHeight)))

			dc.DrawLine(x1, y1, x2, y2)
			dc.Stroke()
		}

		// Draw points with values
		dc.SetRGB(0.8, 0, 0)
		for i, vfe := range vfeValues {
			x := float64(margin + i*plotWidth/(len(vfeValues)-1))
			y := float64(margin + plotHeight - int((vfe-minVal)/(maxVal-minVal)*float64(plotHeight)))
			dc.DrawCircle(x, y, 4)
			dc.Fill()

			// Value label for final point
			if i == len(vfeValues)-1 {
				dc.SetRGB(0.5, 0.5, 0.5)
				dc.LoadFontFace("Arial", 8)
				dc.DrawStringAnchored(fmt.Sprintf("%.3f", vfe), x+8, y-8, 0, 1)
			}
		}
	}

	// Legend
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)

	// Legend box
	legendX, legendY := float64(spv.width-150), float64(30)
	dc.SetRGB(0.95, 0.95, 0.95)
	dc.DrawRectangle(legendX-5, legendY-5, 140, 30)
	dc.Fill()

	dc.SetRGB(0.3, 0.3, 0.3)
	dc.DrawRectangle(legendX-5, legendY-5, 140, 30)
	dc.Stroke()

	// Legend content
	dc.SetRGB(1, 0, 0)
	dc.DrawRectangle(legendX, legendY, 12, 12)
	dc.Fill()
	dc.SetRGB(0, 0, 0)
	dc.DrawString("VFE (Belief Precision)", legendX+18, legendY+10)

	// Title
	dc.LoadFontFace("Arial", 16)
	dc.DrawStringAnchored("Active Inference: Variational Free Energy Evolution", float64(spv.width/2), 20, 0.5, 0.5)

	// Subtitle
	dc.LoadFontFace("Arial", 10)
	dc.SetRGB(0.5, 0.5, 0.5)
	dc.DrawStringAnchored("Measures how well beliefs explain observations - increases with certainty", float64(spv.width/2), 40, 0.5, 0.5)

	// Save the image
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(spv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	return dc.SavePNG(fullPath)
}

// GenerateEFECurve generates a simple line chart of EFE over time
func (spv *SimplePNGVisualizer) GenerateEFECurve(efeValues []float64, filename string) error {
	if len(efeValues) == 0 {
		return fmt.Errorf("no EFE values to visualize")
	}

	dc := gg.NewContext(spv.width, spv.height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	margin := 80
	plotWidth := spv.width - 2*margin
	plotHeight := spv.height - 2*margin

	// Find data range for proper scaling
	minVal, maxVal := efeValues[0], efeValues[0]
	for _, v := range efeValues {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	rangeVal := maxVal - minVal
	if rangeVal == 0 {
		rangeVal = 0.01 // Handle flat EFE curves
	}
	minVal -= rangeVal * 0.1
	maxVal += rangeVal * 0.1

	// Draw grid lines
	dc.SetRGB(0.9, 0.9, 0.9)
	dc.SetLineWidth(1)
	for i := 0; i <= 10; i++ {
		x := float64(margin + i*plotWidth/10)
		y := float64(margin + i*plotHeight/10)
		dc.DrawLine(x, float64(margin), x, float64(margin+plotHeight))
		dc.DrawLine(float64(margin), y, float64(margin+plotWidth), y)
	}
	dc.Stroke()

	// Draw axes
	dc.SetRGB(0.3, 0.3, 0.3)
	dc.SetLineWidth(2)
	dc.DrawLine(float64(margin), float64(margin), float64(margin), float64(margin+plotHeight))
	dc.DrawLine(float64(margin), float64(margin+plotHeight), float64(margin+plotWidth), float64(margin+plotHeight))
	dc.Stroke()

	// Draw axis labels and ticks
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)

	// Y-axis ticks and labels
	for i := 0; i <= 5; i++ {
		y := float64(margin + plotHeight - i*plotHeight/5)
		val := minVal + float64(i)*(maxVal-minVal)/5

		// Tick mark
		dc.SetLineWidth(1)
		dc.DrawLine(float64(margin-3), y, float64(margin+3), y)
		dc.Stroke()

		// Label
		dc.DrawStringAnchored(fmt.Sprintf("%.4f", val), float64(margin-8), y, 1, 0.5)
	}

	// X-axis ticks and labels
	for i := 0; i <= len(efeValues)-1; i++ {
		if i%1 == 0 { // Label every step
			x := float64(margin + i*plotWidth/(len(efeValues)-1))

			// Tick mark
			dc.SetLineWidth(1)
			dc.DrawLine(x, float64(margin+plotHeight-3), x, float64(margin+plotHeight+3))
			dc.Stroke()

			// Label
			dc.DrawStringAnchored(fmt.Sprintf("%d", i+1), x, float64(margin+plotHeight+12), 0.5, 0)
		}
	}

	// Axis titles
	dc.LoadFontFace("Arial", 12)
	dc.DrawStringAnchored("Simulation Step", float64(margin+plotWidth/2), float64(margin+plotHeight+35), 0.5, 0.5)
	dc.DrawStringAnchored("Expected Free Energy (EFE)", float64(margin-50), float64(margin+plotHeight/2), 0.5, 0.5)
	dc.RotateAbout(-math.Pi/2, float64(margin-50), float64(margin+plotHeight/2))
	dc.DrawStringAnchored("Expected Free Energy (EFE)", 0, 0, 0.5, 0.5)
	dc.RotateAbout(math.Pi/2, float64(margin-50), float64(margin+plotHeight/2))

	// Plot EFE curve with area fill
	if len(efeValues) > 0 {
		// Fill area under curve
		dc.SetRGB(0.8, 0.8, 1) // Light blue fill
		dc.MoveTo(float64(margin), float64(margin+plotHeight))
		for i, efe := range efeValues {
			x := float64(margin + i*plotWidth/(len(efeValues)-1))
			y := float64(margin + plotHeight - int((efe-minVal)/(maxVal-minVal)*float64(plotHeight)))
			dc.LineTo(x, y)
		}
		dc.LineTo(float64(margin+plotWidth), float64(margin+plotHeight))
		dc.ClosePath()
		dc.Fill()

		// Draw line
		dc.SetRGB(0, 0, 1) // Blue for EFE line
		dc.SetLineWidth(3)
		for i := 0; i < len(efeValues)-1; i++ {
			x1 := float64(margin + i*plotWidth/(len(efeValues)-1))
			y1 := float64(margin + plotHeight - int((efeValues[i]-minVal)/(maxVal-minVal)*float64(plotHeight)))
			x2 := float64(margin + (i+1)*plotWidth/(len(efeValues)-1))
			y2 := float64(margin + plotHeight - int((efeValues[i+1]-minVal)/(maxVal-minVal)*float64(plotHeight)))

			dc.DrawLine(x1, y1, x2, y2)
			dc.Stroke()
		}

		// Draw points with values
		dc.SetRGB(0, 0, 0.8)
		for i, efe := range efeValues {
			x := float64(margin + i*plotWidth/(len(efeValues)-1))
			y := float64(margin + plotHeight - int((efe-minVal)/(maxVal-minVal)*float64(plotHeight)))
			dc.DrawCircle(x, y, 4)
			dc.Fill()

			// Value label for final point
			if i == len(efeValues)-1 {
				dc.SetRGB(0.5, 0.5, 0.5)
				dc.LoadFontFace("Arial", 8)
				dc.DrawStringAnchored(fmt.Sprintf("%.4f", efe), x+8, y-8, 0, 1)
			}
		}
	}

	// Legend
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)

	// Legend box
	legendX, legendY := float64(spv.width-160), float64(30)
	dc.SetRGB(0.95, 0.95, 0.95)
	dc.DrawRectangle(legendX-5, legendY-5, 150, 30)
	dc.Fill()

	dc.SetRGB(0.3, 0.3, 0.3)
	dc.DrawRectangle(legendX-5, legendY-5, 150, 30)
	dc.Stroke()

	// Legend content
	dc.SetRGB(0, 0, 1)
	dc.DrawRectangle(legendX, legendY, 12, 12)
	dc.Fill()
	dc.SetRGB(0, 0, 0)
	dc.DrawString("EFE (Policy Evaluation)", legendX+18, legendY+10)

	// Title
	dc.LoadFontFace("Arial", 16)
	dc.DrawStringAnchored("Active Inference: Expected Free Energy Evolution", float64(spv.width/2), 20, 0.5, 0.5)

	// Subtitle
	dc.LoadFontFace("Arial", 10)
	dc.SetRGB(0.5, 0.5, 0.5)
	dc.DrawStringAnchored("Evaluates expected future free energy under different policies - lower is better", float64(spv.width/2), 40, 0.5, 0.5)

	// Save the image
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(spv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	return dc.SavePNG(fullPath)
}

// GenerateBeliefEvolution generates a simple line chart of belief evolution
func (spv *SimplePNGVisualizer) GenerateBeliefEvolution(beliefs []core.Belief, states []string, filename string) error {
	if len(beliefs) == 0 || len(states) == 0 {
		return fmt.Errorf("no beliefs or states to visualize")
	}

	dc := gg.NewContext(spv.width, spv.height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	margin := 50
	plotWidth := spv.width - 2*margin
	plotHeight := spv.height - 2*margin

	// Draw axes
	dc.SetRGB(0.2, 0.2, 0.2)
	dc.SetLineWidth(2)
	dc.DrawLine(float64(margin), float64(margin), float64(margin), float64(margin+plotHeight))
	dc.DrawLine(float64(margin), float64(margin+plotHeight), float64(margin+plotWidth), float64(margin+plotHeight))
	dc.Stroke()

	// Draw labels
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 12)
	dc.DrawStringAnchored("Step", float64(margin+plotWidth/2), float64(margin+plotHeight+20), 0.5, 0.5)
	dc.DrawStringAnchored("Belief", float64(margin-20), float64(margin+plotHeight/2), 0.5, 0.5)
	dc.RotateAbout(-math.Pi/2, float64(margin-20), float64(margin+plotHeight/2))
	dc.DrawStringAnchored("Belief", 0, 0, 0.5, 0.5)
	dc.RotateAbout(math.Pi/2, float64(margin-20), float64(margin+plotHeight/2))

	// Colors for different states with better contrast
	colors := []struct{ r, g, b float64 }{
		{0.8, 0.2, 0.2}, // Red
		{0.2, 0.8, 0.2}, // Green
		{0.2, 0.2, 0.8}, // Blue
		{0.8, 0.8, 0.2}, // Yellow
		{0.8, 0.2, 0.8}, // Magenta
		{0.2, 0.8, 0.8}, // Cyan
	}

	// Plot each state's belief evolution
	for stateIdx := range states {
		if stateIdx >= len(colors) {
			break
		}

		color := colors[stateIdx]
		dc.SetRGB(color.r, color.g, color.b)
		dc.SetLineWidth(2)

		for i := 0; i < len(beliefs)-1; i++ {
			if stateIdx < len(beliefs[i]) && stateIdx < len(beliefs[i+1]) {
				x1 := float64(margin + i*plotWidth/(len(beliefs)-1))
				y1 := float64(margin + plotHeight - int(beliefs[i][stateIdx]*float64(plotHeight)))
				x2 := float64(margin + (i+1)*plotWidth/(len(beliefs)-1))
				y2 := float64(margin + plotHeight - int(beliefs[i+1][stateIdx]*float64(plotHeight)))

				dc.DrawLine(x1, y1, x2, y2)
				dc.Stroke()

				// Draw points
				dc.DrawCircle(x1, y1, 3)
				dc.Fill()
			}
		}

		// Draw final point
		if len(beliefs) > 0 && stateIdx < len(beliefs[len(beliefs)-1]) {
			xFinal := float64(margin + (len(beliefs)-1)*plotWidth/(len(beliefs)-1))
			yFinal := float64(margin + plotHeight - int(beliefs[len(beliefs)-1][stateIdx]*float64(plotHeight)))
			dc.DrawCircle(xFinal, yFinal, 3)
			dc.Fill()
		}

		// Legend
		dc.SetRGB(color.r, color.g, color.b)
		dc.DrawRectangle(20, float64(50+stateIdx*20), 15, 15)
		dc.Fill()
		dc.SetRGB(0, 0, 0)
		dc.DrawString(states[stateIdx], float64(40), float64(62+stateIdx*20))
	}

	// Title
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 16)
	dc.DrawStringAnchored("Belief Evolution Over Time", float64(spv.width/2), 25, 0.5, 0.5)

	// Save the image
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(spv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	return dc.SavePNG(fullPath)
}

// GenerateMultiPanel generates a multi-panel visualization
func (spv *SimplePNGVisualizer) GenerateMultiPanel(
	title string,
	beliefs []core.Belief,
	vfeValues []float64,
	efeValues []float64,
	observations []string,
	filename string,
) error {
	// Create 2x2 grid
	panelWidth := spv.width / 2
	panelHeight := spv.height / 2

	dc := gg.NewContext(spv.width, spv.height)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// Draw title
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 18)
	dc.DrawStringAnchored(title, float64(spv.width/2), 25, 0.5, 0.5)

	// Draw panel borders
	dc.SetRGB(0.8, 0.8, 0.8)
	dc.SetLineWidth(1)
	dc.DrawRectangle(0, 0, float64(panelWidth), float64(panelHeight))
	dc.Stroke()
	dc.DrawRectangle(float64(panelWidth), 0, float64(panelWidth), float64(panelHeight))
	dc.Stroke()
	dc.DrawRectangle(0, float64(panelHeight), float64(panelWidth), float64(panelHeight))
	dc.Stroke()
	dc.DrawRectangle(float64(panelWidth), float64(panelHeight), float64(panelWidth), float64(panelHeight))
	dc.Stroke()

	// Panel 1: Current Belief (Top Left)
	if len(beliefs) > 0 {
		spv.drawBeliefPanel(dc, beliefs[len(beliefs)-1], []string{"Safe", "Dangerous"}, 0, 0, panelWidth, panelHeight, "Current Belief")
	}

	// Panel 2: VFE Evolution (Top Right)
	if len(vfeValues) > 0 {
		spv.drawSimplePlot(dc, vfeValues, panelWidth, 0, panelWidth, panelHeight, "VFE Evolution")
	}

	// Panel 3: Observations (Bottom Left)
	if len(observations) > 0 {
		spv.drawObservationsPanel(dc, observations, 0, panelHeight, panelWidth, panelHeight, "Observations")
	}

	// Panel 4: Belief Evolution (Bottom Right)
	if len(beliefs) > 0 {
		spv.drawBeliefEvolutionPanel(dc, beliefs, panelWidth, panelHeight, panelWidth, panelHeight, "Belief Evolution")
	}

	// Save the image
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	fullPath := filepath.Join(spv.outputDir, fmt.Sprintf("%s_%s.png", filename, timestamp))

	return dc.SavePNG(fullPath)
}

// Helper methods
func (spv *SimplePNGVisualizer) drawBeliefPanel(dc *gg.Context, belief core.Belief, states []string, x, y, w, h int, title string) {
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)
	dc.DrawStringAnchored(title, float64(x+w/2), float64(y+15), 0.5, 0.5)

	barWidth := w / len(belief)
	barHeight := h - 40

	for i, prob := range belief {
		height := int(prob * float64(barHeight))

		// Choose color
		if prob > 0.6 {
			dc.SetRGB(0, 1, 0)
		} else if prob > 0.3 {
			dc.SetRGB(1, 1, 0)
		} else {
			dc.SetRGB(1, 0, 0)
		}

		dc.DrawRectangle(float64(x+i*barWidth), float64(y+h-20-height), float64(barWidth-5), float64(height))
		dc.Fill()

		dc.SetRGB(0, 0, 0)
		dc.DrawStringAnchored(fmt.Sprintf("%.2f", prob), float64(x+i*barWidth+barWidth/2), float64(y+h-25-height), 0.5, 0.5)
	}
}

func (spv *SimplePNGVisualizer) drawSimplePlot(dc *gg.Context, data []float64, x, y, w, h int, title string) {
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)
	dc.DrawStringAnchored(title, float64(x+w/2), float64(y+15), 0.5, 0.5)

	margin := 10
	plotWidth := w - 2*margin
	plotHeight := h - 2*margin

	// Draw axes
	dc.SetRGB(0.2, 0.2, 0.2)
	dc.SetLineWidth(1)
	dc.DrawLine(float64(x+margin), float64(y+margin), float64(x+margin), float64(y+margin+plotHeight))
	dc.DrawLine(float64(x+margin), float64(y+margin+plotHeight), float64(x+margin+plotWidth), float64(y+margin+plotHeight))
	dc.Stroke()

	// Plot data
	dc.SetRGB(1, 0, 0)
	dc.SetLineWidth(2)

	for i := 0; i < len(data)-1; i++ {
		x1 := float64(x + margin + i*plotWidth/(len(data)-1))
		y1 := float64(y + margin + plotHeight - int(data[i]*float64(plotHeight)))
		x2 := float64(x + margin + (i+1)*plotWidth/(len(data)-1))
		y2 := float64(y + margin + plotHeight - int(data[i+1]*float64(plotHeight)))

		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}
}

func (spv *SimplePNGVisualizer) drawObservationsPanel(dc *gg.Context, observations []string, x, y, w, h int, title string) {
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)
	dc.DrawStringAnchored(title, float64(x+w/2), float64(y+15), 0.5, 0.5)

	margin := 20
	for i, obs := range observations {
		if i >= 5 { // Limit display
			break
		}
		yPos := float64(y + margin + 20 + i*15)
		dc.DrawString(fmt.Sprintf("%d: %s", i+1, obs), float64(x+margin), yPos)
	}
}

func (spv *SimplePNGVisualizer) drawBeliefEvolutionPanel(dc *gg.Context, beliefs []core.Belief, x, y, w, h int, title string) {
	dc.SetRGB(0, 0, 0)
	dc.LoadFontFace("Arial", 10)
	dc.DrawStringAnchored(title, float64(x+w/2), float64(y+15), 0.5, 0.5)

	margin := 20
	plotWidth := w - 2*margin
	plotHeight := h - 2*margin

	// Colors for different states
	colors := []struct{ r, g, b float64 }{{1, 0, 0}, {0, 0, 1}}

	// Plot belief evolution for first 2 states
	for stateIdx := 0; stateIdx < 2 && stateIdx < len(beliefs[0]); stateIdx++ {
		color := colors[stateIdx]
		dc.SetRGB(color.r, color.g, color.b)
		dc.SetLineWidth(2)

		for i := 0; i < len(beliefs)-1; i++ {
			if stateIdx < len(beliefs[i]) && stateIdx < len(beliefs[i+1]) {
				x1 := float64(x + margin + i*plotWidth/(len(beliefs)-1))
				y1 := float64(y + margin + plotHeight - int(beliefs[i][stateIdx]*float64(plotHeight)))
				x2 := float64(x + margin + (i+1)*plotWidth/(len(beliefs)-1))
				y2 := float64(y + margin + plotHeight - int(beliefs[i+1][stateIdx]*float64(plotHeight)))

				dc.DrawLine(x1, y1, x2, y2)
				dc.Stroke()
			}
		}
	}
}

// Global simple PNG visualizer instance
var SimplePNGVisualizerInstance = NewSimplePNGVisualizer("output/png_simple", 800, 600)
