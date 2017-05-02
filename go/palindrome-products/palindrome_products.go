package palindrome

import (
	"errors"
	"strconv"
)

const testVersion = 1

type Product struct {
	Product        int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	pmin, pmax := Product{}, Product{}
	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax")
	}
	for x := fmin; x <= fmax; x++ {
		for y := x; y <= fmax; y++ {
			if isPalindromic(x * y) {
				compare(x, y, &pmin, x*y < pmin.Product)
				compare(x, y, &pmax, x*y > pmax.Product)
			}
		}
	}
	if pmin.Factorizations == nil {
		return pmin, pmax, errors.New("No palindromes.")
	}
	return pmin, pmax, nil
}

func compare(x, y int, current *Product, better bool) {
	switch {
	case current.Factorizations == nil || better:
		*current = Product{x * y, [][2]int{{x, y}}}
	case x*y == current.Product:
		current.Factorizations = append(current.Factorizations, [2]int{x, y})
	}
}

func isPalindromic(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
