package acronym

import "unicode"

const testVersion = 2

func Abbreviate(s string) string {
	var acronym []rune
	var previous rune

	for _, r := range s {
		if (unicode.IsUpper(r) && !unicode.IsUpper(previous)) ||
			(unicode.IsLetter(r) && !unicode.IsLetter(previous)) {
			acronym = append(acronym, unicode.ToUpper(r))
		}
		previous = r
	}
	return string(acronym)
}
