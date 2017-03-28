package raindrops

import "strconv"

const testVersion = 2

func Convert(i int) string {
	result := drop(i, 3, "Pling") + drop(i, 5, "Plang") + drop(i, 7, "Plong")

	if result == "" {
		return strconv.Itoa(i)
	}
	return result
}

func drop(i, factor int, sound string) string {
	if i%factor == 0 {
		return sound
	}
	return ""
}
