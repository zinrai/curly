package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zinrai/curly/internal/curl"
	"github.com/zinrai/curly/internal/output"
)

var headersCmd = &cobra.Command{
	Use:   "headers <url> [-- <curl-options>...]",
	Short: "Get response headers only",
	Long: `Send HEAD request to get response headers only.
This is useful for checking server configuration, content type, 
or redirect locations without downloading the body.
Additional curl options can be passed after '--'.`,
	Example: `  curly headers https://example.com
  curly headers https://example.com --follow
  curly headers https://example.com -- --dump-header headers.txt`,
	Args: cobra.ArbitraryArgs,
	RunE: runHeaders,
}

func runHeaders(cmd *cobra.Command, args []string) error {
	url, curlArgs, err := parseArgsWithCurlArgs(args)
	if err != nil {
		return err
	}

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

	// Add curl arguments
	builder.AddCurlArgs(curlArgs)

	// Execute
	executor := curl.NewExecutor()
	if showCommand {
		output.PrintCommand(builder.Build())
	}

	return executor.Run(builder)
}
