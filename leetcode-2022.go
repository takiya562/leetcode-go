package main

func construct2DArray(original []int, m int, n int) (ans [][]int) {
	if len(original) != m*n {
		return ans
	}
	ans = make([][]int, m)
	for l := 0; l < m; l++ {
		ans[l] = make([]int, n)
		for i := 0; i < n; i++ {
			ans[l][i] = original[l*n+i]
		}
	}
	return
}

func main() {

}
