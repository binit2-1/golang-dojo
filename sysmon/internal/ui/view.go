package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	cpuBar := MakeProgressBar(m.CPUUsage, 20)
	cpuText := fmt.Sprintf("CPU Load:\t\t%.1f%%\n%s", m.CPUUsage, cpuBar)

	memBar := MakeProgressBar(m.MemoryUsage, 20)
	memText := fmt.Sprintf("Memory Usage:\t\t%.1f%%\n%s", m.MemoryUsage, memBar)

	cpuBox := BoxStyle.Render(cpuText)
	memBox := BoxStyle.Render(memText)

	dashboard := lipgloss.JoinHorizontal(lipgloss.Top, cpuBox, memBox)

	title := TitleStyle.Render("SYSTEM MONITOR DASHBOARD")

	footerText := "Instructions: Press 'q' or 'ctrl+c' to quit."
	footerRendered := FooterStyle.Width(m.Width).Render(footerText)

	// Render top area (title + dashboard)
	topRendered := lipgloss.JoinVertical(lipgloss.Center, title, dashboard)

	// Measure heights and compute spacer to push footer to bottom
	topHeight := lipgloss.Height(topRendered)
	footerHeight := lipgloss.Height(footerRendered)

	remaining := m.Height - topHeight - footerHeight
	if remaining < 0 {
		remaining = 0
	}
	spacer := strings.Repeat("\n", remaining)

	finalContents := topRendered + spacer + footerRendered
	finalUI := ContainerStyle.Width(m.Width).Render(finalContents)

	return finalUI
}

func MakeProgressBar(percent float64, width int) string {
	filled := int((percent / 100.0) * float64(width))
	empty := width - filled

	filledStr := strings.Repeat("█", filled)
	emptyStr := strings.Repeat("░", empty)
	return "[" + filledStr + emptyStr + "]"
}
