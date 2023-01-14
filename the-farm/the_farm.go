package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (silly *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", silly.cows)
}

var ErrNegativeFodder = errors.New("negative fodder")
var ErrDivisionByZero = errors.New("division by zero")

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	var division float64

	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			if fodder > 0 {
				fodder = fodder * 2
				division = fodder / float64(cows)

				return division, nil
			} else if fodder < 0 {
				return 0, ErrNegativeFodder
			}
		}

		return 0, err
	}

	if fodder < 0 {
		return 0, ErrNegativeFodder
	}

	if cows == 0 {
		return 0, ErrDivisionByZero
	} else if cows < 0 {
		return 0, &SillyNephewError{cows: cows}
	}

	division = fodder / float64(cows)

	return division, nil
}
