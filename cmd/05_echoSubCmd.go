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
)

func init() {
	RootCmd.AddCommand(ParentCmd)
}
