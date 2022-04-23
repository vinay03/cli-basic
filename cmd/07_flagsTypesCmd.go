package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	testCmd = &cobra.Command{
		Use:   "test",
		Short: "Test input data",
		Run: func(c *cobra.Command, args []string) {
			boolValue, _ := c.Flags().GetBool("boolean")
			fmt.Println("boolean value: ", boolValue)

		},
	}
)

func init() {

	// boolean data
	var booleanVar bool
	testCmd.Flags().BoolVarP(&booleanVar, "boolean", "b", false, "Check a boolean value")

	RootCmd.AddCommand(testCmd)

}
