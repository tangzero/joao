package main

import (
	"log"
	"os"

	"github.com/tangzero/joao/commands"
	"github.com/urfave/cli/v2"
)

var (
	version = "dev"
	date    = "unknown"
)

func main() {
	app := &cli.App{
		Name:  "joao",
		Usage: "simple password recovery tool",
		Commands: []*cli.Command{
			commands.Identify,
			commands.Decrypt,
			commands.Version(version),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
