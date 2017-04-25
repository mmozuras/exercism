package sieve

const testVersion = 1

func Sieve(limit int) []int {
	result := []int{}

	sieve := make([]bool, limit)
	for i := 2; i < limit; i++ {
		if !sieve[i] {
			result = append(result, i)
			for j := 2 * i; j < limit; j += i {
				sieve[j] = true
			}
		}
	}
	return result
}
