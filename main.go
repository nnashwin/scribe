package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/mitchellh/go-homedir"
	"os"
	"path"
)

var dirName = ".scribe/links.json"

func main() {
	respCol := color.New(color.FgMagenta).SprintFunc()
	errCol := color.New(color.FgRed).SprintFunc()

	homeDir, err := homedir.Dir()
	if err != nil {
		fmt.Errorf(errCol(fmt.Sprintf("The homedir could not be found with the following message %s", err)))
		return
	}

	linkPath := path.Join(homeDir, dirName)

	resp, err := StartCli(os.Args, linkPath)
	if err != nil {
		fmt.Println(errCol(fmt.Sprintf("%s", err)))
		return
	}

	// Print all output to the response
	for _, str := range resp {
		fmt.Printf("\n%s", respCol(fmt.Sprintf(str)))
	}
}
