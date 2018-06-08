package main

import (
	"github.com/fatih/color"
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

	app.Commands = []cli.Command{}

	err = app.Run(args)
	if err != nil {
		return resp, err
	}

	return
}
