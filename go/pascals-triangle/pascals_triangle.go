package pascal

const testVersion = 1

func Triangle(number int) [][]int {
	result := make([][]int, number)
	for row := 0; row < number; row++ {
		result[row] = make([]int, row+1)
		left, right := 1, row-1
		result[row][left-1], result[row][right+1] = 1, 1
		for i := left; i <= right; i++ {
			result[row][i] = result[row-1][i-1] + result[row-1][i]
		}
	}
	return result
}
