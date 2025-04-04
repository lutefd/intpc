package ui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	if m.Error != nil {
		return errorStyle.Render(fmt.Sprintf("Error: %v", m.Error))
	}

	if m.Completed {
		return successStyle.Render(fmt.Sprintf("✨ Conversion completed! ✨\n\nOutput saved to: %s",
			filepathStyle.Render(m.OutputFile)))
	}

	var s strings.Builder

	title := "Insomnia ↔ Postman Converter"
	s.WriteString(titleStyle.Render(title) + "\n\n")

	inputBase := filepath.Base(m.InputFile)
	outputBase := filepath.Base(m.OutputFile)
	s.WriteString(fmt.Sprintf("Converting: %s → %s\n\n",
		filepathStyle.Render(inputBase),
		filepathStyle.Render(outputBase)))

	for _, step := range m.Steps {
		var prefix string
		var nameStyle lipgloss.Style

		switch step.Status {
		case "completed":
			prefix = completedStepStyle.String()
			nameStyle = stepNameStyle.Copy().Faint(true)
		case "running":
			prefix = m.Spinner.View()
			nameStyle = stepNameStyle
		case "failed":
			prefix = failedStepStyle.String()
			nameStyle = errorStyle
		default:
			prefix = pendingStepStyle.String()
			nameStyle = stepNameStyle.Copy().Faint(true)
		}

		s.WriteString(fmt.Sprintf("%s%s", prefix, nameStyle.Render(step.Name)))

		if step.Status == "running" {
			s.WriteString(fmt.Sprintf(": %s", stepMessageStyle.Render(step.Message)))
		}

		if step.Error != nil {
			s.WriteString(fmt.Sprintf(": %s", errorStyle.Render(step.Error.Error())))
		}

		s.WriteString("\n")
	}

	s.WriteString("\nPress Ctrl+C to exit\n")

	return s.String()
}
