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
	expectedClue := "search"
	expectedLink := "www.google.com"
	// test addLink
	resp, err := StartCli([]string{"./scribe", "al", expectedClue, expectedLink}, testDir)
	if err != nil {
		t.Errorf("The addLink command encountered the following error: %s", err)
	}

	// test getLink
	resp, err = StartCli([]string{"./scribe", "gl", "search"}, testDir)
	if err != nil {
		t.Errorf("The getLink command encountered the following error: %s", err)
	}

	text, err := clipboard.ReadAll()
	if err != nil {
		t.Errorf("There was an error reading the string from the clipboard: %s", err)
	}

	if text != expectedLink {
		t.Errorf("The getLink command did not return the expected output\n Expected: %s\n Actual: %s", expectedLink, text)
	}

	// test listLinks
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

	// Adds link in order to be changed
	_, err = StartCli([]string{"./scribe", "al", expectedClue, expectedLink}, testDir)
	if err != nil {
		t.Errorf("The addLink command encountered the following error: %s", err)
	}

	newLink := "www.amazon.com"
	resp, err = StartCli([]string{"./scribe", "cl", "search", newLink}, testDir)
	if err != nil {
		t.Errorf("The changeLink command encountered the following error: %s", err)
	}

	// Gets the link again to overwrite the clipboard with the new link
	resp, _ = StartCli([]string{"./scribe", "gl", "search"}, testDir)

	text, err = clipboard.ReadAll()
	if err != nil {
		t.Errorf("There was an error reading the string from the clipboard: %s", err)
	}

	if text != newLink {
		t.Errorf("The changeLink command did not return the expected output\n Expected: %s\n Actual: %s", newLink, text)
	}

	err = os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Cleanup in the StartCli / AddLink test failed with the following error: %s", err)
	}
}
