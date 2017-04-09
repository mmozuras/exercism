package grains

import "errors"

const testVersion = 1

var invalidSquare = errors.New("Invalid square, must be in 1-64 range.")

func Square(square int) (uint64, error) {
	if square <= 0 || square > 64 {
		return 0, invalidSquare
	}
	return 1 << uint64(square-1), nil
}

func Total() (result uint64) {
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		result += square
	}
	return
}
