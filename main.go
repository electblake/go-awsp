package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/manifoldco/promptui"
)

const defaultProfileChoice = "default"

func readAwsProfiles() ([]byte, error) {
	dat, err := os.ReadFile(path.Join(os.Getenv("HOME"), ".aws/config"))
	if err != nil {
		panic(err)
	}
	if string(dat) == "" {
		return nil, errors.New("empty data")
	}

	return dat, nil
}

func writeProfileChoice(profile string) {
	if profile == defaultProfileChoice {
		profile = ""
	}
	// fmt.Printf("save %s to ~/.awsp\n", profile)
	err := os.WriteFile(path.Join(os.Getenv("HOME"), ".awsp"), []byte(profile), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func promptProfileChoice(data []byte) string {
	fmt.Println("AWS Profile Switcher")

	re, err := regexp.Compile(`\[profile (.*)]`)
	if err != nil {
		log.Fatal(err)
	}

	profiles := re.FindAll(data, -1)
	if len(profiles) == 0 {
		fmt.Printf("No profiles found.")
		fmt.Printf("Refer to this guide for help on setting up a new AWS profile:")
		fmt.Printf("https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html")

		log.Fatal("no profiles found")
	}

	// first and default choice
	choices := []string{defaultProfileChoice}

	for _, val := range profiles {
		choice := string(val)
		choice = strings.ReplaceAll(choice, "[profile ", "")
		choice = strings.ReplaceAll(choice, "]", "")
		// add to string map for promptui
		choices = append(choices, choice)
	}
	// fmt.Printf("profiles: %s", choices)

	prompt := promptui.Select{
		Label: "Choose a profile",
		Items: choices,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatal(err)
	}

	return result
}

func main() {

	dat, err := readAwsProfiles()

	fmt.Printf("loaded profiles (%d bytes)\n", len(dat))
	if err != nil {
		log.Fatal(err)
	}

	choice := promptProfileChoice(dat)
	writeProfileChoice(choice)

}
