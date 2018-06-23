package main

import (
	"fmt"
	"os"
	"path"
	"testing"
)

var testDir = path.Join("./fixtures", "Links.json")

func TestStartCli(t *testing.T) {
	// test add Link

	_, err := StartCli([]string{"./scribe", "al", "goog", "www.google.com"}, testDir)
	if err != nil {
		fmt.Printf("The addLink command encountered the follwing error: %s", err)
		t.Fail()
	}

	err = os.RemoveAll(testDir)
	if err != nil {
		t.Errorf("Cleanup in the StartCli / AddLink test failed with the following error: %s", err)
	}
}
