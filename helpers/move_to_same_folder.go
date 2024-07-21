package helpers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type CommonFileInfo struct {
	Count int
	Path  []string
}
// a variable that store if we made any changes to files 
var MadeChanges bool

func MoveToCommonFolder(folderToLookFor string, deepSearch, removeEmptyDirs bool) error {
	// get file in a directory
	filesPath, err := ReadFiles(folderToLookFor, deepSearch, removeEmptyDirs)

	// creating an instance for saving files with the same name and their addressses
	commonFiles := map[string]*CommonFileInfo{}

	if err != nil {
		return err
	}
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

	for key, value := range commonFiles {
		// again getting last part of path of same file names
		lastPartOfPath := filepath.Base(value.Path[0])
		// getting the parent path of file
		parentPath := value.Path[0][:strings.LastIndex(value.Path[0], lastPartOfPath)]
		// getting parent path file + file name
		folderPath := path.Join(parentPath, key)

		// checking if files is not in the folder with same name of files with common name
		if filepath.Base(folderPath) == filepath.Base(parentPath) {
			continue
		}
		// check if folder not exists and if it does creates one
		if CheckFileOrFolderNotExist(folderPath) {
			err := os.Mkdir(folderPath, 07000)
			if err != nil {
				return nil
			}
		}

		GreenLog(fmt.Sprintf("📁 Making Folder For %s", folderPath))
		// looping over files with same names
		for _, file := range commonFiles[key].Path {
			// moving them in one directory with their name
			err := MoveFile(file, folderPath)

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
