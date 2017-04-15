package cryptosquare

import "strings"

const testVersion = 2

func Encode(s string) string {
	s = normalize(s)
	c := root(len(s))

	result := make([]string, c)
	for i, r := range s {
		result[i%c] += string(r)
	}

	return strings.Join(result, " ")
}

func normalize(s string) string {
	result := []rune{}
	for _, r := range strings.ToLower(s) {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result = append(result, r)
		}
	}
	return string(result)
}

func root(i int) int {
	result := 1
	for result*result < i {
		result++
	}
	return result
}
