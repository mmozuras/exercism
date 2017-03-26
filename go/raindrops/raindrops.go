package raindrops

import "strconv"

const testVersion = 2

func Convert(number int) string {
	result := drop(number, 3, "Pling") +
		drop(number, 5, "Plang") +
		drop(number, 7, "Plong")

	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}

func drop(number, factor int, sound string) string {
	if number%factor == 0 {
		return sound
	}
	return ""
}
