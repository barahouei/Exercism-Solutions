// Package prime checks for prime numbers.
package prime

import "errors"

// Nth returns the nth prime number.
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("not valid number")
	}

	nthPrime := 0
	counter := 0

	for i := 2; counter <= n; i++ {
		if isPrime(i) {
			counter++

			if counter == n {
				nthPrime = i
			}
		}
	}

	return nthPrime, nil
}

// isPrime checks if a number is prime.
func isPrime(n int) bool {
	if n < 1 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
