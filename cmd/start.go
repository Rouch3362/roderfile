/*
Copyright © 2024 Amirali Ashoori <rouchashoori@gmail.com>
*/
package cmd

import (
	"github.com/Rouch3362/roderfile/helpers"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "a cli tool that you can organize you files with",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := helpers.ReadFiles(args[0])
		return err
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
