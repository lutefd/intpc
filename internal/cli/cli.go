package cli

import (
	"fmt"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lutefd/intpc/internal/app"
	"github.com/lutefd/intpc/internal/ui"
	"github.com/spf13/cobra"
)

var (
	outputFile string
	toPostman  bool
	toInsomnia bool
)

func Execute() error {
	var rootCmd = &cobra.Command{
		Use:   "intpc [file]",
		Short: "Convert between Insomnia and Postman formats",
		Long: `Insomnia to Postman Converter (intpc) is a command-line tool that allows you to 
convert between Insomnia v5 YAML exports and Postman collection JSON files.

You can convert in either direction: from Insomnia to Postman or from Postman to Insomnia.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			inputFile := args[0]

			if !app.FileExists(inputFile) {
				return fmt.Errorf("input file does not exist: %s", inputFile)
			}

			var targetFormat string
			if toPostman {
				targetFormat = "postman"
			} else if toInsomnia {
				targetFormat = "insomnia"
			} else {
				ext := filepath.Ext(inputFile)
				if ext == ".yaml" || ext == ".yml" {
					targetFormat = "postman"
				} else if ext == ".json" {
					targetFormat = "insomnia"
				} else {
					return fmt.Errorf("could not determine target format, please specify with --to-postman or --to-insomnia")
				}
			}

			if outputFile == "" {
				inputBase := filepath.Base(inputFile)
				inputExt := filepath.Ext(inputFile)
				inputName := inputBase[:len(inputBase)-len(inputExt)]

				if targetFormat == "postman" {
					outputFile = inputName + ".postman.json"
				} else {
					outputFile = inputName + ".insomnia.yaml"
				}
			}

			converter := app.NewConverter(inputFile, outputFile, targetFormat)

			model := ui.NewModel(converter)
			program := tea.NewProgram(model)

			if _, err := program.Run(); err != nil {
				return fmt.Errorf("error running UI: %w", err)
			}

			return nil
		},
	}

	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "", "output file (default: derived from input filename)")
	rootCmd.Flags().BoolVar(&toPostman, "to-postman", false, "convert to Postman format")
	rootCmd.Flags().BoolVar(&toInsomnia, "to-insomnia", false, "convert to Insomnia format")

	return rootCmd.Execute()
}
