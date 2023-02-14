// Package grains show the number of grains in chessboard squares.
package grains

import (
	"errors"
)

// Square returns the number of grains in a given chessboard square,
// and an error if it accrues.
func Square(number int) (uint64, error) {
	var grains uint64 = 1

	if number < 1 || number > 64 {
		return 0, errors.New("wrong number")
	}

	for i := 1; i < number; i++ {
		grains *= 2
	}

	return grains, nil
}

// Total returns the total of grains in chessboard.
func Total() uint64 {
	var grains uint64 = 1

	for i := 1; i < 64; i++ {
		grains *= 2
	}

	return grains<<grains - 1
}
