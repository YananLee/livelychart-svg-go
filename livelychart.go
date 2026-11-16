// Package livelychart renders a simple animated horizontal bar chart SVG.
package livelychart

import (
	"fmt"
	"html"
	"math"
	"strings"
)

type Bar struct {
	Label string
	Value float64
}
type Options struct {
	Width, Height      int
	Accent, Background string
	Duration           float64
}

// Render creates a standalone SVG from nonnegative values.
func Render(bars []Bar, options Options) (string, error) {
	if len(bars) == 0 || len(bars) > 30 {
		return "", fmt.Errorf("provide 1 to 30 bars")
	}
	if options.Width == 0 {
		options.Width = 800
	}
	if options.Height == 0 {
		options.Height = 480
	}
	if options.Accent == "" {
		options.Accent = "#635bff"
	}
	if options.Background == "" {
		options.Background = "#ffffff"
	}
	if options.Duration == 0 {
		options.Duration = 1.2
	}
	if options.Width < 300 || options.Height < 100 || options.Duration <= 0 || math.IsNaN(options.Duration) || math.IsInf(options.Duration, 0) || !hexColor(options.Accent) || !hexColor(options.Background) {
		return "", fmt.Errorf("invalid options")
	}
	max := 0.0
	for _, bar := range bars {
		if strings.TrimSpace(bar.Label) == "" || bar.Value < 0 || math.IsNaN(bar.Value) || math.IsInf(bar.Value, 0) {
			return "", fmt.Errorf("bars need labels and finite nonnegative values")
		}
		if bar.Value > max {
			max = bar.Value
		}
	}
	if max == 0 {
		max = 1
	}
	var out strings.Builder
	fmt.Fprintf(&out, `<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d" role="img" aria-label="Animated bar chart">`, options.Width, options.Height, options.Width, options.Height)
	fmt.Fprintf(&out, `<rect width="100%%" height="100%%" fill="%s"/>`, options.Background)
	row := float64(options.Height-40) / float64(len(bars))
	barHeight := math.Min(32, row*0.55)
	plotWidth := float64(options.Width - 270)
	for i, bar := range bars {
		y := 20 + row*float64(i) + (row-barHeight)/2
		w := plotWidth * bar.Value / max
		fmt.Fprintf(&out, `<text x="20" y="%.1f" dominant-baseline="middle" font-family="Arial, sans-serif" font-size="16" fill="#222222">%s</text>`, y+barHeight/2, html.EscapeString(bar.Label))
		fmt.Fprintf(&out, `<rect x="180" y="%.1f" width="0" height="%.1f" rx="5" fill="%s"><animate attributeName="width" from="0" to="%.1f" dur="%.2fs" begin="%.2fs" fill="freeze"/></rect>`, y, barHeight, options.Accent, w, options.Duration, float64(i)*0.1)
		fmt.Fprintf(&out, `<text x="%.1f" y="%.1f" dominant-baseline="middle" font-family="Arial, sans-serif" font-size="15" fill="#222222">%g</text>`, 190+w, y+barHeight/2, bar.Value)
	}
	out.WriteString(`</svg>`)
	return out.String(), nil
}

func hexColor(value string) bool {
	if len(value) != 7 || value[0] != '#' {
		return false
	}
	for _, c := range value[1:] {
		if !strings.ContainsRune("0123456789abcdefABCDEF", c) {
			return false
		}
	}
	return true
}
