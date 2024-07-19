package prompts

import (
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