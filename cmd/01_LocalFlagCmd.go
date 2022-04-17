package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var LocalFlagCmd = &cobra.Command{
	Use:   "local",
	Short: "Demonstrates local flag",
	Run: func(c *cobra.Command, args []string) {
		flagValue, _ := c.Flags().GetBool("local")
		fmt.Println("local flag value : ", flagValue)
	},
}

var LocalFlag bool

func init() {
	LocalFlagCmd.Flags().BoolVarP(&LocalFlag, "local", "l", false, "Specify local flag")
	RootCmd.AddCommand(LocalFlagCmd)
}
