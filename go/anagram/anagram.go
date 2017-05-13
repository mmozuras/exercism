package anagram

import "strings"

const testVersion = 1

func Detect(word string, candidates []string) []string {
	result := []string{}
	word = strings.ToLower(word)

	for _, candidate := range candidates {
		candidate = strings.ToLower(candidate)
		if isAnagram(word, candidate) {
			result = append(result, candidate)
		}
	}
	return result
}

func isAnagram(word string, candidate string) bool {
	if candidate == word || len(word) > len(candidate) {
		return false
	}
	for _, r := range word {
		candidate = strings.Replace(candidate, string(r), "", 1)
	}
	return candidate == ""
}
