package perfect

import "errors"

const testVersion = 1

type Classification int

const (
	ClassificationPerfect Classification = iota
	ClassificationAbundant
	ClassificationDeficient
)

var ErrOnlyPositive = errors.New("Can't classify zero or negative.")

func Classify(number uint64) (Classification, error) {
	if number < 1 {
		return 0, ErrOnlyPositive
	}

	sum := aliquotSum(number)
	if sum == number {
		return ClassificationPerfect, nil
	} else if sum > number {
		return ClassificationAbundant, nil
	}
	return ClassificationDeficient, nil
}

func aliquotSum(number uint64) uint64 {
	var result uint64
	for i := uint64(1); i <= number/2; i++ {
		if number%i == 0 {
			result += i
		}
	}
	return result
}
