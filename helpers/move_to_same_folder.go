package helpers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Rouch3362/roderfile/prompts"
)

type CommonFileInfo struct {
	Count int
	Path  []string
}
// a variable that store if we made any changes to files 
var MadeChanges bool


func CreateCommonFileMap(folderToLookFor string, deepSearch, removeEmptyDirs bool) (map[string]*CommonFileInfo, error ){
	// get file in a directory
	filesPath, err := ReadFiles(folderToLookFor, deepSearch, removeEmptyDirs)

	if err != nil {
		return nil,err
	}

	// creating an instance for saving files with the same name and their addressses
	commonFiles := map[string]*CommonFileInfo{}

	// looping over founded files
	for _, filePath := range *filesPath {

		GreenLog("🔍 Finding Relation Between Files...")

		// getting the file name from file path
		lastPartOfPath := filepath.Base(filePath)
		// getting file extension from file path
		fileExt := path.Ext(filePath)
		// extracting pure name of file from last part of path (filename.ext)
		filename := lastPartOfPath[:strings.LastIndex(lastPartOfPath, fileExt)]

		// if the file name not exists in the commonFiles map
		if _, ok := commonFiles[filename]; !ok {
			commonFiles[filename] = &CommonFileInfo{
				Count: 0,
			}
		}
		// saves path of file with same name file
		commonFiles[filename].Path = append(commonFiles[filename].Path, filePath)
		// increasing count of it
		commonFiles[filename].Count++
	}

	return commonFiles, nil
}




func MoveToCommonFolder(folderToLookFor string, deepSearch, removeEmptyDirs bool) error {
	
	commonFiles, err := CreateCommonFileMap(folderToLookFor, deepSearch, removeEmptyDirs)

	if err != nil {
		return err
	}

	for key, value := range commonFiles {

		if value.Count < 2 {
			continue
		}


		// store the destination folder we want to store new folder in it
		parentPath := folderToLookFor
		// getting parent path file + file name
		folderPath := path.Join(parentPath, key)
		// getting user ideal path
		result,err := prompts.GetUserPrompt(fmt.Sprintf("The Folder We'll Be Saved On This Path %s Changed It Or Just Hit Enter To Continue" , folderPath), false)

		if err != nil {
			return err
		}
		// if use not hit enter the user ideal path we'll we saved
		if result != "" {
			folderPath = path.Join(result,key)
		}


		// checking if files is not in the folder with same name of files with common name
		if strings.Contains(folderPath[:strings.LastIndex(folderPath,"/")], key) {
			continue
		}

		// check if folder not exists and if it does creates one
		if CheckFileOrFolderNotExist(folderPath) {
			err := os.Mkdir(folderPath, 0700)
			if err != nil {
				return nil
			}
		}

		GreenLog(fmt.Sprintf("📁 Making Folder For %s", folderPath))
		// looping over files with same names
		for index, file := range commonFiles[key].Path {
			// moving them in one directory with their name
			err := MoveFile(file, folderPath)
			
			// checks if we are in the last iteration and then executes below codes
			if index == len(commonFiles[key].Path)-1 && removeEmptyDirs {
				// getting files parent path that is moving
				fileParent := file[:strings.LastIndex(file,"/")]
				// getting the content in those folders
				folderContents, _ := os.ReadDir(fileParent)

				// if they not contain any file or folders just removes them as empty folders 

				if len(folderContents) < 1 {
					os.Remove(fileParent)
				}
			}

			if err != nil {
				return err
			}
		}
		MadeChanges = true
		GreenLog(fmt.Sprintf("🆗 Made Common Folder For Your Files In %s", folderPath))
	}
	if !MadeChanges {
		GreenLog("🎊 Hooooray!! Everthing Is Clean")
	}
	return nil
}
