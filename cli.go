package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"github.com/urfave/cli"
)

func StartCli(args []string) (resp []string, err error) {
	app := cli.NewApp()
	app.Name = "scribe"
	app.Version = "0.0.1"
	app.Usage = "Quick and easy storage / retrieval of links from keywords"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Tyler Boright",
			Email: "ru.lai.development@gmail.com",
		},
	}

	app.Action = func(c *cli.Context) error {
		color.Magenta("Add a link with scribe add <linkName> <link>")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "addLink",
			Aliases: []string{"al"},
			Usage:   "adds a link to your link repository",
			Action: func(c *cli.Context) error {
				homeDir, err := homedir.Dir()
				if err != nil {
					return fmt.Errorf("The homedir could not be found with the following message %s", err)
				}

				resp = append(resp, homeDir)

				return nil
			},
		},
	}

	err = app.Run(args)
	return
}
