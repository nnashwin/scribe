package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	pf "github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"path"
)

var Links = struct {
	Entries map[string]Link `json:"entries,omitempty"`
}{}

var dirName = ".scribe/links.json"

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

				linkPath := path.Join(homeDir, dirName)

				// Create file if it doesn't exist
				if pf.DoesExist(linkPath) == false {
					err = pf.CreateFile(linkPath)
					if err != nil {
						return err
					}
				}

				links, err := ioutil.ReadFile(linkPath)
				if err != nil {
					return fmt.Errorf("%s", err)
				}

				if len(links) > 0 {
					err = json.Unmarshal(links, &Links)
					if err != nil {
						return fmt.Errorf("%s", err)
					}
				}

				if Links.Entries == nil {
					Links.Entries = make(map[string]Link)
				}

				// map the arg desc to the url
				if _, ok := Links.Entries[c.Args().First()]; ok == false {
					Links.Entries[c.Args().First()] = Link{c.Args().Get(1)}
				} else {
					return fmt.Errorf("The keyword '%s' is already recorded in your list of links", c.Args().First())
				}

				b, err := json.Marshal(Links)
				if err != nil {
					return fmt.Errorf("%s", err)
				}

				ioutil.WriteFile(linkPath, b, os.ModePerm)

				fmt.Printf("Enscribed a link '%s' to your records.\n", c.Args().Get(1))
				fmt.Println(Links.Entries)

				resp = append(resp, homeDir)

				return nil
			},
		},
	}

	err = app.Run(args)
	return
}
