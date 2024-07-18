package helpers

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/Rouch3362/roderfile/prompts"
	"github.com/fatih/color"
)



func RemoveDuplicates(dirPath string , filePaths []string) error {
	color.Green("🔍 Searching For Duplicated Files in %s", dirPath)
	// getting duplicated files
	duplicatedPaths,  err := CheckDuplicate(dirPath, filePaths)

	if err != nil {
		return err
	}
    // if found any duplicated files asks for permission and removes files
	if len(duplicatedPaths) > 0 {
		color.Red("❗ Found %d Duplicated File(s) In %s And Its Sub Dirs" ,len(duplicatedPaths),dirPath)
	
        accessGranted,err := prompts.RunConfirmDeletePrompt()
		
		if err != nil{
			return err
		}

		if accessGranted {
		
			for _, path := range duplicatedPaths {
				
				err := os.Remove(path)
	
				if err != nil {
					return err
				}
	
				color.Green("✅ Deleted Your Duplicated File Located: %s", path)
			}
		}
	}

	return nil
	
}

func CheckDuplicate(dirPath string,files []string) ([]string,error) {
	// create an instance of a map that save hash as key and file path as value
	filesSearched := map[string]string{}

	// a slice of duplicated files path
	var duplicatedFilesPath []string


	for _, filePath := range files {

		// getting hash of file
		hash, err := HashFile(filePath)

		if err != nil {
			return nil, err
		}

		// if hash as a key does not exist creates one and stores its file path
		if _,ok := filesSearched[hash]; !ok {
			filesSearched[hash] = filePath
			continue
		}
		// if hash already exists on map appends its value to duplicatedFilesPath
		duplicatedFilesPath = append(duplicatedFilesPath, filePath)
	}


	return duplicatedFilesPath, nil

}

func HashFile(filePath string) (string, error) {
	// open file for getting its content
	file , err := os.Open(filePath)

	if err != nil {
		return "", err
	}
	
	// created a hash instance
	hash := sha256.New()

	// copy the content of file into hash variable
	if _, err := io.Copy(hash, file); err != nil {
		return "",err
	}
	// after done with file we always close it
	defer file.Close()
	// returning the hash value
	return fmt.Sprintf("%x", hash), nil
}