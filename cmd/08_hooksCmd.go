package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	hooksCmd = &cobra.Command{
		Use:   "hooks",
		Short: "demonstrate the sequence of hooks calling",
		PersistentPreRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooks command 'PersistentPreRun' hook")
		},
		PreRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooks command 'PreRun' hook")
		},
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("From hooks command 'Run' hook")
		},
		PostRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooks command 'PostRun' hook")
		},
		PersistentPostRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooks command 'PersistentPostRun' hook")
		},
	}

	hooksChildCmd = &cobra.Command{
		Use:   "hookschild",
		Short: "demonstrate the sequence of hooks calling",
		PersistentPreRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooksChild command 'PersistentPreRun' hook")
		},
		PreRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooksChild command 'PreRun' hook")
		},
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("From hooksChild command 'Run' hook")
		},
		PostRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooksChild command 'PostRun' hook")
		},
		PersistentPostRun: func(c *cobra.Command, args []string) {
			fmt.Println("From hooksChild command 'PersistentPostRun' hook")
		},
	}
)

func init() {
	hooksCmd.AddCommand(hooksChildCmd)
	RootCmd.AddCommand(hooksCmd)
}
