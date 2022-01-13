package leetcode

func DominantIndex(nums []int) int {
	return dominantIndex(nums)
}

func dominantIndex(nums []int) int {
	m1, m2, idx := -1, -1, 0
	for i, num := range nums {
		if num > m1 {
			m1, m2, idx = num, m1, i
		} else if num > m2 {
			m2 = num
		}
	}
	if m2*2 <= m1 {
		return idx
	}
	return -1
}
