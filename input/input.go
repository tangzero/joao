package input

import (
	"os"

	"github.com/urfave/cli/v2"
)

func GetInput(c *cli.Context) (*os.File, error) {
	if isInputFromPipe() {
		return os.Stdin, nil
	}
	if c.Args().Len() > 0 {
		return os.Open(c.Args().Get(0))
	}
	return nil, nil
}

func isInputFromPipe() bool {
	info, _ := os.Stdin.Stat()
	return info.Mode()&os.ModeCharDevice == 0
}
