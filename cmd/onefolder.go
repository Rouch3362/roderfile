/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Rouch3362/roderfile/helpers"
	"github.com/Rouch3362/roderfile/prompts"
	"github.com/spf13/cobra"
)

// onefolderCmd represents the onefolder command
var onefolderCmd = &cobra.Command{
	Use:   "onefolder",
	Short: "creates a folder for two or more files with same name but different file types.",
	RunE: func(cmd *cobra.Command, args []string) error {
		deepSearch , _ := cmd.Flags().GetBool("deep")
		rmd, _ := cmd.Flags().GetBool("remove-empty-dirs")

		result ,err := prompts.GetUserPrompt("Enter the path" , true)

		if err != nil {
			return err
		}

		err = helpers.MoveToCommonFolder(result , deepSearch, rmd)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(onefolderCmd)
	onefolderCmd.Flags().BoolVarP(&deep , "deep", "d" , true, "for deep folder search. if you don't set this to false, the search will look through all sub directories.")
	onefolderCmd.Flags().BoolVarP(&removeEmptyDir, "remove-empty-dirs", "r",true , "for removing empty folders. if you don't set this to false, this tool will remove all empty folder in given directory and it is also affected by deep flag")
}
