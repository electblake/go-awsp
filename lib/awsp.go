package awsp

import (
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

const defaultProfileChoice = "default"

func ReadAwsProfiles() ([]byte, error) {
	dat, err := os.ReadFile(path.Join(os.Getenv("HOME"), ".aws/config"))
	if err != nil {
		panic(err)
	}
	if string(dat) == "" {
		return nil, errors.New("no profile data in ~/.aws/config")
	}

	return dat, nil
}

func WriteProfileChoice(profile string) error {
	if profile == defaultProfileChoice {
		profile = ""
	}
	// fmt.Printf("save %s to ~/.awsp\n", profile)
	err := os.WriteFile(path.Join(os.Getenv("HOME"), ".awsp"), []byte(profile), 0644)
	if err != nil {
		return err
	}
	return nil
}

func FindProfiles(data []byte) ([]string, error) {
	re, err := regexp.Compile(`\[profile (.*)]`)
	if err != nil {
		return nil, err
	}

	profiles := re.FindAll(data, -1)
	if len(profiles) == 0 {
		fmt.Printf("No profiles found.")
		fmt.Printf("Refer to this guide for help on setting up a new AWS profile:")
		fmt.Printf("https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-profiles.html")

		return nil, errors.New("no profiles found")
	}

	// first default choice
	choices := []string{defaultProfileChoice}

	for _, val := range profiles {
		choice := string(val)
		choice = strings.ReplaceAll(choice, "[profile ", "")
		choice = strings.ReplaceAll(choice, "]", "")
		// add to string map for promptui
		choices = append(choices, choice)
	}
	return choices, nil
}

func PromptProfileChoice(profiles []string) (string, error) {

	prompt := promptui.Select{
		Label: "Choose a profile",
		Items: profiles,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return "", err
	}

	return result, nil
}

func SaveProfileChoice(profile string) error {
	err := WriteProfileChoice(profile)
	os.Setenv("AWS_PROFILE", profile)

	fmt.Printf("aws profile: %s\n", profile)

	return err
}
