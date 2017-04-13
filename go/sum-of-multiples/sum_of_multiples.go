package summultiples

const testVersion = 1

func SumMultiples(max int, divisors ...int) int {
	result := 0
	for i := 1; i < max; i++ {
		for _, divisor := range divisors {
			if i%divisor == 0 {
				result += i
				break
			}
		}
	}
	return result
}
