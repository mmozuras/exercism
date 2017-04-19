package lsproduct

import "errors"

const testVersion = 4

func LargestSeriesProduct(digits string, span int) (int, error) {
	if span > len(digits) || span < 0 {
		return -1, errors.New("Inccorect input")
	}

	result := 0
	for i := 0; i <= len(digits)-span; i++ {
		product, e := product(digits[i : i+span])
		if e != nil {
			return product, e
		}
		if product > result {
			result = product
		}
	}
	return result, nil
}

func product(digits string) (int, error) {
	result := 1
	for _, digit := range digits {
		digit := digit - '0'
		if digit > 9 {
			return -1, errors.New("Input contains a non-digit")
		}
		result *= int(digit)
	}
	return result, nil
}
