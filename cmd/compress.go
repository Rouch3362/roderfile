/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Rouch3362/roderfile/helpers"
	"github.com/spf13/cobra"
)

// compressCmd represents the compress command
var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "to compress your large files",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		err := helpers.Compress()


		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(compressCmd)
}
