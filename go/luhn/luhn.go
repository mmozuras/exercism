package luhn

import (
	"strconv"
	"strings"
)

const testVersion = 2

func Valid(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	i, _ := strconv.ParseUint(s, 10, 64)
	if i == 0 {
		return false
	}
	return checksum(i) == 0
}

func checksum(i uint64) int {
	result := 0
	for i != 0 {
		first := int(i % 10)
		second := int((i / 10) % 10)
		result += fold(first) + fold(second*2)
		i /= 100
	}
	return result % 10
}

func fold(i int) int {
	if i >= 10 {
		return i - 9
	}
	return i
}
