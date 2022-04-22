package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	ErrorCmd = &cobra.Command{
		Use:   "error",
		Short: "demonstrate error condition",
		RunE: func(c *cobra.Command, args []string) error {
			return errors.New("dummy error ocurred")
		},
	}
)

func init() {
	RootCmd.AddCommand(ErrorCmd)
}
