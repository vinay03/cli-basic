package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var EchoCmd = &cobra.Command{
	Use:   "echo [strings to echo]",
	Short: "prints given strings to stdout",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Echo: ", strings.Join(args, " "))
	},
}

func init() {
	RootCmd.AddCommand(EchoCmd)
}
