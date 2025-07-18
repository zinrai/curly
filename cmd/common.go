package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zinrai/curly/internal/curl"
	"github.com/zinrai/curly/internal/output"
)

// Basic auth support
func addBasicAuthCommand() *cobra.Command {
	var userPass string

	cmd := &cobra.Command{
		Use:   "basic-auth <url>",
		Short: "Send request with basic authentication",
		Long: `Send request with HTTP Basic Authentication.
The username and password should be provided in the format username:password.`,
		Example: `  curly basic-auth https://api.example.com --user admin:password
  curly basic-auth https://api.example.com --user admin:password --data "key=value"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if userPass == "" {
				return fmt.Errorf("--user flag is required for basic auth")
			}

			url := args[0]
			method := "GET"

			// Check if this is a POST request
			data, _ := cmd.Flags().GetString("data")
			if data != "" {
				method = "POST"
			}

			// Build curl command
			builder := curl.NewBuilder(method, url)
			builder.AddBasicAuth(userPass)

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

			// Add data if provided
			if data != "" {
				builder.AddData(data)
			}

			// Execute
			executor := curl.NewExecutor()
			if showCommand {
				output.PrintCommand(builder.Build())
			}

			return executor.Run(builder)
		},
	}

	cmd.Flags().StringVar(&userPass, "user", "", "Username and password (user:pass)")
	cmd.Flags().String("data", "", "Data to send (makes it a POST request)")
	cmd.MarkFlagRequired("user")

	return cmd
}

func init() {
	rootCmd.AddCommand(addBasicAuthCommand())
}

// Helper function to parse headers
func parseHeader(header string) (key, value string, err error) {
	parts := strings.SplitN(header, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid header format: %s", header)
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), nil
}
