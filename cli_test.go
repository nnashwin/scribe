package main

import (
	"github.com/atotto/clipboard"
	"os"
	"path"
	"strings"
	"testing"
)

var testDir = path.Join("./fixtures", "Links.json")

func TestStartCli(t *testing.T) {
	// test addLink
	resp, err := StartCli([]string{"./scribe", "al", "search", "www.google.com"}, testDir)
	if err != nil {
		t.Errorf("The addLink command encountered the following error: %s", err)
	}

	// test getLink
	expected := "www.google.com"
	resp, err = StartCli([]string{"./scribe", "gl", "search"}, testDir)
	if err != nil {
		t.Errorf("The getLink command encountered the following error: %s", err)
	}

	text, err := clipboard.ReadAll()
	if err != nil {
		t.Errorf("There was an error reading the string from the clipboard: %s", err)
	}

	if text != expected {
		t.Errorf("The getLink command did not return the expected output\n Expected: %s\n Actual: %s", expected, text)
	}

	// test listLinks
	expectedLink := "www.google.com"
	expectedClue := "search"

	resp, err = StartCli([]string{"./scribe", "ll"}, testDir)
	if err != nil {
		t.Errorf("The listLinks method returned an error: %s", err)
	}

	// check to see if the second string in the listLinks slice of strings has the Link and Clue
	if strings.Contains(resp[1], expectedLink) == false || strings.Contains(resp[1], expectedClue) == false {
		t.Errorf("The listLinks method failed to return the list of links and their clues")
	}

	StartCli([]string{"./scribe", "dl", "search"}, testDir)

	// Run getLink to see if www.google.com still exists; if it does this test should fail
	resp, err = StartCli([]string{"./scribe", "gl", "search"}, testDir)
	if len(resp) > 0 {
		t.Errorf("The deleteLinks method failed to delete the link")
	}

	err = os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Cleanup in the StartCli / AddLink test failed with the following error: %s", err)
	}
}
