package output

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

var (
	green  = color.New(color.FgGreen).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
)

// PrintCommand prints the curl command that will be executed
func PrintCommand(args []string) {
	// Build the command string
	quotedArgs := make([]string, len(args))
	for i, arg := range args {
		// Quote arguments that contain spaces or special characters
		if strings.ContainsAny(arg, " \t\n\"'{}[]") {
			quotedArgs[i] = fmt.Sprintf("'%s'", strings.ReplaceAll(arg, "'", "'\"'\"'"))
		} else {
			quotedArgs[i] = arg
		}
	}

	command := "curl " + strings.Join(quotedArgs, " ")

	// Print with color
	fmt.Fprintf(os.Stderr, "%s %s\n", green("Executing:"), cyan(command))
	fmt.Fprintln(os.Stderr, strings.Repeat("-", 80))
}

// PrintError prints an error message
func PrintError(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s %s\n", red("Error:"), msg)
}

// PrintWarning prints a warning message
func PrintWarning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s %s\n", yellow("Warning:"), msg)
}

// PrintInfo prints an info message
func PrintInfo(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(os.Stderr, "%s %s\n", green("Info:"), msg)
}
