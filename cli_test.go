package main

import (
	"fmt"
	"testing"
)

func TestStartCli(t *testing.T) {
	ran, err := StartCli([]string{"./scribe get cookies"})
	if len(ran) < 1 && err != nil {
		fmt.Println("The command cookies did not exist.  Try to add the command and run again")
		t.Fail()
	}
}
