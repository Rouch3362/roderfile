/*
Copyright © 2024 Amirali Ashoori <rouchashoori@gmail.com>
*/
package cmd

import (
	"github.com/Rouch3362/roderfile/helpers"
	"github.com/Rouch3362/roderfile/prompts"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "a cli tool that you can organize you files with",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := start()

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}


func start() error {
	dirPath, err := prompts.GetDirectoryPrompt()

	if err != nil {
		return err
	}

	r, err := helpers.ReadFiles(dirPath)
	if err != nil {
		return err
	}
	err = helpers.RemoveDuplicates(dirPath,r)

	if err != nil {
		return err
	}

	err = helpers.CategorizeFiles(r)

	if err != nil {
		return err
	}

	return nil
}