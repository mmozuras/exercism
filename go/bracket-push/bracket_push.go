package brackets

const testVersion = 4

var brackets = map[string]string{
	"{": "}",
	"[": "]",
	"(": ")",
}

func Bracket(input string) (bool, error) {
	result := "_"
	for _, c := range input {
		value := brackets[result[len(result)-1:]]
		result += string(c)
		if string(c) == value {
			result = result[:len(result)-2]
		}
	}
	return len(result) == 1, nil
}
