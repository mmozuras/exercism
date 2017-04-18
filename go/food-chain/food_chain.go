package foodchain

const testVersion = 3

var parts = []struct {
	name    string
	comment string
}{
	{"", ""},
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", "It wriggled and jiggled and tickled inside her.\n"},
	{"bird", "How absurd to swallow a bird!\n"},
	{"cat", "Imagine that, to swallow a cat!\n"},
	{"dog", "What a hog, to swallow a dog!\n"},
	{"goat", "Just opened her throat and swallowed a goat!\n"},
	{"cow", "I don't know how she swallowed a cow!\n"},
	{"horse", "She's dead, of course!"},
}

func Song() string {
	return Verses(1, 8)
}

func Verses(start int, end int) string {
	verses := Verse(start)
	for start < end {
		start++
		verses += "\n\n" + Verse(start)
	}
	return verses
}

func Verse(i int) string {
	result := "I know an old lady who swallowed a " + parts[i].name + ".\n" + parts[i].comment

	if i == 1 || i == 8 {
		return result
	}

	for ; i > 1 && i < 8; i-- {
		result += "She swallowed the " + parts[i].name
		result += " to catch the " + parts[i-1].name
		if i == 3 {
			result += " that wriggled and jiggled and tickled inside her"
		}
		result += ".\n"
	}
	return result + parts[1].comment
}
