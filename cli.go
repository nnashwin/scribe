package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	pf "github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var Links = struct {
	Entries map[string]Link `json:"entries,omitempty"`
}{}

func StartCli(args []string, linkPath string) (resp []string, err error) {
	respCol := color.New(color.FgMagenta).SprintFunc()

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
		resp = append(resp, respCol(fmt.Sprintf("Add a link with scribe!  Run scribe addLink <linkName> <link> to begin!")))
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "addLink",
			Aliases: []string{"al"},
			Usage:   "adds a link to your link repository",
			Action: func(c *cli.Context) error {
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
				Links.Entries[c.Args().First()] = Link{c.Args().Get(1)}

				b, err := json.Marshal(Links)
				if err != nil {
					return fmt.Errorf("%s", err)
				}

				ioutil.WriteFile(linkPath, b, os.ModePerm)

				fmt.Printf("Enscribed a link '%s' to your records.\n", c.Args().Get(1))

				resp = append(resp, linkPath)

				return nil
			},
		},
		{
			Name:    "getLink",
			Aliases: []string{"gl"},
			Usage:   "retrieves a previously defined link by a mnemonic and pastes it to your clipboard",
			Action: func(c *cli.Context) error {
				// Create file if it doesn't exist
				if pf.DoesExist(linkPath) == false {

					if err != nil {
						return err
					}
				}
				clipboard.WriteAll("There is no cow level")
				text, _ := clipboard.ReadAll()
				fmt.Println(text)
				return nil
			},
		},
	}

	err = app.Run(args)
	return
}
