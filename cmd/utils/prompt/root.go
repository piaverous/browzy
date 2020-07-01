package prompt

import (
	"errors"
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func PromptUrl() (string, error) {
	validateUrl := func(input string) error {
		if len(input) <= 3 {
			return errors.New("URL is too short")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Input bookmark URL",
		Validate: validateUrl,
	}
	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return "", err
	}
	return result, nil
}

func PromptName() (string, error) {
	validateName := func(input string) error {
		if len(input) <= 2 {
			return errors.New("Name is too short")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Choose a name for your bookmark",
		Validate: validateName,
	}
	result, err := prompt.Run()
	result = strings.Title(strings.ToLower(result))

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return "", err
	}
	return result, nil
}

func PromptPassword() (string, error) {
	validatePwd := func(input string) error {
		if len(input) < 8 {
			return errors.New("Password length needs to be >= 8")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Password",
		Validate: validatePwd,
		Mask:     '*',
	}
	pwd, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return "", err
	}
	return pwd, nil
}
