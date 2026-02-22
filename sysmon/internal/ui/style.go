package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	TextGreen   = lipgloss.Color("#4ade80")
	BorderGreen = lipgloss.Color("#22c55e")
)

var ContainerStyle = lipgloss.NewStyle().
			MarginTop(2).
			Align(lipgloss.Center)

var TitleStyle  = lipgloss.NewStyle().
	Foreground(TextGreen).
	Bold(true).
	BorderStyle(lipgloss.NormalBorder()).
	Width(60).Align(lipgloss.Center).
	BorderForeground(BorderGreen).
	Padding(1, 2)
