package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

/*
$ go run main.go minargs
requires at least 1 arg(s), only received 0
exit status 1
*/

var MinArgsCmd = &cobra.Command{
	Use:   "minargs",
	Short: "Demonstates the minimum arguments rule",
	Args:  cobra.MinimumNArgs(1),
	Run: func(c *cobra.Command, args []string) {
		flags, _ := c.Flags().GetBool("simple1")
		fmt.Println("Value of sample1 flag: ", flags)
		flags, _ = c.Flags().GetBool("simple2")
		fmt.Println("Value of sample2 flag: ", flags)
	},
}

var sampleFlag1, sampleFlag2 bool

func init() {

	// Adding a flag
	MinArgsCmd.Flags().BoolVarP(&sampleFlag1, "sample1", "s", false, "Sample flag #1")
	MinArgsCmd.Flags().BoolVarP(&sampleFlag2, "sample2", "a", false, "Sample flag #2")

	RootCmd.AddCommand(MinArgsCmd)
}
