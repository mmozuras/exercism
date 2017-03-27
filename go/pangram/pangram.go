package pangram

import "strings"

const testVersion = 1

func IsPangram(input string) bool {
	result := make(map[rune]bool)
	for _, letter := range strings.ToUpper(input) {
		if letter >= 'A' && letter <= 'Z' {
			result[letter] = true
		}
	}
	return len(result) == 26
}
