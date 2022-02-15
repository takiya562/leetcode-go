package leetcode

func luckyNumbers(matrix [][]int) []int {
	n := len(matrix)
	m := len(matrix[0])
	r := make([]int, n)
	for i := 0; i < n; i++ {
		min := 0
		for j := 0; j < m; j++ {
			if matrix[i][j] < matrix[i][min] {
				min = j
			}
		}
		r[i] = min
	}
	ans := []int{}
	for i := 0; i < m; i++ {
		max := 0
		for j := 0; j < n; j++ {
			if matrix[j][i] > matrix[max][i] {
				max = j
			}
		}
		if i == r[max] {
			ans = append(ans, matrix[max][i])
		}
	}
	return ans
}

func LuckyNumbers(matrix [][]int) []int {
	return luckyNumbers(matrix)
}
