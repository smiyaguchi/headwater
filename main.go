package main

import (
	"os"

	"github.com/smiyaguchi/headwater/cmd"
)

const (
	ExitCodeOK int = iota
	ExitCodeNG
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := cmd.Execute(); err != nil {
		return ExitCodeNG
	}

	return ExitCodeOK
}
