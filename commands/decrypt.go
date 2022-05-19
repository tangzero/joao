//go:generate unzip rockyou.zip
package commands

import (
	"bufio"
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	_ "embed"
	"fmt"
	"strings"

	"github.com/tangzero/joao/input"
	"github.com/urfave/cli/v2"
)

var Decrypt = &cli.Command{
	Name:      "decrypt",
	Aliases:   []string{"d"},
	Usage:     "decrypt a hashed password",
	ArgsUsage: "[file]",
	Flags: []cli.Flag{
		&cli.Int64Flag{Name: "toggle", Aliases: []string{"t"}, Usage: "toggle letter cases"},
	},
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
			alg := identify(hash)

			tryToDecrypt(c, hash, alg)
		}
		return nil
	},
}

//go:embed rockyou.txt
var rockyou string

func tryToDecrypt(c *cli.Context, hash []byte, alg string) {
	algorithm, ok := algorithms[alg]
	if !ok {
		return
	}

	toggle := c.Value("toggle").(int64)

	scanner := bufio.NewScanner(bytes.NewBufferString(rockyou))
	scanner.Split(bufio.ScanLines)

	var found = false
	for scanner.Scan() && !found {
		password := scanner.Bytes()
		try := func(hash []byte, password []byte) {
			if strings.ToUpper(string(hash)) == fmt.Sprintf("%X", algorithm(password)) {
				fmt.Printf("%s [%s] = %s\n", string(hash), alg, password)
				found = true
			}
		}

		try(hash, password)

		for idx := int64(0); idx < toggle && idx < int64(len(password)) && !found; idx++ {
			if isLetter(hash[idx]) {
				toggledPassword := make([]byte, len(password))
				copy(toggledPassword, password)

				toggleFunc := strings.ToUpper
				if isUpper(toggledPassword[idx]) {
					toggleFunc = strings.ToLower
				}

				toggledPassword[idx] = toggleFunc(string([]byte{toggledPassword[idx]}))[0]
				try(hash, toggledPassword)
			}
		}
	}
}

func isLetter(c byte) bool {
	return isUpper(c) || isLower(c)
}

func isUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

var algorithms = map[string]func([]byte) []byte{
	"MD5": func(input []byte) []byte {
		sum := md5.Sum(input)
		return sum[:]
	},
	"SHA1": func(input []byte) []byte {
		sum := sha1.Sum(input)
		return sum[:]
	},
	"SHA256": func(input []byte) []byte {
		sum := sha256.Sum256(input)
		return sum[:]
	},
}
