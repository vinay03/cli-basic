package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ParentCmd = &cobra.Command{
		Use:   "parent",
		Short: "Represents as parent command",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("From Parent command")
		},
	}

	ChildCmd = &cobra.Command{
		Use:   "child",
		Short: "Represents child command",
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("From Child command")
		},
	}
)

func init() {
	ParentCmd.AddCommand(ChildCmd)
	RootCmd.AddCommand(ParentCmd)
}
