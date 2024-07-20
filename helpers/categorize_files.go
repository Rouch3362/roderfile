package helpers

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/Rouch3362/roderfile/prompts"
	"github.com/Rouch3362/roderfile/types"
)

// a variable that turns to true if we made changes to user files
var ORGANIZED bool


func CategorizeFiles(filePaths *[]string) error {

	fileTypesDirectories := map[string][]string{}

	GreenLog("🕵️  Analyzing Your Files...")
	for _, path := range *filePaths {
		// get extensions of a file
		fileExtension := filepath.Ext(path)
		// check if file in that path exists
		if fileExtension != "" {
			// get the kind(type) of file based on its extension
			fileKind := types.FileTypes[fileExtension]
			// save that file path to map as value and its kind as key
			fileTypesDirectories[fileKind] = append(fileTypesDirectories[fileKind], path)
		}
	}
	// created direcetories based on categories
	err := CreateDirectories(fileTypesDirectories)
    
	if err != nil {
		return err
	}
	// if everthing is clean and organized
	if !ORGANIZED {
		GreenLog("🎉 Hooooooray Everything is Already Organized")
	}

	return nil
}



func CreateDirectories(dirs map[string][]string) error {
	// loop over file types and their folder names
	for key , value := range dirs {
		
		// loop over file paths for creating folder for nested files
		for _ , pathFile := range value {
			// removing the filename from path to get pure path
			lastSlashIndex := strings.LastIndex(pathFile,"/")
			parentPath := pathFile[:lastSlashIndex] // example: d:/test/document/new.txt => d:/test/document
			pathToFolder := path.Join(parentPath,key)

			if AlreadyInCategorizedFolder(pathToFolder, parentPath) {
				continue
			}

			// if folder already exists ingnores rest of the code
			if !CheckFileOrFolderNotExist(pathToFolder) {
				err := MoveFile(pathFile, pathToFolder)
				if err != nil {
					fmt.Println(err)
					return err
				}
				continue
			}

			GreenLog(fmt.Sprintf("📁 Creating %s Folder For You...",key))

			if err := os.Mkdir(pathToFolder, 0700); err != nil {
				return err
			}
			
			// after creating folder move them to created folder
			err := MoveFile(pathFile, pathToFolder)
			

			if err != nil {	
				return err
			}
			GreenLog(fmt.Sprintf("✅ %sFolder Created Successfully", key))
		}
		
	}
	return nil
}



func CheckFileOrFolderNotExist(path string) bool {
	// basically as you can see this checks for existance of a folder or file
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return true
		} else {
			return false
		}
	}

	return false
}


func MoveFile(from , to string) error {

	ORGANIZED = true

	// opens file and reads its content
	file , err := os.Open(from)

	if err != nil {
		return err
	}

	GreenLog(fmt.Sprintf("📦 Moving %s File to %s ...",file.Name(), to))


	// extracts file name for example new.txt
	lastSlashIndex := strings.LastIndex(file.Name() , "/")
	newfilePath := file.Name()[lastSlashIndex:]

	// newToFilePath is just => to/filename
	newToFilePath := path.Join(to,newfilePath)

	// checks if file is exists
	if !CheckFileOrFolderNotExist(newToFilePath) {
		promptMessage := fmt.Sprintf("this file name is already exists at %s , select new name or either select ingore to ignore file and not moving it" , newToFilePath)
		// show user option for changing the file name or ingoring it
		userSelected , err :=prompts.CreateSelectPrompt(promptMessage , []string{"New Name","Ignore"})
			
		if err != nil {
			return err
		}
		// if user choose ignore
		if userSelected == "Ignore" {
			return nil
		}
		// get user input for new file name
		result , err := prompts.GetUserPrompt("Type New Name")

		if err != nil {
			return err
		}
		// rename the file
		newToFilePath = RenameFile(newToFilePath , result)
	}

	
	// creates a file in new path 
	newFile , err := os.Create(newToFilePath)

	if err != nil {
		return err
	}

	// copies content of old file to new file
	if _, err := io.Copy(newFile , file); err != nil {
		return err
	}
	
	/* for this usage we don't use defer for closing because defer happens after function done. and if this
	happens we get error for removing because another process accessing it file which is those open 
	statements. so we use this to close immediatly so we don't get that error */
	file.Close()
	newFile.Close()

	// removing the old file
	err = os.Remove(from)

	if err != nil {
		return err
	}

	GreenLog("📦 File Moved Successfully")

	return nil
}


/* checks if the file we want to create folder for it is not in a folder that has
been created before for categorization */
func AlreadyInCategorizedFolder(to, parentFilePath string) bool {
	/* get the folder names of destiantion path and current file path for example:
		to: d:/test/document => /document
		parentFilePath: d:/test => /test
	*/ 
	folderName := to[strings.LastIndex(to, "/"):]
	lastFolderOfParent := parentFilePath[strings.LastIndex(parentFilePath,"/"):]

	return folderName == lastFolderOfParent
}





func RenameFile(filePath , filename string) string {
	// getting file extension
	fileExt := path.Ext(filePath)

	// extracting old file name from path
	oldFileName := filePath[strings.LastIndex(filePath,"/"):]
	
	// creating new file name based on user input and file extension
	newFilename := "/"+filename+fileExt
	// creating new path with new file name
	newFilePath := strings.Replace(filePath,oldFileName, newFilename , -1)

	return newFilePath

}