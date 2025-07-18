package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zinrai/curly/internal/curl"
	"github.com/zinrai/curly/internal/output"
)

var headersCmd = &cobra.Command{
	Use:   "headers <url>",
	Short: "Get response headers only",
	Long: `Send HEAD request to get response headers only.
This is useful for checking server configuration, content type, 
or redirect locations without downloading the body.`,
	Example: `  curly headers https://example.com
  curly headers https://example.com --follow`,
	Args: cobra.ExactArgs(1),
	RunE: runHeaders,
}

func runHeaders(cmd *cobra.Command, args []string) error {
	url := args[0]

	// Build curl command for headers only
	builder := curl.NewBuilder("HEAD", url)
	builder.AddFlag("--head") // Use HEAD method

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

	// Execute
	executor := curl.NewExecutor()
	if showCommand {
		output.PrintCommand(builder.Build())
	}

	return executor.Run(builder)
}
