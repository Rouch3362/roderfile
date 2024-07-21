/*
Copyright © 2024 Amirali Ashoori <rouchashoori@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/Rouch3362/roderfile/helpers"
	"github.com/Rouch3362/roderfile/prompts"
	"github.com/spf13/cobra"
)


var (
	deep bool
	removeEmptyDir bool
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "a cli tool that your can organize you files with",
	RunE: func(cmd *cobra.Command, args []string) error {
		
		deepFlag , _ := cmd.Flags().GetBool("deep")
		rmd , _ := cmd.Flags().GetBool("remove-empty-dirs")
		fmt.Println(rmd)
		err := start(deepFlag , rmd)

		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolVarP(&deep,"deep", "d", true, "for deep folder search. if you don't set this to false, the search will look through all sub directories.")
	startCmd.Flags().BoolVarP(&removeEmptyDir, "remove-empty-dirs", "r",true , "for removing empty folders. if you don't set this to false, this tool will remove all empty folder in given directory and it is also affected by deep flag")
}


func start(deepSearch, rmd bool) error {
	dirPath, err := prompts.GetUserPrompt("Type the path of directory you want to organize", true)
	

	if err != nil {
		return err
	}

	path, err := helpers.ReadFiles(dirPath, deepSearch , rmd)
	if err != nil {
		return err
	}
	err = helpers.RemoveDuplicates(dirPath,path)

	if err != nil {
		return err
	}

	err = helpers.SortFiles(path)
	
	if err != nil {
		return err
	}

	err = helpers.CategorizeFiles(path)

	if err != nil {
		return err
	}
	
	return nil
}