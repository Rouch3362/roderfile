package helpers

import (
	"os"
	"path"

	"github.com/fatih/color"
)

// reads all files in a directory
func ReadFiles(dirPath string) ([]string,error){
	color.Green("📂 Fetching File From %s", dirPath)
	entries, err := os.ReadDir(dirPath)

	if err != nil {
		return nil , err
	}

	var files []string

	for _,entry := range entries{
		// check if the entry is a directory or a file
		if entry.IsDir(){
			// if it is a directory it will use recursion to read all files
			newPath := path.Join(dirPath, entry.Name())
			nestedFiles , _ := ReadFiles(newPath)
			// append returned values from recursion call to files slice
			files = append(files, nestedFiles...)
		} else {
			files = append(files, path.Join(dirPath,entry.Name()))
		}
	}
	color.Green("✅ Got Files From %s" , dirPath)
	return files, nil
}