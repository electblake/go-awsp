package main

import (
	"fmt"
	"log"
	"os"

	awsp "github.com/electblake/homebrew-awsp/lib"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "awsp"
	app.Usage = "AWS Profile Switcher"
	app.Version = "0.1"

	app.Action = func(c *cli.Context) error {
		dat, err := awsp.ReadAwsProfiles()
		if err != nil {
			log.Fatal(err)
		}

		choice := awsp.PromptProfileChoice(dat)
		awsp.WriteProfileChoice(choice)

		os.Setenv("AWS_PROFILE", choice)

		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	app.Run(os.Args)

}
