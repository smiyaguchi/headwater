package main

import (
	"os"

	"github.com/smiyaguchi/headwater/cmd"
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := cmd.Execute(); err != nil {
		return 1
	}

	return 0
}
