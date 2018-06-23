package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

var dirName = ".scribe/links.json"

func main() {
	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Errorf("The homedir could not be found with the following message %s", err)
		return
	}

	linkPath := path.Join(homeDir, dirName)

	resp, err := StartCli(os.Args, linkPath)
	for _, str := range resp {
		fmt.Println(str)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
