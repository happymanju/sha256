package main

import (
	"os"

	"github.com/happymanju/sha256/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
