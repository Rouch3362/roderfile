package helpers

import (
	"fmt"
	"os"
	"path"
)

// its a variable that turns true if searched for empty dirs but no empty dirs not found and turns false if found empty folders
var NotFoundEmptyFolders bool

func RemoveEmptyDirectory(dirPath string) error {
	// check for empty folder in directory user and its sub directories enter
	result , err := CheckForEmptyDirs(dirPath)

	if err != nil {
		return err
	}
	if len(*result) > 0 {
		NotFoundEmptyFolders = false

		for _,r := range *result {
			// removing empty folder
			err := os.Remove(r)
		
			if err != nil {
				return err
			}

			GreenLog(fmt.Sprintf("✅ Empty Folder %s Deleted Successfully", r))
			// check other empty files using recursion and removing them
			err = RemoveEmptyDirectory(dirPath)
		
			if err != nil {
				return err
			}
		}
	}


	

	return nil
}

func CheckForEmptyDirs(dirPath string) (*[]string, error) {
	// get all enteries in a directory user enetered
	enteries , err := os.ReadDir(dirPath)


	if err != nil {
		return nil, err
	}
	// path of empty folders will be saved in this
	var folderPathsToRemove []string

	NotFoundEmptyFolders = true

	for _,entery := range enteries {
		if entery.IsDir() {
			// create the folder path that points to => path_user_entered/folder
			folderPath := path.Join(dirPath, entery.Name())

			
			GreenLog(fmt.Sprintf("🔍 Searching For Empty Folders In %s ...", folderPath))
			
			// open folder for seeing if got anything in it
			f , err := os.Open(folderPath)


			if err != nil {
				return nil, err
			}

			defer f.Close()

			// getting name of files or folders in the searched folder 
			enteryNames , err := f.Readdirnames(-1)

			if err != nil {
				return nil, err
			}
			// if folder is empty
			if len(enteryNames) == 0 {
				// appends its path to variable that stores them
				folderPathsToRemove = append(folderPathsToRemove, folderPath)
			} else {
				// if it is not empty goes deeper and append result to variable for folder paths 
				result , _ := CheckForEmptyDirs(folderPath)
				folderPathsToRemove = append(folderPathsToRemove, *result...)
			}
		}
	}

	return &folderPathsToRemove, nil
}