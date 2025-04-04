package ui

import "github.com/charmbracelet/lipgloss"

var (
	primaryColor   = lipgloss.Color("#4ECDC4")
	successColor   = lipgloss.Color("#59C9A5")
	errorColor     = lipgloss.Color("#FF6B6B")
	warningColor   = lipgloss.Color("#FFB347")
	secondaryColor = lipgloss.Color("#A0A0A0")

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor).
			Border(lipgloss.NormalBorder()).
			BorderForeground(primaryColor).
			Padding(0, 1).
			Align(lipgloss.Center)

	stepNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(primaryColor)

	stepMessageStyle = lipgloss.NewStyle().
				Foreground(secondaryColor).
				Italic(true)

	completedStepStyle = lipgloss.NewStyle().
				Foreground(successColor).
				SetString("✓ ")

	pendingStepStyle = lipgloss.NewStyle().
				Foreground(warningColor).
				SetString("• ")

	failedStepStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			SetString("✗ ")

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true)

	filepathStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Italic(true)

	spinnerStyle = lipgloss.NewStyle().
			Foreground(primaryColor)
)
