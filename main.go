package main

import (
	"log"
	"os"

	"github.com/tangzero/joao/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "joao",
		Usage: "simple password recovery tool",
		Commands: []*cli.Command{
			commands.Identify,
			commands.Decrypt,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
