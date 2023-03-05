// Package reverse reverses a string.
package reverse

// Reverse returns the reverse of a string.
func Reverse(input string) string {
	inputRune := []rune(input)
	result := []rune{}

	for i := len(inputRune) - 1; i >= 0; i-- {
		result = append(result, inputRune[i])
	}

	return string(result)
}
