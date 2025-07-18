package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	showCommand bool // Print curl command to execute
	verbose     bool
	silent      bool
	headers     []string
	userAgent   string
	follow      bool
)

var rootCmd = &cobra.Command{
	Use:   "curly",
	Short: "A curl wrapper for common use cases",
	Long: `curly is a curl wrapper that simplifies common HTTP operations.
It provides an intuitive interface for sending requests, handling authentication,
and managing responses.`,
	Version: "1.0.0",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().BoolVar(&showCommand, "show-command", false, "Show curl command being executed")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Verbose output (curl --verbose)")
	rootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "Silent mode (hide progress)")
	rootCmd.PersistentFlags().StringArrayVar(&headers, "header", []string{}, "Custom header (can be used multiple times)")
	rootCmd.PersistentFlags().StringVar(&userAgent, "user-agent", "", "User-Agent string")
	rootCmd.PersistentFlags().BoolVar(&follow, "follow", false, "Follow redirects")

	// Add subcommands
	rootCmd.AddCommand(postCmd)
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(headersCmd)

	// Disable default completion command
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

// checkURL validates if URL is provided
func checkURL(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("URL is required")
	}
	return nil
}
