package curl

import (
	"fmt"
	"os"
	"os/exec"
)

// Executor runs curl commands
type Executor struct {
	curlPath string
}

// NewExecutor creates a new curl executor
func NewExecutor() *Executor {
	return &Executor{
		curlPath: "curl", // Assume curl is in PATH
	}
}

// Run executes the curl command
func (e *Executor) Run(builder *Builder) error {
	args := builder.Build()

	// Check if curl is available
	if _, err := exec.LookPath(e.curlPath); err != nil {
		return fmt.Errorf("curl not found in PATH: %w", err)
	}

	// Create command
	cmd := exec.Command(e.curlPath, args...)

	// Connect stdin, stdout, stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		// Don't wrap the error to preserve curl's exit code
		return err
	}

	return nil
}

// DryRun returns the command that would be executed
func (e *Executor) DryRun(builder *Builder) string {
	return builder.BuildCommand()
}
