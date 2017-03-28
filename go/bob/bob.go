package bob

import "strings"

const testVersion = 2

func Hey(s string) string {
	s = strings.TrimSpace(s)
	if isSilence(s) {
		return "Fine. Be that way!"
	} else if isYelling(s) {
		return "Whoa, chill out!"
	} else if isQuestion(s) {
		return "Sure."
	}
	return "Whatever."
}

func isSilence(s string) bool {
	return len(s) == 0
}

func isYelling(s string) bool {
	return s == strings.ToUpper(s) && s != strings.ToLower(s)
}

func isQuestion(s string) bool {
	return s[len(s)-1:] == "?"
}
