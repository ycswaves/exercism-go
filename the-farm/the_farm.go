package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fAmount, err := weightFodder.FodderAmount()

	switch {
	case err == ErrScaleMalfunction:
		if fAmount > 0 {
			fAmount *= 2
		} else {
			return 0, errors.New("negative fodder")
		}
	case err != nil: // generic err
		return 0, err
	case fAmount < 0:
		return 0, errors.New("negative fodder")

	case cows == 0:
		return 0, errors.New("division by zero")

	case cows < 0:
		return 0, &SillyNephewError{cows}
	}

	return fAmount / float64(cows), nil

}
