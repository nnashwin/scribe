package main

import (
	"os"
	"path"
	"testing"
)

var testDir = path.Join("./fixtures", "Links.json")

func TestStartCli(t *testing.T) {

	// test add Link
	_, err := StartCli([]string{"./scribe", "al", "goog", "www.google.com"}, testDir)
	if err != nil {
		t.Errorf("The addLink command encountered the following error: %s", err)
	}

	// test get Link
	expected := "www.google.com"
	resp, err := StartCli([]string{"./scribe", "gl", "goog"}, testDir)
	if err != nil {
		t.Errorf("The getLink command encountered the following error: %s", err)
	}

	if resp[0] != expected {
		t.Error("The getLink command did not return the expected output")
	}

	err = os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Cleanup in the StartCli / AddLink test failed with the following error: %s", err)
	}
}
