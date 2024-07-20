package helpers

import (
	"os"
	"sort"
	"github.com/Rouch3362/roderfile/prompts"
)


func SortFiles(files *[]string) error {
	result , err := prompts.SortingPrompt()

	if err != nil {
		return err
	}

	switch result {
		case "By Size Ascending":
			SortBySize(files , false)
		case "By Size Descending":
			SortBySize(files, true)
		case "By Date Modified Ascending":
			SortByDateModified(files , false)
		case "By Date Modified Descending":
			SortByDateModified(files, true)
		case "Don't Sort":
			break
	}

	return nil
}


func SortBySize(files *[]string , descending bool) {
	sort.Slice(*files, func(i, j int) bool {
		// get files info
		file1 , _ := os.Stat((*files)[i])
		file2 , _ := os.Stat((*files)[j])

		if !descending {
			return file1.Size() < file2.Size()
		}
		return file1.Size() > file2.Size()
	})
}



func SortByDateModified(files *[]string , descending bool) {
	sort.Slice(*files, func(i, j int) bool {
		// get files info
		file1 , _ := os.Stat((*files)[i])
		file2 , _ := os.Stat((*files)[j])

		if !descending {
			return file1.ModTime().Before(file2.ModTime())
		}
		return file1.ModTime().After(file2.ModTime())
	})
}