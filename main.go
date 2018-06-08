package main

import (
	"fmt"
	"os"
)

func main() {
	resp, err := StartCli(os.Args)
	for _, str := range resp {
		fmt.Println(str)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
}
