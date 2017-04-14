package pythagorean

const testVersion = 1

type Triplet [3]int

func Range(min, max int) (result []Triplet) {
	for a := min; a <= max; a++ {
		b := a + 1
		c := b + 1
		for c <= max {
			for c*c < a*a+b*b {
				c += 1
			}
			if c*c == a*a+b*b && c <= max {
				result = append(result, Triplet{a, b, c})
			}
			b += 1
		}
	}
	return result
}

func Sum(p int) (result []Triplet) {
	for _, triplet := range Range(1, p) {
		if triplet[0]+triplet[1]+triplet[2] == p {
			result = append(result, triplet)
		}
	}

	return result
}
