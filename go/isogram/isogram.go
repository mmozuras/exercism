package isogram

import "unicode"

const testVersion = 1

func IsIsogram(s string) bool {
	runes := map[rune]bool{}
	for _, r := range s {
		r = unicode.ToLower(r)
		if unicode.IsLetter(r) && runes[r] {
			return false
		}
		runes[r] = true
	}
	return true
}
