package accumulate

const testVersion = 1

func Accumulate(collection []string, converter func(string) string) []string {
	result := []string{}
	for _, element := range collection {
		result = append(result, converter(element))
	}
	return result
}
