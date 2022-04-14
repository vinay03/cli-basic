package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use:   "Example",
		Short: "Example CLI command",
		Long: `A Simple Example of CLI command. 
Description follows to second line`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("hello from the root command")
		},
	}
)
