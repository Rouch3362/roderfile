package prompts

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func CreateSelectPrompt(label string , options []string) (string , error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "✔ {{ . | green }}",
		Inactive: "  {{ . }}",
		Selected: "✔ {{ . | green }}",
	}
	
	prompt := promptui.Select{
		Label: label,
		Items: options,
		Templates: templates,
		HideHelp: true,
	} 

	_,result , err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result , nil
}


func GetUserPrompt(promptMessage string) (string, error) {

	
	templates := &promptui.PromptTemplates{
		Prompt: "{{ . }}: ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}


	validate := func (input string) error {
		if input == ""{
			return errors.New("this prompt can not be empty")
		}
		return nil
	}
	
	prompt := promptui.Prompt{
		Label: promptMessage,
		Validate: validate,
		Templates: templates,
		HideEntered: true,
	}


	

	result , err := prompt.Run()

	if err != nil {
		return "",err
	}

	return result, nil

}


func RunConfirmDeletePrompt(label string) (bool,error) {

	options := []string{"Yes", "No"}

	result , err := CreateSelectPrompt(label , options)
    
	if err != nil {
		return false, err
	}

	if result == "Yes"{
		return true,nil
	} else{
		return false , nil
	}
}


func SortingPrompt() (string , error) {

	label := "How you want me to sort your files"
	options := []string{"Don't Sort", "By Size Ascending", "By Size Descending","By Date Modified Ascending", "By Date Modified Descending"}
	
	result , err := CreateSelectPrompt(label , options)

	if err != nil {
		return "", err
	}

	return result, nil
}


