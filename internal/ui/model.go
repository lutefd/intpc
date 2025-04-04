package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/lutefd/intpc/internal/app"
)

type Step struct {
	Name    string
	Message string
	Status  string
	Error   error
}

type Model struct {
	Converter     *app.Converter
	Steps         []Step
	CurrentStep   int
	Error         error
	Completed     bool
	OutputFile    string
	InputFile     string
	TargetFormat  string
	Spinner       spinner.Model
	Width, Height int
}

func NewModel(converter *app.Converter) Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	steps := []Step{
		{
			Name:    "Detect Format",
			Message: "Detecting input file format...",
			Status:  "pending",
		},
		{
			Name:    "Read Input",
			Message: "Reading input file...",
			Status:  "pending",
		},
		{
			Name:    "Convert",
			Message: "Converting between formats...",
			Status:  "pending",
		},
		{
			Name:    "Write Output",
			Message: "Writing output file...",
			Status:  "pending",
		},
	}

	return Model{
		Converter:    converter,
		Steps:        steps,
		CurrentStep:  0,
		InputFile:    converter.InputFile,
		OutputFile:   converter.OutputFile,
		TargetFormat: converter.Format,
		Spinner:      s,
	}
}

type StartMsg struct{}

type StepCompleteMsg struct {
	StepIndex int
}

type ConversionCompleteMsg struct {
	OutputFile string
}

type ConversionErrorMsg struct {
	Error error
}

type TickMsg time.Time
