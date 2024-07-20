package helpers

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/Rouch3362/roderfile/prompts"
)

func Compress() error {
	// getting user input path for compressing files
	result , err := prompts.GetUserPrompt("Enter folder path you want to compress its files", false)

	if err != nil {
		return err
	}

	// get all files in directory
	filesPath, err := ReadFiles(result)

	if err != nil {
		return err
	}

	filesChoosen , err := GetFileChoosen(filesPath)

	if err != nil {
		return err
	}

	err = CompressToZip(&filesChoosen, result)

	if err != nil {
		return err
	}

	return nil
}

// compresses file to zip
func CompressToZip(filesPath *[]string, dirPath string) error {
	// get user input for naming zip file
	result , err := prompts.GetUserPrompt("Enter zip file name (leave it blank so it will be default.zip)",false)

	if err != nil {
		return err
	}

	// if path not specified by user we store the compressed file(s) to the same path
	if result == "" {
		result = path.Join(dirPath, "default.zip")
	} else {
		// creating a name based on user entered name 
		result = path.Join(dirPath,result+".zip")
	}

	// creating a zip file
	zipFile , err := os.Create(result)

	if err != nil {
		return err
	}
	// creating an instance of zip writer with zip file so we can copy files to it
	zipWriter := zip.NewWriter(zipFile)

	// looping through all files
	for _ , filePath := range *filesPath {
		// get files info like file name, file instance
		file, fileName, _, err := GetFileContent(filePath)
		
		
		if err != nil {
			return err
		}

		defer file.Close()

		

		

		if err != nil {
			return err
		}

		


		zippedFile, err := zipWriter.Create(fileName)

		if err != nil {
			return err
		}

		if _, err := io.Copy(zippedFile, file); err != nil {
			return err
		}

		defer zipWriter.Close()
		
	}

	return nil
}


// gets information we need for comperssing a file like it content in bytes and it name and extension
func GetFileContent(filePath string) (*os.File, string, string , error)  {
	fileExt := path.Ext(filePath)

	file, err := os.Open(filePath)

	if err != nil {
		return nil, "", "", err
	}

	fileInfo , _ := os.Stat(filePath)


	return file, fileInfo.Name(), fileExt, nil
}


func GetFileChoosen(filesPath *[]string) ([]string , error) {
	// an instance of item from prompts Item
	options := []*prompts.Items{}

	// a map to store file name as key and file path as key
	filesPathMap := map[string]string{} 

	for _,path := range *filesPath {
		// getting last part of path => d:/test/example.txt -> example.txt
		fileName := filepath.Base(path)
		// getting file extension
		fileExt := filepath.Ext(path)

		// ignoring all .zip files
		if fileExt == ".zip" {
			continue
		}
		// adding file name as key and its path as value so we can access it later when user choose them by file name
		filesPathMap[fileName] = path
		// creating an item
		item := prompts.Items{
			Name: fileName,
			IsSelected: false,
		}
		// append it to items so we can pass it to prompt
		options = append(options, &item)
	}
	// get results from prompt
	result , err := prompts.MultipleChoicePrompt(0 , "Choose file to compress" , options)

	if err != nil {
		return nil,err
	}
	// getting paths based on user choice by file name
	for index, name := range result {
		result[index] = filesPathMap[name]
	}

	fmt.Println(result)

	return result, nil
}