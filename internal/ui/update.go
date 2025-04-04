package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,
		tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return StartMsg{}
		}),
	)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd

	case StartMsg:
		if m.CurrentStep < len(m.Steps) {
			m.Steps[m.CurrentStep].Status = "running"
			return m, runStep(m)
		}

	case StepCompleteMsg:
		m.Steps[msg.StepIndex].Status = "completed"
		m.CurrentStep++

		if m.CurrentStep < len(m.Steps) {
			m.Steps[m.CurrentStep].Status = "running"
			return m, runStep(m)
		} else {
			m.Completed = true
			return m, tea.Quit
		}

	case ConversionCompleteMsg:
		m.Completed = true
		m.OutputFile = msg.OutputFile
		return m, tea.Quit

	case ConversionErrorMsg:
		if m.CurrentStep < len(m.Steps) {
			m.Steps[m.CurrentStep].Status = "failed"
			m.Steps[m.CurrentStep].Error = msg.Error
		}
		m.Error = msg.Error
		return m, tea.Quit
	}

	return m, nil
}

func runStep(m Model) tea.Cmd {
	return func() tea.Msg {
		currentStep := m.CurrentStep

		switch currentStep {
		case 0:
			format, err := m.Converter.Detect()
			if err != nil {
				return ConversionErrorMsg{Error: err}
			}

			m.Steps[currentStep].Message = "Detected format: " + format

			return StepCompleteMsg{StepIndex: currentStep}

		case 1:

			return StepCompleteMsg{StepIndex: currentStep}

		case 2:
			if err := m.Converter.Convert(); err != nil {
				return ConversionErrorMsg{Error: err}
			}
			return StepCompleteMsg{StepIndex: currentStep}

		case 3:
			return ConversionCompleteMsg{OutputFile: m.Converter.OutputFile}
		}

		return nil
	}
}
