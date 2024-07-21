/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Rouch3362/roderfile/helpers"
	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "to compress your large files",
	RunE: func(cmd *cobra.Command, args []string) error {
		deepFlag , _ := cmd.Flags().GetBool("deep")
		err := helpers.Compress(deepFlag)

		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
	compressCmd.Flags().BoolVarP(&deep,"deep", "d", true, "for deep folder search. if you don't set this to false, the search will look through all sub directories.")
}
