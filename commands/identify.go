package commands

import (
	"bufio"
	"fmt"
	"regexp"

	"github.com/tangzero/joao/input"
	"github.com/urfave/cli/v2"
)

var Identify = &cli.Command{
	Name:      "identify",
	Aliases:   []string{"i"},
	Usage:     "identify the password hash algorithm",
	ArgsUsage: "[file]",
	Action: func(c *cli.Context) error {

		input, err := input.GetInput(c)
		if err != nil {
			return err
		}
		if input == nil {
			return nil
		}

		scanner := bufio.NewScanner(input)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			hash := scanner.Bytes()
			if alg := identify(hash); alg != "" {
				fmt.Printf("%s [%s]\n", hash, alg)
				continue
			}
			fmt.Printf("%s [unknow]\n", hash)
		}
		return nil
	},
}

func identify(hash []byte) string {
	for alg, identifier := range identifiers {
		if identifier(hash) {
			return alg
		}
	}
	return ""
}

var identifiers = map[string]func([]byte) bool{
	"MD5":    regexp.MustCompilePOSIX(`^[[:xdigit:]]{32}$`).Match,
	"SHA1":   regexp.MustCompilePOSIX(`^[[:xdigit:]]{40}$`).Match,
	"SHA256": regexp.MustCompilePOSIX(`^[[:xdigit:]]{64}$`).Match,
}
