package cmd

import (
	"github.com/spf13/cobra"
)

var extractDataCmd = &cobra.Command{
	Use:   "extract-data",
	Short: "Extracts data",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		s, err := getStrategyByName(args[0])
		if err != nil {
			return err
		}

		return s.DataExtraction(cmd.Context())
	},
}
