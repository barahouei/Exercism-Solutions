// Package diffsquares deals with square of sum and sum of squares.
package diffsquares

// SquareOfSum returns the square of the sum of the first n natural numbers.
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	sum *= sum

	return sum
}

// SumOfSquares returns the sum of the squares of the first n natural number.
func SumOfSquares(n int) int {
	return (n * ((n + 1) * (n*2 + 1))) / 6
}

// Difference returns the difference between square of sum and sum of squares of the first n number.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
