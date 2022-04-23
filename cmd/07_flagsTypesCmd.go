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

			intValue, _ := c.Flags().GetBool("int")
			fmt.Println("int value: ", intValue)

			float32Value, _ := c.Flags().GetBool("float")
			fmt.Println("float32 value: ", float32Value)

		},
	}
)

func init() {

	// boolean data
	var booleanVar bool
	testCmd.Flags().BoolVarP(&booleanVar, "boolean", "b", false, "Check a boolean value")

	// int data
	var intVar int
	testCmd.Flags().IntVarP(&intVar, "int", "i", 0, "Check a int value")

	// float data
	var float32Var float32
	testCmd.Flags().Float32VarP(&float32Var, "float", "f", 0.0, "check a float32 value")

	RootCmd.AddCommand(testCmd)

}
