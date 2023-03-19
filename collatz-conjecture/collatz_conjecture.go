// Package collatzconjecture calculate collatz conjecture (3x+1).
package collatzconjecture

import "errors"

// CollatzConjecture returns the steps from n to 1.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("zero or negative numbers are not allowed")
	}

	counter := 0

	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}

		counter++
	}

	return counter, nil
}
