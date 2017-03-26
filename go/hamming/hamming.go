package hamming

import "errors"

const testVersion = 5

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Strands must be of the same length.")
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
