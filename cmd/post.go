package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zinrai/curly/internal/curl"
	"github.com/zinrai/curly/internal/output"
)

var (
	jsonData string
	formData string
	fileData string
)

var postCmd = &cobra.Command{
	Use:   "post <url>",
	Short: "Send POST request",
	Long: `Send POST request with various data formats.
You can send JSON data, form data, or data from a file.`,
	Example: `  curly post https://api.example.com --json '{"key":"value"}'
  curly post https://api.example.com --data "name=john&age=30"
  curly post https://api.example.com --file @data.json`,
	Args: cobra.ExactArgs(1),
	RunE: runPost,
}

func init() {
	postCmd.Flags().StringVar(&jsonData, "json", "", "JSON data to send")
	postCmd.Flags().StringVar(&formData, "data", "", "Form data to send")
	postCmd.Flags().StringVar(&fileData, "file", "", "Read data from file (use @filename)")

	// These flags are mutually exclusive
	postCmd.MarkFlagsMutuallyExclusive("json", "data", "file")
}

func runPost(cmd *cobra.Command, args []string) error {
	url := args[0]

	// Build curl command
	builder := curl.NewBuilder("POST", url)

	// Add global flags
	if verbose {
		builder.AddFlag("--verbose")
	}
	if silent {
		builder.AddFlag("--silent")
	}
	if follow {
		builder.AddFlag("--location")
	}
	if userAgent != "" {
		builder.AddHeader("User-Agent", userAgent)
	}
	for _, h := range headers {
		builder.AddRawHeader(h)
	}

	// Add POST-specific options
	if jsonData != "" {
		builder.AddHeader("Content-Type", "application/json")
		builder.AddData(jsonData)
	} else if formData != "" {
		builder.AddData(formData)
	} else if fileData != "" {
		if fileData[0] != '@' {
			return fmt.Errorf("file parameter should start with @")
		}
		builder.AddDataFile(fileData)
	}

	// Execute
	executor := curl.NewExecutor()
	if showCommand {
		output.PrintCommand(builder.Build())
	}

	return executor.Run(builder)
}
