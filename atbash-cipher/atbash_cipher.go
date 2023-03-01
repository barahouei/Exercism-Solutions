package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	s = strings.ToLower(s)

	cipherText := []rune{}
	counter := 0

	for _, v := range s {
		if counter == 5 && v != '.' && v != ',' {
			cipherText = append(cipherText, ' ')
			counter = 0
		}

		if v >= 'a' && v <= 'z' {
			diff := v - 'a'
			cipherText = append(cipherText, 'z'-diff)

			counter++
		} else if unicode.IsNumber(v) {
			cipherText = append(cipherText, v)

			counter++
		}
	}

	return string(cipherText)
}
