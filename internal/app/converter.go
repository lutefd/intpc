package app

import (
	"fmt"
)

type Converter struct {
	InputFile    string
	OutputFile   string
	Format       string
	SourceFormat string
}

func NewConverter(inputFile, outputFile, format string) *Converter {
	return &Converter{
		InputFile:  inputFile,
		OutputFile: outputFile,
		Format:     format,
	}
}

func (c *Converter) Detect() (string, error) {
	format, err := c.DetectFormat()
	if err != nil {
		return "", err
	}

	c.SourceFormat = format
	return format, nil
}

func (c *Converter) Convert() error {
	if c.SourceFormat == "" {
		format, err := c.Detect()
		if err != nil {
			return fmt.Errorf("error detecting input format: %w", err)
		}
		c.SourceFormat = format
	}

	if c.Format == "" {
		if c.SourceFormat == "insomnia" {
			c.Format = "postman"
		} else {
			c.Format = "insomnia"
		}
	}

	if c.SourceFormat == "insomnia" && c.Format == "postman" {
		return c.InsomniaToPostman()
	} else if c.SourceFormat == "postman" && c.Format == "insomnia" {
		return c.PostmanToInsomnia()
	}

	return fmt.Errorf("unsupported conversion: %s to %s", c.SourceFormat, c.Format)
}

func (c *Converter) InsomniaToPostman() error {
	insomniaData, err := ReadInsomniaFile(c.InputFile)
	if err != nil {
		return fmt.Errorf("error reading Insomnia file: %w", err)
	}

	postmanCollection, err := ConvertInsomniaToPostman(insomniaData)
	if err != nil {
		return fmt.Errorf("error converting to Postman format: %w", err)
	}

	err = WritePostmanFile(c.OutputFile, postmanCollection)
	if err != nil {
		return fmt.Errorf("error writing Postman file: %w", err)
	}

	return nil
}

func (c *Converter) PostmanToInsomnia() error {
	postmanData, err := ReadPostmanFile(c.InputFile)
	if err != nil {
		return fmt.Errorf("error reading Postman file: %w", err)
	}

	insomniaExport, err := ConvertPostmanToInsomnia(postmanData)
	if err != nil {
		return fmt.Errorf("error converting to Insomnia format: %w", err)
	}

	err = WriteInsomniaFile(c.OutputFile, insomniaExport)
	if err != nil {
		return fmt.Errorf("error writing Insomnia file: %w", err)
	}

	return nil
}

func (c *Converter) GetStepCount() int {
	return 4
}

func (c *Converter) GetStepName(step int) string {
	steps := []string{
		"Detect Format",
		"Read Input",
		"Convert",
		"Write Output",
	}

	if step >= 0 && step < len(steps) {
		return steps[step]
	}

	return "Unknown Step"
}

func (c *Converter) GetStepDescription(step int) string {
	descriptions := []string{
		"Detecting input file format...",
		"Reading input file...",
		"Converting between formats...",
		"Writing output file...",
	}

	if step >= 0 && step < len(descriptions) {
		return descriptions[step]
	}

	return "Unknown step"
}
