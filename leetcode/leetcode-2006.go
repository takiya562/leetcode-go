package leetcode

func countKDifference(nums []int, k int) int {
	counts := map[int]int{}
	ans := 0
	for _, num := range nums {
		ans += counts[num-k] + counts[num+k]
		counts[num]++
	}
	return ans
}

func CountKDifference(nums []int, k int) int {
	return countKDifference(nums, k)
}
