// Package acronym makes abbreviation of words.
package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the abbreviation of a string.
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s, "-", " ")
	s = regexp.MustCompile(`[^a-zA-Z ]+`).ReplaceAllString(s, "")

	words := strings.Fields(s)

	var res string

	for _, word := range words {
		res += word[0:1]
	}

	res = strings.ToUpper(res)

	return res
}
