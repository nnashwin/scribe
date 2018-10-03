package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/atotto/clipboard"
	pf "github.com/ru-lai/pathfinder"
	"github.com/urfave/cli"
	"io/ioutil"
	"os"
	"text/template"
)

var Links = struct {
	Entries map[string]Link `json:"entries,omitempty"`
}{}

const listLinkTmpl = `- Clue: {{.Clue}},  Link: {{.Link}}`

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
		resp = append(resp, "Add a link with scribe!  Run scribe addLink <linkClue> <link> to begin!")
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "addLink",
			Aliases: []string{"al"},
			Usage:   "\n      - adds a link to your link repository by clue; \n        Example: scribe addLink search www.google.com\n          //=> Adds www.google.com to your directory of links under the clue 'search'\n",
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
			Name:    "deleteLink",
			Aliases: []string{"dl"},
			Usage:   "\n      - deletes a previously defined link by clue; \n        Example: scribe deleteLink search \n          //=> Deleted the link to 'google.com' from your link directory\n",
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
					delete(Links.Entries, c.Args().First())

					b, err := json.Marshal(Links)
					if err != nil {
						return err
					}

					ioutil.WriteFile(linkPath, b, os.ModePerm)
					resp = append(resp, fmt.Sprintf("Deleted the link to '%s' from your link directory", url))

					return nil
				} else {
					return fmt.Errorf("The keyword '%s' does not exist in your directory of links", c.Args().First())
				}
			},
		},
		{
			Name:    "changeLink",
			Aliases: []string{"cl"},
			Usage:   "\n      - changes a link previously defined to a clue to another; \n        Example: scribe changeLink search www.amazon.com\n          //=> Changes www.google.com to www.amazon.com\n",
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

				clue := c.Args().First()

				if _, ok := Links.Entries[clue]; ok == true {
					l := Links.Entries[clue]
					Links.Entries[clue] = Link{c.Args().Get(1)}

					b, err := json.Marshal(Links)
					if err != nil {
						return err
					}

					ioutil.WriteFile(linkPath, b, os.ModePerm)
					resp = append(resp, fmt.Sprintf("Found the link '%s' and changed it to '%s'", l.Url, Links.Entries[clue].Url))
				} else {
					return fmt.Errorf("The keyword '%s' does not exist in your list of links", clue)
				}

				return nil
			},
		},
		{
			Name:    "getLink",
			Aliases: []string{"gl"},
			Usage:   "\n      - retrieves a previously defined link by clue and pastes it to your clipboard; \n        Example: scribe getLink search\n          //=> Pastes www.google.com to your clipboard\n",
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
			Usage: "\n      - displays all of your stored clues and links; \n        Example: scribe listLinks\n					//=> Printing out your links:\n						- Link: tyler.com, Clue: cookies\n						- Link: google.com, Clue: goog\n",
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

				resp = append(resp, "Printing out your links:\n")
				t := template.Must(template.New("ListLinks").Parse(listLinkTmpl))
				for k, link := range Links.Entries {
					buf := &bytes.Buffer{}
					// adds map to better execute template on
					data := map[string]interface{}{
						"Clue": k,
						"Link": link.Url,
					}

					if err := t.Execute(buf, data); err != nil {
						return err
					}

					resp = append(resp, buf.String())
				}

				return nil
			},
		},
	}

	err = app.Run(args)
	return
}
