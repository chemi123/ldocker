package main

import (
	"fmt"
	"os"

	"github.com/chemi123/ldocker/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
