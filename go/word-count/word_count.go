package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 3

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	result := Frequency{}
	for _, word := range split(phrase) {
		result[word]++
	}
	return result
}

var wordSplitter = regexp.MustCompile("[,:\n ]+")

func split(phrase string) []string {
	result := []string{}
	words := wordSplitter.Split(strings.ToLower(phrase), -1)
	for _, word := range words {
		result = append(result, strings.Trim(word, "'.!@$%^&"))
	}
	return result
}
