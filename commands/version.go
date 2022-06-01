package commands

import (
	"fmt"
	"runtime"

	"github.com/urfave/cli/v2"
)

func Version(version string) *cli.Command {
	return &cli.Command{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "show version",
		Action: func(c *cli.Context) error {
			fmt.Printf("joao version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
			return nil
		},
	}
}
