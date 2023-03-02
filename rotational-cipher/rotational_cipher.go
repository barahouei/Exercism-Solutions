// Package rotationalcipher encodes string based on Caesar cipher.
package rotationalcipher

import "unicode"

// RotationalCipher will move every character of the given string x times on the alphabet.
func RotationalCipher(plain string, shiftKey int) string {
	if shiftKey == 0 || shiftKey == 26 {
		return plain
	}

	shift := rune(shiftKey)
	cipheredText := []rune{}

	for _, v := range plain {
		if unicode.IsSpace(v) || unicode.IsNumber(v) {
			cipheredText = append(cipheredText, v)
		} else if v >= 'a' && v <= 'z' {
			if v+shift > 'z' {
				v = (v + shift) - 'z' - 1 + 'a'

				cipheredText = append(cipheredText, v)
			} else {
				v += shift

				cipheredText = append(cipheredText, v)
			}
		} else if v >= 'A' && v <= 'Z' {
			if v+shift > 'Z' {
				v = (v + shift) - 'Z' - 1 + 'A'

				cipheredText = append(cipheredText, v)
			} else {
				v += shift

				cipheredText = append(cipheredText, v)
			}
		} else {
			cipheredText = append(cipheredText, v)
		}
	}

	return string(cipheredText)
}
