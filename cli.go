package main

import (
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	pf "github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
)

var Links = struct {
	Entries map[string]Link `json:"entries,omitempty"`
}{}

func StartCli(args []string, linkPath string) (resp []string, err error) {
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
		resp = append(resp, "Add a link with scribe!  Run scribe addLink <linkName> <link> to begin!")
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
					return err
				}

				if len(links) > 0 {
					err = json.Unmarshal(links, &Links)
					if err != nil {
						return err
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
					return err
				}

				ioutil.WriteFile(linkPath, b, os.ModePerm)

				resp = append(resp, fmt.Sprintf("Enscribed a link '%s' to your records.\n", c.Args().Get(1)))

				return nil
			},
		},
		{
			Name:    "getLink",
			Aliases: []string{"gl"},
			Usage:   "retrieves a previously defined link by a mnemonic and pastes it to your clipboard",
			Action: func(c *cli.Context) error {
				if pf.DoesExist(linkPath) == false {
					return fmt.Errorf("You have not created any links.  Run the addLink command and start")
				}

				links, err := ioutil.ReadFile(linkPath)
				if err != nil {
					return err
				}

				if len(links) > 0 {
					err = json.Unmarshal(links, &Links)
					if err != nil {
						return err
					}
				}

				if _, ok := Links.Entries[c.Args().First()]; ok == true {
					url := Links.Entries[c.Args().First()].Url

					clipboard.WriteAll(url)
					resp = append(resp, fmt.Sprintf("Found the link '%s' and copied it to your clipboard", url))
				} else {
					return fmt.Errorf("The keyword '%s' does not exist in your list of links", c.Args().First())
				}

				return nil
			},
		},
		{
			Name:    "listLinks",
			Aliases: []string{"ll"},
			Usage:   "displays all of your stored hints and links",
			Action: func(c *cli.Context) error {
				if pf.DoesExist(linkPath) == false {
					return fmt.Errorf("You have not created any links.  Run the addLink command and start")
				}

				links, err := ioutil.ReadFile(linkPath)
				if err != nil {
					return err
				}

				if len(links) > 0 {
					err = json.Unmarshal(links, &Links)
					if err != nil {
						return err
					}
				}

				resp = append(resp, "Printing out your links:")
				for k, link := range Links.Entries {
					resp = append(resp, fmt.Sprintf("- Link: %s, Clue: %s", link.Url, k))
				}

				return nil
			},
		},
	}

	err = app.Run(args)
	return
}
