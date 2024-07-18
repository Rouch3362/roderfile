package helpers

import (
	"os"
	"path"
)

// reads all files in a directory
func ReadFiles(dirPath string) ([]*os.DirEntry,error){
	entries, err := os.ReadDir(dirPath)

	if err != nil {
		return nil , err
	}

	var files []*os.DirEntry

	for _,entry := range entries{
		// check if the entry is a directory or a file
		if entry.IsDir(){
			// if it is a directory it will use recursion to read all files
			newPath := path.Join(dirPath, entry.Name())
			nestedFiles , _ := ReadFiles(newPath)
			// append returned values from recursion call to files slice
			files = append(files, nestedFiles...)
		} else {
			files = append(files, &entry)
		}
	}


	return files, nil
}