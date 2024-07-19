package prompts

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)


func RunConfirmDeletePrompt() (bool,error) {

    templates := &promptui.SelectTemplates{
		Label:    "{{ . }}:",
		Active:   fmt.Sprintf("%s {{ . | green }}", promptui.IconSelect),
		Inactive: "  {{ . }}",
		Selected: "{{ . | red | green }}",
	}


	prompt := promptui.Select{
		Label: "Do you want to delete duplicated files",
		Items: []string{"Yes","No"},
		HideHelp: true,
		Templates: templates,
	}


	_,result , err := prompt.Run()
    


	if err != nil {
		return false, err
	}

	if result == "Yes"{
		return true,nil
	} else{
		return false , nil
	}
}


func GetDirectoryPrompt() (string , error) {

	validate := func (input string) error {
		if input == ""{
			return errors.New("this prompt can not be empty")
		} else if len(input) < 3 {
			return errors.New("invalid path")
		}
		return nil
	}


	prompt := promptui.Prompt{
		Label: "Type the path of directory you want to organize",
		Validate: validate,
	}

	result , err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result,nil
}