package main

import (
	"fmt"
	"io"
	"os"

	awsp "github.com/electblake/homebrew-awsp/lib"
	"github.com/urfave/cli/v2"
)

// https://play.golang.org/p/Qg_uv_inCek
// contains checks if a string is present in a slice
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	app := cli.NewApp()
	app.Name = "awsp"
	app.Usage = "AWS Profile Switcher"
	app.Version = "0.1.2"

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Fprintf(w, "NAME:\n\tawsp - AWS Profile Switcher\n")
		fmt.Fprintf(w, "USAGE:\n\tawsp\t\t\t(prompt)\n\tawsp [profile]\t\t(direct set)\n")
		fmt.Fprintf(w, "COMMANDS:\n")
		fmt.Fprintf(w, "\thelp, h  Shows a list of commands or help for one command\n")
		fmt.Fprintf(w, "FLAGS:\n")
		fmt.Fprintf(w, "\t--version, -v  Shows version\n")

	}

	app.Action = func(c *cli.Context) error {

		input := c.Args().Get(0)

		dat, err := awsp.ReadAwsProfiles()
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}

		profiles, err := awsp.FindProfiles(dat)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}

		if input != "" && contains(profiles, input) {
			awsp.SaveProfileChoice(input)
			return nil
		} else if input != "" {
			return cli.Exit("specified profile does not exist", 1)
		}

		choice, err := awsp.PromptProfileChoice(profiles)
		if err != nil {
			return cli.Exit(err.Error(), 1)
		}

		err2 := awsp.SaveProfileChoice(choice)
		if err2 != nil {
			return cli.Exit(err2.Error(), 1)
		}

		return nil
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	app.Run(os.Args)

}
