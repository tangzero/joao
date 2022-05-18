package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifiers(t *testing.T) {
	for _, c := range []struct {
		algorithm string
		hash      string
		match     bool
	}{
		{"MD5", "9bdf52a483077", false},
		{"MD5", "zbdf52r5143729q1383ac24f1b3ep054", false},
		{"MD5", "84e5f5b9dad87e723b37de2e5c4cf072", true},
		{"SHA1", "AHKSAK", false},
		{"SHA1", "AHJSAJKHS767867868ASHJAGSJAGSJAGSJ678567", false},
		{"SHA1", "A94A8FE5CCB19BA61C4C0873D391E987982FBBD3", true},
		{"SHA256", "ba7816bf8f01cfea414140de5dae2223b0036", false},
		{"SHA256", "sadasg132dfscfea4dahsdjkash23134400361a396177a9cb410ff61f123dfkl", false},
		{"SHA256", "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad", true},
	} {
		identifier := identifiers[c.algorithm]
		assert.Equal(t, c.match, identifier([]byte(c.hash)), fmt.Sprintf("failed to match %q", c.hash))
	}
}
