package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zinrai/curly/internal/curl"
	"github.com/zinrai/curly/internal/output"
)

var (
	outputFile string
)

var getCmd = &cobra.Command{
	Use:   "get <url>",
	Short: "Send GET request",
	Long:  `Send GET request and optionally save the response to a file.`,
	Example: `  curly get https://example.com
  curly get https://example.com --output output.html
  curly get https://example.com/file.pdf --output file.pdf`,
	Args: cobra.ExactArgs(1),
	RunE: runGet,
}

func init() {
	getCmd.Flags().StringVar(&outputFile, "output", "", "Save output to file")
}

func runGet(cmd *cobra.Command, args []string) error {
	url := args[0]

	// Build curl command
	builder := curl.NewBuilder("GET", url)

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

	// Add GET-specific options
	if outputFile != "" {
		builder.AddOutput(outputFile)
	}

	// Execute
	executor := curl.NewExecutor()
	if showCommand {
		output.PrintCommand(builder.Build())
	}

	return executor.Run(builder)
}
