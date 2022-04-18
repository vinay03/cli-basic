package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var PersistFlagCmd = &cobra.Command{
	Use:   "persist",
	Short: "Demonstrates Persistent flag",
	Run: func(c *cobra.Command, args []string) {
		flagValue, _ := c.Flags().GetBool("persist")
		fmt.Println("Persist flag vlaue : ", flagValue)
	},
}

var persistFlag bool

func init() {
	PersistFlagCmd.Flags().BoolVarP(&persistFlag, "persist", "p", false, "Specify a persist flag")
	RootCmd.AddCommand(PersistFlagCmd)
}
