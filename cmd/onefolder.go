/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Rouch3362/roderfile/helpers"
	"github.com/spf13/cobra"
)

// onefolderCmd represents the onefolder command
var onefolderCmd = &cobra.Command{
	Use:   "onefolder",
	Short: "creates a folder for two or more files with same name but different file types.",
	RunE: func(cmd *cobra.Command, args []string) error {
		deepSearch , _ := cmd.Flags().GetBool("deep")

		helpers.MoveToCommonFolder("d:/test/season 1/episode 1" , deepSearch)


		return nil
	},
}

func init() {
	rootCmd.AddCommand(onefolderCmd)
	onefolderCmd.Flags().BoolVarP(&deep , "deep", "d" , true, "for deep folder search. if you don't set this to false, the search will look through all sub directories.")
}
