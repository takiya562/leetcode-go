package leetcode

func findCenter(edges [][]int) int {
	n := len(edges)
	c := make([]int, n+2)
	for _, edge := range edges {
		i, j := edge[0], edge[1]
		c[i]++
		c[j]++
		if c[i] == n {
			return i
		}
		if c[j] == n {
			return j
		}
	}
	return -1
}

func FindCenter(edges [][]int) int {
	return findCenter(edges)
}
