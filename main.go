package main

import (
	"fmt"
	"os"

	"github.com/shiki-tak/shiki-web/cmd/root"
)

func main() {
	err := root.CMD().Execute()
	if err != nil {
		fmt.Printf("Failed executing CLI command: %s, exiting...\n", err)
		os.Exit(1)
	}
}
