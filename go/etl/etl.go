package etl

import "strings"

const testVersion = 1

func Transform(input map[int][]string) map[string]int {
	output := map[string]int{}
	for key, values := range input {
		for _, value := range values {
			output[strings.ToLower(value)] = key
		}
	}
	return output
}
