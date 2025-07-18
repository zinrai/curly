package main

import (
	"os"

	"github.com/zinrai/curly/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
