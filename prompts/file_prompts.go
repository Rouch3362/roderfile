package prompts

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"
)

func CreateSelectPrompt(label string , options []string) (string , error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "→ {{ . | green }}",
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


func GetUserPrompt(promptMessage string, requiredAnswer bool) (string, error) {

	
	templates := &promptui.PromptTemplates{
		Prompt: "{{ . }}: ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}


	validate := func (input string) error {
		if input == "" && requiredAnswer{
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



type Items struct{
	Name       string
	IsSelected bool
}

func MultipleChoicePrompt(selecId int, label string , options []*Items) ([]string, error) {
	
    // Define promptui template
    templates := &promptui.SelectTemplates{
        Label: `{{if .IsSelected}}
                    ✔
                {{end}} {{ .Name }} - label`,
        Active:   "→ {{if .IsSelected}}✔ {{end}}{{ .Name | green }}",
        Inactive: "{{if .IsSelected}}✔ {{end}}{{ .Name | green }}",
    }

	// check for "Done" if not exist so we add it so user can sumbit their choices
	if options[len(options)-1].Name != "Done" {
		options = append(options, &Items{"Done",false})
	}
    
	

    prompt := promptui.Select{
        Label:     label,
        Items:     options,
        Templates: templates,
        Size:      5,
        HideSelected: true,
		// Start the cursor at the currently selected index
		CursorPos: selecId,
    }



    
    selectedIdx, _, err := prompt.Run()
    
	if err != nil {
        return nil, fmt.Errorf("prompt failed: %w", err)
    }
	
	// getting choosen option 
    chosenItem := options[selectedIdx]
    
    if chosenItem.Name != "Done" {
        // If the user selected something other than "Done",
        // toggle selection on this item and run the function again.
        chosenItem.IsSelected = !chosenItem.IsSelected
        return MultipleChoicePrompt(selectedIdx,label, options)
    }
	

    // If the user selected the "Done" item, return
    // all selected items.
    var selectedItems []string
    for _, i := range options {
        if i.IsSelected {
            selectedItems = append(selectedItems, i.Name)
        }
    }

    return selectedItems, nil
}