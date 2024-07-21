package helpers

import (
	"errors"
	"fmt"
	"os"
	"path"
)

// reads all files in a directory
func ReadFiles(dirPath string , deepSearchFlag ,removeEmptyDirsFlag bool) (*[]string,error){

	if !deepSearchFlag {
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
			if deepSearchFlag {
				// if it is a directory it will use recursion to read all files
				newPath := path.Join(dirPath, entry.Name())
				nestedFiles , err := ReadFiles(newPath, deepSearchFlag, removeEmptyDirsFlag)
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
		// removes empty directories
		if removeEmptyDirsFlag {
			err := os.Remove(dirPath)
			if err != nil{
				return nil,err
			}
			GreenLog(fmt.Sprintf("✅ %s Empty Folder Deleted Successfully", dirPath))
		} else {
			RedLog(fmt.Sprintf("🔴 No Files Found In %s", dirPath))
			return nil, errors.New("no files found")
		}

		
	}

	GreenLog(fmt.Sprintf("✅ Got Files From %s" , dirPath))
	return &files, nil
}