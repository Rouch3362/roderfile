package helpers

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// reads all files in a directory
func ReadFiles(dirPath string , deepSearch bool) (*[]string,error){

	if !deepSearch {
		YellowLog("⚠️  You Turned Deep Search Off")
	}

	GreenLog(fmt.Sprintf("📂 Fetching File(s) From %s", dirPath))
	entries, err := os.ReadDir(dirPath)
	
	if err != nil {
		return nil , err
	}

	var files []string

	for _,entry := range entries{
		// check if the entry is a directory or a file
		if entry.IsDir(){
			// if deep search is on i will crawl
			if deepSearch {
				// if it is a directory it will use recursion to read all files
				newPath := path.Join(dirPath, entry.Name())
				nestedFiles , err := ReadFiles(newPath, deepSearch)
				// append returned values from recursion call to files slice
				if err != nil {
					continue
				}
				files = append(files, *nestedFiles...)
			} else {
				continue 
			}
		} else {
			files = append(files, path.Join(dirPath,entry.Name()))
		}
	}

	// if nothing founded
	if len(files) < 1 {
		RedLog(fmt.Sprintf("🔴 No Files Found In %s", dirPath))
		return nil, errors.New("no files found")
	}

	GreenLog(fmt.Sprintf("✅ Got Files From %s" , dirPath))
	return &files, nil
}