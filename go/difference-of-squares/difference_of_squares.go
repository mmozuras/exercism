package diffsquares

const testVersion = 1

func SquareOfSums(n int) int {
	return n * n * (n + 1) * (n + 1) / 4
}

func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
