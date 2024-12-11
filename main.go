package main

import (
	"os"

	"github.com/beamlit/toolkit/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	err := cli.Execute(version, commit, date)
	if err != nil {
		os.Exit(1)
	}
}
